/*
Copyright (C) 2018 Intel Corporation.

SPDX-License-Identifier: Apache-2.0
*/

package oimcontroller

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/intel/oim/pkg/log"
	"github.com/intel/oim/pkg/oim-common"
	"github.com/intel/oim/pkg/spdk"
	"github.com/intel/oim/pkg/spec/oim/v0"
)

// Controller implements oim.Controller.
type Controller struct {
	registryAddress string
	registryDelay   time.Duration
	controllerID    string
	controllerAddr  string
	spdkPath        string
	SPDK            *spdk.Client
	vhostSCSI       string
	vhostDev        *oim.PCIAddress

	wg   sync.WaitGroup
	stop chan<- interface{}
}

func (c *Controller) MapVolume(ctx context.Context, in *oim.MapVolumeRequest) (*oim.MapVolumeReply, error) {
	volumeID := in.GetVolumeId()
	if volumeID == "" {
		return nil, errors.New("empty volume ID")
	}
	if c.SPDK == nil {
		return nil, errors.New("not connected to SPDK")
	}
	if c.vhostSCSI == "" {
		return nil, errors.New("no VHost SCSI controller configured")
	}
	if c.vhostDev == nil {
		return nil, errors.New("no PCI BDF configured")
	}

	// Reuse or create BDev.
	if _, err := spdk.GetBDevs(ctx, c.SPDK, spdk.GetBDevsArgs{Name: volumeID}); err != nil {
		// TODO: check error more carefully instead of assuming that it merely
		// wasn't found.
		switch x := in.Params.(type) {
		case *oim.MapVolumeRequest_Malloc:
			return nil, fmt.Errorf("no existing MallocBDev with name %s found", volumeID)
		case *oim.MapVolumeRequest_Ceph:
			if err := c.mapCeph(ctx, volumeID, x.Ceph); err != nil {
				return nil, err
			}
		case nil:
			return nil, errors.New("missing volume parameters")
		default:
			return nil, errors.New(fmt.Sprintf("unsupported params type %T", x))
		}
	} else {
		// BDev with the intended name already exists. Assume that it is the right one.
		log.FromContext(ctx).Infof("reusing existing BDev %s", volumeID)
	}

	var err error

	// If this BDev is active as LUN, do nothing because a previous MapVolume
	// call must have succeeded (idempotency!).
	controllers, err := spdk.GetVHostControllers(ctx, c.SPDK)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetVHostControllers: %s", err))
	}
	for _, controller := range controllers {
		for key, value := range controller.BackendSpecific {
			switch key {
			case "scsi":
				if scsi, ok := value.(spdk.SCSIControllerSpecific); ok {
					for _, target := range scsi {
						for _, lun := range target.LUNs {
							if lun.BDevName == volumeID {
								// BDev already active.
								return &oim.MapVolumeReply{
									PciAddress: c.vhostDev,
									ScsiDisk: &oim.SCSIDisk{
										Target: target.SCSIDevNum,
										Lun:    0,
									},
								}, nil
							}
						}
					}
				}
			}
		}
	}

	// Create a new SCSI target with a LUN connected to this BDev. We iterate over all available
	// targets and attempt to use them.
	// TODO: we don't know the SPDK limit for targets. 8 is just the default.
	// TODO: let vhost pick an unused one (https://github.com/spdk/spdk/issues/328)
	for target := uint32(0); target < 8; target++ {
		args := spdk.AddVHostSCSILUNArgs{
			Controller:    c.vhostSCSI,
			SCSITargetNum: target,
			BDevName:      volumeID,
		}
		err = spdk.AddVHostSCSILUN(ctx, c.SPDK, args)
		if err == nil {
			// Success!
			return &oim.MapVolumeReply{
				PciAddress: c.vhostDev,
				ScsiDisk: &oim.SCSIDisk{
					Target: target,
					Lun:    0,
				},
			}, nil
		}
	}

	// TODO: document that the BDev is not going to get deleted.
	// To remove it, UnmapVolume must be called.

	// Return the last SPDK error.
	errorResult := errors.New(fmt.Sprintf("AddVHostSCSILUN failed for all LUNs, last error: %s", err))
	return nil, errorResult
}

func (c *Controller) UnmapVolume(ctx context.Context, in *oim.UnmapVolumeRequest) (*oim.UnmapVolumeReply, error) {
	volumeID := in.GetVolumeId()
	if volumeID == "" {
		return nil, errors.New("empty volume ID")
	}
	if c.SPDK == nil {
		return nil, errors.New("not connected to SPDK")
	}

	controllers, err := spdk.GetVHostControllers(ctx, c.SPDK)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetVHostControllers: %s", err))
	}
	// For the sake of completeness we keep iterating even after having found
	// something.
	for _, controller := range controllers {
		for key, value := range controller.BackendSpecific {
			switch key {
			case "scsi":
				if scsi, ok := value.(spdk.SCSIControllerSpecific); ok {
					for _, target := range scsi {
						for _, lun := range target.LUNs {
							if lun.BDevName == volumeID {
								// Found the right SCSI target.
								removeArgs := spdk.RemoveVHostSCSITargetArgs{
									Controller:    controller.Controller,
									SCSITargetNum: target.SCSIDevNum,
								}
								if err := spdk.RemoveVHostSCSITarget(ctx, c.SPDK, removeArgs); err != nil {
									return nil, errors.New(fmt.Sprintf("RemoveVHostSCSITarget: %s", err))
								}
							}
						}
					}
				}
			}
		}
	}

	// Don't fail when the BDev is not found (idempotency).
	// Check whether this is really a BDev created by MapVolume (i.e. everything except MallocBDevs).
	// TODO: detect "not found" errors (https://github.com/spdk/spdk/issues/319)
	if bdev, err := spdk.GetBDevs(ctx, c.SPDK, spdk.GetBDevsArgs{Name: volumeID}); err == nil && len(bdev) > 0 && bdev[0].ProductName != "Malloc disk" {
		if err := spdk.DeleteBDev(ctx, c.SPDK, spdk.DeleteBDevArgs{Name: volumeID}); err != nil {
			// TODO: detect "not found" error (https://github.com/spdk/spdk/issues/319)
		}
	}

	return &oim.UnmapVolumeReply{}, nil
}

func (c *Controller) ProvisionMallocBDev(ctx context.Context, in *oim.ProvisionMallocBDevRequest) (*oim.ProvisionMallocBDevReply, error) {
	bdevName := in.GetBdevName()
	if bdevName == "" {
		return nil, errors.New("empty BDev name")
	}
	if c.SPDK == nil {
		return nil, errors.New("not connected to SPDK")
	}
	size := in.Size_
	if size != 0 {
		bdevs, err := spdk.GetBDevs(ctx, c.SPDK, spdk.GetBDevsArgs{Name: bdevName})
		if err != nil || len(bdevs) != 1 {
			args := spdk.ConstructMallocBDevArgs{
				ConstructBDevArgs: spdk.ConstructBDevArgs{
					NumBlocks: size / 512,
					BlockSize: 512,
					Name:      bdevName,
				},
			}
			// TODO: detect already existing BDev of the same name (https://github.com/spdk/spdk/issues/319)
			if _, err := spdk.ConstructMallocBDev(ctx, c.SPDK, args); err != nil {
				return nil, errors.New(fmt.Sprintf("ConstructMallocBDev failed: %s", err))
			}
		} else {
			// Check that the BDev has the right size.
			actualSize := bdevs[0].NumBlocks * bdevs[0].BlockSize
			if actualSize != size {
				return nil, status.Errorf(codes.AlreadyExists, "Existing BDev %s has wrong size %d", bdevName, actualSize)
			}
		}
	} else {
		if err := spdk.DeleteBDev(ctx, c.SPDK, spdk.DeleteBDevArgs{Name: bdevName}); err != nil {
			// TODO: detect error (https://github.com/spdk/spdk/issues/319)
		}
	}
	return &oim.ProvisionMallocBDevReply{}, nil
}

func (c *Controller) CheckMallocBDev(ctx context.Context, in *oim.CheckMallocBDevRequest) (*oim.CheckMallocBDevReply, error) {
	bdevName := in.GetBdevName()
	if bdevName == "" {
		return nil, errors.New("empty BDev name")
	}
	if c.SPDK == nil {
		return nil, errors.New("not connected to SPDK")
	}

	bdevs, err := spdk.GetBDevs(ctx, c.SPDK, spdk.GetBDevsArgs{Name: bdevName})
	if err == nil && len(bdevs) == 1 {
		return &oim.CheckMallocBDevReply{}, nil
	} else {
		// TODO: detect "not found" error (https://github.com/spdk/spdk/issues/319)
		return nil, status.Error(codes.NotFound, "")
	}
}

func (c *Controller) mapCeph(ctx context.Context, volumeID string, cephParams *oim.CephParams) error {
	if c.SPDK == nil {
		return errors.New("not connected to SPDK")
	}
	request := spdk.ConstructRBDBDevArgs{
		BlockSize: 512,
		Name:      volumeID,
		UserID:    cephParams.UserId,
		PoolName:  cephParams.Pool,
		RBDName:   cephParams.Image,
		Config: map[string]string{
			"mon_host": cephParams.Monitors,
			"key":      cephParams.Secret,
		},
	}
	_, err := spdk.ConstructRBDBDev(ctx, c.SPDK, request)
	return errors.Wrapf(err, "ConstructRBDBDev %q for RBD pool %q and image %q, monitors %q", volumeID, cephParams.Pool, cephParams.Image, cephParams.Monitors)
}

type Option func(c *Controller) error

func WithRegistry(address string) Option {
	return func(c *Controller) error {
		c.registryAddress = address
		return nil
	}
}

func WithRegistryDelay(delay time.Duration) Option {
	return func(c *Controller) error {
		c.registryDelay = delay
		return nil
	}
}

// WithControllerAddress sets the *external* address for the
// controller, i.e. what the OIM registry needs to use for gRPC.Dial
// to contact the controller.
func WithControllerAddress(address string) Option {
	return func(c *Controller) error {
		c.controllerAddr = address
		return nil
	}
}

func WithControllerID(controllerID string) Option {
	return func(c *Controller) error {
		c.controllerID = controllerID
		return nil
	}
}

func WithSPDK(path string) Option {
	return func(c *Controller) error {
		c.spdkPath = path
		return nil
	}
}

func WithVHostController(vhost string) Option {
	return func(c *Controller) error {
		c.vhostSCSI = vhost
		return nil
	}
}

func WithVHostDev(dev string) Option {
	return func(c *Controller) error {
		d, err := oimcommon.ParseBDFString(dev)
		if err != nil {
			return err
		}
		c.vhostDev = d
		return nil
	}
}

func New(options ...Option) (*Controller, error) {
	c := Controller{
		controllerID:  "unset-controller-id",
		registryDelay: time.Minute,
	}
	for _, op := range options {
		err := op(&c)
		if err != nil {
			return nil, err
		}
	}

	if c.spdkPath != "" {
		client, err := spdk.New(c.spdkPath)
		if err != nil {
			return nil, err
		}
		c.SPDK = client
	}

	if c.registryAddress != "" && (c.controllerID == "" || c.controllerAddr == "") {
		return nil, errors.New("Need both controller ID and external controller address for registering  with the OIM registry.")
	}

	return &c, nil
}

// Starts the interaction with the OIM Registry, if one was configured.
func (c *Controller) Start() error {
	if c.registryAddress == "" {
		return nil
	}

	stop := make(chan interface{})
	c.stop = stop
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Register for the first time immediately.
		again := time.After(0 * time.Second)
		done := make(chan bool)
		for {
			select {
			case <-stop:
				return
			case <-done:
				// TODO (?): exponential backoff when registry is down
				again = time.After(c.registryDelay)
			case <-again:
				// Run at most one call at a time by re-arming
				// the time only after we are done.
				go func() {
					c.register(ctx)
					done <- true
				}()
			}
		}
	}()

	return nil
}

func (c *Controller) register(ctx context.Context) {
	// Dial anew, because a) when the registry is down
	// and our address uses Unix domain sockets, dialing
	// will fail permanently and b) we don't want to keep
	// a permanent connection from each controller to
	// the registry.
	log.L().Infof("Registering OIM controller %s at address %s with OIM registry %s", c.controllerID, c.controllerAddr, c.registryAddress)
	// TODO: secure connection
	opts := oimcommon.ChooseDialOpts(c.registryAddress, grpc.WithInsecure())
	conn, err := grpc.DialContext(ctx, c.registryAddress, opts...)
	if err != nil {
		log.L().Infow("connecting to OIM registry", "error", err)
	}
	defer conn.Close()
	registry := oim.NewRegistryClient(conn)
	registry.SetValue(ctx, &oim.SetValueRequest{
		Value: &oim.Value{
			Path:  c.controllerID + "/" + oimcommon.RegistryAddress,
			Value: c.controllerAddr,
		},
	})
}

// Stops the interaction with the OIM Registry, if one was configured.
func (c *Controller) Stop() {
	if c.stop != nil {
		close(c.stop)
		c.wg.Wait()
	}
}

func Server(endpoint string, c oim.ControllerServer) (*oimcommon.NonBlockingGRPCServer, func(*grpc.Server)) {
	service := func(s *grpc.Server) {
		oim.RegisterControllerServer(s, c)
	}
	server := &oimcommon.NonBlockingGRPCServer{
		Endpoint: endpoint,
	}
	return server, service
}

/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/pkg/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/version"
	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/framework/ginkgowrapper"

	"github.com/intel/oim/pkg/log"
	"github.com/intel/oim/test/pkg/qemu"
	"github.com/intel/oim/test/pkg/spdk"
)

var initialized = false

// setupProviderConfig validates and sets up cloudConfig based on framework.TestContext.Provider.
func setupProviderConfig(data *[]byte) error {
	switch framework.TestContext.Provider {
	case "skeleton": // The default.
		fallthrough
	case "":
		if *data == nil {
			if err := spdk.Init(spdk.WithVHostSCSI()); err != nil {
				return err
			}
			if err := qemu.Init(qemu.WithKubernetes()); err != nil {
				return err
			}
			if qemu.VM == nil {
				return errors.New("a QEMU image is required for this test")
			}
			// Tell child nodes about our SPDK path.
			*data = []byte(spdk.SPDKPath)
			initialized = true
		} else {
			if initialized {
				// This gets called twice on the master node, once with data and once without.
				// We don't need to do anything the second time.
				return nil
			}

			if err := qemu.SimpleInit(); err != nil {
				return err
			}
			log.L().Info("using SPDK with path %s", string(*data))
			if err := spdk.Init(spdk.WithSPDKSocket(string(*data))); err != nil {
				return err
			}
		}
		config, err := qemu.KubeConfig()
		if err != nil {
			return err
		}

		// This is no longer enough to make the following code use the config.
		// KUBECONFIG must already be set correctly before invoking the binary.
		framework.TestContext.KubeConfig = config
		abs := func(path string) string {
			p, err := filepath.Abs(path)
			if err != nil {
				return path
			}
			return p
		}
		if abs(os.Getenv("KUBECONFIG")) != abs(config) {
			return errors.Errorf("KUBECONFIG must be set to %s", abs(config))
		}
	}

	return nil
}

// There are certain operations we only want to run once per overall test invocation
// (such as deleting old namespaces, or verifying that all system pods are running.
// Because of the way Ginkgo runs tests in parallel, we must use SynchronizedBeforeSuite
// to ensure that these operations only run on the first parallel Ginkgo node.
//
// This function takes two parameters: one function which runs on only the first Ginkgo node,
// returning an opaque byte array, and then a second function which runs on all Ginkgo nodes,
// accepting the byte array.
var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
	// Run only on Ginkgo node 1
	var data []byte

	log.L().Info("checking config")

	if err := setupProviderConfig(&data); err != nil {
		framework.Failf("Failed to setup provider config: %v", err)
	}

	c, err := framework.LoadClientset()
	if err != nil {
		log.L().Fatalf("Error loading client: ", err)
	}

	// Delete any namespaces except those created by the system. This ensures no
	// lingering resources are left over from a previous test run.
	if framework.TestContext.CleanStart {
		deleted, err := framework.DeleteNamespaces(c, nil, /* deleteFilter */
			[]string{
				metav1.NamespaceSystem,
				metav1.NamespaceDefault,
				metav1.NamespacePublic,
			})
		if err != nil {
			framework.Failf("Error deleting orphaned namespaces: %v", err)
		}
		log.L().Infof("Waiting for deletion of the following namespaces: %v", deleted)
		if err := framework.WaitForNamespacesDeleted(c, deleted, framework.NamespaceCleanupTimeout); err != nil {
			framework.Failf("Failed to delete orphaned namespaces %v: %v", deleted, err)
		}
	}

	// In large clusters we may get to this point but still have a bunch
	// of nodes without Routes created. Since this would make a node
	// unschedulable, we need to wait until all of them are schedulable.
	framework.ExpectNoError(framework.WaitForAllNodesSchedulable(c, framework.TestContext.NodeSchedulableTimeout))

	// Ensure all pods are running and ready before starting tests (otherwise,
	// cluster infrastructure pods that are being pulled or started can block
	// test pods from running, and tests that ensure all pods are running and
	// ready will fail).
	podStartupTimeout := framework.TestContext.SystemPodsStartupTimeout
	// TODO: In large clusters, we often observe a non-starting pods due to
	// #41007. To avoid those pods preventing the whole test runs (and just
	// wasting the whole run), we allow for some not-ready pods (with the
	// number equal to the number of allowed not-ready nodes).
	if err := framework.WaitForPodsRunningReady(c, metav1.NamespaceSystem, int32(framework.TestContext.MinStartupPods), int32(framework.TestContext.AllowedNotReadyNodes), podStartupTimeout, map[string]string{}); err != nil {
		framework.DumpAllNamespaceInfo(c, metav1.NamespaceSystem)
		framework.LogFailedContainers(c, metav1.NamespaceSystem, framework.Logf)
		framework.Failf("Error waiting for all pods to be running and ready: %v", err)
	}

	if err := framework.WaitForDaemonSets(c, metav1.NamespaceSystem, int32(framework.TestContext.AllowedNotReadyNodes), framework.TestContext.SystemDaemonsetStartupTimeout); err != nil {
		framework.Logf("WARNING: Waiting for all daemonsets to be ready failed: %v", err)
	}

	// Log the version of the server and this client.
	framework.Logf("e2e test version: %s", version.Get().GitVersion)

	dc := c.DiscoveryClient

	serverVersion, serverErr := dc.ServerVersion()
	if serverErr != nil {
		framework.Logf("Unexpected server error retrieving version: %v", serverErr)
	}
	if serverVersion != nil {
		framework.Logf("kube-apiserver version: %s", serverVersion.GitVersion)
	}

	// Reference common test to make the import valid.
	// commontest.CurrentSuite = commontest.E2E

	return data

}, func(data []byte) {
	// Run on all Ginkgo nodes

	if err := setupProviderConfig(&data); err != nil {
		framework.Failf("Failed to setup provider config: %v", err)
	}
})

// Similar to SynchornizedBeforeSuite, we want to run some operations only once (such as collecting cluster logs).
// Here, the order of functions is reversed; first, the function which runs everywhere,
// and then the function that only runs on the first Ginkgo node.
var _ = ginkgo.SynchronizedAfterSuite(func() {
	// Run on all Ginkgo nodes
	framework.Logf("Running AfterSuite actions on all node")
	framework.RunCleanupActions()
}, func() {
	// Run only Ginkgo on node 1
	framework.Logf("Running AfterSuite actions on node 1")
	if err := qemu.Finalize(); err != nil {
		framework.Logf("QEMU shutdown: %s", err)
	}
	if err := spdk.Finalize(); err != nil {
		framework.Logf("SPDK shutdown: %s", err)
	}
})

// RunE2ETests checks configuration parameters (specified through flags) and then runs
// E2E tests using the Ginkgo runner.
// This function is called on each Ginkgo node in parallel mode.
func RunE2ETests(t *testing.T) {
	// TODO: with "ginkgo ./test/e2e" we shouldn't get verbose output, but somehow we do.
	gomega.RegisterFailHandler(ginkgowrapper.Fail)
	ginkgo.RunSpecs(t, "OIM E2E suite")
}

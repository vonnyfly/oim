//
// Copyright 2018, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudstack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type AddClusterParams struct {
	p map[string]interface{}
}

func (p *AddClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["guestvswitchname"]; found {
		u.Set("guestvswitchname", v.(string))
	}
	if v, found := p.p["guestvswitchtype"]; found {
		u.Set("guestvswitchtype", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["ovm3cluster"]; found {
		u.Set("ovm3cluster", v.(string))
	}
	if v, found := p.p["ovm3pool"]; found {
		u.Set("ovm3pool", v.(string))
	}
	if v, found := p.p["ovm3vip"]; found {
		u.Set("ovm3vip", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["publicvswitchname"]; found {
		u.Set("publicvswitchname", v.(string))
	}
	if v, found := p.p["publicvswitchtype"]; found {
		u.Set("publicvswitchtype", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["vsmipaddress"]; found {
		u.Set("vsmipaddress", v.(string))
	}
	if v, found := p.p["vsmpassword"]; found {
		u.Set("vsmpassword", v.(string))
	}
	if v, found := p.p["vsmusername"]; found {
		u.Set("vsmusername", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddClusterParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *AddClusterParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
	return
}

func (p *AddClusterParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
	return
}

func (p *AddClusterParams) SetGuestvswitchname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestvswitchname"] = v
	return
}

func (p *AddClusterParams) SetGuestvswitchtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestvswitchtype"] = v
	return
}

func (p *AddClusterParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *AddClusterParams) SetOvm3cluster(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ovm3cluster"] = v
	return
}

func (p *AddClusterParams) SetOvm3pool(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ovm3pool"] = v
	return
}

func (p *AddClusterParams) SetOvm3vip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ovm3vip"] = v
	return
}

func (p *AddClusterParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddClusterParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *AddClusterParams) SetPublicvswitchname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicvswitchname"] = v
	return
}

func (p *AddClusterParams) SetPublicvswitchtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicvswitchtype"] = v
	return
}

func (p *AddClusterParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddClusterParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

func (p *AddClusterParams) SetVsmipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vsmipaddress"] = v
	return
}

func (p *AddClusterParams) SetVsmpassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vsmpassword"] = v
	return
}

func (p *AddClusterParams) SetVsmusername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vsmusername"] = v
	return
}

func (p *AddClusterParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new AddClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewAddClusterParams(clustername string, clustertype string, hypervisor string, podid string, zoneid string) *AddClusterParams {
	p := &AddClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clustername"] = clustername
	p.p["clustertype"] = clustertype
	p.p["hypervisor"] = hypervisor
	p.p["podid"] = podid
	p.p["zoneid"] = zoneid
	return p
}

// Adds a new cluster
func (s *ClusterService) AddCluster(p *AddClusterParams) (*AddClusterResponse, error) {
	resp, err := s.cs.newRequest("addCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddClusterResponse struct {
	Allocationstate string `json:"allocationstate"`
	Capacity        []struct {
		Capacitytotal int64  `json:"capacitytotal"`
		Capacityused  int64  `json:"capacityused"`
		Clusterid     string `json:"clusterid"`
		Clustername   string `json:"clustername"`
		Percentused   string `json:"percentused"`
		Podid         string `json:"podid"`
		Podname       string `json:"podname"`
		Type          int    `json:"type"`
		Zoneid        string `json:"zoneid"`
		Zonename      string `json:"zonename"`
	} `json:"capacity"`
	Clustertype           string            `json:"clustertype"`
	Cpuovercommitratio    string            `json:"cpuovercommitratio"`
	Hypervisortype        string            `json:"hypervisortype"`
	Id                    string            `json:"id"`
	Managedstate          string            `json:"managedstate"`
	Memoryovercommitratio string            `json:"memoryovercommitratio"`
	Name                  string            `json:"name"`
	Ovm3vip               string            `json:"ovm3vip"`
	Podid                 string            `json:"podid"`
	Podname               string            `json:"podname"`
	Resourcedetails       map[string]string `json:"resourcedetails"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}

type DedicateClusterParams struct {
	p map[string]interface{}
}

func (p *DedicateClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	return u
}

func (p *DedicateClusterParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *DedicateClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

func (p *DedicateClusterParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

// You should always use this function to get a new DedicateClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDedicateClusterParams(clusterid string, domainid string) *DedicateClusterParams {
	p := &DedicateClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	p.p["domainid"] = domainid
	return p
}

// Dedicate an existing cluster
func (s *ClusterService) DedicateCluster(p *DedicateClusterParams) (*DedicateClusterResponse, error) {
	resp, err := s.cs.newRequest("dedicateCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicateClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DedicateClusterResponse struct {
	JobID           string `json:"jobid"`
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Clusterid       string `json:"clusterid"`
	Clustername     string `json:"clustername"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
}

type DeleteClusterParams struct {
	p map[string]interface{}
}

func (p *DeleteClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDeleteClusterParams(id string) *DeleteClusterParams {
	p := &DeleteClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a cluster.
func (s *ClusterService) DeleteCluster(p *DeleteClusterParams) (*DeleteClusterResponse, error) {
	resp, err := s.cs.newRequest("deleteCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteClusterResponse struct {
	Displaytext string `json:"displaytext"`
	Success     string `json:"success"`
}

type DisableOutOfBandManagementForClusterParams struct {
	p map[string]interface{}
}

func (p *DisableOutOfBandManagementForClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (p *DisableOutOfBandManagementForClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

// You should always use this function to get a new DisableOutOfBandManagementForClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDisableOutOfBandManagementForClusterParams(clusterid string) *DisableOutOfBandManagementForClusterParams {
	p := &DisableOutOfBandManagementForClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	return p
}

// Disables out-of-band management for a cluster
func (s *ClusterService) DisableOutOfBandManagementForCluster(p *DisableOutOfBandManagementForClusterParams) (*DisableOutOfBandManagementForClusterResponse, error) {
	resp, err := s.cs.newRequest("disableOutOfBandManagementForCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableOutOfBandManagementForClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DisableOutOfBandManagementForClusterResponse struct {
	JobID       string `json:"jobid"`
	Action      string `json:"action"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Driver      string `json:"driver"`
	Enabled     bool   `json:"enabled"`
	Hostid      string `json:"hostid"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	Powerstate  string `json:"powerstate"`
	Status      bool   `json:"status"`
	Username    string `json:"username"`
}

type EnableOutOfBandManagementForClusterParams struct {
	p map[string]interface{}
}

func (p *EnableOutOfBandManagementForClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (p *EnableOutOfBandManagementForClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

// You should always use this function to get a new EnableOutOfBandManagementForClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewEnableOutOfBandManagementForClusterParams(clusterid string) *EnableOutOfBandManagementForClusterParams {
	p := &EnableOutOfBandManagementForClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	return p
}

// Enables out-of-band management for a cluster
func (s *ClusterService) EnableOutOfBandManagementForCluster(p *EnableOutOfBandManagementForClusterParams) (*EnableOutOfBandManagementForClusterResponse, error) {
	resp, err := s.cs.newRequest("enableOutOfBandManagementForCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableOutOfBandManagementForClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type EnableOutOfBandManagementForClusterResponse struct {
	JobID       string `json:"jobid"`
	Action      string `json:"action"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Driver      string `json:"driver"`
	Enabled     bool   `json:"enabled"`
	Hostid      string `json:"hostid"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	Powerstate  string `json:"powerstate"`
	Status      bool   `json:"status"`
	Username    string `json:"username"`
}

type ListClustersParams struct {
	p map[string]interface{}
}

func (p *ListClustersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["managedstate"]; found {
		u.Set("managedstate", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListClustersParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *ListClustersParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
	return
}

func (p *ListClustersParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *ListClustersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListClustersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListClustersParams) SetManagedstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managedstate"] = v
	return
}

func (p *ListClustersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListClustersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListClustersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListClustersParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *ListClustersParams) SetShowcapacities(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showcapacities"] = v
	return
}

func (p *ListClustersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListClustersParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListClustersParams() *ListClustersParams {
	p := &ListClustersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListClustersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListClusters(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Clusters[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Clusters {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterByName(name string, opts ...OptionFunc) (*Cluster, int, error) {
	id, count, err := s.GetClusterID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetClusterByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterByID(id string, opts ...OptionFunc) (*Cluster, int, error) {
	p := &ListClustersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListClusters(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Clusters[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Cluster UUID: %s!", id)
}

// Lists clusters.
func (s *ClusterService) ListClusters(p *ListClustersParams) (*ListClustersResponse, error) {
	resp, err := s.cs.newRequest("listClusters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListClustersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListClustersResponse struct {
	Count    int        `json:"count"`
	Clusters []*Cluster `json:"cluster"`
}

type Cluster struct {
	Allocationstate string `json:"allocationstate"`
	Capacity        []struct {
		Capacitytotal int64  `json:"capacitytotal"`
		Capacityused  int64  `json:"capacityused"`
		Clusterid     string `json:"clusterid"`
		Clustername   string `json:"clustername"`
		Percentused   string `json:"percentused"`
		Podid         string `json:"podid"`
		Podname       string `json:"podname"`
		Type          int    `json:"type"`
		Zoneid        string `json:"zoneid"`
		Zonename      string `json:"zonename"`
	} `json:"capacity"`
	Clustertype           string            `json:"clustertype"`
	Cpuovercommitratio    string            `json:"cpuovercommitratio"`
	Hypervisortype        string            `json:"hypervisortype"`
	Id                    string            `json:"id"`
	Managedstate          string            `json:"managedstate"`
	Memoryovercommitratio string            `json:"memoryovercommitratio"`
	Name                  string            `json:"name"`
	Ovm3vip               string            `json:"ovm3vip"`
	Podid                 string            `json:"podid"`
	Podname               string            `json:"podname"`
	Resourcedetails       map[string]string `json:"resourcedetails"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}

type ListDedicatedClustersParams struct {
	p map[string]interface{}
}

func (p *ListDedicatedClustersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["affinitygroupid"]; found {
		u.Set("affinitygroupid", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListDedicatedClustersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListDedicatedClustersParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
	return
}

func (p *ListDedicatedClustersParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

func (p *ListDedicatedClustersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListDedicatedClustersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListDedicatedClustersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListDedicatedClustersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListDedicatedClustersParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListDedicatedClustersParams() *ListDedicatedClustersParams {
	p := &ListDedicatedClustersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists dedicated clusters.
func (s *ClusterService) ListDedicatedClusters(p *ListDedicatedClustersParams) (*ListDedicatedClustersResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedClusters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedClustersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDedicatedClustersResponse struct {
	Count             int                 `json:"count"`
	DedicatedClusters []*DedicatedCluster `json:"dedicatedcluster"`
}

type DedicatedCluster struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Clusterid       string `json:"clusterid"`
	Clustername     string `json:"clustername"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
}

type ReleaseDedicatedClusterParams struct {
	p map[string]interface{}
}

func (p *ReleaseDedicatedClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (p *ReleaseDedicatedClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

// You should always use this function to get a new ReleaseDedicatedClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewReleaseDedicatedClusterParams(clusterid string) *ReleaseDedicatedClusterParams {
	p := &ReleaseDedicatedClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	return p
}

// Release the dedication for cluster
func (s *ClusterService) ReleaseDedicatedCluster(p *ReleaseDedicatedClusterParams) (*ReleaseDedicatedClusterResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type ReleaseDedicatedClusterResponse struct {
	JobID       string `json:"jobid"`
	Displaytext string `json:"displaytext"`
	Success     bool   `json:"success"`
}

type UpdateClusterParams struct {
	p map[string]interface{}
}

func (p *UpdateClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["managedstate"]; found {
		u.Set("managedstate", v.(string))
	}
	return u
}

func (p *UpdateClusterParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *UpdateClusterParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
	return
}

func (p *UpdateClusterParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
	return
}

func (p *UpdateClusterParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *UpdateClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateClusterParams) SetManagedstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managedstate"] = v
	return
}

// You should always use this function to get a new UpdateClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewUpdateClusterParams(id string) *UpdateClusterParams {
	p := &UpdateClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an existing cluster
func (s *ClusterService) UpdateCluster(p *UpdateClusterParams) (*UpdateClusterResponse, error) {
	resp, err := s.cs.newRequest("updateCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateClusterResponse struct {
	Allocationstate string `json:"allocationstate"`
	Capacity        []struct {
		Capacitytotal int64  `json:"capacitytotal"`
		Capacityused  int64  `json:"capacityused"`
		Clusterid     string `json:"clusterid"`
		Clustername   string `json:"clustername"`
		Percentused   string `json:"percentused"`
		Podid         string `json:"podid"`
		Podname       string `json:"podname"`
		Type          int    `json:"type"`
		Zoneid        string `json:"zoneid"`
		Zonename      string `json:"zonename"`
	} `json:"capacity"`
	Clustertype           string            `json:"clustertype"`
	Cpuovercommitratio    string            `json:"cpuovercommitratio"`
	Hypervisortype        string            `json:"hypervisortype"`
	Id                    string            `json:"id"`
	Managedstate          string            `json:"managedstate"`
	Memoryovercommitratio string            `json:"memoryovercommitratio"`
	Name                  string            `json:"name"`
	Ovm3vip               string            `json:"ovm3vip"`
	Podid                 string            `json:"podid"`
	Podname               string            `json:"podname"`
	Resourcedetails       map[string]string `json:"resourcedetails"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}
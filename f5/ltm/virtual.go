// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import (
	"encoding/json"
	"net/http"

	"e-xpert_solutions/f5-rest-client/f5"
)

// VirtualServerConfig holds a list of virtual server configuration.
type VirtualServerConfig struct {
	Items    []VirtualServerConfigItem `json:"items,omitempty"`
	Kind     string                    `json:"kind,omitempty"`
	SelfLink string                    `json:"selfLink,omitempty"`
}

// VirtualServerConfigItem  holds the configuration of a single virtual server.
type VirtualServerConfigItem struct {
	AddressStatus     string `json:"addressStatus,omitempty"`
	AutoLasthop       string `json:"autoLasthop,omitempty"`
	CmpEnabled        string `json:"cmpEnabled,omitempty"`
	ConnectionLimit   int64  `json:"connectionLimit,omitempty"`
	Description       string `json:"description,omitempty"`
	Destination       string `json:"destination,omitempty"`
	Disabled          bool   `json:"disabled,omitempty"`
	Enabled           bool   `json:"enabled,omitempty"`
	FullPath          string `json:"fullPath,omitempty"`
	Generation        int64  `json:"generation,omitempty"`
	GtmScore          int64  `json:"gtmScore,omitempty"`
	IPProtocol        string `json:"ipProtocol,omitempty"`
	Kind              string `json:"kind,omitempty"`
	Mask              string `json:"mask,omitempty"`
	Mirror            string `json:"mirror,omitempty"`
	MobileAppTunnel   string `json:"mobileAppTunnel,omitempty"`
	Name              string `json:"name,omitempty"`
	Nat64             string `json:"nat64,omitempty"`
	Partition         string `json:"partition,omitempty"`
	PoliciesReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"policiesReference,omitempty"`
	Pool              string `json:"pool,omitempty"`
	ProfilesReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"profilesReference,omitempty"`
	RateLimit                string   `json:"rateLimit,omitempty"`
	RateLimitDstMask         int64    `json:"rateLimitDstMask,omitempty"`
	RateLimitMode            string   `json:"rateLimitMode,omitempty"`
	RateLimitSrcMask         int64    `json:"rateLimitSrcMask,omitempty"`
	Rules                    []string `json:"rules,omitempty"`
	SelfLink                 string   `json:"selfLink,omitempty"`
	Source                   string   `json:"source,omitempty"`
	SourceAddressTranslation struct {
		Type string `json:"type,omitempty"`
	} `json:"sourceAddressTranslation,omitempty"`
	SourcePort       string   `json:"sourcePort,omitempty"`
	SynCookieStatus  string   `json:"synCookieStatus,omitempty"`
	TranslateAddress string   `json:"translateAddress,omitempty"`
	TranslatePort    string   `json:"translatePort,omitempty"`
	Vlans            []string `json:"vlans,omitempty"`
	VlansDisabled    bool     `json:"vlansDisabled,omitempty"`
	VlansEnabled     bool     `json:"vlansEnabled,omitempty"`
	VsIndex          int64    `json:"vsIndex,omitempty"`
}

// VirtualEndpoint represents the REST resource for managing virtual server.
const VirtualEndpoint = "/virtual"

// VirtualResponse provide a simple mechanism to read paginated results.
//
// TODO(gilliek): use VirtualResponse object where pagination is needed.
type VirtualResponse struct {
}

// VirtualResource provides API to manage virtual server configurations.
type VirtualResource struct {
	c f5.Client
}

// ListAll lists all the virtual server configurations.
func (vr *VirtualResource) ListAll() (*VirtualServerConfig, error) {
	resp, err := vr.doRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return nil, err
	}
	var vsc VirtualServerConfig
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&vsc); err != nil {
		return nil, err
	}
	return &vsc, nil
}

// Get a single virtual server configurations identified by id.
func (vr *VirtualResource) Get(id string) (*VirtualServerConfigItem, error) {
	resp, err := vr.doRequest("GET", id, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return nil, err
	}
	var vsci VirtualServerConfigItem
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&vsci); err != nil {
		return nil, err
	}
	return &vsci, nil
}

// Create a new virtual server configuration.
func (vr *VirtualResource) Create(item VirtualServerConfigItem) error {
	resp, err := vr.doRequest("POST", "", item)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Edit a virtual server configuration identified by id.
func (vr *VirtualResource) Edit(id string, item VirtualServerConfigItem) error {
	resp, err := vr.doRequest("PUT", id, item)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Delete a single server configuration identified by id.
func (vr *VirtualResource) Delete(id string) error {
	resp, err := vr.doRequest("DELETE", id, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Rules gets the iRules configuration for a virtual server identified by id.
func (vr *VirtualResource) Rules(id string) ([]Rule, error) {
	resp, err := vr.doRequest("GET", id+"/rule", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return nil, err
	}
	var rules []Rule
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&rules); err != nil {
		return nil, err
	}
	return rules, nil
}

// AddRule adds an iRule to the virtual server identified by id.
func (vr *VirtualResource) AddRule(id string, rule Rule) error {
	resp, err := vr.doRequest("POST", id+"/rule", rule)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return err
	}
	return nil
}

// RemoveRule removes a single  iRule from the virtual server identified by id.
func (vr *VirtualResource) RemoveRule(vsID, ruleID string) error {
	resp, err := vr.doRequest("DELETE", vsID+"/rule/"+ruleID, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return err
	}
	return nil
}

// doRequest creates and send HTTP request using the F5 client.
//
// TODO(gilliek): decorate errors
func (vr *VirtualResource) doRequest(method, id string, data interface{}) (*http.Response, error) {
	req, err := vr.c.MakeRequest(method, BasePath+VirtualEndpoint+"/"+id, data)
	if err != nil {
		return nil, err
	}
	resp, err := vr.c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// readError reads request error object from a HTTP response.
//
// TODO(gilliek): move this function into F5 package.
func (vr *VirtualResource) readError(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		errResp, err := f5.NewRequestError(resp.Body)
		if err != nil {
			return err
		}
		return errResp
	}
	return nil
}

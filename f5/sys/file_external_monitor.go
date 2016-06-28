// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "e-xpert_solutions/f5-rest-client/f5"

// FileExternalMonitorConfigList holds a list of FileExternalMonitor configuration.
type FileExternalMonitorConfigList struct {
	Items    []FileExternalMonitorConfig `json:"items"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

// FileExternalMonitorConfig holds the configuration of a single FileExternalMonitor.
type FileExternalMonitorConfig struct {
	Checksum       string `json:"checksum"`
	CreateTime     string `json:"createTime"`
	CreatedBy      string `json:"createdBy"`
	FullPath       string `json:"fullPath"`
	Generation     int    `json:"generation"`
	Kind           string `json:"kind"`
	LastUpdateTime string `json:"lastUpdateTime"`
	Mode           int    `json:"mode"`
	Name           string `json:"name"`
	Partition      string `json:"partition"`
	Revision       int    `json:"revision"`
	SelfLink       string `json:"selfLink"`
	Size           int    `json:"size"`
	SystemPath     string `json:"systemPath"`
	UpdatedBy      string `json:"updatedBy"`
}

// FileExternalMonitorEndpoint represents the REST resource for managing FileExternalMonitor.
const FileExternalMonitorEndpoint = "/tm/sys/file/external-monitor"

// FileExternalMonitorResource provides an API to manage FileExternalMonitor configurations.
type FileExternalMonitorResource struct {
	c f5.Client
}

// ListAll  lists all the FileExternalMonitor configurations.
func (r *FileExternalMonitorResource) ListAll() (*FileExternalMonitorConfigList, error) {
	var list FileExternalMonitorConfigList
	if err := r.c.ReadQuery(BasePath+FileExternalMonitorEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileExternalMonitor configuration identified by id.
func (r *FileExternalMonitorResource) Get(id string) (*FileExternalMonitorConfig, error) {
	var item FileExternalMonitorConfig
	if err := r.c.ReadQuery(BasePath+FileExternalMonitorEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileExternalMonitor configuration.
func (r *FileExternalMonitorResource) Create(item FileExternalMonitorConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileExternalMonitorEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileExternalMonitor configuration identified by id.
func (r *FileExternalMonitorResource) Edit(id string, item FileExternalMonitorConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileExternalMonitorEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileExternalMonitor configuration identified by id.
func (r *FileExternalMonitorResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileExternalMonitorEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
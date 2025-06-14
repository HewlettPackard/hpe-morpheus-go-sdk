/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 8.0.7
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the UpdateNetworkEdgeClusterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkEdgeClusterRequest{}

// UpdateNetworkEdgeClusterRequest The parameters for update a network Edge Cluster is type dependent. The following lists the common parameters. See get a specific type to list available options for the network server type.
type UpdateNetworkEdgeClusterRequest struct {
	NetworkEdgeCluster   map[string]interface{} `json:"networkEdgeCluster,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _UpdateNetworkEdgeClusterRequest UpdateNetworkEdgeClusterRequest

// NewUpdateNetworkEdgeClusterRequest instantiates a new UpdateNetworkEdgeClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkEdgeClusterRequest() *UpdateNetworkEdgeClusterRequest {
	this := UpdateNetworkEdgeClusterRequest{}
	return &this
}

// NewUpdateNetworkEdgeClusterRequestWithDefaults instantiates a new UpdateNetworkEdgeClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkEdgeClusterRequestWithDefaults() *UpdateNetworkEdgeClusterRequest {
	this := UpdateNetworkEdgeClusterRequest{}
	return &this
}

// GetNetworkEdgeCluster returns the NetworkEdgeCluster field value if set, zero value otherwise.
func (o *UpdateNetworkEdgeClusterRequest) GetNetworkEdgeCluster() map[string]interface{} {
	if o == nil || IsNil(o.NetworkEdgeCluster) {
		var ret map[string]interface{}
		return ret
	}
	return o.NetworkEdgeCluster
}

// GetNetworkEdgeClusterOk returns a tuple with the NetworkEdgeCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkEdgeClusterRequest) GetNetworkEdgeClusterOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.NetworkEdgeCluster) {
		return map[string]interface{}{}, false
	}
	return o.NetworkEdgeCluster, true
}

// IsSetNetworkEdgeCluster returns a boolean if a field has been set.
func (o *UpdateNetworkEdgeClusterRequest) IsSetNetworkEdgeCluster() bool {
	if o != nil && !IsNil(o.NetworkEdgeCluster) {
		return true
	}

	return false
}

// SetNetworkEdgeCluster gets a reference to the given map[string]interface{} and assigns it to the NetworkEdgeCluster field.
func (o *UpdateNetworkEdgeClusterRequest) SetNetworkEdgeCluster(v map[string]interface{}) {
	o.NetworkEdgeCluster = v
}

func (o UpdateNetworkEdgeClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkEdgeClusterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkEdgeCluster) {
		toSerialize["networkEdgeCluster"] = o.NetworkEdgeCluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateNetworkEdgeClusterRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

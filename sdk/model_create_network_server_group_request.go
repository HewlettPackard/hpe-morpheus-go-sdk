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

// checks if the CreateNetworkServerGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkServerGroupRequest{}

// CreateNetworkServerGroupRequest struct for CreateNetworkServerGroupRequest
type CreateNetworkServerGroupRequest struct {
	Group                *CreateNetworkServerGroupRequestGroup `json:"group,omitempty"`
	AdditionalProperties map[string]interface{}                `json:",remain"`
}

type _CreateNetworkServerGroupRequest CreateNetworkServerGroupRequest

// NewCreateNetworkServerGroupRequest instantiates a new CreateNetworkServerGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkServerGroupRequest() *CreateNetworkServerGroupRequest {
	this := CreateNetworkServerGroupRequest{}
	return &this
}

// NewCreateNetworkServerGroupRequestWithDefaults instantiates a new CreateNetworkServerGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkServerGroupRequestWithDefaults() *CreateNetworkServerGroupRequest {
	this := CreateNetworkServerGroupRequest{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *CreateNetworkServerGroupRequest) GetGroup() CreateNetworkServerGroupRequestGroup {
	if o == nil || IsNil(o.Group) {
		var ret CreateNetworkServerGroupRequestGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkServerGroupRequest) GetGroupOk() (*CreateNetworkServerGroupRequestGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// IsSetGroup returns a boolean if a field has been set.
func (o *CreateNetworkServerGroupRequest) IsSetGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given CreateNetworkServerGroupRequestGroup and assigns it to the Group field.
func (o *CreateNetworkServerGroupRequest) SetGroup(v CreateNetworkServerGroupRequestGroup) {
	o.Group = &v
}

func (o CreateNetworkServerGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkServerGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *CreateNetworkServerGroupRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

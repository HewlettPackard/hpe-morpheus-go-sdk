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

// checks if the UpdateLayoutPermissionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLayoutPermissionsRequest{}

// UpdateLayoutPermissionsRequest struct for UpdateLayoutPermissionsRequest
type UpdateLayoutPermissionsRequest struct {
	InstanceTypeLayout   *UpdateLayoutPermissionsRequestInstanceTypeLayout `json:"instanceTypeLayout,omitempty"`
	AdditionalProperties map[string]interface{}                            `json:",remain"`
}

type _UpdateLayoutPermissionsRequest UpdateLayoutPermissionsRequest

// NewUpdateLayoutPermissionsRequest instantiates a new UpdateLayoutPermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLayoutPermissionsRequest() *UpdateLayoutPermissionsRequest {
	this := UpdateLayoutPermissionsRequest{}
	return &this
}

// NewUpdateLayoutPermissionsRequestWithDefaults instantiates a new UpdateLayoutPermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLayoutPermissionsRequestWithDefaults() *UpdateLayoutPermissionsRequest {
	this := UpdateLayoutPermissionsRequest{}
	return &this
}

// GetInstanceTypeLayout returns the InstanceTypeLayout field value if set, zero value otherwise.
func (o *UpdateLayoutPermissionsRequest) GetInstanceTypeLayout() UpdateLayoutPermissionsRequestInstanceTypeLayout {
	if o == nil || IsNil(o.InstanceTypeLayout) {
		var ret UpdateLayoutPermissionsRequestInstanceTypeLayout
		return ret
	}
	return *o.InstanceTypeLayout
}

// GetInstanceTypeLayoutOk returns a tuple with the InstanceTypeLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLayoutPermissionsRequest) GetInstanceTypeLayoutOk() (*UpdateLayoutPermissionsRequestInstanceTypeLayout, bool) {
	if o == nil || IsNil(o.InstanceTypeLayout) {
		return nil, false
	}
	return o.InstanceTypeLayout, true
}

// IsSetInstanceTypeLayout returns a boolean if a field has been set.
func (o *UpdateLayoutPermissionsRequest) IsSetInstanceTypeLayout() bool {
	if o != nil && !IsNil(o.InstanceTypeLayout) {
		return true
	}

	return false
}

// SetInstanceTypeLayout gets a reference to the given UpdateLayoutPermissionsRequestInstanceTypeLayout and assigns it to the InstanceTypeLayout field.
func (o *UpdateLayoutPermissionsRequest) SetInstanceTypeLayout(v UpdateLayoutPermissionsRequestInstanceTypeLayout) {
	o.InstanceTypeLayout = &v
}

func (o UpdateLayoutPermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLayoutPermissionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceTypeLayout) {
		toSerialize["instanceTypeLayout"] = o.InstanceTypeLayout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateLayoutPermissionsRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

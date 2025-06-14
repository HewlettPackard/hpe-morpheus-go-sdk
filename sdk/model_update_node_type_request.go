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

// checks if the UpdateNodeTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNodeTypeRequest{}

// UpdateNodeTypeRequest struct for UpdateNodeTypeRequest
type UpdateNodeTypeRequest struct {
	ContainerType        *UpdateNodeTypeRequestContainerType `json:"containerType,omitempty"`
	AdditionalProperties map[string]interface{}              `json:",remain"`
}

type _UpdateNodeTypeRequest UpdateNodeTypeRequest

// NewUpdateNodeTypeRequest instantiates a new UpdateNodeTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNodeTypeRequest() *UpdateNodeTypeRequest {
	this := UpdateNodeTypeRequest{}
	return &this
}

// NewUpdateNodeTypeRequestWithDefaults instantiates a new UpdateNodeTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNodeTypeRequestWithDefaults() *UpdateNodeTypeRequest {
	this := UpdateNodeTypeRequest{}
	return &this
}

// GetContainerType returns the ContainerType field value if set, zero value otherwise.
func (o *UpdateNodeTypeRequest) GetContainerType() UpdateNodeTypeRequestContainerType {
	if o == nil || IsNil(o.ContainerType) {
		var ret UpdateNodeTypeRequestContainerType
		return ret
	}
	return *o.ContainerType
}

// GetContainerTypeOk returns a tuple with the ContainerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNodeTypeRequest) GetContainerTypeOk() (*UpdateNodeTypeRequestContainerType, bool) {
	if o == nil || IsNil(o.ContainerType) {
		return nil, false
	}
	return o.ContainerType, true
}

// IsSetContainerType returns a boolean if a field has been set.
func (o *UpdateNodeTypeRequest) IsSetContainerType() bool {
	if o != nil && !IsNil(o.ContainerType) {
		return true
	}

	return false
}

// SetContainerType gets a reference to the given UpdateNodeTypeRequestContainerType and assigns it to the ContainerType field.
func (o *UpdateNodeTypeRequest) SetContainerType(v UpdateNodeTypeRequestContainerType) {
	o.ContainerType = &v
}

func (o UpdateNodeTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNodeTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerType) {
		toSerialize["containerType"] = o.ContainerType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateNodeTypeRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

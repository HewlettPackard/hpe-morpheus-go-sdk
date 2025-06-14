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

// checks if the UpdateOsTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOsTypeRequest{}

// UpdateOsTypeRequest struct for UpdateOsTypeRequest
type UpdateOsTypeRequest struct {
	OsType               *UpdateOsTypeRequestOsType `json:"osType,omitempty"`
	AdditionalProperties map[string]interface{}     `json:",remain"`
}

type _UpdateOsTypeRequest UpdateOsTypeRequest

// NewUpdateOsTypeRequest instantiates a new UpdateOsTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOsTypeRequest() *UpdateOsTypeRequest {
	this := UpdateOsTypeRequest{}
	return &this
}

// NewUpdateOsTypeRequestWithDefaults instantiates a new UpdateOsTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOsTypeRequestWithDefaults() *UpdateOsTypeRequest {
	this := UpdateOsTypeRequest{}
	return &this
}

// GetOsType returns the OsType field value if set, zero value otherwise.
func (o *UpdateOsTypeRequest) GetOsType() UpdateOsTypeRequestOsType {
	if o == nil || IsNil(o.OsType) {
		var ret UpdateOsTypeRequestOsType
		return ret
	}
	return *o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOsTypeRequest) GetOsTypeOk() (*UpdateOsTypeRequestOsType, bool) {
	if o == nil || IsNil(o.OsType) {
		return nil, false
	}
	return o.OsType, true
}

// IsSetOsType returns a boolean if a field has been set.
func (o *UpdateOsTypeRequest) IsSetOsType() bool {
	if o != nil && !IsNil(o.OsType) {
		return true
	}

	return false
}

// SetOsType gets a reference to the given UpdateOsTypeRequestOsType and assigns it to the OsType field.
func (o *UpdateOsTypeRequest) SetOsType(v UpdateOsTypeRequestOsType) {
	o.OsType = &v
}

func (o UpdateOsTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOsTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OsType) {
		toSerialize["osType"] = o.OsType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateOsTypeRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

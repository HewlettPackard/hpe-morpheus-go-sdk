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

// checks if the UpdateHostAssignTenantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateHostAssignTenantRequest{}

// UpdateHostAssignTenantRequest struct for UpdateHostAssignTenantRequest
type UpdateHostAssignTenantRequest struct {
	// Move associated instances to specified Tenant account.
	MoveAssociatedInstances *bool `json:"moveAssociatedInstances,omitempty"`
}

// NewUpdateHostAssignTenantRequest instantiates a new UpdateHostAssignTenantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateHostAssignTenantRequest() *UpdateHostAssignTenantRequest {
	this := UpdateHostAssignTenantRequest{}
	var moveAssociatedInstances bool = true
	this.MoveAssociatedInstances = &moveAssociatedInstances
	return &this
}

// NewUpdateHostAssignTenantRequestWithDefaults instantiates a new UpdateHostAssignTenantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateHostAssignTenantRequestWithDefaults() *UpdateHostAssignTenantRequest {
	this := UpdateHostAssignTenantRequest{}
	var moveAssociatedInstances bool = true
	this.MoveAssociatedInstances = &moveAssociatedInstances
	return &this
}

// GetMoveAssociatedInstances returns the MoveAssociatedInstances field value if set, zero value otherwise.
func (o *UpdateHostAssignTenantRequest) GetMoveAssociatedInstances() bool {
	if o == nil || IsNil(o.MoveAssociatedInstances) {
		var ret bool
		return ret
	}
	return *o.MoveAssociatedInstances
}

// GetMoveAssociatedInstancesOk returns a tuple with the MoveAssociatedInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostAssignTenantRequest) GetMoveAssociatedInstancesOk() (*bool, bool) {
	if o == nil || IsNil(o.MoveAssociatedInstances) {
		return nil, false
	}
	return o.MoveAssociatedInstances, true
}

// IsSetMoveAssociatedInstances returns a boolean if a field has been set.
func (o *UpdateHostAssignTenantRequest) IsSetMoveAssociatedInstances() bool {
	if o != nil && !IsNil(o.MoveAssociatedInstances) {
		return true
	}

	return false
}

// SetMoveAssociatedInstances gets a reference to the given bool and assigns it to the MoveAssociatedInstances field.
func (o *UpdateHostAssignTenantRequest) SetMoveAssociatedInstances(v bool) {
	o.MoveAssociatedInstances = &v
}

func (o UpdateHostAssignTenantRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateHostAssignTenantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MoveAssociatedInstances) {
		toSerialize["moveAssociatedInstances"] = o.MoveAssociatedInstances
	}
	return toSerialize, nil
}

type NullableUpdateHostAssignTenantRequest struct {
	value *UpdateHostAssignTenantRequest
	isSet bool
}

func (v NullableUpdateHostAssignTenantRequest) Get() *UpdateHostAssignTenantRequest {
	return v.value
}

func (v *NullableUpdateHostAssignTenantRequest) Set(val *UpdateHostAssignTenantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateHostAssignTenantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateHostAssignTenantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateHostAssignTenantRequest(val *UpdateHostAssignTenantRequest) *NullableUpdateHostAssignTenantRequest {
	return &NullableUpdateHostAssignTenantRequest{value: val, isSet: true}
}

func (v NullableUpdateHostAssignTenantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateHostAssignTenantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



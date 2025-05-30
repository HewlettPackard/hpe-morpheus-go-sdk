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

// checks if the AddClusterRequestClusterServerServerType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddClusterRequestClusterServerServerType{}

// AddClusterRequestClusterServerServerType Server type to create.  See `/api/server-types` for available server types for the cloud.
type AddClusterRequestClusterServerServerType struct {
	Id *int32 `json:"id,omitempty"`
}

// NewAddClusterRequestClusterServerServerType instantiates a new AddClusterRequestClusterServerServerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddClusterRequestClusterServerServerType() *AddClusterRequestClusterServerServerType {
	this := AddClusterRequestClusterServerServerType{}
	return &this
}

// NewAddClusterRequestClusterServerServerTypeWithDefaults instantiates a new AddClusterRequestClusterServerServerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddClusterRequestClusterServerServerTypeWithDefaults() *AddClusterRequestClusterServerServerType {
	this := AddClusterRequestClusterServerServerType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddClusterRequestClusterServerServerType) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddClusterRequestClusterServerServerType) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *AddClusterRequestClusterServerServerType) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AddClusterRequestClusterServerServerType) SetId(v int32) {
	o.Id = &v
}

func (o AddClusterRequestClusterServerServerType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddClusterRequestClusterServerServerType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAddClusterRequestClusterServerServerType struct {
	value *AddClusterRequestClusterServerServerType
	isSet bool
}

func (v NullableAddClusterRequestClusterServerServerType) Get() *AddClusterRequestClusterServerServerType {
	return v.value
}

func (v *NullableAddClusterRequestClusterServerServerType) Set(val *AddClusterRequestClusterServerServerType) {
	v.value = val
	v.isSet = true
}

func (v NullableAddClusterRequestClusterServerServerType) IsSet() bool {
	return v.isSet
}

func (v *NullableAddClusterRequestClusterServerServerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddClusterRequestClusterServerServerType(val *AddClusterRequestClusterServerServerType) *NullableAddClusterRequestClusterServerServerType {
	return &NullableAddClusterRequestClusterServerServerType{value: val, isSet: true}
}

func (v NullableAddClusterRequestClusterServerServerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddClusterRequestClusterServerServerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



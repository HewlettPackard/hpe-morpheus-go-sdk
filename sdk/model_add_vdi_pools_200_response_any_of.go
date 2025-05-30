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

// checks if the AddVDIPools200ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddVDIPools200ResponseAnyOf{}

// AddVDIPools200ResponseAnyOf struct for AddVDIPools200ResponseAnyOf
type AddVDIPools200ResponseAnyOf struct {
	VdiPool *ListVDIPools200ResponseAllOfVdiPoolsInner `json:"vdiPool,omitempty"`
}

// NewAddVDIPools200ResponseAnyOf instantiates a new AddVDIPools200ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVDIPools200ResponseAnyOf() *AddVDIPools200ResponseAnyOf {
	this := AddVDIPools200ResponseAnyOf{}
	return &this
}

// NewAddVDIPools200ResponseAnyOfWithDefaults instantiates a new AddVDIPools200ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVDIPools200ResponseAnyOfWithDefaults() *AddVDIPools200ResponseAnyOf {
	this := AddVDIPools200ResponseAnyOf{}
	return &this
}

// GetVdiPool returns the VdiPool field value if set, zero value otherwise.
func (o *AddVDIPools200ResponseAnyOf) GetVdiPool() ListVDIPools200ResponseAllOfVdiPoolsInner {
	if o == nil || IsNil(o.VdiPool) {
		var ret ListVDIPools200ResponseAllOfVdiPoolsInner
		return ret
	}
	return *o.VdiPool
}

// GetVdiPoolOk returns a tuple with the VdiPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVDIPools200ResponseAnyOf) GetVdiPoolOk() (*ListVDIPools200ResponseAllOfVdiPoolsInner, bool) {
	if o == nil || IsNil(o.VdiPool) {
		return nil, false
	}
	return o.VdiPool, true
}

// IsSetVdiPool returns a boolean if a field has been set.
func (o *AddVDIPools200ResponseAnyOf) IsSetVdiPool() bool {
	if o != nil && !IsNil(o.VdiPool) {
		return true
	}

	return false
}

// SetVdiPool gets a reference to the given ListVDIPools200ResponseAllOfVdiPoolsInner and assigns it to the VdiPool field.
func (o *AddVDIPools200ResponseAnyOf) SetVdiPool(v ListVDIPools200ResponseAllOfVdiPoolsInner) {
	o.VdiPool = &v
}

func (o AddVDIPools200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddVDIPools200ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VdiPool) {
		toSerialize["vdiPool"] = o.VdiPool
	}
	return toSerialize, nil
}

type NullableAddVDIPools200ResponseAnyOf struct {
	value *AddVDIPools200ResponseAnyOf
	isSet bool
}

func (v NullableAddVDIPools200ResponseAnyOf) Get() *AddVDIPools200ResponseAnyOf {
	return v.value
}

func (v *NullableAddVDIPools200ResponseAnyOf) Set(val *AddVDIPools200ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVDIPools200ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVDIPools200ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVDIPools200ResponseAnyOf(val *AddVDIPools200ResponseAnyOf) *NullableAddVDIPools200ResponseAnyOf {
	return &NullableAddVDIPools200ResponseAnyOf{value: val, isSet: true}
}

func (v NullableAddVDIPools200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVDIPools200ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



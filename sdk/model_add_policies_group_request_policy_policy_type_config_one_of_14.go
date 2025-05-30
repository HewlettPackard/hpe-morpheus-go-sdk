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

// checks if the AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14{}

// AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14 - Max Load Balancer Pools 
type AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14 struct {
	MaxPools *string `json:"maxPools,omitempty"`
}

// NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14 instantiates a new AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14() *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14 {
	this := AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14{}
	return &this
}

// NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14WithDefaults instantiates a new AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14WithDefaults() *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14 {
	this := AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14{}
	return &this
}

// GetMaxPools returns the MaxPools field value if set, zero value otherwise.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) GetMaxPools() string {
	if o == nil || IsNil(o.MaxPools) {
		var ret string
		return ret
	}
	return *o.MaxPools
}

// GetMaxPoolsOk returns a tuple with the MaxPools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) GetMaxPoolsOk() (*string, bool) {
	if o == nil || IsNil(o.MaxPools) {
		return nil, false
	}
	return o.MaxPools, true
}

// IsSetMaxPools returns a boolean if a field has been set.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) IsSetMaxPools() bool {
	if o != nil && !IsNil(o.MaxPools) {
		return true
	}

	return false
}

// SetMaxPools gets a reference to the given string and assigns it to the MaxPools field.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) SetMaxPools(v string) {
	o.MaxPools = &v
}

func (o AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxPools) {
		toSerialize["maxPools"] = o.MaxPools
	}
	return toSerialize, nil
}

type NullableAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14 struct {
	value *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14
	isSet bool
}

func (v NullableAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) Get() *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14 {
	return v.value
}

func (v *NullableAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) Set(val *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14(val *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) *NullableAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14 {
	return &NullableAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14{value: val, isSet: true}
}

func (v NullableAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf14) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



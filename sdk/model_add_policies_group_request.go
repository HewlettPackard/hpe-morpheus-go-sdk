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
	"bytes"
	"fmt"
)

// checks if the AddPoliciesGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPoliciesGroupRequest{}

// AddPoliciesGroupRequest struct for AddPoliciesGroupRequest
type AddPoliciesGroupRequest struct {
	Policy AddPoliciesGroupRequestPolicy `json:"policy"`
}

type _AddPoliciesGroupRequest AddPoliciesGroupRequest

// NewAddPoliciesGroupRequest instantiates a new AddPoliciesGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPoliciesGroupRequest(policy AddPoliciesGroupRequestPolicy) *AddPoliciesGroupRequest {
	this := AddPoliciesGroupRequest{}
	this.Policy = policy
	return &this
}

// NewAddPoliciesGroupRequestWithDefaults instantiates a new AddPoliciesGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPoliciesGroupRequestWithDefaults() *AddPoliciesGroupRequest {
	this := AddPoliciesGroupRequest{}
	return &this
}

// GetPolicy returns the Policy field value
func (o *AddPoliciesGroupRequest) GetPolicy() AddPoliciesGroupRequestPolicy {
	if o == nil {
		var ret AddPoliciesGroupRequestPolicy
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *AddPoliciesGroupRequest) GetPolicyOk() (*AddPoliciesGroupRequestPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *AddPoliciesGroupRequest) SetPolicy(v AddPoliciesGroupRequestPolicy) {
	o.Policy = v
}

func (o AddPoliciesGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPoliciesGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["policy"] = o.Policy
	return toSerialize, nil
}

func (o *AddPoliciesGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"policy",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddPoliciesGroupRequest := _AddPoliciesGroupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddPoliciesGroupRequest)

	if err != nil {
		return err
	}

	*o = AddPoliciesGroupRequest(varAddPoliciesGroupRequest)

	return err
}

type NullableAddPoliciesGroupRequest struct {
	value *AddPoliciesGroupRequest
	isSet bool
}

func (v NullableAddPoliciesGroupRequest) Get() *AddPoliciesGroupRequest {
	return v.value
}

func (v *NullableAddPoliciesGroupRequest) Set(val *AddPoliciesGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPoliciesGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPoliciesGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPoliciesGroupRequest(val *AddPoliciesGroupRequest) *NullableAddPoliciesGroupRequest {
	return &NullableAddPoliciesGroupRequest{value: val, isSet: true}
}

func (v NullableAddPoliciesGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPoliciesGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



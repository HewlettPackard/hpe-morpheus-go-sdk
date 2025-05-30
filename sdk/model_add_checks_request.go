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

// checks if the AddChecksRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddChecksRequest{}

// AddChecksRequest struct for AddChecksRequest
type AddChecksRequest struct {
	Check AddChecksRequestCheck `json:"check"`
}

type _AddChecksRequest AddChecksRequest

// NewAddChecksRequest instantiates a new AddChecksRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddChecksRequest(check AddChecksRequestCheck) *AddChecksRequest {
	this := AddChecksRequest{}
	this.Check = check
	return &this
}

// NewAddChecksRequestWithDefaults instantiates a new AddChecksRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddChecksRequestWithDefaults() *AddChecksRequest {
	this := AddChecksRequest{}
	return &this
}

// GetCheck returns the Check field value
func (o *AddChecksRequest) GetCheck() AddChecksRequestCheck {
	if o == nil {
		var ret AddChecksRequestCheck
		return ret
	}

	return o.Check
}

// GetCheckOk returns a tuple with the Check field value
// and a boolean to check if the value has been set.
func (o *AddChecksRequest) GetCheckOk() (*AddChecksRequestCheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Check, true
}

// SetCheck sets field value
func (o *AddChecksRequest) SetCheck(v AddChecksRequestCheck) {
	o.Check = v
}

func (o AddChecksRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddChecksRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["check"] = o.Check
	return toSerialize, nil
}

func (o *AddChecksRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"check",
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

	varAddChecksRequest := _AddChecksRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddChecksRequest)

	if err != nil {
		return err
	}

	*o = AddChecksRequest(varAddChecksRequest)

	return err
}

type NullableAddChecksRequest struct {
	value *AddChecksRequest
	isSet bool
}

func (v NullableAddChecksRequest) Get() *AddChecksRequest {
	return v.value
}

func (v *NullableAddChecksRequest) Set(val *AddChecksRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddChecksRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddChecksRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddChecksRequest(val *AddChecksRequest) *NullableAddChecksRequest {
	return &NullableAddChecksRequest{value: val, isSet: true}
}

func (v NullableAddChecksRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddChecksRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



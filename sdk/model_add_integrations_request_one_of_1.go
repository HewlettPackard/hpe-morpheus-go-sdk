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

// checks if the AddIntegrationsRequestOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddIntegrationsRequestOneOf1{}

// AddIntegrationsRequestOneOf1 struct for AddIntegrationsRequestOneOf1
type AddIntegrationsRequestOneOf1 struct {
	Integration AddIntegrationsRequestOneOf1Integration `json:"integration"`
}

type _AddIntegrationsRequestOneOf1 AddIntegrationsRequestOneOf1

// NewAddIntegrationsRequestOneOf1 instantiates a new AddIntegrationsRequestOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIntegrationsRequestOneOf1(integration AddIntegrationsRequestOneOf1Integration) *AddIntegrationsRequestOneOf1 {
	this := AddIntegrationsRequestOneOf1{}
	this.Integration = integration
	return &this
}

// NewAddIntegrationsRequestOneOf1WithDefaults instantiates a new AddIntegrationsRequestOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIntegrationsRequestOneOf1WithDefaults() *AddIntegrationsRequestOneOf1 {
	this := AddIntegrationsRequestOneOf1{}
	return &this
}

// GetIntegration returns the Integration field value
func (o *AddIntegrationsRequestOneOf1) GetIntegration() AddIntegrationsRequestOneOf1Integration {
	if o == nil {
		var ret AddIntegrationsRequestOneOf1Integration
		return ret
	}

	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *AddIntegrationsRequestOneOf1) GetIntegrationOk() (*AddIntegrationsRequestOneOf1Integration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value
func (o *AddIntegrationsRequestOneOf1) SetIntegration(v AddIntegrationsRequestOneOf1Integration) {
	o.Integration = v
}

func (o AddIntegrationsRequestOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddIntegrationsRequestOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["integration"] = o.Integration
	return toSerialize, nil
}

func (o *AddIntegrationsRequestOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"integration",
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

	varAddIntegrationsRequestOneOf1 := _AddIntegrationsRequestOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddIntegrationsRequestOneOf1)

	if err != nil {
		return err
	}

	*o = AddIntegrationsRequestOneOf1(varAddIntegrationsRequestOneOf1)

	return err
}

type NullableAddIntegrationsRequestOneOf1 struct {
	value *AddIntegrationsRequestOneOf1
	isSet bool
}

func (v NullableAddIntegrationsRequestOneOf1) Get() *AddIntegrationsRequestOneOf1 {
	return v.value
}

func (v *NullableAddIntegrationsRequestOneOf1) Set(val *AddIntegrationsRequestOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableAddIntegrationsRequestOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableAddIntegrationsRequestOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddIntegrationsRequestOneOf1(val *AddIntegrationsRequestOneOf1) *NullableAddIntegrationsRequestOneOf1 {
	return &NullableAddIntegrationsRequestOneOf1{value: val, isSet: true}
}

func (v NullableAddIntegrationsRequestOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddIntegrationsRequestOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



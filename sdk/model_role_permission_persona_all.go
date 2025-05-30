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

// checks if the RolePermissionPersonaAll type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolePermissionPersonaAll{}

// RolePermissionPersonaAll struct for RolePermissionPersonaAll
type RolePermissionPersonaAll struct {
	// Apply to all personas
	AllPersonas bool `json:"allPersonas"`
	// The new access level.
	Access string `json:"access"`
}

type _RolePermissionPersonaAll RolePermissionPersonaAll

// NewRolePermissionPersonaAll instantiates a new RolePermissionPersonaAll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermissionPersonaAll(allPersonas bool, access string) *RolePermissionPersonaAll {
	this := RolePermissionPersonaAll{}
	this.AllPersonas = allPersonas
	this.Access = access
	return &this
}

// NewRolePermissionPersonaAllWithDefaults instantiates a new RolePermissionPersonaAll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionPersonaAllWithDefaults() *RolePermissionPersonaAll {
	this := RolePermissionPersonaAll{}
	return &this
}

// GetAllPersonas returns the AllPersonas field value
func (o *RolePermissionPersonaAll) GetAllPersonas() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllPersonas
}

// GetAllPersonasOk returns a tuple with the AllPersonas field value
// and a boolean to check if the value has been set.
func (o *RolePermissionPersonaAll) GetAllPersonasOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllPersonas, true
}

// SetAllPersonas sets field value
func (o *RolePermissionPersonaAll) SetAllPersonas(v bool) {
	o.AllPersonas = v
}

// GetAccess returns the Access field value
func (o *RolePermissionPersonaAll) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *RolePermissionPersonaAll) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *RolePermissionPersonaAll) SetAccess(v string) {
	o.Access = v
}

func (o RolePermissionPersonaAll) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolePermissionPersonaAll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allPersonas"] = o.AllPersonas
	toSerialize["access"] = o.Access
	return toSerialize, nil
}

func (o *RolePermissionPersonaAll) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allPersonas",
		"access",
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

	varRolePermissionPersonaAll := _RolePermissionPersonaAll{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRolePermissionPersonaAll)

	if err != nil {
		return err
	}

	*o = RolePermissionPersonaAll(varRolePermissionPersonaAll)

	return err
}

type NullableRolePermissionPersonaAll struct {
	value *RolePermissionPersonaAll
	isSet bool
}

func (v NullableRolePermissionPersonaAll) Get() *RolePermissionPersonaAll {
	return v.value
}

func (v *NullableRolePermissionPersonaAll) Set(val *RolePermissionPersonaAll) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePermissionPersonaAll) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePermissionPersonaAll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePermissionPersonaAll(val *RolePermissionPersonaAll) *NullableRolePermissionPersonaAll {
	return &NullableRolePermissionPersonaAll{value: val, isSet: true}
}

func (v NullableRolePermissionPersonaAll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePermissionPersonaAll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



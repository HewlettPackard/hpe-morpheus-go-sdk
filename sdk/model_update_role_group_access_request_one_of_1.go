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

// checks if the UpdateRoleGroupAccessRequestOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoleGroupAccessRequestOneOf1{}

// UpdateRoleGroupAccessRequestOneOf1 struct for UpdateRoleGroupAccessRequestOneOf1
type UpdateRoleGroupAccessRequestOneOf1 struct {
	// Apply to all groups (site)
	AllGroups bool `json:"allGroups"`
	// The new access level.
	Access string `json:"access"`
}

type _UpdateRoleGroupAccessRequestOneOf1 UpdateRoleGroupAccessRequestOneOf1

// NewUpdateRoleGroupAccessRequestOneOf1 instantiates a new UpdateRoleGroupAccessRequestOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleGroupAccessRequestOneOf1(allGroups bool, access string) *UpdateRoleGroupAccessRequestOneOf1 {
	this := UpdateRoleGroupAccessRequestOneOf1{}
	this.AllGroups = allGroups
	this.Access = access
	return &this
}

// NewUpdateRoleGroupAccessRequestOneOf1WithDefaults instantiates a new UpdateRoleGroupAccessRequestOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleGroupAccessRequestOneOf1WithDefaults() *UpdateRoleGroupAccessRequestOneOf1 {
	this := UpdateRoleGroupAccessRequestOneOf1{}
	return &this
}

// GetAllGroups returns the AllGroups field value
func (o *UpdateRoleGroupAccessRequestOneOf1) GetAllGroups() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllGroups
}

// GetAllGroupsOk returns a tuple with the AllGroups field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleGroupAccessRequestOneOf1) GetAllGroupsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllGroups, true
}

// SetAllGroups sets field value
func (o *UpdateRoleGroupAccessRequestOneOf1) SetAllGroups(v bool) {
	o.AllGroups = v
}

// GetAccess returns the Access field value
func (o *UpdateRoleGroupAccessRequestOneOf1) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleGroupAccessRequestOneOf1) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *UpdateRoleGroupAccessRequestOneOf1) SetAccess(v string) {
	o.Access = v
}

func (o UpdateRoleGroupAccessRequestOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoleGroupAccessRequestOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allGroups"] = o.AllGroups
	toSerialize["access"] = o.Access
	return toSerialize, nil
}

func (o *UpdateRoleGroupAccessRequestOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allGroups",
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

	varUpdateRoleGroupAccessRequestOneOf1 := _UpdateRoleGroupAccessRequestOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateRoleGroupAccessRequestOneOf1)

	if err != nil {
		return err
	}

	*o = UpdateRoleGroupAccessRequestOneOf1(varUpdateRoleGroupAccessRequestOneOf1)

	return err
}

type NullableUpdateRoleGroupAccessRequestOneOf1 struct {
	value *UpdateRoleGroupAccessRequestOneOf1
	isSet bool
}

func (v NullableUpdateRoleGroupAccessRequestOneOf1) Get() *UpdateRoleGroupAccessRequestOneOf1 {
	return v.value
}

func (v *NullableUpdateRoleGroupAccessRequestOneOf1) Set(val *UpdateRoleGroupAccessRequestOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoleGroupAccessRequestOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoleGroupAccessRequestOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoleGroupAccessRequestOneOf1(val *UpdateRoleGroupAccessRequestOneOf1) *NullableUpdateRoleGroupAccessRequestOneOf1 {
	return &NullableUpdateRoleGroupAccessRequestOneOf1{value: val, isSet: true}
}

func (v NullableUpdateRoleGroupAccessRequestOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoleGroupAccessRequestOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



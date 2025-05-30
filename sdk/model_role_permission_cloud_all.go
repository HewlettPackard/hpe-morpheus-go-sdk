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

// checks if the RolePermissionCloudAll type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolePermissionCloudAll{}

// RolePermissionCloudAll struct for RolePermissionCloudAll
type RolePermissionCloudAll struct {
	// Apply to all clouds
	AllClouds bool `json:"allClouds"`
	// The new access level.
	Access string `json:"access"`
}

type _RolePermissionCloudAll RolePermissionCloudAll

// NewRolePermissionCloudAll instantiates a new RolePermissionCloudAll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermissionCloudAll(allClouds bool, access string) *RolePermissionCloudAll {
	this := RolePermissionCloudAll{}
	this.AllClouds = allClouds
	this.Access = access
	return &this
}

// NewRolePermissionCloudAllWithDefaults instantiates a new RolePermissionCloudAll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionCloudAllWithDefaults() *RolePermissionCloudAll {
	this := RolePermissionCloudAll{}
	return &this
}

// GetAllClouds returns the AllClouds field value
func (o *RolePermissionCloudAll) GetAllClouds() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllClouds
}

// GetAllCloudsOk returns a tuple with the AllClouds field value
// and a boolean to check if the value has been set.
func (o *RolePermissionCloudAll) GetAllCloudsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllClouds, true
}

// SetAllClouds sets field value
func (o *RolePermissionCloudAll) SetAllClouds(v bool) {
	o.AllClouds = v
}

// GetAccess returns the Access field value
func (o *RolePermissionCloudAll) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *RolePermissionCloudAll) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *RolePermissionCloudAll) SetAccess(v string) {
	o.Access = v
}

func (o RolePermissionCloudAll) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolePermissionCloudAll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allClouds"] = o.AllClouds
	toSerialize["access"] = o.Access
	return toSerialize, nil
}

func (o *RolePermissionCloudAll) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allClouds",
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

	varRolePermissionCloudAll := _RolePermissionCloudAll{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRolePermissionCloudAll)

	if err != nil {
		return err
	}

	*o = RolePermissionCloudAll(varRolePermissionCloudAll)

	return err
}

type NullableRolePermissionCloudAll struct {
	value *RolePermissionCloudAll
	isSet bool
}

func (v NullableRolePermissionCloudAll) Get() *RolePermissionCloudAll {
	return v.value
}

func (v *NullableRolePermissionCloudAll) Set(val *RolePermissionCloudAll) {
	v.value = val
	v.isSet = true
}

func (v NullableRolePermissionCloudAll) IsSet() bool {
	return v.isSet
}

func (v *NullableRolePermissionCloudAll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolePermissionCloudAll(val *RolePermissionCloudAll) *NullableRolePermissionCloudAll {
	return &NullableRolePermissionCloudAll{value: val, isSet: true}
}

func (v NullableRolePermissionCloudAll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolePermissionCloudAll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



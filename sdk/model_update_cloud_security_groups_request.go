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

// checks if the UpdateCloudSecurityGroupsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCloudSecurityGroupsRequest{}

// UpdateCloudSecurityGroupsRequest struct for UpdateCloudSecurityGroupsRequest
type UpdateCloudSecurityGroupsRequest struct {
	SecurityGroupIds []int64 `json:"securityGroupIds"`
}

type _UpdateCloudSecurityGroupsRequest UpdateCloudSecurityGroupsRequest

// NewUpdateCloudSecurityGroupsRequest instantiates a new UpdateCloudSecurityGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCloudSecurityGroupsRequest(securityGroupIds []int64) *UpdateCloudSecurityGroupsRequest {
	this := UpdateCloudSecurityGroupsRequest{}
	this.SecurityGroupIds = securityGroupIds
	return &this
}

// NewUpdateCloudSecurityGroupsRequestWithDefaults instantiates a new UpdateCloudSecurityGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCloudSecurityGroupsRequestWithDefaults() *UpdateCloudSecurityGroupsRequest {
	this := UpdateCloudSecurityGroupsRequest{}
	return &this
}

// GetSecurityGroupIds returns the SecurityGroupIds field value
func (o *UpdateCloudSecurityGroupsRequest) GetSecurityGroupIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value
// and a boolean to check if the value has been set.
func (o *UpdateCloudSecurityGroupsRequest) GetSecurityGroupIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecurityGroupIds, true
}

// SetSecurityGroupIds sets field value
func (o *UpdateCloudSecurityGroupsRequest) SetSecurityGroupIds(v []int64) {
	o.SecurityGroupIds = v
}

func (o UpdateCloudSecurityGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCloudSecurityGroupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["securityGroupIds"] = o.SecurityGroupIds
	return toSerialize, nil
}

func (o *UpdateCloudSecurityGroupsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"securityGroupIds",
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

	varUpdateCloudSecurityGroupsRequest := _UpdateCloudSecurityGroupsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateCloudSecurityGroupsRequest)

	if err != nil {
		return err
	}

	*o = UpdateCloudSecurityGroupsRequest(varUpdateCloudSecurityGroupsRequest)

	return err
}

type NullableUpdateCloudSecurityGroupsRequest struct {
	value *UpdateCloudSecurityGroupsRequest
	isSet bool
}

func (v NullableUpdateCloudSecurityGroupsRequest) Get() *UpdateCloudSecurityGroupsRequest {
	return v.value
}

func (v *NullableUpdateCloudSecurityGroupsRequest) Set(val *UpdateCloudSecurityGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCloudSecurityGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCloudSecurityGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCloudSecurityGroupsRequest(val *UpdateCloudSecurityGroupsRequest) *NullableUpdateCloudSecurityGroupsRequest {
	return &NullableUpdateCloudSecurityGroupsRequest{value: val, isSet: true}
}

func (v NullableUpdateCloudSecurityGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCloudSecurityGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



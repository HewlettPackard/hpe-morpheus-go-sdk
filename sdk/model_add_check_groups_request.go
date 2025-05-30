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

// checks if the AddCheckGroupsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCheckGroupsRequest{}

// AddCheckGroupsRequest struct for AddCheckGroupsRequest
type AddCheckGroupsRequest struct {
	CheckGroup AddCheckGroupsRequestCheckGroup `json:"checkGroup"`
}

type _AddCheckGroupsRequest AddCheckGroupsRequest

// NewAddCheckGroupsRequest instantiates a new AddCheckGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCheckGroupsRequest(checkGroup AddCheckGroupsRequestCheckGroup) *AddCheckGroupsRequest {
	this := AddCheckGroupsRequest{}
	this.CheckGroup = checkGroup
	return &this
}

// NewAddCheckGroupsRequestWithDefaults instantiates a new AddCheckGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCheckGroupsRequestWithDefaults() *AddCheckGroupsRequest {
	this := AddCheckGroupsRequest{}
	return &this
}

// GetCheckGroup returns the CheckGroup field value
func (o *AddCheckGroupsRequest) GetCheckGroup() AddCheckGroupsRequestCheckGroup {
	if o == nil {
		var ret AddCheckGroupsRequestCheckGroup
		return ret
	}

	return o.CheckGroup
}

// GetCheckGroupOk returns a tuple with the CheckGroup field value
// and a boolean to check if the value has been set.
func (o *AddCheckGroupsRequest) GetCheckGroupOk() (*AddCheckGroupsRequestCheckGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckGroup, true
}

// SetCheckGroup sets field value
func (o *AddCheckGroupsRequest) SetCheckGroup(v AddCheckGroupsRequestCheckGroup) {
	o.CheckGroup = v
}

func (o AddCheckGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCheckGroupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["checkGroup"] = o.CheckGroup
	return toSerialize, nil
}

func (o *AddCheckGroupsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"checkGroup",
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

	varAddCheckGroupsRequest := _AddCheckGroupsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddCheckGroupsRequest)

	if err != nil {
		return err
	}

	*o = AddCheckGroupsRequest(varAddCheckGroupsRequest)

	return err
}

type NullableAddCheckGroupsRequest struct {
	value *AddCheckGroupsRequest
	isSet bool
}

func (v NullableAddCheckGroupsRequest) Get() *AddCheckGroupsRequest {
	return v.value
}

func (v *NullableAddCheckGroupsRequest) Set(val *AddCheckGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCheckGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCheckGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCheckGroupsRequest(val *AddCheckGroupsRequest) *NullableAddCheckGroupsRequest {
	return &NullableAddCheckGroupsRequest{value: val, isSet: true}
}

func (v NullableAddCheckGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCheckGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



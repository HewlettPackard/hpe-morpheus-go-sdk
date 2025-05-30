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

// checks if the RemoveAppInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveAppInstanceRequest{}

// RemoveAppInstanceRequest struct for RemoveAppInstanceRequest
type RemoveAppInstanceRequest struct {
	// The ID of the instance being removed
	InstanceId int64 `json:"instanceId"`
}

type _RemoveAppInstanceRequest RemoveAppInstanceRequest

// NewRemoveAppInstanceRequest instantiates a new RemoveAppInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveAppInstanceRequest(instanceId int64) *RemoveAppInstanceRequest {
	this := RemoveAppInstanceRequest{}
	this.InstanceId = instanceId
	return &this
}

// NewRemoveAppInstanceRequestWithDefaults instantiates a new RemoveAppInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveAppInstanceRequestWithDefaults() *RemoveAppInstanceRequest {
	this := RemoveAppInstanceRequest{}
	return &this
}

// GetInstanceId returns the InstanceId field value
func (o *RemoveAppInstanceRequest) GetInstanceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *RemoveAppInstanceRequest) GetInstanceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *RemoveAppInstanceRequest) SetInstanceId(v int64) {
	o.InstanceId = v
}

func (o RemoveAppInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveAppInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceId"] = o.InstanceId
	return toSerialize, nil
}

func (o *RemoveAppInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instanceId",
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

	varRemoveAppInstanceRequest := _RemoveAppInstanceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRemoveAppInstanceRequest)

	if err != nil {
		return err
	}

	*o = RemoveAppInstanceRequest(varRemoveAppInstanceRequest)

	return err
}

type NullableRemoveAppInstanceRequest struct {
	value *RemoveAppInstanceRequest
	isSet bool
}

func (v NullableRemoveAppInstanceRequest) Get() *RemoveAppInstanceRequest {
	return v.value
}

func (v *NullableRemoveAppInstanceRequest) Set(val *RemoveAppInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveAppInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveAppInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveAppInstanceRequest(val *RemoveAppInstanceRequest) *NullableRemoveAppInstanceRequest {
	return &NullableRemoveAppInstanceRequest{value: val, isSet: true}
}

func (v NullableRemoveAppInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveAppInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



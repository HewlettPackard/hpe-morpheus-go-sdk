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

// checks if the UpdateTasksRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTasksRequest{}

// UpdateTasksRequest struct for UpdateTasksRequest
type UpdateTasksRequest struct {
	Task UpdateTasksRequestTask `json:"task"`
}

type _UpdateTasksRequest UpdateTasksRequest

// NewUpdateTasksRequest instantiates a new UpdateTasksRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTasksRequest(task UpdateTasksRequestTask) *UpdateTasksRequest {
	this := UpdateTasksRequest{}
	this.Task = task
	return &this
}

// NewUpdateTasksRequestWithDefaults instantiates a new UpdateTasksRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTasksRequestWithDefaults() *UpdateTasksRequest {
	this := UpdateTasksRequest{}
	return &this
}

// GetTask returns the Task field value
func (o *UpdateTasksRequest) GetTask() UpdateTasksRequestTask {
	if o == nil {
		var ret UpdateTasksRequestTask
		return ret
	}

	return o.Task
}

// GetTaskOk returns a tuple with the Task field value
// and a boolean to check if the value has been set.
func (o *UpdateTasksRequest) GetTaskOk() (*UpdateTasksRequestTask, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Task, true
}

// SetTask sets field value
func (o *UpdateTasksRequest) SetTask(v UpdateTasksRequestTask) {
	o.Task = v
}

func (o UpdateTasksRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTasksRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["task"] = o.Task
	return toSerialize, nil
}

func (o *UpdateTasksRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"task",
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

	varUpdateTasksRequest := _UpdateTasksRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateTasksRequest)

	if err != nil {
		return err
	}

	*o = UpdateTasksRequest(varUpdateTasksRequest)

	return err
}

type NullableUpdateTasksRequest struct {
	value *UpdateTasksRequest
	isSet bool
}

func (v NullableUpdateTasksRequest) Get() *UpdateTasksRequest {
	return v.value
}

func (v *NullableUpdateTasksRequest) Set(val *UpdateTasksRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTasksRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTasksRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTasksRequest(val *UpdateTasksRequest) *NullableUpdateTasksRequest {
	return &NullableUpdateTasksRequest{value: val, isSet: true}
}

func (v NullableUpdateTasksRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTasksRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the AddPowerSchedulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPowerSchedulesRequest{}

// AddPowerSchedulesRequest struct for AddPowerSchedulesRequest
type AddPowerSchedulesRequest struct {
	Schedule AddPowerSchedulesRequestSchedule `json:"schedule"`
}

type _AddPowerSchedulesRequest AddPowerSchedulesRequest

// NewAddPowerSchedulesRequest instantiates a new AddPowerSchedulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPowerSchedulesRequest(schedule AddPowerSchedulesRequestSchedule) *AddPowerSchedulesRequest {
	this := AddPowerSchedulesRequest{}
	this.Schedule = schedule
	return &this
}

// NewAddPowerSchedulesRequestWithDefaults instantiates a new AddPowerSchedulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPowerSchedulesRequestWithDefaults() *AddPowerSchedulesRequest {
	this := AddPowerSchedulesRequest{}
	return &this
}

// GetSchedule returns the Schedule field value
func (o *AddPowerSchedulesRequest) GetSchedule() AddPowerSchedulesRequestSchedule {
	if o == nil {
		var ret AddPowerSchedulesRequestSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *AddPowerSchedulesRequest) GetScheduleOk() (*AddPowerSchedulesRequestSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *AddPowerSchedulesRequest) SetSchedule(v AddPowerSchedulesRequestSchedule) {
	o.Schedule = v
}

func (o AddPowerSchedulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPowerSchedulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schedule"] = o.Schedule
	return toSerialize, nil
}

func (o *AddPowerSchedulesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"schedule",
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

	varAddPowerSchedulesRequest := _AddPowerSchedulesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddPowerSchedulesRequest)

	if err != nil {
		return err
	}

	*o = AddPowerSchedulesRequest(varAddPowerSchedulesRequest)

	return err
}

type NullableAddPowerSchedulesRequest struct {
	value *AddPowerSchedulesRequest
	isSet bool
}

func (v NullableAddPowerSchedulesRequest) Get() *AddPowerSchedulesRequest {
	return v.value
}

func (v *NullableAddPowerSchedulesRequest) Set(val *AddPowerSchedulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPowerSchedulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPowerSchedulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPowerSchedulesRequest(val *AddPowerSchedulesRequest) *NullableAddPowerSchedulesRequest {
	return &NullableAddPowerSchedulesRequest{value: val, isSet: true}
}

func (v NullableAddPowerSchedulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPowerSchedulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the UpdateBackupsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBackupsRequest{}

// UpdateBackupsRequest struct for UpdateBackupsRequest
type UpdateBackupsRequest struct {
	Backup UpdateBackupsRequestBackup `json:"backup"`
}

type _UpdateBackupsRequest UpdateBackupsRequest

// NewUpdateBackupsRequest instantiates a new UpdateBackupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBackupsRequest(backup UpdateBackupsRequestBackup) *UpdateBackupsRequest {
	this := UpdateBackupsRequest{}
	this.Backup = backup
	return &this
}

// NewUpdateBackupsRequestWithDefaults instantiates a new UpdateBackupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBackupsRequestWithDefaults() *UpdateBackupsRequest {
	this := UpdateBackupsRequest{}
	return &this
}

// GetBackup returns the Backup field value
func (o *UpdateBackupsRequest) GetBackup() UpdateBackupsRequestBackup {
	if o == nil {
		var ret UpdateBackupsRequestBackup
		return ret
	}

	return o.Backup
}

// GetBackupOk returns a tuple with the Backup field value
// and a boolean to check if the value has been set.
func (o *UpdateBackupsRequest) GetBackupOk() (*UpdateBackupsRequestBackup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Backup, true
}

// SetBackup sets field value
func (o *UpdateBackupsRequest) SetBackup(v UpdateBackupsRequestBackup) {
	o.Backup = v
}

func (o UpdateBackupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateBackupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["backup"] = o.Backup
	return toSerialize, nil
}

func (o *UpdateBackupsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"backup",
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

	varUpdateBackupsRequest := _UpdateBackupsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateBackupsRequest)

	if err != nil {
		return err
	}

	*o = UpdateBackupsRequest(varUpdateBackupsRequest)

	return err
}

type NullableUpdateBackupsRequest struct {
	value *UpdateBackupsRequest
	isSet bool
}

func (v NullableUpdateBackupsRequest) Get() *UpdateBackupsRequest {
	return v.value
}

func (v *NullableUpdateBackupsRequest) Set(val *UpdateBackupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBackupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBackupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBackupsRequest(val *UpdateBackupsRequest) *NullableUpdateBackupsRequest {
	return &NullableUpdateBackupsRequest{value: val, isSet: true}
}

func (v NullableUpdateBackupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBackupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



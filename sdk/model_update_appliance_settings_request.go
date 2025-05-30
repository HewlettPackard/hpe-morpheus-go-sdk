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
)

// checks if the UpdateApplianceSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateApplianceSettingsRequest{}

// UpdateApplianceSettingsRequest struct for UpdateApplianceSettingsRequest
type UpdateApplianceSettingsRequest struct {
	ApplianceSettings *UpdateApplianceSettingsRequestApplianceSettings `json:"applianceSettings,omitempty"`
}

// NewUpdateApplianceSettingsRequest instantiates a new UpdateApplianceSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApplianceSettingsRequest() *UpdateApplianceSettingsRequest {
	this := UpdateApplianceSettingsRequest{}
	return &this
}

// NewUpdateApplianceSettingsRequestWithDefaults instantiates a new UpdateApplianceSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApplianceSettingsRequestWithDefaults() *UpdateApplianceSettingsRequest {
	this := UpdateApplianceSettingsRequest{}
	return &this
}

// GetApplianceSettings returns the ApplianceSettings field value if set, zero value otherwise.
func (o *UpdateApplianceSettingsRequest) GetApplianceSettings() UpdateApplianceSettingsRequestApplianceSettings {
	if o == nil || IsNil(o.ApplianceSettings) {
		var ret UpdateApplianceSettingsRequestApplianceSettings
		return ret
	}
	return *o.ApplianceSettings
}

// GetApplianceSettingsOk returns a tuple with the ApplianceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApplianceSettingsRequest) GetApplianceSettingsOk() (*UpdateApplianceSettingsRequestApplianceSettings, bool) {
	if o == nil || IsNil(o.ApplianceSettings) {
		return nil, false
	}
	return o.ApplianceSettings, true
}

// IsSetApplianceSettings returns a boolean if a field has been set.
func (o *UpdateApplianceSettingsRequest) IsSetApplianceSettings() bool {
	if o != nil && !IsNil(o.ApplianceSettings) {
		return true
	}

	return false
}

// SetApplianceSettings gets a reference to the given UpdateApplianceSettingsRequestApplianceSettings and assigns it to the ApplianceSettings field.
func (o *UpdateApplianceSettingsRequest) SetApplianceSettings(v UpdateApplianceSettingsRequestApplianceSettings) {
	o.ApplianceSettings = &v
}

func (o UpdateApplianceSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateApplianceSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplianceSettings) {
		toSerialize["applianceSettings"] = o.ApplianceSettings
	}
	return toSerialize, nil
}

type NullableUpdateApplianceSettingsRequest struct {
	value *UpdateApplianceSettingsRequest
	isSet bool
}

func (v NullableUpdateApplianceSettingsRequest) Get() *UpdateApplianceSettingsRequest {
	return v.value
}

func (v *NullableUpdateApplianceSettingsRequest) Set(val *UpdateApplianceSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApplianceSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApplianceSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApplianceSettingsRequest(val *UpdateApplianceSettingsRequest) *NullableUpdateApplianceSettingsRequest {
	return &NullableUpdateApplianceSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateApplianceSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApplianceSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



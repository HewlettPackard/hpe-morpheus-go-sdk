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

// checks if the UpdateCheckApps200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCheckApps200Response{}

// UpdateCheckApps200Response struct for UpdateCheckApps200Response
type UpdateCheckApps200Response struct {
	MonitorApp *GetAlerts200ResponseAllOfAppsInner `json:"monitorApp,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewUpdateCheckApps200Response instantiates a new UpdateCheckApps200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCheckApps200Response() *UpdateCheckApps200Response {
	this := UpdateCheckApps200Response{}
	return &this
}

// NewUpdateCheckApps200ResponseWithDefaults instantiates a new UpdateCheckApps200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCheckApps200ResponseWithDefaults() *UpdateCheckApps200Response {
	this := UpdateCheckApps200Response{}
	return &this
}

// GetMonitorApp returns the MonitorApp field value if set, zero value otherwise.
func (o *UpdateCheckApps200Response) GetMonitorApp() GetAlerts200ResponseAllOfAppsInner {
	if o == nil || IsNil(o.MonitorApp) {
		var ret GetAlerts200ResponseAllOfAppsInner
		return ret
	}
	return *o.MonitorApp
}

// GetMonitorAppOk returns a tuple with the MonitorApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCheckApps200Response) GetMonitorAppOk() (*GetAlerts200ResponseAllOfAppsInner, bool) {
	if o == nil || IsNil(o.MonitorApp) {
		return nil, false
	}
	return o.MonitorApp, true
}

// IsSetMonitorApp returns a boolean if a field has been set.
func (o *UpdateCheckApps200Response) IsSetMonitorApp() bool {
	if o != nil && !IsNil(o.MonitorApp) {
		return true
	}

	return false
}

// SetMonitorApp gets a reference to the given GetAlerts200ResponseAllOfAppsInner and assigns it to the MonitorApp field.
func (o *UpdateCheckApps200Response) SetMonitorApp(v GetAlerts200ResponseAllOfAppsInner) {
	o.MonitorApp = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UpdateCheckApps200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCheckApps200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *UpdateCheckApps200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UpdateCheckApps200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o UpdateCheckApps200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCheckApps200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MonitorApp) {
		toSerialize["monitorApp"] = o.MonitorApp
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableUpdateCheckApps200Response struct {
	value *UpdateCheckApps200Response
	isSet bool
}

func (v NullableUpdateCheckApps200Response) Get() *UpdateCheckApps200Response {
	return v.value
}

func (v *NullableUpdateCheckApps200Response) Set(val *UpdateCheckApps200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCheckApps200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCheckApps200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCheckApps200Response(val *UpdateCheckApps200Response) *NullableUpdateCheckApps200Response {
	return &NullableUpdateCheckApps200Response{value: val, isSet: true}
}

func (v NullableUpdateCheckApps200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCheckApps200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



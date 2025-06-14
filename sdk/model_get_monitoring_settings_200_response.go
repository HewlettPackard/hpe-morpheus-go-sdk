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

// checks if the GetMonitoringSettings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMonitoringSettings200Response{}

// GetMonitoringSettings200Response struct for GetMonitoringSettings200Response
type GetMonitoringSettings200Response struct {
	MonitoringSettings   *GetMonitoringSettings200ResponseMonitoringSettings `json:"monitoringSettings,omitempty"`
	AdditionalProperties map[string]interface{}                              `json:",remain"`
}

type _GetMonitoringSettings200Response GetMonitoringSettings200Response

// NewGetMonitoringSettings200Response instantiates a new GetMonitoringSettings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMonitoringSettings200Response() *GetMonitoringSettings200Response {
	this := GetMonitoringSettings200Response{}
	return &this
}

// NewGetMonitoringSettings200ResponseWithDefaults instantiates a new GetMonitoringSettings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMonitoringSettings200ResponseWithDefaults() *GetMonitoringSettings200Response {
	this := GetMonitoringSettings200Response{}
	return &this
}

// GetMonitoringSettings returns the MonitoringSettings field value if set, zero value otherwise.
func (o *GetMonitoringSettings200Response) GetMonitoringSettings() GetMonitoringSettings200ResponseMonitoringSettings {
	if o == nil || IsNil(o.MonitoringSettings) {
		var ret GetMonitoringSettings200ResponseMonitoringSettings
		return ret
	}
	return *o.MonitoringSettings
}

// GetMonitoringSettingsOk returns a tuple with the MonitoringSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMonitoringSettings200Response) GetMonitoringSettingsOk() (*GetMonitoringSettings200ResponseMonitoringSettings, bool) {
	if o == nil || IsNil(o.MonitoringSettings) {
		return nil, false
	}
	return o.MonitoringSettings, true
}

// IsSetMonitoringSettings returns a boolean if a field has been set.
func (o *GetMonitoringSettings200Response) IsSetMonitoringSettings() bool {
	if o != nil && !IsNil(o.MonitoringSettings) {
		return true
	}

	return false
}

// SetMonitoringSettings gets a reference to the given GetMonitoringSettings200ResponseMonitoringSettings and assigns it to the MonitoringSettings field.
func (o *GetMonitoringSettings200Response) SetMonitoringSettings(v GetMonitoringSettings200ResponseMonitoringSettings) {
	o.MonitoringSettings = &v
}

func (o GetMonitoringSettings200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMonitoringSettings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MonitoringSettings) {
		toSerialize["monitoringSettings"] = o.MonitoringSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetMonitoringSettings200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

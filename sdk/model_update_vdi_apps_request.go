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

// checks if the UpdateVDIAppsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVDIAppsRequest{}

// UpdateVDIAppsRequest Updates VDI App
type UpdateVDIAppsRequest struct {
	VdiApp               UpdateVDIAppsRequestVdiApp `json:"vdiApp"`
	AdditionalProperties map[string]interface{}     `json:",remain"`
}

type _UpdateVDIAppsRequest UpdateVDIAppsRequest

// NewUpdateVDIAppsRequest instantiates a new UpdateVDIAppsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVDIAppsRequest(vdiApp UpdateVDIAppsRequestVdiApp) *UpdateVDIAppsRequest {
	this := UpdateVDIAppsRequest{}
	this.VdiApp = vdiApp
	return &this
}

// NewUpdateVDIAppsRequestWithDefaults instantiates a new UpdateVDIAppsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVDIAppsRequestWithDefaults() *UpdateVDIAppsRequest {
	this := UpdateVDIAppsRequest{}
	return &this
}

// GetVdiApp returns the VdiApp field value
func (o *UpdateVDIAppsRequest) GetVdiApp() UpdateVDIAppsRequestVdiApp {
	if o == nil {
		var ret UpdateVDIAppsRequestVdiApp
		return ret
	}

	return o.VdiApp
}

// GetVdiAppOk returns a tuple with the VdiApp field value
// and a boolean to check if the value has been set.
func (o *UpdateVDIAppsRequest) GetVdiAppOk() (*UpdateVDIAppsRequestVdiApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VdiApp, true
}

// SetVdiApp sets field value
func (o *UpdateVDIAppsRequest) SetVdiApp(v UpdateVDIAppsRequestVdiApp) {
	o.VdiApp = v
}

func (o UpdateVDIAppsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVDIAppsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vdiApp"] = o.VdiApp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateVDIAppsRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

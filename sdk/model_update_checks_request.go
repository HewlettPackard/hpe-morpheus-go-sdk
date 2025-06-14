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

// checks if the UpdateChecksRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateChecksRequest{}

// UpdateChecksRequest struct for UpdateChecksRequest
type UpdateChecksRequest struct {
	Check                UpdateChecksRequestCheck `json:"check"`
	AdditionalProperties map[string]interface{}   `json:",remain"`
}

type _UpdateChecksRequest UpdateChecksRequest

// NewUpdateChecksRequest instantiates a new UpdateChecksRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateChecksRequest(check UpdateChecksRequestCheck) *UpdateChecksRequest {
	this := UpdateChecksRequest{}
	this.Check = check
	return &this
}

// NewUpdateChecksRequestWithDefaults instantiates a new UpdateChecksRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateChecksRequestWithDefaults() *UpdateChecksRequest {
	this := UpdateChecksRequest{}
	return &this
}

// GetCheck returns the Check field value
func (o *UpdateChecksRequest) GetCheck() UpdateChecksRequestCheck {
	if o == nil {
		var ret UpdateChecksRequestCheck
		return ret
	}

	return o.Check
}

// GetCheckOk returns a tuple with the Check field value
// and a boolean to check if the value has been set.
func (o *UpdateChecksRequest) GetCheckOk() (*UpdateChecksRequestCheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Check, true
}

// SetCheck sets field value
func (o *UpdateChecksRequest) SetCheck(v UpdateChecksRequestCheck) {
	o.Check = v
}

func (o UpdateChecksRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateChecksRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["check"] = o.Check

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateChecksRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

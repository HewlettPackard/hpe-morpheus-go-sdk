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

// checks if the AcknowledgeHealthAlarmRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcknowledgeHealthAlarmRequest{}

// AcknowledgeHealthAlarmRequest struct for AcknowledgeHealthAlarmRequest
type AcknowledgeHealthAlarmRequest struct {
	Alarm                AcknowledgeHealthAlarmRequestAlarm `json:"alarm"`
	AdditionalProperties map[string]interface{}             `json:",remain"`
}

type _AcknowledgeHealthAlarmRequest AcknowledgeHealthAlarmRequest

// NewAcknowledgeHealthAlarmRequest instantiates a new AcknowledgeHealthAlarmRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcknowledgeHealthAlarmRequest(alarm AcknowledgeHealthAlarmRequestAlarm) *AcknowledgeHealthAlarmRequest {
	this := AcknowledgeHealthAlarmRequest{}
	this.Alarm = alarm
	return &this
}

// NewAcknowledgeHealthAlarmRequestWithDefaults instantiates a new AcknowledgeHealthAlarmRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcknowledgeHealthAlarmRequestWithDefaults() *AcknowledgeHealthAlarmRequest {
	this := AcknowledgeHealthAlarmRequest{}
	return &this
}

// GetAlarm returns the Alarm field value
func (o *AcknowledgeHealthAlarmRequest) GetAlarm() AcknowledgeHealthAlarmRequestAlarm {
	if o == nil {
		var ret AcknowledgeHealthAlarmRequestAlarm
		return ret
	}

	return o.Alarm
}

// GetAlarmOk returns a tuple with the Alarm field value
// and a boolean to check if the value has been set.
func (o *AcknowledgeHealthAlarmRequest) GetAlarmOk() (*AcknowledgeHealthAlarmRequestAlarm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alarm, true
}

// SetAlarm sets field value
func (o *AcknowledgeHealthAlarmRequest) SetAlarm(v AcknowledgeHealthAlarmRequestAlarm) {
	o.Alarm = v
}

func (o AcknowledgeHealthAlarmRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcknowledgeHealthAlarmRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alarm"] = o.Alarm

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AcknowledgeHealthAlarmRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

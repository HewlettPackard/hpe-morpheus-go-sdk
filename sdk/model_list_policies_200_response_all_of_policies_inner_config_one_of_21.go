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

// checks if the ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21{}

// ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21 - Power Schedule
type ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21 struct {
	PowerScheduleType      string                 `json:"powerScheduleType"`
	PowerSchedule          *string                `json:"powerSchedule,omitempty"`
	PowerScheduleHideFixed *bool                  `json:"powerScheduleHideFixed,omitempty"`
	AdditionalProperties   map[string]interface{} `json:",remain"`
}

type _ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21 ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21

// NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21 instantiates a new ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21(powerScheduleType string) *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21 {
	this := ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21{}
	this.PowerScheduleType = powerScheduleType
	return &this
}

// NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21WithDefaults instantiates a new ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21WithDefaults() *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21 {
	this := ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21{}
	return &this
}

// GetPowerScheduleType returns the PowerScheduleType field value
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) GetPowerScheduleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PowerScheduleType
}

// GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field value
// and a boolean to check if the value has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) GetPowerScheduleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerScheduleType, true
}

// SetPowerScheduleType sets field value
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) SetPowerScheduleType(v string) {
	o.PowerScheduleType = v
}

// GetPowerSchedule returns the PowerSchedule field value if set, zero value otherwise.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) GetPowerSchedule() string {
	if o == nil || IsNil(o.PowerSchedule) {
		var ret string
		return ret
	}
	return *o.PowerSchedule
}

// GetPowerScheduleOk returns a tuple with the PowerSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) GetPowerScheduleOk() (*string, bool) {
	if o == nil || IsNil(o.PowerSchedule) {
		return nil, false
	}
	return o.PowerSchedule, true
}

// IsSetPowerSchedule returns a boolean if a field has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) IsSetPowerSchedule() bool {
	if o != nil && !IsNil(o.PowerSchedule) {
		return true
	}

	return false
}

// SetPowerSchedule gets a reference to the given string and assigns it to the PowerSchedule field.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) SetPowerSchedule(v string) {
	o.PowerSchedule = &v
}

// GetPowerScheduleHideFixed returns the PowerScheduleHideFixed field value if set, zero value otherwise.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) GetPowerScheduleHideFixed() bool {
	if o == nil || IsNil(o.PowerScheduleHideFixed) {
		var ret bool
		return ret
	}
	return *o.PowerScheduleHideFixed
}

// GetPowerScheduleHideFixedOk returns a tuple with the PowerScheduleHideFixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) GetPowerScheduleHideFixedOk() (*bool, bool) {
	if o == nil || IsNil(o.PowerScheduleHideFixed) {
		return nil, false
	}
	return o.PowerScheduleHideFixed, true
}

// IsSetPowerScheduleHideFixed returns a boolean if a field has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) IsSetPowerScheduleHideFixed() bool {
	if o != nil && !IsNil(o.PowerScheduleHideFixed) {
		return true
	}

	return false
}

// SetPowerScheduleHideFixed gets a reference to the given bool and assigns it to the PowerScheduleHideFixed field.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) SetPowerScheduleHideFixed(v bool) {
	o.PowerScheduleHideFixed = &v
}

func (o ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["powerScheduleType"] = o.PowerScheduleType
	if !IsNil(o.PowerSchedule) {
		toSerialize["powerSchedule"] = o.PowerSchedule
	}
	if !IsNil(o.PowerScheduleHideFixed) {
		toSerialize["powerScheduleHideFixed"] = o.PowerScheduleHideFixed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf21) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

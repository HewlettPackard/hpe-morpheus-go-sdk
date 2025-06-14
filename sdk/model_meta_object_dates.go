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
	"time"
)

// checks if the MetaObjectDates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaObjectDates{}

// MetaObjectDates struct for MetaObjectDates
type MetaObjectDates struct {
	StartDate            *time.Time             `json:"startDate,omitempty"`
	EndDate              *time.Time             `json:"endDate,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _MetaObjectDates MetaObjectDates

// NewMetaObjectDates instantiates a new MetaObjectDates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaObjectDates() *MetaObjectDates {
	this := MetaObjectDates{}
	return &this
}

// NewMetaObjectDatesWithDefaults instantiates a new MetaObjectDates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaObjectDatesWithDefaults() *MetaObjectDates {
	this := MetaObjectDates{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *MetaObjectDates) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaObjectDates) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// IsSetStartDate returns a boolean if a field has been set.
func (o *MetaObjectDates) IsSetStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *MetaObjectDates) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *MetaObjectDates) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaObjectDates) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// IsSetEndDate returns a boolean if a field has been set.
func (o *MetaObjectDates) IsSetEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *MetaObjectDates) SetEndDate(v time.Time) {
	o.EndDate = &v
}

func (o MetaObjectDates) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaObjectDates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *MetaObjectDates) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

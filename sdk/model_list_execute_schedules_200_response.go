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

// checks if the ListExecuteSchedules200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListExecuteSchedules200Response{}

// ListExecuteSchedules200Response struct for ListExecuteSchedules200Response
type ListExecuteSchedules200Response struct {
	Schedules            []ListExecuteSchedules200ResponseAllOfSchedulesInner `json:"schedules,omitempty"`
	Meta                 *ListActivity200ResponseAllOfMeta                    `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}                               `json:",remain"`
}

type _ListExecuteSchedules200Response ListExecuteSchedules200Response

// NewListExecuteSchedules200Response instantiates a new ListExecuteSchedules200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListExecuteSchedules200Response() *ListExecuteSchedules200Response {
	this := ListExecuteSchedules200Response{}
	return &this
}

// NewListExecuteSchedules200ResponseWithDefaults instantiates a new ListExecuteSchedules200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListExecuteSchedules200ResponseWithDefaults() *ListExecuteSchedules200Response {
	this := ListExecuteSchedules200Response{}
	return &this
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *ListExecuteSchedules200Response) GetSchedules() []ListExecuteSchedules200ResponseAllOfSchedulesInner {
	if o == nil || IsNil(o.Schedules) {
		var ret []ListExecuteSchedules200ResponseAllOfSchedulesInner
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListExecuteSchedules200Response) GetSchedulesOk() ([]ListExecuteSchedules200ResponseAllOfSchedulesInner, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// IsSetSchedules returns a boolean if a field has been set.
func (o *ListExecuteSchedules200Response) IsSetSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []ListExecuteSchedules200ResponseAllOfSchedulesInner and assigns it to the Schedules field.
func (o *ListExecuteSchedules200Response) SetSchedules(v []ListExecuteSchedules200ResponseAllOfSchedulesInner) {
	o.Schedules = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListExecuteSchedules200Response) GetMeta() ListActivity200ResponseAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListActivity200ResponseAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListExecuteSchedules200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// IsSetMeta returns a boolean if a field has been set.
func (o *ListExecuteSchedules200Response) IsSetMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListActivity200ResponseAllOfMeta and assigns it to the Meta field.
func (o *ListExecuteSchedules200Response) SetMeta(v ListActivity200ResponseAllOfMeta) {
	o.Meta = &v
}

func (o ListExecuteSchedules200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListExecuteSchedules200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schedules) {
		toSerialize["schedules"] = o.Schedules
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListExecuteSchedules200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

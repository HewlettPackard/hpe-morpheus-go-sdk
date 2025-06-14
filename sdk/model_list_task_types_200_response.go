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

// checks if the ListTaskTypes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTaskTypes200Response{}

// ListTaskTypes200Response struct for ListTaskTypes200Response
type ListTaskTypes200Response struct {
	TaskTypes            []ListTaskTypes200ResponseTaskTypesInner `json:"taskTypes,omitempty"`
	AdditionalProperties map[string]interface{}                   `json:",remain"`
}

type _ListTaskTypes200Response ListTaskTypes200Response

// NewListTaskTypes200Response instantiates a new ListTaskTypes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTaskTypes200Response() *ListTaskTypes200Response {
	this := ListTaskTypes200Response{}
	return &this
}

// NewListTaskTypes200ResponseWithDefaults instantiates a new ListTaskTypes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTaskTypes200ResponseWithDefaults() *ListTaskTypes200Response {
	this := ListTaskTypes200Response{}
	return &this
}

// GetTaskTypes returns the TaskTypes field value if set, zero value otherwise.
func (o *ListTaskTypes200Response) GetTaskTypes() []ListTaskTypes200ResponseTaskTypesInner {
	if o == nil || IsNil(o.TaskTypes) {
		var ret []ListTaskTypes200ResponseTaskTypesInner
		return ret
	}
	return o.TaskTypes
}

// GetTaskTypesOk returns a tuple with the TaskTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200Response) GetTaskTypesOk() ([]ListTaskTypes200ResponseTaskTypesInner, bool) {
	if o == nil || IsNil(o.TaskTypes) {
		return nil, false
	}
	return o.TaskTypes, true
}

// IsSetTaskTypes returns a boolean if a field has been set.
func (o *ListTaskTypes200Response) IsSetTaskTypes() bool {
	if o != nil && !IsNil(o.TaskTypes) {
		return true
	}

	return false
}

// SetTaskTypes gets a reference to the given []ListTaskTypes200ResponseTaskTypesInner and assigns it to the TaskTypes field.
func (o *ListTaskTypes200Response) SetTaskTypes(v []ListTaskTypes200ResponseTaskTypesInner) {
	o.TaskTypes = v
}

func (o ListTaskTypes200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTaskTypes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaskTypes) {
		toSerialize["taskTypes"] = o.TaskTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListTaskTypes200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the ExecutionId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionId{}

// ExecutionId struct for ExecutionId
type ExecutionId struct {
	ExecutionId          *string                `json:"executionId,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ExecutionId ExecutionId

// NewExecutionId instantiates a new ExecutionId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionId() *ExecutionId {
	this := ExecutionId{}
	return &this
}

// NewExecutionIdWithDefaults instantiates a new ExecutionId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionIdWithDefaults() *ExecutionId {
	this := ExecutionId{}
	return &this
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *ExecutionId) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId) {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionId) GetExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// IsSetExecutionId returns a boolean if a field has been set.
func (o *ExecutionId) IsSetExecutionId() bool {
	if o != nil && !IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *ExecutionId) SetExecutionId(v string) {
	o.ExecutionId = &v
}

func (o ExecutionId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExecutionId) {
		toSerialize["executionId"] = o.ExecutionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ExecutionId) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

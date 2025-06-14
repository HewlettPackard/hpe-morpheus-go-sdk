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

// checks if the WorkflowJobPayloadTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowJobPayloadTask{}

// WorkflowJobPayloadTask struct for WorkflowJobPayloadTask
type WorkflowJobPayloadTask struct {
	Id                   *int64                 `json:"id,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _WorkflowJobPayloadTask WorkflowJobPayloadTask

// NewWorkflowJobPayloadTask instantiates a new WorkflowJobPayloadTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowJobPayloadTask() *WorkflowJobPayloadTask {
	this := WorkflowJobPayloadTask{}
	return &this
}

// NewWorkflowJobPayloadTaskWithDefaults instantiates a new WorkflowJobPayloadTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowJobPayloadTaskWithDefaults() *WorkflowJobPayloadTask {
	this := WorkflowJobPayloadTask{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowJobPayloadTask) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowJobPayloadTask) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *WorkflowJobPayloadTask) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *WorkflowJobPayloadTask) SetId(v int64) {
	o.Id = &v
}

func (o WorkflowJobPayloadTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowJobPayloadTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *WorkflowJobPayloadTask) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType{}

// ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType struct for ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType
type ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType struct {
	Id                   *int64                 `json:"id,omitempty"`
	Code                 *string                `json:"code,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType

// NewListTasks200ResponseAllOfTasksInnerAnyOf12TaskType instantiates a new ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTasks200ResponseAllOfTasksInnerAnyOf12TaskType() *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType {
	this := ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType{}
	return &this
}

// NewListTasks200ResponseAllOfTasksInnerAnyOf12TaskTypeWithDefaults instantiates a new ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTasks200ResponseAllOfTasksInnerAnyOf12TaskTypeWithDefaults() *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType {
	this := ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) SetName(v string) {
	o.Name = &v
}

func (o ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListTasks200ResponseAllOfTasksInnerAnyOf12TaskType) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

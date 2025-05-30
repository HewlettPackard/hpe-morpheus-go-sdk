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

// checks if the AddJobs200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddJobs200Response{}

// AddJobs200Response struct for AddJobs200Response
type AddJobs200Response struct {
	Job *ListJobs200ResponseAllOfJobsInnerAnyOf `json:"job,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewAddJobs200Response instantiates a new AddJobs200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddJobs200Response() *AddJobs200Response {
	this := AddJobs200Response{}
	return &this
}

// NewAddJobs200ResponseWithDefaults instantiates a new AddJobs200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddJobs200ResponseWithDefaults() *AddJobs200Response {
	this := AddJobs200Response{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *AddJobs200Response) GetJob() ListJobs200ResponseAllOfJobsInnerAnyOf {
	if o == nil || IsNil(o.Job) {
		var ret ListJobs200ResponseAllOfJobsInnerAnyOf
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJobs200Response) GetJobOk() (*ListJobs200ResponseAllOfJobsInnerAnyOf, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// IsSetJob returns a boolean if a field has been set.
func (o *AddJobs200Response) IsSetJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given ListJobs200ResponseAllOfJobsInnerAnyOf and assigns it to the Job field.
func (o *AddJobs200Response) SetJob(v ListJobs200ResponseAllOfJobsInnerAnyOf) {
	o.Job = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AddJobs200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJobs200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *AddJobs200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AddJobs200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o AddJobs200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddJobs200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableAddJobs200Response struct {
	value *AddJobs200Response
	isSet bool
}

func (v NullableAddJobs200Response) Get() *AddJobs200Response {
	return v.value
}

func (v *NullableAddJobs200Response) Set(val *AddJobs200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddJobs200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddJobs200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddJobs200Response(val *AddJobs200Response) *NullableAddJobs200Response {
	return &NullableAddJobs200Response{value: val, isSet: true}
}

func (v NullableAddJobs200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddJobs200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



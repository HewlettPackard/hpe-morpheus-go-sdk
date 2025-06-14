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

// checks if the AddExecuteSchedules200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddExecuteSchedules200Response{}

// AddExecuteSchedules200Response struct for AddExecuteSchedules200Response
type AddExecuteSchedules200Response struct {
	Schedule             *ListExecuteSchedules200ResponseAllOfSchedulesInner `json:"schedule,omitempty"`
	Success              *bool                                               `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}                              `json:",remain"`
}

type _AddExecuteSchedules200Response AddExecuteSchedules200Response

// NewAddExecuteSchedules200Response instantiates a new AddExecuteSchedules200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddExecuteSchedules200Response() *AddExecuteSchedules200Response {
	this := AddExecuteSchedules200Response{}
	return &this
}

// NewAddExecuteSchedules200ResponseWithDefaults instantiates a new AddExecuteSchedules200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddExecuteSchedules200ResponseWithDefaults() *AddExecuteSchedules200Response {
	this := AddExecuteSchedules200Response{}
	return &this
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *AddExecuteSchedules200Response) GetSchedule() ListExecuteSchedules200ResponseAllOfSchedulesInner {
	if o == nil || IsNil(o.Schedule) {
		var ret ListExecuteSchedules200ResponseAllOfSchedulesInner
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddExecuteSchedules200Response) GetScheduleOk() (*ListExecuteSchedules200ResponseAllOfSchedulesInner, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// IsSetSchedule returns a boolean if a field has been set.
func (o *AddExecuteSchedules200Response) IsSetSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ListExecuteSchedules200ResponseAllOfSchedulesInner and assigns it to the Schedule field.
func (o *AddExecuteSchedules200Response) SetSchedule(v ListExecuteSchedules200ResponseAllOfSchedulesInner) {
	o.Schedule = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AddExecuteSchedules200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddExecuteSchedules200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *AddExecuteSchedules200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AddExecuteSchedules200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o AddExecuteSchedules200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddExecuteSchedules200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddExecuteSchedules200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

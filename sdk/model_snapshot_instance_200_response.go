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

// checks if the SnapshotInstance200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotInstance200Response{}

// SnapshotInstance200Response struct for SnapshotInstance200Response
type SnapshotInstance200Response struct {
	Success *bool `json:"success,omitempty"`
	// Process ID(s) to track execution results with, use `/api/processes/$processId`. Instances with more than one server will result in multiple processes, one for each snapshot.
	ProcessIds []int64 `json:"processIds,omitempty"`
}

// NewSnapshotInstance200Response instantiates a new SnapshotInstance200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotInstance200Response() *SnapshotInstance200Response {
	this := SnapshotInstance200Response{}
	return &this
}

// NewSnapshotInstance200ResponseWithDefaults instantiates a new SnapshotInstance200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotInstance200ResponseWithDefaults() *SnapshotInstance200Response {
	this := SnapshotInstance200Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SnapshotInstance200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotInstance200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *SnapshotInstance200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SnapshotInstance200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetProcessIds returns the ProcessIds field value if set, zero value otherwise.
func (o *SnapshotInstance200Response) GetProcessIds() []int64 {
	if o == nil || IsNil(o.ProcessIds) {
		var ret []int64
		return ret
	}
	return o.ProcessIds
}

// GetProcessIdsOk returns a tuple with the ProcessIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotInstance200Response) GetProcessIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ProcessIds) {
		return nil, false
	}
	return o.ProcessIds, true
}

// IsSetProcessIds returns a boolean if a field has been set.
func (o *SnapshotInstance200Response) IsSetProcessIds() bool {
	if o != nil && !IsNil(o.ProcessIds) {
		return true
	}

	return false
}

// SetProcessIds gets a reference to the given []int64 and assigns it to the ProcessIds field.
func (o *SnapshotInstance200Response) SetProcessIds(v []int64) {
	o.ProcessIds = v
}

func (o SnapshotInstance200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotInstance200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.ProcessIds) {
		toSerialize["processIds"] = o.ProcessIds
	}
	return toSerialize, nil
}

type NullableSnapshotInstance200Response struct {
	value *SnapshotInstance200Response
	isSet bool
}

func (v NullableSnapshotInstance200Response) Get() *SnapshotInstance200Response {
	return v.value
}

func (v *NullableSnapshotInstance200Response) Set(val *SnapshotInstance200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotInstance200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotInstance200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotInstance200Response(val *SnapshotInstance200Response) *NullableSnapshotInstance200Response {
	return &NullableSnapshotInstance200Response{value: val, isSet: true}
}

func (v NullableSnapshotInstance200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotInstance200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



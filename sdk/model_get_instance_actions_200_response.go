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

// checks if the GetInstanceActions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInstanceActions200Response{}

// GetInstanceActions200Response struct for GetInstanceActions200Response
type GetInstanceActions200Response struct {
	InstanceIds []int64 `json:"instanceIds,omitempty"`
	Actions []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner `json:"actions,omitempty"`
}

// NewGetInstanceActions200Response instantiates a new GetInstanceActions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInstanceActions200Response() *GetInstanceActions200Response {
	this := GetInstanceActions200Response{}
	return &this
}

// NewGetInstanceActions200ResponseWithDefaults instantiates a new GetInstanceActions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInstanceActions200ResponseWithDefaults() *GetInstanceActions200Response {
	this := GetInstanceActions200Response{}
	return &this
}

// GetInstanceIds returns the InstanceIds field value if set, zero value otherwise.
func (o *GetInstanceActions200Response) GetInstanceIds() []int64 {
	if o == nil || IsNil(o.InstanceIds) {
		var ret []int64
		return ret
	}
	return o.InstanceIds
}

// GetInstanceIdsOk returns a tuple with the InstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceActions200Response) GetInstanceIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.InstanceIds) {
		return nil, false
	}
	return o.InstanceIds, true
}

// IsSetInstanceIds returns a boolean if a field has been set.
func (o *GetInstanceActions200Response) IsSetInstanceIds() bool {
	if o != nil && !IsNil(o.InstanceIds) {
		return true
	}

	return false
}

// SetInstanceIds gets a reference to the given []int64 and assigns it to the InstanceIds field.
func (o *GetInstanceActions200Response) SetInstanceIds(v []int64) {
	o.InstanceIds = v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *GetInstanceActions200Response) GetActions() []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner {
	if o == nil || IsNil(o.Actions) {
		var ret []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceActions200Response) GetActionsOk() ([]ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// IsSetActions returns a boolean if a field has been set.
func (o *GetInstanceActions200Response) IsSetActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner and assigns it to the Actions field.
func (o *GetInstanceActions200Response) SetActions(v []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner) {
	o.Actions = v
}

func (o GetInstanceActions200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInstanceActions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceIds) {
		toSerialize["instanceIds"] = o.InstanceIds
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	return toSerialize, nil
}

type NullableGetInstanceActions200Response struct {
	value *GetInstanceActions200Response
	isSet bool
}

func (v NullableGetInstanceActions200Response) Get() *GetInstanceActions200Response {
	return v.value
}

func (v *NullableGetInstanceActions200Response) Set(val *GetInstanceActions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInstanceActions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInstanceActions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInstanceActions200Response(val *GetInstanceActions200Response) *NullableGetInstanceActions200Response {
	return &NullableGetInstanceActions200Response{value: val, isSet: true}
}

func (v NullableGetInstanceActions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInstanceActions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



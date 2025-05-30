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

// checks if the GetNetworkPoolIps401Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkPoolIps401Response{}

// GetNetworkPoolIps401Response struct for GetNetworkPoolIps401Response
type GetNetworkPoolIps401Response struct {
	Msg *string `json:"msg,omitempty"`
}

// NewGetNetworkPoolIps401Response instantiates a new GetNetworkPoolIps401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkPoolIps401Response() *GetNetworkPoolIps401Response {
	this := GetNetworkPoolIps401Response{}
	return &this
}

// NewGetNetworkPoolIps401ResponseWithDefaults instantiates a new GetNetworkPoolIps401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkPoolIps401ResponseWithDefaults() *GetNetworkPoolIps401Response {
	this := GetNetworkPoolIps401Response{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GetNetworkPoolIps401Response) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkPoolIps401Response) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// IsSetMsg returns a boolean if a field has been set.
func (o *GetNetworkPoolIps401Response) IsSetMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GetNetworkPoolIps401Response) SetMsg(v string) {
	o.Msg = &v
}

func (o GetNetworkPoolIps401Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkPoolIps401Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullableGetNetworkPoolIps401Response struct {
	value *GetNetworkPoolIps401Response
	isSet bool
}

func (v NullableGetNetworkPoolIps401Response) Get() *GetNetworkPoolIps401Response {
	return v.value
}

func (v *NullableGetNetworkPoolIps401Response) Set(val *GetNetworkPoolIps401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkPoolIps401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkPoolIps401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkPoolIps401Response(val *GetNetworkPoolIps401Response) *NullableGetNetworkPoolIps401Response {
	return &NullableGetNetworkPoolIps401Response{value: val, isSet: true}
}

func (v NullableGetNetworkPoolIps401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkPoolIps401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



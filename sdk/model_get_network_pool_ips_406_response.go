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

// checks if the GetNetworkPoolIps406Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkPoolIps406Response{}

// GetNetworkPoolIps406Response struct for GetNetworkPoolIps406Response
type GetNetworkPoolIps406Response struct {
	Msg                  *string                `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _GetNetworkPoolIps406Response GetNetworkPoolIps406Response

// NewGetNetworkPoolIps406Response instantiates a new GetNetworkPoolIps406Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkPoolIps406Response() *GetNetworkPoolIps406Response {
	this := GetNetworkPoolIps406Response{}
	return &this
}

// NewGetNetworkPoolIps406ResponseWithDefaults instantiates a new GetNetworkPoolIps406Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkPoolIps406ResponseWithDefaults() *GetNetworkPoolIps406Response {
	this := GetNetworkPoolIps406Response{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GetNetworkPoolIps406Response) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkPoolIps406Response) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// IsSetMsg returns a boolean if a field has been set.
func (o *GetNetworkPoolIps406Response) IsSetMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GetNetworkPoolIps406Response) SetMsg(v string) {
	o.Msg = &v
}

func (o GetNetworkPoolIps406Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkPoolIps406Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetNetworkPoolIps406Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the GetNetworkDhcpRelay200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkDhcpRelay200Response{}

// GetNetworkDhcpRelay200Response struct for GetNetworkDhcpRelay200Response
type GetNetworkDhcpRelay200Response struct {
	NetworkDhcpRelay *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay `json:"networkDhcpRelay,omitempty"`
}

// NewGetNetworkDhcpRelay200Response instantiates a new GetNetworkDhcpRelay200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkDhcpRelay200Response() *GetNetworkDhcpRelay200Response {
	this := GetNetworkDhcpRelay200Response{}
	return &this
}

// NewGetNetworkDhcpRelay200ResponseWithDefaults instantiates a new GetNetworkDhcpRelay200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkDhcpRelay200ResponseWithDefaults() *GetNetworkDhcpRelay200Response {
	this := GetNetworkDhcpRelay200Response{}
	return &this
}

// GetNetworkDhcpRelay returns the NetworkDhcpRelay field value if set, zero value otherwise.
func (o *GetNetworkDhcpRelay200Response) GetNetworkDhcpRelay() GetNetworkDhcpRelay200ResponseNetworkDhcpRelay {
	if o == nil || IsNil(o.NetworkDhcpRelay) {
		var ret GetNetworkDhcpRelay200ResponseNetworkDhcpRelay
		return ret
	}
	return *o.NetworkDhcpRelay
}

// GetNetworkDhcpRelayOk returns a tuple with the NetworkDhcpRelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpRelay200Response) GetNetworkDhcpRelayOk() (*GetNetworkDhcpRelay200ResponseNetworkDhcpRelay, bool) {
	if o == nil || IsNil(o.NetworkDhcpRelay) {
		return nil, false
	}
	return o.NetworkDhcpRelay, true
}

// IsSetNetworkDhcpRelay returns a boolean if a field has been set.
func (o *GetNetworkDhcpRelay200Response) IsSetNetworkDhcpRelay() bool {
	if o != nil && !IsNil(o.NetworkDhcpRelay) {
		return true
	}

	return false
}

// SetNetworkDhcpRelay gets a reference to the given GetNetworkDhcpRelay200ResponseNetworkDhcpRelay and assigns it to the NetworkDhcpRelay field.
func (o *GetNetworkDhcpRelay200Response) SetNetworkDhcpRelay(v GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) {
	o.NetworkDhcpRelay = &v
}

func (o GetNetworkDhcpRelay200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkDhcpRelay200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkDhcpRelay) {
		toSerialize["networkDhcpRelay"] = o.NetworkDhcpRelay
	}
	return toSerialize, nil
}

type NullableGetNetworkDhcpRelay200Response struct {
	value *GetNetworkDhcpRelay200Response
	isSet bool
}

func (v NullableGetNetworkDhcpRelay200Response) Get() *GetNetworkDhcpRelay200Response {
	return v.value
}

func (v *NullableGetNetworkDhcpRelay200Response) Set(val *GetNetworkDhcpRelay200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkDhcpRelay200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkDhcpRelay200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkDhcpRelay200Response(val *GetNetworkDhcpRelay200Response) *NullableGetNetworkDhcpRelay200Response {
	return &NullableGetNetworkDhcpRelay200Response{value: val, isSet: true}
}

func (v NullableGetNetworkDhcpRelay200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkDhcpRelay200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the GetNetworkRoutersBgpNeighbors200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkRoutersBgpNeighbors200Response{}

// GetNetworkRoutersBgpNeighbors200Response struct for GetNetworkRoutersBgpNeighbors200Response
type GetNetworkRoutersBgpNeighbors200Response struct {
	NetworkRouterBgpNeighbors []GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner `json:"networkRouterBgpNeighbors,omitempty"`
}

// NewGetNetworkRoutersBgpNeighbors200Response instantiates a new GetNetworkRoutersBgpNeighbors200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkRoutersBgpNeighbors200Response() *GetNetworkRoutersBgpNeighbors200Response {
	this := GetNetworkRoutersBgpNeighbors200Response{}
	return &this
}

// NewGetNetworkRoutersBgpNeighbors200ResponseWithDefaults instantiates a new GetNetworkRoutersBgpNeighbors200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkRoutersBgpNeighbors200ResponseWithDefaults() *GetNetworkRoutersBgpNeighbors200Response {
	this := GetNetworkRoutersBgpNeighbors200Response{}
	return &this
}

// GetNetworkRouterBgpNeighbors returns the NetworkRouterBgpNeighbors field value if set, zero value otherwise.
func (o *GetNetworkRoutersBgpNeighbors200Response) GetNetworkRouterBgpNeighbors() []GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner {
	if o == nil || IsNil(o.NetworkRouterBgpNeighbors) {
		var ret []GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner
		return ret
	}
	return o.NetworkRouterBgpNeighbors
}

// GetNetworkRouterBgpNeighborsOk returns a tuple with the NetworkRouterBgpNeighbors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRoutersBgpNeighbors200Response) GetNetworkRouterBgpNeighborsOk() ([]GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner, bool) {
	if o == nil || IsNil(o.NetworkRouterBgpNeighbors) {
		return nil, false
	}
	return o.NetworkRouterBgpNeighbors, true
}

// IsSetNetworkRouterBgpNeighbors returns a boolean if a field has been set.
func (o *GetNetworkRoutersBgpNeighbors200Response) IsSetNetworkRouterBgpNeighbors() bool {
	if o != nil && !IsNil(o.NetworkRouterBgpNeighbors) {
		return true
	}

	return false
}

// SetNetworkRouterBgpNeighbors gets a reference to the given []GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner and assigns it to the NetworkRouterBgpNeighbors field.
func (o *GetNetworkRoutersBgpNeighbors200Response) SetNetworkRouterBgpNeighbors(v []GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) {
	o.NetworkRouterBgpNeighbors = v
}

func (o GetNetworkRoutersBgpNeighbors200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkRoutersBgpNeighbors200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkRouterBgpNeighbors) {
		toSerialize["networkRouterBgpNeighbors"] = o.NetworkRouterBgpNeighbors
	}
	return toSerialize, nil
}

type NullableGetNetworkRoutersBgpNeighbors200Response struct {
	value *GetNetworkRoutersBgpNeighbors200Response
	isSet bool
}

func (v NullableGetNetworkRoutersBgpNeighbors200Response) Get() *GetNetworkRoutersBgpNeighbors200Response {
	return v.value
}

func (v *NullableGetNetworkRoutersBgpNeighbors200Response) Set(val *GetNetworkRoutersBgpNeighbors200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkRoutersBgpNeighbors200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkRoutersBgpNeighbors200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkRoutersBgpNeighbors200Response(val *GetNetworkRoutersBgpNeighbors200Response) *NullableGetNetworkRoutersBgpNeighbors200Response {
	return &NullableGetNetworkRoutersBgpNeighbors200Response{value: val, isSet: true}
}

func (v NullableGetNetworkRoutersBgpNeighbors200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkRoutersBgpNeighbors200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



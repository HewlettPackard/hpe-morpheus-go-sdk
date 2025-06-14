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

// checks if the GetNetworkGroups200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkGroups200Response{}

// GetNetworkGroups200Response struct for GetNetworkGroups200Response
type GetNetworkGroups200Response struct {
	NetworkGroups        []GetNetworkGroups200ResponseNetworkGroupsInner `json:"networkGroups,omitempty"`
	AdditionalProperties map[string]interface{}                          `json:",remain"`
}

type _GetNetworkGroups200Response GetNetworkGroups200Response

// NewGetNetworkGroups200Response instantiates a new GetNetworkGroups200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkGroups200Response() *GetNetworkGroups200Response {
	this := GetNetworkGroups200Response{}
	return &this
}

// NewGetNetworkGroups200ResponseWithDefaults instantiates a new GetNetworkGroups200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkGroups200ResponseWithDefaults() *GetNetworkGroups200Response {
	this := GetNetworkGroups200Response{}
	return &this
}

// GetNetworkGroups returns the NetworkGroups field value if set, zero value otherwise.
func (o *GetNetworkGroups200Response) GetNetworkGroups() []GetNetworkGroups200ResponseNetworkGroupsInner {
	if o == nil || IsNil(o.NetworkGroups) {
		var ret []GetNetworkGroups200ResponseNetworkGroupsInner
		return ret
	}
	return o.NetworkGroups
}

// GetNetworkGroupsOk returns a tuple with the NetworkGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroups200Response) GetNetworkGroupsOk() ([]GetNetworkGroups200ResponseNetworkGroupsInner, bool) {
	if o == nil || IsNil(o.NetworkGroups) {
		return nil, false
	}
	return o.NetworkGroups, true
}

// IsSetNetworkGroups returns a boolean if a field has been set.
func (o *GetNetworkGroups200Response) IsSetNetworkGroups() bool {
	if o != nil && !IsNil(o.NetworkGroups) {
		return true
	}

	return false
}

// SetNetworkGroups gets a reference to the given []GetNetworkGroups200ResponseNetworkGroupsInner and assigns it to the NetworkGroups field.
func (o *GetNetworkGroups200Response) SetNetworkGroups(v []GetNetworkGroups200ResponseNetworkGroupsInner) {
	o.NetworkGroups = v
}

func (o GetNetworkGroups200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkGroups200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkGroups) {
		toSerialize["networkGroups"] = o.NetworkGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetNetworkGroups200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

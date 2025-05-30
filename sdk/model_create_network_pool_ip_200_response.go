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

// checks if the CreateNetworkPoolIp200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkPoolIp200Response{}

// CreateNetworkPoolIp200Response struct for CreateNetworkPoolIp200Response
type CreateNetworkPoolIp200Response struct {
	NetworkPool *CreateNetworkPoolIp200ResponseNetworkPool `json:"networkPool,omitempty"`
}

// NewCreateNetworkPoolIp200Response instantiates a new CreateNetworkPoolIp200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkPoolIp200Response() *CreateNetworkPoolIp200Response {
	this := CreateNetworkPoolIp200Response{}
	return &this
}

// NewCreateNetworkPoolIp200ResponseWithDefaults instantiates a new CreateNetworkPoolIp200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkPoolIp200ResponseWithDefaults() *CreateNetworkPoolIp200Response {
	this := CreateNetworkPoolIp200Response{}
	return &this
}

// GetNetworkPool returns the NetworkPool field value if set, zero value otherwise.
func (o *CreateNetworkPoolIp200Response) GetNetworkPool() CreateNetworkPoolIp200ResponseNetworkPool {
	if o == nil || IsNil(o.NetworkPool) {
		var ret CreateNetworkPoolIp200ResponseNetworkPool
		return ret
	}
	return *o.NetworkPool
}

// GetNetworkPoolOk returns a tuple with the NetworkPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkPoolIp200Response) GetNetworkPoolOk() (*CreateNetworkPoolIp200ResponseNetworkPool, bool) {
	if o == nil || IsNil(o.NetworkPool) {
		return nil, false
	}
	return o.NetworkPool, true
}

// IsSetNetworkPool returns a boolean if a field has been set.
func (o *CreateNetworkPoolIp200Response) IsSetNetworkPool() bool {
	if o != nil && !IsNil(o.NetworkPool) {
		return true
	}

	return false
}

// SetNetworkPool gets a reference to the given CreateNetworkPoolIp200ResponseNetworkPool and assigns it to the NetworkPool field.
func (o *CreateNetworkPoolIp200Response) SetNetworkPool(v CreateNetworkPoolIp200ResponseNetworkPool) {
	o.NetworkPool = &v
}

func (o CreateNetworkPoolIp200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkPoolIp200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkPool) {
		toSerialize["networkPool"] = o.NetworkPool
	}
	return toSerialize, nil
}

type NullableCreateNetworkPoolIp200Response struct {
	value *CreateNetworkPoolIp200Response
	isSet bool
}

func (v NullableCreateNetworkPoolIp200Response) Get() *CreateNetworkPoolIp200Response {
	return v.value
}

func (v *NullableCreateNetworkPoolIp200Response) Set(val *CreateNetworkPoolIp200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkPoolIp200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkPoolIp200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkPoolIp200Response(val *CreateNetworkPoolIp200Response) *NullableCreateNetworkPoolIp200Response {
	return &NullableCreateNetworkPoolIp200Response{value: val, isSet: true}
}

func (v NullableCreateNetworkPoolIp200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkPoolIp200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the AddIntegrations200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddIntegrations200Response{}

// AddIntegrations200Response struct for AddIntegrations200Response
type AddIntegrations200Response struct {
	Integration *AddIntegrations200ResponseAllOfIntegration `json:"integration,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewAddIntegrations200Response instantiates a new AddIntegrations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIntegrations200Response() *AddIntegrations200Response {
	this := AddIntegrations200Response{}
	return &this
}

// NewAddIntegrations200ResponseWithDefaults instantiates a new AddIntegrations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIntegrations200ResponseWithDefaults() *AddIntegrations200Response {
	this := AddIntegrations200Response{}
	return &this
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *AddIntegrations200Response) GetIntegration() AddIntegrations200ResponseAllOfIntegration {
	if o == nil || IsNil(o.Integration) {
		var ret AddIntegrations200ResponseAllOfIntegration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIntegrations200Response) GetIntegrationOk() (*AddIntegrations200ResponseAllOfIntegration, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// IsSetIntegration returns a boolean if a field has been set.
func (o *AddIntegrations200Response) IsSetIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given AddIntegrations200ResponseAllOfIntegration and assigns it to the Integration field.
func (o *AddIntegrations200Response) SetIntegration(v AddIntegrations200ResponseAllOfIntegration) {
	o.Integration = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AddIntegrations200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIntegrations200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *AddIntegrations200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AddIntegrations200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o AddIntegrations200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddIntegrations200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableAddIntegrations200Response struct {
	value *AddIntegrations200Response
	isSet bool
}

func (v NullableAddIntegrations200Response) Get() *AddIntegrations200Response {
	return v.value
}

func (v *NullableAddIntegrations200Response) Set(val *AddIntegrations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddIntegrations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddIntegrations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddIntegrations200Response(val *AddIntegrations200Response) *NullableAddIntegrations200Response {
	return &NullableAddIntegrations200Response{value: val, isSet: true}
}

func (v NullableAddIntegrations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddIntegrations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the UpdateIdentitySourceSubdomains200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIdentitySourceSubdomains200Response{}

// UpdateIdentitySourceSubdomains200Response struct for UpdateIdentitySourceSubdomains200Response
type UpdateIdentitySourceSubdomains200Response struct {
	UserSource *ListIdentitySources200ResponseAllOfUserSourcesInner `json:"userSource,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewUpdateIdentitySourceSubdomains200Response instantiates a new UpdateIdentitySourceSubdomains200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIdentitySourceSubdomains200Response() *UpdateIdentitySourceSubdomains200Response {
	this := UpdateIdentitySourceSubdomains200Response{}
	return &this
}

// NewUpdateIdentitySourceSubdomains200ResponseWithDefaults instantiates a new UpdateIdentitySourceSubdomains200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIdentitySourceSubdomains200ResponseWithDefaults() *UpdateIdentitySourceSubdomains200Response {
	this := UpdateIdentitySourceSubdomains200Response{}
	return &this
}

// GetUserSource returns the UserSource field value if set, zero value otherwise.
func (o *UpdateIdentitySourceSubdomains200Response) GetUserSource() ListIdentitySources200ResponseAllOfUserSourcesInner {
	if o == nil || IsNil(o.UserSource) {
		var ret ListIdentitySources200ResponseAllOfUserSourcesInner
		return ret
	}
	return *o.UserSource
}

// GetUserSourceOk returns a tuple with the UserSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentitySourceSubdomains200Response) GetUserSourceOk() (*ListIdentitySources200ResponseAllOfUserSourcesInner, bool) {
	if o == nil || IsNil(o.UserSource) {
		return nil, false
	}
	return o.UserSource, true
}

// IsSetUserSource returns a boolean if a field has been set.
func (o *UpdateIdentitySourceSubdomains200Response) IsSetUserSource() bool {
	if o != nil && !IsNil(o.UserSource) {
		return true
	}

	return false
}

// SetUserSource gets a reference to the given ListIdentitySources200ResponseAllOfUserSourcesInner and assigns it to the UserSource field.
func (o *UpdateIdentitySourceSubdomains200Response) SetUserSource(v ListIdentitySources200ResponseAllOfUserSourcesInner) {
	o.UserSource = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UpdateIdentitySourceSubdomains200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIdentitySourceSubdomains200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *UpdateIdentitySourceSubdomains200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UpdateIdentitySourceSubdomains200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o UpdateIdentitySourceSubdomains200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIdentitySourceSubdomains200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserSource) {
		toSerialize["userSource"] = o.UserSource
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableUpdateIdentitySourceSubdomains200Response struct {
	value *UpdateIdentitySourceSubdomains200Response
	isSet bool
}

func (v NullableUpdateIdentitySourceSubdomains200Response) Get() *UpdateIdentitySourceSubdomains200Response {
	return v.value
}

func (v *NullableUpdateIdentitySourceSubdomains200Response) Set(val *UpdateIdentitySourceSubdomains200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIdentitySourceSubdomains200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIdentitySourceSubdomains200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIdentitySourceSubdomains200Response(val *UpdateIdentitySourceSubdomains200Response) *NullableUpdateIdentitySourceSubdomains200Response {
	return &NullableUpdateIdentitySourceSubdomains200Response{value: val, isSet: true}
}

func (v NullableUpdateIdentitySourceSubdomains200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIdentitySourceSubdomains200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



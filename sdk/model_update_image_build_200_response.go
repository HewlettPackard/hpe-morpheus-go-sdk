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

// checks if the UpdateImageBuild200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateImageBuild200Response{}

// UpdateImageBuild200Response struct for UpdateImageBuild200Response
type UpdateImageBuild200Response struct {
	ImageBuild *AddImageBuild200ResponseAllOfImageBuild `json:"imageBuild,omitempty"`
}

// NewUpdateImageBuild200Response instantiates a new UpdateImageBuild200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateImageBuild200Response() *UpdateImageBuild200Response {
	this := UpdateImageBuild200Response{}
	return &this
}

// NewUpdateImageBuild200ResponseWithDefaults instantiates a new UpdateImageBuild200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateImageBuild200ResponseWithDefaults() *UpdateImageBuild200Response {
	this := UpdateImageBuild200Response{}
	return &this
}

// GetImageBuild returns the ImageBuild field value if set, zero value otherwise.
func (o *UpdateImageBuild200Response) GetImageBuild() AddImageBuild200ResponseAllOfImageBuild {
	if o == nil || IsNil(o.ImageBuild) {
		var ret AddImageBuild200ResponseAllOfImageBuild
		return ret
	}
	return *o.ImageBuild
}

// GetImageBuildOk returns a tuple with the ImageBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImageBuild200Response) GetImageBuildOk() (*AddImageBuild200ResponseAllOfImageBuild, bool) {
	if o == nil || IsNil(o.ImageBuild) {
		return nil, false
	}
	return o.ImageBuild, true
}

// IsSetImageBuild returns a boolean if a field has been set.
func (o *UpdateImageBuild200Response) IsSetImageBuild() bool {
	if o != nil && !IsNil(o.ImageBuild) {
		return true
	}

	return false
}

// SetImageBuild gets a reference to the given AddImageBuild200ResponseAllOfImageBuild and assigns it to the ImageBuild field.
func (o *UpdateImageBuild200Response) SetImageBuild(v AddImageBuild200ResponseAllOfImageBuild) {
	o.ImageBuild = &v
}

func (o UpdateImageBuild200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateImageBuild200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageBuild) {
		toSerialize["imageBuild"] = o.ImageBuild
	}
	return toSerialize, nil
}

type NullableUpdateImageBuild200Response struct {
	value *UpdateImageBuild200Response
	isSet bool
}

func (v NullableUpdateImageBuild200Response) Get() *UpdateImageBuild200Response {
	return v.value
}

func (v *NullableUpdateImageBuild200Response) Set(val *UpdateImageBuild200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateImageBuild200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateImageBuild200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateImageBuild200Response(val *UpdateImageBuild200Response) *NullableUpdateImageBuild200Response {
	return &NullableUpdateImageBuild200Response{value: val, isSet: true}
}

func (v NullableUpdateImageBuild200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateImageBuild200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



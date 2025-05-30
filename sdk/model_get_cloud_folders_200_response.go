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

// checks if the GetCloudFolders200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCloudFolders200Response{}

// GetCloudFolders200Response struct for GetCloudFolders200Response
type GetCloudFolders200Response struct {
	Folder *ListCloudFolders200ResponseAllOfFoldersInner `json:"folder,omitempty"`
}

// NewGetCloudFolders200Response instantiates a new GetCloudFolders200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCloudFolders200Response() *GetCloudFolders200Response {
	this := GetCloudFolders200Response{}
	return &this
}

// NewGetCloudFolders200ResponseWithDefaults instantiates a new GetCloudFolders200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCloudFolders200ResponseWithDefaults() *GetCloudFolders200Response {
	this := GetCloudFolders200Response{}
	return &this
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *GetCloudFolders200Response) GetFolder() ListCloudFolders200ResponseAllOfFoldersInner {
	if o == nil || IsNil(o.Folder) {
		var ret ListCloudFolders200ResponseAllOfFoldersInner
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudFolders200Response) GetFolderOk() (*ListCloudFolders200ResponseAllOfFoldersInner, bool) {
	if o == nil || IsNil(o.Folder) {
		return nil, false
	}
	return o.Folder, true
}

// IsSetFolder returns a boolean if a field has been set.
func (o *GetCloudFolders200Response) IsSetFolder() bool {
	if o != nil && !IsNil(o.Folder) {
		return true
	}

	return false
}

// SetFolder gets a reference to the given ListCloudFolders200ResponseAllOfFoldersInner and assigns it to the Folder field.
func (o *GetCloudFolders200Response) SetFolder(v ListCloudFolders200ResponseAllOfFoldersInner) {
	o.Folder = &v
}

func (o GetCloudFolders200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCloudFolders200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Folder) {
		toSerialize["folder"] = o.Folder
	}
	return toSerialize, nil
}

type NullableGetCloudFolders200Response struct {
	value *GetCloudFolders200Response
	isSet bool
}

func (v NullableGetCloudFolders200Response) Get() *GetCloudFolders200Response {
	return v.value
}

func (v *NullableGetCloudFolders200Response) Set(val *GetCloudFolders200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCloudFolders200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCloudFolders200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCloudFolders200Response(val *GetCloudFolders200Response) *NullableGetCloudFolders200Response {
	return &NullableGetCloudFolders200Response{value: val, isSet: true}
}

func (v NullableGetCloudFolders200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCloudFolders200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



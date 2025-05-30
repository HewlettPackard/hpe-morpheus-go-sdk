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

// checks if the AddCatalogItemTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCatalogItemTypeRequest{}

// AddCatalogItemTypeRequest struct for AddCatalogItemTypeRequest
type AddCatalogItemTypeRequest struct {
	CatalogItemType *AddCatalogItemTypeRequestCatalogItemType `json:"catalogItemType,omitempty"`
}

// NewAddCatalogItemTypeRequest instantiates a new AddCatalogItemTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCatalogItemTypeRequest() *AddCatalogItemTypeRequest {
	this := AddCatalogItemTypeRequest{}
	return &this
}

// NewAddCatalogItemTypeRequestWithDefaults instantiates a new AddCatalogItemTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCatalogItemTypeRequestWithDefaults() *AddCatalogItemTypeRequest {
	this := AddCatalogItemTypeRequest{}
	return &this
}

// GetCatalogItemType returns the CatalogItemType field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequest) GetCatalogItemType() AddCatalogItemTypeRequestCatalogItemType {
	if o == nil || IsNil(o.CatalogItemType) {
		var ret AddCatalogItemTypeRequestCatalogItemType
		return ret
	}
	return *o.CatalogItemType
}

// GetCatalogItemTypeOk returns a tuple with the CatalogItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequest) GetCatalogItemTypeOk() (*AddCatalogItemTypeRequestCatalogItemType, bool) {
	if o == nil || IsNil(o.CatalogItemType) {
		return nil, false
	}
	return o.CatalogItemType, true
}

// IsSetCatalogItemType returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequest) IsSetCatalogItemType() bool {
	if o != nil && !IsNil(o.CatalogItemType) {
		return true
	}

	return false
}

// SetCatalogItemType gets a reference to the given AddCatalogItemTypeRequestCatalogItemType and assigns it to the CatalogItemType field.
func (o *AddCatalogItemTypeRequest) SetCatalogItemType(v AddCatalogItemTypeRequestCatalogItemType) {
	o.CatalogItemType = &v
}

func (o AddCatalogItemTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCatalogItemTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CatalogItemType) {
		toSerialize["catalogItemType"] = o.CatalogItemType
	}
	return toSerialize, nil
}

type NullableAddCatalogItemTypeRequest struct {
	value *AddCatalogItemTypeRequest
	isSet bool
}

func (v NullableAddCatalogItemTypeRequest) Get() *AddCatalogItemTypeRequest {
	return v.value
}

func (v *NullableAddCatalogItemTypeRequest) Set(val *AddCatalogItemTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCatalogItemTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCatalogItemTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCatalogItemTypeRequest(val *AddCatalogItemTypeRequest) *NullableAddCatalogItemTypeRequest {
	return &NullableAddCatalogItemTypeRequest{value: val, isSet: true}
}

func (v NullableAddCatalogItemTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCatalogItemTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



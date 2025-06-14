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

// checks if the ListStorageVolumeTypes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListStorageVolumeTypes200Response{}

// ListStorageVolumeTypes200Response struct for ListStorageVolumeTypes200Response
type ListStorageVolumeTypes200Response struct {
	StorageVolumeTypes   []ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner `json:"storageVolumeTypes,omitempty"`
	Meta                 *ListActivity200ResponseAllOfMeta                               `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}                                          `json:",remain"`
}

type _ListStorageVolumeTypes200Response ListStorageVolumeTypes200Response

// NewListStorageVolumeTypes200Response instantiates a new ListStorageVolumeTypes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStorageVolumeTypes200Response() *ListStorageVolumeTypes200Response {
	this := ListStorageVolumeTypes200Response{}
	return &this
}

// NewListStorageVolumeTypes200ResponseWithDefaults instantiates a new ListStorageVolumeTypes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStorageVolumeTypes200ResponseWithDefaults() *ListStorageVolumeTypes200Response {
	this := ListStorageVolumeTypes200Response{}
	return &this
}

// GetStorageVolumeTypes returns the StorageVolumeTypes field value if set, zero value otherwise.
func (o *ListStorageVolumeTypes200Response) GetStorageVolumeTypes() []ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner {
	if o == nil || IsNil(o.StorageVolumeTypes) {
		var ret []ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner
		return ret
	}
	return o.StorageVolumeTypes
}

// GetStorageVolumeTypesOk returns a tuple with the StorageVolumeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStorageVolumeTypes200Response) GetStorageVolumeTypesOk() ([]ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner, bool) {
	if o == nil || IsNil(o.StorageVolumeTypes) {
		return nil, false
	}
	return o.StorageVolumeTypes, true
}

// IsSetStorageVolumeTypes returns a boolean if a field has been set.
func (o *ListStorageVolumeTypes200Response) IsSetStorageVolumeTypes() bool {
	if o != nil && !IsNil(o.StorageVolumeTypes) {
		return true
	}

	return false
}

// SetStorageVolumeTypes gets a reference to the given []ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner and assigns it to the StorageVolumeTypes field.
func (o *ListStorageVolumeTypes200Response) SetStorageVolumeTypes(v []ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) {
	o.StorageVolumeTypes = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListStorageVolumeTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListActivity200ResponseAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStorageVolumeTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// IsSetMeta returns a boolean if a field has been set.
func (o *ListStorageVolumeTypes200Response) IsSetMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListActivity200ResponseAllOfMeta and assigns it to the Meta field.
func (o *ListStorageVolumeTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta) {
	o.Meta = &v
}

func (o ListStorageVolumeTypes200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListStorageVolumeTypes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StorageVolumeTypes) {
		toSerialize["storageVolumeTypes"] = o.StorageVolumeTypes
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListStorageVolumeTypes200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

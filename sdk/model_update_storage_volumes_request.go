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

// checks if the UpdateStorageVolumesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStorageVolumesRequest{}

// UpdateStorageVolumesRequest struct for UpdateStorageVolumesRequest
type UpdateStorageVolumesRequest struct {
	StorageVolume        UpdateStorageVolumesRequestStorageVolume `json:"storageVolume"`
	AdditionalProperties map[string]interface{}                   `json:",remain"`
}

type _UpdateStorageVolumesRequest UpdateStorageVolumesRequest

// NewUpdateStorageVolumesRequest instantiates a new UpdateStorageVolumesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStorageVolumesRequest(storageVolume UpdateStorageVolumesRequestStorageVolume) *UpdateStorageVolumesRequest {
	this := UpdateStorageVolumesRequest{}
	this.StorageVolume = storageVolume
	return &this
}

// NewUpdateStorageVolumesRequestWithDefaults instantiates a new UpdateStorageVolumesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStorageVolumesRequestWithDefaults() *UpdateStorageVolumesRequest {
	this := UpdateStorageVolumesRequest{}
	return &this
}

// GetStorageVolume returns the StorageVolume field value
func (o *UpdateStorageVolumesRequest) GetStorageVolume() UpdateStorageVolumesRequestStorageVolume {
	if o == nil {
		var ret UpdateStorageVolumesRequestStorageVolume
		return ret
	}

	return o.StorageVolume
}

// GetStorageVolumeOk returns a tuple with the StorageVolume field value
// and a boolean to check if the value has been set.
func (o *UpdateStorageVolumesRequest) GetStorageVolumeOk() (*UpdateStorageVolumesRequestStorageVolume, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageVolume, true
}

// SetStorageVolume sets field value
func (o *UpdateStorageVolumesRequest) SetStorageVolume(v UpdateStorageVolumesRequestStorageVolume) {
	o.StorageVolume = v
}

func (o UpdateStorageVolumesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStorageVolumesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["storageVolume"] = o.StorageVolume

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateStorageVolumesRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

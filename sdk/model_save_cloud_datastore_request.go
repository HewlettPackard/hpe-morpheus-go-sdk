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

// checks if the SaveCloudDatastoreRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveCloudDatastoreRequest{}

// SaveCloudDatastoreRequest struct for SaveCloudDatastoreRequest
type SaveCloudDatastoreRequest struct {
	Datastore *SaveCloudDatastoreRequestDatastore `json:"datastore,omitempty"`
}

// NewSaveCloudDatastoreRequest instantiates a new SaveCloudDatastoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveCloudDatastoreRequest() *SaveCloudDatastoreRequest {
	this := SaveCloudDatastoreRequest{}
	return &this
}

// NewSaveCloudDatastoreRequestWithDefaults instantiates a new SaveCloudDatastoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveCloudDatastoreRequestWithDefaults() *SaveCloudDatastoreRequest {
	this := SaveCloudDatastoreRequest{}
	return &this
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *SaveCloudDatastoreRequest) GetDatastore() SaveCloudDatastoreRequestDatastore {
	if o == nil || IsNil(o.Datastore) {
		var ret SaveCloudDatastoreRequestDatastore
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveCloudDatastoreRequest) GetDatastoreOk() (*SaveCloudDatastoreRequestDatastore, bool) {
	if o == nil || IsNil(o.Datastore) {
		return nil, false
	}
	return o.Datastore, true
}

// IsSetDatastore returns a boolean if a field has been set.
func (o *SaveCloudDatastoreRequest) IsSetDatastore() bool {
	if o != nil && !IsNil(o.Datastore) {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given SaveCloudDatastoreRequestDatastore and assigns it to the Datastore field.
func (o *SaveCloudDatastoreRequest) SetDatastore(v SaveCloudDatastoreRequestDatastore) {
	o.Datastore = &v
}

func (o SaveCloudDatastoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveCloudDatastoreRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Datastore) {
		toSerialize["datastore"] = o.Datastore
	}
	return toSerialize, nil
}

type NullableSaveCloudDatastoreRequest struct {
	value *SaveCloudDatastoreRequest
	isSet bool
}

func (v NullableSaveCloudDatastoreRequest) Get() *SaveCloudDatastoreRequest {
	return v.value
}

func (v *NullableSaveCloudDatastoreRequest) Set(val *SaveCloudDatastoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveCloudDatastoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveCloudDatastoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveCloudDatastoreRequest(val *SaveCloudDatastoreRequest) *NullableSaveCloudDatastoreRequest {
	return &NullableSaveCloudDatastoreRequest{value: val, isSet: true}
}

func (v NullableSaveCloudDatastoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveCloudDatastoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



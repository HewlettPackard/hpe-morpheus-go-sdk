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
	"bytes"
	"fmt"
)

// checks if the AddStorageVolumesRequestStorageVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddStorageVolumesRequestStorageVolume{}

// AddStorageVolumesRequestStorageVolume struct for AddStorageVolumesRequestStorageVolume
type AddStorageVolumesRequestStorageVolume struct {
	// A unique name scoped to your account for the storage volume
	Name string `json:"name"`
	// Storage Type Code or ID
	Type string `json:"type"`
	// Configuration object with parameters that vary by `type`.
	Config map[string]interface{} `json:"config,omitempty"`
	StorageServer AddClusterLayoutsRequestLayoutMastersInnerContainerType `json:"storageServer"`
	StorageGroup AddClusterLayoutsRequestLayoutMastersInnerContainerType `json:"storageGroup"`
}

type _AddStorageVolumesRequestStorageVolume AddStorageVolumesRequestStorageVolume

// NewAddStorageVolumesRequestStorageVolume instantiates a new AddStorageVolumesRequestStorageVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddStorageVolumesRequestStorageVolume(name string, type_ string, storageServer AddClusterLayoutsRequestLayoutMastersInnerContainerType, storageGroup AddClusterLayoutsRequestLayoutMastersInnerContainerType) *AddStorageVolumesRequestStorageVolume {
	this := AddStorageVolumesRequestStorageVolume{}
	this.Name = name
	this.Type = type_
	this.StorageServer = storageServer
	this.StorageGroup = storageGroup
	return &this
}

// NewAddStorageVolumesRequestStorageVolumeWithDefaults instantiates a new AddStorageVolumesRequestStorageVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddStorageVolumesRequestStorageVolumeWithDefaults() *AddStorageVolumesRequestStorageVolume {
	this := AddStorageVolumesRequestStorageVolume{}
	return &this
}

// GetName returns the Name field value
func (o *AddStorageVolumesRequestStorageVolume) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddStorageVolumesRequestStorageVolume) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddStorageVolumesRequestStorageVolume) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *AddStorageVolumesRequestStorageVolume) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddStorageVolumesRequestStorageVolume) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddStorageVolumesRequestStorageVolume) SetType(v string) {
	o.Type = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AddStorageVolumesRequestStorageVolume) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddStorageVolumesRequestStorageVolume) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *AddStorageVolumesRequestStorageVolume) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *AddStorageVolumesRequestStorageVolume) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetStorageServer returns the StorageServer field value
func (o *AddStorageVolumesRequestStorageVolume) GetStorageServer() AddClusterLayoutsRequestLayoutMastersInnerContainerType {
	if o == nil {
		var ret AddClusterLayoutsRequestLayoutMastersInnerContainerType
		return ret
	}

	return o.StorageServer
}

// GetStorageServerOk returns a tuple with the StorageServer field value
// and a boolean to check if the value has been set.
func (o *AddStorageVolumesRequestStorageVolume) GetStorageServerOk() (*AddClusterLayoutsRequestLayoutMastersInnerContainerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageServer, true
}

// SetStorageServer sets field value
func (o *AddStorageVolumesRequestStorageVolume) SetStorageServer(v AddClusterLayoutsRequestLayoutMastersInnerContainerType) {
	o.StorageServer = v
}

// GetStorageGroup returns the StorageGroup field value
func (o *AddStorageVolumesRequestStorageVolume) GetStorageGroup() AddClusterLayoutsRequestLayoutMastersInnerContainerType {
	if o == nil {
		var ret AddClusterLayoutsRequestLayoutMastersInnerContainerType
		return ret
	}

	return o.StorageGroup
}

// GetStorageGroupOk returns a tuple with the StorageGroup field value
// and a boolean to check if the value has been set.
func (o *AddStorageVolumesRequestStorageVolume) GetStorageGroupOk() (*AddClusterLayoutsRequestLayoutMastersInnerContainerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageGroup, true
}

// SetStorageGroup sets field value
func (o *AddStorageVolumesRequestStorageVolume) SetStorageGroup(v AddClusterLayoutsRequestLayoutMastersInnerContainerType) {
	o.StorageGroup = v
}

func (o AddStorageVolumesRequestStorageVolume) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddStorageVolumesRequestStorageVolume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	toSerialize["storageServer"] = o.StorageServer
	toSerialize["storageGroup"] = o.StorageGroup
	return toSerialize, nil
}

func (o *AddStorageVolumesRequestStorageVolume) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"storageServer",
		"storageGroup",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddStorageVolumesRequestStorageVolume := _AddStorageVolumesRequestStorageVolume{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddStorageVolumesRequestStorageVolume)

	if err != nil {
		return err
	}

	*o = AddStorageVolumesRequestStorageVolume(varAddStorageVolumesRequestStorageVolume)

	return err
}

type NullableAddStorageVolumesRequestStorageVolume struct {
	value *AddStorageVolumesRequestStorageVolume
	isSet bool
}

func (v NullableAddStorageVolumesRequestStorageVolume) Get() *AddStorageVolumesRequestStorageVolume {
	return v.value
}

func (v *NullableAddStorageVolumesRequestStorageVolume) Set(val *AddStorageVolumesRequestStorageVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableAddStorageVolumesRequestStorageVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableAddStorageVolumesRequestStorageVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddStorageVolumesRequestStorageVolume(val *AddStorageVolumesRequestStorageVolume) *NullableAddStorageVolumesRequestStorageVolume {
	return &NullableAddStorageVolumesRequestStorageVolume{value: val, isSet: true}
}

func (v NullableAddStorageVolumesRequestStorageVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddStorageVolumesRequestStorageVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



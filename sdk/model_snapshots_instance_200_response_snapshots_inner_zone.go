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

// checks if the SnapshotsInstance200ResponseSnapshotsInnerZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotsInstance200ResponseSnapshotsInnerZone{}

// SnapshotsInstance200ResponseSnapshotsInnerZone struct for SnapshotsInstance200ResponseSnapshotsInnerZone
type SnapshotsInstance200ResponseSnapshotsInnerZone struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewSnapshotsInstance200ResponseSnapshotsInnerZone instantiates a new SnapshotsInstance200ResponseSnapshotsInnerZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotsInstance200ResponseSnapshotsInnerZone() *SnapshotsInstance200ResponseSnapshotsInnerZone {
	this := SnapshotsInstance200ResponseSnapshotsInnerZone{}
	return &this
}

// NewSnapshotsInstance200ResponseSnapshotsInnerZoneWithDefaults instantiates a new SnapshotsInstance200ResponseSnapshotsInnerZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotsInstance200ResponseSnapshotsInnerZoneWithDefaults() *SnapshotsInstance200ResponseSnapshotsInnerZone {
	this := SnapshotsInstance200ResponseSnapshotsInnerZone{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SnapshotsInstance200ResponseSnapshotsInnerZone) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotsInstance200ResponseSnapshotsInnerZone) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *SnapshotsInstance200ResponseSnapshotsInnerZone) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SnapshotsInstance200ResponseSnapshotsInnerZone) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SnapshotsInstance200ResponseSnapshotsInnerZone) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotsInstance200ResponseSnapshotsInnerZone) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *SnapshotsInstance200ResponseSnapshotsInnerZone) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SnapshotsInstance200ResponseSnapshotsInnerZone) SetName(v string) {
	o.Name = &v
}

func (o SnapshotsInstance200ResponseSnapshotsInnerZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotsInstance200ResponseSnapshotsInnerZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableSnapshotsInstance200ResponseSnapshotsInnerZone struct {
	value *SnapshotsInstance200ResponseSnapshotsInnerZone
	isSet bool
}

func (v NullableSnapshotsInstance200ResponseSnapshotsInnerZone) Get() *SnapshotsInstance200ResponseSnapshotsInnerZone {
	return v.value
}

func (v *NullableSnapshotsInstance200ResponseSnapshotsInnerZone) Set(val *SnapshotsInstance200ResponseSnapshotsInnerZone) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotsInstance200ResponseSnapshotsInnerZone) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotsInstance200ResponseSnapshotsInnerZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotsInstance200ResponseSnapshotsInnerZone(val *SnapshotsInstance200ResponseSnapshotsInnerZone) *NullableSnapshotsInstance200ResponseSnapshotsInnerZone {
	return &NullableSnapshotsInstance200ResponseSnapshotsInnerZone{value: val, isSet: true}
}

func (v NullableSnapshotsInstance200ResponseSnapshotsInnerZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotsInstance200ResponseSnapshotsInnerZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



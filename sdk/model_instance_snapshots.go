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

// checks if the InstanceSnapshots type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceSnapshots{}

// InstanceSnapshots struct for InstanceSnapshots
type InstanceSnapshots struct {
	// List of snapshot objects
	Snapshots []SnapshotsInstance200ResponseSnapshotsInner `json:"snapshots,omitempty"`
}

// NewInstanceSnapshots instantiates a new InstanceSnapshots object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceSnapshots() *InstanceSnapshots {
	this := InstanceSnapshots{}
	return &this
}

// NewInstanceSnapshotsWithDefaults instantiates a new InstanceSnapshots object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceSnapshotsWithDefaults() *InstanceSnapshots {
	this := InstanceSnapshots{}
	return &this
}

// GetSnapshots returns the Snapshots field value if set, zero value otherwise.
func (o *InstanceSnapshots) GetSnapshots() []SnapshotsInstance200ResponseSnapshotsInner {
	if o == nil || IsNil(o.Snapshots) {
		var ret []SnapshotsInstance200ResponseSnapshotsInner
		return ret
	}
	return o.Snapshots
}

// GetSnapshotsOk returns a tuple with the Snapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceSnapshots) GetSnapshotsOk() ([]SnapshotsInstance200ResponseSnapshotsInner, bool) {
	if o == nil || IsNil(o.Snapshots) {
		return nil, false
	}
	return o.Snapshots, true
}

// IsSetSnapshots returns a boolean if a field has been set.
func (o *InstanceSnapshots) IsSetSnapshots() bool {
	if o != nil && !IsNil(o.Snapshots) {
		return true
	}

	return false
}

// SetSnapshots gets a reference to the given []SnapshotsInstance200ResponseSnapshotsInner and assigns it to the Snapshots field.
func (o *InstanceSnapshots) SetSnapshots(v []SnapshotsInstance200ResponseSnapshotsInner) {
	o.Snapshots = v
}

func (o InstanceSnapshots) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceSnapshots) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Snapshots) {
		toSerialize["snapshots"] = o.Snapshots
	}
	return toSerialize, nil
}

type NullableInstanceSnapshots struct {
	value *InstanceSnapshots
	isSet bool
}

func (v NullableInstanceSnapshots) Get() *InstanceSnapshots {
	return v.value
}

func (v *NullableInstanceSnapshots) Set(val *InstanceSnapshots) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceSnapshots) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceSnapshots) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceSnapshots(val *InstanceSnapshots) *NullableInstanceSnapshots {
	return &NullableInstanceSnapshots{value: val, isSet: true}
}

func (v NullableInstanceSnapshots) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceSnapshots) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



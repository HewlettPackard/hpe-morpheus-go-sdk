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

// checks if the BackupInstanceBackupJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupInstanceBackupJob{}

// BackupInstanceBackupJob struct for BackupInstanceBackupJob
type BackupInstanceBackupJob struct {
	// Enable synthetic full backups on this backup jobAction. Only applies to `kvm` backup type.
	SyntheticFullEnabled *bool `json:"syntheticFullEnabled,omitempty"`
	// the ID of the execute schedule for the synthetic full backup to be created. Only applies to `kvm` backup type.
	SyntheticFullSchedule *int64                 `json:"syntheticFullSchedule,omitempty"`
	AdditionalProperties  map[string]interface{} `json:",remain"`
}

type _BackupInstanceBackupJob BackupInstanceBackupJob

// NewBackupInstanceBackupJob instantiates a new BackupInstanceBackupJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupInstanceBackupJob() *BackupInstanceBackupJob {
	this := BackupInstanceBackupJob{}
	return &this
}

// NewBackupInstanceBackupJobWithDefaults instantiates a new BackupInstanceBackupJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupInstanceBackupJobWithDefaults() *BackupInstanceBackupJob {
	this := BackupInstanceBackupJob{}
	return &this
}

// GetSyntheticFullEnabled returns the SyntheticFullEnabled field value if set, zero value otherwise.
func (o *BackupInstanceBackupJob) GetSyntheticFullEnabled() bool {
	if o == nil || IsNil(o.SyntheticFullEnabled) {
		var ret bool
		return ret
	}
	return *o.SyntheticFullEnabled
}

// GetSyntheticFullEnabledOk returns a tuple with the SyntheticFullEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupInstanceBackupJob) GetSyntheticFullEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SyntheticFullEnabled) {
		return nil, false
	}
	return o.SyntheticFullEnabled, true
}

// IsSetSyntheticFullEnabled returns a boolean if a field has been set.
func (o *BackupInstanceBackupJob) IsSetSyntheticFullEnabled() bool {
	if o != nil && !IsNil(o.SyntheticFullEnabled) {
		return true
	}

	return false
}

// SetSyntheticFullEnabled gets a reference to the given bool and assigns it to the SyntheticFullEnabled field.
func (o *BackupInstanceBackupJob) SetSyntheticFullEnabled(v bool) {
	o.SyntheticFullEnabled = &v
}

// GetSyntheticFullSchedule returns the SyntheticFullSchedule field value if set, zero value otherwise.
func (o *BackupInstanceBackupJob) GetSyntheticFullSchedule() int64 {
	if o == nil || IsNil(o.SyntheticFullSchedule) {
		var ret int64
		return ret
	}
	return *o.SyntheticFullSchedule
}

// GetSyntheticFullScheduleOk returns a tuple with the SyntheticFullSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupInstanceBackupJob) GetSyntheticFullScheduleOk() (*int64, bool) {
	if o == nil || IsNil(o.SyntheticFullSchedule) {
		return nil, false
	}
	return o.SyntheticFullSchedule, true
}

// IsSetSyntheticFullSchedule returns a boolean if a field has been set.
func (o *BackupInstanceBackupJob) IsSetSyntheticFullSchedule() bool {
	if o != nil && !IsNil(o.SyntheticFullSchedule) {
		return true
	}

	return false
}

// SetSyntheticFullSchedule gets a reference to the given int64 and assigns it to the SyntheticFullSchedule field.
func (o *BackupInstanceBackupJob) SetSyntheticFullSchedule(v int64) {
	o.SyntheticFullSchedule = &v
}

func (o BackupInstanceBackupJob) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupInstanceBackupJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SyntheticFullEnabled) {
		toSerialize["syntheticFullEnabled"] = o.SyntheticFullEnabled
	}
	if !IsNil(o.SyntheticFullSchedule) {
		toSerialize["syntheticFullSchedule"] = o.SyntheticFullSchedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *BackupInstanceBackupJob) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

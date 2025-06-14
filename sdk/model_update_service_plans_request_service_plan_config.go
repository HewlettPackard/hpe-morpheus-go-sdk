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

// checks if the UpdateServicePlansRequestServicePlanConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServicePlansRequestServicePlanConfig{}

// UpdateServicePlansRequestServicePlanConfig struct for UpdateServicePlansRequestServicePlanConfig
type UpdateServicePlansRequestServicePlanConfig struct {
	// Specifies range min / max storage multiplier
	StorageSizeType *string `json:"storageSizeType,omitempty"`
	// Specifies range min / max memory multiplier
	MemorySizeType       *string                                           `json:"memorySizeType,omitempty"`
	Ranges               *UpdateServicePlansRequestServicePlanConfigRanges `json:"ranges,omitempty"`
	AdditionalProperties map[string]interface{}                            `json:",remain"`
}

type _UpdateServicePlansRequestServicePlanConfig UpdateServicePlansRequestServicePlanConfig

// NewUpdateServicePlansRequestServicePlanConfig instantiates a new UpdateServicePlansRequestServicePlanConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServicePlansRequestServicePlanConfig() *UpdateServicePlansRequestServicePlanConfig {
	this := UpdateServicePlansRequestServicePlanConfig{}
	var storageSizeType string = "gb"
	this.StorageSizeType = &storageSizeType
	var memorySizeType string = "mb"
	this.MemorySizeType = &memorySizeType
	return &this
}

// NewUpdateServicePlansRequestServicePlanConfigWithDefaults instantiates a new UpdateServicePlansRequestServicePlanConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServicePlansRequestServicePlanConfigWithDefaults() *UpdateServicePlansRequestServicePlanConfig {
	this := UpdateServicePlansRequestServicePlanConfig{}
	var storageSizeType string = "gb"
	this.StorageSizeType = &storageSizeType
	var memorySizeType string = "mb"
	this.MemorySizeType = &memorySizeType
	return &this
}

// GetStorageSizeType returns the StorageSizeType field value if set, zero value otherwise.
func (o *UpdateServicePlansRequestServicePlanConfig) GetStorageSizeType() string {
	if o == nil || IsNil(o.StorageSizeType) {
		var ret string
		return ret
	}
	return *o.StorageSizeType
}

// GetStorageSizeTypeOk returns a tuple with the StorageSizeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServicePlansRequestServicePlanConfig) GetStorageSizeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StorageSizeType) {
		return nil, false
	}
	return o.StorageSizeType, true
}

// IsSetStorageSizeType returns a boolean if a field has been set.
func (o *UpdateServicePlansRequestServicePlanConfig) IsSetStorageSizeType() bool {
	if o != nil && !IsNil(o.StorageSizeType) {
		return true
	}

	return false
}

// SetStorageSizeType gets a reference to the given string and assigns it to the StorageSizeType field.
func (o *UpdateServicePlansRequestServicePlanConfig) SetStorageSizeType(v string) {
	o.StorageSizeType = &v
}

// GetMemorySizeType returns the MemorySizeType field value if set, zero value otherwise.
func (o *UpdateServicePlansRequestServicePlanConfig) GetMemorySizeType() string {
	if o == nil || IsNil(o.MemorySizeType) {
		var ret string
		return ret
	}
	return *o.MemorySizeType
}

// GetMemorySizeTypeOk returns a tuple with the MemorySizeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServicePlansRequestServicePlanConfig) GetMemorySizeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MemorySizeType) {
		return nil, false
	}
	return o.MemorySizeType, true
}

// IsSetMemorySizeType returns a boolean if a field has been set.
func (o *UpdateServicePlansRequestServicePlanConfig) IsSetMemorySizeType() bool {
	if o != nil && !IsNil(o.MemorySizeType) {
		return true
	}

	return false
}

// SetMemorySizeType gets a reference to the given string and assigns it to the MemorySizeType field.
func (o *UpdateServicePlansRequestServicePlanConfig) SetMemorySizeType(v string) {
	o.MemorySizeType = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *UpdateServicePlansRequestServicePlanConfig) GetRanges() UpdateServicePlansRequestServicePlanConfigRanges {
	if o == nil || IsNil(o.Ranges) {
		var ret UpdateServicePlansRequestServicePlanConfigRanges
		return ret
	}
	return *o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServicePlansRequestServicePlanConfig) GetRangesOk() (*UpdateServicePlansRequestServicePlanConfigRanges, bool) {
	if o == nil || IsNil(o.Ranges) {
		return nil, false
	}
	return o.Ranges, true
}

// IsSetRanges returns a boolean if a field has been set.
func (o *UpdateServicePlansRequestServicePlanConfig) IsSetRanges() bool {
	if o != nil && !IsNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given UpdateServicePlansRequestServicePlanConfigRanges and assigns it to the Ranges field.
func (o *UpdateServicePlansRequestServicePlanConfig) SetRanges(v UpdateServicePlansRequestServicePlanConfigRanges) {
	o.Ranges = &v
}

func (o UpdateServicePlansRequestServicePlanConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServicePlansRequestServicePlanConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StorageSizeType) {
		toSerialize["storageSizeType"] = o.StorageSizeType
	}
	if !IsNil(o.MemorySizeType) {
		toSerialize["memorySizeType"] = o.MemorySizeType
	}
	if !IsNil(o.Ranges) {
		toSerialize["ranges"] = o.Ranges
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateServicePlansRequestServicePlanConfig) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the MaxCoresPolicyTypeConfiguration1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaxCoresPolicyTypeConfiguration1{}

// MaxCoresPolicyTypeConfiguration1 Configuration settings for the following policy types: - Max Cores
type MaxCoresPolicyTypeConfiguration1 struct {
	MaxCores             *string                `json:"maxCores,omitempty"`
	ExcludeContainers    *bool                  `json:"excludeContainers,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _MaxCoresPolicyTypeConfiguration1 MaxCoresPolicyTypeConfiguration1

// NewMaxCoresPolicyTypeConfiguration1 instantiates a new MaxCoresPolicyTypeConfiguration1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxCoresPolicyTypeConfiguration1() *MaxCoresPolicyTypeConfiguration1 {
	this := MaxCoresPolicyTypeConfiguration1{}
	return &this
}

// NewMaxCoresPolicyTypeConfiguration1WithDefaults instantiates a new MaxCoresPolicyTypeConfiguration1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxCoresPolicyTypeConfiguration1WithDefaults() *MaxCoresPolicyTypeConfiguration1 {
	this := MaxCoresPolicyTypeConfiguration1{}
	return &this
}

// GetMaxCores returns the MaxCores field value if set, zero value otherwise.
func (o *MaxCoresPolicyTypeConfiguration1) GetMaxCores() string {
	if o == nil || IsNil(o.MaxCores) {
		var ret string
		return ret
	}
	return *o.MaxCores
}

// GetMaxCoresOk returns a tuple with the MaxCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxCoresPolicyTypeConfiguration1) GetMaxCoresOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCores) {
		return nil, false
	}
	return o.MaxCores, true
}

// IsSetMaxCores returns a boolean if a field has been set.
func (o *MaxCoresPolicyTypeConfiguration1) IsSetMaxCores() bool {
	if o != nil && !IsNil(o.MaxCores) {
		return true
	}

	return false
}

// SetMaxCores gets a reference to the given string and assigns it to the MaxCores field.
func (o *MaxCoresPolicyTypeConfiguration1) SetMaxCores(v string) {
	o.MaxCores = &v
}

// GetExcludeContainers returns the ExcludeContainers field value if set, zero value otherwise.
func (o *MaxCoresPolicyTypeConfiguration1) GetExcludeContainers() bool {
	if o == nil || IsNil(o.ExcludeContainers) {
		var ret bool
		return ret
	}
	return *o.ExcludeContainers
}

// GetExcludeContainersOk returns a tuple with the ExcludeContainers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxCoresPolicyTypeConfiguration1) GetExcludeContainersOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeContainers) {
		return nil, false
	}
	return o.ExcludeContainers, true
}

// IsSetExcludeContainers returns a boolean if a field has been set.
func (o *MaxCoresPolicyTypeConfiguration1) IsSetExcludeContainers() bool {
	if o != nil && !IsNil(o.ExcludeContainers) {
		return true
	}

	return false
}

// SetExcludeContainers gets a reference to the given bool and assigns it to the ExcludeContainers field.
func (o *MaxCoresPolicyTypeConfiguration1) SetExcludeContainers(v bool) {
	o.ExcludeContainers = &v
}

func (o MaxCoresPolicyTypeConfiguration1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxCoresPolicyTypeConfiguration1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxCores) {
		toSerialize["maxCores"] = o.MaxCores
	}
	if !IsNil(o.ExcludeContainers) {
		toSerialize["excludeContainers"] = o.ExcludeContainers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *MaxCoresPolicyTypeConfiguration1) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

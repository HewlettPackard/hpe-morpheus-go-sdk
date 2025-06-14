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

// checks if the MaxVirtualServersPolicyTypeConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaxVirtualServersPolicyTypeConfiguration{}

// MaxVirtualServersPolicyTypeConfiguration Configuration settings for the following policy types: - Max Virtual Servers
type MaxVirtualServersPolicyTypeConfiguration struct {
	MaxVirtualServers    *string                `json:"maxVirtualServers,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _MaxVirtualServersPolicyTypeConfiguration MaxVirtualServersPolicyTypeConfiguration

// NewMaxVirtualServersPolicyTypeConfiguration instantiates a new MaxVirtualServersPolicyTypeConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxVirtualServersPolicyTypeConfiguration() *MaxVirtualServersPolicyTypeConfiguration {
	this := MaxVirtualServersPolicyTypeConfiguration{}
	return &this
}

// NewMaxVirtualServersPolicyTypeConfigurationWithDefaults instantiates a new MaxVirtualServersPolicyTypeConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxVirtualServersPolicyTypeConfigurationWithDefaults() *MaxVirtualServersPolicyTypeConfiguration {
	this := MaxVirtualServersPolicyTypeConfiguration{}
	return &this
}

// GetMaxVirtualServers returns the MaxVirtualServers field value if set, zero value otherwise.
func (o *MaxVirtualServersPolicyTypeConfiguration) GetMaxVirtualServers() string {
	if o == nil || IsNil(o.MaxVirtualServers) {
		var ret string
		return ret
	}
	return *o.MaxVirtualServers
}

// GetMaxVirtualServersOk returns a tuple with the MaxVirtualServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxVirtualServersPolicyTypeConfiguration) GetMaxVirtualServersOk() (*string, bool) {
	if o == nil || IsNil(o.MaxVirtualServers) {
		return nil, false
	}
	return o.MaxVirtualServers, true
}

// IsSetMaxVirtualServers returns a boolean if a field has been set.
func (o *MaxVirtualServersPolicyTypeConfiguration) IsSetMaxVirtualServers() bool {
	if o != nil && !IsNil(o.MaxVirtualServers) {
		return true
	}

	return false
}

// SetMaxVirtualServers gets a reference to the given string and assigns it to the MaxVirtualServers field.
func (o *MaxVirtualServersPolicyTypeConfiguration) SetMaxVirtualServers(v string) {
	o.MaxVirtualServers = &v
}

func (o MaxVirtualServersPolicyTypeConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxVirtualServersPolicyTypeConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxVirtualServers) {
		toSerialize["maxVirtualServers"] = o.MaxVirtualServers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *MaxVirtualServersPolicyTypeConfiguration) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

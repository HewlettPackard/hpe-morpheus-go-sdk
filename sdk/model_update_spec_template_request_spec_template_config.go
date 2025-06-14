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

// checks if the UpdateSpecTemplateRequestSpecTemplateConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSpecTemplateRequestSpecTemplateConfig{}

// UpdateSpecTemplateRequestSpecTemplateConfig The Cloud Formation type supports some additional configuration parameters
type UpdateSpecTemplateRequestSpecTemplateConfig struct {
	Cloudformation       *UpdateSpecTemplateRequestSpecTemplateConfigCloudformation `json:"cloudformation,omitempty"`
	AdditionalProperties map[string]interface{}                                     `json:",remain"`
}

type _UpdateSpecTemplateRequestSpecTemplateConfig UpdateSpecTemplateRequestSpecTemplateConfig

// NewUpdateSpecTemplateRequestSpecTemplateConfig instantiates a new UpdateSpecTemplateRequestSpecTemplateConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSpecTemplateRequestSpecTemplateConfig() *UpdateSpecTemplateRequestSpecTemplateConfig {
	this := UpdateSpecTemplateRequestSpecTemplateConfig{}
	return &this
}

// NewUpdateSpecTemplateRequestSpecTemplateConfigWithDefaults instantiates a new UpdateSpecTemplateRequestSpecTemplateConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSpecTemplateRequestSpecTemplateConfigWithDefaults() *UpdateSpecTemplateRequestSpecTemplateConfig {
	this := UpdateSpecTemplateRequestSpecTemplateConfig{}
	return &this
}

// GetCloudformation returns the Cloudformation field value if set, zero value otherwise.
func (o *UpdateSpecTemplateRequestSpecTemplateConfig) GetCloudformation() UpdateSpecTemplateRequestSpecTemplateConfigCloudformation {
	if o == nil || IsNil(o.Cloudformation) {
		var ret UpdateSpecTemplateRequestSpecTemplateConfigCloudformation
		return ret
	}
	return *o.Cloudformation
}

// GetCloudformationOk returns a tuple with the Cloudformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSpecTemplateRequestSpecTemplateConfig) GetCloudformationOk() (*UpdateSpecTemplateRequestSpecTemplateConfigCloudformation, bool) {
	if o == nil || IsNil(o.Cloudformation) {
		return nil, false
	}
	return o.Cloudformation, true
}

// IsSetCloudformation returns a boolean if a field has been set.
func (o *UpdateSpecTemplateRequestSpecTemplateConfig) IsSetCloudformation() bool {
	if o != nil && !IsNil(o.Cloudformation) {
		return true
	}

	return false
}

// SetCloudformation gets a reference to the given UpdateSpecTemplateRequestSpecTemplateConfigCloudformation and assigns it to the Cloudformation field.
func (o *UpdateSpecTemplateRequestSpecTemplateConfig) SetCloudformation(v UpdateSpecTemplateRequestSpecTemplateConfigCloudformation) {
	o.Cloudformation = &v
}

func (o UpdateSpecTemplateRequestSpecTemplateConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSpecTemplateRequestSpecTemplateConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cloudformation) {
		toSerialize["cloudformation"] = o.Cloudformation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateSpecTemplateRequestSpecTemplateConfig) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

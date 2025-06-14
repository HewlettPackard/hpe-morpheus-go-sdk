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

// checks if the AddSpecTemplateRequestSpecTemplateConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSpecTemplateRequestSpecTemplateConfig{}

// AddSpecTemplateRequestSpecTemplateConfig The Cloud Formation type supports some additional configuration parameters
type AddSpecTemplateRequestSpecTemplateConfig struct {
	Cloudformation       *AddSpecTemplateRequestSpecTemplateConfigCloudformation `json:"cloudformation,omitempty"`
	AdditionalProperties map[string]interface{}                                  `json:",remain"`
}

type _AddSpecTemplateRequestSpecTemplateConfig AddSpecTemplateRequestSpecTemplateConfig

// NewAddSpecTemplateRequestSpecTemplateConfig instantiates a new AddSpecTemplateRequestSpecTemplateConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSpecTemplateRequestSpecTemplateConfig() *AddSpecTemplateRequestSpecTemplateConfig {
	this := AddSpecTemplateRequestSpecTemplateConfig{}
	return &this
}

// NewAddSpecTemplateRequestSpecTemplateConfigWithDefaults instantiates a new AddSpecTemplateRequestSpecTemplateConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSpecTemplateRequestSpecTemplateConfigWithDefaults() *AddSpecTemplateRequestSpecTemplateConfig {
	this := AddSpecTemplateRequestSpecTemplateConfig{}
	return &this
}

// GetCloudformation returns the Cloudformation field value if set, zero value otherwise.
func (o *AddSpecTemplateRequestSpecTemplateConfig) GetCloudformation() AddSpecTemplateRequestSpecTemplateConfigCloudformation {
	if o == nil || IsNil(o.Cloudformation) {
		var ret AddSpecTemplateRequestSpecTemplateConfigCloudformation
		return ret
	}
	return *o.Cloudformation
}

// GetCloudformationOk returns a tuple with the Cloudformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSpecTemplateRequestSpecTemplateConfig) GetCloudformationOk() (*AddSpecTemplateRequestSpecTemplateConfigCloudformation, bool) {
	if o == nil || IsNil(o.Cloudformation) {
		return nil, false
	}
	return o.Cloudformation, true
}

// IsSetCloudformation returns a boolean if a field has been set.
func (o *AddSpecTemplateRequestSpecTemplateConfig) IsSetCloudformation() bool {
	if o != nil && !IsNil(o.Cloudformation) {
		return true
	}

	return false
}

// SetCloudformation gets a reference to the given AddSpecTemplateRequestSpecTemplateConfigCloudformation and assigns it to the Cloudformation field.
func (o *AddSpecTemplateRequestSpecTemplateConfig) SetCloudformation(v AddSpecTemplateRequestSpecTemplateConfigCloudformation) {
	o.Cloudformation = &v
}

func (o AddSpecTemplateRequestSpecTemplateConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSpecTemplateRequestSpecTemplateConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cloudformation) {
		toSerialize["cloudformation"] = o.Cloudformation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddSpecTemplateRequestSpecTemplateConfig) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

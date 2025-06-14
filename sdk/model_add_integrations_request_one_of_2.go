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

// checks if the AddIntegrationsRequestOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddIntegrationsRequestOneOf2{}

// AddIntegrationsRequestOneOf2 struct for AddIntegrationsRequestOneOf2
type AddIntegrationsRequestOneOf2 struct {
	Integration          AddIntegrationsRequestOneOf2Integration `json:"integration"`
	AdditionalProperties map[string]interface{}                  `json:",remain"`
}

type _AddIntegrationsRequestOneOf2 AddIntegrationsRequestOneOf2

// NewAddIntegrationsRequestOneOf2 instantiates a new AddIntegrationsRequestOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIntegrationsRequestOneOf2(integration AddIntegrationsRequestOneOf2Integration) *AddIntegrationsRequestOneOf2 {
	this := AddIntegrationsRequestOneOf2{}
	this.Integration = integration
	return &this
}

// NewAddIntegrationsRequestOneOf2WithDefaults instantiates a new AddIntegrationsRequestOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIntegrationsRequestOneOf2WithDefaults() *AddIntegrationsRequestOneOf2 {
	this := AddIntegrationsRequestOneOf2{}
	return &this
}

// GetIntegration returns the Integration field value
func (o *AddIntegrationsRequestOneOf2) GetIntegration() AddIntegrationsRequestOneOf2Integration {
	if o == nil {
		var ret AddIntegrationsRequestOneOf2Integration
		return ret
	}

	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *AddIntegrationsRequestOneOf2) GetIntegrationOk() (*AddIntegrationsRequestOneOf2Integration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value
func (o *AddIntegrationsRequestOneOf2) SetIntegration(v AddIntegrationsRequestOneOf2Integration) {
	o.Integration = v
}

func (o AddIntegrationsRequestOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddIntegrationsRequestOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["integration"] = o.Integration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddIntegrationsRequestOneOf2) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

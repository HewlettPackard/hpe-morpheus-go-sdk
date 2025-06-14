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

// checks if the ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig{}

// ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig struct for ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig
type ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig struct {
	CredentialTypes      []string               `json:"credentialTypes,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig

// NewListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig instantiates a new ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig() *ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig {
	this := ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig{}
	return &this
}

// NewListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfigWithDefaults instantiates a new ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfigWithDefaults() *ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig {
	this := ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig{}
	return &this
}

// GetCredentialTypes returns the CredentialTypes field value if set, zero value otherwise.
func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig) GetCredentialTypes() []string {
	if o == nil || IsNil(o.CredentialTypes) {
		var ret []string
		return ret
	}
	return o.CredentialTypes
}

// GetCredentialTypesOk returns a tuple with the CredentialTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig) GetCredentialTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.CredentialTypes) {
		return nil, false
	}
	return o.CredentialTypes, true
}

// IsSetCredentialTypes returns a boolean if a field has been set.
func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig) IsSetCredentialTypes() bool {
	if o != nil && !IsNil(o.CredentialTypes) {
		return true
	}

	return false
}

// SetCredentialTypes gets a reference to the given []string and assigns it to the CredentialTypes field.
func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig) SetCredentialTypes(v []string) {
	o.CredentialTypes = v
}

func (o ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CredentialTypes) {
		toSerialize["credentialTypes"] = o.CredentialTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerOptionTypesInnerConfig) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

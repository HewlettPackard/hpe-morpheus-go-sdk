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

// checks if the ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config{}

// ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config struct for ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config
type ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config struct {
	Zones []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner `json:"zones,omitempty"`
}

// NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config {
	this := ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config{}
	return &this
}

// NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigWithDefaults instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigWithDefaults() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config {
	this := ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config{}
	return &this
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) GetZones() []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner {
	if o == nil || IsNil(o.Zones) {
		var ret []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner
		return ret
	}
	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) GetZonesOk() ([]ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner, bool) {
	if o == nil || IsNil(o.Zones) {
		return nil, false
	}
	return o.Zones, true
}

// IsSetZones returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) IsSetZones() bool {
	if o != nil && !IsNil(o.Zones) {
		return true
	}

	return false
}

// SetZones gets a reference to the given []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner and assigns it to the Zones field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) SetZones(v []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner) {
	o.Zones = v
}

func (o ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Zones) {
		toSerialize["zones"] = o.Zones
	}
	return toSerialize, nil
}

type NullableListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config struct {
	value *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config
	isSet bool
}

func (v NullableListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) Get() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config {
	return v.value
}

func (v *NullableListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) Set(val *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) {
	v.value = val
	v.isSet = true
}

func (v NullableListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) IsSet() bool {
	return v.isSet
}

func (v *NullableListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config(val *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) *NullableListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config {
	return &NullableListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config{value: val, isSet: true}
}

func (v NullableListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2Config) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



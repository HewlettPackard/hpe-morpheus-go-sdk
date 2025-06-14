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

// checks if the ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner{}

// ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner struct for ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner
type ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner struct {
	ZoneName             *string                `json:"zoneName,omitempty"`
	ReverseZone          *string                `json:"reverseZone,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner

// NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner {
	this := ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner{}
	return &this
}

// NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInnerWithDefaults instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInnerWithDefaults() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner {
	this := ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner{}
	return &this
}

// GetZoneName returns the ZoneName field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner) GetZoneName() string {
	if o == nil || IsNil(o.ZoneName) {
		var ret string
		return ret
	}
	return *o.ZoneName
}

// GetZoneNameOk returns a tuple with the ZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner) GetZoneNameOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneName) {
		return nil, false
	}
	return o.ZoneName, true
}

// IsSetZoneName returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner) IsSetZoneName() bool {
	if o != nil && !IsNil(o.ZoneName) {
		return true
	}

	return false
}

// SetZoneName gets a reference to the given string and assigns it to the ZoneName field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner) SetZoneName(v string) {
	o.ZoneName = &v
}

// GetReverseZone returns the ReverseZone field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner) GetReverseZone() string {
	if o == nil || IsNil(o.ReverseZone) {
		var ret string
		return ret
	}
	return *o.ReverseZone
}

// GetReverseZoneOk returns a tuple with the ReverseZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner) GetReverseZoneOk() (*string, bool) {
	if o == nil || IsNil(o.ReverseZone) {
		return nil, false
	}
	return o.ReverseZone, true
}

// IsSetReverseZone returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner) IsSetReverseZone() bool {
	if o != nil && !IsNil(o.ReverseZone) {
		return true
	}

	return false
}

// SetReverseZone gets a reference to the given string and assigns it to the ReverseZone field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner) SetReverseZone(v string) {
	o.ReverseZone = &v
}

func (o ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ZoneName) {
		toSerialize["zoneName"] = o.ZoneName
	}
	if !IsNil(o.ReverseZone) {
		toSerialize["reverseZone"] = o.ReverseZone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf2ConfigZonesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

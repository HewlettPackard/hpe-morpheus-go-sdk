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

// checks if the ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings{}

// ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings struct for ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings
type ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings struct {
	EntityId             *string                `json:"entityId,omitempty"`
	AcsUrl               *string                `json:"acsUrl,omitempty"`
	SpMetadata           *string                `json:"spMetadata,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings

// NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings instantiates a new ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings() *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings {
	this := ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings{}
	return &this
}

// NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettingsWithDefaults instantiates a new ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettingsWithDefaults() *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings {
	this := ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// IsSetEntityId returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) IsSetEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) SetEntityId(v string) {
	o.EntityId = &v
}

// GetAcsUrl returns the AcsUrl field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) GetAcsUrl() string {
	if o == nil || IsNil(o.AcsUrl) {
		var ret string
		return ret
	}
	return *o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) GetAcsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AcsUrl) {
		return nil, false
	}
	return o.AcsUrl, true
}

// IsSetAcsUrl returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) IsSetAcsUrl() bool {
	if o != nil && !IsNil(o.AcsUrl) {
		return true
	}

	return false
}

// SetAcsUrl gets a reference to the given string and assigns it to the AcsUrl field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) SetAcsUrl(v string) {
	o.AcsUrl = &v
}

// GetSpMetadata returns the SpMetadata field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) GetSpMetadata() string {
	if o == nil || IsNil(o.SpMetadata) {
		var ret string
		return ret
	}
	return *o.SpMetadata
}

// GetSpMetadataOk returns a tuple with the SpMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) GetSpMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.SpMetadata) {
		return nil, false
	}
	return o.SpMetadata, true
}

// IsSetSpMetadata returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) IsSetSpMetadata() bool {
	if o != nil && !IsNil(o.SpMetadata) {
		return true
	}

	return false
}

// SetSpMetadata gets a reference to the given string and assigns it to the SpMetadata field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) SetSpMetadata(v string) {
	o.SpMetadata = &v
}

func (o ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.AcsUrl) {
		toSerialize["acsUrl"] = o.AcsUrl
	}
	if !IsNil(o.SpMetadata) {
		toSerialize["spMetadata"] = o.SpMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf5ProviderSettings) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

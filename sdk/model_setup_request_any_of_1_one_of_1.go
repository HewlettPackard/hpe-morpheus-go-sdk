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

// checks if the SetupRequestAnyOf1OneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetupRequestAnyOf1OneOf1{}

// SetupRequestAnyOf1OneOf1 Object for registering with the [Morpheus Hub](https://morpheushub.com). This is only required for `hubmode=register`.
type SetupRequestAnyOf1OneOf1 struct {
	Hub                  SetupRequestAnyOf1OneOf1Hub `json:"hub"`
	AdditionalProperties map[string]interface{}      `json:",remain"`
}

type _SetupRequestAnyOf1OneOf1 SetupRequestAnyOf1OneOf1

// NewSetupRequestAnyOf1OneOf1 instantiates a new SetupRequestAnyOf1OneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetupRequestAnyOf1OneOf1(hub SetupRequestAnyOf1OneOf1Hub) *SetupRequestAnyOf1OneOf1 {
	this := SetupRequestAnyOf1OneOf1{}
	this.Hub = hub
	return &this
}

// NewSetupRequestAnyOf1OneOf1WithDefaults instantiates a new SetupRequestAnyOf1OneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetupRequestAnyOf1OneOf1WithDefaults() *SetupRequestAnyOf1OneOf1 {
	this := SetupRequestAnyOf1OneOf1{}
	return &this
}

// GetHub returns the Hub field value
func (o *SetupRequestAnyOf1OneOf1) GetHub() SetupRequestAnyOf1OneOf1Hub {
	if o == nil {
		var ret SetupRequestAnyOf1OneOf1Hub
		return ret
	}

	return o.Hub
}

// GetHubOk returns a tuple with the Hub field value
// and a boolean to check if the value has been set.
func (o *SetupRequestAnyOf1OneOf1) GetHubOk() (*SetupRequestAnyOf1OneOf1Hub, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hub, true
}

// SetHub sets field value
func (o *SetupRequestAnyOf1OneOf1) SetHub(v SetupRequestAnyOf1OneOf1Hub) {
	o.Hub = v
}

func (o SetupRequestAnyOf1OneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetupRequestAnyOf1OneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hub"] = o.Hub

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *SetupRequestAnyOf1OneOf1) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

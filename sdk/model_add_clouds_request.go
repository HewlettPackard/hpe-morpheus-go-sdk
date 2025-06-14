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

// checks if the AddCloudsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCloudsRequest{}

// AddCloudsRequest struct for AddCloudsRequest
type AddCloudsRequest struct {
	Zone                 AddCloudsRequestZone   `json:"zone"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddCloudsRequest AddCloudsRequest

// NewAddCloudsRequest instantiates a new AddCloudsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCloudsRequest(zone AddCloudsRequestZone) *AddCloudsRequest {
	this := AddCloudsRequest{}
	this.Zone = zone
	return &this
}

// NewAddCloudsRequestWithDefaults instantiates a new AddCloudsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCloudsRequestWithDefaults() *AddCloudsRequest {
	this := AddCloudsRequest{}
	return &this
}

// GetZone returns the Zone field value
func (o *AddCloudsRequest) GetZone() AddCloudsRequestZone {
	if o == nil {
		var ret AddCloudsRequestZone
		return ret
	}

	return o.Zone
}

// GetZoneOk returns a tuple with the Zone field value
// and a boolean to check if the value has been set.
func (o *AddCloudsRequest) GetZoneOk() (*AddCloudsRequestZone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zone, true
}

// SetZone sets field value
func (o *AddCloudsRequest) SetZone(v AddCloudsRequestZone) {
	o.Zone = v
}

func (o AddCloudsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCloudsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["zone"] = o.Zone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddCloudsRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

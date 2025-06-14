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

// checks if the UpdateInstance200ResponseAllOfOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInstance200ResponseAllOfOneOf{}

// UpdateInstance200ResponseAllOfOneOf struct for UpdateInstance200ResponseAllOfOneOf
type UpdateInstance200ResponseAllOfOneOf struct {
	Instance UpdateInstance200ResponseAllOfOneOfInstance `json:"instance"`
	// The Cloud ID to provision the instance onto.
	ZoneId               int64                  `json:"zoneId"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _UpdateInstance200ResponseAllOfOneOf UpdateInstance200ResponseAllOfOneOf

// NewUpdateInstance200ResponseAllOfOneOf instantiates a new UpdateInstance200ResponseAllOfOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInstance200ResponseAllOfOneOf(instance UpdateInstance200ResponseAllOfOneOfInstance, zoneId int64) *UpdateInstance200ResponseAllOfOneOf {
	this := UpdateInstance200ResponseAllOfOneOf{}
	this.Instance = instance
	this.ZoneId = zoneId
	return &this
}

// NewUpdateInstance200ResponseAllOfOneOfWithDefaults instantiates a new UpdateInstance200ResponseAllOfOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInstance200ResponseAllOfOneOfWithDefaults() *UpdateInstance200ResponseAllOfOneOf {
	this := UpdateInstance200ResponseAllOfOneOf{}
	return &this
}

// GetInstance returns the Instance field value
func (o *UpdateInstance200ResponseAllOfOneOf) GetInstance() UpdateInstance200ResponseAllOfOneOfInstance {
	if o == nil {
		var ret UpdateInstance200ResponseAllOfOneOfInstance
		return ret
	}

	return o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value
// and a boolean to check if the value has been set.
func (o *UpdateInstance200ResponseAllOfOneOf) GetInstanceOk() (*UpdateInstance200ResponseAllOfOneOfInstance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instance, true
}

// SetInstance sets field value
func (o *UpdateInstance200ResponseAllOfOneOf) SetInstance(v UpdateInstance200ResponseAllOfOneOfInstance) {
	o.Instance = v
}

// GetZoneId returns the ZoneId field value
func (o *UpdateInstance200ResponseAllOfOneOf) GetZoneId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value
// and a boolean to check if the value has been set.
func (o *UpdateInstance200ResponseAllOfOneOf) GetZoneIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneId, true
}

// SetZoneId sets field value
func (o *UpdateInstance200ResponseAllOfOneOf) SetZoneId(v int64) {
	o.ZoneId = v
}

func (o UpdateInstance200ResponseAllOfOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInstance200ResponseAllOfOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instance"] = o.Instance
	toSerialize["zoneId"] = o.ZoneId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateInstance200ResponseAllOfOneOf) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

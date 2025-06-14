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

// checks if the AddAppInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddAppInstanceRequest{}

// AddAppInstanceRequest struct for AddAppInstanceRequest
type AddAppInstanceRequest struct {
	// The ID of the instance being added
	InstanceId int64 `json:"instanceId"`
	// The Name of the Tier
	TierName             string                 `json:"tierName"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddAppInstanceRequest AddAppInstanceRequest

// NewAddAppInstanceRequest instantiates a new AddAppInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAppInstanceRequest(instanceId int64, tierName string) *AddAppInstanceRequest {
	this := AddAppInstanceRequest{}
	this.InstanceId = instanceId
	this.TierName = tierName
	return &this
}

// NewAddAppInstanceRequestWithDefaults instantiates a new AddAppInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAppInstanceRequestWithDefaults() *AddAppInstanceRequest {
	this := AddAppInstanceRequest{}
	return &this
}

// GetInstanceId returns the InstanceId field value
func (o *AddAppInstanceRequest) GetInstanceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *AddAppInstanceRequest) GetInstanceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *AddAppInstanceRequest) SetInstanceId(v int64) {
	o.InstanceId = v
}

// GetTierName returns the TierName field value
func (o *AddAppInstanceRequest) GetTierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TierName
}

// GetTierNameOk returns a tuple with the TierName field value
// and a boolean to check if the value has been set.
func (o *AddAppInstanceRequest) GetTierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TierName, true
}

// SetTierName sets field value
func (o *AddAppInstanceRequest) SetTierName(v string) {
	o.TierName = v
}

func (o AddAppInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddAppInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["tierName"] = o.TierName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddAppInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

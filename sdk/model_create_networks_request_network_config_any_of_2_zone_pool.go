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

// checks if the CreateNetworksRequestNetworkConfigAnyOf2ZonePool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworksRequestNetworkConfigAnyOf2ZonePool{}

// CreateNetworksRequestNetworkConfigAnyOf2ZonePool struct for CreateNetworksRequestNetworkConfigAnyOf2ZonePool
type CreateNetworksRequestNetworkConfigAnyOf2ZonePool struct {
	// Morpheus resource pool ID of the GCP Project for the network.
	Id                   int32                  `json:"id"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _CreateNetworksRequestNetworkConfigAnyOf2ZonePool CreateNetworksRequestNetworkConfigAnyOf2ZonePool

// NewCreateNetworksRequestNetworkConfigAnyOf2ZonePool instantiates a new CreateNetworksRequestNetworkConfigAnyOf2ZonePool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworksRequestNetworkConfigAnyOf2ZonePool(id int32) *CreateNetworksRequestNetworkConfigAnyOf2ZonePool {
	this := CreateNetworksRequestNetworkConfigAnyOf2ZonePool{}
	this.Id = id
	return &this
}

// NewCreateNetworksRequestNetworkConfigAnyOf2ZonePoolWithDefaults instantiates a new CreateNetworksRequestNetworkConfigAnyOf2ZonePool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworksRequestNetworkConfigAnyOf2ZonePoolWithDefaults() *CreateNetworksRequestNetworkConfigAnyOf2ZonePool {
	this := CreateNetworksRequestNetworkConfigAnyOf2ZonePool{}
	return &this
}

// GetId returns the Id field value
func (o *CreateNetworksRequestNetworkConfigAnyOf2ZonePool) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateNetworksRequestNetworkConfigAnyOf2ZonePool) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateNetworksRequestNetworkConfigAnyOf2ZonePool) SetId(v int32) {
	o.Id = v
}

func (o CreateNetworksRequestNetworkConfigAnyOf2ZonePool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworksRequestNetworkConfigAnyOf2ZonePool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *CreateNetworksRequestNetworkConfigAnyOf2ZonePool) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

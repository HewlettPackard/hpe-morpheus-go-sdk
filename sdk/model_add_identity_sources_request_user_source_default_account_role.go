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

// checks if the AddIdentitySourcesRequestUserSourceDefaultAccountRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddIdentitySourcesRequestUserSourceDefaultAccountRole{}

// AddIdentitySourcesRequestUserSourceDefaultAccountRole struct for AddIdentitySourcesRequestUserSourceDefaultAccountRole
type AddIdentitySourcesRequestUserSourceDefaultAccountRole struct {
	// Default `Role ID`
	Id                   int64                  `json:"id"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddIdentitySourcesRequestUserSourceDefaultAccountRole AddIdentitySourcesRequestUserSourceDefaultAccountRole

// NewAddIdentitySourcesRequestUserSourceDefaultAccountRole instantiates a new AddIdentitySourcesRequestUserSourceDefaultAccountRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIdentitySourcesRequestUserSourceDefaultAccountRole(id int64) *AddIdentitySourcesRequestUserSourceDefaultAccountRole {
	this := AddIdentitySourcesRequestUserSourceDefaultAccountRole{}
	this.Id = id
	return &this
}

// NewAddIdentitySourcesRequestUserSourceDefaultAccountRoleWithDefaults instantiates a new AddIdentitySourcesRequestUserSourceDefaultAccountRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIdentitySourcesRequestUserSourceDefaultAccountRoleWithDefaults() *AddIdentitySourcesRequestUserSourceDefaultAccountRole {
	this := AddIdentitySourcesRequestUserSourceDefaultAccountRole{}
	return &this
}

// GetId returns the Id field value
func (o *AddIdentitySourcesRequestUserSourceDefaultAccountRole) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddIdentitySourcesRequestUserSourceDefaultAccountRole) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddIdentitySourcesRequestUserSourceDefaultAccountRole) SetId(v int64) {
	o.Id = v
}

func (o AddIdentitySourcesRequestUserSourceDefaultAccountRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddIdentitySourcesRequestUserSourceDefaultAccountRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddIdentitySourcesRequestUserSourceDefaultAccountRole) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

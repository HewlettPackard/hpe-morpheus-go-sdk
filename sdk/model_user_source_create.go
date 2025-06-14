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

// checks if the UserSourceCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSourceCreate{}

// UserSourceCreate struct for UserSourceCreate
type UserSourceCreate struct {
	UserSource           AddIdentitySourcesRequestUserSource `json:"userSource"`
	AdditionalProperties map[string]interface{}              `json:",remain"`
}

type _UserSourceCreate UserSourceCreate

// NewUserSourceCreate instantiates a new UserSourceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSourceCreate(userSource AddIdentitySourcesRequestUserSource) *UserSourceCreate {
	this := UserSourceCreate{}
	this.UserSource = userSource
	return &this
}

// NewUserSourceCreateWithDefaults instantiates a new UserSourceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSourceCreateWithDefaults() *UserSourceCreate {
	this := UserSourceCreate{}
	return &this
}

// GetUserSource returns the UserSource field value
func (o *UserSourceCreate) GetUserSource() AddIdentitySourcesRequestUserSource {
	if o == nil {
		var ret AddIdentitySourcesRequestUserSource
		return ret
	}

	return o.UserSource
}

// GetUserSourceOk returns a tuple with the UserSource field value
// and a boolean to check if the value has been set.
func (o *UserSourceCreate) GetUserSourceOk() (*AddIdentitySourcesRequestUserSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserSource, true
}

// SetUserSource sets field value
func (o *UserSourceCreate) SetUserSource(v AddIdentitySourcesRequestUserSource) {
	o.UserSource = v
}

func (o UserSourceCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSourceCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userSource"] = o.UserSource

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UserSourceCreate) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

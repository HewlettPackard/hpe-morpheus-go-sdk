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

// checks if the AddWikiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddWikiRequest{}

// AddWikiRequest struct for AddWikiRequest
type AddWikiRequest struct {
	Page                 AddWikiRequestPage     `json:"page"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddWikiRequest AddWikiRequest

// NewAddWikiRequest instantiates a new AddWikiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddWikiRequest(page AddWikiRequestPage) *AddWikiRequest {
	this := AddWikiRequest{}
	this.Page = page
	return &this
}

// NewAddWikiRequestWithDefaults instantiates a new AddWikiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddWikiRequestWithDefaults() *AddWikiRequest {
	this := AddWikiRequest{}
	return &this
}

// GetPage returns the Page field value
func (o *AddWikiRequest) GetPage() AddWikiRequestPage {
	if o == nil {
		var ret AddWikiRequestPage
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *AddWikiRequest) GetPageOk() (*AddWikiRequestPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *AddWikiRequest) SetPage(v AddWikiRequestPage) {
	o.Page = v
}

func (o AddWikiRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddWikiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["page"] = o.Page

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddWikiRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

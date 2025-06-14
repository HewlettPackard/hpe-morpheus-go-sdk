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

// checks if the AddClusterLayoutsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddClusterLayoutsRequest{}

// AddClusterLayoutsRequest struct for AddClusterLayoutsRequest
type AddClusterLayoutsRequest struct {
	Layout               *AddClusterLayoutsRequestLayout `json:"layout,omitempty"`
	AdditionalProperties map[string]interface{}          `json:",remain"`
}

type _AddClusterLayoutsRequest AddClusterLayoutsRequest

// NewAddClusterLayoutsRequest instantiates a new AddClusterLayoutsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddClusterLayoutsRequest() *AddClusterLayoutsRequest {
	this := AddClusterLayoutsRequest{}
	return &this
}

// NewAddClusterLayoutsRequestWithDefaults instantiates a new AddClusterLayoutsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddClusterLayoutsRequestWithDefaults() *AddClusterLayoutsRequest {
	this := AddClusterLayoutsRequest{}
	return &this
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *AddClusterLayoutsRequest) GetLayout() AddClusterLayoutsRequestLayout {
	if o == nil || IsNil(o.Layout) {
		var ret AddClusterLayoutsRequestLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddClusterLayoutsRequest) GetLayoutOk() (*AddClusterLayoutsRequestLayout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// IsSetLayout returns a boolean if a field has been set.
func (o *AddClusterLayoutsRequest) IsSetLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given AddClusterLayoutsRequestLayout and assigns it to the Layout field.
func (o *AddClusterLayoutsRequest) SetLayout(v AddClusterLayoutsRequestLayout) {
	o.Layout = &v
}

func (o AddClusterLayoutsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddClusterLayoutsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddClusterLayoutsRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

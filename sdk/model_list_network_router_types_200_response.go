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

// checks if the ListNetworkRouterTypes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListNetworkRouterTypes200Response{}

// ListNetworkRouterTypes200Response struct for ListNetworkRouterTypes200Response
type ListNetworkRouterTypes200Response struct {
	NetworkRouterTypes   []ListNetworkRouterTypes200ResponseNetworkRouterTypesInner `json:"networkRouterTypes,omitempty"`
	AdditionalProperties map[string]interface{}                                     `json:",remain"`
}

type _ListNetworkRouterTypes200Response ListNetworkRouterTypes200Response

// NewListNetworkRouterTypes200Response instantiates a new ListNetworkRouterTypes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListNetworkRouterTypes200Response() *ListNetworkRouterTypes200Response {
	this := ListNetworkRouterTypes200Response{}
	return &this
}

// NewListNetworkRouterTypes200ResponseWithDefaults instantiates a new ListNetworkRouterTypes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListNetworkRouterTypes200ResponseWithDefaults() *ListNetworkRouterTypes200Response {
	this := ListNetworkRouterTypes200Response{}
	return &this
}

// GetNetworkRouterTypes returns the NetworkRouterTypes field value if set, zero value otherwise.
func (o *ListNetworkRouterTypes200Response) GetNetworkRouterTypes() []ListNetworkRouterTypes200ResponseNetworkRouterTypesInner {
	if o == nil || IsNil(o.NetworkRouterTypes) {
		var ret []ListNetworkRouterTypes200ResponseNetworkRouterTypesInner
		return ret
	}
	return o.NetworkRouterTypes
}

// GetNetworkRouterTypesOk returns a tuple with the NetworkRouterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkRouterTypes200Response) GetNetworkRouterTypesOk() ([]ListNetworkRouterTypes200ResponseNetworkRouterTypesInner, bool) {
	if o == nil || IsNil(o.NetworkRouterTypes) {
		return nil, false
	}
	return o.NetworkRouterTypes, true
}

// IsSetNetworkRouterTypes returns a boolean if a field has been set.
func (o *ListNetworkRouterTypes200Response) IsSetNetworkRouterTypes() bool {
	if o != nil && !IsNil(o.NetworkRouterTypes) {
		return true
	}

	return false
}

// SetNetworkRouterTypes gets a reference to the given []ListNetworkRouterTypes200ResponseNetworkRouterTypesInner and assigns it to the NetworkRouterTypes field.
func (o *ListNetworkRouterTypes200Response) SetNetworkRouterTypes(v []ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) {
	o.NetworkRouterTypes = v
}

func (o ListNetworkRouterTypes200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListNetworkRouterTypes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkRouterTypes) {
		toSerialize["networkRouterTypes"] = o.NetworkRouterTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListNetworkRouterTypes200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

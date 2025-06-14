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

// checks if the ListHealth200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHealth200Response{}

// ListHealth200Response struct for ListHealth200Response
type ListHealth200Response struct {
	Health               *ListHealth200ResponseAllOfHealth `json:"health,omitempty"`
	Meta                 *ListActivity200ResponseAllOfMeta `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}            `json:",remain"`
}

type _ListHealth200Response ListHealth200Response

// NewListHealth200Response instantiates a new ListHealth200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHealth200Response() *ListHealth200Response {
	this := ListHealth200Response{}
	return &this
}

// NewListHealth200ResponseWithDefaults instantiates a new ListHealth200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHealth200ResponseWithDefaults() *ListHealth200Response {
	this := ListHealth200Response{}
	return &this
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *ListHealth200Response) GetHealth() ListHealth200ResponseAllOfHealth {
	if o == nil || IsNil(o.Health) {
		var ret ListHealth200ResponseAllOfHealth
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200Response) GetHealthOk() (*ListHealth200ResponseAllOfHealth, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// IsSetHealth returns a boolean if a field has been set.
func (o *ListHealth200Response) IsSetHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given ListHealth200ResponseAllOfHealth and assigns it to the Health field.
func (o *ListHealth200Response) SetHealth(v ListHealth200ResponseAllOfHealth) {
	o.Health = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListHealth200Response) GetMeta() ListActivity200ResponseAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListActivity200ResponseAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// IsSetMeta returns a boolean if a field has been set.
func (o *ListHealth200Response) IsSetMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListActivity200ResponseAllOfMeta and assigns it to the Meta field.
func (o *ListHealth200Response) SetMeta(v ListActivity200ResponseAllOfMeta) {
	o.Meta = &v
}

func (o ListHealth200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHealth200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListHealth200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

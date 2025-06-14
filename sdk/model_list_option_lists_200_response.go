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

// checks if the ListOptionLists200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOptionLists200Response{}

// ListOptionLists200Response struct for ListOptionLists200Response
type ListOptionLists200Response struct {
	OptionTypes          []ListOptionLists200ResponseAllOfOptionTypesInner `json:"optionTypes,omitempty"`
	Meta                 *ListActivity200ResponseAllOfMeta                 `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}                            `json:",remain"`
}

type _ListOptionLists200Response ListOptionLists200Response

// NewListOptionLists200Response instantiates a new ListOptionLists200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOptionLists200Response() *ListOptionLists200Response {
	this := ListOptionLists200Response{}
	return &this
}

// NewListOptionLists200ResponseWithDefaults instantiates a new ListOptionLists200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOptionLists200ResponseWithDefaults() *ListOptionLists200Response {
	this := ListOptionLists200Response{}
	return &this
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *ListOptionLists200Response) GetOptionTypes() []ListOptionLists200ResponseAllOfOptionTypesInner {
	if o == nil || IsNil(o.OptionTypes) {
		var ret []ListOptionLists200ResponseAllOfOptionTypesInner
		return ret
	}
	return o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOptionLists200Response) GetOptionTypesOk() ([]ListOptionLists200ResponseAllOfOptionTypesInner, bool) {
	if o == nil || IsNil(o.OptionTypes) {
		return nil, false
	}
	return o.OptionTypes, true
}

// IsSetOptionTypes returns a boolean if a field has been set.
func (o *ListOptionLists200Response) IsSetOptionTypes() bool {
	if o != nil && !IsNil(o.OptionTypes) {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []ListOptionLists200ResponseAllOfOptionTypesInner and assigns it to the OptionTypes field.
func (o *ListOptionLists200Response) SetOptionTypes(v []ListOptionLists200ResponseAllOfOptionTypesInner) {
	o.OptionTypes = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListOptionLists200Response) GetMeta() ListActivity200ResponseAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListActivity200ResponseAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOptionLists200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// IsSetMeta returns a boolean if a field has been set.
func (o *ListOptionLists200Response) IsSetMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListActivity200ResponseAllOfMeta and assigns it to the Meta field.
func (o *ListOptionLists200Response) SetMeta(v ListActivity200ResponseAllOfMeta) {
	o.Meta = &v
}

func (o ListOptionLists200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOptionLists200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptionTypes) {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListOptionLists200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

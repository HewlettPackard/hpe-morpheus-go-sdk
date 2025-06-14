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

// checks if the GetOptionForm200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOptionForm200Response{}

// GetOptionForm200Response struct for GetOptionForm200Response
type GetOptionForm200Response struct {
	OptionTypes          []ListOptionForms200ResponseAllOfOptionTypesInner `json:"optionTypes,omitempty"`
	AdditionalProperties map[string]interface{}                            `json:",remain"`
}

type _GetOptionForm200Response GetOptionForm200Response

// NewGetOptionForm200Response instantiates a new GetOptionForm200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOptionForm200Response() *GetOptionForm200Response {
	this := GetOptionForm200Response{}
	return &this
}

// NewGetOptionForm200ResponseWithDefaults instantiates a new GetOptionForm200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOptionForm200ResponseWithDefaults() *GetOptionForm200Response {
	this := GetOptionForm200Response{}
	return &this
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *GetOptionForm200Response) GetOptionTypes() []ListOptionForms200ResponseAllOfOptionTypesInner {
	if o == nil || IsNil(o.OptionTypes) {
		var ret []ListOptionForms200ResponseAllOfOptionTypesInner
		return ret
	}
	return o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionForm200Response) GetOptionTypesOk() ([]ListOptionForms200ResponseAllOfOptionTypesInner, bool) {
	if o == nil || IsNil(o.OptionTypes) {
		return nil, false
	}
	return o.OptionTypes, true
}

// IsSetOptionTypes returns a boolean if a field has been set.
func (o *GetOptionForm200Response) IsSetOptionTypes() bool {
	if o != nil && !IsNil(o.OptionTypes) {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []ListOptionForms200ResponseAllOfOptionTypesInner and assigns it to the OptionTypes field.
func (o *GetOptionForm200Response) SetOptionTypes(v []ListOptionForms200ResponseAllOfOptionTypesInner) {
	o.OptionTypes = v
}

func (o GetOptionForm200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOptionForm200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptionTypes) {
		toSerialize["optionTypes"] = o.OptionTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetOptionForm200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

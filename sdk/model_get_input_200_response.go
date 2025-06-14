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

// checks if the GetInput200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInput200Response{}

// GetInput200Response struct for GetInput200Response
type GetInput200Response struct {
	OptionType           *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner `json:"optionType,omitempty"`
	AdditionalProperties map[string]interface{}                                                     `json:",remain"`
}

type _GetInput200Response GetInput200Response

// NewGetInput200Response instantiates a new GetInput200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInput200Response() *GetInput200Response {
	this := GetInput200Response{}
	return &this
}

// NewGetInput200ResponseWithDefaults instantiates a new GetInput200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInput200ResponseWithDefaults() *GetInput200Response {
	this := GetInput200Response{}
	return &this
}

// GetOptionType returns the OptionType field value if set, zero value otherwise.
func (o *GetInput200Response) GetOptionType() ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner {
	if o == nil || IsNil(o.OptionType) {
		var ret ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner
		return ret
	}
	return *o.OptionType
}

// GetOptionTypeOk returns a tuple with the OptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInput200Response) GetOptionTypeOk() (*ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool) {
	if o == nil || IsNil(o.OptionType) {
		return nil, false
	}
	return o.OptionType, true
}

// IsSetOptionType returns a boolean if a field has been set.
func (o *GetInput200Response) IsSetOptionType() bool {
	if o != nil && !IsNil(o.OptionType) {
		return true
	}

	return false
}

// SetOptionType gets a reference to the given ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner and assigns it to the OptionType field.
func (o *GetInput200Response) SetOptionType(v ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner) {
	o.OptionType = &v
}

func (o GetInput200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInput200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptionType) {
		toSerialize["optionType"] = o.OptionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetInput200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

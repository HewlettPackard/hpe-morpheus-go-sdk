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

// checks if the GetSpecTemplate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSpecTemplate200Response{}

// GetSpecTemplate200Response struct for GetSpecTemplate200Response
type GetSpecTemplate200Response struct {
	SpecTemplate         *ListSpecTemplates200ResponseAllOfSpecTemplatesInner `json:"specTemplate,omitempty"`
	AdditionalProperties map[string]interface{}                               `json:",remain"`
}

type _GetSpecTemplate200Response GetSpecTemplate200Response

// NewGetSpecTemplate200Response instantiates a new GetSpecTemplate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSpecTemplate200Response() *GetSpecTemplate200Response {
	this := GetSpecTemplate200Response{}
	return &this
}

// NewGetSpecTemplate200ResponseWithDefaults instantiates a new GetSpecTemplate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSpecTemplate200ResponseWithDefaults() *GetSpecTemplate200Response {
	this := GetSpecTemplate200Response{}
	return &this
}

// GetSpecTemplate returns the SpecTemplate field value if set, zero value otherwise.
func (o *GetSpecTemplate200Response) GetSpecTemplate() ListSpecTemplates200ResponseAllOfSpecTemplatesInner {
	if o == nil || IsNil(o.SpecTemplate) {
		var ret ListSpecTemplates200ResponseAllOfSpecTemplatesInner
		return ret
	}
	return *o.SpecTemplate
}

// GetSpecTemplateOk returns a tuple with the SpecTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSpecTemplate200Response) GetSpecTemplateOk() (*ListSpecTemplates200ResponseAllOfSpecTemplatesInner, bool) {
	if o == nil || IsNil(o.SpecTemplate) {
		return nil, false
	}
	return o.SpecTemplate, true
}

// IsSetSpecTemplate returns a boolean if a field has been set.
func (o *GetSpecTemplate200Response) IsSetSpecTemplate() bool {
	if o != nil && !IsNil(o.SpecTemplate) {
		return true
	}

	return false
}

// SetSpecTemplate gets a reference to the given ListSpecTemplates200ResponseAllOfSpecTemplatesInner and assigns it to the SpecTemplate field.
func (o *GetSpecTemplate200Response) SetSpecTemplate(v ListSpecTemplates200ResponseAllOfSpecTemplatesInner) {
	o.SpecTemplate = &v
}

func (o GetSpecTemplate200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSpecTemplate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpecTemplate) {
		toSerialize["specTemplate"] = o.SpecTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetSpecTemplate200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

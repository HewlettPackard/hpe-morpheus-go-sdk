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

// checks if the AddOptionListRequestOptionTypeListConfigSourceHeadersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOptionListRequestOptionTypeListConfigSourceHeadersInner{}

// AddOptionListRequestOptionTypeListConfigSourceHeadersInner struct for AddOptionListRequestOptionTypeListConfigSourceHeadersInner
type AddOptionListRequestOptionTypeListConfigSourceHeadersInner struct {
	// Header name
	Name string `json:"name"`
	// Header value
	Value *string `json:"value,omitempty"`
	// Can be used to enable / disable masking of value
	Masked               *bool                  `json:"masked,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddOptionListRequestOptionTypeListConfigSourceHeadersInner AddOptionListRequestOptionTypeListConfigSourceHeadersInner

// NewAddOptionListRequestOptionTypeListConfigSourceHeadersInner instantiates a new AddOptionListRequestOptionTypeListConfigSourceHeadersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOptionListRequestOptionTypeListConfigSourceHeadersInner(name string) *AddOptionListRequestOptionTypeListConfigSourceHeadersInner {
	this := AddOptionListRequestOptionTypeListConfigSourceHeadersInner{}
	this.Name = name
	var masked bool = false
	this.Masked = &masked
	return &this
}

// NewAddOptionListRequestOptionTypeListConfigSourceHeadersInnerWithDefaults instantiates a new AddOptionListRequestOptionTypeListConfigSourceHeadersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOptionListRequestOptionTypeListConfigSourceHeadersInnerWithDefaults() *AddOptionListRequestOptionTypeListConfigSourceHeadersInner {
	this := AddOptionListRequestOptionTypeListConfigSourceHeadersInner{}
	var masked bool = false
	this.Masked = &masked
	return &this
}

// GetName returns the Name field value
func (o *AddOptionListRequestOptionTypeListConfigSourceHeadersInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddOptionListRequestOptionTypeListConfigSourceHeadersInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddOptionListRequestOptionTypeListConfigSourceHeadersInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AddOptionListRequestOptionTypeListConfigSourceHeadersInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOptionListRequestOptionTypeListConfigSourceHeadersInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// IsSetValue returns a boolean if a field has been set.
func (o *AddOptionListRequestOptionTypeListConfigSourceHeadersInner) IsSetValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AddOptionListRequestOptionTypeListConfigSourceHeadersInner) SetValue(v string) {
	o.Value = &v
}

// GetMasked returns the Masked field value if set, zero value otherwise.
func (o *AddOptionListRequestOptionTypeListConfigSourceHeadersInner) GetMasked() bool {
	if o == nil || IsNil(o.Masked) {
		var ret bool
		return ret
	}
	return *o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOptionListRequestOptionTypeListConfigSourceHeadersInner) GetMaskedOk() (*bool, bool) {
	if o == nil || IsNil(o.Masked) {
		return nil, false
	}
	return o.Masked, true
}

// IsSetMasked returns a boolean if a field has been set.
func (o *AddOptionListRequestOptionTypeListConfigSourceHeadersInner) IsSetMasked() bool {
	if o != nil && !IsNil(o.Masked) {
		return true
	}

	return false
}

// SetMasked gets a reference to the given bool and assigns it to the Masked field.
func (o *AddOptionListRequestOptionTypeListConfigSourceHeadersInner) SetMasked(v bool) {
	o.Masked = &v
}

func (o AddOptionListRequestOptionTypeListConfigSourceHeadersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOptionListRequestOptionTypeListConfigSourceHeadersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Masked) {
		toSerialize["masked"] = o.Masked
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddOptionListRequestOptionTypeListConfigSourceHeadersInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

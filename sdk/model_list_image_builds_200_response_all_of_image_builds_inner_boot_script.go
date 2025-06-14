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

// checks if the ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript{}

// ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript struct for ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript
type ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript struct {
	Id                   *int64                 `json:"id,omitempty"`
	FileName             *string                `json:"fileName,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript

// NewListImageBuilds200ResponseAllOfImageBuildsInnerBootScript instantiates a new ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListImageBuilds200ResponseAllOfImageBuildsInnerBootScript() *ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript {
	this := ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript{}
	return &this
}

// NewListImageBuilds200ResponseAllOfImageBuildsInnerBootScriptWithDefaults instantiates a new ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListImageBuilds200ResponseAllOfImageBuildsInnerBootScriptWithDefaults() *ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript {
	this := ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript) SetId(v int64) {
	o.Id = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// IsSetFileName returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript) IsSetFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript) SetFileName(v string) {
	o.FileName = &v
}

func (o ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

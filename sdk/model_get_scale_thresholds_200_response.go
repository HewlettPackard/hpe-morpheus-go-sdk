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

// checks if the GetScaleThresholds200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetScaleThresholds200Response{}

// GetScaleThresholds200Response struct for GetScaleThresholds200Response
type GetScaleThresholds200Response struct {
	ScaleThreshold       *ListScaleThresholds200ResponseAllOfScaleThresholdsInner `json:"scaleThreshold,omitempty"`
	AdditionalProperties map[string]interface{}                                   `json:",remain"`
}

type _GetScaleThresholds200Response GetScaleThresholds200Response

// NewGetScaleThresholds200Response instantiates a new GetScaleThresholds200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetScaleThresholds200Response() *GetScaleThresholds200Response {
	this := GetScaleThresholds200Response{}
	return &this
}

// NewGetScaleThresholds200ResponseWithDefaults instantiates a new GetScaleThresholds200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetScaleThresholds200ResponseWithDefaults() *GetScaleThresholds200Response {
	this := GetScaleThresholds200Response{}
	return &this
}

// GetScaleThreshold returns the ScaleThreshold field value if set, zero value otherwise.
func (o *GetScaleThresholds200Response) GetScaleThreshold() ListScaleThresholds200ResponseAllOfScaleThresholdsInner {
	if o == nil || IsNil(o.ScaleThreshold) {
		var ret ListScaleThresholds200ResponseAllOfScaleThresholdsInner
		return ret
	}
	return *o.ScaleThreshold
}

// GetScaleThresholdOk returns a tuple with the ScaleThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScaleThresholds200Response) GetScaleThresholdOk() (*ListScaleThresholds200ResponseAllOfScaleThresholdsInner, bool) {
	if o == nil || IsNil(o.ScaleThreshold) {
		return nil, false
	}
	return o.ScaleThreshold, true
}

// IsSetScaleThreshold returns a boolean if a field has been set.
func (o *GetScaleThresholds200Response) IsSetScaleThreshold() bool {
	if o != nil && !IsNil(o.ScaleThreshold) {
		return true
	}

	return false
}

// SetScaleThreshold gets a reference to the given ListScaleThresholds200ResponseAllOfScaleThresholdsInner and assigns it to the ScaleThreshold field.
func (o *GetScaleThresholds200Response) SetScaleThreshold(v ListScaleThresholds200ResponseAllOfScaleThresholdsInner) {
	o.ScaleThreshold = &v
}

func (o GetScaleThresholds200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetScaleThresholds200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScaleThreshold) {
		toSerialize["scaleThreshold"] = o.ScaleThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetScaleThresholds200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

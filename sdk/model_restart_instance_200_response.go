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

// checks if the RestartInstance200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestartInstance200Response{}

// RestartInstance200Response struct for RestartInstance200Response
type RestartInstance200Response struct {
	Results              map[string]interface{} `json:"results,omitempty"`
	Success              *bool                  `json:"success,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _RestartInstance200Response RestartInstance200Response

// NewRestartInstance200Response instantiates a new RestartInstance200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestartInstance200Response() *RestartInstance200Response {
	this := RestartInstance200Response{}
	return &this
}

// NewRestartInstance200ResponseWithDefaults instantiates a new RestartInstance200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestartInstance200ResponseWithDefaults() *RestartInstance200Response {
	this := RestartInstance200Response{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *RestartInstance200Response) GetResults() map[string]interface{} {
	if o == nil || IsNil(o.Results) {
		var ret map[string]interface{}
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestartInstance200Response) GetResultsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Results) {
		return map[string]interface{}{}, false
	}
	return o.Results, true
}

// IsSetResults returns a boolean if a field has been set.
func (o *RestartInstance200Response) IsSetResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given map[string]interface{} and assigns it to the Results field.
func (o *RestartInstance200Response) SetResults(v map[string]interface{}) {
	o.Results = v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *RestartInstance200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestartInstance200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *RestartInstance200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *RestartInstance200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o RestartInstance200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestartInstance200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *RestartInstance200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the ValidateAppState200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateAppState200Response{}

// ValidateAppState200Response struct for ValidateAppState200Response
type ValidateAppState200Response struct {
	ExecutionId          *string                `json:"executionId,omitempty"`
	Success              *bool                  `json:"success,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ValidateAppState200Response ValidateAppState200Response

// NewValidateAppState200Response instantiates a new ValidateAppState200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateAppState200Response() *ValidateAppState200Response {
	this := ValidateAppState200Response{}
	return &this
}

// NewValidateAppState200ResponseWithDefaults instantiates a new ValidateAppState200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateAppState200ResponseWithDefaults() *ValidateAppState200Response {
	this := ValidateAppState200Response{}
	return &this
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *ValidateAppState200Response) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId) {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateAppState200Response) GetExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// IsSetExecutionId returns a boolean if a field has been set.
func (o *ValidateAppState200Response) IsSetExecutionId() bool {
	if o != nil && !IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *ValidateAppState200Response) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ValidateAppState200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateAppState200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *ValidateAppState200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ValidateAppState200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o ValidateAppState200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateAppState200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExecutionId) {
		toSerialize["executionId"] = o.ExecutionId
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ValidateAppState200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

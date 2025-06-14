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

// checks if the AddEnvironments200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddEnvironments200Response{}

// AddEnvironments200Response struct for AddEnvironments200Response
type AddEnvironments200Response struct {
	Environment          *ListEnvironments200ResponseAllOfEnvironmentsInner `json:"environment,omitempty"`
	Success              *bool                                              `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}                             `json:",remain"`
}

type _AddEnvironments200Response AddEnvironments200Response

// NewAddEnvironments200Response instantiates a new AddEnvironments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddEnvironments200Response() *AddEnvironments200Response {
	this := AddEnvironments200Response{}
	return &this
}

// NewAddEnvironments200ResponseWithDefaults instantiates a new AddEnvironments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddEnvironments200ResponseWithDefaults() *AddEnvironments200Response {
	this := AddEnvironments200Response{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *AddEnvironments200Response) GetEnvironment() ListEnvironments200ResponseAllOfEnvironmentsInner {
	if o == nil || IsNil(o.Environment) {
		var ret ListEnvironments200ResponseAllOfEnvironmentsInner
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEnvironments200Response) GetEnvironmentOk() (*ListEnvironments200ResponseAllOfEnvironmentsInner, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// IsSetEnvironment returns a boolean if a field has been set.
func (o *AddEnvironments200Response) IsSetEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ListEnvironments200ResponseAllOfEnvironmentsInner and assigns it to the Environment field.
func (o *AddEnvironments200Response) SetEnvironment(v ListEnvironments200ResponseAllOfEnvironmentsInner) {
	o.Environment = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AddEnvironments200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEnvironments200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *AddEnvironments200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AddEnvironments200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o AddEnvironments200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddEnvironments200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddEnvironments200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the UpdateDeployment200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeployment200Response{}

// UpdateDeployment200Response struct for UpdateDeployment200Response
type UpdateDeployment200Response struct {
	Deployment           *GetDeployment200ResponseDeployment `json:"deployment,omitempty"`
	Success              *bool                               `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}              `json:",remain"`
}

type _UpdateDeployment200Response UpdateDeployment200Response

// NewUpdateDeployment200Response instantiates a new UpdateDeployment200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeployment200Response() *UpdateDeployment200Response {
	this := UpdateDeployment200Response{}
	return &this
}

// NewUpdateDeployment200ResponseWithDefaults instantiates a new UpdateDeployment200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeployment200ResponseWithDefaults() *UpdateDeployment200Response {
	this := UpdateDeployment200Response{}
	return &this
}

// GetDeployment returns the Deployment field value if set, zero value otherwise.
func (o *UpdateDeployment200Response) GetDeployment() GetDeployment200ResponseDeployment {
	if o == nil || IsNil(o.Deployment) {
		var ret GetDeployment200ResponseDeployment
		return ret
	}
	return *o.Deployment
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeployment200Response) GetDeploymentOk() (*GetDeployment200ResponseDeployment, bool) {
	if o == nil || IsNil(o.Deployment) {
		return nil, false
	}
	return o.Deployment, true
}

// IsSetDeployment returns a boolean if a field has been set.
func (o *UpdateDeployment200Response) IsSetDeployment() bool {
	if o != nil && !IsNil(o.Deployment) {
		return true
	}

	return false
}

// SetDeployment gets a reference to the given GetDeployment200ResponseDeployment and assigns it to the Deployment field.
func (o *UpdateDeployment200Response) SetDeployment(v GetDeployment200ResponseDeployment) {
	o.Deployment = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UpdateDeployment200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeployment200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *UpdateDeployment200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UpdateDeployment200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o UpdateDeployment200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeployment200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deployment) {
		toSerialize["deployment"] = o.Deployment
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateDeployment200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

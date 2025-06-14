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

// checks if the UpdateFileTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFileTemplateRequest{}

// UpdateFileTemplateRequest struct for UpdateFileTemplateRequest
type UpdateFileTemplateRequest struct {
	ContainerTemplate    *UpdateFileTemplateRequestContainerTemplate `json:"containerTemplate,omitempty"`
	AdditionalProperties map[string]interface{}                      `json:",remain"`
}

type _UpdateFileTemplateRequest UpdateFileTemplateRequest

// NewUpdateFileTemplateRequest instantiates a new UpdateFileTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFileTemplateRequest() *UpdateFileTemplateRequest {
	this := UpdateFileTemplateRequest{}
	return &this
}

// NewUpdateFileTemplateRequestWithDefaults instantiates a new UpdateFileTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFileTemplateRequestWithDefaults() *UpdateFileTemplateRequest {
	this := UpdateFileTemplateRequest{}
	return &this
}

// GetContainerTemplate returns the ContainerTemplate field value if set, zero value otherwise.
func (o *UpdateFileTemplateRequest) GetContainerTemplate() UpdateFileTemplateRequestContainerTemplate {
	if o == nil || IsNil(o.ContainerTemplate) {
		var ret UpdateFileTemplateRequestContainerTemplate
		return ret
	}
	return *o.ContainerTemplate
}

// GetContainerTemplateOk returns a tuple with the ContainerTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileTemplateRequest) GetContainerTemplateOk() (*UpdateFileTemplateRequestContainerTemplate, bool) {
	if o == nil || IsNil(o.ContainerTemplate) {
		return nil, false
	}
	return o.ContainerTemplate, true
}

// IsSetContainerTemplate returns a boolean if a field has been set.
func (o *UpdateFileTemplateRequest) IsSetContainerTemplate() bool {
	if o != nil && !IsNil(o.ContainerTemplate) {
		return true
	}

	return false
}

// SetContainerTemplate gets a reference to the given UpdateFileTemplateRequestContainerTemplate and assigns it to the ContainerTemplate field.
func (o *UpdateFileTemplateRequest) SetContainerTemplate(v UpdateFileTemplateRequestContainerTemplate) {
	o.ContainerTemplate = &v
}

func (o UpdateFileTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFileTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerTemplate) {
		toSerialize["containerTemplate"] = o.ContainerTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateFileTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

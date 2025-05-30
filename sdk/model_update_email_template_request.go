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

// checks if the UpdateEmailTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEmailTemplateRequest{}

// UpdateEmailTemplateRequest struct for UpdateEmailTemplateRequest
type UpdateEmailTemplateRequest struct {
	Policy *ListEmailTemplates200ResponseAllOfEmailTemplatesInner `json:"policy,omitempty"`
}

// NewUpdateEmailTemplateRequest instantiates a new UpdateEmailTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEmailTemplateRequest() *UpdateEmailTemplateRequest {
	this := UpdateEmailTemplateRequest{}
	return &this
}

// NewUpdateEmailTemplateRequestWithDefaults instantiates a new UpdateEmailTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEmailTemplateRequestWithDefaults() *UpdateEmailTemplateRequest {
	this := UpdateEmailTemplateRequest{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *UpdateEmailTemplateRequest) GetPolicy() ListEmailTemplates200ResponseAllOfEmailTemplatesInner {
	if o == nil || IsNil(o.Policy) {
		var ret ListEmailTemplates200ResponseAllOfEmailTemplatesInner
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEmailTemplateRequest) GetPolicyOk() (*ListEmailTemplates200ResponseAllOfEmailTemplatesInner, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// IsSetPolicy returns a boolean if a field has been set.
func (o *UpdateEmailTemplateRequest) IsSetPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given ListEmailTemplates200ResponseAllOfEmailTemplatesInner and assigns it to the Policy field.
func (o *UpdateEmailTemplateRequest) SetPolicy(v ListEmailTemplates200ResponseAllOfEmailTemplatesInner) {
	o.Policy = &v
}

func (o UpdateEmailTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEmailTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	return toSerialize, nil
}

type NullableUpdateEmailTemplateRequest struct {
	value *UpdateEmailTemplateRequest
	isSet bool
}

func (v NullableUpdateEmailTemplateRequest) Get() *UpdateEmailTemplateRequest {
	return v.value
}

func (v *NullableUpdateEmailTemplateRequest) Set(val *UpdateEmailTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEmailTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEmailTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEmailTemplateRequest(val *UpdateEmailTemplateRequest) *NullableUpdateEmailTemplateRequest {
	return &NullableUpdateEmailTemplateRequest{value: val, isSet: true}
}

func (v NullableUpdateEmailTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEmailTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



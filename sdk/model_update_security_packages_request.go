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

// checks if the UpdateSecurityPackagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSecurityPackagesRequest{}

// UpdateSecurityPackagesRequest struct for UpdateSecurityPackagesRequest
type UpdateSecurityPackagesRequest struct {
	SecurityPackage      UpdateSecurityPackagesRequestSecurityPackage `json:"securityPackage"`
	AdditionalProperties map[string]interface{}                       `json:",remain"`
}

type _UpdateSecurityPackagesRequest UpdateSecurityPackagesRequest

// NewUpdateSecurityPackagesRequest instantiates a new UpdateSecurityPackagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSecurityPackagesRequest(securityPackage UpdateSecurityPackagesRequestSecurityPackage) *UpdateSecurityPackagesRequest {
	this := UpdateSecurityPackagesRequest{}
	this.SecurityPackage = securityPackage
	return &this
}

// NewUpdateSecurityPackagesRequestWithDefaults instantiates a new UpdateSecurityPackagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSecurityPackagesRequestWithDefaults() *UpdateSecurityPackagesRequest {
	this := UpdateSecurityPackagesRequest{}
	return &this
}

// GetSecurityPackage returns the SecurityPackage field value
func (o *UpdateSecurityPackagesRequest) GetSecurityPackage() UpdateSecurityPackagesRequestSecurityPackage {
	if o == nil {
		var ret UpdateSecurityPackagesRequestSecurityPackage
		return ret
	}

	return o.SecurityPackage
}

// GetSecurityPackageOk returns a tuple with the SecurityPackage field value
// and a boolean to check if the value has been set.
func (o *UpdateSecurityPackagesRequest) GetSecurityPackageOk() (*UpdateSecurityPackagesRequestSecurityPackage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityPackage, true
}

// SetSecurityPackage sets field value
func (o *UpdateSecurityPackagesRequest) SetSecurityPackage(v UpdateSecurityPackagesRequestSecurityPackage) {
	o.SecurityPackage = v
}

func (o UpdateSecurityPackagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSecurityPackagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["securityPackage"] = o.SecurityPackage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateSecurityPackagesRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

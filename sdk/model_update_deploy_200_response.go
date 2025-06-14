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

// checks if the UpdateDeploy200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeploy200Response{}

// UpdateDeploy200Response struct for UpdateDeploy200Response
type UpdateDeploy200Response struct {
	AppDeploy            *ListDeploys200ResponseAllOfAppDeploysInner `json:"appDeploy,omitempty"`
	AdditionalProperties map[string]interface{}                      `json:",remain"`
}

type _UpdateDeploy200Response UpdateDeploy200Response

// NewUpdateDeploy200Response instantiates a new UpdateDeploy200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeploy200Response() *UpdateDeploy200Response {
	this := UpdateDeploy200Response{}
	return &this
}

// NewUpdateDeploy200ResponseWithDefaults instantiates a new UpdateDeploy200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeploy200ResponseWithDefaults() *UpdateDeploy200Response {
	this := UpdateDeploy200Response{}
	return &this
}

// GetAppDeploy returns the AppDeploy field value if set, zero value otherwise.
func (o *UpdateDeploy200Response) GetAppDeploy() ListDeploys200ResponseAllOfAppDeploysInner {
	if o == nil || IsNil(o.AppDeploy) {
		var ret ListDeploys200ResponseAllOfAppDeploysInner
		return ret
	}
	return *o.AppDeploy
}

// GetAppDeployOk returns a tuple with the AppDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeploy200Response) GetAppDeployOk() (*ListDeploys200ResponseAllOfAppDeploysInner, bool) {
	if o == nil || IsNil(o.AppDeploy) {
		return nil, false
	}
	return o.AppDeploy, true
}

// IsSetAppDeploy returns a boolean if a field has been set.
func (o *UpdateDeploy200Response) IsSetAppDeploy() bool {
	if o != nil && !IsNil(o.AppDeploy) {
		return true
	}

	return false
}

// SetAppDeploy gets a reference to the given ListDeploys200ResponseAllOfAppDeploysInner and assigns it to the AppDeploy field.
func (o *UpdateDeploy200Response) SetAppDeploy(v ListDeploys200ResponseAllOfAppDeploysInner) {
	o.AppDeploy = &v
}

func (o UpdateDeploy200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeploy200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppDeploy) {
		toSerialize["appDeploy"] = o.AppDeploy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateDeploy200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

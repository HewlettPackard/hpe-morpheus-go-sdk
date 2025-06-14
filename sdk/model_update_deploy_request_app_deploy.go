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

// checks if the UpdateDeployRequestAppDeploy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeployRequestAppDeploy{}

// UpdateDeployRequestAppDeploy struct for UpdateDeployRequestAppDeploy
type UpdateDeployRequestAppDeploy struct {
	// JSON encoded list of parameters that varies by instance type.
	Config               map[string]interface{} `json:"config,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _UpdateDeployRequestAppDeploy UpdateDeployRequestAppDeploy

// NewUpdateDeployRequestAppDeploy instantiates a new UpdateDeployRequestAppDeploy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeployRequestAppDeploy() *UpdateDeployRequestAppDeploy {
	this := UpdateDeployRequestAppDeploy{}
	return &this
}

// NewUpdateDeployRequestAppDeployWithDefaults instantiates a new UpdateDeployRequestAppDeploy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeployRequestAppDeployWithDefaults() *UpdateDeployRequestAppDeploy {
	this := UpdateDeployRequestAppDeploy{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UpdateDeployRequestAppDeploy) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeployRequestAppDeploy) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *UpdateDeployRequestAppDeploy) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *UpdateDeployRequestAppDeploy) SetConfig(v map[string]interface{}) {
	o.Config = v
}

func (o UpdateDeployRequestAppDeploy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeployRequestAppDeploy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateDeployRequestAppDeploy) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

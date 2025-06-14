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

// checks if the AddBlueprintRequestOneOf2Helm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddBlueprintRequestOneOf2Helm{}

// AddBlueprintRequestOneOf2Helm struct for AddBlueprintRequestOneOf2Helm
type AddBlueprintRequestOneOf2Helm struct {
	// Configuration Type
	ConfigType           string                            `json:"configType"`
	Git                  *AddBlueprintRequestOneOf2HelmGit `json:"git,omitempty"`
	AdditionalProperties map[string]interface{}            `json:",remain"`
}

type _AddBlueprintRequestOneOf2Helm AddBlueprintRequestOneOf2Helm

// NewAddBlueprintRequestOneOf2Helm instantiates a new AddBlueprintRequestOneOf2Helm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddBlueprintRequestOneOf2Helm(configType string) *AddBlueprintRequestOneOf2Helm {
	this := AddBlueprintRequestOneOf2Helm{}
	this.ConfigType = configType
	return &this
}

// NewAddBlueprintRequestOneOf2HelmWithDefaults instantiates a new AddBlueprintRequestOneOf2Helm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddBlueprintRequestOneOf2HelmWithDefaults() *AddBlueprintRequestOneOf2Helm {
	this := AddBlueprintRequestOneOf2Helm{}
	return &this
}

// GetConfigType returns the ConfigType field value
func (o *AddBlueprintRequestOneOf2Helm) GetConfigType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigType
}

// GetConfigTypeOk returns a tuple with the ConfigType field value
// and a boolean to check if the value has been set.
func (o *AddBlueprintRequestOneOf2Helm) GetConfigTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigType, true
}

// SetConfigType sets field value
func (o *AddBlueprintRequestOneOf2Helm) SetConfigType(v string) {
	o.ConfigType = v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *AddBlueprintRequestOneOf2Helm) GetGit() AddBlueprintRequestOneOf2HelmGit {
	if o == nil || IsNil(o.Git) {
		var ret AddBlueprintRequestOneOf2HelmGit
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprintRequestOneOf2Helm) GetGitOk() (*AddBlueprintRequestOneOf2HelmGit, bool) {
	if o == nil || IsNil(o.Git) {
		return nil, false
	}
	return o.Git, true
}

// IsSetGit returns a boolean if a field has been set.
func (o *AddBlueprintRequestOneOf2Helm) IsSetGit() bool {
	if o != nil && !IsNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given AddBlueprintRequestOneOf2HelmGit and assigns it to the Git field.
func (o *AddBlueprintRequestOneOf2Helm) SetGit(v AddBlueprintRequestOneOf2HelmGit) {
	o.Git = &v
}

func (o AddBlueprintRequestOneOf2Helm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddBlueprintRequestOneOf2Helm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configType"] = o.ConfigType
	if !IsNil(o.Git) {
		toSerialize["git"] = o.Git
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddBlueprintRequestOneOf2Helm) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the AddBlueprintRequestOneOf3Kubernetes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddBlueprintRequestOneOf3Kubernetes{}

// AddBlueprintRequestOneOf3Kubernetes struct for AddBlueprintRequestOneOf3Kubernetes
type AddBlueprintRequestOneOf3Kubernetes struct {
	// Configuration Type
	ConfigType string `json:"configType"`
	// Kubernetes Spec in YAML
	Yaml                 *string                                 `json:"yaml,omitempty"`
	Git                  *AddBlueprintRequestOneOf3KubernetesGit `json:"git,omitempty"`
	AdditionalProperties map[string]interface{}                  `json:",remain"`
}

type _AddBlueprintRequestOneOf3Kubernetes AddBlueprintRequestOneOf3Kubernetes

// NewAddBlueprintRequestOneOf3Kubernetes instantiates a new AddBlueprintRequestOneOf3Kubernetes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddBlueprintRequestOneOf3Kubernetes(configType string) *AddBlueprintRequestOneOf3Kubernetes {
	this := AddBlueprintRequestOneOf3Kubernetes{}
	this.ConfigType = configType
	return &this
}

// NewAddBlueprintRequestOneOf3KubernetesWithDefaults instantiates a new AddBlueprintRequestOneOf3Kubernetes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddBlueprintRequestOneOf3KubernetesWithDefaults() *AddBlueprintRequestOneOf3Kubernetes {
	this := AddBlueprintRequestOneOf3Kubernetes{}
	return &this
}

// GetConfigType returns the ConfigType field value
func (o *AddBlueprintRequestOneOf3Kubernetes) GetConfigType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigType
}

// GetConfigTypeOk returns a tuple with the ConfigType field value
// and a boolean to check if the value has been set.
func (o *AddBlueprintRequestOneOf3Kubernetes) GetConfigTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigType, true
}

// SetConfigType sets field value
func (o *AddBlueprintRequestOneOf3Kubernetes) SetConfigType(v string) {
	o.ConfigType = v
}

// GetYaml returns the Yaml field value if set, zero value otherwise.
func (o *AddBlueprintRequestOneOf3Kubernetes) GetYaml() string {
	if o == nil || IsNil(o.Yaml) {
		var ret string
		return ret
	}
	return *o.Yaml
}

// GetYamlOk returns a tuple with the Yaml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprintRequestOneOf3Kubernetes) GetYamlOk() (*string, bool) {
	if o == nil || IsNil(o.Yaml) {
		return nil, false
	}
	return o.Yaml, true
}

// IsSetYaml returns a boolean if a field has been set.
func (o *AddBlueprintRequestOneOf3Kubernetes) IsSetYaml() bool {
	if o != nil && !IsNil(o.Yaml) {
		return true
	}

	return false
}

// SetYaml gets a reference to the given string and assigns it to the Yaml field.
func (o *AddBlueprintRequestOneOf3Kubernetes) SetYaml(v string) {
	o.Yaml = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *AddBlueprintRequestOneOf3Kubernetes) GetGit() AddBlueprintRequestOneOf3KubernetesGit {
	if o == nil || IsNil(o.Git) {
		var ret AddBlueprintRequestOneOf3KubernetesGit
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprintRequestOneOf3Kubernetes) GetGitOk() (*AddBlueprintRequestOneOf3KubernetesGit, bool) {
	if o == nil || IsNil(o.Git) {
		return nil, false
	}
	return o.Git, true
}

// IsSetGit returns a boolean if a field has been set.
func (o *AddBlueprintRequestOneOf3Kubernetes) IsSetGit() bool {
	if o != nil && !IsNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given AddBlueprintRequestOneOf3KubernetesGit and assigns it to the Git field.
func (o *AddBlueprintRequestOneOf3Kubernetes) SetGit(v AddBlueprintRequestOneOf3KubernetesGit) {
	o.Git = &v
}

func (o AddBlueprintRequestOneOf3Kubernetes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddBlueprintRequestOneOf3Kubernetes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configType"] = o.ConfigType
	if !IsNil(o.Yaml) {
		toSerialize["yaml"] = o.Yaml
	}
	if !IsNil(o.Git) {
		toSerialize["git"] = o.Git
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddBlueprintRequestOneOf3Kubernetes) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

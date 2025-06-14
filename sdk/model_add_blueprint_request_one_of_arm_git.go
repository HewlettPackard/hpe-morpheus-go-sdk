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

// checks if the AddBlueprintRequestOneOfArmGit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddBlueprintRequestOneOfArmGit{}

// AddBlueprintRequestOneOfArmGit struct for AddBlueprintRequestOneOfArmGit
type AddBlueprintRequestOneOfArmGit struct {
	// Morpheus SCM Repository ID
	RepoId int64 `json:"repoId"`
	// Path to ARM Files in the Repository
	Path string `json:"path"`
	// Morpheus SCM Integration ID
	IntegrationId int64 `json:"integrationId"`
	// Branch Name
	Branch               string                 `json:"branch"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddBlueprintRequestOneOfArmGit AddBlueprintRequestOneOfArmGit

// NewAddBlueprintRequestOneOfArmGit instantiates a new AddBlueprintRequestOneOfArmGit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddBlueprintRequestOneOfArmGit(repoId int64, path string, integrationId int64, branch string) *AddBlueprintRequestOneOfArmGit {
	this := AddBlueprintRequestOneOfArmGit{}
	this.RepoId = repoId
	this.Path = path
	this.IntegrationId = integrationId
	this.Branch = branch
	return &this
}

// NewAddBlueprintRequestOneOfArmGitWithDefaults instantiates a new AddBlueprintRequestOneOfArmGit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddBlueprintRequestOneOfArmGitWithDefaults() *AddBlueprintRequestOneOfArmGit {
	this := AddBlueprintRequestOneOfArmGit{}
	return &this
}

// GetRepoId returns the RepoId field value
func (o *AddBlueprintRequestOneOfArmGit) GetRepoId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RepoId
}

// GetRepoIdOk returns a tuple with the RepoId field value
// and a boolean to check if the value has been set.
func (o *AddBlueprintRequestOneOfArmGit) GetRepoIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepoId, true
}

// SetRepoId sets field value
func (o *AddBlueprintRequestOneOfArmGit) SetRepoId(v int64) {
	o.RepoId = v
}

// GetPath returns the Path field value
func (o *AddBlueprintRequestOneOfArmGit) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *AddBlueprintRequestOneOfArmGit) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *AddBlueprintRequestOneOfArmGit) SetPath(v string) {
	o.Path = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *AddBlueprintRequestOneOfArmGit) GetIntegrationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *AddBlueprintRequestOneOfArmGit) GetIntegrationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *AddBlueprintRequestOneOfArmGit) SetIntegrationId(v int64) {
	o.IntegrationId = v
}

// GetBranch returns the Branch field value
func (o *AddBlueprintRequestOneOfArmGit) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *AddBlueprintRequestOneOfArmGit) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *AddBlueprintRequestOneOfArmGit) SetBranch(v string) {
	o.Branch = v
}

func (o AddBlueprintRequestOneOfArmGit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddBlueprintRequestOneOfArmGit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["repoId"] = o.RepoId
	toSerialize["path"] = o.Path
	toSerialize["integrationId"] = o.IntegrationId
	toSerialize["branch"] = o.Branch

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddBlueprintRequestOneOfArmGit) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

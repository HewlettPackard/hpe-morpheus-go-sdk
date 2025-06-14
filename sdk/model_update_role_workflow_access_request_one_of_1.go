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

// checks if the UpdateRoleWorkflowAccessRequestOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoleWorkflowAccessRequestOneOf1{}

// UpdateRoleWorkflowAccessRequestOneOf1 struct for UpdateRoleWorkflowAccessRequestOneOf1
type UpdateRoleWorkflowAccessRequestOneOf1 struct {
	// Apply to all workflows (taskSets)
	AllTaskSets bool `json:"allTaskSets"`
	// The new access level.
	Access               string                 `json:"access"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _UpdateRoleWorkflowAccessRequestOneOf1 UpdateRoleWorkflowAccessRequestOneOf1

// NewUpdateRoleWorkflowAccessRequestOneOf1 instantiates a new UpdateRoleWorkflowAccessRequestOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleWorkflowAccessRequestOneOf1(allTaskSets bool, access string) *UpdateRoleWorkflowAccessRequestOneOf1 {
	this := UpdateRoleWorkflowAccessRequestOneOf1{}
	this.AllTaskSets = allTaskSets
	this.Access = access
	return &this
}

// NewUpdateRoleWorkflowAccessRequestOneOf1WithDefaults instantiates a new UpdateRoleWorkflowAccessRequestOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleWorkflowAccessRequestOneOf1WithDefaults() *UpdateRoleWorkflowAccessRequestOneOf1 {
	this := UpdateRoleWorkflowAccessRequestOneOf1{}
	return &this
}

// GetAllTaskSets returns the AllTaskSets field value
func (o *UpdateRoleWorkflowAccessRequestOneOf1) GetAllTaskSets() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllTaskSets
}

// GetAllTaskSetsOk returns a tuple with the AllTaskSets field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleWorkflowAccessRequestOneOf1) GetAllTaskSetsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllTaskSets, true
}

// SetAllTaskSets sets field value
func (o *UpdateRoleWorkflowAccessRequestOneOf1) SetAllTaskSets(v bool) {
	o.AllTaskSets = v
}

// GetAccess returns the Access field value
func (o *UpdateRoleWorkflowAccessRequestOneOf1) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleWorkflowAccessRequestOneOf1) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *UpdateRoleWorkflowAccessRequestOneOf1) SetAccess(v string) {
	o.Access = v
}

func (o UpdateRoleWorkflowAccessRequestOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoleWorkflowAccessRequestOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allTaskSets"] = o.AllTaskSets
	toSerialize["access"] = o.Access

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateRoleWorkflowAccessRequestOneOf1) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

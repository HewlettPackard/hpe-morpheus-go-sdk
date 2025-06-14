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

// checks if the RolePermissionGroupAll type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolePermissionGroupAll{}

// RolePermissionGroupAll struct for RolePermissionGroupAll
type RolePermissionGroupAll struct {
	// Apply to all groups (site)
	AllGroups bool `json:"allGroups"`
	// The new access level.
	Access               string                 `json:"access"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _RolePermissionGroupAll RolePermissionGroupAll

// NewRolePermissionGroupAll instantiates a new RolePermissionGroupAll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermissionGroupAll(allGroups bool, access string) *RolePermissionGroupAll {
	this := RolePermissionGroupAll{}
	this.AllGroups = allGroups
	this.Access = access
	return &this
}

// NewRolePermissionGroupAllWithDefaults instantiates a new RolePermissionGroupAll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionGroupAllWithDefaults() *RolePermissionGroupAll {
	this := RolePermissionGroupAll{}
	return &this
}

// GetAllGroups returns the AllGroups field value
func (o *RolePermissionGroupAll) GetAllGroups() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllGroups
}

// GetAllGroupsOk returns a tuple with the AllGroups field value
// and a boolean to check if the value has been set.
func (o *RolePermissionGroupAll) GetAllGroupsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllGroups, true
}

// SetAllGroups sets field value
func (o *RolePermissionGroupAll) SetAllGroups(v bool) {
	o.AllGroups = v
}

// GetAccess returns the Access field value
func (o *RolePermissionGroupAll) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *RolePermissionGroupAll) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *RolePermissionGroupAll) SetAccess(v string) {
	o.Access = v
}

func (o RolePermissionGroupAll) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolePermissionGroupAll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allGroups"] = o.AllGroups
	toSerialize["access"] = o.Access

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *RolePermissionGroupAll) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

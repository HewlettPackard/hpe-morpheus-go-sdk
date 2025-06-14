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

// checks if the RolePermissionDefaultVDIPool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolePermissionDefaultVDIPool{}

// RolePermissionDefaultVDIPool struct for RolePermissionDefaultVDIPool
type RolePermissionDefaultVDIPool struct {
	// `VdiPools` is the code for Default VDI Pool Access
	PermissionCode string `json:"permissionCode"`
	// The new access level.
	Access               string                 `json:"access"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _RolePermissionDefaultVDIPool RolePermissionDefaultVDIPool

// NewRolePermissionDefaultVDIPool instantiates a new RolePermissionDefaultVDIPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolePermissionDefaultVDIPool(permissionCode string, access string) *RolePermissionDefaultVDIPool {
	this := RolePermissionDefaultVDIPool{}
	this.PermissionCode = permissionCode
	this.Access = access
	return &this
}

// NewRolePermissionDefaultVDIPoolWithDefaults instantiates a new RolePermissionDefaultVDIPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolePermissionDefaultVDIPoolWithDefaults() *RolePermissionDefaultVDIPool {
	this := RolePermissionDefaultVDIPool{}
	return &this
}

// GetPermissionCode returns the PermissionCode field value
func (o *RolePermissionDefaultVDIPool) GetPermissionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PermissionCode
}

// GetPermissionCodeOk returns a tuple with the PermissionCode field value
// and a boolean to check if the value has been set.
func (o *RolePermissionDefaultVDIPool) GetPermissionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionCode, true
}

// SetPermissionCode sets field value
func (o *RolePermissionDefaultVDIPool) SetPermissionCode(v string) {
	o.PermissionCode = v
}

// GetAccess returns the Access field value
func (o *RolePermissionDefaultVDIPool) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *RolePermissionDefaultVDIPool) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *RolePermissionDefaultVDIPool) SetAccess(v string) {
	o.Access = v
}

func (o RolePermissionDefaultVDIPool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolePermissionDefaultVDIPool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissionCode"] = o.PermissionCode
	toSerialize["access"] = o.Access

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *RolePermissionDefaultVDIPool) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

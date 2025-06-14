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

// checks if the AddRolesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddRolesRequest{}

// AddRolesRequest struct for AddRolesRequest
type AddRolesRequest struct {
	Role                 AddRolesRequestRole    `json:"role"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddRolesRequest AddRolesRequest

// NewAddRolesRequest instantiates a new AddRolesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddRolesRequest(role AddRolesRequestRole) *AddRolesRequest {
	this := AddRolesRequest{}
	this.Role = role
	return &this
}

// NewAddRolesRequestWithDefaults instantiates a new AddRolesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddRolesRequestWithDefaults() *AddRolesRequest {
	this := AddRolesRequest{}
	return &this
}

// GetRole returns the Role field value
func (o *AddRolesRequest) GetRole() AddRolesRequestRole {
	if o == nil {
		var ret AddRolesRequestRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *AddRolesRequest) GetRoleOk() (*AddRolesRequestRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *AddRolesRequest) SetRole(v AddRolesRequestRole) {
	o.Role = v
}

func (o AddRolesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddRolesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddRolesRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

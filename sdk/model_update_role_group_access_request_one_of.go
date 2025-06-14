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

// checks if the UpdateRoleGroupAccessRequestOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoleGroupAccessRequestOneOf{}

// UpdateRoleGroupAccessRequestOneOf struct for UpdateRoleGroupAccessRequestOneOf
type UpdateRoleGroupAccessRequestOneOf struct {
	// `id` of the group (site)
	GroupId int32 `json:"groupId"`
	// The new access level.
	Access               string                 `json:"access"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _UpdateRoleGroupAccessRequestOneOf UpdateRoleGroupAccessRequestOneOf

// NewUpdateRoleGroupAccessRequestOneOf instantiates a new UpdateRoleGroupAccessRequestOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleGroupAccessRequestOneOf(groupId int32, access string) *UpdateRoleGroupAccessRequestOneOf {
	this := UpdateRoleGroupAccessRequestOneOf{}
	this.GroupId = groupId
	this.Access = access
	return &this
}

// NewUpdateRoleGroupAccessRequestOneOfWithDefaults instantiates a new UpdateRoleGroupAccessRequestOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleGroupAccessRequestOneOfWithDefaults() *UpdateRoleGroupAccessRequestOneOf {
	this := UpdateRoleGroupAccessRequestOneOf{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *UpdateRoleGroupAccessRequestOneOf) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleGroupAccessRequestOneOf) GetGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *UpdateRoleGroupAccessRequestOneOf) SetGroupId(v int32) {
	o.GroupId = v
}

// GetAccess returns the Access field value
func (o *UpdateRoleGroupAccessRequestOneOf) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleGroupAccessRequestOneOf) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *UpdateRoleGroupAccessRequestOneOf) SetAccess(v string) {
	o.Access = v
}

func (o UpdateRoleGroupAccessRequestOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoleGroupAccessRequestOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupId"] = o.GroupId
	toSerialize["access"] = o.Access

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateRoleGroupAccessRequestOneOf) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27{}

// AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27 - User Creation
type AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27 struct {
	CreateUserType       *string                `json:"createUserType,omitempty"`
	CreateUser           *bool                  `json:"createUser,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27 AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27

// NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27 instantiates a new AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27() *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27 {
	this := AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27{}
	return &this
}

// NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27WithDefaults instantiates a new AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27WithDefaults() *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27 {
	this := AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27{}
	return &this
}

// GetCreateUserType returns the CreateUserType field value if set, zero value otherwise.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27) GetCreateUserType() string {
	if o == nil || IsNil(o.CreateUserType) {
		var ret string
		return ret
	}
	return *o.CreateUserType
}

// GetCreateUserTypeOk returns a tuple with the CreateUserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27) GetCreateUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateUserType) {
		return nil, false
	}
	return o.CreateUserType, true
}

// IsSetCreateUserType returns a boolean if a field has been set.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27) IsSetCreateUserType() bool {
	if o != nil && !IsNil(o.CreateUserType) {
		return true
	}

	return false
}

// SetCreateUserType gets a reference to the given string and assigns it to the CreateUserType field.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27) SetCreateUserType(v string) {
	o.CreateUserType = &v
}

// GetCreateUser returns the CreateUser field value if set, zero value otherwise.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27) GetCreateUser() bool {
	if o == nil || IsNil(o.CreateUser) {
		var ret bool
		return ret
	}
	return *o.CreateUser
}

// GetCreateUserOk returns a tuple with the CreateUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27) GetCreateUserOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateUser) {
		return nil, false
	}
	return o.CreateUser, true
}

// IsSetCreateUser returns a boolean if a field has been set.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27) IsSetCreateUser() bool {
	if o != nil && !IsNil(o.CreateUser) {
		return true
	}

	return false
}

// SetCreateUser gets a reference to the given bool and assigns it to the CreateUser field.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27) SetCreateUser(v bool) {
	o.CreateUser = &v
}

func (o AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateUserType) {
		toSerialize["createUserType"] = o.CreateUserType
	}
	if !IsNil(o.CreateUser) {
		toSerialize["createUser"] = o.CreateUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf27) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

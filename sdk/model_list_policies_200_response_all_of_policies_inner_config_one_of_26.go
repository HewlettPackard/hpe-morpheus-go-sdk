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

// checks if the ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26{}

// ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26 - User Creation
type ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26 struct {
	CreateUserType       string                 `json:"createUserType"`
	CreateUser           *bool                  `json:"createUser,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26 ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26

// NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26 instantiates a new ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26(createUserType string) *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26 {
	this := ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26{}
	this.CreateUserType = createUserType
	return &this
}

// NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26WithDefaults instantiates a new ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26WithDefaults() *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26 {
	this := ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26{}
	return &this
}

// GetCreateUserType returns the CreateUserType field value
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26) GetCreateUserType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreateUserType
}

// GetCreateUserTypeOk returns a tuple with the CreateUserType field value
// and a boolean to check if the value has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26) GetCreateUserTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateUserType, true
}

// SetCreateUserType sets field value
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26) SetCreateUserType(v string) {
	o.CreateUserType = v
}

// GetCreateUser returns the CreateUser field value if set, zero value otherwise.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26) GetCreateUser() bool {
	if o == nil || IsNil(o.CreateUser) {
		var ret bool
		return ret
	}
	return *o.CreateUser
}

// GetCreateUserOk returns a tuple with the CreateUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26) GetCreateUserOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateUser) {
		return nil, false
	}
	return o.CreateUser, true
}

// IsSetCreateUser returns a boolean if a field has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26) IsSetCreateUser() bool {
	if o != nil && !IsNil(o.CreateUser) {
		return true
	}

	return false
}

// SetCreateUser gets a reference to the given bool and assigns it to the CreateUser field.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26) SetCreateUser(v bool) {
	o.CreateUser = &v
}

func (o ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createUserType"] = o.CreateUserType
	if !IsNil(o.CreateUser) {
		toSerialize["createUser"] = o.CreateUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf26) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

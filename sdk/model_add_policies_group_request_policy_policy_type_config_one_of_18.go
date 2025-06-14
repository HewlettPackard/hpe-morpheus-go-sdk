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

// checks if the AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18{}

// AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18 - Max Virtual Servers
type AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18 struct {
	MaxVirtualServers    *string                `json:"maxVirtualServers,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18 AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18

// NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18 instantiates a new AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18() *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18 {
	this := AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18{}
	return &this
}

// NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18WithDefaults instantiates a new AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18WithDefaults() *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18 {
	this := AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18{}
	return &this
}

// GetMaxVirtualServers returns the MaxVirtualServers field value if set, zero value otherwise.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18) GetMaxVirtualServers() string {
	if o == nil || IsNil(o.MaxVirtualServers) {
		var ret string
		return ret
	}
	return *o.MaxVirtualServers
}

// GetMaxVirtualServersOk returns a tuple with the MaxVirtualServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18) GetMaxVirtualServersOk() (*string, bool) {
	if o == nil || IsNil(o.MaxVirtualServers) {
		return nil, false
	}
	return o.MaxVirtualServers, true
}

// IsSetMaxVirtualServers returns a boolean if a field has been set.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18) IsSetMaxVirtualServers() bool {
	if o != nil && !IsNil(o.MaxVirtualServers) {
		return true
	}

	return false
}

// SetMaxVirtualServers gets a reference to the given string and assigns it to the MaxVirtualServers field.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18) SetMaxVirtualServers(v string) {
	o.MaxVirtualServers = &v
}

func (o AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxVirtualServers) {
		toSerialize["maxVirtualServers"] = o.MaxVirtualServers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf18) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

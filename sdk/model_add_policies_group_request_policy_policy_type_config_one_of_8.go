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

// checks if the AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8{}

// AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8 - File Share Storage Quota
type AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8 struct {
	MaxStorage           *string                `json:"maxStorage,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8 AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8

// NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8 instantiates a new AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8() *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8 {
	this := AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8{}
	return &this
}

// NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8WithDefaults instantiates a new AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8WithDefaults() *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8 {
	this := AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8{}
	return &this
}

// GetMaxStorage returns the MaxStorage field value if set, zero value otherwise.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8) GetMaxStorage() string {
	if o == nil || IsNil(o.MaxStorage) {
		var ret string
		return ret
	}
	return *o.MaxStorage
}

// GetMaxStorageOk returns a tuple with the MaxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8) GetMaxStorageOk() (*string, bool) {
	if o == nil || IsNil(o.MaxStorage) {
		return nil, false
	}
	return o.MaxStorage, true
}

// IsSetMaxStorage returns a boolean if a field has been set.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8) IsSetMaxStorage() bool {
	if o != nil && !IsNil(o.MaxStorage) {
		return true
	}

	return false
}

// SetMaxStorage gets a reference to the given string and assigns it to the MaxStorage field.
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8) SetMaxStorage(v string) {
	o.MaxStorage = &v
}

func (o AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxStorage) {
		toSerialize["maxStorage"] = o.MaxStorage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddPoliciesGroupRequestPolicyPolicyTypeConfigOneOf8) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

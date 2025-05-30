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

// checks if the CreateNetworkFirewallRuleRequestRuleSources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirewallRuleRequestRuleSources{}

// CreateNetworkFirewallRuleRequestRuleSources struct for CreateNetworkFirewallRuleRequestRuleSources
type CreateNetworkFirewallRuleRequestRuleSources struct {
	Id []string `json:"id,omitempty"`
}

// NewCreateNetworkFirewallRuleRequestRuleSources instantiates a new CreateNetworkFirewallRuleRequestRuleSources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirewallRuleRequestRuleSources() *CreateNetworkFirewallRuleRequestRuleSources {
	this := CreateNetworkFirewallRuleRequestRuleSources{}
	return &this
}

// NewCreateNetworkFirewallRuleRequestRuleSourcesWithDefaults instantiates a new CreateNetworkFirewallRuleRequestRuleSources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirewallRuleRequestRuleSourcesWithDefaults() *CreateNetworkFirewallRuleRequestRuleSources {
	this := CreateNetworkFirewallRuleRequestRuleSources{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateNetworkFirewallRuleRequestRuleSources) GetId() []string {
	if o == nil || IsNil(o.Id) {
		var ret []string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirewallRuleRequestRuleSources) GetIdOk() ([]string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *CreateNetworkFirewallRuleRequestRuleSources) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given []string and assigns it to the Id field.
func (o *CreateNetworkFirewallRuleRequestRuleSources) SetId(v []string) {
	o.Id = v
}

func (o CreateNetworkFirewallRuleRequestRuleSources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirewallRuleRequestRuleSources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreateNetworkFirewallRuleRequestRuleSources struct {
	value *CreateNetworkFirewallRuleRequestRuleSources
	isSet bool
}

func (v NullableCreateNetworkFirewallRuleRequestRuleSources) Get() *CreateNetworkFirewallRuleRequestRuleSources {
	return v.value
}

func (v *NullableCreateNetworkFirewallRuleRequestRuleSources) Set(val *CreateNetworkFirewallRuleRequestRuleSources) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirewallRuleRequestRuleSources) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirewallRuleRequestRuleSources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirewallRuleRequestRuleSources(val *CreateNetworkFirewallRuleRequestRuleSources) *NullableCreateNetworkFirewallRuleRequestRuleSources {
	return &NullableCreateNetworkFirewallRuleRequestRuleSources{value: val, isSet: true}
}

func (v NullableCreateNetworkFirewallRuleRequestRuleSources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirewallRuleRequestRuleSources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



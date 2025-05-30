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

// checks if the CreateNetworkFirewallRuleRequestRuleRuleGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirewallRuleRequestRuleRuleGroup{}

// CreateNetworkFirewallRuleRequestRuleRuleGroup struct for CreateNetworkFirewallRuleRequestRuleRuleGroup
type CreateNetworkFirewallRuleRequestRuleRuleGroup struct {
	// Firewall rule group for rule (*applicable to select network servers).
	Id *int32 `json:"id,omitempty"`
}

// NewCreateNetworkFirewallRuleRequestRuleRuleGroup instantiates a new CreateNetworkFirewallRuleRequestRuleRuleGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirewallRuleRequestRuleRuleGroup() *CreateNetworkFirewallRuleRequestRuleRuleGroup {
	this := CreateNetworkFirewallRuleRequestRuleRuleGroup{}
	return &this
}

// NewCreateNetworkFirewallRuleRequestRuleRuleGroupWithDefaults instantiates a new CreateNetworkFirewallRuleRequestRuleRuleGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirewallRuleRequestRuleRuleGroupWithDefaults() *CreateNetworkFirewallRuleRequestRuleRuleGroup {
	this := CreateNetworkFirewallRuleRequestRuleRuleGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateNetworkFirewallRuleRequestRuleRuleGroup) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirewallRuleRequestRuleRuleGroup) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *CreateNetworkFirewallRuleRequestRuleRuleGroup) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CreateNetworkFirewallRuleRequestRuleRuleGroup) SetId(v int32) {
	o.Id = &v
}

func (o CreateNetworkFirewallRuleRequestRuleRuleGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirewallRuleRequestRuleRuleGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreateNetworkFirewallRuleRequestRuleRuleGroup struct {
	value *CreateNetworkFirewallRuleRequestRuleRuleGroup
	isSet bool
}

func (v NullableCreateNetworkFirewallRuleRequestRuleRuleGroup) Get() *CreateNetworkFirewallRuleRequestRuleRuleGroup {
	return v.value
}

func (v *NullableCreateNetworkFirewallRuleRequestRuleRuleGroup) Set(val *CreateNetworkFirewallRuleRequestRuleRuleGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirewallRuleRequestRuleRuleGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirewallRuleRequestRuleRuleGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirewallRuleRequestRuleRuleGroup(val *CreateNetworkFirewallRuleRequestRuleRuleGroup) *NullableCreateNetworkFirewallRuleRequestRuleRuleGroup {
	return &NullableCreateNetworkFirewallRuleRequestRuleRuleGroup{value: val, isSet: true}
}

func (v NullableCreateNetworkFirewallRuleRequestRuleRuleGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirewallRuleRequestRuleRuleGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



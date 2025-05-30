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

// checks if the SecurityGroupLocationAzureCustomOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupLocationAzureCustomOptions{}

// SecurityGroupLocationAzureCustomOptions struct for SecurityGroupLocationAzureCustomOptions
type SecurityGroupLocationAzureCustomOptions struct {
	// External ID of Azure Resource Group
	ResourceGroup *string `json:"resourceGroup,omitempty"`
}

// NewSecurityGroupLocationAzureCustomOptions instantiates a new SecurityGroupLocationAzureCustomOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupLocationAzureCustomOptions() *SecurityGroupLocationAzureCustomOptions {
	this := SecurityGroupLocationAzureCustomOptions{}
	return &this
}

// NewSecurityGroupLocationAzureCustomOptionsWithDefaults instantiates a new SecurityGroupLocationAzureCustomOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupLocationAzureCustomOptionsWithDefaults() *SecurityGroupLocationAzureCustomOptions {
	this := SecurityGroupLocationAzureCustomOptions{}
	return &this
}

// GetResourceGroup returns the ResourceGroup field value if set, zero value otherwise.
func (o *SecurityGroupLocationAzureCustomOptions) GetResourceGroup() string {
	if o == nil || IsNil(o.ResourceGroup) {
		var ret string
		return ret
	}
	return *o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLocationAzureCustomOptions) GetResourceGroupOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceGroup) {
		return nil, false
	}
	return o.ResourceGroup, true
}

// IsSetResourceGroup returns a boolean if a field has been set.
func (o *SecurityGroupLocationAzureCustomOptions) IsSetResourceGroup() bool {
	if o != nil && !IsNil(o.ResourceGroup) {
		return true
	}

	return false
}

// SetResourceGroup gets a reference to the given string and assigns it to the ResourceGroup field.
func (o *SecurityGroupLocationAzureCustomOptions) SetResourceGroup(v string) {
	o.ResourceGroup = &v
}

func (o SecurityGroupLocationAzureCustomOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityGroupLocationAzureCustomOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceGroup) {
		toSerialize["resourceGroup"] = o.ResourceGroup
	}
	return toSerialize, nil
}

type NullableSecurityGroupLocationAzureCustomOptions struct {
	value *SecurityGroupLocationAzureCustomOptions
	isSet bool
}

func (v NullableSecurityGroupLocationAzureCustomOptions) Get() *SecurityGroupLocationAzureCustomOptions {
	return v.value
}

func (v *NullableSecurityGroupLocationAzureCustomOptions) Set(val *SecurityGroupLocationAzureCustomOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupLocationAzureCustomOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupLocationAzureCustomOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupLocationAzureCustomOptions(val *SecurityGroupLocationAzureCustomOptions) *NullableSecurityGroupLocationAzureCustomOptions {
	return &NullableSecurityGroupLocationAzureCustomOptions{value: val, isSet: true}
}

func (v NullableSecurityGroupLocationAzureCustomOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupLocationAzureCustomOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig{}

// GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig struct for GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig
type GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig struct {
	EdgeCluster *string `json:"edgeCluster,omitempty"`
	PreferredEdgeNode1 *string `json:"preferredEdgeNode1,omitempty"`
	PreferredEdgeNode2 *string `json:"preferredEdgeNode2,omitempty"`
}

// NewGetNetworkDhcpServer200ResponseNetworkDhcpServerConfig instantiates a new GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkDhcpServer200ResponseNetworkDhcpServerConfig() *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig {
	this := GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig{}
	return &this
}

// NewGetNetworkDhcpServer200ResponseNetworkDhcpServerConfigWithDefaults instantiates a new GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkDhcpServer200ResponseNetworkDhcpServerConfigWithDefaults() *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig {
	this := GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig{}
	return &this
}

// GetEdgeCluster returns the EdgeCluster field value if set, zero value otherwise.
func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) GetEdgeCluster() string {
	if o == nil || IsNil(o.EdgeCluster) {
		var ret string
		return ret
	}
	return *o.EdgeCluster
}

// GetEdgeClusterOk returns a tuple with the EdgeCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) GetEdgeClusterOk() (*string, bool) {
	if o == nil || IsNil(o.EdgeCluster) {
		return nil, false
	}
	return o.EdgeCluster, true
}

// IsSetEdgeCluster returns a boolean if a field has been set.
func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) IsSetEdgeCluster() bool {
	if o != nil && !IsNil(o.EdgeCluster) {
		return true
	}

	return false
}

// SetEdgeCluster gets a reference to the given string and assigns it to the EdgeCluster field.
func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) SetEdgeCluster(v string) {
	o.EdgeCluster = &v
}

// GetPreferredEdgeNode1 returns the PreferredEdgeNode1 field value if set, zero value otherwise.
func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) GetPreferredEdgeNode1() string {
	if o == nil || IsNil(o.PreferredEdgeNode1) {
		var ret string
		return ret
	}
	return *o.PreferredEdgeNode1
}

// GetPreferredEdgeNode1Ok returns a tuple with the PreferredEdgeNode1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) GetPreferredEdgeNode1Ok() (*string, bool) {
	if o == nil || IsNil(o.PreferredEdgeNode1) {
		return nil, false
	}
	return o.PreferredEdgeNode1, true
}

// IsSetPreferredEdgeNode1 returns a boolean if a field has been set.
func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) IsSetPreferredEdgeNode1() bool {
	if o != nil && !IsNil(o.PreferredEdgeNode1) {
		return true
	}

	return false
}

// SetPreferredEdgeNode1 gets a reference to the given string and assigns it to the PreferredEdgeNode1 field.
func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) SetPreferredEdgeNode1(v string) {
	o.PreferredEdgeNode1 = &v
}

// GetPreferredEdgeNode2 returns the PreferredEdgeNode2 field value if set, zero value otherwise.
func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) GetPreferredEdgeNode2() string {
	if o == nil || IsNil(o.PreferredEdgeNode2) {
		var ret string
		return ret
	}
	return *o.PreferredEdgeNode2
}

// GetPreferredEdgeNode2Ok returns a tuple with the PreferredEdgeNode2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) GetPreferredEdgeNode2Ok() (*string, bool) {
	if o == nil || IsNil(o.PreferredEdgeNode2) {
		return nil, false
	}
	return o.PreferredEdgeNode2, true
}

// IsSetPreferredEdgeNode2 returns a boolean if a field has been set.
func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) IsSetPreferredEdgeNode2() bool {
	if o != nil && !IsNil(o.PreferredEdgeNode2) {
		return true
	}

	return false
}

// SetPreferredEdgeNode2 gets a reference to the given string and assigns it to the PreferredEdgeNode2 field.
func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) SetPreferredEdgeNode2(v string) {
	o.PreferredEdgeNode2 = &v
}

func (o GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EdgeCluster) {
		toSerialize["edgeCluster"] = o.EdgeCluster
	}
	if !IsNil(o.PreferredEdgeNode1) {
		toSerialize["preferredEdgeNode1"] = o.PreferredEdgeNode1
	}
	if !IsNil(o.PreferredEdgeNode2) {
		toSerialize["preferredEdgeNode2"] = o.PreferredEdgeNode2
	}
	return toSerialize, nil
}

type NullableGetNetworkDhcpServer200ResponseNetworkDhcpServerConfig struct {
	value *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig
	isSet bool
}

func (v NullableGetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) Get() *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig {
	return v.value
}

func (v *NullableGetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) Set(val *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkDhcpServer200ResponseNetworkDhcpServerConfig(val *GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) *NullableGetNetworkDhcpServer200ResponseNetworkDhcpServerConfig {
	return &NullableGetNetworkDhcpServer200ResponseNetworkDhcpServerConfig{value: val, isSet: true}
}

func (v NullableGetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkDhcpServer200ResponseNetworkDhcpServerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



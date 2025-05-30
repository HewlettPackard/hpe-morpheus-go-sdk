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
	"bytes"
	"fmt"
)

// checks if the ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf{}

// ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf Type AWS
type ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf struct {
	// Provide the base CIDR Block to use for this VPC (must be between a /16 and /28 Block)
	CidrBlock string `json:"cidrBlock"`
	Tenancy string `json:"tenancy"`
}

type _ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf

// NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf instantiates a new ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf(cidrBlock string, tenancy string) *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf {
	this := ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf{}
	this.CidrBlock = cidrBlock
	this.Tenancy = tenancy
	return &this
}

// NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOfWithDefaults instantiates a new ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOfWithDefaults() *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf {
	this := ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf{}
	var tenancy string = "default"
	this.Tenancy = tenancy
	return &this
}

// GetCidrBlock returns the CidrBlock field value
func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) GetCidrBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CidrBlock
}

// GetCidrBlockOk returns a tuple with the CidrBlock field value
// and a boolean to check if the value has been set.
func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) GetCidrBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CidrBlock, true
}

// SetCidrBlock sets field value
func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) SetCidrBlock(v string) {
	o.CidrBlock = v
}

// GetTenancy returns the Tenancy field value
func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) GetTenancy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenancy
}

// GetTenancyOk returns a tuple with the Tenancy field value
// and a boolean to check if the value has been set.
func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) GetTenancyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenancy, true
}

// SetTenancy sets field value
func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) SetTenancy(v string) {
	o.Tenancy = v
}

func (o ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cidrBlock"] = o.CidrBlock
	toSerialize["tenancy"] = o.Tenancy
	return toSerialize, nil
}

func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cidrBlock",
		"tenancy",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf := _ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf)

	if err != nil {
		return err
	}

	*o = ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf(varListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf)

	return err
}

type NullableListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf struct {
	value *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf
	isSet bool
}

func (v NullableListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) Get() *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf {
	return v.value
}

func (v *NullableListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) Set(val *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf(val *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) *NullableListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf {
	return &NullableListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf{value: val, isSet: true}
}

func (v NullableListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



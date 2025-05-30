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
	"fmt"
)

// checks if the ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22{}

// ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 - Router Quota 
type ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 struct {
	MaxRouters string `json:"maxRouters"`
	AdditionalProperties map[string]interface{}
}

type _ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22

// NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 instantiates a new ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22(maxRouters string) *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 {
	this := ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22{}
	this.MaxRouters = maxRouters
	return &this
}

// NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22WithDefaults instantiates a new ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22WithDefaults() *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 {
	this := ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22{}
	return &this
}

// GetMaxRouters returns the MaxRouters field value
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) GetMaxRouters() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxRouters
}

// GetMaxRoutersOk returns a tuple with the MaxRouters field value
// and a boolean to check if the value has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) GetMaxRoutersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxRouters, true
}

// SetMaxRouters sets field value
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) SetMaxRouters(v string) {
	o.MaxRouters = v
}

func (o ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxRouters"] = o.MaxRouters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"maxRouters",
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

	varListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 := _ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22{}

	err = json.Unmarshal(data, &varListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22)

	if err != nil {
		return err
	}

	*o = ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22(varListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maxRouters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 struct {
	value *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22
	isSet bool
}

func (v NullableListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) Get() *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 {
	return v.value
}

func (v *NullableListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) Set(val *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) {
	v.value = val
	v.isSet = true
}

func (v NullableListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) IsSet() bool {
	return v.isSet
}

func (v *NullableListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22(val *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) *NullableListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22 {
	return &NullableListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22{value: val, isSet: true}
}

func (v NullableListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPolicies200ResponseAllOfPoliciesInnerConfigOneOf22) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



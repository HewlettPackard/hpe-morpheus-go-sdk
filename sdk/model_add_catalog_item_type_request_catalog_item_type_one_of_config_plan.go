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

// checks if the AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan{}

// AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan struct for AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan
type AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan struct {
	Id AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlanId `json:"id"`
}

type _AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan

// NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan(id AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlanId) *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan {
	this := AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan{}
	this.Id = id
	return &this
}

// NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlanWithDefaults instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlanWithDefaults() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan {
	this := AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan{}
	return &this
}

// GetId returns the Id field value
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) GetId() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlanId {
	if o == nil {
		var ret AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlanId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) GetIdOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlanId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) SetId(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlanId) {
	o.Id = v
}

func (o AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan := _AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan)

	if err != nil {
		return err
	}

	*o = AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan(varAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan)

	return err
}

type NullableAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan struct {
	value *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan
	isSet bool
}

func (v NullableAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) Get() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan {
	return v.value
}

func (v *NullableAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) Set(val *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan(val *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) *NullableAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan {
	return &NullableAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan{value: val, isSet: true}
}

func (v NullableAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the AddBudgetsRequestBudgetForecastType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddBudgetsRequestBudgetForecastType{}

// AddBudgetsRequestBudgetForecastType Forecast Model
type AddBudgetsRequestBudgetForecastType struct {
	// The ID of the Forecast Model type. See GET /api/options/forecastTypes
	Id int64 `json:"id"`
}

type _AddBudgetsRequestBudgetForecastType AddBudgetsRequestBudgetForecastType

// NewAddBudgetsRequestBudgetForecastType instantiates a new AddBudgetsRequestBudgetForecastType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddBudgetsRequestBudgetForecastType(id int64) *AddBudgetsRequestBudgetForecastType {
	this := AddBudgetsRequestBudgetForecastType{}
	this.Id = id
	return &this
}

// NewAddBudgetsRequestBudgetForecastTypeWithDefaults instantiates a new AddBudgetsRequestBudgetForecastType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddBudgetsRequestBudgetForecastTypeWithDefaults() *AddBudgetsRequestBudgetForecastType {
	this := AddBudgetsRequestBudgetForecastType{}
	return &this
}

// GetId returns the Id field value
func (o *AddBudgetsRequestBudgetForecastType) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddBudgetsRequestBudgetForecastType) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddBudgetsRequestBudgetForecastType) SetId(v int64) {
	o.Id = v
}

func (o AddBudgetsRequestBudgetForecastType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddBudgetsRequestBudgetForecastType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AddBudgetsRequestBudgetForecastType) UnmarshalJSON(data []byte) (err error) {
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

	varAddBudgetsRequestBudgetForecastType := _AddBudgetsRequestBudgetForecastType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddBudgetsRequestBudgetForecastType)

	if err != nil {
		return err
	}

	*o = AddBudgetsRequestBudgetForecastType(varAddBudgetsRequestBudgetForecastType)

	return err
}

type NullableAddBudgetsRequestBudgetForecastType struct {
	value *AddBudgetsRequestBudgetForecastType
	isSet bool
}

func (v NullableAddBudgetsRequestBudgetForecastType) Get() *AddBudgetsRequestBudgetForecastType {
	return v.value
}

func (v *NullableAddBudgetsRequestBudgetForecastType) Set(val *AddBudgetsRequestBudgetForecastType) {
	v.value = val
	v.isSet = true
}

func (v NullableAddBudgetsRequestBudgetForecastType) IsSet() bool {
	return v.isSet
}

func (v *NullableAddBudgetsRequestBudgetForecastType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddBudgetsRequestBudgetForecastType(val *AddBudgetsRequestBudgetForecastType) *NullableAddBudgetsRequestBudgetForecastType {
	return &NullableAddBudgetsRequestBudgetForecastType{value: val, isSet: true}
}

func (v NullableAddBudgetsRequestBudgetForecastType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddBudgetsRequestBudgetForecastType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



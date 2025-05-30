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

// checks if the AddServicePlansRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddServicePlansRequest{}

// AddServicePlansRequest struct for AddServicePlansRequest
type AddServicePlansRequest struct {
	ServicePlan AddServicePlansRequestServicePlan `json:"servicePlan"`
}

type _AddServicePlansRequest AddServicePlansRequest

// NewAddServicePlansRequest instantiates a new AddServicePlansRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddServicePlansRequest(servicePlan AddServicePlansRequestServicePlan) *AddServicePlansRequest {
	this := AddServicePlansRequest{}
	this.ServicePlan = servicePlan
	return &this
}

// NewAddServicePlansRequestWithDefaults instantiates a new AddServicePlansRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddServicePlansRequestWithDefaults() *AddServicePlansRequest {
	this := AddServicePlansRequest{}
	return &this
}

// GetServicePlan returns the ServicePlan field value
func (o *AddServicePlansRequest) GetServicePlan() AddServicePlansRequestServicePlan {
	if o == nil {
		var ret AddServicePlansRequestServicePlan
		return ret
	}

	return o.ServicePlan
}

// GetServicePlanOk returns a tuple with the ServicePlan field value
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequest) GetServicePlanOk() (*AddServicePlansRequestServicePlan, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServicePlan, true
}

// SetServicePlan sets field value
func (o *AddServicePlansRequest) SetServicePlan(v AddServicePlansRequestServicePlan) {
	o.ServicePlan = v
}

func (o AddServicePlansRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddServicePlansRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["servicePlan"] = o.ServicePlan
	return toSerialize, nil
}

func (o *AddServicePlansRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"servicePlan",
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

	varAddServicePlansRequest := _AddServicePlansRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddServicePlansRequest)

	if err != nil {
		return err
	}

	*o = AddServicePlansRequest(varAddServicePlansRequest)

	return err
}

type NullableAddServicePlansRequest struct {
	value *AddServicePlansRequest
	isSet bool
}

func (v NullableAddServicePlansRequest) Get() *AddServicePlansRequest {
	return v.value
}

func (v *NullableAddServicePlansRequest) Set(val *AddServicePlansRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddServicePlansRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddServicePlansRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddServicePlansRequest(val *AddServicePlansRequest) *NullableAddServicePlansRequest {
	return &NullableAddServicePlansRequest{value: val, isSet: true}
}

func (v NullableAddServicePlansRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddServicePlansRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



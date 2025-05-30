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

// checks if the UpdateIntegrationInventoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIntegrationInventoryRequest{}

// UpdateIntegrationInventoryRequest struct for UpdateIntegrationInventoryRequest
type UpdateIntegrationInventoryRequest struct {
	Inventory UpdateIntegrationInventoryRequestInventory `json:"inventory"`
}

type _UpdateIntegrationInventoryRequest UpdateIntegrationInventoryRequest

// NewUpdateIntegrationInventoryRequest instantiates a new UpdateIntegrationInventoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIntegrationInventoryRequest(inventory UpdateIntegrationInventoryRequestInventory) *UpdateIntegrationInventoryRequest {
	this := UpdateIntegrationInventoryRequest{}
	this.Inventory = inventory
	return &this
}

// NewUpdateIntegrationInventoryRequestWithDefaults instantiates a new UpdateIntegrationInventoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIntegrationInventoryRequestWithDefaults() *UpdateIntegrationInventoryRequest {
	this := UpdateIntegrationInventoryRequest{}
	return &this
}

// GetInventory returns the Inventory field value
func (o *UpdateIntegrationInventoryRequest) GetInventory() UpdateIntegrationInventoryRequestInventory {
	if o == nil {
		var ret UpdateIntegrationInventoryRequestInventory
		return ret
	}

	return o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value
// and a boolean to check if the value has been set.
func (o *UpdateIntegrationInventoryRequest) GetInventoryOk() (*UpdateIntegrationInventoryRequestInventory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inventory, true
}

// SetInventory sets field value
func (o *UpdateIntegrationInventoryRequest) SetInventory(v UpdateIntegrationInventoryRequestInventory) {
	o.Inventory = v
}

func (o UpdateIntegrationInventoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIntegrationInventoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inventory"] = o.Inventory
	return toSerialize, nil
}

func (o *UpdateIntegrationInventoryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inventory",
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

	varUpdateIntegrationInventoryRequest := _UpdateIntegrationInventoryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateIntegrationInventoryRequest)

	if err != nil {
		return err
	}

	*o = UpdateIntegrationInventoryRequest(varUpdateIntegrationInventoryRequest)

	return err
}

type NullableUpdateIntegrationInventoryRequest struct {
	value *UpdateIntegrationInventoryRequest
	isSet bool
}

func (v NullableUpdateIntegrationInventoryRequest) Get() *UpdateIntegrationInventoryRequest {
	return v.value
}

func (v *NullableUpdateIntegrationInventoryRequest) Set(val *UpdateIntegrationInventoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIntegrationInventoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIntegrationInventoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIntegrationInventoryRequest(val *UpdateIntegrationInventoryRequest) *NullableUpdateIntegrationInventoryRequest {
	return &NullableUpdateIntegrationInventoryRequest{value: val, isSet: true}
}

func (v NullableUpdateIntegrationInventoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIntegrationInventoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



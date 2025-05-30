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

// checks if the GenerateKeyPairsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateKeyPairsRequest{}

// GenerateKeyPairsRequest struct for GenerateKeyPairsRequest
type GenerateKeyPairsRequest struct {
	KeyPair GenerateKeyPairsRequestKeyPair `json:"keyPair"`
}

type _GenerateKeyPairsRequest GenerateKeyPairsRequest

// NewGenerateKeyPairsRequest instantiates a new GenerateKeyPairsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateKeyPairsRequest(keyPair GenerateKeyPairsRequestKeyPair) *GenerateKeyPairsRequest {
	this := GenerateKeyPairsRequest{}
	this.KeyPair = keyPair
	return &this
}

// NewGenerateKeyPairsRequestWithDefaults instantiates a new GenerateKeyPairsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateKeyPairsRequestWithDefaults() *GenerateKeyPairsRequest {
	this := GenerateKeyPairsRequest{}
	return &this
}

// GetKeyPair returns the KeyPair field value
func (o *GenerateKeyPairsRequest) GetKeyPair() GenerateKeyPairsRequestKeyPair {
	if o == nil {
		var ret GenerateKeyPairsRequestKeyPair
		return ret
	}

	return o.KeyPair
}

// GetKeyPairOk returns a tuple with the KeyPair field value
// and a boolean to check if the value has been set.
func (o *GenerateKeyPairsRequest) GetKeyPairOk() (*GenerateKeyPairsRequestKeyPair, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyPair, true
}

// SetKeyPair sets field value
func (o *GenerateKeyPairsRequest) SetKeyPair(v GenerateKeyPairsRequestKeyPair) {
	o.KeyPair = v
}

func (o GenerateKeyPairsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateKeyPairsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keyPair"] = o.KeyPair
	return toSerialize, nil
}

func (o *GenerateKeyPairsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"keyPair",
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

	varGenerateKeyPairsRequest := _GenerateKeyPairsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenerateKeyPairsRequest)

	if err != nil {
		return err
	}

	*o = GenerateKeyPairsRequest(varGenerateKeyPairsRequest)

	return err
}

type NullableGenerateKeyPairsRequest struct {
	value *GenerateKeyPairsRequest
	isSet bool
}

func (v NullableGenerateKeyPairsRequest) Get() *GenerateKeyPairsRequest {
	return v.value
}

func (v *NullableGenerateKeyPairsRequest) Set(val *GenerateKeyPairsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateKeyPairsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateKeyPairsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateKeyPairsRequest(val *GenerateKeyPairsRequest) *NullableGenerateKeyPairsRequest {
	return &NullableGenerateKeyPairsRequest{value: val, isSet: true}
}

func (v NullableGenerateKeyPairsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateKeyPairsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



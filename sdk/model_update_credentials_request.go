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

// checks if the UpdateCredentialsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCredentialsRequest{}

// UpdateCredentialsRequest struct for UpdateCredentialsRequest
type UpdateCredentialsRequest struct {
	Credential UpdateCredentialsRequestCredential `json:"credential"`
}

type _UpdateCredentialsRequest UpdateCredentialsRequest

// NewUpdateCredentialsRequest instantiates a new UpdateCredentialsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCredentialsRequest(credential UpdateCredentialsRequestCredential) *UpdateCredentialsRequest {
	this := UpdateCredentialsRequest{}
	this.Credential = credential
	return &this
}

// NewUpdateCredentialsRequestWithDefaults instantiates a new UpdateCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCredentialsRequestWithDefaults() *UpdateCredentialsRequest {
	this := UpdateCredentialsRequest{}
	return &this
}

// GetCredential returns the Credential field value
func (o *UpdateCredentialsRequest) GetCredential() UpdateCredentialsRequestCredential {
	if o == nil {
		var ret UpdateCredentialsRequestCredential
		return ret
	}

	return o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value
// and a boolean to check if the value has been set.
func (o *UpdateCredentialsRequest) GetCredentialOk() (*UpdateCredentialsRequestCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credential, true
}

// SetCredential sets field value
func (o *UpdateCredentialsRequest) SetCredential(v UpdateCredentialsRequestCredential) {
	o.Credential = v
}

func (o UpdateCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCredentialsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credential"] = o.Credential
	return toSerialize, nil
}

func (o *UpdateCredentialsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"credential",
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

	varUpdateCredentialsRequest := _UpdateCredentialsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateCredentialsRequest)

	if err != nil {
		return err
	}

	*o = UpdateCredentialsRequest(varUpdateCredentialsRequest)

	return err
}

type NullableUpdateCredentialsRequest struct {
	value *UpdateCredentialsRequest
	isSet bool
}

func (v NullableUpdateCredentialsRequest) Get() *UpdateCredentialsRequest {
	return v.value
}

func (v *NullableUpdateCredentialsRequest) Set(val *UpdateCredentialsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCredentialsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCredentialsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCredentialsRequest(val *UpdateCredentialsRequest) *NullableUpdateCredentialsRequest {
	return &NullableUpdateCredentialsRequest{value: val, isSet: true}
}

func (v NullableUpdateCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCredentialsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



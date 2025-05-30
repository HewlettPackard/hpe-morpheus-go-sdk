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

// checks if the AddLogSettingsSyslogRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddLogSettingsSyslogRulesRequest{}

// AddLogSettingsSyslogRulesRequest struct for AddLogSettingsSyslogRulesRequest
type AddLogSettingsSyslogRulesRequest struct {
	SyslogRule AddLogSettingsSyslogRulesRequestSyslogRule `json:"syslogRule"`
}

type _AddLogSettingsSyslogRulesRequest AddLogSettingsSyslogRulesRequest

// NewAddLogSettingsSyslogRulesRequest instantiates a new AddLogSettingsSyslogRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLogSettingsSyslogRulesRequest(syslogRule AddLogSettingsSyslogRulesRequestSyslogRule) *AddLogSettingsSyslogRulesRequest {
	this := AddLogSettingsSyslogRulesRequest{}
	this.SyslogRule = syslogRule
	return &this
}

// NewAddLogSettingsSyslogRulesRequestWithDefaults instantiates a new AddLogSettingsSyslogRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLogSettingsSyslogRulesRequestWithDefaults() *AddLogSettingsSyslogRulesRequest {
	this := AddLogSettingsSyslogRulesRequest{}
	return &this
}

// GetSyslogRule returns the SyslogRule field value
func (o *AddLogSettingsSyslogRulesRequest) GetSyslogRule() AddLogSettingsSyslogRulesRequestSyslogRule {
	if o == nil {
		var ret AddLogSettingsSyslogRulesRequestSyslogRule
		return ret
	}

	return o.SyslogRule
}

// GetSyslogRuleOk returns a tuple with the SyslogRule field value
// and a boolean to check if the value has been set.
func (o *AddLogSettingsSyslogRulesRequest) GetSyslogRuleOk() (*AddLogSettingsSyslogRulesRequestSyslogRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyslogRule, true
}

// SetSyslogRule sets field value
func (o *AddLogSettingsSyslogRulesRequest) SetSyslogRule(v AddLogSettingsSyslogRulesRequestSyslogRule) {
	o.SyslogRule = v
}

func (o AddLogSettingsSyslogRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddLogSettingsSyslogRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["syslogRule"] = o.SyslogRule
	return toSerialize, nil
}

func (o *AddLogSettingsSyslogRulesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"syslogRule",
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

	varAddLogSettingsSyslogRulesRequest := _AddLogSettingsSyslogRulesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddLogSettingsSyslogRulesRequest)

	if err != nil {
		return err
	}

	*o = AddLogSettingsSyslogRulesRequest(varAddLogSettingsSyslogRulesRequest)

	return err
}

type NullableAddLogSettingsSyslogRulesRequest struct {
	value *AddLogSettingsSyslogRulesRequest
	isSet bool
}

func (v NullableAddLogSettingsSyslogRulesRequest) Get() *AddLogSettingsSyslogRulesRequest {
	return v.value
}

func (v *NullableAddLogSettingsSyslogRulesRequest) Set(val *AddLogSettingsSyslogRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLogSettingsSyslogRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLogSettingsSyslogRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLogSettingsSyslogRulesRequest(val *AddLogSettingsSyslogRulesRequest) *NullableAddLogSettingsSyslogRulesRequest {
	return &NullableAddLogSettingsSyslogRulesRequest{value: val, isSet: true}
}

func (v NullableAddLogSettingsSyslogRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLogSettingsSyslogRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

	"gopkg.in/validator.v2"
)

// very silly way of avoiding `"fmt" imported and not used` errors
var _ fmt.Stringer

// AddClusterRequestClusterType - struct for AddClusterRequestClusterType
type AddClusterRequestClusterType struct {
	AddClusterRequestClusterTypeOneOf *AddClusterRequestClusterTypeOneOf
	String                            *string
}

// AddClusterRequestClusterTypeOneOfAsAddClusterRequestClusterType is a convenience function that returns AddClusterRequestClusterTypeOneOf wrapped in AddClusterRequestClusterType
func AddClusterRequestClusterTypeOneOfAsAddClusterRequestClusterType(v *AddClusterRequestClusterTypeOneOf) AddClusterRequestClusterType {
	return AddClusterRequestClusterType{
		AddClusterRequestClusterTypeOneOf: v,
	}
}

// stringAsAddClusterRequestClusterType is a convenience function that returns string wrapped in AddClusterRequestClusterType
func StringAsAddClusterRequestClusterType(v *string) AddClusterRequestClusterType {
	return AddClusterRequestClusterType{
		String: v,
	}
}

func (dst *AddClusterRequestClusterType) UnmarshalMapstructure(data any) (any, error) {
	if dst == nil {
		dst = &AddClusterRequestClusterType{}
	}

	if out, ok := data.(AddClusterRequestClusterTypeOneOf); ok {
		dst.AddClusterRequestClusterTypeOneOf = &out
	}

	if out, ok := data.(string); ok {
		dst.String = &out
	}

	return dst, nil
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddClusterRequestClusterType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddClusterRequestClusterTypeOneOf
	err = newStrictDecoder(data).Decode(&dst.AddClusterRequestClusterTypeOneOf)
	if err == nil {
		jsonAddClusterRequestClusterTypeOneOf, _ := json.Marshal(dst.AddClusterRequestClusterTypeOneOf)
		if string(jsonAddClusterRequestClusterTypeOneOf) == "{}" { // empty struct
			dst.AddClusterRequestClusterTypeOneOf = nil
		} else {
			if err = validator.Validate(dst.AddClusterRequestClusterTypeOneOf); err != nil {
				dst.AddClusterRequestClusterTypeOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddClusterRequestClusterTypeOneOf = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddClusterRequestClusterTypeOneOf = nil
		dst.String = nil

		return NewResponseValidationError("data matches more than one schema in oneOf(AddClusterRequestClusterType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return NewResponseValidationError("data failed to match schemas in oneOf(AddClusterRequestClusterType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddClusterRequestClusterType) MarshalJSON() ([]byte, error) {
	if src.AddClusterRequestClusterTypeOneOf != nil {
		return json.Marshal(&src.AddClusterRequestClusterTypeOneOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddClusterRequestClusterType) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddClusterRequestClusterTypeOneOf != nil {
		return obj.AddClusterRequestClusterTypeOneOf
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AddClusterRequestClusterType) GetActualInstanceValue() interface{} {
	if obj.AddClusterRequestClusterTypeOneOf != nil {
		return *obj.AddClusterRequestClusterTypeOneOf
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableAddClusterRequestClusterType struct {
	value *AddClusterRequestClusterType
	isSet bool
}

func (v NullableAddClusterRequestClusterType) Get() *AddClusterRequestClusterType {
	return v.value
}

func (v *NullableAddClusterRequestClusterType) Set(val *AddClusterRequestClusterType) {
	v.value = val
	v.isSet = true
}

func (v NullableAddClusterRequestClusterType) IsSet() bool {
	return v.isSet
}

func (v *NullableAddClusterRequestClusterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddClusterRequestClusterType(val *AddClusterRequestClusterType) *NullableAddClusterRequestClusterType {
	return &NullableAddClusterRequestClusterType{value: val, isSet: true}
}

func (v NullableAddClusterRequestClusterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddClusterRequestClusterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

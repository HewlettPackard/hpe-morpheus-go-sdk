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

// AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId - struct for AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId
type AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId struct {
	Int64  *int64
	String *string
}

// int64AsAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId is a convenience function that returns int64 wrapped in AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId
func Int64AsAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId(v *int64) AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId {
	return AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId{
		Int64: v,
	}
}

// stringAsAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId is a convenience function that returns string wrapped in AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId
func StringAsAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId(v *string) AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId {
	return AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId{
		String: v,
	}
}

func (dst *AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) UnmarshalMapstructure(data any) (any, error) {
	if dst == nil {
		dst = &AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId{}
	}

	if out, ok := data.(int64); ok {
		dst.Int64 = &out
	}

	if out, ok := data.(string); ok {
		dst.String = &out
	}

	return dst, nil
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int64
	err = newStrictDecoder(data).Decode(&dst.Int64)
	if err == nil {
		jsonInt64, _ := json.Marshal(dst.Int64)
		if string(jsonInt64) == "{}" { // empty struct
			dst.Int64 = nil
		} else {
			if err = validator.Validate(dst.Int64); err != nil {
				dst.Int64 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int64 = nil
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
		dst.Int64 = nil
		dst.String = nil

		return NewResponseValidationError("data matches more than one schema in oneOf(AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return NewResponseValidationError("data failed to match schemas in oneOf(AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Int64 != nil {
		return obj.Int64
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) GetActualInstanceValue() interface{} {
	if obj.Int64 != nil {
		return *obj.Int64
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId struct {
	value *AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId
	isSet bool
}

func (v NullableAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) Get() *AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId {
	return v.value
}

func (v *NullableAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) Set(val *AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) {
	v.value = val
	v.isSet = true
}

func (v NullableAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) IsSet() bool {
	return v.isSet
}

func (v *NullableAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId(val *AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) *NullableAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId {
	return &NullableAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId{value: val, isSet: true}
}

func (v NullableAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

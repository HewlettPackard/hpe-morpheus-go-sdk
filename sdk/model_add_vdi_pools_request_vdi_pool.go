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

// AddVDIPoolsRequestVdiPool - struct for AddVDIPoolsRequestVdiPool
type AddVDIPoolsRequestVdiPool struct {
	AddVDIPoolsRequestVdiPoolOneOf  *AddVDIPoolsRequestVdiPoolOneOf
	AddVDIPoolsRequestVdiPoolOneOf1 *AddVDIPoolsRequestVdiPoolOneOf1
}

// AddVDIPoolsRequestVdiPoolOneOfAsAddVDIPoolsRequestVdiPool is a convenience function that returns AddVDIPoolsRequestVdiPoolOneOf wrapped in AddVDIPoolsRequestVdiPool
func AddVDIPoolsRequestVdiPoolOneOfAsAddVDIPoolsRequestVdiPool(v *AddVDIPoolsRequestVdiPoolOneOf) AddVDIPoolsRequestVdiPool {
	return AddVDIPoolsRequestVdiPool{
		AddVDIPoolsRequestVdiPoolOneOf: v,
	}
}

// AddVDIPoolsRequestVdiPoolOneOf1AsAddVDIPoolsRequestVdiPool is a convenience function that returns AddVDIPoolsRequestVdiPoolOneOf1 wrapped in AddVDIPoolsRequestVdiPool
func AddVDIPoolsRequestVdiPoolOneOf1AsAddVDIPoolsRequestVdiPool(v *AddVDIPoolsRequestVdiPoolOneOf1) AddVDIPoolsRequestVdiPool {
	return AddVDIPoolsRequestVdiPool{
		AddVDIPoolsRequestVdiPoolOneOf1: v,
	}
}

func (dst *AddVDIPoolsRequestVdiPool) UnmarshalMapstructure(data any) (any, error) {
	if dst == nil {
		dst = &AddVDIPoolsRequestVdiPool{}
	}

	if out, ok := data.(AddVDIPoolsRequestVdiPoolOneOf); ok {
		dst.AddVDIPoolsRequestVdiPoolOneOf = &out
	}

	if out, ok := data.(AddVDIPoolsRequestVdiPoolOneOf1); ok {
		dst.AddVDIPoolsRequestVdiPoolOneOf1 = &out
	}

	return dst, nil
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddVDIPoolsRequestVdiPool) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddVDIPoolsRequestVdiPoolOneOf
	err = newStrictDecoder(data).Decode(&dst.AddVDIPoolsRequestVdiPoolOneOf)
	if err == nil {
		jsonAddVDIPoolsRequestVdiPoolOneOf, _ := json.Marshal(dst.AddVDIPoolsRequestVdiPoolOneOf)
		if string(jsonAddVDIPoolsRequestVdiPoolOneOf) == "{}" { // empty struct
			dst.AddVDIPoolsRequestVdiPoolOneOf = nil
		} else {
			if err = validator.Validate(dst.AddVDIPoolsRequestVdiPoolOneOf); err != nil {
				dst.AddVDIPoolsRequestVdiPoolOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddVDIPoolsRequestVdiPoolOneOf = nil
	}

	// try to unmarshal data into AddVDIPoolsRequestVdiPoolOneOf1
	err = newStrictDecoder(data).Decode(&dst.AddVDIPoolsRequestVdiPoolOneOf1)
	if err == nil {
		jsonAddVDIPoolsRequestVdiPoolOneOf1, _ := json.Marshal(dst.AddVDIPoolsRequestVdiPoolOneOf1)
		if string(jsonAddVDIPoolsRequestVdiPoolOneOf1) == "{}" { // empty struct
			dst.AddVDIPoolsRequestVdiPoolOneOf1 = nil
		} else {
			if err = validator.Validate(dst.AddVDIPoolsRequestVdiPoolOneOf1); err != nil {
				dst.AddVDIPoolsRequestVdiPoolOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddVDIPoolsRequestVdiPoolOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddVDIPoolsRequestVdiPoolOneOf = nil
		dst.AddVDIPoolsRequestVdiPoolOneOf1 = nil

		return NewResponseValidationError("data matches more than one schema in oneOf(AddVDIPoolsRequestVdiPool)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return NewResponseValidationError("data failed to match schemas in oneOf(AddVDIPoolsRequestVdiPool)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddVDIPoolsRequestVdiPool) MarshalJSON() ([]byte, error) {
	if src.AddVDIPoolsRequestVdiPoolOneOf != nil {
		return json.Marshal(&src.AddVDIPoolsRequestVdiPoolOneOf)
	}

	if src.AddVDIPoolsRequestVdiPoolOneOf1 != nil {
		return json.Marshal(&src.AddVDIPoolsRequestVdiPoolOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddVDIPoolsRequestVdiPool) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddVDIPoolsRequestVdiPoolOneOf != nil {
		return obj.AddVDIPoolsRequestVdiPoolOneOf
	}

	if obj.AddVDIPoolsRequestVdiPoolOneOf1 != nil {
		return obj.AddVDIPoolsRequestVdiPoolOneOf1
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AddVDIPoolsRequestVdiPool) GetActualInstanceValue() interface{} {
	if obj.AddVDIPoolsRequestVdiPoolOneOf != nil {
		return *obj.AddVDIPoolsRequestVdiPoolOneOf
	}

	if obj.AddVDIPoolsRequestVdiPoolOneOf1 != nil {
		return *obj.AddVDIPoolsRequestVdiPoolOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableAddVDIPoolsRequestVdiPool struct {
	value *AddVDIPoolsRequestVdiPool
	isSet bool
}

func (v NullableAddVDIPoolsRequestVdiPool) Get() *AddVDIPoolsRequestVdiPool {
	return v.value
}

func (v *NullableAddVDIPoolsRequestVdiPool) Set(val *AddVDIPoolsRequestVdiPool) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVDIPoolsRequestVdiPool) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVDIPoolsRequestVdiPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVDIPoolsRequestVdiPool(val *AddVDIPoolsRequestVdiPool) *NullableAddVDIPoolsRequestVdiPool {
	return &NullableAddVDIPoolsRequestVdiPool{value: val, isSet: true}
}

func (v NullableAddVDIPoolsRequestVdiPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVDIPoolsRequestVdiPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

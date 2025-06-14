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

// AddBlueprintRequestOneOfArmInstallAgent - Install Morpheus Agent
type AddBlueprintRequestOneOfArmInstallAgent struct {
	Bool   *bool
	String *string
}

// boolAsAddBlueprintRequestOneOfArmInstallAgent is a convenience function that returns bool wrapped in AddBlueprintRequestOneOfArmInstallAgent
func BoolAsAddBlueprintRequestOneOfArmInstallAgent(v *bool) AddBlueprintRequestOneOfArmInstallAgent {
	return AddBlueprintRequestOneOfArmInstallAgent{
		Bool: v,
	}
}

// stringAsAddBlueprintRequestOneOfArmInstallAgent is a convenience function that returns string wrapped in AddBlueprintRequestOneOfArmInstallAgent
func StringAsAddBlueprintRequestOneOfArmInstallAgent(v *string) AddBlueprintRequestOneOfArmInstallAgent {
	return AddBlueprintRequestOneOfArmInstallAgent{
		String: v,
	}
}

func (dst *AddBlueprintRequestOneOfArmInstallAgent) UnmarshalMapstructure(data any) (any, error) {
	if dst == nil {
		dst = &AddBlueprintRequestOneOfArmInstallAgent{}
	}

	if out, ok := data.(bool); ok {
		dst.Bool = &out
	}

	if out, ok := data.(string); ok {
		dst.String = &out
	}

	return dst, nil
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddBlueprintRequestOneOfArmInstallAgent) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			if err = validator.Validate(dst.Bool); err != nil {
				dst.Bool = nil
			} else {
				match++
			}
		}
	} else {
		dst.Bool = nil
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
		dst.Bool = nil
		dst.String = nil

		return NewResponseValidationError("data matches more than one schema in oneOf(AddBlueprintRequestOneOfArmInstallAgent)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return NewResponseValidationError("data failed to match schemas in oneOf(AddBlueprintRequestOneOfArmInstallAgent)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddBlueprintRequestOneOfArmInstallAgent) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddBlueprintRequestOneOfArmInstallAgent) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AddBlueprintRequestOneOfArmInstallAgent) GetActualInstanceValue() interface{} {
	if obj.Bool != nil {
		return *obj.Bool
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableAddBlueprintRequestOneOfArmInstallAgent struct {
	value *AddBlueprintRequestOneOfArmInstallAgent
	isSet bool
}

func (v NullableAddBlueprintRequestOneOfArmInstallAgent) Get() *AddBlueprintRequestOneOfArmInstallAgent {
	return v.value
}

func (v *NullableAddBlueprintRequestOneOfArmInstallAgent) Set(val *AddBlueprintRequestOneOfArmInstallAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableAddBlueprintRequestOneOfArmInstallAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAddBlueprintRequestOneOfArmInstallAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddBlueprintRequestOneOfArmInstallAgent(val *AddBlueprintRequestOneOfArmInstallAgent) *NullableAddBlueprintRequestOneOfArmInstallAgent {
	return &NullableAddBlueprintRequestOneOfArmInstallAgent{value: val, isSet: true}
}

func (v NullableAddBlueprintRequestOneOfArmInstallAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddBlueprintRequestOneOfArmInstallAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

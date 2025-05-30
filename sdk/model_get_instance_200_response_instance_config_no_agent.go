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

// GetInstance200ResponseInstanceConfigNoAgent - struct for GetInstance200ResponseInstanceConfigNoAgent
type GetInstance200ResponseInstanceConfigNoAgent struct {
	Bool *bool
	String *string
}

// boolAsGetInstance200ResponseInstanceConfigNoAgent is a convenience function that returns bool wrapped in GetInstance200ResponseInstanceConfigNoAgent
func BoolAsGetInstance200ResponseInstanceConfigNoAgent(v *bool) GetInstance200ResponseInstanceConfigNoAgent {
	return GetInstance200ResponseInstanceConfigNoAgent{
		Bool: v,
	}
}

// stringAsGetInstance200ResponseInstanceConfigNoAgent is a convenience function that returns string wrapped in GetInstance200ResponseInstanceConfigNoAgent
func StringAsGetInstance200ResponseInstanceConfigNoAgent(v *string) GetInstance200ResponseInstanceConfigNoAgent {
	return GetInstance200ResponseInstanceConfigNoAgent{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetInstance200ResponseInstanceConfigNoAgent) UnmarshalJSON(data []byte) error {
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

		return NewResponseValidationError("data matches more than one schema in oneOf(GetInstance200ResponseInstanceConfigNoAgent)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return NewResponseValidationError("data failed to match schemas in oneOf(GetInstance200ResponseInstanceConfigNoAgent)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetInstance200ResponseInstanceConfigNoAgent) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetInstance200ResponseInstanceConfigNoAgent) GetActualInstance() (interface{}) {
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
func (obj GetInstance200ResponseInstanceConfigNoAgent) GetActualInstanceValue() (interface{}) {
	if obj.Bool != nil {
		return *obj.Bool
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableGetInstance200ResponseInstanceConfigNoAgent struct {
	value *GetInstance200ResponseInstanceConfigNoAgent
	isSet bool
}

func (v NullableGetInstance200ResponseInstanceConfigNoAgent) Get() *GetInstance200ResponseInstanceConfigNoAgent {
	return v.value
}

func (v *NullableGetInstance200ResponseInstanceConfigNoAgent) Set(val *GetInstance200ResponseInstanceConfigNoAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInstance200ResponseInstanceConfigNoAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInstance200ResponseInstanceConfigNoAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInstance200ResponseInstanceConfigNoAgent(val *GetInstance200ResponseInstanceConfigNoAgent) *NullableGetInstance200ResponseInstanceConfigNoAgent {
	return &NullableGetInstance200ResponseInstanceConfigNoAgent{value: val, isSet: true}
}

func (v NullableGetInstance200ResponseInstanceConfigNoAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInstance200ResponseInstanceConfigNoAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// UpdateCredentialsRequestCredential - Payload for updating a credential
type UpdateCredentialsRequestCredential struct {
	AddCredentialsRequestCredentialOneOf  *AddCredentialsRequestCredentialOneOf
	AddCredentialsRequestCredentialOneOf1 *AddCredentialsRequestCredentialOneOf1
	AddCredentialsRequestCredentialOneOf2 *AddCredentialsRequestCredentialOneOf2
	AddCredentialsRequestCredentialOneOf3 *AddCredentialsRequestCredentialOneOf3
	AddCredentialsRequestCredentialOneOf4 *AddCredentialsRequestCredentialOneOf4
	AddCredentialsRequestCredentialOneOf5 *AddCredentialsRequestCredentialOneOf5
	AddCredentialsRequestCredentialOneOf6 *AddCredentialsRequestCredentialOneOf6
	AddCredentialsRequestCredentialOneOf7 *AddCredentialsRequestCredentialOneOf7
	AddCredentialsRequestCredentialOneOf8 *AddCredentialsRequestCredentialOneOf8
}

// AddCredentialsRequestCredentialOneOfAsUpdateCredentialsRequestCredential is a convenience function that returns AddCredentialsRequestCredentialOneOf wrapped in UpdateCredentialsRequestCredential
func AddCredentialsRequestCredentialOneOfAsUpdateCredentialsRequestCredential(v *AddCredentialsRequestCredentialOneOf) UpdateCredentialsRequestCredential {
	return UpdateCredentialsRequestCredential{
		AddCredentialsRequestCredentialOneOf: v,
	}
}

// AddCredentialsRequestCredentialOneOf1AsUpdateCredentialsRequestCredential is a convenience function that returns AddCredentialsRequestCredentialOneOf1 wrapped in UpdateCredentialsRequestCredential
func AddCredentialsRequestCredentialOneOf1AsUpdateCredentialsRequestCredential(v *AddCredentialsRequestCredentialOneOf1) UpdateCredentialsRequestCredential {
	return UpdateCredentialsRequestCredential{
		AddCredentialsRequestCredentialOneOf1: v,
	}
}

// AddCredentialsRequestCredentialOneOf2AsUpdateCredentialsRequestCredential is a convenience function that returns AddCredentialsRequestCredentialOneOf2 wrapped in UpdateCredentialsRequestCredential
func AddCredentialsRequestCredentialOneOf2AsUpdateCredentialsRequestCredential(v *AddCredentialsRequestCredentialOneOf2) UpdateCredentialsRequestCredential {
	return UpdateCredentialsRequestCredential{
		AddCredentialsRequestCredentialOneOf2: v,
	}
}

// AddCredentialsRequestCredentialOneOf3AsUpdateCredentialsRequestCredential is a convenience function that returns AddCredentialsRequestCredentialOneOf3 wrapped in UpdateCredentialsRequestCredential
func AddCredentialsRequestCredentialOneOf3AsUpdateCredentialsRequestCredential(v *AddCredentialsRequestCredentialOneOf3) UpdateCredentialsRequestCredential {
	return UpdateCredentialsRequestCredential{
		AddCredentialsRequestCredentialOneOf3: v,
	}
}

// AddCredentialsRequestCredentialOneOf4AsUpdateCredentialsRequestCredential is a convenience function that returns AddCredentialsRequestCredentialOneOf4 wrapped in UpdateCredentialsRequestCredential
func AddCredentialsRequestCredentialOneOf4AsUpdateCredentialsRequestCredential(v *AddCredentialsRequestCredentialOneOf4) UpdateCredentialsRequestCredential {
	return UpdateCredentialsRequestCredential{
		AddCredentialsRequestCredentialOneOf4: v,
	}
}

// AddCredentialsRequestCredentialOneOf5AsUpdateCredentialsRequestCredential is a convenience function that returns AddCredentialsRequestCredentialOneOf5 wrapped in UpdateCredentialsRequestCredential
func AddCredentialsRequestCredentialOneOf5AsUpdateCredentialsRequestCredential(v *AddCredentialsRequestCredentialOneOf5) UpdateCredentialsRequestCredential {
	return UpdateCredentialsRequestCredential{
		AddCredentialsRequestCredentialOneOf5: v,
	}
}

// AddCredentialsRequestCredentialOneOf6AsUpdateCredentialsRequestCredential is a convenience function that returns AddCredentialsRequestCredentialOneOf6 wrapped in UpdateCredentialsRequestCredential
func AddCredentialsRequestCredentialOneOf6AsUpdateCredentialsRequestCredential(v *AddCredentialsRequestCredentialOneOf6) UpdateCredentialsRequestCredential {
	return UpdateCredentialsRequestCredential{
		AddCredentialsRequestCredentialOneOf6: v,
	}
}

// AddCredentialsRequestCredentialOneOf7AsUpdateCredentialsRequestCredential is a convenience function that returns AddCredentialsRequestCredentialOneOf7 wrapped in UpdateCredentialsRequestCredential
func AddCredentialsRequestCredentialOneOf7AsUpdateCredentialsRequestCredential(v *AddCredentialsRequestCredentialOneOf7) UpdateCredentialsRequestCredential {
	return UpdateCredentialsRequestCredential{
		AddCredentialsRequestCredentialOneOf7: v,
	}
}

// AddCredentialsRequestCredentialOneOf8AsUpdateCredentialsRequestCredential is a convenience function that returns AddCredentialsRequestCredentialOneOf8 wrapped in UpdateCredentialsRequestCredential
func AddCredentialsRequestCredentialOneOf8AsUpdateCredentialsRequestCredential(v *AddCredentialsRequestCredentialOneOf8) UpdateCredentialsRequestCredential {
	return UpdateCredentialsRequestCredential{
		AddCredentialsRequestCredentialOneOf8: v,
	}
}

func (dst *UpdateCredentialsRequestCredential) UnmarshalMapstructure(data any) (any, error) {
	if dst == nil {
		dst = &UpdateCredentialsRequestCredential{}
	}

	if out, ok := data.(AddCredentialsRequestCredentialOneOf); ok {
		dst.AddCredentialsRequestCredentialOneOf = &out
	}

	if out, ok := data.(AddCredentialsRequestCredentialOneOf1); ok {
		dst.AddCredentialsRequestCredentialOneOf1 = &out
	}

	if out, ok := data.(AddCredentialsRequestCredentialOneOf2); ok {
		dst.AddCredentialsRequestCredentialOneOf2 = &out
	}

	if out, ok := data.(AddCredentialsRequestCredentialOneOf3); ok {
		dst.AddCredentialsRequestCredentialOneOf3 = &out
	}

	if out, ok := data.(AddCredentialsRequestCredentialOneOf4); ok {
		dst.AddCredentialsRequestCredentialOneOf4 = &out
	}

	if out, ok := data.(AddCredentialsRequestCredentialOneOf5); ok {
		dst.AddCredentialsRequestCredentialOneOf5 = &out
	}

	if out, ok := data.(AddCredentialsRequestCredentialOneOf6); ok {
		dst.AddCredentialsRequestCredentialOneOf6 = &out
	}

	if out, ok := data.(AddCredentialsRequestCredentialOneOf7); ok {
		dst.AddCredentialsRequestCredentialOneOf7 = &out
	}

	if out, ok := data.(AddCredentialsRequestCredentialOneOf8); ok {
		dst.AddCredentialsRequestCredentialOneOf8 = &out
	}

	return dst, nil
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateCredentialsRequestCredential) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddCredentialsRequestCredentialOneOf
	err = newStrictDecoder(data).Decode(&dst.AddCredentialsRequestCredentialOneOf)
	if err == nil {
		jsonAddCredentialsRequestCredentialOneOf, _ := json.Marshal(dst.AddCredentialsRequestCredentialOneOf)
		if string(jsonAddCredentialsRequestCredentialOneOf) == "{}" { // empty struct
			dst.AddCredentialsRequestCredentialOneOf = nil
		} else {
			if err = validator.Validate(dst.AddCredentialsRequestCredentialOneOf); err != nil {
				dst.AddCredentialsRequestCredentialOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddCredentialsRequestCredentialOneOf = nil
	}

	// try to unmarshal data into AddCredentialsRequestCredentialOneOf1
	err = newStrictDecoder(data).Decode(&dst.AddCredentialsRequestCredentialOneOf1)
	if err == nil {
		jsonAddCredentialsRequestCredentialOneOf1, _ := json.Marshal(dst.AddCredentialsRequestCredentialOneOf1)
		if string(jsonAddCredentialsRequestCredentialOneOf1) == "{}" { // empty struct
			dst.AddCredentialsRequestCredentialOneOf1 = nil
		} else {
			if err = validator.Validate(dst.AddCredentialsRequestCredentialOneOf1); err != nil {
				dst.AddCredentialsRequestCredentialOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddCredentialsRequestCredentialOneOf1 = nil
	}

	// try to unmarshal data into AddCredentialsRequestCredentialOneOf2
	err = newStrictDecoder(data).Decode(&dst.AddCredentialsRequestCredentialOneOf2)
	if err == nil {
		jsonAddCredentialsRequestCredentialOneOf2, _ := json.Marshal(dst.AddCredentialsRequestCredentialOneOf2)
		if string(jsonAddCredentialsRequestCredentialOneOf2) == "{}" { // empty struct
			dst.AddCredentialsRequestCredentialOneOf2 = nil
		} else {
			if err = validator.Validate(dst.AddCredentialsRequestCredentialOneOf2); err != nil {
				dst.AddCredentialsRequestCredentialOneOf2 = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddCredentialsRequestCredentialOneOf2 = nil
	}

	// try to unmarshal data into AddCredentialsRequestCredentialOneOf3
	err = newStrictDecoder(data).Decode(&dst.AddCredentialsRequestCredentialOneOf3)
	if err == nil {
		jsonAddCredentialsRequestCredentialOneOf3, _ := json.Marshal(dst.AddCredentialsRequestCredentialOneOf3)
		if string(jsonAddCredentialsRequestCredentialOneOf3) == "{}" { // empty struct
			dst.AddCredentialsRequestCredentialOneOf3 = nil
		} else {
			if err = validator.Validate(dst.AddCredentialsRequestCredentialOneOf3); err != nil {
				dst.AddCredentialsRequestCredentialOneOf3 = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddCredentialsRequestCredentialOneOf3 = nil
	}

	// try to unmarshal data into AddCredentialsRequestCredentialOneOf4
	err = newStrictDecoder(data).Decode(&dst.AddCredentialsRequestCredentialOneOf4)
	if err == nil {
		jsonAddCredentialsRequestCredentialOneOf4, _ := json.Marshal(dst.AddCredentialsRequestCredentialOneOf4)
		if string(jsonAddCredentialsRequestCredentialOneOf4) == "{}" { // empty struct
			dst.AddCredentialsRequestCredentialOneOf4 = nil
		} else {
			if err = validator.Validate(dst.AddCredentialsRequestCredentialOneOf4); err != nil {
				dst.AddCredentialsRequestCredentialOneOf4 = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddCredentialsRequestCredentialOneOf4 = nil
	}

	// try to unmarshal data into AddCredentialsRequestCredentialOneOf5
	err = newStrictDecoder(data).Decode(&dst.AddCredentialsRequestCredentialOneOf5)
	if err == nil {
		jsonAddCredentialsRequestCredentialOneOf5, _ := json.Marshal(dst.AddCredentialsRequestCredentialOneOf5)
		if string(jsonAddCredentialsRequestCredentialOneOf5) == "{}" { // empty struct
			dst.AddCredentialsRequestCredentialOneOf5 = nil
		} else {
			if err = validator.Validate(dst.AddCredentialsRequestCredentialOneOf5); err != nil {
				dst.AddCredentialsRequestCredentialOneOf5 = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddCredentialsRequestCredentialOneOf5 = nil
	}

	// try to unmarshal data into AddCredentialsRequestCredentialOneOf6
	err = newStrictDecoder(data).Decode(&dst.AddCredentialsRequestCredentialOneOf6)
	if err == nil {
		jsonAddCredentialsRequestCredentialOneOf6, _ := json.Marshal(dst.AddCredentialsRequestCredentialOneOf6)
		if string(jsonAddCredentialsRequestCredentialOneOf6) == "{}" { // empty struct
			dst.AddCredentialsRequestCredentialOneOf6 = nil
		} else {
			if err = validator.Validate(dst.AddCredentialsRequestCredentialOneOf6); err != nil {
				dst.AddCredentialsRequestCredentialOneOf6 = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddCredentialsRequestCredentialOneOf6 = nil
	}

	// try to unmarshal data into AddCredentialsRequestCredentialOneOf7
	err = newStrictDecoder(data).Decode(&dst.AddCredentialsRequestCredentialOneOf7)
	if err == nil {
		jsonAddCredentialsRequestCredentialOneOf7, _ := json.Marshal(dst.AddCredentialsRequestCredentialOneOf7)
		if string(jsonAddCredentialsRequestCredentialOneOf7) == "{}" { // empty struct
			dst.AddCredentialsRequestCredentialOneOf7 = nil
		} else {
			if err = validator.Validate(dst.AddCredentialsRequestCredentialOneOf7); err != nil {
				dst.AddCredentialsRequestCredentialOneOf7 = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddCredentialsRequestCredentialOneOf7 = nil
	}

	// try to unmarshal data into AddCredentialsRequestCredentialOneOf8
	err = newStrictDecoder(data).Decode(&dst.AddCredentialsRequestCredentialOneOf8)
	if err == nil {
		jsonAddCredentialsRequestCredentialOneOf8, _ := json.Marshal(dst.AddCredentialsRequestCredentialOneOf8)
		if string(jsonAddCredentialsRequestCredentialOneOf8) == "{}" { // empty struct
			dst.AddCredentialsRequestCredentialOneOf8 = nil
		} else {
			if err = validator.Validate(dst.AddCredentialsRequestCredentialOneOf8); err != nil {
				dst.AddCredentialsRequestCredentialOneOf8 = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddCredentialsRequestCredentialOneOf8 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddCredentialsRequestCredentialOneOf = nil
		dst.AddCredentialsRequestCredentialOneOf1 = nil
		dst.AddCredentialsRequestCredentialOneOf2 = nil
		dst.AddCredentialsRequestCredentialOneOf3 = nil
		dst.AddCredentialsRequestCredentialOneOf4 = nil
		dst.AddCredentialsRequestCredentialOneOf5 = nil
		dst.AddCredentialsRequestCredentialOneOf6 = nil
		dst.AddCredentialsRequestCredentialOneOf7 = nil
		dst.AddCredentialsRequestCredentialOneOf8 = nil

		return NewResponseValidationError("data matches more than one schema in oneOf(UpdateCredentialsRequestCredential)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return NewResponseValidationError("data failed to match schemas in oneOf(UpdateCredentialsRequestCredential)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateCredentialsRequestCredential) MarshalJSON() ([]byte, error) {
	if src.AddCredentialsRequestCredentialOneOf != nil {
		return json.Marshal(&src.AddCredentialsRequestCredentialOneOf)
	}

	if src.AddCredentialsRequestCredentialOneOf1 != nil {
		return json.Marshal(&src.AddCredentialsRequestCredentialOneOf1)
	}

	if src.AddCredentialsRequestCredentialOneOf2 != nil {
		return json.Marshal(&src.AddCredentialsRequestCredentialOneOf2)
	}

	if src.AddCredentialsRequestCredentialOneOf3 != nil {
		return json.Marshal(&src.AddCredentialsRequestCredentialOneOf3)
	}

	if src.AddCredentialsRequestCredentialOneOf4 != nil {
		return json.Marshal(&src.AddCredentialsRequestCredentialOneOf4)
	}

	if src.AddCredentialsRequestCredentialOneOf5 != nil {
		return json.Marshal(&src.AddCredentialsRequestCredentialOneOf5)
	}

	if src.AddCredentialsRequestCredentialOneOf6 != nil {
		return json.Marshal(&src.AddCredentialsRequestCredentialOneOf6)
	}

	if src.AddCredentialsRequestCredentialOneOf7 != nil {
		return json.Marshal(&src.AddCredentialsRequestCredentialOneOf7)
	}

	if src.AddCredentialsRequestCredentialOneOf8 != nil {
		return json.Marshal(&src.AddCredentialsRequestCredentialOneOf8)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateCredentialsRequestCredential) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddCredentialsRequestCredentialOneOf != nil {
		return obj.AddCredentialsRequestCredentialOneOf
	}

	if obj.AddCredentialsRequestCredentialOneOf1 != nil {
		return obj.AddCredentialsRequestCredentialOneOf1
	}

	if obj.AddCredentialsRequestCredentialOneOf2 != nil {
		return obj.AddCredentialsRequestCredentialOneOf2
	}

	if obj.AddCredentialsRequestCredentialOneOf3 != nil {
		return obj.AddCredentialsRequestCredentialOneOf3
	}

	if obj.AddCredentialsRequestCredentialOneOf4 != nil {
		return obj.AddCredentialsRequestCredentialOneOf4
	}

	if obj.AddCredentialsRequestCredentialOneOf5 != nil {
		return obj.AddCredentialsRequestCredentialOneOf5
	}

	if obj.AddCredentialsRequestCredentialOneOf6 != nil {
		return obj.AddCredentialsRequestCredentialOneOf6
	}

	if obj.AddCredentialsRequestCredentialOneOf7 != nil {
		return obj.AddCredentialsRequestCredentialOneOf7
	}

	if obj.AddCredentialsRequestCredentialOneOf8 != nil {
		return obj.AddCredentialsRequestCredentialOneOf8
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj UpdateCredentialsRequestCredential) GetActualInstanceValue() interface{} {
	if obj.AddCredentialsRequestCredentialOneOf != nil {
		return *obj.AddCredentialsRequestCredentialOneOf
	}

	if obj.AddCredentialsRequestCredentialOneOf1 != nil {
		return *obj.AddCredentialsRequestCredentialOneOf1
	}

	if obj.AddCredentialsRequestCredentialOneOf2 != nil {
		return *obj.AddCredentialsRequestCredentialOneOf2
	}

	if obj.AddCredentialsRequestCredentialOneOf3 != nil {
		return *obj.AddCredentialsRequestCredentialOneOf3
	}

	if obj.AddCredentialsRequestCredentialOneOf4 != nil {
		return *obj.AddCredentialsRequestCredentialOneOf4
	}

	if obj.AddCredentialsRequestCredentialOneOf5 != nil {
		return *obj.AddCredentialsRequestCredentialOneOf5
	}

	if obj.AddCredentialsRequestCredentialOneOf6 != nil {
		return *obj.AddCredentialsRequestCredentialOneOf6
	}

	if obj.AddCredentialsRequestCredentialOneOf7 != nil {
		return *obj.AddCredentialsRequestCredentialOneOf7
	}

	if obj.AddCredentialsRequestCredentialOneOf8 != nil {
		return *obj.AddCredentialsRequestCredentialOneOf8
	}

	// all schemas are nil
	return nil
}

type NullableUpdateCredentialsRequestCredential struct {
	value *UpdateCredentialsRequestCredential
	isSet bool
}

func (v NullableUpdateCredentialsRequestCredential) Get() *UpdateCredentialsRequestCredential {
	return v.value
}

func (v *NullableUpdateCredentialsRequestCredential) Set(val *UpdateCredentialsRequestCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCredentialsRequestCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCredentialsRequestCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCredentialsRequestCredential(val *UpdateCredentialsRequestCredential) *NullableUpdateCredentialsRequestCredential {
	return &NullableUpdateCredentialsRequestCredential{value: val, isSet: true}
}

func (v NullableUpdateCredentialsRequestCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCredentialsRequestCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

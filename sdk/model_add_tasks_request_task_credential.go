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

// AddTasksRequestTaskCredential - Map containing Credential ID or the default `{\"type\": \"local\"}` which means use the values set in the local task options username and password instead of associating a credential.
type AddTasksRequestTaskCredential struct {
	AddIntegrationsRequestOneOfIntegrationCredentialOneOf *AddIntegrationsRequestOneOfIntegrationCredentialOneOf
	GetAlerts200ResponseAllOfChecksInnerAccount           *GetAlerts200ResponseAllOfChecksInnerAccount
}

// AddIntegrationsRequestOneOfIntegrationCredentialOneOfAsAddTasksRequestTaskCredential is a convenience function that returns AddIntegrationsRequestOneOfIntegrationCredentialOneOf wrapped in AddTasksRequestTaskCredential
func AddIntegrationsRequestOneOfIntegrationCredentialOneOfAsAddTasksRequestTaskCredential(v *AddIntegrationsRequestOneOfIntegrationCredentialOneOf) AddTasksRequestTaskCredential {
	return AddTasksRequestTaskCredential{
		AddIntegrationsRequestOneOfIntegrationCredentialOneOf: v,
	}
}

// GetAlerts200ResponseAllOfChecksInnerAccountAsAddTasksRequestTaskCredential is a convenience function that returns GetAlerts200ResponseAllOfChecksInnerAccount wrapped in AddTasksRequestTaskCredential
func GetAlerts200ResponseAllOfChecksInnerAccountAsAddTasksRequestTaskCredential(v *GetAlerts200ResponseAllOfChecksInnerAccount) AddTasksRequestTaskCredential {
	return AddTasksRequestTaskCredential{
		GetAlerts200ResponseAllOfChecksInnerAccount: v,
	}
}

func (dst *AddTasksRequestTaskCredential) UnmarshalMapstructure(data any) (any, error) {
	if dst == nil {
		dst = &AddTasksRequestTaskCredential{}
	}

	if out, ok := data.(AddIntegrationsRequestOneOfIntegrationCredentialOneOf); ok {
		dst.AddIntegrationsRequestOneOfIntegrationCredentialOneOf = &out
	}

	if out, ok := data.(GetAlerts200ResponseAllOfChecksInnerAccount); ok {
		dst.GetAlerts200ResponseAllOfChecksInnerAccount = &out
	}

	return dst, nil
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AddTasksRequestTaskCredential) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AddIntegrationsRequestOneOfIntegrationCredentialOneOf
	err = newStrictDecoder(data).Decode(&dst.AddIntegrationsRequestOneOfIntegrationCredentialOneOf)
	if err == nil {
		jsonAddIntegrationsRequestOneOfIntegrationCredentialOneOf, _ := json.Marshal(dst.AddIntegrationsRequestOneOfIntegrationCredentialOneOf)
		if string(jsonAddIntegrationsRequestOneOfIntegrationCredentialOneOf) == "{}" { // empty struct
			dst.AddIntegrationsRequestOneOfIntegrationCredentialOneOf = nil
		} else {
			if err = validator.Validate(dst.AddIntegrationsRequestOneOfIntegrationCredentialOneOf); err != nil {
				dst.AddIntegrationsRequestOneOfIntegrationCredentialOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.AddIntegrationsRequestOneOfIntegrationCredentialOneOf = nil
	}

	// try to unmarshal data into GetAlerts200ResponseAllOfChecksInnerAccount
	err = newStrictDecoder(data).Decode(&dst.GetAlerts200ResponseAllOfChecksInnerAccount)
	if err == nil {
		jsonGetAlerts200ResponseAllOfChecksInnerAccount, _ := json.Marshal(dst.GetAlerts200ResponseAllOfChecksInnerAccount)
		if string(jsonGetAlerts200ResponseAllOfChecksInnerAccount) == "{}" { // empty struct
			dst.GetAlerts200ResponseAllOfChecksInnerAccount = nil
		} else {
			if err = validator.Validate(dst.GetAlerts200ResponseAllOfChecksInnerAccount); err != nil {
				dst.GetAlerts200ResponseAllOfChecksInnerAccount = nil
			} else {
				match++
			}
		}
	} else {
		dst.GetAlerts200ResponseAllOfChecksInnerAccount = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AddIntegrationsRequestOneOfIntegrationCredentialOneOf = nil
		dst.GetAlerts200ResponseAllOfChecksInnerAccount = nil

		return NewResponseValidationError("data matches more than one schema in oneOf(AddTasksRequestTaskCredential)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return NewResponseValidationError("data failed to match schemas in oneOf(AddTasksRequestTaskCredential)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AddTasksRequestTaskCredential) MarshalJSON() ([]byte, error) {
	if src.AddIntegrationsRequestOneOfIntegrationCredentialOneOf != nil {
		return json.Marshal(&src.AddIntegrationsRequestOneOfIntegrationCredentialOneOf)
	}

	if src.GetAlerts200ResponseAllOfChecksInnerAccount != nil {
		return json.Marshal(&src.GetAlerts200ResponseAllOfChecksInnerAccount)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AddTasksRequestTaskCredential) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AddIntegrationsRequestOneOfIntegrationCredentialOneOf != nil {
		return obj.AddIntegrationsRequestOneOfIntegrationCredentialOneOf
	}

	if obj.GetAlerts200ResponseAllOfChecksInnerAccount != nil {
		return obj.GetAlerts200ResponseAllOfChecksInnerAccount
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AddTasksRequestTaskCredential) GetActualInstanceValue() interface{} {
	if obj.AddIntegrationsRequestOneOfIntegrationCredentialOneOf != nil {
		return *obj.AddIntegrationsRequestOneOfIntegrationCredentialOneOf
	}

	if obj.GetAlerts200ResponseAllOfChecksInnerAccount != nil {
		return *obj.GetAlerts200ResponseAllOfChecksInnerAccount
	}

	// all schemas are nil
	return nil
}

type NullableAddTasksRequestTaskCredential struct {
	value *AddTasksRequestTaskCredential
	isSet bool
}

func (v NullableAddTasksRequestTaskCredential) Get() *AddTasksRequestTaskCredential {
	return v.value
}

func (v *NullableAddTasksRequestTaskCredential) Set(val *AddTasksRequestTaskCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTasksRequestTaskCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTasksRequestTaskCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTasksRequestTaskCredential(val *AddTasksRequestTaskCredential) *NullableAddTasksRequestTaskCredential {
	return &NullableAddTasksRequestTaskCredential{value: val, isSet: true}
}

func (v NullableAddTasksRequestTaskCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTasksRequestTaskCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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
)

// checks if the GetApprovals200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetApprovals200Response{}

// GetApprovals200Response struct for GetApprovals200Response
type GetApprovals200Response struct {
	Approval             *GetApprovals200ResponseApproval `json:"approval,omitempty"`
	AdditionalProperties map[string]interface{}           `json:",remain"`
}

type _GetApprovals200Response GetApprovals200Response

// NewGetApprovals200Response instantiates a new GetApprovals200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApprovals200Response() *GetApprovals200Response {
	this := GetApprovals200Response{}
	return &this
}

// NewGetApprovals200ResponseWithDefaults instantiates a new GetApprovals200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApprovals200ResponseWithDefaults() *GetApprovals200Response {
	this := GetApprovals200Response{}
	return &this
}

// GetApproval returns the Approval field value if set, zero value otherwise.
func (o *GetApprovals200Response) GetApproval() GetApprovals200ResponseApproval {
	if o == nil || IsNil(o.Approval) {
		var ret GetApprovals200ResponseApproval
		return ret
	}
	return *o.Approval
}

// GetApprovalOk returns a tuple with the Approval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApprovals200Response) GetApprovalOk() (*GetApprovals200ResponseApproval, bool) {
	if o == nil || IsNil(o.Approval) {
		return nil, false
	}
	return o.Approval, true
}

// IsSetApproval returns a boolean if a field has been set.
func (o *GetApprovals200Response) IsSetApproval() bool {
	if o != nil && !IsNil(o.Approval) {
		return true
	}

	return false
}

// SetApproval gets a reference to the given GetApprovals200ResponseApproval and assigns it to the Approval field.
func (o *GetApprovals200Response) SetApproval(v GetApprovals200ResponseApproval) {
	o.Approval = &v
}

func (o GetApprovals200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApprovals200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Approval) {
		toSerialize["approval"] = o.Approval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetApprovals200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

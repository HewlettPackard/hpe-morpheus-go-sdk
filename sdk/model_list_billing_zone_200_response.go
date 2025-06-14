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

// checks if the ListBillingZone200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBillingZone200Response{}

// ListBillingZone200Response struct for ListBillingZone200Response
type ListBillingZone200Response struct {
	BillingInfo          *ListBillingZone200ResponseAllOfBillingInfo `json:"billingInfo,omitempty"`
	Success              *bool                                       `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}                      `json:",remain"`
}

type _ListBillingZone200Response ListBillingZone200Response

// NewListBillingZone200Response instantiates a new ListBillingZone200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBillingZone200Response() *ListBillingZone200Response {
	this := ListBillingZone200Response{}
	return &this
}

// NewListBillingZone200ResponseWithDefaults instantiates a new ListBillingZone200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBillingZone200ResponseWithDefaults() *ListBillingZone200Response {
	this := ListBillingZone200Response{}
	return &this
}

// GetBillingInfo returns the BillingInfo field value if set, zero value otherwise.
func (o *ListBillingZone200Response) GetBillingInfo() ListBillingZone200ResponseAllOfBillingInfo {
	if o == nil || IsNil(o.BillingInfo) {
		var ret ListBillingZone200ResponseAllOfBillingInfo
		return ret
	}
	return *o.BillingInfo
}

// GetBillingInfoOk returns a tuple with the BillingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBillingZone200Response) GetBillingInfoOk() (*ListBillingZone200ResponseAllOfBillingInfo, bool) {
	if o == nil || IsNil(o.BillingInfo) {
		return nil, false
	}
	return o.BillingInfo, true
}

// IsSetBillingInfo returns a boolean if a field has been set.
func (o *ListBillingZone200Response) IsSetBillingInfo() bool {
	if o != nil && !IsNil(o.BillingInfo) {
		return true
	}

	return false
}

// SetBillingInfo gets a reference to the given ListBillingZone200ResponseAllOfBillingInfo and assigns it to the BillingInfo field.
func (o *ListBillingZone200Response) SetBillingInfo(v ListBillingZone200ResponseAllOfBillingInfo) {
	o.BillingInfo = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ListBillingZone200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBillingZone200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *ListBillingZone200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ListBillingZone200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o ListBillingZone200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListBillingZone200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingInfo) {
		toSerialize["billingInfo"] = o.BillingInfo
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListBillingZone200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

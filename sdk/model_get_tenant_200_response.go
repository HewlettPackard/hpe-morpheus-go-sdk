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

// checks if the GetTenant200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTenant200Response{}

// GetTenant200Response struct for GetTenant200Response
type GetTenant200Response struct {
	Account              *ListTenants200ResponseAllOfAccountsInner `json:"account,omitempty"`
	AdditionalProperties map[string]interface{}                    `json:",remain"`
}

type _GetTenant200Response GetTenant200Response

// NewGetTenant200Response instantiates a new GetTenant200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTenant200Response() *GetTenant200Response {
	this := GetTenant200Response{}
	return &this
}

// NewGetTenant200ResponseWithDefaults instantiates a new GetTenant200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTenant200ResponseWithDefaults() *GetTenant200Response {
	this := GetTenant200Response{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *GetTenant200Response) GetAccount() ListTenants200ResponseAllOfAccountsInner {
	if o == nil || IsNil(o.Account) {
		var ret ListTenants200ResponseAllOfAccountsInner
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTenant200Response) GetAccountOk() (*ListTenants200ResponseAllOfAccountsInner, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// IsSetAccount returns a boolean if a field has been set.
func (o *GetTenant200Response) IsSetAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ListTenants200ResponseAllOfAccountsInner and assigns it to the Account field.
func (o *GetTenant200Response) SetAccount(v ListTenants200ResponseAllOfAccountsInner) {
	o.Account = &v
}

func (o GetTenant200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTenant200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetTenant200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

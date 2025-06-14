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

// checks if the ListTenants200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTenants200Response{}

// ListTenants200Response struct for ListTenants200Response
type ListTenants200Response struct {
	Accounts             []ListTenants200ResponseAllOfAccountsInner `json:"accounts,omitempty"`
	Meta                 *ListActivity200ResponseAllOfMeta          `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}                     `json:",remain"`
}

type _ListTenants200Response ListTenants200Response

// NewListTenants200Response instantiates a new ListTenants200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTenants200Response() *ListTenants200Response {
	this := ListTenants200Response{}
	return &this
}

// NewListTenants200ResponseWithDefaults instantiates a new ListTenants200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTenants200ResponseWithDefaults() *ListTenants200Response {
	this := ListTenants200Response{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *ListTenants200Response) GetAccounts() []ListTenants200ResponseAllOfAccountsInner {
	if o == nil || IsNil(o.Accounts) {
		var ret []ListTenants200ResponseAllOfAccountsInner
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTenants200Response) GetAccountsOk() ([]ListTenants200ResponseAllOfAccountsInner, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// IsSetAccounts returns a boolean if a field has been set.
func (o *ListTenants200Response) IsSetAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []ListTenants200ResponseAllOfAccountsInner and assigns it to the Accounts field.
func (o *ListTenants200Response) SetAccounts(v []ListTenants200ResponseAllOfAccountsInner) {
	o.Accounts = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListTenants200Response) GetMeta() ListActivity200ResponseAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListActivity200ResponseAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTenants200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// IsSetMeta returns a boolean if a field has been set.
func (o *ListTenants200Response) IsSetMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListActivity200ResponseAllOfMeta and assigns it to the Meta field.
func (o *ListTenants200Response) SetMeta(v ListActivity200ResponseAllOfMeta) {
	o.Meta = &v
}

func (o ListTenants200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTenants200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListTenants200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

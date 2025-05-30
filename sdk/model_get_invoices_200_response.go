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

// checks if the GetInvoices200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInvoices200Response{}

// GetInvoices200Response struct for GetInvoices200Response
type GetInvoices200Response struct {
	Invoice *ListInvoices200ResponseAllOfInvoicesInner `json:"invoice,omitempty"`
	MasterAccount *bool `json:"masterAccount,omitempty"`
}

// NewGetInvoices200Response instantiates a new GetInvoices200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInvoices200Response() *GetInvoices200Response {
	this := GetInvoices200Response{}
	return &this
}

// NewGetInvoices200ResponseWithDefaults instantiates a new GetInvoices200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInvoices200ResponseWithDefaults() *GetInvoices200Response {
	this := GetInvoices200Response{}
	return &this
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *GetInvoices200Response) GetInvoice() ListInvoices200ResponseAllOfInvoicesInner {
	if o == nil || IsNil(o.Invoice) {
		var ret ListInvoices200ResponseAllOfInvoicesInner
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoices200Response) GetInvoiceOk() (*ListInvoices200ResponseAllOfInvoicesInner, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// IsSetInvoice returns a boolean if a field has been set.
func (o *GetInvoices200Response) IsSetInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given ListInvoices200ResponseAllOfInvoicesInner and assigns it to the Invoice field.
func (o *GetInvoices200Response) SetInvoice(v ListInvoices200ResponseAllOfInvoicesInner) {
	o.Invoice = &v
}

// GetMasterAccount returns the MasterAccount field value if set, zero value otherwise.
func (o *GetInvoices200Response) GetMasterAccount() bool {
	if o == nil || IsNil(o.MasterAccount) {
		var ret bool
		return ret
	}
	return *o.MasterAccount
}

// GetMasterAccountOk returns a tuple with the MasterAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoices200Response) GetMasterAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.MasterAccount) {
		return nil, false
	}
	return o.MasterAccount, true
}

// IsSetMasterAccount returns a boolean if a field has been set.
func (o *GetInvoices200Response) IsSetMasterAccount() bool {
	if o != nil && !IsNil(o.MasterAccount) {
		return true
	}

	return false
}

// SetMasterAccount gets a reference to the given bool and assigns it to the MasterAccount field.
func (o *GetInvoices200Response) SetMasterAccount(v bool) {
	o.MasterAccount = &v
}

func (o GetInvoices200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInvoices200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.MasterAccount) {
		toSerialize["masterAccount"] = o.MasterAccount
	}
	return toSerialize, nil
}

type NullableGetInvoices200Response struct {
	value *GetInvoices200Response
	isSet bool
}

func (v NullableGetInvoices200Response) Get() *GetInvoices200Response {
	return v.value
}

func (v *NullableGetInvoices200Response) Set(val *GetInvoices200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvoices200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvoices200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvoices200Response(val *GetInvoices200Response) *NullableGetInvoices200Response {
	return &NullableGetInvoices200Response{value: val, isSet: true}
}

func (v NullableGetInvoices200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInvoices200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



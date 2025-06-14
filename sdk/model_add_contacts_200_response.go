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

// checks if the AddContacts200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddContacts200Response{}

// AddContacts200Response struct for AddContacts200Response
type AddContacts200Response struct {
	Contact              *ListContacts200ResponseAllOfContactsInner `json:"contact,omitempty"`
	Success              *bool                                      `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}                     `json:",remain"`
}

type _AddContacts200Response AddContacts200Response

// NewAddContacts200Response instantiates a new AddContacts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddContacts200Response() *AddContacts200Response {
	this := AddContacts200Response{}
	return &this
}

// NewAddContacts200ResponseWithDefaults instantiates a new AddContacts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddContacts200ResponseWithDefaults() *AddContacts200Response {
	this := AddContacts200Response{}
	return &this
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *AddContacts200Response) GetContact() ListContacts200ResponseAllOfContactsInner {
	if o == nil || IsNil(o.Contact) {
		var ret ListContacts200ResponseAllOfContactsInner
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContacts200Response) GetContactOk() (*ListContacts200ResponseAllOfContactsInner, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// IsSetContact returns a boolean if a field has been set.
func (o *AddContacts200Response) IsSetContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given ListContacts200ResponseAllOfContactsInner and assigns it to the Contact field.
func (o *AddContacts200Response) SetContact(v ListContacts200ResponseAllOfContactsInner) {
	o.Contact = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AddContacts200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContacts200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *AddContacts200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AddContacts200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o AddContacts200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddContacts200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddContacts200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the AddContactsRequestContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddContactsRequestContact{}

// AddContactsRequestContact Payload for creating a new monitoring contact
type AddContactsRequestContact struct {
	// Unique name scoped to your account for the contact
	Name string `json:"name"`
	// Email notification address
	EmailAddress *string `json:"emailAddress,omitempty"`
	// SMS notification address
	SmsAddress *string `json:"smsAddress,omitempty"`
	// Slack Hook
	SlackHook            *string                `json:"slackHook,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddContactsRequestContact AddContactsRequestContact

// NewAddContactsRequestContact instantiates a new AddContactsRequestContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddContactsRequestContact(name string) *AddContactsRequestContact {
	this := AddContactsRequestContact{}
	this.Name = name
	return &this
}

// NewAddContactsRequestContactWithDefaults instantiates a new AddContactsRequestContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddContactsRequestContactWithDefaults() *AddContactsRequestContact {
	this := AddContactsRequestContact{}
	return &this
}

// GetName returns the Name field value
func (o *AddContactsRequestContact) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddContactsRequestContact) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddContactsRequestContact) SetName(v string) {
	o.Name = v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *AddContactsRequestContact) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContactsRequestContact) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// IsSetEmailAddress returns a boolean if a field has been set.
func (o *AddContactsRequestContact) IsSetEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *AddContactsRequestContact) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetSmsAddress returns the SmsAddress field value if set, zero value otherwise.
func (o *AddContactsRequestContact) GetSmsAddress() string {
	if o == nil || IsNil(o.SmsAddress) {
		var ret string
		return ret
	}
	return *o.SmsAddress
}

// GetSmsAddressOk returns a tuple with the SmsAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContactsRequestContact) GetSmsAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SmsAddress) {
		return nil, false
	}
	return o.SmsAddress, true
}

// IsSetSmsAddress returns a boolean if a field has been set.
func (o *AddContactsRequestContact) IsSetSmsAddress() bool {
	if o != nil && !IsNil(o.SmsAddress) {
		return true
	}

	return false
}

// SetSmsAddress gets a reference to the given string and assigns it to the SmsAddress field.
func (o *AddContactsRequestContact) SetSmsAddress(v string) {
	o.SmsAddress = &v
}

// GetSlackHook returns the SlackHook field value if set, zero value otherwise.
func (o *AddContactsRequestContact) GetSlackHook() string {
	if o == nil || IsNil(o.SlackHook) {
		var ret string
		return ret
	}
	return *o.SlackHook
}

// GetSlackHookOk returns a tuple with the SlackHook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContactsRequestContact) GetSlackHookOk() (*string, bool) {
	if o == nil || IsNil(o.SlackHook) {
		return nil, false
	}
	return o.SlackHook, true
}

// IsSetSlackHook returns a boolean if a field has been set.
func (o *AddContactsRequestContact) IsSetSlackHook() bool {
	if o != nil && !IsNil(o.SlackHook) {
		return true
	}

	return false
}

// SetSlackHook gets a reference to the given string and assigns it to the SlackHook field.
func (o *AddContactsRequestContact) SetSlackHook(v string) {
	o.SlackHook = &v
}

func (o AddContactsRequestContact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddContactsRequestContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if !IsNil(o.SmsAddress) {
		toSerialize["smsAddress"] = o.SmsAddress
	}
	if !IsNil(o.SlackHook) {
		toSerialize["slackHook"] = o.SlackHook
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddContactsRequestContact) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

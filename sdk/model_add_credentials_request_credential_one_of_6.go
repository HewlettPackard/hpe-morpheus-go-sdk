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

// checks if the AddCredentialsRequestCredentialOneOf6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCredentialsRequestCredentialOneOf6{}

// AddCredentialsRequestCredentialOneOf6 struct for AddCredentialsRequestCredentialOneOf6
type AddCredentialsRequestCredentialOneOf6 struct {
	// Credential Type Code
	Type string `json:"type"`
	// A unique name scoped to your account for the credential
	Name string `json:"name"`
	// Optional Description
	Description *string `json:"description,omitempty"`
	// Credential enabled
	Enabled     *bool                                            `json:"enabled,omitempty"`
	Integration *AddCredentialsRequestCredentialOneOfIntegration `json:"integration,omitempty"`
	// Username
	Username string `json:"username"`
	// Password
	Password             string                 `json:"password"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddCredentialsRequestCredentialOneOf6 AddCredentialsRequestCredentialOneOf6

// NewAddCredentialsRequestCredentialOneOf6 instantiates a new AddCredentialsRequestCredentialOneOf6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCredentialsRequestCredentialOneOf6(type_ string, name string, username string, password string) *AddCredentialsRequestCredentialOneOf6 {
	this := AddCredentialsRequestCredentialOneOf6{}
	this.Type = type_
	this.Name = name
	var enabled bool = true
	this.Enabled = &enabled
	this.Username = username
	this.Password = password
	return &this
}

// NewAddCredentialsRequestCredentialOneOf6WithDefaults instantiates a new AddCredentialsRequestCredentialOneOf6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCredentialsRequestCredentialOneOf6WithDefaults() *AddCredentialsRequestCredentialOneOf6 {
	this := AddCredentialsRequestCredentialOneOf6{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetType returns the Type field value
func (o *AddCredentialsRequestCredentialOneOf6) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddCredentialsRequestCredentialOneOf6) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddCredentialsRequestCredentialOneOf6) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *AddCredentialsRequestCredentialOneOf6) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddCredentialsRequestCredentialOneOf6) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddCredentialsRequestCredentialOneOf6) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddCredentialsRequestCredentialOneOf6) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCredentialsRequestCredentialOneOf6) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *AddCredentialsRequestCredentialOneOf6) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddCredentialsRequestCredentialOneOf6) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddCredentialsRequestCredentialOneOf6) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCredentialsRequestCredentialOneOf6) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *AddCredentialsRequestCredentialOneOf6) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddCredentialsRequestCredentialOneOf6) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *AddCredentialsRequestCredentialOneOf6) GetIntegration() AddCredentialsRequestCredentialOneOfIntegration {
	if o == nil || IsNil(o.Integration) {
		var ret AddCredentialsRequestCredentialOneOfIntegration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCredentialsRequestCredentialOneOf6) GetIntegrationOk() (*AddCredentialsRequestCredentialOneOfIntegration, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// IsSetIntegration returns a boolean if a field has been set.
func (o *AddCredentialsRequestCredentialOneOf6) IsSetIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given AddCredentialsRequestCredentialOneOfIntegration and assigns it to the Integration field.
func (o *AddCredentialsRequestCredentialOneOf6) SetIntegration(v AddCredentialsRequestCredentialOneOfIntegration) {
	o.Integration = &v
}

// GetUsername returns the Username field value
func (o *AddCredentialsRequestCredentialOneOf6) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AddCredentialsRequestCredentialOneOf6) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AddCredentialsRequestCredentialOneOf6) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *AddCredentialsRequestCredentialOneOf6) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AddCredentialsRequestCredentialOneOf6) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AddCredentialsRequestCredentialOneOf6) SetPassword(v string) {
	o.Password = v
}

func (o AddCredentialsRequestCredentialOneOf6) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCredentialsRequestCredentialOneOf6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddCredentialsRequestCredentialOneOf6) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the Whoami200ResponsePermissionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Whoami200ResponsePermissionsInner{}

// Whoami200ResponsePermissionsInner struct for Whoami200ResponsePermissionsInner
type Whoami200ResponsePermissionsInner struct {
	Name                 *string                `json:"name,omitempty"`
	Code                 *string                `json:"code,omitempty"`
	Access               *string                `json:"access,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _Whoami200ResponsePermissionsInner Whoami200ResponsePermissionsInner

// NewWhoami200ResponsePermissionsInner instantiates a new Whoami200ResponsePermissionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhoami200ResponsePermissionsInner() *Whoami200ResponsePermissionsInner {
	this := Whoami200ResponsePermissionsInner{}
	return &this
}

// NewWhoami200ResponsePermissionsInnerWithDefaults instantiates a new Whoami200ResponsePermissionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhoami200ResponsePermissionsInnerWithDefaults() *Whoami200ResponsePermissionsInner {
	this := Whoami200ResponsePermissionsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Whoami200ResponsePermissionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Whoami200ResponsePermissionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *Whoami200ResponsePermissionsInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Whoami200ResponsePermissionsInner) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Whoami200ResponsePermissionsInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Whoami200ResponsePermissionsInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *Whoami200ResponsePermissionsInner) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Whoami200ResponsePermissionsInner) SetCode(v string) {
	o.Code = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *Whoami200ResponsePermissionsInner) GetAccess() string {
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Whoami200ResponsePermissionsInner) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// IsSetAccess returns a boolean if a field has been set.
func (o *Whoami200ResponsePermissionsInner) IsSetAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *Whoami200ResponsePermissionsInner) SetAccess(v string) {
	o.Access = &v
}

func (o Whoami200ResponsePermissionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Whoami200ResponsePermissionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *Whoami200ResponsePermissionsInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

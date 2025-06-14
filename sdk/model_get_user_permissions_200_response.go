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

// checks if the GetUserPermissions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserPermissions200Response{}

// GetUserPermissions200Response struct for GetUserPermissions200Response
type GetUserPermissions200Response struct {
	Access               *AddUserTenant200ResponseAllOfUserAccess `json:"access,omitempty"`
	AdditionalProperties map[string]interface{}                   `json:",remain"`
}

type _GetUserPermissions200Response GetUserPermissions200Response

// NewGetUserPermissions200Response instantiates a new GetUserPermissions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserPermissions200Response() *GetUserPermissions200Response {
	this := GetUserPermissions200Response{}
	return &this
}

// NewGetUserPermissions200ResponseWithDefaults instantiates a new GetUserPermissions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserPermissions200ResponseWithDefaults() *GetUserPermissions200Response {
	this := GetUserPermissions200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GetUserPermissions200Response) GetAccess() AddUserTenant200ResponseAllOfUserAccess {
	if o == nil || IsNil(o.Access) {
		var ret AddUserTenant200ResponseAllOfUserAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserPermissions200Response) GetAccessOk() (*AddUserTenant200ResponseAllOfUserAccess, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// IsSetAccess returns a boolean if a field has been set.
func (o *GetUserPermissions200Response) IsSetAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AddUserTenant200ResponseAllOfUserAccess and assigns it to the Access field.
func (o *GetUserPermissions200Response) SetAccess(v AddUserTenant200ResponseAllOfUserAccess) {
	o.Access = &v
}

func (o GetUserPermissions200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserPermissions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetUserPermissions200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the ManageHostPlacementRequestServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManageHostPlacementRequestServer{}

// ManageHostPlacementRequestServer struct for ManageHostPlacementRequestServer
type ManageHostPlacementRequestServer struct {
	// Placement Strategy
	Name *string `json:"name,omitempty"`
	PreferredServer *ManageHostPlacementRequestServerPreferredServer `json:"preferredServer,omitempty"`
}

// NewManageHostPlacementRequestServer instantiates a new ManageHostPlacementRequestServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageHostPlacementRequestServer() *ManageHostPlacementRequestServer {
	this := ManageHostPlacementRequestServer{}
	return &this
}

// NewManageHostPlacementRequestServerWithDefaults instantiates a new ManageHostPlacementRequestServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageHostPlacementRequestServerWithDefaults() *ManageHostPlacementRequestServer {
	this := ManageHostPlacementRequestServer{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ManageHostPlacementRequestServer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageHostPlacementRequestServer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ManageHostPlacementRequestServer) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ManageHostPlacementRequestServer) SetName(v string) {
	o.Name = &v
}

// GetPreferredServer returns the PreferredServer field value if set, zero value otherwise.
func (o *ManageHostPlacementRequestServer) GetPreferredServer() ManageHostPlacementRequestServerPreferredServer {
	if o == nil || IsNil(o.PreferredServer) {
		var ret ManageHostPlacementRequestServerPreferredServer
		return ret
	}
	return *o.PreferredServer
}

// GetPreferredServerOk returns a tuple with the PreferredServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageHostPlacementRequestServer) GetPreferredServerOk() (*ManageHostPlacementRequestServerPreferredServer, bool) {
	if o == nil || IsNil(o.PreferredServer) {
		return nil, false
	}
	return o.PreferredServer, true
}

// IsSetPreferredServer returns a boolean if a field has been set.
func (o *ManageHostPlacementRequestServer) IsSetPreferredServer() bool {
	if o != nil && !IsNil(o.PreferredServer) {
		return true
	}

	return false
}

// SetPreferredServer gets a reference to the given ManageHostPlacementRequestServerPreferredServer and assigns it to the PreferredServer field.
func (o *ManageHostPlacementRequestServer) SetPreferredServer(v ManageHostPlacementRequestServerPreferredServer) {
	o.PreferredServer = &v
}

func (o ManageHostPlacementRequestServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManageHostPlacementRequestServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PreferredServer) {
		toSerialize["preferredServer"] = o.PreferredServer
	}
	return toSerialize, nil
}

type NullableManageHostPlacementRequestServer struct {
	value *ManageHostPlacementRequestServer
	isSet bool
}

func (v NullableManageHostPlacementRequestServer) Get() *ManageHostPlacementRequestServer {
	return v.value
}

func (v *NullableManageHostPlacementRequestServer) Set(val *ManageHostPlacementRequestServer) {
	v.value = val
	v.isSet = true
}

func (v NullableManageHostPlacementRequestServer) IsSet() bool {
	return v.isSet
}

func (v *NullableManageHostPlacementRequestServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageHostPlacementRequestServer(val *ManageHostPlacementRequestServer) *NullableManageHostPlacementRequestServer {
	return &NullableManageHostPlacementRequestServer{value: val, isSet: true}
}

func (v NullableManageHostPlacementRequestServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageHostPlacementRequestServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



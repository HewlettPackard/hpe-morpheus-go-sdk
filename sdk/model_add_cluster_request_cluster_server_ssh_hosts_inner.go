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

// checks if the AddClusterRequestClusterServerSshHostsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddClusterRequestClusterServerSshHostsInner{}

// AddClusterRequestClusterServerSshHostsInner struct for AddClusterRequestClusterServerSshHostsInner
type AddClusterRequestClusterServerSshHostsInner struct {
	// Host IP
	Ip *string `json:"ip,omitempty"`
	// Host Name
	Name                 *string                `json:"name,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddClusterRequestClusterServerSshHostsInner AddClusterRequestClusterServerSshHostsInner

// NewAddClusterRequestClusterServerSshHostsInner instantiates a new AddClusterRequestClusterServerSshHostsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddClusterRequestClusterServerSshHostsInner() *AddClusterRequestClusterServerSshHostsInner {
	this := AddClusterRequestClusterServerSshHostsInner{}
	return &this
}

// NewAddClusterRequestClusterServerSshHostsInnerWithDefaults instantiates a new AddClusterRequestClusterServerSshHostsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddClusterRequestClusterServerSshHostsInnerWithDefaults() *AddClusterRequestClusterServerSshHostsInner {
	this := AddClusterRequestClusterServerSshHostsInner{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *AddClusterRequestClusterServerSshHostsInner) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddClusterRequestClusterServerSshHostsInner) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// IsSetIp returns a boolean if a field has been set.
func (o *AddClusterRequestClusterServerSshHostsInner) IsSetIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *AddClusterRequestClusterServerSshHostsInner) SetIp(v string) {
	o.Ip = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddClusterRequestClusterServerSshHostsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddClusterRequestClusterServerSshHostsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *AddClusterRequestClusterServerSshHostsInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddClusterRequestClusterServerSshHostsInner) SetName(v string) {
	o.Name = &v
}

func (o AddClusterRequestClusterServerSshHostsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddClusterRequestClusterServerSshHostsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddClusterRequestClusterServerSshHostsInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

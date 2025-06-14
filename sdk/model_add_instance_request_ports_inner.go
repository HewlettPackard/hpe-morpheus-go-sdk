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

// checks if the AddInstanceRequestPortsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddInstanceRequestPortsInner{}

// AddInstanceRequestPortsInner struct for AddInstanceRequestPortsInner
type AddInstanceRequestPortsInner struct {
	// Port number.
	Port int64 `json:"port"`
	// A name for the port.
	Name *string `json:"name,omitempty"`
	// The load balancer protocol. HTTP, HTTPS, or TCP.
	Lb                   *string                `json:"lb,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddInstanceRequestPortsInner AddInstanceRequestPortsInner

// NewAddInstanceRequestPortsInner instantiates a new AddInstanceRequestPortsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddInstanceRequestPortsInner(port int64) *AddInstanceRequestPortsInner {
	this := AddInstanceRequestPortsInner{}
	this.Port = port
	return &this
}

// NewAddInstanceRequestPortsInnerWithDefaults instantiates a new AddInstanceRequestPortsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddInstanceRequestPortsInnerWithDefaults() *AddInstanceRequestPortsInner {
	this := AddInstanceRequestPortsInner{}
	return &this
}

// GetPort returns the Port field value
func (o *AddInstanceRequestPortsInner) GetPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *AddInstanceRequestPortsInner) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *AddInstanceRequestPortsInner) SetPort(v int64) {
	o.Port = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddInstanceRequestPortsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddInstanceRequestPortsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *AddInstanceRequestPortsInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddInstanceRequestPortsInner) SetName(v string) {
	o.Name = &v
}

// GetLb returns the Lb field value if set, zero value otherwise.
func (o *AddInstanceRequestPortsInner) GetLb() string {
	if o == nil || IsNil(o.Lb) {
		var ret string
		return ret
	}
	return *o.Lb
}

// GetLbOk returns a tuple with the Lb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddInstanceRequestPortsInner) GetLbOk() (*string, bool) {
	if o == nil || IsNil(o.Lb) {
		return nil, false
	}
	return o.Lb, true
}

// IsSetLb returns a boolean if a field has been set.
func (o *AddInstanceRequestPortsInner) IsSetLb() bool {
	if o != nil && !IsNil(o.Lb) {
		return true
	}

	return false
}

// SetLb gets a reference to the given string and assigns it to the Lb field.
func (o *AddInstanceRequestPortsInner) SetLb(v string) {
	o.Lb = &v
}

func (o AddInstanceRequestPortsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddInstanceRequestPortsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["port"] = o.Port
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Lb) {
		toSerialize["lb"] = o.Lb
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddInstanceRequestPortsInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

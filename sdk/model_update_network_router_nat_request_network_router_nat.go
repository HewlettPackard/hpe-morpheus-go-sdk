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

// checks if the UpdateNetworkRouterNatRequestNetworkRouterNAT type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkRouterNatRequestNetworkRouterNAT{}

// UpdateNetworkRouterNatRequestNetworkRouterNAT For a full list of available NAT options, see natOptionTypes in the specific Network Router Type
type UpdateNetworkRouterNatRequestNetworkRouterNAT struct {
	Name                 interface{}            `json:"name,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _UpdateNetworkRouterNatRequestNetworkRouterNAT UpdateNetworkRouterNatRequestNetworkRouterNAT

// NewUpdateNetworkRouterNatRequestNetworkRouterNAT instantiates a new UpdateNetworkRouterNatRequestNetworkRouterNAT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkRouterNatRequestNetworkRouterNAT() *UpdateNetworkRouterNatRequestNetworkRouterNAT {
	this := UpdateNetworkRouterNatRequestNetworkRouterNAT{}
	return &this
}

// NewUpdateNetworkRouterNatRequestNetworkRouterNATWithDefaults instantiates a new UpdateNetworkRouterNatRequestNetworkRouterNAT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkRouterNatRequestNetworkRouterNATWithDefaults() *UpdateNetworkRouterNatRequestNetworkRouterNAT {
	this := UpdateNetworkRouterNatRequestNetworkRouterNAT{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateNetworkRouterNatRequestNetworkRouterNAT) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateNetworkRouterNatRequestNetworkRouterNAT) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *UpdateNetworkRouterNatRequestNetworkRouterNAT) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *UpdateNetworkRouterNatRequestNetworkRouterNAT) SetName(v interface{}) {
	o.Name = v
}

func (o UpdateNetworkRouterNatRequestNetworkRouterNAT) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkRouterNatRequestNetworkRouterNAT) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateNetworkRouterNatRequestNetworkRouterNAT) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

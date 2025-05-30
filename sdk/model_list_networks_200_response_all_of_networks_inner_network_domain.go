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

// checks if the ListNetworks200ResponseAllOfNetworksInnerNetworkDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListNetworks200ResponseAllOfNetworksInnerNetworkDomain{}

// ListNetworks200ResponseAllOfNetworksInnerNetworkDomain struct for ListNetworks200ResponseAllOfNetworksInnerNetworkDomain
type ListNetworks200ResponseAllOfNetworksInnerNetworkDomain struct {
	// Network Domain ID
	Id *int64 `json:"id,omitempty"`
}

// NewListNetworks200ResponseAllOfNetworksInnerNetworkDomain instantiates a new ListNetworks200ResponseAllOfNetworksInnerNetworkDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListNetworks200ResponseAllOfNetworksInnerNetworkDomain() *ListNetworks200ResponseAllOfNetworksInnerNetworkDomain {
	this := ListNetworks200ResponseAllOfNetworksInnerNetworkDomain{}
	return &this
}

// NewListNetworks200ResponseAllOfNetworksInnerNetworkDomainWithDefaults instantiates a new ListNetworks200ResponseAllOfNetworksInnerNetworkDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListNetworks200ResponseAllOfNetworksInnerNetworkDomainWithDefaults() *ListNetworks200ResponseAllOfNetworksInnerNetworkDomain {
	this := ListNetworks200ResponseAllOfNetworksInnerNetworkDomain{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListNetworks200ResponseAllOfNetworksInnerNetworkDomain) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworks200ResponseAllOfNetworksInnerNetworkDomain) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListNetworks200ResponseAllOfNetworksInnerNetworkDomain) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListNetworks200ResponseAllOfNetworksInnerNetworkDomain) SetId(v int64) {
	o.Id = &v
}

func (o ListNetworks200ResponseAllOfNetworksInnerNetworkDomain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListNetworks200ResponseAllOfNetworksInnerNetworkDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableListNetworks200ResponseAllOfNetworksInnerNetworkDomain struct {
	value *ListNetworks200ResponseAllOfNetworksInnerNetworkDomain
	isSet bool
}

func (v NullableListNetworks200ResponseAllOfNetworksInnerNetworkDomain) Get() *ListNetworks200ResponseAllOfNetworksInnerNetworkDomain {
	return v.value
}

func (v *NullableListNetworks200ResponseAllOfNetworksInnerNetworkDomain) Set(val *ListNetworks200ResponseAllOfNetworksInnerNetworkDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableListNetworks200ResponseAllOfNetworksInnerNetworkDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableListNetworks200ResponseAllOfNetworksInnerNetworkDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListNetworks200ResponseAllOfNetworksInnerNetworkDomain(val *ListNetworks200ResponseAllOfNetworksInnerNetworkDomain) *NullableListNetworks200ResponseAllOfNetworksInnerNetworkDomain {
	return &NullableListNetworks200ResponseAllOfNetworksInnerNetworkDomain{value: val, isSet: true}
}

func (v NullableListNetworks200ResponseAllOfNetworksInnerNetworkDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListNetworks200ResponseAllOfNetworksInnerNetworkDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



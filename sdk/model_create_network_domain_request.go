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

// checks if the CreateNetworkDomainRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkDomainRequest{}

// CreateNetworkDomainRequest struct for CreateNetworkDomainRequest
type CreateNetworkDomainRequest struct {
	NetworkDomain *CreateNetworkDomainRequestNetworkDomain `json:"networkDomain,omitempty"`
}

// NewCreateNetworkDomainRequest instantiates a new CreateNetworkDomainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkDomainRequest() *CreateNetworkDomainRequest {
	this := CreateNetworkDomainRequest{}
	return &this
}

// NewCreateNetworkDomainRequestWithDefaults instantiates a new CreateNetworkDomainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkDomainRequestWithDefaults() *CreateNetworkDomainRequest {
	this := CreateNetworkDomainRequest{}
	return &this
}

// GetNetworkDomain returns the NetworkDomain field value if set, zero value otherwise.
func (o *CreateNetworkDomainRequest) GetNetworkDomain() CreateNetworkDomainRequestNetworkDomain {
	if o == nil || IsNil(o.NetworkDomain) {
		var ret CreateNetworkDomainRequestNetworkDomain
		return ret
	}
	return *o.NetworkDomain
}

// GetNetworkDomainOk returns a tuple with the NetworkDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkDomainRequest) GetNetworkDomainOk() (*CreateNetworkDomainRequestNetworkDomain, bool) {
	if o == nil || IsNil(o.NetworkDomain) {
		return nil, false
	}
	return o.NetworkDomain, true
}

// IsSetNetworkDomain returns a boolean if a field has been set.
func (o *CreateNetworkDomainRequest) IsSetNetworkDomain() bool {
	if o != nil && !IsNil(o.NetworkDomain) {
		return true
	}

	return false
}

// SetNetworkDomain gets a reference to the given CreateNetworkDomainRequestNetworkDomain and assigns it to the NetworkDomain field.
func (o *CreateNetworkDomainRequest) SetNetworkDomain(v CreateNetworkDomainRequestNetworkDomain) {
	o.NetworkDomain = &v
}

func (o CreateNetworkDomainRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkDomainRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkDomain) {
		toSerialize["networkDomain"] = o.NetworkDomain
	}
	return toSerialize, nil
}

type NullableCreateNetworkDomainRequest struct {
	value *CreateNetworkDomainRequest
	isSet bool
}

func (v NullableCreateNetworkDomainRequest) Get() *CreateNetworkDomainRequest {
	return v.value
}

func (v *NullableCreateNetworkDomainRequest) Set(val *CreateNetworkDomainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkDomainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkDomainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkDomainRequest(val *CreateNetworkDomainRequest) *NullableCreateNetworkDomainRequest {
	return &NullableCreateNetworkDomainRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkDomainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkDomainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the CreateNetworkTransportZoneRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkTransportZoneRequest{}

// CreateNetworkTransportZoneRequest The parameters for creating a network transport zone is type dependent. The following lists the common parameters. See get a specific type to list available options for the network server type.
type CreateNetworkTransportZoneRequest struct {
	NetworkScope *CreateNetworkTransportZoneRequestNetworkScope `json:"networkScope,omitempty"`
}

// NewCreateNetworkTransportZoneRequest instantiates a new CreateNetworkTransportZoneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkTransportZoneRequest() *CreateNetworkTransportZoneRequest {
	this := CreateNetworkTransportZoneRequest{}
	return &this
}

// NewCreateNetworkTransportZoneRequestWithDefaults instantiates a new CreateNetworkTransportZoneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkTransportZoneRequestWithDefaults() *CreateNetworkTransportZoneRequest {
	this := CreateNetworkTransportZoneRequest{}
	return &this
}

// GetNetworkScope returns the NetworkScope field value if set, zero value otherwise.
func (o *CreateNetworkTransportZoneRequest) GetNetworkScope() CreateNetworkTransportZoneRequestNetworkScope {
	if o == nil || IsNil(o.NetworkScope) {
		var ret CreateNetworkTransportZoneRequestNetworkScope
		return ret
	}
	return *o.NetworkScope
}

// GetNetworkScopeOk returns a tuple with the NetworkScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkTransportZoneRequest) GetNetworkScopeOk() (*CreateNetworkTransportZoneRequestNetworkScope, bool) {
	if o == nil || IsNil(o.NetworkScope) {
		return nil, false
	}
	return o.NetworkScope, true
}

// IsSetNetworkScope returns a boolean if a field has been set.
func (o *CreateNetworkTransportZoneRequest) IsSetNetworkScope() bool {
	if o != nil && !IsNil(o.NetworkScope) {
		return true
	}

	return false
}

// SetNetworkScope gets a reference to the given CreateNetworkTransportZoneRequestNetworkScope and assigns it to the NetworkScope field.
func (o *CreateNetworkTransportZoneRequest) SetNetworkScope(v CreateNetworkTransportZoneRequestNetworkScope) {
	o.NetworkScope = &v
}

func (o CreateNetworkTransportZoneRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkTransportZoneRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkScope) {
		toSerialize["networkScope"] = o.NetworkScope
	}
	return toSerialize, nil
}

type NullableCreateNetworkTransportZoneRequest struct {
	value *CreateNetworkTransportZoneRequest
	isSet bool
}

func (v NullableCreateNetworkTransportZoneRequest) Get() *CreateNetworkTransportZoneRequest {
	return v.value
}

func (v *NullableCreateNetworkTransportZoneRequest) Set(val *CreateNetworkTransportZoneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkTransportZoneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkTransportZoneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkTransportZoneRequest(val *CreateNetworkTransportZoneRequest) *NullableCreateNetworkTransportZoneRequest {
	return &NullableCreateNetworkTransportZoneRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkTransportZoneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkTransportZoneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



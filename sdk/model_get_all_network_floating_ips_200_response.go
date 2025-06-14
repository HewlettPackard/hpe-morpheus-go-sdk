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

// checks if the GetAllNetworkFloatingIps200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllNetworkFloatingIps200Response{}

// GetAllNetworkFloatingIps200Response struct for GetAllNetworkFloatingIps200Response
type GetAllNetworkFloatingIps200Response struct {
	NetworkFloatingIps   interface{}                       `json:"networkFloatingIps,omitempty"`
	Meta                 *ListActivity200ResponseAllOfMeta `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}            `json:",remain"`
}

type _GetAllNetworkFloatingIps200Response GetAllNetworkFloatingIps200Response

// NewGetAllNetworkFloatingIps200Response instantiates a new GetAllNetworkFloatingIps200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllNetworkFloatingIps200Response() *GetAllNetworkFloatingIps200Response {
	this := GetAllNetworkFloatingIps200Response{}
	return &this
}

// NewGetAllNetworkFloatingIps200ResponseWithDefaults instantiates a new GetAllNetworkFloatingIps200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllNetworkFloatingIps200ResponseWithDefaults() *GetAllNetworkFloatingIps200Response {
	this := GetAllNetworkFloatingIps200Response{}
	return &this
}

// GetNetworkFloatingIps returns the NetworkFloatingIps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetAllNetworkFloatingIps200Response) GetNetworkFloatingIps() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NetworkFloatingIps
}

// GetNetworkFloatingIpsOk returns a tuple with the NetworkFloatingIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetAllNetworkFloatingIps200Response) GetNetworkFloatingIpsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NetworkFloatingIps) {
		return nil, false
	}
	return &o.NetworkFloatingIps, true
}

// IsSetNetworkFloatingIps returns a boolean if a field has been set.
func (o *GetAllNetworkFloatingIps200Response) IsSetNetworkFloatingIps() bool {
	if o != nil && !IsNil(o.NetworkFloatingIps) {
		return true
	}

	return false
}

// SetNetworkFloatingIps gets a reference to the given interface{} and assigns it to the NetworkFloatingIps field.
func (o *GetAllNetworkFloatingIps200Response) SetNetworkFloatingIps(v interface{}) {
	o.NetworkFloatingIps = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetAllNetworkFloatingIps200Response) GetMeta() ListActivity200ResponseAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListActivity200ResponseAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllNetworkFloatingIps200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// IsSetMeta returns a boolean if a field has been set.
func (o *GetAllNetworkFloatingIps200Response) IsSetMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListActivity200ResponseAllOfMeta and assigns it to the Meta field.
func (o *GetAllNetworkFloatingIps200Response) SetMeta(v ListActivity200ResponseAllOfMeta) {
	o.Meta = &v
}

func (o GetAllNetworkFloatingIps200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllNetworkFloatingIps200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkFloatingIps != nil {
		toSerialize["networkFloatingIps"] = o.NetworkFloatingIps
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetAllNetworkFloatingIps200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the GetNetworkRouter200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkRouter200Response{}

// GetNetworkRouter200Response struct for GetNetworkRouter200Response
type GetNetworkRouter200Response struct {
	NetworkRouter        *GetNetworkRouter200ResponseNetworkRouter `json:"networkRouter,omitempty"`
	AdditionalProperties map[string]interface{}                    `json:",remain"`
}

type _GetNetworkRouter200Response GetNetworkRouter200Response

// NewGetNetworkRouter200Response instantiates a new GetNetworkRouter200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkRouter200Response() *GetNetworkRouter200Response {
	this := GetNetworkRouter200Response{}
	return &this
}

// NewGetNetworkRouter200ResponseWithDefaults instantiates a new GetNetworkRouter200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkRouter200ResponseWithDefaults() *GetNetworkRouter200Response {
	this := GetNetworkRouter200Response{}
	return &this
}

// GetNetworkRouter returns the NetworkRouter field value if set, zero value otherwise.
func (o *GetNetworkRouter200Response) GetNetworkRouter() GetNetworkRouter200ResponseNetworkRouter {
	if o == nil || IsNil(o.NetworkRouter) {
		var ret GetNetworkRouter200ResponseNetworkRouter
		return ret
	}
	return *o.NetworkRouter
}

// GetNetworkRouterOk returns a tuple with the NetworkRouter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200Response) GetNetworkRouterOk() (*GetNetworkRouter200ResponseNetworkRouter, bool) {
	if o == nil || IsNil(o.NetworkRouter) {
		return nil, false
	}
	return o.NetworkRouter, true
}

// IsSetNetworkRouter returns a boolean if a field has been set.
func (o *GetNetworkRouter200Response) IsSetNetworkRouter() bool {
	if o != nil && !IsNil(o.NetworkRouter) {
		return true
	}

	return false
}

// SetNetworkRouter gets a reference to the given GetNetworkRouter200ResponseNetworkRouter and assigns it to the NetworkRouter field.
func (o *GetNetworkRouter200Response) SetNetworkRouter(v GetNetworkRouter200ResponseNetworkRouter) {
	o.NetworkRouter = &v
}

func (o GetNetworkRouter200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkRouter200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkRouter) {
		toSerialize["networkRouter"] = o.NetworkRouter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetNetworkRouter200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the CreateNetworkRouterRouteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkRouterRouteRequest{}

// CreateNetworkRouterRouteRequest struct for CreateNetworkRouterRouteRequest
type CreateNetworkRouterRouteRequest struct {
	NetworkRoute         *CreateNetworkRouterRouteRequestNetworkRoute `json:"networkRoute,omitempty"`
	AdditionalProperties map[string]interface{}                       `json:",remain"`
}

type _CreateNetworkRouterRouteRequest CreateNetworkRouterRouteRequest

// NewCreateNetworkRouterRouteRequest instantiates a new CreateNetworkRouterRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkRouterRouteRequest() *CreateNetworkRouterRouteRequest {
	this := CreateNetworkRouterRouteRequest{}
	return &this
}

// NewCreateNetworkRouterRouteRequestWithDefaults instantiates a new CreateNetworkRouterRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkRouterRouteRequestWithDefaults() *CreateNetworkRouterRouteRequest {
	this := CreateNetworkRouterRouteRequest{}
	return &this
}

// GetNetworkRoute returns the NetworkRoute field value if set, zero value otherwise.
func (o *CreateNetworkRouterRouteRequest) GetNetworkRoute() CreateNetworkRouterRouteRequestNetworkRoute {
	if o == nil || IsNil(o.NetworkRoute) {
		var ret CreateNetworkRouterRouteRequestNetworkRoute
		return ret
	}
	return *o.NetworkRoute
}

// GetNetworkRouteOk returns a tuple with the NetworkRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkRouterRouteRequest) GetNetworkRouteOk() (*CreateNetworkRouterRouteRequestNetworkRoute, bool) {
	if o == nil || IsNil(o.NetworkRoute) {
		return nil, false
	}
	return o.NetworkRoute, true
}

// IsSetNetworkRoute returns a boolean if a field has been set.
func (o *CreateNetworkRouterRouteRequest) IsSetNetworkRoute() bool {
	if o != nil && !IsNil(o.NetworkRoute) {
		return true
	}

	return false
}

// SetNetworkRoute gets a reference to the given CreateNetworkRouterRouteRequestNetworkRoute and assigns it to the NetworkRoute field.
func (o *CreateNetworkRouterRouteRequest) SetNetworkRoute(v CreateNetworkRouterRouteRequestNetworkRoute) {
	o.NetworkRoute = &v
}

func (o CreateNetworkRouterRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkRouterRouteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkRoute) {
		toSerialize["networkRoute"] = o.NetworkRoute
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *CreateNetworkRouterRouteRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

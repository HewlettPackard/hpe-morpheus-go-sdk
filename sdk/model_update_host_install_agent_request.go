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

// checks if the UpdateHostInstallAgentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateHostInstallAgentRequest{}

// UpdateHostInstallAgentRequest struct for UpdateHostInstallAgentRequest
type UpdateHostInstallAgentRequest struct {
	Server               *UpdateHostInstallAgentRequestServer `json:"server,omitempty"`
	AdditionalProperties map[string]interface{}               `json:",remain"`
}

type _UpdateHostInstallAgentRequest UpdateHostInstallAgentRequest

// NewUpdateHostInstallAgentRequest instantiates a new UpdateHostInstallAgentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateHostInstallAgentRequest() *UpdateHostInstallAgentRequest {
	this := UpdateHostInstallAgentRequest{}
	return &this
}

// NewUpdateHostInstallAgentRequestWithDefaults instantiates a new UpdateHostInstallAgentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateHostInstallAgentRequestWithDefaults() *UpdateHostInstallAgentRequest {
	this := UpdateHostInstallAgentRequest{}
	return &this
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *UpdateHostInstallAgentRequest) GetServer() UpdateHostInstallAgentRequestServer {
	if o == nil || IsNil(o.Server) {
		var ret UpdateHostInstallAgentRequestServer
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostInstallAgentRequest) GetServerOk() (*UpdateHostInstallAgentRequestServer, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// IsSetServer returns a boolean if a field has been set.
func (o *UpdateHostInstallAgentRequest) IsSetServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given UpdateHostInstallAgentRequestServer and assigns it to the Server field.
func (o *UpdateHostInstallAgentRequest) SetServer(v UpdateHostInstallAgentRequestServer) {
	o.Server = &v
}

func (o UpdateHostInstallAgentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateHostInstallAgentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateHostInstallAgentRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

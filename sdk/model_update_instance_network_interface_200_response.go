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

// checks if the UpdateInstanceNetworkInterface200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInstanceNetworkInterface200Response{}

// UpdateInstanceNetworkInterface200Response struct for UpdateInstanceNetworkInterface200Response
type UpdateInstanceNetworkInterface200Response struct {
	NetworkInterface     *UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface `json:"networkInterface,omitempty"`
	InterfaceType        *string                                                              `json:"interfaceType,omitempty"`
	NetId                *int64                                                               `json:"netId,omitempty"`
	Server               *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer           `json:"server,omitempty"`
	Success              *bool                                                                `json:"success,omitempty"`
	Errors               map[string]interface{}                                               `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}                                               `json:",remain"`
}

type _UpdateInstanceNetworkInterface200Response UpdateInstanceNetworkInterface200Response

// NewUpdateInstanceNetworkInterface200Response instantiates a new UpdateInstanceNetworkInterface200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInstanceNetworkInterface200Response() *UpdateInstanceNetworkInterface200Response {
	this := UpdateInstanceNetworkInterface200Response{}
	return &this
}

// NewUpdateInstanceNetworkInterface200ResponseWithDefaults instantiates a new UpdateInstanceNetworkInterface200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInstanceNetworkInterface200ResponseWithDefaults() *UpdateInstanceNetworkInterface200Response {
	this := UpdateInstanceNetworkInterface200Response{}
	return &this
}

// GetNetworkInterface returns the NetworkInterface field value if set, zero value otherwise.
func (o *UpdateInstanceNetworkInterface200Response) GetNetworkInterface() UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface {
	if o == nil || IsNil(o.NetworkInterface) {
		var ret UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface
		return ret
	}
	return *o.NetworkInterface
}

// GetNetworkInterfaceOk returns a tuple with the NetworkInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceNetworkInterface200Response) GetNetworkInterfaceOk() (*UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface, bool) {
	if o == nil || IsNil(o.NetworkInterface) {
		return nil, false
	}
	return o.NetworkInterface, true
}

// IsSetNetworkInterface returns a boolean if a field has been set.
func (o *UpdateInstanceNetworkInterface200Response) IsSetNetworkInterface() bool {
	if o != nil && !IsNil(o.NetworkInterface) {
		return true
	}

	return false
}

// SetNetworkInterface gets a reference to the given UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface and assigns it to the NetworkInterface field.
func (o *UpdateInstanceNetworkInterface200Response) SetNetworkInterface(v UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface) {
	o.NetworkInterface = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *UpdateInstanceNetworkInterface200Response) GetInterfaceType() string {
	if o == nil || IsNil(o.InterfaceType) {
		var ret string
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceNetworkInterface200Response) GetInterfaceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceType) {
		return nil, false
	}
	return o.InterfaceType, true
}

// IsSetInterfaceType returns a boolean if a field has been set.
func (o *UpdateInstanceNetworkInterface200Response) IsSetInterfaceType() bool {
	if o != nil && !IsNil(o.InterfaceType) {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.
func (o *UpdateInstanceNetworkInterface200Response) SetInterfaceType(v string) {
	o.InterfaceType = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *UpdateInstanceNetworkInterface200Response) GetNetId() int64 {
	if o == nil || IsNil(o.NetId) {
		var ret int64
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceNetworkInterface200Response) GetNetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.NetId) {
		return nil, false
	}
	return o.NetId, true
}

// IsSetNetId returns a boolean if a field has been set.
func (o *UpdateInstanceNetworkInterface200Response) IsSetNetId() bool {
	if o != nil && !IsNil(o.NetId) {
		return true
	}

	return false
}

// SetNetId gets a reference to the given int64 and assigns it to the NetId field.
func (o *UpdateInstanceNetworkInterface200Response) SetNetId(v int64) {
	o.NetId = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *UpdateInstanceNetworkInterface200Response) GetServer() UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer {
	if o == nil || IsNil(o.Server) {
		var ret UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceNetworkInterface200Response) GetServerOk() (*UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// IsSetServer returns a boolean if a field has been set.
func (o *UpdateInstanceNetworkInterface200Response) IsSetServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer and assigns it to the Server field.
func (o *UpdateInstanceNetworkInterface200Response) SetServer(v UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer) {
	o.Server = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UpdateInstanceNetworkInterface200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceNetworkInterface200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *UpdateInstanceNetworkInterface200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UpdateInstanceNetworkInterface200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *UpdateInstanceNetworkInterface200Response) GetErrors() map[string]interface{} {
	if o == nil || IsNil(o.Errors) {
		var ret map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceNetworkInterface200Response) GetErrorsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Errors) {
		return map[string]interface{}{}, false
	}
	return o.Errors, true
}

// IsSetErrors returns a boolean if a field has been set.
func (o *UpdateInstanceNetworkInterface200Response) IsSetErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string]interface{} and assigns it to the Errors field.
func (o *UpdateInstanceNetworkInterface200Response) SetErrors(v map[string]interface{}) {
	o.Errors = v
}

func (o UpdateInstanceNetworkInterface200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInstanceNetworkInterface200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkInterface) {
		toSerialize["networkInterface"] = o.NetworkInterface
	}
	if !IsNil(o.InterfaceType) {
		toSerialize["interfaceType"] = o.InterfaceType
	}
	if !IsNil(o.NetId) {
		toSerialize["netId"] = o.NetId
	}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateInstanceNetworkInterface200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

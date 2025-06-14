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

// checks if the CreateNetworks200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworks200Response{}

// CreateNetworks200Response struct for CreateNetworks200Response
type CreateNetworks200Response struct {
	Network              *ListNetworks200ResponseAllOfNetworksInner `json:"network,omitempty"`
	Errors               map[string]interface{}                     `json:"errors,omitempty"`
	Success              *bool                                      `json:"success,omitempty"`
	Msg                  *string                                    `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}                     `json:",remain"`
}

type _CreateNetworks200Response CreateNetworks200Response

// NewCreateNetworks200Response instantiates a new CreateNetworks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworks200Response() *CreateNetworks200Response {
	this := CreateNetworks200Response{}
	return &this
}

// NewCreateNetworks200ResponseWithDefaults instantiates a new CreateNetworks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworks200ResponseWithDefaults() *CreateNetworks200Response {
	this := CreateNetworks200Response{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CreateNetworks200Response) GetNetwork() ListNetworks200ResponseAllOfNetworksInner {
	if o == nil || IsNil(o.Network) {
		var ret ListNetworks200ResponseAllOfNetworksInner
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworks200Response) GetNetworkOk() (*ListNetworks200ResponseAllOfNetworksInner, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// IsSetNetwork returns a boolean if a field has been set.
func (o *CreateNetworks200Response) IsSetNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ListNetworks200ResponseAllOfNetworksInner and assigns it to the Network field.
func (o *CreateNetworks200Response) SetNetwork(v ListNetworks200ResponseAllOfNetworksInner) {
	o.Network = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateNetworks200Response) GetErrors() map[string]interface{} {
	if o == nil || IsNil(o.Errors) {
		var ret map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworks200Response) GetErrorsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Errors) {
		return map[string]interface{}{}, false
	}
	return o.Errors, true
}

// IsSetErrors returns a boolean if a field has been set.
func (o *CreateNetworks200Response) IsSetErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string]interface{} and assigns it to the Errors field.
func (o *CreateNetworks200Response) SetErrors(v map[string]interface{}) {
	o.Errors = v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CreateNetworks200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworks200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *CreateNetworks200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CreateNetworks200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CreateNetworks200Response) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworks200Response) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// IsSetMsg returns a boolean if a field has been set.
func (o *CreateNetworks200Response) IsSetMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CreateNetworks200Response) SetMsg(v string) {
	o.Msg = &v
}

func (o CreateNetworks200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworks200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *CreateNetworks200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

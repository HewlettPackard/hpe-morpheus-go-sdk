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

// checks if the CreateLoadBalancerPool200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLoadBalancerPool200Response{}

// CreateLoadBalancerPool200Response struct for CreateLoadBalancerPool200Response
type CreateLoadBalancerPool200Response struct {
	LoadBalancerPool     *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner `json:"loadBalancerPool,omitempty"`
	Success              *bool                                                        `json:"success,omitempty"`
	Msg                  *string                                                      `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}                                       `json:",remain"`
}

type _CreateLoadBalancerPool200Response CreateLoadBalancerPool200Response

// NewCreateLoadBalancerPool200Response instantiates a new CreateLoadBalancerPool200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLoadBalancerPool200Response() *CreateLoadBalancerPool200Response {
	this := CreateLoadBalancerPool200Response{}
	return &this
}

// NewCreateLoadBalancerPool200ResponseWithDefaults instantiates a new CreateLoadBalancerPool200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLoadBalancerPool200ResponseWithDefaults() *CreateLoadBalancerPool200Response {
	this := CreateLoadBalancerPool200Response{}
	return &this
}

// GetLoadBalancerPool returns the LoadBalancerPool field value if set, zero value otherwise.
func (o *CreateLoadBalancerPool200Response) GetLoadBalancerPool() ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner {
	if o == nil || IsNil(o.LoadBalancerPool) {
		var ret ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner
		return ret
	}
	return *o.LoadBalancerPool
}

// GetLoadBalancerPoolOk returns a tuple with the LoadBalancerPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPool200Response) GetLoadBalancerPoolOk() (*ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner, bool) {
	if o == nil || IsNil(o.LoadBalancerPool) {
		return nil, false
	}
	return o.LoadBalancerPool, true
}

// IsSetLoadBalancerPool returns a boolean if a field has been set.
func (o *CreateLoadBalancerPool200Response) IsSetLoadBalancerPool() bool {
	if o != nil && !IsNil(o.LoadBalancerPool) {
		return true
	}

	return false
}

// SetLoadBalancerPool gets a reference to the given ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner and assigns it to the LoadBalancerPool field.
func (o *CreateLoadBalancerPool200Response) SetLoadBalancerPool(v ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) {
	o.LoadBalancerPool = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CreateLoadBalancerPool200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPool200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *CreateLoadBalancerPool200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CreateLoadBalancerPool200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CreateLoadBalancerPool200Response) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPool200Response) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// IsSetMsg returns a boolean if a field has been set.
func (o *CreateLoadBalancerPool200Response) IsSetMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CreateLoadBalancerPool200Response) SetMsg(v string) {
	o.Msg = &v
}

func (o CreateLoadBalancerPool200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLoadBalancerPool200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LoadBalancerPool) {
		toSerialize["loadBalancerPool"] = o.LoadBalancerPool
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
func (o *CreateLoadBalancerPool200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

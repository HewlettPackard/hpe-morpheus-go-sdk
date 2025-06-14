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

// checks if the CreateLoadBalancerMonitor200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLoadBalancerMonitor200Response{}

// CreateLoadBalancerMonitor200Response struct for CreateLoadBalancerMonitor200Response
type CreateLoadBalancerMonitor200Response struct {
	LoadBalancerMonitor  *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner `json:"loadBalancerMonitor,omitempty"`
	Success              *bool                                                              `json:"success,omitempty"`
	Msg                  *string                                                            `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}                                             `json:",remain"`
}

type _CreateLoadBalancerMonitor200Response CreateLoadBalancerMonitor200Response

// NewCreateLoadBalancerMonitor200Response instantiates a new CreateLoadBalancerMonitor200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLoadBalancerMonitor200Response() *CreateLoadBalancerMonitor200Response {
	this := CreateLoadBalancerMonitor200Response{}
	return &this
}

// NewCreateLoadBalancerMonitor200ResponseWithDefaults instantiates a new CreateLoadBalancerMonitor200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLoadBalancerMonitor200ResponseWithDefaults() *CreateLoadBalancerMonitor200Response {
	this := CreateLoadBalancerMonitor200Response{}
	return &this
}

// GetLoadBalancerMonitor returns the LoadBalancerMonitor field value if set, zero value otherwise.
func (o *CreateLoadBalancerMonitor200Response) GetLoadBalancerMonitor() ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner {
	if o == nil || IsNil(o.LoadBalancerMonitor) {
		var ret ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner
		return ret
	}
	return *o.LoadBalancerMonitor
}

// GetLoadBalancerMonitorOk returns a tuple with the LoadBalancerMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerMonitor200Response) GetLoadBalancerMonitorOk() (*ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner, bool) {
	if o == nil || IsNil(o.LoadBalancerMonitor) {
		return nil, false
	}
	return o.LoadBalancerMonitor, true
}

// IsSetLoadBalancerMonitor returns a boolean if a field has been set.
func (o *CreateLoadBalancerMonitor200Response) IsSetLoadBalancerMonitor() bool {
	if o != nil && !IsNil(o.LoadBalancerMonitor) {
		return true
	}

	return false
}

// SetLoadBalancerMonitor gets a reference to the given ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner and assigns it to the LoadBalancerMonitor field.
func (o *CreateLoadBalancerMonitor200Response) SetLoadBalancerMonitor(v ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner) {
	o.LoadBalancerMonitor = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CreateLoadBalancerMonitor200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerMonitor200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *CreateLoadBalancerMonitor200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CreateLoadBalancerMonitor200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CreateLoadBalancerMonitor200Response) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerMonitor200Response) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// IsSetMsg returns a boolean if a field has been set.
func (o *CreateLoadBalancerMonitor200Response) IsSetMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CreateLoadBalancerMonitor200Response) SetMsg(v string) {
	o.Msg = &v
}

func (o CreateLoadBalancerMonitor200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLoadBalancerMonitor200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LoadBalancerMonitor) {
		toSerialize["loadBalancerMonitor"] = o.LoadBalancerMonitor
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
func (o *CreateLoadBalancerMonitor200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

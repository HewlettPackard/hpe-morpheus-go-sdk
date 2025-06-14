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

// checks if the CreateLoadBalancerVirtualServer200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLoadBalancerVirtualServer200Response{}

// CreateLoadBalancerVirtualServer200Response struct for CreateLoadBalancerVirtualServer200Response
type CreateLoadBalancerVirtualServer200Response struct {
	LoadBalancerInstance *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner `json:"loadBalancerInstance,omitempty"`
	AdditionalProperties map[string]interface{}                                                    `json:",remain"`
}

type _CreateLoadBalancerVirtualServer200Response CreateLoadBalancerVirtualServer200Response

// NewCreateLoadBalancerVirtualServer200Response instantiates a new CreateLoadBalancerVirtualServer200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLoadBalancerVirtualServer200Response() *CreateLoadBalancerVirtualServer200Response {
	this := CreateLoadBalancerVirtualServer200Response{}
	return &this
}

// NewCreateLoadBalancerVirtualServer200ResponseWithDefaults instantiates a new CreateLoadBalancerVirtualServer200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLoadBalancerVirtualServer200ResponseWithDefaults() *CreateLoadBalancerVirtualServer200Response {
	this := CreateLoadBalancerVirtualServer200Response{}
	return &this
}

// GetLoadBalancerInstance returns the LoadBalancerInstance field value if set, zero value otherwise.
func (o *CreateLoadBalancerVirtualServer200Response) GetLoadBalancerInstance() ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner {
	if o == nil || IsNil(o.LoadBalancerInstance) {
		var ret ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner
		return ret
	}
	return *o.LoadBalancerInstance
}

// GetLoadBalancerInstanceOk returns a tuple with the LoadBalancerInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerVirtualServer200Response) GetLoadBalancerInstanceOk() (*ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner, bool) {
	if o == nil || IsNil(o.LoadBalancerInstance) {
		return nil, false
	}
	return o.LoadBalancerInstance, true
}

// IsSetLoadBalancerInstance returns a boolean if a field has been set.
func (o *CreateLoadBalancerVirtualServer200Response) IsSetLoadBalancerInstance() bool {
	if o != nil && !IsNil(o.LoadBalancerInstance) {
		return true
	}

	return false
}

// SetLoadBalancerInstance gets a reference to the given ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner and assigns it to the LoadBalancerInstance field.
func (o *CreateLoadBalancerVirtualServer200Response) SetLoadBalancerInstance(v ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) {
	o.LoadBalancerInstance = &v
}

func (o CreateLoadBalancerVirtualServer200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLoadBalancerVirtualServer200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LoadBalancerInstance) {
		toSerialize["loadBalancerInstance"] = o.LoadBalancerInstance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *CreateLoadBalancerVirtualServer200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

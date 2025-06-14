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

// checks if the ListLoadBalancerVirtualServers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListLoadBalancerVirtualServers200Response{}

// ListLoadBalancerVirtualServers200Response struct for ListLoadBalancerVirtualServers200Response
type ListLoadBalancerVirtualServers200Response struct {
	LoadBalancerInstances []ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner `json:"loadBalancerInstances,omitempty"`
	Meta                  *ListActivity200ResponseAllOfMeta                                          `json:"meta,omitempty"`
	AdditionalProperties  map[string]interface{}                                                     `json:",remain"`
}

type _ListLoadBalancerVirtualServers200Response ListLoadBalancerVirtualServers200Response

// NewListLoadBalancerVirtualServers200Response instantiates a new ListLoadBalancerVirtualServers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLoadBalancerVirtualServers200Response() *ListLoadBalancerVirtualServers200Response {
	this := ListLoadBalancerVirtualServers200Response{}
	return &this
}

// NewListLoadBalancerVirtualServers200ResponseWithDefaults instantiates a new ListLoadBalancerVirtualServers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLoadBalancerVirtualServers200ResponseWithDefaults() *ListLoadBalancerVirtualServers200Response {
	this := ListLoadBalancerVirtualServers200Response{}
	return &this
}

// GetLoadBalancerInstances returns the LoadBalancerInstances field value if set, zero value otherwise.
func (o *ListLoadBalancerVirtualServers200Response) GetLoadBalancerInstances() []ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner {
	if o == nil || IsNil(o.LoadBalancerInstances) {
		var ret []ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner
		return ret
	}
	return o.LoadBalancerInstances
}

// GetLoadBalancerInstancesOk returns a tuple with the LoadBalancerInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLoadBalancerVirtualServers200Response) GetLoadBalancerInstancesOk() ([]ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner, bool) {
	if o == nil || IsNil(o.LoadBalancerInstances) {
		return nil, false
	}
	return o.LoadBalancerInstances, true
}

// IsSetLoadBalancerInstances returns a boolean if a field has been set.
func (o *ListLoadBalancerVirtualServers200Response) IsSetLoadBalancerInstances() bool {
	if o != nil && !IsNil(o.LoadBalancerInstances) {
		return true
	}

	return false
}

// SetLoadBalancerInstances gets a reference to the given []ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner and assigns it to the LoadBalancerInstances field.
func (o *ListLoadBalancerVirtualServers200Response) SetLoadBalancerInstances(v []ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) {
	o.LoadBalancerInstances = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListLoadBalancerVirtualServers200Response) GetMeta() ListActivity200ResponseAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListActivity200ResponseAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLoadBalancerVirtualServers200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// IsSetMeta returns a boolean if a field has been set.
func (o *ListLoadBalancerVirtualServers200Response) IsSetMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListActivity200ResponseAllOfMeta and assigns it to the Meta field.
func (o *ListLoadBalancerVirtualServers200Response) SetMeta(v ListActivity200ResponseAllOfMeta) {
	o.Meta = &v
}

func (o ListLoadBalancerVirtualServers200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListLoadBalancerVirtualServers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LoadBalancerInstances) {
		toSerialize["loadBalancerInstances"] = o.LoadBalancerInstances
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListLoadBalancerVirtualServers200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

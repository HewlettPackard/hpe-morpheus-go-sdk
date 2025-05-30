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

// checks if the GetLoadBalancerType200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLoadBalancerType200Response{}

// GetLoadBalancerType200Response struct for GetLoadBalancerType200Response
type GetLoadBalancerType200Response struct {
	LoadBalancerType *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner `json:"loadBalancerType,omitempty"`
}

// NewGetLoadBalancerType200Response instantiates a new GetLoadBalancerType200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLoadBalancerType200Response() *GetLoadBalancerType200Response {
	this := GetLoadBalancerType200Response{}
	return &this
}

// NewGetLoadBalancerType200ResponseWithDefaults instantiates a new GetLoadBalancerType200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLoadBalancerType200ResponseWithDefaults() *GetLoadBalancerType200Response {
	this := GetLoadBalancerType200Response{}
	return &this
}

// GetLoadBalancerType returns the LoadBalancerType field value if set, zero value otherwise.
func (o *GetLoadBalancerType200Response) GetLoadBalancerType() ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner {
	if o == nil || IsNil(o.LoadBalancerType) {
		var ret ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner
		return ret
	}
	return *o.LoadBalancerType
}

// GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoadBalancerType200Response) GetLoadBalancerTypeOk() (*ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner, bool) {
	if o == nil || IsNil(o.LoadBalancerType) {
		return nil, false
	}
	return o.LoadBalancerType, true
}

// IsSetLoadBalancerType returns a boolean if a field has been set.
func (o *GetLoadBalancerType200Response) IsSetLoadBalancerType() bool {
	if o != nil && !IsNil(o.LoadBalancerType) {
		return true
	}

	return false
}

// SetLoadBalancerType gets a reference to the given ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner and assigns it to the LoadBalancerType field.
func (o *GetLoadBalancerType200Response) SetLoadBalancerType(v ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) {
	o.LoadBalancerType = &v
}

func (o GetLoadBalancerType200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLoadBalancerType200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LoadBalancerType) {
		toSerialize["loadBalancerType"] = o.LoadBalancerType
	}
	return toSerialize, nil
}

type NullableGetLoadBalancerType200Response struct {
	value *GetLoadBalancerType200Response
	isSet bool
}

func (v NullableGetLoadBalancerType200Response) Get() *GetLoadBalancerType200Response {
	return v.value
}

func (v *NullableGetLoadBalancerType200Response) Set(val *GetLoadBalancerType200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLoadBalancerType200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLoadBalancerType200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLoadBalancerType200Response(val *GetLoadBalancerType200Response) *NullableGetLoadBalancerType200Response {
	return &NullableGetLoadBalancerType200Response{value: val, isSet: true}
}

func (v NullableGetLoadBalancerType200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLoadBalancerType200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



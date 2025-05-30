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

// checks if the GetServicePlans200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServicePlans200Response{}

// GetServicePlans200Response struct for GetServicePlans200Response
type GetServicePlans200Response struct {
	ServicePlan *ListServicePlans200ResponseAllOfServicePlansInner `json:"servicePlan,omitempty"`
}

// NewGetServicePlans200Response instantiates a new GetServicePlans200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServicePlans200Response() *GetServicePlans200Response {
	this := GetServicePlans200Response{}
	return &this
}

// NewGetServicePlans200ResponseWithDefaults instantiates a new GetServicePlans200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServicePlans200ResponseWithDefaults() *GetServicePlans200Response {
	this := GetServicePlans200Response{}
	return &this
}

// GetServicePlan returns the ServicePlan field value if set, zero value otherwise.
func (o *GetServicePlans200Response) GetServicePlan() ListServicePlans200ResponseAllOfServicePlansInner {
	if o == nil || IsNil(o.ServicePlan) {
		var ret ListServicePlans200ResponseAllOfServicePlansInner
		return ret
	}
	return *o.ServicePlan
}

// GetServicePlanOk returns a tuple with the ServicePlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServicePlans200Response) GetServicePlanOk() (*ListServicePlans200ResponseAllOfServicePlansInner, bool) {
	if o == nil || IsNil(o.ServicePlan) {
		return nil, false
	}
	return o.ServicePlan, true
}

// IsSetServicePlan returns a boolean if a field has been set.
func (o *GetServicePlans200Response) IsSetServicePlan() bool {
	if o != nil && !IsNil(o.ServicePlan) {
		return true
	}

	return false
}

// SetServicePlan gets a reference to the given ListServicePlans200ResponseAllOfServicePlansInner and assigns it to the ServicePlan field.
func (o *GetServicePlans200Response) SetServicePlan(v ListServicePlans200ResponseAllOfServicePlansInner) {
	o.ServicePlan = &v
}

func (o GetServicePlans200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServicePlans200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServicePlan) {
		toSerialize["servicePlan"] = o.ServicePlan
	}
	return toSerialize, nil
}

type NullableGetServicePlans200Response struct {
	value *GetServicePlans200Response
	isSet bool
}

func (v NullableGetServicePlans200Response) Get() *GetServicePlans200Response {
	return v.value
}

func (v *NullableGetServicePlans200Response) Set(val *GetServicePlans200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServicePlans200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServicePlans200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServicePlans200Response(val *GetServicePlans200Response) *NullableGetServicePlans200Response {
	return &NullableGetServicePlans200Response{value: val, isSet: true}
}

func (v NullableGetServicePlans200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServicePlans200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



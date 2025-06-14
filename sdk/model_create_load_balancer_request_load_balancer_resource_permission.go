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

// checks if the CreateLoadBalancerRequestLoadBalancerResourcePermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLoadBalancerRequestLoadBalancerResourcePermission{}

// CreateLoadBalancerRequestLoadBalancerResourcePermission struct for CreateLoadBalancerRequestLoadBalancerResourcePermission
type CreateLoadBalancerRequestLoadBalancerResourcePermission struct {
	// Pass true to allow access to all groups
	All *bool `json:"all,omitempty"`
	// Array of groups that are allowed access
	Sites                []int64                `json:"sites,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _CreateLoadBalancerRequestLoadBalancerResourcePermission CreateLoadBalancerRequestLoadBalancerResourcePermission

// NewCreateLoadBalancerRequestLoadBalancerResourcePermission instantiates a new CreateLoadBalancerRequestLoadBalancerResourcePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLoadBalancerRequestLoadBalancerResourcePermission() *CreateLoadBalancerRequestLoadBalancerResourcePermission {
	this := CreateLoadBalancerRequestLoadBalancerResourcePermission{}
	return &this
}

// NewCreateLoadBalancerRequestLoadBalancerResourcePermissionWithDefaults instantiates a new CreateLoadBalancerRequestLoadBalancerResourcePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLoadBalancerRequestLoadBalancerResourcePermissionWithDefaults() *CreateLoadBalancerRequestLoadBalancerResourcePermission {
	this := CreateLoadBalancerRequestLoadBalancerResourcePermission{}
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *CreateLoadBalancerRequestLoadBalancerResourcePermission) GetAll() bool {
	if o == nil || IsNil(o.All) {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequestLoadBalancerResourcePermission) GetAllOk() (*bool, bool) {
	if o == nil || IsNil(o.All) {
		return nil, false
	}
	return o.All, true
}

// IsSetAll returns a boolean if a field has been set.
func (o *CreateLoadBalancerRequestLoadBalancerResourcePermission) IsSetAll() bool {
	if o != nil && !IsNil(o.All) {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *CreateLoadBalancerRequestLoadBalancerResourcePermission) SetAll(v bool) {
	o.All = &v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *CreateLoadBalancerRequestLoadBalancerResourcePermission) GetSites() []int64 {
	if o == nil || IsNil(o.Sites) {
		var ret []int64
		return ret
	}
	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequestLoadBalancerResourcePermission) GetSitesOk() ([]int64, bool) {
	if o == nil || IsNil(o.Sites) {
		return nil, false
	}
	return o.Sites, true
}

// IsSetSites returns a boolean if a field has been set.
func (o *CreateLoadBalancerRequestLoadBalancerResourcePermission) IsSetSites() bool {
	if o != nil && !IsNil(o.Sites) {
		return true
	}

	return false
}

// SetSites gets a reference to the given []int64 and assigns it to the Sites field.
func (o *CreateLoadBalancerRequestLoadBalancerResourcePermission) SetSites(v []int64) {
	o.Sites = v
}

func (o CreateLoadBalancerRequestLoadBalancerResourcePermission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLoadBalancerRequestLoadBalancerResourcePermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.All) {
		toSerialize["all"] = o.All
	}
	if !IsNil(o.Sites) {
		toSerialize["sites"] = o.Sites
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *CreateLoadBalancerRequestLoadBalancerResourcePermission) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

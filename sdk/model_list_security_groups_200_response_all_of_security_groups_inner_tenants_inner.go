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

// checks if the ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner{}

// ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner struct for ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner
type ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner struct {
	Id                   *int64                 `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	CanManage            *bool                  `json:"canManage,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner

// NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner instantiates a new ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner() *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner {
	this := ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner{}
	return &this
}

// NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInnerWithDefaults instantiates a new ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInnerWithDefaults() *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner {
	this := ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) SetName(v string) {
	o.Name = &v
}

// GetCanManage returns the CanManage field value if set, zero value otherwise.
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) GetCanManage() bool {
	if o == nil || IsNil(o.CanManage) {
		var ret bool
		return ret
	}
	return *o.CanManage
}

// GetCanManageOk returns a tuple with the CanManage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) GetCanManageOk() (*bool, bool) {
	if o == nil || IsNil(o.CanManage) {
		return nil, false
	}
	return o.CanManage, true
}

// IsSetCanManage returns a boolean if a field has been set.
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) IsSetCanManage() bool {
	if o != nil && !IsNil(o.CanManage) {
		return true
	}

	return false
}

// SetCanManage gets a reference to the given bool and assigns it to the CanManage field.
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) SetCanManage(v bool) {
	o.CanManage = &v
}

func (o ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CanManage) {
		toSerialize["canManage"] = o.CanManage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerTenantsInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

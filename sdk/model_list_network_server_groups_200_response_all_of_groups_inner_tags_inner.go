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

// checks if the ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner{}

// ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner struct for ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner
type ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner struct {
	Id                   *int64                 `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Value                *string                `json:"value,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner

// NewListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner instantiates a new ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner() *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner {
	this := ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner{}
	return &this
}

// NewListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInnerWithDefaults instantiates a new ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInnerWithDefaults() *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner {
	this := ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// IsSetValue returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) IsSetValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) SetValue(v string) {
	o.Value = &v
}

func (o ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

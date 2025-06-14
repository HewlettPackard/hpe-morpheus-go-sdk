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

// checks if the ListNetworkServerGroups200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListNetworkServerGroups200Response{}

// ListNetworkServerGroups200Response struct for ListNetworkServerGroups200Response
type ListNetworkServerGroups200Response struct {
	Groups               []ListNetworkServerGroups200ResponseAllOfGroupsInner `json:"groups,omitempty"`
	Meta                 *ListActivity200ResponseAllOfMeta                    `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}                               `json:",remain"`
}

type _ListNetworkServerGroups200Response ListNetworkServerGroups200Response

// NewListNetworkServerGroups200Response instantiates a new ListNetworkServerGroups200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListNetworkServerGroups200Response() *ListNetworkServerGroups200Response {
	this := ListNetworkServerGroups200Response{}
	return &this
}

// NewListNetworkServerGroups200ResponseWithDefaults instantiates a new ListNetworkServerGroups200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListNetworkServerGroups200ResponseWithDefaults() *ListNetworkServerGroups200Response {
	this := ListNetworkServerGroups200Response{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200Response) GetGroups() []ListNetworkServerGroups200ResponseAllOfGroupsInner {
	if o == nil || IsNil(o.Groups) {
		var ret []ListNetworkServerGroups200ResponseAllOfGroupsInner
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200Response) GetGroupsOk() ([]ListNetworkServerGroups200ResponseAllOfGroupsInner, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// IsSetGroups returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200Response) IsSetGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []ListNetworkServerGroups200ResponseAllOfGroupsInner and assigns it to the Groups field.
func (o *ListNetworkServerGroups200Response) SetGroups(v []ListNetworkServerGroups200ResponseAllOfGroupsInner) {
	o.Groups = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200Response) GetMeta() ListActivity200ResponseAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListActivity200ResponseAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// IsSetMeta returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200Response) IsSetMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListActivity200ResponseAllOfMeta and assigns it to the Meta field.
func (o *ListNetworkServerGroups200Response) SetMeta(v ListActivity200ResponseAllOfMeta) {
	o.Meta = &v
}

func (o ListNetworkServerGroups200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListNetworkServerGroups200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListNetworkServerGroups200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

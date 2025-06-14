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

// checks if the ListClusterContainers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListClusterContainers200Response{}

// ListClusterContainers200Response struct for ListClusterContainers200Response
type ListClusterContainers200Response struct {
	Containers           []ListClusterContainers200ResponseAllOfContainersInner `json:"containers,omitempty"`
	Meta                 *ListActivity200ResponseAllOfMeta                      `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}                                 `json:",remain"`
}

type _ListClusterContainers200Response ListClusterContainers200Response

// NewListClusterContainers200Response instantiates a new ListClusterContainers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListClusterContainers200Response() *ListClusterContainers200Response {
	this := ListClusterContainers200Response{}
	return &this
}

// NewListClusterContainers200ResponseWithDefaults instantiates a new ListClusterContainers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListClusterContainers200ResponseWithDefaults() *ListClusterContainers200Response {
	this := ListClusterContainers200Response{}
	return &this
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *ListClusterContainers200Response) GetContainers() []ListClusterContainers200ResponseAllOfContainersInner {
	if o == nil || IsNil(o.Containers) {
		var ret []ListClusterContainers200ResponseAllOfContainersInner
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterContainers200Response) GetContainersOk() ([]ListClusterContainers200ResponseAllOfContainersInner, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// IsSetContainers returns a boolean if a field has been set.
func (o *ListClusterContainers200Response) IsSetContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []ListClusterContainers200ResponseAllOfContainersInner and assigns it to the Containers field.
func (o *ListClusterContainers200Response) SetContainers(v []ListClusterContainers200ResponseAllOfContainersInner) {
	o.Containers = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListClusterContainers200Response) GetMeta() ListActivity200ResponseAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListActivity200ResponseAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterContainers200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// IsSetMeta returns a boolean if a field has been set.
func (o *ListClusterContainers200Response) IsSetMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListActivity200ResponseAllOfMeta and assigns it to the Meta field.
func (o *ListClusterContainers200Response) SetMeta(v ListActivity200ResponseAllOfMeta) {
	o.Meta = &v
}

func (o ListClusterContainers200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListClusterContainers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListClusterContainers200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the ListStorageBuckets200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListStorageBuckets200Response{}

// ListStorageBuckets200Response struct for ListStorageBuckets200Response
type ListStorageBuckets200Response struct {
	StorageBuckets       []ListStorageBuckets200ResponseAllOfStorageBucketsInner `json:"storageBuckets,omitempty"`
	Meta                 *ListActivity200ResponseAllOfMeta                       `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}                                  `json:",remain"`
}

type _ListStorageBuckets200Response ListStorageBuckets200Response

// NewListStorageBuckets200Response instantiates a new ListStorageBuckets200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStorageBuckets200Response() *ListStorageBuckets200Response {
	this := ListStorageBuckets200Response{}
	return &this
}

// NewListStorageBuckets200ResponseWithDefaults instantiates a new ListStorageBuckets200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStorageBuckets200ResponseWithDefaults() *ListStorageBuckets200Response {
	this := ListStorageBuckets200Response{}
	return &this
}

// GetStorageBuckets returns the StorageBuckets field value if set, zero value otherwise.
func (o *ListStorageBuckets200Response) GetStorageBuckets() []ListStorageBuckets200ResponseAllOfStorageBucketsInner {
	if o == nil || IsNil(o.StorageBuckets) {
		var ret []ListStorageBuckets200ResponseAllOfStorageBucketsInner
		return ret
	}
	return o.StorageBuckets
}

// GetStorageBucketsOk returns a tuple with the StorageBuckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStorageBuckets200Response) GetStorageBucketsOk() ([]ListStorageBuckets200ResponseAllOfStorageBucketsInner, bool) {
	if o == nil || IsNil(o.StorageBuckets) {
		return nil, false
	}
	return o.StorageBuckets, true
}

// IsSetStorageBuckets returns a boolean if a field has been set.
func (o *ListStorageBuckets200Response) IsSetStorageBuckets() bool {
	if o != nil && !IsNil(o.StorageBuckets) {
		return true
	}

	return false
}

// SetStorageBuckets gets a reference to the given []ListStorageBuckets200ResponseAllOfStorageBucketsInner and assigns it to the StorageBuckets field.
func (o *ListStorageBuckets200Response) SetStorageBuckets(v []ListStorageBuckets200ResponseAllOfStorageBucketsInner) {
	o.StorageBuckets = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListStorageBuckets200Response) GetMeta() ListActivity200ResponseAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListActivity200ResponseAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStorageBuckets200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// IsSetMeta returns a boolean if a field has been set.
func (o *ListStorageBuckets200Response) IsSetMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListActivity200ResponseAllOfMeta and assigns it to the Meta field.
func (o *ListStorageBuckets200Response) SetMeta(v ListActivity200ResponseAllOfMeta) {
	o.Meta = &v
}

func (o ListStorageBuckets200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListStorageBuckets200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StorageBuckets) {
		toSerialize["storageBuckets"] = o.StorageBuckets
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListStorageBuckets200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the ListClusterDatastores200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListClusterDatastores200Response{}

// ListClusterDatastores200Response struct for ListClusterDatastores200Response
type ListClusterDatastores200Response struct {
	Datastores           []ListClusterDatastores200ResponseAllOfDatastoresInner `json:"datastores,omitempty"`
	Meta                 *ListActivity200ResponseAllOfMeta                      `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}                                 `json:",remain"`
}

type _ListClusterDatastores200Response ListClusterDatastores200Response

// NewListClusterDatastores200Response instantiates a new ListClusterDatastores200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListClusterDatastores200Response() *ListClusterDatastores200Response {
	this := ListClusterDatastores200Response{}
	return &this
}

// NewListClusterDatastores200ResponseWithDefaults instantiates a new ListClusterDatastores200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListClusterDatastores200ResponseWithDefaults() *ListClusterDatastores200Response {
	this := ListClusterDatastores200Response{}
	return &this
}

// GetDatastores returns the Datastores field value if set, zero value otherwise.
func (o *ListClusterDatastores200Response) GetDatastores() []ListClusterDatastores200ResponseAllOfDatastoresInner {
	if o == nil || IsNil(o.Datastores) {
		var ret []ListClusterDatastores200ResponseAllOfDatastoresInner
		return ret
	}
	return o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterDatastores200Response) GetDatastoresOk() ([]ListClusterDatastores200ResponseAllOfDatastoresInner, bool) {
	if o == nil || IsNil(o.Datastores) {
		return nil, false
	}
	return o.Datastores, true
}

// IsSetDatastores returns a boolean if a field has been set.
func (o *ListClusterDatastores200Response) IsSetDatastores() bool {
	if o != nil && !IsNil(o.Datastores) {
		return true
	}

	return false
}

// SetDatastores gets a reference to the given []ListClusterDatastores200ResponseAllOfDatastoresInner and assigns it to the Datastores field.
func (o *ListClusterDatastores200Response) SetDatastores(v []ListClusterDatastores200ResponseAllOfDatastoresInner) {
	o.Datastores = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListClusterDatastores200Response) GetMeta() ListActivity200ResponseAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListActivity200ResponseAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterDatastores200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// IsSetMeta returns a boolean if a field has been set.
func (o *ListClusterDatastores200Response) IsSetMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListActivity200ResponseAllOfMeta and assigns it to the Meta field.
func (o *ListClusterDatastores200Response) SetMeta(v ListActivity200ResponseAllOfMeta) {
	o.Meta = &v
}

func (o ListClusterDatastores200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListClusterDatastores200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Datastores) {
		toSerialize["datastores"] = o.Datastores
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListClusterDatastores200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

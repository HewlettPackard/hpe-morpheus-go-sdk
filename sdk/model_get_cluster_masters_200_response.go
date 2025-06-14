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

// checks if the GetClusterMasters200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetClusterMasters200Response{}

// GetClusterMasters200Response struct for GetClusterMasters200Response
type GetClusterMasters200Response struct {
	Masters              []GetClusterMasters200ResponseMastersInner `json:"masters,omitempty"`
	AdditionalProperties map[string]interface{}                     `json:",remain"`
}

type _GetClusterMasters200Response GetClusterMasters200Response

// NewGetClusterMasters200Response instantiates a new GetClusterMasters200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetClusterMasters200Response() *GetClusterMasters200Response {
	this := GetClusterMasters200Response{}
	return &this
}

// NewGetClusterMasters200ResponseWithDefaults instantiates a new GetClusterMasters200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetClusterMasters200ResponseWithDefaults() *GetClusterMasters200Response {
	this := GetClusterMasters200Response{}
	return &this
}

// GetMasters returns the Masters field value if set, zero value otherwise.
func (o *GetClusterMasters200Response) GetMasters() []GetClusterMasters200ResponseMastersInner {
	if o == nil || IsNil(o.Masters) {
		var ret []GetClusterMasters200ResponseMastersInner
		return ret
	}
	return o.Masters
}

// GetMastersOk returns a tuple with the Masters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200Response) GetMastersOk() ([]GetClusterMasters200ResponseMastersInner, bool) {
	if o == nil || IsNil(o.Masters) {
		return nil, false
	}
	return o.Masters, true
}

// IsSetMasters returns a boolean if a field has been set.
func (o *GetClusterMasters200Response) IsSetMasters() bool {
	if o != nil && !IsNil(o.Masters) {
		return true
	}

	return false
}

// SetMasters gets a reference to the given []GetClusterMasters200ResponseMastersInner and assigns it to the Masters field.
func (o *GetClusterMasters200Response) SetMasters(v []GetClusterMasters200ResponseMastersInner) {
	o.Masters = v
}

func (o GetClusterMasters200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetClusterMasters200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Masters) {
		toSerialize["masters"] = o.Masters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetClusterMasters200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

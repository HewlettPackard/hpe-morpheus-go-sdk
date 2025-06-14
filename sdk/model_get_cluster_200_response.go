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

// checks if the GetCluster200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCluster200Response{}

// GetCluster200Response struct for GetCluster200Response
type GetCluster200Response struct {
	Cluster              *AddCluster200ResponseAllOfCluster `json:"cluster,omitempty"`
	AdditionalProperties map[string]interface{}             `json:",remain"`
}

type _GetCluster200Response GetCluster200Response

// NewGetCluster200Response instantiates a new GetCluster200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCluster200Response() *GetCluster200Response {
	this := GetCluster200Response{}
	return &this
}

// NewGetCluster200ResponseWithDefaults instantiates a new GetCluster200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCluster200ResponseWithDefaults() *GetCluster200Response {
	this := GetCluster200Response{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *GetCluster200Response) GetCluster() AddCluster200ResponseAllOfCluster {
	if o == nil || IsNil(o.Cluster) {
		var ret AddCluster200ResponseAllOfCluster
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCluster200Response) GetClusterOk() (*AddCluster200ResponseAllOfCluster, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// IsSetCluster returns a boolean if a field has been set.
func (o *GetCluster200Response) IsSetCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given AddCluster200ResponseAllOfCluster and assigns it to the Cluster field.
func (o *GetCluster200Response) SetCluster(v AddCluster200ResponseAllOfCluster) {
	o.Cluster = &v
}

func (o GetCluster200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCluster200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetCluster200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

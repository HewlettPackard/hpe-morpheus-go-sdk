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

// checks if the AddCluster200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCluster200Response{}

// AddCluster200Response struct for AddCluster200Response
type AddCluster200Response struct {
	Cluster              *AddCluster200ResponseAllOfCluster `json:"cluster,omitempty"`
	Success              *bool                              `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}             `json:",remain"`
}

type _AddCluster200Response AddCluster200Response

// NewAddCluster200Response instantiates a new AddCluster200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCluster200Response() *AddCluster200Response {
	this := AddCluster200Response{}
	return &this
}

// NewAddCluster200ResponseWithDefaults instantiates a new AddCluster200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCluster200ResponseWithDefaults() *AddCluster200Response {
	this := AddCluster200Response{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *AddCluster200Response) GetCluster() AddCluster200ResponseAllOfCluster {
	if o == nil || IsNil(o.Cluster) {
		var ret AddCluster200ResponseAllOfCluster
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCluster200Response) GetClusterOk() (*AddCluster200ResponseAllOfCluster, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// IsSetCluster returns a boolean if a field has been set.
func (o *AddCluster200Response) IsSetCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given AddCluster200ResponseAllOfCluster and assigns it to the Cluster field.
func (o *AddCluster200Response) SetCluster(v AddCluster200ResponseAllOfCluster) {
	o.Cluster = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AddCluster200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCluster200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *AddCluster200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AddCluster200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o AddCluster200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCluster200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddCluster200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

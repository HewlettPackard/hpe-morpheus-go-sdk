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

// checks if the GetClusterContainer200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetClusterContainer200Response{}

// GetClusterContainer200Response struct for GetClusterContainer200Response
type GetClusterContainer200Response struct {
	Resource             *GetClusterContainer200ResponseResource `json:"resource,omitempty"`
	AdditionalProperties map[string]interface{}                  `json:",remain"`
}

type _GetClusterContainer200Response GetClusterContainer200Response

// NewGetClusterContainer200Response instantiates a new GetClusterContainer200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetClusterContainer200Response() *GetClusterContainer200Response {
	this := GetClusterContainer200Response{}
	return &this
}

// NewGetClusterContainer200ResponseWithDefaults instantiates a new GetClusterContainer200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetClusterContainer200ResponseWithDefaults() *GetClusterContainer200Response {
	this := GetClusterContainer200Response{}
	return &this
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *GetClusterContainer200Response) GetResource() GetClusterContainer200ResponseResource {
	if o == nil || IsNil(o.Resource) {
		var ret GetClusterContainer200ResponseResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterContainer200Response) GetResourceOk() (*GetClusterContainer200ResponseResource, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// IsSetResource returns a boolean if a field has been set.
func (o *GetClusterContainer200Response) IsSetResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given GetClusterContainer200ResponseResource and assigns it to the Resource field.
func (o *GetClusterContainer200Response) SetResource(v GetClusterContainer200ResponseResource) {
	o.Resource = &v
}

func (o GetClusterContainer200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetClusterContainer200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetClusterContainer200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

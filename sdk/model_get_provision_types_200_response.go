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

// checks if the GetProvisionTypes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProvisionTypes200Response{}

// GetProvisionTypes200Response struct for GetProvisionTypes200Response
type GetProvisionTypes200Response struct {
	ProvisionType        *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType `json:"provisionType,omitempty"`
	AdditionalProperties map[string]interface{}                                                                        `json:",remain"`
}

type _GetProvisionTypes200Response GetProvisionTypes200Response

// NewGetProvisionTypes200Response instantiates a new GetProvisionTypes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProvisionTypes200Response() *GetProvisionTypes200Response {
	this := GetProvisionTypes200Response{}
	return &this
}

// NewGetProvisionTypes200ResponseWithDefaults instantiates a new GetProvisionTypes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProvisionTypes200ResponseWithDefaults() *GetProvisionTypes200Response {
	this := GetProvisionTypes200Response{}
	return &this
}

// GetProvisionType returns the ProvisionType field value if set, zero value otherwise.
func (o *GetProvisionTypes200Response) GetProvisionType() GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType {
	if o == nil || IsNil(o.ProvisionType) {
		var ret GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType
		return ret
	}
	return *o.ProvisionType
}

// GetProvisionTypeOk returns a tuple with the ProvisionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProvisionTypes200Response) GetProvisionTypeOk() (*GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType, bool) {
	if o == nil || IsNil(o.ProvisionType) {
		return nil, false
	}
	return o.ProvisionType, true
}

// IsSetProvisionType returns a boolean if a field has been set.
func (o *GetProvisionTypes200Response) IsSetProvisionType() bool {
	if o != nil && !IsNil(o.ProvisionType) {
		return true
	}

	return false
}

// SetProvisionType gets a reference to the given GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType and assigns it to the ProvisionType field.
func (o *GetProvisionTypes200Response) SetProvisionType(v GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType) {
	o.ProvisionType = &v
}

func (o GetProvisionTypes200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProvisionTypes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProvisionType) {
		toSerialize["provisionType"] = o.ProvisionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetProvisionTypes200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

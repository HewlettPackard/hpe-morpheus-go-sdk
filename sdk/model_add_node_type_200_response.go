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

// checks if the AddNodeType200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddNodeType200Response{}

// AddNodeType200Response struct for AddNodeType200Response
type AddNodeType200Response struct {
	ContainerType        *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner `json:"containerType,omitempty"`
	Success              *bool                                                                                               `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}                                                                              `json:",remain"`
}

type _AddNodeType200Response AddNodeType200Response

// NewAddNodeType200Response instantiates a new AddNodeType200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddNodeType200Response() *AddNodeType200Response {
	this := AddNodeType200Response{}
	return &this
}

// NewAddNodeType200ResponseWithDefaults instantiates a new AddNodeType200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddNodeType200ResponseWithDefaults() *AddNodeType200Response {
	this := AddNodeType200Response{}
	return &this
}

// GetContainerType returns the ContainerType field value if set, zero value otherwise.
func (o *AddNodeType200Response) GetContainerType() GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner {
	if o == nil || IsNil(o.ContainerType) {
		var ret GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner
		return ret
	}
	return *o.ContainerType
}

// GetContainerTypeOk returns a tuple with the ContainerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNodeType200Response) GetContainerTypeOk() (*GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner, bool) {
	if o == nil || IsNil(o.ContainerType) {
		return nil, false
	}
	return o.ContainerType, true
}

// IsSetContainerType returns a boolean if a field has been set.
func (o *AddNodeType200Response) IsSetContainerType() bool {
	if o != nil && !IsNil(o.ContainerType) {
		return true
	}

	return false
}

// SetContainerType gets a reference to the given GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner and assigns it to the ContainerType field.
func (o *AddNodeType200Response) SetContainerType(v GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) {
	o.ContainerType = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AddNodeType200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNodeType200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *AddNodeType200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AddNodeType200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o AddNodeType200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddNodeType200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerType) {
		toSerialize["containerType"] = o.ContainerType
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddNodeType200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

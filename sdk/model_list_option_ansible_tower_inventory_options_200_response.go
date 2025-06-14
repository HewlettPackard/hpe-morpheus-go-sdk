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

// checks if the ListOptionAnsibleTowerInventoryOptions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOptionAnsibleTowerInventoryOptions200Response{}

// ListOptionAnsibleTowerInventoryOptions200Response struct for ListOptionAnsibleTowerInventoryOptions200Response
type ListOptionAnsibleTowerInventoryOptions200Response struct {
	Data                 []ListCodeRepositories200ResponseAllOfDataInner `json:"data,omitempty"`
	Success              *bool                                           `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}                          `json:",remain"`
}

type _ListOptionAnsibleTowerInventoryOptions200Response ListOptionAnsibleTowerInventoryOptions200Response

// NewListOptionAnsibleTowerInventoryOptions200Response instantiates a new ListOptionAnsibleTowerInventoryOptions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOptionAnsibleTowerInventoryOptions200Response() *ListOptionAnsibleTowerInventoryOptions200Response {
	this := ListOptionAnsibleTowerInventoryOptions200Response{}
	return &this
}

// NewListOptionAnsibleTowerInventoryOptions200ResponseWithDefaults instantiates a new ListOptionAnsibleTowerInventoryOptions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOptionAnsibleTowerInventoryOptions200ResponseWithDefaults() *ListOptionAnsibleTowerInventoryOptions200Response {
	this := ListOptionAnsibleTowerInventoryOptions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListOptionAnsibleTowerInventoryOptions200Response) GetData() []ListCodeRepositories200ResponseAllOfDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []ListCodeRepositories200ResponseAllOfDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOptionAnsibleTowerInventoryOptions200Response) GetDataOk() ([]ListCodeRepositories200ResponseAllOfDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// IsSetData returns a boolean if a field has been set.
func (o *ListOptionAnsibleTowerInventoryOptions200Response) IsSetData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ListCodeRepositories200ResponseAllOfDataInner and assigns it to the Data field.
func (o *ListOptionAnsibleTowerInventoryOptions200Response) SetData(v []ListCodeRepositories200ResponseAllOfDataInner) {
	o.Data = v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ListOptionAnsibleTowerInventoryOptions200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOptionAnsibleTowerInventoryOptions200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *ListOptionAnsibleTowerInventoryOptions200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ListOptionAnsibleTowerInventoryOptions200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o ListOptionAnsibleTowerInventoryOptions200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOptionAnsibleTowerInventoryOptions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListOptionAnsibleTowerInventoryOptions200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

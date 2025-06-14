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

// checks if the UpdateIntegrationInventory200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIntegrationInventory200Response{}

// UpdateIntegrationInventory200Response struct for UpdateIntegrationInventory200Response
type UpdateIntegrationInventory200Response struct {
	Inventory            *ListIntegrationInventory200ResponseAllOfInventoryInner `json:"inventory,omitempty"`
	Success              *bool                                                   `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}                                  `json:",remain"`
}

type _UpdateIntegrationInventory200Response UpdateIntegrationInventory200Response

// NewUpdateIntegrationInventory200Response instantiates a new UpdateIntegrationInventory200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIntegrationInventory200Response() *UpdateIntegrationInventory200Response {
	this := UpdateIntegrationInventory200Response{}
	return &this
}

// NewUpdateIntegrationInventory200ResponseWithDefaults instantiates a new UpdateIntegrationInventory200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIntegrationInventory200ResponseWithDefaults() *UpdateIntegrationInventory200Response {
	this := UpdateIntegrationInventory200Response{}
	return &this
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *UpdateIntegrationInventory200Response) GetInventory() ListIntegrationInventory200ResponseAllOfInventoryInner {
	if o == nil || IsNil(o.Inventory) {
		var ret ListIntegrationInventory200ResponseAllOfInventoryInner
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIntegrationInventory200Response) GetInventoryOk() (*ListIntegrationInventory200ResponseAllOfInventoryInner, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// IsSetInventory returns a boolean if a field has been set.
func (o *UpdateIntegrationInventory200Response) IsSetInventory() bool {
	if o != nil && !IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given ListIntegrationInventory200ResponseAllOfInventoryInner and assigns it to the Inventory field.
func (o *UpdateIntegrationInventory200Response) SetInventory(v ListIntegrationInventory200ResponseAllOfInventoryInner) {
	o.Inventory = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UpdateIntegrationInventory200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIntegrationInventory200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *UpdateIntegrationInventory200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UpdateIntegrationInventory200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o UpdateIntegrationInventory200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIntegrationInventory200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Inventory) {
		toSerialize["inventory"] = o.Inventory
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateIntegrationInventory200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

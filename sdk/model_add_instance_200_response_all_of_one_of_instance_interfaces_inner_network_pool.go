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

// checks if the AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool{}

// AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool struct for AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool
type AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool struct {
	Id                   *int64                 `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool

// NewAddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool instantiates a new AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool() *AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool {
	this := AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool{}
	return &this
}

// NewAddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPoolWithDefaults instantiates a new AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPoolWithDefaults() *AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool {
	this := AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool) SetName(v string) {
	o.Name = &v
}

func (o AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddInstance200ResponseAllOfOneOfInstanceInterfacesInnerNetworkPool) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

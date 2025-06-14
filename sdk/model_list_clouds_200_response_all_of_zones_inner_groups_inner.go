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

// checks if the ListClouds200ResponseAllOfZonesInnerGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListClouds200ResponseAllOfZonesInnerGroupsInner{}

// ListClouds200ResponseAllOfZonesInnerGroupsInner struct for ListClouds200ResponseAllOfZonesInnerGroupsInner
type ListClouds200ResponseAllOfZonesInnerGroupsInner struct {
	Id                   *int64                 `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	AccountId            *int64                 `json:"accountId,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListClouds200ResponseAllOfZonesInnerGroupsInner ListClouds200ResponseAllOfZonesInnerGroupsInner

// NewListClouds200ResponseAllOfZonesInnerGroupsInner instantiates a new ListClouds200ResponseAllOfZonesInnerGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListClouds200ResponseAllOfZonesInnerGroupsInner() *ListClouds200ResponseAllOfZonesInnerGroupsInner {
	this := ListClouds200ResponseAllOfZonesInnerGroupsInner{}
	return &this
}

// NewListClouds200ResponseAllOfZonesInnerGroupsInnerWithDefaults instantiates a new ListClouds200ResponseAllOfZonesInnerGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListClouds200ResponseAllOfZonesInnerGroupsInnerWithDefaults() *ListClouds200ResponseAllOfZonesInnerGroupsInner {
	this := ListClouds200ResponseAllOfZonesInnerGroupsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) SetName(v string) {
	o.Name = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) GetAccountId() int64 {
	if o == nil || IsNil(o.AccountId) {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) GetAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// IsSetAccountId returns a boolean if a field has been set.
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) IsSetAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) SetAccountId(v int64) {
	o.AccountId = &v
}

func (o ListClouds200ResponseAllOfZonesInnerGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListClouds200ResponseAllOfZonesInnerGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListClouds200ResponseAllOfZonesInnerGroupsInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

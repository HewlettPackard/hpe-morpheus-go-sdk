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

// checks if the ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner{}

// ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner struct for ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner
type ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner struct {
	Id                   *int32                 `json:"id,omitempty"`
	AccountId            *int32                 `json:"accountId,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Description          *string                `json:"description,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner

// NewListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner instantiates a new ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner() *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner {
	this := ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner{}
	return &this
}

// NewListCloudSecurityGroups200ResponseAllOfSecurityGroupsInnerWithDefaults instantiates a new ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCloudSecurityGroups200ResponseAllOfSecurityGroupsInnerWithDefaults() *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner {
	this := ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) SetId(v int32) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) GetAccountId() int32 {
	if o == nil || IsNil(o.AccountId) {
		var ret int32
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) GetAccountIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// IsSetAccountId returns a boolean if a field has been set.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) IsSetAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) SetAccountId(v int32) {
	o.AccountId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) SetDescription(v string) {
	o.Description = &v
}

func (o ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

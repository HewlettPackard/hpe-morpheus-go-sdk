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

// checks if the UpdateRoleCatalogItemTypeAccessRequestOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoleCatalogItemTypeAccessRequestOneOf1{}

// UpdateRoleCatalogItemTypeAccessRequestOneOf1 struct for UpdateRoleCatalogItemTypeAccessRequestOneOf1
type UpdateRoleCatalogItemTypeAccessRequestOneOf1 struct {
	// Apply to all catalog item types
	AllCatalogItemTypes bool `json:"allCatalogItemTypes"`
	// The new access level.
	Access               string                 `json:"access"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _UpdateRoleCatalogItemTypeAccessRequestOneOf1 UpdateRoleCatalogItemTypeAccessRequestOneOf1

// NewUpdateRoleCatalogItemTypeAccessRequestOneOf1 instantiates a new UpdateRoleCatalogItemTypeAccessRequestOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleCatalogItemTypeAccessRequestOneOf1(allCatalogItemTypes bool, access string) *UpdateRoleCatalogItemTypeAccessRequestOneOf1 {
	this := UpdateRoleCatalogItemTypeAccessRequestOneOf1{}
	this.AllCatalogItemTypes = allCatalogItemTypes
	this.Access = access
	return &this
}

// NewUpdateRoleCatalogItemTypeAccessRequestOneOf1WithDefaults instantiates a new UpdateRoleCatalogItemTypeAccessRequestOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleCatalogItemTypeAccessRequestOneOf1WithDefaults() *UpdateRoleCatalogItemTypeAccessRequestOneOf1 {
	this := UpdateRoleCatalogItemTypeAccessRequestOneOf1{}
	return &this
}

// GetAllCatalogItemTypes returns the AllCatalogItemTypes field value
func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf1) GetAllCatalogItemTypes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllCatalogItemTypes
}

// GetAllCatalogItemTypesOk returns a tuple with the AllCatalogItemTypes field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf1) GetAllCatalogItemTypesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllCatalogItemTypes, true
}

// SetAllCatalogItemTypes sets field value
func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf1) SetAllCatalogItemTypes(v bool) {
	o.AllCatalogItemTypes = v
}

// GetAccess returns the Access field value
func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf1) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf1) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf1) SetAccess(v string) {
	o.Access = v
}

func (o UpdateRoleCatalogItemTypeAccessRequestOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoleCatalogItemTypeAccessRequestOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allCatalogItemTypes"] = o.AllCatalogItemTypes
	toSerialize["access"] = o.Access

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateRoleCatalogItemTypeAccessRequestOneOf1) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

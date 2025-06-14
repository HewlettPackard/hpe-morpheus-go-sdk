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

// checks if the UpdateTenantSubtenantGroupZonesRequestGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTenantSubtenantGroupZonesRequestGroup{}

// UpdateTenantSubtenantGroupZonesRequestGroup struct for UpdateTenantSubtenantGroupZonesRequestGroup
type UpdateTenantSubtenantGroupZonesRequestGroup struct {
	// An array of all the zones assigned to this group.
	Zones                []AddClusterLayoutsRequestLayoutMastersInnerContainerType `json:"zones"`
	AdditionalProperties map[string]interface{}                                    `json:",remain"`
}

type _UpdateTenantSubtenantGroupZonesRequestGroup UpdateTenantSubtenantGroupZonesRequestGroup

// NewUpdateTenantSubtenantGroupZonesRequestGroup instantiates a new UpdateTenantSubtenantGroupZonesRequestGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTenantSubtenantGroupZonesRequestGroup(zones []AddClusterLayoutsRequestLayoutMastersInnerContainerType) *UpdateTenantSubtenantGroupZonesRequestGroup {
	this := UpdateTenantSubtenantGroupZonesRequestGroup{}
	this.Zones = zones
	return &this
}

// NewUpdateTenantSubtenantGroupZonesRequestGroupWithDefaults instantiates a new UpdateTenantSubtenantGroupZonesRequestGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTenantSubtenantGroupZonesRequestGroupWithDefaults() *UpdateTenantSubtenantGroupZonesRequestGroup {
	this := UpdateTenantSubtenantGroupZonesRequestGroup{}
	return &this
}

// GetZones returns the Zones field value
func (o *UpdateTenantSubtenantGroupZonesRequestGroup) GetZones() []AddClusterLayoutsRequestLayoutMastersInnerContainerType {
	if o == nil {
		var ret []AddClusterLayoutsRequestLayoutMastersInnerContainerType
		return ret
	}

	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value
// and a boolean to check if the value has been set.
func (o *UpdateTenantSubtenantGroupZonesRequestGroup) GetZonesOk() ([]AddClusterLayoutsRequestLayoutMastersInnerContainerType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zones, true
}

// SetZones sets field value
func (o *UpdateTenantSubtenantGroupZonesRequestGroup) SetZones(v []AddClusterLayoutsRequestLayoutMastersInnerContainerType) {
	o.Zones = v
}

func (o UpdateTenantSubtenantGroupZonesRequestGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTenantSubtenantGroupZonesRequestGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["zones"] = o.Zones

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateTenantSubtenantGroupZonesRequestGroup) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the GetResourcePoolGroups200ResponseResourcePoolGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetResourcePoolGroups200ResponseResourcePoolGroupsInner{}

// GetResourcePoolGroups200ResponseResourcePoolGroupsInner struct for GetResourcePoolGroups200ResponseResourcePoolGroupsInner
type GetResourcePoolGroups200ResponseResourcePoolGroupsInner struct {
	Id          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Visibility  *string `json:"visibility,omitempty"`
	// Pool selection mode. Valid values are `roundrobin` or `availablecapacity`.
	Mode *string `json:"mode,omitempty"`
	// Array of Resource Pool IDs
	Pools                []int64                                                                    `json:"pools,omitempty"`
	Tenants              []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner   `json:"tenants,omitempty"`
	ResourcePermission   *GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission `json:"resourcePermission,omitempty"`
	AdditionalProperties map[string]interface{}                                                     `json:",remain"`
}

type _GetResourcePoolGroups200ResponseResourcePoolGroupsInner GetResourcePoolGroups200ResponseResourcePoolGroupsInner

// NewGetResourcePoolGroups200ResponseResourcePoolGroupsInner instantiates a new GetResourcePoolGroups200ResponseResourcePoolGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetResourcePoolGroups200ResponseResourcePoolGroupsInner() *GetResourcePoolGroups200ResponseResourcePoolGroupsInner {
	this := GetResourcePoolGroups200ResponseResourcePoolGroupsInner{}
	return &this
}

// NewGetResourcePoolGroups200ResponseResourcePoolGroupsInnerWithDefaults instantiates a new GetResourcePoolGroups200ResponseResourcePoolGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetResourcePoolGroups200ResponseResourcePoolGroupsInnerWithDefaults() *GetResourcePoolGroups200ResponseResourcePoolGroupsInner {
	this := GetResourcePoolGroups200ResponseResourcePoolGroupsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetDescription(v string) {
	o.Description = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetVisibility(v string) {
	o.Visibility = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// IsSetMode returns a boolean if a field has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) IsSetMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetMode(v string) {
	o.Mode = &v
}

// GetPools returns the Pools field value if set, zero value otherwise.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetPools() []int64 {
	if o == nil || IsNil(o.Pools) {
		var ret []int64
		return ret
	}
	return o.Pools
}

// GetPoolsOk returns a tuple with the Pools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetPoolsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Pools) {
		return nil, false
	}
	return o.Pools, true
}

// IsSetPools returns a boolean if a field has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) IsSetPools() bool {
	if o != nil && !IsNil(o.Pools) {
		return true
	}

	return false
}

// SetPools gets a reference to the given []int64 and assigns it to the Pools field.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetPools(v []int64) {
	o.Pools = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Tenants) {
		var ret []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetTenantsOk() ([]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// IsSetTenants returns a boolean if a field has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) IsSetTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Tenants field.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Tenants = v
}

// GetResourcePermission returns the ResourcePermission field value if set, zero value otherwise.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetResourcePermission() GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission {
	if o == nil || IsNil(o.ResourcePermission) {
		var ret GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission
		return ret
	}
	return *o.ResourcePermission
}

// GetResourcePermissionOk returns a tuple with the ResourcePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetResourcePermissionOk() (*GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission, bool) {
	if o == nil || IsNil(o.ResourcePermission) {
		return nil, false
	}
	return o.ResourcePermission, true
}

// IsSetResourcePermission returns a boolean if a field has been set.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) IsSetResourcePermission() bool {
	if o != nil && !IsNil(o.ResourcePermission) {
		return true
	}

	return false
}

// SetResourcePermission gets a reference to the given GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission and assigns it to the ResourcePermission field.
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetResourcePermission(v GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission) {
	o.ResourcePermission = &v
}

func (o GetResourcePoolGroups200ResponseResourcePoolGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetResourcePoolGroups200ResponseResourcePoolGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Pools) {
		toSerialize["pools"] = o.Pools
	}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}
	if !IsNil(o.ResourcePermission) {
		toSerialize["resourcePermission"] = o.ResourcePermission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

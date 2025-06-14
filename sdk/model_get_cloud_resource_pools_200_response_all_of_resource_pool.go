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

// checks if the GetCloudResourcePools200ResponseAllOfResourcePool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCloudResourcePools200ResponseAllOfResourcePool{}

// GetCloudResourcePools200ResponseAllOfResourcePool struct for GetCloudResourcePools200ResponseAllOfResourcePool
type GetCloudResourcePools200ResponseAllOfResourcePool struct {
	Id                   *int64                                                                   `json:"id,omitempty"`
	Description          *string                                                                  `json:"description,omitempty"`
	Zone                 *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner  `json:"zone,omitempty"`
	Parent               *GetAlerts200ResponseAllOfCheckGroupsInnerInstance                       `json:"parent,omitempty"`
	Type                 *string                                                                  `json:"type,omitempty"`
	ExternalId           *string                                                                  `json:"externalId,omitempty"`
	RegionCode           *string                                                                  `json:"regionCode,omitempty"`
	IacId                *string                                                                  `json:"iacId,omitempty"`
	Visibility           *string                                                                  `json:"visibility,omitempty"`
	ReadOnly             *bool                                                                    `json:"readOnly,omitempty"`
	DefaultPool          *bool                                                                    `json:"defaultPool,omitempty"`
	Active               *bool                                                                    `json:"active,omitempty"`
	Status               *string                                                                  `json:"status,omitempty"`
	Inventory            *bool                                                                    `json:"inventory,omitempty"`
	Config               *AddCloudResourcePool200ResponseResourcePoolAllOfConfig                  `json:"config,omitempty"`
	Name                 *string                                                                  `json:"name,omitempty"`
	DisplayName          *string                                                                  `json:"displayName,omitempty"`
	Tenants              []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"tenants,omitempty"`
	ResourcePermission   *ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission    `json:"resourcePermission,omitempty"`
	Depth                *int64                                                                   `json:"depth,omitempty"`
	AdditionalProperties map[string]interface{}                                                   `json:",remain"`
}

type _GetCloudResourcePools200ResponseAllOfResourcePool GetCloudResourcePools200ResponseAllOfResourcePool

// NewGetCloudResourcePools200ResponseAllOfResourcePool instantiates a new GetCloudResourcePools200ResponseAllOfResourcePool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCloudResourcePools200ResponseAllOfResourcePool() *GetCloudResourcePools200ResponseAllOfResourcePool {
	this := GetCloudResourcePools200ResponseAllOfResourcePool{}
	return &this
}

// NewGetCloudResourcePools200ResponseAllOfResourcePoolWithDefaults instantiates a new GetCloudResourcePools200ResponseAllOfResourcePool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCloudResourcePools200ResponseAllOfResourcePoolWithDefaults() *GetCloudResourcePools200ResponseAllOfResourcePool {
	this := GetCloudResourcePools200ResponseAllOfResourcePool{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetId(v int64) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetDescription(v string) {
	o.Description = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Zone) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// IsSetZone returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Zone field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Zone = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetParent() GetAlerts200ResponseAllOfCheckGroupsInnerInstance {
	if o == nil || IsNil(o.Parent) {
		var ret GetAlerts200ResponseAllOfCheckGroupsInnerInstance
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetParentOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// IsSetParent returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given GetAlerts200ResponseAllOfCheckGroupsInnerInstance and assigns it to the Parent field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetParent(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance) {
	o.Parent = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetType(v string) {
	o.Type = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// IsSetExternalId returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetRegionCode() string {
	if o == nil || IsNil(o.RegionCode) {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetRegionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RegionCode) {
		return nil, false
	}
	return o.RegionCode, true
}

// IsSetRegionCode returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetRegionCode() bool {
	if o != nil && !IsNil(o.RegionCode) {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetIacId returns the IacId field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetIacId() string {
	if o == nil || IsNil(o.IacId) {
		var ret string
		return ret
	}
	return *o.IacId
}

// GetIacIdOk returns a tuple with the IacId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetIacIdOk() (*string, bool) {
	if o == nil || IsNil(o.IacId) {
		return nil, false
	}
	return o.IacId, true
}

// IsSetIacId returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetIacId() bool {
	if o != nil && !IsNil(o.IacId) {
		return true
	}

	return false
}

// SetIacId gets a reference to the given string and assigns it to the IacId field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetIacId(v string) {
	o.IacId = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetVisibility(v string) {
	o.Visibility = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// IsSetReadOnly returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetDefaultPool returns the DefaultPool field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDefaultPool() bool {
	if o == nil || IsNil(o.DefaultPool) {
		var ret bool
		return ret
	}
	return *o.DefaultPool
}

// GetDefaultPoolOk returns a tuple with the DefaultPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDefaultPoolOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultPool) {
		return nil, false
	}
	return o.DefaultPool, true
}

// IsSetDefaultPool returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetDefaultPool() bool {
	if o != nil && !IsNil(o.DefaultPool) {
		return true
	}

	return false
}

// SetDefaultPool gets a reference to the given bool and assigns it to the DefaultPool field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetDefaultPool(v bool) {
	o.DefaultPool = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// IsSetActive returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetActive(v bool) {
	o.Active = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// IsSetStatus returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetStatus(v string) {
	o.Status = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetInventory() bool {
	if o == nil || IsNil(o.Inventory) {
		var ret bool
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetInventoryOk() (*bool, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// IsSetInventory returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetInventory() bool {
	if o != nil && !IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given bool and assigns it to the Inventory field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetInventory(v bool) {
	o.Inventory = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetConfig() AddCloudResourcePool200ResponseResourcePoolAllOfConfig {
	if o == nil || IsNil(o.Config) {
		var ret AddCloudResourcePool200ResponseResourcePoolAllOfConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetConfigOk() (*AddCloudResourcePool200ResponseResourcePoolAllOfConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given AddCloudResourcePool200ResponseResourcePoolAllOfConfig and assigns it to the Config field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetConfig(v AddCloudResourcePool200ResponseResourcePoolAllOfConfig) {
	o.Config = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// IsSetDisplayName returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Tenants) {
		var ret []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetTenantsOk() ([]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// IsSetTenants returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Tenants field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Tenants = v
}

// GetResourcePermission returns the ResourcePermission field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetResourcePermission() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission {
	if o == nil || IsNil(o.ResourcePermission) {
		var ret ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission
		return ret
	}
	return *o.ResourcePermission
}

// GetResourcePermissionOk returns a tuple with the ResourcePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetResourcePermissionOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool) {
	if o == nil || IsNil(o.ResourcePermission) {
		return nil, false
	}
	return o.ResourcePermission, true
}

// IsSetResourcePermission returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetResourcePermission() bool {
	if o != nil && !IsNil(o.ResourcePermission) {
		return true
	}

	return false
}

// SetResourcePermission gets a reference to the given ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission and assigns it to the ResourcePermission field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetResourcePermission(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission) {
	o.ResourcePermission = &v
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDepth() int64 {
	if o == nil || IsNil(o.Depth) {
		var ret int64
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDepthOk() (*int64, bool) {
	if o == nil || IsNil(o.Depth) {
		return nil, false
	}
	return o.Depth, true
}

// IsSetDepth returns a boolean if a field has been set.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) IsSetDepth() bool {
	if o != nil && !IsNil(o.Depth) {
		return true
	}

	return false
}

// SetDepth gets a reference to the given int64 and assigns it to the Depth field.
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetDepth(v int64) {
	o.Depth = &v
}

func (o GetCloudResourcePools200ResponseAllOfResourcePool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCloudResourcePools200ResponseAllOfResourcePool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.RegionCode) {
		toSerialize["regionCode"] = o.RegionCode
	}
	if !IsNil(o.IacId) {
		toSerialize["iacId"] = o.IacId
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.DefaultPool) {
		toSerialize["defaultPool"] = o.DefaultPool
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Inventory) {
		toSerialize["inventory"] = o.Inventory
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}
	if !IsNil(o.ResourcePermission) {
		toSerialize["resourcePermission"] = o.ResourcePermission
	}
	if !IsNil(o.Depth) {
		toSerialize["depth"] = o.Depth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetCloudResourcePools200ResponseAllOfResourcePool) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

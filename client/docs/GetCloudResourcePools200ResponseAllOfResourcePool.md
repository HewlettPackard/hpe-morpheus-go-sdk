# GetCloudResourcePools200ResponseAllOfResourcePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Parent** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**IacId** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**DefaultPool** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Inventory** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**AddCloudResourcePool200ResponseResourcePoolAllOfConfig**](AddCloudResourcePool200ResponseResourcePoolAllOfConfig.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Tenants** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission**](ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission.md) |  | [optional] 
**Depth** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetCloudResourcePools200ResponseAllOfResourcePool

`func NewGetCloudResourcePools200ResponseAllOfResourcePool() *GetCloudResourcePools200ResponseAllOfResourcePool`

NewGetCloudResourcePools200ResponseAllOfResourcePool instantiates a new GetCloudResourcePools200ResponseAllOfResourcePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCloudResourcePools200ResponseAllOfResourcePoolWithDefaults

`func NewGetCloudResourcePools200ResponseAllOfResourcePoolWithDefaults() *GetCloudResourcePools200ResponseAllOfResourcePool`

NewGetCloudResourcePools200ResponseAllOfResourcePoolWithDefaults instantiates a new GetCloudResourcePools200ResponseAllOfResourcePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZone

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetParent

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetParent() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetParentOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetParent(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRegionCode

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetIacId

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetIacId() string`

GetIacId returns the IacId field if non-nil, zero value otherwise.

### GetIacIdOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetIacIdOk() (*string, bool)`

GetIacIdOk returns a tuple with the IacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacId

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetIacId(v string)`

SetIacId sets IacId field to given value.

### HasIacId

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasIacId() bool`

HasIacId returns a boolean if a field has been set.

### GetVisibility

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetReadOnly

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetDefaultPool

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDefaultPool() bool`

GetDefaultPool returns the DefaultPool field if non-nil, zero value otherwise.

### GetDefaultPoolOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDefaultPoolOk() (*bool, bool)`

GetDefaultPoolOk returns a tuple with the DefaultPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPool

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetDefaultPool(v bool)`

SetDefaultPool sets DefaultPool field to given value.

### HasDefaultPool

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasDefaultPool() bool`

HasDefaultPool returns a boolean if a field has been set.

### GetActive

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStatus

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInventory

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetInventory() bool`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetInventoryOk() (*bool, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetInventory(v bool)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetConfig

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetConfig() AddCloudResourcePool200ResponseResourcePoolAllOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetConfigOk() (*AddCloudResourcePool200ResponseResourcePoolAllOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetConfig(v AddCloudResourcePool200ResponseResourcePoolAllOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetName

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTenants

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetResourcePermission() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetResourcePermissionOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetResourcePermission(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetDepth

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) SetDepth(v int64)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *GetCloudResourcePools200ResponseAllOfResourcePool) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



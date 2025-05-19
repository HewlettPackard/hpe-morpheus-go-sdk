# ListCloudResourcePools200ResponseAllOfResourcePoolsInner

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
**Config** | Pointer to [**ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig**](ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Tenants** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission**](ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission.md) |  | [optional] 
**Depth** | Pointer to **int64** |  | [optional] 

## Methods

### NewListCloudResourcePools200ResponseAllOfResourcePoolsInner

`func NewListCloudResourcePools200ResponseAllOfResourcePoolsInner() *ListCloudResourcePools200ResponseAllOfResourcePoolsInner`

NewListCloudResourcePools200ResponseAllOfResourcePoolsInner instantiates a new ListCloudResourcePools200ResponseAllOfResourcePoolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerWithDefaults

`func NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerWithDefaults() *ListCloudResourcePools200ResponseAllOfResourcePoolsInner`

NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerWithDefaults instantiates a new ListCloudResourcePools200ResponseAllOfResourcePoolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZone

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetParent

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetParent() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetParentOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetParent(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRegionCode

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetIacId

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetIacId() string`

GetIacId returns the IacId field if non-nil, zero value otherwise.

### GetIacIdOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetIacIdOk() (*string, bool)`

GetIacIdOk returns a tuple with the IacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacId

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetIacId(v string)`

SetIacId sets IacId field to given value.

### HasIacId

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasIacId() bool`

HasIacId returns a boolean if a field has been set.

### GetVisibility

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetReadOnly

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetDefaultPool

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetDefaultPool() bool`

GetDefaultPool returns the DefaultPool field if non-nil, zero value otherwise.

### GetDefaultPoolOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetDefaultPoolOk() (*bool, bool)`

GetDefaultPoolOk returns a tuple with the DefaultPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPool

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetDefaultPool(v bool)`

SetDefaultPool sets DefaultPool field to given value.

### HasDefaultPool

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasDefaultPool() bool`

HasDefaultPool returns a boolean if a field has been set.

### GetActive

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStatus

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInventory

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetInventory() bool`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetInventoryOk() (*bool, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetInventory(v bool)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetConfig

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetConfig() ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetConfigOk() (*ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetConfig(v ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetName

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTenants

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetResourcePermission() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetResourcePermissionOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetResourcePermission(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetDepth

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) SetDepth(v int64)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInner) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



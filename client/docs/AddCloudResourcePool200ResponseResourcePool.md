# AddCloudResourcePool200ResponseResourcePool

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
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddCloudResourcePool200ResponseResourcePool

`func NewAddCloudResourcePool200ResponseResourcePool() *AddCloudResourcePool200ResponseResourcePool`

NewAddCloudResourcePool200ResponseResourcePool instantiates a new AddCloudResourcePool200ResponseResourcePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCloudResourcePool200ResponseResourcePoolWithDefaults

`func NewAddCloudResourcePool200ResponseResourcePoolWithDefaults() *AddCloudResourcePool200ResponseResourcePool`

NewAddCloudResourcePool200ResponseResourcePoolWithDefaults instantiates a new AddCloudResourcePool200ResponseResourcePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddCloudResourcePool200ResponseResourcePool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCloudResourcePool200ResponseResourcePool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddCloudResourcePool200ResponseResourcePool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *AddCloudResourcePool200ResponseResourcePool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCloudResourcePool200ResponseResourcePool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCloudResourcePool200ResponseResourcePool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZone

`func (o *AddCloudResourcePool200ResponseResourcePool) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddCloudResourcePool200ResponseResourcePool) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddCloudResourcePool200ResponseResourcePool) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetParent

`func (o *AddCloudResourcePool200ResponseResourcePool) GetParent() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetParentOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AddCloudResourcePool200ResponseResourcePool) SetParent(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AddCloudResourcePool200ResponseResourcePool) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *AddCloudResourcePool200ResponseResourcePool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCloudResourcePool200ResponseResourcePool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddCloudResourcePool200ResponseResourcePool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *AddCloudResourcePool200ResponseResourcePool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddCloudResourcePool200ResponseResourcePool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddCloudResourcePool200ResponseResourcePool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRegionCode

`func (o *AddCloudResourcePool200ResponseResourcePool) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *AddCloudResourcePool200ResponseResourcePool) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *AddCloudResourcePool200ResponseResourcePool) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetIacId

`func (o *AddCloudResourcePool200ResponseResourcePool) GetIacId() string`

GetIacId returns the IacId field if non-nil, zero value otherwise.

### GetIacIdOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetIacIdOk() (*string, bool)`

GetIacIdOk returns a tuple with the IacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacId

`func (o *AddCloudResourcePool200ResponseResourcePool) SetIacId(v string)`

SetIacId sets IacId field to given value.

### HasIacId

`func (o *AddCloudResourcePool200ResponseResourcePool) HasIacId() bool`

HasIacId returns a boolean if a field has been set.

### GetVisibility

`func (o *AddCloudResourcePool200ResponseResourcePool) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddCloudResourcePool200ResponseResourcePool) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddCloudResourcePool200ResponseResourcePool) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetReadOnly

`func (o *AddCloudResourcePool200ResponseResourcePool) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *AddCloudResourcePool200ResponseResourcePool) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *AddCloudResourcePool200ResponseResourcePool) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetDefaultPool

`func (o *AddCloudResourcePool200ResponseResourcePool) GetDefaultPool() bool`

GetDefaultPool returns the DefaultPool field if non-nil, zero value otherwise.

### GetDefaultPoolOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetDefaultPoolOk() (*bool, bool)`

GetDefaultPoolOk returns a tuple with the DefaultPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPool

`func (o *AddCloudResourcePool200ResponseResourcePool) SetDefaultPool(v bool)`

SetDefaultPool sets DefaultPool field to given value.

### HasDefaultPool

`func (o *AddCloudResourcePool200ResponseResourcePool) HasDefaultPool() bool`

HasDefaultPool returns a boolean if a field has been set.

### GetActive

`func (o *AddCloudResourcePool200ResponseResourcePool) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddCloudResourcePool200ResponseResourcePool) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddCloudResourcePool200ResponseResourcePool) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStatus

`func (o *AddCloudResourcePool200ResponseResourcePool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddCloudResourcePool200ResponseResourcePool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddCloudResourcePool200ResponseResourcePool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInventory

`func (o *AddCloudResourcePool200ResponseResourcePool) GetInventory() bool`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetInventoryOk() (*bool, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *AddCloudResourcePool200ResponseResourcePool) SetInventory(v bool)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *AddCloudResourcePool200ResponseResourcePool) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetConfig

`func (o *AddCloudResourcePool200ResponseResourcePool) GetConfig() AddCloudResourcePool200ResponseResourcePoolAllOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetConfigOk() (*AddCloudResourcePool200ResponseResourcePoolAllOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddCloudResourcePool200ResponseResourcePool) SetConfig(v AddCloudResourcePool200ResponseResourcePoolAllOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddCloudResourcePool200ResponseResourcePool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetName

`func (o *AddCloudResourcePool200ResponseResourcePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCloudResourcePool200ResponseResourcePool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddCloudResourcePool200ResponseResourcePool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *AddCloudResourcePool200ResponseResourcePool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddCloudResourcePool200ResponseResourcePool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AddCloudResourcePool200ResponseResourcePool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTenants

`func (o *AddCloudResourcePool200ResponseResourcePool) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *AddCloudResourcePool200ResponseResourcePool) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *AddCloudResourcePool200ResponseResourcePool) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *AddCloudResourcePool200ResponseResourcePool) GetResourcePermission() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetResourcePermissionOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *AddCloudResourcePool200ResponseResourcePool) SetResourcePermission(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *AddCloudResourcePool200ResponseResourcePool) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetDepth

`func (o *AddCloudResourcePool200ResponseResourcePool) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *AddCloudResourcePool200ResponseResourcePool) SetDepth(v int64)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *AddCloudResourcePool200ResponseResourcePool) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetSuccess

`func (o *AddCloudResourcePool200ResponseResourcePool) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddCloudResourcePool200ResponseResourcePool) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddCloudResourcePool200ResponseResourcePool) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddCloudResourcePool200ResponseResourcePool) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



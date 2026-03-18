# UpdateCloudResourcePool200ResponseResourcePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to [**NullableUpdateCloudResourcePool200ResponseResourcePoolAllOfZone**](UpdateCloudResourcePool200ResponseResourcePoolAllOfZone.md) |  | [optional] 
**Parent** | Pointer to [**UpdateCloudResourcePool200ResponseResourcePoolAllOfParent**](UpdateCloudResourcePool200ResponseResourcePoolAllOfParent.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**RegionCode** | Pointer to **NullableString** |  | [optional] 
**IacId** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**DefaultPool** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Inventory** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig**](UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Tenants** | Pointer to [**[]UpdateCloudResourcePool200ResponseResourcePoolAllOfTenantsInner**](UpdateCloudResourcePool200ResponseResourcePoolAllOfTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission**](UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission.md) |  | [optional] 
**Depth** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateCloudResourcePool200ResponseResourcePool

`func NewUpdateCloudResourcePool200ResponseResourcePool() *UpdateCloudResourcePool200ResponseResourcePool`

NewUpdateCloudResourcePool200ResponseResourcePool instantiates a new UpdateCloudResourcePool200ResponseResourcePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCloudResourcePool200ResponseResourcePoolWithDefaults

`func NewUpdateCloudResourcePool200ResponseResourcePoolWithDefaults() *UpdateCloudResourcePool200ResponseResourcePool`

NewUpdateCloudResourcePool200ResponseResourcePoolWithDefaults instantiates a new UpdateCloudResourcePool200ResponseResourcePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateCloudResourcePool200ResponseResourcePool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetZone

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetZone() UpdateCloudResourcePool200ResponseResourcePoolAllOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetZoneOk() (*UpdateCloudResourcePool200ResponseResourcePoolAllOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetZone(v UpdateCloudResourcePool200ResponseResourcePoolAllOfZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *UpdateCloudResourcePool200ResponseResourcePool) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetParent

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetParent() UpdateCloudResourcePool200ResponseResourcePoolAllOfParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetParentOk() (*UpdateCloudResourcePool200ResponseResourcePoolAllOfParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetParent(v UpdateCloudResourcePool200ResponseResourcePoolAllOfParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRegionCode

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *UpdateCloudResourcePool200ResponseResourcePool) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetIacId

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetIacId() string`

GetIacId returns the IacId field if non-nil, zero value otherwise.

### GetIacIdOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetIacIdOk() (*string, bool)`

GetIacIdOk returns a tuple with the IacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacId

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetIacId(v string)`

SetIacId sets IacId field to given value.

### HasIacId

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasIacId() bool`

HasIacId returns a boolean if a field has been set.

### SetIacIdNil

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetIacIdNil(b bool)`

 SetIacIdNil sets the value for IacId to be an explicit nil

### UnsetIacId
`func (o *UpdateCloudResourcePool200ResponseResourcePool) UnsetIacId()`

UnsetIacId ensures that no value is present for IacId, not even an explicit nil
### GetVisibility

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetReadOnly

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetDefaultPool

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetDefaultPool() bool`

GetDefaultPool returns the DefaultPool field if non-nil, zero value otherwise.

### GetDefaultPoolOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetDefaultPoolOk() (*bool, bool)`

GetDefaultPoolOk returns a tuple with the DefaultPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPool

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetDefaultPool(v bool)`

SetDefaultPool sets DefaultPool field to given value.

### HasDefaultPool

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasDefaultPool() bool`

HasDefaultPool returns a boolean if a field has been set.

### GetActive

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInventory

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetInventory() bool`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetInventoryOk() (*bool, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetInventory(v bool)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetConfig() UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetConfigOk() (*UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetConfig(v UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetName

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UpdateCloudResourcePool200ResponseResourcePool) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetTenants

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetTenants() []UpdateCloudResourcePool200ResponseResourcePoolAllOfTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetTenantsOk() (*[]UpdateCloudResourcePool200ResponseResourcePoolAllOfTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetTenants(v []UpdateCloudResourcePool200ResponseResourcePoolAllOfTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetResourcePermission() UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetResourcePermissionOk() (*UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetResourcePermission(v UpdateCloudResourcePool200ResponseResourcePoolAllOfResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetDepth

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetDepth(v int64)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateCloudResourcePool200ResponseResourcePool) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateCloudResourcePool200ResponseResourcePool) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateCloudResourcePool200ResponseResourcePool) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



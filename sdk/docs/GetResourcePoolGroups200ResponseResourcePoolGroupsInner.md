# GetResourcePoolGroups200ResponseResourcePoolGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** | Pool selection mode. Valid values are &#x60;roundrobin&#x60; or &#x60;availablecapacity&#x60;. | [optional] 
**Pools** | Pointer to **[]int64** | Array of Resource Pool IDs | [optional] 
**Tenants** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission**](GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission.md) |  | [optional] 

## Methods

### NewGetResourcePoolGroups200ResponseResourcePoolGroupsInner

`func NewGetResourcePoolGroups200ResponseResourcePoolGroupsInner() *GetResourcePoolGroups200ResponseResourcePoolGroupsInner`

NewGetResourcePoolGroups200ResponseResourcePoolGroupsInner instantiates a new GetResourcePoolGroups200ResponseResourcePoolGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResourcePoolGroups200ResponseResourcePoolGroupsInnerWithDefaults

`func NewGetResourcePoolGroups200ResponseResourcePoolGroupsInnerWithDefaults() *GetResourcePoolGroups200ResponseResourcePoolGroupsInner`

NewGetResourcePoolGroups200ResponseResourcePoolGroupsInnerWithDefaults instantiates a new GetResourcePoolGroups200ResponseResourcePoolGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMode

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPools

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetPools() []int64`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetPoolsOk() (*[]int64, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetPools(v []int64)`

SetPools sets Pools field to given value.

### HasPools

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetTenants

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetResourcePermission() GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) GetResourcePermissionOk() (*GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) SetResourcePermission(v GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInner) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



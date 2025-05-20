# CreateResourcePoolGroupRequestResourcePoolGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** | Pool selection mode. Valid values are &#x60;roundrobin&#x60; or &#x60;availablecapacity&#x60;. | [optional] 
**Pools** | Pointer to **[]int64** |  | [optional] 
**Tenants** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission**](GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission.md) |  | [optional] 

## Methods

### NewCreateResourcePoolGroupRequestResourcePoolGroup

`func NewCreateResourcePoolGroupRequestResourcePoolGroup() *CreateResourcePoolGroupRequestResourcePoolGroup`

NewCreateResourcePoolGroupRequestResourcePoolGroup instantiates a new CreateResourcePoolGroupRequestResourcePoolGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourcePoolGroupRequestResourcePoolGroupWithDefaults

`func NewCreateResourcePoolGroupRequestResourcePoolGroupWithDefaults() *CreateResourcePoolGroupRequestResourcePoolGroup`

NewCreateResourcePoolGroupRequestResourcePoolGroupWithDefaults instantiates a new CreateResourcePoolGroupRequestResourcePoolGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMode

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPools

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetPools() []int64`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetPoolsOk() (*[]int64, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) SetPools(v []int64)`

SetPools sets Pools field to given value.

### HasPools

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetTenants

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetResourcePermission() GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) GetResourcePermissionOk() (*GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) SetResourcePermission(v GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *CreateResourcePoolGroupRequestResourcePoolGroup) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



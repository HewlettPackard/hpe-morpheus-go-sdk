# UpdateResourcePoolGroupRequestResourcePoolGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** | Pool selection mode. Valid values are &#x60;roundrobin&#x60; or &#x60;availablecapacity&#x60;. | [optional] 
**Pools** | Pointer to **[]int64** |  | [optional] 
**Tenants** | Pointer to [**[]UpdateResourcePoolGroupRequestResourcePoolGroupTenantsInner**](UpdateResourcePoolGroupRequestResourcePoolGroupTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission**](UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission.md) |  | [optional] 

## Methods

### NewUpdateResourcePoolGroupRequestResourcePoolGroup

`func NewUpdateResourcePoolGroupRequestResourcePoolGroup() *UpdateResourcePoolGroupRequestResourcePoolGroup`

NewUpdateResourcePoolGroupRequestResourcePoolGroup instantiates a new UpdateResourcePoolGroupRequestResourcePoolGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMode

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPools

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetPools() []int64`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetPoolsOk() (*[]int64, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) SetPools(v []int64)`

SetPools sets Pools field to given value.

### HasPools

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetTenants() []UpdateResourcePoolGroupRequestResourcePoolGroupTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetTenantsOk() (*[]UpdateResourcePoolGroupRequestResourcePoolGroupTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) SetTenants(v []UpdateResourcePoolGroupRequestResourcePoolGroupTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetResourcePermission() UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) GetResourcePermissionOk() (*UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) SetResourcePermission(v UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroup) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



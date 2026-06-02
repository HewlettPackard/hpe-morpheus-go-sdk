# CreateResourcePoolGroup200ResponseResourcePoolGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** | Pool selection mode. Valid values are &#x60;roundrobin&#x60; or &#x60;availablecapacity&#x60;. | [optional] 
**Pools** | Pointer to **[]int64** | Array of Resource Pool IDs | [optional] 
**Tenants** | Pointer to [**[]CreateResourcePoolGroup200ResponseResourcePoolGroupTenantsInner**](CreateResourcePoolGroup200ResponseResourcePoolGroupTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission**](CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission.md) |  | [optional] 

## Methods

### NewCreateResourcePoolGroup200ResponseResourcePoolGroup

`func NewCreateResourcePoolGroup200ResponseResourcePoolGroup() *CreateResourcePoolGroup200ResponseResourcePoolGroup`

NewCreateResourcePoolGroup200ResponseResourcePoolGroup instantiates a new CreateResourcePoolGroup200ResponseResourcePoolGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMode

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPools

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetPools() []int64`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetPoolsOk() (*[]int64, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) SetPools(v []int64)`

SetPools sets Pools field to given value.

### HasPools

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetTenants

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetTenants() []CreateResourcePoolGroup200ResponseResourcePoolGroupTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetTenantsOk() (*[]CreateResourcePoolGroup200ResponseResourcePoolGroupTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) SetTenants(v []CreateResourcePoolGroup200ResponseResourcePoolGroupTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetResourcePermission() CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) GetResourcePermissionOk() (*CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) SetResourcePermission(v CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroup) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



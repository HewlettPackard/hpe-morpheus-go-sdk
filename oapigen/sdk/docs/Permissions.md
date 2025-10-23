# Permissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePool** | Pointer to [**AddCluster200ResponseAllOfClusterPermissionsResourcePool**](AddCluster200ResponseAllOfClusterPermissionsResourcePool.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission**](ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission.md) |  | [optional] 
**TenantPermissions** | Pointer to [**AddCloudResourcePoolRequestResourcePoolTenantPermissions**](AddCloudResourcePoolRequestResourcePoolTenantPermissions.md) |  | [optional] 

## Methods

### NewPermissions

`func NewPermissions() *Permissions`

NewPermissions instantiates a new Permissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsWithDefaults

`func NewPermissionsWithDefaults() *Permissions`

NewPermissionsWithDefaults instantiates a new Permissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePool

`func (o *Permissions) GetResourcePool() AddCluster200ResponseAllOfClusterPermissionsResourcePool`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *Permissions) GetResourcePoolOk() (*AddCluster200ResponseAllOfClusterPermissionsResourcePool, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *Permissions) SetResourcePool(v AddCluster200ResponseAllOfClusterPermissionsResourcePool)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *Permissions) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *Permissions) GetResourcePermissions() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *Permissions) GetResourcePermissionsOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *Permissions) SetResourcePermissions(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *Permissions) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *Permissions) GetTenantPermissions() AddCloudResourcePoolRequestResourcePoolTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *Permissions) GetTenantPermissionsOk() (*AddCloudResourcePoolRequestResourcePoolTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *Permissions) SetTenantPermissions(v AddCloudResourcePoolRequestResourcePoolTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *Permissions) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



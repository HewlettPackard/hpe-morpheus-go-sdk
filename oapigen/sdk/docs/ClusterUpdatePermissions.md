# ClusterUpdatePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePermissions** | Pointer to [**ClusterUpdatePermissionsResourcePermissions**](ClusterUpdatePermissionsResourcePermissions.md) |  | [optional] 
**ResourcePool** | Pointer to [**ClusterUpdatePermissionsResourcePool**](ClusterUpdatePermissionsResourcePool.md) |  | [optional] 
**TenantPermissions** | Pointer to [**ClusterUpdatePermissionsTenantPermissions**](ClusterUpdatePermissionsTenantPermissions.md) |  | [optional] 

## Methods

### NewClusterUpdatePermissions

`func NewClusterUpdatePermissions() *ClusterUpdatePermissions`

NewClusterUpdatePermissions instantiates a new ClusterUpdatePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterUpdatePermissionsWithDefaults

`func NewClusterUpdatePermissionsWithDefaults() *ClusterUpdatePermissions`

NewClusterUpdatePermissionsWithDefaults instantiates a new ClusterUpdatePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePermissions

`func (o *ClusterUpdatePermissions) GetResourcePermissions() ClusterUpdatePermissionsResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ClusterUpdatePermissions) GetResourcePermissionsOk() (*ClusterUpdatePermissionsResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ClusterUpdatePermissions) SetResourcePermissions(v ClusterUpdatePermissionsResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ClusterUpdatePermissions) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetResourcePool

`func (o *ClusterUpdatePermissions) GetResourcePool() ClusterUpdatePermissionsResourcePool`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ClusterUpdatePermissions) GetResourcePoolOk() (*ClusterUpdatePermissionsResourcePool, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ClusterUpdatePermissions) SetResourcePool(v ClusterUpdatePermissionsResourcePool)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ClusterUpdatePermissions) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *ClusterUpdatePermissions) GetTenantPermissions() ClusterUpdatePermissionsTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *ClusterUpdatePermissions) GetTenantPermissionsOk() (*ClusterUpdatePermissionsTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *ClusterUpdatePermissions) SetTenantPermissions(v ClusterUpdatePermissionsTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *ClusterUpdatePermissions) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



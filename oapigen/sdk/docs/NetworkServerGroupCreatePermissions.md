# NetworkServerGroupCreatePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePool** | Pointer to [**NetworkServerGroupCreatePermissionsResourcePool**](NetworkServerGroupCreatePermissionsResourcePool.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**NetworkServerGroupCreatePermissionsResourcePermissions**](NetworkServerGroupCreatePermissionsResourcePermissions.md) |  | [optional] 
**TenantPermissions** | Pointer to [**NetworkServerGroupCreatePermissionsTenantPermissions**](NetworkServerGroupCreatePermissionsTenantPermissions.md) |  | [optional] 

## Methods

### NewNetworkServerGroupCreatePermissions

`func NewNetworkServerGroupCreatePermissions() *NetworkServerGroupCreatePermissions`

NewNetworkServerGroupCreatePermissions instantiates a new NetworkServerGroupCreatePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkServerGroupCreatePermissionsWithDefaults

`func NewNetworkServerGroupCreatePermissionsWithDefaults() *NetworkServerGroupCreatePermissions`

NewNetworkServerGroupCreatePermissionsWithDefaults instantiates a new NetworkServerGroupCreatePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePool

`func (o *NetworkServerGroupCreatePermissions) GetResourcePool() NetworkServerGroupCreatePermissionsResourcePool`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *NetworkServerGroupCreatePermissions) GetResourcePoolOk() (*NetworkServerGroupCreatePermissionsResourcePool, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *NetworkServerGroupCreatePermissions) SetResourcePool(v NetworkServerGroupCreatePermissionsResourcePool)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *NetworkServerGroupCreatePermissions) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *NetworkServerGroupCreatePermissions) GetResourcePermissions() NetworkServerGroupCreatePermissionsResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *NetworkServerGroupCreatePermissions) GetResourcePermissionsOk() (*NetworkServerGroupCreatePermissionsResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *NetworkServerGroupCreatePermissions) SetResourcePermissions(v NetworkServerGroupCreatePermissionsResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *NetworkServerGroupCreatePermissions) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *NetworkServerGroupCreatePermissions) GetTenantPermissions() NetworkServerGroupCreatePermissionsTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *NetworkServerGroupCreatePermissions) GetTenantPermissionsOk() (*NetworkServerGroupCreatePermissionsTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *NetworkServerGroupCreatePermissions) SetTenantPermissions(v NetworkServerGroupCreatePermissionsTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *NetworkServerGroupCreatePermissions) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ClusterPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePool** | Pointer to [**ClusterPermissionsResourcePool**](ClusterPermissionsResourcePool.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**ClusterPermissionsResourcePermissions**](ClusterPermissionsResourcePermissions.md) |  | [optional] 

## Methods

### NewClusterPermissions

`func NewClusterPermissions() *ClusterPermissions`

NewClusterPermissions instantiates a new ClusterPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPermissionsWithDefaults

`func NewClusterPermissionsWithDefaults() *ClusterPermissions`

NewClusterPermissionsWithDefaults instantiates a new ClusterPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePool

`func (o *ClusterPermissions) GetResourcePool() ClusterPermissionsResourcePool`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ClusterPermissions) GetResourcePoolOk() (*ClusterPermissionsResourcePool, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ClusterPermissions) SetResourcePool(v ClusterPermissionsResourcePool)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ClusterPermissions) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ClusterPermissions) GetResourcePermissions() ClusterPermissionsResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ClusterPermissions) GetResourcePermissionsOk() (*ClusterPermissionsResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ClusterPermissions) SetResourcePermissions(v ClusterPermissionsResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ClusterPermissions) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



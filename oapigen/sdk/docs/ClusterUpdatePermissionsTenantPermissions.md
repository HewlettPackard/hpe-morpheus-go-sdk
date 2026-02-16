# ClusterUpdatePermissionsTenantPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **[]int64** | Array of tenant account ids that are allowed access | [optional] 
**DefaultTarget** | Pointer to **[]int64** | Array of tenant account ids which should use the data store as the Default | [optional] 
**DefaultStore** | Pointer to **[]int64** | Array of tenant account ids which should use the data store as the Image Target | [optional] 

## Methods

### NewClusterUpdatePermissionsTenantPermissions

`func NewClusterUpdatePermissionsTenantPermissions() *ClusterUpdatePermissionsTenantPermissions`

NewClusterUpdatePermissionsTenantPermissions instantiates a new ClusterUpdatePermissionsTenantPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterUpdatePermissionsTenantPermissionsWithDefaults

`func NewClusterUpdatePermissionsTenantPermissionsWithDefaults() *ClusterUpdatePermissionsTenantPermissions`

NewClusterUpdatePermissionsTenantPermissionsWithDefaults instantiates a new ClusterUpdatePermissionsTenantPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ClusterUpdatePermissionsTenantPermissions) GetAccounts() []int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ClusterUpdatePermissionsTenantPermissions) GetAccountsOk() (*[]int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ClusterUpdatePermissionsTenantPermissions) SetAccounts(v []int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ClusterUpdatePermissionsTenantPermissions) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *ClusterUpdatePermissionsTenantPermissions) GetDefaultTarget() []int64`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *ClusterUpdatePermissionsTenantPermissions) GetDefaultTargetOk() (*[]int64, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *ClusterUpdatePermissionsTenantPermissions) SetDefaultTarget(v []int64)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *ClusterUpdatePermissionsTenantPermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetDefaultStore

`func (o *ClusterUpdatePermissionsTenantPermissions) GetDefaultStore() []int64`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ClusterUpdatePermissionsTenantPermissions) GetDefaultStoreOk() (*[]int64, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ClusterUpdatePermissionsTenantPermissions) SetDefaultStore(v []int64)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ClusterUpdatePermissionsTenantPermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



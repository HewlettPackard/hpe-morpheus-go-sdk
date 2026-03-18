# RolePermissionClusterType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterTypeId** | **int32** | &#x60;id&#x60; of the cluster type | 
**Access** | **string** | The new access level. | 

## Methods

### NewRolePermissionClusterType

`func NewRolePermissionClusterType(clusterTypeId int32, access string, ) *RolePermissionClusterType`

NewRolePermissionClusterType instantiates a new RolePermissionClusterType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionClusterTypeWithDefaults

`func NewRolePermissionClusterTypeWithDefaults() *RolePermissionClusterType`

NewRolePermissionClusterTypeWithDefaults instantiates a new RolePermissionClusterType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterTypeId

`func (o *RolePermissionClusterType) GetClusterTypeId() int32`

GetClusterTypeId returns the ClusterTypeId field if non-nil, zero value otherwise.

### GetClusterTypeIdOk

`func (o *RolePermissionClusterType) GetClusterTypeIdOk() (*int32, bool)`

GetClusterTypeIdOk returns a tuple with the ClusterTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTypeId

`func (o *RolePermissionClusterType) SetClusterTypeId(v int32)`

SetClusterTypeId sets ClusterTypeId field to given value.


### GetAccess

`func (o *RolePermissionClusterType) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *RolePermissionClusterType) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *RolePermissionClusterType) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RolePermissionGroupAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllGroups** | **bool** | Apply to all groups (site) | 
**Access** | **string** | The new access level. | 

## Methods

### NewRolePermissionGroupAll

`func NewRolePermissionGroupAll(allGroups bool, access string, ) *RolePermissionGroupAll`

NewRolePermissionGroupAll instantiates a new RolePermissionGroupAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionGroupAllWithDefaults

`func NewRolePermissionGroupAllWithDefaults() *RolePermissionGroupAll`

NewRolePermissionGroupAllWithDefaults instantiates a new RolePermissionGroupAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllGroups

`func (o *RolePermissionGroupAll) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *RolePermissionGroupAll) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *RolePermissionGroupAll) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.


### GetAccess

`func (o *RolePermissionGroupAll) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *RolePermissionGroupAll) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *RolePermissionGroupAll) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



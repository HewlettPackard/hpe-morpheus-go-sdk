# RolePermissionTaskSetAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllTaskSets** | **bool** | Apply to all workflows (taskSets) | 
**Access** | **string** | The new access level. | 

## Methods

### NewRolePermissionTaskSetAll

`func NewRolePermissionTaskSetAll(allTaskSets bool, access string, ) *RolePermissionTaskSetAll`

NewRolePermissionTaskSetAll instantiates a new RolePermissionTaskSetAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionTaskSetAllWithDefaults

`func NewRolePermissionTaskSetAllWithDefaults() *RolePermissionTaskSetAll`

NewRolePermissionTaskSetAllWithDefaults instantiates a new RolePermissionTaskSetAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllTaskSets

`func (o *RolePermissionTaskSetAll) GetAllTaskSets() bool`

GetAllTaskSets returns the AllTaskSets field if non-nil, zero value otherwise.

### GetAllTaskSetsOk

`func (o *RolePermissionTaskSetAll) GetAllTaskSetsOk() (*bool, bool)`

GetAllTaskSetsOk returns a tuple with the AllTaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTaskSets

`func (o *RolePermissionTaskSetAll) SetAllTaskSets(v bool)`

SetAllTaskSets sets AllTaskSets field to given value.


### GetAccess

`func (o *RolePermissionTaskSetAll) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *RolePermissionTaskSetAll) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *RolePermissionTaskSetAll) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



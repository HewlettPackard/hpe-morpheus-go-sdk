# RolePermissionTaskSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskSetId** | **int32** | &#x60;id&#x60; of the workflow (taskSet) | 
**Access** | **string** | The new access level. | 

## Methods

### NewRolePermissionTaskSet

`func NewRolePermissionTaskSet(taskSetId int32, access string, ) *RolePermissionTaskSet`

NewRolePermissionTaskSet instantiates a new RolePermissionTaskSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionTaskSetWithDefaults

`func NewRolePermissionTaskSetWithDefaults() *RolePermissionTaskSet`

NewRolePermissionTaskSetWithDefaults instantiates a new RolePermissionTaskSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskSetId

`func (o *RolePermissionTaskSet) GetTaskSetId() int32`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *RolePermissionTaskSet) GetTaskSetIdOk() (*int32, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *RolePermissionTaskSet) SetTaskSetId(v int32)`

SetTaskSetId sets TaskSetId field to given value.


### GetAccess

`func (o *RolePermissionTaskSet) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *RolePermissionTaskSet) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *RolePermissionTaskSet) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



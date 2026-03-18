# UpdateRoleTaskAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **int32** | &#x60;id&#x60; of the task | 
**Access** | **string** | The new access level. | 
**AllTasks** | **bool** | Apply to all tasks | 

## Methods

### NewUpdateRoleTaskAccessRequest

`func NewUpdateRoleTaskAccessRequest(taskId int32, access string, allTasks bool, ) *UpdateRoleTaskAccessRequest`

NewUpdateRoleTaskAccessRequest instantiates a new UpdateRoleTaskAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleTaskAccessRequestWithDefaults

`func NewUpdateRoleTaskAccessRequestWithDefaults() *UpdateRoleTaskAccessRequest`

NewUpdateRoleTaskAccessRequestWithDefaults instantiates a new UpdateRoleTaskAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *UpdateRoleTaskAccessRequest) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *UpdateRoleTaskAccessRequest) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *UpdateRoleTaskAccessRequest) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.


### GetAccess

`func (o *UpdateRoleTaskAccessRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleTaskAccessRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleTaskAccessRequest) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAllTasks

`func (o *UpdateRoleTaskAccessRequest) GetAllTasks() bool`

GetAllTasks returns the AllTasks field if non-nil, zero value otherwise.

### GetAllTasksOk

`func (o *UpdateRoleTaskAccessRequest) GetAllTasksOk() (*bool, bool)`

GetAllTasksOk returns a tuple with the AllTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTasks

`func (o *UpdateRoleTaskAccessRequest) SetAllTasks(v bool)`

SetAllTasks sets AllTasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateRoleWorkflowAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskSetId** | **int32** | &#x60;id&#x60; of the workflow (taskSet) | 
**Access** | **string** | The new access level. | 
**AllTaskSets** | **bool** | Apply to all workflows (taskSets) | 

## Methods

### NewUpdateRoleWorkflowAccessRequest

`func NewUpdateRoleWorkflowAccessRequest(taskSetId int32, access string, allTaskSets bool, ) *UpdateRoleWorkflowAccessRequest`

NewUpdateRoleWorkflowAccessRequest instantiates a new UpdateRoleWorkflowAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleWorkflowAccessRequestWithDefaults

`func NewUpdateRoleWorkflowAccessRequestWithDefaults() *UpdateRoleWorkflowAccessRequest`

NewUpdateRoleWorkflowAccessRequestWithDefaults instantiates a new UpdateRoleWorkflowAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskSetId

`func (o *UpdateRoleWorkflowAccessRequest) GetTaskSetId() int32`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *UpdateRoleWorkflowAccessRequest) GetTaskSetIdOk() (*int32, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *UpdateRoleWorkflowAccessRequest) SetTaskSetId(v int32)`

SetTaskSetId sets TaskSetId field to given value.


### GetAccess

`func (o *UpdateRoleWorkflowAccessRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleWorkflowAccessRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleWorkflowAccessRequest) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAllTaskSets

`func (o *UpdateRoleWorkflowAccessRequest) GetAllTaskSets() bool`

GetAllTaskSets returns the AllTaskSets field if non-nil, zero value otherwise.

### GetAllTaskSetsOk

`func (o *UpdateRoleWorkflowAccessRequest) GetAllTaskSetsOk() (*bool, bool)`

GetAllTaskSetsOk returns a tuple with the AllTaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTaskSets

`func (o *UpdateRoleWorkflowAccessRequest) SetAllTaskSets(v bool)`

SetAllTaskSets sets AllTaskSets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



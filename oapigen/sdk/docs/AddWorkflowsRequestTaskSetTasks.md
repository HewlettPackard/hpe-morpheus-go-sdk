# AddWorkflowsRequestTaskSetTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **int64** | Task ID | [optional] 
**TaskPhase** | Pointer to **string** | Task Phase. Pass operation for &#x60;operational&#x60; workflows. | [optional] [default to "provision"]

## Methods

### NewAddWorkflowsRequestTaskSetTasks

`func NewAddWorkflowsRequestTaskSetTasks() *AddWorkflowsRequestTaskSetTasks`

NewAddWorkflowsRequestTaskSetTasks instantiates a new AddWorkflowsRequestTaskSetTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddWorkflowsRequestTaskSetTasksWithDefaults

`func NewAddWorkflowsRequestTaskSetTasksWithDefaults() *AddWorkflowsRequestTaskSetTasks`

NewAddWorkflowsRequestTaskSetTasksWithDefaults instantiates a new AddWorkflowsRequestTaskSetTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *AddWorkflowsRequestTaskSetTasks) GetTaskId() int64`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *AddWorkflowsRequestTaskSetTasks) GetTaskIdOk() (*int64, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *AddWorkflowsRequestTaskSetTasks) SetTaskId(v int64)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *AddWorkflowsRequestTaskSetTasks) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTaskPhase

`func (o *AddWorkflowsRequestTaskSetTasks) GetTaskPhase() string`

GetTaskPhase returns the TaskPhase field if non-nil, zero value otherwise.

### GetTaskPhaseOk

`func (o *AddWorkflowsRequestTaskSetTasks) GetTaskPhaseOk() (*string, bool)`

GetTaskPhaseOk returns a tuple with the TaskPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPhase

`func (o *AddWorkflowsRequestTaskSetTasks) SetTaskPhase(v string)`

SetTaskPhase sets TaskPhase field to given value.

### HasTaskPhase

`func (o *AddWorkflowsRequestTaskSetTasks) HasTaskPhase() bool`

HasTaskPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



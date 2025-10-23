# UpdateHostExecuteWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskSet** | Pointer to [**UpdateHostExecuteWorkflowRequestTaskSet**](UpdateHostExecuteWorkflowRequestTaskSet.md) |  | [optional] 
**TaskPhase** | Pointer to **string** | Task Phase to run for Provisioning workflows. The default is &#x60;provision&#x60;. | [optional] [default to "provision"]

## Methods

### NewUpdateHostExecuteWorkflowRequest

`func NewUpdateHostExecuteWorkflowRequest() *UpdateHostExecuteWorkflowRequest`

NewUpdateHostExecuteWorkflowRequest instantiates a new UpdateHostExecuteWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostExecuteWorkflowRequestWithDefaults

`func NewUpdateHostExecuteWorkflowRequestWithDefaults() *UpdateHostExecuteWorkflowRequest`

NewUpdateHostExecuteWorkflowRequestWithDefaults instantiates a new UpdateHostExecuteWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskSet

`func (o *UpdateHostExecuteWorkflowRequest) GetTaskSet() UpdateHostExecuteWorkflowRequestTaskSet`

GetTaskSet returns the TaskSet field if non-nil, zero value otherwise.

### GetTaskSetOk

`func (o *UpdateHostExecuteWorkflowRequest) GetTaskSetOk() (*UpdateHostExecuteWorkflowRequestTaskSet, bool)`

GetTaskSetOk returns a tuple with the TaskSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSet

`func (o *UpdateHostExecuteWorkflowRequest) SetTaskSet(v UpdateHostExecuteWorkflowRequestTaskSet)`

SetTaskSet sets TaskSet field to given value.

### HasTaskSet

`func (o *UpdateHostExecuteWorkflowRequest) HasTaskSet() bool`

HasTaskSet returns a boolean if a field has been set.

### GetTaskPhase

`func (o *UpdateHostExecuteWorkflowRequest) GetTaskPhase() string`

GetTaskPhase returns the TaskPhase field if non-nil, zero value otherwise.

### GetTaskPhaseOk

`func (o *UpdateHostExecuteWorkflowRequest) GetTaskPhaseOk() (*string, bool)`

GetTaskPhaseOk returns a tuple with the TaskPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPhase

`func (o *UpdateHostExecuteWorkflowRequest) SetTaskPhase(v string)`

SetTaskPhase sets TaskPhase field to given value.

### HasTaskPhase

`func (o *UpdateHostExecuteWorkflowRequest) HasTaskPhase() bool`

HasTaskPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InstanceWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskSet** | Pointer to [**RunWorkflowInstanceRequestTaskSet**](RunWorkflowInstanceRequestTaskSet.md) |  | [optional] 
**TaskPhase** | Pointer to **string** | Task Phase to run for Provisioning workflows. The default is &#x60;provision&#x60;. | [optional] [default to "provision"]

## Methods

### NewInstanceWorkflow

`func NewInstanceWorkflow() *InstanceWorkflow`

NewInstanceWorkflow instantiates a new InstanceWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWorkflowWithDefaults

`func NewInstanceWorkflowWithDefaults() *InstanceWorkflow`

NewInstanceWorkflowWithDefaults instantiates a new InstanceWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskSet

`func (o *InstanceWorkflow) GetTaskSet() RunWorkflowInstanceRequestTaskSet`

GetTaskSet returns the TaskSet field if non-nil, zero value otherwise.

### GetTaskSetOk

`func (o *InstanceWorkflow) GetTaskSetOk() (*RunWorkflowInstanceRequestTaskSet, bool)`

GetTaskSetOk returns a tuple with the TaskSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSet

`func (o *InstanceWorkflow) SetTaskSet(v RunWorkflowInstanceRequestTaskSet)`

SetTaskSet sets TaskSet field to given value.

### HasTaskSet

`func (o *InstanceWorkflow) HasTaskSet() bool`

HasTaskSet returns a boolean if a field has been set.

### GetTaskPhase

`func (o *InstanceWorkflow) GetTaskPhase() string`

GetTaskPhase returns the TaskPhase field if non-nil, zero value otherwise.

### GetTaskPhaseOk

`func (o *InstanceWorkflow) GetTaskPhaseOk() (*string, bool)`

GetTaskPhaseOk returns a tuple with the TaskPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPhase

`func (o *InstanceWorkflow) SetTaskPhase(v string)`

SetTaskPhase sets TaskPhase field to given value.

### HasTaskPhase

`func (o *InstanceWorkflow) HasTaskPhase() bool`

HasTaskPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



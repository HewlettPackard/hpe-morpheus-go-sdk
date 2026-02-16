# TaskConditionalWorkflowConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionalScript** | Pointer to **NullableString** | Allows the user to set JavaScript logic. If it resolves to true, HPE Morpheus Enterprise will run the Operational Workflow set as the “IF OPERATIONAL WORKFLOW” and if it resolves to false, HPE Morpheus Enterprise will run the “ELSE OPERATIONAL WORKFLOW” | [optional] 
**IfOperationalWorkflowId** | Pointer to **NullableInt64** | If Operational Workflow ID | [optional] 
**IfOperationalWorkflowName** | Pointer to **NullableString** | If Operational Workflow Name | [optional] 
**ElseOperationalWorkflowId** | Pointer to **NullableInt64** | else Operational Workflow ID | [optional] 
**ElseOperationalWorkflowName** | Pointer to **NullableString** | Else Operational Workflow Name | [optional] 

## Methods

### NewTaskConditionalWorkflowConfig

`func NewTaskConditionalWorkflowConfig() *TaskConditionalWorkflowConfig`

NewTaskConditionalWorkflowConfig instantiates a new TaskConditionalWorkflowConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskConditionalWorkflowConfigWithDefaults

`func NewTaskConditionalWorkflowConfigWithDefaults() *TaskConditionalWorkflowConfig`

NewTaskConditionalWorkflowConfigWithDefaults instantiates a new TaskConditionalWorkflowConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionalScript

`func (o *TaskConditionalWorkflowConfig) GetConditionalScript() string`

GetConditionalScript returns the ConditionalScript field if non-nil, zero value otherwise.

### GetConditionalScriptOk

`func (o *TaskConditionalWorkflowConfig) GetConditionalScriptOk() (*string, bool)`

GetConditionalScriptOk returns a tuple with the ConditionalScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalScript

`func (o *TaskConditionalWorkflowConfig) SetConditionalScript(v string)`

SetConditionalScript sets ConditionalScript field to given value.

### HasConditionalScript

`func (o *TaskConditionalWorkflowConfig) HasConditionalScript() bool`

HasConditionalScript returns a boolean if a field has been set.

### SetConditionalScriptNil

`func (o *TaskConditionalWorkflowConfig) SetConditionalScriptNil(b bool)`

 SetConditionalScriptNil sets the value for ConditionalScript to be an explicit nil

### UnsetConditionalScript
`func (o *TaskConditionalWorkflowConfig) UnsetConditionalScript()`

UnsetConditionalScript ensures that no value is present for ConditionalScript, not even an explicit nil
### GetIfOperationalWorkflowId

`func (o *TaskConditionalWorkflowConfig) GetIfOperationalWorkflowId() int64`

GetIfOperationalWorkflowId returns the IfOperationalWorkflowId field if non-nil, zero value otherwise.

### GetIfOperationalWorkflowIdOk

`func (o *TaskConditionalWorkflowConfig) GetIfOperationalWorkflowIdOk() (*int64, bool)`

GetIfOperationalWorkflowIdOk returns a tuple with the IfOperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfOperationalWorkflowId

`func (o *TaskConditionalWorkflowConfig) SetIfOperationalWorkflowId(v int64)`

SetIfOperationalWorkflowId sets IfOperationalWorkflowId field to given value.

### HasIfOperationalWorkflowId

`func (o *TaskConditionalWorkflowConfig) HasIfOperationalWorkflowId() bool`

HasIfOperationalWorkflowId returns a boolean if a field has been set.

### SetIfOperationalWorkflowIdNil

`func (o *TaskConditionalWorkflowConfig) SetIfOperationalWorkflowIdNil(b bool)`

 SetIfOperationalWorkflowIdNil sets the value for IfOperationalWorkflowId to be an explicit nil

### UnsetIfOperationalWorkflowId
`func (o *TaskConditionalWorkflowConfig) UnsetIfOperationalWorkflowId()`

UnsetIfOperationalWorkflowId ensures that no value is present for IfOperationalWorkflowId, not even an explicit nil
### GetIfOperationalWorkflowName

`func (o *TaskConditionalWorkflowConfig) GetIfOperationalWorkflowName() string`

GetIfOperationalWorkflowName returns the IfOperationalWorkflowName field if non-nil, zero value otherwise.

### GetIfOperationalWorkflowNameOk

`func (o *TaskConditionalWorkflowConfig) GetIfOperationalWorkflowNameOk() (*string, bool)`

GetIfOperationalWorkflowNameOk returns a tuple with the IfOperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfOperationalWorkflowName

`func (o *TaskConditionalWorkflowConfig) SetIfOperationalWorkflowName(v string)`

SetIfOperationalWorkflowName sets IfOperationalWorkflowName field to given value.

### HasIfOperationalWorkflowName

`func (o *TaskConditionalWorkflowConfig) HasIfOperationalWorkflowName() bool`

HasIfOperationalWorkflowName returns a boolean if a field has been set.

### SetIfOperationalWorkflowNameNil

`func (o *TaskConditionalWorkflowConfig) SetIfOperationalWorkflowNameNil(b bool)`

 SetIfOperationalWorkflowNameNil sets the value for IfOperationalWorkflowName to be an explicit nil

### UnsetIfOperationalWorkflowName
`func (o *TaskConditionalWorkflowConfig) UnsetIfOperationalWorkflowName()`

UnsetIfOperationalWorkflowName ensures that no value is present for IfOperationalWorkflowName, not even an explicit nil
### GetElseOperationalWorkflowId

`func (o *TaskConditionalWorkflowConfig) GetElseOperationalWorkflowId() int64`

GetElseOperationalWorkflowId returns the ElseOperationalWorkflowId field if non-nil, zero value otherwise.

### GetElseOperationalWorkflowIdOk

`func (o *TaskConditionalWorkflowConfig) GetElseOperationalWorkflowIdOk() (*int64, bool)`

GetElseOperationalWorkflowIdOk returns a tuple with the ElseOperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseOperationalWorkflowId

`func (o *TaskConditionalWorkflowConfig) SetElseOperationalWorkflowId(v int64)`

SetElseOperationalWorkflowId sets ElseOperationalWorkflowId field to given value.

### HasElseOperationalWorkflowId

`func (o *TaskConditionalWorkflowConfig) HasElseOperationalWorkflowId() bool`

HasElseOperationalWorkflowId returns a boolean if a field has been set.

### SetElseOperationalWorkflowIdNil

`func (o *TaskConditionalWorkflowConfig) SetElseOperationalWorkflowIdNil(b bool)`

 SetElseOperationalWorkflowIdNil sets the value for ElseOperationalWorkflowId to be an explicit nil

### UnsetElseOperationalWorkflowId
`func (o *TaskConditionalWorkflowConfig) UnsetElseOperationalWorkflowId()`

UnsetElseOperationalWorkflowId ensures that no value is present for ElseOperationalWorkflowId, not even an explicit nil
### GetElseOperationalWorkflowName

`func (o *TaskConditionalWorkflowConfig) GetElseOperationalWorkflowName() string`

GetElseOperationalWorkflowName returns the ElseOperationalWorkflowName field if non-nil, zero value otherwise.

### GetElseOperationalWorkflowNameOk

`func (o *TaskConditionalWorkflowConfig) GetElseOperationalWorkflowNameOk() (*string, bool)`

GetElseOperationalWorkflowNameOk returns a tuple with the ElseOperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseOperationalWorkflowName

`func (o *TaskConditionalWorkflowConfig) SetElseOperationalWorkflowName(v string)`

SetElseOperationalWorkflowName sets ElseOperationalWorkflowName field to given value.

### HasElseOperationalWorkflowName

`func (o *TaskConditionalWorkflowConfig) HasElseOperationalWorkflowName() bool`

HasElseOperationalWorkflowName returns a boolean if a field has been set.

### SetElseOperationalWorkflowNameNil

`func (o *TaskConditionalWorkflowConfig) SetElseOperationalWorkflowNameNil(b bool)`

 SetElseOperationalWorkflowNameNil sets the value for ElseOperationalWorkflowName to be an explicit nil

### UnsetElseOperationalWorkflowName
`func (o *TaskConditionalWorkflowConfig) UnsetElseOperationalWorkflowName()`

UnsetElseOperationalWorkflowName ensures that no value is present for ElseOperationalWorkflowName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



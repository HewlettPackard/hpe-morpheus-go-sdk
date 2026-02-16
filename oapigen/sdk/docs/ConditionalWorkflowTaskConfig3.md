# ConditionalWorkflowTaskConfig3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionalScript** | Pointer to **NullableString** | Allows the user to set JavaScript logic. If it resolves to true, HPE Morpheus Enterprise will run the Operational Workflow set as the “IF OPERATIONAL WORKFLOW” and if it resolves to false, HPE Morpheus Enterprise will run the “ELSE OPERATIONAL WORKFLOW” | [optional] 
**IfOperationalWorkflowId** | Pointer to **NullableInt64** | If Operational Workflow ID | [optional] 
**IfOperationalWorkflowName** | Pointer to **NullableString** | If Operational Workflow Name | [optional] 
**ElseOperationalWorkflowId** | Pointer to **NullableInt64** | else Operational Workflow ID | [optional] 
**ElseOperationalWorkflowName** | Pointer to **NullableString** | Else Operational Workflow Name | [optional] 

## Methods

### NewConditionalWorkflowTaskConfig3

`func NewConditionalWorkflowTaskConfig3() *ConditionalWorkflowTaskConfig3`

NewConditionalWorkflowTaskConfig3 instantiates a new ConditionalWorkflowTaskConfig3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionalWorkflowTaskConfig3WithDefaults

`func NewConditionalWorkflowTaskConfig3WithDefaults() *ConditionalWorkflowTaskConfig3`

NewConditionalWorkflowTaskConfig3WithDefaults instantiates a new ConditionalWorkflowTaskConfig3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionalScript

`func (o *ConditionalWorkflowTaskConfig3) GetConditionalScript() string`

GetConditionalScript returns the ConditionalScript field if non-nil, zero value otherwise.

### GetConditionalScriptOk

`func (o *ConditionalWorkflowTaskConfig3) GetConditionalScriptOk() (*string, bool)`

GetConditionalScriptOk returns a tuple with the ConditionalScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalScript

`func (o *ConditionalWorkflowTaskConfig3) SetConditionalScript(v string)`

SetConditionalScript sets ConditionalScript field to given value.

### HasConditionalScript

`func (o *ConditionalWorkflowTaskConfig3) HasConditionalScript() bool`

HasConditionalScript returns a boolean if a field has been set.

### SetConditionalScriptNil

`func (o *ConditionalWorkflowTaskConfig3) SetConditionalScriptNil(b bool)`

 SetConditionalScriptNil sets the value for ConditionalScript to be an explicit nil

### UnsetConditionalScript
`func (o *ConditionalWorkflowTaskConfig3) UnsetConditionalScript()`

UnsetConditionalScript ensures that no value is present for ConditionalScript, not even an explicit nil
### GetIfOperationalWorkflowId

`func (o *ConditionalWorkflowTaskConfig3) GetIfOperationalWorkflowId() int64`

GetIfOperationalWorkflowId returns the IfOperationalWorkflowId field if non-nil, zero value otherwise.

### GetIfOperationalWorkflowIdOk

`func (o *ConditionalWorkflowTaskConfig3) GetIfOperationalWorkflowIdOk() (*int64, bool)`

GetIfOperationalWorkflowIdOk returns a tuple with the IfOperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfOperationalWorkflowId

`func (o *ConditionalWorkflowTaskConfig3) SetIfOperationalWorkflowId(v int64)`

SetIfOperationalWorkflowId sets IfOperationalWorkflowId field to given value.

### HasIfOperationalWorkflowId

`func (o *ConditionalWorkflowTaskConfig3) HasIfOperationalWorkflowId() bool`

HasIfOperationalWorkflowId returns a boolean if a field has been set.

### SetIfOperationalWorkflowIdNil

`func (o *ConditionalWorkflowTaskConfig3) SetIfOperationalWorkflowIdNil(b bool)`

 SetIfOperationalWorkflowIdNil sets the value for IfOperationalWorkflowId to be an explicit nil

### UnsetIfOperationalWorkflowId
`func (o *ConditionalWorkflowTaskConfig3) UnsetIfOperationalWorkflowId()`

UnsetIfOperationalWorkflowId ensures that no value is present for IfOperationalWorkflowId, not even an explicit nil
### GetIfOperationalWorkflowName

`func (o *ConditionalWorkflowTaskConfig3) GetIfOperationalWorkflowName() string`

GetIfOperationalWorkflowName returns the IfOperationalWorkflowName field if non-nil, zero value otherwise.

### GetIfOperationalWorkflowNameOk

`func (o *ConditionalWorkflowTaskConfig3) GetIfOperationalWorkflowNameOk() (*string, bool)`

GetIfOperationalWorkflowNameOk returns a tuple with the IfOperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfOperationalWorkflowName

`func (o *ConditionalWorkflowTaskConfig3) SetIfOperationalWorkflowName(v string)`

SetIfOperationalWorkflowName sets IfOperationalWorkflowName field to given value.

### HasIfOperationalWorkflowName

`func (o *ConditionalWorkflowTaskConfig3) HasIfOperationalWorkflowName() bool`

HasIfOperationalWorkflowName returns a boolean if a field has been set.

### SetIfOperationalWorkflowNameNil

`func (o *ConditionalWorkflowTaskConfig3) SetIfOperationalWorkflowNameNil(b bool)`

 SetIfOperationalWorkflowNameNil sets the value for IfOperationalWorkflowName to be an explicit nil

### UnsetIfOperationalWorkflowName
`func (o *ConditionalWorkflowTaskConfig3) UnsetIfOperationalWorkflowName()`

UnsetIfOperationalWorkflowName ensures that no value is present for IfOperationalWorkflowName, not even an explicit nil
### GetElseOperationalWorkflowId

`func (o *ConditionalWorkflowTaskConfig3) GetElseOperationalWorkflowId() int64`

GetElseOperationalWorkflowId returns the ElseOperationalWorkflowId field if non-nil, zero value otherwise.

### GetElseOperationalWorkflowIdOk

`func (o *ConditionalWorkflowTaskConfig3) GetElseOperationalWorkflowIdOk() (*int64, bool)`

GetElseOperationalWorkflowIdOk returns a tuple with the ElseOperationalWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseOperationalWorkflowId

`func (o *ConditionalWorkflowTaskConfig3) SetElseOperationalWorkflowId(v int64)`

SetElseOperationalWorkflowId sets ElseOperationalWorkflowId field to given value.

### HasElseOperationalWorkflowId

`func (o *ConditionalWorkflowTaskConfig3) HasElseOperationalWorkflowId() bool`

HasElseOperationalWorkflowId returns a boolean if a field has been set.

### SetElseOperationalWorkflowIdNil

`func (o *ConditionalWorkflowTaskConfig3) SetElseOperationalWorkflowIdNil(b bool)`

 SetElseOperationalWorkflowIdNil sets the value for ElseOperationalWorkflowId to be an explicit nil

### UnsetElseOperationalWorkflowId
`func (o *ConditionalWorkflowTaskConfig3) UnsetElseOperationalWorkflowId()`

UnsetElseOperationalWorkflowId ensures that no value is present for ElseOperationalWorkflowId, not even an explicit nil
### GetElseOperationalWorkflowName

`func (o *ConditionalWorkflowTaskConfig3) GetElseOperationalWorkflowName() string`

GetElseOperationalWorkflowName returns the ElseOperationalWorkflowName field if non-nil, zero value otherwise.

### GetElseOperationalWorkflowNameOk

`func (o *ConditionalWorkflowTaskConfig3) GetElseOperationalWorkflowNameOk() (*string, bool)`

GetElseOperationalWorkflowNameOk returns a tuple with the ElseOperationalWorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseOperationalWorkflowName

`func (o *ConditionalWorkflowTaskConfig3) SetElseOperationalWorkflowName(v string)`

SetElseOperationalWorkflowName sets ElseOperationalWorkflowName field to given value.

### HasElseOperationalWorkflowName

`func (o *ConditionalWorkflowTaskConfig3) HasElseOperationalWorkflowName() bool`

HasElseOperationalWorkflowName returns a boolean if a field has been set.

### SetElseOperationalWorkflowNameNil

`func (o *ConditionalWorkflowTaskConfig3) SetElseOperationalWorkflowNameNil(b bool)`

 SetElseOperationalWorkflowNameNil sets the value for ElseOperationalWorkflowName to be an explicit nil

### UnsetElseOperationalWorkflowName
`func (o *ConditionalWorkflowTaskConfig3) UnsetElseOperationalWorkflowName()`

UnsetElseOperationalWorkflowName ensures that no value is present for ElseOperationalWorkflowName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WorkflowJobPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the Job | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Enabled** | Pointer to **bool** | Use this to set enabled state | [optional] [default to true]
**Task** | Pointer to [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | [optional] 
**Workflow** | [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | 
**TargetType** | **string** | Target type where job will execute | 
**Targets** | Pointer to [**[]WorkflowJobPayloadTargetsInner**](WorkflowJobPayloadTargetsInner.md) |  | [optional] 
**InstanceLabel** | Pointer to **string** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional] 
**ServerLabel** | Pointer to **string** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional] 
**ScheduleMode** | [**WorkflowJobPayloadScheduleMode**](WorkflowJobPayloadScheduleMode.md) |  | 
**CustomOptions** | Pointer to **map[string]interface{}** | Map of options to be used as values in the workflow tasks. These correspond to option types. | [optional] 
**CustomConfig** | Pointer to **string** | Job custom configuration (String in JSON format) | [optional] 
**DateTime** | Pointer to **time.Time** | Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is &#39;dateTime&#39;. | [optional] 
**Run** | Pointer to **bool** | If true, executes job | [optional] 

## Methods

### NewWorkflowJobPayload

`func NewWorkflowJobPayload(name string, workflow WorkflowJobPayloadTask, targetType string, scheduleMode WorkflowJobPayloadScheduleMode, ) *WorkflowJobPayload`

NewWorkflowJobPayload instantiates a new WorkflowJobPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowJobPayloadWithDefaults

`func NewWorkflowJobPayloadWithDefaults() *WorkflowJobPayload`

NewWorkflowJobPayloadWithDefaults instantiates a new WorkflowJobPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkflowJobPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowJobPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowJobPayload) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *WorkflowJobPayload) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WorkflowJobPayload) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WorkflowJobPayload) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WorkflowJobPayload) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnabled

`func (o *WorkflowJobPayload) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowJobPayload) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowJobPayload) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowJobPayload) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTask

`func (o *WorkflowJobPayload) GetTask() WorkflowJobPayloadTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *WorkflowJobPayload) GetTaskOk() (*WorkflowJobPayloadTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *WorkflowJobPayload) SetTask(v WorkflowJobPayloadTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *WorkflowJobPayload) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetWorkflow

`func (o *WorkflowJobPayload) GetWorkflow() WorkflowJobPayloadTask`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WorkflowJobPayload) GetWorkflowOk() (*WorkflowJobPayloadTask, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WorkflowJobPayload) SetWorkflow(v WorkflowJobPayloadTask)`

SetWorkflow sets Workflow field to given value.


### GetTargetType

`func (o *WorkflowJobPayload) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *WorkflowJobPayload) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *WorkflowJobPayload) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetTargets

`func (o *WorkflowJobPayload) GetTargets() []WorkflowJobPayloadTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *WorkflowJobPayload) GetTargetsOk() (*[]WorkflowJobPayloadTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *WorkflowJobPayload) SetTargets(v []WorkflowJobPayloadTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *WorkflowJobPayload) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetInstanceLabel

`func (o *WorkflowJobPayload) GetInstanceLabel() string`

GetInstanceLabel returns the InstanceLabel field if non-nil, zero value otherwise.

### GetInstanceLabelOk

`func (o *WorkflowJobPayload) GetInstanceLabelOk() (*string, bool)`

GetInstanceLabelOk returns a tuple with the InstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLabel

`func (o *WorkflowJobPayload) SetInstanceLabel(v string)`

SetInstanceLabel sets InstanceLabel field to given value.

### HasInstanceLabel

`func (o *WorkflowJobPayload) HasInstanceLabel() bool`

HasInstanceLabel returns a boolean if a field has been set.

### GetServerLabel

`func (o *WorkflowJobPayload) GetServerLabel() string`

GetServerLabel returns the ServerLabel field if non-nil, zero value otherwise.

### GetServerLabelOk

`func (o *WorkflowJobPayload) GetServerLabelOk() (*string, bool)`

GetServerLabelOk returns a tuple with the ServerLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLabel

`func (o *WorkflowJobPayload) SetServerLabel(v string)`

SetServerLabel sets ServerLabel field to given value.

### HasServerLabel

`func (o *WorkflowJobPayload) HasServerLabel() bool`

HasServerLabel returns a boolean if a field has been set.

### GetScheduleMode

`func (o *WorkflowJobPayload) GetScheduleMode() WorkflowJobPayloadScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *WorkflowJobPayload) GetScheduleModeOk() (*WorkflowJobPayloadScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *WorkflowJobPayload) SetScheduleMode(v WorkflowJobPayloadScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.


### GetCustomOptions

`func (o *WorkflowJobPayload) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *WorkflowJobPayload) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *WorkflowJobPayload) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *WorkflowJobPayload) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetCustomConfig

`func (o *WorkflowJobPayload) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *WorkflowJobPayload) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *WorkflowJobPayload) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *WorkflowJobPayload) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetDateTime

`func (o *WorkflowJobPayload) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *WorkflowJobPayload) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *WorkflowJobPayload) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *WorkflowJobPayload) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetRun

`func (o *WorkflowJobPayload) GetRun() bool`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *WorkflowJobPayload) GetRunOk() (*bool, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *WorkflowJobPayload) SetRun(v bool)`

SetRun sets Run field to given value.

### HasRun

`func (o *WorkflowJobPayload) HasRun() bool`

HasRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



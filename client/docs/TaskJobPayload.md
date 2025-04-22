# TaskJobPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the Job | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Enabled** | Pointer to **bool** | Use this to set enabled state | [optional] [default to true]
**Task** | [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | 
**TargetType** | **string** | Target type where job will execute | 
**Targets** | Pointer to [**[]WorkflowJobPayloadTargetsInner**](WorkflowJobPayloadTargetsInner.md) |  | [optional] 
**ScheduleMode** | [**WorkflowJobPayloadScheduleMode**](WorkflowJobPayloadScheduleMode.md) |  | 
**CustomOptions** | Pointer to **map[string]interface{}** | Map of options to be used as values in the workflow tasks. These correspond to option types. | [optional] 
**CustomConfig** | Pointer to **string** | Job custom configuration (String in JSON format) | [optional] 
**DateTime** | Pointer to **time.Time** | Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is &#39;dateTime&#39;. | [optional] 
**Run** | Pointer to **bool** | If true, executes job | [optional] 

## Methods

### NewTaskJobPayload

`func NewTaskJobPayload(name string, task WorkflowJobPayloadTask, targetType string, scheduleMode WorkflowJobPayloadScheduleMode, ) *TaskJobPayload`

NewTaskJobPayload instantiates a new TaskJobPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskJobPayloadWithDefaults

`func NewTaskJobPayloadWithDefaults() *TaskJobPayload`

NewTaskJobPayloadWithDefaults instantiates a new TaskJobPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TaskJobPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskJobPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskJobPayload) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *TaskJobPayload) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *TaskJobPayload) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *TaskJobPayload) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *TaskJobPayload) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnabled

`func (o *TaskJobPayload) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TaskJobPayload) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TaskJobPayload) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TaskJobPayload) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTask

`func (o *TaskJobPayload) GetTask() WorkflowJobPayloadTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *TaskJobPayload) GetTaskOk() (*WorkflowJobPayloadTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *TaskJobPayload) SetTask(v WorkflowJobPayloadTask)`

SetTask sets Task field to given value.


### GetTargetType

`func (o *TaskJobPayload) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *TaskJobPayload) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *TaskJobPayload) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetTargets

`func (o *TaskJobPayload) GetTargets() []WorkflowJobPayloadTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *TaskJobPayload) GetTargetsOk() (*[]WorkflowJobPayloadTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *TaskJobPayload) SetTargets(v []WorkflowJobPayloadTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *TaskJobPayload) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetScheduleMode

`func (o *TaskJobPayload) GetScheduleMode() WorkflowJobPayloadScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *TaskJobPayload) GetScheduleModeOk() (*WorkflowJobPayloadScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *TaskJobPayload) SetScheduleMode(v WorkflowJobPayloadScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.


### GetCustomOptions

`func (o *TaskJobPayload) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *TaskJobPayload) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *TaskJobPayload) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *TaskJobPayload) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetCustomConfig

`func (o *TaskJobPayload) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *TaskJobPayload) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *TaskJobPayload) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *TaskJobPayload) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetDateTime

`func (o *TaskJobPayload) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *TaskJobPayload) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *TaskJobPayload) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *TaskJobPayload) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetRun

`func (o *TaskJobPayload) GetRun() bool`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *TaskJobPayload) GetRunOk() (*bool, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *TaskJobPayload) SetRun(v bool)`

SetRun sets Run field to given value.

### HasRun

`func (o *TaskJobPayload) HasRun() bool`

HasRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



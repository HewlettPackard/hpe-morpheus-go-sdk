# UpdateJobsRequestJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the Job | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Enabled** | Pointer to **bool** | Use this to set enabled state | [optional] [default to true]
**Task** | Pointer to [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | [optional] 
**Workflow** | Pointer to [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | [optional] 
**ScanPath** | Pointer to **string** | Scan Checklist. Only applies to type scap-package. | [optional] 
**SecurityProfile** | Pointer to **string** | Security Profile. Only applies to type scap-package. | [optional] 
**TargetType** | Pointer to **string** | Target type where job will execute | [optional] 
**Targets** | Pointer to [**[]WorkflowJobPayloadTargetsInner**](WorkflowJobPayloadTargetsInner.md) |  | [optional] 
**InstanceLabel** | Pointer to **string** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional] 
**ServerLabel** | Pointer to **string** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional] 
**ScheduleMode** | Pointer to [**WorkflowJobPayloadScheduleMode**](WorkflowJobPayloadScheduleMode.md) |  | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** | Map of options to be used as values in the workflow tasks. These correspond to option types. | [optional] 
**CustomConfig** | Pointer to **string** | Job custom configuration (String in JSON format) | [optional] 
**DateTime** | Pointer to **time.Time** | Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is &#39;dateTime&#39;. | [optional] 
**Run** | Pointer to **bool** | If true, executes job | [optional] 

## Methods

### NewUpdateJobsRequestJob

`func NewUpdateJobsRequestJob() *UpdateJobsRequestJob`

NewUpdateJobsRequestJob instantiates a new UpdateJobsRequestJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateJobsRequestJobWithDefaults

`func NewUpdateJobsRequestJobWithDefaults() *UpdateJobsRequestJob`

NewUpdateJobsRequestJobWithDefaults instantiates a new UpdateJobsRequestJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateJobsRequestJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateJobsRequestJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateJobsRequestJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateJobsRequestJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateJobsRequestJob) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateJobsRequestJob) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateJobsRequestJob) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateJobsRequestJob) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateJobsRequestJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateJobsRequestJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateJobsRequestJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateJobsRequestJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTask

`func (o *UpdateJobsRequestJob) GetTask() WorkflowJobPayloadTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *UpdateJobsRequestJob) GetTaskOk() (*WorkflowJobPayloadTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *UpdateJobsRequestJob) SetTask(v WorkflowJobPayloadTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *UpdateJobsRequestJob) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetWorkflow

`func (o *UpdateJobsRequestJob) GetWorkflow() WorkflowJobPayloadTask`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *UpdateJobsRequestJob) GetWorkflowOk() (*WorkflowJobPayloadTask, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *UpdateJobsRequestJob) SetWorkflow(v WorkflowJobPayloadTask)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *UpdateJobsRequestJob) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetScanPath

`func (o *UpdateJobsRequestJob) GetScanPath() string`

GetScanPath returns the ScanPath field if non-nil, zero value otherwise.

### GetScanPathOk

`func (o *UpdateJobsRequestJob) GetScanPathOk() (*string, bool)`

GetScanPathOk returns a tuple with the ScanPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPath

`func (o *UpdateJobsRequestJob) SetScanPath(v string)`

SetScanPath sets ScanPath field to given value.

### HasScanPath

`func (o *UpdateJobsRequestJob) HasScanPath() bool`

HasScanPath returns a boolean if a field has been set.

### GetSecurityProfile

`func (o *UpdateJobsRequestJob) GetSecurityProfile() string`

GetSecurityProfile returns the SecurityProfile field if non-nil, zero value otherwise.

### GetSecurityProfileOk

`func (o *UpdateJobsRequestJob) GetSecurityProfileOk() (*string, bool)`

GetSecurityProfileOk returns a tuple with the SecurityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProfile

`func (o *UpdateJobsRequestJob) SetSecurityProfile(v string)`

SetSecurityProfile sets SecurityProfile field to given value.

### HasSecurityProfile

`func (o *UpdateJobsRequestJob) HasSecurityProfile() bool`

HasSecurityProfile returns a boolean if a field has been set.

### GetTargetType

`func (o *UpdateJobsRequestJob) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *UpdateJobsRequestJob) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *UpdateJobsRequestJob) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *UpdateJobsRequestJob) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargets

`func (o *UpdateJobsRequestJob) GetTargets() []WorkflowJobPayloadTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *UpdateJobsRequestJob) GetTargetsOk() (*[]WorkflowJobPayloadTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *UpdateJobsRequestJob) SetTargets(v []WorkflowJobPayloadTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *UpdateJobsRequestJob) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetInstanceLabel

`func (o *UpdateJobsRequestJob) GetInstanceLabel() string`

GetInstanceLabel returns the InstanceLabel field if non-nil, zero value otherwise.

### GetInstanceLabelOk

`func (o *UpdateJobsRequestJob) GetInstanceLabelOk() (*string, bool)`

GetInstanceLabelOk returns a tuple with the InstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLabel

`func (o *UpdateJobsRequestJob) SetInstanceLabel(v string)`

SetInstanceLabel sets InstanceLabel field to given value.

### HasInstanceLabel

`func (o *UpdateJobsRequestJob) HasInstanceLabel() bool`

HasInstanceLabel returns a boolean if a field has been set.

### GetServerLabel

`func (o *UpdateJobsRequestJob) GetServerLabel() string`

GetServerLabel returns the ServerLabel field if non-nil, zero value otherwise.

### GetServerLabelOk

`func (o *UpdateJobsRequestJob) GetServerLabelOk() (*string, bool)`

GetServerLabelOk returns a tuple with the ServerLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLabel

`func (o *UpdateJobsRequestJob) SetServerLabel(v string)`

SetServerLabel sets ServerLabel field to given value.

### HasServerLabel

`func (o *UpdateJobsRequestJob) HasServerLabel() bool`

HasServerLabel returns a boolean if a field has been set.

### GetScheduleMode

`func (o *UpdateJobsRequestJob) GetScheduleMode() WorkflowJobPayloadScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *UpdateJobsRequestJob) GetScheduleModeOk() (*WorkflowJobPayloadScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *UpdateJobsRequestJob) SetScheduleMode(v WorkflowJobPayloadScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *UpdateJobsRequestJob) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetCustomOptions

`func (o *UpdateJobsRequestJob) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *UpdateJobsRequestJob) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *UpdateJobsRequestJob) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *UpdateJobsRequestJob) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetCustomConfig

`func (o *UpdateJobsRequestJob) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *UpdateJobsRequestJob) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *UpdateJobsRequestJob) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *UpdateJobsRequestJob) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetDateTime

`func (o *UpdateJobsRequestJob) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *UpdateJobsRequestJob) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *UpdateJobsRequestJob) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *UpdateJobsRequestJob) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetRun

`func (o *UpdateJobsRequestJob) GetRun() bool`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *UpdateJobsRequestJob) GetRunOk() (*bool, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *UpdateJobsRequestJob) SetRun(v bool)`

SetRun sets Run field to given value.

### HasRun

`func (o *UpdateJobsRequestJob) HasRun() bool`

HasRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



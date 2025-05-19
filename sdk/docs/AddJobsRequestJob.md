# AddJobsRequestJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the Job | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Enabled** | Pointer to **bool** | Use this to set enabled state | [optional] [default to true]
**Task** | [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | 
**Workflow** | [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | 
**TargetType** | **string** | Target type where job will execute | 
**Targets** | [**[]WorkflowJobPayloadTargetsInner**](WorkflowJobPayloadTargetsInner.md) |  | 
**InstanceLabel** | Pointer to **string** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional] 
**ServerLabel** | Pointer to **string** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional] 
**ScheduleMode** | [**WorkflowJobPayloadScheduleMode**](WorkflowJobPayloadScheduleMode.md) |  | 
**CustomOptions** | Pointer to **map[string]interface{}** | Map of options to be used as values in the workflow tasks. These correspond to option types. | [optional] 
**CustomConfig** | Pointer to **string** | Job custom configuration (String in JSON format) | [optional] 
**DateTime** | Pointer to **time.Time** | Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is &#39;dateTime&#39;. | [optional] 
**Run** | Pointer to **bool** | If true, executes job | [optional] 
**SecurityPackage** | [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | 
**ScanPath** | Pointer to **string** | Scan Checklist | [optional] 
**SecurityProfile** | Pointer to **string** | Security Profile | [optional] 

## Methods

### NewAddJobsRequestJob

`func NewAddJobsRequestJob(name string, task WorkflowJobPayloadTask, workflow WorkflowJobPayloadTask, targetType string, targets []WorkflowJobPayloadTargetsInner, scheduleMode WorkflowJobPayloadScheduleMode, securityPackage WorkflowJobPayloadTask, ) *AddJobsRequestJob`

NewAddJobsRequestJob instantiates a new AddJobsRequestJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddJobsRequestJobWithDefaults

`func NewAddJobsRequestJobWithDefaults() *AddJobsRequestJob`

NewAddJobsRequestJobWithDefaults instantiates a new AddJobsRequestJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddJobsRequestJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddJobsRequestJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddJobsRequestJob) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *AddJobsRequestJob) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddJobsRequestJob) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddJobsRequestJob) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddJobsRequestJob) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnabled

`func (o *AddJobsRequestJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddJobsRequestJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddJobsRequestJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddJobsRequestJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTask

`func (o *AddJobsRequestJob) GetTask() WorkflowJobPayloadTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *AddJobsRequestJob) GetTaskOk() (*WorkflowJobPayloadTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *AddJobsRequestJob) SetTask(v WorkflowJobPayloadTask)`

SetTask sets Task field to given value.


### GetWorkflow

`func (o *AddJobsRequestJob) GetWorkflow() WorkflowJobPayloadTask`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *AddJobsRequestJob) GetWorkflowOk() (*WorkflowJobPayloadTask, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *AddJobsRequestJob) SetWorkflow(v WorkflowJobPayloadTask)`

SetWorkflow sets Workflow field to given value.


### GetTargetType

`func (o *AddJobsRequestJob) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AddJobsRequestJob) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AddJobsRequestJob) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetTargets

`func (o *AddJobsRequestJob) GetTargets() []WorkflowJobPayloadTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AddJobsRequestJob) GetTargetsOk() (*[]WorkflowJobPayloadTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AddJobsRequestJob) SetTargets(v []WorkflowJobPayloadTargetsInner)`

SetTargets sets Targets field to given value.


### GetInstanceLabel

`func (o *AddJobsRequestJob) GetInstanceLabel() string`

GetInstanceLabel returns the InstanceLabel field if non-nil, zero value otherwise.

### GetInstanceLabelOk

`func (o *AddJobsRequestJob) GetInstanceLabelOk() (*string, bool)`

GetInstanceLabelOk returns a tuple with the InstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLabel

`func (o *AddJobsRequestJob) SetInstanceLabel(v string)`

SetInstanceLabel sets InstanceLabel field to given value.

### HasInstanceLabel

`func (o *AddJobsRequestJob) HasInstanceLabel() bool`

HasInstanceLabel returns a boolean if a field has been set.

### GetServerLabel

`func (o *AddJobsRequestJob) GetServerLabel() string`

GetServerLabel returns the ServerLabel field if non-nil, zero value otherwise.

### GetServerLabelOk

`func (o *AddJobsRequestJob) GetServerLabelOk() (*string, bool)`

GetServerLabelOk returns a tuple with the ServerLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLabel

`func (o *AddJobsRequestJob) SetServerLabel(v string)`

SetServerLabel sets ServerLabel field to given value.

### HasServerLabel

`func (o *AddJobsRequestJob) HasServerLabel() bool`

HasServerLabel returns a boolean if a field has been set.

### GetScheduleMode

`func (o *AddJobsRequestJob) GetScheduleMode() WorkflowJobPayloadScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *AddJobsRequestJob) GetScheduleModeOk() (*WorkflowJobPayloadScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *AddJobsRequestJob) SetScheduleMode(v WorkflowJobPayloadScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.


### GetCustomOptions

`func (o *AddJobsRequestJob) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *AddJobsRequestJob) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *AddJobsRequestJob) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *AddJobsRequestJob) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetCustomConfig

`func (o *AddJobsRequestJob) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *AddJobsRequestJob) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *AddJobsRequestJob) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *AddJobsRequestJob) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetDateTime

`func (o *AddJobsRequestJob) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *AddJobsRequestJob) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *AddJobsRequestJob) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *AddJobsRequestJob) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetRun

`func (o *AddJobsRequestJob) GetRun() bool`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *AddJobsRequestJob) GetRunOk() (*bool, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *AddJobsRequestJob) SetRun(v bool)`

SetRun sets Run field to given value.

### HasRun

`func (o *AddJobsRequestJob) HasRun() bool`

HasRun returns a boolean if a field has been set.

### GetSecurityPackage

`func (o *AddJobsRequestJob) GetSecurityPackage() WorkflowJobPayloadTask`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *AddJobsRequestJob) GetSecurityPackageOk() (*WorkflowJobPayloadTask, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *AddJobsRequestJob) SetSecurityPackage(v WorkflowJobPayloadTask)`

SetSecurityPackage sets SecurityPackage field to given value.


### GetScanPath

`func (o *AddJobsRequestJob) GetScanPath() string`

GetScanPath returns the ScanPath field if non-nil, zero value otherwise.

### GetScanPathOk

`func (o *AddJobsRequestJob) GetScanPathOk() (*string, bool)`

GetScanPathOk returns a tuple with the ScanPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPath

`func (o *AddJobsRequestJob) SetScanPath(v string)`

SetScanPath sets ScanPath field to given value.

### HasScanPath

`func (o *AddJobsRequestJob) HasScanPath() bool`

HasScanPath returns a boolean if a field has been set.

### GetSecurityProfile

`func (o *AddJobsRequestJob) GetSecurityProfile() string`

GetSecurityProfile returns the SecurityProfile field if non-nil, zero value otherwise.

### GetSecurityProfileOk

`func (o *AddJobsRequestJob) GetSecurityProfileOk() (*string, bool)`

GetSecurityProfileOk returns a tuple with the SecurityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProfile

`func (o *AddJobsRequestJob) SetSecurityProfile(v string)`

SetSecurityProfile sets SecurityProfile field to given value.

### HasSecurityProfile

`func (o *AddJobsRequestJob) HasSecurityProfile() bool`

HasSecurityProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



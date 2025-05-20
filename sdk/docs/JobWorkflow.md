# JobWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Workflow** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**JobSummary** | Pointer to **string** |  | [optional] 
**ScheduleMode** | Pointer to [**ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode**](ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode.md) |  | [optional] 
**DateTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastRun** | Pointer to **time.Time** |  | [optional] 
**LastResult** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy**](ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy.md) |  | [optional] 
**TargetType** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to [**[]ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner**](ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner.md) |  | [optional] 
**CustomConfig** | Pointer to **string** |  | [optional] 
**CustomOptions** | Pointer to [**ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions**](ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions.md) |  | [optional] 

## Methods

### NewJobWorkflow

`func NewJobWorkflow() *JobWorkflow`

NewJobWorkflow instantiates a new JobWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWorkflowWithDefaults

`func NewJobWorkflowWithDefaults() *JobWorkflow`

NewJobWorkflowWithDefaults instantiates a new JobWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobWorkflow) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobWorkflow) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobWorkflow) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *JobWorkflow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *JobWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobWorkflow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JobWorkflow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *JobWorkflow) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *JobWorkflow) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *JobWorkflow) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *JobWorkflow) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *JobWorkflow) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobWorkflow) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobWorkflow) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *JobWorkflow) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkflow

`func (o *JobWorkflow) GetWorkflow() GetAlerts200ResponseAllOfChecksInnerAccount`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *JobWorkflow) GetWorkflowOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *JobWorkflow) SetWorkflow(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *JobWorkflow) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetJobSummary

`func (o *JobWorkflow) GetJobSummary() string`

GetJobSummary returns the JobSummary field if non-nil, zero value otherwise.

### GetJobSummaryOk

`func (o *JobWorkflow) GetJobSummaryOk() (*string, bool)`

GetJobSummaryOk returns a tuple with the JobSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSummary

`func (o *JobWorkflow) SetJobSummary(v string)`

SetJobSummary sets JobSummary field to given value.

### HasJobSummary

`func (o *JobWorkflow) HasJobSummary() bool`

HasJobSummary returns a boolean if a field has been set.

### GetScheduleMode

`func (o *JobWorkflow) GetScheduleMode() ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *JobWorkflow) GetScheduleModeOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *JobWorkflow) SetScheduleMode(v ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *JobWorkflow) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetDateTime

`func (o *JobWorkflow) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *JobWorkflow) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *JobWorkflow) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *JobWorkflow) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *JobWorkflow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobWorkflow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobWorkflow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JobWorkflow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNamespace

`func (o *JobWorkflow) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *JobWorkflow) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *JobWorkflow) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *JobWorkflow) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCategory

`func (o *JobWorkflow) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *JobWorkflow) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *JobWorkflow) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *JobWorkflow) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *JobWorkflow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobWorkflow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobWorkflow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JobWorkflow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *JobWorkflow) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JobWorkflow) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JobWorkflow) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JobWorkflow) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *JobWorkflow) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *JobWorkflow) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *JobWorkflow) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *JobWorkflow) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *JobWorkflow) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *JobWorkflow) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *JobWorkflow) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *JobWorkflow) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastRun

`func (o *JobWorkflow) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *JobWorkflow) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *JobWorkflow) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *JobWorkflow) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastResult

`func (o *JobWorkflow) GetLastResult() string`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *JobWorkflow) GetLastResultOk() (*string, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *JobWorkflow) SetLastResult(v string)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *JobWorkflow) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetCreatedBy

`func (o *JobWorkflow) GetCreatedBy() ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *JobWorkflow) GetCreatedByOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *JobWorkflow) SetCreatedBy(v ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *JobWorkflow) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTargetType

`func (o *JobWorkflow) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *JobWorkflow) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *JobWorkflow) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *JobWorkflow) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargets

`func (o *JobWorkflow) GetTargets() []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *JobWorkflow) GetTargetsOk() (*[]ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *JobWorkflow) SetTargets(v []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *JobWorkflow) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetCustomConfig

`func (o *JobWorkflow) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *JobWorkflow) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *JobWorkflow) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *JobWorkflow) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetCustomOptions

`func (o *JobWorkflow) GetCustomOptions() ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *JobWorkflow) GetCustomOptionsOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *JobWorkflow) SetCustomOptions(v ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *JobWorkflow) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



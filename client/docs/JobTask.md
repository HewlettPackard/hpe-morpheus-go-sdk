# JobTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Task** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
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

### NewJobTask

`func NewJobTask() *JobTask`

NewJobTask instantiates a new JobTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobTaskWithDefaults

`func NewJobTaskWithDefaults() *JobTask`

NewJobTaskWithDefaults instantiates a new JobTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *JobTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *JobTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JobTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *JobTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *JobTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *JobTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *JobTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *JobTask) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobTask) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobTask) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *JobTask) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTask

`func (o *JobTask) GetTask() GetAlerts200ResponseAllOfChecksInnerAccount`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *JobTask) GetTaskOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *JobTask) SetTask(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTask sets Task field to given value.

### HasTask

`func (o *JobTask) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetJobSummary

`func (o *JobTask) GetJobSummary() string`

GetJobSummary returns the JobSummary field if non-nil, zero value otherwise.

### GetJobSummaryOk

`func (o *JobTask) GetJobSummaryOk() (*string, bool)`

GetJobSummaryOk returns a tuple with the JobSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSummary

`func (o *JobTask) SetJobSummary(v string)`

SetJobSummary sets JobSummary field to given value.

### HasJobSummary

`func (o *JobTask) HasJobSummary() bool`

HasJobSummary returns a boolean if a field has been set.

### GetScheduleMode

`func (o *JobTask) GetScheduleMode() ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *JobTask) GetScheduleModeOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *JobTask) SetScheduleMode(v ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *JobTask) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetDateTime

`func (o *JobTask) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *JobTask) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *JobTask) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *JobTask) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *JobTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JobTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNamespace

`func (o *JobTask) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *JobTask) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *JobTask) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *JobTask) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCategory

`func (o *JobTask) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *JobTask) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *JobTask) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *JobTask) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *JobTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JobTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *JobTask) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JobTask) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JobTask) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *JobTask) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *JobTask) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *JobTask) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *JobTask) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *JobTask) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *JobTask) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *JobTask) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *JobTask) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *JobTask) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastRun

`func (o *JobTask) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *JobTask) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *JobTask) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *JobTask) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastResult

`func (o *JobTask) GetLastResult() string`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *JobTask) GetLastResultOk() (*string, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *JobTask) SetLastResult(v string)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *JobTask) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetCreatedBy

`func (o *JobTask) GetCreatedBy() ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *JobTask) GetCreatedByOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *JobTask) SetCreatedBy(v ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *JobTask) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTargetType

`func (o *JobTask) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *JobTask) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *JobTask) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *JobTask) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargets

`func (o *JobTask) GetTargets() []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *JobTask) GetTargetsOk() (*[]ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *JobTask) SetTargets(v []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *JobTask) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetCustomConfig

`func (o *JobTask) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *JobTask) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *JobTask) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *JobTask) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetCustomOptions

`func (o *JobTask) GetCustomOptions() ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *JobTask) GetCustomOptionsOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *JobTask) SetCustomOptions(v ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *JobTask) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



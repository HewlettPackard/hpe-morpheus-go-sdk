# ListJobs200ResponseAllOfJobsInnerAnyOf2

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

### NewListJobs200ResponseAllOfJobsInnerAnyOf2

`func NewListJobs200ResponseAllOfJobsInnerAnyOf2() *ListJobs200ResponseAllOfJobsInnerAnyOf2`

NewListJobs200ResponseAllOfJobsInnerAnyOf2 instantiates a new ListJobs200ResponseAllOfJobsInnerAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListJobs200ResponseAllOfJobsInnerAnyOf2WithDefaults

`func NewListJobs200ResponseAllOfJobsInnerAnyOf2WithDefaults() *ListJobs200ResponseAllOfJobsInnerAnyOf2`

NewListJobs200ResponseAllOfJobsInnerAnyOf2WithDefaults instantiates a new ListJobs200ResponseAllOfJobsInnerAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkflow

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetWorkflow() GetAlerts200ResponseAllOfChecksInnerAccount`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetWorkflowOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetWorkflow(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetJobSummary

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetJobSummary() string`

GetJobSummary returns the JobSummary field if non-nil, zero value otherwise.

### GetJobSummaryOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetJobSummaryOk() (*string, bool)`

GetJobSummaryOk returns a tuple with the JobSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSummary

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetJobSummary(v string)`

SetJobSummary sets JobSummary field to given value.

### HasJobSummary

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasJobSummary() bool`

HasJobSummary returns a boolean if a field has been set.

### GetScheduleMode

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetScheduleMode() ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetScheduleModeOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetScheduleMode(v ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetDateTime

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNamespace

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCategory

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastRun

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastResult

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetLastResult() string`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetLastResultOk() (*string, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetLastResult(v string)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetCreatedBy() ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetCreatedByOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetCreatedBy(v ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTargetType

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargets

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetTargets() []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetTargetsOk() (*[]ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetTargets(v []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetCustomConfig

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetCustomOptions

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetCustomOptions() ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) GetCustomOptionsOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) SetCustomOptions(v ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *ListJobs200ResponseAllOfJobsInnerAnyOf2) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



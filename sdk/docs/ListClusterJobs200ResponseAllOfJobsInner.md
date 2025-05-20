# ListClusterJobs200ResponseAllOfJobsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**JobSummary** | Pointer to **string** |  | [optional] 
**ScheduleMode** | Pointer to **string** |  | [optional] 
**DateTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastRun** | Pointer to **time.Time** |  | [optional] 
**LastResult** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerCreatedBy**](GetAlerts200ResponseAllOfChecksInnerCreatedBy.md) |  | [optional] 
**TargetType** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CustomConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewListClusterJobs200ResponseAllOfJobsInner

`func NewListClusterJobs200ResponseAllOfJobsInner() *ListClusterJobs200ResponseAllOfJobsInner`

NewListClusterJobs200ResponseAllOfJobsInner instantiates a new ListClusterJobs200ResponseAllOfJobsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterJobs200ResponseAllOfJobsInnerWithDefaults

`func NewListClusterJobs200ResponseAllOfJobsInnerWithDefaults() *ListClusterJobs200ResponseAllOfJobsInner`

NewListClusterJobs200ResponseAllOfJobsInnerWithDefaults instantiates a new ListClusterJobs200ResponseAllOfJobsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetJobSummary

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetJobSummary() string`

GetJobSummary returns the JobSummary field if non-nil, zero value otherwise.

### GetJobSummaryOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetJobSummaryOk() (*string, bool)`

GetJobSummaryOk returns a tuple with the JobSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSummary

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetJobSummary(v string)`

SetJobSummary sets JobSummary field to given value.

### HasJobSummary

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasJobSummary() bool`

HasJobSummary returns a boolean if a field has been set.

### GetScheduleMode

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetScheduleMode() string`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetScheduleModeOk() (*string, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetScheduleMode(v string)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetDateTime

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNamespace

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCategory

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastRun

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastResult

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetLastResult() string`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetLastResultOk() (*string, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetLastResult(v string)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetCreatedBy() GetAlerts200ResponseAllOfChecksInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetCreatedByOk() (*GetAlerts200ResponseAllOfChecksInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetCreatedBy(v GetAlerts200ResponseAllOfChecksInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTargetType

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargets

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetTargets() []map[string]interface{}`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetTargetsOk() (*[]map[string]interface{}, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetTargets(v []map[string]interface{})`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetCustomConfig

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetCustomConfig() map[string]interface{}`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetCustomConfigOk() (*map[string]interface{}, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetCustomConfig(v map[string]interface{})`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetCustomOptions

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *ListClusterJobs200ResponseAllOfJobsInner) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *ListClusterJobs200ResponseAllOfJobsInner) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *ListClusterJobs200ResponseAllOfJobsInner) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



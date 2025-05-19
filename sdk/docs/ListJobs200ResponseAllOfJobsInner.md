# ListJobs200ResponseAllOfJobsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Workflow** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Task** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**SecurityPackage** | Pointer to [**ListJobs200ResponseAllOfJobsInnerAnyOfSecurityPackage**](ListJobs200ResponseAllOfJobsInnerAnyOfSecurityPackage.md) |  | [optional] 
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
**ScanPath** | Pointer to **string** | Scan Checklist. Only applies to type scap-package. | [optional] 
**SecurityProfile** | Pointer to **string** | Security Profile. Only applies to type scap-package. | [optional] 
**CustomConfig** | Pointer to **string** |  | [optional] 
**CustomOptions** | Pointer to [**ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions**](ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions.md) |  | [optional] 

## Methods

### NewListJobs200ResponseAllOfJobsInner

`func NewListJobs200ResponseAllOfJobsInner() *ListJobs200ResponseAllOfJobsInner`

NewListJobs200ResponseAllOfJobsInner instantiates a new ListJobs200ResponseAllOfJobsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListJobs200ResponseAllOfJobsInnerWithDefaults

`func NewListJobs200ResponseAllOfJobsInnerWithDefaults() *ListJobs200ResponseAllOfJobsInner`

NewListJobs200ResponseAllOfJobsInnerWithDefaults instantiates a new ListJobs200ResponseAllOfJobsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListJobs200ResponseAllOfJobsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListJobs200ResponseAllOfJobsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListJobs200ResponseAllOfJobsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListJobs200ResponseAllOfJobsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListJobs200ResponseAllOfJobsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListJobs200ResponseAllOfJobsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListJobs200ResponseAllOfJobsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListJobs200ResponseAllOfJobsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListJobs200ResponseAllOfJobsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *ListJobs200ResponseAllOfJobsInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListJobs200ResponseAllOfJobsInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListJobs200ResponseAllOfJobsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkflow

`func (o *ListJobs200ResponseAllOfJobsInner) GetWorkflow() GetAlerts200ResponseAllOfChecksInnerAccount`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetWorkflowOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ListJobs200ResponseAllOfJobsInner) SetWorkflow(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ListJobs200ResponseAllOfJobsInner) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetTask

`func (o *ListJobs200ResponseAllOfJobsInner) GetTask() GetAlerts200ResponseAllOfChecksInnerAccount`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetTaskOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ListJobs200ResponseAllOfJobsInner) SetTask(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTask sets Task field to given value.

### HasTask

`func (o *ListJobs200ResponseAllOfJobsInner) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetSecurityPackage

`func (o *ListJobs200ResponseAllOfJobsInner) GetSecurityPackage() ListJobs200ResponseAllOfJobsInnerAnyOfSecurityPackage`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetSecurityPackageOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfSecurityPackage, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *ListJobs200ResponseAllOfJobsInner) SetSecurityPackage(v ListJobs200ResponseAllOfJobsInnerAnyOfSecurityPackage)`

SetSecurityPackage sets SecurityPackage field to given value.

### HasSecurityPackage

`func (o *ListJobs200ResponseAllOfJobsInner) HasSecurityPackage() bool`

HasSecurityPackage returns a boolean if a field has been set.

### GetJobSummary

`func (o *ListJobs200ResponseAllOfJobsInner) GetJobSummary() string`

GetJobSummary returns the JobSummary field if non-nil, zero value otherwise.

### GetJobSummaryOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetJobSummaryOk() (*string, bool)`

GetJobSummaryOk returns a tuple with the JobSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSummary

`func (o *ListJobs200ResponseAllOfJobsInner) SetJobSummary(v string)`

SetJobSummary sets JobSummary field to given value.

### HasJobSummary

`func (o *ListJobs200ResponseAllOfJobsInner) HasJobSummary() bool`

HasJobSummary returns a boolean if a field has been set.

### GetScheduleMode

`func (o *ListJobs200ResponseAllOfJobsInner) GetScheduleMode() ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetScheduleModeOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *ListJobs200ResponseAllOfJobsInner) SetScheduleMode(v ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *ListJobs200ResponseAllOfJobsInner) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetDateTime

`func (o *ListJobs200ResponseAllOfJobsInner) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *ListJobs200ResponseAllOfJobsInner) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *ListJobs200ResponseAllOfJobsInner) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *ListJobs200ResponseAllOfJobsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListJobs200ResponseAllOfJobsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListJobs200ResponseAllOfJobsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNamespace

`func (o *ListJobs200ResponseAllOfJobsInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListJobs200ResponseAllOfJobsInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ListJobs200ResponseAllOfJobsInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCategory

`func (o *ListJobs200ResponseAllOfJobsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListJobs200ResponseAllOfJobsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListJobs200ResponseAllOfJobsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *ListJobs200ResponseAllOfJobsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListJobs200ResponseAllOfJobsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListJobs200ResponseAllOfJobsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ListJobs200ResponseAllOfJobsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListJobs200ResponseAllOfJobsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListJobs200ResponseAllOfJobsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListJobs200ResponseAllOfJobsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListJobs200ResponseAllOfJobsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListJobs200ResponseAllOfJobsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListJobs200ResponseAllOfJobsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListJobs200ResponseAllOfJobsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListJobs200ResponseAllOfJobsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastRun

`func (o *ListJobs200ResponseAllOfJobsInner) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *ListJobs200ResponseAllOfJobsInner) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *ListJobs200ResponseAllOfJobsInner) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastResult

`func (o *ListJobs200ResponseAllOfJobsInner) GetLastResult() string`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetLastResultOk() (*string, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *ListJobs200ResponseAllOfJobsInner) SetLastResult(v string)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *ListJobs200ResponseAllOfJobsInner) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListJobs200ResponseAllOfJobsInner) GetCreatedBy() ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetCreatedByOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListJobs200ResponseAllOfJobsInner) SetCreatedBy(v ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListJobs200ResponseAllOfJobsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTargetType

`func (o *ListJobs200ResponseAllOfJobsInner) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ListJobs200ResponseAllOfJobsInner) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ListJobs200ResponseAllOfJobsInner) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargets

`func (o *ListJobs200ResponseAllOfJobsInner) GetTargets() []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetTargetsOk() (*[]ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ListJobs200ResponseAllOfJobsInner) SetTargets(v []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ListJobs200ResponseAllOfJobsInner) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetScanPath

`func (o *ListJobs200ResponseAllOfJobsInner) GetScanPath() string`

GetScanPath returns the ScanPath field if non-nil, zero value otherwise.

### GetScanPathOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetScanPathOk() (*string, bool)`

GetScanPathOk returns a tuple with the ScanPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPath

`func (o *ListJobs200ResponseAllOfJobsInner) SetScanPath(v string)`

SetScanPath sets ScanPath field to given value.

### HasScanPath

`func (o *ListJobs200ResponseAllOfJobsInner) HasScanPath() bool`

HasScanPath returns a boolean if a field has been set.

### GetSecurityProfile

`func (o *ListJobs200ResponseAllOfJobsInner) GetSecurityProfile() string`

GetSecurityProfile returns the SecurityProfile field if non-nil, zero value otherwise.

### GetSecurityProfileOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetSecurityProfileOk() (*string, bool)`

GetSecurityProfileOk returns a tuple with the SecurityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProfile

`func (o *ListJobs200ResponseAllOfJobsInner) SetSecurityProfile(v string)`

SetSecurityProfile sets SecurityProfile field to given value.

### HasSecurityProfile

`func (o *ListJobs200ResponseAllOfJobsInner) HasSecurityProfile() bool`

HasSecurityProfile returns a boolean if a field has been set.

### GetCustomConfig

`func (o *ListJobs200ResponseAllOfJobsInner) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *ListJobs200ResponseAllOfJobsInner) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *ListJobs200ResponseAllOfJobsInner) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetCustomOptions

`func (o *ListJobs200ResponseAllOfJobsInner) GetCustomOptions() ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *ListJobs200ResponseAllOfJobsInner) GetCustomOptionsOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *ListJobs200ResponseAllOfJobsInner) SetCustomOptions(v ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *ListJobs200ResponseAllOfJobsInner) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



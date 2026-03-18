# GetJobs200ResponseAllOfJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**GetJobs200ResponseAllOfJobAnyOf2Type**](GetJobs200ResponseAllOfJobAnyOf2Type.md) |  | [optional] 
**Workflow** | Pointer to [**GetJobs200ResponseAllOfJobAnyOf2Workflow**](GetJobs200ResponseAllOfJobAnyOf2Workflow.md) |  | [optional] 
**Task** | Pointer to [**GetJobs200ResponseAllOfJobAnyOf1Task**](GetJobs200ResponseAllOfJobAnyOf1Task.md) |  | [optional] 
**SecurityPackage** | Pointer to [**GetJobs200ResponseAllOfJobAnyOfSecurityPackage**](GetJobs200ResponseAllOfJobAnyOfSecurityPackage.md) |  | [optional] 
**JobSummary** | Pointer to **string** |  | [optional] 
**ScheduleMode** | Pointer to [**GetJobs200ResponseAllOfJobAnyOf2ScheduleMode**](GetJobs200ResponseAllOfJobAnyOf2ScheduleMode.md) |  | [optional] 
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
**CreatedBy** | Pointer to [**GetJobs200ResponseAllOfJobAnyOf2CreatedBy**](GetJobs200ResponseAllOfJobAnyOf2CreatedBy.md) |  | [optional] 
**TargetType** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to [**[]GetJobs200ResponseAllOfJobAnyOf2TargetsInner**](GetJobs200ResponseAllOfJobAnyOf2TargetsInner.md) |  | [optional] 
**ScanPath** | Pointer to **string** | Scan Checklist. Only applies to type scap-package. | [optional] 
**SecurityProfile** | Pointer to **string** | Security Profile. Only applies to type scap-package. | [optional] 
**CustomConfig** | Pointer to **string** |  | [optional] 
**CustomOptions** | Pointer to [**GetJobs200ResponseAllOfJobAnyOf2CustomOptions**](GetJobs200ResponseAllOfJobAnyOf2CustomOptions.md) |  | [optional] 

## Methods

### NewGetJobs200ResponseAllOfJob

`func NewGetJobs200ResponseAllOfJob() *GetJobs200ResponseAllOfJob`

NewGetJobs200ResponseAllOfJob instantiates a new GetJobs200ResponseAllOfJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJobs200ResponseAllOfJobWithDefaults

`func NewGetJobs200ResponseAllOfJobWithDefaults() *GetJobs200ResponseAllOfJob`

NewGetJobs200ResponseAllOfJobWithDefaults instantiates a new GetJobs200ResponseAllOfJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetJobs200ResponseAllOfJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetJobs200ResponseAllOfJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetJobs200ResponseAllOfJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetJobs200ResponseAllOfJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetJobs200ResponseAllOfJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetJobs200ResponseAllOfJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetJobs200ResponseAllOfJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetJobs200ResponseAllOfJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetJobs200ResponseAllOfJob) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetJobs200ResponseAllOfJob) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetJobs200ResponseAllOfJob) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetJobs200ResponseAllOfJob) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *GetJobs200ResponseAllOfJob) GetType() GetJobs200ResponseAllOfJobAnyOf2Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetJobs200ResponseAllOfJob) GetTypeOk() (*GetJobs200ResponseAllOfJobAnyOf2Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetJobs200ResponseAllOfJob) SetType(v GetJobs200ResponseAllOfJobAnyOf2Type)`

SetType sets Type field to given value.

### HasType

`func (o *GetJobs200ResponseAllOfJob) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkflow

`func (o *GetJobs200ResponseAllOfJob) GetWorkflow() GetJobs200ResponseAllOfJobAnyOf2Workflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *GetJobs200ResponseAllOfJob) GetWorkflowOk() (*GetJobs200ResponseAllOfJobAnyOf2Workflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *GetJobs200ResponseAllOfJob) SetWorkflow(v GetJobs200ResponseAllOfJobAnyOf2Workflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *GetJobs200ResponseAllOfJob) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetTask

`func (o *GetJobs200ResponseAllOfJob) GetTask() GetJobs200ResponseAllOfJobAnyOf1Task`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *GetJobs200ResponseAllOfJob) GetTaskOk() (*GetJobs200ResponseAllOfJobAnyOf1Task, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *GetJobs200ResponseAllOfJob) SetTask(v GetJobs200ResponseAllOfJobAnyOf1Task)`

SetTask sets Task field to given value.

### HasTask

`func (o *GetJobs200ResponseAllOfJob) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetSecurityPackage

`func (o *GetJobs200ResponseAllOfJob) GetSecurityPackage() GetJobs200ResponseAllOfJobAnyOfSecurityPackage`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *GetJobs200ResponseAllOfJob) GetSecurityPackageOk() (*GetJobs200ResponseAllOfJobAnyOfSecurityPackage, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *GetJobs200ResponseAllOfJob) SetSecurityPackage(v GetJobs200ResponseAllOfJobAnyOfSecurityPackage)`

SetSecurityPackage sets SecurityPackage field to given value.

### HasSecurityPackage

`func (o *GetJobs200ResponseAllOfJob) HasSecurityPackage() bool`

HasSecurityPackage returns a boolean if a field has been set.

### GetJobSummary

`func (o *GetJobs200ResponseAllOfJob) GetJobSummary() string`

GetJobSummary returns the JobSummary field if non-nil, zero value otherwise.

### GetJobSummaryOk

`func (o *GetJobs200ResponseAllOfJob) GetJobSummaryOk() (*string, bool)`

GetJobSummaryOk returns a tuple with the JobSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSummary

`func (o *GetJobs200ResponseAllOfJob) SetJobSummary(v string)`

SetJobSummary sets JobSummary field to given value.

### HasJobSummary

`func (o *GetJobs200ResponseAllOfJob) HasJobSummary() bool`

HasJobSummary returns a boolean if a field has been set.

### GetScheduleMode

`func (o *GetJobs200ResponseAllOfJob) GetScheduleMode() GetJobs200ResponseAllOfJobAnyOf2ScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *GetJobs200ResponseAllOfJob) GetScheduleModeOk() (*GetJobs200ResponseAllOfJobAnyOf2ScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *GetJobs200ResponseAllOfJob) SetScheduleMode(v GetJobs200ResponseAllOfJobAnyOf2ScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *GetJobs200ResponseAllOfJob) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetDateTime

`func (o *GetJobs200ResponseAllOfJob) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *GetJobs200ResponseAllOfJob) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *GetJobs200ResponseAllOfJob) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *GetJobs200ResponseAllOfJob) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *GetJobs200ResponseAllOfJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetJobs200ResponseAllOfJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetJobs200ResponseAllOfJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetJobs200ResponseAllOfJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNamespace

`func (o *GetJobs200ResponseAllOfJob) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GetJobs200ResponseAllOfJob) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GetJobs200ResponseAllOfJob) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GetJobs200ResponseAllOfJob) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCategory

`func (o *GetJobs200ResponseAllOfJob) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetJobs200ResponseAllOfJob) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetJobs200ResponseAllOfJob) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetJobs200ResponseAllOfJob) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *GetJobs200ResponseAllOfJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetJobs200ResponseAllOfJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetJobs200ResponseAllOfJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetJobs200ResponseAllOfJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetJobs200ResponseAllOfJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetJobs200ResponseAllOfJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetJobs200ResponseAllOfJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetJobs200ResponseAllOfJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetJobs200ResponseAllOfJob) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetJobs200ResponseAllOfJob) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetJobs200ResponseAllOfJob) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetJobs200ResponseAllOfJob) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetJobs200ResponseAllOfJob) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetJobs200ResponseAllOfJob) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetJobs200ResponseAllOfJob) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetJobs200ResponseAllOfJob) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastRun

`func (o *GetJobs200ResponseAllOfJob) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *GetJobs200ResponseAllOfJob) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *GetJobs200ResponseAllOfJob) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *GetJobs200ResponseAllOfJob) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastResult

`func (o *GetJobs200ResponseAllOfJob) GetLastResult() string`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *GetJobs200ResponseAllOfJob) GetLastResultOk() (*string, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *GetJobs200ResponseAllOfJob) SetLastResult(v string)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *GetJobs200ResponseAllOfJob) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetJobs200ResponseAllOfJob) GetCreatedBy() GetJobs200ResponseAllOfJobAnyOf2CreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetJobs200ResponseAllOfJob) GetCreatedByOk() (*GetJobs200ResponseAllOfJobAnyOf2CreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetJobs200ResponseAllOfJob) SetCreatedBy(v GetJobs200ResponseAllOfJobAnyOf2CreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetJobs200ResponseAllOfJob) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTargetType

`func (o *GetJobs200ResponseAllOfJob) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *GetJobs200ResponseAllOfJob) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *GetJobs200ResponseAllOfJob) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *GetJobs200ResponseAllOfJob) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargets

`func (o *GetJobs200ResponseAllOfJob) GetTargets() []GetJobs200ResponseAllOfJobAnyOf2TargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *GetJobs200ResponseAllOfJob) GetTargetsOk() (*[]GetJobs200ResponseAllOfJobAnyOf2TargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *GetJobs200ResponseAllOfJob) SetTargets(v []GetJobs200ResponseAllOfJobAnyOf2TargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *GetJobs200ResponseAllOfJob) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetScanPath

`func (o *GetJobs200ResponseAllOfJob) GetScanPath() string`

GetScanPath returns the ScanPath field if non-nil, zero value otherwise.

### GetScanPathOk

`func (o *GetJobs200ResponseAllOfJob) GetScanPathOk() (*string, bool)`

GetScanPathOk returns a tuple with the ScanPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPath

`func (o *GetJobs200ResponseAllOfJob) SetScanPath(v string)`

SetScanPath sets ScanPath field to given value.

### HasScanPath

`func (o *GetJobs200ResponseAllOfJob) HasScanPath() bool`

HasScanPath returns a boolean if a field has been set.

### GetSecurityProfile

`func (o *GetJobs200ResponseAllOfJob) GetSecurityProfile() string`

GetSecurityProfile returns the SecurityProfile field if non-nil, zero value otherwise.

### GetSecurityProfileOk

`func (o *GetJobs200ResponseAllOfJob) GetSecurityProfileOk() (*string, bool)`

GetSecurityProfileOk returns a tuple with the SecurityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProfile

`func (o *GetJobs200ResponseAllOfJob) SetSecurityProfile(v string)`

SetSecurityProfile sets SecurityProfile field to given value.

### HasSecurityProfile

`func (o *GetJobs200ResponseAllOfJob) HasSecurityProfile() bool`

HasSecurityProfile returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetJobs200ResponseAllOfJob) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetJobs200ResponseAllOfJob) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetJobs200ResponseAllOfJob) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetJobs200ResponseAllOfJob) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetCustomOptions

`func (o *GetJobs200ResponseAllOfJob) GetCustomOptions() GetJobs200ResponseAllOfJobAnyOf2CustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *GetJobs200ResponseAllOfJob) GetCustomOptionsOk() (*GetJobs200ResponseAllOfJobAnyOf2CustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *GetJobs200ResponseAllOfJob) SetCustomOptions(v GetJobs200ResponseAllOfJobAnyOf2CustomOptions)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *GetJobs200ResponseAllOfJob) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



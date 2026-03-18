# AddJobs200ResponseAllOfJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**AddJobs200ResponseAllOfJobType**](AddJobs200ResponseAllOfJobType.md) |  | [optional] 
**Workflow** | Pointer to [**AddJobs200ResponseAllOfJobWorkflow**](AddJobs200ResponseAllOfJobWorkflow.md) |  | [optional] 
**Task** | Pointer to [**AddJobs200ResponseAllOfJobTask**](AddJobs200ResponseAllOfJobTask.md) |  | [optional] 
**SecurityPackage** | Pointer to [**AddJobs200ResponseAllOfJobSecurityPackage**](AddJobs200ResponseAllOfJobSecurityPackage.md) |  | [optional] 
**JobSummary** | Pointer to **NullableString** |  | [optional] 
**ScheduleMode** | Pointer to [**AddJobs200ResponseAllOfJobScheduleMode**](AddJobs200ResponseAllOfJobScheduleMode.md) |  | [optional] 
**DateTime** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastRun** | Pointer to **time.Time** |  | [optional] 
**LastResult** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**AddJobs200ResponseAllOfJobCreatedBy**](AddJobs200ResponseAllOfJobCreatedBy.md) |  | [optional] 
**TargetType** | Pointer to **NullableString** |  | [optional] 
**Targets** | Pointer to [**[]AddJobs200ResponseAllOfJobTargetsInner**](AddJobs200ResponseAllOfJobTargetsInner.md) |  | [optional] 
**ScanPath** | Pointer to **NullableString** | Scan Checklist. Only applies to type scap-package. | [optional] 
**SecurityProfile** | Pointer to **NullableString** | Security Profile. Only applies to type scap-package. | [optional] 
**CustomConfig** | Pointer to **NullableString** |  | [optional] 
**CustomOptions** | Pointer to [**AddJobs200ResponseAllOfJobCustomOptions**](AddJobs200ResponseAllOfJobCustomOptions.md) |  | [optional] 

## Methods

### NewAddJobs200ResponseAllOfJob

`func NewAddJobs200ResponseAllOfJob() *AddJobs200ResponseAllOfJob`

NewAddJobs200ResponseAllOfJob instantiates a new AddJobs200ResponseAllOfJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddJobs200ResponseAllOfJobWithDefaults

`func NewAddJobs200ResponseAllOfJobWithDefaults() *AddJobs200ResponseAllOfJob`

NewAddJobs200ResponseAllOfJobWithDefaults instantiates a new AddJobs200ResponseAllOfJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddJobs200ResponseAllOfJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddJobs200ResponseAllOfJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddJobs200ResponseAllOfJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddJobs200ResponseAllOfJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddJobs200ResponseAllOfJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddJobs200ResponseAllOfJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddJobs200ResponseAllOfJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddJobs200ResponseAllOfJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *AddJobs200ResponseAllOfJob) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddJobs200ResponseAllOfJob) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddJobs200ResponseAllOfJob) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddJobs200ResponseAllOfJob) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AddJobs200ResponseAllOfJob) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AddJobs200ResponseAllOfJob) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *AddJobs200ResponseAllOfJob) GetType() AddJobs200ResponseAllOfJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddJobs200ResponseAllOfJob) GetTypeOk() (*AddJobs200ResponseAllOfJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddJobs200ResponseAllOfJob) SetType(v AddJobs200ResponseAllOfJobType)`

SetType sets Type field to given value.

### HasType

`func (o *AddJobs200ResponseAllOfJob) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkflow

`func (o *AddJobs200ResponseAllOfJob) GetWorkflow() AddJobs200ResponseAllOfJobWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *AddJobs200ResponseAllOfJob) GetWorkflowOk() (*AddJobs200ResponseAllOfJobWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *AddJobs200ResponseAllOfJob) SetWorkflow(v AddJobs200ResponseAllOfJobWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *AddJobs200ResponseAllOfJob) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetTask

`func (o *AddJobs200ResponseAllOfJob) GetTask() AddJobs200ResponseAllOfJobTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *AddJobs200ResponseAllOfJob) GetTaskOk() (*AddJobs200ResponseAllOfJobTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *AddJobs200ResponseAllOfJob) SetTask(v AddJobs200ResponseAllOfJobTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *AddJobs200ResponseAllOfJob) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetSecurityPackage

`func (o *AddJobs200ResponseAllOfJob) GetSecurityPackage() AddJobs200ResponseAllOfJobSecurityPackage`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *AddJobs200ResponseAllOfJob) GetSecurityPackageOk() (*AddJobs200ResponseAllOfJobSecurityPackage, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *AddJobs200ResponseAllOfJob) SetSecurityPackage(v AddJobs200ResponseAllOfJobSecurityPackage)`

SetSecurityPackage sets SecurityPackage field to given value.

### HasSecurityPackage

`func (o *AddJobs200ResponseAllOfJob) HasSecurityPackage() bool`

HasSecurityPackage returns a boolean if a field has been set.

### GetJobSummary

`func (o *AddJobs200ResponseAllOfJob) GetJobSummary() string`

GetJobSummary returns the JobSummary field if non-nil, zero value otherwise.

### GetJobSummaryOk

`func (o *AddJobs200ResponseAllOfJob) GetJobSummaryOk() (*string, bool)`

GetJobSummaryOk returns a tuple with the JobSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSummary

`func (o *AddJobs200ResponseAllOfJob) SetJobSummary(v string)`

SetJobSummary sets JobSummary field to given value.

### HasJobSummary

`func (o *AddJobs200ResponseAllOfJob) HasJobSummary() bool`

HasJobSummary returns a boolean if a field has been set.

### SetJobSummaryNil

`func (o *AddJobs200ResponseAllOfJob) SetJobSummaryNil(b bool)`

 SetJobSummaryNil sets the value for JobSummary to be an explicit nil

### UnsetJobSummary
`func (o *AddJobs200ResponseAllOfJob) UnsetJobSummary()`

UnsetJobSummary ensures that no value is present for JobSummary, not even an explicit nil
### GetScheduleMode

`func (o *AddJobs200ResponseAllOfJob) GetScheduleMode() AddJobs200ResponseAllOfJobScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *AddJobs200ResponseAllOfJob) GetScheduleModeOk() (*AddJobs200ResponseAllOfJobScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *AddJobs200ResponseAllOfJob) SetScheduleMode(v AddJobs200ResponseAllOfJobScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *AddJobs200ResponseAllOfJob) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetDateTime

`func (o *AddJobs200ResponseAllOfJob) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *AddJobs200ResponseAllOfJob) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *AddJobs200ResponseAllOfJob) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *AddJobs200ResponseAllOfJob) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTimeNil

`func (o *AddJobs200ResponseAllOfJob) SetDateTimeNil(b bool)`

 SetDateTimeNil sets the value for DateTime to be an explicit nil

### UnsetDateTime
`func (o *AddJobs200ResponseAllOfJob) UnsetDateTime()`

UnsetDateTime ensures that no value is present for DateTime, not even an explicit nil
### GetStatus

`func (o *AddJobs200ResponseAllOfJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddJobs200ResponseAllOfJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddJobs200ResponseAllOfJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddJobs200ResponseAllOfJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AddJobs200ResponseAllOfJob) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AddJobs200ResponseAllOfJob) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNamespace

`func (o *AddJobs200ResponseAllOfJob) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AddJobs200ResponseAllOfJob) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AddJobs200ResponseAllOfJob) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AddJobs200ResponseAllOfJob) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AddJobs200ResponseAllOfJob) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AddJobs200ResponseAllOfJob) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetCategory

`func (o *AddJobs200ResponseAllOfJob) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AddJobs200ResponseAllOfJob) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AddJobs200ResponseAllOfJob) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AddJobs200ResponseAllOfJob) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *AddJobs200ResponseAllOfJob) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *AddJobs200ResponseAllOfJob) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *AddJobs200ResponseAllOfJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddJobs200ResponseAllOfJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddJobs200ResponseAllOfJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddJobs200ResponseAllOfJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddJobs200ResponseAllOfJob) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddJobs200ResponseAllOfJob) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *AddJobs200ResponseAllOfJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddJobs200ResponseAllOfJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddJobs200ResponseAllOfJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddJobs200ResponseAllOfJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddJobs200ResponseAllOfJob) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddJobs200ResponseAllOfJob) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddJobs200ResponseAllOfJob) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddJobs200ResponseAllOfJob) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddJobs200ResponseAllOfJob) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddJobs200ResponseAllOfJob) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddJobs200ResponseAllOfJob) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddJobs200ResponseAllOfJob) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastRun

`func (o *AddJobs200ResponseAllOfJob) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *AddJobs200ResponseAllOfJob) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *AddJobs200ResponseAllOfJob) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *AddJobs200ResponseAllOfJob) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastResult

`func (o *AddJobs200ResponseAllOfJob) GetLastResult() string`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *AddJobs200ResponseAllOfJob) GetLastResultOk() (*string, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *AddJobs200ResponseAllOfJob) SetLastResult(v string)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *AddJobs200ResponseAllOfJob) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### SetLastResultNil

`func (o *AddJobs200ResponseAllOfJob) SetLastResultNil(b bool)`

 SetLastResultNil sets the value for LastResult to be an explicit nil

### UnsetLastResult
`func (o *AddJobs200ResponseAllOfJob) UnsetLastResult()`

UnsetLastResult ensures that no value is present for LastResult, not even an explicit nil
### GetCreatedBy

`func (o *AddJobs200ResponseAllOfJob) GetCreatedBy() AddJobs200ResponseAllOfJobCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AddJobs200ResponseAllOfJob) GetCreatedByOk() (*AddJobs200ResponseAllOfJobCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AddJobs200ResponseAllOfJob) SetCreatedBy(v AddJobs200ResponseAllOfJobCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AddJobs200ResponseAllOfJob) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTargetType

`func (o *AddJobs200ResponseAllOfJob) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AddJobs200ResponseAllOfJob) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AddJobs200ResponseAllOfJob) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *AddJobs200ResponseAllOfJob) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *AddJobs200ResponseAllOfJob) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *AddJobs200ResponseAllOfJob) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTargets

`func (o *AddJobs200ResponseAllOfJob) GetTargets() []AddJobs200ResponseAllOfJobTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AddJobs200ResponseAllOfJob) GetTargetsOk() (*[]AddJobs200ResponseAllOfJobTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AddJobs200ResponseAllOfJob) SetTargets(v []AddJobs200ResponseAllOfJobTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *AddJobs200ResponseAllOfJob) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *AddJobs200ResponseAllOfJob) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *AddJobs200ResponseAllOfJob) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetScanPath

`func (o *AddJobs200ResponseAllOfJob) GetScanPath() string`

GetScanPath returns the ScanPath field if non-nil, zero value otherwise.

### GetScanPathOk

`func (o *AddJobs200ResponseAllOfJob) GetScanPathOk() (*string, bool)`

GetScanPathOk returns a tuple with the ScanPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPath

`func (o *AddJobs200ResponseAllOfJob) SetScanPath(v string)`

SetScanPath sets ScanPath field to given value.

### HasScanPath

`func (o *AddJobs200ResponseAllOfJob) HasScanPath() bool`

HasScanPath returns a boolean if a field has been set.

### SetScanPathNil

`func (o *AddJobs200ResponseAllOfJob) SetScanPathNil(b bool)`

 SetScanPathNil sets the value for ScanPath to be an explicit nil

### UnsetScanPath
`func (o *AddJobs200ResponseAllOfJob) UnsetScanPath()`

UnsetScanPath ensures that no value is present for ScanPath, not even an explicit nil
### GetSecurityProfile

`func (o *AddJobs200ResponseAllOfJob) GetSecurityProfile() string`

GetSecurityProfile returns the SecurityProfile field if non-nil, zero value otherwise.

### GetSecurityProfileOk

`func (o *AddJobs200ResponseAllOfJob) GetSecurityProfileOk() (*string, bool)`

GetSecurityProfileOk returns a tuple with the SecurityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProfile

`func (o *AddJobs200ResponseAllOfJob) SetSecurityProfile(v string)`

SetSecurityProfile sets SecurityProfile field to given value.

### HasSecurityProfile

`func (o *AddJobs200ResponseAllOfJob) HasSecurityProfile() bool`

HasSecurityProfile returns a boolean if a field has been set.

### SetSecurityProfileNil

`func (o *AddJobs200ResponseAllOfJob) SetSecurityProfileNil(b bool)`

 SetSecurityProfileNil sets the value for SecurityProfile to be an explicit nil

### UnsetSecurityProfile
`func (o *AddJobs200ResponseAllOfJob) UnsetSecurityProfile()`

UnsetSecurityProfile ensures that no value is present for SecurityProfile, not even an explicit nil
### GetCustomConfig

`func (o *AddJobs200ResponseAllOfJob) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *AddJobs200ResponseAllOfJob) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *AddJobs200ResponseAllOfJob) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *AddJobs200ResponseAllOfJob) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### SetCustomConfigNil

`func (o *AddJobs200ResponseAllOfJob) SetCustomConfigNil(b bool)`

 SetCustomConfigNil sets the value for CustomConfig to be an explicit nil

### UnsetCustomConfig
`func (o *AddJobs200ResponseAllOfJob) UnsetCustomConfig()`

UnsetCustomConfig ensures that no value is present for CustomConfig, not even an explicit nil
### GetCustomOptions

`func (o *AddJobs200ResponseAllOfJob) GetCustomOptions() AddJobs200ResponseAllOfJobCustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *AddJobs200ResponseAllOfJob) GetCustomOptionsOk() (*AddJobs200ResponseAllOfJobCustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *AddJobs200ResponseAllOfJob) SetCustomOptions(v AddJobs200ResponseAllOfJobCustomOptions)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *AddJobs200ResponseAllOfJob) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



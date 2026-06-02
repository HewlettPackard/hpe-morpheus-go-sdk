# UpdateJobs200ResponseAllOfJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**UpdateJobs200ResponseAllOfJobType**](UpdateJobs200ResponseAllOfJobType.md) |  | [optional] 
**Workflow** | Pointer to [**UpdateJobs200ResponseAllOfJobWorkflow**](UpdateJobs200ResponseAllOfJobWorkflow.md) |  | [optional] 
**Task** | Pointer to [**UpdateJobs200ResponseAllOfJobTask**](UpdateJobs200ResponseAllOfJobTask.md) |  | [optional] 
**SecurityPackage** | Pointer to [**UpdateJobs200ResponseAllOfJobSecurityPackage**](UpdateJobs200ResponseAllOfJobSecurityPackage.md) |  | [optional] 
**JobSummary** | Pointer to **NullableString** |  | [optional] 
**ScheduleMode** | Pointer to [**UpdateJobs200ResponseAllOfJobScheduleMode**](UpdateJobs200ResponseAllOfJobScheduleMode.md) |  | [optional] 
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
**CreatedBy** | Pointer to [**UpdateJobs200ResponseAllOfJobCreatedBy**](UpdateJobs200ResponseAllOfJobCreatedBy.md) |  | [optional] 
**TargetType** | Pointer to **NullableString** |  | [optional] 
**Targets** | Pointer to [**[]UpdateJobs200ResponseAllOfJobTargetsInner**](UpdateJobs200ResponseAllOfJobTargetsInner.md) |  | [optional] 
**ScanPath** | Pointer to **NullableString** | Scan Checklist. Only applies to type scap-package. | [optional] 
**SecurityProfile** | Pointer to **NullableString** | Security Profile. Only applies to type scap-package. | [optional] 
**CustomConfig** | Pointer to **NullableString** |  | [optional] 
**CustomOptions** | Pointer to [**UpdateJobs200ResponseAllOfJobCustomOptions**](UpdateJobs200ResponseAllOfJobCustomOptions.md) |  | [optional] 

## Methods

### NewUpdateJobs200ResponseAllOfJob

`func NewUpdateJobs200ResponseAllOfJob() *UpdateJobs200ResponseAllOfJob`

NewUpdateJobs200ResponseAllOfJob instantiates a new UpdateJobs200ResponseAllOfJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateJobs200ResponseAllOfJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateJobs200ResponseAllOfJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateJobs200ResponseAllOfJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateJobs200ResponseAllOfJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateJobs200ResponseAllOfJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateJobs200ResponseAllOfJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateJobs200ResponseAllOfJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateJobs200ResponseAllOfJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateJobs200ResponseAllOfJob) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateJobs200ResponseAllOfJob) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateJobs200ResponseAllOfJob) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateJobs200ResponseAllOfJob) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateJobs200ResponseAllOfJob) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateJobs200ResponseAllOfJob) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *UpdateJobs200ResponseAllOfJob) GetType() UpdateJobs200ResponseAllOfJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateJobs200ResponseAllOfJob) GetTypeOk() (*UpdateJobs200ResponseAllOfJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateJobs200ResponseAllOfJob) SetType(v UpdateJobs200ResponseAllOfJobType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateJobs200ResponseAllOfJob) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkflow

`func (o *UpdateJobs200ResponseAllOfJob) GetWorkflow() UpdateJobs200ResponseAllOfJobWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *UpdateJobs200ResponseAllOfJob) GetWorkflowOk() (*UpdateJobs200ResponseAllOfJobWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *UpdateJobs200ResponseAllOfJob) SetWorkflow(v UpdateJobs200ResponseAllOfJobWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *UpdateJobs200ResponseAllOfJob) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetTask

`func (o *UpdateJobs200ResponseAllOfJob) GetTask() UpdateJobs200ResponseAllOfJobTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *UpdateJobs200ResponseAllOfJob) GetTaskOk() (*UpdateJobs200ResponseAllOfJobTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *UpdateJobs200ResponseAllOfJob) SetTask(v UpdateJobs200ResponseAllOfJobTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *UpdateJobs200ResponseAllOfJob) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetSecurityPackage

`func (o *UpdateJobs200ResponseAllOfJob) GetSecurityPackage() UpdateJobs200ResponseAllOfJobSecurityPackage`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *UpdateJobs200ResponseAllOfJob) GetSecurityPackageOk() (*UpdateJobs200ResponseAllOfJobSecurityPackage, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *UpdateJobs200ResponseAllOfJob) SetSecurityPackage(v UpdateJobs200ResponseAllOfJobSecurityPackage)`

SetSecurityPackage sets SecurityPackage field to given value.

### HasSecurityPackage

`func (o *UpdateJobs200ResponseAllOfJob) HasSecurityPackage() bool`

HasSecurityPackage returns a boolean if a field has been set.

### GetJobSummary

`func (o *UpdateJobs200ResponseAllOfJob) GetJobSummary() string`

GetJobSummary returns the JobSummary field if non-nil, zero value otherwise.

### GetJobSummaryOk

`func (o *UpdateJobs200ResponseAllOfJob) GetJobSummaryOk() (*string, bool)`

GetJobSummaryOk returns a tuple with the JobSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSummary

`func (o *UpdateJobs200ResponseAllOfJob) SetJobSummary(v string)`

SetJobSummary sets JobSummary field to given value.

### HasJobSummary

`func (o *UpdateJobs200ResponseAllOfJob) HasJobSummary() bool`

HasJobSummary returns a boolean if a field has been set.

### SetJobSummaryNil

`func (o *UpdateJobs200ResponseAllOfJob) SetJobSummaryNil(b bool)`

 SetJobSummaryNil sets the value for JobSummary to be an explicit nil

### UnsetJobSummary
`func (o *UpdateJobs200ResponseAllOfJob) UnsetJobSummary()`

UnsetJobSummary ensures that no value is present for JobSummary, not even an explicit nil
### GetScheduleMode

`func (o *UpdateJobs200ResponseAllOfJob) GetScheduleMode() UpdateJobs200ResponseAllOfJobScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *UpdateJobs200ResponseAllOfJob) GetScheduleModeOk() (*UpdateJobs200ResponseAllOfJobScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *UpdateJobs200ResponseAllOfJob) SetScheduleMode(v UpdateJobs200ResponseAllOfJobScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *UpdateJobs200ResponseAllOfJob) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetDateTime

`func (o *UpdateJobs200ResponseAllOfJob) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *UpdateJobs200ResponseAllOfJob) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *UpdateJobs200ResponseAllOfJob) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *UpdateJobs200ResponseAllOfJob) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTimeNil

`func (o *UpdateJobs200ResponseAllOfJob) SetDateTimeNil(b bool)`

 SetDateTimeNil sets the value for DateTime to be an explicit nil

### UnsetDateTime
`func (o *UpdateJobs200ResponseAllOfJob) UnsetDateTime()`

UnsetDateTime ensures that no value is present for DateTime, not even an explicit nil
### GetStatus

`func (o *UpdateJobs200ResponseAllOfJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateJobs200ResponseAllOfJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateJobs200ResponseAllOfJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateJobs200ResponseAllOfJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UpdateJobs200ResponseAllOfJob) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UpdateJobs200ResponseAllOfJob) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNamespace

`func (o *UpdateJobs200ResponseAllOfJob) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UpdateJobs200ResponseAllOfJob) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UpdateJobs200ResponseAllOfJob) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UpdateJobs200ResponseAllOfJob) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *UpdateJobs200ResponseAllOfJob) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *UpdateJobs200ResponseAllOfJob) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetCategory

`func (o *UpdateJobs200ResponseAllOfJob) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateJobs200ResponseAllOfJob) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateJobs200ResponseAllOfJob) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateJobs200ResponseAllOfJob) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UpdateJobs200ResponseAllOfJob) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UpdateJobs200ResponseAllOfJob) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *UpdateJobs200ResponseAllOfJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateJobs200ResponseAllOfJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateJobs200ResponseAllOfJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateJobs200ResponseAllOfJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateJobs200ResponseAllOfJob) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateJobs200ResponseAllOfJob) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *UpdateJobs200ResponseAllOfJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateJobs200ResponseAllOfJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateJobs200ResponseAllOfJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateJobs200ResponseAllOfJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateJobs200ResponseAllOfJob) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateJobs200ResponseAllOfJob) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateJobs200ResponseAllOfJob) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateJobs200ResponseAllOfJob) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateJobs200ResponseAllOfJob) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateJobs200ResponseAllOfJob) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateJobs200ResponseAllOfJob) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateJobs200ResponseAllOfJob) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastRun

`func (o *UpdateJobs200ResponseAllOfJob) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *UpdateJobs200ResponseAllOfJob) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *UpdateJobs200ResponseAllOfJob) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *UpdateJobs200ResponseAllOfJob) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastResult

`func (o *UpdateJobs200ResponseAllOfJob) GetLastResult() string`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *UpdateJobs200ResponseAllOfJob) GetLastResultOk() (*string, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *UpdateJobs200ResponseAllOfJob) SetLastResult(v string)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *UpdateJobs200ResponseAllOfJob) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### SetLastResultNil

`func (o *UpdateJobs200ResponseAllOfJob) SetLastResultNil(b bool)`

 SetLastResultNil sets the value for LastResult to be an explicit nil

### UnsetLastResult
`func (o *UpdateJobs200ResponseAllOfJob) UnsetLastResult()`

UnsetLastResult ensures that no value is present for LastResult, not even an explicit nil
### GetCreatedBy

`func (o *UpdateJobs200ResponseAllOfJob) GetCreatedBy() UpdateJobs200ResponseAllOfJobCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateJobs200ResponseAllOfJob) GetCreatedByOk() (*UpdateJobs200ResponseAllOfJobCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateJobs200ResponseAllOfJob) SetCreatedBy(v UpdateJobs200ResponseAllOfJobCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateJobs200ResponseAllOfJob) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTargetType

`func (o *UpdateJobs200ResponseAllOfJob) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *UpdateJobs200ResponseAllOfJob) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *UpdateJobs200ResponseAllOfJob) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *UpdateJobs200ResponseAllOfJob) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *UpdateJobs200ResponseAllOfJob) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *UpdateJobs200ResponseAllOfJob) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTargets

`func (o *UpdateJobs200ResponseAllOfJob) GetTargets() []UpdateJobs200ResponseAllOfJobTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *UpdateJobs200ResponseAllOfJob) GetTargetsOk() (*[]UpdateJobs200ResponseAllOfJobTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *UpdateJobs200ResponseAllOfJob) SetTargets(v []UpdateJobs200ResponseAllOfJobTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *UpdateJobs200ResponseAllOfJob) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *UpdateJobs200ResponseAllOfJob) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *UpdateJobs200ResponseAllOfJob) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetScanPath

`func (o *UpdateJobs200ResponseAllOfJob) GetScanPath() string`

GetScanPath returns the ScanPath field if non-nil, zero value otherwise.

### GetScanPathOk

`func (o *UpdateJobs200ResponseAllOfJob) GetScanPathOk() (*string, bool)`

GetScanPathOk returns a tuple with the ScanPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPath

`func (o *UpdateJobs200ResponseAllOfJob) SetScanPath(v string)`

SetScanPath sets ScanPath field to given value.

### HasScanPath

`func (o *UpdateJobs200ResponseAllOfJob) HasScanPath() bool`

HasScanPath returns a boolean if a field has been set.

### SetScanPathNil

`func (o *UpdateJobs200ResponseAllOfJob) SetScanPathNil(b bool)`

 SetScanPathNil sets the value for ScanPath to be an explicit nil

### UnsetScanPath
`func (o *UpdateJobs200ResponseAllOfJob) UnsetScanPath()`

UnsetScanPath ensures that no value is present for ScanPath, not even an explicit nil
### GetSecurityProfile

`func (o *UpdateJobs200ResponseAllOfJob) GetSecurityProfile() string`

GetSecurityProfile returns the SecurityProfile field if non-nil, zero value otherwise.

### GetSecurityProfileOk

`func (o *UpdateJobs200ResponseAllOfJob) GetSecurityProfileOk() (*string, bool)`

GetSecurityProfileOk returns a tuple with the SecurityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProfile

`func (o *UpdateJobs200ResponseAllOfJob) SetSecurityProfile(v string)`

SetSecurityProfile sets SecurityProfile field to given value.

### HasSecurityProfile

`func (o *UpdateJobs200ResponseAllOfJob) HasSecurityProfile() bool`

HasSecurityProfile returns a boolean if a field has been set.

### SetSecurityProfileNil

`func (o *UpdateJobs200ResponseAllOfJob) SetSecurityProfileNil(b bool)`

 SetSecurityProfileNil sets the value for SecurityProfile to be an explicit nil

### UnsetSecurityProfile
`func (o *UpdateJobs200ResponseAllOfJob) UnsetSecurityProfile()`

UnsetSecurityProfile ensures that no value is present for SecurityProfile, not even an explicit nil
### GetCustomConfig

`func (o *UpdateJobs200ResponseAllOfJob) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *UpdateJobs200ResponseAllOfJob) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *UpdateJobs200ResponseAllOfJob) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *UpdateJobs200ResponseAllOfJob) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### SetCustomConfigNil

`func (o *UpdateJobs200ResponseAllOfJob) SetCustomConfigNil(b bool)`

 SetCustomConfigNil sets the value for CustomConfig to be an explicit nil

### UnsetCustomConfig
`func (o *UpdateJobs200ResponseAllOfJob) UnsetCustomConfig()`

UnsetCustomConfig ensures that no value is present for CustomConfig, not even an explicit nil
### GetCustomOptions

`func (o *UpdateJobs200ResponseAllOfJob) GetCustomOptions() UpdateJobs200ResponseAllOfJobCustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *UpdateJobs200ResponseAllOfJob) GetCustomOptionsOk() (*UpdateJobs200ResponseAllOfJobCustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *UpdateJobs200ResponseAllOfJob) SetCustomOptions(v UpdateJobs200ResponseAllOfJobCustomOptions)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *UpdateJobs200ResponseAllOfJob) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



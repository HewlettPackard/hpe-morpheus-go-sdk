# GetJobs200ResponseAllOfJobAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**GetJobs200ResponseAllOfJobAnyOfType**](GetJobs200ResponseAllOfJobAnyOfType.md) |  | [optional] 
**Workflow** | Pointer to [**GetJobs200ResponseAllOfJobAnyOfWorkflow**](GetJobs200ResponseAllOfJobAnyOfWorkflow.md) |  | [optional] 
**Task** | Pointer to [**GetJobs200ResponseAllOfJobAnyOfTask**](GetJobs200ResponseAllOfJobAnyOfTask.md) |  | [optional] 
**SecurityPackage** | Pointer to [**GetJobs200ResponseAllOfJobAnyOfSecurityPackage**](GetJobs200ResponseAllOfJobAnyOfSecurityPackage.md) |  | [optional] 
**JobSummary** | Pointer to **NullableString** |  | [optional] 
**ScheduleMode** | Pointer to [**GetJobs200ResponseAllOfJobAnyOfScheduleMode**](GetJobs200ResponseAllOfJobAnyOfScheduleMode.md) |  | [optional] 
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
**CreatedBy** | Pointer to [**GetJobs200ResponseAllOfJobAnyOfCreatedBy**](GetJobs200ResponseAllOfJobAnyOfCreatedBy.md) |  | [optional] 
**TargetType** | Pointer to **NullableString** |  | [optional] 
**Targets** | Pointer to [**[]GetJobs200ResponseAllOfJobAnyOfTargetsInner**](GetJobs200ResponseAllOfJobAnyOfTargetsInner.md) |  | [optional] 
**ScanPath** | Pointer to **NullableString** | Scan Checklist. Only applies to type scap-package. | [optional] 
**SecurityProfile** | Pointer to **NullableString** | Security Profile. Only applies to type scap-package. | [optional] 
**CustomConfig** | Pointer to **NullableString** |  | [optional] 
**CustomOptions** | Pointer to [**GetJobs200ResponseAllOfJobAnyOfCustomOptions**](GetJobs200ResponseAllOfJobAnyOfCustomOptions.md) |  | [optional] 

## Methods

### NewGetJobs200ResponseAllOfJobAnyOf

`func NewGetJobs200ResponseAllOfJobAnyOf() *GetJobs200ResponseAllOfJobAnyOf`

NewGetJobs200ResponseAllOfJobAnyOf instantiates a new GetJobs200ResponseAllOfJobAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetType() GetJobs200ResponseAllOfJobAnyOfType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetTypeOk() (*GetJobs200ResponseAllOfJobAnyOfType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetType(v GetJobs200ResponseAllOfJobAnyOfType)`

SetType sets Type field to given value.

### HasType

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkflow

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetWorkflow() GetJobs200ResponseAllOfJobAnyOfWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetWorkflowOk() (*GetJobs200ResponseAllOfJobAnyOfWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetWorkflow(v GetJobs200ResponseAllOfJobAnyOfWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetTask

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetTask() GetJobs200ResponseAllOfJobAnyOfTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetTaskOk() (*GetJobs200ResponseAllOfJobAnyOfTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetTask(v GetJobs200ResponseAllOfJobAnyOfTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetSecurityPackage

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetSecurityPackage() GetJobs200ResponseAllOfJobAnyOfSecurityPackage`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetSecurityPackageOk() (*GetJobs200ResponseAllOfJobAnyOfSecurityPackage, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetSecurityPackage(v GetJobs200ResponseAllOfJobAnyOfSecurityPackage)`

SetSecurityPackage sets SecurityPackage field to given value.

### HasSecurityPackage

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasSecurityPackage() bool`

HasSecurityPackage returns a boolean if a field has been set.

### GetJobSummary

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetJobSummary() string`

GetJobSummary returns the JobSummary field if non-nil, zero value otherwise.

### GetJobSummaryOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetJobSummaryOk() (*string, bool)`

GetJobSummaryOk returns a tuple with the JobSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSummary

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetJobSummary(v string)`

SetJobSummary sets JobSummary field to given value.

### HasJobSummary

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasJobSummary() bool`

HasJobSummary returns a boolean if a field has been set.

### SetJobSummaryNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetJobSummaryNil(b bool)`

 SetJobSummaryNil sets the value for JobSummary to be an explicit nil

### UnsetJobSummary
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetJobSummary()`

UnsetJobSummary ensures that no value is present for JobSummary, not even an explicit nil
### GetScheduleMode

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetScheduleMode() GetJobs200ResponseAllOfJobAnyOfScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetScheduleModeOk() (*GetJobs200ResponseAllOfJobAnyOfScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetScheduleMode(v GetJobs200ResponseAllOfJobAnyOfScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetDateTime

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTimeNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetDateTimeNil(b bool)`

 SetDateTimeNil sets the value for DateTime to be an explicit nil

### UnsetDateTime
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetDateTime()`

UnsetDateTime ensures that no value is present for DateTime, not even an explicit nil
### GetStatus

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNamespace

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetCategory

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastRun

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastResult

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetLastResult() string`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetLastResultOk() (*string, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetLastResult(v string)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### SetLastResultNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetLastResultNil(b bool)`

 SetLastResultNil sets the value for LastResult to be an explicit nil

### UnsetLastResult
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetLastResult()`

UnsetLastResult ensures that no value is present for LastResult, not even an explicit nil
### GetCreatedBy

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetCreatedBy() GetJobs200ResponseAllOfJobAnyOfCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetCreatedByOk() (*GetJobs200ResponseAllOfJobAnyOfCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetCreatedBy(v GetJobs200ResponseAllOfJobAnyOfCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTargetType

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTargets

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetTargets() []GetJobs200ResponseAllOfJobAnyOfTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetTargetsOk() (*[]GetJobs200ResponseAllOfJobAnyOfTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetTargets(v []GetJobs200ResponseAllOfJobAnyOfTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetScanPath

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetScanPath() string`

GetScanPath returns the ScanPath field if non-nil, zero value otherwise.

### GetScanPathOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetScanPathOk() (*string, bool)`

GetScanPathOk returns a tuple with the ScanPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPath

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetScanPath(v string)`

SetScanPath sets ScanPath field to given value.

### HasScanPath

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasScanPath() bool`

HasScanPath returns a boolean if a field has been set.

### SetScanPathNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetScanPathNil(b bool)`

 SetScanPathNil sets the value for ScanPath to be an explicit nil

### UnsetScanPath
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetScanPath()`

UnsetScanPath ensures that no value is present for ScanPath, not even an explicit nil
### GetSecurityProfile

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetSecurityProfile() string`

GetSecurityProfile returns the SecurityProfile field if non-nil, zero value otherwise.

### GetSecurityProfileOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetSecurityProfileOk() (*string, bool)`

GetSecurityProfileOk returns a tuple with the SecurityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProfile

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetSecurityProfile(v string)`

SetSecurityProfile sets SecurityProfile field to given value.

### HasSecurityProfile

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasSecurityProfile() bool`

HasSecurityProfile returns a boolean if a field has been set.

### SetSecurityProfileNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetSecurityProfileNil(b bool)`

 SetSecurityProfileNil sets the value for SecurityProfile to be an explicit nil

### UnsetSecurityProfile
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetSecurityProfile()`

UnsetSecurityProfile ensures that no value is present for SecurityProfile, not even an explicit nil
### GetCustomConfig

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### SetCustomConfigNil

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetCustomConfigNil(b bool)`

 SetCustomConfigNil sets the value for CustomConfig to be an explicit nil

### UnsetCustomConfig
`func (o *GetJobs200ResponseAllOfJobAnyOf) UnsetCustomConfig()`

UnsetCustomConfig ensures that no value is present for CustomConfig, not even an explicit nil
### GetCustomOptions

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetCustomOptions() GetJobs200ResponseAllOfJobAnyOfCustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *GetJobs200ResponseAllOfJobAnyOf) GetCustomOptionsOk() (*GetJobs200ResponseAllOfJobAnyOfCustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *GetJobs200ResponseAllOfJobAnyOf) SetCustomOptions(v GetJobs200ResponseAllOfJobAnyOfCustomOptions)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *GetJobs200ResponseAllOfJobAnyOf) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



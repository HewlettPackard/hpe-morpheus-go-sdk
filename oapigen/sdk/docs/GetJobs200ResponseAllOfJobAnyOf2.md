# GetJobs200ResponseAllOfJobAnyOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**GetJobs200ResponseAllOfJobAnyOf2Type**](GetJobs200ResponseAllOfJobAnyOf2Type.md) |  | [optional] 
**Workflow** | Pointer to [**GetJobs200ResponseAllOfJobAnyOf2Workflow**](GetJobs200ResponseAllOfJobAnyOf2Workflow.md) |  | [optional] 
**JobSummary** | Pointer to **NullableString** |  | [optional] 
**ScheduleMode** | Pointer to [**GetJobs200ResponseAllOfJobAnyOf2ScheduleMode**](GetJobs200ResponseAllOfJobAnyOf2ScheduleMode.md) |  | [optional] 
**DateTime** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastRun** | Pointer to **time.Time** |  | [optional] 
**LastResult** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**GetJobs200ResponseAllOfJobAnyOf2CreatedBy**](GetJobs200ResponseAllOfJobAnyOf2CreatedBy.md) |  | [optional] 
**TargetType** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to [**[]GetJobs200ResponseAllOfJobAnyOf2TargetsInner**](GetJobs200ResponseAllOfJobAnyOf2TargetsInner.md) |  | [optional] 
**CustomConfig** | Pointer to **NullableString** |  | [optional] 
**CustomOptions** | Pointer to [**GetJobs200ResponseAllOfJobAnyOf2CustomOptions**](GetJobs200ResponseAllOfJobAnyOf2CustomOptions.md) |  | [optional] 

## Methods

### NewGetJobs200ResponseAllOfJobAnyOf2

`func NewGetJobs200ResponseAllOfJobAnyOf2() *GetJobs200ResponseAllOfJobAnyOf2`

NewGetJobs200ResponseAllOfJobAnyOf2 instantiates a new GetJobs200ResponseAllOfJobAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetJobs200ResponseAllOfJobAnyOf2) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetType() GetJobs200ResponseAllOfJobAnyOf2Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetTypeOk() (*GetJobs200ResponseAllOfJobAnyOf2Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetType(v GetJobs200ResponseAllOfJobAnyOf2Type)`

SetType sets Type field to given value.

### HasType

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkflow

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetWorkflow() GetJobs200ResponseAllOfJobAnyOf2Workflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetWorkflowOk() (*GetJobs200ResponseAllOfJobAnyOf2Workflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetWorkflow(v GetJobs200ResponseAllOfJobAnyOf2Workflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetJobSummary

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetJobSummary() string`

GetJobSummary returns the JobSummary field if non-nil, zero value otherwise.

### GetJobSummaryOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetJobSummaryOk() (*string, bool)`

GetJobSummaryOk returns a tuple with the JobSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSummary

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetJobSummary(v string)`

SetJobSummary sets JobSummary field to given value.

### HasJobSummary

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasJobSummary() bool`

HasJobSummary returns a boolean if a field has been set.

### SetJobSummaryNil

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetJobSummaryNil(b bool)`

 SetJobSummaryNil sets the value for JobSummary to be an explicit nil

### UnsetJobSummary
`func (o *GetJobs200ResponseAllOfJobAnyOf2) UnsetJobSummary()`

UnsetJobSummary ensures that no value is present for JobSummary, not even an explicit nil
### GetScheduleMode

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetScheduleMode() GetJobs200ResponseAllOfJobAnyOf2ScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetScheduleModeOk() (*GetJobs200ResponseAllOfJobAnyOf2ScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetScheduleMode(v GetJobs200ResponseAllOfJobAnyOf2ScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetDateTime

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTimeNil

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetDateTimeNil(b bool)`

 SetDateTimeNil sets the value for DateTime to be an explicit nil

### UnsetDateTime
`func (o *GetJobs200ResponseAllOfJobAnyOf2) UnsetDateTime()`

UnsetDateTime ensures that no value is present for DateTime, not even an explicit nil
### GetStatus

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetJobs200ResponseAllOfJobAnyOf2) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNamespace

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *GetJobs200ResponseAllOfJobAnyOf2) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetCategory

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetJobs200ResponseAllOfJobAnyOf2) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetJobs200ResponseAllOfJobAnyOf2) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastRun

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastResult

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetLastResult() string`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetLastResultOk() (*string, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetLastResult(v string)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetCreatedBy() GetJobs200ResponseAllOfJobAnyOf2CreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetCreatedByOk() (*GetJobs200ResponseAllOfJobAnyOf2CreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetCreatedBy(v GetJobs200ResponseAllOfJobAnyOf2CreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTargetType

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargets

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetTargets() []GetJobs200ResponseAllOfJobAnyOf2TargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetTargetsOk() (*[]GetJobs200ResponseAllOfJobAnyOf2TargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetTargets(v []GetJobs200ResponseAllOfJobAnyOf2TargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *GetJobs200ResponseAllOfJobAnyOf2) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetCustomConfig

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### SetCustomConfigNil

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetCustomConfigNil(b bool)`

 SetCustomConfigNil sets the value for CustomConfig to be an explicit nil

### UnsetCustomConfig
`func (o *GetJobs200ResponseAllOfJobAnyOf2) UnsetCustomConfig()`

UnsetCustomConfig ensures that no value is present for CustomConfig, not even an explicit nil
### GetCustomOptions

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetCustomOptions() GetJobs200ResponseAllOfJobAnyOf2CustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *GetJobs200ResponseAllOfJobAnyOf2) GetCustomOptionsOk() (*GetJobs200ResponseAllOfJobAnyOf2CustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *GetJobs200ResponseAllOfJobAnyOf2) SetCustomOptions(v GetJobs200ResponseAllOfJobAnyOf2CustomOptions)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *GetJobs200ResponseAllOfJobAnyOf2) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SecurityScanJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the Job | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Enabled** | Pointer to **bool** | Use this to set enabled state | [optional] [default to true]
**SecurityPackage** | [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | 
**ScanPath** | Pointer to **string** | Scan Checklist | [optional] 
**SecurityProfile** | Pointer to **string** | Security Profile | [optional] 
**TargetType** | **string** | Target type where job will execute | 
**Targets** | [**[]WorkflowJobPayloadTargetsInner**](WorkflowJobPayloadTargetsInner.md) |  | 
**ScheduleMode** | [**WorkflowJobPayloadScheduleMode**](WorkflowJobPayloadScheduleMode.md) |  | 
**CustomOptions** | Pointer to **map[string]interface{}** | Map of options to be used as values in the workflow tasks. These correspond to option types. | [optional] 
**CustomConfig** | Pointer to **string** | Job custom configuration (String in JSON format) | [optional] 
**DateTime** | Pointer to **time.Time** | Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is &#39;dateTime&#39;. | [optional] 
**Run** | Pointer to **bool** | If true, executes job | [optional] 

## Methods

### NewSecurityScanJob

`func NewSecurityScanJob(name string, securityPackage WorkflowJobPayloadTask, targetType string, targets []WorkflowJobPayloadTargetsInner, scheduleMode WorkflowJobPayloadScheduleMode, ) *SecurityScanJob`

NewSecurityScanJob instantiates a new SecurityScanJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityScanJobWithDefaults

`func NewSecurityScanJobWithDefaults() *SecurityScanJob`

NewSecurityScanJobWithDefaults instantiates a new SecurityScanJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecurityScanJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityScanJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityScanJob) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *SecurityScanJob) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SecurityScanJob) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SecurityScanJob) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SecurityScanJob) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEnabled

`func (o *SecurityScanJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityScanJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityScanJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SecurityScanJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSecurityPackage

`func (o *SecurityScanJob) GetSecurityPackage() WorkflowJobPayloadTask`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *SecurityScanJob) GetSecurityPackageOk() (*WorkflowJobPayloadTask, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *SecurityScanJob) SetSecurityPackage(v WorkflowJobPayloadTask)`

SetSecurityPackage sets SecurityPackage field to given value.


### GetScanPath

`func (o *SecurityScanJob) GetScanPath() string`

GetScanPath returns the ScanPath field if non-nil, zero value otherwise.

### GetScanPathOk

`func (o *SecurityScanJob) GetScanPathOk() (*string, bool)`

GetScanPathOk returns a tuple with the ScanPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPath

`func (o *SecurityScanJob) SetScanPath(v string)`

SetScanPath sets ScanPath field to given value.

### HasScanPath

`func (o *SecurityScanJob) HasScanPath() bool`

HasScanPath returns a boolean if a field has been set.

### GetSecurityProfile

`func (o *SecurityScanJob) GetSecurityProfile() string`

GetSecurityProfile returns the SecurityProfile field if non-nil, zero value otherwise.

### GetSecurityProfileOk

`func (o *SecurityScanJob) GetSecurityProfileOk() (*string, bool)`

GetSecurityProfileOk returns a tuple with the SecurityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProfile

`func (o *SecurityScanJob) SetSecurityProfile(v string)`

SetSecurityProfile sets SecurityProfile field to given value.

### HasSecurityProfile

`func (o *SecurityScanJob) HasSecurityProfile() bool`

HasSecurityProfile returns a boolean if a field has been set.

### GetTargetType

`func (o *SecurityScanJob) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *SecurityScanJob) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *SecurityScanJob) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetTargets

`func (o *SecurityScanJob) GetTargets() []WorkflowJobPayloadTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *SecurityScanJob) GetTargetsOk() (*[]WorkflowJobPayloadTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *SecurityScanJob) SetTargets(v []WorkflowJobPayloadTargetsInner)`

SetTargets sets Targets field to given value.


### GetScheduleMode

`func (o *SecurityScanJob) GetScheduleMode() WorkflowJobPayloadScheduleMode`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *SecurityScanJob) GetScheduleModeOk() (*WorkflowJobPayloadScheduleMode, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *SecurityScanJob) SetScheduleMode(v WorkflowJobPayloadScheduleMode)`

SetScheduleMode sets ScheduleMode field to given value.


### GetCustomOptions

`func (o *SecurityScanJob) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *SecurityScanJob) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *SecurityScanJob) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *SecurityScanJob) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetCustomConfig

`func (o *SecurityScanJob) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *SecurityScanJob) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *SecurityScanJob) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *SecurityScanJob) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetDateTime

`func (o *SecurityScanJob) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *SecurityScanJob) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *SecurityScanJob) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *SecurityScanJob) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetRun

`func (o *SecurityScanJob) GetRun() bool`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *SecurityScanJob) GetRunOk() (*bool, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *SecurityScanJob) SetRun(v bool)`

SetRun sets Run field to given value.

### HasRun

`func (o *SecurityScanJob) HasRun() bool`

HasRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



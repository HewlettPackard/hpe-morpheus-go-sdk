# ListBackupSettings200ResponseBackupSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupsEnabled** | Pointer to **bool** |  | [optional] 
**CreateBackups** | Pointer to **bool** |  | [optional] 
**BackupAppliance** | Pointer to **bool** |  | [optional] 
**DefaultStorageBucket** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**DefaultSchedule** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**RetentionCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewListBackupSettings200ResponseBackupSettings

`func NewListBackupSettings200ResponseBackupSettings() *ListBackupSettings200ResponseBackupSettings`

NewListBackupSettings200ResponseBackupSettings instantiates a new ListBackupSettings200ResponseBackupSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupSettings200ResponseBackupSettingsWithDefaults

`func NewListBackupSettings200ResponseBackupSettingsWithDefaults() *ListBackupSettings200ResponseBackupSettings`

NewListBackupSettings200ResponseBackupSettingsWithDefaults instantiates a new ListBackupSettings200ResponseBackupSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupsEnabled

`func (o *ListBackupSettings200ResponseBackupSettings) GetBackupsEnabled() bool`

GetBackupsEnabled returns the BackupsEnabled field if non-nil, zero value otherwise.

### GetBackupsEnabledOk

`func (o *ListBackupSettings200ResponseBackupSettings) GetBackupsEnabledOk() (*bool, bool)`

GetBackupsEnabledOk returns a tuple with the BackupsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsEnabled

`func (o *ListBackupSettings200ResponseBackupSettings) SetBackupsEnabled(v bool)`

SetBackupsEnabled sets BackupsEnabled field to given value.

### HasBackupsEnabled

`func (o *ListBackupSettings200ResponseBackupSettings) HasBackupsEnabled() bool`

HasBackupsEnabled returns a boolean if a field has been set.

### GetCreateBackups

`func (o *ListBackupSettings200ResponseBackupSettings) GetCreateBackups() bool`

GetCreateBackups returns the CreateBackups field if non-nil, zero value otherwise.

### GetCreateBackupsOk

`func (o *ListBackupSettings200ResponseBackupSettings) GetCreateBackupsOk() (*bool, bool)`

GetCreateBackupsOk returns a tuple with the CreateBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackups

`func (o *ListBackupSettings200ResponseBackupSettings) SetCreateBackups(v bool)`

SetCreateBackups sets CreateBackups field to given value.

### HasCreateBackups

`func (o *ListBackupSettings200ResponseBackupSettings) HasCreateBackups() bool`

HasCreateBackups returns a boolean if a field has been set.

### GetBackupAppliance

`func (o *ListBackupSettings200ResponseBackupSettings) GetBackupAppliance() bool`

GetBackupAppliance returns the BackupAppliance field if non-nil, zero value otherwise.

### GetBackupApplianceOk

`func (o *ListBackupSettings200ResponseBackupSettings) GetBackupApplianceOk() (*bool, bool)`

GetBackupApplianceOk returns a tuple with the BackupAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAppliance

`func (o *ListBackupSettings200ResponseBackupSettings) SetBackupAppliance(v bool)`

SetBackupAppliance sets BackupAppliance field to given value.

### HasBackupAppliance

`func (o *ListBackupSettings200ResponseBackupSettings) HasBackupAppliance() bool`

HasBackupAppliance returns a boolean if a field has been set.

### GetDefaultStorageBucket

`func (o *ListBackupSettings200ResponseBackupSettings) GetDefaultStorageBucket() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetDefaultStorageBucket returns the DefaultStorageBucket field if non-nil, zero value otherwise.

### GetDefaultStorageBucketOk

`func (o *ListBackupSettings200ResponseBackupSettings) GetDefaultStorageBucketOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetDefaultStorageBucketOk returns a tuple with the DefaultStorageBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStorageBucket

`func (o *ListBackupSettings200ResponseBackupSettings) SetDefaultStorageBucket(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetDefaultStorageBucket sets DefaultStorageBucket field to given value.

### HasDefaultStorageBucket

`func (o *ListBackupSettings200ResponseBackupSettings) HasDefaultStorageBucket() bool`

HasDefaultStorageBucket returns a boolean if a field has been set.

### GetDefaultSchedule

`func (o *ListBackupSettings200ResponseBackupSettings) GetDefaultSchedule() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetDefaultSchedule returns the DefaultSchedule field if non-nil, zero value otherwise.

### GetDefaultScheduleOk

`func (o *ListBackupSettings200ResponseBackupSettings) GetDefaultScheduleOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetDefaultScheduleOk returns a tuple with the DefaultSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSchedule

`func (o *ListBackupSettings200ResponseBackupSettings) SetDefaultSchedule(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetDefaultSchedule sets DefaultSchedule field to given value.

### HasDefaultSchedule

`func (o *ListBackupSettings200ResponseBackupSettings) HasDefaultSchedule() bool`

HasDefaultSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *ListBackupSettings200ResponseBackupSettings) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *ListBackupSettings200ResponseBackupSettings) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *ListBackupSettings200ResponseBackupSettings) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *ListBackupSettings200ResponseBackupSettings) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



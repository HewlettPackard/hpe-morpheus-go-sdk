# BackupSettings

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

### NewBackupSettings

`func NewBackupSettings() *BackupSettings`

NewBackupSettings instantiates a new BackupSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSettingsWithDefaults

`func NewBackupSettingsWithDefaults() *BackupSettings`

NewBackupSettingsWithDefaults instantiates a new BackupSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupsEnabled

`func (o *BackupSettings) GetBackupsEnabled() bool`

GetBackupsEnabled returns the BackupsEnabled field if non-nil, zero value otherwise.

### GetBackupsEnabledOk

`func (o *BackupSettings) GetBackupsEnabledOk() (*bool, bool)`

GetBackupsEnabledOk returns a tuple with the BackupsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsEnabled

`func (o *BackupSettings) SetBackupsEnabled(v bool)`

SetBackupsEnabled sets BackupsEnabled field to given value.

### HasBackupsEnabled

`func (o *BackupSettings) HasBackupsEnabled() bool`

HasBackupsEnabled returns a boolean if a field has been set.

### GetCreateBackups

`func (o *BackupSettings) GetCreateBackups() bool`

GetCreateBackups returns the CreateBackups field if non-nil, zero value otherwise.

### GetCreateBackupsOk

`func (o *BackupSettings) GetCreateBackupsOk() (*bool, bool)`

GetCreateBackupsOk returns a tuple with the CreateBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackups

`func (o *BackupSettings) SetCreateBackups(v bool)`

SetCreateBackups sets CreateBackups field to given value.

### HasCreateBackups

`func (o *BackupSettings) HasCreateBackups() bool`

HasCreateBackups returns a boolean if a field has been set.

### GetBackupAppliance

`func (o *BackupSettings) GetBackupAppliance() bool`

GetBackupAppliance returns the BackupAppliance field if non-nil, zero value otherwise.

### GetBackupApplianceOk

`func (o *BackupSettings) GetBackupApplianceOk() (*bool, bool)`

GetBackupApplianceOk returns a tuple with the BackupAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAppliance

`func (o *BackupSettings) SetBackupAppliance(v bool)`

SetBackupAppliance sets BackupAppliance field to given value.

### HasBackupAppliance

`func (o *BackupSettings) HasBackupAppliance() bool`

HasBackupAppliance returns a boolean if a field has been set.

### GetDefaultStorageBucket

`func (o *BackupSettings) GetDefaultStorageBucket() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetDefaultStorageBucket returns the DefaultStorageBucket field if non-nil, zero value otherwise.

### GetDefaultStorageBucketOk

`func (o *BackupSettings) GetDefaultStorageBucketOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetDefaultStorageBucketOk returns a tuple with the DefaultStorageBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStorageBucket

`func (o *BackupSettings) SetDefaultStorageBucket(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetDefaultStorageBucket sets DefaultStorageBucket field to given value.

### HasDefaultStorageBucket

`func (o *BackupSettings) HasDefaultStorageBucket() bool`

HasDefaultStorageBucket returns a boolean if a field has been set.

### GetDefaultSchedule

`func (o *BackupSettings) GetDefaultSchedule() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetDefaultSchedule returns the DefaultSchedule field if non-nil, zero value otherwise.

### GetDefaultScheduleOk

`func (o *BackupSettings) GetDefaultScheduleOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetDefaultScheduleOk returns a tuple with the DefaultSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSchedule

`func (o *BackupSettings) SetDefaultSchedule(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetDefaultSchedule sets DefaultSchedule field to given value.

### HasDefaultSchedule

`func (o *BackupSettings) HasDefaultSchedule() bool`

HasDefaultSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *BackupSettings) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *BackupSettings) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *BackupSettings) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *BackupSettings) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



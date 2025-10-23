# BackupSettingsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupsEnabled** | Pointer to **bool** | Use this to enable / disable scheduled backups | [optional] 
**RetentionCount** | Pointer to **int64** | Maximum number of successful backups to retain | [optional] 
**CreateBackups** | Pointer to **bool** | Use this to enable / disable create backups | [optional] 
**BackupAppliance** | Pointer to **bool** | When enabled, a Backup will be created to backup the Morpheus appliance database | [optional] 
**UpdateExisting** | Pointer to **bool** | Use this to update existing backups with new settings | [optional] 
**DefaultSchedule** | Pointer to [**UpdateBackupSettingsRequestBackupSettingsDefaultSchedule**](UpdateBackupSettingsRequestBackupSettingsDefaultSchedule.md) |  | [optional] 
**ClearDefaultSchedule** | Pointer to **bool** | Use this to clear existing default backup schedule | [optional] 
**DefaultStorageBucket** | Pointer to [**UpdateBackupSettingsRequestBackupSettingsDefaultStorageBucket**](UpdateBackupSettingsRequestBackupSettingsDefaultStorageBucket.md) |  | [optional] 
**ClearDefaultStorageBucket** | Pointer to **bool** | Use this to clear default store bucket | [optional] 

## Methods

### NewBackupSettingsUpdate

`func NewBackupSettingsUpdate() *BackupSettingsUpdate`

NewBackupSettingsUpdate instantiates a new BackupSettingsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSettingsUpdateWithDefaults

`func NewBackupSettingsUpdateWithDefaults() *BackupSettingsUpdate`

NewBackupSettingsUpdateWithDefaults instantiates a new BackupSettingsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupsEnabled

`func (o *BackupSettingsUpdate) GetBackupsEnabled() bool`

GetBackupsEnabled returns the BackupsEnabled field if non-nil, zero value otherwise.

### GetBackupsEnabledOk

`func (o *BackupSettingsUpdate) GetBackupsEnabledOk() (*bool, bool)`

GetBackupsEnabledOk returns a tuple with the BackupsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsEnabled

`func (o *BackupSettingsUpdate) SetBackupsEnabled(v bool)`

SetBackupsEnabled sets BackupsEnabled field to given value.

### HasBackupsEnabled

`func (o *BackupSettingsUpdate) HasBackupsEnabled() bool`

HasBackupsEnabled returns a boolean if a field has been set.

### GetRetentionCount

`func (o *BackupSettingsUpdate) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *BackupSettingsUpdate) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *BackupSettingsUpdate) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *BackupSettingsUpdate) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### GetCreateBackups

`func (o *BackupSettingsUpdate) GetCreateBackups() bool`

GetCreateBackups returns the CreateBackups field if non-nil, zero value otherwise.

### GetCreateBackupsOk

`func (o *BackupSettingsUpdate) GetCreateBackupsOk() (*bool, bool)`

GetCreateBackupsOk returns a tuple with the CreateBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackups

`func (o *BackupSettingsUpdate) SetCreateBackups(v bool)`

SetCreateBackups sets CreateBackups field to given value.

### HasCreateBackups

`func (o *BackupSettingsUpdate) HasCreateBackups() bool`

HasCreateBackups returns a boolean if a field has been set.

### GetBackupAppliance

`func (o *BackupSettingsUpdate) GetBackupAppliance() bool`

GetBackupAppliance returns the BackupAppliance field if non-nil, zero value otherwise.

### GetBackupApplianceOk

`func (o *BackupSettingsUpdate) GetBackupApplianceOk() (*bool, bool)`

GetBackupApplianceOk returns a tuple with the BackupAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAppliance

`func (o *BackupSettingsUpdate) SetBackupAppliance(v bool)`

SetBackupAppliance sets BackupAppliance field to given value.

### HasBackupAppliance

`func (o *BackupSettingsUpdate) HasBackupAppliance() bool`

HasBackupAppliance returns a boolean if a field has been set.

### GetUpdateExisting

`func (o *BackupSettingsUpdate) GetUpdateExisting() bool`

GetUpdateExisting returns the UpdateExisting field if non-nil, zero value otherwise.

### GetUpdateExistingOk

`func (o *BackupSettingsUpdate) GetUpdateExistingOk() (*bool, bool)`

GetUpdateExistingOk returns a tuple with the UpdateExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateExisting

`func (o *BackupSettingsUpdate) SetUpdateExisting(v bool)`

SetUpdateExisting sets UpdateExisting field to given value.

### HasUpdateExisting

`func (o *BackupSettingsUpdate) HasUpdateExisting() bool`

HasUpdateExisting returns a boolean if a field has been set.

### GetDefaultSchedule

`func (o *BackupSettingsUpdate) GetDefaultSchedule() UpdateBackupSettingsRequestBackupSettingsDefaultSchedule`

GetDefaultSchedule returns the DefaultSchedule field if non-nil, zero value otherwise.

### GetDefaultScheduleOk

`func (o *BackupSettingsUpdate) GetDefaultScheduleOk() (*UpdateBackupSettingsRequestBackupSettingsDefaultSchedule, bool)`

GetDefaultScheduleOk returns a tuple with the DefaultSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSchedule

`func (o *BackupSettingsUpdate) SetDefaultSchedule(v UpdateBackupSettingsRequestBackupSettingsDefaultSchedule)`

SetDefaultSchedule sets DefaultSchedule field to given value.

### HasDefaultSchedule

`func (o *BackupSettingsUpdate) HasDefaultSchedule() bool`

HasDefaultSchedule returns a boolean if a field has been set.

### GetClearDefaultSchedule

`func (o *BackupSettingsUpdate) GetClearDefaultSchedule() bool`

GetClearDefaultSchedule returns the ClearDefaultSchedule field if non-nil, zero value otherwise.

### GetClearDefaultScheduleOk

`func (o *BackupSettingsUpdate) GetClearDefaultScheduleOk() (*bool, bool)`

GetClearDefaultScheduleOk returns a tuple with the ClearDefaultSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearDefaultSchedule

`func (o *BackupSettingsUpdate) SetClearDefaultSchedule(v bool)`

SetClearDefaultSchedule sets ClearDefaultSchedule field to given value.

### HasClearDefaultSchedule

`func (o *BackupSettingsUpdate) HasClearDefaultSchedule() bool`

HasClearDefaultSchedule returns a boolean if a field has been set.

### GetDefaultStorageBucket

`func (o *BackupSettingsUpdate) GetDefaultStorageBucket() UpdateBackupSettingsRequestBackupSettingsDefaultStorageBucket`

GetDefaultStorageBucket returns the DefaultStorageBucket field if non-nil, zero value otherwise.

### GetDefaultStorageBucketOk

`func (o *BackupSettingsUpdate) GetDefaultStorageBucketOk() (*UpdateBackupSettingsRequestBackupSettingsDefaultStorageBucket, bool)`

GetDefaultStorageBucketOk returns a tuple with the DefaultStorageBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStorageBucket

`func (o *BackupSettingsUpdate) SetDefaultStorageBucket(v UpdateBackupSettingsRequestBackupSettingsDefaultStorageBucket)`

SetDefaultStorageBucket sets DefaultStorageBucket field to given value.

### HasDefaultStorageBucket

`func (o *BackupSettingsUpdate) HasDefaultStorageBucket() bool`

HasDefaultStorageBucket returns a boolean if a field has been set.

### GetClearDefaultStorageBucket

`func (o *BackupSettingsUpdate) GetClearDefaultStorageBucket() bool`

GetClearDefaultStorageBucket returns the ClearDefaultStorageBucket field if non-nil, zero value otherwise.

### GetClearDefaultStorageBucketOk

`func (o *BackupSettingsUpdate) GetClearDefaultStorageBucketOk() (*bool, bool)`

GetClearDefaultStorageBucketOk returns a tuple with the ClearDefaultStorageBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearDefaultStorageBucket

`func (o *BackupSettingsUpdate) SetClearDefaultStorageBucket(v bool)`

SetClearDefaultStorageBucket sets ClearDefaultStorageBucket field to given value.

### HasClearDefaultStorageBucket

`func (o *BackupSettingsUpdate) HasClearDefaultStorageBucket() bool`

HasClearDefaultStorageBucket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



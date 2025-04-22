# Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**LocationType** | Pointer to **string** | Source Type (instance, server, storage) | [optional] 
**Instance** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerInstance**](ListBackups200ResponseAllOfBackupsInnerInstance.md) |  | [optional] 
**ContainerId** | Pointer to **int64** |  | [optional] 
**Job** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerJob**](ListBackups200ResponseAllOfBackupsInnerJob.md) |  | [optional] 
**Schedule** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerSchedule**](ListBackups200ResponseAllOfBackupsInnerSchedule.md) |  | [optional] 
**RetentionCount** | Pointer to **int64** |  | [optional] 
**BackupType** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerBackupType**](ListBackups200ResponseAllOfBackupsInnerBackupType.md) |  | [optional] 
**StorageProvider** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerStorageProvider**](ListBackups200ResponseAllOfBackupsInnerStorageProvider.md) |  | [optional] 
**BackupProvider** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerBackupProvider**](ListBackups200ResponseAllOfBackupsInnerBackupProvider.md) |  | [optional] 
**BackupRespository** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerBackupRespository**](ListBackups200ResponseAllOfBackupsInnerBackupRespository.md) |  | [optional] 
**CronExpression** | Pointer to **string** | Cron Expression | [optional] 
**NextFire** | Pointer to **time.Time** | Next Fire | [optional] 
**LastStatus** | Pointer to **string** | Last Status | [optional] 
**LastResult** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerLastResult**](ListBackups200ResponseAllOfBackupsInnerLastResult.md) |  | [optional] 
**Stats** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerStats**](ListBackups200ResponseAllOfBackupsInnerStats.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Methods

### NewBackup

`func NewBackup() *Backup`

NewBackup instantiates a new Backup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWithDefaults

`func NewBackupWithDefaults() *Backup`

NewBackupWithDefaults instantiates a new Backup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Backup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Backup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Backup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Backup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Backup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Backup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Backup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Backup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocationType

`func (o *Backup) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *Backup) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *Backup) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *Backup) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetInstance

`func (o *Backup) GetInstance() ListBackups200ResponseAllOfBackupsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *Backup) GetInstanceOk() (*ListBackups200ResponseAllOfBackupsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *Backup) SetInstance(v ListBackups200ResponseAllOfBackupsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *Backup) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerId

`func (o *Backup) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *Backup) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *Backup) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *Backup) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetJob

`func (o *Backup) GetJob() ListBackups200ResponseAllOfBackupsInnerJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *Backup) GetJobOk() (*ListBackups200ResponseAllOfBackupsInnerJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *Backup) SetJob(v ListBackups200ResponseAllOfBackupsInnerJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *Backup) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetSchedule

`func (o *Backup) GetSchedule() ListBackups200ResponseAllOfBackupsInnerSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Backup) GetScheduleOk() (*ListBackups200ResponseAllOfBackupsInnerSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Backup) SetSchedule(v ListBackups200ResponseAllOfBackupsInnerSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Backup) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *Backup) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *Backup) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *Backup) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *Backup) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### GetBackupType

`func (o *Backup) GetBackupType() ListBackups200ResponseAllOfBackupsInnerBackupType`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *Backup) GetBackupTypeOk() (*ListBackups200ResponseAllOfBackupsInnerBackupType, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *Backup) SetBackupType(v ListBackups200ResponseAllOfBackupsInnerBackupType)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *Backup) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetStorageProvider

`func (o *Backup) GetStorageProvider() ListBackups200ResponseAllOfBackupsInnerStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *Backup) GetStorageProviderOk() (*ListBackups200ResponseAllOfBackupsInnerStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *Backup) SetStorageProvider(v ListBackups200ResponseAllOfBackupsInnerStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *Backup) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetBackupProvider

`func (o *Backup) GetBackupProvider() ListBackups200ResponseAllOfBackupsInnerBackupProvider`

GetBackupProvider returns the BackupProvider field if non-nil, zero value otherwise.

### GetBackupProviderOk

`func (o *Backup) GetBackupProviderOk() (*ListBackups200ResponseAllOfBackupsInnerBackupProvider, bool)`

GetBackupProviderOk returns a tuple with the BackupProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProvider

`func (o *Backup) SetBackupProvider(v ListBackups200ResponseAllOfBackupsInnerBackupProvider)`

SetBackupProvider sets BackupProvider field to given value.

### HasBackupProvider

`func (o *Backup) HasBackupProvider() bool`

HasBackupProvider returns a boolean if a field has been set.

### GetBackupRespository

`func (o *Backup) GetBackupRespository() ListBackups200ResponseAllOfBackupsInnerBackupRespository`

GetBackupRespository returns the BackupRespository field if non-nil, zero value otherwise.

### GetBackupRespositoryOk

`func (o *Backup) GetBackupRespositoryOk() (*ListBackups200ResponseAllOfBackupsInnerBackupRespository, bool)`

GetBackupRespositoryOk returns a tuple with the BackupRespository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRespository

`func (o *Backup) SetBackupRespository(v ListBackups200ResponseAllOfBackupsInnerBackupRespository)`

SetBackupRespository sets BackupRespository field to given value.

### HasBackupRespository

`func (o *Backup) HasBackupRespository() bool`

HasBackupRespository returns a boolean if a field has been set.

### GetCronExpression

`func (o *Backup) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *Backup) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *Backup) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *Backup) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetNextFire

`func (o *Backup) GetNextFire() time.Time`

GetNextFire returns the NextFire field if non-nil, zero value otherwise.

### GetNextFireOk

`func (o *Backup) GetNextFireOk() (*time.Time, bool)`

GetNextFireOk returns a tuple with the NextFire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFire

`func (o *Backup) SetNextFire(v time.Time)`

SetNextFire sets NextFire field to given value.

### HasNextFire

`func (o *Backup) HasNextFire() bool`

HasNextFire returns a boolean if a field has been set.

### GetLastStatus

`func (o *Backup) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *Backup) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *Backup) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *Backup) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetLastResult

`func (o *Backup) GetLastResult() ListBackups200ResponseAllOfBackupsInnerLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *Backup) GetLastResultOk() (*ListBackups200ResponseAllOfBackupsInnerLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *Backup) SetLastResult(v ListBackups200ResponseAllOfBackupsInnerLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *Backup) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetStats

`func (o *Backup) GetStats() ListBackups200ResponseAllOfBackupsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Backup) GetStatsOk() (*ListBackups200ResponseAllOfBackupsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Backup) SetStats(v ListBackups200ResponseAllOfBackupsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Backup) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetEnabled

`func (o *Backup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Backup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Backup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Backup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *Backup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Backup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Backup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Backup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Backup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Backup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Backup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Backup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ExecuteBackups200ResponseAllOfBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**LocationType** | Pointer to **string** | Source Type (instance, server, storage) | [optional] 
**Instance** | Pointer to [**ExecuteBackups200ResponseAllOfBackupInstance**](ExecuteBackups200ResponseAllOfBackupInstance.md) |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**Job** | Pointer to [**ExecuteBackups200ResponseAllOfBackupJob**](ExecuteBackups200ResponseAllOfBackupJob.md) |  | [optional] 
**Schedule** | Pointer to [**ExecuteBackups200ResponseAllOfBackupSchedule**](ExecuteBackups200ResponseAllOfBackupSchedule.md) |  | [optional] 
**RetentionCount** | Pointer to **NullableInt64** |  | [optional] 
**BackupType** | Pointer to [**ExecuteBackups200ResponseAllOfBackupBackupType**](ExecuteBackups200ResponseAllOfBackupBackupType.md) |  | [optional] 
**StorageProvider** | Pointer to [**ExecuteBackups200ResponseAllOfBackupStorageProvider**](ExecuteBackups200ResponseAllOfBackupStorageProvider.md) |  | [optional] 
**BackupProvider** | Pointer to [**ExecuteBackups200ResponseAllOfBackupBackupProvider**](ExecuteBackups200ResponseAllOfBackupBackupProvider.md) |  | [optional] 
**BackupRespository** | Pointer to [**ExecuteBackups200ResponseAllOfBackupBackupRespository**](ExecuteBackups200ResponseAllOfBackupBackupRespository.md) |  | [optional] 
**CronExpression** | Pointer to **NullableString** | Cron Expression | [optional] 
**NextFire** | Pointer to **NullableTime** | Next Fire | [optional] 
**LastStatus** | Pointer to **NullableString** | Last Status | [optional] 
**LastResult** | Pointer to [**ExecuteBackups200ResponseAllOfBackupLastResult**](ExecuteBackups200ResponseAllOfBackupLastResult.md) |  | [optional] 
**Stats** | Pointer to [**ExecuteBackups200ResponseAllOfBackupStats**](ExecuteBackups200ResponseAllOfBackupStats.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Methods

### NewExecuteBackups200ResponseAllOfBackup

`func NewExecuteBackups200ResponseAllOfBackup() *ExecuteBackups200ResponseAllOfBackup`

NewExecuteBackups200ResponseAllOfBackup instantiates a new ExecuteBackups200ResponseAllOfBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *ExecuteBackups200ResponseAllOfBackup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecuteBackups200ResponseAllOfBackup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ExecuteBackups200ResponseAllOfBackup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ExecuteBackups200ResponseAllOfBackup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecuteBackups200ResponseAllOfBackup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExecuteBackups200ResponseAllOfBackup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocationType

`func (o *ExecuteBackups200ResponseAllOfBackup) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *ExecuteBackups200ResponseAllOfBackup) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *ExecuteBackups200ResponseAllOfBackup) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetInstance

`func (o *ExecuteBackups200ResponseAllOfBackup) GetInstance() ExecuteBackups200ResponseAllOfBackupInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetInstanceOk() (*ExecuteBackups200ResponseAllOfBackupInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ExecuteBackups200ResponseAllOfBackup) SetInstance(v ExecuteBackups200ResponseAllOfBackupInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ExecuteBackups200ResponseAllOfBackup) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerId

`func (o *ExecuteBackups200ResponseAllOfBackup) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ExecuteBackups200ResponseAllOfBackup) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ExecuteBackups200ResponseAllOfBackup) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *ExecuteBackups200ResponseAllOfBackup) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *ExecuteBackups200ResponseAllOfBackup) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetJob

`func (o *ExecuteBackups200ResponseAllOfBackup) GetJob() ExecuteBackups200ResponseAllOfBackupJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetJobOk() (*ExecuteBackups200ResponseAllOfBackupJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *ExecuteBackups200ResponseAllOfBackup) SetJob(v ExecuteBackups200ResponseAllOfBackupJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *ExecuteBackups200ResponseAllOfBackup) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetSchedule

`func (o *ExecuteBackups200ResponseAllOfBackup) GetSchedule() ExecuteBackups200ResponseAllOfBackupSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetScheduleOk() (*ExecuteBackups200ResponseAllOfBackupSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ExecuteBackups200ResponseAllOfBackup) SetSchedule(v ExecuteBackups200ResponseAllOfBackupSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ExecuteBackups200ResponseAllOfBackup) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *ExecuteBackups200ResponseAllOfBackup) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *ExecuteBackups200ResponseAllOfBackup) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *ExecuteBackups200ResponseAllOfBackup) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### SetRetentionCountNil

`func (o *ExecuteBackups200ResponseAllOfBackup) SetRetentionCountNil(b bool)`

 SetRetentionCountNil sets the value for RetentionCount to be an explicit nil

### UnsetRetentionCount
`func (o *ExecuteBackups200ResponseAllOfBackup) UnsetRetentionCount()`

UnsetRetentionCount ensures that no value is present for RetentionCount, not even an explicit nil
### GetBackupType

`func (o *ExecuteBackups200ResponseAllOfBackup) GetBackupType() ExecuteBackups200ResponseAllOfBackupBackupType`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetBackupTypeOk() (*ExecuteBackups200ResponseAllOfBackupBackupType, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *ExecuteBackups200ResponseAllOfBackup) SetBackupType(v ExecuteBackups200ResponseAllOfBackupBackupType)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *ExecuteBackups200ResponseAllOfBackup) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetStorageProvider

`func (o *ExecuteBackups200ResponseAllOfBackup) GetStorageProvider() ExecuteBackups200ResponseAllOfBackupStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetStorageProviderOk() (*ExecuteBackups200ResponseAllOfBackupStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ExecuteBackups200ResponseAllOfBackup) SetStorageProvider(v ExecuteBackups200ResponseAllOfBackupStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *ExecuteBackups200ResponseAllOfBackup) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetBackupProvider

`func (o *ExecuteBackups200ResponseAllOfBackup) GetBackupProvider() ExecuteBackups200ResponseAllOfBackupBackupProvider`

GetBackupProvider returns the BackupProvider field if non-nil, zero value otherwise.

### GetBackupProviderOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetBackupProviderOk() (*ExecuteBackups200ResponseAllOfBackupBackupProvider, bool)`

GetBackupProviderOk returns a tuple with the BackupProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProvider

`func (o *ExecuteBackups200ResponseAllOfBackup) SetBackupProvider(v ExecuteBackups200ResponseAllOfBackupBackupProvider)`

SetBackupProvider sets BackupProvider field to given value.

### HasBackupProvider

`func (o *ExecuteBackups200ResponseAllOfBackup) HasBackupProvider() bool`

HasBackupProvider returns a boolean if a field has been set.

### GetBackupRespository

`func (o *ExecuteBackups200ResponseAllOfBackup) GetBackupRespository() ExecuteBackups200ResponseAllOfBackupBackupRespository`

GetBackupRespository returns the BackupRespository field if non-nil, zero value otherwise.

### GetBackupRespositoryOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetBackupRespositoryOk() (*ExecuteBackups200ResponseAllOfBackupBackupRespository, bool)`

GetBackupRespositoryOk returns a tuple with the BackupRespository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRespository

`func (o *ExecuteBackups200ResponseAllOfBackup) SetBackupRespository(v ExecuteBackups200ResponseAllOfBackupBackupRespository)`

SetBackupRespository sets BackupRespository field to given value.

### HasBackupRespository

`func (o *ExecuteBackups200ResponseAllOfBackup) HasBackupRespository() bool`

HasBackupRespository returns a boolean if a field has been set.

### GetCronExpression

`func (o *ExecuteBackups200ResponseAllOfBackup) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *ExecuteBackups200ResponseAllOfBackup) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *ExecuteBackups200ResponseAllOfBackup) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### SetCronExpressionNil

`func (o *ExecuteBackups200ResponseAllOfBackup) SetCronExpressionNil(b bool)`

 SetCronExpressionNil sets the value for CronExpression to be an explicit nil

### UnsetCronExpression
`func (o *ExecuteBackups200ResponseAllOfBackup) UnsetCronExpression()`

UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
### GetNextFire

`func (o *ExecuteBackups200ResponseAllOfBackup) GetNextFire() time.Time`

GetNextFire returns the NextFire field if non-nil, zero value otherwise.

### GetNextFireOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetNextFireOk() (*time.Time, bool)`

GetNextFireOk returns a tuple with the NextFire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFire

`func (o *ExecuteBackups200ResponseAllOfBackup) SetNextFire(v time.Time)`

SetNextFire sets NextFire field to given value.

### HasNextFire

`func (o *ExecuteBackups200ResponseAllOfBackup) HasNextFire() bool`

HasNextFire returns a boolean if a field has been set.

### SetNextFireNil

`func (o *ExecuteBackups200ResponseAllOfBackup) SetNextFireNil(b bool)`

 SetNextFireNil sets the value for NextFire to be an explicit nil

### UnsetNextFire
`func (o *ExecuteBackups200ResponseAllOfBackup) UnsetNextFire()`

UnsetNextFire ensures that no value is present for NextFire, not even an explicit nil
### GetLastStatus

`func (o *ExecuteBackups200ResponseAllOfBackup) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *ExecuteBackups200ResponseAllOfBackup) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *ExecuteBackups200ResponseAllOfBackup) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### SetLastStatusNil

`func (o *ExecuteBackups200ResponseAllOfBackup) SetLastStatusNil(b bool)`

 SetLastStatusNil sets the value for LastStatus to be an explicit nil

### UnsetLastStatus
`func (o *ExecuteBackups200ResponseAllOfBackup) UnsetLastStatus()`

UnsetLastStatus ensures that no value is present for LastStatus, not even an explicit nil
### GetLastResult

`func (o *ExecuteBackups200ResponseAllOfBackup) GetLastResult() ExecuteBackups200ResponseAllOfBackupLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetLastResultOk() (*ExecuteBackups200ResponseAllOfBackupLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *ExecuteBackups200ResponseAllOfBackup) SetLastResult(v ExecuteBackups200ResponseAllOfBackupLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *ExecuteBackups200ResponseAllOfBackup) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetStats

`func (o *ExecuteBackups200ResponseAllOfBackup) GetStats() ExecuteBackups200ResponseAllOfBackupStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetStatsOk() (*ExecuteBackups200ResponseAllOfBackupStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ExecuteBackups200ResponseAllOfBackup) SetStats(v ExecuteBackups200ResponseAllOfBackupStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ExecuteBackups200ResponseAllOfBackup) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetEnabled

`func (o *ExecuteBackups200ResponseAllOfBackup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExecuteBackups200ResponseAllOfBackup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExecuteBackups200ResponseAllOfBackup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *ExecuteBackups200ResponseAllOfBackup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ExecuteBackups200ResponseAllOfBackup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ExecuteBackups200ResponseAllOfBackup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ExecuteBackups200ResponseAllOfBackup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ExecuteBackups200ResponseAllOfBackup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ExecuteBackups200ResponseAllOfBackup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ExecuteBackups200ResponseAllOfBackup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetBackups200ResponseBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**LocationType** | Pointer to **string** | Source Type (instance, server, storage) | [optional] 
**Instance** | Pointer to [**GetBackups200ResponseBackupInstance**](GetBackups200ResponseBackupInstance.md) |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**Job** | Pointer to [**GetBackups200ResponseBackupJob**](GetBackups200ResponseBackupJob.md) |  | [optional] 
**Schedule** | Pointer to [**GetBackups200ResponseBackupSchedule**](GetBackups200ResponseBackupSchedule.md) |  | [optional] 
**RetentionCount** | Pointer to **NullableInt64** |  | [optional] 
**BackupType** | Pointer to [**GetBackups200ResponseBackupBackupType**](GetBackups200ResponseBackupBackupType.md) |  | [optional] 
**StorageProvider** | Pointer to [**GetBackups200ResponseBackupStorageProvider**](GetBackups200ResponseBackupStorageProvider.md) |  | [optional] 
**BackupProvider** | Pointer to [**GetBackups200ResponseBackupBackupProvider**](GetBackups200ResponseBackupBackupProvider.md) |  | [optional] 
**BackupRespository** | Pointer to [**GetBackups200ResponseBackupBackupRespository**](GetBackups200ResponseBackupBackupRespository.md) |  | [optional] 
**CronExpression** | Pointer to **NullableString** | Cron Expression | [optional] 
**NextFire** | Pointer to **NullableTime** | Next Fire | [optional] 
**LastStatus** | Pointer to **NullableString** | Last Status | [optional] 
**LastResult** | Pointer to [**GetBackups200ResponseBackupLastResult**](GetBackups200ResponseBackupLastResult.md) |  | [optional] 
**Stats** | Pointer to [**GetBackups200ResponseBackupStats**](GetBackups200ResponseBackupStats.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Methods

### NewGetBackups200ResponseBackup

`func NewGetBackups200ResponseBackup() *GetBackups200ResponseBackup`

NewGetBackups200ResponseBackup instantiates a new GetBackups200ResponseBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetBackups200ResponseBackup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBackups200ResponseBackup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBackups200ResponseBackup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetBackups200ResponseBackup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetBackups200ResponseBackup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetBackups200ResponseBackup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetBackups200ResponseBackup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetBackups200ResponseBackup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocationType

`func (o *GetBackups200ResponseBackup) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *GetBackups200ResponseBackup) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *GetBackups200ResponseBackup) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *GetBackups200ResponseBackup) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetInstance

`func (o *GetBackups200ResponseBackup) GetInstance() GetBackups200ResponseBackupInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetBackups200ResponseBackup) GetInstanceOk() (*GetBackups200ResponseBackupInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetBackups200ResponseBackup) SetInstance(v GetBackups200ResponseBackupInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetBackups200ResponseBackup) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerId

`func (o *GetBackups200ResponseBackup) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GetBackups200ResponseBackup) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GetBackups200ResponseBackup) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *GetBackups200ResponseBackup) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *GetBackups200ResponseBackup) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *GetBackups200ResponseBackup) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetJob

`func (o *GetBackups200ResponseBackup) GetJob() GetBackups200ResponseBackupJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *GetBackups200ResponseBackup) GetJobOk() (*GetBackups200ResponseBackupJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *GetBackups200ResponseBackup) SetJob(v GetBackups200ResponseBackupJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *GetBackups200ResponseBackup) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetSchedule

`func (o *GetBackups200ResponseBackup) GetSchedule() GetBackups200ResponseBackupSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *GetBackups200ResponseBackup) GetScheduleOk() (*GetBackups200ResponseBackupSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *GetBackups200ResponseBackup) SetSchedule(v GetBackups200ResponseBackupSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *GetBackups200ResponseBackup) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *GetBackups200ResponseBackup) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *GetBackups200ResponseBackup) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *GetBackups200ResponseBackup) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *GetBackups200ResponseBackup) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### SetRetentionCountNil

`func (o *GetBackups200ResponseBackup) SetRetentionCountNil(b bool)`

 SetRetentionCountNil sets the value for RetentionCount to be an explicit nil

### UnsetRetentionCount
`func (o *GetBackups200ResponseBackup) UnsetRetentionCount()`

UnsetRetentionCount ensures that no value is present for RetentionCount, not even an explicit nil
### GetBackupType

`func (o *GetBackups200ResponseBackup) GetBackupType() GetBackups200ResponseBackupBackupType`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *GetBackups200ResponseBackup) GetBackupTypeOk() (*GetBackups200ResponseBackupBackupType, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *GetBackups200ResponseBackup) SetBackupType(v GetBackups200ResponseBackupBackupType)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *GetBackups200ResponseBackup) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetStorageProvider

`func (o *GetBackups200ResponseBackup) GetStorageProvider() GetBackups200ResponseBackupStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *GetBackups200ResponseBackup) GetStorageProviderOk() (*GetBackups200ResponseBackupStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *GetBackups200ResponseBackup) SetStorageProvider(v GetBackups200ResponseBackupStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *GetBackups200ResponseBackup) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetBackupProvider

`func (o *GetBackups200ResponseBackup) GetBackupProvider() GetBackups200ResponseBackupBackupProvider`

GetBackupProvider returns the BackupProvider field if non-nil, zero value otherwise.

### GetBackupProviderOk

`func (o *GetBackups200ResponseBackup) GetBackupProviderOk() (*GetBackups200ResponseBackupBackupProvider, bool)`

GetBackupProviderOk returns a tuple with the BackupProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProvider

`func (o *GetBackups200ResponseBackup) SetBackupProvider(v GetBackups200ResponseBackupBackupProvider)`

SetBackupProvider sets BackupProvider field to given value.

### HasBackupProvider

`func (o *GetBackups200ResponseBackup) HasBackupProvider() bool`

HasBackupProvider returns a boolean if a field has been set.

### GetBackupRespository

`func (o *GetBackups200ResponseBackup) GetBackupRespository() GetBackups200ResponseBackupBackupRespository`

GetBackupRespository returns the BackupRespository field if non-nil, zero value otherwise.

### GetBackupRespositoryOk

`func (o *GetBackups200ResponseBackup) GetBackupRespositoryOk() (*GetBackups200ResponseBackupBackupRespository, bool)`

GetBackupRespositoryOk returns a tuple with the BackupRespository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRespository

`func (o *GetBackups200ResponseBackup) SetBackupRespository(v GetBackups200ResponseBackupBackupRespository)`

SetBackupRespository sets BackupRespository field to given value.

### HasBackupRespository

`func (o *GetBackups200ResponseBackup) HasBackupRespository() bool`

HasBackupRespository returns a boolean if a field has been set.

### GetCronExpression

`func (o *GetBackups200ResponseBackup) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *GetBackups200ResponseBackup) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *GetBackups200ResponseBackup) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *GetBackups200ResponseBackup) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### SetCronExpressionNil

`func (o *GetBackups200ResponseBackup) SetCronExpressionNil(b bool)`

 SetCronExpressionNil sets the value for CronExpression to be an explicit nil

### UnsetCronExpression
`func (o *GetBackups200ResponseBackup) UnsetCronExpression()`

UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
### GetNextFire

`func (o *GetBackups200ResponseBackup) GetNextFire() time.Time`

GetNextFire returns the NextFire field if non-nil, zero value otherwise.

### GetNextFireOk

`func (o *GetBackups200ResponseBackup) GetNextFireOk() (*time.Time, bool)`

GetNextFireOk returns a tuple with the NextFire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFire

`func (o *GetBackups200ResponseBackup) SetNextFire(v time.Time)`

SetNextFire sets NextFire field to given value.

### HasNextFire

`func (o *GetBackups200ResponseBackup) HasNextFire() bool`

HasNextFire returns a boolean if a field has been set.

### SetNextFireNil

`func (o *GetBackups200ResponseBackup) SetNextFireNil(b bool)`

 SetNextFireNil sets the value for NextFire to be an explicit nil

### UnsetNextFire
`func (o *GetBackups200ResponseBackup) UnsetNextFire()`

UnsetNextFire ensures that no value is present for NextFire, not even an explicit nil
### GetLastStatus

`func (o *GetBackups200ResponseBackup) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *GetBackups200ResponseBackup) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *GetBackups200ResponseBackup) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *GetBackups200ResponseBackup) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### SetLastStatusNil

`func (o *GetBackups200ResponseBackup) SetLastStatusNil(b bool)`

 SetLastStatusNil sets the value for LastStatus to be an explicit nil

### UnsetLastStatus
`func (o *GetBackups200ResponseBackup) UnsetLastStatus()`

UnsetLastStatus ensures that no value is present for LastStatus, not even an explicit nil
### GetLastResult

`func (o *GetBackups200ResponseBackup) GetLastResult() GetBackups200ResponseBackupLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *GetBackups200ResponseBackup) GetLastResultOk() (*GetBackups200ResponseBackupLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *GetBackups200ResponseBackup) SetLastResult(v GetBackups200ResponseBackupLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *GetBackups200ResponseBackup) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetStats

`func (o *GetBackups200ResponseBackup) GetStats() GetBackups200ResponseBackupStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetBackups200ResponseBackup) GetStatsOk() (*GetBackups200ResponseBackupStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetBackups200ResponseBackup) SetStats(v GetBackups200ResponseBackupStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetBackups200ResponseBackup) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetEnabled

`func (o *GetBackups200ResponseBackup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetBackups200ResponseBackup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetBackups200ResponseBackup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetBackups200ResponseBackup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetBackups200ResponseBackup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetBackups200ResponseBackup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetBackups200ResponseBackup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetBackups200ResponseBackup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetBackups200ResponseBackup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetBackups200ResponseBackup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetBackups200ResponseBackup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetBackups200ResponseBackup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



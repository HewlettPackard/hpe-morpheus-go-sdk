# UpdateBackups200ResponseAllOfBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**LocationType** | Pointer to **string** | Source Type (instance, server, storage) | [optional] 
**Instance** | Pointer to [**UpdateBackups200ResponseAllOfBackupInstance**](UpdateBackups200ResponseAllOfBackupInstance.md) |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**Job** | Pointer to [**UpdateBackups200ResponseAllOfBackupJob**](UpdateBackups200ResponseAllOfBackupJob.md) |  | [optional] 
**Schedule** | Pointer to [**UpdateBackups200ResponseAllOfBackupSchedule**](UpdateBackups200ResponseAllOfBackupSchedule.md) |  | [optional] 
**RetentionCount** | Pointer to **NullableInt64** |  | [optional] 
**BackupType** | Pointer to [**UpdateBackups200ResponseAllOfBackupBackupType**](UpdateBackups200ResponseAllOfBackupBackupType.md) |  | [optional] 
**StorageProvider** | Pointer to [**UpdateBackups200ResponseAllOfBackupStorageProvider**](UpdateBackups200ResponseAllOfBackupStorageProvider.md) |  | [optional] 
**BackupProvider** | Pointer to [**UpdateBackups200ResponseAllOfBackupBackupProvider**](UpdateBackups200ResponseAllOfBackupBackupProvider.md) |  | [optional] 
**BackupRespository** | Pointer to [**UpdateBackups200ResponseAllOfBackupBackupRespository**](UpdateBackups200ResponseAllOfBackupBackupRespository.md) |  | [optional] 
**CronExpression** | Pointer to **NullableString** | Cron Expression | [optional] 
**NextFire** | Pointer to **NullableTime** | Next Fire | [optional] 
**LastStatus** | Pointer to **NullableString** | Last Status | [optional] 
**LastResult** | Pointer to [**UpdateBackups200ResponseAllOfBackupLastResult**](UpdateBackups200ResponseAllOfBackupLastResult.md) |  | [optional] 
**Stats** | Pointer to [**UpdateBackups200ResponseAllOfBackupStats**](UpdateBackups200ResponseAllOfBackupStats.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Methods

### NewUpdateBackups200ResponseAllOfBackup

`func NewUpdateBackups200ResponseAllOfBackup() *UpdateBackups200ResponseAllOfBackup`

NewUpdateBackups200ResponseAllOfBackup instantiates a new UpdateBackups200ResponseAllOfBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBackups200ResponseAllOfBackupWithDefaults

`func NewUpdateBackups200ResponseAllOfBackupWithDefaults() *UpdateBackups200ResponseAllOfBackup`

NewUpdateBackups200ResponseAllOfBackupWithDefaults instantiates a new UpdateBackups200ResponseAllOfBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateBackups200ResponseAllOfBackup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateBackups200ResponseAllOfBackup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateBackups200ResponseAllOfBackup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateBackups200ResponseAllOfBackup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBackups200ResponseAllOfBackup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBackups200ResponseAllOfBackup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocationType

`func (o *UpdateBackups200ResponseAllOfBackup) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *UpdateBackups200ResponseAllOfBackup) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *UpdateBackups200ResponseAllOfBackup) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetInstance

`func (o *UpdateBackups200ResponseAllOfBackup) GetInstance() UpdateBackups200ResponseAllOfBackupInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetInstanceOk() (*UpdateBackups200ResponseAllOfBackupInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *UpdateBackups200ResponseAllOfBackup) SetInstance(v UpdateBackups200ResponseAllOfBackupInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *UpdateBackups200ResponseAllOfBackup) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerId

`func (o *UpdateBackups200ResponseAllOfBackup) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *UpdateBackups200ResponseAllOfBackup) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *UpdateBackups200ResponseAllOfBackup) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *UpdateBackups200ResponseAllOfBackup) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *UpdateBackups200ResponseAllOfBackup) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetJob

`func (o *UpdateBackups200ResponseAllOfBackup) GetJob() UpdateBackups200ResponseAllOfBackupJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetJobOk() (*UpdateBackups200ResponseAllOfBackupJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *UpdateBackups200ResponseAllOfBackup) SetJob(v UpdateBackups200ResponseAllOfBackupJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *UpdateBackups200ResponseAllOfBackup) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetSchedule

`func (o *UpdateBackups200ResponseAllOfBackup) GetSchedule() UpdateBackups200ResponseAllOfBackupSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetScheduleOk() (*UpdateBackups200ResponseAllOfBackupSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdateBackups200ResponseAllOfBackup) SetSchedule(v UpdateBackups200ResponseAllOfBackupSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpdateBackups200ResponseAllOfBackup) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *UpdateBackups200ResponseAllOfBackup) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *UpdateBackups200ResponseAllOfBackup) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *UpdateBackups200ResponseAllOfBackup) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### SetRetentionCountNil

`func (o *UpdateBackups200ResponseAllOfBackup) SetRetentionCountNil(b bool)`

 SetRetentionCountNil sets the value for RetentionCount to be an explicit nil

### UnsetRetentionCount
`func (o *UpdateBackups200ResponseAllOfBackup) UnsetRetentionCount()`

UnsetRetentionCount ensures that no value is present for RetentionCount, not even an explicit nil
### GetBackupType

`func (o *UpdateBackups200ResponseAllOfBackup) GetBackupType() UpdateBackups200ResponseAllOfBackupBackupType`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetBackupTypeOk() (*UpdateBackups200ResponseAllOfBackupBackupType, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *UpdateBackups200ResponseAllOfBackup) SetBackupType(v UpdateBackups200ResponseAllOfBackupBackupType)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *UpdateBackups200ResponseAllOfBackup) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetStorageProvider

`func (o *UpdateBackups200ResponseAllOfBackup) GetStorageProvider() UpdateBackups200ResponseAllOfBackupStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetStorageProviderOk() (*UpdateBackups200ResponseAllOfBackupStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *UpdateBackups200ResponseAllOfBackup) SetStorageProvider(v UpdateBackups200ResponseAllOfBackupStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *UpdateBackups200ResponseAllOfBackup) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetBackupProvider

`func (o *UpdateBackups200ResponseAllOfBackup) GetBackupProvider() UpdateBackups200ResponseAllOfBackupBackupProvider`

GetBackupProvider returns the BackupProvider field if non-nil, zero value otherwise.

### GetBackupProviderOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetBackupProviderOk() (*UpdateBackups200ResponseAllOfBackupBackupProvider, bool)`

GetBackupProviderOk returns a tuple with the BackupProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProvider

`func (o *UpdateBackups200ResponseAllOfBackup) SetBackupProvider(v UpdateBackups200ResponseAllOfBackupBackupProvider)`

SetBackupProvider sets BackupProvider field to given value.

### HasBackupProvider

`func (o *UpdateBackups200ResponseAllOfBackup) HasBackupProvider() bool`

HasBackupProvider returns a boolean if a field has been set.

### GetBackupRespository

`func (o *UpdateBackups200ResponseAllOfBackup) GetBackupRespository() UpdateBackups200ResponseAllOfBackupBackupRespository`

GetBackupRespository returns the BackupRespository field if non-nil, zero value otherwise.

### GetBackupRespositoryOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetBackupRespositoryOk() (*UpdateBackups200ResponseAllOfBackupBackupRespository, bool)`

GetBackupRespositoryOk returns a tuple with the BackupRespository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRespository

`func (o *UpdateBackups200ResponseAllOfBackup) SetBackupRespository(v UpdateBackups200ResponseAllOfBackupBackupRespository)`

SetBackupRespository sets BackupRespository field to given value.

### HasBackupRespository

`func (o *UpdateBackups200ResponseAllOfBackup) HasBackupRespository() bool`

HasBackupRespository returns a boolean if a field has been set.

### GetCronExpression

`func (o *UpdateBackups200ResponseAllOfBackup) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *UpdateBackups200ResponseAllOfBackup) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *UpdateBackups200ResponseAllOfBackup) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### SetCronExpressionNil

`func (o *UpdateBackups200ResponseAllOfBackup) SetCronExpressionNil(b bool)`

 SetCronExpressionNil sets the value for CronExpression to be an explicit nil

### UnsetCronExpression
`func (o *UpdateBackups200ResponseAllOfBackup) UnsetCronExpression()`

UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
### GetNextFire

`func (o *UpdateBackups200ResponseAllOfBackup) GetNextFire() time.Time`

GetNextFire returns the NextFire field if non-nil, zero value otherwise.

### GetNextFireOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetNextFireOk() (*time.Time, bool)`

GetNextFireOk returns a tuple with the NextFire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFire

`func (o *UpdateBackups200ResponseAllOfBackup) SetNextFire(v time.Time)`

SetNextFire sets NextFire field to given value.

### HasNextFire

`func (o *UpdateBackups200ResponseAllOfBackup) HasNextFire() bool`

HasNextFire returns a boolean if a field has been set.

### SetNextFireNil

`func (o *UpdateBackups200ResponseAllOfBackup) SetNextFireNil(b bool)`

 SetNextFireNil sets the value for NextFire to be an explicit nil

### UnsetNextFire
`func (o *UpdateBackups200ResponseAllOfBackup) UnsetNextFire()`

UnsetNextFire ensures that no value is present for NextFire, not even an explicit nil
### GetLastStatus

`func (o *UpdateBackups200ResponseAllOfBackup) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *UpdateBackups200ResponseAllOfBackup) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *UpdateBackups200ResponseAllOfBackup) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### SetLastStatusNil

`func (o *UpdateBackups200ResponseAllOfBackup) SetLastStatusNil(b bool)`

 SetLastStatusNil sets the value for LastStatus to be an explicit nil

### UnsetLastStatus
`func (o *UpdateBackups200ResponseAllOfBackup) UnsetLastStatus()`

UnsetLastStatus ensures that no value is present for LastStatus, not even an explicit nil
### GetLastResult

`func (o *UpdateBackups200ResponseAllOfBackup) GetLastResult() UpdateBackups200ResponseAllOfBackupLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetLastResultOk() (*UpdateBackups200ResponseAllOfBackupLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *UpdateBackups200ResponseAllOfBackup) SetLastResult(v UpdateBackups200ResponseAllOfBackupLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *UpdateBackups200ResponseAllOfBackup) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetStats

`func (o *UpdateBackups200ResponseAllOfBackup) GetStats() UpdateBackups200ResponseAllOfBackupStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetStatsOk() (*UpdateBackups200ResponseAllOfBackupStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *UpdateBackups200ResponseAllOfBackup) SetStats(v UpdateBackups200ResponseAllOfBackupStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *UpdateBackups200ResponseAllOfBackup) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateBackups200ResponseAllOfBackup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateBackups200ResponseAllOfBackup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateBackups200ResponseAllOfBackup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateBackups200ResponseAllOfBackup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateBackups200ResponseAllOfBackup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateBackups200ResponseAllOfBackup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateBackups200ResponseAllOfBackup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateBackups200ResponseAllOfBackup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateBackups200ResponseAllOfBackup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateBackups200ResponseAllOfBackup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



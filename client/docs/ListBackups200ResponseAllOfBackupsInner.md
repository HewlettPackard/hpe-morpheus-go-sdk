# ListBackups200ResponseAllOfBackupsInner

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

### NewListBackups200ResponseAllOfBackupsInner

`func NewListBackups200ResponseAllOfBackupsInner() *ListBackups200ResponseAllOfBackupsInner`

NewListBackups200ResponseAllOfBackupsInner instantiates a new ListBackups200ResponseAllOfBackupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackups200ResponseAllOfBackupsInnerWithDefaults

`func NewListBackups200ResponseAllOfBackupsInnerWithDefaults() *ListBackups200ResponseAllOfBackupsInner`

NewListBackups200ResponseAllOfBackupsInnerWithDefaults instantiates a new ListBackups200ResponseAllOfBackupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBackups200ResponseAllOfBackupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBackups200ResponseAllOfBackupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListBackups200ResponseAllOfBackupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListBackups200ResponseAllOfBackupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListBackups200ResponseAllOfBackupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListBackups200ResponseAllOfBackupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocationType

`func (o *ListBackups200ResponseAllOfBackupsInner) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *ListBackups200ResponseAllOfBackupsInner) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *ListBackups200ResponseAllOfBackupsInner) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetInstance

`func (o *ListBackups200ResponseAllOfBackupsInner) GetInstance() ListBackups200ResponseAllOfBackupsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetInstanceOk() (*ListBackups200ResponseAllOfBackupsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ListBackups200ResponseAllOfBackupsInner) SetInstance(v ListBackups200ResponseAllOfBackupsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ListBackups200ResponseAllOfBackupsInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerId

`func (o *ListBackups200ResponseAllOfBackupsInner) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ListBackups200ResponseAllOfBackupsInner) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ListBackups200ResponseAllOfBackupsInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetJob

`func (o *ListBackups200ResponseAllOfBackupsInner) GetJob() ListBackups200ResponseAllOfBackupsInnerJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetJobOk() (*ListBackups200ResponseAllOfBackupsInnerJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *ListBackups200ResponseAllOfBackupsInner) SetJob(v ListBackups200ResponseAllOfBackupsInnerJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *ListBackups200ResponseAllOfBackupsInner) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetSchedule

`func (o *ListBackups200ResponseAllOfBackupsInner) GetSchedule() ListBackups200ResponseAllOfBackupsInnerSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetScheduleOk() (*ListBackups200ResponseAllOfBackupsInnerSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ListBackups200ResponseAllOfBackupsInner) SetSchedule(v ListBackups200ResponseAllOfBackupsInnerSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ListBackups200ResponseAllOfBackupsInner) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *ListBackups200ResponseAllOfBackupsInner) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *ListBackups200ResponseAllOfBackupsInner) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *ListBackups200ResponseAllOfBackupsInner) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### GetBackupType

`func (o *ListBackups200ResponseAllOfBackupsInner) GetBackupType() ListBackups200ResponseAllOfBackupsInnerBackupType`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetBackupTypeOk() (*ListBackups200ResponseAllOfBackupsInnerBackupType, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *ListBackups200ResponseAllOfBackupsInner) SetBackupType(v ListBackups200ResponseAllOfBackupsInnerBackupType)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *ListBackups200ResponseAllOfBackupsInner) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetStorageProvider

`func (o *ListBackups200ResponseAllOfBackupsInner) GetStorageProvider() ListBackups200ResponseAllOfBackupsInnerStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetStorageProviderOk() (*ListBackups200ResponseAllOfBackupsInnerStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ListBackups200ResponseAllOfBackupsInner) SetStorageProvider(v ListBackups200ResponseAllOfBackupsInnerStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *ListBackups200ResponseAllOfBackupsInner) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetBackupProvider

`func (o *ListBackups200ResponseAllOfBackupsInner) GetBackupProvider() ListBackups200ResponseAllOfBackupsInnerBackupProvider`

GetBackupProvider returns the BackupProvider field if non-nil, zero value otherwise.

### GetBackupProviderOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetBackupProviderOk() (*ListBackups200ResponseAllOfBackupsInnerBackupProvider, bool)`

GetBackupProviderOk returns a tuple with the BackupProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProvider

`func (o *ListBackups200ResponseAllOfBackupsInner) SetBackupProvider(v ListBackups200ResponseAllOfBackupsInnerBackupProvider)`

SetBackupProvider sets BackupProvider field to given value.

### HasBackupProvider

`func (o *ListBackups200ResponseAllOfBackupsInner) HasBackupProvider() bool`

HasBackupProvider returns a boolean if a field has been set.

### GetBackupRespository

`func (o *ListBackups200ResponseAllOfBackupsInner) GetBackupRespository() ListBackups200ResponseAllOfBackupsInnerBackupRespository`

GetBackupRespository returns the BackupRespository field if non-nil, zero value otherwise.

### GetBackupRespositoryOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetBackupRespositoryOk() (*ListBackups200ResponseAllOfBackupsInnerBackupRespository, bool)`

GetBackupRespositoryOk returns a tuple with the BackupRespository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRespository

`func (o *ListBackups200ResponseAllOfBackupsInner) SetBackupRespository(v ListBackups200ResponseAllOfBackupsInnerBackupRespository)`

SetBackupRespository sets BackupRespository field to given value.

### HasBackupRespository

`func (o *ListBackups200ResponseAllOfBackupsInner) HasBackupRespository() bool`

HasBackupRespository returns a boolean if a field has been set.

### GetCronExpression

`func (o *ListBackups200ResponseAllOfBackupsInner) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *ListBackups200ResponseAllOfBackupsInner) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *ListBackups200ResponseAllOfBackupsInner) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetNextFire

`func (o *ListBackups200ResponseAllOfBackupsInner) GetNextFire() time.Time`

GetNextFire returns the NextFire field if non-nil, zero value otherwise.

### GetNextFireOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetNextFireOk() (*time.Time, bool)`

GetNextFireOk returns a tuple with the NextFire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFire

`func (o *ListBackups200ResponseAllOfBackupsInner) SetNextFire(v time.Time)`

SetNextFire sets NextFire field to given value.

### HasNextFire

`func (o *ListBackups200ResponseAllOfBackupsInner) HasNextFire() bool`

HasNextFire returns a boolean if a field has been set.

### GetLastStatus

`func (o *ListBackups200ResponseAllOfBackupsInner) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *ListBackups200ResponseAllOfBackupsInner) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *ListBackups200ResponseAllOfBackupsInner) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetLastResult

`func (o *ListBackups200ResponseAllOfBackupsInner) GetLastResult() ListBackups200ResponseAllOfBackupsInnerLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetLastResultOk() (*ListBackups200ResponseAllOfBackupsInnerLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *ListBackups200ResponseAllOfBackupsInner) SetLastResult(v ListBackups200ResponseAllOfBackupsInnerLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *ListBackups200ResponseAllOfBackupsInner) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetStats

`func (o *ListBackups200ResponseAllOfBackupsInner) GetStats() ListBackups200ResponseAllOfBackupsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetStatsOk() (*ListBackups200ResponseAllOfBackupsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListBackups200ResponseAllOfBackupsInner) SetStats(v ListBackups200ResponseAllOfBackupsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListBackups200ResponseAllOfBackupsInner) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetEnabled

`func (o *ListBackups200ResponseAllOfBackupsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListBackups200ResponseAllOfBackupsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListBackups200ResponseAllOfBackupsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListBackups200ResponseAllOfBackupsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListBackups200ResponseAllOfBackupsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListBackups200ResponseAllOfBackupsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListBackups200ResponseAllOfBackupsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListBackups200ResponseAllOfBackupsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListBackups200ResponseAllOfBackupsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListBackups200ResponseAllOfBackupsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



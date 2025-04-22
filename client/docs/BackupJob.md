# BackupJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Schedule** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerSchedule**](ListBackups200ResponseAllOfBackupsInnerSchedule.md) |  | [optional] 
**RetentionCount** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**BackupProvider** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerBackupProvider**](ListBackups200ResponseAllOfBackupsInnerBackupProvider.md) |  | [optional] 
**BackupRespository** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerBackupRespository**](ListBackups200ResponseAllOfBackupsInnerBackupRespository.md) |  | [optional] 
**CronExpression** | Pointer to **string** | Cron Expression | [optional] 
**NextFire** | Pointer to **time.Time** | Next Fire is the datetime the job will next occur on according to its schedule | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ListBackupJobs200ResponseAllOfJobsInnerAccount**](ListBackupJobs200ResponseAllOfJobsInnerAccount.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 
**Backups** | Pointer to [**[]ListBackupJobs200ResponseAllOfJobsInnerBackupsInner**](ListBackupJobs200ResponseAllOfJobsInnerBackupsInner.md) | Backups associated with this job | [optional] 

## Methods

### NewBackupJob

`func NewBackupJob() *BackupJob`

NewBackupJob instantiates a new BackupJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobWithDefaults

`func NewBackupJobWithDefaults() *BackupJob`

NewBackupJobWithDefaults instantiates a new BackupJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BackupJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BackupJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackupJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *BackupJob) GetSchedule() ListBackups200ResponseAllOfBackupsInnerSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *BackupJob) GetScheduleOk() (*ListBackups200ResponseAllOfBackupsInnerSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *BackupJob) SetSchedule(v ListBackups200ResponseAllOfBackupsInnerSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *BackupJob) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *BackupJob) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *BackupJob) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *BackupJob) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *BackupJob) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### GetExternalId

`func (o *BackupJob) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BackupJob) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BackupJob) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BackupJob) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetBackupProvider

`func (o *BackupJob) GetBackupProvider() ListBackups200ResponseAllOfBackupsInnerBackupProvider`

GetBackupProvider returns the BackupProvider field if non-nil, zero value otherwise.

### GetBackupProviderOk

`func (o *BackupJob) GetBackupProviderOk() (*ListBackups200ResponseAllOfBackupsInnerBackupProvider, bool)`

GetBackupProviderOk returns a tuple with the BackupProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProvider

`func (o *BackupJob) SetBackupProvider(v ListBackups200ResponseAllOfBackupsInnerBackupProvider)`

SetBackupProvider sets BackupProvider field to given value.

### HasBackupProvider

`func (o *BackupJob) HasBackupProvider() bool`

HasBackupProvider returns a boolean if a field has been set.

### GetBackupRespository

`func (o *BackupJob) GetBackupRespository() ListBackups200ResponseAllOfBackupsInnerBackupRespository`

GetBackupRespository returns the BackupRespository field if non-nil, zero value otherwise.

### GetBackupRespositoryOk

`func (o *BackupJob) GetBackupRespositoryOk() (*ListBackups200ResponseAllOfBackupsInnerBackupRespository, bool)`

GetBackupRespositoryOk returns a tuple with the BackupRespository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRespository

`func (o *BackupJob) SetBackupRespository(v ListBackups200ResponseAllOfBackupsInnerBackupRespository)`

SetBackupRespository sets BackupRespository field to given value.

### HasBackupRespository

`func (o *BackupJob) HasBackupRespository() bool`

HasBackupRespository returns a boolean if a field has been set.

### GetCronExpression

`func (o *BackupJob) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *BackupJob) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *BackupJob) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *BackupJob) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetNextFire

`func (o *BackupJob) GetNextFire() time.Time`

GetNextFire returns the NextFire field if non-nil, zero value otherwise.

### GetNextFireOk

`func (o *BackupJob) GetNextFireOk() (*time.Time, bool)`

GetNextFireOk returns a tuple with the NextFire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFire

`func (o *BackupJob) SetNextFire(v time.Time)`

SetNextFire sets NextFire field to given value.

### HasNextFire

`func (o *BackupJob) HasNextFire() bool`

HasNextFire returns a boolean if a field has been set.

### GetSource

`func (o *BackupJob) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BackupJob) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BackupJob) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BackupJob) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVisibility

`func (o *BackupJob) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BackupJob) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BackupJob) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BackupJob) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *BackupJob) GetAccount() ListBackupJobs200ResponseAllOfJobsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BackupJob) GetAccountOk() (*ListBackupJobs200ResponseAllOfJobsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BackupJob) SetAccount(v ListBackupJobs200ResponseAllOfJobsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BackupJob) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEnabled

`func (o *BackupJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BackupJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BackupJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BackupJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *BackupJob) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BackupJob) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BackupJob) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *BackupJob) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *BackupJob) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *BackupJob) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *BackupJob) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *BackupJob) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetBackups

`func (o *BackupJob) GetBackups() []ListBackupJobs200ResponseAllOfJobsInnerBackupsInner`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *BackupJob) GetBackupsOk() (*[]ListBackupJobs200ResponseAllOfJobsInnerBackupsInner, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *BackupJob) SetBackups(v []ListBackupJobs200ResponseAllOfJobsInnerBackupsInner)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *BackupJob) HasBackups() bool`

HasBackups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



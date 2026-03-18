# AddBackupJobs200ResponseAllOfJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Schedule** | Pointer to [**AddBackupJobs200ResponseAllOfJobSchedule**](AddBackupJobs200ResponseAllOfJobSchedule.md) |  | [optional] 
**RetentionCount** | Pointer to **NullableInt64** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**BackupProvider** | Pointer to [**AddBackupJobs200ResponseAllOfJobBackupProvider**](AddBackupJobs200ResponseAllOfJobBackupProvider.md) |  | [optional] 
**BackupRespository** | Pointer to [**AddBackupJobs200ResponseAllOfJobBackupRespository**](AddBackupJobs200ResponseAllOfJobBackupRespository.md) |  | [optional] 
**CronExpression** | Pointer to **NullableString** | Cron Expression | [optional] 
**NextFire** | Pointer to **NullableTime** | Next Fire is the datetime the job will next occur on according to its schedule | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**AddBackupJobs200ResponseAllOfJobAccount**](AddBackupJobs200ResponseAllOfJobAccount.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 
**Backups** | Pointer to [**[]AddBackupJobs200ResponseAllOfJobBackupsInner**](AddBackupJobs200ResponseAllOfJobBackupsInner.md) | Backups associated with this job | [optional] 

## Methods

### NewAddBackupJobs200ResponseAllOfJob

`func NewAddBackupJobs200ResponseAllOfJob() *AddBackupJobs200ResponseAllOfJob`

NewAddBackupJobs200ResponseAllOfJob instantiates a new AddBackupJobs200ResponseAllOfJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBackupJobs200ResponseAllOfJobWithDefaults

`func NewAddBackupJobs200ResponseAllOfJobWithDefaults() *AddBackupJobs200ResponseAllOfJob`

NewAddBackupJobs200ResponseAllOfJobWithDefaults instantiates a new AddBackupJobs200ResponseAllOfJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddBackupJobs200ResponseAllOfJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddBackupJobs200ResponseAllOfJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddBackupJobs200ResponseAllOfJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddBackupJobs200ResponseAllOfJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBackupJobs200ResponseAllOfJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddBackupJobs200ResponseAllOfJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *AddBackupJobs200ResponseAllOfJob) GetSchedule() AddBackupJobs200ResponseAllOfJobSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetScheduleOk() (*AddBackupJobs200ResponseAllOfJobSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AddBackupJobs200ResponseAllOfJob) SetSchedule(v AddBackupJobs200ResponseAllOfJobSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AddBackupJobs200ResponseAllOfJob) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *AddBackupJobs200ResponseAllOfJob) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *AddBackupJobs200ResponseAllOfJob) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *AddBackupJobs200ResponseAllOfJob) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### SetRetentionCountNil

`func (o *AddBackupJobs200ResponseAllOfJob) SetRetentionCountNil(b bool)`

 SetRetentionCountNil sets the value for RetentionCount to be an explicit nil

### UnsetRetentionCount
`func (o *AddBackupJobs200ResponseAllOfJob) UnsetRetentionCount()`

UnsetRetentionCount ensures that no value is present for RetentionCount, not even an explicit nil
### GetExternalId

`func (o *AddBackupJobs200ResponseAllOfJob) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddBackupJobs200ResponseAllOfJob) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddBackupJobs200ResponseAllOfJob) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AddBackupJobs200ResponseAllOfJob) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AddBackupJobs200ResponseAllOfJob) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetBackupProvider

`func (o *AddBackupJobs200ResponseAllOfJob) GetBackupProvider() AddBackupJobs200ResponseAllOfJobBackupProvider`

GetBackupProvider returns the BackupProvider field if non-nil, zero value otherwise.

### GetBackupProviderOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetBackupProviderOk() (*AddBackupJobs200ResponseAllOfJobBackupProvider, bool)`

GetBackupProviderOk returns a tuple with the BackupProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProvider

`func (o *AddBackupJobs200ResponseAllOfJob) SetBackupProvider(v AddBackupJobs200ResponseAllOfJobBackupProvider)`

SetBackupProvider sets BackupProvider field to given value.

### HasBackupProvider

`func (o *AddBackupJobs200ResponseAllOfJob) HasBackupProvider() bool`

HasBackupProvider returns a boolean if a field has been set.

### GetBackupRespository

`func (o *AddBackupJobs200ResponseAllOfJob) GetBackupRespository() AddBackupJobs200ResponseAllOfJobBackupRespository`

GetBackupRespository returns the BackupRespository field if non-nil, zero value otherwise.

### GetBackupRespositoryOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetBackupRespositoryOk() (*AddBackupJobs200ResponseAllOfJobBackupRespository, bool)`

GetBackupRespositoryOk returns a tuple with the BackupRespository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRespository

`func (o *AddBackupJobs200ResponseAllOfJob) SetBackupRespository(v AddBackupJobs200ResponseAllOfJobBackupRespository)`

SetBackupRespository sets BackupRespository field to given value.

### HasBackupRespository

`func (o *AddBackupJobs200ResponseAllOfJob) HasBackupRespository() bool`

HasBackupRespository returns a boolean if a field has been set.

### GetCronExpression

`func (o *AddBackupJobs200ResponseAllOfJob) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *AddBackupJobs200ResponseAllOfJob) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *AddBackupJobs200ResponseAllOfJob) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### SetCronExpressionNil

`func (o *AddBackupJobs200ResponseAllOfJob) SetCronExpressionNil(b bool)`

 SetCronExpressionNil sets the value for CronExpression to be an explicit nil

### UnsetCronExpression
`func (o *AddBackupJobs200ResponseAllOfJob) UnsetCronExpression()`

UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
### GetNextFire

`func (o *AddBackupJobs200ResponseAllOfJob) GetNextFire() time.Time`

GetNextFire returns the NextFire field if non-nil, zero value otherwise.

### GetNextFireOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetNextFireOk() (*time.Time, bool)`

GetNextFireOk returns a tuple with the NextFire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFire

`func (o *AddBackupJobs200ResponseAllOfJob) SetNextFire(v time.Time)`

SetNextFire sets NextFire field to given value.

### HasNextFire

`func (o *AddBackupJobs200ResponseAllOfJob) HasNextFire() bool`

HasNextFire returns a boolean if a field has been set.

### SetNextFireNil

`func (o *AddBackupJobs200ResponseAllOfJob) SetNextFireNil(b bool)`

 SetNextFireNil sets the value for NextFire to be an explicit nil

### UnsetNextFire
`func (o *AddBackupJobs200ResponseAllOfJob) UnsetNextFire()`

UnsetNextFire ensures that no value is present for NextFire, not even an explicit nil
### GetSource

`func (o *AddBackupJobs200ResponseAllOfJob) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AddBackupJobs200ResponseAllOfJob) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AddBackupJobs200ResponseAllOfJob) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVisibility

`func (o *AddBackupJobs200ResponseAllOfJob) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddBackupJobs200ResponseAllOfJob) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddBackupJobs200ResponseAllOfJob) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *AddBackupJobs200ResponseAllOfJob) GetAccount() AddBackupJobs200ResponseAllOfJobAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetAccountOk() (*AddBackupJobs200ResponseAllOfJobAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddBackupJobs200ResponseAllOfJob) SetAccount(v AddBackupJobs200ResponseAllOfJobAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddBackupJobs200ResponseAllOfJob) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEnabled

`func (o *AddBackupJobs200ResponseAllOfJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddBackupJobs200ResponseAllOfJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddBackupJobs200ResponseAllOfJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddBackupJobs200ResponseAllOfJob) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddBackupJobs200ResponseAllOfJob) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddBackupJobs200ResponseAllOfJob) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddBackupJobs200ResponseAllOfJob) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddBackupJobs200ResponseAllOfJob) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddBackupJobs200ResponseAllOfJob) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetBackups

`func (o *AddBackupJobs200ResponseAllOfJob) GetBackups() []AddBackupJobs200ResponseAllOfJobBackupsInner`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *AddBackupJobs200ResponseAllOfJob) GetBackupsOk() (*[]AddBackupJobs200ResponseAllOfJobBackupsInner, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *AddBackupJobs200ResponseAllOfJob) SetBackups(v []AddBackupJobs200ResponseAllOfJobBackupsInner)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *AddBackupJobs200ResponseAllOfJob) HasBackups() bool`

HasBackups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



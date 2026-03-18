# GetBackupJobs200ResponseJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Schedule** | Pointer to [**GetBackupJobs200ResponseJobSchedule**](GetBackupJobs200ResponseJobSchedule.md) |  | [optional] 
**RetentionCount** | Pointer to **NullableInt64** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**BackupProvider** | Pointer to [**GetBackupJobs200ResponseJobBackupProvider**](GetBackupJobs200ResponseJobBackupProvider.md) |  | [optional] 
**BackupRespository** | Pointer to [**GetBackupJobs200ResponseJobBackupRespository**](GetBackupJobs200ResponseJobBackupRespository.md) |  | [optional] 
**CronExpression** | Pointer to **NullableString** | Cron Expression | [optional] 
**NextFire** | Pointer to **NullableTime** | Next Fire is the datetime the job will next occur on according to its schedule | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**GetBackupJobs200ResponseJobAccount**](GetBackupJobs200ResponseJobAccount.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 
**Backups** | Pointer to [**[]GetBackupJobs200ResponseJobBackupsInner**](GetBackupJobs200ResponseJobBackupsInner.md) | Backups associated with this job | [optional] 

## Methods

### NewGetBackupJobs200ResponseJob

`func NewGetBackupJobs200ResponseJob() *GetBackupJobs200ResponseJob`

NewGetBackupJobs200ResponseJob instantiates a new GetBackupJobs200ResponseJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBackupJobs200ResponseJobWithDefaults

`func NewGetBackupJobs200ResponseJobWithDefaults() *GetBackupJobs200ResponseJob`

NewGetBackupJobs200ResponseJobWithDefaults instantiates a new GetBackupJobs200ResponseJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetBackupJobs200ResponseJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBackupJobs200ResponseJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBackupJobs200ResponseJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetBackupJobs200ResponseJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetBackupJobs200ResponseJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetBackupJobs200ResponseJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetBackupJobs200ResponseJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetBackupJobs200ResponseJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *GetBackupJobs200ResponseJob) GetSchedule() GetBackupJobs200ResponseJobSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *GetBackupJobs200ResponseJob) GetScheduleOk() (*GetBackupJobs200ResponseJobSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *GetBackupJobs200ResponseJob) SetSchedule(v GetBackupJobs200ResponseJobSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *GetBackupJobs200ResponseJob) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *GetBackupJobs200ResponseJob) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *GetBackupJobs200ResponseJob) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *GetBackupJobs200ResponseJob) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *GetBackupJobs200ResponseJob) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### SetRetentionCountNil

`func (o *GetBackupJobs200ResponseJob) SetRetentionCountNil(b bool)`

 SetRetentionCountNil sets the value for RetentionCount to be an explicit nil

### UnsetRetentionCount
`func (o *GetBackupJobs200ResponseJob) UnsetRetentionCount()`

UnsetRetentionCount ensures that no value is present for RetentionCount, not even an explicit nil
### GetExternalId

`func (o *GetBackupJobs200ResponseJob) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetBackupJobs200ResponseJob) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetBackupJobs200ResponseJob) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetBackupJobs200ResponseJob) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetBackupJobs200ResponseJob) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetBackupJobs200ResponseJob) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetBackupProvider

`func (o *GetBackupJobs200ResponseJob) GetBackupProvider() GetBackupJobs200ResponseJobBackupProvider`

GetBackupProvider returns the BackupProvider field if non-nil, zero value otherwise.

### GetBackupProviderOk

`func (o *GetBackupJobs200ResponseJob) GetBackupProviderOk() (*GetBackupJobs200ResponseJobBackupProvider, bool)`

GetBackupProviderOk returns a tuple with the BackupProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProvider

`func (o *GetBackupJobs200ResponseJob) SetBackupProvider(v GetBackupJobs200ResponseJobBackupProvider)`

SetBackupProvider sets BackupProvider field to given value.

### HasBackupProvider

`func (o *GetBackupJobs200ResponseJob) HasBackupProvider() bool`

HasBackupProvider returns a boolean if a field has been set.

### GetBackupRespository

`func (o *GetBackupJobs200ResponseJob) GetBackupRespository() GetBackupJobs200ResponseJobBackupRespository`

GetBackupRespository returns the BackupRespository field if non-nil, zero value otherwise.

### GetBackupRespositoryOk

`func (o *GetBackupJobs200ResponseJob) GetBackupRespositoryOk() (*GetBackupJobs200ResponseJobBackupRespository, bool)`

GetBackupRespositoryOk returns a tuple with the BackupRespository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRespository

`func (o *GetBackupJobs200ResponseJob) SetBackupRespository(v GetBackupJobs200ResponseJobBackupRespository)`

SetBackupRespository sets BackupRespository field to given value.

### HasBackupRespository

`func (o *GetBackupJobs200ResponseJob) HasBackupRespository() bool`

HasBackupRespository returns a boolean if a field has been set.

### GetCronExpression

`func (o *GetBackupJobs200ResponseJob) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *GetBackupJobs200ResponseJob) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *GetBackupJobs200ResponseJob) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *GetBackupJobs200ResponseJob) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### SetCronExpressionNil

`func (o *GetBackupJobs200ResponseJob) SetCronExpressionNil(b bool)`

 SetCronExpressionNil sets the value for CronExpression to be an explicit nil

### UnsetCronExpression
`func (o *GetBackupJobs200ResponseJob) UnsetCronExpression()`

UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
### GetNextFire

`func (o *GetBackupJobs200ResponseJob) GetNextFire() time.Time`

GetNextFire returns the NextFire field if non-nil, zero value otherwise.

### GetNextFireOk

`func (o *GetBackupJobs200ResponseJob) GetNextFireOk() (*time.Time, bool)`

GetNextFireOk returns a tuple with the NextFire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFire

`func (o *GetBackupJobs200ResponseJob) SetNextFire(v time.Time)`

SetNextFire sets NextFire field to given value.

### HasNextFire

`func (o *GetBackupJobs200ResponseJob) HasNextFire() bool`

HasNextFire returns a boolean if a field has been set.

### SetNextFireNil

`func (o *GetBackupJobs200ResponseJob) SetNextFireNil(b bool)`

 SetNextFireNil sets the value for NextFire to be an explicit nil

### UnsetNextFire
`func (o *GetBackupJobs200ResponseJob) UnsetNextFire()`

UnsetNextFire ensures that no value is present for NextFire, not even an explicit nil
### GetSource

`func (o *GetBackupJobs200ResponseJob) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetBackupJobs200ResponseJob) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetBackupJobs200ResponseJob) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetBackupJobs200ResponseJob) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVisibility

`func (o *GetBackupJobs200ResponseJob) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetBackupJobs200ResponseJob) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetBackupJobs200ResponseJob) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetBackupJobs200ResponseJob) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *GetBackupJobs200ResponseJob) GetAccount() GetBackupJobs200ResponseJobAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetBackupJobs200ResponseJob) GetAccountOk() (*GetBackupJobs200ResponseJobAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetBackupJobs200ResponseJob) SetAccount(v GetBackupJobs200ResponseJobAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetBackupJobs200ResponseJob) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEnabled

`func (o *GetBackupJobs200ResponseJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetBackupJobs200ResponseJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetBackupJobs200ResponseJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetBackupJobs200ResponseJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetBackupJobs200ResponseJob) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetBackupJobs200ResponseJob) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetBackupJobs200ResponseJob) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetBackupJobs200ResponseJob) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetBackupJobs200ResponseJob) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetBackupJobs200ResponseJob) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetBackupJobs200ResponseJob) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetBackupJobs200ResponseJob) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetBackups

`func (o *GetBackupJobs200ResponseJob) GetBackups() []GetBackupJobs200ResponseJobBackupsInner`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *GetBackupJobs200ResponseJob) GetBackupsOk() (*[]GetBackupJobs200ResponseJobBackupsInner, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *GetBackupJobs200ResponseJob) SetBackups(v []GetBackupJobs200ResponseJobBackupsInner)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *GetBackupJobs200ResponseJob) HasBackups() bool`

HasBackups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



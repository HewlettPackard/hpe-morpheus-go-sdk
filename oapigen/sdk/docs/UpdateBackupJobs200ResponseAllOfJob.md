# UpdateBackupJobs200ResponseAllOfJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Schedule** | Pointer to [**UpdateBackupJobs200ResponseAllOfJobSchedule**](UpdateBackupJobs200ResponseAllOfJobSchedule.md) |  | [optional] 
**RetentionCount** | Pointer to **NullableInt64** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**BackupProvider** | Pointer to [**UpdateBackupJobs200ResponseAllOfJobBackupProvider**](UpdateBackupJobs200ResponseAllOfJobBackupProvider.md) |  | [optional] 
**BackupRespository** | Pointer to [**UpdateBackupJobs200ResponseAllOfJobBackupRespository**](UpdateBackupJobs200ResponseAllOfJobBackupRespository.md) |  | [optional] 
**CronExpression** | Pointer to **NullableString** | Cron Expression | [optional] 
**NextFire** | Pointer to **NullableTime** | Next Fire is the datetime the job will next occur on according to its schedule | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**UpdateBackupJobs200ResponseAllOfJobAccount**](UpdateBackupJobs200ResponseAllOfJobAccount.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 
**Backups** | Pointer to [**[]UpdateBackupJobs200ResponseAllOfJobBackupsInner**](UpdateBackupJobs200ResponseAllOfJobBackupsInner.md) | Backups associated with this job | [optional] 

## Methods

### NewUpdateBackupJobs200ResponseAllOfJob

`func NewUpdateBackupJobs200ResponseAllOfJob() *UpdateBackupJobs200ResponseAllOfJob`

NewUpdateBackupJobs200ResponseAllOfJob instantiates a new UpdateBackupJobs200ResponseAllOfJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBackupJobs200ResponseAllOfJobWithDefaults

`func NewUpdateBackupJobs200ResponseAllOfJobWithDefaults() *UpdateBackupJobs200ResponseAllOfJob`

NewUpdateBackupJobs200ResponseAllOfJobWithDefaults instantiates a new UpdateBackupJobs200ResponseAllOfJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetSchedule() UpdateBackupJobs200ResponseAllOfJobSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetScheduleOk() (*UpdateBackupJobs200ResponseAllOfJobSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetSchedule(v UpdateBackupJobs200ResponseAllOfJobSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### SetRetentionCountNil

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetRetentionCountNil(b bool)`

 SetRetentionCountNil sets the value for RetentionCount to be an explicit nil

### UnsetRetentionCount
`func (o *UpdateBackupJobs200ResponseAllOfJob) UnsetRetentionCount()`

UnsetRetentionCount ensures that no value is present for RetentionCount, not even an explicit nil
### GetExternalId

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateBackupJobs200ResponseAllOfJob) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetBackupProvider

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetBackupProvider() UpdateBackupJobs200ResponseAllOfJobBackupProvider`

GetBackupProvider returns the BackupProvider field if non-nil, zero value otherwise.

### GetBackupProviderOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetBackupProviderOk() (*UpdateBackupJobs200ResponseAllOfJobBackupProvider, bool)`

GetBackupProviderOk returns a tuple with the BackupProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProvider

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetBackupProvider(v UpdateBackupJobs200ResponseAllOfJobBackupProvider)`

SetBackupProvider sets BackupProvider field to given value.

### HasBackupProvider

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasBackupProvider() bool`

HasBackupProvider returns a boolean if a field has been set.

### GetBackupRespository

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetBackupRespository() UpdateBackupJobs200ResponseAllOfJobBackupRespository`

GetBackupRespository returns the BackupRespository field if non-nil, zero value otherwise.

### GetBackupRespositoryOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetBackupRespositoryOk() (*UpdateBackupJobs200ResponseAllOfJobBackupRespository, bool)`

GetBackupRespositoryOk returns a tuple with the BackupRespository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRespository

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetBackupRespository(v UpdateBackupJobs200ResponseAllOfJobBackupRespository)`

SetBackupRespository sets BackupRespository field to given value.

### HasBackupRespository

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasBackupRespository() bool`

HasBackupRespository returns a boolean if a field has been set.

### GetCronExpression

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### SetCronExpressionNil

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetCronExpressionNil(b bool)`

 SetCronExpressionNil sets the value for CronExpression to be an explicit nil

### UnsetCronExpression
`func (o *UpdateBackupJobs200ResponseAllOfJob) UnsetCronExpression()`

UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
### GetNextFire

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetNextFire() time.Time`

GetNextFire returns the NextFire field if non-nil, zero value otherwise.

### GetNextFireOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetNextFireOk() (*time.Time, bool)`

GetNextFireOk returns a tuple with the NextFire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFire

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetNextFire(v time.Time)`

SetNextFire sets NextFire field to given value.

### HasNextFire

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasNextFire() bool`

HasNextFire returns a boolean if a field has been set.

### SetNextFireNil

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetNextFireNil(b bool)`

 SetNextFireNil sets the value for NextFire to be an explicit nil

### UnsetNextFire
`func (o *UpdateBackupJobs200ResponseAllOfJob) UnsetNextFire()`

UnsetNextFire ensures that no value is present for NextFire, not even an explicit nil
### GetSource

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetAccount() UpdateBackupJobs200ResponseAllOfJobAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetAccountOk() (*UpdateBackupJobs200ResponseAllOfJobAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetAccount(v UpdateBackupJobs200ResponseAllOfJobAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetBackups

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetBackups() []UpdateBackupJobs200ResponseAllOfJobBackupsInner`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *UpdateBackupJobs200ResponseAllOfJob) GetBackupsOk() (*[]UpdateBackupJobs200ResponseAllOfJobBackupsInner, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *UpdateBackupJobs200ResponseAllOfJob) SetBackups(v []UpdateBackupJobs200ResponseAllOfJobBackupsInner)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *UpdateBackupJobs200ResponseAllOfJob) HasBackups() bool`

HasBackups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



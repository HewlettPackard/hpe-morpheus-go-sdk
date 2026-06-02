# BackupTypeInstanceBackupJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyntheticFullEnabled** | Pointer to **bool** | Enable synthetic full backups on this backup jobAction. Only applies to &#x60;kvm&#x60; backup type. | [optional] 
**SyntheticFullSchedule** | Pointer to **int64** | the ID of the execute schedule for the synthetic full backup to be created. Only applies to &#x60;kvm&#x60; backup type. | [optional] 

## Methods

### NewBackupTypeInstanceBackupJob

`func NewBackupTypeInstanceBackupJob() *BackupTypeInstanceBackupJob`

NewBackupTypeInstanceBackupJob instantiates a new BackupTypeInstanceBackupJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetSyntheticFullEnabled

`func (o *BackupTypeInstanceBackupJob) GetSyntheticFullEnabled() bool`

GetSyntheticFullEnabled returns the SyntheticFullEnabled field if non-nil, zero value otherwise.

### GetSyntheticFullEnabledOk

`func (o *BackupTypeInstanceBackupJob) GetSyntheticFullEnabledOk() (*bool, bool)`

GetSyntheticFullEnabledOk returns a tuple with the SyntheticFullEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticFullEnabled

`func (o *BackupTypeInstanceBackupJob) SetSyntheticFullEnabled(v bool)`

SetSyntheticFullEnabled sets SyntheticFullEnabled field to given value.

### HasSyntheticFullEnabled

`func (o *BackupTypeInstanceBackupJob) HasSyntheticFullEnabled() bool`

HasSyntheticFullEnabled returns a boolean if a field has been set.

### GetSyntheticFullSchedule

`func (o *BackupTypeInstanceBackupJob) GetSyntheticFullSchedule() int64`

GetSyntheticFullSchedule returns the SyntheticFullSchedule field if non-nil, zero value otherwise.

### GetSyntheticFullScheduleOk

`func (o *BackupTypeInstanceBackupJob) GetSyntheticFullScheduleOk() (*int64, bool)`

GetSyntheticFullScheduleOk returns a tuple with the SyntheticFullSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticFullSchedule

`func (o *BackupTypeInstanceBackupJob) SetSyntheticFullSchedule(v int64)`

SetSyntheticFullSchedule sets SyntheticFullSchedule field to given value.

### HasSyntheticFullSchedule

`func (o *BackupTypeInstanceBackupJob) HasSyntheticFullSchedule() bool`

HasSyntheticFullSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



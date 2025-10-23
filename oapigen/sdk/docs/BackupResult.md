# BackupResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup Result ID | [optional] 
**Backup** | Pointer to [**ListBackupJobs200ResponseAllOfJobsInnerBackupsInner**](ListBackupJobs200ResponseAllOfJobsInnerBackupsInner.md) |  | [optional] 
**BackupSetId** | Pointer to **NullableString** |  | [optional] 
**InstanceId** | Pointer to **NullableInt64** |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**ServerId** | Pointer to **NullableInt64** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**DurationMillis** | Pointer to **NullableInt64** |  | [optional] 
**SizeInBytes** | Pointer to **NullableInt64** |  | [optional] 
**SizeInMb** | Pointer to **NullableInt64** |  | [optional] 
**VolumePath** | Pointer to **NullableString** |  | [optional] 
**ResultArchive** | Pointer to **NullableString** |  | [optional] 
**ResultPath** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**SnapshotId** | Pointer to **NullableString** |  | [optional] 
**SnapshotExternalId** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**ListBackupResults200ResponseAllOfResultsInnerCreatedBy**](ListBackupResults200ResponseAllOfResultsInnerCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Methods

### NewBackupResult

`func NewBackupResult() *BackupResult`

NewBackupResult instantiates a new BackupResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupResultWithDefaults

`func NewBackupResultWithDefaults() *BackupResult`

NewBackupResultWithDefaults instantiates a new BackupResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BackupResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBackup

`func (o *BackupResult) GetBackup() ListBackupJobs200ResponseAllOfJobsInnerBackupsInner`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *BackupResult) GetBackupOk() (*ListBackupJobs200ResponseAllOfJobsInnerBackupsInner, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *BackupResult) SetBackup(v ListBackupJobs200ResponseAllOfJobsInnerBackupsInner)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *BackupResult) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetBackupSetId

`func (o *BackupResult) GetBackupSetId() string`

GetBackupSetId returns the BackupSetId field if non-nil, zero value otherwise.

### GetBackupSetIdOk

`func (o *BackupResult) GetBackupSetIdOk() (*string, bool)`

GetBackupSetIdOk returns a tuple with the BackupSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSetId

`func (o *BackupResult) SetBackupSetId(v string)`

SetBackupSetId sets BackupSetId field to given value.

### HasBackupSetId

`func (o *BackupResult) HasBackupSetId() bool`

HasBackupSetId returns a boolean if a field has been set.

### SetBackupSetIdNil

`func (o *BackupResult) SetBackupSetIdNil(b bool)`

 SetBackupSetIdNil sets the value for BackupSetId to be an explicit nil

### UnsetBackupSetId
`func (o *BackupResult) UnsetBackupSetId()`

UnsetBackupSetId ensures that no value is present for BackupSetId, not even an explicit nil
### GetInstanceId

`func (o *BackupResult) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *BackupResult) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *BackupResult) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *BackupResult) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *BackupResult) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *BackupResult) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetContainerId

`func (o *BackupResult) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *BackupResult) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *BackupResult) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *BackupResult) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *BackupResult) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *BackupResult) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetServerId

`func (o *BackupResult) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *BackupResult) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *BackupResult) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *BackupResult) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *BackupResult) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *BackupResult) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetStatus

`func (o *BackupResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BackupResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BackupResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetErrorMessage

`func (o *BackupResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BackupResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BackupResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BackupResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *BackupResult) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *BackupResult) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStartDate

`func (o *BackupResult) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BackupResult) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BackupResult) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BackupResult) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *BackupResult) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *BackupResult) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *BackupResult) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BackupResult) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BackupResult) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BackupResult) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *BackupResult) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *BackupResult) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDurationMillis

`func (o *BackupResult) GetDurationMillis() int64`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *BackupResult) GetDurationMillisOk() (*int64, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *BackupResult) SetDurationMillis(v int64)`

SetDurationMillis sets DurationMillis field to given value.

### HasDurationMillis

`func (o *BackupResult) HasDurationMillis() bool`

HasDurationMillis returns a boolean if a field has been set.

### SetDurationMillisNil

`func (o *BackupResult) SetDurationMillisNil(b bool)`

 SetDurationMillisNil sets the value for DurationMillis to be an explicit nil

### UnsetDurationMillis
`func (o *BackupResult) UnsetDurationMillis()`

UnsetDurationMillis ensures that no value is present for DurationMillis, not even an explicit nil
### GetSizeInBytes

`func (o *BackupResult) GetSizeInBytes() int64`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *BackupResult) GetSizeInBytesOk() (*int64, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *BackupResult) SetSizeInBytes(v int64)`

SetSizeInBytes sets SizeInBytes field to given value.

### HasSizeInBytes

`func (o *BackupResult) HasSizeInBytes() bool`

HasSizeInBytes returns a boolean if a field has been set.

### SetSizeInBytesNil

`func (o *BackupResult) SetSizeInBytesNil(b bool)`

 SetSizeInBytesNil sets the value for SizeInBytes to be an explicit nil

### UnsetSizeInBytes
`func (o *BackupResult) UnsetSizeInBytes()`

UnsetSizeInBytes ensures that no value is present for SizeInBytes, not even an explicit nil
### GetSizeInMb

`func (o *BackupResult) GetSizeInMb() int64`

GetSizeInMb returns the SizeInMb field if non-nil, zero value otherwise.

### GetSizeInMbOk

`func (o *BackupResult) GetSizeInMbOk() (*int64, bool)`

GetSizeInMbOk returns a tuple with the SizeInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInMb

`func (o *BackupResult) SetSizeInMb(v int64)`

SetSizeInMb sets SizeInMb field to given value.

### HasSizeInMb

`func (o *BackupResult) HasSizeInMb() bool`

HasSizeInMb returns a boolean if a field has been set.

### SetSizeInMbNil

`func (o *BackupResult) SetSizeInMbNil(b bool)`

 SetSizeInMbNil sets the value for SizeInMb to be an explicit nil

### UnsetSizeInMb
`func (o *BackupResult) UnsetSizeInMb()`

UnsetSizeInMb ensures that no value is present for SizeInMb, not even an explicit nil
### GetVolumePath

`func (o *BackupResult) GetVolumePath() string`

GetVolumePath returns the VolumePath field if non-nil, zero value otherwise.

### GetVolumePathOk

`func (o *BackupResult) GetVolumePathOk() (*string, bool)`

GetVolumePathOk returns a tuple with the VolumePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePath

`func (o *BackupResult) SetVolumePath(v string)`

SetVolumePath sets VolumePath field to given value.

### HasVolumePath

`func (o *BackupResult) HasVolumePath() bool`

HasVolumePath returns a boolean if a field has been set.

### SetVolumePathNil

`func (o *BackupResult) SetVolumePathNil(b bool)`

 SetVolumePathNil sets the value for VolumePath to be an explicit nil

### UnsetVolumePath
`func (o *BackupResult) UnsetVolumePath()`

UnsetVolumePath ensures that no value is present for VolumePath, not even an explicit nil
### GetResultArchive

`func (o *BackupResult) GetResultArchive() string`

GetResultArchive returns the ResultArchive field if non-nil, zero value otherwise.

### GetResultArchiveOk

`func (o *BackupResult) GetResultArchiveOk() (*string, bool)`

GetResultArchiveOk returns a tuple with the ResultArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultArchive

`func (o *BackupResult) SetResultArchive(v string)`

SetResultArchive sets ResultArchive field to given value.

### HasResultArchive

`func (o *BackupResult) HasResultArchive() bool`

HasResultArchive returns a boolean if a field has been set.

### SetResultArchiveNil

`func (o *BackupResult) SetResultArchiveNil(b bool)`

 SetResultArchiveNil sets the value for ResultArchive to be an explicit nil

### UnsetResultArchive
`func (o *BackupResult) UnsetResultArchive()`

UnsetResultArchive ensures that no value is present for ResultArchive, not even an explicit nil
### GetResultPath

`func (o *BackupResult) GetResultPath() string`

GetResultPath returns the ResultPath field if non-nil, zero value otherwise.

### GetResultPathOk

`func (o *BackupResult) GetResultPathOk() (*string, bool)`

GetResultPathOk returns a tuple with the ResultPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultPath

`func (o *BackupResult) SetResultPath(v string)`

SetResultPath sets ResultPath field to given value.

### HasResultPath

`func (o *BackupResult) HasResultPath() bool`

HasResultPath returns a boolean if a field has been set.

### SetResultPathNil

`func (o *BackupResult) SetResultPathNil(b bool)`

 SetResultPathNil sets the value for ResultPath to be an explicit nil

### UnsetResultPath
`func (o *BackupResult) UnsetResultPath()`

UnsetResultPath ensures that no value is present for ResultPath, not even an explicit nil
### GetExternalId

`func (o *BackupResult) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BackupResult) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BackupResult) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BackupResult) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *BackupResult) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *BackupResult) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetSnapshotId

`func (o *BackupResult) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *BackupResult) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *BackupResult) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *BackupResult) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *BackupResult) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *BackupResult) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetSnapshotExternalId

`func (o *BackupResult) GetSnapshotExternalId() string`

GetSnapshotExternalId returns the SnapshotExternalId field if non-nil, zero value otherwise.

### GetSnapshotExternalIdOk

`func (o *BackupResult) GetSnapshotExternalIdOk() (*string, bool)`

GetSnapshotExternalIdOk returns a tuple with the SnapshotExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotExternalId

`func (o *BackupResult) SetSnapshotExternalId(v string)`

SetSnapshotExternalId sets SnapshotExternalId field to given value.

### HasSnapshotExternalId

`func (o *BackupResult) HasSnapshotExternalId() bool`

HasSnapshotExternalId returns a boolean if a field has been set.

### SetSnapshotExternalIdNil

`func (o *BackupResult) SetSnapshotExternalIdNil(b bool)`

 SetSnapshotExternalIdNil sets the value for SnapshotExternalId to be an explicit nil

### UnsetSnapshotExternalId
`func (o *BackupResult) UnsetSnapshotExternalId()`

UnsetSnapshotExternalId ensures that no value is present for SnapshotExternalId, not even an explicit nil
### GetCreatedBy

`func (o *BackupResult) GetCreatedBy() ListBackupResults200ResponseAllOfResultsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BackupResult) GetCreatedByOk() (*ListBackupResults200ResponseAllOfResultsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BackupResult) SetCreatedBy(v ListBackupResults200ResponseAllOfResultsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BackupResult) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *BackupResult) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BackupResult) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BackupResult) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *BackupResult) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *BackupResult) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *BackupResult) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *BackupResult) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *BackupResult) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



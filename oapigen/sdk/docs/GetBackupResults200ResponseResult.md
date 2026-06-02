# GetBackupResults200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup Result ID | [optional] 
**Backup** | Pointer to [**GetBackupResults200ResponseResultBackup**](GetBackupResults200ResponseResultBackup.md) |  | [optional] 
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
**CreatedBy** | Pointer to [**GetBackupResults200ResponseResultCreatedBy**](GetBackupResults200ResponseResultCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Methods

### NewGetBackupResults200ResponseResult

`func NewGetBackupResults200ResponseResult() *GetBackupResults200ResponseResult`

NewGetBackupResults200ResponseResult instantiates a new GetBackupResults200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetBackupResults200ResponseResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBackupResults200ResponseResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBackupResults200ResponseResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetBackupResults200ResponseResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBackup

`func (o *GetBackupResults200ResponseResult) GetBackup() GetBackupResults200ResponseResultBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *GetBackupResults200ResponseResult) GetBackupOk() (*GetBackupResults200ResponseResultBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *GetBackupResults200ResponseResult) SetBackup(v GetBackupResults200ResponseResultBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *GetBackupResults200ResponseResult) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetBackupSetId

`func (o *GetBackupResults200ResponseResult) GetBackupSetId() string`

GetBackupSetId returns the BackupSetId field if non-nil, zero value otherwise.

### GetBackupSetIdOk

`func (o *GetBackupResults200ResponseResult) GetBackupSetIdOk() (*string, bool)`

GetBackupSetIdOk returns a tuple with the BackupSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSetId

`func (o *GetBackupResults200ResponseResult) SetBackupSetId(v string)`

SetBackupSetId sets BackupSetId field to given value.

### HasBackupSetId

`func (o *GetBackupResults200ResponseResult) HasBackupSetId() bool`

HasBackupSetId returns a boolean if a field has been set.

### SetBackupSetIdNil

`func (o *GetBackupResults200ResponseResult) SetBackupSetIdNil(b bool)`

 SetBackupSetIdNil sets the value for BackupSetId to be an explicit nil

### UnsetBackupSetId
`func (o *GetBackupResults200ResponseResult) UnsetBackupSetId()`

UnsetBackupSetId ensures that no value is present for BackupSetId, not even an explicit nil
### GetInstanceId

`func (o *GetBackupResults200ResponseResult) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetBackupResults200ResponseResult) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetBackupResults200ResponseResult) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GetBackupResults200ResponseResult) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *GetBackupResults200ResponseResult) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *GetBackupResults200ResponseResult) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetContainerId

`func (o *GetBackupResults200ResponseResult) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GetBackupResults200ResponseResult) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GetBackupResults200ResponseResult) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *GetBackupResults200ResponseResult) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *GetBackupResults200ResponseResult) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *GetBackupResults200ResponseResult) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetServerId

`func (o *GetBackupResults200ResponseResult) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetBackupResults200ResponseResult) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetBackupResults200ResponseResult) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetBackupResults200ResponseResult) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *GetBackupResults200ResponseResult) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *GetBackupResults200ResponseResult) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetStatus

`func (o *GetBackupResults200ResponseResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBackupResults200ResponseResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBackupResults200ResponseResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetBackupResults200ResponseResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetBackupResults200ResponseResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetBackupResults200ResponseResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetErrorMessage

`func (o *GetBackupResults200ResponseResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetBackupResults200ResponseResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetBackupResults200ResponseResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetBackupResults200ResponseResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetBackupResults200ResponseResult) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetBackupResults200ResponseResult) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStartDate

`func (o *GetBackupResults200ResponseResult) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetBackupResults200ResponseResult) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetBackupResults200ResponseResult) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetBackupResults200ResponseResult) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *GetBackupResults200ResponseResult) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *GetBackupResults200ResponseResult) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *GetBackupResults200ResponseResult) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetBackupResults200ResponseResult) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetBackupResults200ResponseResult) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetBackupResults200ResponseResult) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *GetBackupResults200ResponseResult) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *GetBackupResults200ResponseResult) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDurationMillis

`func (o *GetBackupResults200ResponseResult) GetDurationMillis() int64`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *GetBackupResults200ResponseResult) GetDurationMillisOk() (*int64, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *GetBackupResults200ResponseResult) SetDurationMillis(v int64)`

SetDurationMillis sets DurationMillis field to given value.

### HasDurationMillis

`func (o *GetBackupResults200ResponseResult) HasDurationMillis() bool`

HasDurationMillis returns a boolean if a field has been set.

### SetDurationMillisNil

`func (o *GetBackupResults200ResponseResult) SetDurationMillisNil(b bool)`

 SetDurationMillisNil sets the value for DurationMillis to be an explicit nil

### UnsetDurationMillis
`func (o *GetBackupResults200ResponseResult) UnsetDurationMillis()`

UnsetDurationMillis ensures that no value is present for DurationMillis, not even an explicit nil
### GetSizeInBytes

`func (o *GetBackupResults200ResponseResult) GetSizeInBytes() int64`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *GetBackupResults200ResponseResult) GetSizeInBytesOk() (*int64, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *GetBackupResults200ResponseResult) SetSizeInBytes(v int64)`

SetSizeInBytes sets SizeInBytes field to given value.

### HasSizeInBytes

`func (o *GetBackupResults200ResponseResult) HasSizeInBytes() bool`

HasSizeInBytes returns a boolean if a field has been set.

### SetSizeInBytesNil

`func (o *GetBackupResults200ResponseResult) SetSizeInBytesNil(b bool)`

 SetSizeInBytesNil sets the value for SizeInBytes to be an explicit nil

### UnsetSizeInBytes
`func (o *GetBackupResults200ResponseResult) UnsetSizeInBytes()`

UnsetSizeInBytes ensures that no value is present for SizeInBytes, not even an explicit nil
### GetSizeInMb

`func (o *GetBackupResults200ResponseResult) GetSizeInMb() int64`

GetSizeInMb returns the SizeInMb field if non-nil, zero value otherwise.

### GetSizeInMbOk

`func (o *GetBackupResults200ResponseResult) GetSizeInMbOk() (*int64, bool)`

GetSizeInMbOk returns a tuple with the SizeInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInMb

`func (o *GetBackupResults200ResponseResult) SetSizeInMb(v int64)`

SetSizeInMb sets SizeInMb field to given value.

### HasSizeInMb

`func (o *GetBackupResults200ResponseResult) HasSizeInMb() bool`

HasSizeInMb returns a boolean if a field has been set.

### SetSizeInMbNil

`func (o *GetBackupResults200ResponseResult) SetSizeInMbNil(b bool)`

 SetSizeInMbNil sets the value for SizeInMb to be an explicit nil

### UnsetSizeInMb
`func (o *GetBackupResults200ResponseResult) UnsetSizeInMb()`

UnsetSizeInMb ensures that no value is present for SizeInMb, not even an explicit nil
### GetVolumePath

`func (o *GetBackupResults200ResponseResult) GetVolumePath() string`

GetVolumePath returns the VolumePath field if non-nil, zero value otherwise.

### GetVolumePathOk

`func (o *GetBackupResults200ResponseResult) GetVolumePathOk() (*string, bool)`

GetVolumePathOk returns a tuple with the VolumePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePath

`func (o *GetBackupResults200ResponseResult) SetVolumePath(v string)`

SetVolumePath sets VolumePath field to given value.

### HasVolumePath

`func (o *GetBackupResults200ResponseResult) HasVolumePath() bool`

HasVolumePath returns a boolean if a field has been set.

### SetVolumePathNil

`func (o *GetBackupResults200ResponseResult) SetVolumePathNil(b bool)`

 SetVolumePathNil sets the value for VolumePath to be an explicit nil

### UnsetVolumePath
`func (o *GetBackupResults200ResponseResult) UnsetVolumePath()`

UnsetVolumePath ensures that no value is present for VolumePath, not even an explicit nil
### GetResultArchive

`func (o *GetBackupResults200ResponseResult) GetResultArchive() string`

GetResultArchive returns the ResultArchive field if non-nil, zero value otherwise.

### GetResultArchiveOk

`func (o *GetBackupResults200ResponseResult) GetResultArchiveOk() (*string, bool)`

GetResultArchiveOk returns a tuple with the ResultArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultArchive

`func (o *GetBackupResults200ResponseResult) SetResultArchive(v string)`

SetResultArchive sets ResultArchive field to given value.

### HasResultArchive

`func (o *GetBackupResults200ResponseResult) HasResultArchive() bool`

HasResultArchive returns a boolean if a field has been set.

### SetResultArchiveNil

`func (o *GetBackupResults200ResponseResult) SetResultArchiveNil(b bool)`

 SetResultArchiveNil sets the value for ResultArchive to be an explicit nil

### UnsetResultArchive
`func (o *GetBackupResults200ResponseResult) UnsetResultArchive()`

UnsetResultArchive ensures that no value is present for ResultArchive, not even an explicit nil
### GetResultPath

`func (o *GetBackupResults200ResponseResult) GetResultPath() string`

GetResultPath returns the ResultPath field if non-nil, zero value otherwise.

### GetResultPathOk

`func (o *GetBackupResults200ResponseResult) GetResultPathOk() (*string, bool)`

GetResultPathOk returns a tuple with the ResultPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultPath

`func (o *GetBackupResults200ResponseResult) SetResultPath(v string)`

SetResultPath sets ResultPath field to given value.

### HasResultPath

`func (o *GetBackupResults200ResponseResult) HasResultPath() bool`

HasResultPath returns a boolean if a field has been set.

### SetResultPathNil

`func (o *GetBackupResults200ResponseResult) SetResultPathNil(b bool)`

 SetResultPathNil sets the value for ResultPath to be an explicit nil

### UnsetResultPath
`func (o *GetBackupResults200ResponseResult) UnsetResultPath()`

UnsetResultPath ensures that no value is present for ResultPath, not even an explicit nil
### GetExternalId

`func (o *GetBackupResults200ResponseResult) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetBackupResults200ResponseResult) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetBackupResults200ResponseResult) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetBackupResults200ResponseResult) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetBackupResults200ResponseResult) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetBackupResults200ResponseResult) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetSnapshotId

`func (o *GetBackupResults200ResponseResult) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *GetBackupResults200ResponseResult) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *GetBackupResults200ResponseResult) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *GetBackupResults200ResponseResult) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *GetBackupResults200ResponseResult) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *GetBackupResults200ResponseResult) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetSnapshotExternalId

`func (o *GetBackupResults200ResponseResult) GetSnapshotExternalId() string`

GetSnapshotExternalId returns the SnapshotExternalId field if non-nil, zero value otherwise.

### GetSnapshotExternalIdOk

`func (o *GetBackupResults200ResponseResult) GetSnapshotExternalIdOk() (*string, bool)`

GetSnapshotExternalIdOk returns a tuple with the SnapshotExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotExternalId

`func (o *GetBackupResults200ResponseResult) SetSnapshotExternalId(v string)`

SetSnapshotExternalId sets SnapshotExternalId field to given value.

### HasSnapshotExternalId

`func (o *GetBackupResults200ResponseResult) HasSnapshotExternalId() bool`

HasSnapshotExternalId returns a boolean if a field has been set.

### SetSnapshotExternalIdNil

`func (o *GetBackupResults200ResponseResult) SetSnapshotExternalIdNil(b bool)`

 SetSnapshotExternalIdNil sets the value for SnapshotExternalId to be an explicit nil

### UnsetSnapshotExternalId
`func (o *GetBackupResults200ResponseResult) UnsetSnapshotExternalId()`

UnsetSnapshotExternalId ensures that no value is present for SnapshotExternalId, not even an explicit nil
### GetCreatedBy

`func (o *GetBackupResults200ResponseResult) GetCreatedBy() GetBackupResults200ResponseResultCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetBackupResults200ResponseResult) GetCreatedByOk() (*GetBackupResults200ResponseResultCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetBackupResults200ResponseResult) SetCreatedBy(v GetBackupResults200ResponseResultCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetBackupResults200ResponseResult) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetBackupResults200ResponseResult) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetBackupResults200ResponseResult) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetBackupResults200ResponseResult) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetBackupResults200ResponseResult) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetBackupResults200ResponseResult) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetBackupResults200ResponseResult) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetBackupResults200ResponseResult) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetBackupResults200ResponseResult) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



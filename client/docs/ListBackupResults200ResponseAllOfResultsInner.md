# ListBackupResults200ResponseAllOfResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup Result ID | [optional] 
**Backup** | Pointer to [**ListBackupJobs200ResponseAllOfJobsInnerBackupsInner**](ListBackupJobs200ResponseAllOfJobsInnerBackupsInner.md) |  | [optional] 
**BackupSetId** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **int64** |  | [optional] 
**ContainerId** | Pointer to **int64** |  | [optional] 
**ServerId** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**DurationMillis** | Pointer to **int64** |  | [optional] 
**SizeInBytes** | Pointer to **int64** |  | [optional] 
**SizeInMb** | Pointer to **int64** |  | [optional] 
**VolumePath** | Pointer to **string** |  | [optional] 
**ResultArchive** | Pointer to **string** |  | [optional] 
**ResultPath** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**SnapshotId** | Pointer to **string** |  | [optional] 
**SnapshotExternalId** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**ListBackupResults200ResponseAllOfResultsInnerCreatedBy**](ListBackupResults200ResponseAllOfResultsInnerCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Methods

### NewListBackupResults200ResponseAllOfResultsInner

`func NewListBackupResults200ResponseAllOfResultsInner() *ListBackupResults200ResponseAllOfResultsInner`

NewListBackupResults200ResponseAllOfResultsInner instantiates a new ListBackupResults200ResponseAllOfResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupResults200ResponseAllOfResultsInnerWithDefaults

`func NewListBackupResults200ResponseAllOfResultsInnerWithDefaults() *ListBackupResults200ResponseAllOfResultsInner`

NewListBackupResults200ResponseAllOfResultsInnerWithDefaults instantiates a new ListBackupResults200ResponseAllOfResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBackup

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetBackup() ListBackupJobs200ResponseAllOfJobsInnerBackupsInner`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetBackupOk() (*ListBackupJobs200ResponseAllOfJobsInnerBackupsInner, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetBackup(v ListBackupJobs200ResponseAllOfJobsInnerBackupsInner)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetBackupSetId

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetBackupSetId() string`

GetBackupSetId returns the BackupSetId field if non-nil, zero value otherwise.

### GetBackupSetIdOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetBackupSetIdOk() (*string, bool)`

GetBackupSetIdOk returns a tuple with the BackupSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSetId

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetBackupSetId(v string)`

SetBackupSetId sets BackupSetId field to given value.

### HasBackupSetId

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasBackupSetId() bool`

HasBackupSetId returns a boolean if a field has been set.

### GetInstanceId

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetContainerId

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetServerId

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetStatus

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStartDate

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDurationMillis

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetDurationMillis() int64`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetDurationMillisOk() (*int64, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetDurationMillis(v int64)`

SetDurationMillis sets DurationMillis field to given value.

### HasDurationMillis

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasDurationMillis() bool`

HasDurationMillis returns a boolean if a field has been set.

### GetSizeInBytes

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetSizeInBytes() int64`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetSizeInBytesOk() (*int64, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetSizeInBytes(v int64)`

SetSizeInBytes sets SizeInBytes field to given value.

### HasSizeInBytes

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasSizeInBytes() bool`

HasSizeInBytes returns a boolean if a field has been set.

### GetSizeInMb

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetSizeInMb() int64`

GetSizeInMb returns the SizeInMb field if non-nil, zero value otherwise.

### GetSizeInMbOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetSizeInMbOk() (*int64, bool)`

GetSizeInMbOk returns a tuple with the SizeInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInMb

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetSizeInMb(v int64)`

SetSizeInMb sets SizeInMb field to given value.

### HasSizeInMb

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasSizeInMb() bool`

HasSizeInMb returns a boolean if a field has been set.

### GetVolumePath

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetVolumePath() string`

GetVolumePath returns the VolumePath field if non-nil, zero value otherwise.

### GetVolumePathOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetVolumePathOk() (*string, bool)`

GetVolumePathOk returns a tuple with the VolumePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePath

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetVolumePath(v string)`

SetVolumePath sets VolumePath field to given value.

### HasVolumePath

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasVolumePath() bool`

HasVolumePath returns a boolean if a field has been set.

### GetResultArchive

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetResultArchive() string`

GetResultArchive returns the ResultArchive field if non-nil, zero value otherwise.

### GetResultArchiveOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetResultArchiveOk() (*string, bool)`

GetResultArchiveOk returns a tuple with the ResultArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultArchive

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetResultArchive(v string)`

SetResultArchive sets ResultArchive field to given value.

### HasResultArchive

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasResultArchive() bool`

HasResultArchive returns a boolean if a field has been set.

### GetResultPath

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetResultPath() string`

GetResultPath returns the ResultPath field if non-nil, zero value otherwise.

### GetResultPathOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetResultPathOk() (*string, bool)`

GetResultPathOk returns a tuple with the ResultPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultPath

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetResultPath(v string)`

SetResultPath sets ResultPath field to given value.

### HasResultPath

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasResultPath() bool`

HasResultPath returns a boolean if a field has been set.

### GetExternalId

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetSnapshotId

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetSnapshotExternalId

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetSnapshotExternalId() string`

GetSnapshotExternalId returns the SnapshotExternalId field if non-nil, zero value otherwise.

### GetSnapshotExternalIdOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetSnapshotExternalIdOk() (*string, bool)`

GetSnapshotExternalIdOk returns a tuple with the SnapshotExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotExternalId

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetSnapshotExternalId(v string)`

SetSnapshotExternalId sets SnapshotExternalId field to given value.

### HasSnapshotExternalId

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasSnapshotExternalId() bool`

HasSnapshotExternalId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetCreatedBy() ListBackupResults200ResponseAllOfResultsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetCreatedByOk() (*ListBackupResults200ResponseAllOfResultsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetCreatedBy(v ListBackupResults200ResponseAllOfResultsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListBackupResults200ResponseAllOfResultsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListBackupResults200ResponseAllOfResultsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListBackupResults200ResponseAllOfResultsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



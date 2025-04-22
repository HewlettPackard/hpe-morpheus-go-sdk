# BackupRestore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup Result ID | [optional] 
**BackupResultId** | Pointer to **int64** |  | [optional] 
**BackupId** | Pointer to **int64** |  | [optional] 
**Backup** | Pointer to [**ListBackupRestores200ResponseAllOfRestoresInnerBackup**](ListBackupRestores200ResponseAllOfRestoresInnerBackup.md) |  | [optional] 
**ContainerId** | Pointer to **int64** |  | [optional] 
**Container** | Pointer to [**ListBackupRestores200ResponseAllOfRestoresInnerContainer**](ListBackupRestores200ResponseAllOfRestoresInnerContainer.md) |  | [optional] 
**Instance** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerInstance**](ListBackups200ResponseAllOfBackupsInnerInstance.md) |  | [optional] 
**RestoreToNew** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**DurationMillis** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalStatusRef** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Methods

### NewBackupRestore

`func NewBackupRestore() *BackupRestore`

NewBackupRestore instantiates a new BackupRestore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreWithDefaults

`func NewBackupRestoreWithDefaults() *BackupRestore`

NewBackupRestoreWithDefaults instantiates a new BackupRestore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupRestore) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupRestore) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupRestore) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BackupRestore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBackupResultId

`func (o *BackupRestore) GetBackupResultId() int64`

GetBackupResultId returns the BackupResultId field if non-nil, zero value otherwise.

### GetBackupResultIdOk

`func (o *BackupRestore) GetBackupResultIdOk() (*int64, bool)`

GetBackupResultIdOk returns a tuple with the BackupResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupResultId

`func (o *BackupRestore) SetBackupResultId(v int64)`

SetBackupResultId sets BackupResultId field to given value.

### HasBackupResultId

`func (o *BackupRestore) HasBackupResultId() bool`

HasBackupResultId returns a boolean if a field has been set.

### GetBackupId

`func (o *BackupRestore) GetBackupId() int64`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *BackupRestore) GetBackupIdOk() (*int64, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *BackupRestore) SetBackupId(v int64)`

SetBackupId sets BackupId field to given value.

### HasBackupId

`func (o *BackupRestore) HasBackupId() bool`

HasBackupId returns a boolean if a field has been set.

### GetBackup

`func (o *BackupRestore) GetBackup() ListBackupRestores200ResponseAllOfRestoresInnerBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *BackupRestore) GetBackupOk() (*ListBackupRestores200ResponseAllOfRestoresInnerBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *BackupRestore) SetBackup(v ListBackupRestores200ResponseAllOfRestoresInnerBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *BackupRestore) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetContainerId

`func (o *BackupRestore) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *BackupRestore) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *BackupRestore) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *BackupRestore) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetContainer

`func (o *BackupRestore) GetContainer() ListBackupRestores200ResponseAllOfRestoresInnerContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *BackupRestore) GetContainerOk() (*ListBackupRestores200ResponseAllOfRestoresInnerContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *BackupRestore) SetContainer(v ListBackupRestores200ResponseAllOfRestoresInnerContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *BackupRestore) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetInstance

`func (o *BackupRestore) GetInstance() ListBackups200ResponseAllOfBackupsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *BackupRestore) GetInstanceOk() (*ListBackups200ResponseAllOfBackupsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *BackupRestore) SetInstance(v ListBackups200ResponseAllOfBackupsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *BackupRestore) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetRestoreToNew

`func (o *BackupRestore) GetRestoreToNew() bool`

GetRestoreToNew returns the RestoreToNew field if non-nil, zero value otherwise.

### GetRestoreToNewOk

`func (o *BackupRestore) GetRestoreToNewOk() (*bool, bool)`

GetRestoreToNewOk returns a tuple with the RestoreToNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToNew

`func (o *BackupRestore) SetRestoreToNew(v bool)`

SetRestoreToNew sets RestoreToNew field to given value.

### HasRestoreToNew

`func (o *BackupRestore) HasRestoreToNew() bool`

HasRestoreToNew returns a boolean if a field has been set.

### GetStatus

`func (o *BackupRestore) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupRestore) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupRestore) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupRestore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *BackupRestore) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BackupRestore) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BackupRestore) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BackupRestore) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStartDate

`func (o *BackupRestore) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BackupRestore) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BackupRestore) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BackupRestore) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BackupRestore) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BackupRestore) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BackupRestore) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BackupRestore) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDurationMillis

`func (o *BackupRestore) GetDurationMillis() int64`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *BackupRestore) GetDurationMillisOk() (*int64, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *BackupRestore) SetDurationMillis(v int64)`

SetDurationMillis sets DurationMillis field to given value.

### HasDurationMillis

`func (o *BackupRestore) HasDurationMillis() bool`

HasDurationMillis returns a boolean if a field has been set.

### GetExternalId

`func (o *BackupRestore) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BackupRestore) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BackupRestore) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BackupRestore) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalStatusRef

`func (o *BackupRestore) GetExternalStatusRef() string`

GetExternalStatusRef returns the ExternalStatusRef field if non-nil, zero value otherwise.

### GetExternalStatusRefOk

`func (o *BackupRestore) GetExternalStatusRefOk() (*string, bool)`

GetExternalStatusRefOk returns a tuple with the ExternalStatusRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStatusRef

`func (o *BackupRestore) SetExternalStatusRef(v string)`

SetExternalStatusRef sets ExternalStatusRef field to given value.

### HasExternalStatusRef

`func (o *BackupRestore) HasExternalStatusRef() bool`

HasExternalStatusRef returns a boolean if a field has been set.

### GetDateCreated

`func (o *BackupRestore) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BackupRestore) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BackupRestore) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *BackupRestore) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *BackupRestore) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *BackupRestore) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *BackupRestore) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *BackupRestore) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



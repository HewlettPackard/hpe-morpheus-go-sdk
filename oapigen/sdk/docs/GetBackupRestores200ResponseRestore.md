# GetBackupRestores200ResponseRestore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup Result ID | [optional] 
**BackupResultId** | Pointer to **int64** |  | [optional] 
**BackupId** | Pointer to **int64** |  | [optional] 
**Backup** | Pointer to [**GetBackupRestores200ResponseRestoreBackup**](GetBackupRestores200ResponseRestoreBackup.md) |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**Container** | Pointer to [**GetBackupRestores200ResponseRestoreContainer**](GetBackupRestores200ResponseRestoreContainer.md) |  | [optional] 
**Instance** | Pointer to [**GetBackupRestores200ResponseRestoreInstance**](GetBackupRestores200ResponseRestoreInstance.md) |  | [optional] 
**RestoreToNew** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**DurationMillis** | Pointer to **NullableInt64** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalStatusRef** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Methods

### NewGetBackupRestores200ResponseRestore

`func NewGetBackupRestores200ResponseRestore() *GetBackupRestores200ResponseRestore`

NewGetBackupRestores200ResponseRestore instantiates a new GetBackupRestores200ResponseRestore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetBackupRestores200ResponseRestore) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBackupRestores200ResponseRestore) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBackupRestores200ResponseRestore) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetBackupRestores200ResponseRestore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBackupResultId

`func (o *GetBackupRestores200ResponseRestore) GetBackupResultId() int64`

GetBackupResultId returns the BackupResultId field if non-nil, zero value otherwise.

### GetBackupResultIdOk

`func (o *GetBackupRestores200ResponseRestore) GetBackupResultIdOk() (*int64, bool)`

GetBackupResultIdOk returns a tuple with the BackupResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupResultId

`func (o *GetBackupRestores200ResponseRestore) SetBackupResultId(v int64)`

SetBackupResultId sets BackupResultId field to given value.

### HasBackupResultId

`func (o *GetBackupRestores200ResponseRestore) HasBackupResultId() bool`

HasBackupResultId returns a boolean if a field has been set.

### GetBackupId

`func (o *GetBackupRestores200ResponseRestore) GetBackupId() int64`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *GetBackupRestores200ResponseRestore) GetBackupIdOk() (*int64, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *GetBackupRestores200ResponseRestore) SetBackupId(v int64)`

SetBackupId sets BackupId field to given value.

### HasBackupId

`func (o *GetBackupRestores200ResponseRestore) HasBackupId() bool`

HasBackupId returns a boolean if a field has been set.

### GetBackup

`func (o *GetBackupRestores200ResponseRestore) GetBackup() GetBackupRestores200ResponseRestoreBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *GetBackupRestores200ResponseRestore) GetBackupOk() (*GetBackupRestores200ResponseRestoreBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *GetBackupRestores200ResponseRestore) SetBackup(v GetBackupRestores200ResponseRestoreBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *GetBackupRestores200ResponseRestore) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetContainerId

`func (o *GetBackupRestores200ResponseRestore) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GetBackupRestores200ResponseRestore) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GetBackupRestores200ResponseRestore) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *GetBackupRestores200ResponseRestore) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *GetBackupRestores200ResponseRestore) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *GetBackupRestores200ResponseRestore) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetContainer

`func (o *GetBackupRestores200ResponseRestore) GetContainer() GetBackupRestores200ResponseRestoreContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *GetBackupRestores200ResponseRestore) GetContainerOk() (*GetBackupRestores200ResponseRestoreContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *GetBackupRestores200ResponseRestore) SetContainer(v GetBackupRestores200ResponseRestoreContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *GetBackupRestores200ResponseRestore) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetInstance

`func (o *GetBackupRestores200ResponseRestore) GetInstance() GetBackupRestores200ResponseRestoreInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetBackupRestores200ResponseRestore) GetInstanceOk() (*GetBackupRestores200ResponseRestoreInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetBackupRestores200ResponseRestore) SetInstance(v GetBackupRestores200ResponseRestoreInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetBackupRestores200ResponseRestore) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetRestoreToNew

`func (o *GetBackupRestores200ResponseRestore) GetRestoreToNew() bool`

GetRestoreToNew returns the RestoreToNew field if non-nil, zero value otherwise.

### GetRestoreToNewOk

`func (o *GetBackupRestores200ResponseRestore) GetRestoreToNewOk() (*bool, bool)`

GetRestoreToNewOk returns a tuple with the RestoreToNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToNew

`func (o *GetBackupRestores200ResponseRestore) SetRestoreToNew(v bool)`

SetRestoreToNew sets RestoreToNew field to given value.

### HasRestoreToNew

`func (o *GetBackupRestores200ResponseRestore) HasRestoreToNew() bool`

HasRestoreToNew returns a boolean if a field has been set.

### GetStatus

`func (o *GetBackupRestores200ResponseRestore) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBackupRestores200ResponseRestore) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBackupRestores200ResponseRestore) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetBackupRestores200ResponseRestore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetBackupRestores200ResponseRestore) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetBackupRestores200ResponseRestore) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetErrorMessage

`func (o *GetBackupRestores200ResponseRestore) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetBackupRestores200ResponseRestore) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetBackupRestores200ResponseRestore) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetBackupRestores200ResponseRestore) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetBackupRestores200ResponseRestore) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetBackupRestores200ResponseRestore) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStartDate

`func (o *GetBackupRestores200ResponseRestore) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetBackupRestores200ResponseRestore) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetBackupRestores200ResponseRestore) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetBackupRestores200ResponseRestore) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *GetBackupRestores200ResponseRestore) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *GetBackupRestores200ResponseRestore) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *GetBackupRestores200ResponseRestore) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetBackupRestores200ResponseRestore) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetBackupRestores200ResponseRestore) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetBackupRestores200ResponseRestore) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *GetBackupRestores200ResponseRestore) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *GetBackupRestores200ResponseRestore) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDurationMillis

`func (o *GetBackupRestores200ResponseRestore) GetDurationMillis() int64`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *GetBackupRestores200ResponseRestore) GetDurationMillisOk() (*int64, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *GetBackupRestores200ResponseRestore) SetDurationMillis(v int64)`

SetDurationMillis sets DurationMillis field to given value.

### HasDurationMillis

`func (o *GetBackupRestores200ResponseRestore) HasDurationMillis() bool`

HasDurationMillis returns a boolean if a field has been set.

### SetDurationMillisNil

`func (o *GetBackupRestores200ResponseRestore) SetDurationMillisNil(b bool)`

 SetDurationMillisNil sets the value for DurationMillis to be an explicit nil

### UnsetDurationMillis
`func (o *GetBackupRestores200ResponseRestore) UnsetDurationMillis()`

UnsetDurationMillis ensures that no value is present for DurationMillis, not even an explicit nil
### GetExternalId

`func (o *GetBackupRestores200ResponseRestore) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetBackupRestores200ResponseRestore) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetBackupRestores200ResponseRestore) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetBackupRestores200ResponseRestore) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetBackupRestores200ResponseRestore) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetBackupRestores200ResponseRestore) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalStatusRef

`func (o *GetBackupRestores200ResponseRestore) GetExternalStatusRef() string`

GetExternalStatusRef returns the ExternalStatusRef field if non-nil, zero value otherwise.

### GetExternalStatusRefOk

`func (o *GetBackupRestores200ResponseRestore) GetExternalStatusRefOk() (*string, bool)`

GetExternalStatusRefOk returns a tuple with the ExternalStatusRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStatusRef

`func (o *GetBackupRestores200ResponseRestore) SetExternalStatusRef(v string)`

SetExternalStatusRef sets ExternalStatusRef field to given value.

### HasExternalStatusRef

`func (o *GetBackupRestores200ResponseRestore) HasExternalStatusRef() bool`

HasExternalStatusRef returns a boolean if a field has been set.

### SetExternalStatusRefNil

`func (o *GetBackupRestores200ResponseRestore) SetExternalStatusRefNil(b bool)`

 SetExternalStatusRefNil sets the value for ExternalStatusRef to be an explicit nil

### UnsetExternalStatusRef
`func (o *GetBackupRestores200ResponseRestore) UnsetExternalStatusRef()`

UnsetExternalStatusRef ensures that no value is present for ExternalStatusRef, not even an explicit nil
### GetDateCreated

`func (o *GetBackupRestores200ResponseRestore) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetBackupRestores200ResponseRestore) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetBackupRestores200ResponseRestore) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetBackupRestores200ResponseRestore) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetBackupRestores200ResponseRestore) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetBackupRestores200ResponseRestore) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetBackupRestores200ResponseRestore) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetBackupRestores200ResponseRestore) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



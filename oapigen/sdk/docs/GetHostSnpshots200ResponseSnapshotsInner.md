# GetHostSnpshots200ResponseSnapshotsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**SnapshotType** | Pointer to **string** |  | [optional] 
**SnapshotCreated** | Pointer to **NullableTime** |  | [optional] 
**Zone** | Pointer to [**GetHostSnpshots200ResponseSnapshotsInnerZone**](GetHostSnpshots200ResponseSnapshotsInnerZone.md) |  | [optional] 
**Datastore** | Pointer to **NullableString** |  | [optional] 
**ParentSnapshot** | Pointer to **NullableString** |  | [optional] 
**SnapshotFiles** | Pointer to [**[]GetHostSnpshots200ResponseSnapshotsInnerSnapshotFilesInner**](GetHostSnpshots200ResponseSnapshotsInnerSnapshotFilesInner.md) |  | [optional] 
**CurrentlyActive** | Pointer to **bool** |  | [optional] 
**MemorySnapshot** | Pointer to **bool** |  | [optional] 
**ForExport** | Pointer to **bool** |  | [optional] 
**ForBackup** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetHostSnpshots200ResponseSnapshotsInner

`func NewGetHostSnpshots200ResponseSnapshotsInner() *GetHostSnpshots200ResponseSnapshotsInner`

NewGetHostSnpshots200ResponseSnapshotsInner instantiates a new GetHostSnpshots200ResponseSnapshotsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHostSnpshots200ResponseSnapshotsInnerWithDefaults

`func NewGetHostSnpshots200ResponseSnapshotsInnerWithDefaults() *GetHostSnpshots200ResponseSnapshotsInner`

NewGetHostSnpshots200ResponseSnapshotsInnerWithDefaults instantiates a new GetHostSnpshots200ResponseSnapshotsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetHostSnpshots200ResponseSnapshotsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetHostSnpshots200ResponseSnapshotsInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetStatus

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetState

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *GetHostSnpshots200ResponseSnapshotsInner) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetSnapshotType

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetSnapshotCreated

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetSnapshotCreated() time.Time`

GetSnapshotCreated returns the SnapshotCreated field if non-nil, zero value otherwise.

### GetSnapshotCreatedOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetSnapshotCreatedOk() (*time.Time, bool)`

GetSnapshotCreatedOk returns a tuple with the SnapshotCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCreated

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetSnapshotCreated(v time.Time)`

SetSnapshotCreated sets SnapshotCreated field to given value.

### HasSnapshotCreated

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasSnapshotCreated() bool`

HasSnapshotCreated returns a boolean if a field has been set.

### SetSnapshotCreatedNil

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetSnapshotCreatedNil(b bool)`

 SetSnapshotCreatedNil sets the value for SnapshotCreated to be an explicit nil

### UnsetSnapshotCreated
`func (o *GetHostSnpshots200ResponseSnapshotsInner) UnsetSnapshotCreated()`

UnsetSnapshotCreated ensures that no value is present for SnapshotCreated, not even an explicit nil
### GetZone

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetZone() GetHostSnpshots200ResponseSnapshotsInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetZoneOk() (*GetHostSnpshots200ResponseSnapshotsInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetZone(v GetHostSnpshots200ResponseSnapshotsInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetDatastore

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### SetDatastoreNil

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetDatastoreNil(b bool)`

 SetDatastoreNil sets the value for Datastore to be an explicit nil

### UnsetDatastore
`func (o *GetHostSnpshots200ResponseSnapshotsInner) UnsetDatastore()`

UnsetDatastore ensures that no value is present for Datastore, not even an explicit nil
### GetParentSnapshot

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetParentSnapshot() string`

GetParentSnapshot returns the ParentSnapshot field if non-nil, zero value otherwise.

### GetParentSnapshotOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetParentSnapshotOk() (*string, bool)`

GetParentSnapshotOk returns a tuple with the ParentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSnapshot

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetParentSnapshot(v string)`

SetParentSnapshot sets ParentSnapshot field to given value.

### HasParentSnapshot

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasParentSnapshot() bool`

HasParentSnapshot returns a boolean if a field has been set.

### SetParentSnapshotNil

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetParentSnapshotNil(b bool)`

 SetParentSnapshotNil sets the value for ParentSnapshot to be an explicit nil

### UnsetParentSnapshot
`func (o *GetHostSnpshots200ResponseSnapshotsInner) UnsetParentSnapshot()`

UnsetParentSnapshot ensures that no value is present for ParentSnapshot, not even an explicit nil
### GetSnapshotFiles

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetSnapshotFiles() []GetHostSnpshots200ResponseSnapshotsInnerSnapshotFilesInner`

GetSnapshotFiles returns the SnapshotFiles field if non-nil, zero value otherwise.

### GetSnapshotFilesOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetSnapshotFilesOk() (*[]GetHostSnpshots200ResponseSnapshotsInnerSnapshotFilesInner, bool)`

GetSnapshotFilesOk returns a tuple with the SnapshotFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotFiles

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetSnapshotFiles(v []GetHostSnpshots200ResponseSnapshotsInnerSnapshotFilesInner)`

SetSnapshotFiles sets SnapshotFiles field to given value.

### HasSnapshotFiles

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasSnapshotFiles() bool`

HasSnapshotFiles returns a boolean if a field has been set.

### GetCurrentlyActive

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetCurrentlyActive() bool`

GetCurrentlyActive returns the CurrentlyActive field if non-nil, zero value otherwise.

### GetCurrentlyActiveOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetCurrentlyActiveOk() (*bool, bool)`

GetCurrentlyActiveOk returns a tuple with the CurrentlyActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyActive

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetCurrentlyActive(v bool)`

SetCurrentlyActive sets CurrentlyActive field to given value.

### HasCurrentlyActive

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasCurrentlyActive() bool`

HasCurrentlyActive returns a boolean if a field has been set.

### GetMemorySnapshot

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetMemorySnapshot() bool`

GetMemorySnapshot returns the MemorySnapshot field if non-nil, zero value otherwise.

### GetMemorySnapshotOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetMemorySnapshotOk() (*bool, bool)`

GetMemorySnapshotOk returns a tuple with the MemorySnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySnapshot

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetMemorySnapshot(v bool)`

SetMemorySnapshot sets MemorySnapshot field to given value.

### HasMemorySnapshot

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasMemorySnapshot() bool`

HasMemorySnapshot returns a boolean if a field has been set.

### GetForExport

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetForExport() bool`

GetForExport returns the ForExport field if non-nil, zero value otherwise.

### GetForExportOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetForExportOk() (*bool, bool)`

GetForExportOk returns a tuple with the ForExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForExport

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetForExport(v bool)`

SetForExport sets ForExport field to given value.

### HasForExport

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasForExport() bool`

HasForExport returns a boolean if a field has been set.

### GetForBackup

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetForBackup() bool`

GetForBackup returns the ForBackup field if non-nil, zero value otherwise.

### GetForBackupOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetForBackupOk() (*bool, bool)`

GetForBackupOk returns a tuple with the ForBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForBackup

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetForBackup(v bool)`

SetForBackup sets ForBackup field to given value.

### HasForBackup

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasForBackup() bool`

HasForBackup returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetHostSnpshots200ResponseSnapshotsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetHostSnpshots200ResponseSnapshotsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetHostSnpshots200ResponseSnapshotsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



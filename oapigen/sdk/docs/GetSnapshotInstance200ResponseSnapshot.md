# GetSnapshotInstance200ResponseSnapshot

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

### NewGetSnapshotInstance200ResponseSnapshot

`func NewGetSnapshotInstance200ResponseSnapshot() *GetSnapshotInstance200ResponseSnapshot`

NewGetSnapshotInstance200ResponseSnapshot instantiates a new GetSnapshotInstance200ResponseSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetSnapshotInstance200ResponseSnapshot) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSnapshotInstance200ResponseSnapshot) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetSnapshotInstance200ResponseSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetSnapshotInstance200ResponseSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSnapshotInstance200ResponseSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSnapshotInstance200ResponseSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetSnapshotInstance200ResponseSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSnapshotInstance200ResponseSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetSnapshotInstance200ResponseSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetSnapshotInstance200ResponseSnapshot) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetSnapshotInstance200ResponseSnapshot) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *GetSnapshotInstance200ResponseSnapshot) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetSnapshotInstance200ResponseSnapshot) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetSnapshotInstance200ResponseSnapshot) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetSnapshotInstance200ResponseSnapshot) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetSnapshotInstance200ResponseSnapshot) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetStatus

`func (o *GetSnapshotInstance200ResponseSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSnapshotInstance200ResponseSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSnapshotInstance200ResponseSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetState

`func (o *GetSnapshotInstance200ResponseSnapshot) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetSnapshotInstance200ResponseSnapshot) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetSnapshotInstance200ResponseSnapshot) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *GetSnapshotInstance200ResponseSnapshot) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *GetSnapshotInstance200ResponseSnapshot) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetSnapshotType

`func (o *GetSnapshotInstance200ResponseSnapshot) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *GetSnapshotInstance200ResponseSnapshot) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *GetSnapshotInstance200ResponseSnapshot) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetSnapshotCreated

`func (o *GetSnapshotInstance200ResponseSnapshot) GetSnapshotCreated() time.Time`

GetSnapshotCreated returns the SnapshotCreated field if non-nil, zero value otherwise.

### GetSnapshotCreatedOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetSnapshotCreatedOk() (*time.Time, bool)`

GetSnapshotCreatedOk returns a tuple with the SnapshotCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCreated

`func (o *GetSnapshotInstance200ResponseSnapshot) SetSnapshotCreated(v time.Time)`

SetSnapshotCreated sets SnapshotCreated field to given value.

### HasSnapshotCreated

`func (o *GetSnapshotInstance200ResponseSnapshot) HasSnapshotCreated() bool`

HasSnapshotCreated returns a boolean if a field has been set.

### SetSnapshotCreatedNil

`func (o *GetSnapshotInstance200ResponseSnapshot) SetSnapshotCreatedNil(b bool)`

 SetSnapshotCreatedNil sets the value for SnapshotCreated to be an explicit nil

### UnsetSnapshotCreated
`func (o *GetSnapshotInstance200ResponseSnapshot) UnsetSnapshotCreated()`

UnsetSnapshotCreated ensures that no value is present for SnapshotCreated, not even an explicit nil
### GetZone

`func (o *GetSnapshotInstance200ResponseSnapshot) GetZone() GetHostSnpshots200ResponseSnapshotsInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetZoneOk() (*GetHostSnpshots200ResponseSnapshotsInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetSnapshotInstance200ResponseSnapshot) SetZone(v GetHostSnpshots200ResponseSnapshotsInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetSnapshotInstance200ResponseSnapshot) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetDatastore

`func (o *GetSnapshotInstance200ResponseSnapshot) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *GetSnapshotInstance200ResponseSnapshot) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *GetSnapshotInstance200ResponseSnapshot) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### SetDatastoreNil

`func (o *GetSnapshotInstance200ResponseSnapshot) SetDatastoreNil(b bool)`

 SetDatastoreNil sets the value for Datastore to be an explicit nil

### UnsetDatastore
`func (o *GetSnapshotInstance200ResponseSnapshot) UnsetDatastore()`

UnsetDatastore ensures that no value is present for Datastore, not even an explicit nil
### GetParentSnapshot

`func (o *GetSnapshotInstance200ResponseSnapshot) GetParentSnapshot() string`

GetParentSnapshot returns the ParentSnapshot field if non-nil, zero value otherwise.

### GetParentSnapshotOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetParentSnapshotOk() (*string, bool)`

GetParentSnapshotOk returns a tuple with the ParentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSnapshot

`func (o *GetSnapshotInstance200ResponseSnapshot) SetParentSnapshot(v string)`

SetParentSnapshot sets ParentSnapshot field to given value.

### HasParentSnapshot

`func (o *GetSnapshotInstance200ResponseSnapshot) HasParentSnapshot() bool`

HasParentSnapshot returns a boolean if a field has been set.

### SetParentSnapshotNil

`func (o *GetSnapshotInstance200ResponseSnapshot) SetParentSnapshotNil(b bool)`

 SetParentSnapshotNil sets the value for ParentSnapshot to be an explicit nil

### UnsetParentSnapshot
`func (o *GetSnapshotInstance200ResponseSnapshot) UnsetParentSnapshot()`

UnsetParentSnapshot ensures that no value is present for ParentSnapshot, not even an explicit nil
### GetSnapshotFiles

`func (o *GetSnapshotInstance200ResponseSnapshot) GetSnapshotFiles() []GetHostSnpshots200ResponseSnapshotsInnerSnapshotFilesInner`

GetSnapshotFiles returns the SnapshotFiles field if non-nil, zero value otherwise.

### GetSnapshotFilesOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetSnapshotFilesOk() (*[]GetHostSnpshots200ResponseSnapshotsInnerSnapshotFilesInner, bool)`

GetSnapshotFilesOk returns a tuple with the SnapshotFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotFiles

`func (o *GetSnapshotInstance200ResponseSnapshot) SetSnapshotFiles(v []GetHostSnpshots200ResponseSnapshotsInnerSnapshotFilesInner)`

SetSnapshotFiles sets SnapshotFiles field to given value.

### HasSnapshotFiles

`func (o *GetSnapshotInstance200ResponseSnapshot) HasSnapshotFiles() bool`

HasSnapshotFiles returns a boolean if a field has been set.

### GetCurrentlyActive

`func (o *GetSnapshotInstance200ResponseSnapshot) GetCurrentlyActive() bool`

GetCurrentlyActive returns the CurrentlyActive field if non-nil, zero value otherwise.

### GetCurrentlyActiveOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetCurrentlyActiveOk() (*bool, bool)`

GetCurrentlyActiveOk returns a tuple with the CurrentlyActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyActive

`func (o *GetSnapshotInstance200ResponseSnapshot) SetCurrentlyActive(v bool)`

SetCurrentlyActive sets CurrentlyActive field to given value.

### HasCurrentlyActive

`func (o *GetSnapshotInstance200ResponseSnapshot) HasCurrentlyActive() bool`

HasCurrentlyActive returns a boolean if a field has been set.

### GetMemorySnapshot

`func (o *GetSnapshotInstance200ResponseSnapshot) GetMemorySnapshot() bool`

GetMemorySnapshot returns the MemorySnapshot field if non-nil, zero value otherwise.

### GetMemorySnapshotOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetMemorySnapshotOk() (*bool, bool)`

GetMemorySnapshotOk returns a tuple with the MemorySnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySnapshot

`func (o *GetSnapshotInstance200ResponseSnapshot) SetMemorySnapshot(v bool)`

SetMemorySnapshot sets MemorySnapshot field to given value.

### HasMemorySnapshot

`func (o *GetSnapshotInstance200ResponseSnapshot) HasMemorySnapshot() bool`

HasMemorySnapshot returns a boolean if a field has been set.

### GetForExport

`func (o *GetSnapshotInstance200ResponseSnapshot) GetForExport() bool`

GetForExport returns the ForExport field if non-nil, zero value otherwise.

### GetForExportOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetForExportOk() (*bool, bool)`

GetForExportOk returns a tuple with the ForExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForExport

`func (o *GetSnapshotInstance200ResponseSnapshot) SetForExport(v bool)`

SetForExport sets ForExport field to given value.

### HasForExport

`func (o *GetSnapshotInstance200ResponseSnapshot) HasForExport() bool`

HasForExport returns a boolean if a field has been set.

### GetForBackup

`func (o *GetSnapshotInstance200ResponseSnapshot) GetForBackup() bool`

GetForBackup returns the ForBackup field if non-nil, zero value otherwise.

### GetForBackupOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetForBackupOk() (*bool, bool)`

GetForBackupOk returns a tuple with the ForBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForBackup

`func (o *GetSnapshotInstance200ResponseSnapshot) SetForBackup(v bool)`

SetForBackup sets ForBackup field to given value.

### HasForBackup

`func (o *GetSnapshotInstance200ResponseSnapshot) HasForBackup() bool`

HasForBackup returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetSnapshotInstance200ResponseSnapshot) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetSnapshotInstance200ResponseSnapshot) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetSnapshotInstance200ResponseSnapshot) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetSnapshotInstance200ResponseSnapshot) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



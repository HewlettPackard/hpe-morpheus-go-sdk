# InstanceSnapshotsSnapshotsInner

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
**Zone** | Pointer to [**InstanceSnapshotsSnapshotsInnerZone**](InstanceSnapshotsSnapshotsInnerZone.md) |  | [optional] 
**Datastore** | Pointer to **NullableString** |  | [optional] 
**ParentSnapshot** | Pointer to **NullableString** |  | [optional] 
**SnapshotFiles** | Pointer to [**[]InstanceSnapshotsSnapshotsInnerSnapshotFilesInner**](InstanceSnapshotsSnapshotsInnerSnapshotFilesInner.md) |  | [optional] 
**CurrentlyActive** | Pointer to **bool** |  | [optional] 
**MemorySnapshot** | Pointer to **bool** |  | [optional] 
**ForExport** | Pointer to **bool** |  | [optional] 
**ForBackup** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInstanceSnapshotsSnapshotsInner

`func NewInstanceSnapshotsSnapshotsInner() *InstanceSnapshotsSnapshotsInner`

NewInstanceSnapshotsSnapshotsInner instantiates a new InstanceSnapshotsSnapshotsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *InstanceSnapshotsSnapshotsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceSnapshotsSnapshotsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceSnapshotsSnapshotsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceSnapshotsSnapshotsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InstanceSnapshotsSnapshotsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceSnapshotsSnapshotsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceSnapshotsSnapshotsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceSnapshotsSnapshotsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceSnapshotsSnapshotsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceSnapshotsSnapshotsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceSnapshotsSnapshotsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceSnapshotsSnapshotsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InstanceSnapshotsSnapshotsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InstanceSnapshotsSnapshotsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *InstanceSnapshotsSnapshotsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InstanceSnapshotsSnapshotsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InstanceSnapshotsSnapshotsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InstanceSnapshotsSnapshotsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *InstanceSnapshotsSnapshotsInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *InstanceSnapshotsSnapshotsInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetStatus

`func (o *InstanceSnapshotsSnapshotsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceSnapshotsSnapshotsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceSnapshotsSnapshotsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstanceSnapshotsSnapshotsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetState

`func (o *InstanceSnapshotsSnapshotsInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InstanceSnapshotsSnapshotsInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InstanceSnapshotsSnapshotsInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InstanceSnapshotsSnapshotsInner) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *InstanceSnapshotsSnapshotsInner) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *InstanceSnapshotsSnapshotsInner) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetSnapshotType

`func (o *InstanceSnapshotsSnapshotsInner) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *InstanceSnapshotsSnapshotsInner) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *InstanceSnapshotsSnapshotsInner) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *InstanceSnapshotsSnapshotsInner) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetSnapshotCreated

`func (o *InstanceSnapshotsSnapshotsInner) GetSnapshotCreated() time.Time`

GetSnapshotCreated returns the SnapshotCreated field if non-nil, zero value otherwise.

### GetSnapshotCreatedOk

`func (o *InstanceSnapshotsSnapshotsInner) GetSnapshotCreatedOk() (*time.Time, bool)`

GetSnapshotCreatedOk returns a tuple with the SnapshotCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCreated

`func (o *InstanceSnapshotsSnapshotsInner) SetSnapshotCreated(v time.Time)`

SetSnapshotCreated sets SnapshotCreated field to given value.

### HasSnapshotCreated

`func (o *InstanceSnapshotsSnapshotsInner) HasSnapshotCreated() bool`

HasSnapshotCreated returns a boolean if a field has been set.

### SetSnapshotCreatedNil

`func (o *InstanceSnapshotsSnapshotsInner) SetSnapshotCreatedNil(b bool)`

 SetSnapshotCreatedNil sets the value for SnapshotCreated to be an explicit nil

### UnsetSnapshotCreated
`func (o *InstanceSnapshotsSnapshotsInner) UnsetSnapshotCreated()`

UnsetSnapshotCreated ensures that no value is present for SnapshotCreated, not even an explicit nil
### GetZone

`func (o *InstanceSnapshotsSnapshotsInner) GetZone() InstanceSnapshotsSnapshotsInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *InstanceSnapshotsSnapshotsInner) GetZoneOk() (*InstanceSnapshotsSnapshotsInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *InstanceSnapshotsSnapshotsInner) SetZone(v InstanceSnapshotsSnapshotsInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *InstanceSnapshotsSnapshotsInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetDatastore

`func (o *InstanceSnapshotsSnapshotsInner) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *InstanceSnapshotsSnapshotsInner) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *InstanceSnapshotsSnapshotsInner) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *InstanceSnapshotsSnapshotsInner) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### SetDatastoreNil

`func (o *InstanceSnapshotsSnapshotsInner) SetDatastoreNil(b bool)`

 SetDatastoreNil sets the value for Datastore to be an explicit nil

### UnsetDatastore
`func (o *InstanceSnapshotsSnapshotsInner) UnsetDatastore()`

UnsetDatastore ensures that no value is present for Datastore, not even an explicit nil
### GetParentSnapshot

`func (o *InstanceSnapshotsSnapshotsInner) GetParentSnapshot() string`

GetParentSnapshot returns the ParentSnapshot field if non-nil, zero value otherwise.

### GetParentSnapshotOk

`func (o *InstanceSnapshotsSnapshotsInner) GetParentSnapshotOk() (*string, bool)`

GetParentSnapshotOk returns a tuple with the ParentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSnapshot

`func (o *InstanceSnapshotsSnapshotsInner) SetParentSnapshot(v string)`

SetParentSnapshot sets ParentSnapshot field to given value.

### HasParentSnapshot

`func (o *InstanceSnapshotsSnapshotsInner) HasParentSnapshot() bool`

HasParentSnapshot returns a boolean if a field has been set.

### SetParentSnapshotNil

`func (o *InstanceSnapshotsSnapshotsInner) SetParentSnapshotNil(b bool)`

 SetParentSnapshotNil sets the value for ParentSnapshot to be an explicit nil

### UnsetParentSnapshot
`func (o *InstanceSnapshotsSnapshotsInner) UnsetParentSnapshot()`

UnsetParentSnapshot ensures that no value is present for ParentSnapshot, not even an explicit nil
### GetSnapshotFiles

`func (o *InstanceSnapshotsSnapshotsInner) GetSnapshotFiles() []InstanceSnapshotsSnapshotsInnerSnapshotFilesInner`

GetSnapshotFiles returns the SnapshotFiles field if non-nil, zero value otherwise.

### GetSnapshotFilesOk

`func (o *InstanceSnapshotsSnapshotsInner) GetSnapshotFilesOk() (*[]InstanceSnapshotsSnapshotsInnerSnapshotFilesInner, bool)`

GetSnapshotFilesOk returns a tuple with the SnapshotFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotFiles

`func (o *InstanceSnapshotsSnapshotsInner) SetSnapshotFiles(v []InstanceSnapshotsSnapshotsInnerSnapshotFilesInner)`

SetSnapshotFiles sets SnapshotFiles field to given value.

### HasSnapshotFiles

`func (o *InstanceSnapshotsSnapshotsInner) HasSnapshotFiles() bool`

HasSnapshotFiles returns a boolean if a field has been set.

### GetCurrentlyActive

`func (o *InstanceSnapshotsSnapshotsInner) GetCurrentlyActive() bool`

GetCurrentlyActive returns the CurrentlyActive field if non-nil, zero value otherwise.

### GetCurrentlyActiveOk

`func (o *InstanceSnapshotsSnapshotsInner) GetCurrentlyActiveOk() (*bool, bool)`

GetCurrentlyActiveOk returns a tuple with the CurrentlyActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyActive

`func (o *InstanceSnapshotsSnapshotsInner) SetCurrentlyActive(v bool)`

SetCurrentlyActive sets CurrentlyActive field to given value.

### HasCurrentlyActive

`func (o *InstanceSnapshotsSnapshotsInner) HasCurrentlyActive() bool`

HasCurrentlyActive returns a boolean if a field has been set.

### GetMemorySnapshot

`func (o *InstanceSnapshotsSnapshotsInner) GetMemorySnapshot() bool`

GetMemorySnapshot returns the MemorySnapshot field if non-nil, zero value otherwise.

### GetMemorySnapshotOk

`func (o *InstanceSnapshotsSnapshotsInner) GetMemorySnapshotOk() (*bool, bool)`

GetMemorySnapshotOk returns a tuple with the MemorySnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySnapshot

`func (o *InstanceSnapshotsSnapshotsInner) SetMemorySnapshot(v bool)`

SetMemorySnapshot sets MemorySnapshot field to given value.

### HasMemorySnapshot

`func (o *InstanceSnapshotsSnapshotsInner) HasMemorySnapshot() bool`

HasMemorySnapshot returns a boolean if a field has been set.

### GetForExport

`func (o *InstanceSnapshotsSnapshotsInner) GetForExport() bool`

GetForExport returns the ForExport field if non-nil, zero value otherwise.

### GetForExportOk

`func (o *InstanceSnapshotsSnapshotsInner) GetForExportOk() (*bool, bool)`

GetForExportOk returns a tuple with the ForExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForExport

`func (o *InstanceSnapshotsSnapshotsInner) SetForExport(v bool)`

SetForExport sets ForExport field to given value.

### HasForExport

`func (o *InstanceSnapshotsSnapshotsInner) HasForExport() bool`

HasForExport returns a boolean if a field has been set.

### GetForBackup

`func (o *InstanceSnapshotsSnapshotsInner) GetForBackup() bool`

GetForBackup returns the ForBackup field if non-nil, zero value otherwise.

### GetForBackupOk

`func (o *InstanceSnapshotsSnapshotsInner) GetForBackupOk() (*bool, bool)`

GetForBackupOk returns a tuple with the ForBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForBackup

`func (o *InstanceSnapshotsSnapshotsInner) SetForBackup(v bool)`

SetForBackup sets ForBackup field to given value.

### HasForBackup

`func (o *InstanceSnapshotsSnapshotsInner) HasForBackup() bool`

HasForBackup returns a boolean if a field has been set.

### GetDateCreated

`func (o *InstanceSnapshotsSnapshotsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InstanceSnapshotsSnapshotsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InstanceSnapshotsSnapshotsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InstanceSnapshotsSnapshotsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



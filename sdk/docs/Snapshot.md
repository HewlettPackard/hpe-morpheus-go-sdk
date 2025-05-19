# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SnapshotType** | Pointer to **string** |  | [optional] 
**SnapshotCreated** | Pointer to **time.Time** |  | [optional] 
**Zone** | Pointer to [**SnapshotsInstance200ResponseSnapshotsInnerZone**](SnapshotsInstance200ResponseSnapshotsInnerZone.md) |  | [optional] 
**Datastore** | Pointer to **string** |  | [optional] 
**ParentSnapshot** | Pointer to **string** |  | [optional] 
**SnapshotFiles** | Pointer to [**[]SnapshotsInstance200ResponseSnapshotsInnerSnapshotFilesInner**](SnapshotsInstance200ResponseSnapshotsInnerSnapshotFilesInner.md) |  | [optional] 
**CurrentlyActive** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSnapshot

`func NewSnapshot() *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Snapshot) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Snapshot) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Snapshot) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Snapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Snapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Snapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Snapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Snapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Snapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Snapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Snapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Snapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *Snapshot) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Snapshot) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Snapshot) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Snapshot) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetStatus

`func (o *Snapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Snapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Snapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Snapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetState

`func (o *Snapshot) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Snapshot) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Snapshot) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Snapshot) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSnapshotType

`func (o *Snapshot) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *Snapshot) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *Snapshot) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *Snapshot) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetSnapshotCreated

`func (o *Snapshot) GetSnapshotCreated() time.Time`

GetSnapshotCreated returns the SnapshotCreated field if non-nil, zero value otherwise.

### GetSnapshotCreatedOk

`func (o *Snapshot) GetSnapshotCreatedOk() (*time.Time, bool)`

GetSnapshotCreatedOk returns a tuple with the SnapshotCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCreated

`func (o *Snapshot) SetSnapshotCreated(v time.Time)`

SetSnapshotCreated sets SnapshotCreated field to given value.

### HasSnapshotCreated

`func (o *Snapshot) HasSnapshotCreated() bool`

HasSnapshotCreated returns a boolean if a field has been set.

### GetZone

`func (o *Snapshot) GetZone() SnapshotsInstance200ResponseSnapshotsInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Snapshot) GetZoneOk() (*SnapshotsInstance200ResponseSnapshotsInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Snapshot) SetZone(v SnapshotsInstance200ResponseSnapshotsInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Snapshot) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetDatastore

`func (o *Snapshot) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *Snapshot) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *Snapshot) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *Snapshot) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetParentSnapshot

`func (o *Snapshot) GetParentSnapshot() string`

GetParentSnapshot returns the ParentSnapshot field if non-nil, zero value otherwise.

### GetParentSnapshotOk

`func (o *Snapshot) GetParentSnapshotOk() (*string, bool)`

GetParentSnapshotOk returns a tuple with the ParentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSnapshot

`func (o *Snapshot) SetParentSnapshot(v string)`

SetParentSnapshot sets ParentSnapshot field to given value.

### HasParentSnapshot

`func (o *Snapshot) HasParentSnapshot() bool`

HasParentSnapshot returns a boolean if a field has been set.

### GetSnapshotFiles

`func (o *Snapshot) GetSnapshotFiles() []SnapshotsInstance200ResponseSnapshotsInnerSnapshotFilesInner`

GetSnapshotFiles returns the SnapshotFiles field if non-nil, zero value otherwise.

### GetSnapshotFilesOk

`func (o *Snapshot) GetSnapshotFilesOk() (*[]SnapshotsInstance200ResponseSnapshotsInnerSnapshotFilesInner, bool)`

GetSnapshotFilesOk returns a tuple with the SnapshotFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotFiles

`func (o *Snapshot) SetSnapshotFiles(v []SnapshotsInstance200ResponseSnapshotsInnerSnapshotFilesInner)`

SetSnapshotFiles sets SnapshotFiles field to given value.

### HasSnapshotFiles

`func (o *Snapshot) HasSnapshotFiles() bool`

HasSnapshotFiles returns a boolean if a field has been set.

### GetCurrentlyActive

`func (o *Snapshot) GetCurrentlyActive() bool`

GetCurrentlyActive returns the CurrentlyActive field if non-nil, zero value otherwise.

### GetCurrentlyActiveOk

`func (o *Snapshot) GetCurrentlyActiveOk() (*bool, bool)`

GetCurrentlyActiveOk returns a tuple with the CurrentlyActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyActive

`func (o *Snapshot) SetCurrentlyActive(v bool)`

SetCurrentlyActive sets CurrentlyActive field to given value.

### HasCurrentlyActive

`func (o *Snapshot) HasCurrentlyActive() bool`

HasCurrentlyActive returns a boolean if a field has been set.

### GetDateCreated

`func (o *Snapshot) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Snapshot) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Snapshot) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Snapshot) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



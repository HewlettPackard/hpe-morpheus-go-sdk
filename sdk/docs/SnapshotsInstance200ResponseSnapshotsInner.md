# SnapshotsInstance200ResponseSnapshotsInner

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

### NewSnapshotsInstance200ResponseSnapshotsInner

`func NewSnapshotsInstance200ResponseSnapshotsInner() *SnapshotsInstance200ResponseSnapshotsInner`

NewSnapshotsInstance200ResponseSnapshotsInner instantiates a new SnapshotsInstance200ResponseSnapshotsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotsInstance200ResponseSnapshotsInnerWithDefaults

`func NewSnapshotsInstance200ResponseSnapshotsInnerWithDefaults() *SnapshotsInstance200ResponseSnapshotsInner`

NewSnapshotsInstance200ResponseSnapshotsInnerWithDefaults instantiates a new SnapshotsInstance200ResponseSnapshotsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetStatus

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetState

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSnapshotType

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetSnapshotCreated

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetSnapshotCreated() time.Time`

GetSnapshotCreated returns the SnapshotCreated field if non-nil, zero value otherwise.

### GetSnapshotCreatedOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetSnapshotCreatedOk() (*time.Time, bool)`

GetSnapshotCreatedOk returns a tuple with the SnapshotCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCreated

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetSnapshotCreated(v time.Time)`

SetSnapshotCreated sets SnapshotCreated field to given value.

### HasSnapshotCreated

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasSnapshotCreated() bool`

HasSnapshotCreated returns a boolean if a field has been set.

### GetZone

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetZone() SnapshotsInstance200ResponseSnapshotsInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetZoneOk() (*SnapshotsInstance200ResponseSnapshotsInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetZone(v SnapshotsInstance200ResponseSnapshotsInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetDatastore

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetParentSnapshot

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetParentSnapshot() string`

GetParentSnapshot returns the ParentSnapshot field if non-nil, zero value otherwise.

### GetParentSnapshotOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetParentSnapshotOk() (*string, bool)`

GetParentSnapshotOk returns a tuple with the ParentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSnapshot

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetParentSnapshot(v string)`

SetParentSnapshot sets ParentSnapshot field to given value.

### HasParentSnapshot

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasParentSnapshot() bool`

HasParentSnapshot returns a boolean if a field has been set.

### GetSnapshotFiles

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetSnapshotFiles() []SnapshotsInstance200ResponseSnapshotsInnerSnapshotFilesInner`

GetSnapshotFiles returns the SnapshotFiles field if non-nil, zero value otherwise.

### GetSnapshotFilesOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetSnapshotFilesOk() (*[]SnapshotsInstance200ResponseSnapshotsInnerSnapshotFilesInner, bool)`

GetSnapshotFilesOk returns a tuple with the SnapshotFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotFiles

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetSnapshotFiles(v []SnapshotsInstance200ResponseSnapshotsInnerSnapshotFilesInner)`

SetSnapshotFiles sets SnapshotFiles field to given value.

### HasSnapshotFiles

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasSnapshotFiles() bool`

HasSnapshotFiles returns a boolean if a field has been set.

### GetCurrentlyActive

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetCurrentlyActive() bool`

GetCurrentlyActive returns the CurrentlyActive field if non-nil, zero value otherwise.

### GetCurrentlyActiveOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetCurrentlyActiveOk() (*bool, bool)`

GetCurrentlyActiveOk returns a tuple with the CurrentlyActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyActive

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetCurrentlyActive(v bool)`

SetCurrentlyActive sets CurrentlyActive field to given value.

### HasCurrentlyActive

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasCurrentlyActive() bool`

HasCurrentlyActive returns a boolean if a field has been set.

### GetDateCreated

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SnapshotsInstance200ResponseSnapshotsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SnapshotsInstance200ResponseSnapshotsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SnapshotsInstance200ResponseSnapshotsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SaveDatastore200ResponseAllOfDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Code** | Pointer to **NullableString** |  | [optional] 
**DatastoreType** | [**SaveDatastore200ResponseAllOfDatastoreDatastoreType**](SaveDatastore200ResponseAllOfDatastoreDatastoreType.md) |  | 
**Config** | Pointer to **map[string]interface{}** | Configuration object. Settings vary by type. | [optional] 
**StorageServer** | Pointer to [**SaveDatastore200ResponseAllOfDatastoreStorageServer**](SaveDatastore200ResponseAllOfDatastoreStorageServer.md) |  | [optional] 
**Type** | **string** | The underlying type of the datastore, e.g. &#39;generic&#39;, &#39;cluster&#39;, &#39;nfs&#39;, &#39;vmfs&#39;, &#39;NFS Pool&#39;, &#39;Directory Pool&#39;, &#39;GFS2 Pool (Global File System 2)&#39;, &#39;storage-profile&#39;, &#39;ext&#39; | 
**Status** | **string** | The current status of the datastore, e.g. &#39;provisioned&#39;, &#39;provisioning&#39;, &#39;failed&#39;, &#39;warning&#39; | 
**StatusMessage** | Pointer to **string** | Additional details about the current status of the datastore | [optional] 
**Visibility** | Pointer to **string** | Visibility level of the datastore, can be &#39;private&#39; or &#39;public&#39;. If not specified, defaults to &#39;private&#39;. | [optional] 
**StorageSize** | Pointer to **NullableInt64** |  | [optional] 
**FreeSpace** | Pointer to **NullableInt64** |  | [optional] 
**DrsEnabled** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**AllowWrite** | Pointer to **bool** |  | [optional] 
**DefaultStore** | Pointer to **bool** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**AllowRead** | Pointer to **bool** |  | [optional] 
**AllowProvision** | Pointer to **bool** |  | [optional] 
**HeartBeatTarget** | Pointer to **bool** |  | [optional] 
**RefType** | Pointer to **string** | Type of the resource this datastore is associated with, e.g. &#39;ComputeZone&#39;, &#39;ComputeServerGroup&#39; | [optional] 
**RefId** | Pointer to **int64** | The ID of the resource this datastore is associated with, e.g. ComputeZone, ComputeServerGroup | [optional] 
**ExternalId** | Pointer to **string** | UUID for the datastore | [optional] 
**ExternalPath** | Pointer to **string** | External path for the datastore, e.g. mount path, datastore path, etc. | [optional] 
**ExternalType** | Pointer to **string** | External type for the datastore, e.g. rbd, netfs, dir:gfs2 | [optional] 
**Zone** | Pointer to [**SaveDatastore200ResponseAllOfDatastoreZone**](SaveDatastore200ResponseAllOfDatastoreZone.md) |  | [optional] 
**ZonePool** | Pointer to [**SaveDatastore200ResponseAllOfDatastoreZonePool**](SaveDatastore200ResponseAllOfDatastoreZonePool.md) |  | [optional] 
**Owner** | Pointer to [**SaveDatastore200ResponseAllOfDatastoreOwner**](SaveDatastore200ResponseAllOfDatastoreOwner.md) |  | [optional] 
**Datastores** | Pointer to [**[]SaveDatastore200ResponseAllOfDatastoreDatastoresInner**](SaveDatastore200ResponseAllOfDatastoreDatastoresInner.md) | List of datastores associated with this datastore, for use with vSphere clouds. | [optional] 
**Locations** | Pointer to [**[]SaveDatastore200ResponseAllOfDatastoreLocationsInner**](SaveDatastore200ResponseAllOfDatastoreLocationsInner.md) |  | [optional] 

## Methods

### NewSaveDatastore200ResponseAllOfDatastore

`func NewSaveDatastore200ResponseAllOfDatastore(id int64, name string, datastoreType SaveDatastore200ResponseAllOfDatastoreDatastoreType, type_ string, status string, ) *SaveDatastore200ResponseAllOfDatastore`

NewSaveDatastore200ResponseAllOfDatastore instantiates a new SaveDatastore200ResponseAllOfDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveDatastore200ResponseAllOfDatastoreWithDefaults

`func NewSaveDatastore200ResponseAllOfDatastoreWithDefaults() *SaveDatastore200ResponseAllOfDatastore`

NewSaveDatastore200ResponseAllOfDatastoreWithDefaults instantiates a new SaveDatastore200ResponseAllOfDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SaveDatastore200ResponseAllOfDatastore) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaveDatastore200ResponseAllOfDatastore) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *SaveDatastore200ResponseAllOfDatastore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaveDatastore200ResponseAllOfDatastore) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *SaveDatastore200ResponseAllOfDatastore) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SaveDatastore200ResponseAllOfDatastore) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SaveDatastore200ResponseAllOfDatastore) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *SaveDatastore200ResponseAllOfDatastore) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *SaveDatastore200ResponseAllOfDatastore) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDatastoreType

`func (o *SaveDatastore200ResponseAllOfDatastore) GetDatastoreType() SaveDatastore200ResponseAllOfDatastoreDatastoreType`

GetDatastoreType returns the DatastoreType field if non-nil, zero value otherwise.

### GetDatastoreTypeOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetDatastoreTypeOk() (*SaveDatastore200ResponseAllOfDatastoreDatastoreType, bool)`

GetDatastoreTypeOk returns a tuple with the DatastoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreType

`func (o *SaveDatastore200ResponseAllOfDatastore) SetDatastoreType(v SaveDatastore200ResponseAllOfDatastoreDatastoreType)`

SetDatastoreType sets DatastoreType field to given value.


### GetConfig

`func (o *SaveDatastore200ResponseAllOfDatastore) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SaveDatastore200ResponseAllOfDatastore) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SaveDatastore200ResponseAllOfDatastore) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStorageServer

`func (o *SaveDatastore200ResponseAllOfDatastore) GetStorageServer() SaveDatastore200ResponseAllOfDatastoreStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetStorageServerOk() (*SaveDatastore200ResponseAllOfDatastoreStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *SaveDatastore200ResponseAllOfDatastore) SetStorageServer(v SaveDatastore200ResponseAllOfDatastoreStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *SaveDatastore200ResponseAllOfDatastore) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetType

`func (o *SaveDatastore200ResponseAllOfDatastore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SaveDatastore200ResponseAllOfDatastore) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *SaveDatastore200ResponseAllOfDatastore) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SaveDatastore200ResponseAllOfDatastore) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *SaveDatastore200ResponseAllOfDatastore) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *SaveDatastore200ResponseAllOfDatastore) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *SaveDatastore200ResponseAllOfDatastore) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetVisibility

`func (o *SaveDatastore200ResponseAllOfDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SaveDatastore200ResponseAllOfDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SaveDatastore200ResponseAllOfDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStorageSize

`func (o *SaveDatastore200ResponseAllOfDatastore) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *SaveDatastore200ResponseAllOfDatastore) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *SaveDatastore200ResponseAllOfDatastore) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### SetStorageSizeNil

`func (o *SaveDatastore200ResponseAllOfDatastore) SetStorageSizeNil(b bool)`

 SetStorageSizeNil sets the value for StorageSize to be an explicit nil

### UnsetStorageSize
`func (o *SaveDatastore200ResponseAllOfDatastore) UnsetStorageSize()`

UnsetStorageSize ensures that no value is present for StorageSize, not even an explicit nil
### GetFreeSpace

`func (o *SaveDatastore200ResponseAllOfDatastore) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *SaveDatastore200ResponseAllOfDatastore) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *SaveDatastore200ResponseAllOfDatastore) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *SaveDatastore200ResponseAllOfDatastore) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *SaveDatastore200ResponseAllOfDatastore) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
### GetDrsEnabled

`func (o *SaveDatastore200ResponseAllOfDatastore) GetDrsEnabled() bool`

GetDrsEnabled returns the DrsEnabled field if non-nil, zero value otherwise.

### GetDrsEnabledOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetDrsEnabledOk() (*bool, bool)`

GetDrsEnabledOk returns a tuple with the DrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrsEnabled

`func (o *SaveDatastore200ResponseAllOfDatastore) SetDrsEnabled(v bool)`

SetDrsEnabled sets DrsEnabled field to given value.

### HasDrsEnabled

`func (o *SaveDatastore200ResponseAllOfDatastore) HasDrsEnabled() bool`

HasDrsEnabled returns a boolean if a field has been set.

### GetActive

`func (o *SaveDatastore200ResponseAllOfDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SaveDatastore200ResponseAllOfDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SaveDatastore200ResponseAllOfDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowWrite

`func (o *SaveDatastore200ResponseAllOfDatastore) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *SaveDatastore200ResponseAllOfDatastore) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *SaveDatastore200ResponseAllOfDatastore) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.

### GetDefaultStore

`func (o *SaveDatastore200ResponseAllOfDatastore) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *SaveDatastore200ResponseAllOfDatastore) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *SaveDatastore200ResponseAllOfDatastore) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetOnline

`func (o *SaveDatastore200ResponseAllOfDatastore) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SaveDatastore200ResponseAllOfDatastore) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *SaveDatastore200ResponseAllOfDatastore) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetAllowRead

`func (o *SaveDatastore200ResponseAllOfDatastore) GetAllowRead() bool`

GetAllowRead returns the AllowRead field if non-nil, zero value otherwise.

### GetAllowReadOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetAllowReadOk() (*bool, bool)`

GetAllowReadOk returns a tuple with the AllowRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRead

`func (o *SaveDatastore200ResponseAllOfDatastore) SetAllowRead(v bool)`

SetAllowRead sets AllowRead field to given value.

### HasAllowRead

`func (o *SaveDatastore200ResponseAllOfDatastore) HasAllowRead() bool`

HasAllowRead returns a boolean if a field has been set.

### GetAllowProvision

`func (o *SaveDatastore200ResponseAllOfDatastore) GetAllowProvision() bool`

GetAllowProvision returns the AllowProvision field if non-nil, zero value otherwise.

### GetAllowProvisionOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetAllowProvisionOk() (*bool, bool)`

GetAllowProvisionOk returns a tuple with the AllowProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvision

`func (o *SaveDatastore200ResponseAllOfDatastore) SetAllowProvision(v bool)`

SetAllowProvision sets AllowProvision field to given value.

### HasAllowProvision

`func (o *SaveDatastore200ResponseAllOfDatastore) HasAllowProvision() bool`

HasAllowProvision returns a boolean if a field has been set.

### GetHeartBeatTarget

`func (o *SaveDatastore200ResponseAllOfDatastore) GetHeartBeatTarget() bool`

GetHeartBeatTarget returns the HeartBeatTarget field if non-nil, zero value otherwise.

### GetHeartBeatTargetOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetHeartBeatTargetOk() (*bool, bool)`

GetHeartBeatTargetOk returns a tuple with the HeartBeatTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartBeatTarget

`func (o *SaveDatastore200ResponseAllOfDatastore) SetHeartBeatTarget(v bool)`

SetHeartBeatTarget sets HeartBeatTarget field to given value.

### HasHeartBeatTarget

`func (o *SaveDatastore200ResponseAllOfDatastore) HasHeartBeatTarget() bool`

HasHeartBeatTarget returns a boolean if a field has been set.

### GetRefType

`func (o *SaveDatastore200ResponseAllOfDatastore) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *SaveDatastore200ResponseAllOfDatastore) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *SaveDatastore200ResponseAllOfDatastore) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *SaveDatastore200ResponseAllOfDatastore) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *SaveDatastore200ResponseAllOfDatastore) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *SaveDatastore200ResponseAllOfDatastore) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetExternalId

`func (o *SaveDatastore200ResponseAllOfDatastore) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SaveDatastore200ResponseAllOfDatastore) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SaveDatastore200ResponseAllOfDatastore) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalPath

`func (o *SaveDatastore200ResponseAllOfDatastore) GetExternalPath() string`

GetExternalPath returns the ExternalPath field if non-nil, zero value otherwise.

### GetExternalPathOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetExternalPathOk() (*string, bool)`

GetExternalPathOk returns a tuple with the ExternalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPath

`func (o *SaveDatastore200ResponseAllOfDatastore) SetExternalPath(v string)`

SetExternalPath sets ExternalPath field to given value.

### HasExternalPath

`func (o *SaveDatastore200ResponseAllOfDatastore) HasExternalPath() bool`

HasExternalPath returns a boolean if a field has been set.

### GetExternalType

`func (o *SaveDatastore200ResponseAllOfDatastore) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *SaveDatastore200ResponseAllOfDatastore) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *SaveDatastore200ResponseAllOfDatastore) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetZone

`func (o *SaveDatastore200ResponseAllOfDatastore) GetZone() SaveDatastore200ResponseAllOfDatastoreZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetZoneOk() (*SaveDatastore200ResponseAllOfDatastoreZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *SaveDatastore200ResponseAllOfDatastore) SetZone(v SaveDatastore200ResponseAllOfDatastoreZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *SaveDatastore200ResponseAllOfDatastore) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *SaveDatastore200ResponseAllOfDatastore) GetZonePool() SaveDatastore200ResponseAllOfDatastoreZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetZonePoolOk() (*SaveDatastore200ResponseAllOfDatastoreZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *SaveDatastore200ResponseAllOfDatastore) SetZonePool(v SaveDatastore200ResponseAllOfDatastoreZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *SaveDatastore200ResponseAllOfDatastore) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetOwner

`func (o *SaveDatastore200ResponseAllOfDatastore) GetOwner() SaveDatastore200ResponseAllOfDatastoreOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetOwnerOk() (*SaveDatastore200ResponseAllOfDatastoreOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SaveDatastore200ResponseAllOfDatastore) SetOwner(v SaveDatastore200ResponseAllOfDatastoreOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SaveDatastore200ResponseAllOfDatastore) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDatastores

`func (o *SaveDatastore200ResponseAllOfDatastore) GetDatastores() []SaveDatastore200ResponseAllOfDatastoreDatastoresInner`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetDatastoresOk() (*[]SaveDatastore200ResponseAllOfDatastoreDatastoresInner, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *SaveDatastore200ResponseAllOfDatastore) SetDatastores(v []SaveDatastore200ResponseAllOfDatastoreDatastoresInner)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *SaveDatastore200ResponseAllOfDatastore) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetLocations

`func (o *SaveDatastore200ResponseAllOfDatastore) GetLocations() []SaveDatastore200ResponseAllOfDatastoreLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SaveDatastore200ResponseAllOfDatastore) GetLocationsOk() (*[]SaveDatastore200ResponseAllOfDatastoreLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SaveDatastore200ResponseAllOfDatastore) SetLocations(v []SaveDatastore200ResponseAllOfDatastoreLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *SaveDatastore200ResponseAllOfDatastore) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StorageDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Code** | Pointer to **NullableString** |  | [optional] 
**DatastoreType** | [**StorageDatastoreDatastoreType**](StorageDatastoreDatastoreType.md) |  | 
**Config** | Pointer to **map[string]interface{}** | Configuration object. Settings vary by type. | [optional] 
**StorageServer** | Pointer to [**StorageDatastoreStorageServer**](StorageDatastoreStorageServer.md) |  | [optional] 
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
**Zone** | Pointer to [**StorageDatastoreZone**](StorageDatastoreZone.md) |  | [optional] 
**ZonePool** | Pointer to [**StorageDatastoreZonePool**](StorageDatastoreZonePool.md) |  | [optional] 
**Owner** | Pointer to [**StorageDatastoreOwner**](StorageDatastoreOwner.md) |  | [optional] 
**Datastores** | Pointer to [**[]StorageDatastoreDatastoresInner**](StorageDatastoreDatastoresInner.md) | List of datastores associated with this datastore, for use with vSphere clouds. | [optional] 
**Locations** | Pointer to [**[]StorageDatastoreLocationsInner**](StorageDatastoreLocationsInner.md) |  | [optional] 

## Methods

### NewStorageDatastore

`func NewStorageDatastore(id int64, name string, datastoreType StorageDatastoreDatastoreType, type_ string, status string, ) *StorageDatastore`

NewStorageDatastore instantiates a new StorageDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDatastoreWithDefaults

`func NewStorageDatastoreWithDefaults() *StorageDatastore`

NewStorageDatastoreWithDefaults instantiates a new StorageDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageDatastore) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageDatastore) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageDatastore) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *StorageDatastore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageDatastore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageDatastore) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *StorageDatastore) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StorageDatastore) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StorageDatastore) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *StorageDatastore) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *StorageDatastore) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *StorageDatastore) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDatastoreType

`func (o *StorageDatastore) GetDatastoreType() StorageDatastoreDatastoreType`

GetDatastoreType returns the DatastoreType field if non-nil, zero value otherwise.

### GetDatastoreTypeOk

`func (o *StorageDatastore) GetDatastoreTypeOk() (*StorageDatastoreDatastoreType, bool)`

GetDatastoreTypeOk returns a tuple with the DatastoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreType

`func (o *StorageDatastore) SetDatastoreType(v StorageDatastoreDatastoreType)`

SetDatastoreType sets DatastoreType field to given value.


### GetConfig

`func (o *StorageDatastore) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *StorageDatastore) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *StorageDatastore) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *StorageDatastore) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStorageServer

`func (o *StorageDatastore) GetStorageServer() StorageDatastoreStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *StorageDatastore) GetStorageServerOk() (*StorageDatastoreStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *StorageDatastore) SetStorageServer(v StorageDatastoreStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *StorageDatastore) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetType

`func (o *StorageDatastore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageDatastore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageDatastore) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *StorageDatastore) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageDatastore) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageDatastore) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *StorageDatastore) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *StorageDatastore) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *StorageDatastore) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *StorageDatastore) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetVisibility

`func (o *StorageDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *StorageDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *StorageDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *StorageDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStorageSize

`func (o *StorageDatastore) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *StorageDatastore) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *StorageDatastore) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *StorageDatastore) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### SetStorageSizeNil

`func (o *StorageDatastore) SetStorageSizeNil(b bool)`

 SetStorageSizeNil sets the value for StorageSize to be an explicit nil

### UnsetStorageSize
`func (o *StorageDatastore) UnsetStorageSize()`

UnsetStorageSize ensures that no value is present for StorageSize, not even an explicit nil
### GetFreeSpace

`func (o *StorageDatastore) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *StorageDatastore) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *StorageDatastore) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *StorageDatastore) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *StorageDatastore) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *StorageDatastore) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
### GetDrsEnabled

`func (o *StorageDatastore) GetDrsEnabled() bool`

GetDrsEnabled returns the DrsEnabled field if non-nil, zero value otherwise.

### GetDrsEnabledOk

`func (o *StorageDatastore) GetDrsEnabledOk() (*bool, bool)`

GetDrsEnabledOk returns a tuple with the DrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrsEnabled

`func (o *StorageDatastore) SetDrsEnabled(v bool)`

SetDrsEnabled sets DrsEnabled field to given value.

### HasDrsEnabled

`func (o *StorageDatastore) HasDrsEnabled() bool`

HasDrsEnabled returns a boolean if a field has been set.

### GetActive

`func (o *StorageDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StorageDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StorageDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *StorageDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowWrite

`func (o *StorageDatastore) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *StorageDatastore) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *StorageDatastore) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *StorageDatastore) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.

### GetDefaultStore

`func (o *StorageDatastore) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *StorageDatastore) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *StorageDatastore) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *StorageDatastore) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetOnline

`func (o *StorageDatastore) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *StorageDatastore) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *StorageDatastore) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *StorageDatastore) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetAllowRead

`func (o *StorageDatastore) GetAllowRead() bool`

GetAllowRead returns the AllowRead field if non-nil, zero value otherwise.

### GetAllowReadOk

`func (o *StorageDatastore) GetAllowReadOk() (*bool, bool)`

GetAllowReadOk returns a tuple with the AllowRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRead

`func (o *StorageDatastore) SetAllowRead(v bool)`

SetAllowRead sets AllowRead field to given value.

### HasAllowRead

`func (o *StorageDatastore) HasAllowRead() bool`

HasAllowRead returns a boolean if a field has been set.

### GetAllowProvision

`func (o *StorageDatastore) GetAllowProvision() bool`

GetAllowProvision returns the AllowProvision field if non-nil, zero value otherwise.

### GetAllowProvisionOk

`func (o *StorageDatastore) GetAllowProvisionOk() (*bool, bool)`

GetAllowProvisionOk returns a tuple with the AllowProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvision

`func (o *StorageDatastore) SetAllowProvision(v bool)`

SetAllowProvision sets AllowProvision field to given value.

### HasAllowProvision

`func (o *StorageDatastore) HasAllowProvision() bool`

HasAllowProvision returns a boolean if a field has been set.

### GetHeartBeatTarget

`func (o *StorageDatastore) GetHeartBeatTarget() bool`

GetHeartBeatTarget returns the HeartBeatTarget field if non-nil, zero value otherwise.

### GetHeartBeatTargetOk

`func (o *StorageDatastore) GetHeartBeatTargetOk() (*bool, bool)`

GetHeartBeatTargetOk returns a tuple with the HeartBeatTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartBeatTarget

`func (o *StorageDatastore) SetHeartBeatTarget(v bool)`

SetHeartBeatTarget sets HeartBeatTarget field to given value.

### HasHeartBeatTarget

`func (o *StorageDatastore) HasHeartBeatTarget() bool`

HasHeartBeatTarget returns a boolean if a field has been set.

### GetRefType

`func (o *StorageDatastore) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *StorageDatastore) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *StorageDatastore) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *StorageDatastore) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *StorageDatastore) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *StorageDatastore) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *StorageDatastore) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *StorageDatastore) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetExternalId

`func (o *StorageDatastore) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *StorageDatastore) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *StorageDatastore) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *StorageDatastore) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalPath

`func (o *StorageDatastore) GetExternalPath() string`

GetExternalPath returns the ExternalPath field if non-nil, zero value otherwise.

### GetExternalPathOk

`func (o *StorageDatastore) GetExternalPathOk() (*string, bool)`

GetExternalPathOk returns a tuple with the ExternalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPath

`func (o *StorageDatastore) SetExternalPath(v string)`

SetExternalPath sets ExternalPath field to given value.

### HasExternalPath

`func (o *StorageDatastore) HasExternalPath() bool`

HasExternalPath returns a boolean if a field has been set.

### GetExternalType

`func (o *StorageDatastore) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *StorageDatastore) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *StorageDatastore) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *StorageDatastore) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetZone

`func (o *StorageDatastore) GetZone() StorageDatastoreZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *StorageDatastore) GetZoneOk() (*StorageDatastoreZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *StorageDatastore) SetZone(v StorageDatastoreZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *StorageDatastore) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *StorageDatastore) GetZonePool() StorageDatastoreZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *StorageDatastore) GetZonePoolOk() (*StorageDatastoreZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *StorageDatastore) SetZonePool(v StorageDatastoreZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *StorageDatastore) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetOwner

`func (o *StorageDatastore) GetOwner() StorageDatastoreOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StorageDatastore) GetOwnerOk() (*StorageDatastoreOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StorageDatastore) SetOwner(v StorageDatastoreOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *StorageDatastore) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDatastores

`func (o *StorageDatastore) GetDatastores() []StorageDatastoreDatastoresInner`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *StorageDatastore) GetDatastoresOk() (*[]StorageDatastoreDatastoresInner, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *StorageDatastore) SetDatastores(v []StorageDatastoreDatastoresInner)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *StorageDatastore) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetLocations

`func (o *StorageDatastore) GetLocations() []StorageDatastoreLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *StorageDatastore) GetLocationsOk() (*[]StorageDatastoreLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *StorageDatastore) SetLocations(v []StorageDatastoreLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *StorageDatastore) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



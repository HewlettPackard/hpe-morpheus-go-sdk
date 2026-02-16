# GetDatastores200ResponseAllOfDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Code** | Pointer to **NullableString** |  | [optional] 
**DatastoreType** | [**GetDatastores200ResponseAllOfDatastoreDatastoreType**](GetDatastores200ResponseAllOfDatastoreDatastoreType.md) |  | 
**Config** | Pointer to **map[string]interface{}** | Configuration object. Settings vary by type. | [optional] 
**StorageServer** | Pointer to [**GetDatastores200ResponseAllOfDatastoreStorageServer**](GetDatastores200ResponseAllOfDatastoreStorageServer.md) |  | [optional] 
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
**Zone** | Pointer to [**GetDatastores200ResponseAllOfDatastoreZone**](GetDatastores200ResponseAllOfDatastoreZone.md) |  | [optional] 
**ZonePool** | Pointer to [**GetDatastores200ResponseAllOfDatastoreZonePool**](GetDatastores200ResponseAllOfDatastoreZonePool.md) |  | [optional] 
**Owner** | Pointer to [**GetDatastores200ResponseAllOfDatastoreOwner**](GetDatastores200ResponseAllOfDatastoreOwner.md) |  | [optional] 
**Datastores** | Pointer to [**[]GetDatastores200ResponseAllOfDatastoreDatastoresInner**](GetDatastores200ResponseAllOfDatastoreDatastoresInner.md) | List of datastores associated with this datastore, for use with vSphere clouds. | [optional] 
**Locations** | Pointer to [**[]GetDatastores200ResponseAllOfDatastoreLocationsInner**](GetDatastores200ResponseAllOfDatastoreLocationsInner.md) |  | [optional] 

## Methods

### NewGetDatastores200ResponseAllOfDatastore

`func NewGetDatastores200ResponseAllOfDatastore(id int64, name string, datastoreType GetDatastores200ResponseAllOfDatastoreDatastoreType, type_ string, status string, ) *GetDatastores200ResponseAllOfDatastore`

NewGetDatastores200ResponseAllOfDatastore instantiates a new GetDatastores200ResponseAllOfDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatastores200ResponseAllOfDatastoreWithDefaults

`func NewGetDatastores200ResponseAllOfDatastoreWithDefaults() *GetDatastores200ResponseAllOfDatastore`

NewGetDatastores200ResponseAllOfDatastoreWithDefaults instantiates a new GetDatastores200ResponseAllOfDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDatastores200ResponseAllOfDatastore) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDatastores200ResponseAllOfDatastore) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *GetDatastores200ResponseAllOfDatastore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDatastores200ResponseAllOfDatastore) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *GetDatastores200ResponseAllOfDatastore) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetDatastores200ResponseAllOfDatastore) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetDatastores200ResponseAllOfDatastore) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetDatastores200ResponseAllOfDatastore) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetDatastores200ResponseAllOfDatastore) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDatastoreType

`func (o *GetDatastores200ResponseAllOfDatastore) GetDatastoreType() GetDatastores200ResponseAllOfDatastoreDatastoreType`

GetDatastoreType returns the DatastoreType field if non-nil, zero value otherwise.

### GetDatastoreTypeOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetDatastoreTypeOk() (*GetDatastores200ResponseAllOfDatastoreDatastoreType, bool)`

GetDatastoreTypeOk returns a tuple with the DatastoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreType

`func (o *GetDatastores200ResponseAllOfDatastore) SetDatastoreType(v GetDatastores200ResponseAllOfDatastoreDatastoreType)`

SetDatastoreType sets DatastoreType field to given value.


### GetConfig

`func (o *GetDatastores200ResponseAllOfDatastore) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetDatastores200ResponseAllOfDatastore) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetDatastores200ResponseAllOfDatastore) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStorageServer

`func (o *GetDatastores200ResponseAllOfDatastore) GetStorageServer() GetDatastores200ResponseAllOfDatastoreStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetStorageServerOk() (*GetDatastores200ResponseAllOfDatastoreStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *GetDatastores200ResponseAllOfDatastore) SetStorageServer(v GetDatastores200ResponseAllOfDatastoreStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *GetDatastores200ResponseAllOfDatastore) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetType

`func (o *GetDatastores200ResponseAllOfDatastore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDatastores200ResponseAllOfDatastore) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *GetDatastores200ResponseAllOfDatastore) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDatastores200ResponseAllOfDatastore) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *GetDatastores200ResponseAllOfDatastore) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetDatastores200ResponseAllOfDatastore) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetDatastores200ResponseAllOfDatastore) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetVisibility

`func (o *GetDatastores200ResponseAllOfDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetDatastores200ResponseAllOfDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetDatastores200ResponseAllOfDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStorageSize

`func (o *GetDatastores200ResponseAllOfDatastore) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *GetDatastores200ResponseAllOfDatastore) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *GetDatastores200ResponseAllOfDatastore) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### SetStorageSizeNil

`func (o *GetDatastores200ResponseAllOfDatastore) SetStorageSizeNil(b bool)`

 SetStorageSizeNil sets the value for StorageSize to be an explicit nil

### UnsetStorageSize
`func (o *GetDatastores200ResponseAllOfDatastore) UnsetStorageSize()`

UnsetStorageSize ensures that no value is present for StorageSize, not even an explicit nil
### GetFreeSpace

`func (o *GetDatastores200ResponseAllOfDatastore) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *GetDatastores200ResponseAllOfDatastore) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *GetDatastores200ResponseAllOfDatastore) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *GetDatastores200ResponseAllOfDatastore) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *GetDatastores200ResponseAllOfDatastore) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
### GetDrsEnabled

`func (o *GetDatastores200ResponseAllOfDatastore) GetDrsEnabled() bool`

GetDrsEnabled returns the DrsEnabled field if non-nil, zero value otherwise.

### GetDrsEnabledOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetDrsEnabledOk() (*bool, bool)`

GetDrsEnabledOk returns a tuple with the DrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrsEnabled

`func (o *GetDatastores200ResponseAllOfDatastore) SetDrsEnabled(v bool)`

SetDrsEnabled sets DrsEnabled field to given value.

### HasDrsEnabled

`func (o *GetDatastores200ResponseAllOfDatastore) HasDrsEnabled() bool`

HasDrsEnabled returns a boolean if a field has been set.

### GetActive

`func (o *GetDatastores200ResponseAllOfDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetDatastores200ResponseAllOfDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetDatastores200ResponseAllOfDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowWrite

`func (o *GetDatastores200ResponseAllOfDatastore) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *GetDatastores200ResponseAllOfDatastore) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *GetDatastores200ResponseAllOfDatastore) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.

### GetDefaultStore

`func (o *GetDatastores200ResponseAllOfDatastore) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *GetDatastores200ResponseAllOfDatastore) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *GetDatastores200ResponseAllOfDatastore) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetOnline

`func (o *GetDatastores200ResponseAllOfDatastore) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *GetDatastores200ResponseAllOfDatastore) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *GetDatastores200ResponseAllOfDatastore) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetAllowRead

`func (o *GetDatastores200ResponseAllOfDatastore) GetAllowRead() bool`

GetAllowRead returns the AllowRead field if non-nil, zero value otherwise.

### GetAllowReadOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetAllowReadOk() (*bool, bool)`

GetAllowReadOk returns a tuple with the AllowRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRead

`func (o *GetDatastores200ResponseAllOfDatastore) SetAllowRead(v bool)`

SetAllowRead sets AllowRead field to given value.

### HasAllowRead

`func (o *GetDatastores200ResponseAllOfDatastore) HasAllowRead() bool`

HasAllowRead returns a boolean if a field has been set.

### GetAllowProvision

`func (o *GetDatastores200ResponseAllOfDatastore) GetAllowProvision() bool`

GetAllowProvision returns the AllowProvision field if non-nil, zero value otherwise.

### GetAllowProvisionOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetAllowProvisionOk() (*bool, bool)`

GetAllowProvisionOk returns a tuple with the AllowProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvision

`func (o *GetDatastores200ResponseAllOfDatastore) SetAllowProvision(v bool)`

SetAllowProvision sets AllowProvision field to given value.

### HasAllowProvision

`func (o *GetDatastores200ResponseAllOfDatastore) HasAllowProvision() bool`

HasAllowProvision returns a boolean if a field has been set.

### GetHeartBeatTarget

`func (o *GetDatastores200ResponseAllOfDatastore) GetHeartBeatTarget() bool`

GetHeartBeatTarget returns the HeartBeatTarget field if non-nil, zero value otherwise.

### GetHeartBeatTargetOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetHeartBeatTargetOk() (*bool, bool)`

GetHeartBeatTargetOk returns a tuple with the HeartBeatTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartBeatTarget

`func (o *GetDatastores200ResponseAllOfDatastore) SetHeartBeatTarget(v bool)`

SetHeartBeatTarget sets HeartBeatTarget field to given value.

### HasHeartBeatTarget

`func (o *GetDatastores200ResponseAllOfDatastore) HasHeartBeatTarget() bool`

HasHeartBeatTarget returns a boolean if a field has been set.

### GetRefType

`func (o *GetDatastores200ResponseAllOfDatastore) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetDatastores200ResponseAllOfDatastore) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetDatastores200ResponseAllOfDatastore) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetDatastores200ResponseAllOfDatastore) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetDatastores200ResponseAllOfDatastore) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetDatastores200ResponseAllOfDatastore) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetExternalId

`func (o *GetDatastores200ResponseAllOfDatastore) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetDatastores200ResponseAllOfDatastore) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetDatastores200ResponseAllOfDatastore) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalPath

`func (o *GetDatastores200ResponseAllOfDatastore) GetExternalPath() string`

GetExternalPath returns the ExternalPath field if non-nil, zero value otherwise.

### GetExternalPathOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetExternalPathOk() (*string, bool)`

GetExternalPathOk returns a tuple with the ExternalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPath

`func (o *GetDatastores200ResponseAllOfDatastore) SetExternalPath(v string)`

SetExternalPath sets ExternalPath field to given value.

### HasExternalPath

`func (o *GetDatastores200ResponseAllOfDatastore) HasExternalPath() bool`

HasExternalPath returns a boolean if a field has been set.

### GetExternalType

`func (o *GetDatastores200ResponseAllOfDatastore) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *GetDatastores200ResponseAllOfDatastore) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *GetDatastores200ResponseAllOfDatastore) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetZone

`func (o *GetDatastores200ResponseAllOfDatastore) GetZone() GetDatastores200ResponseAllOfDatastoreZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetZoneOk() (*GetDatastores200ResponseAllOfDatastoreZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetDatastores200ResponseAllOfDatastore) SetZone(v GetDatastores200ResponseAllOfDatastoreZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetDatastores200ResponseAllOfDatastore) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *GetDatastores200ResponseAllOfDatastore) GetZonePool() GetDatastores200ResponseAllOfDatastoreZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetZonePoolOk() (*GetDatastores200ResponseAllOfDatastoreZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *GetDatastores200ResponseAllOfDatastore) SetZonePool(v GetDatastores200ResponseAllOfDatastoreZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *GetDatastores200ResponseAllOfDatastore) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetOwner

`func (o *GetDatastores200ResponseAllOfDatastore) GetOwner() GetDatastores200ResponseAllOfDatastoreOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetOwnerOk() (*GetDatastores200ResponseAllOfDatastoreOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetDatastores200ResponseAllOfDatastore) SetOwner(v GetDatastores200ResponseAllOfDatastoreOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetDatastores200ResponseAllOfDatastore) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDatastores

`func (o *GetDatastores200ResponseAllOfDatastore) GetDatastores() []GetDatastores200ResponseAllOfDatastoreDatastoresInner`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetDatastoresOk() (*[]GetDatastores200ResponseAllOfDatastoreDatastoresInner, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *GetDatastores200ResponseAllOfDatastore) SetDatastores(v []GetDatastores200ResponseAllOfDatastoreDatastoresInner)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *GetDatastores200ResponseAllOfDatastore) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetLocations

`func (o *GetDatastores200ResponseAllOfDatastore) GetLocations() []GetDatastores200ResponseAllOfDatastoreLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *GetDatastores200ResponseAllOfDatastore) GetLocationsOk() (*[]GetDatastores200ResponseAllOfDatastoreLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *GetDatastores200ResponseAllOfDatastore) SetLocations(v []GetDatastores200ResponseAllOfDatastoreLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *GetDatastores200ResponseAllOfDatastore) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ListDatastores200ResponseAllOfDatastoresInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Code** | Pointer to **NullableString** |  | [optional] 
**DatastoreType** | [**ListDatastores200ResponseAllOfDatastoresInnerDatastoreType**](ListDatastores200ResponseAllOfDatastoresInnerDatastoreType.md) |  | 
**Config** | Pointer to **map[string]interface{}** | Configuration object. Settings vary by type. | [optional] 
**StorageServer** | Pointer to [**ListDatastores200ResponseAllOfDatastoresInnerStorageServer**](ListDatastores200ResponseAllOfDatastoresInnerStorageServer.md) |  | [optional] 
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
**Zone** | Pointer to [**ListDatastores200ResponseAllOfDatastoresInnerZone**](ListDatastores200ResponseAllOfDatastoresInnerZone.md) |  | [optional] 
**ZonePool** | Pointer to [**ListDatastores200ResponseAllOfDatastoresInnerZonePool**](ListDatastores200ResponseAllOfDatastoresInnerZonePool.md) |  | [optional] 
**Owner** | Pointer to [**ListDatastores200ResponseAllOfDatastoresInnerOwner**](ListDatastores200ResponseAllOfDatastoresInnerOwner.md) |  | [optional] 
**Datastores** | Pointer to [**[]ListDatastores200ResponseAllOfDatastoresInnerDatastoresInner**](ListDatastores200ResponseAllOfDatastoresInnerDatastoresInner.md) | List of datastores associated with this datastore, for use with vSphere clouds. | [optional] 
**Locations** | Pointer to [**[]ListDatastores200ResponseAllOfDatastoresInnerLocationsInner**](ListDatastores200ResponseAllOfDatastoresInnerLocationsInner.md) |  | [optional] 

## Methods

### NewListDatastores200ResponseAllOfDatastoresInner

`func NewListDatastores200ResponseAllOfDatastoresInner(id int64, name string, datastoreType ListDatastores200ResponseAllOfDatastoresInnerDatastoreType, type_ string, status string, ) *ListDatastores200ResponseAllOfDatastoresInner`

NewListDatastores200ResponseAllOfDatastoresInner instantiates a new ListDatastores200ResponseAllOfDatastoresInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDatastores200ResponseAllOfDatastoresInnerWithDefaults

`func NewListDatastores200ResponseAllOfDatastoresInnerWithDefaults() *ListDatastores200ResponseAllOfDatastoresInner`

NewListDatastores200ResponseAllOfDatastoresInnerWithDefaults instantiates a new ListDatastores200ResponseAllOfDatastoresInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ListDatastores200ResponseAllOfDatastoresInner) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDatastoreType

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetDatastoreType() ListDatastores200ResponseAllOfDatastoresInnerDatastoreType`

GetDatastoreType returns the DatastoreType field if non-nil, zero value otherwise.

### GetDatastoreTypeOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetDatastoreTypeOk() (*ListDatastores200ResponseAllOfDatastoresInnerDatastoreType, bool)`

GetDatastoreTypeOk returns a tuple with the DatastoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreType

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetDatastoreType(v ListDatastores200ResponseAllOfDatastoresInnerDatastoreType)`

SetDatastoreType sets DatastoreType field to given value.


### GetConfig

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStorageServer

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetStorageServer() ListDatastores200ResponseAllOfDatastoresInnerStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetStorageServerOk() (*ListDatastores200ResponseAllOfDatastoresInnerStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetStorageServer(v ListDatastores200ResponseAllOfDatastoresInnerStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetType

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetVisibility

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStorageSize

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### SetStorageSizeNil

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetStorageSizeNil(b bool)`

 SetStorageSizeNil sets the value for StorageSize to be an explicit nil

### UnsetStorageSize
`func (o *ListDatastores200ResponseAllOfDatastoresInner) UnsetStorageSize()`

UnsetStorageSize ensures that no value is present for StorageSize, not even an explicit nil
### GetFreeSpace

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *ListDatastores200ResponseAllOfDatastoresInner) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
### GetDrsEnabled

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetDrsEnabled() bool`

GetDrsEnabled returns the DrsEnabled field if non-nil, zero value otherwise.

### GetDrsEnabledOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetDrsEnabledOk() (*bool, bool)`

GetDrsEnabledOk returns a tuple with the DrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrsEnabled

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetDrsEnabled(v bool)`

SetDrsEnabled sets DrsEnabled field to given value.

### HasDrsEnabled

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasDrsEnabled() bool`

HasDrsEnabled returns a boolean if a field has been set.

### GetActive

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowWrite

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.

### GetDefaultStore

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetOnline

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetAllowRead

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetAllowRead() bool`

GetAllowRead returns the AllowRead field if non-nil, zero value otherwise.

### GetAllowReadOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetAllowReadOk() (*bool, bool)`

GetAllowReadOk returns a tuple with the AllowRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRead

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetAllowRead(v bool)`

SetAllowRead sets AllowRead field to given value.

### HasAllowRead

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasAllowRead() bool`

HasAllowRead returns a boolean if a field has been set.

### GetAllowProvision

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetAllowProvision() bool`

GetAllowProvision returns the AllowProvision field if non-nil, zero value otherwise.

### GetAllowProvisionOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetAllowProvisionOk() (*bool, bool)`

GetAllowProvisionOk returns a tuple with the AllowProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvision

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetAllowProvision(v bool)`

SetAllowProvision sets AllowProvision field to given value.

### HasAllowProvision

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasAllowProvision() bool`

HasAllowProvision returns a boolean if a field has been set.

### GetHeartBeatTarget

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetHeartBeatTarget() bool`

GetHeartBeatTarget returns the HeartBeatTarget field if non-nil, zero value otherwise.

### GetHeartBeatTargetOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetHeartBeatTargetOk() (*bool, bool)`

GetHeartBeatTargetOk returns a tuple with the HeartBeatTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartBeatTarget

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetHeartBeatTarget(v bool)`

SetHeartBeatTarget sets HeartBeatTarget field to given value.

### HasHeartBeatTarget

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasHeartBeatTarget() bool`

HasHeartBeatTarget returns a boolean if a field has been set.

### GetRefType

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalPath

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetExternalPath() string`

GetExternalPath returns the ExternalPath field if non-nil, zero value otherwise.

### GetExternalPathOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetExternalPathOk() (*string, bool)`

GetExternalPathOk returns a tuple with the ExternalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPath

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetExternalPath(v string)`

SetExternalPath sets ExternalPath field to given value.

### HasExternalPath

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasExternalPath() bool`

HasExternalPath returns a boolean if a field has been set.

### GetExternalType

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetZone

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetZone() ListDatastores200ResponseAllOfDatastoresInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetZoneOk() (*ListDatastores200ResponseAllOfDatastoresInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetZone(v ListDatastores200ResponseAllOfDatastoresInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetZonePool() ListDatastores200ResponseAllOfDatastoresInnerZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetZonePoolOk() (*ListDatastores200ResponseAllOfDatastoresInnerZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetZonePool(v ListDatastores200ResponseAllOfDatastoresInnerZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetOwner

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetOwner() ListDatastores200ResponseAllOfDatastoresInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetOwnerOk() (*ListDatastores200ResponseAllOfDatastoresInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetOwner(v ListDatastores200ResponseAllOfDatastoresInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDatastores

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetDatastores() []ListDatastores200ResponseAllOfDatastoresInnerDatastoresInner`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetDatastoresOk() (*[]ListDatastores200ResponseAllOfDatastoresInnerDatastoresInner, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetDatastores(v []ListDatastores200ResponseAllOfDatastoresInnerDatastoresInner)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetLocations

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetLocations() []ListDatastores200ResponseAllOfDatastoresInnerLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ListDatastores200ResponseAllOfDatastoresInner) GetLocationsOk() (*[]ListDatastores200ResponseAllOfDatastoresInnerLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ListDatastores200ResponseAllOfDatastoresInner) SetLocations(v []ListDatastores200ResponseAllOfDatastoresInnerLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ListDatastores200ResponseAllOfDatastoresInner) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



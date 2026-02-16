# UpdateDatastores200ResponseDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Code** | Pointer to **NullableString** |  | [optional] 
**DatastoreType** | [**UpdateDatastores200ResponseDatastoreAllOfDatastoreType**](UpdateDatastores200ResponseDatastoreAllOfDatastoreType.md) |  | 
**Config** | Pointer to **map[string]interface{}** | Configuration object. Settings vary by type. | [optional] 
**StorageServer** | Pointer to [**UpdateDatastores200ResponseDatastoreAllOfStorageServer**](UpdateDatastores200ResponseDatastoreAllOfStorageServer.md) |  | [optional] 
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
**Zone** | Pointer to [**UpdateDatastores200ResponseDatastoreAllOfZone**](UpdateDatastores200ResponseDatastoreAllOfZone.md) |  | [optional] 
**ZonePool** | Pointer to [**UpdateDatastores200ResponseDatastoreAllOfZonePool**](UpdateDatastores200ResponseDatastoreAllOfZonePool.md) |  | [optional] 
**Owner** | Pointer to [**UpdateDatastores200ResponseDatastoreAllOfOwner**](UpdateDatastores200ResponseDatastoreAllOfOwner.md) |  | [optional] 
**Datastores** | Pointer to [**[]UpdateDatastores200ResponseDatastoreAllOfDatastoresInner**](UpdateDatastores200ResponseDatastoreAllOfDatastoresInner.md) | List of datastores associated with this datastore, for use with vSphere clouds. | [optional] 
**Locations** | Pointer to [**[]UpdateDatastores200ResponseDatastoreAllOfLocationsInner**](UpdateDatastores200ResponseDatastoreAllOfLocationsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateDatastores200ResponseDatastore

`func NewUpdateDatastores200ResponseDatastore(id int64, name string, datastoreType UpdateDatastores200ResponseDatastoreAllOfDatastoreType, type_ string, status string, ) *UpdateDatastores200ResponseDatastore`

NewUpdateDatastores200ResponseDatastore instantiates a new UpdateDatastores200ResponseDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatastores200ResponseDatastoreWithDefaults

`func NewUpdateDatastores200ResponseDatastoreWithDefaults() *UpdateDatastores200ResponseDatastore`

NewUpdateDatastores200ResponseDatastoreWithDefaults instantiates a new UpdateDatastores200ResponseDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateDatastores200ResponseDatastore) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateDatastores200ResponseDatastore) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateDatastores200ResponseDatastore) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateDatastores200ResponseDatastore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDatastores200ResponseDatastore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDatastores200ResponseDatastore) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *UpdateDatastores200ResponseDatastore) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateDatastores200ResponseDatastore) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateDatastores200ResponseDatastore) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateDatastores200ResponseDatastore) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *UpdateDatastores200ResponseDatastore) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *UpdateDatastores200ResponseDatastore) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDatastoreType

`func (o *UpdateDatastores200ResponseDatastore) GetDatastoreType() UpdateDatastores200ResponseDatastoreAllOfDatastoreType`

GetDatastoreType returns the DatastoreType field if non-nil, zero value otherwise.

### GetDatastoreTypeOk

`func (o *UpdateDatastores200ResponseDatastore) GetDatastoreTypeOk() (*UpdateDatastores200ResponseDatastoreAllOfDatastoreType, bool)`

GetDatastoreTypeOk returns a tuple with the DatastoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreType

`func (o *UpdateDatastores200ResponseDatastore) SetDatastoreType(v UpdateDatastores200ResponseDatastoreAllOfDatastoreType)`

SetDatastoreType sets DatastoreType field to given value.


### GetConfig

`func (o *UpdateDatastores200ResponseDatastore) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateDatastores200ResponseDatastore) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateDatastores200ResponseDatastore) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateDatastores200ResponseDatastore) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStorageServer

`func (o *UpdateDatastores200ResponseDatastore) GetStorageServer() UpdateDatastores200ResponseDatastoreAllOfStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *UpdateDatastores200ResponseDatastore) GetStorageServerOk() (*UpdateDatastores200ResponseDatastoreAllOfStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *UpdateDatastores200ResponseDatastore) SetStorageServer(v UpdateDatastores200ResponseDatastoreAllOfStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *UpdateDatastores200ResponseDatastore) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetType

`func (o *UpdateDatastores200ResponseDatastore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateDatastores200ResponseDatastore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateDatastores200ResponseDatastore) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *UpdateDatastores200ResponseDatastore) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateDatastores200ResponseDatastore) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateDatastores200ResponseDatastore) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *UpdateDatastores200ResponseDatastore) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateDatastores200ResponseDatastore) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateDatastores200ResponseDatastore) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateDatastores200ResponseDatastore) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateDatastores200ResponseDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateDatastores200ResponseDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateDatastores200ResponseDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateDatastores200ResponseDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStorageSize

`func (o *UpdateDatastores200ResponseDatastore) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *UpdateDatastores200ResponseDatastore) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *UpdateDatastores200ResponseDatastore) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *UpdateDatastores200ResponseDatastore) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### SetStorageSizeNil

`func (o *UpdateDatastores200ResponseDatastore) SetStorageSizeNil(b bool)`

 SetStorageSizeNil sets the value for StorageSize to be an explicit nil

### UnsetStorageSize
`func (o *UpdateDatastores200ResponseDatastore) UnsetStorageSize()`

UnsetStorageSize ensures that no value is present for StorageSize, not even an explicit nil
### GetFreeSpace

`func (o *UpdateDatastores200ResponseDatastore) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *UpdateDatastores200ResponseDatastore) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *UpdateDatastores200ResponseDatastore) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *UpdateDatastores200ResponseDatastore) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *UpdateDatastores200ResponseDatastore) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *UpdateDatastores200ResponseDatastore) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
### GetDrsEnabled

`func (o *UpdateDatastores200ResponseDatastore) GetDrsEnabled() bool`

GetDrsEnabled returns the DrsEnabled field if non-nil, zero value otherwise.

### GetDrsEnabledOk

`func (o *UpdateDatastores200ResponseDatastore) GetDrsEnabledOk() (*bool, bool)`

GetDrsEnabledOk returns a tuple with the DrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrsEnabled

`func (o *UpdateDatastores200ResponseDatastore) SetDrsEnabled(v bool)`

SetDrsEnabled sets DrsEnabled field to given value.

### HasDrsEnabled

`func (o *UpdateDatastores200ResponseDatastore) HasDrsEnabled() bool`

HasDrsEnabled returns a boolean if a field has been set.

### GetActive

`func (o *UpdateDatastores200ResponseDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateDatastores200ResponseDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateDatastores200ResponseDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateDatastores200ResponseDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowWrite

`func (o *UpdateDatastores200ResponseDatastore) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *UpdateDatastores200ResponseDatastore) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *UpdateDatastores200ResponseDatastore) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *UpdateDatastores200ResponseDatastore) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.

### GetDefaultStore

`func (o *UpdateDatastores200ResponseDatastore) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *UpdateDatastores200ResponseDatastore) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *UpdateDatastores200ResponseDatastore) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *UpdateDatastores200ResponseDatastore) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetOnline

`func (o *UpdateDatastores200ResponseDatastore) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *UpdateDatastores200ResponseDatastore) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *UpdateDatastores200ResponseDatastore) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *UpdateDatastores200ResponseDatastore) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetAllowRead

`func (o *UpdateDatastores200ResponseDatastore) GetAllowRead() bool`

GetAllowRead returns the AllowRead field if non-nil, zero value otherwise.

### GetAllowReadOk

`func (o *UpdateDatastores200ResponseDatastore) GetAllowReadOk() (*bool, bool)`

GetAllowReadOk returns a tuple with the AllowRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRead

`func (o *UpdateDatastores200ResponseDatastore) SetAllowRead(v bool)`

SetAllowRead sets AllowRead field to given value.

### HasAllowRead

`func (o *UpdateDatastores200ResponseDatastore) HasAllowRead() bool`

HasAllowRead returns a boolean if a field has been set.

### GetAllowProvision

`func (o *UpdateDatastores200ResponseDatastore) GetAllowProvision() bool`

GetAllowProvision returns the AllowProvision field if non-nil, zero value otherwise.

### GetAllowProvisionOk

`func (o *UpdateDatastores200ResponseDatastore) GetAllowProvisionOk() (*bool, bool)`

GetAllowProvisionOk returns a tuple with the AllowProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvision

`func (o *UpdateDatastores200ResponseDatastore) SetAllowProvision(v bool)`

SetAllowProvision sets AllowProvision field to given value.

### HasAllowProvision

`func (o *UpdateDatastores200ResponseDatastore) HasAllowProvision() bool`

HasAllowProvision returns a boolean if a field has been set.

### GetHeartBeatTarget

`func (o *UpdateDatastores200ResponseDatastore) GetHeartBeatTarget() bool`

GetHeartBeatTarget returns the HeartBeatTarget field if non-nil, zero value otherwise.

### GetHeartBeatTargetOk

`func (o *UpdateDatastores200ResponseDatastore) GetHeartBeatTargetOk() (*bool, bool)`

GetHeartBeatTargetOk returns a tuple with the HeartBeatTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartBeatTarget

`func (o *UpdateDatastores200ResponseDatastore) SetHeartBeatTarget(v bool)`

SetHeartBeatTarget sets HeartBeatTarget field to given value.

### HasHeartBeatTarget

`func (o *UpdateDatastores200ResponseDatastore) HasHeartBeatTarget() bool`

HasHeartBeatTarget returns a boolean if a field has been set.

### GetRefType

`func (o *UpdateDatastores200ResponseDatastore) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdateDatastores200ResponseDatastore) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdateDatastores200ResponseDatastore) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdateDatastores200ResponseDatastore) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *UpdateDatastores200ResponseDatastore) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateDatastores200ResponseDatastore) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateDatastores200ResponseDatastore) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateDatastores200ResponseDatastore) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateDatastores200ResponseDatastore) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateDatastores200ResponseDatastore) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateDatastores200ResponseDatastore) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateDatastores200ResponseDatastore) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalPath

`func (o *UpdateDatastores200ResponseDatastore) GetExternalPath() string`

GetExternalPath returns the ExternalPath field if non-nil, zero value otherwise.

### GetExternalPathOk

`func (o *UpdateDatastores200ResponseDatastore) GetExternalPathOk() (*string, bool)`

GetExternalPathOk returns a tuple with the ExternalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPath

`func (o *UpdateDatastores200ResponseDatastore) SetExternalPath(v string)`

SetExternalPath sets ExternalPath field to given value.

### HasExternalPath

`func (o *UpdateDatastores200ResponseDatastore) HasExternalPath() bool`

HasExternalPath returns a boolean if a field has been set.

### GetExternalType

`func (o *UpdateDatastores200ResponseDatastore) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *UpdateDatastores200ResponseDatastore) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *UpdateDatastores200ResponseDatastore) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *UpdateDatastores200ResponseDatastore) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetZone

`func (o *UpdateDatastores200ResponseDatastore) GetZone() UpdateDatastores200ResponseDatastoreAllOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateDatastores200ResponseDatastore) GetZoneOk() (*UpdateDatastores200ResponseDatastoreAllOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateDatastores200ResponseDatastore) SetZone(v UpdateDatastores200ResponseDatastoreAllOfZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateDatastores200ResponseDatastore) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *UpdateDatastores200ResponseDatastore) GetZonePool() UpdateDatastores200ResponseDatastoreAllOfZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *UpdateDatastores200ResponseDatastore) GetZonePoolOk() (*UpdateDatastores200ResponseDatastoreAllOfZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *UpdateDatastores200ResponseDatastore) SetZonePool(v UpdateDatastores200ResponseDatastoreAllOfZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *UpdateDatastores200ResponseDatastore) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateDatastores200ResponseDatastore) GetOwner() UpdateDatastores200ResponseDatastoreAllOfOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateDatastores200ResponseDatastore) GetOwnerOk() (*UpdateDatastores200ResponseDatastoreAllOfOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateDatastores200ResponseDatastore) SetOwner(v UpdateDatastores200ResponseDatastoreAllOfOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateDatastores200ResponseDatastore) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDatastores

`func (o *UpdateDatastores200ResponseDatastore) GetDatastores() []UpdateDatastores200ResponseDatastoreAllOfDatastoresInner`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *UpdateDatastores200ResponseDatastore) GetDatastoresOk() (*[]UpdateDatastores200ResponseDatastoreAllOfDatastoresInner, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *UpdateDatastores200ResponseDatastore) SetDatastores(v []UpdateDatastores200ResponseDatastoreAllOfDatastoresInner)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *UpdateDatastores200ResponseDatastore) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetLocations

`func (o *UpdateDatastores200ResponseDatastore) GetLocations() []UpdateDatastores200ResponseDatastoreAllOfLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *UpdateDatastores200ResponseDatastore) GetLocationsOk() (*[]UpdateDatastores200ResponseDatastoreAllOfLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *UpdateDatastores200ResponseDatastore) SetLocations(v []UpdateDatastores200ResponseDatastoreAllOfLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *UpdateDatastores200ResponseDatastore) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateDatastores200ResponseDatastore) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateDatastores200ResponseDatastore) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateDatastores200ResponseDatastore) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateDatastores200ResponseDatastore) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



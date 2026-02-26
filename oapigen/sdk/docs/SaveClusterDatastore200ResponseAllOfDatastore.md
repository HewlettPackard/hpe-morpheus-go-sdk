# SaveClusterDatastore200ResponseAllOfDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**DatastoreType** | Pointer to [**SaveClusterDatastore200ResponseAllOfDatastoreDatastoreType**](SaveClusterDatastore200ResponseAllOfDatastoreDatastoreType.md) |  | [optional] 
**StorageServer** | Pointer to [**SaveClusterDatastore200ResponseAllOfDatastoreStorageServer**](SaveClusterDatastore200ResponseAllOfDatastoreStorageServer.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**StorageSize** | Pointer to **NullableInt64** |  | [optional] 
**FreeSpace** | Pointer to **NullableInt64** |  | [optional] 
**DrsEnabled** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**AllowWrite** | Pointer to **bool** |  | [optional] 
**DefaultStore** | Pointer to **bool** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**AllowRead** | Pointer to **bool** |  | [optional] 
**AllowProvision** | Pointer to **bool** |  | [optional] 
**HeartbeatTarget** | Pointer to **bool** | Heartbeat Target | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**SaveClusterDatastore200ResponseAllOfDatastoreZone**](SaveClusterDatastore200ResponseAllOfDatastoreZone.md) |  | [optional] 
**ZonePool** | Pointer to [**SaveClusterDatastore200ResponseAllOfDatastoreZonePool**](SaveClusterDatastore200ResponseAllOfDatastoreZonePool.md) |  | [optional] 
**Owner** | Pointer to [**SaveClusterDatastore200ResponseAllOfDatastoreOwner**](SaveClusterDatastore200ResponseAllOfDatastoreOwner.md) |  | [optional] 
**Tenants** | Pointer to [**[]SaveClusterDatastore200ResponseAllOfDatastoreTenantsInner**](SaveClusterDatastore200ResponseAllOfDatastoreTenantsInner.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**SaveClusterDatastore200ResponseAllOfDatastoreResourcePermissions**](SaveClusterDatastore200ResponseAllOfDatastoreResourcePermissions.md) |  | [optional] 
**Datastores** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewSaveClusterDatastore200ResponseAllOfDatastore

`func NewSaveClusterDatastore200ResponseAllOfDatastore() *SaveClusterDatastore200ResponseAllOfDatastore`

NewSaveClusterDatastore200ResponseAllOfDatastore instantiates a new SaveClusterDatastore200ResponseAllOfDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveClusterDatastore200ResponseAllOfDatastoreWithDefaults

`func NewSaveClusterDatastore200ResponseAllOfDatastoreWithDefaults() *SaveClusterDatastore200ResponseAllOfDatastore`

NewSaveClusterDatastore200ResponseAllOfDatastoreWithDefaults instantiates a new SaveClusterDatastore200ResponseAllOfDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *SaveClusterDatastore200ResponseAllOfDatastore) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDatastoreType

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetDatastoreType() SaveClusterDatastore200ResponseAllOfDatastoreDatastoreType`

GetDatastoreType returns the DatastoreType field if non-nil, zero value otherwise.

### GetDatastoreTypeOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetDatastoreTypeOk() (*SaveClusterDatastore200ResponseAllOfDatastoreDatastoreType, bool)`

GetDatastoreTypeOk returns a tuple with the DatastoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreType

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetDatastoreType(v SaveClusterDatastore200ResponseAllOfDatastoreDatastoreType)`

SetDatastoreType sets DatastoreType field to given value.

### HasDatastoreType

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasDatastoreType() bool`

HasDatastoreType returns a boolean if a field has been set.

### GetStorageServer

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetStorageServer() SaveClusterDatastore200ResponseAllOfDatastoreStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetStorageServerOk() (*SaveClusterDatastore200ResponseAllOfDatastoreStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetStorageServer(v SaveClusterDatastore200ResponseAllOfDatastoreStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetType

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStorageSize

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### SetStorageSizeNil

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetStorageSizeNil(b bool)`

 SetStorageSizeNil sets the value for StorageSize to be an explicit nil

### UnsetStorageSize
`func (o *SaveClusterDatastore200ResponseAllOfDatastore) UnsetStorageSize()`

UnsetStorageSize ensures that no value is present for StorageSize, not even an explicit nil
### GetFreeSpace

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *SaveClusterDatastore200ResponseAllOfDatastore) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
### GetDrsEnabled

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetDrsEnabled() bool`

GetDrsEnabled returns the DrsEnabled field if non-nil, zero value otherwise.

### GetDrsEnabledOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetDrsEnabledOk() (*bool, bool)`

GetDrsEnabledOk returns a tuple with the DrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrsEnabled

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetDrsEnabled(v bool)`

SetDrsEnabled sets DrsEnabled field to given value.

### HasDrsEnabled

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasDrsEnabled() bool`

HasDrsEnabled returns a boolean if a field has been set.

### GetActive

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowWrite

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.

### GetDefaultStore

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetOnline

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetAllowRead

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetAllowRead() bool`

GetAllowRead returns the AllowRead field if non-nil, zero value otherwise.

### GetAllowReadOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetAllowReadOk() (*bool, bool)`

GetAllowReadOk returns a tuple with the AllowRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRead

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetAllowRead(v bool)`

SetAllowRead sets AllowRead field to given value.

### HasAllowRead

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasAllowRead() bool`

HasAllowRead returns a boolean if a field has been set.

### GetAllowProvision

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetAllowProvision() bool`

GetAllowProvision returns the AllowProvision field if non-nil, zero value otherwise.

### GetAllowProvisionOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetAllowProvisionOk() (*bool, bool)`

GetAllowProvisionOk returns a tuple with the AllowProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvision

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetAllowProvision(v bool)`

SetAllowProvision sets AllowProvision field to given value.

### HasAllowProvision

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasAllowProvision() bool`

HasAllowProvision returns a boolean if a field has been set.

### GetHeartbeatTarget

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetHeartbeatTarget() bool`

GetHeartbeatTarget returns the HeartbeatTarget field if non-nil, zero value otherwise.

### GetHeartbeatTargetOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetHeartbeatTargetOk() (*bool, bool)`

GetHeartbeatTargetOk returns a tuple with the HeartbeatTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatTarget

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetHeartbeatTarget(v bool)`

SetHeartbeatTarget sets HeartbeatTarget field to given value.

### HasHeartbeatTarget

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasHeartbeatTarget() bool`

HasHeartbeatTarget returns a boolean if a field has been set.

### GetRefType

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetExternalId

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetZone

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetZone() SaveClusterDatastore200ResponseAllOfDatastoreZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetZoneOk() (*SaveClusterDatastore200ResponseAllOfDatastoreZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetZone(v SaveClusterDatastore200ResponseAllOfDatastoreZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetZonePool() SaveClusterDatastore200ResponseAllOfDatastoreZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetZonePoolOk() (*SaveClusterDatastore200ResponseAllOfDatastoreZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetZonePool(v SaveClusterDatastore200ResponseAllOfDatastoreZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetOwner

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetOwner() SaveClusterDatastore200ResponseAllOfDatastoreOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetOwnerOk() (*SaveClusterDatastore200ResponseAllOfDatastoreOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetOwner(v SaveClusterDatastore200ResponseAllOfDatastoreOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenants

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetTenants() []SaveClusterDatastore200ResponseAllOfDatastoreTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetTenantsOk() (*[]SaveClusterDatastore200ResponseAllOfDatastoreTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetTenants(v []SaveClusterDatastore200ResponseAllOfDatastoreTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetResourcePermissions() SaveClusterDatastore200ResponseAllOfDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetResourcePermissionsOk() (*SaveClusterDatastore200ResponseAllOfDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetResourcePermissions(v SaveClusterDatastore200ResponseAllOfDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetDatastores

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetDatastores() []map[string]interface{}`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) GetDatastoresOk() (*[]map[string]interface{}, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) SetDatastores(v []map[string]interface{})`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *SaveClusterDatastore200ResponseAllOfDatastore) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



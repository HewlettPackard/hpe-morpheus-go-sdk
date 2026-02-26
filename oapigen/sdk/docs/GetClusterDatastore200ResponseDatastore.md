# GetClusterDatastore200ResponseDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**DatastoreType** | Pointer to [**GetClusterDatastore200ResponseDatastoreDatastoreType**](GetClusterDatastore200ResponseDatastoreDatastoreType.md) |  | [optional] 
**StorageServer** | Pointer to [**GetClusterDatastore200ResponseDatastoreStorageServer**](GetClusterDatastore200ResponseDatastoreStorageServer.md) |  | [optional] 
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
**Zone** | Pointer to [**GetClusterDatastore200ResponseDatastoreZone**](GetClusterDatastore200ResponseDatastoreZone.md) |  | [optional] 
**ZonePool** | Pointer to [**GetClusterDatastore200ResponseDatastoreZonePool**](GetClusterDatastore200ResponseDatastoreZonePool.md) |  | [optional] 
**Owner** | Pointer to [**GetClusterDatastore200ResponseDatastoreOwner**](GetClusterDatastore200ResponseDatastoreOwner.md) |  | [optional] 
**Tenants** | Pointer to [**[]GetClusterDatastore200ResponseDatastoreTenantsInner**](GetClusterDatastore200ResponseDatastoreTenantsInner.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**GetClusterDatastore200ResponseDatastoreResourcePermissions**](GetClusterDatastore200ResponseDatastoreResourcePermissions.md) |  | [optional] 
**Datastores** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetClusterDatastore200ResponseDatastore

`func NewGetClusterDatastore200ResponseDatastore() *GetClusterDatastore200ResponseDatastore`

NewGetClusterDatastore200ResponseDatastore instantiates a new GetClusterDatastore200ResponseDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterDatastore200ResponseDatastoreWithDefaults

`func NewGetClusterDatastore200ResponseDatastoreWithDefaults() *GetClusterDatastore200ResponseDatastore`

NewGetClusterDatastore200ResponseDatastoreWithDefaults instantiates a new GetClusterDatastore200ResponseDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClusterDatastore200ResponseDatastore) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterDatastore200ResponseDatastore) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterDatastore200ResponseDatastore) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterDatastore200ResponseDatastore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetClusterDatastore200ResponseDatastore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterDatastore200ResponseDatastore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterDatastore200ResponseDatastore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterDatastore200ResponseDatastore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetClusterDatastore200ResponseDatastore) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetClusterDatastore200ResponseDatastore) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetClusterDatastore200ResponseDatastore) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetClusterDatastore200ResponseDatastore) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetClusterDatastore200ResponseDatastore) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetClusterDatastore200ResponseDatastore) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDatastoreType

`func (o *GetClusterDatastore200ResponseDatastore) GetDatastoreType() GetClusterDatastore200ResponseDatastoreDatastoreType`

GetDatastoreType returns the DatastoreType field if non-nil, zero value otherwise.

### GetDatastoreTypeOk

`func (o *GetClusterDatastore200ResponseDatastore) GetDatastoreTypeOk() (*GetClusterDatastore200ResponseDatastoreDatastoreType, bool)`

GetDatastoreTypeOk returns a tuple with the DatastoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreType

`func (o *GetClusterDatastore200ResponseDatastore) SetDatastoreType(v GetClusterDatastore200ResponseDatastoreDatastoreType)`

SetDatastoreType sets DatastoreType field to given value.

### HasDatastoreType

`func (o *GetClusterDatastore200ResponseDatastore) HasDatastoreType() bool`

HasDatastoreType returns a boolean if a field has been set.

### GetStorageServer

`func (o *GetClusterDatastore200ResponseDatastore) GetStorageServer() GetClusterDatastore200ResponseDatastoreStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *GetClusterDatastore200ResponseDatastore) GetStorageServerOk() (*GetClusterDatastore200ResponseDatastoreStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *GetClusterDatastore200ResponseDatastore) SetStorageServer(v GetClusterDatastore200ResponseDatastoreStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *GetClusterDatastore200ResponseDatastore) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetType

`func (o *GetClusterDatastore200ResponseDatastore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetClusterDatastore200ResponseDatastore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetClusterDatastore200ResponseDatastore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetClusterDatastore200ResponseDatastore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *GetClusterDatastore200ResponseDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetClusterDatastore200ResponseDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetClusterDatastore200ResponseDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetClusterDatastore200ResponseDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStorageSize

`func (o *GetClusterDatastore200ResponseDatastore) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *GetClusterDatastore200ResponseDatastore) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *GetClusterDatastore200ResponseDatastore) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *GetClusterDatastore200ResponseDatastore) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### SetStorageSizeNil

`func (o *GetClusterDatastore200ResponseDatastore) SetStorageSizeNil(b bool)`

 SetStorageSizeNil sets the value for StorageSize to be an explicit nil

### UnsetStorageSize
`func (o *GetClusterDatastore200ResponseDatastore) UnsetStorageSize()`

UnsetStorageSize ensures that no value is present for StorageSize, not even an explicit nil
### GetFreeSpace

`func (o *GetClusterDatastore200ResponseDatastore) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *GetClusterDatastore200ResponseDatastore) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *GetClusterDatastore200ResponseDatastore) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *GetClusterDatastore200ResponseDatastore) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *GetClusterDatastore200ResponseDatastore) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *GetClusterDatastore200ResponseDatastore) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
### GetDrsEnabled

`func (o *GetClusterDatastore200ResponseDatastore) GetDrsEnabled() bool`

GetDrsEnabled returns the DrsEnabled field if non-nil, zero value otherwise.

### GetDrsEnabledOk

`func (o *GetClusterDatastore200ResponseDatastore) GetDrsEnabledOk() (*bool, bool)`

GetDrsEnabledOk returns a tuple with the DrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrsEnabled

`func (o *GetClusterDatastore200ResponseDatastore) SetDrsEnabled(v bool)`

SetDrsEnabled sets DrsEnabled field to given value.

### HasDrsEnabled

`func (o *GetClusterDatastore200ResponseDatastore) HasDrsEnabled() bool`

HasDrsEnabled returns a boolean if a field has been set.

### GetActive

`func (o *GetClusterDatastore200ResponseDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetClusterDatastore200ResponseDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetClusterDatastore200ResponseDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetClusterDatastore200ResponseDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowWrite

`func (o *GetClusterDatastore200ResponseDatastore) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *GetClusterDatastore200ResponseDatastore) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *GetClusterDatastore200ResponseDatastore) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *GetClusterDatastore200ResponseDatastore) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.

### GetDefaultStore

`func (o *GetClusterDatastore200ResponseDatastore) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *GetClusterDatastore200ResponseDatastore) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *GetClusterDatastore200ResponseDatastore) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *GetClusterDatastore200ResponseDatastore) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetOnline

`func (o *GetClusterDatastore200ResponseDatastore) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *GetClusterDatastore200ResponseDatastore) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *GetClusterDatastore200ResponseDatastore) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *GetClusterDatastore200ResponseDatastore) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetAllowRead

`func (o *GetClusterDatastore200ResponseDatastore) GetAllowRead() bool`

GetAllowRead returns the AllowRead field if non-nil, zero value otherwise.

### GetAllowReadOk

`func (o *GetClusterDatastore200ResponseDatastore) GetAllowReadOk() (*bool, bool)`

GetAllowReadOk returns a tuple with the AllowRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRead

`func (o *GetClusterDatastore200ResponseDatastore) SetAllowRead(v bool)`

SetAllowRead sets AllowRead field to given value.

### HasAllowRead

`func (o *GetClusterDatastore200ResponseDatastore) HasAllowRead() bool`

HasAllowRead returns a boolean if a field has been set.

### GetAllowProvision

`func (o *GetClusterDatastore200ResponseDatastore) GetAllowProvision() bool`

GetAllowProvision returns the AllowProvision field if non-nil, zero value otherwise.

### GetAllowProvisionOk

`func (o *GetClusterDatastore200ResponseDatastore) GetAllowProvisionOk() (*bool, bool)`

GetAllowProvisionOk returns a tuple with the AllowProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvision

`func (o *GetClusterDatastore200ResponseDatastore) SetAllowProvision(v bool)`

SetAllowProvision sets AllowProvision field to given value.

### HasAllowProvision

`func (o *GetClusterDatastore200ResponseDatastore) HasAllowProvision() bool`

HasAllowProvision returns a boolean if a field has been set.

### GetHeartbeatTarget

`func (o *GetClusterDatastore200ResponseDatastore) GetHeartbeatTarget() bool`

GetHeartbeatTarget returns the HeartbeatTarget field if non-nil, zero value otherwise.

### GetHeartbeatTargetOk

`func (o *GetClusterDatastore200ResponseDatastore) GetHeartbeatTargetOk() (*bool, bool)`

GetHeartbeatTargetOk returns a tuple with the HeartbeatTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatTarget

`func (o *GetClusterDatastore200ResponseDatastore) SetHeartbeatTarget(v bool)`

SetHeartbeatTarget sets HeartbeatTarget field to given value.

### HasHeartbeatTarget

`func (o *GetClusterDatastore200ResponseDatastore) HasHeartbeatTarget() bool`

HasHeartbeatTarget returns a boolean if a field has been set.

### GetRefType

`func (o *GetClusterDatastore200ResponseDatastore) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetClusterDatastore200ResponseDatastore) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetClusterDatastore200ResponseDatastore) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetClusterDatastore200ResponseDatastore) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetClusterDatastore200ResponseDatastore) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetClusterDatastore200ResponseDatastore) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetClusterDatastore200ResponseDatastore) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetClusterDatastore200ResponseDatastore) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetExternalId

`func (o *GetClusterDatastore200ResponseDatastore) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetClusterDatastore200ResponseDatastore) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetClusterDatastore200ResponseDatastore) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetClusterDatastore200ResponseDatastore) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetZone

`func (o *GetClusterDatastore200ResponseDatastore) GetZone() GetClusterDatastore200ResponseDatastoreZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetClusterDatastore200ResponseDatastore) GetZoneOk() (*GetClusterDatastore200ResponseDatastoreZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetClusterDatastore200ResponseDatastore) SetZone(v GetClusterDatastore200ResponseDatastoreZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetClusterDatastore200ResponseDatastore) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *GetClusterDatastore200ResponseDatastore) GetZonePool() GetClusterDatastore200ResponseDatastoreZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *GetClusterDatastore200ResponseDatastore) GetZonePoolOk() (*GetClusterDatastore200ResponseDatastoreZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *GetClusterDatastore200ResponseDatastore) SetZonePool(v GetClusterDatastore200ResponseDatastoreZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *GetClusterDatastore200ResponseDatastore) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetOwner

`func (o *GetClusterDatastore200ResponseDatastore) GetOwner() GetClusterDatastore200ResponseDatastoreOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetClusterDatastore200ResponseDatastore) GetOwnerOk() (*GetClusterDatastore200ResponseDatastoreOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetClusterDatastore200ResponseDatastore) SetOwner(v GetClusterDatastore200ResponseDatastoreOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetClusterDatastore200ResponseDatastore) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenants

`func (o *GetClusterDatastore200ResponseDatastore) GetTenants() []GetClusterDatastore200ResponseDatastoreTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetClusterDatastore200ResponseDatastore) GetTenantsOk() (*[]GetClusterDatastore200ResponseDatastoreTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetClusterDatastore200ResponseDatastore) SetTenants(v []GetClusterDatastore200ResponseDatastoreTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetClusterDatastore200ResponseDatastore) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *GetClusterDatastore200ResponseDatastore) GetResourcePermissions() GetClusterDatastore200ResponseDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *GetClusterDatastore200ResponseDatastore) GetResourcePermissionsOk() (*GetClusterDatastore200ResponseDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *GetClusterDatastore200ResponseDatastore) SetResourcePermissions(v GetClusterDatastore200ResponseDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *GetClusterDatastore200ResponseDatastore) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetDatastores

`func (o *GetClusterDatastore200ResponseDatastore) GetDatastores() []map[string]interface{}`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *GetClusterDatastore200ResponseDatastore) GetDatastoresOk() (*[]map[string]interface{}, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *GetClusterDatastore200ResponseDatastore) SetDatastores(v []map[string]interface{})`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *GetClusterDatastore200ResponseDatastore) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



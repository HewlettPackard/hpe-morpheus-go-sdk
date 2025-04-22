# SaveClusterDatastore200ResponseAllOfCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**DatastoreType** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**StorageServer** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**StorageSize** | Pointer to **int64** |  | [optional] 
**FreeSpace** | Pointer to **int64** |  | [optional] 
**DrsEnabled** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**AllowWrite** | Pointer to **bool** |  | [optional] 
**DefaultStore** | Pointer to **bool** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**AllowRead** | Pointer to **bool** |  | [optional] 
**AllowProvision** | Pointer to **bool** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**ZonePool** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Tenants** | Pointer to [**[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner**](ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**SaveClusterDatastoreRequestClusterResourcePermissions**](SaveClusterDatastoreRequestClusterResourcePermissions.md) |  | [optional] 
**Datastores** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewSaveClusterDatastore200ResponseAllOfCluster

`func NewSaveClusterDatastore200ResponseAllOfCluster() *SaveClusterDatastore200ResponseAllOfCluster`

NewSaveClusterDatastore200ResponseAllOfCluster instantiates a new SaveClusterDatastore200ResponseAllOfCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveClusterDatastore200ResponseAllOfClusterWithDefaults

`func NewSaveClusterDatastore200ResponseAllOfClusterWithDefaults() *SaveClusterDatastore200ResponseAllOfCluster`

NewSaveClusterDatastore200ResponseAllOfClusterWithDefaults instantiates a new SaveClusterDatastore200ResponseAllOfCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDatastoreType

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetDatastoreType() GetAlerts200ResponseAllOfChecksInnerAccount`

GetDatastoreType returns the DatastoreType field if non-nil, zero value otherwise.

### GetDatastoreTypeOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetDatastoreTypeOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetDatastoreTypeOk returns a tuple with the DatastoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreType

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetDatastoreType(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetDatastoreType sets DatastoreType field to given value.

### HasDatastoreType

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasDatastoreType() bool`

HasDatastoreType returns a boolean if a field has been set.

### GetStorageServer

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetStorageServer() GetAlerts200ResponseAllOfChecksInnerAccount`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetStorageServerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetStorageServer(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetType

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStorageSize

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetFreeSpace

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### GetDrsEnabled

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetDrsEnabled() bool`

GetDrsEnabled returns the DrsEnabled field if non-nil, zero value otherwise.

### GetDrsEnabledOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetDrsEnabledOk() (*bool, bool)`

GetDrsEnabledOk returns a tuple with the DrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrsEnabled

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetDrsEnabled(v bool)`

SetDrsEnabled sets DrsEnabled field to given value.

### HasDrsEnabled

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasDrsEnabled() bool`

HasDrsEnabled returns a boolean if a field has been set.

### GetActive

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllowWrite

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.

### GetDefaultStore

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetOnline

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetAllowRead

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetAllowRead() bool`

GetAllowRead returns the AllowRead field if non-nil, zero value otherwise.

### GetAllowReadOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetAllowReadOk() (*bool, bool)`

GetAllowReadOk returns a tuple with the AllowRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRead

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetAllowRead(v bool)`

SetAllowRead sets AllowRead field to given value.

### HasAllowRead

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasAllowRead() bool`

HasAllowRead returns a boolean if a field has been set.

### GetAllowProvision

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetAllowProvision() bool`

GetAllowProvision returns the AllowProvision field if non-nil, zero value otherwise.

### GetAllowProvisionOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetAllowProvisionOk() (*bool, bool)`

GetAllowProvisionOk returns a tuple with the AllowProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvision

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetAllowProvision(v bool)`

SetAllowProvision sets AllowProvision field to given value.

### HasAllowProvision

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasAllowProvision() bool`

HasAllowProvision returns a boolean if a field has been set.

### GetRefType

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetExternalId

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetZone

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetZone() GetAlerts200ResponseAllOfChecksInnerAccount`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetZoneOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetZone(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetZone sets Zone field to given value.

### HasZone

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetZonePool() GetAlerts200ResponseAllOfChecksInnerAccount`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetZonePoolOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetZonePool(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetOwner

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetOwner() GetAlerts200ResponseAllOfChecksInnerAccount`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetOwner(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenants

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetTenants() []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetTenantsOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetTenants(v []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetResourcePermissions() SaveClusterDatastoreRequestClusterResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetResourcePermissionsOk() (*SaveClusterDatastoreRequestClusterResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetResourcePermissions(v SaveClusterDatastoreRequestClusterResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetDatastores

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetDatastores() []map[string]interface{}`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *SaveClusterDatastore200ResponseAllOfCluster) GetDatastoresOk() (*[]map[string]interface{}, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *SaveClusterDatastore200ResponseAllOfCluster) SetDatastores(v []map[string]interface{})`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *SaveClusterDatastore200ResponseAllOfCluster) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



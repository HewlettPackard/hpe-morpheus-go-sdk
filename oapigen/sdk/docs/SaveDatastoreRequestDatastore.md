# SaveDatastoreRequestDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the datastore to be created. | 
**DatastoreType** | **string** | The code of the datatoreType | 
**Config** | [**SaveDatastoreRequestDatastoreConfig**](SaveDatastoreRequestDatastoreConfig.md) |  | 
**RefType** | **string** | Type of the resource this datastore is associated with, can be &#39;ComputeZone&#39; (&#39;Cloud&#39;) or &#39;ComputeServerGroup&#39; (&#39;Cluster&#39;) | 
**RefId** | **int64** | The ID of the resource this datastore is associated with, e.g. ComputeZone, ComputeServerGroup | 
**StorageServer** | Pointer to [**SaveDatastoreRequestDatastoreStorageServer**](SaveDatastoreRequestDatastoreStorageServer.md) |  | [optional] 
**Visibility** | Pointer to **string** | Visibility level of the datastore, can be &#39;private&#39; or &#39;public&#39;. If not specified, defaults to &#39;private&#39;. | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DefaultStore** | Pointer to **bool** |  | [optional] 
**TenantPermissions** | Pointer to [**SaveDatastoreRequestDatastoreTenantPermissions**](SaveDatastoreRequestDatastoreTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**SaveDatastoreRequestDatastoreResourcePermissions**](SaveDatastoreRequestDatastoreResourcePermissions.md) |  | [optional] 
**Datastores** | Pointer to **[]map[string]interface{}** | List of datastores associated with this datastore, for use with vSphere clouds. | [optional] 

## Methods

### NewSaveDatastoreRequestDatastore

`func NewSaveDatastoreRequestDatastore(name string, datastoreType string, config SaveDatastoreRequestDatastoreConfig, refType string, refId int64, ) *SaveDatastoreRequestDatastore`

NewSaveDatastoreRequestDatastore instantiates a new SaveDatastoreRequestDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveDatastoreRequestDatastoreWithDefaults

`func NewSaveDatastoreRequestDatastoreWithDefaults() *SaveDatastoreRequestDatastore`

NewSaveDatastoreRequestDatastoreWithDefaults instantiates a new SaveDatastoreRequestDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SaveDatastoreRequestDatastore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaveDatastoreRequestDatastore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaveDatastoreRequestDatastore) SetName(v string)`

SetName sets Name field to given value.


### GetDatastoreType

`func (o *SaveDatastoreRequestDatastore) GetDatastoreType() string`

GetDatastoreType returns the DatastoreType field if non-nil, zero value otherwise.

### GetDatastoreTypeOk

`func (o *SaveDatastoreRequestDatastore) GetDatastoreTypeOk() (*string, bool)`

GetDatastoreTypeOk returns a tuple with the DatastoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreType

`func (o *SaveDatastoreRequestDatastore) SetDatastoreType(v string)`

SetDatastoreType sets DatastoreType field to given value.


### GetConfig

`func (o *SaveDatastoreRequestDatastore) GetConfig() SaveDatastoreRequestDatastoreConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SaveDatastoreRequestDatastore) GetConfigOk() (*SaveDatastoreRequestDatastoreConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SaveDatastoreRequestDatastore) SetConfig(v SaveDatastoreRequestDatastoreConfig)`

SetConfig sets Config field to given value.


### GetRefType

`func (o *SaveDatastoreRequestDatastore) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *SaveDatastoreRequestDatastore) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *SaveDatastoreRequestDatastore) SetRefType(v string)`

SetRefType sets RefType field to given value.


### GetRefId

`func (o *SaveDatastoreRequestDatastore) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *SaveDatastoreRequestDatastore) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *SaveDatastoreRequestDatastore) SetRefId(v int64)`

SetRefId sets RefId field to given value.


### GetStorageServer

`func (o *SaveDatastoreRequestDatastore) GetStorageServer() SaveDatastoreRequestDatastoreStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *SaveDatastoreRequestDatastore) GetStorageServerOk() (*SaveDatastoreRequestDatastoreStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *SaveDatastoreRequestDatastore) SetStorageServer(v SaveDatastoreRequestDatastoreStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *SaveDatastoreRequestDatastore) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetVisibility

`func (o *SaveDatastoreRequestDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SaveDatastoreRequestDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SaveDatastoreRequestDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SaveDatastoreRequestDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *SaveDatastoreRequestDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SaveDatastoreRequestDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SaveDatastoreRequestDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SaveDatastoreRequestDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefaultStore

`func (o *SaveDatastoreRequestDatastore) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *SaveDatastoreRequestDatastore) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *SaveDatastoreRequestDatastore) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *SaveDatastoreRequestDatastore) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *SaveDatastoreRequestDatastore) GetTenantPermissions() SaveDatastoreRequestDatastoreTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *SaveDatastoreRequestDatastore) GetTenantPermissionsOk() (*SaveDatastoreRequestDatastoreTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *SaveDatastoreRequestDatastore) SetTenantPermissions(v SaveDatastoreRequestDatastoreTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *SaveDatastoreRequestDatastore) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *SaveDatastoreRequestDatastore) GetResourcePermissions() SaveDatastoreRequestDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *SaveDatastoreRequestDatastore) GetResourcePermissionsOk() (*SaveDatastoreRequestDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *SaveDatastoreRequestDatastore) SetResourcePermissions(v SaveDatastoreRequestDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *SaveDatastoreRequestDatastore) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetDatastores

`func (o *SaveDatastoreRequestDatastore) GetDatastores() []map[string]interface{}`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *SaveDatastoreRequestDatastore) GetDatastoresOk() (*[]map[string]interface{}, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *SaveDatastoreRequestDatastore) SetDatastores(v []map[string]interface{})`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *SaveDatastoreRequestDatastore) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



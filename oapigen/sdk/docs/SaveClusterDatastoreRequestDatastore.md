# SaveClusterDatastoreRequestDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DatastoreType** | Pointer to [**SaveClusterDatastoreRequestDatastoreDatastoreType**](SaveClusterDatastoreRequestDatastoreDatastoreType.md) |  | [optional] 
**StorageServer** | Pointer to [**SaveClusterDatastoreRequestDatastoreStorageServer**](SaveClusterDatastoreRequestDatastoreStorageServer.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DefaultStore** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**SaveClusterDatastoreRequestDatastoreConfig**](SaveClusterDatastoreRequestDatastoreConfig.md) |  | [optional] 
**Tenants** | Pointer to [**[]SaveClusterDatastoreRequestDatastoreTenantsInner**](SaveClusterDatastoreRequestDatastoreTenantsInner.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**SaveClusterDatastoreRequestDatastoreResourcePermissions**](SaveClusterDatastoreRequestDatastoreResourcePermissions.md) |  | [optional] 
**Datastores** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewSaveClusterDatastoreRequestDatastore

`func NewSaveClusterDatastoreRequestDatastore() *SaveClusterDatastoreRequestDatastore`

NewSaveClusterDatastoreRequestDatastore instantiates a new SaveClusterDatastoreRequestDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveClusterDatastoreRequestDatastoreWithDefaults

`func NewSaveClusterDatastoreRequestDatastoreWithDefaults() *SaveClusterDatastoreRequestDatastore`

NewSaveClusterDatastoreRequestDatastoreWithDefaults instantiates a new SaveClusterDatastoreRequestDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SaveClusterDatastoreRequestDatastore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaveClusterDatastoreRequestDatastore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaveClusterDatastoreRequestDatastore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaveClusterDatastoreRequestDatastore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDatastoreType

`func (o *SaveClusterDatastoreRequestDatastore) GetDatastoreType() SaveClusterDatastoreRequestDatastoreDatastoreType`

GetDatastoreType returns the DatastoreType field if non-nil, zero value otherwise.

### GetDatastoreTypeOk

`func (o *SaveClusterDatastoreRequestDatastore) GetDatastoreTypeOk() (*SaveClusterDatastoreRequestDatastoreDatastoreType, bool)`

GetDatastoreTypeOk returns a tuple with the DatastoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreType

`func (o *SaveClusterDatastoreRequestDatastore) SetDatastoreType(v SaveClusterDatastoreRequestDatastoreDatastoreType)`

SetDatastoreType sets DatastoreType field to given value.

### HasDatastoreType

`func (o *SaveClusterDatastoreRequestDatastore) HasDatastoreType() bool`

HasDatastoreType returns a boolean if a field has been set.

### GetStorageServer

`func (o *SaveClusterDatastoreRequestDatastore) GetStorageServer() SaveClusterDatastoreRequestDatastoreStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *SaveClusterDatastoreRequestDatastore) GetStorageServerOk() (*SaveClusterDatastoreRequestDatastoreStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *SaveClusterDatastoreRequestDatastore) SetStorageServer(v SaveClusterDatastoreRequestDatastoreStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *SaveClusterDatastoreRequestDatastore) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetVisibility

`func (o *SaveClusterDatastoreRequestDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SaveClusterDatastoreRequestDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SaveClusterDatastoreRequestDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SaveClusterDatastoreRequestDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *SaveClusterDatastoreRequestDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SaveClusterDatastoreRequestDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SaveClusterDatastoreRequestDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SaveClusterDatastoreRequestDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefaultStore

`func (o *SaveClusterDatastoreRequestDatastore) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *SaveClusterDatastoreRequestDatastore) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *SaveClusterDatastoreRequestDatastore) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *SaveClusterDatastoreRequestDatastore) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetConfig

`func (o *SaveClusterDatastoreRequestDatastore) GetConfig() SaveClusterDatastoreRequestDatastoreConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SaveClusterDatastoreRequestDatastore) GetConfigOk() (*SaveClusterDatastoreRequestDatastoreConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SaveClusterDatastoreRequestDatastore) SetConfig(v SaveClusterDatastoreRequestDatastoreConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SaveClusterDatastoreRequestDatastore) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTenants

`func (o *SaveClusterDatastoreRequestDatastore) GetTenants() []SaveClusterDatastoreRequestDatastoreTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *SaveClusterDatastoreRequestDatastore) GetTenantsOk() (*[]SaveClusterDatastoreRequestDatastoreTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *SaveClusterDatastoreRequestDatastore) SetTenants(v []SaveClusterDatastoreRequestDatastoreTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *SaveClusterDatastoreRequestDatastore) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *SaveClusterDatastoreRequestDatastore) GetResourcePermissions() SaveClusterDatastoreRequestDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *SaveClusterDatastoreRequestDatastore) GetResourcePermissionsOk() (*SaveClusterDatastoreRequestDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *SaveClusterDatastoreRequestDatastore) SetResourcePermissions(v SaveClusterDatastoreRequestDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *SaveClusterDatastoreRequestDatastore) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetDatastores

`func (o *SaveClusterDatastoreRequestDatastore) GetDatastores() []map[string]interface{}`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *SaveClusterDatastoreRequestDatastore) GetDatastoresOk() (*[]map[string]interface{}, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *SaveClusterDatastoreRequestDatastore) SetDatastores(v []map[string]interface{})`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *SaveClusterDatastoreRequestDatastore) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



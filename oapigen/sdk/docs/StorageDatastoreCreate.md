# StorageDatastoreCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the datastore to be created. | 
**DatastoreType** | **string** | The code of the datatoreType | 
**Config** | [**StorageDatastoreCreateConfig**](StorageDatastoreCreateConfig.md) |  | 
**RefType** | **string** | Type of the resource this datastore is associated with, can be &#39;ComputeZone&#39; (&#39;Cloud&#39;) or &#39;ComputeServerGroup&#39; (&#39;Cluster&#39;) | 
**RefId** | **int64** | The ID of the resource this datastore is associated with, e.g. ComputeZone, ComputeServerGroup | 
**StorageServer** | Pointer to [**StorageDatastoreCreateStorageServer**](StorageDatastoreCreateStorageServer.md) |  | [optional] 
**Visibility** | Pointer to **string** | Visibility level of the datastore, can be &#39;private&#39; or &#39;public&#39;. If not specified, defaults to &#39;private&#39;. | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DefaultStore** | Pointer to **bool** |  | [optional] 
**TenantPermissions** | Pointer to [**StorageDatastoreCreateTenantPermissions**](StorageDatastoreCreateTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**StorageDatastoreCreateResourcePermissions**](StorageDatastoreCreateResourcePermissions.md) |  | [optional] 
**Datastores** | Pointer to **[]map[string]interface{}** | List of datastores associated with this datastore, for use with vSphere clouds. | [optional] 

## Methods

### NewStorageDatastoreCreate

`func NewStorageDatastoreCreate(name string, datastoreType string, config StorageDatastoreCreateConfig, refType string, refId int64, ) *StorageDatastoreCreate`

NewStorageDatastoreCreate instantiates a new StorageDatastoreCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDatastoreCreateWithDefaults

`func NewStorageDatastoreCreateWithDefaults() *StorageDatastoreCreate`

NewStorageDatastoreCreateWithDefaults instantiates a new StorageDatastoreCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageDatastoreCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageDatastoreCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageDatastoreCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDatastoreType

`func (o *StorageDatastoreCreate) GetDatastoreType() string`

GetDatastoreType returns the DatastoreType field if non-nil, zero value otherwise.

### GetDatastoreTypeOk

`func (o *StorageDatastoreCreate) GetDatastoreTypeOk() (*string, bool)`

GetDatastoreTypeOk returns a tuple with the DatastoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreType

`func (o *StorageDatastoreCreate) SetDatastoreType(v string)`

SetDatastoreType sets DatastoreType field to given value.


### GetConfig

`func (o *StorageDatastoreCreate) GetConfig() StorageDatastoreCreateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *StorageDatastoreCreate) GetConfigOk() (*StorageDatastoreCreateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *StorageDatastoreCreate) SetConfig(v StorageDatastoreCreateConfig)`

SetConfig sets Config field to given value.


### GetRefType

`func (o *StorageDatastoreCreate) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *StorageDatastoreCreate) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *StorageDatastoreCreate) SetRefType(v string)`

SetRefType sets RefType field to given value.


### GetRefId

`func (o *StorageDatastoreCreate) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *StorageDatastoreCreate) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *StorageDatastoreCreate) SetRefId(v int64)`

SetRefId sets RefId field to given value.


### GetStorageServer

`func (o *StorageDatastoreCreate) GetStorageServer() StorageDatastoreCreateStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *StorageDatastoreCreate) GetStorageServerOk() (*StorageDatastoreCreateStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *StorageDatastoreCreate) SetStorageServer(v StorageDatastoreCreateStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *StorageDatastoreCreate) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetVisibility

`func (o *StorageDatastoreCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *StorageDatastoreCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *StorageDatastoreCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *StorageDatastoreCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *StorageDatastoreCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StorageDatastoreCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StorageDatastoreCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *StorageDatastoreCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefaultStore

`func (o *StorageDatastoreCreate) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *StorageDatastoreCreate) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *StorageDatastoreCreate) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *StorageDatastoreCreate) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *StorageDatastoreCreate) GetTenantPermissions() StorageDatastoreCreateTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *StorageDatastoreCreate) GetTenantPermissionsOk() (*StorageDatastoreCreateTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *StorageDatastoreCreate) SetTenantPermissions(v StorageDatastoreCreateTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *StorageDatastoreCreate) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *StorageDatastoreCreate) GetResourcePermissions() StorageDatastoreCreateResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *StorageDatastoreCreate) GetResourcePermissionsOk() (*StorageDatastoreCreateResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *StorageDatastoreCreate) SetResourcePermissions(v StorageDatastoreCreateResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *StorageDatastoreCreate) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetDatastores

`func (o *StorageDatastoreCreate) GetDatastores() []map[string]interface{}`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *StorageDatastoreCreate) GetDatastoresOk() (*[]map[string]interface{}, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *StorageDatastoreCreate) SetDatastores(v []map[string]interface{})`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *StorageDatastoreCreate) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ZoneDatastoreCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DatastoreType** | Pointer to **string** | The code of the datatoreType | [optional] 
**StorageServer** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DefaultStore** | Pointer to **bool** |  | [optional] 
**Tenants** | Pointer to [**[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner**](ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**SaveCloudDatastoreRequestDatastoreResourcePermissions**](SaveCloudDatastoreRequestDatastoreResourcePermissions.md) |  | [optional] 
**Datastores** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewZoneDatastoreCreate

`func NewZoneDatastoreCreate() *ZoneDatastoreCreate`

NewZoneDatastoreCreate instantiates a new ZoneDatastoreCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneDatastoreCreateWithDefaults

`func NewZoneDatastoreCreateWithDefaults() *ZoneDatastoreCreate`

NewZoneDatastoreCreateWithDefaults instantiates a new ZoneDatastoreCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ZoneDatastoreCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneDatastoreCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneDatastoreCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneDatastoreCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDatastoreType

`func (o *ZoneDatastoreCreate) GetDatastoreType() string`

GetDatastoreType returns the DatastoreType field if non-nil, zero value otherwise.

### GetDatastoreTypeOk

`func (o *ZoneDatastoreCreate) GetDatastoreTypeOk() (*string, bool)`

GetDatastoreTypeOk returns a tuple with the DatastoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreType

`func (o *ZoneDatastoreCreate) SetDatastoreType(v string)`

SetDatastoreType sets DatastoreType field to given value.

### HasDatastoreType

`func (o *ZoneDatastoreCreate) HasDatastoreType() bool`

HasDatastoreType returns a boolean if a field has been set.

### GetStorageServer

`func (o *ZoneDatastoreCreate) GetStorageServer() GetAlerts200ResponseAllOfChecksInnerAccount`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *ZoneDatastoreCreate) GetStorageServerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *ZoneDatastoreCreate) SetStorageServer(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *ZoneDatastoreCreate) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetVisibility

`func (o *ZoneDatastoreCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ZoneDatastoreCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ZoneDatastoreCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ZoneDatastoreCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *ZoneDatastoreCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ZoneDatastoreCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ZoneDatastoreCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ZoneDatastoreCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefaultStore

`func (o *ZoneDatastoreCreate) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ZoneDatastoreCreate) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ZoneDatastoreCreate) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ZoneDatastoreCreate) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetTenants

`func (o *ZoneDatastoreCreate) GetTenants() []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ZoneDatastoreCreate) GetTenantsOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ZoneDatastoreCreate) SetTenants(v []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ZoneDatastoreCreate) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ZoneDatastoreCreate) GetResourcePermissions() SaveCloudDatastoreRequestDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ZoneDatastoreCreate) GetResourcePermissionsOk() (*SaveCloudDatastoreRequestDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ZoneDatastoreCreate) SetResourcePermissions(v SaveCloudDatastoreRequestDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ZoneDatastoreCreate) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetDatastores

`func (o *ZoneDatastoreCreate) GetDatastores() []map[string]interface{}`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *ZoneDatastoreCreate) GetDatastoresOk() (*[]map[string]interface{}, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *ZoneDatastoreCreate) SetDatastores(v []map[string]interface{})`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *ZoneDatastoreCreate) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



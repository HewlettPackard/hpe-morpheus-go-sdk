# ZoneDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**FreeSpace** | Pointer to **int64** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Tenants** | Pointer to [**[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner**](ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission**](ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission.md) |  | [optional] 

## Methods

### NewZoneDatastore

`func NewZoneDatastore() *ZoneDatastore`

NewZoneDatastore instantiates a new ZoneDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneDatastoreWithDefaults

`func NewZoneDatastoreWithDefaults() *ZoneDatastore`

NewZoneDatastoreWithDefaults instantiates a new ZoneDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneDatastore) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneDatastore) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneDatastore) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneDatastore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ZoneDatastore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneDatastore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneDatastore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneDatastore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZone

`func (o *ZoneDatastore) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ZoneDatastore) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ZoneDatastore) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ZoneDatastore) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetType

`func (o *ZoneDatastore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneDatastore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneDatastore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ZoneDatastore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFreeSpace

`func (o *ZoneDatastore) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *ZoneDatastore) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *ZoneDatastore) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *ZoneDatastore) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### GetOnline

`func (o *ZoneDatastore) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ZoneDatastore) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ZoneDatastore) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ZoneDatastore) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetActive

`func (o *ZoneDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ZoneDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ZoneDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ZoneDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVisibility

`func (o *ZoneDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ZoneDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ZoneDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ZoneDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *ZoneDatastore) GetTenants() []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ZoneDatastore) GetTenantsOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ZoneDatastore) SetTenants(v []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ZoneDatastore) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *ZoneDatastore) GetResourcePermission() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *ZoneDatastore) GetResourcePermissionOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *ZoneDatastore) SetResourcePermission(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *ZoneDatastore) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



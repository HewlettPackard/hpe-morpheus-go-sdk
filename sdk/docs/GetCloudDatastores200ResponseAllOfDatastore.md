# GetCloudDatastores200ResponseAllOfDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**FreeSpace** | Pointer to **int64** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Tenants** | Pointer to [**[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner**](ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission**](ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission.md) |  | [optional] 

## Methods

### NewGetCloudDatastores200ResponseAllOfDatastore

`func NewGetCloudDatastores200ResponseAllOfDatastore() *GetCloudDatastores200ResponseAllOfDatastore`

NewGetCloudDatastores200ResponseAllOfDatastore instantiates a new GetCloudDatastores200ResponseAllOfDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCloudDatastores200ResponseAllOfDatastoreWithDefaults

`func NewGetCloudDatastores200ResponseAllOfDatastoreWithDefaults() *GetCloudDatastores200ResponseAllOfDatastore`

NewGetCloudDatastores200ResponseAllOfDatastoreWithDefaults instantiates a new GetCloudDatastores200ResponseAllOfDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCloudDatastores200ResponseAllOfDatastore) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCloudDatastores200ResponseAllOfDatastore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCloudDatastores200ResponseAllOfDatastore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCloudDatastores200ResponseAllOfDatastore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZone

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetZone() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetZoneOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetCloudDatastores200ResponseAllOfDatastore) SetZone(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetCloudDatastores200ResponseAllOfDatastore) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetType

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCloudDatastores200ResponseAllOfDatastore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCloudDatastores200ResponseAllOfDatastore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFreeSpace

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *GetCloudDatastores200ResponseAllOfDatastore) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *GetCloudDatastores200ResponseAllOfDatastore) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### GetOnline

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *GetCloudDatastores200ResponseAllOfDatastore) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *GetCloudDatastores200ResponseAllOfDatastore) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetActive

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetCloudDatastores200ResponseAllOfDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetCloudDatastores200ResponseAllOfDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVisibility

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetCloudDatastores200ResponseAllOfDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetCloudDatastores200ResponseAllOfDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetTenants() []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetTenantsOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetCloudDatastores200ResponseAllOfDatastore) SetTenants(v []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetCloudDatastores200ResponseAllOfDatastore) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *GetCloudDatastores200ResponseAllOfDatastore) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *GetCloudDatastores200ResponseAllOfDatastore) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetResourcePermission

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetResourcePermission() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *GetCloudDatastores200ResponseAllOfDatastore) GetResourcePermissionOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *GetCloudDatastores200ResponseAllOfDatastore) SetResourcePermission(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *GetCloudDatastores200ResponseAllOfDatastore) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



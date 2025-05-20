# UpdateCloudDatastores200ResponseDatastore

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
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateCloudDatastores200ResponseDatastore

`func NewUpdateCloudDatastores200ResponseDatastore() *UpdateCloudDatastores200ResponseDatastore`

NewUpdateCloudDatastores200ResponseDatastore instantiates a new UpdateCloudDatastores200ResponseDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCloudDatastores200ResponseDatastoreWithDefaults

`func NewUpdateCloudDatastores200ResponseDatastoreWithDefaults() *UpdateCloudDatastores200ResponseDatastore`

NewUpdateCloudDatastores200ResponseDatastoreWithDefaults instantiates a new UpdateCloudDatastores200ResponseDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateCloudDatastores200ResponseDatastore) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCloudDatastores200ResponseDatastore) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCloudDatastores200ResponseDatastore) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCloudDatastores200ResponseDatastore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateCloudDatastores200ResponseDatastore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCloudDatastores200ResponseDatastore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCloudDatastores200ResponseDatastore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCloudDatastores200ResponseDatastore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZone

`func (o *UpdateCloudDatastores200ResponseDatastore) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateCloudDatastores200ResponseDatastore) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateCloudDatastores200ResponseDatastore) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateCloudDatastores200ResponseDatastore) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetType

`func (o *UpdateCloudDatastores200ResponseDatastore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCloudDatastores200ResponseDatastore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCloudDatastores200ResponseDatastore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateCloudDatastores200ResponseDatastore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFreeSpace

`func (o *UpdateCloudDatastores200ResponseDatastore) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *UpdateCloudDatastores200ResponseDatastore) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *UpdateCloudDatastores200ResponseDatastore) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *UpdateCloudDatastores200ResponseDatastore) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### GetOnline

`func (o *UpdateCloudDatastores200ResponseDatastore) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *UpdateCloudDatastores200ResponseDatastore) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *UpdateCloudDatastores200ResponseDatastore) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *UpdateCloudDatastores200ResponseDatastore) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetActive

`func (o *UpdateCloudDatastores200ResponseDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateCloudDatastores200ResponseDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateCloudDatastores200ResponseDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateCloudDatastores200ResponseDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateCloudDatastores200ResponseDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateCloudDatastores200ResponseDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateCloudDatastores200ResponseDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateCloudDatastores200ResponseDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateCloudDatastores200ResponseDatastore) GetTenants() []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateCloudDatastores200ResponseDatastore) GetTenantsOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateCloudDatastores200ResponseDatastore) SetTenants(v []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateCloudDatastores200ResponseDatastore) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *UpdateCloudDatastores200ResponseDatastore) GetResourcePermission() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *UpdateCloudDatastores200ResponseDatastore) GetResourcePermissionOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *UpdateCloudDatastores200ResponseDatastore) SetResourcePermission(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *UpdateCloudDatastores200ResponseDatastore) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateCloudDatastores200ResponseDatastore) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateCloudDatastores200ResponseDatastore) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateCloudDatastores200ResponseDatastore) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateCloudDatastores200ResponseDatastore) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



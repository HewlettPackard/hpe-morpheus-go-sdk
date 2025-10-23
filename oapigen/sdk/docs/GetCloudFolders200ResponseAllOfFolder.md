# GetCloudFolders200ResponseAllOfFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Parent** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**DefaultFolder** | Pointer to **bool** |  | [optional] 
**DefaultStore** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Tenants** | Pointer to [**[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner**](ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission**](ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission.md) |  | [optional] 
**Depth** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetCloudFolders200ResponseAllOfFolder

`func NewGetCloudFolders200ResponseAllOfFolder() *GetCloudFolders200ResponseAllOfFolder`

NewGetCloudFolders200ResponseAllOfFolder instantiates a new GetCloudFolders200ResponseAllOfFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCloudFolders200ResponseAllOfFolderWithDefaults

`func NewGetCloudFolders200ResponseAllOfFolderWithDefaults() *GetCloudFolders200ResponseAllOfFolder`

NewGetCloudFolders200ResponseAllOfFolderWithDefaults instantiates a new GetCloudFolders200ResponseAllOfFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCloudFolders200ResponseAllOfFolder) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCloudFolders200ResponseAllOfFolder) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCloudFolders200ResponseAllOfFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetCloudFolders200ResponseAllOfFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCloudFolders200ResponseAllOfFolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCloudFolders200ResponseAllOfFolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZone

`func (o *GetCloudFolders200ResponseAllOfFolder) GetZone() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetZoneOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetCloudFolders200ResponseAllOfFolder) SetZone(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetCloudFolders200ResponseAllOfFolder) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetParent

`func (o *GetCloudFolders200ResponseAllOfFolder) GetParent() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetParentOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetCloudFolders200ResponseAllOfFolder) SetParent(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetCloudFolders200ResponseAllOfFolder) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *GetCloudFolders200ResponseAllOfFolder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCloudFolders200ResponseAllOfFolder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCloudFolders200ResponseAllOfFolder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *GetCloudFolders200ResponseAllOfFolder) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetCloudFolders200ResponseAllOfFolder) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetCloudFolders200ResponseAllOfFolder) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVisibility

`func (o *GetCloudFolders200ResponseAllOfFolder) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetCloudFolders200ResponseAllOfFolder) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetCloudFolders200ResponseAllOfFolder) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetReadOnly

`func (o *GetCloudFolders200ResponseAllOfFolder) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *GetCloudFolders200ResponseAllOfFolder) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *GetCloudFolders200ResponseAllOfFolder) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetDefaultFolder

`func (o *GetCloudFolders200ResponseAllOfFolder) GetDefaultFolder() bool`

GetDefaultFolder returns the DefaultFolder field if non-nil, zero value otherwise.

### GetDefaultFolderOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetDefaultFolderOk() (*bool, bool)`

GetDefaultFolderOk returns a tuple with the DefaultFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFolder

`func (o *GetCloudFolders200ResponseAllOfFolder) SetDefaultFolder(v bool)`

SetDefaultFolder sets DefaultFolder field to given value.

### HasDefaultFolder

`func (o *GetCloudFolders200ResponseAllOfFolder) HasDefaultFolder() bool`

HasDefaultFolder returns a boolean if a field has been set.

### GetDefaultStore

`func (o *GetCloudFolders200ResponseAllOfFolder) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *GetCloudFolders200ResponseAllOfFolder) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *GetCloudFolders200ResponseAllOfFolder) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetActive

`func (o *GetCloudFolders200ResponseAllOfFolder) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetCloudFolders200ResponseAllOfFolder) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetCloudFolders200ResponseAllOfFolder) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetTenants

`func (o *GetCloudFolders200ResponseAllOfFolder) GetTenants() []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetTenantsOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetCloudFolders200ResponseAllOfFolder) SetTenants(v []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetCloudFolders200ResponseAllOfFolder) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *GetCloudFolders200ResponseAllOfFolder) GetResourcePermissions() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetResourcePermissionsOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *GetCloudFolders200ResponseAllOfFolder) SetResourcePermissions(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *GetCloudFolders200ResponseAllOfFolder) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetDepth

`func (o *GetCloudFolders200ResponseAllOfFolder) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *GetCloudFolders200ResponseAllOfFolder) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *GetCloudFolders200ResponseAllOfFolder) SetDepth(v int64)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *GetCloudFolders200ResponseAllOfFolder) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ListCloudFolders200ResponseAllOfFoldersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
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

### NewListCloudFolders200ResponseAllOfFoldersInner

`func NewListCloudFolders200ResponseAllOfFoldersInner() *ListCloudFolders200ResponseAllOfFoldersInner`

NewListCloudFolders200ResponseAllOfFoldersInner instantiates a new ListCloudFolders200ResponseAllOfFoldersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCloudFolders200ResponseAllOfFoldersInnerWithDefaults

`func NewListCloudFolders200ResponseAllOfFoldersInnerWithDefaults() *ListCloudFolders200ResponseAllOfFoldersInner`

NewListCloudFolders200ResponseAllOfFoldersInnerWithDefaults instantiates a new ListCloudFolders200ResponseAllOfFoldersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZone

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetParent

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetParent() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetParentOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetParent(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVisibility

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetReadOnly

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetDefaultFolder

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetDefaultFolder() bool`

GetDefaultFolder returns the DefaultFolder field if non-nil, zero value otherwise.

### GetDefaultFolderOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetDefaultFolderOk() (*bool, bool)`

GetDefaultFolderOk returns a tuple with the DefaultFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFolder

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetDefaultFolder(v bool)`

SetDefaultFolder sets DefaultFolder field to given value.

### HasDefaultFolder

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasDefaultFolder() bool`

HasDefaultFolder returns a boolean if a field has been set.

### GetDefaultStore

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetActive

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetTenants

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetTenants() []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetTenantsOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetTenants(v []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetResourcePermissions() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetResourcePermissionsOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetResourcePermissions(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetDepth

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) SetDepth(v int64)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *ListCloudFolders200ResponseAllOfFoldersInner) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



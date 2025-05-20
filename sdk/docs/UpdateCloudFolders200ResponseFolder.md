# UpdateCloudFolders200ResponseFolder

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
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateCloudFolders200ResponseFolder

`func NewUpdateCloudFolders200ResponseFolder() *UpdateCloudFolders200ResponseFolder`

NewUpdateCloudFolders200ResponseFolder instantiates a new UpdateCloudFolders200ResponseFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCloudFolders200ResponseFolderWithDefaults

`func NewUpdateCloudFolders200ResponseFolderWithDefaults() *UpdateCloudFolders200ResponseFolder`

NewUpdateCloudFolders200ResponseFolderWithDefaults instantiates a new UpdateCloudFolders200ResponseFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateCloudFolders200ResponseFolder) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCloudFolders200ResponseFolder) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCloudFolders200ResponseFolder) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCloudFolders200ResponseFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateCloudFolders200ResponseFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCloudFolders200ResponseFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCloudFolders200ResponseFolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCloudFolders200ResponseFolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZone

`func (o *UpdateCloudFolders200ResponseFolder) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateCloudFolders200ResponseFolder) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateCloudFolders200ResponseFolder) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateCloudFolders200ResponseFolder) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetParent

`func (o *UpdateCloudFolders200ResponseFolder) GetParent() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *UpdateCloudFolders200ResponseFolder) GetParentOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *UpdateCloudFolders200ResponseFolder) SetParent(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetParent sets Parent field to given value.

### HasParent

`func (o *UpdateCloudFolders200ResponseFolder) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *UpdateCloudFolders200ResponseFolder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCloudFolders200ResponseFolder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCloudFolders200ResponseFolder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateCloudFolders200ResponseFolder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateCloudFolders200ResponseFolder) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateCloudFolders200ResponseFolder) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateCloudFolders200ResponseFolder) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateCloudFolders200ResponseFolder) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateCloudFolders200ResponseFolder) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateCloudFolders200ResponseFolder) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateCloudFolders200ResponseFolder) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateCloudFolders200ResponseFolder) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetReadOnly

`func (o *UpdateCloudFolders200ResponseFolder) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *UpdateCloudFolders200ResponseFolder) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *UpdateCloudFolders200ResponseFolder) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *UpdateCloudFolders200ResponseFolder) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetDefaultFolder

`func (o *UpdateCloudFolders200ResponseFolder) GetDefaultFolder() bool`

GetDefaultFolder returns the DefaultFolder field if non-nil, zero value otherwise.

### GetDefaultFolderOk

`func (o *UpdateCloudFolders200ResponseFolder) GetDefaultFolderOk() (*bool, bool)`

GetDefaultFolderOk returns a tuple with the DefaultFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFolder

`func (o *UpdateCloudFolders200ResponseFolder) SetDefaultFolder(v bool)`

SetDefaultFolder sets DefaultFolder field to given value.

### HasDefaultFolder

`func (o *UpdateCloudFolders200ResponseFolder) HasDefaultFolder() bool`

HasDefaultFolder returns a boolean if a field has been set.

### GetDefaultStore

`func (o *UpdateCloudFolders200ResponseFolder) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *UpdateCloudFolders200ResponseFolder) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *UpdateCloudFolders200ResponseFolder) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *UpdateCloudFolders200ResponseFolder) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetActive

`func (o *UpdateCloudFolders200ResponseFolder) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateCloudFolders200ResponseFolder) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateCloudFolders200ResponseFolder) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateCloudFolders200ResponseFolder) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateCloudFolders200ResponseFolder) GetTenants() []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateCloudFolders200ResponseFolder) GetTenantsOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateCloudFolders200ResponseFolder) SetTenants(v []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateCloudFolders200ResponseFolder) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *UpdateCloudFolders200ResponseFolder) GetResourcePermissions() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *UpdateCloudFolders200ResponseFolder) GetResourcePermissionsOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *UpdateCloudFolders200ResponseFolder) SetResourcePermissions(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *UpdateCloudFolders200ResponseFolder) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetDepth

`func (o *UpdateCloudFolders200ResponseFolder) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *UpdateCloudFolders200ResponseFolder) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *UpdateCloudFolders200ResponseFolder) SetDepth(v int64)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *UpdateCloudFolders200ResponseFolder) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateCloudFolders200ResponseFolder) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateCloudFolders200ResponseFolder) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateCloudFolders200ResponseFolder) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateCloudFolders200ResponseFolder) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateCloudFoldersRequestFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultFolder** | Pointer to **bool** |  | [optional] [default to false]
**DefaultImage** | Pointer to **bool** |  | [optional] [default to false]
**Active** | Pointer to **bool** | Activate &#x60;true&#x60; or disable &#x60;false&#x60; the folder | [optional] 
**Visibility** | Pointer to **string** | Setting &#x60;private&#x60; or &#x60;public&#x60; | [optional] [default to "private"]
**TenantPermissions** | Pointer to [**UpdateCloudFoldersRequestFolderTenantPermissions**](UpdateCloudFoldersRequestFolderTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**UpdateCloudDatastoresRequestDatastoreResourcePermissions**](UpdateCloudDatastoresRequestDatastoreResourcePermissions.md) |  | [optional] 

## Methods

### NewUpdateCloudFoldersRequestFolder

`func NewUpdateCloudFoldersRequestFolder() *UpdateCloudFoldersRequestFolder`

NewUpdateCloudFoldersRequestFolder instantiates a new UpdateCloudFoldersRequestFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCloudFoldersRequestFolderWithDefaults

`func NewUpdateCloudFoldersRequestFolderWithDefaults() *UpdateCloudFoldersRequestFolder`

NewUpdateCloudFoldersRequestFolderWithDefaults instantiates a new UpdateCloudFoldersRequestFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultFolder

`func (o *UpdateCloudFoldersRequestFolder) GetDefaultFolder() bool`

GetDefaultFolder returns the DefaultFolder field if non-nil, zero value otherwise.

### GetDefaultFolderOk

`func (o *UpdateCloudFoldersRequestFolder) GetDefaultFolderOk() (*bool, bool)`

GetDefaultFolderOk returns a tuple with the DefaultFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFolder

`func (o *UpdateCloudFoldersRequestFolder) SetDefaultFolder(v bool)`

SetDefaultFolder sets DefaultFolder field to given value.

### HasDefaultFolder

`func (o *UpdateCloudFoldersRequestFolder) HasDefaultFolder() bool`

HasDefaultFolder returns a boolean if a field has been set.

### GetDefaultImage

`func (o *UpdateCloudFoldersRequestFolder) GetDefaultImage() bool`

GetDefaultImage returns the DefaultImage field if non-nil, zero value otherwise.

### GetDefaultImageOk

`func (o *UpdateCloudFoldersRequestFolder) GetDefaultImageOk() (*bool, bool)`

GetDefaultImageOk returns a tuple with the DefaultImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImage

`func (o *UpdateCloudFoldersRequestFolder) SetDefaultImage(v bool)`

SetDefaultImage sets DefaultImage field to given value.

### HasDefaultImage

`func (o *UpdateCloudFoldersRequestFolder) HasDefaultImage() bool`

HasDefaultImage returns a boolean if a field has been set.

### GetActive

`func (o *UpdateCloudFoldersRequestFolder) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateCloudFoldersRequestFolder) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateCloudFoldersRequestFolder) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateCloudFoldersRequestFolder) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateCloudFoldersRequestFolder) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateCloudFoldersRequestFolder) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateCloudFoldersRequestFolder) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateCloudFoldersRequestFolder) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *UpdateCloudFoldersRequestFolder) GetTenantPermissions() UpdateCloudFoldersRequestFolderTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *UpdateCloudFoldersRequestFolder) GetTenantPermissionsOk() (*UpdateCloudFoldersRequestFolderTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *UpdateCloudFoldersRequestFolder) SetTenantPermissions(v UpdateCloudFoldersRequestFolderTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *UpdateCloudFoldersRequestFolder) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *UpdateCloudFoldersRequestFolder) GetResourcePermissions() UpdateCloudDatastoresRequestDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *UpdateCloudFoldersRequestFolder) GetResourcePermissionsOk() (*UpdateCloudDatastoresRequestDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *UpdateCloudFoldersRequestFolder) SetResourcePermissions(v UpdateCloudDatastoresRequestDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *UpdateCloudFoldersRequestFolder) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



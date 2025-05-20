# NetworkRouterPermissionsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visibility** | Pointer to **string** | Sets visibility - public, private | [optional] 
**TenantPermissions** | Pointer to [**UpdateNetworkRouterPermissionsRequestPermissionsTenantPermissions**](UpdateNetworkRouterPermissionsRequestPermissionsTenantPermissions.md) |  | [optional] 

## Methods

### NewNetworkRouterPermissionsUpdate

`func NewNetworkRouterPermissionsUpdate() *NetworkRouterPermissionsUpdate`

NewNetworkRouterPermissionsUpdate instantiates a new NetworkRouterPermissionsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRouterPermissionsUpdateWithDefaults

`func NewNetworkRouterPermissionsUpdateWithDefaults() *NetworkRouterPermissionsUpdate`

NewNetworkRouterPermissionsUpdateWithDefaults instantiates a new NetworkRouterPermissionsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisibility

`func (o *NetworkRouterPermissionsUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *NetworkRouterPermissionsUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *NetworkRouterPermissionsUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *NetworkRouterPermissionsUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *NetworkRouterPermissionsUpdate) GetTenantPermissions() UpdateNetworkRouterPermissionsRequestPermissionsTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *NetworkRouterPermissionsUpdate) GetTenantPermissionsOk() (*UpdateNetworkRouterPermissionsRequestPermissionsTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *NetworkRouterPermissionsUpdate) SetTenantPermissions(v UpdateNetworkRouterPermissionsRequestPermissionsTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *NetworkRouterPermissionsUpdate) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



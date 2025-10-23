# UpdateCloudResourcePoolRequestResourcePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Activate &#x60;true&#x60; or disable &#x60;false&#x60; the datastore | [optional] 
**Visibility** | Pointer to **string** | Setting &#x60;private&#x60; or &#x60;public&#x60; | [optional] [default to "private"]
**DisplayName** | Pointer to **string** | Optional Display Name (VMware only) | [optional] 
**Inventory** | Pointer to **bool** | Enable &#x60;True&#x60; or disable &#x60;False&#x60; inventory sync for resource pool during cloud refresh | [optional] 
**TenantPermissions** | Pointer to [**AddCloudResourcePoolRequestResourcePoolTenantPermissions**](AddCloudResourcePoolRequestResourcePoolTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**UpdateCloudDatastoresRequestDatastoreResourcePermissions**](UpdateCloudDatastoresRequestDatastoreResourcePermissions.md) |  | [optional] 

## Methods

### NewUpdateCloudResourcePoolRequestResourcePool

`func NewUpdateCloudResourcePoolRequestResourcePool() *UpdateCloudResourcePoolRequestResourcePool`

NewUpdateCloudResourcePoolRequestResourcePool instantiates a new UpdateCloudResourcePoolRequestResourcePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCloudResourcePoolRequestResourcePoolWithDefaults

`func NewUpdateCloudResourcePoolRequestResourcePoolWithDefaults() *UpdateCloudResourcePoolRequestResourcePool`

NewUpdateCloudResourcePoolRequestResourcePoolWithDefaults instantiates a new UpdateCloudResourcePoolRequestResourcePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateCloudResourcePoolRequestResourcePool) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateCloudResourcePoolRequestResourcePool) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateCloudResourcePoolRequestResourcePool) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateCloudResourcePoolRequestResourcePool) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateCloudResourcePoolRequestResourcePool) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateCloudResourcePoolRequestResourcePool) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateCloudResourcePoolRequestResourcePool) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateCloudResourcePoolRequestResourcePool) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateCloudResourcePoolRequestResourcePool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateCloudResourcePoolRequestResourcePool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateCloudResourcePoolRequestResourcePool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateCloudResourcePoolRequestResourcePool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetInventory

`func (o *UpdateCloudResourcePoolRequestResourcePool) GetInventory() bool`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *UpdateCloudResourcePoolRequestResourcePool) GetInventoryOk() (*bool, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *UpdateCloudResourcePoolRequestResourcePool) SetInventory(v bool)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *UpdateCloudResourcePoolRequestResourcePool) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *UpdateCloudResourcePoolRequestResourcePool) GetTenantPermissions() AddCloudResourcePoolRequestResourcePoolTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *UpdateCloudResourcePoolRequestResourcePool) GetTenantPermissionsOk() (*AddCloudResourcePoolRequestResourcePoolTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *UpdateCloudResourcePoolRequestResourcePool) SetTenantPermissions(v AddCloudResourcePoolRequestResourcePoolTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *UpdateCloudResourcePoolRequestResourcePool) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *UpdateCloudResourcePoolRequestResourcePool) GetResourcePermissions() UpdateCloudDatastoresRequestDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *UpdateCloudResourcePoolRequestResourcePool) GetResourcePermissionsOk() (*UpdateCloudDatastoresRequestDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *UpdateCloudResourcePoolRequestResourcePool) SetResourcePermissions(v UpdateCloudDatastoresRequestDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *UpdateCloudResourcePoolRequestResourcePool) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



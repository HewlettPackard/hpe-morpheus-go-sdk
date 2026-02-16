# UpdateDatastoresRequestDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Activate &#x60;true&#x60; or disable &#x60;false&#x60; the datastore | [optional] 
**Visibility** | Pointer to **string** | Setting &#x60;private&#x60; or &#x60;public&#x60; | [optional] [default to "private"]
**TenantPermissions** | Pointer to [**UpdateDatastoresRequestDatastoreTenantPermissions**](UpdateDatastoresRequestDatastoreTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**UpdateDatastoresRequestDatastoreResourcePermissions**](UpdateDatastoresRequestDatastoreResourcePermissions.md) |  | [optional] 

## Methods

### NewUpdateDatastoresRequestDatastore

`func NewUpdateDatastoresRequestDatastore() *UpdateDatastoresRequestDatastore`

NewUpdateDatastoresRequestDatastore instantiates a new UpdateDatastoresRequestDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatastoresRequestDatastoreWithDefaults

`func NewUpdateDatastoresRequestDatastoreWithDefaults() *UpdateDatastoresRequestDatastore`

NewUpdateDatastoresRequestDatastoreWithDefaults instantiates a new UpdateDatastoresRequestDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateDatastoresRequestDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateDatastoresRequestDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateDatastoresRequestDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateDatastoresRequestDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateDatastoresRequestDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateDatastoresRequestDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateDatastoresRequestDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateDatastoresRequestDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *UpdateDatastoresRequestDatastore) GetTenantPermissions() UpdateDatastoresRequestDatastoreTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *UpdateDatastoresRequestDatastore) GetTenantPermissionsOk() (*UpdateDatastoresRequestDatastoreTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *UpdateDatastoresRequestDatastore) SetTenantPermissions(v UpdateDatastoresRequestDatastoreTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *UpdateDatastoresRequestDatastore) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *UpdateDatastoresRequestDatastore) GetResourcePermissions() UpdateDatastoresRequestDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *UpdateDatastoresRequestDatastore) GetResourcePermissionsOk() (*UpdateDatastoresRequestDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *UpdateDatastoresRequestDatastore) SetResourcePermissions(v UpdateDatastoresRequestDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *UpdateDatastoresRequestDatastore) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



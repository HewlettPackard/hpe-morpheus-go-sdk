# UpdateNetworkServerRequestNetworkServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**ServiceUrl** | Pointer to **string** | URL | [optional] 
**ServiceUsername** | Pointer to **string** | Username | [optional] 
**ServicePassword** | Pointer to **string** | Password | [optional] 
**Visibility** | Pointer to **string** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**Tenants** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of tenant account ids that are allowed access | [optional] 

## Methods

### NewUpdateNetworkServerRequestNetworkServer

`func NewUpdateNetworkServerRequestNetworkServer() *UpdateNetworkServerRequestNetworkServer`

NewUpdateNetworkServerRequestNetworkServer instantiates a new UpdateNetworkServerRequestNetworkServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkServerRequestNetworkServerWithDefaults

`func NewUpdateNetworkServerRequestNetworkServerWithDefaults() *UpdateNetworkServerRequestNetworkServer`

NewUpdateNetworkServerRequestNetworkServerWithDefaults instantiates a new UpdateNetworkServerRequestNetworkServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkServerRequestNetworkServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkServerRequestNetworkServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkServerRequestNetworkServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkServerRequestNetworkServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceUrl

`func (o *UpdateNetworkServerRequestNetworkServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *UpdateNetworkServerRequestNetworkServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *UpdateNetworkServerRequestNetworkServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *UpdateNetworkServerRequestNetworkServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceUsername

`func (o *UpdateNetworkServerRequestNetworkServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *UpdateNetworkServerRequestNetworkServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *UpdateNetworkServerRequestNetworkServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *UpdateNetworkServerRequestNetworkServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *UpdateNetworkServerRequestNetworkServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *UpdateNetworkServerRequestNetworkServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *UpdateNetworkServerRequestNetworkServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *UpdateNetworkServerRequestNetworkServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateNetworkServerRequestNetworkServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateNetworkServerRequestNetworkServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateNetworkServerRequestNetworkServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateNetworkServerRequestNetworkServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateNetworkServerRequestNetworkServer) GetTenants() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateNetworkServerRequestNetworkServer) GetTenantsOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateNetworkServerRequestNetworkServer) SetTenants(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateNetworkServerRequestNetworkServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



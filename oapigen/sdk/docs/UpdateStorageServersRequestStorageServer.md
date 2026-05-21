# UpdateStorageServersRequestStorageServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Type** | Pointer to **string** | The &#x60;Storage Type&#x60; Code or ID | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**Enabled** | Pointer to **bool** | The enabled flag | [optional] [default to true]
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by &#x60;type&#x60; | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**ServiceHost** | Pointer to **string** | Storage server host | [optional] 
**ServiceUrl** | Pointer to **string** | Storage server URL | [optional] 
**ServiceUsername** | Pointer to **string** | Service username for authentication | [optional] 
**ServicePassword** | Pointer to **string** | Service password for authentication | [optional] 
**ServicePort** | Pointer to **int32** | Service port | [optional] 
**Credential** | Pointer to [**UpdateStorageServersRequestStorageServerCredential**](UpdateStorageServersRequestStorageServerCredential.md) |  | [optional] 
**Tenants** | Pointer to [**[]UpdateStorageServersRequestStorageServerTenantsInner**](UpdateStorageServersRequestStorageServerTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 

## Methods

### NewUpdateStorageServersRequestStorageServer

`func NewUpdateStorageServersRequestStorageServer() *UpdateStorageServersRequestStorageServer`

NewUpdateStorageServersRequestStorageServer instantiates a new UpdateStorageServersRequestStorageServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageServersRequestStorageServerWithDefaults

`func NewUpdateStorageServersRequestStorageServerWithDefaults() *UpdateStorageServersRequestStorageServer`

NewUpdateStorageServersRequestStorageServerWithDefaults instantiates a new UpdateStorageServersRequestStorageServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateStorageServersRequestStorageServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStorageServersRequestStorageServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStorageServersRequestStorageServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateStorageServersRequestStorageServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *UpdateStorageServersRequestStorageServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateStorageServersRequestStorageServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateStorageServersRequestStorageServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateStorageServersRequestStorageServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateStorageServersRequestStorageServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateStorageServersRequestStorageServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateStorageServersRequestStorageServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateStorageServersRequestStorageServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateStorageServersRequestStorageServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateStorageServersRequestStorageServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateStorageServersRequestStorageServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateStorageServersRequestStorageServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateStorageServersRequestStorageServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateStorageServersRequestStorageServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateStorageServersRequestStorageServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateStorageServersRequestStorageServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateStorageServersRequestStorageServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateStorageServersRequestStorageServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateStorageServersRequestStorageServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateStorageServersRequestStorageServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetServiceHost

`func (o *UpdateStorageServersRequestStorageServer) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *UpdateStorageServersRequestStorageServer) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *UpdateStorageServersRequestStorageServer) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *UpdateStorageServersRequestStorageServer) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### GetServiceUrl

`func (o *UpdateStorageServersRequestStorageServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *UpdateStorageServersRequestStorageServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *UpdateStorageServersRequestStorageServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *UpdateStorageServersRequestStorageServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceUsername

`func (o *UpdateStorageServersRequestStorageServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *UpdateStorageServersRequestStorageServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *UpdateStorageServersRequestStorageServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *UpdateStorageServersRequestStorageServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *UpdateStorageServersRequestStorageServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *UpdateStorageServersRequestStorageServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *UpdateStorageServersRequestStorageServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *UpdateStorageServersRequestStorageServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServicePort

`func (o *UpdateStorageServersRequestStorageServer) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *UpdateStorageServersRequestStorageServer) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *UpdateStorageServersRequestStorageServer) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *UpdateStorageServersRequestStorageServer) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetCredential

`func (o *UpdateStorageServersRequestStorageServer) GetCredential() UpdateStorageServersRequestStorageServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *UpdateStorageServersRequestStorageServer) GetCredentialOk() (*UpdateStorageServersRequestStorageServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *UpdateStorageServersRequestStorageServer) SetCredential(v UpdateStorageServersRequestStorageServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *UpdateStorageServersRequestStorageServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateStorageServersRequestStorageServer) GetTenants() []UpdateStorageServersRequestStorageServerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateStorageServersRequestStorageServer) GetTenantsOk() (*[]UpdateStorageServersRequestStorageServerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateStorageServersRequestStorageServer) SetTenants(v []UpdateStorageServersRequestStorageServerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateStorageServersRequestStorageServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



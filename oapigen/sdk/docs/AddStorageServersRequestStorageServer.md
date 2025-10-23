# AddStorageServersRequestStorageServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**Type** | **string** | The &#x60;Storage Type&#x60; Code or ID | 
**Description** | Pointer to **string** | description | [optional] 
**Enabled** | Pointer to **bool** | The enabled flag | [optional] [default to true]
**Config** | **map[string]interface{}** | Configuration object with parameters that vary by &#x60;type&#x60; | 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**Tenants** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of tenant account ids that are allowed access | [optional] 

## Methods

### NewAddStorageServersRequestStorageServer

`func NewAddStorageServersRequestStorageServer(name string, type_ string, config map[string]interface{}, ) *AddStorageServersRequestStorageServer`

NewAddStorageServersRequestStorageServer instantiates a new AddStorageServersRequestStorageServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStorageServersRequestStorageServerWithDefaults

`func NewAddStorageServersRequestStorageServerWithDefaults() *AddStorageServersRequestStorageServer`

NewAddStorageServersRequestStorageServerWithDefaults instantiates a new AddStorageServersRequestStorageServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddStorageServersRequestStorageServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddStorageServersRequestStorageServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddStorageServersRequestStorageServer) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddStorageServersRequestStorageServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddStorageServersRequestStorageServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddStorageServersRequestStorageServer) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *AddStorageServersRequestStorageServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddStorageServersRequestStorageServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddStorageServersRequestStorageServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddStorageServersRequestStorageServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddStorageServersRequestStorageServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddStorageServersRequestStorageServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddStorageServersRequestStorageServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddStorageServersRequestStorageServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfig

`func (o *AddStorageServersRequestStorageServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddStorageServersRequestStorageServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddStorageServersRequestStorageServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetVisibility

`func (o *AddStorageServersRequestStorageServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddStorageServersRequestStorageServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddStorageServersRequestStorageServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddStorageServersRequestStorageServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *AddStorageServersRequestStorageServer) GetTenants() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *AddStorageServersRequestStorageServer) GetTenantsOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *AddStorageServersRequestStorageServer) SetTenants(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *AddStorageServersRequestStorageServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



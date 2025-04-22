# UpdateTenantRequestAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Role** | Pointer to [**UpdateTenantRequestAccountRole**](UpdateTenantRequestAccountRole.md) |  | [optional] 
**Subdomain** | Pointer to **string** | The subdomain. This will be part of the login URL and username for sub tenant users. | [optional] 
**Currency** | Pointer to **string** | Currency Code (ISO 4217) | [optional] [default to "USD"]

## Methods

### NewUpdateTenantRequestAccount

`func NewUpdateTenantRequestAccount() *UpdateTenantRequestAccount`

NewUpdateTenantRequestAccount instantiates a new UpdateTenantRequestAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantRequestAccountWithDefaults

`func NewUpdateTenantRequestAccountWithDefaults() *UpdateTenantRequestAccount`

NewUpdateTenantRequestAccountWithDefaults instantiates a new UpdateTenantRequestAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTenantRequestAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTenantRequestAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTenantRequestAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTenantRequestAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateTenantRequestAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTenantRequestAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTenantRequestAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTenantRequestAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *UpdateTenantRequestAccount) GetRole() UpdateTenantRequestAccountRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateTenantRequestAccount) GetRoleOk() (*UpdateTenantRequestAccountRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateTenantRequestAccount) SetRole(v UpdateTenantRequestAccountRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateTenantRequestAccount) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSubdomain

`func (o *UpdateTenantRequestAccount) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *UpdateTenantRequestAccount) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *UpdateTenantRequestAccount) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *UpdateTenantRequestAccount) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateTenantRequestAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateTenantRequestAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateTenantRequestAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateTenantRequestAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



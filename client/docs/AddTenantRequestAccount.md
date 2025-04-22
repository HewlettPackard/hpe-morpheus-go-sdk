# AddTenantRequestAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**Description** | Pointer to **string** | Description | [optional] 
**Role** | Pointer to [**AddTenantRequestAccountRole**](AddTenantRequestAccountRole.md) |  | [optional] 
**Subdomain** | Pointer to **string** | The subdomain. This will be part of the login URL and username for sub tenant users. | [optional] 
**Currency** | Pointer to **string** | Currency Code (ISO 4217) | [optional] [default to "USD"]

## Methods

### NewAddTenantRequestAccount

`func NewAddTenantRequestAccount(name string, ) *AddTenantRequestAccount`

NewAddTenantRequestAccount instantiates a new AddTenantRequestAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTenantRequestAccountWithDefaults

`func NewAddTenantRequestAccountWithDefaults() *AddTenantRequestAccount`

NewAddTenantRequestAccountWithDefaults instantiates a new AddTenantRequestAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddTenantRequestAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddTenantRequestAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddTenantRequestAccount) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddTenantRequestAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddTenantRequestAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddTenantRequestAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddTenantRequestAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *AddTenantRequestAccount) GetRole() AddTenantRequestAccountRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddTenantRequestAccount) GetRoleOk() (*AddTenantRequestAccountRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddTenantRequestAccount) SetRole(v AddTenantRequestAccountRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *AddTenantRequestAccount) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSubdomain

`func (o *AddTenantRequestAccount) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *AddTenantRequestAccount) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *AddTenantRequestAccount) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *AddTenantRequestAccount) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetCurrency

`func (o *AddTenantRequestAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AddTenantRequestAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AddTenantRequestAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AddTenantRequestAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



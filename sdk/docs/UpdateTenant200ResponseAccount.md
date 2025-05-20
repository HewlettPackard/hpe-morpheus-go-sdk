# UpdateTenant200ResponseAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**CustomerNumber** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Master** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to [**ListTenants200ResponseAllOfAccountsInnerRole**](ListTenants200ResponseAllOfAccountsInnerRole.md) |  | [optional] 
**Stats** | Pointer to [**ListTenants200ResponseAllOfAccountsInnerStats**](ListTenants200ResponseAllOfAccountsInnerStats.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateTenant200ResponseAccount

`func NewUpdateTenant200ResponseAccount() *UpdateTenant200ResponseAccount`

NewUpdateTenant200ResponseAccount instantiates a new UpdateTenant200ResponseAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenant200ResponseAccountWithDefaults

`func NewUpdateTenant200ResponseAccountWithDefaults() *UpdateTenant200ResponseAccount`

NewUpdateTenant200ResponseAccountWithDefaults instantiates a new UpdateTenant200ResponseAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateTenant200ResponseAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateTenant200ResponseAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateTenant200ResponseAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateTenant200ResponseAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateTenant200ResponseAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTenant200ResponseAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTenant200ResponseAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTenant200ResponseAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateTenant200ResponseAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTenant200ResponseAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTenant200ResponseAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTenant200ResponseAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubdomain

`func (o *UpdateTenant200ResponseAccount) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *UpdateTenant200ResponseAccount) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *UpdateTenant200ResponseAccount) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *UpdateTenant200ResponseAccount) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateTenant200ResponseAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateTenant200ResponseAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateTenant200ResponseAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateTenant200ResponseAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateTenant200ResponseAccount) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateTenant200ResponseAccount) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateTenant200ResponseAccount) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateTenant200ResponseAccount) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCustomerNumber

`func (o *UpdateTenant200ResponseAccount) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *UpdateTenant200ResponseAccount) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *UpdateTenant200ResponseAccount) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.

### HasCustomerNumber

`func (o *UpdateTenant200ResponseAccount) HasCustomerNumber() bool`

HasCustomerNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *UpdateTenant200ResponseAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *UpdateTenant200ResponseAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *UpdateTenant200ResponseAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *UpdateTenant200ResponseAccount) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountName

`func (o *UpdateTenant200ResponseAccount) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *UpdateTenant200ResponseAccount) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *UpdateTenant200ResponseAccount) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *UpdateTenant200ResponseAccount) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetActive

`func (o *UpdateTenant200ResponseAccount) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateTenant200ResponseAccount) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateTenant200ResponseAccount) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateTenant200ResponseAccount) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMaster

`func (o *UpdateTenant200ResponseAccount) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *UpdateTenant200ResponseAccount) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *UpdateTenant200ResponseAccount) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *UpdateTenant200ResponseAccount) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetRole

`func (o *UpdateTenant200ResponseAccount) GetRole() ListTenants200ResponseAllOfAccountsInnerRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateTenant200ResponseAccount) GetRoleOk() (*ListTenants200ResponseAllOfAccountsInnerRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateTenant200ResponseAccount) SetRole(v ListTenants200ResponseAllOfAccountsInnerRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateTenant200ResponseAccount) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStats

`func (o *UpdateTenant200ResponseAccount) GetStats() ListTenants200ResponseAllOfAccountsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *UpdateTenant200ResponseAccount) GetStatsOk() (*ListTenants200ResponseAllOfAccountsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *UpdateTenant200ResponseAccount) SetStats(v ListTenants200ResponseAllOfAccountsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *UpdateTenant200ResponseAccount) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateTenant200ResponseAccount) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateTenant200ResponseAccount) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateTenant200ResponseAccount) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateTenant200ResponseAccount) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateTenant200ResponseAccount) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateTenant200ResponseAccount) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateTenant200ResponseAccount) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateTenant200ResponseAccount) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateTenant200ResponseAccount) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateTenant200ResponseAccount) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateTenant200ResponseAccount) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateTenant200ResponseAccount) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



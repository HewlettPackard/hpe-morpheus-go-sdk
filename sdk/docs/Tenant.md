# Tenant

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

## Methods

### NewTenant

`func NewTenant() *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tenant) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tenant) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tenant) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Tenant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Tenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tenant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Tenant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tenant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tenant) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tenant) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubdomain

`func (o *Tenant) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *Tenant) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *Tenant) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *Tenant) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetCurrency

`func (o *Tenant) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Tenant) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Tenant) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Tenant) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalId

`func (o *Tenant) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Tenant) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Tenant) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Tenant) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCustomerNumber

`func (o *Tenant) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *Tenant) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *Tenant) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.

### HasCustomerNumber

`func (o *Tenant) HasCustomerNumber() bool`

HasCustomerNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *Tenant) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Tenant) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Tenant) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Tenant) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountName

`func (o *Tenant) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Tenant) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Tenant) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *Tenant) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetActive

`func (o *Tenant) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Tenant) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Tenant) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Tenant) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMaster

`func (o *Tenant) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *Tenant) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *Tenant) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *Tenant) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetRole

`func (o *Tenant) GetRole() ListTenants200ResponseAllOfAccountsInnerRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Tenant) GetRoleOk() (*ListTenants200ResponseAllOfAccountsInnerRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Tenant) SetRole(v ListTenants200ResponseAllOfAccountsInnerRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *Tenant) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStats

`func (o *Tenant) GetStats() ListTenants200ResponseAllOfAccountsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Tenant) GetStatsOk() (*ListTenants200ResponseAllOfAccountsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Tenant) SetStats(v ListTenants200ResponseAllOfAccountsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Tenant) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetDateCreated

`func (o *Tenant) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Tenant) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Tenant) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Tenant) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Tenant) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Tenant) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Tenant) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Tenant) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



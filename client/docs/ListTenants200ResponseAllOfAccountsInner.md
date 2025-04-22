# ListTenants200ResponseAllOfAccountsInner

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

### NewListTenants200ResponseAllOfAccountsInner

`func NewListTenants200ResponseAllOfAccountsInner() *ListTenants200ResponseAllOfAccountsInner`

NewListTenants200ResponseAllOfAccountsInner instantiates a new ListTenants200ResponseAllOfAccountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTenants200ResponseAllOfAccountsInnerWithDefaults

`func NewListTenants200ResponseAllOfAccountsInnerWithDefaults() *ListTenants200ResponseAllOfAccountsInner`

NewListTenants200ResponseAllOfAccountsInnerWithDefaults instantiates a new ListTenants200ResponseAllOfAccountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListTenants200ResponseAllOfAccountsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTenants200ResponseAllOfAccountsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListTenants200ResponseAllOfAccountsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListTenants200ResponseAllOfAccountsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTenants200ResponseAllOfAccountsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListTenants200ResponseAllOfAccountsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListTenants200ResponseAllOfAccountsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListTenants200ResponseAllOfAccountsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListTenants200ResponseAllOfAccountsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubdomain

`func (o *ListTenants200ResponseAllOfAccountsInner) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ListTenants200ResponseAllOfAccountsInner) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ListTenants200ResponseAllOfAccountsInner) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetCurrency

`func (o *ListTenants200ResponseAllOfAccountsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListTenants200ResponseAllOfAccountsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListTenants200ResponseAllOfAccountsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalId

`func (o *ListTenants200ResponseAllOfAccountsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListTenants200ResponseAllOfAccountsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListTenants200ResponseAllOfAccountsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCustomerNumber

`func (o *ListTenants200ResponseAllOfAccountsInner) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *ListTenants200ResponseAllOfAccountsInner) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.

### HasCustomerNumber

`func (o *ListTenants200ResponseAllOfAccountsInner) HasCustomerNumber() bool`

HasCustomerNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *ListTenants200ResponseAllOfAccountsInner) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *ListTenants200ResponseAllOfAccountsInner) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *ListTenants200ResponseAllOfAccountsInner) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountName

`func (o *ListTenants200ResponseAllOfAccountsInner) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ListTenants200ResponseAllOfAccountsInner) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *ListTenants200ResponseAllOfAccountsInner) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetActive

`func (o *ListTenants200ResponseAllOfAccountsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListTenants200ResponseAllOfAccountsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListTenants200ResponseAllOfAccountsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMaster

`func (o *ListTenants200ResponseAllOfAccountsInner) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ListTenants200ResponseAllOfAccountsInner) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ListTenants200ResponseAllOfAccountsInner) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetRole

`func (o *ListTenants200ResponseAllOfAccountsInner) GetRole() ListTenants200ResponseAllOfAccountsInnerRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetRoleOk() (*ListTenants200ResponseAllOfAccountsInnerRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ListTenants200ResponseAllOfAccountsInner) SetRole(v ListTenants200ResponseAllOfAccountsInnerRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ListTenants200ResponseAllOfAccountsInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStats

`func (o *ListTenants200ResponseAllOfAccountsInner) GetStats() ListTenants200ResponseAllOfAccountsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetStatsOk() (*ListTenants200ResponseAllOfAccountsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListTenants200ResponseAllOfAccountsInner) SetStats(v ListTenants200ResponseAllOfAccountsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListTenants200ResponseAllOfAccountsInner) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListTenants200ResponseAllOfAccountsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListTenants200ResponseAllOfAccountsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListTenants200ResponseAllOfAccountsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListTenants200ResponseAllOfAccountsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListTenants200ResponseAllOfAccountsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListTenants200ResponseAllOfAccountsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListTenants200ResponseAllOfAccountsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



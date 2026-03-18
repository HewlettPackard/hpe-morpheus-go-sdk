# AddTenant200ResponseAllOfAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**CustomerNumber** | Pointer to **NullableString** |  | [optional] 
**AccountNumber** | Pointer to **NullableString** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Master** | Pointer to **bool** |  | [optional] 
**Parent** | Pointer to [**AddTenant200ResponseAllOfAccountParent**](AddTenant200ResponseAllOfAccountParent.md) |  | [optional] 
**Role** | Pointer to [**AddTenant200ResponseAllOfAccountRole**](AddTenant200ResponseAllOfAccountRole.md) |  | [optional] 
**Stats** | Pointer to [**AddTenant200ResponseAllOfAccountStats**](AddTenant200ResponseAllOfAccountStats.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddTenant200ResponseAllOfAccount

`func NewAddTenant200ResponseAllOfAccount() *AddTenant200ResponseAllOfAccount`

NewAddTenant200ResponseAllOfAccount instantiates a new AddTenant200ResponseAllOfAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTenant200ResponseAllOfAccountWithDefaults

`func NewAddTenant200ResponseAllOfAccountWithDefaults() *AddTenant200ResponseAllOfAccount`

NewAddTenant200ResponseAllOfAccountWithDefaults instantiates a new AddTenant200ResponseAllOfAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddTenant200ResponseAllOfAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddTenant200ResponseAllOfAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddTenant200ResponseAllOfAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddTenant200ResponseAllOfAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddTenant200ResponseAllOfAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddTenant200ResponseAllOfAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddTenant200ResponseAllOfAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddTenant200ResponseAllOfAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddTenant200ResponseAllOfAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddTenant200ResponseAllOfAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddTenant200ResponseAllOfAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddTenant200ResponseAllOfAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddTenant200ResponseAllOfAccount) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddTenant200ResponseAllOfAccount) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSubdomain

`func (o *AddTenant200ResponseAllOfAccount) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *AddTenant200ResponseAllOfAccount) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *AddTenant200ResponseAllOfAccount) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *AddTenant200ResponseAllOfAccount) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetCurrency

`func (o *AddTenant200ResponseAllOfAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AddTenant200ResponseAllOfAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AddTenant200ResponseAllOfAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AddTenant200ResponseAllOfAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalId

`func (o *AddTenant200ResponseAllOfAccount) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddTenant200ResponseAllOfAccount) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddTenant200ResponseAllOfAccount) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddTenant200ResponseAllOfAccount) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AddTenant200ResponseAllOfAccount) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AddTenant200ResponseAllOfAccount) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetCustomerNumber

`func (o *AddTenant200ResponseAllOfAccount) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *AddTenant200ResponseAllOfAccount) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *AddTenant200ResponseAllOfAccount) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.

### HasCustomerNumber

`func (o *AddTenant200ResponseAllOfAccount) HasCustomerNumber() bool`

HasCustomerNumber returns a boolean if a field has been set.

### SetCustomerNumberNil

`func (o *AddTenant200ResponseAllOfAccount) SetCustomerNumberNil(b bool)`

 SetCustomerNumberNil sets the value for CustomerNumber to be an explicit nil

### UnsetCustomerNumber
`func (o *AddTenant200ResponseAllOfAccount) UnsetCustomerNumber()`

UnsetCustomerNumber ensures that no value is present for CustomerNumber, not even an explicit nil
### GetAccountNumber

`func (o *AddTenant200ResponseAllOfAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AddTenant200ResponseAllOfAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AddTenant200ResponseAllOfAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AddTenant200ResponseAllOfAccount) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *AddTenant200ResponseAllOfAccount) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *AddTenant200ResponseAllOfAccount) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetAccountName

`func (o *AddTenant200ResponseAllOfAccount) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AddTenant200ResponseAllOfAccount) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AddTenant200ResponseAllOfAccount) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *AddTenant200ResponseAllOfAccount) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetActive

`func (o *AddTenant200ResponseAllOfAccount) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddTenant200ResponseAllOfAccount) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddTenant200ResponseAllOfAccount) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddTenant200ResponseAllOfAccount) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMaster

`func (o *AddTenant200ResponseAllOfAccount) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *AddTenant200ResponseAllOfAccount) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *AddTenant200ResponseAllOfAccount) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *AddTenant200ResponseAllOfAccount) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetParent

`func (o *AddTenant200ResponseAllOfAccount) GetParent() AddTenant200ResponseAllOfAccountParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AddTenant200ResponseAllOfAccount) GetParentOk() (*AddTenant200ResponseAllOfAccountParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AddTenant200ResponseAllOfAccount) SetParent(v AddTenant200ResponseAllOfAccountParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AddTenant200ResponseAllOfAccount) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetRole

`func (o *AddTenant200ResponseAllOfAccount) GetRole() AddTenant200ResponseAllOfAccountRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddTenant200ResponseAllOfAccount) GetRoleOk() (*AddTenant200ResponseAllOfAccountRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddTenant200ResponseAllOfAccount) SetRole(v AddTenant200ResponseAllOfAccountRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *AddTenant200ResponseAllOfAccount) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStats

`func (o *AddTenant200ResponseAllOfAccount) GetStats() AddTenant200ResponseAllOfAccountStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AddTenant200ResponseAllOfAccount) GetStatsOk() (*AddTenant200ResponseAllOfAccountStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AddTenant200ResponseAllOfAccount) SetStats(v AddTenant200ResponseAllOfAccountStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *AddTenant200ResponseAllOfAccount) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddTenant200ResponseAllOfAccount) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddTenant200ResponseAllOfAccount) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddTenant200ResponseAllOfAccount) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddTenant200ResponseAllOfAccount) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddTenant200ResponseAllOfAccount) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddTenant200ResponseAllOfAccount) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddTenant200ResponseAllOfAccount) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddTenant200ResponseAllOfAccount) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



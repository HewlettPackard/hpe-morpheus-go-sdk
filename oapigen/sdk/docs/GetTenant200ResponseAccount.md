# GetTenant200ResponseAccount

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
**Parent** | Pointer to [**GetTenant200ResponseAccountParent**](GetTenant200ResponseAccountParent.md) |  | [optional] 
**Role** | Pointer to [**GetTenant200ResponseAccountRole**](GetTenant200ResponseAccountRole.md) |  | [optional] 
**Stats** | Pointer to [**GetTenant200ResponseAccountStats**](GetTenant200ResponseAccountStats.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetTenant200ResponseAccount

`func NewGetTenant200ResponseAccount() *GetTenant200ResponseAccount`

NewGetTenant200ResponseAccount instantiates a new GetTenant200ResponseAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenant200ResponseAccountWithDefaults

`func NewGetTenant200ResponseAccountWithDefaults() *GetTenant200ResponseAccount`

NewGetTenant200ResponseAccountWithDefaults instantiates a new GetTenant200ResponseAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTenant200ResponseAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTenant200ResponseAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTenant200ResponseAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetTenant200ResponseAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetTenant200ResponseAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTenant200ResponseAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTenant200ResponseAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTenant200ResponseAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetTenant200ResponseAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetTenant200ResponseAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetTenant200ResponseAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetTenant200ResponseAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetTenant200ResponseAccount) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetTenant200ResponseAccount) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSubdomain

`func (o *GetTenant200ResponseAccount) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *GetTenant200ResponseAccount) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *GetTenant200ResponseAccount) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *GetTenant200ResponseAccount) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetCurrency

`func (o *GetTenant200ResponseAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetTenant200ResponseAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetTenant200ResponseAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetTenant200ResponseAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalId

`func (o *GetTenant200ResponseAccount) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetTenant200ResponseAccount) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetTenant200ResponseAccount) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetTenant200ResponseAccount) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetTenant200ResponseAccount) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetTenant200ResponseAccount) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetCustomerNumber

`func (o *GetTenant200ResponseAccount) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *GetTenant200ResponseAccount) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *GetTenant200ResponseAccount) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.

### HasCustomerNumber

`func (o *GetTenant200ResponseAccount) HasCustomerNumber() bool`

HasCustomerNumber returns a boolean if a field has been set.

### SetCustomerNumberNil

`func (o *GetTenant200ResponseAccount) SetCustomerNumberNil(b bool)`

 SetCustomerNumberNil sets the value for CustomerNumber to be an explicit nil

### UnsetCustomerNumber
`func (o *GetTenant200ResponseAccount) UnsetCustomerNumber()`

UnsetCustomerNumber ensures that no value is present for CustomerNumber, not even an explicit nil
### GetAccountNumber

`func (o *GetTenant200ResponseAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GetTenant200ResponseAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GetTenant200ResponseAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GetTenant200ResponseAccount) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *GetTenant200ResponseAccount) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *GetTenant200ResponseAccount) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetAccountName

`func (o *GetTenant200ResponseAccount) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *GetTenant200ResponseAccount) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *GetTenant200ResponseAccount) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *GetTenant200ResponseAccount) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetActive

`func (o *GetTenant200ResponseAccount) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetTenant200ResponseAccount) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetTenant200ResponseAccount) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetTenant200ResponseAccount) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMaster

`func (o *GetTenant200ResponseAccount) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *GetTenant200ResponseAccount) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *GetTenant200ResponseAccount) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *GetTenant200ResponseAccount) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetParent

`func (o *GetTenant200ResponseAccount) GetParent() GetTenant200ResponseAccountParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetTenant200ResponseAccount) GetParentOk() (*GetTenant200ResponseAccountParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetTenant200ResponseAccount) SetParent(v GetTenant200ResponseAccountParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetTenant200ResponseAccount) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetRole

`func (o *GetTenant200ResponseAccount) GetRole() GetTenant200ResponseAccountRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetTenant200ResponseAccount) GetRoleOk() (*GetTenant200ResponseAccountRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetTenant200ResponseAccount) SetRole(v GetTenant200ResponseAccountRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetTenant200ResponseAccount) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStats

`func (o *GetTenant200ResponseAccount) GetStats() GetTenant200ResponseAccountStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetTenant200ResponseAccount) GetStatsOk() (*GetTenant200ResponseAccountStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetTenant200ResponseAccount) SetStats(v GetTenant200ResponseAccountStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetTenant200ResponseAccount) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetTenant200ResponseAccount) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetTenant200ResponseAccount) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetTenant200ResponseAccount) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetTenant200ResponseAccount) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetTenant200ResponseAccount) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetTenant200ResponseAccount) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetTenant200ResponseAccount) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetTenant200ResponseAccount) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateBudgets200ResponseAllOfBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to [**UpdateBudgets200ResponseAllOfBudgetAccount**](UpdateBudgets200ResponseAllOfBudgetAccount.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**RefScope** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**RefName** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Year** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Costs** | Pointer to **[]int64** |  | [optional] 
**IsFiscal** | Pointer to **bool** |  | [optional] 
**AverageCost** | Pointer to **int64** |  | [optional] 
**TotalCost** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Rollover** | Pointer to **bool** |  | [optional] 
**WarningLimit** | Pointer to **NullableString** |  | [optional] 
**OverLimit** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**CreatedById** | Pointer to **int64** |  | [optional] 
**CreatedByName** | Pointer to **string** |  | [optional] 
**UpdatedById** | Pointer to **NullableString** |  | [optional] 
**UpdatedByName** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateBudgets200ResponseAllOfBudget

`func NewUpdateBudgets200ResponseAllOfBudget() *UpdateBudgets200ResponseAllOfBudget`

NewUpdateBudgets200ResponseAllOfBudget instantiates a new UpdateBudgets200ResponseAllOfBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBudgets200ResponseAllOfBudgetWithDefaults

`func NewUpdateBudgets200ResponseAllOfBudgetWithDefaults() *UpdateBudgets200ResponseAllOfBudget`

NewUpdateBudgets200ResponseAllOfBudgetWithDefaults instantiates a new UpdateBudgets200ResponseAllOfBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateBudgets200ResponseAllOfBudget) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateBudgets200ResponseAllOfBudget) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateBudgets200ResponseAllOfBudget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateBudgets200ResponseAllOfBudget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBudgets200ResponseAllOfBudget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBudgets200ResponseAllOfBudget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateBudgets200ResponseAllOfBudget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateBudgets200ResponseAllOfBudget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateBudgets200ResponseAllOfBudget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateBudgets200ResponseAllOfBudget) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateBudgets200ResponseAllOfBudget) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAccount

`func (o *UpdateBudgets200ResponseAllOfBudget) GetAccount() UpdateBudgets200ResponseAllOfBudgetAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetAccountOk() (*UpdateBudgets200ResponseAllOfBudgetAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateBudgets200ResponseAllOfBudget) SetAccount(v UpdateBudgets200ResponseAllOfBudgetAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateBudgets200ResponseAllOfBudget) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateBudgets200ResponseAllOfBudget) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateBudgets200ResponseAllOfBudget) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateBudgets200ResponseAllOfBudget) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefScope

`func (o *UpdateBudgets200ResponseAllOfBudget) GetRefScope() string`

GetRefScope returns the RefScope field if non-nil, zero value otherwise.

### GetRefScopeOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetRefScopeOk() (*string, bool)`

GetRefScopeOk returns a tuple with the RefScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefScope

`func (o *UpdateBudgets200ResponseAllOfBudget) SetRefScope(v string)`

SetRefScope sets RefScope field to given value.

### HasRefScope

`func (o *UpdateBudgets200ResponseAllOfBudget) HasRefScope() bool`

HasRefScope returns a boolean if a field has been set.

### GetRefType

`func (o *UpdateBudgets200ResponseAllOfBudget) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdateBudgets200ResponseAllOfBudget) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdateBudgets200ResponseAllOfBudget) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *UpdateBudgets200ResponseAllOfBudget) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateBudgets200ResponseAllOfBudget) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateBudgets200ResponseAllOfBudget) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *UpdateBudgets200ResponseAllOfBudget) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *UpdateBudgets200ResponseAllOfBudget) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *UpdateBudgets200ResponseAllOfBudget) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetPeriod

`func (o *UpdateBudgets200ResponseAllOfBudget) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *UpdateBudgets200ResponseAllOfBudget) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *UpdateBudgets200ResponseAllOfBudget) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetYear

`func (o *UpdateBudgets200ResponseAllOfBudget) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *UpdateBudgets200ResponseAllOfBudget) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *UpdateBudgets200ResponseAllOfBudget) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetResourceType

`func (o *UpdateBudgets200ResponseAllOfBudget) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *UpdateBudgets200ResponseAllOfBudget) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *UpdateBudgets200ResponseAllOfBudget) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTimezone

`func (o *UpdateBudgets200ResponseAllOfBudget) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateBudgets200ResponseAllOfBudget) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdateBudgets200ResponseAllOfBudget) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateBudgets200ResponseAllOfBudget) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateBudgets200ResponseAllOfBudget) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateBudgets200ResponseAllOfBudget) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UpdateBudgets200ResponseAllOfBudget) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateBudgets200ResponseAllOfBudget) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateBudgets200ResponseAllOfBudget) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInterval

`func (o *UpdateBudgets200ResponseAllOfBudget) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UpdateBudgets200ResponseAllOfBudget) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UpdateBudgets200ResponseAllOfBudget) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCosts

`func (o *UpdateBudgets200ResponseAllOfBudget) GetCosts() []int64`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetCostsOk() (*[]int64, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *UpdateBudgets200ResponseAllOfBudget) SetCosts(v []int64)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *UpdateBudgets200ResponseAllOfBudget) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetIsFiscal

`func (o *UpdateBudgets200ResponseAllOfBudget) GetIsFiscal() bool`

GetIsFiscal returns the IsFiscal field if non-nil, zero value otherwise.

### GetIsFiscalOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetIsFiscalOk() (*bool, bool)`

GetIsFiscalOk returns a tuple with the IsFiscal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFiscal

`func (o *UpdateBudgets200ResponseAllOfBudget) SetIsFiscal(v bool)`

SetIsFiscal sets IsFiscal field to given value.

### HasIsFiscal

`func (o *UpdateBudgets200ResponseAllOfBudget) HasIsFiscal() bool`

HasIsFiscal returns a boolean if a field has been set.

### GetAverageCost

`func (o *UpdateBudgets200ResponseAllOfBudget) GetAverageCost() int64`

GetAverageCost returns the AverageCost field if non-nil, zero value otherwise.

### GetAverageCostOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetAverageCostOk() (*int64, bool)`

GetAverageCostOk returns a tuple with the AverageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCost

`func (o *UpdateBudgets200ResponseAllOfBudget) SetAverageCost(v int64)`

SetAverageCost sets AverageCost field to given value.

### HasAverageCost

`func (o *UpdateBudgets200ResponseAllOfBudget) HasAverageCost() bool`

HasAverageCost returns a boolean if a field has been set.

### GetTotalCost

`func (o *UpdateBudgets200ResponseAllOfBudget) GetTotalCost() int64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetTotalCostOk() (*int64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *UpdateBudgets200ResponseAllOfBudget) SetTotalCost(v int64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *UpdateBudgets200ResponseAllOfBudget) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateBudgets200ResponseAllOfBudget) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateBudgets200ResponseAllOfBudget) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateBudgets200ResponseAllOfBudget) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRollover

`func (o *UpdateBudgets200ResponseAllOfBudget) GetRollover() bool`

GetRollover returns the Rollover field if non-nil, zero value otherwise.

### GetRolloverOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetRolloverOk() (*bool, bool)`

GetRolloverOk returns a tuple with the Rollover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollover

`func (o *UpdateBudgets200ResponseAllOfBudget) SetRollover(v bool)`

SetRollover sets Rollover field to given value.

### HasRollover

`func (o *UpdateBudgets200ResponseAllOfBudget) HasRollover() bool`

HasRollover returns a boolean if a field has been set.

### GetWarningLimit

`func (o *UpdateBudgets200ResponseAllOfBudget) GetWarningLimit() string`

GetWarningLimit returns the WarningLimit field if non-nil, zero value otherwise.

### GetWarningLimitOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetWarningLimitOk() (*string, bool)`

GetWarningLimitOk returns a tuple with the WarningLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningLimit

`func (o *UpdateBudgets200ResponseAllOfBudget) SetWarningLimit(v string)`

SetWarningLimit sets WarningLimit field to given value.

### HasWarningLimit

`func (o *UpdateBudgets200ResponseAllOfBudget) HasWarningLimit() bool`

HasWarningLimit returns a boolean if a field has been set.

### SetWarningLimitNil

`func (o *UpdateBudgets200ResponseAllOfBudget) SetWarningLimitNil(b bool)`

 SetWarningLimitNil sets the value for WarningLimit to be an explicit nil

### UnsetWarningLimit
`func (o *UpdateBudgets200ResponseAllOfBudget) UnsetWarningLimit()`

UnsetWarningLimit ensures that no value is present for WarningLimit, not even an explicit nil
### GetOverLimit

`func (o *UpdateBudgets200ResponseAllOfBudget) GetOverLimit() string`

GetOverLimit returns the OverLimit field if non-nil, zero value otherwise.

### GetOverLimitOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetOverLimitOk() (*string, bool)`

GetOverLimitOk returns a tuple with the OverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverLimit

`func (o *UpdateBudgets200ResponseAllOfBudget) SetOverLimit(v string)`

SetOverLimit sets OverLimit field to given value.

### HasOverLimit

`func (o *UpdateBudgets200ResponseAllOfBudget) HasOverLimit() bool`

HasOverLimit returns a boolean if a field has been set.

### SetOverLimitNil

`func (o *UpdateBudgets200ResponseAllOfBudget) SetOverLimitNil(b bool)`

 SetOverLimitNil sets the value for OverLimit to be an explicit nil

### UnsetOverLimit
`func (o *UpdateBudgets200ResponseAllOfBudget) UnsetOverLimit()`

UnsetOverLimit ensures that no value is present for OverLimit, not even an explicit nil
### GetExternalId

`func (o *UpdateBudgets200ResponseAllOfBudget) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateBudgets200ResponseAllOfBudget) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateBudgets200ResponseAllOfBudget) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateBudgets200ResponseAllOfBudget) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateBudgets200ResponseAllOfBudget) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInternalId

`func (o *UpdateBudgets200ResponseAllOfBudget) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateBudgets200ResponseAllOfBudget) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateBudgets200ResponseAllOfBudget) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *UpdateBudgets200ResponseAllOfBudget) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *UpdateBudgets200ResponseAllOfBudget) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetCreatedById

`func (o *UpdateBudgets200ResponseAllOfBudget) GetCreatedById() int64`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetCreatedByIdOk() (*int64, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *UpdateBudgets200ResponseAllOfBudget) SetCreatedById(v int64)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *UpdateBudgets200ResponseAllOfBudget) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *UpdateBudgets200ResponseAllOfBudget) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *UpdateBudgets200ResponseAllOfBudget) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *UpdateBudgets200ResponseAllOfBudget) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetUpdatedById

`func (o *UpdateBudgets200ResponseAllOfBudget) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *UpdateBudgets200ResponseAllOfBudget) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.

### HasUpdatedById

`func (o *UpdateBudgets200ResponseAllOfBudget) HasUpdatedById() bool`

HasUpdatedById returns a boolean if a field has been set.

### SetUpdatedByIdNil

`func (o *UpdateBudgets200ResponseAllOfBudget) SetUpdatedByIdNil(b bool)`

 SetUpdatedByIdNil sets the value for UpdatedById to be an explicit nil

### UnsetUpdatedById
`func (o *UpdateBudgets200ResponseAllOfBudget) UnsetUpdatedById()`

UnsetUpdatedById ensures that no value is present for UpdatedById, not even an explicit nil
### GetUpdatedByName

`func (o *UpdateBudgets200ResponseAllOfBudget) GetUpdatedByName() string`

GetUpdatedByName returns the UpdatedByName field if non-nil, zero value otherwise.

### GetUpdatedByNameOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetUpdatedByNameOk() (*string, bool)`

GetUpdatedByNameOk returns a tuple with the UpdatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByName

`func (o *UpdateBudgets200ResponseAllOfBudget) SetUpdatedByName(v string)`

SetUpdatedByName sets UpdatedByName field to given value.

### HasUpdatedByName

`func (o *UpdateBudgets200ResponseAllOfBudget) HasUpdatedByName() bool`

HasUpdatedByName returns a boolean if a field has been set.

### SetUpdatedByNameNil

`func (o *UpdateBudgets200ResponseAllOfBudget) SetUpdatedByNameNil(b bool)`

 SetUpdatedByNameNil sets the value for UpdatedByName to be an explicit nil

### UnsetUpdatedByName
`func (o *UpdateBudgets200ResponseAllOfBudget) UnsetUpdatedByName()`

UnsetUpdatedByName ensures that no value is present for UpdatedByName, not even an explicit nil
### GetDateCreated

`func (o *UpdateBudgets200ResponseAllOfBudget) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateBudgets200ResponseAllOfBudget) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateBudgets200ResponseAllOfBudget) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateBudgets200ResponseAllOfBudget) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateBudgets200ResponseAllOfBudget) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateBudgets200ResponseAllOfBudget) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateBudgets200ResponseAllOfBudget) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Budgets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
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
**WarningLimit** | Pointer to **string** |  | [optional] 
**OverLimit** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**CreatedById** | Pointer to **int64** |  | [optional] 
**CreatedByName** | Pointer to **string** |  | [optional] 
**UpdatedById** | Pointer to **string** |  | [optional] 
**UpdatedByName** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBudgets

`func NewBudgets() *Budgets`

NewBudgets instantiates a new Budgets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetsWithDefaults

`func NewBudgetsWithDefaults() *Budgets`

NewBudgetsWithDefaults instantiates a new Budgets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Budgets) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Budgets) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Budgets) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Budgets) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Budgets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Budgets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Budgets) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Budgets) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Budgets) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Budgets) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Budgets) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Budgets) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccount

`func (o *Budgets) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Budgets) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Budgets) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Budgets) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEnabled

`func (o *Budgets) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Budgets) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Budgets) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Budgets) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefScope

`func (o *Budgets) GetRefScope() string`

GetRefScope returns the RefScope field if non-nil, zero value otherwise.

### GetRefScopeOk

`func (o *Budgets) GetRefScopeOk() (*string, bool)`

GetRefScopeOk returns a tuple with the RefScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefScope

`func (o *Budgets) SetRefScope(v string)`

SetRefScope sets RefScope field to given value.

### HasRefScope

`func (o *Budgets) HasRefScope() bool`

HasRefScope returns a boolean if a field has been set.

### GetRefType

`func (o *Budgets) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *Budgets) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *Budgets) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *Budgets) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *Budgets) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Budgets) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Budgets) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Budgets) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *Budgets) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *Budgets) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *Budgets) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *Budgets) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetPeriod

`func (o *Budgets) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Budgets) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Budgets) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *Budgets) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetYear

`func (o *Budgets) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Budgets) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Budgets) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *Budgets) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetResourceType

`func (o *Budgets) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Budgets) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Budgets) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Budgets) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTimezone

`func (o *Budgets) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Budgets) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Budgets) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Budgets) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetStartDate

`func (o *Budgets) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Budgets) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Budgets) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Budgets) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Budgets) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Budgets) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Budgets) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Budgets) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInterval

`func (o *Budgets) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Budgets) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Budgets) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Budgets) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCosts

`func (o *Budgets) GetCosts() []int64`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *Budgets) GetCostsOk() (*[]int64, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *Budgets) SetCosts(v []int64)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *Budgets) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetIsFiscal

`func (o *Budgets) GetIsFiscal() bool`

GetIsFiscal returns the IsFiscal field if non-nil, zero value otherwise.

### GetIsFiscalOk

`func (o *Budgets) GetIsFiscalOk() (*bool, bool)`

GetIsFiscalOk returns a tuple with the IsFiscal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFiscal

`func (o *Budgets) SetIsFiscal(v bool)`

SetIsFiscal sets IsFiscal field to given value.

### HasIsFiscal

`func (o *Budgets) HasIsFiscal() bool`

HasIsFiscal returns a boolean if a field has been set.

### GetAverageCost

`func (o *Budgets) GetAverageCost() int64`

GetAverageCost returns the AverageCost field if non-nil, zero value otherwise.

### GetAverageCostOk

`func (o *Budgets) GetAverageCostOk() (*int64, bool)`

GetAverageCostOk returns a tuple with the AverageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCost

`func (o *Budgets) SetAverageCost(v int64)`

SetAverageCost sets AverageCost field to given value.

### HasAverageCost

`func (o *Budgets) HasAverageCost() bool`

HasAverageCost returns a boolean if a field has been set.

### GetTotalCost

`func (o *Budgets) GetTotalCost() int64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *Budgets) GetTotalCostOk() (*int64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *Budgets) SetTotalCost(v int64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *Budgets) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetCurrency

`func (o *Budgets) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Budgets) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Budgets) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Budgets) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRollover

`func (o *Budgets) GetRollover() bool`

GetRollover returns the Rollover field if non-nil, zero value otherwise.

### GetRolloverOk

`func (o *Budgets) GetRolloverOk() (*bool, bool)`

GetRolloverOk returns a tuple with the Rollover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollover

`func (o *Budgets) SetRollover(v bool)`

SetRollover sets Rollover field to given value.

### HasRollover

`func (o *Budgets) HasRollover() bool`

HasRollover returns a boolean if a field has been set.

### GetWarningLimit

`func (o *Budgets) GetWarningLimit() string`

GetWarningLimit returns the WarningLimit field if non-nil, zero value otherwise.

### GetWarningLimitOk

`func (o *Budgets) GetWarningLimitOk() (*string, bool)`

GetWarningLimitOk returns a tuple with the WarningLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningLimit

`func (o *Budgets) SetWarningLimit(v string)`

SetWarningLimit sets WarningLimit field to given value.

### HasWarningLimit

`func (o *Budgets) HasWarningLimit() bool`

HasWarningLimit returns a boolean if a field has been set.

### GetOverLimit

`func (o *Budgets) GetOverLimit() string`

GetOverLimit returns the OverLimit field if non-nil, zero value otherwise.

### GetOverLimitOk

`func (o *Budgets) GetOverLimitOk() (*string, bool)`

GetOverLimitOk returns a tuple with the OverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverLimit

`func (o *Budgets) SetOverLimit(v string)`

SetOverLimit sets OverLimit field to given value.

### HasOverLimit

`func (o *Budgets) HasOverLimit() bool`

HasOverLimit returns a boolean if a field has been set.

### GetExternalId

`func (o *Budgets) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Budgets) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Budgets) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Budgets) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInternalId

`func (o *Budgets) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *Budgets) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *Budgets) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *Budgets) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetCreatedById

`func (o *Budgets) GetCreatedById() int64`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Budgets) GetCreatedByIdOk() (*int64, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Budgets) SetCreatedById(v int64)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *Budgets) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *Budgets) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *Budgets) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *Budgets) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *Budgets) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetUpdatedById

`func (o *Budgets) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Budgets) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Budgets) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.

### HasUpdatedById

`func (o *Budgets) HasUpdatedById() bool`

HasUpdatedById returns a boolean if a field has been set.

### GetUpdatedByName

`func (o *Budgets) GetUpdatedByName() string`

GetUpdatedByName returns the UpdatedByName field if non-nil, zero value otherwise.

### GetUpdatedByNameOk

`func (o *Budgets) GetUpdatedByNameOk() (*string, bool)`

GetUpdatedByNameOk returns a tuple with the UpdatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByName

`func (o *Budgets) SetUpdatedByName(v string)`

SetUpdatedByName sets UpdatedByName field to given value.

### HasUpdatedByName

`func (o *Budgets) HasUpdatedByName() bool`

HasUpdatedByName returns a boolean if a field has been set.

### GetDateCreated

`func (o *Budgets) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Budgets) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Budgets) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Budgets) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Budgets) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Budgets) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Budgets) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Budgets) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



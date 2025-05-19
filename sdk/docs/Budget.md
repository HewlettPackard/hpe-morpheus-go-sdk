# Budget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**ForecastType** | Pointer to [**GetBudgets200ResponseAllOfBudgetForecastType**](GetBudgets200ResponseAllOfBudgetForecastType.md) |  | [optional] 
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
**Stats** | Pointer to [**GetBudgets200ResponseAllOfBudgetStats**](GetBudgets200ResponseAllOfBudgetStats.md) |  | [optional] 

## Methods

### NewBudget

`func NewBudget() *Budget`

NewBudget instantiates a new Budget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetWithDefaults

`func NewBudgetWithDefaults() *Budget`

NewBudgetWithDefaults instantiates a new Budget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Budget) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Budget) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Budget) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Budget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Budget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Budget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Budget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Budget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Budget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Budget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Budget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Budget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccount

`func (o *Budget) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Budget) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Budget) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Budget) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetForecastType

`func (o *Budget) GetForecastType() GetBudgets200ResponseAllOfBudgetForecastType`

GetForecastType returns the ForecastType field if non-nil, zero value otherwise.

### GetForecastTypeOk

`func (o *Budget) GetForecastTypeOk() (*GetBudgets200ResponseAllOfBudgetForecastType, bool)`

GetForecastTypeOk returns a tuple with the ForecastType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastType

`func (o *Budget) SetForecastType(v GetBudgets200ResponseAllOfBudgetForecastType)`

SetForecastType sets ForecastType field to given value.

### HasForecastType

`func (o *Budget) HasForecastType() bool`

HasForecastType returns a boolean if a field has been set.

### GetEnabled

`func (o *Budget) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Budget) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Budget) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Budget) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefScope

`func (o *Budget) GetRefScope() string`

GetRefScope returns the RefScope field if non-nil, zero value otherwise.

### GetRefScopeOk

`func (o *Budget) GetRefScopeOk() (*string, bool)`

GetRefScopeOk returns a tuple with the RefScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefScope

`func (o *Budget) SetRefScope(v string)`

SetRefScope sets RefScope field to given value.

### HasRefScope

`func (o *Budget) HasRefScope() bool`

HasRefScope returns a boolean if a field has been set.

### GetRefType

`func (o *Budget) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *Budget) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *Budget) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *Budget) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *Budget) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Budget) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Budget) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Budget) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *Budget) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *Budget) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *Budget) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *Budget) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetPeriod

`func (o *Budget) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Budget) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Budget) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *Budget) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetYear

`func (o *Budget) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Budget) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Budget) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *Budget) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetResourceType

`func (o *Budget) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Budget) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Budget) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Budget) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTimezone

`func (o *Budget) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Budget) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Budget) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Budget) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetStartDate

`func (o *Budget) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Budget) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Budget) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Budget) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Budget) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Budget) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Budget) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Budget) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInterval

`func (o *Budget) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Budget) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Budget) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Budget) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCosts

`func (o *Budget) GetCosts() []int64`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *Budget) GetCostsOk() (*[]int64, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *Budget) SetCosts(v []int64)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *Budget) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetIsFiscal

`func (o *Budget) GetIsFiscal() bool`

GetIsFiscal returns the IsFiscal field if non-nil, zero value otherwise.

### GetIsFiscalOk

`func (o *Budget) GetIsFiscalOk() (*bool, bool)`

GetIsFiscalOk returns a tuple with the IsFiscal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFiscal

`func (o *Budget) SetIsFiscal(v bool)`

SetIsFiscal sets IsFiscal field to given value.

### HasIsFiscal

`func (o *Budget) HasIsFiscal() bool`

HasIsFiscal returns a boolean if a field has been set.

### GetAverageCost

`func (o *Budget) GetAverageCost() int64`

GetAverageCost returns the AverageCost field if non-nil, zero value otherwise.

### GetAverageCostOk

`func (o *Budget) GetAverageCostOk() (*int64, bool)`

GetAverageCostOk returns a tuple with the AverageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCost

`func (o *Budget) SetAverageCost(v int64)`

SetAverageCost sets AverageCost field to given value.

### HasAverageCost

`func (o *Budget) HasAverageCost() bool`

HasAverageCost returns a boolean if a field has been set.

### GetTotalCost

`func (o *Budget) GetTotalCost() int64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *Budget) GetTotalCostOk() (*int64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *Budget) SetTotalCost(v int64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *Budget) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetCurrency

`func (o *Budget) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Budget) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Budget) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Budget) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRollover

`func (o *Budget) GetRollover() bool`

GetRollover returns the Rollover field if non-nil, zero value otherwise.

### GetRolloverOk

`func (o *Budget) GetRolloverOk() (*bool, bool)`

GetRolloverOk returns a tuple with the Rollover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollover

`func (o *Budget) SetRollover(v bool)`

SetRollover sets Rollover field to given value.

### HasRollover

`func (o *Budget) HasRollover() bool`

HasRollover returns a boolean if a field has been set.

### GetWarningLimit

`func (o *Budget) GetWarningLimit() string`

GetWarningLimit returns the WarningLimit field if non-nil, zero value otherwise.

### GetWarningLimitOk

`func (o *Budget) GetWarningLimitOk() (*string, bool)`

GetWarningLimitOk returns a tuple with the WarningLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningLimit

`func (o *Budget) SetWarningLimit(v string)`

SetWarningLimit sets WarningLimit field to given value.

### HasWarningLimit

`func (o *Budget) HasWarningLimit() bool`

HasWarningLimit returns a boolean if a field has been set.

### GetOverLimit

`func (o *Budget) GetOverLimit() string`

GetOverLimit returns the OverLimit field if non-nil, zero value otherwise.

### GetOverLimitOk

`func (o *Budget) GetOverLimitOk() (*string, bool)`

GetOverLimitOk returns a tuple with the OverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverLimit

`func (o *Budget) SetOverLimit(v string)`

SetOverLimit sets OverLimit field to given value.

### HasOverLimit

`func (o *Budget) HasOverLimit() bool`

HasOverLimit returns a boolean if a field has been set.

### GetExternalId

`func (o *Budget) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Budget) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Budget) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Budget) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInternalId

`func (o *Budget) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *Budget) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *Budget) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *Budget) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetCreatedById

`func (o *Budget) GetCreatedById() int64`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Budget) GetCreatedByIdOk() (*int64, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Budget) SetCreatedById(v int64)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *Budget) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *Budget) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *Budget) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *Budget) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *Budget) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetUpdatedById

`func (o *Budget) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Budget) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Budget) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.

### HasUpdatedById

`func (o *Budget) HasUpdatedById() bool`

HasUpdatedById returns a boolean if a field has been set.

### GetUpdatedByName

`func (o *Budget) GetUpdatedByName() string`

GetUpdatedByName returns the UpdatedByName field if non-nil, zero value otherwise.

### GetUpdatedByNameOk

`func (o *Budget) GetUpdatedByNameOk() (*string, bool)`

GetUpdatedByNameOk returns a tuple with the UpdatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByName

`func (o *Budget) SetUpdatedByName(v string)`

SetUpdatedByName sets UpdatedByName field to given value.

### HasUpdatedByName

`func (o *Budget) HasUpdatedByName() bool`

HasUpdatedByName returns a boolean if a field has been set.

### GetDateCreated

`func (o *Budget) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Budget) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Budget) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Budget) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Budget) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Budget) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Budget) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Budget) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStats

`func (o *Budget) GetStats() GetBudgets200ResponseAllOfBudgetStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Budget) GetStatsOk() (*GetBudgets200ResponseAllOfBudgetStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Budget) SetStats(v GetBudgets200ResponseAllOfBudgetStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Budget) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



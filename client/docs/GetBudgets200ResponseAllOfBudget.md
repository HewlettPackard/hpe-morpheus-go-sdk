# GetBudgets200ResponseAllOfBudget

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

### NewGetBudgets200ResponseAllOfBudget

`func NewGetBudgets200ResponseAllOfBudget() *GetBudgets200ResponseAllOfBudget`

NewGetBudgets200ResponseAllOfBudget instantiates a new GetBudgets200ResponseAllOfBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBudgets200ResponseAllOfBudgetWithDefaults

`func NewGetBudgets200ResponseAllOfBudgetWithDefaults() *GetBudgets200ResponseAllOfBudget`

NewGetBudgets200ResponseAllOfBudgetWithDefaults instantiates a new GetBudgets200ResponseAllOfBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetBudgets200ResponseAllOfBudget) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBudgets200ResponseAllOfBudget) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBudgets200ResponseAllOfBudget) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetBudgets200ResponseAllOfBudget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetBudgets200ResponseAllOfBudget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetBudgets200ResponseAllOfBudget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetBudgets200ResponseAllOfBudget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetBudgets200ResponseAllOfBudget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetBudgets200ResponseAllOfBudget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetBudgets200ResponseAllOfBudget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetBudgets200ResponseAllOfBudget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetBudgets200ResponseAllOfBudget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccount

`func (o *GetBudgets200ResponseAllOfBudget) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetBudgets200ResponseAllOfBudget) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetBudgets200ResponseAllOfBudget) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetBudgets200ResponseAllOfBudget) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetForecastType

`func (o *GetBudgets200ResponseAllOfBudget) GetForecastType() GetBudgets200ResponseAllOfBudgetForecastType`

GetForecastType returns the ForecastType field if non-nil, zero value otherwise.

### GetForecastTypeOk

`func (o *GetBudgets200ResponseAllOfBudget) GetForecastTypeOk() (*GetBudgets200ResponseAllOfBudgetForecastType, bool)`

GetForecastTypeOk returns a tuple with the ForecastType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastType

`func (o *GetBudgets200ResponseAllOfBudget) SetForecastType(v GetBudgets200ResponseAllOfBudgetForecastType)`

SetForecastType sets ForecastType field to given value.

### HasForecastType

`func (o *GetBudgets200ResponseAllOfBudget) HasForecastType() bool`

HasForecastType returns a boolean if a field has been set.

### GetEnabled

`func (o *GetBudgets200ResponseAllOfBudget) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetBudgets200ResponseAllOfBudget) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetBudgets200ResponseAllOfBudget) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetBudgets200ResponseAllOfBudget) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefScope

`func (o *GetBudgets200ResponseAllOfBudget) GetRefScope() string`

GetRefScope returns the RefScope field if non-nil, zero value otherwise.

### GetRefScopeOk

`func (o *GetBudgets200ResponseAllOfBudget) GetRefScopeOk() (*string, bool)`

GetRefScopeOk returns a tuple with the RefScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefScope

`func (o *GetBudgets200ResponseAllOfBudget) SetRefScope(v string)`

SetRefScope sets RefScope field to given value.

### HasRefScope

`func (o *GetBudgets200ResponseAllOfBudget) HasRefScope() bool`

HasRefScope returns a boolean if a field has been set.

### GetRefType

`func (o *GetBudgets200ResponseAllOfBudget) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetBudgets200ResponseAllOfBudget) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetBudgets200ResponseAllOfBudget) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetBudgets200ResponseAllOfBudget) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetBudgets200ResponseAllOfBudget) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetBudgets200ResponseAllOfBudget) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetBudgets200ResponseAllOfBudget) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetBudgets200ResponseAllOfBudget) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *GetBudgets200ResponseAllOfBudget) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *GetBudgets200ResponseAllOfBudget) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *GetBudgets200ResponseAllOfBudget) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *GetBudgets200ResponseAllOfBudget) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetPeriod

`func (o *GetBudgets200ResponseAllOfBudget) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetBudgets200ResponseAllOfBudget) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetBudgets200ResponseAllOfBudget) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GetBudgets200ResponseAllOfBudget) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetYear

`func (o *GetBudgets200ResponseAllOfBudget) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *GetBudgets200ResponseAllOfBudget) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *GetBudgets200ResponseAllOfBudget) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *GetBudgets200ResponseAllOfBudget) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetResourceType

`func (o *GetBudgets200ResponseAllOfBudget) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *GetBudgets200ResponseAllOfBudget) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *GetBudgets200ResponseAllOfBudget) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *GetBudgets200ResponseAllOfBudget) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTimezone

`func (o *GetBudgets200ResponseAllOfBudget) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetBudgets200ResponseAllOfBudget) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetBudgets200ResponseAllOfBudget) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetBudgets200ResponseAllOfBudget) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetStartDate

`func (o *GetBudgets200ResponseAllOfBudget) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetBudgets200ResponseAllOfBudget) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetBudgets200ResponseAllOfBudget) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetBudgets200ResponseAllOfBudget) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetBudgets200ResponseAllOfBudget) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetBudgets200ResponseAllOfBudget) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetBudgets200ResponseAllOfBudget) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetBudgets200ResponseAllOfBudget) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInterval

`func (o *GetBudgets200ResponseAllOfBudget) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetBudgets200ResponseAllOfBudget) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetBudgets200ResponseAllOfBudget) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GetBudgets200ResponseAllOfBudget) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCosts

`func (o *GetBudgets200ResponseAllOfBudget) GetCosts() []int64`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *GetBudgets200ResponseAllOfBudget) GetCostsOk() (*[]int64, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *GetBudgets200ResponseAllOfBudget) SetCosts(v []int64)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *GetBudgets200ResponseAllOfBudget) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetIsFiscal

`func (o *GetBudgets200ResponseAllOfBudget) GetIsFiscal() bool`

GetIsFiscal returns the IsFiscal field if non-nil, zero value otherwise.

### GetIsFiscalOk

`func (o *GetBudgets200ResponseAllOfBudget) GetIsFiscalOk() (*bool, bool)`

GetIsFiscalOk returns a tuple with the IsFiscal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFiscal

`func (o *GetBudgets200ResponseAllOfBudget) SetIsFiscal(v bool)`

SetIsFiscal sets IsFiscal field to given value.

### HasIsFiscal

`func (o *GetBudgets200ResponseAllOfBudget) HasIsFiscal() bool`

HasIsFiscal returns a boolean if a field has been set.

### GetAverageCost

`func (o *GetBudgets200ResponseAllOfBudget) GetAverageCost() int64`

GetAverageCost returns the AverageCost field if non-nil, zero value otherwise.

### GetAverageCostOk

`func (o *GetBudgets200ResponseAllOfBudget) GetAverageCostOk() (*int64, bool)`

GetAverageCostOk returns a tuple with the AverageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCost

`func (o *GetBudgets200ResponseAllOfBudget) SetAverageCost(v int64)`

SetAverageCost sets AverageCost field to given value.

### HasAverageCost

`func (o *GetBudgets200ResponseAllOfBudget) HasAverageCost() bool`

HasAverageCost returns a boolean if a field has been set.

### GetTotalCost

`func (o *GetBudgets200ResponseAllOfBudget) GetTotalCost() int64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *GetBudgets200ResponseAllOfBudget) GetTotalCostOk() (*int64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *GetBudgets200ResponseAllOfBudget) SetTotalCost(v int64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *GetBudgets200ResponseAllOfBudget) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetCurrency

`func (o *GetBudgets200ResponseAllOfBudget) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBudgets200ResponseAllOfBudget) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBudgets200ResponseAllOfBudget) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetBudgets200ResponseAllOfBudget) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRollover

`func (o *GetBudgets200ResponseAllOfBudget) GetRollover() bool`

GetRollover returns the Rollover field if non-nil, zero value otherwise.

### GetRolloverOk

`func (o *GetBudgets200ResponseAllOfBudget) GetRolloverOk() (*bool, bool)`

GetRolloverOk returns a tuple with the Rollover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollover

`func (o *GetBudgets200ResponseAllOfBudget) SetRollover(v bool)`

SetRollover sets Rollover field to given value.

### HasRollover

`func (o *GetBudgets200ResponseAllOfBudget) HasRollover() bool`

HasRollover returns a boolean if a field has been set.

### GetWarningLimit

`func (o *GetBudgets200ResponseAllOfBudget) GetWarningLimit() string`

GetWarningLimit returns the WarningLimit field if non-nil, zero value otherwise.

### GetWarningLimitOk

`func (o *GetBudgets200ResponseAllOfBudget) GetWarningLimitOk() (*string, bool)`

GetWarningLimitOk returns a tuple with the WarningLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningLimit

`func (o *GetBudgets200ResponseAllOfBudget) SetWarningLimit(v string)`

SetWarningLimit sets WarningLimit field to given value.

### HasWarningLimit

`func (o *GetBudgets200ResponseAllOfBudget) HasWarningLimit() bool`

HasWarningLimit returns a boolean if a field has been set.

### GetOverLimit

`func (o *GetBudgets200ResponseAllOfBudget) GetOverLimit() string`

GetOverLimit returns the OverLimit field if non-nil, zero value otherwise.

### GetOverLimitOk

`func (o *GetBudgets200ResponseAllOfBudget) GetOverLimitOk() (*string, bool)`

GetOverLimitOk returns a tuple with the OverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverLimit

`func (o *GetBudgets200ResponseAllOfBudget) SetOverLimit(v string)`

SetOverLimit sets OverLimit field to given value.

### HasOverLimit

`func (o *GetBudgets200ResponseAllOfBudget) HasOverLimit() bool`

HasOverLimit returns a boolean if a field has been set.

### GetExternalId

`func (o *GetBudgets200ResponseAllOfBudget) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetBudgets200ResponseAllOfBudget) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetBudgets200ResponseAllOfBudget) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetBudgets200ResponseAllOfBudget) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInternalId

`func (o *GetBudgets200ResponseAllOfBudget) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetBudgets200ResponseAllOfBudget) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetBudgets200ResponseAllOfBudget) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetBudgets200ResponseAllOfBudget) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetCreatedById

`func (o *GetBudgets200ResponseAllOfBudget) GetCreatedById() int64`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *GetBudgets200ResponseAllOfBudget) GetCreatedByIdOk() (*int64, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *GetBudgets200ResponseAllOfBudget) SetCreatedById(v int64)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *GetBudgets200ResponseAllOfBudget) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *GetBudgets200ResponseAllOfBudget) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *GetBudgets200ResponseAllOfBudget) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *GetBudgets200ResponseAllOfBudget) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *GetBudgets200ResponseAllOfBudget) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetUpdatedById

`func (o *GetBudgets200ResponseAllOfBudget) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *GetBudgets200ResponseAllOfBudget) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *GetBudgets200ResponseAllOfBudget) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.

### HasUpdatedById

`func (o *GetBudgets200ResponseAllOfBudget) HasUpdatedById() bool`

HasUpdatedById returns a boolean if a field has been set.

### GetUpdatedByName

`func (o *GetBudgets200ResponseAllOfBudget) GetUpdatedByName() string`

GetUpdatedByName returns the UpdatedByName field if non-nil, zero value otherwise.

### GetUpdatedByNameOk

`func (o *GetBudgets200ResponseAllOfBudget) GetUpdatedByNameOk() (*string, bool)`

GetUpdatedByNameOk returns a tuple with the UpdatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByName

`func (o *GetBudgets200ResponseAllOfBudget) SetUpdatedByName(v string)`

SetUpdatedByName sets UpdatedByName field to given value.

### HasUpdatedByName

`func (o *GetBudgets200ResponseAllOfBudget) HasUpdatedByName() bool`

HasUpdatedByName returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetBudgets200ResponseAllOfBudget) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetBudgets200ResponseAllOfBudget) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetBudgets200ResponseAllOfBudget) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetBudgets200ResponseAllOfBudget) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetBudgets200ResponseAllOfBudget) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetBudgets200ResponseAllOfBudget) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetBudgets200ResponseAllOfBudget) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetBudgets200ResponseAllOfBudget) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStats

`func (o *GetBudgets200ResponseAllOfBudget) GetStats() GetBudgets200ResponseAllOfBudgetStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetBudgets200ResponseAllOfBudget) GetStatsOk() (*GetBudgets200ResponseAllOfBudgetStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetBudgets200ResponseAllOfBudget) SetStats(v GetBudgets200ResponseAllOfBudgetStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetBudgets200ResponseAllOfBudget) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



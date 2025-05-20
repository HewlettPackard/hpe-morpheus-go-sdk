# ListBudgets200ResponseAllOfBudgetsInner

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

### NewListBudgets200ResponseAllOfBudgetsInner

`func NewListBudgets200ResponseAllOfBudgetsInner() *ListBudgets200ResponseAllOfBudgetsInner`

NewListBudgets200ResponseAllOfBudgetsInner instantiates a new ListBudgets200ResponseAllOfBudgetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBudgets200ResponseAllOfBudgetsInnerWithDefaults

`func NewListBudgets200ResponseAllOfBudgetsInnerWithDefaults() *ListBudgets200ResponseAllOfBudgetsInner`

NewListBudgets200ResponseAllOfBudgetsInnerWithDefaults instantiates a new ListBudgets200ResponseAllOfBudgetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccount

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEnabled

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefScope

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetRefScope() string`

GetRefScope returns the RefScope field if non-nil, zero value otherwise.

### GetRefScopeOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetRefScopeOk() (*string, bool)`

GetRefScopeOk returns a tuple with the RefScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefScope

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetRefScope(v string)`

SetRefScope sets RefScope field to given value.

### HasRefScope

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasRefScope() bool`

HasRefScope returns a boolean if a field has been set.

### GetRefType

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetPeriod

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetYear

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetResourceType

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTimezone

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetStartDate

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInterval

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCosts

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetCosts() []int64`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetCostsOk() (*[]int64, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetCosts(v []int64)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetIsFiscal

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetIsFiscal() bool`

GetIsFiscal returns the IsFiscal field if non-nil, zero value otherwise.

### GetIsFiscalOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetIsFiscalOk() (*bool, bool)`

GetIsFiscalOk returns a tuple with the IsFiscal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFiscal

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetIsFiscal(v bool)`

SetIsFiscal sets IsFiscal field to given value.

### HasIsFiscal

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasIsFiscal() bool`

HasIsFiscal returns a boolean if a field has been set.

### GetAverageCost

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetAverageCost() int64`

GetAverageCost returns the AverageCost field if non-nil, zero value otherwise.

### GetAverageCostOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetAverageCostOk() (*int64, bool)`

GetAverageCostOk returns a tuple with the AverageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCost

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetAverageCost(v int64)`

SetAverageCost sets AverageCost field to given value.

### HasAverageCost

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasAverageCost() bool`

HasAverageCost returns a boolean if a field has been set.

### GetTotalCost

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetTotalCost() int64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetTotalCostOk() (*int64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetTotalCost(v int64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetCurrency

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRollover

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetRollover() bool`

GetRollover returns the Rollover field if non-nil, zero value otherwise.

### GetRolloverOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetRolloverOk() (*bool, bool)`

GetRolloverOk returns a tuple with the Rollover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollover

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetRollover(v bool)`

SetRollover sets Rollover field to given value.

### HasRollover

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasRollover() bool`

HasRollover returns a boolean if a field has been set.

### GetWarningLimit

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetWarningLimit() string`

GetWarningLimit returns the WarningLimit field if non-nil, zero value otherwise.

### GetWarningLimitOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetWarningLimitOk() (*string, bool)`

GetWarningLimitOk returns a tuple with the WarningLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningLimit

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetWarningLimit(v string)`

SetWarningLimit sets WarningLimit field to given value.

### HasWarningLimit

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasWarningLimit() bool`

HasWarningLimit returns a boolean if a field has been set.

### GetOverLimit

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetOverLimit() string`

GetOverLimit returns the OverLimit field if non-nil, zero value otherwise.

### GetOverLimitOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetOverLimitOk() (*string, bool)`

GetOverLimitOk returns a tuple with the OverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverLimit

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetOverLimit(v string)`

SetOverLimit sets OverLimit field to given value.

### HasOverLimit

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasOverLimit() bool`

HasOverLimit returns a boolean if a field has been set.

### GetExternalId

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInternalId

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetCreatedById

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetCreatedById() int64`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetCreatedByIdOk() (*int64, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetCreatedById(v int64)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedByName

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetUpdatedById

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.

### HasUpdatedById

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasUpdatedById() bool`

HasUpdatedById returns a boolean if a field has been set.

### GetUpdatedByName

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetUpdatedByName() string`

GetUpdatedByName returns the UpdatedByName field if non-nil, zero value otherwise.

### GetUpdatedByNameOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetUpdatedByNameOk() (*string, bool)`

GetUpdatedByNameOk returns a tuple with the UpdatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByName

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetUpdatedByName(v string)`

SetUpdatedByName sets UpdatedByName field to given value.

### HasUpdatedByName

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasUpdatedByName() bool`

HasUpdatedByName returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListBudgets200ResponseAllOfBudgetsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListBudgets200ResponseAllOfBudgetsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListBudgets200ResponseAllOfBudgetsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



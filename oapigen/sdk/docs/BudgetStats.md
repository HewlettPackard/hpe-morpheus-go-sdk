# BudgetStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**ConversionRate** | Pointer to **int64** |  | [optional] 
**Intervals** | Pointer to [**[]BudgetStatsIntervalsInner**](BudgetStatsIntervalsInner.md) |  | [optional] 
**Current** | Pointer to [**BudgetStatsCurrent**](BudgetStatsCurrent.md) |  | [optional] 

## Methods

### NewBudgetStats

`func NewBudgetStats() *BudgetStats`

NewBudgetStats instantiates a new BudgetStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetStatsWithDefaults

`func NewBudgetStatsWithDefaults() *BudgetStats`

NewBudgetStatsWithDefaults instantiates a new BudgetStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *BudgetStats) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BudgetStats) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BudgetStats) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BudgetStats) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetConversionRate

`func (o *BudgetStats) GetConversionRate() int64`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *BudgetStats) GetConversionRateOk() (*int64, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *BudgetStats) SetConversionRate(v int64)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *BudgetStats) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetIntervals

`func (o *BudgetStats) GetIntervals() []BudgetStatsIntervalsInner`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *BudgetStats) GetIntervalsOk() (*[]BudgetStatsIntervalsInner, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *BudgetStats) SetIntervals(v []BudgetStatsIntervalsInner)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *BudgetStats) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetCurrent

`func (o *BudgetStats) GetCurrent() BudgetStatsCurrent`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *BudgetStats) GetCurrentOk() (*BudgetStatsCurrent, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *BudgetStats) SetCurrent(v BudgetStatsCurrent)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *BudgetStats) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



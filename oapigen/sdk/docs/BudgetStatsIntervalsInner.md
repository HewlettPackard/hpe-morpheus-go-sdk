# BudgetStatsIntervalsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int64** |  | [optional] 
**Month** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**ChartName** | Pointer to **string** |  | [optional] 
**Budget** | Pointer to **int64** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBudgetStatsIntervalsInner

`func NewBudgetStatsIntervalsInner() *BudgetStatsIntervalsInner`

NewBudgetStatsIntervalsInner instantiates a new BudgetStatsIntervalsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetStatsIntervalsInnerWithDefaults

`func NewBudgetStatsIntervalsInnerWithDefaults() *BudgetStatsIntervalsInner`

NewBudgetStatsIntervalsInnerWithDefaults instantiates a new BudgetStatsIntervalsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *BudgetStatsIntervalsInner) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BudgetStatsIntervalsInner) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BudgetStatsIntervalsInner) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *BudgetStatsIntervalsInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetMonth

`func (o *BudgetStatsIntervalsInner) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *BudgetStatsIntervalsInner) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *BudgetStatsIntervalsInner) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *BudgetStatsIntervalsInner) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetShortName

`func (o *BudgetStatsIntervalsInner) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *BudgetStatsIntervalsInner) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *BudgetStatsIntervalsInner) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *BudgetStatsIntervalsInner) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetChartName

`func (o *BudgetStatsIntervalsInner) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *BudgetStatsIntervalsInner) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *BudgetStatsIntervalsInner) SetChartName(v string)`

SetChartName sets ChartName field to given value.

### HasChartName

`func (o *BudgetStatsIntervalsInner) HasChartName() bool`

HasChartName returns a boolean if a field has been set.

### GetBudget

`func (o *BudgetStatsIntervalsInner) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *BudgetStatsIntervalsInner) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *BudgetStatsIntervalsInner) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *BudgetStatsIntervalsInner) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCost

`func (o *BudgetStatsIntervalsInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BudgetStatsIntervalsInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BudgetStatsIntervalsInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BudgetStatsIntervalsInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetStartDate

`func (o *BudgetStatsIntervalsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BudgetStatsIntervalsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BudgetStatsIntervalsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BudgetStatsIntervalsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BudgetStatsIntervalsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BudgetStatsIntervalsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BudgetStatsIntervalsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BudgetStatsIntervalsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



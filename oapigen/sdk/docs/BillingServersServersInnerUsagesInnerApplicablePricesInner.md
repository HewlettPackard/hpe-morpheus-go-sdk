# BillingServersServersInnerUsagesInnerApplicablePricesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**NumUnits** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Prices** | Pointer to [**[]BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner**](BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner.md) |  | [optional] 

## Methods

### NewBillingServersServersInnerUsagesInnerApplicablePricesInner

`func NewBillingServersServersInnerUsagesInnerApplicablePricesInner() *BillingServersServersInnerUsagesInnerApplicablePricesInner`

NewBillingServersServersInnerUsagesInnerApplicablePricesInner instantiates a new BillingServersServersInnerUsagesInnerApplicablePricesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingServersServersInnerUsagesInnerApplicablePricesInnerWithDefaults

`func NewBillingServersServersInnerUsagesInnerApplicablePricesInnerWithDefaults() *BillingServersServersInnerUsagesInnerApplicablePricesInner`

NewBillingServersServersInnerUsagesInnerApplicablePricesInnerWithDefaults instantiates a new BillingServersServersInnerUsagesInnerApplicablePricesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNumUnits

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetNumUnits() float32`

GetNumUnits returns the NumUnits field if non-nil, zero value otherwise.

### GetNumUnitsOk

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetNumUnitsOk() (*float32, bool)`

GetNumUnitsOk returns a tuple with the NumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnits

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) SetNumUnits(v float32)`

SetNumUnits sets NumUnits field to given value.

### HasNumUnits

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) HasNumUnits() bool`

HasNumUnits returns a boolean if a field has been set.

### GetCost

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPrices

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetPrices() []BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) GetPricesOk() (*[]BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) SetPrices(v []BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *BillingServersServersInnerUsagesInnerApplicablePricesInner) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



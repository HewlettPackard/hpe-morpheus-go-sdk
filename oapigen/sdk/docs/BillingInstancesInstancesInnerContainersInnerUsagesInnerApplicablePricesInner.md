# BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**NumUnits** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Prices** | Pointer to [**[]BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner**](BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner.md) |  | [optional] 

## Methods

### NewBillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner

`func NewBillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner() *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner`

NewBillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner instantiates a new BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerWithDefaults

`func NewBillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerWithDefaults() *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner`

NewBillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerWithDefaults instantiates a new BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNumUnits

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetNumUnits() float32`

GetNumUnits returns the NumUnits field if non-nil, zero value otherwise.

### GetNumUnitsOk

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetNumUnitsOk() (*float32, bool)`

GetNumUnitsOk returns a tuple with the NumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnits

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetNumUnits(v float32)`

SetNumUnits sets NumUnits field to given value.

### HasNumUnits

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasNumUnits() bool`

HasNumUnits returns a boolean if a field has been set.

### GetCost

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPrices

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetPrices() []BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetPricesOk() (*[]BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetPrices(v []BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



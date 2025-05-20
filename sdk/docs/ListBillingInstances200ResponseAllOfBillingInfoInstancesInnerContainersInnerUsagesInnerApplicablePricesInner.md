# ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**NumUnits** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Prices** | Pointer to [**[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner**](ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner.md) |  | [optional] 

## Methods

### NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner

`func NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner() *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner`

NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner instantiates a new ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerWithDefaults

`func NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerWithDefaults() *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner`

NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerWithDefaults instantiates a new ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNumUnits

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetNumUnits() float32`

GetNumUnits returns the NumUnits field if non-nil, zero value otherwise.

### GetNumUnitsOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetNumUnitsOk() (*float32, bool)`

GetNumUnitsOk returns a tuple with the NumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnits

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetNumUnits(v float32)`

SetNumUnits sets NumUnits field to given value.

### HasNumUnits

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasNumUnits() bool`

HasNumUnits returns a boolean if a field has been set.

### GetCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPrices

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetPrices() []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) GetPricesOk() (*[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) SetPrices(v []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



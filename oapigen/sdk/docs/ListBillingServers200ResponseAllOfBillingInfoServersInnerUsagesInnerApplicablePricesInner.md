# ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**NumUnits** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Prices** | Pointer to [**[]ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInnerPricesInner**](ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInnerPricesInner.md) |  | [optional] 

## Methods

### NewListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner

`func NewListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner() *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner`

NewListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner instantiates a new ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInnerWithDefaults

`func NewListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInnerWithDefaults() *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner`

NewListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInnerWithDefaults instantiates a new ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNumUnits

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetNumUnits() float32`

GetNumUnits returns the NumUnits field if non-nil, zero value otherwise.

### GetNumUnitsOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetNumUnitsOk() (*float32, bool)`

GetNumUnitsOk returns a tuple with the NumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnits

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) SetNumUnits(v float32)`

SetNumUnits sets NumUnits field to given value.

### HasNumUnits

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) HasNumUnits() bool`

HasNumUnits returns a boolean if a field has been set.

### GetCost

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPrices

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetPrices() []ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInnerPricesInner`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) GetPricesOk() (*[]ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInnerPricesInner, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) SetPrices(v []ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInnerPricesInner)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



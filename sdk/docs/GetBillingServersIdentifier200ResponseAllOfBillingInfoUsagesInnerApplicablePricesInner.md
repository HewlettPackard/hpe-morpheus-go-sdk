# GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**NumUnits** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Prices** | Pointer to [**[]GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInnerPricesInner**](GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInnerPricesInner.md) |  | [optional] 

## Methods

### NewGetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner

`func NewGetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner() *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner`

NewGetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner instantiates a new GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInnerWithDefaults

`func NewGetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInnerWithDefaults() *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner`

NewGetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInnerWithDefaults instantiates a new GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNumUnits

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetNumUnits() float32`

GetNumUnits returns the NumUnits field if non-nil, zero value otherwise.

### GetNumUnitsOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetNumUnitsOk() (*float32, bool)`

GetNumUnitsOk returns a tuple with the NumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnits

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) SetNumUnits(v float32)`

SetNumUnits sets NumUnits field to given value.

### HasNumUnits

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) HasNumUnits() bool`

HasNumUnits returns a boolean if a field has been set.

### GetCost

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPrices

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetPrices() []GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInnerPricesInner`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) GetPricesOk() (*[]GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInnerPricesInner, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) SetPrices(v []GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInnerPricesInner)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



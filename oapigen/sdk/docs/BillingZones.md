# BillingZones

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Zones** | Pointer to [**[]ListBillingAccount200ResponseAllOfBillingInfoZonesInner**](ListBillingAccount200ResponseAllOfBillingInfoZonesInner.md) |  | [optional] 

## Methods

### NewBillingZones

`func NewBillingZones() *BillingZones`

NewBillingZones instantiates a new BillingZones object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingZonesWithDefaults

`func NewBillingZonesWithDefaults() *BillingZones`

NewBillingZonesWithDefaults instantiates a new BillingZones object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *BillingZones) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingZones) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingZones) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingZones) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *BillingZones) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingZones) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingZones) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingZones) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetStartDate

`func (o *BillingZones) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingZones) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingZones) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingZones) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingZones) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingZones) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingZones) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingZones) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetZones

`func (o *BillingZones) GetZones() []ListBillingAccount200ResponseAllOfBillingInfoZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *BillingZones) GetZonesOk() (*[]ListBillingAccount200ResponseAllOfBillingInfoZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *BillingZones) SetZones(v []ListBillingAccount200ResponseAllOfBillingInfoZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *BillingZones) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BillingServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Servers** | Pointer to [**[]ListBillingServers200ResponseAllOfBillingInfoServersInner**](ListBillingServers200ResponseAllOfBillingInfoServersInner.md) |  | [optional] 

## Methods

### NewBillingServers

`func NewBillingServers() *BillingServers`

NewBillingServers instantiates a new BillingServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingServersWithDefaults

`func NewBillingServersWithDefaults() *BillingServers`

NewBillingServersWithDefaults instantiates a new BillingServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *BillingServers) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingServers) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingServers) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingServers) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *BillingServers) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingServers) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingServers) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingServers) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetStartDate

`func (o *BillingServers) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingServers) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingServers) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingServers) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingServers) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingServers) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingServers) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingServers) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetServers

`func (o *BillingServers) GetServers() []ListBillingServers200ResponseAllOfBillingInfoServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *BillingServers) GetServersOk() (*[]ListBillingServers200ResponseAllOfBillingInfoServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *BillingServers) SetServers(v []ListBillingServers200ResponseAllOfBillingInfoServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *BillingServers) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



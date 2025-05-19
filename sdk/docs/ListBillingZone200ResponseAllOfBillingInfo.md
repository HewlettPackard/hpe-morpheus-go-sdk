# ListBillingZone200ResponseAllOfBillingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Zones** | Pointer to [**[]ListBillingAccount200ResponseAllOfBillingInfoZonesInner**](ListBillingAccount200ResponseAllOfBillingInfoZonesInner.md) |  | [optional] 

## Methods

### NewListBillingZone200ResponseAllOfBillingInfo

`func NewListBillingZone200ResponseAllOfBillingInfo() *ListBillingZone200ResponseAllOfBillingInfo`

NewListBillingZone200ResponseAllOfBillingInfo instantiates a new ListBillingZone200ResponseAllOfBillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingZone200ResponseAllOfBillingInfoWithDefaults

`func NewListBillingZone200ResponseAllOfBillingInfoWithDefaults() *ListBillingZone200ResponseAllOfBillingInfo`

NewListBillingZone200ResponseAllOfBillingInfoWithDefaults instantiates a new ListBillingZone200ResponseAllOfBillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *ListBillingZone200ResponseAllOfBillingInfo) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListBillingZone200ResponseAllOfBillingInfo) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListBillingZone200ResponseAllOfBillingInfo) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListBillingZone200ResponseAllOfBillingInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *ListBillingZone200ResponseAllOfBillingInfo) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListBillingZone200ResponseAllOfBillingInfo) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListBillingZone200ResponseAllOfBillingInfo) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListBillingZone200ResponseAllOfBillingInfo) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetStartDate

`func (o *ListBillingZone200ResponseAllOfBillingInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBillingZone200ResponseAllOfBillingInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBillingZone200ResponseAllOfBillingInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBillingZone200ResponseAllOfBillingInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBillingZone200ResponseAllOfBillingInfo) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBillingZone200ResponseAllOfBillingInfo) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBillingZone200ResponseAllOfBillingInfo) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBillingZone200ResponseAllOfBillingInfo) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetZones

`func (o *ListBillingZone200ResponseAllOfBillingInfo) GetZones() []ListBillingAccount200ResponseAllOfBillingInfoZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ListBillingZone200ResponseAllOfBillingInfo) GetZonesOk() (*[]ListBillingAccount200ResponseAllOfBillingInfoZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ListBillingZone200ResponseAllOfBillingInfo) SetZones(v []ListBillingAccount200ResponseAllOfBillingInfoZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ListBillingZone200ResponseAllOfBillingInfo) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



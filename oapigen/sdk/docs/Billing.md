# Billing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int64** |  | [optional] 
**AccountUUID** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**PriceUnit** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Zones** | Pointer to [**[]ListBillingAccount200ResponseAllOfBillingInfoZonesInner**](ListBillingAccount200ResponseAllOfBillingInfoZonesInner.md) |  | [optional] 

## Methods

### NewBilling

`func NewBilling() *Billing`

NewBilling instantiates a new Billing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingWithDefaults

`func NewBillingWithDefaults() *Billing`

NewBillingWithDefaults instantiates a new Billing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Billing) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Billing) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Billing) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Billing) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountUUID

`func (o *Billing) GetAccountUUID() string`

GetAccountUUID returns the AccountUUID field if non-nil, zero value otherwise.

### GetAccountUUIDOk

`func (o *Billing) GetAccountUUIDOk() (*string, bool)`

GetAccountUUIDOk returns a tuple with the AccountUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUUID

`func (o *Billing) SetAccountUUID(v string)`

SetAccountUUID sets AccountUUID field to given value.

### HasAccountUUID

`func (o *Billing) HasAccountUUID() bool`

HasAccountUUID returns a boolean if a field has been set.

### GetName

`func (o *Billing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Billing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Billing) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Billing) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartDate

`func (o *Billing) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Billing) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Billing) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Billing) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Billing) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Billing) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Billing) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Billing) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriceUnit

`func (o *Billing) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *Billing) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *Billing) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *Billing) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetPrice

`func (o *Billing) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Billing) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Billing) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Billing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *Billing) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Billing) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Billing) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Billing) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetZones

`func (o *Billing) GetZones() []ListBillingAccount200ResponseAllOfBillingInfoZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *Billing) GetZonesOk() (*[]ListBillingAccount200ResponseAllOfBillingInfoZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *Billing) SetZones(v []ListBillingAccount200ResponseAllOfBillingInfoZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *Billing) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



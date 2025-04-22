# ListBillingAccount200ResponseAllOfBillingInfo

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

### NewListBillingAccount200ResponseAllOfBillingInfo

`func NewListBillingAccount200ResponseAllOfBillingInfo() *ListBillingAccount200ResponseAllOfBillingInfo`

NewListBillingAccount200ResponseAllOfBillingInfo instantiates a new ListBillingAccount200ResponseAllOfBillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingAccount200ResponseAllOfBillingInfoWithDefaults

`func NewListBillingAccount200ResponseAllOfBillingInfoWithDefaults() *ListBillingAccount200ResponseAllOfBillingInfo`

NewListBillingAccount200ResponseAllOfBillingInfoWithDefaults instantiates a new ListBillingAccount200ResponseAllOfBillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountUUID

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetAccountUUID() string`

GetAccountUUID returns the AccountUUID field if non-nil, zero value otherwise.

### GetAccountUUIDOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetAccountUUIDOk() (*string, bool)`

GetAccountUUIDOk returns a tuple with the AccountUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUUID

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) SetAccountUUID(v string)`

SetAccountUUID sets AccountUUID field to given value.

### HasAccountUUID

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) HasAccountUUID() bool`

HasAccountUUID returns a boolean if a field has been set.

### GetName

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartDate

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriceUnit

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetPrice

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetZones

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetZones() []ListBillingAccount200ResponseAllOfBillingInfoZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) GetZonesOk() (*[]ListBillingAccount200ResponseAllOfBillingInfoZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) SetZones(v []ListBillingAccount200ResponseAllOfBillingInfoZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ListBillingAccount200ResponseAllOfBillingInfo) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



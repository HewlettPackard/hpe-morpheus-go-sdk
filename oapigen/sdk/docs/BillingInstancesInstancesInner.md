# BillingInstancesInstancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | Pointer to **int64** |  | [optional] 
**InstanceUUID** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Containers** | Pointer to [**[]BillingInstancesInstancesInnerContainersInner**](BillingInstancesInstancesInnerContainersInner.md) |  | [optional] 

## Methods

### NewBillingInstancesInstancesInner

`func NewBillingInstancesInstancesInner() *BillingInstancesInstancesInner`

NewBillingInstancesInstancesInner instantiates a new BillingInstancesInstancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInstancesInstancesInnerWithDefaults

`func NewBillingInstancesInstancesInnerWithDefaults() *BillingInstancesInstancesInner`

NewBillingInstancesInstancesInnerWithDefaults instantiates a new BillingInstancesInstancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *BillingInstancesInstancesInner) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *BillingInstancesInstancesInner) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *BillingInstancesInstancesInner) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *BillingInstancesInstancesInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceUUID

`func (o *BillingInstancesInstancesInner) GetInstanceUUID() string`

GetInstanceUUID returns the InstanceUUID field if non-nil, zero value otherwise.

### GetInstanceUUIDOk

`func (o *BillingInstancesInstancesInner) GetInstanceUUIDOk() (*string, bool)`

GetInstanceUUIDOk returns a tuple with the InstanceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUUID

`func (o *BillingInstancesInstancesInner) SetInstanceUUID(v string)`

SetInstanceUUID sets InstanceUUID field to given value.

### HasInstanceUUID

`func (o *BillingInstancesInstancesInner) HasInstanceUUID() bool`

HasInstanceUUID returns a boolean if a field has been set.

### GetStartDate

`func (o *BillingInstancesInstancesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingInstancesInstancesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingInstancesInstancesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingInstancesInstancesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingInstancesInstancesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingInstancesInstancesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingInstancesInstancesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingInstancesInstancesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetName

`func (o *BillingInstancesInstancesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingInstancesInstancesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingInstancesInstancesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingInstancesInstancesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *BillingInstancesInstancesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingInstancesInstancesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingInstancesInstancesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingInstancesInstancesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *BillingInstancesInstancesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingInstancesInstancesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingInstancesInstancesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingInstancesInstancesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingInstancesInstancesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingInstancesInstancesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingInstancesInstancesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingInstancesInstancesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetContainers

`func (o *BillingInstancesInstancesInner) GetContainers() []BillingInstancesInstancesInnerContainersInner`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *BillingInstancesInstancesInner) GetContainersOk() (*[]BillingInstancesInstancesInnerContainersInner, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *BillingInstancesInstancesInner) SetContainers(v []BillingInstancesInstancesInnerContainersInner)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *BillingInstancesInstancesInner) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



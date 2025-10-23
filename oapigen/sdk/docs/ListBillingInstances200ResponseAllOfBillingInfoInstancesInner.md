# ListBillingInstances200ResponseAllOfBillingInfoInstancesInner

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
**Containers** | Pointer to [**[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner**](ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner.md) |  | [optional] 

## Methods

### NewListBillingInstances200ResponseAllOfBillingInfoInstancesInner

`func NewListBillingInstances200ResponseAllOfBillingInfoInstancesInner() *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner`

NewListBillingInstances200ResponseAllOfBillingInfoInstancesInner instantiates a new ListBillingInstances200ResponseAllOfBillingInfoInstancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerWithDefaults

`func NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerWithDefaults() *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner`

NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerWithDefaults instantiates a new ListBillingInstances200ResponseAllOfBillingInfoInstancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceUUID

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetInstanceUUID() string`

GetInstanceUUID returns the InstanceUUID field if non-nil, zero value otherwise.

### GetInstanceUUIDOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetInstanceUUIDOk() (*string, bool)`

GetInstanceUUIDOk returns a tuple with the InstanceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUUID

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) SetInstanceUUID(v string)`

SetInstanceUUID sets InstanceUUID field to given value.

### HasInstanceUUID

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) HasInstanceUUID() bool`

HasInstanceUUID returns a boolean if a field has been set.

### GetStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetContainers

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetContainers() []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) GetContainersOk() (*[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) SetContainers(v []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



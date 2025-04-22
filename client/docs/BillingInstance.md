# BillingInstance

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
**Containers** | Pointer to [**[]GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner**](GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner.md) |  | [optional] 

## Methods

### NewBillingInstance

`func NewBillingInstance() *BillingInstance`

NewBillingInstance instantiates a new BillingInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInstanceWithDefaults

`func NewBillingInstanceWithDefaults() *BillingInstance`

NewBillingInstanceWithDefaults instantiates a new BillingInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *BillingInstance) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *BillingInstance) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *BillingInstance) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *BillingInstance) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceUUID

`func (o *BillingInstance) GetInstanceUUID() string`

GetInstanceUUID returns the InstanceUUID field if non-nil, zero value otherwise.

### GetInstanceUUIDOk

`func (o *BillingInstance) GetInstanceUUIDOk() (*string, bool)`

GetInstanceUUIDOk returns a tuple with the InstanceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUUID

`func (o *BillingInstance) SetInstanceUUID(v string)`

SetInstanceUUID sets InstanceUUID field to given value.

### HasInstanceUUID

`func (o *BillingInstance) HasInstanceUUID() bool`

HasInstanceUUID returns a boolean if a field has been set.

### GetStartDate

`func (o *BillingInstance) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingInstance) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingInstance) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingInstance) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingInstance) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingInstance) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingInstance) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingInstance) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetName

`func (o *BillingInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *BillingInstance) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingInstance) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingInstance) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingInstance) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *BillingInstance) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingInstance) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingInstance) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingInstance) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingInstance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingInstance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingInstance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingInstance) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetContainers

`func (o *BillingInstance) GetContainers() []GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *BillingInstance) GetContainersOk() (*[]GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *BillingInstance) SetContainers(v []GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *BillingInstance) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



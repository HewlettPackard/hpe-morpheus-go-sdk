# GetBillingInstancesIdentifier200ResponseAllOfBillingInfo

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

### NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfo

`func NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfo() *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo`

NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfo instantiates a new GetBillingInstancesIdentifier200ResponseAllOfBillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoWithDefaults

`func NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoWithDefaults() *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo`

NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoWithDefaults instantiates a new GetBillingInstancesIdentifier200ResponseAllOfBillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceUUID

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetInstanceUUID() string`

GetInstanceUUID returns the InstanceUUID field if non-nil, zero value otherwise.

### GetInstanceUUIDOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetInstanceUUIDOk() (*string, bool)`

GetInstanceUUIDOk returns a tuple with the InstanceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUUID

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) SetInstanceUUID(v string)`

SetInstanceUUID sets InstanceUUID field to given value.

### HasInstanceUUID

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) HasInstanceUUID() bool`

HasInstanceUUID returns a boolean if a field has been set.

### GetStartDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetContainers

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetContainers() []GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) GetContainersOk() (*[]GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) SetContainers(v []GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfo) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



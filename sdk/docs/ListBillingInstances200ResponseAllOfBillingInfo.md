# ListBillingInstances200ResponseAllOfBillingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Instances** | Pointer to [**[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInner**](ListBillingInstances200ResponseAllOfBillingInfoInstancesInner.md) |  | [optional] 

## Methods

### NewListBillingInstances200ResponseAllOfBillingInfo

`func NewListBillingInstances200ResponseAllOfBillingInfo() *ListBillingInstances200ResponseAllOfBillingInfo`

NewListBillingInstances200ResponseAllOfBillingInfo instantiates a new ListBillingInstances200ResponseAllOfBillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingInstances200ResponseAllOfBillingInfoWithDefaults

`func NewListBillingInstances200ResponseAllOfBillingInfoWithDefaults() *ListBillingInstances200ResponseAllOfBillingInfo`

NewListBillingInstances200ResponseAllOfBillingInfoWithDefaults instantiates a new ListBillingInstances200ResponseAllOfBillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInstances

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) GetInstances() []ListBillingInstances200ResponseAllOfBillingInfoInstancesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) GetInstancesOk() (*[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) SetInstances(v []ListBillingInstances200ResponseAllOfBillingInfoInstancesInner)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ListBillingInstances200ResponseAllOfBillingInfo) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



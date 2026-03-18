# BillingInstancesInstancesInnerContainersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | Pointer to **string** |  | [optional] 
**RefUUID** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**NumUnits** | Pointer to **float32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Usages** | Pointer to [**[]BillingInstancesInstancesInnerContainersInnerUsagesInner**](BillingInstancesInstancesInnerContainersInnerUsagesInner.md) |  | [optional] 
**NumUsages** | Pointer to **int64** |  | [optional] 
**TotalUsages** | Pointer to **int64** |  | [optional] 
**HasMoreUsages** | Pointer to **bool** |  | [optional] 
**FoundPricing** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **int64** |  | [optional] 
**ServerUUID** | Pointer to **string** |  | [optional] 
**ServerUniqueId** | Pointer to **string** |  | [optional] 
**ServerExternalId** | Pointer to **string** |  | [optional] 
**ServerInternalId** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**ResourcePoolName** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingInstancesInstancesInnerContainersInner

`func NewBillingInstancesInstancesInnerContainersInner() *BillingInstancesInstancesInnerContainersInner`

NewBillingInstancesInstancesInnerContainersInner instantiates a new BillingInstancesInstancesInnerContainersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInstancesInstancesInnerContainersInnerWithDefaults

`func NewBillingInstancesInstancesInnerContainersInnerWithDefaults() *BillingInstancesInstancesInnerContainersInner`

NewBillingInstancesInstancesInnerContainersInnerWithDefaults instantiates a new BillingInstancesInstancesInnerContainersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *BillingInstancesInstancesInnerContainersInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *BillingInstancesInstancesInnerContainersInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *BillingInstancesInstancesInnerContainersInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefUUID

`func (o *BillingInstancesInstancesInnerContainersInner) GetRefUUID() string`

GetRefUUID returns the RefUUID field if non-nil, zero value otherwise.

### GetRefUUIDOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetRefUUIDOk() (*string, bool)`

GetRefUUIDOk returns a tuple with the RefUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUUID

`func (o *BillingInstancesInstancesInnerContainersInner) SetRefUUID(v string)`

SetRefUUID sets RefUUID field to given value.

### HasRefUUID

`func (o *BillingInstancesInstancesInnerContainersInner) HasRefUUID() bool`

HasRefUUID returns a boolean if a field has been set.

### GetRefId

`func (o *BillingInstancesInstancesInnerContainersInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *BillingInstancesInstancesInnerContainersInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *BillingInstancesInstancesInnerContainersInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStartDate

`func (o *BillingInstancesInstancesInnerContainersInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingInstancesInstancesInnerContainersInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingInstancesInstancesInnerContainersInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingInstancesInstancesInnerContainersInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingInstancesInstancesInnerContainersInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingInstancesInstancesInnerContainersInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCost

`func (o *BillingInstancesInstancesInnerContainersInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingInstancesInstancesInnerContainersInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingInstancesInstancesInnerContainersInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *BillingInstancesInstancesInnerContainersInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingInstancesInstancesInnerContainersInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingInstancesInstancesInnerContainersInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetNumUnits

`func (o *BillingInstancesInstancesInnerContainersInner) GetNumUnits() float32`

GetNumUnits returns the NumUnits field if non-nil, zero value otherwise.

### GetNumUnitsOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetNumUnitsOk() (*float32, bool)`

GetNumUnitsOk returns a tuple with the NumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnits

`func (o *BillingInstancesInstancesInnerContainersInner) SetNumUnits(v float32)`

SetNumUnits sets NumUnits field to given value.

### HasNumUnits

`func (o *BillingInstancesInstancesInnerContainersInner) HasNumUnits() bool`

HasNumUnits returns a boolean if a field has been set.

### GetUnit

`func (o *BillingInstancesInstancesInnerContainersInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BillingInstancesInstancesInnerContainersInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BillingInstancesInstancesInnerContainersInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingInstancesInstancesInnerContainersInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingInstancesInstancesInnerContainersInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingInstancesInstancesInnerContainersInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUsages

`func (o *BillingInstancesInstancesInnerContainersInner) GetUsages() []BillingInstancesInstancesInnerContainersInnerUsagesInner`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetUsagesOk() (*[]BillingInstancesInstancesInnerContainersInnerUsagesInner, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *BillingInstancesInstancesInnerContainersInner) SetUsages(v []BillingInstancesInstancesInnerContainersInnerUsagesInner)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *BillingInstancesInstancesInnerContainersInner) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetNumUsages

`func (o *BillingInstancesInstancesInnerContainersInner) GetNumUsages() int64`

GetNumUsages returns the NumUsages field if non-nil, zero value otherwise.

### GetNumUsagesOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetNumUsagesOk() (*int64, bool)`

GetNumUsagesOk returns a tuple with the NumUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsages

`func (o *BillingInstancesInstancesInnerContainersInner) SetNumUsages(v int64)`

SetNumUsages sets NumUsages field to given value.

### HasNumUsages

`func (o *BillingInstancesInstancesInnerContainersInner) HasNumUsages() bool`

HasNumUsages returns a boolean if a field has been set.

### GetTotalUsages

`func (o *BillingInstancesInstancesInnerContainersInner) GetTotalUsages() int64`

GetTotalUsages returns the TotalUsages field if non-nil, zero value otherwise.

### GetTotalUsagesOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetTotalUsagesOk() (*int64, bool)`

GetTotalUsagesOk returns a tuple with the TotalUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsages

`func (o *BillingInstancesInstancesInnerContainersInner) SetTotalUsages(v int64)`

SetTotalUsages sets TotalUsages field to given value.

### HasTotalUsages

`func (o *BillingInstancesInstancesInnerContainersInner) HasTotalUsages() bool`

HasTotalUsages returns a boolean if a field has been set.

### GetHasMoreUsages

`func (o *BillingInstancesInstancesInnerContainersInner) GetHasMoreUsages() bool`

GetHasMoreUsages returns the HasMoreUsages field if non-nil, zero value otherwise.

### GetHasMoreUsagesOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetHasMoreUsagesOk() (*bool, bool)`

GetHasMoreUsagesOk returns a tuple with the HasMoreUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreUsages

`func (o *BillingInstancesInstancesInnerContainersInner) SetHasMoreUsages(v bool)`

SetHasMoreUsages sets HasMoreUsages field to given value.

### HasHasMoreUsages

`func (o *BillingInstancesInstancesInnerContainersInner) HasHasMoreUsages() bool`

HasHasMoreUsages returns a boolean if a field has been set.

### GetFoundPricing

`func (o *BillingInstancesInstancesInnerContainersInner) GetFoundPricing() bool`

GetFoundPricing returns the FoundPricing field if non-nil, zero value otherwise.

### GetFoundPricingOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetFoundPricingOk() (*bool, bool)`

GetFoundPricingOk returns a tuple with the FoundPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundPricing

`func (o *BillingInstancesInstancesInnerContainersInner) SetFoundPricing(v bool)`

SetFoundPricing sets FoundPricing field to given value.

### HasFoundPricing

`func (o *BillingInstancesInstancesInnerContainersInner) HasFoundPricing() bool`

HasFoundPricing returns a boolean if a field has been set.

### GetName

`func (o *BillingInstancesInstancesInnerContainersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingInstancesInstancesInnerContainersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingInstancesInstancesInnerContainersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerId

`func (o *BillingInstancesInstancesInnerContainersInner) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *BillingInstancesInstancesInnerContainersInner) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *BillingInstancesInstancesInnerContainersInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerUUID

`func (o *BillingInstancesInstancesInnerContainersInner) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *BillingInstancesInstancesInnerContainersInner) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *BillingInstancesInstancesInnerContainersInner) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetServerUniqueId

`func (o *BillingInstancesInstancesInnerContainersInner) GetServerUniqueId() string`

GetServerUniqueId returns the ServerUniqueId field if non-nil, zero value otherwise.

### GetServerUniqueIdOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetServerUniqueIdOk() (*string, bool)`

GetServerUniqueIdOk returns a tuple with the ServerUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUniqueId

`func (o *BillingInstancesInstancesInnerContainersInner) SetServerUniqueId(v string)`

SetServerUniqueId sets ServerUniqueId field to given value.

### HasServerUniqueId

`func (o *BillingInstancesInstancesInnerContainersInner) HasServerUniqueId() bool`

HasServerUniqueId returns a boolean if a field has been set.

### GetServerExternalId

`func (o *BillingInstancesInstancesInnerContainersInner) GetServerExternalId() string`

GetServerExternalId returns the ServerExternalId field if non-nil, zero value otherwise.

### GetServerExternalIdOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetServerExternalIdOk() (*string, bool)`

GetServerExternalIdOk returns a tuple with the ServerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExternalId

`func (o *BillingInstancesInstancesInnerContainersInner) SetServerExternalId(v string)`

SetServerExternalId sets ServerExternalId field to given value.

### HasServerExternalId

`func (o *BillingInstancesInstancesInnerContainersInner) HasServerExternalId() bool`

HasServerExternalId returns a boolean if a field has been set.

### GetServerInternalId

`func (o *BillingInstancesInstancesInnerContainersInner) GetServerInternalId() string`

GetServerInternalId returns the ServerInternalId field if non-nil, zero value otherwise.

### GetServerInternalIdOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetServerInternalIdOk() (*string, bool)`

GetServerInternalIdOk returns a tuple with the ServerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternalId

`func (o *BillingInstancesInstancesInnerContainersInner) SetServerInternalId(v string)`

SetServerInternalId sets ServerInternalId field to given value.

### HasServerInternalId

`func (o *BillingInstancesInstancesInnerContainersInner) HasServerInternalId() bool`

HasServerInternalId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *BillingInstancesInstancesInnerContainersInner) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *BillingInstancesInstancesInnerContainersInner) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *BillingInstancesInstancesInnerContainersInner) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePoolName

`func (o *BillingInstancesInstancesInnerContainersInner) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *BillingInstancesInstancesInnerContainersInner) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *BillingInstancesInstancesInnerContainersInner) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *BillingInstancesInstancesInnerContainersInner) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



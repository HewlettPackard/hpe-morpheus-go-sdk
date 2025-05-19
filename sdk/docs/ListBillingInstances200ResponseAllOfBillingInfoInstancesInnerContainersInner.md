# ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner

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
**Usages** | Pointer to [**[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner**](ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner.md) |  | [optional] 
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

### NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner

`func NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner() *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner`

NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner instantiates a new ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerWithDefaults

`func NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerWithDefaults() *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner`

NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerWithDefaults instantiates a new ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefUUID

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetRefUUID() string`

GetRefUUID returns the RefUUID field if non-nil, zero value otherwise.

### GetRefUUIDOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetRefUUIDOk() (*string, bool)`

GetRefUUIDOk returns a tuple with the RefUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUUID

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetRefUUID(v string)`

SetRefUUID sets RefUUID field to given value.

### HasRefUUID

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasRefUUID() bool`

HasRefUUID returns a boolean if a field has been set.

### GetRefId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetNumUnits

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetNumUnits() float32`

GetNumUnits returns the NumUnits field if non-nil, zero value otherwise.

### GetNumUnitsOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetNumUnitsOk() (*float32, bool)`

GetNumUnitsOk returns a tuple with the NumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnits

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetNumUnits(v float32)`

SetNumUnits sets NumUnits field to given value.

### HasNumUnits

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasNumUnits() bool`

HasNumUnits returns a boolean if a field has been set.

### GetUnit

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUsages

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetUsages() []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetUsagesOk() (*[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetUsages(v []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetNumUsages

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetNumUsages() int64`

GetNumUsages returns the NumUsages field if non-nil, zero value otherwise.

### GetNumUsagesOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetNumUsagesOk() (*int64, bool)`

GetNumUsagesOk returns a tuple with the NumUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsages

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetNumUsages(v int64)`

SetNumUsages sets NumUsages field to given value.

### HasNumUsages

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasNumUsages() bool`

HasNumUsages returns a boolean if a field has been set.

### GetTotalUsages

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetTotalUsages() int64`

GetTotalUsages returns the TotalUsages field if non-nil, zero value otherwise.

### GetTotalUsagesOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetTotalUsagesOk() (*int64, bool)`

GetTotalUsagesOk returns a tuple with the TotalUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsages

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetTotalUsages(v int64)`

SetTotalUsages sets TotalUsages field to given value.

### HasTotalUsages

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasTotalUsages() bool`

HasTotalUsages returns a boolean if a field has been set.

### GetHasMoreUsages

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetHasMoreUsages() bool`

GetHasMoreUsages returns the HasMoreUsages field if non-nil, zero value otherwise.

### GetHasMoreUsagesOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetHasMoreUsagesOk() (*bool, bool)`

GetHasMoreUsagesOk returns a tuple with the HasMoreUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreUsages

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetHasMoreUsages(v bool)`

SetHasMoreUsages sets HasMoreUsages field to given value.

### HasHasMoreUsages

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasHasMoreUsages() bool`

HasHasMoreUsages returns a boolean if a field has been set.

### GetFoundPricing

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetFoundPricing() bool`

GetFoundPricing returns the FoundPricing field if non-nil, zero value otherwise.

### GetFoundPricingOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetFoundPricingOk() (*bool, bool)`

GetFoundPricingOk returns a tuple with the FoundPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundPricing

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetFoundPricing(v bool)`

SetFoundPricing sets FoundPricing field to given value.

### HasFoundPricing

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasFoundPricing() bool`

HasFoundPricing returns a boolean if a field has been set.

### GetName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerUUID

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetServerUniqueId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetServerUniqueId() string`

GetServerUniqueId returns the ServerUniqueId field if non-nil, zero value otherwise.

### GetServerUniqueIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetServerUniqueIdOk() (*string, bool)`

GetServerUniqueIdOk returns a tuple with the ServerUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUniqueId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetServerUniqueId(v string)`

SetServerUniqueId sets ServerUniqueId field to given value.

### HasServerUniqueId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasServerUniqueId() bool`

HasServerUniqueId returns a boolean if a field has been set.

### GetServerExternalId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetServerExternalId() string`

GetServerExternalId returns the ServerExternalId field if non-nil, zero value otherwise.

### GetServerExternalIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetServerExternalIdOk() (*string, bool)`

GetServerExternalIdOk returns a tuple with the ServerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExternalId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetServerExternalId(v string)`

SetServerExternalId sets ServerExternalId field to given value.

### HasServerExternalId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasServerExternalId() bool`

HasServerExternalId returns a boolean if a field has been set.

### GetServerInternalId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetServerInternalId() string`

GetServerInternalId returns the ServerInternalId field if non-nil, zero value otherwise.

### GetServerInternalIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetServerInternalIdOk() (*string, bool)`

GetServerInternalIdOk returns a tuple with the ServerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternalId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetServerInternalId(v string)`

SetServerInternalId sets ServerInternalId field to given value.

### HasServerInternalId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasServerInternalId() bool`

HasServerInternalId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePoolName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInner) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



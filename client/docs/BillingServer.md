# BillingServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | Pointer to **string** |  | [optional] 
**RefUUID** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**NumUnits** | Pointer to **float32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Usages** | Pointer to [**[]GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner**](GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner.md) |  | [optional] 
**NumUsages** | Pointer to **int64** |  | [optional] 
**TotalUsages** | Pointer to **int64** |  | [optional] 
**HasMoreUsages** | Pointer to **bool** |  | [optional] 
**FoundPricing** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ServerUUID** | Pointer to **string** |  | [optional] 
**ServerUniqueId** | Pointer to **string** |  | [optional] 
**ServerExternalId** | Pointer to **string** |  | [optional] 
**ServerInternalId** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**ResourcePoolName** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingServer

`func NewBillingServer() *BillingServer`

NewBillingServer instantiates a new BillingServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingServerWithDefaults

`func NewBillingServerWithDefaults() *BillingServer`

NewBillingServerWithDefaults instantiates a new BillingServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *BillingServer) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *BillingServer) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *BillingServer) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *BillingServer) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefUUID

`func (o *BillingServer) GetRefUUID() string`

GetRefUUID returns the RefUUID field if non-nil, zero value otherwise.

### GetRefUUIDOk

`func (o *BillingServer) GetRefUUIDOk() (*string, bool)`

GetRefUUIDOk returns a tuple with the RefUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUUID

`func (o *BillingServer) SetRefUUID(v string)`

SetRefUUID sets RefUUID field to given value.

### HasRefUUID

`func (o *BillingServer) HasRefUUID() bool`

HasRefUUID returns a boolean if a field has been set.

### GetRefId

`func (o *BillingServer) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *BillingServer) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *BillingServer) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *BillingServer) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStartDate

`func (o *BillingServer) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingServer) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingServer) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingServer) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingServer) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingServer) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingServer) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingServer) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCost

`func (o *BillingServer) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingServer) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingServer) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingServer) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *BillingServer) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingServer) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingServer) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingServer) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetNumUnits

`func (o *BillingServer) GetNumUnits() float32`

GetNumUnits returns the NumUnits field if non-nil, zero value otherwise.

### GetNumUnitsOk

`func (o *BillingServer) GetNumUnitsOk() (*float32, bool)`

GetNumUnitsOk returns a tuple with the NumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnits

`func (o *BillingServer) SetNumUnits(v float32)`

SetNumUnits sets NumUnits field to given value.

### HasNumUnits

`func (o *BillingServer) HasNumUnits() bool`

HasNumUnits returns a boolean if a field has been set.

### GetUnit

`func (o *BillingServer) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BillingServer) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BillingServer) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BillingServer) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingServer) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingServer) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingServer) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingServer) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUsages

`func (o *BillingServer) GetUsages() []GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *BillingServer) GetUsagesOk() (*[]GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *BillingServer) SetUsages(v []GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *BillingServer) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetNumUsages

`func (o *BillingServer) GetNumUsages() int64`

GetNumUsages returns the NumUsages field if non-nil, zero value otherwise.

### GetNumUsagesOk

`func (o *BillingServer) GetNumUsagesOk() (*int64, bool)`

GetNumUsagesOk returns a tuple with the NumUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsages

`func (o *BillingServer) SetNumUsages(v int64)`

SetNumUsages sets NumUsages field to given value.

### HasNumUsages

`func (o *BillingServer) HasNumUsages() bool`

HasNumUsages returns a boolean if a field has been set.

### GetTotalUsages

`func (o *BillingServer) GetTotalUsages() int64`

GetTotalUsages returns the TotalUsages field if non-nil, zero value otherwise.

### GetTotalUsagesOk

`func (o *BillingServer) GetTotalUsagesOk() (*int64, bool)`

GetTotalUsagesOk returns a tuple with the TotalUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsages

`func (o *BillingServer) SetTotalUsages(v int64)`

SetTotalUsages sets TotalUsages field to given value.

### HasTotalUsages

`func (o *BillingServer) HasTotalUsages() bool`

HasTotalUsages returns a boolean if a field has been set.

### GetHasMoreUsages

`func (o *BillingServer) GetHasMoreUsages() bool`

GetHasMoreUsages returns the HasMoreUsages field if non-nil, zero value otherwise.

### GetHasMoreUsagesOk

`func (o *BillingServer) GetHasMoreUsagesOk() (*bool, bool)`

GetHasMoreUsagesOk returns a tuple with the HasMoreUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreUsages

`func (o *BillingServer) SetHasMoreUsages(v bool)`

SetHasMoreUsages sets HasMoreUsages field to given value.

### HasHasMoreUsages

`func (o *BillingServer) HasHasMoreUsages() bool`

HasHasMoreUsages returns a boolean if a field has been set.

### GetFoundPricing

`func (o *BillingServer) GetFoundPricing() bool`

GetFoundPricing returns the FoundPricing field if non-nil, zero value otherwise.

### GetFoundPricingOk

`func (o *BillingServer) GetFoundPricingOk() (*bool, bool)`

GetFoundPricingOk returns a tuple with the FoundPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundPricing

`func (o *BillingServer) SetFoundPricing(v bool)`

SetFoundPricing sets FoundPricing field to given value.

### HasFoundPricing

`func (o *BillingServer) HasFoundPricing() bool`

HasFoundPricing returns a boolean if a field has been set.

### GetName

`func (o *BillingServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerUUID

`func (o *BillingServer) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *BillingServer) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *BillingServer) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *BillingServer) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetServerUniqueId

`func (o *BillingServer) GetServerUniqueId() string`

GetServerUniqueId returns the ServerUniqueId field if non-nil, zero value otherwise.

### GetServerUniqueIdOk

`func (o *BillingServer) GetServerUniqueIdOk() (*string, bool)`

GetServerUniqueIdOk returns a tuple with the ServerUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUniqueId

`func (o *BillingServer) SetServerUniqueId(v string)`

SetServerUniqueId sets ServerUniqueId field to given value.

### HasServerUniqueId

`func (o *BillingServer) HasServerUniqueId() bool`

HasServerUniqueId returns a boolean if a field has been set.

### GetServerExternalId

`func (o *BillingServer) GetServerExternalId() string`

GetServerExternalId returns the ServerExternalId field if non-nil, zero value otherwise.

### GetServerExternalIdOk

`func (o *BillingServer) GetServerExternalIdOk() (*string, bool)`

GetServerExternalIdOk returns a tuple with the ServerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExternalId

`func (o *BillingServer) SetServerExternalId(v string)`

SetServerExternalId sets ServerExternalId field to given value.

### HasServerExternalId

`func (o *BillingServer) HasServerExternalId() bool`

HasServerExternalId returns a boolean if a field has been set.

### GetServerInternalId

`func (o *BillingServer) GetServerInternalId() string`

GetServerInternalId returns the ServerInternalId field if non-nil, zero value otherwise.

### GetServerInternalIdOk

`func (o *BillingServer) GetServerInternalIdOk() (*string, bool)`

GetServerInternalIdOk returns a tuple with the ServerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternalId

`func (o *BillingServer) SetServerInternalId(v string)`

SetServerInternalId sets ServerInternalId field to given value.

### HasServerInternalId

`func (o *BillingServer) HasServerInternalId() bool`

HasServerInternalId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *BillingServer) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *BillingServer) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *BillingServer) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *BillingServer) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePoolName

`func (o *BillingServer) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *BillingServer) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *BillingServer) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *BillingServer) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



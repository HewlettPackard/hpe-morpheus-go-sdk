# ListBillingServers200ResponseAllOfBillingInfoServersInner

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
**Usages** | Pointer to [**[]ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner**](ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner.md) |  | [optional] 
**NumUsages** | Pointer to **int64** |  | [optional] 
**TotalUsages** | Pointer to **int64** |  | [optional] 
**HasMoreUsages** | Pointer to **bool** |  | [optional] 
**FoundPricing** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ServerUUID** | Pointer to **string** |  | [optional] 
**ServerUniqueId** | Pointer to **string** |  | [optional] 
**ServerExternalId** | Pointer to **string** |  | [optional] 
**ServerInternalId** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **string** |  | [optional] 
**ResourcePoolName** | Pointer to **string** |  | [optional] 

## Methods

### NewListBillingServers200ResponseAllOfBillingInfoServersInner

`func NewListBillingServers200ResponseAllOfBillingInfoServersInner() *ListBillingServers200ResponseAllOfBillingInfoServersInner`

NewListBillingServers200ResponseAllOfBillingInfoServersInner instantiates a new ListBillingServers200ResponseAllOfBillingInfoServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingServers200ResponseAllOfBillingInfoServersInnerWithDefaults

`func NewListBillingServers200ResponseAllOfBillingInfoServersInnerWithDefaults() *ListBillingServers200ResponseAllOfBillingInfoServersInner`

NewListBillingServers200ResponseAllOfBillingInfoServersInnerWithDefaults instantiates a new ListBillingServers200ResponseAllOfBillingInfoServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefUUID

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetRefUUID() string`

GetRefUUID returns the RefUUID field if non-nil, zero value otherwise.

### GetRefUUIDOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetRefUUIDOk() (*string, bool)`

GetRefUUIDOk returns a tuple with the RefUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUUID

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetRefUUID(v string)`

SetRefUUID sets RefUUID field to given value.

### HasRefUUID

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasRefUUID() bool`

HasRefUUID returns a boolean if a field has been set.

### GetRefId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStartDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCost

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetNumUnits

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetNumUnits() float32`

GetNumUnits returns the NumUnits field if non-nil, zero value otherwise.

### GetNumUnitsOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetNumUnitsOk() (*float32, bool)`

GetNumUnitsOk returns a tuple with the NumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnits

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetNumUnits(v float32)`

SetNumUnits sets NumUnits field to given value.

### HasNumUnits

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasNumUnits() bool`

HasNumUnits returns a boolean if a field has been set.

### GetUnit

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUsages

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetUsages() []ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetUsagesOk() (*[]ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetUsages(v []ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetNumUsages

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetNumUsages() int64`

GetNumUsages returns the NumUsages field if non-nil, zero value otherwise.

### GetNumUsagesOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetNumUsagesOk() (*int64, bool)`

GetNumUsagesOk returns a tuple with the NumUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsages

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetNumUsages(v int64)`

SetNumUsages sets NumUsages field to given value.

### HasNumUsages

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasNumUsages() bool`

HasNumUsages returns a boolean if a field has been set.

### GetTotalUsages

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetTotalUsages() int64`

GetTotalUsages returns the TotalUsages field if non-nil, zero value otherwise.

### GetTotalUsagesOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetTotalUsagesOk() (*int64, bool)`

GetTotalUsagesOk returns a tuple with the TotalUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsages

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetTotalUsages(v int64)`

SetTotalUsages sets TotalUsages field to given value.

### HasTotalUsages

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasTotalUsages() bool`

HasTotalUsages returns a boolean if a field has been set.

### GetHasMoreUsages

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetHasMoreUsages() bool`

GetHasMoreUsages returns the HasMoreUsages field if non-nil, zero value otherwise.

### GetHasMoreUsagesOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetHasMoreUsagesOk() (*bool, bool)`

GetHasMoreUsagesOk returns a tuple with the HasMoreUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreUsages

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetHasMoreUsages(v bool)`

SetHasMoreUsages sets HasMoreUsages field to given value.

### HasHasMoreUsages

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasHasMoreUsages() bool`

HasHasMoreUsages returns a boolean if a field has been set.

### GetFoundPricing

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetFoundPricing() bool`

GetFoundPricing returns the FoundPricing field if non-nil, zero value otherwise.

### GetFoundPricingOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetFoundPricingOk() (*bool, bool)`

GetFoundPricingOk returns a tuple with the FoundPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundPricing

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetFoundPricing(v bool)`

SetFoundPricing sets FoundPricing field to given value.

### HasFoundPricing

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasFoundPricing() bool`

HasFoundPricing returns a boolean if a field has been set.

### GetName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerUUID

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetServerUniqueId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetServerUniqueId() string`

GetServerUniqueId returns the ServerUniqueId field if non-nil, zero value otherwise.

### GetServerUniqueIdOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetServerUniqueIdOk() (*string, bool)`

GetServerUniqueIdOk returns a tuple with the ServerUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUniqueId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetServerUniqueId(v string)`

SetServerUniqueId sets ServerUniqueId field to given value.

### HasServerUniqueId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasServerUniqueId() bool`

HasServerUniqueId returns a boolean if a field has been set.

### GetServerExternalId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetServerExternalId() string`

GetServerExternalId returns the ServerExternalId field if non-nil, zero value otherwise.

### GetServerExternalIdOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetServerExternalIdOk() (*string, bool)`

GetServerExternalIdOk returns a tuple with the ServerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExternalId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetServerExternalId(v string)`

SetServerExternalId sets ServerExternalId field to given value.

### HasServerExternalId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasServerExternalId() bool`

HasServerExternalId returns a boolean if a field has been set.

### GetServerInternalId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetServerInternalId() string`

GetServerInternalId returns the ServerInternalId field if non-nil, zero value otherwise.

### GetServerInternalIdOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetServerInternalIdOk() (*string, bool)`

GetServerInternalIdOk returns a tuple with the ServerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternalId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetServerInternalId(v string)`

SetServerInternalId sets ServerInternalId field to given value.

### HasServerInternalId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasServerInternalId() bool`

HasServerInternalId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePoolName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInner) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



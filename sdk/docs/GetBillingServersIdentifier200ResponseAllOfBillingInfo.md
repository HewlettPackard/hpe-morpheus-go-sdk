# GetBillingServersIdentifier200ResponseAllOfBillingInfo

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

### NewGetBillingServersIdentifier200ResponseAllOfBillingInfo

`func NewGetBillingServersIdentifier200ResponseAllOfBillingInfo() *GetBillingServersIdentifier200ResponseAllOfBillingInfo`

NewGetBillingServersIdentifier200ResponseAllOfBillingInfo instantiates a new GetBillingServersIdentifier200ResponseAllOfBillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBillingServersIdentifier200ResponseAllOfBillingInfoWithDefaults

`func NewGetBillingServersIdentifier200ResponseAllOfBillingInfoWithDefaults() *GetBillingServersIdentifier200ResponseAllOfBillingInfo`

NewGetBillingServersIdentifier200ResponseAllOfBillingInfoWithDefaults instantiates a new GetBillingServersIdentifier200ResponseAllOfBillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefUUID

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetRefUUID() string`

GetRefUUID returns the RefUUID field if non-nil, zero value otherwise.

### GetRefUUIDOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetRefUUIDOk() (*string, bool)`

GetRefUUIDOk returns a tuple with the RefUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUUID

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetRefUUID(v string)`

SetRefUUID sets RefUUID field to given value.

### HasRefUUID

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasRefUUID() bool`

HasRefUUID returns a boolean if a field has been set.

### GetRefId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStartDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCost

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetNumUnits

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetNumUnits() float32`

GetNumUnits returns the NumUnits field if non-nil, zero value otherwise.

### GetNumUnitsOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetNumUnitsOk() (*float32, bool)`

GetNumUnitsOk returns a tuple with the NumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnits

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetNumUnits(v float32)`

SetNumUnits sets NumUnits field to given value.

### HasNumUnits

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasNumUnits() bool`

HasNumUnits returns a boolean if a field has been set.

### GetUnit

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUsages

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetUsages() []GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetUsagesOk() (*[]GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetUsages(v []GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetNumUsages

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetNumUsages() int64`

GetNumUsages returns the NumUsages field if non-nil, zero value otherwise.

### GetNumUsagesOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetNumUsagesOk() (*int64, bool)`

GetNumUsagesOk returns a tuple with the NumUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsages

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetNumUsages(v int64)`

SetNumUsages sets NumUsages field to given value.

### HasNumUsages

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasNumUsages() bool`

HasNumUsages returns a boolean if a field has been set.

### GetTotalUsages

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetTotalUsages() int64`

GetTotalUsages returns the TotalUsages field if non-nil, zero value otherwise.

### GetTotalUsagesOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetTotalUsagesOk() (*int64, bool)`

GetTotalUsagesOk returns a tuple with the TotalUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsages

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetTotalUsages(v int64)`

SetTotalUsages sets TotalUsages field to given value.

### HasTotalUsages

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasTotalUsages() bool`

HasTotalUsages returns a boolean if a field has been set.

### GetHasMoreUsages

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetHasMoreUsages() bool`

GetHasMoreUsages returns the HasMoreUsages field if non-nil, zero value otherwise.

### GetHasMoreUsagesOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetHasMoreUsagesOk() (*bool, bool)`

GetHasMoreUsagesOk returns a tuple with the HasMoreUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreUsages

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetHasMoreUsages(v bool)`

SetHasMoreUsages sets HasMoreUsages field to given value.

### HasHasMoreUsages

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasHasMoreUsages() bool`

HasHasMoreUsages returns a boolean if a field has been set.

### GetFoundPricing

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetFoundPricing() bool`

GetFoundPricing returns the FoundPricing field if non-nil, zero value otherwise.

### GetFoundPricingOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetFoundPricingOk() (*bool, bool)`

GetFoundPricingOk returns a tuple with the FoundPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundPricing

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetFoundPricing(v bool)`

SetFoundPricing sets FoundPricing field to given value.

### HasFoundPricing

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasFoundPricing() bool`

HasFoundPricing returns a boolean if a field has been set.

### GetName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerUUID

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetServerUniqueId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetServerUniqueId() string`

GetServerUniqueId returns the ServerUniqueId field if non-nil, zero value otherwise.

### GetServerUniqueIdOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetServerUniqueIdOk() (*string, bool)`

GetServerUniqueIdOk returns a tuple with the ServerUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUniqueId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetServerUniqueId(v string)`

SetServerUniqueId sets ServerUniqueId field to given value.

### HasServerUniqueId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasServerUniqueId() bool`

HasServerUniqueId returns a boolean if a field has been set.

### GetServerExternalId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetServerExternalId() string`

GetServerExternalId returns the ServerExternalId field if non-nil, zero value otherwise.

### GetServerExternalIdOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetServerExternalIdOk() (*string, bool)`

GetServerExternalIdOk returns a tuple with the ServerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExternalId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetServerExternalId(v string)`

SetServerExternalId sets ServerExternalId field to given value.

### HasServerExternalId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasServerExternalId() bool`

HasServerExternalId returns a boolean if a field has been set.

### GetServerInternalId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetServerInternalId() string`

GetServerInternalId returns the ServerInternalId field if non-nil, zero value otherwise.

### GetServerInternalIdOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetServerInternalIdOk() (*string, bool)`

GetServerInternalIdOk returns a tuple with the ServerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternalId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetServerInternalId(v string)`

SetServerInternalId sets ServerInternalId field to given value.

### HasServerInternalId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasServerInternalId() bool`

HasServerInternalId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePoolName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfo) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



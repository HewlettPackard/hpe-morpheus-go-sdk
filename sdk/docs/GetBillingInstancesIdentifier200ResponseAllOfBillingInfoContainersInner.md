# GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner

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
**Usages** | Pointer to [**[]GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner**](GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner.md) |  | [optional] 
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

### NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner

`func NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner() *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner`

NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner instantiates a new GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerWithDefaults

`func NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerWithDefaults() *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner`

NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerWithDefaults instantiates a new GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefUUID

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetRefUUID() string`

GetRefUUID returns the RefUUID field if non-nil, zero value otherwise.

### GetRefUUIDOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetRefUUIDOk() (*string, bool)`

GetRefUUIDOk returns a tuple with the RefUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUUID

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetRefUUID(v string)`

SetRefUUID sets RefUUID field to given value.

### HasRefUUID

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasRefUUID() bool`

HasRefUUID returns a boolean if a field has been set.

### GetRefId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStartDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCost

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetNumUnits

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetNumUnits() float32`

GetNumUnits returns the NumUnits field if non-nil, zero value otherwise.

### GetNumUnitsOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetNumUnitsOk() (*float32, bool)`

GetNumUnitsOk returns a tuple with the NumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnits

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetNumUnits(v float32)`

SetNumUnits sets NumUnits field to given value.

### HasNumUnits

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasNumUnits() bool`

HasNumUnits returns a boolean if a field has been set.

### GetUnit

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCurrency

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUsages

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetUsages() []GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetUsagesOk() (*[]GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetUsages(v []GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetNumUsages

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetNumUsages() int64`

GetNumUsages returns the NumUsages field if non-nil, zero value otherwise.

### GetNumUsagesOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetNumUsagesOk() (*int64, bool)`

GetNumUsagesOk returns a tuple with the NumUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsages

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetNumUsages(v int64)`

SetNumUsages sets NumUsages field to given value.

### HasNumUsages

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasNumUsages() bool`

HasNumUsages returns a boolean if a field has been set.

### GetTotalUsages

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetTotalUsages() int64`

GetTotalUsages returns the TotalUsages field if non-nil, zero value otherwise.

### GetTotalUsagesOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetTotalUsagesOk() (*int64, bool)`

GetTotalUsagesOk returns a tuple with the TotalUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsages

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetTotalUsages(v int64)`

SetTotalUsages sets TotalUsages field to given value.

### HasTotalUsages

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasTotalUsages() bool`

HasTotalUsages returns a boolean if a field has been set.

### GetHasMoreUsages

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetHasMoreUsages() bool`

GetHasMoreUsages returns the HasMoreUsages field if non-nil, zero value otherwise.

### GetHasMoreUsagesOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetHasMoreUsagesOk() (*bool, bool)`

GetHasMoreUsagesOk returns a tuple with the HasMoreUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreUsages

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetHasMoreUsages(v bool)`

SetHasMoreUsages sets HasMoreUsages field to given value.

### HasHasMoreUsages

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasHasMoreUsages() bool`

HasHasMoreUsages returns a boolean if a field has been set.

### GetFoundPricing

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetFoundPricing() bool`

GetFoundPricing returns the FoundPricing field if non-nil, zero value otherwise.

### GetFoundPricingOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetFoundPricingOk() (*bool, bool)`

GetFoundPricingOk returns a tuple with the FoundPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundPricing

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetFoundPricing(v bool)`

SetFoundPricing sets FoundPricing field to given value.

### HasFoundPricing

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasFoundPricing() bool`

HasFoundPricing returns a boolean if a field has been set.

### GetName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerUUID

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetServerUniqueId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetServerUniqueId() string`

GetServerUniqueId returns the ServerUniqueId field if non-nil, zero value otherwise.

### GetServerUniqueIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetServerUniqueIdOk() (*string, bool)`

GetServerUniqueIdOk returns a tuple with the ServerUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUniqueId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetServerUniqueId(v string)`

SetServerUniqueId sets ServerUniqueId field to given value.

### HasServerUniqueId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasServerUniqueId() bool`

HasServerUniqueId returns a boolean if a field has been set.

### GetServerExternalId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetServerExternalId() string`

GetServerExternalId returns the ServerExternalId field if non-nil, zero value otherwise.

### GetServerExternalIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetServerExternalIdOk() (*string, bool)`

GetServerExternalIdOk returns a tuple with the ServerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExternalId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetServerExternalId(v string)`

SetServerExternalId sets ServerExternalId field to given value.

### HasServerExternalId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasServerExternalId() bool`

HasServerExternalId returns a boolean if a field has been set.

### GetServerInternalId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetServerInternalId() string`

GetServerInternalId returns the ServerInternalId field if non-nil, zero value otherwise.

### GetServerInternalIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetServerInternalIdOk() (*string, bool)`

GetServerInternalIdOk returns a tuple with the ServerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternalId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetServerInternalId(v string)`

SetServerInternalId sets ServerInternalId field to given value.

### HasServerInternalId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasServerInternalId() bool`

HasServerInternalId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePoolName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInner) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**CreatedByUser** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 
**SiteUUID** | Pointer to **string** |  | [optional] 
**SiteCode** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ApplicablePrices** | Pointer to [**[]GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner**](GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner.md) |  | [optional] 
**ServicePlanId** | Pointer to **int64** |  | [optional] 
**ServicePlanName** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**ResourcePoolName** | Pointer to **string** |  | [optional] 

## Methods

### NewGetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner

`func NewGetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner() *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner`

NewGetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner instantiates a new GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerWithDefaults

`func NewGetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerWithDefaults() *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner`

NewGetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerWithDefaults instantiates a new GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetCreatedByUser() string`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetCreatedByUserOk() (*string, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetCreatedByUser(v string)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetCreatedByUserId() int64`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetCreatedByUserIdOk() (*int64, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetCreatedByUserId(v int64)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteUUID

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetSiteUUID() string`

GetSiteUUID returns the SiteUUID field if non-nil, zero value otherwise.

### GetSiteUUIDOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetSiteUUIDOk() (*string, bool)`

GetSiteUUIDOk returns a tuple with the SiteUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUUID

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetSiteUUID(v string)`

SetSiteUUID sets SiteUUID field to given value.

### HasSiteUUID

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasSiteUUID() bool`

HasSiteUUID returns a boolean if a field has been set.

### GetSiteCode

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetSiteCode() string`

GetSiteCode returns the SiteCode field if non-nil, zero value otherwise.

### GetSiteCodeOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetSiteCodeOk() (*string, bool)`

GetSiteCodeOk returns a tuple with the SiteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCode

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetSiteCode(v string)`

SetSiteCode sets SiteCode field to given value.

### HasSiteCode

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasSiteCode() bool`

HasSiteCode returns a boolean if a field has been set.

### GetCurrency

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStartDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetApplicablePrices

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetApplicablePrices() []GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner`

GetApplicablePrices returns the ApplicablePrices field if non-nil, zero value otherwise.

### GetApplicablePricesOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetApplicablePricesOk() (*[]GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner, bool)`

GetApplicablePricesOk returns a tuple with the ApplicablePrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicablePrices

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetApplicablePrices(v []GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInnerApplicablePricesInner)`

SetApplicablePrices sets ApplicablePrices field to given value.

### HasApplicablePrices

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasApplicablePrices() bool`

HasApplicablePrices returns a boolean if a field has been set.

### GetServicePlanId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetServicePlanId() int64`

GetServicePlanId returns the ServicePlanId field if non-nil, zero value otherwise.

### GetServicePlanIdOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetServicePlanIdOk() (*int64, bool)`

GetServicePlanIdOk returns a tuple with the ServicePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetServicePlanId(v int64)`

SetServicePlanId sets ServicePlanId field to given value.

### HasServicePlanId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasServicePlanId() bool`

HasServicePlanId returns a boolean if a field has been set.

### GetServicePlanName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.

### HasServicePlanName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasServicePlanName() bool`

HasServicePlanName returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePoolName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *GetBillingServersIdentifier200ResponseAllOfBillingInfoUsagesInner) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



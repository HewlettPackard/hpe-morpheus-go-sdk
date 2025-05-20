# ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ZoneName** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerVolumesInner**](ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerVolumesInner.md) |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **string** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**ServerExternalId** | Pointer to **string** |  | [optional] 
**ServerInternalId** | Pointer to **string** |  | [optional] 
**PlanName** | Pointer to **string** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**HourlyCost** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**PricesUsed** | Pointer to [**[]ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerPricesUsedInner**](ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerPricesUsedInner.md) |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**CreatedByUser** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 
**SiteUUID** | Pointer to **string** |  | [optional] 
**SiteCode** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ApplicablePrices** | Pointer to [**[]ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner**](ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner.md) |  | [optional] 
**ServicePlanId** | Pointer to **int64** |  | [optional] 
**ServicePlanName** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **string** |  | [optional] 
**ResourcePoolName** | Pointer to **string** |  | [optional] 

## Methods

### NewListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner

`func NewListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner() *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner`

NewListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner instantiates a new ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerWithDefaults

`func NewListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerWithDefaults() *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner`

NewListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerWithDefaults instantiates a new ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZoneName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetAccountName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetVolumes

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetVolumes() []ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetVolumesOk() (*[]ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetVolumes(v []ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxCores

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetServerExternalId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetServerExternalId() string`

GetServerExternalId returns the ServerExternalId field if non-nil, zero value otherwise.

### GetServerExternalIdOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetServerExternalIdOk() (*string, bool)`

GetServerExternalIdOk returns a tuple with the ServerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExternalId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetServerExternalId(v string)`

SetServerExternalId sets ServerExternalId field to given value.

### HasServerExternalId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasServerExternalId() bool`

HasServerExternalId returns a boolean if a field has been set.

### GetServerInternalId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetServerInternalId() string`

GetServerInternalId returns the ServerInternalId field if non-nil, zero value otherwise.

### GetServerInternalIdOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetServerInternalIdOk() (*string, bool)`

GetServerInternalIdOk returns a tuple with the ServerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternalId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetServerInternalId(v string)`

SetServerInternalId sets ServerInternalId field to given value.

### HasServerInternalId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasServerInternalId() bool`

HasServerInternalId returns a boolean if a field has been set.

### GetPlanName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetHourlyCost

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetCurrency

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPricesUsed

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetPricesUsed() []ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerPricesUsedInner`

GetPricesUsed returns the PricesUsed field if non-nil, zero value otherwise.

### GetPricesUsedOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetPricesUsedOk() (*[]ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerPricesUsedInner, bool)`

GetPricesUsedOk returns a tuple with the PricesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricesUsed

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetPricesUsed(v []ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerPricesUsedInner)`

SetPricesUsed sets PricesUsed field to given value.

### HasPricesUsed

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasPricesUsed() bool`

HasPricesUsed returns a boolean if a field has been set.

### GetCost

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetCreatedByUser() string`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetCreatedByUserOk() (*string, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetCreatedByUser(v string)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetCreatedByUserId() int64`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetCreatedByUserIdOk() (*int64, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetCreatedByUserId(v int64)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteUUID

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetSiteUUID() string`

GetSiteUUID returns the SiteUUID field if non-nil, zero value otherwise.

### GetSiteUUIDOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetSiteUUIDOk() (*string, bool)`

GetSiteUUIDOk returns a tuple with the SiteUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUUID

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetSiteUUID(v string)`

SetSiteUUID sets SiteUUID field to given value.

### HasSiteUUID

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasSiteUUID() bool`

HasSiteUUID returns a boolean if a field has been set.

### GetSiteCode

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetSiteCode() string`

GetSiteCode returns the SiteCode field if non-nil, zero value otherwise.

### GetSiteCodeOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetSiteCodeOk() (*string, bool)`

GetSiteCodeOk returns a tuple with the SiteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCode

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetSiteCode(v string)`

SetSiteCode sets SiteCode field to given value.

### HasSiteCode

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasSiteCode() bool`

HasSiteCode returns a boolean if a field has been set.

### GetStartDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetApplicablePrices

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetApplicablePrices() []ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner`

GetApplicablePrices returns the ApplicablePrices field if non-nil, zero value otherwise.

### GetApplicablePricesOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetApplicablePricesOk() (*[]ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner, bool)`

GetApplicablePricesOk returns a tuple with the ApplicablePrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicablePrices

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetApplicablePrices(v []ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInnerApplicablePricesInner)`

SetApplicablePrices sets ApplicablePrices field to given value.

### HasApplicablePrices

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasApplicablePrices() bool`

HasApplicablePrices returns a boolean if a field has been set.

### GetServicePlanId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetServicePlanId() int64`

GetServicePlanId returns the ServicePlanId field if non-nil, zero value otherwise.

### GetServicePlanIdOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetServicePlanIdOk() (*int64, bool)`

GetServicePlanIdOk returns a tuple with the ServicePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetServicePlanId(v int64)`

SetServicePlanId sets ServicePlanId field to given value.

### HasServicePlanId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasServicePlanId() bool`

HasServicePlanId returns a boolean if a field has been set.

### GetServicePlanName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.

### HasServicePlanName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasServicePlanName() bool`

HasServicePlanName returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePoolName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *ListBillingServers200ResponseAllOfBillingInfoServersInnerUsagesInner) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



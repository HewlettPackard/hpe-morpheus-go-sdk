# ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**InstanceName** | Pointer to **string** |  | [optional] 
**ZoneName** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner**](ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner.md) |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **string** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**ServerExternalId** | Pointer to **string** |  | [optional] 
**ServerInternalId** | Pointer to **string** |  | [optional] 
**PlanName** | Pointer to **string** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**HourlyCost** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**PricesUsed** | Pointer to [**[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerPricesUsedInner**](ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerPricesUsedInner.md) |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**CreatedByUser** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 
**SiteUUID** | Pointer to **string** |  | [optional] 
**SiteCode** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ApplicablePrices** | Pointer to [**[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner**](ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner.md) |  | [optional] 
**ServicePlanId** | Pointer to **int64** |  | [optional] 
**ServicePlanName** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**ResourcePoolName** | Pointer to **string** |  | [optional] 

## Methods

### NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner

`func NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner() *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner`

NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner instantiates a new ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerWithDefaults

`func NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerWithDefaults() *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner`

NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerWithDefaults instantiates a new ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInstanceName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetZoneName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetAccountName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetVolumes

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetVolumes() []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetVolumesOk() (*[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetVolumes(v []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxCores

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetServerExternalId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetServerExternalId() string`

GetServerExternalId returns the ServerExternalId field if non-nil, zero value otherwise.

### GetServerExternalIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetServerExternalIdOk() (*string, bool)`

GetServerExternalIdOk returns a tuple with the ServerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExternalId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetServerExternalId(v string)`

SetServerExternalId sets ServerExternalId field to given value.

### HasServerExternalId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasServerExternalId() bool`

HasServerExternalId returns a boolean if a field has been set.

### GetServerInternalId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetServerInternalId() string`

GetServerInternalId returns the ServerInternalId field if non-nil, zero value otherwise.

### GetServerInternalIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetServerInternalIdOk() (*string, bool)`

GetServerInternalIdOk returns a tuple with the ServerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternalId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetServerInternalId(v string)`

SetServerInternalId sets ServerInternalId field to given value.

### HasServerInternalId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasServerInternalId() bool`

HasServerInternalId returns a boolean if a field has been set.

### GetPlanName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetHourlyCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetCurrency

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPricesUsed

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetPricesUsed() []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerPricesUsedInner`

GetPricesUsed returns the PricesUsed field if non-nil, zero value otherwise.

### GetPricesUsedOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetPricesUsedOk() (*[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerPricesUsedInner, bool)`

GetPricesUsedOk returns a tuple with the PricesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricesUsed

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetPricesUsed(v []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerPricesUsedInner)`

SetPricesUsed sets PricesUsed field to given value.

### HasPricesUsed

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasPricesUsed() bool`

HasPricesUsed returns a boolean if a field has been set.

### GetCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetCreatedByUser() string`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetCreatedByUserOk() (*string, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetCreatedByUser(v string)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetCreatedByUserId() int64`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetCreatedByUserIdOk() (*int64, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetCreatedByUserId(v int64)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteUUID

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetSiteUUID() string`

GetSiteUUID returns the SiteUUID field if non-nil, zero value otherwise.

### GetSiteUUIDOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetSiteUUIDOk() (*string, bool)`

GetSiteUUIDOk returns a tuple with the SiteUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUUID

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetSiteUUID(v string)`

SetSiteUUID sets SiteUUID field to given value.

### HasSiteUUID

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasSiteUUID() bool`

HasSiteUUID returns a boolean if a field has been set.

### GetSiteCode

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetSiteCode() string`

GetSiteCode returns the SiteCode field if non-nil, zero value otherwise.

### GetSiteCodeOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetSiteCodeOk() (*string, bool)`

GetSiteCodeOk returns a tuple with the SiteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCode

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetSiteCode(v string)`

SetSiteCode sets SiteCode field to given value.

### HasSiteCode

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasSiteCode() bool`

HasSiteCode returns a boolean if a field has been set.

### GetStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetApplicablePrices

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetApplicablePrices() []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner`

GetApplicablePrices returns the ApplicablePrices field if non-nil, zero value otherwise.

### GetApplicablePricesOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetApplicablePricesOk() (*[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner, bool)`

GetApplicablePricesOk returns a tuple with the ApplicablePrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicablePrices

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetApplicablePrices(v []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerApplicablePricesInner)`

SetApplicablePrices sets ApplicablePrices field to given value.

### HasApplicablePrices

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasApplicablePrices() bool`

HasApplicablePrices returns a boolean if a field has been set.

### GetServicePlanId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetServicePlanId() int64`

GetServicePlanId returns the ServicePlanId field if non-nil, zero value otherwise.

### GetServicePlanIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetServicePlanIdOk() (*int64, bool)`

GetServicePlanIdOk returns a tuple with the ServicePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetServicePlanId(v int64)`

SetServicePlanId sets ServicePlanId field to given value.

### HasServicePlanId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasServicePlanId() bool`

HasServicePlanId returns a boolean if a field has been set.

### GetServicePlanName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.

### HasServicePlanName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasServicePlanName() bool`

HasServicePlanName returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePoolName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInner) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



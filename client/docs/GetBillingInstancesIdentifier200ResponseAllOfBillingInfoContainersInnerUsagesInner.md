# GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner

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
**ApplicablePrices** | Pointer to [**[]GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInner**](GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInner.md) |  | [optional] 
**ServicePlanId** | Pointer to **int64** |  | [optional] 
**ServicePlanName** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**ResourcePoolName** | Pointer to **string** |  | [optional] 

## Methods

### NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner

`func NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner() *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner`

NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner instantiates a new GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerWithDefaults

`func NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerWithDefaults() *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner`

NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerWithDefaults instantiates a new GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInstanceName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetZoneName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetAccountName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetVolumes

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetVolumes() []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetVolumesOk() (*[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetVolumes(v []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxCores

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetServerExternalId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetServerExternalId() string`

GetServerExternalId returns the ServerExternalId field if non-nil, zero value otherwise.

### GetServerExternalIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetServerExternalIdOk() (*string, bool)`

GetServerExternalIdOk returns a tuple with the ServerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExternalId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetServerExternalId(v string)`

SetServerExternalId sets ServerExternalId field to given value.

### HasServerExternalId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasServerExternalId() bool`

HasServerExternalId returns a boolean if a field has been set.

### GetServerInternalId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetServerInternalId() string`

GetServerInternalId returns the ServerInternalId field if non-nil, zero value otherwise.

### GetServerInternalIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetServerInternalIdOk() (*string, bool)`

GetServerInternalIdOk returns a tuple with the ServerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternalId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetServerInternalId(v string)`

SetServerInternalId sets ServerInternalId field to given value.

### HasServerInternalId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasServerInternalId() bool`

HasServerInternalId returns a boolean if a field has been set.

### GetPlanName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetHourlyCost

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetCurrency

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPricesUsed

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetPricesUsed() []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerPricesUsedInner`

GetPricesUsed returns the PricesUsed field if non-nil, zero value otherwise.

### GetPricesUsedOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetPricesUsedOk() (*[]ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerPricesUsedInner, bool)`

GetPricesUsedOk returns a tuple with the PricesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricesUsed

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetPricesUsed(v []ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerPricesUsedInner)`

SetPricesUsed sets PricesUsed field to given value.

### HasPricesUsed

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasPricesUsed() bool`

HasPricesUsed returns a boolean if a field has been set.

### GetCost

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetCreatedByUser() string`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetCreatedByUserOk() (*string, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetCreatedByUser(v string)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetCreatedByUserId() int64`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetCreatedByUserIdOk() (*int64, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetCreatedByUserId(v int64)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteUUID

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetSiteUUID() string`

GetSiteUUID returns the SiteUUID field if non-nil, zero value otherwise.

### GetSiteUUIDOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetSiteUUIDOk() (*string, bool)`

GetSiteUUIDOk returns a tuple with the SiteUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUUID

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetSiteUUID(v string)`

SetSiteUUID sets SiteUUID field to given value.

### HasSiteUUID

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasSiteUUID() bool`

HasSiteUUID returns a boolean if a field has been set.

### GetSiteCode

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetSiteCode() string`

GetSiteCode returns the SiteCode field if non-nil, zero value otherwise.

### GetSiteCodeOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetSiteCodeOk() (*string, bool)`

GetSiteCodeOk returns a tuple with the SiteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCode

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetSiteCode(v string)`

SetSiteCode sets SiteCode field to given value.

### HasSiteCode

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasSiteCode() bool`

HasSiteCode returns a boolean if a field has been set.

### GetStartDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetApplicablePrices

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetApplicablePrices() []GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInner`

GetApplicablePrices returns the ApplicablePrices field if non-nil, zero value otherwise.

### GetApplicablePricesOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetApplicablePricesOk() (*[]GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInner, bool)`

GetApplicablePricesOk returns a tuple with the ApplicablePrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicablePrices

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetApplicablePrices(v []GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInner)`

SetApplicablePrices sets ApplicablePrices field to given value.

### HasApplicablePrices

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasApplicablePrices() bool`

HasApplicablePrices returns a boolean if a field has been set.

### GetServicePlanId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetServicePlanId() int64`

GetServicePlanId returns the ServicePlanId field if non-nil, zero value otherwise.

### GetServicePlanIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetServicePlanIdOk() (*int64, bool)`

GetServicePlanIdOk returns a tuple with the ServicePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetServicePlanId(v int64)`

SetServicePlanId sets ServicePlanId field to given value.

### HasServicePlanId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasServicePlanId() bool`

HasServicePlanId returns a boolean if a field has been set.

### GetServicePlanName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.

### HasServicePlanName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasServicePlanName() bool`

HasServicePlanName returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePoolName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetResourcePoolName() string`

GetResourcePoolName returns the ResourcePoolName field if non-nil, zero value otherwise.

### GetResourcePoolNameOk

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) GetResourcePoolNameOk() (*string, bool)`

GetResourcePoolNameOk returns a tuple with the ResourcePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) SetResourcePoolName(v string)`

SetResourcePoolName sets ResourcePoolName field to given value.

### HasResourcePoolName

`func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInner) HasResourcePoolName() bool`

HasResourcePoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



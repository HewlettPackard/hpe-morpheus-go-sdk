# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**OwnerId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Group** | Pointer to **map[string]interface{}** |  | [optional] 
**Cloud** | Pointer to [**ListInvoices200ResponseAllOfInvoicesInnerCloud**](ListInvoices200ResponseAllOfInvoicesInnerCloud.md) |  | [optional] 
**Instance** | Pointer to **map[string]interface{}** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**User** | Pointer to **map[string]interface{}** |  | [optional] 
**Plan** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**RefUuid** | Pointer to **string** |  | [optional] 
**RefName** | Pointer to **string** |  | [optional] 
**RefCategory** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 
**ResourceUuid** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**ResourceExternalId** | Pointer to **string** |  | [optional] 
**ResourceInternalId** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Estimate** | Pointer to **bool** |  | [optional] 
**SummaryInvoice** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**RefStart** | Pointer to **time.Time** |  | [optional] 
**RefEnd** | Pointer to **time.Time** |  | [optional] 
**EstimatedComputePrice** | Pointer to **float32** |  | [optional] 
**EstimatedComputeCost** | Pointer to **float32** |  | [optional] 
**EstimatedMemoryPrice** | Pointer to **float32** |  | [optional] 
**EstimatedMemoryCost** | Pointer to **float32** |  | [optional] 
**EstimatedStoragePrice** | Pointer to **float32** |  | [optional] 
**EstimatedStorageCost** | Pointer to **float32** |  | [optional] 
**EstimatedNetworkPrice** | Pointer to **float32** |  | [optional] 
**EstimatedNetworkCost** | Pointer to **float32** |  | [optional] 
**EstimatedLicensePrice** | Pointer to **float32** |  | [optional] 
**EstimatedLicenseCost** | Pointer to **float32** |  | [optional] 
**EstimatedExtraPrice** | Pointer to **float32** |  | [optional] 
**EstimatedExtraCost** | Pointer to **float32** |  | [optional] 
**EstimatedTotalPrice** | Pointer to **float32** |  | [optional] 
**EstimatedTotalCost** | Pointer to **float32** |  | [optional] 
**EstimatedRunningPrice** | Pointer to **float32** |  | [optional] 
**EstimatedRunningCost** | Pointer to **float32** |  | [optional] 
**EstimatedCurrency** | Pointer to **string** |  | [optional] 
**EstimatedConversionRate** | Pointer to **float32** |  | [optional] 
**ActualComputePrice** | Pointer to **float32** |  | [optional] 
**ActualComputeCost** | Pointer to **float32** |  | [optional] 
**ActualMemoryPrice** | Pointer to **float32** |  | [optional] 
**ActualMemoryCost** | Pointer to **float32** |  | [optional] 
**ActualStoragePrice** | Pointer to **float32** |  | [optional] 
**ActualStorageCost** | Pointer to **float32** |  | [optional] 
**ActualNetworkPrice** | Pointer to **float32** |  | [optional] 
**ActualNetworkCost** | Pointer to **float32** |  | [optional] 
**ActualLicensePrice** | Pointer to **float32** |  | [optional] 
**ActualLicenseCost** | Pointer to **float32** |  | [optional] 
**ActualExtraPrice** | Pointer to **float32** |  | [optional] 
**ActualExtraCost** | Pointer to **float32** |  | [optional] 
**ActualTotalPrice** | Pointer to **float32** |  | [optional] 
**ActualTotalCost** | Pointer to **float32** |  | [optional] 
**ActualRunningPrice** | Pointer to **float32** |  | [optional] 
**ActualRunningCost** | Pointer to **float32** |  | [optional] 
**ActualCurrency** | Pointer to **string** |  | [optional] 
**ActualConversionRate** | Pointer to **float32** |  | [optional] 
**ComputePrice** | Pointer to **float32** |  | [optional] 
**ComputeCost** | Pointer to **float32** |  | [optional] 
**MemoryPrice** | Pointer to **float32** |  | [optional] 
**MemoryCost** | Pointer to **float32** |  | [optional] 
**StoragePrice** | Pointer to **float32** |  | [optional] 
**StorageCost** | Pointer to **float32** |  | [optional] 
**NetworkPrice** | Pointer to **float32** |  | [optional] 
**NetworkCost** | Pointer to **float32** |  | [optional] 
**LicensePrice** | Pointer to **float32** |  | [optional] 
**LicenseCost** | Pointer to **float32** |  | [optional] 
**ExtraPrice** | Pointer to **float32** |  | [optional] 
**ExtraCost** | Pointer to **float32** |  | [optional] 
**TotalPrice** | Pointer to **float32** |  | [optional] 
**TotalCost** | Pointer to **float32** |  | [optional] 
**RunningPrice** | Pointer to **float32** |  | [optional] 
**RunningCost** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ConversionRate** | Pointer to **float32** |  | [optional] 
**CostType** | Pointer to **string** |  | [optional] 
**OffTime** | Pointer to **int64** |  | [optional] 
**PowerState** | Pointer to **string** |  | [optional] 
**PowerDate** | Pointer to **time.Time** |  | [optional] 
**RunningMultiplier** | Pointer to **float32** |  | [optional] 
**UsageType** | Pointer to **string** |  | [optional] 
**UsageCategory** | Pointer to **string** |  | [optional] 
**LastCostDate** | Pointer to **time.Time** |  | [optional] 
**LastActualDate** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LineItemCount** | Pointer to **int64** |  | [optional] 
**LineItems** | Pointer to [**[]ListInvoices200ResponseAllOfInvoicesInnerLineItemsInner**](ListInvoices200ResponseAllOfInvoicesInnerLineItemsInner.md) |  | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwnerId

`func (o *Invoice) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Invoice) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Invoice) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Invoice) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetAccount

`func (o *Invoice) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Invoice) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Invoice) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Invoice) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetGroup

`func (o *Invoice) GetGroup() map[string]interface{}`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Invoice) GetGroupOk() (*map[string]interface{}, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Invoice) SetGroup(v map[string]interface{})`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Invoice) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *Invoice) GetCloud() ListInvoices200ResponseAllOfInvoicesInnerCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *Invoice) GetCloudOk() (*ListInvoices200ResponseAllOfInvoicesInnerCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *Invoice) SetCloud(v ListInvoices200ResponseAllOfInvoicesInnerCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *Invoice) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetInstance

`func (o *Invoice) GetInstance() map[string]interface{}`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *Invoice) GetInstanceOk() (*map[string]interface{}, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *Invoice) SetInstance(v map[string]interface{})`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *Invoice) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetServer

`func (o *Invoice) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *Invoice) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *Invoice) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *Invoice) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetCluster

`func (o *Invoice) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Invoice) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Invoice) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Invoice) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetUser

`func (o *Invoice) GetUser() map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Invoice) GetUserOk() (*map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Invoice) SetUser(v map[string]interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *Invoice) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPlan

`func (o *Invoice) GetPlan() map[string]interface{}`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Invoice) GetPlanOk() (*map[string]interface{}, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Invoice) SetPlan(v map[string]interface{})`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Invoice) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetTags

`func (o *Invoice) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Invoice) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Invoice) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Invoice) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProject

`func (o *Invoice) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Invoice) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Invoice) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Invoice) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRefType

`func (o *Invoice) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *Invoice) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *Invoice) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *Invoice) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *Invoice) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Invoice) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Invoice) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Invoice) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefUuid

`func (o *Invoice) GetRefUuid() string`

GetRefUuid returns the RefUuid field if non-nil, zero value otherwise.

### GetRefUuidOk

`func (o *Invoice) GetRefUuidOk() (*string, bool)`

GetRefUuidOk returns a tuple with the RefUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUuid

`func (o *Invoice) SetRefUuid(v string)`

SetRefUuid sets RefUuid field to given value.

### HasRefUuid

`func (o *Invoice) HasRefUuid() bool`

HasRefUuid returns a boolean if a field has been set.

### GetRefName

`func (o *Invoice) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *Invoice) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *Invoice) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *Invoice) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetRefCategory

`func (o *Invoice) GetRefCategory() string`

GetRefCategory returns the RefCategory field if non-nil, zero value otherwise.

### GetRefCategoryOk

`func (o *Invoice) GetRefCategoryOk() (*string, bool)`

GetRefCategoryOk returns a tuple with the RefCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCategory

`func (o *Invoice) SetRefCategory(v string)`

SetRefCategory sets RefCategory field to given value.

### HasRefCategory

`func (o *Invoice) HasRefCategory() bool`

HasRefCategory returns a boolean if a field has been set.

### GetResourceId

`func (o *Invoice) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Invoice) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Invoice) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Invoice) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceUuid

`func (o *Invoice) GetResourceUuid() string`

GetResourceUuid returns the ResourceUuid field if non-nil, zero value otherwise.

### GetResourceUuidOk

`func (o *Invoice) GetResourceUuidOk() (*string, bool)`

GetResourceUuidOk returns a tuple with the ResourceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUuid

`func (o *Invoice) SetResourceUuid(v string)`

SetResourceUuid sets ResourceUuid field to given value.

### HasResourceUuid

`func (o *Invoice) HasResourceUuid() bool`

HasResourceUuid returns a boolean if a field has been set.

### GetResourceType

`func (o *Invoice) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Invoice) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Invoice) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Invoice) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceName

`func (o *Invoice) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Invoice) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Invoice) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *Invoice) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceExternalId

`func (o *Invoice) GetResourceExternalId() string`

GetResourceExternalId returns the ResourceExternalId field if non-nil, zero value otherwise.

### GetResourceExternalIdOk

`func (o *Invoice) GetResourceExternalIdOk() (*string, bool)`

GetResourceExternalIdOk returns a tuple with the ResourceExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceExternalId

`func (o *Invoice) SetResourceExternalId(v string)`

SetResourceExternalId sets ResourceExternalId field to given value.

### HasResourceExternalId

`func (o *Invoice) HasResourceExternalId() bool`

HasResourceExternalId returns a boolean if a field has been set.

### GetResourceInternalId

`func (o *Invoice) GetResourceInternalId() string`

GetResourceInternalId returns the ResourceInternalId field if non-nil, zero value otherwise.

### GetResourceInternalIdOk

`func (o *Invoice) GetResourceInternalIdOk() (*string, bool)`

GetResourceInternalIdOk returns a tuple with the ResourceInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInternalId

`func (o *Invoice) SetResourceInternalId(v string)`

SetResourceInternalId sets ResourceInternalId field to given value.

### HasResourceInternalId

`func (o *Invoice) HasResourceInternalId() bool`

HasResourceInternalId returns a boolean if a field has been set.

### GetInterval

`func (o *Invoice) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Invoice) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Invoice) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Invoice) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetPeriod

`func (o *Invoice) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Invoice) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Invoice) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *Invoice) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetEstimate

`func (o *Invoice) GetEstimate() bool`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *Invoice) GetEstimateOk() (*bool, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *Invoice) SetEstimate(v bool)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *Invoice) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### GetSummaryInvoice

`func (o *Invoice) GetSummaryInvoice() bool`

GetSummaryInvoice returns the SummaryInvoice field if non-nil, zero value otherwise.

### GetSummaryInvoiceOk

`func (o *Invoice) GetSummaryInvoiceOk() (*bool, bool)`

GetSummaryInvoiceOk returns a tuple with the SummaryInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryInvoice

`func (o *Invoice) SetSummaryInvoice(v bool)`

SetSummaryInvoice sets SummaryInvoice field to given value.

### HasSummaryInvoice

`func (o *Invoice) HasSummaryInvoice() bool`

HasSummaryInvoice returns a boolean if a field has been set.

### GetActive

`func (o *Invoice) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Invoice) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Invoice) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Invoice) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStartDate

`func (o *Invoice) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Invoice) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Invoice) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Invoice) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Invoice) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Invoice) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Invoice) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Invoice) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetRefStart

`func (o *Invoice) GetRefStart() time.Time`

GetRefStart returns the RefStart field if non-nil, zero value otherwise.

### GetRefStartOk

`func (o *Invoice) GetRefStartOk() (*time.Time, bool)`

GetRefStartOk returns a tuple with the RefStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefStart

`func (o *Invoice) SetRefStart(v time.Time)`

SetRefStart sets RefStart field to given value.

### HasRefStart

`func (o *Invoice) HasRefStart() bool`

HasRefStart returns a boolean if a field has been set.

### GetRefEnd

`func (o *Invoice) GetRefEnd() time.Time`

GetRefEnd returns the RefEnd field if non-nil, zero value otherwise.

### GetRefEndOk

`func (o *Invoice) GetRefEndOk() (*time.Time, bool)`

GetRefEndOk returns a tuple with the RefEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefEnd

`func (o *Invoice) SetRefEnd(v time.Time)`

SetRefEnd sets RefEnd field to given value.

### HasRefEnd

`func (o *Invoice) HasRefEnd() bool`

HasRefEnd returns a boolean if a field has been set.

### GetEstimatedComputePrice

`func (o *Invoice) GetEstimatedComputePrice() float32`

GetEstimatedComputePrice returns the EstimatedComputePrice field if non-nil, zero value otherwise.

### GetEstimatedComputePriceOk

`func (o *Invoice) GetEstimatedComputePriceOk() (*float32, bool)`

GetEstimatedComputePriceOk returns a tuple with the EstimatedComputePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedComputePrice

`func (o *Invoice) SetEstimatedComputePrice(v float32)`

SetEstimatedComputePrice sets EstimatedComputePrice field to given value.

### HasEstimatedComputePrice

`func (o *Invoice) HasEstimatedComputePrice() bool`

HasEstimatedComputePrice returns a boolean if a field has been set.

### GetEstimatedComputeCost

`func (o *Invoice) GetEstimatedComputeCost() float32`

GetEstimatedComputeCost returns the EstimatedComputeCost field if non-nil, zero value otherwise.

### GetEstimatedComputeCostOk

`func (o *Invoice) GetEstimatedComputeCostOk() (*float32, bool)`

GetEstimatedComputeCostOk returns a tuple with the EstimatedComputeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedComputeCost

`func (o *Invoice) SetEstimatedComputeCost(v float32)`

SetEstimatedComputeCost sets EstimatedComputeCost field to given value.

### HasEstimatedComputeCost

`func (o *Invoice) HasEstimatedComputeCost() bool`

HasEstimatedComputeCost returns a boolean if a field has been set.

### GetEstimatedMemoryPrice

`func (o *Invoice) GetEstimatedMemoryPrice() float32`

GetEstimatedMemoryPrice returns the EstimatedMemoryPrice field if non-nil, zero value otherwise.

### GetEstimatedMemoryPriceOk

`func (o *Invoice) GetEstimatedMemoryPriceOk() (*float32, bool)`

GetEstimatedMemoryPriceOk returns a tuple with the EstimatedMemoryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMemoryPrice

`func (o *Invoice) SetEstimatedMemoryPrice(v float32)`

SetEstimatedMemoryPrice sets EstimatedMemoryPrice field to given value.

### HasEstimatedMemoryPrice

`func (o *Invoice) HasEstimatedMemoryPrice() bool`

HasEstimatedMemoryPrice returns a boolean if a field has been set.

### GetEstimatedMemoryCost

`func (o *Invoice) GetEstimatedMemoryCost() float32`

GetEstimatedMemoryCost returns the EstimatedMemoryCost field if non-nil, zero value otherwise.

### GetEstimatedMemoryCostOk

`func (o *Invoice) GetEstimatedMemoryCostOk() (*float32, bool)`

GetEstimatedMemoryCostOk returns a tuple with the EstimatedMemoryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMemoryCost

`func (o *Invoice) SetEstimatedMemoryCost(v float32)`

SetEstimatedMemoryCost sets EstimatedMemoryCost field to given value.

### HasEstimatedMemoryCost

`func (o *Invoice) HasEstimatedMemoryCost() bool`

HasEstimatedMemoryCost returns a boolean if a field has been set.

### GetEstimatedStoragePrice

`func (o *Invoice) GetEstimatedStoragePrice() float32`

GetEstimatedStoragePrice returns the EstimatedStoragePrice field if non-nil, zero value otherwise.

### GetEstimatedStoragePriceOk

`func (o *Invoice) GetEstimatedStoragePriceOk() (*float32, bool)`

GetEstimatedStoragePriceOk returns a tuple with the EstimatedStoragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStoragePrice

`func (o *Invoice) SetEstimatedStoragePrice(v float32)`

SetEstimatedStoragePrice sets EstimatedStoragePrice field to given value.

### HasEstimatedStoragePrice

`func (o *Invoice) HasEstimatedStoragePrice() bool`

HasEstimatedStoragePrice returns a boolean if a field has been set.

### GetEstimatedStorageCost

`func (o *Invoice) GetEstimatedStorageCost() float32`

GetEstimatedStorageCost returns the EstimatedStorageCost field if non-nil, zero value otherwise.

### GetEstimatedStorageCostOk

`func (o *Invoice) GetEstimatedStorageCostOk() (*float32, bool)`

GetEstimatedStorageCostOk returns a tuple with the EstimatedStorageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStorageCost

`func (o *Invoice) SetEstimatedStorageCost(v float32)`

SetEstimatedStorageCost sets EstimatedStorageCost field to given value.

### HasEstimatedStorageCost

`func (o *Invoice) HasEstimatedStorageCost() bool`

HasEstimatedStorageCost returns a boolean if a field has been set.

### GetEstimatedNetworkPrice

`func (o *Invoice) GetEstimatedNetworkPrice() float32`

GetEstimatedNetworkPrice returns the EstimatedNetworkPrice field if non-nil, zero value otherwise.

### GetEstimatedNetworkPriceOk

`func (o *Invoice) GetEstimatedNetworkPriceOk() (*float32, bool)`

GetEstimatedNetworkPriceOk returns a tuple with the EstimatedNetworkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedNetworkPrice

`func (o *Invoice) SetEstimatedNetworkPrice(v float32)`

SetEstimatedNetworkPrice sets EstimatedNetworkPrice field to given value.

### HasEstimatedNetworkPrice

`func (o *Invoice) HasEstimatedNetworkPrice() bool`

HasEstimatedNetworkPrice returns a boolean if a field has been set.

### GetEstimatedNetworkCost

`func (o *Invoice) GetEstimatedNetworkCost() float32`

GetEstimatedNetworkCost returns the EstimatedNetworkCost field if non-nil, zero value otherwise.

### GetEstimatedNetworkCostOk

`func (o *Invoice) GetEstimatedNetworkCostOk() (*float32, bool)`

GetEstimatedNetworkCostOk returns a tuple with the EstimatedNetworkCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedNetworkCost

`func (o *Invoice) SetEstimatedNetworkCost(v float32)`

SetEstimatedNetworkCost sets EstimatedNetworkCost field to given value.

### HasEstimatedNetworkCost

`func (o *Invoice) HasEstimatedNetworkCost() bool`

HasEstimatedNetworkCost returns a boolean if a field has been set.

### GetEstimatedLicensePrice

`func (o *Invoice) GetEstimatedLicensePrice() float32`

GetEstimatedLicensePrice returns the EstimatedLicensePrice field if non-nil, zero value otherwise.

### GetEstimatedLicensePriceOk

`func (o *Invoice) GetEstimatedLicensePriceOk() (*float32, bool)`

GetEstimatedLicensePriceOk returns a tuple with the EstimatedLicensePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedLicensePrice

`func (o *Invoice) SetEstimatedLicensePrice(v float32)`

SetEstimatedLicensePrice sets EstimatedLicensePrice field to given value.

### HasEstimatedLicensePrice

`func (o *Invoice) HasEstimatedLicensePrice() bool`

HasEstimatedLicensePrice returns a boolean if a field has been set.

### GetEstimatedLicenseCost

`func (o *Invoice) GetEstimatedLicenseCost() float32`

GetEstimatedLicenseCost returns the EstimatedLicenseCost field if non-nil, zero value otherwise.

### GetEstimatedLicenseCostOk

`func (o *Invoice) GetEstimatedLicenseCostOk() (*float32, bool)`

GetEstimatedLicenseCostOk returns a tuple with the EstimatedLicenseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedLicenseCost

`func (o *Invoice) SetEstimatedLicenseCost(v float32)`

SetEstimatedLicenseCost sets EstimatedLicenseCost field to given value.

### HasEstimatedLicenseCost

`func (o *Invoice) HasEstimatedLicenseCost() bool`

HasEstimatedLicenseCost returns a boolean if a field has been set.

### GetEstimatedExtraPrice

`func (o *Invoice) GetEstimatedExtraPrice() float32`

GetEstimatedExtraPrice returns the EstimatedExtraPrice field if non-nil, zero value otherwise.

### GetEstimatedExtraPriceOk

`func (o *Invoice) GetEstimatedExtraPriceOk() (*float32, bool)`

GetEstimatedExtraPriceOk returns a tuple with the EstimatedExtraPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedExtraPrice

`func (o *Invoice) SetEstimatedExtraPrice(v float32)`

SetEstimatedExtraPrice sets EstimatedExtraPrice field to given value.

### HasEstimatedExtraPrice

`func (o *Invoice) HasEstimatedExtraPrice() bool`

HasEstimatedExtraPrice returns a boolean if a field has been set.

### GetEstimatedExtraCost

`func (o *Invoice) GetEstimatedExtraCost() float32`

GetEstimatedExtraCost returns the EstimatedExtraCost field if non-nil, zero value otherwise.

### GetEstimatedExtraCostOk

`func (o *Invoice) GetEstimatedExtraCostOk() (*float32, bool)`

GetEstimatedExtraCostOk returns a tuple with the EstimatedExtraCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedExtraCost

`func (o *Invoice) SetEstimatedExtraCost(v float32)`

SetEstimatedExtraCost sets EstimatedExtraCost field to given value.

### HasEstimatedExtraCost

`func (o *Invoice) HasEstimatedExtraCost() bool`

HasEstimatedExtraCost returns a boolean if a field has been set.

### GetEstimatedTotalPrice

`func (o *Invoice) GetEstimatedTotalPrice() float32`

GetEstimatedTotalPrice returns the EstimatedTotalPrice field if non-nil, zero value otherwise.

### GetEstimatedTotalPriceOk

`func (o *Invoice) GetEstimatedTotalPriceOk() (*float32, bool)`

GetEstimatedTotalPriceOk returns a tuple with the EstimatedTotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTotalPrice

`func (o *Invoice) SetEstimatedTotalPrice(v float32)`

SetEstimatedTotalPrice sets EstimatedTotalPrice field to given value.

### HasEstimatedTotalPrice

`func (o *Invoice) HasEstimatedTotalPrice() bool`

HasEstimatedTotalPrice returns a boolean if a field has been set.

### GetEstimatedTotalCost

`func (o *Invoice) GetEstimatedTotalCost() float32`

GetEstimatedTotalCost returns the EstimatedTotalCost field if non-nil, zero value otherwise.

### GetEstimatedTotalCostOk

`func (o *Invoice) GetEstimatedTotalCostOk() (*float32, bool)`

GetEstimatedTotalCostOk returns a tuple with the EstimatedTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTotalCost

`func (o *Invoice) SetEstimatedTotalCost(v float32)`

SetEstimatedTotalCost sets EstimatedTotalCost field to given value.

### HasEstimatedTotalCost

`func (o *Invoice) HasEstimatedTotalCost() bool`

HasEstimatedTotalCost returns a boolean if a field has been set.

### GetEstimatedRunningPrice

`func (o *Invoice) GetEstimatedRunningPrice() float32`

GetEstimatedRunningPrice returns the EstimatedRunningPrice field if non-nil, zero value otherwise.

### GetEstimatedRunningPriceOk

`func (o *Invoice) GetEstimatedRunningPriceOk() (*float32, bool)`

GetEstimatedRunningPriceOk returns a tuple with the EstimatedRunningPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRunningPrice

`func (o *Invoice) SetEstimatedRunningPrice(v float32)`

SetEstimatedRunningPrice sets EstimatedRunningPrice field to given value.

### HasEstimatedRunningPrice

`func (o *Invoice) HasEstimatedRunningPrice() bool`

HasEstimatedRunningPrice returns a boolean if a field has been set.

### GetEstimatedRunningCost

`func (o *Invoice) GetEstimatedRunningCost() float32`

GetEstimatedRunningCost returns the EstimatedRunningCost field if non-nil, zero value otherwise.

### GetEstimatedRunningCostOk

`func (o *Invoice) GetEstimatedRunningCostOk() (*float32, bool)`

GetEstimatedRunningCostOk returns a tuple with the EstimatedRunningCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRunningCost

`func (o *Invoice) SetEstimatedRunningCost(v float32)`

SetEstimatedRunningCost sets EstimatedRunningCost field to given value.

### HasEstimatedRunningCost

`func (o *Invoice) HasEstimatedRunningCost() bool`

HasEstimatedRunningCost returns a boolean if a field has been set.

### GetEstimatedCurrency

`func (o *Invoice) GetEstimatedCurrency() string`

GetEstimatedCurrency returns the EstimatedCurrency field if non-nil, zero value otherwise.

### GetEstimatedCurrencyOk

`func (o *Invoice) GetEstimatedCurrencyOk() (*string, bool)`

GetEstimatedCurrencyOk returns a tuple with the EstimatedCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCurrency

`func (o *Invoice) SetEstimatedCurrency(v string)`

SetEstimatedCurrency sets EstimatedCurrency field to given value.

### HasEstimatedCurrency

`func (o *Invoice) HasEstimatedCurrency() bool`

HasEstimatedCurrency returns a boolean if a field has been set.

### GetEstimatedConversionRate

`func (o *Invoice) GetEstimatedConversionRate() float32`

GetEstimatedConversionRate returns the EstimatedConversionRate field if non-nil, zero value otherwise.

### GetEstimatedConversionRateOk

`func (o *Invoice) GetEstimatedConversionRateOk() (*float32, bool)`

GetEstimatedConversionRateOk returns a tuple with the EstimatedConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedConversionRate

`func (o *Invoice) SetEstimatedConversionRate(v float32)`

SetEstimatedConversionRate sets EstimatedConversionRate field to given value.

### HasEstimatedConversionRate

`func (o *Invoice) HasEstimatedConversionRate() bool`

HasEstimatedConversionRate returns a boolean if a field has been set.

### GetActualComputePrice

`func (o *Invoice) GetActualComputePrice() float32`

GetActualComputePrice returns the ActualComputePrice field if non-nil, zero value otherwise.

### GetActualComputePriceOk

`func (o *Invoice) GetActualComputePriceOk() (*float32, bool)`

GetActualComputePriceOk returns a tuple with the ActualComputePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualComputePrice

`func (o *Invoice) SetActualComputePrice(v float32)`

SetActualComputePrice sets ActualComputePrice field to given value.

### HasActualComputePrice

`func (o *Invoice) HasActualComputePrice() bool`

HasActualComputePrice returns a boolean if a field has been set.

### GetActualComputeCost

`func (o *Invoice) GetActualComputeCost() float32`

GetActualComputeCost returns the ActualComputeCost field if non-nil, zero value otherwise.

### GetActualComputeCostOk

`func (o *Invoice) GetActualComputeCostOk() (*float32, bool)`

GetActualComputeCostOk returns a tuple with the ActualComputeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualComputeCost

`func (o *Invoice) SetActualComputeCost(v float32)`

SetActualComputeCost sets ActualComputeCost field to given value.

### HasActualComputeCost

`func (o *Invoice) HasActualComputeCost() bool`

HasActualComputeCost returns a boolean if a field has been set.

### GetActualMemoryPrice

`func (o *Invoice) GetActualMemoryPrice() float32`

GetActualMemoryPrice returns the ActualMemoryPrice field if non-nil, zero value otherwise.

### GetActualMemoryPriceOk

`func (o *Invoice) GetActualMemoryPriceOk() (*float32, bool)`

GetActualMemoryPriceOk returns a tuple with the ActualMemoryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualMemoryPrice

`func (o *Invoice) SetActualMemoryPrice(v float32)`

SetActualMemoryPrice sets ActualMemoryPrice field to given value.

### HasActualMemoryPrice

`func (o *Invoice) HasActualMemoryPrice() bool`

HasActualMemoryPrice returns a boolean if a field has been set.

### GetActualMemoryCost

`func (o *Invoice) GetActualMemoryCost() float32`

GetActualMemoryCost returns the ActualMemoryCost field if non-nil, zero value otherwise.

### GetActualMemoryCostOk

`func (o *Invoice) GetActualMemoryCostOk() (*float32, bool)`

GetActualMemoryCostOk returns a tuple with the ActualMemoryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualMemoryCost

`func (o *Invoice) SetActualMemoryCost(v float32)`

SetActualMemoryCost sets ActualMemoryCost field to given value.

### HasActualMemoryCost

`func (o *Invoice) HasActualMemoryCost() bool`

HasActualMemoryCost returns a boolean if a field has been set.

### GetActualStoragePrice

`func (o *Invoice) GetActualStoragePrice() float32`

GetActualStoragePrice returns the ActualStoragePrice field if non-nil, zero value otherwise.

### GetActualStoragePriceOk

`func (o *Invoice) GetActualStoragePriceOk() (*float32, bool)`

GetActualStoragePriceOk returns a tuple with the ActualStoragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualStoragePrice

`func (o *Invoice) SetActualStoragePrice(v float32)`

SetActualStoragePrice sets ActualStoragePrice field to given value.

### HasActualStoragePrice

`func (o *Invoice) HasActualStoragePrice() bool`

HasActualStoragePrice returns a boolean if a field has been set.

### GetActualStorageCost

`func (o *Invoice) GetActualStorageCost() float32`

GetActualStorageCost returns the ActualStorageCost field if non-nil, zero value otherwise.

### GetActualStorageCostOk

`func (o *Invoice) GetActualStorageCostOk() (*float32, bool)`

GetActualStorageCostOk returns a tuple with the ActualStorageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualStorageCost

`func (o *Invoice) SetActualStorageCost(v float32)`

SetActualStorageCost sets ActualStorageCost field to given value.

### HasActualStorageCost

`func (o *Invoice) HasActualStorageCost() bool`

HasActualStorageCost returns a boolean if a field has been set.

### GetActualNetworkPrice

`func (o *Invoice) GetActualNetworkPrice() float32`

GetActualNetworkPrice returns the ActualNetworkPrice field if non-nil, zero value otherwise.

### GetActualNetworkPriceOk

`func (o *Invoice) GetActualNetworkPriceOk() (*float32, bool)`

GetActualNetworkPriceOk returns a tuple with the ActualNetworkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualNetworkPrice

`func (o *Invoice) SetActualNetworkPrice(v float32)`

SetActualNetworkPrice sets ActualNetworkPrice field to given value.

### HasActualNetworkPrice

`func (o *Invoice) HasActualNetworkPrice() bool`

HasActualNetworkPrice returns a boolean if a field has been set.

### GetActualNetworkCost

`func (o *Invoice) GetActualNetworkCost() float32`

GetActualNetworkCost returns the ActualNetworkCost field if non-nil, zero value otherwise.

### GetActualNetworkCostOk

`func (o *Invoice) GetActualNetworkCostOk() (*float32, bool)`

GetActualNetworkCostOk returns a tuple with the ActualNetworkCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualNetworkCost

`func (o *Invoice) SetActualNetworkCost(v float32)`

SetActualNetworkCost sets ActualNetworkCost field to given value.

### HasActualNetworkCost

`func (o *Invoice) HasActualNetworkCost() bool`

HasActualNetworkCost returns a boolean if a field has been set.

### GetActualLicensePrice

`func (o *Invoice) GetActualLicensePrice() float32`

GetActualLicensePrice returns the ActualLicensePrice field if non-nil, zero value otherwise.

### GetActualLicensePriceOk

`func (o *Invoice) GetActualLicensePriceOk() (*float32, bool)`

GetActualLicensePriceOk returns a tuple with the ActualLicensePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualLicensePrice

`func (o *Invoice) SetActualLicensePrice(v float32)`

SetActualLicensePrice sets ActualLicensePrice field to given value.

### HasActualLicensePrice

`func (o *Invoice) HasActualLicensePrice() bool`

HasActualLicensePrice returns a boolean if a field has been set.

### GetActualLicenseCost

`func (o *Invoice) GetActualLicenseCost() float32`

GetActualLicenseCost returns the ActualLicenseCost field if non-nil, zero value otherwise.

### GetActualLicenseCostOk

`func (o *Invoice) GetActualLicenseCostOk() (*float32, bool)`

GetActualLicenseCostOk returns a tuple with the ActualLicenseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualLicenseCost

`func (o *Invoice) SetActualLicenseCost(v float32)`

SetActualLicenseCost sets ActualLicenseCost field to given value.

### HasActualLicenseCost

`func (o *Invoice) HasActualLicenseCost() bool`

HasActualLicenseCost returns a boolean if a field has been set.

### GetActualExtraPrice

`func (o *Invoice) GetActualExtraPrice() float32`

GetActualExtraPrice returns the ActualExtraPrice field if non-nil, zero value otherwise.

### GetActualExtraPriceOk

`func (o *Invoice) GetActualExtraPriceOk() (*float32, bool)`

GetActualExtraPriceOk returns a tuple with the ActualExtraPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualExtraPrice

`func (o *Invoice) SetActualExtraPrice(v float32)`

SetActualExtraPrice sets ActualExtraPrice field to given value.

### HasActualExtraPrice

`func (o *Invoice) HasActualExtraPrice() bool`

HasActualExtraPrice returns a boolean if a field has been set.

### GetActualExtraCost

`func (o *Invoice) GetActualExtraCost() float32`

GetActualExtraCost returns the ActualExtraCost field if non-nil, zero value otherwise.

### GetActualExtraCostOk

`func (o *Invoice) GetActualExtraCostOk() (*float32, bool)`

GetActualExtraCostOk returns a tuple with the ActualExtraCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualExtraCost

`func (o *Invoice) SetActualExtraCost(v float32)`

SetActualExtraCost sets ActualExtraCost field to given value.

### HasActualExtraCost

`func (o *Invoice) HasActualExtraCost() bool`

HasActualExtraCost returns a boolean if a field has been set.

### GetActualTotalPrice

`func (o *Invoice) GetActualTotalPrice() float32`

GetActualTotalPrice returns the ActualTotalPrice field if non-nil, zero value otherwise.

### GetActualTotalPriceOk

`func (o *Invoice) GetActualTotalPriceOk() (*float32, bool)`

GetActualTotalPriceOk returns a tuple with the ActualTotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualTotalPrice

`func (o *Invoice) SetActualTotalPrice(v float32)`

SetActualTotalPrice sets ActualTotalPrice field to given value.

### HasActualTotalPrice

`func (o *Invoice) HasActualTotalPrice() bool`

HasActualTotalPrice returns a boolean if a field has been set.

### GetActualTotalCost

`func (o *Invoice) GetActualTotalCost() float32`

GetActualTotalCost returns the ActualTotalCost field if non-nil, zero value otherwise.

### GetActualTotalCostOk

`func (o *Invoice) GetActualTotalCostOk() (*float32, bool)`

GetActualTotalCostOk returns a tuple with the ActualTotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualTotalCost

`func (o *Invoice) SetActualTotalCost(v float32)`

SetActualTotalCost sets ActualTotalCost field to given value.

### HasActualTotalCost

`func (o *Invoice) HasActualTotalCost() bool`

HasActualTotalCost returns a boolean if a field has been set.

### GetActualRunningPrice

`func (o *Invoice) GetActualRunningPrice() float32`

GetActualRunningPrice returns the ActualRunningPrice field if non-nil, zero value otherwise.

### GetActualRunningPriceOk

`func (o *Invoice) GetActualRunningPriceOk() (*float32, bool)`

GetActualRunningPriceOk returns a tuple with the ActualRunningPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualRunningPrice

`func (o *Invoice) SetActualRunningPrice(v float32)`

SetActualRunningPrice sets ActualRunningPrice field to given value.

### HasActualRunningPrice

`func (o *Invoice) HasActualRunningPrice() bool`

HasActualRunningPrice returns a boolean if a field has been set.

### GetActualRunningCost

`func (o *Invoice) GetActualRunningCost() float32`

GetActualRunningCost returns the ActualRunningCost field if non-nil, zero value otherwise.

### GetActualRunningCostOk

`func (o *Invoice) GetActualRunningCostOk() (*float32, bool)`

GetActualRunningCostOk returns a tuple with the ActualRunningCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualRunningCost

`func (o *Invoice) SetActualRunningCost(v float32)`

SetActualRunningCost sets ActualRunningCost field to given value.

### HasActualRunningCost

`func (o *Invoice) HasActualRunningCost() bool`

HasActualRunningCost returns a boolean if a field has been set.

### GetActualCurrency

`func (o *Invoice) GetActualCurrency() string`

GetActualCurrency returns the ActualCurrency field if non-nil, zero value otherwise.

### GetActualCurrencyOk

`func (o *Invoice) GetActualCurrencyOk() (*string, bool)`

GetActualCurrencyOk returns a tuple with the ActualCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCurrency

`func (o *Invoice) SetActualCurrency(v string)`

SetActualCurrency sets ActualCurrency field to given value.

### HasActualCurrency

`func (o *Invoice) HasActualCurrency() bool`

HasActualCurrency returns a boolean if a field has been set.

### GetActualConversionRate

`func (o *Invoice) GetActualConversionRate() float32`

GetActualConversionRate returns the ActualConversionRate field if non-nil, zero value otherwise.

### GetActualConversionRateOk

`func (o *Invoice) GetActualConversionRateOk() (*float32, bool)`

GetActualConversionRateOk returns a tuple with the ActualConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualConversionRate

`func (o *Invoice) SetActualConversionRate(v float32)`

SetActualConversionRate sets ActualConversionRate field to given value.

### HasActualConversionRate

`func (o *Invoice) HasActualConversionRate() bool`

HasActualConversionRate returns a boolean if a field has been set.

### GetComputePrice

`func (o *Invoice) GetComputePrice() float32`

GetComputePrice returns the ComputePrice field if non-nil, zero value otherwise.

### GetComputePriceOk

`func (o *Invoice) GetComputePriceOk() (*float32, bool)`

GetComputePriceOk returns a tuple with the ComputePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePrice

`func (o *Invoice) SetComputePrice(v float32)`

SetComputePrice sets ComputePrice field to given value.

### HasComputePrice

`func (o *Invoice) HasComputePrice() bool`

HasComputePrice returns a boolean if a field has been set.

### GetComputeCost

`func (o *Invoice) GetComputeCost() float32`

GetComputeCost returns the ComputeCost field if non-nil, zero value otherwise.

### GetComputeCostOk

`func (o *Invoice) GetComputeCostOk() (*float32, bool)`

GetComputeCostOk returns a tuple with the ComputeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeCost

`func (o *Invoice) SetComputeCost(v float32)`

SetComputeCost sets ComputeCost field to given value.

### HasComputeCost

`func (o *Invoice) HasComputeCost() bool`

HasComputeCost returns a boolean if a field has been set.

### GetMemoryPrice

`func (o *Invoice) GetMemoryPrice() float32`

GetMemoryPrice returns the MemoryPrice field if non-nil, zero value otherwise.

### GetMemoryPriceOk

`func (o *Invoice) GetMemoryPriceOk() (*float32, bool)`

GetMemoryPriceOk returns a tuple with the MemoryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPrice

`func (o *Invoice) SetMemoryPrice(v float32)`

SetMemoryPrice sets MemoryPrice field to given value.

### HasMemoryPrice

`func (o *Invoice) HasMemoryPrice() bool`

HasMemoryPrice returns a boolean if a field has been set.

### GetMemoryCost

`func (o *Invoice) GetMemoryCost() float32`

GetMemoryCost returns the MemoryCost field if non-nil, zero value otherwise.

### GetMemoryCostOk

`func (o *Invoice) GetMemoryCostOk() (*float32, bool)`

GetMemoryCostOk returns a tuple with the MemoryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCost

`func (o *Invoice) SetMemoryCost(v float32)`

SetMemoryCost sets MemoryCost field to given value.

### HasMemoryCost

`func (o *Invoice) HasMemoryCost() bool`

HasMemoryCost returns a boolean if a field has been set.

### GetStoragePrice

`func (o *Invoice) GetStoragePrice() float32`

GetStoragePrice returns the StoragePrice field if non-nil, zero value otherwise.

### GetStoragePriceOk

`func (o *Invoice) GetStoragePriceOk() (*float32, bool)`

GetStoragePriceOk returns a tuple with the StoragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePrice

`func (o *Invoice) SetStoragePrice(v float32)`

SetStoragePrice sets StoragePrice field to given value.

### HasStoragePrice

`func (o *Invoice) HasStoragePrice() bool`

HasStoragePrice returns a boolean if a field has been set.

### GetStorageCost

`func (o *Invoice) GetStorageCost() float32`

GetStorageCost returns the StorageCost field if non-nil, zero value otherwise.

### GetStorageCostOk

`func (o *Invoice) GetStorageCostOk() (*float32, bool)`

GetStorageCostOk returns a tuple with the StorageCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCost

`func (o *Invoice) SetStorageCost(v float32)`

SetStorageCost sets StorageCost field to given value.

### HasStorageCost

`func (o *Invoice) HasStorageCost() bool`

HasStorageCost returns a boolean if a field has been set.

### GetNetworkPrice

`func (o *Invoice) GetNetworkPrice() float32`

GetNetworkPrice returns the NetworkPrice field if non-nil, zero value otherwise.

### GetNetworkPriceOk

`func (o *Invoice) GetNetworkPriceOk() (*float32, bool)`

GetNetworkPriceOk returns a tuple with the NetworkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPrice

`func (o *Invoice) SetNetworkPrice(v float32)`

SetNetworkPrice sets NetworkPrice field to given value.

### HasNetworkPrice

`func (o *Invoice) HasNetworkPrice() bool`

HasNetworkPrice returns a boolean if a field has been set.

### GetNetworkCost

`func (o *Invoice) GetNetworkCost() float32`

GetNetworkCost returns the NetworkCost field if non-nil, zero value otherwise.

### GetNetworkCostOk

`func (o *Invoice) GetNetworkCostOk() (*float32, bool)`

GetNetworkCostOk returns a tuple with the NetworkCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCost

`func (o *Invoice) SetNetworkCost(v float32)`

SetNetworkCost sets NetworkCost field to given value.

### HasNetworkCost

`func (o *Invoice) HasNetworkCost() bool`

HasNetworkCost returns a boolean if a field has been set.

### GetLicensePrice

`func (o *Invoice) GetLicensePrice() float32`

GetLicensePrice returns the LicensePrice field if non-nil, zero value otherwise.

### GetLicensePriceOk

`func (o *Invoice) GetLicensePriceOk() (*float32, bool)`

GetLicensePriceOk returns a tuple with the LicensePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePrice

`func (o *Invoice) SetLicensePrice(v float32)`

SetLicensePrice sets LicensePrice field to given value.

### HasLicensePrice

`func (o *Invoice) HasLicensePrice() bool`

HasLicensePrice returns a boolean if a field has been set.

### GetLicenseCost

`func (o *Invoice) GetLicenseCost() float32`

GetLicenseCost returns the LicenseCost field if non-nil, zero value otherwise.

### GetLicenseCostOk

`func (o *Invoice) GetLicenseCostOk() (*float32, bool)`

GetLicenseCostOk returns a tuple with the LicenseCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCost

`func (o *Invoice) SetLicenseCost(v float32)`

SetLicenseCost sets LicenseCost field to given value.

### HasLicenseCost

`func (o *Invoice) HasLicenseCost() bool`

HasLicenseCost returns a boolean if a field has been set.

### GetExtraPrice

`func (o *Invoice) GetExtraPrice() float32`

GetExtraPrice returns the ExtraPrice field if non-nil, zero value otherwise.

### GetExtraPriceOk

`func (o *Invoice) GetExtraPriceOk() (*float32, bool)`

GetExtraPriceOk returns a tuple with the ExtraPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraPrice

`func (o *Invoice) SetExtraPrice(v float32)`

SetExtraPrice sets ExtraPrice field to given value.

### HasExtraPrice

`func (o *Invoice) HasExtraPrice() bool`

HasExtraPrice returns a boolean if a field has been set.

### GetExtraCost

`func (o *Invoice) GetExtraCost() float32`

GetExtraCost returns the ExtraCost field if non-nil, zero value otherwise.

### GetExtraCostOk

`func (o *Invoice) GetExtraCostOk() (*float32, bool)`

GetExtraCostOk returns a tuple with the ExtraCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraCost

`func (o *Invoice) SetExtraCost(v float32)`

SetExtraCost sets ExtraCost field to given value.

### HasExtraCost

`func (o *Invoice) HasExtraCost() bool`

HasExtraCost returns a boolean if a field has been set.

### GetTotalPrice

`func (o *Invoice) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *Invoice) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *Invoice) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *Invoice) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetTotalCost

`func (o *Invoice) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *Invoice) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *Invoice) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *Invoice) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetRunningPrice

`func (o *Invoice) GetRunningPrice() float32`

GetRunningPrice returns the RunningPrice field if non-nil, zero value otherwise.

### GetRunningPriceOk

`func (o *Invoice) GetRunningPriceOk() (*float32, bool)`

GetRunningPriceOk returns a tuple with the RunningPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningPrice

`func (o *Invoice) SetRunningPrice(v float32)`

SetRunningPrice sets RunningPrice field to given value.

### HasRunningPrice

`func (o *Invoice) HasRunningPrice() bool`

HasRunningPrice returns a boolean if a field has been set.

### GetRunningCost

`func (o *Invoice) GetRunningCost() float32`

GetRunningCost returns the RunningCost field if non-nil, zero value otherwise.

### GetRunningCostOk

`func (o *Invoice) GetRunningCostOk() (*float32, bool)`

GetRunningCostOk returns a tuple with the RunningCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCost

`func (o *Invoice) SetRunningCost(v float32)`

SetRunningCost sets RunningCost field to given value.

### HasRunningCost

`func (o *Invoice) HasRunningCost() bool`

HasRunningCost returns a boolean if a field has been set.

### GetCurrency

`func (o *Invoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Invoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetConversionRate

`func (o *Invoice) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *Invoice) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *Invoice) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *Invoice) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetCostType

`func (o *Invoice) GetCostType() string`

GetCostType returns the CostType field if non-nil, zero value otherwise.

### GetCostTypeOk

`func (o *Invoice) GetCostTypeOk() (*string, bool)`

GetCostTypeOk returns a tuple with the CostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostType

`func (o *Invoice) SetCostType(v string)`

SetCostType sets CostType field to given value.

### HasCostType

`func (o *Invoice) HasCostType() bool`

HasCostType returns a boolean if a field has been set.

### GetOffTime

`func (o *Invoice) GetOffTime() int64`

GetOffTime returns the OffTime field if non-nil, zero value otherwise.

### GetOffTimeOk

`func (o *Invoice) GetOffTimeOk() (*int64, bool)`

GetOffTimeOk returns a tuple with the OffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffTime

`func (o *Invoice) SetOffTime(v int64)`

SetOffTime sets OffTime field to given value.

### HasOffTime

`func (o *Invoice) HasOffTime() bool`

HasOffTime returns a boolean if a field has been set.

### GetPowerState

`func (o *Invoice) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *Invoice) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *Invoice) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *Invoice) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetPowerDate

`func (o *Invoice) GetPowerDate() time.Time`

GetPowerDate returns the PowerDate field if non-nil, zero value otherwise.

### GetPowerDateOk

`func (o *Invoice) GetPowerDateOk() (*time.Time, bool)`

GetPowerDateOk returns a tuple with the PowerDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerDate

`func (o *Invoice) SetPowerDate(v time.Time)`

SetPowerDate sets PowerDate field to given value.

### HasPowerDate

`func (o *Invoice) HasPowerDate() bool`

HasPowerDate returns a boolean if a field has been set.

### GetRunningMultiplier

`func (o *Invoice) GetRunningMultiplier() float32`

GetRunningMultiplier returns the RunningMultiplier field if non-nil, zero value otherwise.

### GetRunningMultiplierOk

`func (o *Invoice) GetRunningMultiplierOk() (*float32, bool)`

GetRunningMultiplierOk returns a tuple with the RunningMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningMultiplier

`func (o *Invoice) SetRunningMultiplier(v float32)`

SetRunningMultiplier sets RunningMultiplier field to given value.

### HasRunningMultiplier

`func (o *Invoice) HasRunningMultiplier() bool`

HasRunningMultiplier returns a boolean if a field has been set.

### GetUsageType

`func (o *Invoice) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *Invoice) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *Invoice) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *Invoice) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetUsageCategory

`func (o *Invoice) GetUsageCategory() string`

GetUsageCategory returns the UsageCategory field if non-nil, zero value otherwise.

### GetUsageCategoryOk

`func (o *Invoice) GetUsageCategoryOk() (*string, bool)`

GetUsageCategoryOk returns a tuple with the UsageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCategory

`func (o *Invoice) SetUsageCategory(v string)`

SetUsageCategory sets UsageCategory field to given value.

### HasUsageCategory

`func (o *Invoice) HasUsageCategory() bool`

HasUsageCategory returns a boolean if a field has been set.

### GetLastCostDate

`func (o *Invoice) GetLastCostDate() time.Time`

GetLastCostDate returns the LastCostDate field if non-nil, zero value otherwise.

### GetLastCostDateOk

`func (o *Invoice) GetLastCostDateOk() (*time.Time, bool)`

GetLastCostDateOk returns a tuple with the LastCostDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCostDate

`func (o *Invoice) SetLastCostDate(v time.Time)`

SetLastCostDate sets LastCostDate field to given value.

### HasLastCostDate

`func (o *Invoice) HasLastCostDate() bool`

HasLastCostDate returns a boolean if a field has been set.

### GetLastActualDate

`func (o *Invoice) GetLastActualDate() time.Time`

GetLastActualDate returns the LastActualDate field if non-nil, zero value otherwise.

### GetLastActualDateOk

`func (o *Invoice) GetLastActualDateOk() (*time.Time, bool)`

GetLastActualDateOk returns a tuple with the LastActualDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActualDate

`func (o *Invoice) SetLastActualDate(v time.Time)`

SetLastActualDate sets LastActualDate field to given value.

### HasLastActualDate

`func (o *Invoice) HasLastActualDate() bool`

HasLastActualDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *Invoice) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Invoice) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Invoice) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Invoice) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Invoice) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Invoice) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Invoice) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Invoice) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLineItemCount

`func (o *Invoice) GetLineItemCount() int64`

GetLineItemCount returns the LineItemCount field if non-nil, zero value otherwise.

### GetLineItemCountOk

`func (o *Invoice) GetLineItemCountOk() (*int64, bool)`

GetLineItemCountOk returns a tuple with the LineItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemCount

`func (o *Invoice) SetLineItemCount(v int64)`

SetLineItemCount sets LineItemCount field to given value.

### HasLineItemCount

`func (o *Invoice) HasLineItemCount() bool`

HasLineItemCount returns a boolean if a field has been set.

### GetLineItems

`func (o *Invoice) GetLineItems() []ListInvoices200ResponseAllOfInvoicesInnerLineItemsInner`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *Invoice) GetLineItemsOk() (*[]ListInvoices200ResponseAllOfInvoicesInnerLineItemsInner, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *Invoice) SetLineItems(v []ListInvoices200ResponseAllOfInvoicesInnerLineItemsInner)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *Invoice) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



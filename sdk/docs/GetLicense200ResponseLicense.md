# GetLicense200ResponseLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID | [optional] 
**KeyId** | Pointer to **string** | Key ID (only the first 8 characters are required to identify license to uninstall) | [optional] 
**Hash** | Pointer to **string** | Hash of the license content which can be used as a fingerprint identifier | [optional] 
**ProductTier** | Pointer to **string** | Product Tier | [optional] 
**StartDate** | Pointer to **time.Time** | The start date of the applied license. | [optional] 
**EndDate** | Pointer to **time.Time** | The expiration date of the applied license. | [optional] 
**MaxInstances** | Pointer to **int64** | Workload Limit. 0 is used for unlimited. | [optional] 
**MaxMemory** | Pointer to **int64** | Memory Limit. 0 is used for unlimited. | [optional] 
**MaxStorage** | Pointer to **int64** | Storage Limit. 0 is used for unlimited. | [optional] 
**LimitType** | Pointer to **string** | The limit type determines which limits apply to the license, the new &#39;standard&#39; or legacy &#39;workload&#39;. | [optional] 
**MaxManagedServers** | Pointer to **int64** | Managed Servers Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxDiscoveredServers** | Pointer to **int64** | Discovered Servers Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxHosts** | Pointer to **int64** | Host Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxMvm** | Pointer to **int64** | HPE VM Host Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxMvmSockets** | Pointer to **int64** | HPE VM Host Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxIac** | Pointer to **int64** | IAC Deployments Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxXaas** | Pointer to **int64** | Xaas Instances Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxExecutions** | Pointer to **int64** | Execution Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxDistributedWorkers** | Pointer to **int64** | Distributed Workers Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxDiscoveredObjects** | Pointer to **int64** | Discovered Objects Limit. Not yet enforced. | [optional] 
**HardLimit** | Pointer to **bool** | Hard Limit | [optional] 
**FreeTrial** | Pointer to **bool** | Free Trial (Community License) | [optional] 
**MultiTenant** | Pointer to **bool** | Multi-Tenant Enabled | [optional] 
**Whitelabel** | Pointer to **bool** | White Label Enabled | [optional] 
**ReportStatus** | Pointer to **bool** | Stats Reporting. This is true when the appliance registers and sends usage stats to the hub. | [optional] 
**SupportLevel** | Pointer to **string** | Support Level | [optional] 
**AccountName** | Pointer to **string** | Account Name | [optional] 
**Config** | Pointer to **map[string]interface{}** | License Configuration Object | [optional] 
**AmazonProductCodes** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**GetLicense200ResponseLicenseFeatures**](GetLicense200ResponseLicenseFeatures.md) |  | [optional] 
**ZoneTypes** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**RecalculationDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetLicense200ResponseLicense

`func NewGetLicense200ResponseLicense() *GetLicense200ResponseLicense`

NewGetLicense200ResponseLicense instantiates a new GetLicense200ResponseLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLicense200ResponseLicenseWithDefaults

`func NewGetLicense200ResponseLicenseWithDefaults() *GetLicense200ResponseLicense`

NewGetLicense200ResponseLicenseWithDefaults instantiates a new GetLicense200ResponseLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLicense200ResponseLicense) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLicense200ResponseLicense) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLicense200ResponseLicense) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetLicense200ResponseLicense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeyId

`func (o *GetLicense200ResponseLicense) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *GetLicense200ResponseLicense) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *GetLicense200ResponseLicense) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *GetLicense200ResponseLicense) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetHash

`func (o *GetLicense200ResponseLicense) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetLicense200ResponseLicense) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetLicense200ResponseLicense) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GetLicense200ResponseLicense) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetProductTier

`func (o *GetLicense200ResponseLicense) GetProductTier() string`

GetProductTier returns the ProductTier field if non-nil, zero value otherwise.

### GetProductTierOk

`func (o *GetLicense200ResponseLicense) GetProductTierOk() (*string, bool)`

GetProductTierOk returns a tuple with the ProductTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTier

`func (o *GetLicense200ResponseLicense) SetProductTier(v string)`

SetProductTier sets ProductTier field to given value.

### HasProductTier

`func (o *GetLicense200ResponseLicense) HasProductTier() bool`

HasProductTier returns a boolean if a field has been set.

### GetStartDate

`func (o *GetLicense200ResponseLicense) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetLicense200ResponseLicense) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetLicense200ResponseLicense) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetLicense200ResponseLicense) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetLicense200ResponseLicense) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetLicense200ResponseLicense) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetLicense200ResponseLicense) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetLicense200ResponseLicense) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMaxInstances

`func (o *GetLicense200ResponseLicense) GetMaxInstances() int64`

GetMaxInstances returns the MaxInstances field if non-nil, zero value otherwise.

### GetMaxInstancesOk

`func (o *GetLicense200ResponseLicense) GetMaxInstancesOk() (*int64, bool)`

GetMaxInstancesOk returns a tuple with the MaxInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstances

`func (o *GetLicense200ResponseLicense) SetMaxInstances(v int64)`

SetMaxInstances sets MaxInstances field to given value.

### HasMaxInstances

`func (o *GetLicense200ResponseLicense) HasMaxInstances() bool`

HasMaxInstances returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetLicense200ResponseLicense) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetLicense200ResponseLicense) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetLicense200ResponseLicense) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetLicense200ResponseLicense) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GetLicense200ResponseLicense) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetLicense200ResponseLicense) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetLicense200ResponseLicense) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetLicense200ResponseLicense) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetLimitType

`func (o *GetLicense200ResponseLicense) GetLimitType() string`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *GetLicense200ResponseLicense) GetLimitTypeOk() (*string, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *GetLicense200ResponseLicense) SetLimitType(v string)`

SetLimitType sets LimitType field to given value.

### HasLimitType

`func (o *GetLicense200ResponseLicense) HasLimitType() bool`

HasLimitType returns a boolean if a field has been set.

### GetMaxManagedServers

`func (o *GetLicense200ResponseLicense) GetMaxManagedServers() int64`

GetMaxManagedServers returns the MaxManagedServers field if non-nil, zero value otherwise.

### GetMaxManagedServersOk

`func (o *GetLicense200ResponseLicense) GetMaxManagedServersOk() (*int64, bool)`

GetMaxManagedServersOk returns a tuple with the MaxManagedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxManagedServers

`func (o *GetLicense200ResponseLicense) SetMaxManagedServers(v int64)`

SetMaxManagedServers sets MaxManagedServers field to given value.

### HasMaxManagedServers

`func (o *GetLicense200ResponseLicense) HasMaxManagedServers() bool`

HasMaxManagedServers returns a boolean if a field has been set.

### GetMaxDiscoveredServers

`func (o *GetLicense200ResponseLicense) GetMaxDiscoveredServers() int64`

GetMaxDiscoveredServers returns the MaxDiscoveredServers field if non-nil, zero value otherwise.

### GetMaxDiscoveredServersOk

`func (o *GetLicense200ResponseLicense) GetMaxDiscoveredServersOk() (*int64, bool)`

GetMaxDiscoveredServersOk returns a tuple with the MaxDiscoveredServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDiscoveredServers

`func (o *GetLicense200ResponseLicense) SetMaxDiscoveredServers(v int64)`

SetMaxDiscoveredServers sets MaxDiscoveredServers field to given value.

### HasMaxDiscoveredServers

`func (o *GetLicense200ResponseLicense) HasMaxDiscoveredServers() bool`

HasMaxDiscoveredServers returns a boolean if a field has been set.

### GetMaxHosts

`func (o *GetLicense200ResponseLicense) GetMaxHosts() int64`

GetMaxHosts returns the MaxHosts field if non-nil, zero value otherwise.

### GetMaxHostsOk

`func (o *GetLicense200ResponseLicense) GetMaxHostsOk() (*int64, bool)`

GetMaxHostsOk returns a tuple with the MaxHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHosts

`func (o *GetLicense200ResponseLicense) SetMaxHosts(v int64)`

SetMaxHosts sets MaxHosts field to given value.

### HasMaxHosts

`func (o *GetLicense200ResponseLicense) HasMaxHosts() bool`

HasMaxHosts returns a boolean if a field has been set.

### GetMaxMvm

`func (o *GetLicense200ResponseLicense) GetMaxMvm() int64`

GetMaxMvm returns the MaxMvm field if non-nil, zero value otherwise.

### GetMaxMvmOk

`func (o *GetLicense200ResponseLicense) GetMaxMvmOk() (*int64, bool)`

GetMaxMvmOk returns a tuple with the MaxMvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMvm

`func (o *GetLicense200ResponseLicense) SetMaxMvm(v int64)`

SetMaxMvm sets MaxMvm field to given value.

### HasMaxMvm

`func (o *GetLicense200ResponseLicense) HasMaxMvm() bool`

HasMaxMvm returns a boolean if a field has been set.

### GetMaxMvmSockets

`func (o *GetLicense200ResponseLicense) GetMaxMvmSockets() int64`

GetMaxMvmSockets returns the MaxMvmSockets field if non-nil, zero value otherwise.

### GetMaxMvmSocketsOk

`func (o *GetLicense200ResponseLicense) GetMaxMvmSocketsOk() (*int64, bool)`

GetMaxMvmSocketsOk returns a tuple with the MaxMvmSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMvmSockets

`func (o *GetLicense200ResponseLicense) SetMaxMvmSockets(v int64)`

SetMaxMvmSockets sets MaxMvmSockets field to given value.

### HasMaxMvmSockets

`func (o *GetLicense200ResponseLicense) HasMaxMvmSockets() bool`

HasMaxMvmSockets returns a boolean if a field has been set.

### GetMaxIac

`func (o *GetLicense200ResponseLicense) GetMaxIac() int64`

GetMaxIac returns the MaxIac field if non-nil, zero value otherwise.

### GetMaxIacOk

`func (o *GetLicense200ResponseLicense) GetMaxIacOk() (*int64, bool)`

GetMaxIacOk returns a tuple with the MaxIac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIac

`func (o *GetLicense200ResponseLicense) SetMaxIac(v int64)`

SetMaxIac sets MaxIac field to given value.

### HasMaxIac

`func (o *GetLicense200ResponseLicense) HasMaxIac() bool`

HasMaxIac returns a boolean if a field has been set.

### GetMaxXaas

`func (o *GetLicense200ResponseLicense) GetMaxXaas() int64`

GetMaxXaas returns the MaxXaas field if non-nil, zero value otherwise.

### GetMaxXaasOk

`func (o *GetLicense200ResponseLicense) GetMaxXaasOk() (*int64, bool)`

GetMaxXaasOk returns a tuple with the MaxXaas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxXaas

`func (o *GetLicense200ResponseLicense) SetMaxXaas(v int64)`

SetMaxXaas sets MaxXaas field to given value.

### HasMaxXaas

`func (o *GetLicense200ResponseLicense) HasMaxXaas() bool`

HasMaxXaas returns a boolean if a field has been set.

### GetMaxExecutions

`func (o *GetLicense200ResponseLicense) GetMaxExecutions() int64`

GetMaxExecutions returns the MaxExecutions field if non-nil, zero value otherwise.

### GetMaxExecutionsOk

`func (o *GetLicense200ResponseLicense) GetMaxExecutionsOk() (*int64, bool)`

GetMaxExecutionsOk returns a tuple with the MaxExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExecutions

`func (o *GetLicense200ResponseLicense) SetMaxExecutions(v int64)`

SetMaxExecutions sets MaxExecutions field to given value.

### HasMaxExecutions

`func (o *GetLicense200ResponseLicense) HasMaxExecutions() bool`

HasMaxExecutions returns a boolean if a field has been set.

### GetMaxDistributedWorkers

`func (o *GetLicense200ResponseLicense) GetMaxDistributedWorkers() int64`

GetMaxDistributedWorkers returns the MaxDistributedWorkers field if non-nil, zero value otherwise.

### GetMaxDistributedWorkersOk

`func (o *GetLicense200ResponseLicense) GetMaxDistributedWorkersOk() (*int64, bool)`

GetMaxDistributedWorkersOk returns a tuple with the MaxDistributedWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistributedWorkers

`func (o *GetLicense200ResponseLicense) SetMaxDistributedWorkers(v int64)`

SetMaxDistributedWorkers sets MaxDistributedWorkers field to given value.

### HasMaxDistributedWorkers

`func (o *GetLicense200ResponseLicense) HasMaxDistributedWorkers() bool`

HasMaxDistributedWorkers returns a boolean if a field has been set.

### GetMaxDiscoveredObjects

`func (o *GetLicense200ResponseLicense) GetMaxDiscoveredObjects() int64`

GetMaxDiscoveredObjects returns the MaxDiscoveredObjects field if non-nil, zero value otherwise.

### GetMaxDiscoveredObjectsOk

`func (o *GetLicense200ResponseLicense) GetMaxDiscoveredObjectsOk() (*int64, bool)`

GetMaxDiscoveredObjectsOk returns a tuple with the MaxDiscoveredObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDiscoveredObjects

`func (o *GetLicense200ResponseLicense) SetMaxDiscoveredObjects(v int64)`

SetMaxDiscoveredObjects sets MaxDiscoveredObjects field to given value.

### HasMaxDiscoveredObjects

`func (o *GetLicense200ResponseLicense) HasMaxDiscoveredObjects() bool`

HasMaxDiscoveredObjects returns a boolean if a field has been set.

### GetHardLimit

`func (o *GetLicense200ResponseLicense) GetHardLimit() bool`

GetHardLimit returns the HardLimit field if non-nil, zero value otherwise.

### GetHardLimitOk

`func (o *GetLicense200ResponseLicense) GetHardLimitOk() (*bool, bool)`

GetHardLimitOk returns a tuple with the HardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardLimit

`func (o *GetLicense200ResponseLicense) SetHardLimit(v bool)`

SetHardLimit sets HardLimit field to given value.

### HasHardLimit

`func (o *GetLicense200ResponseLicense) HasHardLimit() bool`

HasHardLimit returns a boolean if a field has been set.

### GetFreeTrial

`func (o *GetLicense200ResponseLicense) GetFreeTrial() bool`

GetFreeTrial returns the FreeTrial field if non-nil, zero value otherwise.

### GetFreeTrialOk

`func (o *GetLicense200ResponseLicense) GetFreeTrialOk() (*bool, bool)`

GetFreeTrialOk returns a tuple with the FreeTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTrial

`func (o *GetLicense200ResponseLicense) SetFreeTrial(v bool)`

SetFreeTrial sets FreeTrial field to given value.

### HasFreeTrial

`func (o *GetLicense200ResponseLicense) HasFreeTrial() bool`

HasFreeTrial returns a boolean if a field has been set.

### GetMultiTenant

`func (o *GetLicense200ResponseLicense) GetMultiTenant() bool`

GetMultiTenant returns the MultiTenant field if non-nil, zero value otherwise.

### GetMultiTenantOk

`func (o *GetLicense200ResponseLicense) GetMultiTenantOk() (*bool, bool)`

GetMultiTenantOk returns a tuple with the MultiTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenant

`func (o *GetLicense200ResponseLicense) SetMultiTenant(v bool)`

SetMultiTenant sets MultiTenant field to given value.

### HasMultiTenant

`func (o *GetLicense200ResponseLicense) HasMultiTenant() bool`

HasMultiTenant returns a boolean if a field has been set.

### GetWhitelabel

`func (o *GetLicense200ResponseLicense) GetWhitelabel() bool`

GetWhitelabel returns the Whitelabel field if non-nil, zero value otherwise.

### GetWhitelabelOk

`func (o *GetLicense200ResponseLicense) GetWhitelabelOk() (*bool, bool)`

GetWhitelabelOk returns a tuple with the Whitelabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelabel

`func (o *GetLicense200ResponseLicense) SetWhitelabel(v bool)`

SetWhitelabel sets Whitelabel field to given value.

### HasWhitelabel

`func (o *GetLicense200ResponseLicense) HasWhitelabel() bool`

HasWhitelabel returns a boolean if a field has been set.

### GetReportStatus

`func (o *GetLicense200ResponseLicense) GetReportStatus() bool`

GetReportStatus returns the ReportStatus field if non-nil, zero value otherwise.

### GetReportStatusOk

`func (o *GetLicense200ResponseLicense) GetReportStatusOk() (*bool, bool)`

GetReportStatusOk returns a tuple with the ReportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportStatus

`func (o *GetLicense200ResponseLicense) SetReportStatus(v bool)`

SetReportStatus sets ReportStatus field to given value.

### HasReportStatus

`func (o *GetLicense200ResponseLicense) HasReportStatus() bool`

HasReportStatus returns a boolean if a field has been set.

### GetSupportLevel

`func (o *GetLicense200ResponseLicense) GetSupportLevel() string`

GetSupportLevel returns the SupportLevel field if non-nil, zero value otherwise.

### GetSupportLevelOk

`func (o *GetLicense200ResponseLicense) GetSupportLevelOk() (*string, bool)`

GetSupportLevelOk returns a tuple with the SupportLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLevel

`func (o *GetLicense200ResponseLicense) SetSupportLevel(v string)`

SetSupportLevel sets SupportLevel field to given value.

### HasSupportLevel

`func (o *GetLicense200ResponseLicense) HasSupportLevel() bool`

HasSupportLevel returns a boolean if a field has been set.

### GetAccountName

`func (o *GetLicense200ResponseLicense) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *GetLicense200ResponseLicense) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *GetLicense200ResponseLicense) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *GetLicense200ResponseLicense) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetConfig

`func (o *GetLicense200ResponseLicense) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetLicense200ResponseLicense) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetLicense200ResponseLicense) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetLicense200ResponseLicense) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetAmazonProductCodes

`func (o *GetLicense200ResponseLicense) GetAmazonProductCodes() string`

GetAmazonProductCodes returns the AmazonProductCodes field if non-nil, zero value otherwise.

### GetAmazonProductCodesOk

`func (o *GetLicense200ResponseLicense) GetAmazonProductCodesOk() (*string, bool)`

GetAmazonProductCodesOk returns a tuple with the AmazonProductCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonProductCodes

`func (o *GetLicense200ResponseLicense) SetAmazonProductCodes(v string)`

SetAmazonProductCodes sets AmazonProductCodes field to given value.

### HasAmazonProductCodes

`func (o *GetLicense200ResponseLicense) HasAmazonProductCodes() bool`

HasAmazonProductCodes returns a boolean if a field has been set.

### GetFeatures

`func (o *GetLicense200ResponseLicense) GetFeatures() GetLicense200ResponseLicenseFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GetLicense200ResponseLicense) GetFeaturesOk() (*GetLicense200ResponseLicenseFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GetLicense200ResponseLicense) SetFeatures(v GetLicense200ResponseLicenseFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *GetLicense200ResponseLicense) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetZoneTypes

`func (o *GetLicense200ResponseLicense) GetZoneTypes() string`

GetZoneTypes returns the ZoneTypes field if non-nil, zero value otherwise.

### GetZoneTypesOk

`func (o *GetLicense200ResponseLicense) GetZoneTypesOk() (*string, bool)`

GetZoneTypesOk returns a tuple with the ZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTypes

`func (o *GetLicense200ResponseLicense) SetZoneTypes(v string)`

SetZoneTypes sets ZoneTypes field to given value.

### HasZoneTypes

`func (o *GetLicense200ResponseLicense) HasZoneTypes() bool`

HasZoneTypes returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetLicense200ResponseLicense) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetLicense200ResponseLicense) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetLicense200ResponseLicense) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetLicense200ResponseLicense) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetLicense200ResponseLicense) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetLicense200ResponseLicense) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetLicense200ResponseLicense) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetLicense200ResponseLicense) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetRecalculationDate

`func (o *GetLicense200ResponseLicense) GetRecalculationDate() time.Time`

GetRecalculationDate returns the RecalculationDate field if non-nil, zero value otherwise.

### GetRecalculationDateOk

`func (o *GetLicense200ResponseLicense) GetRecalculationDateOk() (*time.Time, bool)`

GetRecalculationDateOk returns a tuple with the RecalculationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecalculationDate

`func (o *GetLicense200ResponseLicense) SetRecalculationDate(v time.Time)`

SetRecalculationDate sets RecalculationDate field to given value.

### HasRecalculationDate

`func (o *GetLicense200ResponseLicense) HasRecalculationDate() bool`

HasRecalculationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InstallLicense200ResponseInstalledLicensesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID | [optional] 
**KeyId** | Pointer to **string** | Key ID (only the first 8 characters are required to identify license to uninstall) | [optional] 
**Hash** | Pointer to **string** | Hash of the license content which can be used as a fingerprint identifier | [optional] 
**LicenseVersion** | Pointer to **int64** | License Version which determines the required appliance version to install this license. | [optional] 
**ProductTier** | Pointer to **string** | Product Tier | [optional] 
**StartDate** | Pointer to **time.Time** | The start date of the applied license. | [optional] 
**EndDate** | Pointer to **time.Time** | The expiration date of the applied license. | [optional] 
**MaxInstances** | Pointer to **int64** | Workload Limit. 0 is used for unlimited. | [optional] 
**MaxMemory** | Pointer to **int64** | Memory Limit. 0 is used for unlimited. | [optional] 
**MaxStorage** | Pointer to **int64** | Storage Limit. 0 is used for unlimited. | [optional] 
**LimitType** | Pointer to **string** | The limit type determines which limits apply to the license, the new &#39;standard&#39; or legacy &#39;workload&#39;. | [optional] 
**MaxManagedServers** | Pointer to **NullableInt64** | Managed Servers Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxDiscoveredServers** | Pointer to **NullableInt64** | Discovered Servers Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxHosts** | Pointer to **NullableInt64** | Host Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxMvm** | Pointer to **NullableInt64** | HPE VM Host Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxMvmSockets** | Pointer to **NullableInt64** | HPE VM Host Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxSockets** | Pointer to **NullableInt64** | Global Socket Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxIac** | Pointer to **NullableInt64** | IAC Deployments Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxXaas** | Pointer to **NullableInt64** | Xaas Instances Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxExecutions** | Pointer to **NullableInt64** | Execution Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxDistributedWorkers** | Pointer to **NullableInt64** | Distributed Workers Limit. 0 is enforced and null is used for unlimited. | [optional] 
**MaxDiscoveredObjects** | Pointer to **NullableInt64** | Discovered Objects Limit. Not yet enforced. | [optional] 
**HardLimit** | Pointer to **bool** | Hard Limit | [optional] 
**FreeTrial** | Pointer to **bool** | Free Trial (Community License) | [optional] 
**MultiTenant** | Pointer to **bool** | Multi-Tenant Enabled | [optional] 
**Whitelabel** | Pointer to **bool** | White Label Enabled | [optional] 
**ReportStatus** | Pointer to **bool** | Stats Reporting. This is true when the appliance registers and sends usage stats to the hub. | [optional] 
**SupportLevel** | Pointer to **string** | Support Level | [optional] 
**AccountName** | Pointer to **string** | Account Name | [optional] 
**Config** | Pointer to **map[string]interface{}** | License Configuration Object | [optional] 
**AmazonProductCodes** | Pointer to **NullableString** |  | [optional] 
**Features** | Pointer to [**InstallLicense200ResponseInstalledLicensesInnerFeatures**](InstallLicense200ResponseInstalledLicensesInnerFeatures.md) |  | [optional] 
**ZoneTypes** | Pointer to **NullableString** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**RecalculationDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewInstallLicense200ResponseInstalledLicensesInner

`func NewInstallLicense200ResponseInstalledLicensesInner() *InstallLicense200ResponseInstalledLicensesInner`

NewInstallLicense200ResponseInstalledLicensesInner instantiates a new InstallLicense200ResponseInstalledLicensesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallLicense200ResponseInstalledLicensesInnerWithDefaults

`func NewInstallLicense200ResponseInstalledLicensesInnerWithDefaults() *InstallLicense200ResponseInstalledLicensesInner`

NewInstallLicense200ResponseInstalledLicensesInnerWithDefaults instantiates a new InstallLicense200ResponseInstalledLicensesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeyId

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetHash

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetLicenseVersion

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetLicenseVersion() int64`

GetLicenseVersion returns the LicenseVersion field if non-nil, zero value otherwise.

### GetLicenseVersionOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetLicenseVersionOk() (*int64, bool)`

GetLicenseVersionOk returns a tuple with the LicenseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseVersion

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetLicenseVersion(v int64)`

SetLicenseVersion sets LicenseVersion field to given value.

### HasLicenseVersion

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasLicenseVersion() bool`

HasLicenseVersion returns a boolean if a field has been set.

### GetProductTier

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetProductTier() string`

GetProductTier returns the ProductTier field if non-nil, zero value otherwise.

### GetProductTierOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetProductTierOk() (*string, bool)`

GetProductTierOk returns a tuple with the ProductTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTier

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetProductTier(v string)`

SetProductTier sets ProductTier field to given value.

### HasProductTier

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasProductTier() bool`

HasProductTier returns a boolean if a field has been set.

### GetStartDate

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetMaxInstances

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxInstances() int64`

GetMaxInstances returns the MaxInstances field if non-nil, zero value otherwise.

### GetMaxInstancesOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxInstancesOk() (*int64, bool)`

GetMaxInstancesOk returns a tuple with the MaxInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstances

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxInstances(v int64)`

SetMaxInstances sets MaxInstances field to given value.

### HasMaxInstances

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxInstances() bool`

HasMaxInstances returns a boolean if a field has been set.

### GetMaxMemory

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetLimitType

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetLimitType() string`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetLimitTypeOk() (*string, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetLimitType(v string)`

SetLimitType sets LimitType field to given value.

### HasLimitType

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasLimitType() bool`

HasLimitType returns a boolean if a field has been set.

### GetMaxManagedServers

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxManagedServers() int64`

GetMaxManagedServers returns the MaxManagedServers field if non-nil, zero value otherwise.

### GetMaxManagedServersOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxManagedServersOk() (*int64, bool)`

GetMaxManagedServersOk returns a tuple with the MaxManagedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxManagedServers

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxManagedServers(v int64)`

SetMaxManagedServers sets MaxManagedServers field to given value.

### HasMaxManagedServers

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxManagedServers() bool`

HasMaxManagedServers returns a boolean if a field has been set.

### SetMaxManagedServersNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxManagedServersNil(b bool)`

 SetMaxManagedServersNil sets the value for MaxManagedServers to be an explicit nil

### UnsetMaxManagedServers
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetMaxManagedServers()`

UnsetMaxManagedServers ensures that no value is present for MaxManagedServers, not even an explicit nil
### GetMaxDiscoveredServers

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxDiscoveredServers() int64`

GetMaxDiscoveredServers returns the MaxDiscoveredServers field if non-nil, zero value otherwise.

### GetMaxDiscoveredServersOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxDiscoveredServersOk() (*int64, bool)`

GetMaxDiscoveredServersOk returns a tuple with the MaxDiscoveredServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDiscoveredServers

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxDiscoveredServers(v int64)`

SetMaxDiscoveredServers sets MaxDiscoveredServers field to given value.

### HasMaxDiscoveredServers

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxDiscoveredServers() bool`

HasMaxDiscoveredServers returns a boolean if a field has been set.

### SetMaxDiscoveredServersNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxDiscoveredServersNil(b bool)`

 SetMaxDiscoveredServersNil sets the value for MaxDiscoveredServers to be an explicit nil

### UnsetMaxDiscoveredServers
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetMaxDiscoveredServers()`

UnsetMaxDiscoveredServers ensures that no value is present for MaxDiscoveredServers, not even an explicit nil
### GetMaxHosts

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxHosts() int64`

GetMaxHosts returns the MaxHosts field if non-nil, zero value otherwise.

### GetMaxHostsOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxHostsOk() (*int64, bool)`

GetMaxHostsOk returns a tuple with the MaxHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHosts

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxHosts(v int64)`

SetMaxHosts sets MaxHosts field to given value.

### HasMaxHosts

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxHosts() bool`

HasMaxHosts returns a boolean if a field has been set.

### SetMaxHostsNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxHostsNil(b bool)`

 SetMaxHostsNil sets the value for MaxHosts to be an explicit nil

### UnsetMaxHosts
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetMaxHosts()`

UnsetMaxHosts ensures that no value is present for MaxHosts, not even an explicit nil
### GetMaxMvm

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxMvm() int64`

GetMaxMvm returns the MaxMvm field if non-nil, zero value otherwise.

### GetMaxMvmOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxMvmOk() (*int64, bool)`

GetMaxMvmOk returns a tuple with the MaxMvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMvm

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxMvm(v int64)`

SetMaxMvm sets MaxMvm field to given value.

### HasMaxMvm

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxMvm() bool`

HasMaxMvm returns a boolean if a field has been set.

### SetMaxMvmNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxMvmNil(b bool)`

 SetMaxMvmNil sets the value for MaxMvm to be an explicit nil

### UnsetMaxMvm
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetMaxMvm()`

UnsetMaxMvm ensures that no value is present for MaxMvm, not even an explicit nil
### GetMaxMvmSockets

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxMvmSockets() int64`

GetMaxMvmSockets returns the MaxMvmSockets field if non-nil, zero value otherwise.

### GetMaxMvmSocketsOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxMvmSocketsOk() (*int64, bool)`

GetMaxMvmSocketsOk returns a tuple with the MaxMvmSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMvmSockets

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxMvmSockets(v int64)`

SetMaxMvmSockets sets MaxMvmSockets field to given value.

### HasMaxMvmSockets

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxMvmSockets() bool`

HasMaxMvmSockets returns a boolean if a field has been set.

### SetMaxMvmSocketsNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxMvmSocketsNil(b bool)`

 SetMaxMvmSocketsNil sets the value for MaxMvmSockets to be an explicit nil

### UnsetMaxMvmSockets
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetMaxMvmSockets()`

UnsetMaxMvmSockets ensures that no value is present for MaxMvmSockets, not even an explicit nil
### GetMaxSockets

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxSockets() int64`

GetMaxSockets returns the MaxSockets field if non-nil, zero value otherwise.

### GetMaxSocketsOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxSocketsOk() (*int64, bool)`

GetMaxSocketsOk returns a tuple with the MaxSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSockets

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxSockets(v int64)`

SetMaxSockets sets MaxSockets field to given value.

### HasMaxSockets

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxSockets() bool`

HasMaxSockets returns a boolean if a field has been set.

### SetMaxSocketsNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxSocketsNil(b bool)`

 SetMaxSocketsNil sets the value for MaxSockets to be an explicit nil

### UnsetMaxSockets
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetMaxSockets()`

UnsetMaxSockets ensures that no value is present for MaxSockets, not even an explicit nil
### GetMaxIac

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxIac() int64`

GetMaxIac returns the MaxIac field if non-nil, zero value otherwise.

### GetMaxIacOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxIacOk() (*int64, bool)`

GetMaxIacOk returns a tuple with the MaxIac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIac

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxIac(v int64)`

SetMaxIac sets MaxIac field to given value.

### HasMaxIac

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxIac() bool`

HasMaxIac returns a boolean if a field has been set.

### SetMaxIacNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxIacNil(b bool)`

 SetMaxIacNil sets the value for MaxIac to be an explicit nil

### UnsetMaxIac
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetMaxIac()`

UnsetMaxIac ensures that no value is present for MaxIac, not even an explicit nil
### GetMaxXaas

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxXaas() int64`

GetMaxXaas returns the MaxXaas field if non-nil, zero value otherwise.

### GetMaxXaasOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxXaasOk() (*int64, bool)`

GetMaxXaasOk returns a tuple with the MaxXaas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxXaas

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxXaas(v int64)`

SetMaxXaas sets MaxXaas field to given value.

### HasMaxXaas

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxXaas() bool`

HasMaxXaas returns a boolean if a field has been set.

### SetMaxXaasNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxXaasNil(b bool)`

 SetMaxXaasNil sets the value for MaxXaas to be an explicit nil

### UnsetMaxXaas
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetMaxXaas()`

UnsetMaxXaas ensures that no value is present for MaxXaas, not even an explicit nil
### GetMaxExecutions

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxExecutions() int64`

GetMaxExecutions returns the MaxExecutions field if non-nil, zero value otherwise.

### GetMaxExecutionsOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxExecutionsOk() (*int64, bool)`

GetMaxExecutionsOk returns a tuple with the MaxExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExecutions

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxExecutions(v int64)`

SetMaxExecutions sets MaxExecutions field to given value.

### HasMaxExecutions

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxExecutions() bool`

HasMaxExecutions returns a boolean if a field has been set.

### SetMaxExecutionsNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxExecutionsNil(b bool)`

 SetMaxExecutionsNil sets the value for MaxExecutions to be an explicit nil

### UnsetMaxExecutions
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetMaxExecutions()`

UnsetMaxExecutions ensures that no value is present for MaxExecutions, not even an explicit nil
### GetMaxDistributedWorkers

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxDistributedWorkers() int64`

GetMaxDistributedWorkers returns the MaxDistributedWorkers field if non-nil, zero value otherwise.

### GetMaxDistributedWorkersOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxDistributedWorkersOk() (*int64, bool)`

GetMaxDistributedWorkersOk returns a tuple with the MaxDistributedWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistributedWorkers

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxDistributedWorkers(v int64)`

SetMaxDistributedWorkers sets MaxDistributedWorkers field to given value.

### HasMaxDistributedWorkers

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxDistributedWorkers() bool`

HasMaxDistributedWorkers returns a boolean if a field has been set.

### SetMaxDistributedWorkersNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxDistributedWorkersNil(b bool)`

 SetMaxDistributedWorkersNil sets the value for MaxDistributedWorkers to be an explicit nil

### UnsetMaxDistributedWorkers
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetMaxDistributedWorkers()`

UnsetMaxDistributedWorkers ensures that no value is present for MaxDistributedWorkers, not even an explicit nil
### GetMaxDiscoveredObjects

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxDiscoveredObjects() int64`

GetMaxDiscoveredObjects returns the MaxDiscoveredObjects field if non-nil, zero value otherwise.

### GetMaxDiscoveredObjectsOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMaxDiscoveredObjectsOk() (*int64, bool)`

GetMaxDiscoveredObjectsOk returns a tuple with the MaxDiscoveredObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDiscoveredObjects

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxDiscoveredObjects(v int64)`

SetMaxDiscoveredObjects sets MaxDiscoveredObjects field to given value.

### HasMaxDiscoveredObjects

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMaxDiscoveredObjects() bool`

HasMaxDiscoveredObjects returns a boolean if a field has been set.

### SetMaxDiscoveredObjectsNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMaxDiscoveredObjectsNil(b bool)`

 SetMaxDiscoveredObjectsNil sets the value for MaxDiscoveredObjects to be an explicit nil

### UnsetMaxDiscoveredObjects
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetMaxDiscoveredObjects()`

UnsetMaxDiscoveredObjects ensures that no value is present for MaxDiscoveredObjects, not even an explicit nil
### GetHardLimit

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetHardLimit() bool`

GetHardLimit returns the HardLimit field if non-nil, zero value otherwise.

### GetHardLimitOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetHardLimitOk() (*bool, bool)`

GetHardLimitOk returns a tuple with the HardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardLimit

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetHardLimit(v bool)`

SetHardLimit sets HardLimit field to given value.

### HasHardLimit

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasHardLimit() bool`

HasHardLimit returns a boolean if a field has been set.

### GetFreeTrial

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetFreeTrial() bool`

GetFreeTrial returns the FreeTrial field if non-nil, zero value otherwise.

### GetFreeTrialOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetFreeTrialOk() (*bool, bool)`

GetFreeTrialOk returns a tuple with the FreeTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTrial

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetFreeTrial(v bool)`

SetFreeTrial sets FreeTrial field to given value.

### HasFreeTrial

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasFreeTrial() bool`

HasFreeTrial returns a boolean if a field has been set.

### GetMultiTenant

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMultiTenant() bool`

GetMultiTenant returns the MultiTenant field if non-nil, zero value otherwise.

### GetMultiTenantOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetMultiTenantOk() (*bool, bool)`

GetMultiTenantOk returns a tuple with the MultiTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenant

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetMultiTenant(v bool)`

SetMultiTenant sets MultiTenant field to given value.

### HasMultiTenant

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasMultiTenant() bool`

HasMultiTenant returns a boolean if a field has been set.

### GetWhitelabel

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetWhitelabel() bool`

GetWhitelabel returns the Whitelabel field if non-nil, zero value otherwise.

### GetWhitelabelOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetWhitelabelOk() (*bool, bool)`

GetWhitelabelOk returns a tuple with the Whitelabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelabel

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetWhitelabel(v bool)`

SetWhitelabel sets Whitelabel field to given value.

### HasWhitelabel

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasWhitelabel() bool`

HasWhitelabel returns a boolean if a field has been set.

### GetReportStatus

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetReportStatus() bool`

GetReportStatus returns the ReportStatus field if non-nil, zero value otherwise.

### GetReportStatusOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetReportStatusOk() (*bool, bool)`

GetReportStatusOk returns a tuple with the ReportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportStatus

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetReportStatus(v bool)`

SetReportStatus sets ReportStatus field to given value.

### HasReportStatus

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasReportStatus() bool`

HasReportStatus returns a boolean if a field has been set.

### GetSupportLevel

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetSupportLevel() string`

GetSupportLevel returns the SupportLevel field if non-nil, zero value otherwise.

### GetSupportLevelOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetSupportLevelOk() (*string, bool)`

GetSupportLevelOk returns a tuple with the SupportLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLevel

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetSupportLevel(v string)`

SetSupportLevel sets SupportLevel field to given value.

### HasSupportLevel

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasSupportLevel() bool`

HasSupportLevel returns a boolean if a field has been set.

### GetAccountName

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetConfig

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetAmazonProductCodes

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetAmazonProductCodes() string`

GetAmazonProductCodes returns the AmazonProductCodes field if non-nil, zero value otherwise.

### GetAmazonProductCodesOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetAmazonProductCodesOk() (*string, bool)`

GetAmazonProductCodesOk returns a tuple with the AmazonProductCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonProductCodes

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetAmazonProductCodes(v string)`

SetAmazonProductCodes sets AmazonProductCodes field to given value.

### HasAmazonProductCodes

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasAmazonProductCodes() bool`

HasAmazonProductCodes returns a boolean if a field has been set.

### SetAmazonProductCodesNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetAmazonProductCodesNil(b bool)`

 SetAmazonProductCodesNil sets the value for AmazonProductCodes to be an explicit nil

### UnsetAmazonProductCodes
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetAmazonProductCodes()`

UnsetAmazonProductCodes ensures that no value is present for AmazonProductCodes, not even an explicit nil
### GetFeatures

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetFeatures() InstallLicense200ResponseInstalledLicensesInnerFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetFeaturesOk() (*InstallLicense200ResponseInstalledLicensesInnerFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetFeatures(v InstallLicense200ResponseInstalledLicensesInnerFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetZoneTypes

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetZoneTypes() string`

GetZoneTypes returns the ZoneTypes field if non-nil, zero value otherwise.

### GetZoneTypesOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetZoneTypesOk() (*string, bool)`

GetZoneTypesOk returns a tuple with the ZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTypes

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetZoneTypes(v string)`

SetZoneTypes sets ZoneTypes field to given value.

### HasZoneTypes

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasZoneTypes() bool`

HasZoneTypes returns a boolean if a field has been set.

### SetZoneTypesNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetZoneTypesNil(b bool)`

 SetZoneTypesNil sets the value for ZoneTypes to be an explicit nil

### UnsetZoneTypes
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetZoneTypes()`

UnsetZoneTypes ensures that no value is present for ZoneTypes, not even an explicit nil
### GetLastUpdated

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDateCreated

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetRecalculationDate

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetRecalculationDate() time.Time`

GetRecalculationDate returns the RecalculationDate field if non-nil, zero value otherwise.

### GetRecalculationDateOk

`func (o *InstallLicense200ResponseInstalledLicensesInner) GetRecalculationDateOk() (*time.Time, bool)`

GetRecalculationDateOk returns a tuple with the RecalculationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecalculationDate

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetRecalculationDate(v time.Time)`

SetRecalculationDate sets RecalculationDate field to given value.

### HasRecalculationDate

`func (o *InstallLicense200ResponseInstalledLicensesInner) HasRecalculationDate() bool`

HasRecalculationDate returns a boolean if a field has been set.

### SetRecalculationDateNil

`func (o *InstallLicense200ResponseInstalledLicensesInner) SetRecalculationDateNil(b bool)`

 SetRecalculationDateNil sets the value for RecalculationDate to be an explicit nil

### UnsetRecalculationDate
`func (o *InstallLicense200ResponseInstalledLicensesInner) UnsetRecalculationDate()`

UnsetRecalculationDate ensures that no value is present for RecalculationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetContainer200ResponseContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int32** |  | [optional] 
**Instance** | Pointer to [**GetContainer200ResponseContainerInstance**](GetContainer200ResponseContainerInstance.md) |  | [optional] 
**ContainerType** | Pointer to [**GetContainer200ResponseContainerContainerType**](GetContainer200ResponseContainerContainerType.md) |  | [optional] 
**ContainerTypeSet** | Pointer to [**GetContainer200ResponseContainerContainerTypeSet**](GetContainer200ResponseContainerContainerTypeSet.md) |  | [optional] 
**Server** | Pointer to [**GetContainer200ResponseContainerInstance**](GetContainer200ResponseContainerInstance.md) |  | [optional] 
**Cloud** | Pointer to [**GetContainer200ResponseContainerInstance**](GetContainer200ResponseContainerInstance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**InternalHostname** | Pointer to **string** |  | [optional] 
**ExternalHostname** | Pointer to **string** |  | [optional] 
**ExternalDomain** | Pointer to **string** |  | [optional] 
**ExternalFqdn** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to [**[]GetContainer200ResponseContainerPortsInner**](GetContainer200ResponseContainerPortsInner.md) |  | [optional] 
**Plan** | Pointer to [**GetContainer200ResponseContainerPlan**](GetContainer200ResponseContainerPlan.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**StatsEnabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UserStatus** | Pointer to **string** |  | [optional] 
**EnvironmentPrefix** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**GetContainer200ResponseContainerStats**](GetContainer200ResponseContainerStats.md) |  | [optional] 
**RuntimeInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerVersion** | Pointer to **string** |  | [optional] 
**RepositoryImage** | Pointer to **string** |  | [optional] 
**PlanCategory** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**VolumeCreated** | Pointer to **bool** |  | [optional] 
**ContainerCreated** | Pointer to **bool** |  | [optional] 
**MaxStorage** | Pointer to **int32** |  | [optional] 
**MaxMemory** | Pointer to **int32** |  | [optional] 
**MaxCores** | Pointer to **int32** |  | [optional] 
**MaxCpu** | Pointer to **int32** |  | [optional] 
**AvailableActions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ConfigGroup** | Pointer to **string** |  | [optional] 
**ConfigId** | Pointer to **string** |  | [optional] 
**ConfigRole** | Pointer to **string** |  | [optional] 
**HourlyCost** | Pointer to **float64** |  | [optional] 
**HourlyPrice** | Pointer to **float64** |  | [optional] 

## Methods

### NewGetContainer200ResponseContainer

`func NewGetContainer200ResponseContainer() *GetContainer200ResponseContainer`

NewGetContainer200ResponseContainer instantiates a new GetContainer200ResponseContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainer200ResponseContainerWithDefaults

`func NewGetContainer200ResponseContainerWithDefaults() *GetContainer200ResponseContainer`

NewGetContainer200ResponseContainerWithDefaults instantiates a new GetContainer200ResponseContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetContainer200ResponseContainer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetContainer200ResponseContainer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetContainer200ResponseContainer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetContainer200ResponseContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GetContainer200ResponseContainer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetContainer200ResponseContainer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetContainer200ResponseContainer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetContainer200ResponseContainer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccountId

`func (o *GetContainer200ResponseContainer) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetContainer200ResponseContainer) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetContainer200ResponseContainer) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetContainer200ResponseContainer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInstance

`func (o *GetContainer200ResponseContainer) GetInstance() GetContainer200ResponseContainerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetContainer200ResponseContainer) GetInstanceOk() (*GetContainer200ResponseContainerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetContainer200ResponseContainer) SetInstance(v GetContainer200ResponseContainerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetContainer200ResponseContainer) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerType

`func (o *GetContainer200ResponseContainer) GetContainerType() GetContainer200ResponseContainerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *GetContainer200ResponseContainer) GetContainerTypeOk() (*GetContainer200ResponseContainerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *GetContainer200ResponseContainer) SetContainerType(v GetContainer200ResponseContainerContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *GetContainer200ResponseContainer) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetContainerTypeSet

`func (o *GetContainer200ResponseContainer) GetContainerTypeSet() GetContainer200ResponseContainerContainerTypeSet`

GetContainerTypeSet returns the ContainerTypeSet field if non-nil, zero value otherwise.

### GetContainerTypeSetOk

`func (o *GetContainer200ResponseContainer) GetContainerTypeSetOk() (*GetContainer200ResponseContainerContainerTypeSet, bool)`

GetContainerTypeSetOk returns a tuple with the ContainerTypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypeSet

`func (o *GetContainer200ResponseContainer) SetContainerTypeSet(v GetContainer200ResponseContainerContainerTypeSet)`

SetContainerTypeSet sets ContainerTypeSet field to given value.

### HasContainerTypeSet

`func (o *GetContainer200ResponseContainer) HasContainerTypeSet() bool`

HasContainerTypeSet returns a boolean if a field has been set.

### GetServer

`func (o *GetContainer200ResponseContainer) GetServer() GetContainer200ResponseContainerInstance`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetContainer200ResponseContainer) GetServerOk() (*GetContainer200ResponseContainerInstance, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetContainer200ResponseContainer) SetServer(v GetContainer200ResponseContainerInstance)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetContainer200ResponseContainer) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetCloud

`func (o *GetContainer200ResponseContainer) GetCloud() GetContainer200ResponseContainerInstance`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *GetContainer200ResponseContainer) GetCloudOk() (*GetContainer200ResponseContainerInstance, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *GetContainer200ResponseContainer) SetCloud(v GetContainer200ResponseContainerInstance)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *GetContainer200ResponseContainer) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetName

`func (o *GetContainer200ResponseContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetContainer200ResponseContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetContainer200ResponseContainer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetContainer200ResponseContainer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *GetContainer200ResponseContainer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetContainer200ResponseContainer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetContainer200ResponseContainer) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetContainer200ResponseContainer) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *GetContainer200ResponseContainer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *GetContainer200ResponseContainer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *GetContainer200ResponseContainer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *GetContainer200ResponseContainer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetInternalHostname

`func (o *GetContainer200ResponseContainer) GetInternalHostname() string`

GetInternalHostname returns the InternalHostname field if non-nil, zero value otherwise.

### GetInternalHostnameOk

`func (o *GetContainer200ResponseContainer) GetInternalHostnameOk() (*string, bool)`

GetInternalHostnameOk returns a tuple with the InternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostname

`func (o *GetContainer200ResponseContainer) SetInternalHostname(v string)`

SetInternalHostname sets InternalHostname field to given value.

### HasInternalHostname

`func (o *GetContainer200ResponseContainer) HasInternalHostname() bool`

HasInternalHostname returns a boolean if a field has been set.

### GetExternalHostname

`func (o *GetContainer200ResponseContainer) GetExternalHostname() string`

GetExternalHostname returns the ExternalHostname field if non-nil, zero value otherwise.

### GetExternalHostnameOk

`func (o *GetContainer200ResponseContainer) GetExternalHostnameOk() (*string, bool)`

GetExternalHostnameOk returns a tuple with the ExternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHostname

`func (o *GetContainer200ResponseContainer) SetExternalHostname(v string)`

SetExternalHostname sets ExternalHostname field to given value.

### HasExternalHostname

`func (o *GetContainer200ResponseContainer) HasExternalHostname() bool`

HasExternalHostname returns a boolean if a field has been set.

### GetExternalDomain

`func (o *GetContainer200ResponseContainer) GetExternalDomain() string`

GetExternalDomain returns the ExternalDomain field if non-nil, zero value otherwise.

### GetExternalDomainOk

`func (o *GetContainer200ResponseContainer) GetExternalDomainOk() (*string, bool)`

GetExternalDomainOk returns a tuple with the ExternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDomain

`func (o *GetContainer200ResponseContainer) SetExternalDomain(v string)`

SetExternalDomain sets ExternalDomain field to given value.

### HasExternalDomain

`func (o *GetContainer200ResponseContainer) HasExternalDomain() bool`

HasExternalDomain returns a boolean if a field has been set.

### GetExternalFqdn

`func (o *GetContainer200ResponseContainer) GetExternalFqdn() string`

GetExternalFqdn returns the ExternalFqdn field if non-nil, zero value otherwise.

### GetExternalFqdnOk

`func (o *GetContainer200ResponseContainer) GetExternalFqdnOk() (*string, bool)`

GetExternalFqdnOk returns a tuple with the ExternalFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqdn

`func (o *GetContainer200ResponseContainer) SetExternalFqdn(v string)`

SetExternalFqdn sets ExternalFqdn field to given value.

### HasExternalFqdn

`func (o *GetContainer200ResponseContainer) HasExternalFqdn() bool`

HasExternalFqdn returns a boolean if a field has been set.

### GetPorts

`func (o *GetContainer200ResponseContainer) GetPorts() []GetContainer200ResponseContainerPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GetContainer200ResponseContainer) GetPortsOk() (*[]GetContainer200ResponseContainerPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GetContainer200ResponseContainer) SetPorts(v []GetContainer200ResponseContainerPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *GetContainer200ResponseContainer) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPlan

`func (o *GetContainer200ResponseContainer) GetPlan() GetContainer200ResponseContainerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetContainer200ResponseContainer) GetPlanOk() (*GetContainer200ResponseContainerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetContainer200ResponseContainer) SetPlan(v GetContainer200ResponseContainerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetContainer200ResponseContainer) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetContainer200ResponseContainer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetContainer200ResponseContainer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetContainer200ResponseContainer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetContainer200ResponseContainer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetContainer200ResponseContainer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetContainer200ResponseContainer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetContainer200ResponseContainer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetContainer200ResponseContainer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatsEnabled

`func (o *GetContainer200ResponseContainer) GetStatsEnabled() bool`

GetStatsEnabled returns the StatsEnabled field if non-nil, zero value otherwise.

### GetStatsEnabledOk

`func (o *GetContainer200ResponseContainer) GetStatsEnabledOk() (*bool, bool)`

GetStatsEnabledOk returns a tuple with the StatsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsEnabled

`func (o *GetContainer200ResponseContainer) SetStatsEnabled(v bool)`

SetStatsEnabled sets StatsEnabled field to given value.

### HasStatsEnabled

`func (o *GetContainer200ResponseContainer) HasStatsEnabled() bool`

HasStatsEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *GetContainer200ResponseContainer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetContainer200ResponseContainer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetContainer200ResponseContainer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetContainer200ResponseContainer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserStatus

`func (o *GetContainer200ResponseContainer) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *GetContainer200ResponseContainer) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *GetContainer200ResponseContainer) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *GetContainer200ResponseContainer) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *GetContainer200ResponseContainer) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *GetContainer200ResponseContainer) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *GetContainer200ResponseContainer) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *GetContainer200ResponseContainer) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetStats

`func (o *GetContainer200ResponseContainer) GetStats() GetContainer200ResponseContainerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetContainer200ResponseContainer) GetStatsOk() (*GetContainer200ResponseContainerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetContainer200ResponseContainer) SetStats(v GetContainer200ResponseContainerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetContainer200ResponseContainer) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetRuntimeInfo

`func (o *GetContainer200ResponseContainer) GetRuntimeInfo() map[string]interface{}`

GetRuntimeInfo returns the RuntimeInfo field if non-nil, zero value otherwise.

### GetRuntimeInfoOk

`func (o *GetContainer200ResponseContainer) GetRuntimeInfoOk() (*map[string]interface{}, bool)`

GetRuntimeInfoOk returns a tuple with the RuntimeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeInfo

`func (o *GetContainer200ResponseContainer) SetRuntimeInfo(v map[string]interface{})`

SetRuntimeInfo sets RuntimeInfo field to given value.

### HasRuntimeInfo

`func (o *GetContainer200ResponseContainer) HasRuntimeInfo() bool`

HasRuntimeInfo returns a boolean if a field has been set.

### GetContainerVersion

`func (o *GetContainer200ResponseContainer) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *GetContainer200ResponseContainer) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *GetContainer200ResponseContainer) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *GetContainer200ResponseContainer) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetRepositoryImage

`func (o *GetContainer200ResponseContainer) GetRepositoryImage() string`

GetRepositoryImage returns the RepositoryImage field if non-nil, zero value otherwise.

### GetRepositoryImageOk

`func (o *GetContainer200ResponseContainer) GetRepositoryImageOk() (*string, bool)`

GetRepositoryImageOk returns a tuple with the RepositoryImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryImage

`func (o *GetContainer200ResponseContainer) SetRepositoryImage(v string)`

SetRepositoryImage sets RepositoryImage field to given value.

### HasRepositoryImage

`func (o *GetContainer200ResponseContainer) HasRepositoryImage() bool`

HasRepositoryImage returns a boolean if a field has been set.

### GetPlanCategory

`func (o *GetContainer200ResponseContainer) GetPlanCategory() string`

GetPlanCategory returns the PlanCategory field if non-nil, zero value otherwise.

### GetPlanCategoryOk

`func (o *GetContainer200ResponseContainer) GetPlanCategoryOk() (*string, bool)`

GetPlanCategoryOk returns a tuple with the PlanCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCategory

`func (o *GetContainer200ResponseContainer) SetPlanCategory(v string)`

SetPlanCategory sets PlanCategory field to given value.

### HasPlanCategory

`func (o *GetContainer200ResponseContainer) HasPlanCategory() bool`

HasPlanCategory returns a boolean if a field has been set.

### GetHostname

`func (o *GetContainer200ResponseContainer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetContainer200ResponseContainer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetContainer200ResponseContainer) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetContainer200ResponseContainer) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDomainName

`func (o *GetContainer200ResponseContainer) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *GetContainer200ResponseContainer) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *GetContainer200ResponseContainer) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *GetContainer200ResponseContainer) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetVolumeCreated

`func (o *GetContainer200ResponseContainer) GetVolumeCreated() bool`

GetVolumeCreated returns the VolumeCreated field if non-nil, zero value otherwise.

### GetVolumeCreatedOk

`func (o *GetContainer200ResponseContainer) GetVolumeCreatedOk() (*bool, bool)`

GetVolumeCreatedOk returns a tuple with the VolumeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCreated

`func (o *GetContainer200ResponseContainer) SetVolumeCreated(v bool)`

SetVolumeCreated sets VolumeCreated field to given value.

### HasVolumeCreated

`func (o *GetContainer200ResponseContainer) HasVolumeCreated() bool`

HasVolumeCreated returns a boolean if a field has been set.

### GetContainerCreated

`func (o *GetContainer200ResponseContainer) GetContainerCreated() bool`

GetContainerCreated returns the ContainerCreated field if non-nil, zero value otherwise.

### GetContainerCreatedOk

`func (o *GetContainer200ResponseContainer) GetContainerCreatedOk() (*bool, bool)`

GetContainerCreatedOk returns a tuple with the ContainerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCreated

`func (o *GetContainer200ResponseContainer) SetContainerCreated(v bool)`

SetContainerCreated sets ContainerCreated field to given value.

### HasContainerCreated

`func (o *GetContainer200ResponseContainer) HasContainerCreated() bool`

HasContainerCreated returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GetContainer200ResponseContainer) GetMaxStorage() int32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetContainer200ResponseContainer) GetMaxStorageOk() (*int32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetContainer200ResponseContainer) SetMaxStorage(v int32)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetContainer200ResponseContainer) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetContainer200ResponseContainer) GetMaxMemory() int32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetContainer200ResponseContainer) GetMaxMemoryOk() (*int32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetContainer200ResponseContainer) SetMaxMemory(v int32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetContainer200ResponseContainer) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCores

`func (o *GetContainer200ResponseContainer) GetMaxCores() int32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *GetContainer200ResponseContainer) GetMaxCoresOk() (*int32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *GetContainer200ResponseContainer) SetMaxCores(v int32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *GetContainer200ResponseContainer) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxCpu

`func (o *GetContainer200ResponseContainer) GetMaxCpu() int32`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GetContainer200ResponseContainer) GetMaxCpuOk() (*int32, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GetContainer200ResponseContainer) SetMaxCpu(v int32)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GetContainer200ResponseContainer) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetAvailableActions

`func (o *GetContainer200ResponseContainer) GetAvailableActions() []map[string]interface{}`

GetAvailableActions returns the AvailableActions field if non-nil, zero value otherwise.

### GetAvailableActionsOk

`func (o *GetContainer200ResponseContainer) GetAvailableActionsOk() (*[]map[string]interface{}, bool)`

GetAvailableActionsOk returns a tuple with the AvailableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActions

`func (o *GetContainer200ResponseContainer) SetAvailableActions(v []map[string]interface{})`

SetAvailableActions sets AvailableActions field to given value.

### HasAvailableActions

`func (o *GetContainer200ResponseContainer) HasAvailableActions() bool`

HasAvailableActions returns a boolean if a field has been set.

### GetConfigGroup

`func (o *GetContainer200ResponseContainer) GetConfigGroup() string`

GetConfigGroup returns the ConfigGroup field if non-nil, zero value otherwise.

### GetConfigGroupOk

`func (o *GetContainer200ResponseContainer) GetConfigGroupOk() (*string, bool)`

GetConfigGroupOk returns a tuple with the ConfigGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGroup

`func (o *GetContainer200ResponseContainer) SetConfigGroup(v string)`

SetConfigGroup sets ConfigGroup field to given value.

### HasConfigGroup

`func (o *GetContainer200ResponseContainer) HasConfigGroup() bool`

HasConfigGroup returns a boolean if a field has been set.

### GetConfigId

`func (o *GetContainer200ResponseContainer) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *GetContainer200ResponseContainer) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *GetContainer200ResponseContainer) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *GetContainer200ResponseContainer) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetConfigRole

`func (o *GetContainer200ResponseContainer) GetConfigRole() string`

GetConfigRole returns the ConfigRole field if non-nil, zero value otherwise.

### GetConfigRoleOk

`func (o *GetContainer200ResponseContainer) GetConfigRoleOk() (*string, bool)`

GetConfigRoleOk returns a tuple with the ConfigRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRole

`func (o *GetContainer200ResponseContainer) SetConfigRole(v string)`

SetConfigRole sets ConfigRole field to given value.

### HasConfigRole

`func (o *GetContainer200ResponseContainer) HasConfigRole() bool`

HasConfigRole returns a boolean if a field has been set.

### GetHourlyCost

`func (o *GetContainer200ResponseContainer) GetHourlyCost() float64`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *GetContainer200ResponseContainer) GetHourlyCostOk() (*float64, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *GetContainer200ResponseContainer) SetHourlyCost(v float64)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *GetContainer200ResponseContainer) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *GetContainer200ResponseContainer) GetHourlyPrice() float64`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *GetContainer200ResponseContainer) GetHourlyPriceOk() (*float64, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *GetContainer200ResponseContainer) SetHourlyPrice(v float64)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *GetContainer200ResponseContainer) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ClusterContainers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Instance** | Pointer to **string** |  | [optional] 
**ContainerType** | Pointer to [**ListClusterContainers200ResponseAllOfContainersInnerContainerType**](ListClusterContainers200ResponseAllOfContainersInnerContainerType.md) |  | [optional] 
**ContainerTypeSet** | Pointer to [**ListClusterContainers200ResponseAllOfContainersInnerContainerTypeSet**](ListClusterContainers200ResponseAllOfContainersInnerContainerTypeSet.md) |  | [optional] 
**Server** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Cloud** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**InternalHostname** | Pointer to **string** |  | [optional] 
**ExternalHostname** | Pointer to **string** |  | [optional] 
**ExternalDomain** | Pointer to **string** |  | [optional] 
**ExternalFqdn** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Plan** | Pointer to [**ListClusterContainers200ResponseAllOfContainersInnerPlan**](ListClusterContainers200ResponseAllOfContainersInnerPlan.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**StatsEnabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UserStatus** | Pointer to **string** |  | [optional] 
**EnvironmentPrefix** | Pointer to **string** |  | [optional] 
**ConfigGroup** | Pointer to **string** |  | [optional] 
**ConfigId** | Pointer to **string** |  | [optional] 
**ConfigRole** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**ListClusterContainers200ResponseAllOfContainersInnerStats**](ListClusterContainers200ResponseAllOfContainersInnerStats.md) |  | [optional] 
**RuntimeInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerVersion** | Pointer to **string** |  | [optional] 
**RepositoryImage** | Pointer to **string** |  | [optional] 
**PlanCategory** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**VolumeCreated** | Pointer to **bool** |  | [optional] 
**ContainerCreated** | Pointer to **bool** |  | [optional] 
**MaxStorage** | Pointer to **string** |  | [optional] 
**MaxMemory** | Pointer to **string** |  | [optional] 
**MaxCores** | Pointer to **string** |  | [optional] 
**MaxCpu** | Pointer to **string** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**AvailableActions** | Pointer to [**[]ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner**](ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner.md) |  | [optional] 

## Methods

### NewClusterContainers

`func NewClusterContainers() *ClusterContainers`

NewClusterContainers instantiates a new ClusterContainers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterContainersWithDefaults

`func NewClusterContainersWithDefaults() *ClusterContainers`

NewClusterContainersWithDefaults instantiates a new ClusterContainers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterContainers) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterContainers) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterContainers) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterContainers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ClusterContainers) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ClusterContainers) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ClusterContainers) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ClusterContainers) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccountId

`func (o *ClusterContainers) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ClusterContainers) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ClusterContainers) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ClusterContainers) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInstance

`func (o *ClusterContainers) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ClusterContainers) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ClusterContainers) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ClusterContainers) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerType

`func (o *ClusterContainers) GetContainerType() ListClusterContainers200ResponseAllOfContainersInnerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *ClusterContainers) GetContainerTypeOk() (*ListClusterContainers200ResponseAllOfContainersInnerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *ClusterContainers) SetContainerType(v ListClusterContainers200ResponseAllOfContainersInnerContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *ClusterContainers) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetContainerTypeSet

`func (o *ClusterContainers) GetContainerTypeSet() ListClusterContainers200ResponseAllOfContainersInnerContainerTypeSet`

GetContainerTypeSet returns the ContainerTypeSet field if non-nil, zero value otherwise.

### GetContainerTypeSetOk

`func (o *ClusterContainers) GetContainerTypeSetOk() (*ListClusterContainers200ResponseAllOfContainersInnerContainerTypeSet, bool)`

GetContainerTypeSetOk returns a tuple with the ContainerTypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypeSet

`func (o *ClusterContainers) SetContainerTypeSet(v ListClusterContainers200ResponseAllOfContainersInnerContainerTypeSet)`

SetContainerTypeSet sets ContainerTypeSet field to given value.

### HasContainerTypeSet

`func (o *ClusterContainers) HasContainerTypeSet() bool`

HasContainerTypeSet returns a boolean if a field has been set.

### GetServer

`func (o *ClusterContainers) GetServer() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ClusterContainers) GetServerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ClusterContainers) SetServer(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetServer sets Server field to given value.

### HasServer

`func (o *ClusterContainers) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetCloud

`func (o *ClusterContainers) GetCloud() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ClusterContainers) GetCloudOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ClusterContainers) SetCloud(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ClusterContainers) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetName

`func (o *ClusterContainers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterContainers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterContainers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterContainers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *ClusterContainers) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClusterContainers) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClusterContainers) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClusterContainers) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *ClusterContainers) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ClusterContainers) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ClusterContainers) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ClusterContainers) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetInternalHostname

`func (o *ClusterContainers) GetInternalHostname() string`

GetInternalHostname returns the InternalHostname field if non-nil, zero value otherwise.

### GetInternalHostnameOk

`func (o *ClusterContainers) GetInternalHostnameOk() (*string, bool)`

GetInternalHostnameOk returns a tuple with the InternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostname

`func (o *ClusterContainers) SetInternalHostname(v string)`

SetInternalHostname sets InternalHostname field to given value.

### HasInternalHostname

`func (o *ClusterContainers) HasInternalHostname() bool`

HasInternalHostname returns a boolean if a field has been set.

### GetExternalHostname

`func (o *ClusterContainers) GetExternalHostname() string`

GetExternalHostname returns the ExternalHostname field if non-nil, zero value otherwise.

### GetExternalHostnameOk

`func (o *ClusterContainers) GetExternalHostnameOk() (*string, bool)`

GetExternalHostnameOk returns a tuple with the ExternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHostname

`func (o *ClusterContainers) SetExternalHostname(v string)`

SetExternalHostname sets ExternalHostname field to given value.

### HasExternalHostname

`func (o *ClusterContainers) HasExternalHostname() bool`

HasExternalHostname returns a boolean if a field has been set.

### GetExternalDomain

`func (o *ClusterContainers) GetExternalDomain() string`

GetExternalDomain returns the ExternalDomain field if non-nil, zero value otherwise.

### GetExternalDomainOk

`func (o *ClusterContainers) GetExternalDomainOk() (*string, bool)`

GetExternalDomainOk returns a tuple with the ExternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDomain

`func (o *ClusterContainers) SetExternalDomain(v string)`

SetExternalDomain sets ExternalDomain field to given value.

### HasExternalDomain

`func (o *ClusterContainers) HasExternalDomain() bool`

HasExternalDomain returns a boolean if a field has been set.

### GetExternalFqdn

`func (o *ClusterContainers) GetExternalFqdn() string`

GetExternalFqdn returns the ExternalFqdn field if non-nil, zero value otherwise.

### GetExternalFqdnOk

`func (o *ClusterContainers) GetExternalFqdnOk() (*string, bool)`

GetExternalFqdnOk returns a tuple with the ExternalFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqdn

`func (o *ClusterContainers) SetExternalFqdn(v string)`

SetExternalFqdn sets ExternalFqdn field to given value.

### HasExternalFqdn

`func (o *ClusterContainers) HasExternalFqdn() bool`

HasExternalFqdn returns a boolean if a field has been set.

### GetPorts

`func (o *ClusterContainers) GetPorts() []map[string]interface{}`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ClusterContainers) GetPortsOk() (*[]map[string]interface{}, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ClusterContainers) SetPorts(v []map[string]interface{})`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ClusterContainers) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPlan

`func (o *ClusterContainers) GetPlan() ListClusterContainers200ResponseAllOfContainersInnerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ClusterContainers) GetPlanOk() (*ListClusterContainers200ResponseAllOfContainersInnerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ClusterContainers) SetPlan(v ListClusterContainers200ResponseAllOfContainersInnerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ClusterContainers) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetDateCreated

`func (o *ClusterContainers) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ClusterContainers) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ClusterContainers) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ClusterContainers) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ClusterContainers) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ClusterContainers) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ClusterContainers) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ClusterContainers) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatsEnabled

`func (o *ClusterContainers) GetStatsEnabled() bool`

GetStatsEnabled returns the StatsEnabled field if non-nil, zero value otherwise.

### GetStatsEnabledOk

`func (o *ClusterContainers) GetStatsEnabledOk() (*bool, bool)`

GetStatsEnabledOk returns a tuple with the StatsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsEnabled

`func (o *ClusterContainers) SetStatsEnabled(v bool)`

SetStatsEnabled sets StatsEnabled field to given value.

### HasStatsEnabled

`func (o *ClusterContainers) HasStatsEnabled() bool`

HasStatsEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterContainers) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterContainers) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterContainers) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterContainers) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserStatus

`func (o *ClusterContainers) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *ClusterContainers) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *ClusterContainers) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *ClusterContainers) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *ClusterContainers) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *ClusterContainers) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *ClusterContainers) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *ClusterContainers) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetConfigGroup

`func (o *ClusterContainers) GetConfigGroup() string`

GetConfigGroup returns the ConfigGroup field if non-nil, zero value otherwise.

### GetConfigGroupOk

`func (o *ClusterContainers) GetConfigGroupOk() (*string, bool)`

GetConfigGroupOk returns a tuple with the ConfigGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGroup

`func (o *ClusterContainers) SetConfigGroup(v string)`

SetConfigGroup sets ConfigGroup field to given value.

### HasConfigGroup

`func (o *ClusterContainers) HasConfigGroup() bool`

HasConfigGroup returns a boolean if a field has been set.

### GetConfigId

`func (o *ClusterContainers) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ClusterContainers) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ClusterContainers) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ClusterContainers) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetConfigRole

`func (o *ClusterContainers) GetConfigRole() string`

GetConfigRole returns the ConfigRole field if non-nil, zero value otherwise.

### GetConfigRoleOk

`func (o *ClusterContainers) GetConfigRoleOk() (*string, bool)`

GetConfigRoleOk returns a tuple with the ConfigRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRole

`func (o *ClusterContainers) SetConfigRole(v string)`

SetConfigRole sets ConfigRole field to given value.

### HasConfigRole

`func (o *ClusterContainers) HasConfigRole() bool`

HasConfigRole returns a boolean if a field has been set.

### GetStats

`func (o *ClusterContainers) GetStats() ListClusterContainers200ResponseAllOfContainersInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ClusterContainers) GetStatsOk() (*ListClusterContainers200ResponseAllOfContainersInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ClusterContainers) SetStats(v ListClusterContainers200ResponseAllOfContainersInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ClusterContainers) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetRuntimeInfo

`func (o *ClusterContainers) GetRuntimeInfo() map[string]interface{}`

GetRuntimeInfo returns the RuntimeInfo field if non-nil, zero value otherwise.

### GetRuntimeInfoOk

`func (o *ClusterContainers) GetRuntimeInfoOk() (*map[string]interface{}, bool)`

GetRuntimeInfoOk returns a tuple with the RuntimeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeInfo

`func (o *ClusterContainers) SetRuntimeInfo(v map[string]interface{})`

SetRuntimeInfo sets RuntimeInfo field to given value.

### HasRuntimeInfo

`func (o *ClusterContainers) HasRuntimeInfo() bool`

HasRuntimeInfo returns a boolean if a field has been set.

### GetContainerVersion

`func (o *ClusterContainers) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *ClusterContainers) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *ClusterContainers) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *ClusterContainers) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetRepositoryImage

`func (o *ClusterContainers) GetRepositoryImage() string`

GetRepositoryImage returns the RepositoryImage field if non-nil, zero value otherwise.

### GetRepositoryImageOk

`func (o *ClusterContainers) GetRepositoryImageOk() (*string, bool)`

GetRepositoryImageOk returns a tuple with the RepositoryImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryImage

`func (o *ClusterContainers) SetRepositoryImage(v string)`

SetRepositoryImage sets RepositoryImage field to given value.

### HasRepositoryImage

`func (o *ClusterContainers) HasRepositoryImage() bool`

HasRepositoryImage returns a boolean if a field has been set.

### GetPlanCategory

`func (o *ClusterContainers) GetPlanCategory() string`

GetPlanCategory returns the PlanCategory field if non-nil, zero value otherwise.

### GetPlanCategoryOk

`func (o *ClusterContainers) GetPlanCategoryOk() (*string, bool)`

GetPlanCategoryOk returns a tuple with the PlanCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCategory

`func (o *ClusterContainers) SetPlanCategory(v string)`

SetPlanCategory sets PlanCategory field to given value.

### HasPlanCategory

`func (o *ClusterContainers) HasPlanCategory() bool`

HasPlanCategory returns a boolean if a field has been set.

### GetHostname

`func (o *ClusterContainers) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ClusterContainers) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ClusterContainers) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ClusterContainers) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDomainName

`func (o *ClusterContainers) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ClusterContainers) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ClusterContainers) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ClusterContainers) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetVolumeCreated

`func (o *ClusterContainers) GetVolumeCreated() bool`

GetVolumeCreated returns the VolumeCreated field if non-nil, zero value otherwise.

### GetVolumeCreatedOk

`func (o *ClusterContainers) GetVolumeCreatedOk() (*bool, bool)`

GetVolumeCreatedOk returns a tuple with the VolumeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCreated

`func (o *ClusterContainers) SetVolumeCreated(v bool)`

SetVolumeCreated sets VolumeCreated field to given value.

### HasVolumeCreated

`func (o *ClusterContainers) HasVolumeCreated() bool`

HasVolumeCreated returns a boolean if a field has been set.

### GetContainerCreated

`func (o *ClusterContainers) GetContainerCreated() bool`

GetContainerCreated returns the ContainerCreated field if non-nil, zero value otherwise.

### GetContainerCreatedOk

`func (o *ClusterContainers) GetContainerCreatedOk() (*bool, bool)`

GetContainerCreatedOk returns a tuple with the ContainerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCreated

`func (o *ClusterContainers) SetContainerCreated(v bool)`

SetContainerCreated sets ContainerCreated field to given value.

### HasContainerCreated

`func (o *ClusterContainers) HasContainerCreated() bool`

HasContainerCreated returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ClusterContainers) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ClusterContainers) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ClusterContainers) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ClusterContainers) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ClusterContainers) GetMaxMemory() string`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ClusterContainers) GetMaxMemoryOk() (*string, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ClusterContainers) SetMaxMemory(v string)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ClusterContainers) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCores

`func (o *ClusterContainers) GetMaxCores() string`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ClusterContainers) GetMaxCoresOk() (*string, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ClusterContainers) SetMaxCores(v string)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ClusterContainers) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ClusterContainers) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ClusterContainers) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ClusterContainers) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ClusterContainers) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *ClusterContainers) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *ClusterContainers) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *ClusterContainers) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *ClusterContainers) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetAvailableActions

`func (o *ClusterContainers) GetAvailableActions() []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner`

GetAvailableActions returns the AvailableActions field if non-nil, zero value otherwise.

### GetAvailableActionsOk

`func (o *ClusterContainers) GetAvailableActionsOk() (*[]ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner, bool)`

GetAvailableActionsOk returns a tuple with the AvailableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActions

`func (o *ClusterContainers) SetAvailableActions(v []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner)`

SetAvailableActions sets AvailableActions field to given value.

### HasAvailableActions

`func (o *ClusterContainers) HasAvailableActions() bool`

HasAvailableActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



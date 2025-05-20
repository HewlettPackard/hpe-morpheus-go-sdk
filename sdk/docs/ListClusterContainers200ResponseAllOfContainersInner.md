# ListClusterContainers200ResponseAllOfContainersInner

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

### NewListClusterContainers200ResponseAllOfContainersInner

`func NewListClusterContainers200ResponseAllOfContainersInner() *ListClusterContainers200ResponseAllOfContainersInner`

NewListClusterContainers200ResponseAllOfContainersInner instantiates a new ListClusterContainers200ResponseAllOfContainersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterContainers200ResponseAllOfContainersInnerWithDefaults

`func NewListClusterContainers200ResponseAllOfContainersInnerWithDefaults() *ListClusterContainers200ResponseAllOfContainersInner`

NewListClusterContainers200ResponseAllOfContainersInnerWithDefaults instantiates a new ListClusterContainers200ResponseAllOfContainersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccountId

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInstance

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerType

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetContainerType() ListClusterContainers200ResponseAllOfContainersInnerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetContainerTypeOk() (*ListClusterContainers200ResponseAllOfContainersInnerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetContainerType(v ListClusterContainers200ResponseAllOfContainersInnerContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetContainerTypeSet

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetContainerTypeSet() ListClusterContainers200ResponseAllOfContainersInnerContainerTypeSet`

GetContainerTypeSet returns the ContainerTypeSet field if non-nil, zero value otherwise.

### GetContainerTypeSetOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetContainerTypeSetOk() (*ListClusterContainers200ResponseAllOfContainersInnerContainerTypeSet, bool)`

GetContainerTypeSetOk returns a tuple with the ContainerTypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypeSet

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetContainerTypeSet(v ListClusterContainers200ResponseAllOfContainersInnerContainerTypeSet)`

SetContainerTypeSet sets ContainerTypeSet field to given value.

### HasContainerTypeSet

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasContainerTypeSet() bool`

HasContainerTypeSet returns a boolean if a field has been set.

### GetServer

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetServer() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetServerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetServer(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetServer sets Server field to given value.

### HasServer

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetCloud

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetCloud() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetCloudOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetCloud(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetName

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetInternalHostname

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetInternalHostname() string`

GetInternalHostname returns the InternalHostname field if non-nil, zero value otherwise.

### GetInternalHostnameOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetInternalHostnameOk() (*string, bool)`

GetInternalHostnameOk returns a tuple with the InternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostname

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetInternalHostname(v string)`

SetInternalHostname sets InternalHostname field to given value.

### HasInternalHostname

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasInternalHostname() bool`

HasInternalHostname returns a boolean if a field has been set.

### GetExternalHostname

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetExternalHostname() string`

GetExternalHostname returns the ExternalHostname field if non-nil, zero value otherwise.

### GetExternalHostnameOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetExternalHostnameOk() (*string, bool)`

GetExternalHostnameOk returns a tuple with the ExternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHostname

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetExternalHostname(v string)`

SetExternalHostname sets ExternalHostname field to given value.

### HasExternalHostname

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasExternalHostname() bool`

HasExternalHostname returns a boolean if a field has been set.

### GetExternalDomain

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetExternalDomain() string`

GetExternalDomain returns the ExternalDomain field if non-nil, zero value otherwise.

### GetExternalDomainOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetExternalDomainOk() (*string, bool)`

GetExternalDomainOk returns a tuple with the ExternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDomain

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetExternalDomain(v string)`

SetExternalDomain sets ExternalDomain field to given value.

### HasExternalDomain

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasExternalDomain() bool`

HasExternalDomain returns a boolean if a field has been set.

### GetExternalFqdn

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetExternalFqdn() string`

GetExternalFqdn returns the ExternalFqdn field if non-nil, zero value otherwise.

### GetExternalFqdnOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetExternalFqdnOk() (*string, bool)`

GetExternalFqdnOk returns a tuple with the ExternalFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqdn

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetExternalFqdn(v string)`

SetExternalFqdn sets ExternalFqdn field to given value.

### HasExternalFqdn

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasExternalFqdn() bool`

HasExternalFqdn returns a boolean if a field has been set.

### GetPorts

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetPorts() []map[string]interface{}`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetPortsOk() (*[]map[string]interface{}, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetPorts(v []map[string]interface{})`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPlan

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetPlan() ListClusterContainers200ResponseAllOfContainersInnerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetPlanOk() (*ListClusterContainers200ResponseAllOfContainersInnerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetPlan(v ListClusterContainers200ResponseAllOfContainersInnerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatsEnabled

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetStatsEnabled() bool`

GetStatsEnabled returns the StatsEnabled field if non-nil, zero value otherwise.

### GetStatsEnabledOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetStatsEnabledOk() (*bool, bool)`

GetStatsEnabledOk returns a tuple with the StatsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsEnabled

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetStatsEnabled(v bool)`

SetStatsEnabled sets StatsEnabled field to given value.

### HasStatsEnabled

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasStatsEnabled() bool`

HasStatsEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserStatus

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetConfigGroup

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetConfigGroup() string`

GetConfigGroup returns the ConfigGroup field if non-nil, zero value otherwise.

### GetConfigGroupOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetConfigGroupOk() (*string, bool)`

GetConfigGroupOk returns a tuple with the ConfigGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGroup

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetConfigGroup(v string)`

SetConfigGroup sets ConfigGroup field to given value.

### HasConfigGroup

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasConfigGroup() bool`

HasConfigGroup returns a boolean if a field has been set.

### GetConfigId

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetConfigRole

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetConfigRole() string`

GetConfigRole returns the ConfigRole field if non-nil, zero value otherwise.

### GetConfigRoleOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetConfigRoleOk() (*string, bool)`

GetConfigRoleOk returns a tuple with the ConfigRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRole

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetConfigRole(v string)`

SetConfigRole sets ConfigRole field to given value.

### HasConfigRole

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasConfigRole() bool`

HasConfigRole returns a boolean if a field has been set.

### GetStats

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetStats() ListClusterContainers200ResponseAllOfContainersInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetStatsOk() (*ListClusterContainers200ResponseAllOfContainersInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetStats(v ListClusterContainers200ResponseAllOfContainersInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetRuntimeInfo

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetRuntimeInfo() map[string]interface{}`

GetRuntimeInfo returns the RuntimeInfo field if non-nil, zero value otherwise.

### GetRuntimeInfoOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetRuntimeInfoOk() (*map[string]interface{}, bool)`

GetRuntimeInfoOk returns a tuple with the RuntimeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeInfo

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetRuntimeInfo(v map[string]interface{})`

SetRuntimeInfo sets RuntimeInfo field to given value.

### HasRuntimeInfo

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasRuntimeInfo() bool`

HasRuntimeInfo returns a boolean if a field has been set.

### GetContainerVersion

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetRepositoryImage

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetRepositoryImage() string`

GetRepositoryImage returns the RepositoryImage field if non-nil, zero value otherwise.

### GetRepositoryImageOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetRepositoryImageOk() (*string, bool)`

GetRepositoryImageOk returns a tuple with the RepositoryImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryImage

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetRepositoryImage(v string)`

SetRepositoryImage sets RepositoryImage field to given value.

### HasRepositoryImage

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasRepositoryImage() bool`

HasRepositoryImage returns a boolean if a field has been set.

### GetPlanCategory

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetPlanCategory() string`

GetPlanCategory returns the PlanCategory field if non-nil, zero value otherwise.

### GetPlanCategoryOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetPlanCategoryOk() (*string, bool)`

GetPlanCategoryOk returns a tuple with the PlanCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCategory

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetPlanCategory(v string)`

SetPlanCategory sets PlanCategory field to given value.

### HasPlanCategory

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasPlanCategory() bool`

HasPlanCategory returns a boolean if a field has been set.

### GetHostname

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDomainName

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetVolumeCreated

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetVolumeCreated() bool`

GetVolumeCreated returns the VolumeCreated field if non-nil, zero value otherwise.

### GetVolumeCreatedOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetVolumeCreatedOk() (*bool, bool)`

GetVolumeCreatedOk returns a tuple with the VolumeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCreated

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetVolumeCreated(v bool)`

SetVolumeCreated sets VolumeCreated field to given value.

### HasVolumeCreated

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasVolumeCreated() bool`

HasVolumeCreated returns a boolean if a field has been set.

### GetContainerCreated

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetContainerCreated() bool`

GetContainerCreated returns the ContainerCreated field if non-nil, zero value otherwise.

### GetContainerCreatedOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetContainerCreatedOk() (*bool, bool)`

GetContainerCreatedOk returns a tuple with the ContainerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCreated

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetContainerCreated(v bool)`

SetContainerCreated sets ContainerCreated field to given value.

### HasContainerCreated

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasContainerCreated() bool`

HasContainerCreated returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetMaxMemory() string`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetMaxMemoryOk() (*string, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetMaxMemory(v string)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCores

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetMaxCores() string`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetMaxCoresOk() (*string, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetMaxCores(v string)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetAvailableActions

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetAvailableActions() []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner`

GetAvailableActions returns the AvailableActions field if non-nil, zero value otherwise.

### GetAvailableActionsOk

`func (o *ListClusterContainers200ResponseAllOfContainersInner) GetAvailableActionsOk() (*[]ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner, bool)`

GetAvailableActionsOk returns a tuple with the AvailableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActions

`func (o *ListClusterContainers200ResponseAllOfContainersInner) SetAvailableActions(v []ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner)`

SetAvailableActions sets AvailableActions field to given value.

### HasAvailableActions

`func (o *ListClusterContainers200ResponseAllOfContainersInner) HasAvailableActions() bool`

HasAvailableActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



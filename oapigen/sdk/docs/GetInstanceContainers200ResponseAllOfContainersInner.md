# GetInstanceContainers200ResponseAllOfContainersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int32** |  | [optional] 
**Instance** | Pointer to [**GetInstanceContainers200ResponseAllOfContainersInnerInstance**](GetInstanceContainers200ResponseAllOfContainersInnerInstance.md) |  | [optional] 
**ContainerType** | Pointer to [**GetInstanceContainers200ResponseAllOfContainersInnerContainerType**](GetInstanceContainers200ResponseAllOfContainersInnerContainerType.md) |  | [optional] 
**ContainerTypeSet** | Pointer to [**GetInstanceContainers200ResponseAllOfContainersInnerContainerTypeSet**](GetInstanceContainers200ResponseAllOfContainersInnerContainerTypeSet.md) |  | [optional] 
**Server** | Pointer to [**GetInstanceContainers200ResponseAllOfContainersInnerServer**](GetInstanceContainers200ResponseAllOfContainersInnerServer.md) |  | [optional] 
**Cloud** | Pointer to [**GetInstanceContainers200ResponseAllOfContainersInnerCloud**](GetInstanceContainers200ResponseAllOfContainersInnerCloud.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**InternalHostname** | Pointer to **string** |  | [optional] 
**ExternalHostname** | Pointer to **string** |  | [optional] 
**ExternalDomain** | Pointer to **string** |  | [optional] 
**ExternalFqdn** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to [**[]GetInstanceContainers200ResponseAllOfContainersInnerPortsInner**](GetInstanceContainers200ResponseAllOfContainersInnerPortsInner.md) |  | [optional] 
**Plan** | Pointer to [**GetInstanceContainers200ResponseAllOfContainersInnerPlan**](GetInstanceContainers200ResponseAllOfContainersInnerPlan.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**StatsEnabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UserStatus** | Pointer to **string** |  | [optional] 
**EnvironmentPrefix** | Pointer to **NullableString** |  | [optional] 
**Stats** | Pointer to [**GetInstanceContainers200ResponseAllOfContainersInnerStats**](GetInstanceContainers200ResponseAllOfContainersInnerStats.md) |  | [optional] 
**RuntimeInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerVersion** | Pointer to **NullableString** |  | [optional] 
**RepositoryImage** | Pointer to **NullableString** |  | [optional] 
**PlanCategory** | Pointer to **NullableString** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **NullableString** |  | [optional] 
**VolumeCreated** | Pointer to **bool** |  | [optional] 
**ContainerCreated** | Pointer to **bool** |  | [optional] 
**MaxStorage** | Pointer to **int32** |  | [optional] 
**MaxMemory** | Pointer to **int32** |  | [optional] 
**MaxCores** | Pointer to **int32** |  | [optional] 
**MaxCpu** | Pointer to **NullableInt32** |  | [optional] 
**AvailableActions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ConfigGroup** | Pointer to **NullableString** |  | [optional] 
**ConfigId** | Pointer to **NullableString** |  | [optional] 
**ConfigRole** | Pointer to **NullableString** |  | [optional] 
**HourlyCost** | Pointer to **float64** |  | [optional] 
**HourlyPrice** | Pointer to **float64** |  | [optional] 

## Methods

### NewGetInstanceContainers200ResponseAllOfContainersInner

`func NewGetInstanceContainers200ResponseAllOfContainersInner() *GetInstanceContainers200ResponseAllOfContainersInner`

NewGetInstanceContainers200ResponseAllOfContainersInner instantiates a new GetInstanceContainers200ResponseAllOfContainersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccountId

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInstance

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetInstance() GetInstanceContainers200ResponseAllOfContainersInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetInstanceOk() (*GetInstanceContainers200ResponseAllOfContainersInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetInstance(v GetInstanceContainers200ResponseAllOfContainersInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetContainerType

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetContainerType() GetInstanceContainers200ResponseAllOfContainersInnerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetContainerTypeOk() (*GetInstanceContainers200ResponseAllOfContainersInnerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetContainerType(v GetInstanceContainers200ResponseAllOfContainersInnerContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetContainerTypeSet

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetContainerTypeSet() GetInstanceContainers200ResponseAllOfContainersInnerContainerTypeSet`

GetContainerTypeSet returns the ContainerTypeSet field if non-nil, zero value otherwise.

### GetContainerTypeSetOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetContainerTypeSetOk() (*GetInstanceContainers200ResponseAllOfContainersInnerContainerTypeSet, bool)`

GetContainerTypeSetOk returns a tuple with the ContainerTypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypeSet

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetContainerTypeSet(v GetInstanceContainers200ResponseAllOfContainersInnerContainerTypeSet)`

SetContainerTypeSet sets ContainerTypeSet field to given value.

### HasContainerTypeSet

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasContainerTypeSet() bool`

HasContainerTypeSet returns a boolean if a field has been set.

### GetServer

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetServer() GetInstanceContainers200ResponseAllOfContainersInnerServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetServerOk() (*GetInstanceContainers200ResponseAllOfContainersInnerServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetServer(v GetInstanceContainers200ResponseAllOfContainersInnerServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetCloud

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetCloud() GetInstanceContainers200ResponseAllOfContainersInnerCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetCloudOk() (*GetInstanceContainers200ResponseAllOfContainersInnerCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetCloud(v GetInstanceContainers200ResponseAllOfContainersInnerCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetName

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetInternalHostname

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetInternalHostname() string`

GetInternalHostname returns the InternalHostname field if non-nil, zero value otherwise.

### GetInternalHostnameOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetInternalHostnameOk() (*string, bool)`

GetInternalHostnameOk returns a tuple with the InternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostname

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetInternalHostname(v string)`

SetInternalHostname sets InternalHostname field to given value.

### HasInternalHostname

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasInternalHostname() bool`

HasInternalHostname returns a boolean if a field has been set.

### GetExternalHostname

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetExternalHostname() string`

GetExternalHostname returns the ExternalHostname field if non-nil, zero value otherwise.

### GetExternalHostnameOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetExternalHostnameOk() (*string, bool)`

GetExternalHostnameOk returns a tuple with the ExternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHostname

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetExternalHostname(v string)`

SetExternalHostname sets ExternalHostname field to given value.

### HasExternalHostname

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasExternalHostname() bool`

HasExternalHostname returns a boolean if a field has been set.

### GetExternalDomain

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetExternalDomain() string`

GetExternalDomain returns the ExternalDomain field if non-nil, zero value otherwise.

### GetExternalDomainOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetExternalDomainOk() (*string, bool)`

GetExternalDomainOk returns a tuple with the ExternalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDomain

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetExternalDomain(v string)`

SetExternalDomain sets ExternalDomain field to given value.

### HasExternalDomain

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasExternalDomain() bool`

HasExternalDomain returns a boolean if a field has been set.

### GetExternalFqdn

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetExternalFqdn() string`

GetExternalFqdn returns the ExternalFqdn field if non-nil, zero value otherwise.

### GetExternalFqdnOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetExternalFqdnOk() (*string, bool)`

GetExternalFqdnOk returns a tuple with the ExternalFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFqdn

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetExternalFqdn(v string)`

SetExternalFqdn sets ExternalFqdn field to given value.

### HasExternalFqdn

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasExternalFqdn() bool`

HasExternalFqdn returns a boolean if a field has been set.

### GetPorts

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetPorts() []GetInstanceContainers200ResponseAllOfContainersInnerPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetPortsOk() (*[]GetInstanceContainers200ResponseAllOfContainersInnerPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetPorts(v []GetInstanceContainers200ResponseAllOfContainersInnerPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *GetInstanceContainers200ResponseAllOfContainersInner) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetPlan

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetPlan() GetInstanceContainers200ResponseAllOfContainersInnerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetPlanOk() (*GetInstanceContainers200ResponseAllOfContainersInnerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetPlan(v GetInstanceContainers200ResponseAllOfContainersInnerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatsEnabled

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetStatsEnabled() bool`

GetStatsEnabled returns the StatsEnabled field if non-nil, zero value otherwise.

### GetStatsEnabledOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetStatsEnabledOk() (*bool, bool)`

GetStatsEnabledOk returns a tuple with the StatsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsEnabled

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetStatsEnabled(v bool)`

SetStatsEnabled sets StatsEnabled field to given value.

### HasStatsEnabled

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasStatsEnabled() bool`

HasStatsEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserStatus

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### SetEnvironmentPrefixNil

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetEnvironmentPrefixNil(b bool)`

 SetEnvironmentPrefixNil sets the value for EnvironmentPrefix to be an explicit nil

### UnsetEnvironmentPrefix
`func (o *GetInstanceContainers200ResponseAllOfContainersInner) UnsetEnvironmentPrefix()`

UnsetEnvironmentPrefix ensures that no value is present for EnvironmentPrefix, not even an explicit nil
### GetStats

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetStats() GetInstanceContainers200ResponseAllOfContainersInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetStatsOk() (*GetInstanceContainers200ResponseAllOfContainersInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetStats(v GetInstanceContainers200ResponseAllOfContainersInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetRuntimeInfo

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetRuntimeInfo() map[string]interface{}`

GetRuntimeInfo returns the RuntimeInfo field if non-nil, zero value otherwise.

### GetRuntimeInfoOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetRuntimeInfoOk() (*map[string]interface{}, bool)`

GetRuntimeInfoOk returns a tuple with the RuntimeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeInfo

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetRuntimeInfo(v map[string]interface{})`

SetRuntimeInfo sets RuntimeInfo field to given value.

### HasRuntimeInfo

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasRuntimeInfo() bool`

HasRuntimeInfo returns a boolean if a field has been set.

### GetContainerVersion

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### SetContainerVersionNil

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetContainerVersionNil(b bool)`

 SetContainerVersionNil sets the value for ContainerVersion to be an explicit nil

### UnsetContainerVersion
`func (o *GetInstanceContainers200ResponseAllOfContainersInner) UnsetContainerVersion()`

UnsetContainerVersion ensures that no value is present for ContainerVersion, not even an explicit nil
### GetRepositoryImage

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetRepositoryImage() string`

GetRepositoryImage returns the RepositoryImage field if non-nil, zero value otherwise.

### GetRepositoryImageOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetRepositoryImageOk() (*string, bool)`

GetRepositoryImageOk returns a tuple with the RepositoryImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryImage

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetRepositoryImage(v string)`

SetRepositoryImage sets RepositoryImage field to given value.

### HasRepositoryImage

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasRepositoryImage() bool`

HasRepositoryImage returns a boolean if a field has been set.

### SetRepositoryImageNil

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetRepositoryImageNil(b bool)`

 SetRepositoryImageNil sets the value for RepositoryImage to be an explicit nil

### UnsetRepositoryImage
`func (o *GetInstanceContainers200ResponseAllOfContainersInner) UnsetRepositoryImage()`

UnsetRepositoryImage ensures that no value is present for RepositoryImage, not even an explicit nil
### GetPlanCategory

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetPlanCategory() string`

GetPlanCategory returns the PlanCategory field if non-nil, zero value otherwise.

### GetPlanCategoryOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetPlanCategoryOk() (*string, bool)`

GetPlanCategoryOk returns a tuple with the PlanCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCategory

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetPlanCategory(v string)`

SetPlanCategory sets PlanCategory field to given value.

### HasPlanCategory

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasPlanCategory() bool`

HasPlanCategory returns a boolean if a field has been set.

### SetPlanCategoryNil

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetPlanCategoryNil(b bool)`

 SetPlanCategoryNil sets the value for PlanCategory to be an explicit nil

### UnsetPlanCategory
`func (o *GetInstanceContainers200ResponseAllOfContainersInner) UnsetPlanCategory()`

UnsetPlanCategory ensures that no value is present for PlanCategory, not even an explicit nil
### GetHostname

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDomainName

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *GetInstanceContainers200ResponseAllOfContainersInner) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetVolumeCreated

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetVolumeCreated() bool`

GetVolumeCreated returns the VolumeCreated field if non-nil, zero value otherwise.

### GetVolumeCreatedOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetVolumeCreatedOk() (*bool, bool)`

GetVolumeCreatedOk returns a tuple with the VolumeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCreated

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetVolumeCreated(v bool)`

SetVolumeCreated sets VolumeCreated field to given value.

### HasVolumeCreated

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasVolumeCreated() bool`

HasVolumeCreated returns a boolean if a field has been set.

### GetContainerCreated

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetContainerCreated() bool`

GetContainerCreated returns the ContainerCreated field if non-nil, zero value otherwise.

### GetContainerCreatedOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetContainerCreatedOk() (*bool, bool)`

GetContainerCreatedOk returns a tuple with the ContainerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCreated

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetContainerCreated(v bool)`

SetContainerCreated sets ContainerCreated field to given value.

### HasContainerCreated

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasContainerCreated() bool`

HasContainerCreated returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetMaxStorage() int32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetMaxStorageOk() (*int32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetMaxStorage(v int32)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetMaxMemory() int32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetMaxMemoryOk() (*int32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetMaxMemory(v int32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCores

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetMaxCores() int32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetMaxCoresOk() (*int32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetMaxCores(v int32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxCpu

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetMaxCpu() int32`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetMaxCpuOk() (*int32, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetMaxCpu(v int32)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *GetInstanceContainers200ResponseAllOfContainersInner) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetAvailableActions

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetAvailableActions() []map[string]interface{}`

GetAvailableActions returns the AvailableActions field if non-nil, zero value otherwise.

### GetAvailableActionsOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetAvailableActionsOk() (*[]map[string]interface{}, bool)`

GetAvailableActionsOk returns a tuple with the AvailableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActions

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetAvailableActions(v []map[string]interface{})`

SetAvailableActions sets AvailableActions field to given value.

### HasAvailableActions

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasAvailableActions() bool`

HasAvailableActions returns a boolean if a field has been set.

### SetAvailableActionsNil

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetAvailableActionsNil(b bool)`

 SetAvailableActionsNil sets the value for AvailableActions to be an explicit nil

### UnsetAvailableActions
`func (o *GetInstanceContainers200ResponseAllOfContainersInner) UnsetAvailableActions()`

UnsetAvailableActions ensures that no value is present for AvailableActions, not even an explicit nil
### GetConfigGroup

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetConfigGroup() string`

GetConfigGroup returns the ConfigGroup field if non-nil, zero value otherwise.

### GetConfigGroupOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetConfigGroupOk() (*string, bool)`

GetConfigGroupOk returns a tuple with the ConfigGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGroup

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetConfigGroup(v string)`

SetConfigGroup sets ConfigGroup field to given value.

### HasConfigGroup

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasConfigGroup() bool`

HasConfigGroup returns a boolean if a field has been set.

### SetConfigGroupNil

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetConfigGroupNil(b bool)`

 SetConfigGroupNil sets the value for ConfigGroup to be an explicit nil

### UnsetConfigGroup
`func (o *GetInstanceContainers200ResponseAllOfContainersInner) UnsetConfigGroup()`

UnsetConfigGroup ensures that no value is present for ConfigGroup, not even an explicit nil
### GetConfigId

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### SetConfigIdNil

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetConfigIdNil(b bool)`

 SetConfigIdNil sets the value for ConfigId to be an explicit nil

### UnsetConfigId
`func (o *GetInstanceContainers200ResponseAllOfContainersInner) UnsetConfigId()`

UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
### GetConfigRole

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetConfigRole() string`

GetConfigRole returns the ConfigRole field if non-nil, zero value otherwise.

### GetConfigRoleOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetConfigRoleOk() (*string, bool)`

GetConfigRoleOk returns a tuple with the ConfigRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRole

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetConfigRole(v string)`

SetConfigRole sets ConfigRole field to given value.

### HasConfigRole

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasConfigRole() bool`

HasConfigRole returns a boolean if a field has been set.

### SetConfigRoleNil

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetConfigRoleNil(b bool)`

 SetConfigRoleNil sets the value for ConfigRole to be an explicit nil

### UnsetConfigRole
`func (o *GetInstanceContainers200ResponseAllOfContainersInner) UnsetConfigRole()`

UnsetConfigRole ensures that no value is present for ConfigRole, not even an explicit nil
### GetHourlyCost

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetHourlyCost() float64`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetHourlyCostOk() (*float64, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetHourlyCost(v float64)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetHourlyPrice() float64`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) GetHourlyPriceOk() (*float64, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) SetHourlyPrice(v float64)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *GetInstanceContainers200ResponseAllOfContainersInner) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Tenant** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**InstanceType** | Pointer to [**ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner**](ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner.md) |  | [optional] 
**Group** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Cloud** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Containers** | Pointer to **[]int64** |  | [optional] 
**Servers** | Pointer to **[]int64** |  | [optional] 
**ConnectionInfo** | Pointer to [**[]ListInstances200ResponseAllOfInstancesInnerConnectionInfoInner**](ListInstances200ResponseAllOfInstancesInnerConnectionInfoInner.md) |  | [optional] 
**Layout** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerLayout**](ListInstances200ResponseAllOfInstancesInnerLayout.md) |  | [optional] 
**Plan** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerConfig**](ListInstances200ResponseAllOfInstancesInnerConfig.md) |  | [optional] 
**ConfigGroup** | Pointer to **string** |  | [optional] 
**ConfigId** | Pointer to **string** |  | [optional] 
**ConfigRole** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]ListInstances200ResponseAllOfInstancesInnerVolumesInner**](ListInstances200ResponseAllOfInstancesInnerVolumesInner.md) |  | [optional] 
**Controllers** | Pointer to [**[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner.md) |  | [optional] 
**Interfaces** | Pointer to [**[]ListInstances200ResponseAllOfInstancesInnerInterfacesInner**](ListInstances200ResponseAllOfInstancesInnerInterfacesInner.md) |  | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**InstanceVersion** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Evars** | Pointer to **[]map[string]interface{}** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**CoresPerSocket** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **string** |  | [optional] 
**HourlyCost** | Pointer to **float32** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**InstancePrice** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerInstancePrice**](ListInstances200ResponseAllOfInstancesInnerInstancePrice.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**EnvironmentPrefix** | Pointer to **string** |  | [optional] 
**FirewallEnabled** | Pointer to **bool** |  | [optional] 
**NetworkLevel** | Pointer to **string** |  | [optional] 
**AutoScale** | Pointer to **bool** |  | [optional] 
**InstanceContext** | Pointer to **string** |  | [optional] 
**CurrentDeployId** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusPercent** | Pointer to **string** |  | [optional] 
**StatusEta** | Pointer to **string** |  | [optional] 
**UserStatus** | Pointer to **string** |  | [optional] 
**ExpireDays** | Pointer to **int64** |  | [optional] 
**RenewDays** | Pointer to **int64** |  | [optional] 
**ExpireCount** | Pointer to **int64** |  | [optional] 
**ExpireDate** | Pointer to **time.Time** |  | [optional] 
**ExpireWarningDate** | Pointer to **time.Time** |  | [optional] 
**ExpireWarningSent** | Pointer to **bool** |  | [optional] 
**ShutdownDays** | Pointer to **int64** |  | [optional] 
**ShutdownRenewDays** | Pointer to **int64** |  | [optional] 
**ShutdownCount** | Pointer to **int64** |  | [optional] 
**ShutdownDate** | Pointer to **time.Time** |  | [optional] 
**ShutdownWarningDate** | Pointer to **time.Time** |  | [optional] 
**ShutdownWarningSent** | Pointer to **bool** |  | [optional] 
**RemovalDate** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**Owner** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerStats**](ListInstances200ResponseAllOfInstancesInnerStats.md) |  | [optional] 
**PowerSchedule** | Pointer to **string** |  | [optional] 
**IsScalable** | Pointer to **bool** |  | [optional] 
**InstanceThreshold** | Pointer to **map[string]interface{}** |  | [optional] 
**IsBusy** | Pointer to **bool** |  | [optional] 
**Apps** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewInstance

`func NewInstance() *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Instance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *Instance) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Instance) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Instance) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Instance) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccountId

`func (o *Instance) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Instance) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Instance) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Instance) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetTenant

`func (o *Instance) GetTenant() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Instance) GetTenantOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Instance) SetTenant(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Instance) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetInstanceType

`func (o *Instance) GetInstanceType() ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Instance) GetInstanceTypeOk() (*ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Instance) SetInstanceType(v ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *Instance) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetGroup

`func (o *Instance) GetGroup() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Instance) GetGroupOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Instance) SetGroup(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Instance) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *Instance) GetCloud() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *Instance) GetCloudOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *Instance) SetCloud(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *Instance) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetContainers

`func (o *Instance) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *Instance) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *Instance) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *Instance) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetServers

`func (o *Instance) GetServers() []int64`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *Instance) GetServersOk() (*[]int64, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *Instance) SetServers(v []int64)`

SetServers sets Servers field to given value.

### HasServers

`func (o *Instance) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetConnectionInfo

`func (o *Instance) GetConnectionInfo() []ListInstances200ResponseAllOfInstancesInnerConnectionInfoInner`

GetConnectionInfo returns the ConnectionInfo field if non-nil, zero value otherwise.

### GetConnectionInfoOk

`func (o *Instance) GetConnectionInfoOk() (*[]ListInstances200ResponseAllOfInstancesInnerConnectionInfoInner, bool)`

GetConnectionInfoOk returns a tuple with the ConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionInfo

`func (o *Instance) SetConnectionInfo(v []ListInstances200ResponseAllOfInstancesInnerConnectionInfoInner)`

SetConnectionInfo sets ConnectionInfo field to given value.

### HasConnectionInfo

`func (o *Instance) HasConnectionInfo() bool`

HasConnectionInfo returns a boolean if a field has been set.

### GetLayout

`func (o *Instance) GetLayout() ListInstances200ResponseAllOfInstancesInnerLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Instance) GetLayoutOk() (*ListInstances200ResponseAllOfInstancesInnerLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Instance) SetLayout(v ListInstances200ResponseAllOfInstancesInnerLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *Instance) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPlan

`func (o *Instance) GetPlan() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Instance) GetPlanOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Instance) SetPlan(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Instance) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetName

`func (o *Instance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Instance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Instance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Instance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Instance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Instance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *Instance) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Instance) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Instance) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Instance) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetConfig

`func (o *Instance) GetConfig() ListInstances200ResponseAllOfInstancesInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Instance) GetConfigOk() (*ListInstances200ResponseAllOfInstancesInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Instance) SetConfig(v ListInstances200ResponseAllOfInstancesInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Instance) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConfigGroup

`func (o *Instance) GetConfigGroup() string`

GetConfigGroup returns the ConfigGroup field if non-nil, zero value otherwise.

### GetConfigGroupOk

`func (o *Instance) GetConfigGroupOk() (*string, bool)`

GetConfigGroupOk returns a tuple with the ConfigGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGroup

`func (o *Instance) SetConfigGroup(v string)`

SetConfigGroup sets ConfigGroup field to given value.

### HasConfigGroup

`func (o *Instance) HasConfigGroup() bool`

HasConfigGroup returns a boolean if a field has been set.

### GetConfigId

`func (o *Instance) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *Instance) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *Instance) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *Instance) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetConfigRole

`func (o *Instance) GetConfigRole() string`

GetConfigRole returns the ConfigRole field if non-nil, zero value otherwise.

### GetConfigRoleOk

`func (o *Instance) GetConfigRoleOk() (*string, bool)`

GetConfigRoleOk returns a tuple with the ConfigRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRole

`func (o *Instance) SetConfigRole(v string)`

SetConfigRole sets ConfigRole field to given value.

### HasConfigRole

`func (o *Instance) HasConfigRole() bool`

HasConfigRole returns a boolean if a field has been set.

### GetVolumes

`func (o *Instance) GetVolumes() []ListInstances200ResponseAllOfInstancesInnerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *Instance) GetVolumesOk() (*[]ListInstances200ResponseAllOfInstancesInnerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *Instance) SetVolumes(v []ListInstances200ResponseAllOfInstancesInnerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *Instance) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *Instance) GetControllers() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *Instance) GetControllersOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *Instance) SetControllers(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *Instance) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *Instance) GetInterfaces() []ListInstances200ResponseAllOfInstancesInnerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *Instance) GetInterfacesOk() (*[]ListInstances200ResponseAllOfInstancesInnerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *Instance) SetInterfaces(v []ListInstances200ResponseAllOfInstancesInnerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *Instance) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetCustomOptions

`func (o *Instance) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *Instance) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *Instance) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *Instance) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetInstanceVersion

`func (o *Instance) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *Instance) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *Instance) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *Instance) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetLabels

`func (o *Instance) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Instance) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Instance) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Instance) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *Instance) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Instance) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Instance) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Instance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEvars

`func (o *Instance) GetEvars() []map[string]interface{}`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *Instance) GetEvarsOk() (*[]map[string]interface{}, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *Instance) SetEvars(v []map[string]interface{})`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *Instance) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetMaxMemory

`func (o *Instance) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *Instance) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *Instance) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *Instance) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *Instance) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *Instance) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *Instance) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *Instance) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCores

`func (o *Instance) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *Instance) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *Instance) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *Instance) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *Instance) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *Instance) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *Instance) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *Instance) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetMaxCpu

`func (o *Instance) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *Instance) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *Instance) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *Instance) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetHourlyCost

`func (o *Instance) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *Instance) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *Instance) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *Instance) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *Instance) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *Instance) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *Instance) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *Instance) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetInstancePrice

`func (o *Instance) GetInstancePrice() ListInstances200ResponseAllOfInstancesInnerInstancePrice`

GetInstancePrice returns the InstancePrice field if non-nil, zero value otherwise.

### GetInstancePriceOk

`func (o *Instance) GetInstancePriceOk() (*ListInstances200ResponseAllOfInstancesInnerInstancePrice, bool)`

GetInstancePriceOk returns a tuple with the InstancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePrice

`func (o *Instance) SetInstancePrice(v ListInstances200ResponseAllOfInstancesInnerInstancePrice)`

SetInstancePrice sets InstancePrice field to given value.

### HasInstancePrice

`func (o *Instance) HasInstancePrice() bool`

HasInstancePrice returns a boolean if a field has been set.

### GetDateCreated

`func (o *Instance) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Instance) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Instance) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Instance) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Instance) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Instance) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Instance) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Instance) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHostName

`func (o *Instance) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *Instance) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *Instance) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *Instance) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetDomainName

`func (o *Instance) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Instance) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Instance) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *Instance) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *Instance) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *Instance) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *Instance) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *Instance) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetFirewallEnabled

`func (o *Instance) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *Instance) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *Instance) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *Instance) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetNetworkLevel

`func (o *Instance) GetNetworkLevel() string`

GetNetworkLevel returns the NetworkLevel field if non-nil, zero value otherwise.

### GetNetworkLevelOk

`func (o *Instance) GetNetworkLevelOk() (*string, bool)`

GetNetworkLevelOk returns a tuple with the NetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLevel

`func (o *Instance) SetNetworkLevel(v string)`

SetNetworkLevel sets NetworkLevel field to given value.

### HasNetworkLevel

`func (o *Instance) HasNetworkLevel() bool`

HasNetworkLevel returns a boolean if a field has been set.

### GetAutoScale

`func (o *Instance) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *Instance) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *Instance) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *Instance) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetInstanceContext

`func (o *Instance) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *Instance) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *Instance) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *Instance) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### GetCurrentDeployId

`func (o *Instance) GetCurrentDeployId() string`

GetCurrentDeployId returns the CurrentDeployId field if non-nil, zero value otherwise.

### GetCurrentDeployIdOk

`func (o *Instance) GetCurrentDeployIdOk() (*string, bool)`

GetCurrentDeployIdOk returns a tuple with the CurrentDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeployId

`func (o *Instance) SetCurrentDeployId(v string)`

SetCurrentDeployId sets CurrentDeployId field to given value.

### HasCurrentDeployId

`func (o *Instance) HasCurrentDeployId() bool`

HasCurrentDeployId returns a boolean if a field has been set.

### GetLocked

`func (o *Instance) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Instance) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Instance) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Instance) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetStatus

`func (o *Instance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Instance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Instance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Instance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Instance) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Instance) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Instance) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Instance) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Instance) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Instance) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Instance) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Instance) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *Instance) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *Instance) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *Instance) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *Instance) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *Instance) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *Instance) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *Instance) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *Instance) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *Instance) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *Instance) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *Instance) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *Instance) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetUserStatus

`func (o *Instance) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *Instance) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *Instance) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *Instance) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetExpireDays

`func (o *Instance) GetExpireDays() int64`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *Instance) GetExpireDaysOk() (*int64, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *Instance) SetExpireDays(v int64)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *Instance) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetRenewDays

`func (o *Instance) GetRenewDays() int64`

GetRenewDays returns the RenewDays field if non-nil, zero value otherwise.

### GetRenewDaysOk

`func (o *Instance) GetRenewDaysOk() (*int64, bool)`

GetRenewDaysOk returns a tuple with the RenewDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewDays

`func (o *Instance) SetRenewDays(v int64)`

SetRenewDays sets RenewDays field to given value.

### HasRenewDays

`func (o *Instance) HasRenewDays() bool`

HasRenewDays returns a boolean if a field has been set.

### GetExpireCount

`func (o *Instance) GetExpireCount() int64`

GetExpireCount returns the ExpireCount field if non-nil, zero value otherwise.

### GetExpireCountOk

`func (o *Instance) GetExpireCountOk() (*int64, bool)`

GetExpireCountOk returns a tuple with the ExpireCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireCount

`func (o *Instance) SetExpireCount(v int64)`

SetExpireCount sets ExpireCount field to given value.

### HasExpireCount

`func (o *Instance) HasExpireCount() bool`

HasExpireCount returns a boolean if a field has been set.

### GetExpireDate

`func (o *Instance) GetExpireDate() time.Time`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *Instance) GetExpireDateOk() (*time.Time, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *Instance) SetExpireDate(v time.Time)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *Instance) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetExpireWarningDate

`func (o *Instance) GetExpireWarningDate() time.Time`

GetExpireWarningDate returns the ExpireWarningDate field if non-nil, zero value otherwise.

### GetExpireWarningDateOk

`func (o *Instance) GetExpireWarningDateOk() (*time.Time, bool)`

GetExpireWarningDateOk returns a tuple with the ExpireWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWarningDate

`func (o *Instance) SetExpireWarningDate(v time.Time)`

SetExpireWarningDate sets ExpireWarningDate field to given value.

### HasExpireWarningDate

`func (o *Instance) HasExpireWarningDate() bool`

HasExpireWarningDate returns a boolean if a field has been set.

### GetExpireWarningSent

`func (o *Instance) GetExpireWarningSent() bool`

GetExpireWarningSent returns the ExpireWarningSent field if non-nil, zero value otherwise.

### GetExpireWarningSentOk

`func (o *Instance) GetExpireWarningSentOk() (*bool, bool)`

GetExpireWarningSentOk returns a tuple with the ExpireWarningSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWarningSent

`func (o *Instance) SetExpireWarningSent(v bool)`

SetExpireWarningSent sets ExpireWarningSent field to given value.

### HasExpireWarningSent

`func (o *Instance) HasExpireWarningSent() bool`

HasExpireWarningSent returns a boolean if a field has been set.

### GetShutdownDays

`func (o *Instance) GetShutdownDays() int64`

GetShutdownDays returns the ShutdownDays field if non-nil, zero value otherwise.

### GetShutdownDaysOk

`func (o *Instance) GetShutdownDaysOk() (*int64, bool)`

GetShutdownDaysOk returns a tuple with the ShutdownDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDays

`func (o *Instance) SetShutdownDays(v int64)`

SetShutdownDays sets ShutdownDays field to given value.

### HasShutdownDays

`func (o *Instance) HasShutdownDays() bool`

HasShutdownDays returns a boolean if a field has been set.

### GetShutdownRenewDays

`func (o *Instance) GetShutdownRenewDays() int64`

GetShutdownRenewDays returns the ShutdownRenewDays field if non-nil, zero value otherwise.

### GetShutdownRenewDaysOk

`func (o *Instance) GetShutdownRenewDaysOk() (*int64, bool)`

GetShutdownRenewDaysOk returns a tuple with the ShutdownRenewDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewDays

`func (o *Instance) SetShutdownRenewDays(v int64)`

SetShutdownRenewDays sets ShutdownRenewDays field to given value.

### HasShutdownRenewDays

`func (o *Instance) HasShutdownRenewDays() bool`

HasShutdownRenewDays returns a boolean if a field has been set.

### GetShutdownCount

`func (o *Instance) GetShutdownCount() int64`

GetShutdownCount returns the ShutdownCount field if non-nil, zero value otherwise.

### GetShutdownCountOk

`func (o *Instance) GetShutdownCountOk() (*int64, bool)`

GetShutdownCountOk returns a tuple with the ShutdownCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownCount

`func (o *Instance) SetShutdownCount(v int64)`

SetShutdownCount sets ShutdownCount field to given value.

### HasShutdownCount

`func (o *Instance) HasShutdownCount() bool`

HasShutdownCount returns a boolean if a field has been set.

### GetShutdownDate

`func (o *Instance) GetShutdownDate() time.Time`

GetShutdownDate returns the ShutdownDate field if non-nil, zero value otherwise.

### GetShutdownDateOk

`func (o *Instance) GetShutdownDateOk() (*time.Time, bool)`

GetShutdownDateOk returns a tuple with the ShutdownDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDate

`func (o *Instance) SetShutdownDate(v time.Time)`

SetShutdownDate sets ShutdownDate field to given value.

### HasShutdownDate

`func (o *Instance) HasShutdownDate() bool`

HasShutdownDate returns a boolean if a field has been set.

### GetShutdownWarningDate

`func (o *Instance) GetShutdownWarningDate() time.Time`

GetShutdownWarningDate returns the ShutdownWarningDate field if non-nil, zero value otherwise.

### GetShutdownWarningDateOk

`func (o *Instance) GetShutdownWarningDateOk() (*time.Time, bool)`

GetShutdownWarningDateOk returns a tuple with the ShutdownWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWarningDate

`func (o *Instance) SetShutdownWarningDate(v time.Time)`

SetShutdownWarningDate sets ShutdownWarningDate field to given value.

### HasShutdownWarningDate

`func (o *Instance) HasShutdownWarningDate() bool`

HasShutdownWarningDate returns a boolean if a field has been set.

### GetShutdownWarningSent

`func (o *Instance) GetShutdownWarningSent() bool`

GetShutdownWarningSent returns the ShutdownWarningSent field if non-nil, zero value otherwise.

### GetShutdownWarningSentOk

`func (o *Instance) GetShutdownWarningSentOk() (*bool, bool)`

GetShutdownWarningSentOk returns a tuple with the ShutdownWarningSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWarningSent

`func (o *Instance) SetShutdownWarningSent(v bool)`

SetShutdownWarningSent sets ShutdownWarningSent field to given value.

### HasShutdownWarningSent

`func (o *Instance) HasShutdownWarningSent() bool`

HasShutdownWarningSent returns a boolean if a field has been set.

### GetRemovalDate

`func (o *Instance) GetRemovalDate() time.Time`

GetRemovalDate returns the RemovalDate field if non-nil, zero value otherwise.

### GetRemovalDateOk

`func (o *Instance) GetRemovalDateOk() (*time.Time, bool)`

GetRemovalDateOk returns a tuple with the RemovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalDate

`func (o *Instance) SetRemovalDate(v time.Time)`

SetRemovalDate sets RemovalDate field to given value.

### HasRemovalDate

`func (o *Instance) HasRemovalDate() bool`

HasRemovalDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Instance) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Instance) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Instance) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Instance) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetOwner

`func (o *Instance) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Instance) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Instance) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Instance) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNotes

`func (o *Instance) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Instance) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Instance) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Instance) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetStats

`func (o *Instance) GetStats() ListInstances200ResponseAllOfInstancesInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Instance) GetStatsOk() (*ListInstances200ResponseAllOfInstancesInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Instance) SetStats(v ListInstances200ResponseAllOfInstancesInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Instance) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetPowerSchedule

`func (o *Instance) GetPowerSchedule() string`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *Instance) GetPowerScheduleOk() (*string, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *Instance) SetPowerSchedule(v string)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *Instance) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetIsScalable

`func (o *Instance) GetIsScalable() bool`

GetIsScalable returns the IsScalable field if non-nil, zero value otherwise.

### GetIsScalableOk

`func (o *Instance) GetIsScalableOk() (*bool, bool)`

GetIsScalableOk returns a tuple with the IsScalable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScalable

`func (o *Instance) SetIsScalable(v bool)`

SetIsScalable sets IsScalable field to given value.

### HasIsScalable

`func (o *Instance) HasIsScalable() bool`

HasIsScalable returns a boolean if a field has been set.

### GetInstanceThreshold

`func (o *Instance) GetInstanceThreshold() map[string]interface{}`

GetInstanceThreshold returns the InstanceThreshold field if non-nil, zero value otherwise.

### GetInstanceThresholdOk

`func (o *Instance) GetInstanceThresholdOk() (*map[string]interface{}, bool)`

GetInstanceThresholdOk returns a tuple with the InstanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceThreshold

`func (o *Instance) SetInstanceThreshold(v map[string]interface{})`

SetInstanceThreshold sets InstanceThreshold field to given value.

### HasInstanceThreshold

`func (o *Instance) HasInstanceThreshold() bool`

HasInstanceThreshold returns a boolean if a field has been set.

### GetIsBusy

`func (o *Instance) GetIsBusy() bool`

GetIsBusy returns the IsBusy field if non-nil, zero value otherwise.

### GetIsBusyOk

`func (o *Instance) GetIsBusyOk() (*bool, bool)`

GetIsBusyOk returns a tuple with the IsBusy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBusy

`func (o *Instance) SetIsBusy(v bool)`

SetIsBusy sets IsBusy field to given value.

### HasIsBusy

`func (o *Instance) HasIsBusy() bool`

HasIsBusy returns a boolean if a field has been set.

### GetApps

`func (o *Instance) GetApps() []map[string]interface{}`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *Instance) GetAppsOk() (*[]map[string]interface{}, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *Instance) SetApps(v []map[string]interface{})`

SetApps sets Apps field to given value.

### HasApps

`func (o *Instance) HasApps() bool`

HasApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



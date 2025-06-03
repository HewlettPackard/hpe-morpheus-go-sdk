# ListInstances200ResponseAllOfInstancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Tenant** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**InstanceType** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerInstanceType**](ListInstances200ResponseAllOfInstancesInnerInstanceType.md) |  | [optional] 
**Group** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Cloud** | Pointer to [**ListApps200ResponseAllOfAppsInnerBlueprint**](ListApps200ResponseAllOfAppsInnerBlueprint.md) |  | [optional] 
**Cluster** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerCluster**](ListInstances200ResponseAllOfInstancesInnerCluster.md) |  | [optional] 
**Containers** | Pointer to **[]int64** |  | [optional] 
**Servers** | Pointer to **[]int64** |  | [optional] 
**ConnectionInfo** | Pointer to [**[]ListInstances200ResponseAllOfInstancesInnerConnectionInfoInner**](ListInstances200ResponseAllOfInstancesInnerConnectionInfoInner.md) |  | [optional] 
**Layout** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerLayout**](ListInstances200ResponseAllOfInstancesInnerLayout.md) |  | [optional] 
**Plan** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
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
**Tags** | Pointer to [**[]ListInstances200ResponseAllOfInstancesInnerTagsInner**](ListInstances200ResponseAllOfInstancesInnerTagsInner.md) |  | [optional] 
**Evars** | Pointer to [**[]ListInstances200ResponseAllOfInstancesInnerEvarsInner**](ListInstances200ResponseAllOfInstancesInnerEvarsInner.md) |  | [optional] 
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
**Owner** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerCreatedBy**](GetAlerts200ResponseAllOfChecksInnerCreatedBy.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerStats**](ListInstances200ResponseAllOfInstancesInnerStats.md) |  | [optional] 
**PowerSchedule** | Pointer to **string** |  | [optional] 
**IsScalable** | Pointer to **bool** |  | [optional] 
**InstanceThreshold** | Pointer to **map[string]interface{}** |  | [optional] 
**IsBusy** | Pointer to **bool** |  | [optional] 
**Apps** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListInstances200ResponseAllOfInstancesInner

`func NewListInstances200ResponseAllOfInstancesInner() *ListInstances200ResponseAllOfInstancesInner`

NewListInstances200ResponseAllOfInstancesInner instantiates a new ListInstances200ResponseAllOfInstancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstances200ResponseAllOfInstancesInnerWithDefaults

`func NewListInstances200ResponseAllOfInstancesInnerWithDefaults() *ListInstances200ResponseAllOfInstancesInner`

NewListInstances200ResponseAllOfInstancesInnerWithDefaults instantiates a new ListInstances200ResponseAllOfInstancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInstances200ResponseAllOfInstancesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInstances200ResponseAllOfInstancesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListInstances200ResponseAllOfInstancesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ListInstances200ResponseAllOfInstancesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListInstances200ResponseAllOfInstancesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListInstances200ResponseAllOfInstancesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccountId

`func (o *ListInstances200ResponseAllOfInstancesInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListInstances200ResponseAllOfInstancesInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListInstances200ResponseAllOfInstancesInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetTenant

`func (o *ListInstances200ResponseAllOfInstancesInner) GetTenant() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetTenantOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ListInstances200ResponseAllOfInstancesInner) SetTenant(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ListInstances200ResponseAllOfInstancesInner) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetInstanceType

`func (o *ListInstances200ResponseAllOfInstancesInner) GetInstanceType() ListInstances200ResponseAllOfInstancesInnerInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetInstanceTypeOk() (*ListInstances200ResponseAllOfInstancesInnerInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ListInstances200ResponseAllOfInstancesInner) SetInstanceType(v ListInstances200ResponseAllOfInstancesInnerInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ListInstances200ResponseAllOfInstancesInner) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetGroup

`func (o *ListInstances200ResponseAllOfInstancesInner) GetGroup() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetGroupOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ListInstances200ResponseAllOfInstancesInner) SetGroup(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ListInstances200ResponseAllOfInstancesInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *ListInstances200ResponseAllOfInstancesInner) GetCloud() ListApps200ResponseAllOfAppsInnerBlueprint`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetCloudOk() (*ListApps200ResponseAllOfAppsInnerBlueprint, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ListInstances200ResponseAllOfInstancesInner) SetCloud(v ListApps200ResponseAllOfAppsInnerBlueprint)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ListInstances200ResponseAllOfInstancesInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCluster

`func (o *ListInstances200ResponseAllOfInstancesInner) GetCluster() ListInstances200ResponseAllOfInstancesInnerCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetClusterOk() (*ListInstances200ResponseAllOfInstancesInnerCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ListInstances200ResponseAllOfInstancesInner) SetCluster(v ListInstances200ResponseAllOfInstancesInnerCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ListInstances200ResponseAllOfInstancesInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetContainers

`func (o *ListInstances200ResponseAllOfInstancesInner) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ListInstances200ResponseAllOfInstancesInner) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *ListInstances200ResponseAllOfInstancesInner) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetServers

`func (o *ListInstances200ResponseAllOfInstancesInner) GetServers() []int64`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetServersOk() (*[]int64, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ListInstances200ResponseAllOfInstancesInner) SetServers(v []int64)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ListInstances200ResponseAllOfInstancesInner) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetConnectionInfo

`func (o *ListInstances200ResponseAllOfInstancesInner) GetConnectionInfo() []ListInstances200ResponseAllOfInstancesInnerConnectionInfoInner`

GetConnectionInfo returns the ConnectionInfo field if non-nil, zero value otherwise.

### GetConnectionInfoOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetConnectionInfoOk() (*[]ListInstances200ResponseAllOfInstancesInnerConnectionInfoInner, bool)`

GetConnectionInfoOk returns a tuple with the ConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionInfo

`func (o *ListInstances200ResponseAllOfInstancesInner) SetConnectionInfo(v []ListInstances200ResponseAllOfInstancesInnerConnectionInfoInner)`

SetConnectionInfo sets ConnectionInfo field to given value.

### HasConnectionInfo

`func (o *ListInstances200ResponseAllOfInstancesInner) HasConnectionInfo() bool`

HasConnectionInfo returns a boolean if a field has been set.

### GetLayout

`func (o *ListInstances200ResponseAllOfInstancesInner) GetLayout() ListInstances200ResponseAllOfInstancesInnerLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetLayoutOk() (*ListInstances200ResponseAllOfInstancesInnerLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ListInstances200ResponseAllOfInstancesInner) SetLayout(v ListInstances200ResponseAllOfInstancesInnerLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *ListInstances200ResponseAllOfInstancesInner) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPlan

`func (o *ListInstances200ResponseAllOfInstancesInner) GetPlan() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetPlanOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ListInstances200ResponseAllOfInstancesInner) SetPlan(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ListInstances200ResponseAllOfInstancesInner) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetName

`func (o *ListInstances200ResponseAllOfInstancesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListInstances200ResponseAllOfInstancesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListInstances200ResponseAllOfInstancesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListInstances200ResponseAllOfInstancesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListInstances200ResponseAllOfInstancesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListInstances200ResponseAllOfInstancesInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *ListInstances200ResponseAllOfInstancesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListInstances200ResponseAllOfInstancesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListInstances200ResponseAllOfInstancesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *ListInstances200ResponseAllOfInstancesInner) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ListInstances200ResponseAllOfInstancesInner) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ListInstances200ResponseAllOfInstancesInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetConfig

`func (o *ListInstances200ResponseAllOfInstancesInner) GetConfig() ListInstances200ResponseAllOfInstancesInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetConfigOk() (*ListInstances200ResponseAllOfInstancesInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListInstances200ResponseAllOfInstancesInner) SetConfig(v ListInstances200ResponseAllOfInstancesInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListInstances200ResponseAllOfInstancesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConfigGroup

`func (o *ListInstances200ResponseAllOfInstancesInner) GetConfigGroup() string`

GetConfigGroup returns the ConfigGroup field if non-nil, zero value otherwise.

### GetConfigGroupOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetConfigGroupOk() (*string, bool)`

GetConfigGroupOk returns a tuple with the ConfigGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGroup

`func (o *ListInstances200ResponseAllOfInstancesInner) SetConfigGroup(v string)`

SetConfigGroup sets ConfigGroup field to given value.

### HasConfigGroup

`func (o *ListInstances200ResponseAllOfInstancesInner) HasConfigGroup() bool`

HasConfigGroup returns a boolean if a field has been set.

### GetConfigId

`func (o *ListInstances200ResponseAllOfInstancesInner) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ListInstances200ResponseAllOfInstancesInner) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ListInstances200ResponseAllOfInstancesInner) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetConfigRole

`func (o *ListInstances200ResponseAllOfInstancesInner) GetConfigRole() string`

GetConfigRole returns the ConfigRole field if non-nil, zero value otherwise.

### GetConfigRoleOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetConfigRoleOk() (*string, bool)`

GetConfigRoleOk returns a tuple with the ConfigRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRole

`func (o *ListInstances200ResponseAllOfInstancesInner) SetConfigRole(v string)`

SetConfigRole sets ConfigRole field to given value.

### HasConfigRole

`func (o *ListInstances200ResponseAllOfInstancesInner) HasConfigRole() bool`

HasConfigRole returns a boolean if a field has been set.

### GetVolumes

`func (o *ListInstances200ResponseAllOfInstancesInner) GetVolumes() []ListInstances200ResponseAllOfInstancesInnerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetVolumesOk() (*[]ListInstances200ResponseAllOfInstancesInnerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ListInstances200ResponseAllOfInstancesInner) SetVolumes(v []ListInstances200ResponseAllOfInstancesInnerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ListInstances200ResponseAllOfInstancesInner) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *ListInstances200ResponseAllOfInstancesInner) GetControllers() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetControllersOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *ListInstances200ResponseAllOfInstancesInner) SetControllers(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *ListInstances200ResponseAllOfInstancesInner) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *ListInstances200ResponseAllOfInstancesInner) GetInterfaces() []ListInstances200ResponseAllOfInstancesInnerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetInterfacesOk() (*[]ListInstances200ResponseAllOfInstancesInnerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ListInstances200ResponseAllOfInstancesInner) SetInterfaces(v []ListInstances200ResponseAllOfInstancesInnerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ListInstances200ResponseAllOfInstancesInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetCustomOptions

`func (o *ListInstances200ResponseAllOfInstancesInner) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *ListInstances200ResponseAllOfInstancesInner) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *ListInstances200ResponseAllOfInstancesInner) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetInstanceVersion

`func (o *ListInstances200ResponseAllOfInstancesInner) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *ListInstances200ResponseAllOfInstancesInner) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *ListInstances200ResponseAllOfInstancesInner) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetLabels

`func (o *ListInstances200ResponseAllOfInstancesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListInstances200ResponseAllOfInstancesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListInstances200ResponseAllOfInstancesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *ListInstances200ResponseAllOfInstancesInner) GetTags() []ListInstances200ResponseAllOfInstancesInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetTagsOk() (*[]ListInstances200ResponseAllOfInstancesInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListInstances200ResponseAllOfInstancesInner) SetTags(v []ListInstances200ResponseAllOfInstancesInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListInstances200ResponseAllOfInstancesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEvars

`func (o *ListInstances200ResponseAllOfInstancesInner) GetEvars() []ListInstances200ResponseAllOfInstancesInnerEvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetEvarsOk() (*[]ListInstances200ResponseAllOfInstancesInnerEvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *ListInstances200ResponseAllOfInstancesInner) SetEvars(v []ListInstances200ResponseAllOfInstancesInnerEvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *ListInstances200ResponseAllOfInstancesInner) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListInstances200ResponseAllOfInstancesInner) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListInstances200ResponseAllOfInstancesInner) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListInstances200ResponseAllOfInstancesInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListInstances200ResponseAllOfInstancesInner) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListInstances200ResponseAllOfInstancesInner) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListInstances200ResponseAllOfInstancesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCores

`func (o *ListInstances200ResponseAllOfInstancesInner) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListInstances200ResponseAllOfInstancesInner) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ListInstances200ResponseAllOfInstancesInner) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *ListInstances200ResponseAllOfInstancesInner) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *ListInstances200ResponseAllOfInstancesInner) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *ListInstances200ResponseAllOfInstancesInner) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ListInstances200ResponseAllOfInstancesInner) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ListInstances200ResponseAllOfInstancesInner) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ListInstances200ResponseAllOfInstancesInner) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetHourlyCost

`func (o *ListInstances200ResponseAllOfInstancesInner) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *ListInstances200ResponseAllOfInstancesInner) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *ListInstances200ResponseAllOfInstancesInner) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *ListInstances200ResponseAllOfInstancesInner) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *ListInstances200ResponseAllOfInstancesInner) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *ListInstances200ResponseAllOfInstancesInner) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetInstancePrice

`func (o *ListInstances200ResponseAllOfInstancesInner) GetInstancePrice() ListInstances200ResponseAllOfInstancesInnerInstancePrice`

GetInstancePrice returns the InstancePrice field if non-nil, zero value otherwise.

### GetInstancePriceOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetInstancePriceOk() (*ListInstances200ResponseAllOfInstancesInnerInstancePrice, bool)`

GetInstancePriceOk returns a tuple with the InstancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePrice

`func (o *ListInstances200ResponseAllOfInstancesInner) SetInstancePrice(v ListInstances200ResponseAllOfInstancesInnerInstancePrice)`

SetInstancePrice sets InstancePrice field to given value.

### HasInstancePrice

`func (o *ListInstances200ResponseAllOfInstancesInner) HasInstancePrice() bool`

HasInstancePrice returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListInstances200ResponseAllOfInstancesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListInstances200ResponseAllOfInstancesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListInstances200ResponseAllOfInstancesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListInstances200ResponseAllOfInstancesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListInstances200ResponseAllOfInstancesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListInstances200ResponseAllOfInstancesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHostName

`func (o *ListInstances200ResponseAllOfInstancesInner) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ListInstances200ResponseAllOfInstancesInner) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ListInstances200ResponseAllOfInstancesInner) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetDomainName

`func (o *ListInstances200ResponseAllOfInstancesInner) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ListInstances200ResponseAllOfInstancesInner) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ListInstances200ResponseAllOfInstancesInner) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *ListInstances200ResponseAllOfInstancesInner) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *ListInstances200ResponseAllOfInstancesInner) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *ListInstances200ResponseAllOfInstancesInner) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetFirewallEnabled

`func (o *ListInstances200ResponseAllOfInstancesInner) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *ListInstances200ResponseAllOfInstancesInner) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *ListInstances200ResponseAllOfInstancesInner) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetNetworkLevel

`func (o *ListInstances200ResponseAllOfInstancesInner) GetNetworkLevel() string`

GetNetworkLevel returns the NetworkLevel field if non-nil, zero value otherwise.

### GetNetworkLevelOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetNetworkLevelOk() (*string, bool)`

GetNetworkLevelOk returns a tuple with the NetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLevel

`func (o *ListInstances200ResponseAllOfInstancesInner) SetNetworkLevel(v string)`

SetNetworkLevel sets NetworkLevel field to given value.

### HasNetworkLevel

`func (o *ListInstances200ResponseAllOfInstancesInner) HasNetworkLevel() bool`

HasNetworkLevel returns a boolean if a field has been set.

### GetAutoScale

`func (o *ListInstances200ResponseAllOfInstancesInner) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *ListInstances200ResponseAllOfInstancesInner) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *ListInstances200ResponseAllOfInstancesInner) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetInstanceContext

`func (o *ListInstances200ResponseAllOfInstancesInner) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *ListInstances200ResponseAllOfInstancesInner) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *ListInstances200ResponseAllOfInstancesInner) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### GetCurrentDeployId

`func (o *ListInstances200ResponseAllOfInstancesInner) GetCurrentDeployId() string`

GetCurrentDeployId returns the CurrentDeployId field if non-nil, zero value otherwise.

### GetCurrentDeployIdOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetCurrentDeployIdOk() (*string, bool)`

GetCurrentDeployIdOk returns a tuple with the CurrentDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeployId

`func (o *ListInstances200ResponseAllOfInstancesInner) SetCurrentDeployId(v string)`

SetCurrentDeployId sets CurrentDeployId field to given value.

### HasCurrentDeployId

`func (o *ListInstances200ResponseAllOfInstancesInner) HasCurrentDeployId() bool`

HasCurrentDeployId returns a boolean if a field has been set.

### GetLocked

`func (o *ListInstances200ResponseAllOfInstancesInner) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ListInstances200ResponseAllOfInstancesInner) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ListInstances200ResponseAllOfInstancesInner) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetStatus

`func (o *ListInstances200ResponseAllOfInstancesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListInstances200ResponseAllOfInstancesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListInstances200ResponseAllOfInstancesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListInstances200ResponseAllOfInstancesInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListInstances200ResponseAllOfInstancesInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListInstances200ResponseAllOfInstancesInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ListInstances200ResponseAllOfInstancesInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListInstances200ResponseAllOfInstancesInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListInstances200ResponseAllOfInstancesInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *ListInstances200ResponseAllOfInstancesInner) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListInstances200ResponseAllOfInstancesInner) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListInstances200ResponseAllOfInstancesInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *ListInstances200ResponseAllOfInstancesInner) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *ListInstances200ResponseAllOfInstancesInner) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *ListInstances200ResponseAllOfInstancesInner) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *ListInstances200ResponseAllOfInstancesInner) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ListInstances200ResponseAllOfInstancesInner) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ListInstances200ResponseAllOfInstancesInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetUserStatus

`func (o *ListInstances200ResponseAllOfInstancesInner) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *ListInstances200ResponseAllOfInstancesInner) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *ListInstances200ResponseAllOfInstancesInner) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetExpireDays

`func (o *ListInstances200ResponseAllOfInstancesInner) GetExpireDays() int64`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetExpireDaysOk() (*int64, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *ListInstances200ResponseAllOfInstancesInner) SetExpireDays(v int64)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *ListInstances200ResponseAllOfInstancesInner) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetRenewDays

`func (o *ListInstances200ResponseAllOfInstancesInner) GetRenewDays() int64`

GetRenewDays returns the RenewDays field if non-nil, zero value otherwise.

### GetRenewDaysOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetRenewDaysOk() (*int64, bool)`

GetRenewDaysOk returns a tuple with the RenewDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewDays

`func (o *ListInstances200ResponseAllOfInstancesInner) SetRenewDays(v int64)`

SetRenewDays sets RenewDays field to given value.

### HasRenewDays

`func (o *ListInstances200ResponseAllOfInstancesInner) HasRenewDays() bool`

HasRenewDays returns a boolean if a field has been set.

### GetExpireCount

`func (o *ListInstances200ResponseAllOfInstancesInner) GetExpireCount() int64`

GetExpireCount returns the ExpireCount field if non-nil, zero value otherwise.

### GetExpireCountOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetExpireCountOk() (*int64, bool)`

GetExpireCountOk returns a tuple with the ExpireCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireCount

`func (o *ListInstances200ResponseAllOfInstancesInner) SetExpireCount(v int64)`

SetExpireCount sets ExpireCount field to given value.

### HasExpireCount

`func (o *ListInstances200ResponseAllOfInstancesInner) HasExpireCount() bool`

HasExpireCount returns a boolean if a field has been set.

### GetExpireDate

`func (o *ListInstances200ResponseAllOfInstancesInner) GetExpireDate() time.Time`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetExpireDateOk() (*time.Time, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *ListInstances200ResponseAllOfInstancesInner) SetExpireDate(v time.Time)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *ListInstances200ResponseAllOfInstancesInner) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetExpireWarningDate

`func (o *ListInstances200ResponseAllOfInstancesInner) GetExpireWarningDate() time.Time`

GetExpireWarningDate returns the ExpireWarningDate field if non-nil, zero value otherwise.

### GetExpireWarningDateOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetExpireWarningDateOk() (*time.Time, bool)`

GetExpireWarningDateOk returns a tuple with the ExpireWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWarningDate

`func (o *ListInstances200ResponseAllOfInstancesInner) SetExpireWarningDate(v time.Time)`

SetExpireWarningDate sets ExpireWarningDate field to given value.

### HasExpireWarningDate

`func (o *ListInstances200ResponseAllOfInstancesInner) HasExpireWarningDate() bool`

HasExpireWarningDate returns a boolean if a field has been set.

### GetExpireWarningSent

`func (o *ListInstances200ResponseAllOfInstancesInner) GetExpireWarningSent() bool`

GetExpireWarningSent returns the ExpireWarningSent field if non-nil, zero value otherwise.

### GetExpireWarningSentOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetExpireWarningSentOk() (*bool, bool)`

GetExpireWarningSentOk returns a tuple with the ExpireWarningSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWarningSent

`func (o *ListInstances200ResponseAllOfInstancesInner) SetExpireWarningSent(v bool)`

SetExpireWarningSent sets ExpireWarningSent field to given value.

### HasExpireWarningSent

`func (o *ListInstances200ResponseAllOfInstancesInner) HasExpireWarningSent() bool`

HasExpireWarningSent returns a boolean if a field has been set.

### GetShutdownDays

`func (o *ListInstances200ResponseAllOfInstancesInner) GetShutdownDays() int64`

GetShutdownDays returns the ShutdownDays field if non-nil, zero value otherwise.

### GetShutdownDaysOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetShutdownDaysOk() (*int64, bool)`

GetShutdownDaysOk returns a tuple with the ShutdownDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDays

`func (o *ListInstances200ResponseAllOfInstancesInner) SetShutdownDays(v int64)`

SetShutdownDays sets ShutdownDays field to given value.

### HasShutdownDays

`func (o *ListInstances200ResponseAllOfInstancesInner) HasShutdownDays() bool`

HasShutdownDays returns a boolean if a field has been set.

### GetShutdownRenewDays

`func (o *ListInstances200ResponseAllOfInstancesInner) GetShutdownRenewDays() int64`

GetShutdownRenewDays returns the ShutdownRenewDays field if non-nil, zero value otherwise.

### GetShutdownRenewDaysOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetShutdownRenewDaysOk() (*int64, bool)`

GetShutdownRenewDaysOk returns a tuple with the ShutdownRenewDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewDays

`func (o *ListInstances200ResponseAllOfInstancesInner) SetShutdownRenewDays(v int64)`

SetShutdownRenewDays sets ShutdownRenewDays field to given value.

### HasShutdownRenewDays

`func (o *ListInstances200ResponseAllOfInstancesInner) HasShutdownRenewDays() bool`

HasShutdownRenewDays returns a boolean if a field has been set.

### GetShutdownCount

`func (o *ListInstances200ResponseAllOfInstancesInner) GetShutdownCount() int64`

GetShutdownCount returns the ShutdownCount field if non-nil, zero value otherwise.

### GetShutdownCountOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetShutdownCountOk() (*int64, bool)`

GetShutdownCountOk returns a tuple with the ShutdownCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownCount

`func (o *ListInstances200ResponseAllOfInstancesInner) SetShutdownCount(v int64)`

SetShutdownCount sets ShutdownCount field to given value.

### HasShutdownCount

`func (o *ListInstances200ResponseAllOfInstancesInner) HasShutdownCount() bool`

HasShutdownCount returns a boolean if a field has been set.

### GetShutdownDate

`func (o *ListInstances200ResponseAllOfInstancesInner) GetShutdownDate() time.Time`

GetShutdownDate returns the ShutdownDate field if non-nil, zero value otherwise.

### GetShutdownDateOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetShutdownDateOk() (*time.Time, bool)`

GetShutdownDateOk returns a tuple with the ShutdownDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDate

`func (o *ListInstances200ResponseAllOfInstancesInner) SetShutdownDate(v time.Time)`

SetShutdownDate sets ShutdownDate field to given value.

### HasShutdownDate

`func (o *ListInstances200ResponseAllOfInstancesInner) HasShutdownDate() bool`

HasShutdownDate returns a boolean if a field has been set.

### GetShutdownWarningDate

`func (o *ListInstances200ResponseAllOfInstancesInner) GetShutdownWarningDate() time.Time`

GetShutdownWarningDate returns the ShutdownWarningDate field if non-nil, zero value otherwise.

### GetShutdownWarningDateOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetShutdownWarningDateOk() (*time.Time, bool)`

GetShutdownWarningDateOk returns a tuple with the ShutdownWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWarningDate

`func (o *ListInstances200ResponseAllOfInstancesInner) SetShutdownWarningDate(v time.Time)`

SetShutdownWarningDate sets ShutdownWarningDate field to given value.

### HasShutdownWarningDate

`func (o *ListInstances200ResponseAllOfInstancesInner) HasShutdownWarningDate() bool`

HasShutdownWarningDate returns a boolean if a field has been set.

### GetShutdownWarningSent

`func (o *ListInstances200ResponseAllOfInstancesInner) GetShutdownWarningSent() bool`

GetShutdownWarningSent returns the ShutdownWarningSent field if non-nil, zero value otherwise.

### GetShutdownWarningSentOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetShutdownWarningSentOk() (*bool, bool)`

GetShutdownWarningSentOk returns a tuple with the ShutdownWarningSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWarningSent

`func (o *ListInstances200ResponseAllOfInstancesInner) SetShutdownWarningSent(v bool)`

SetShutdownWarningSent sets ShutdownWarningSent field to given value.

### HasShutdownWarningSent

`func (o *ListInstances200ResponseAllOfInstancesInner) HasShutdownWarningSent() bool`

HasShutdownWarningSent returns a boolean if a field has been set.

### GetRemovalDate

`func (o *ListInstances200ResponseAllOfInstancesInner) GetRemovalDate() time.Time`

GetRemovalDate returns the RemovalDate field if non-nil, zero value otherwise.

### GetRemovalDateOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetRemovalDateOk() (*time.Time, bool)`

GetRemovalDateOk returns a tuple with the RemovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalDate

`func (o *ListInstances200ResponseAllOfInstancesInner) SetRemovalDate(v time.Time)`

SetRemovalDate sets RemovalDate field to given value.

### HasRemovalDate

`func (o *ListInstances200ResponseAllOfInstancesInner) HasRemovalDate() bool`

HasRemovalDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListInstances200ResponseAllOfInstancesInner) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListInstances200ResponseAllOfInstancesInner) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListInstances200ResponseAllOfInstancesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetOwner

`func (o *ListInstances200ResponseAllOfInstancesInner) GetOwner() GetAlerts200ResponseAllOfChecksInnerCreatedBy`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerCreatedBy, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListInstances200ResponseAllOfInstancesInner) SetOwner(v GetAlerts200ResponseAllOfChecksInnerCreatedBy)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListInstances200ResponseAllOfInstancesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNotes

`func (o *ListInstances200ResponseAllOfInstancesInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ListInstances200ResponseAllOfInstancesInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ListInstances200ResponseAllOfInstancesInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetStats

`func (o *ListInstances200ResponseAllOfInstancesInner) GetStats() ListInstances200ResponseAllOfInstancesInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetStatsOk() (*ListInstances200ResponseAllOfInstancesInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListInstances200ResponseAllOfInstancesInner) SetStats(v ListInstances200ResponseAllOfInstancesInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListInstances200ResponseAllOfInstancesInner) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetPowerSchedule

`func (o *ListInstances200ResponseAllOfInstancesInner) GetPowerSchedule() string`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetPowerScheduleOk() (*string, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *ListInstances200ResponseAllOfInstancesInner) SetPowerSchedule(v string)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *ListInstances200ResponseAllOfInstancesInner) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetIsScalable

`func (o *ListInstances200ResponseAllOfInstancesInner) GetIsScalable() bool`

GetIsScalable returns the IsScalable field if non-nil, zero value otherwise.

### GetIsScalableOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetIsScalableOk() (*bool, bool)`

GetIsScalableOk returns a tuple with the IsScalable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScalable

`func (o *ListInstances200ResponseAllOfInstancesInner) SetIsScalable(v bool)`

SetIsScalable sets IsScalable field to given value.

### HasIsScalable

`func (o *ListInstances200ResponseAllOfInstancesInner) HasIsScalable() bool`

HasIsScalable returns a boolean if a field has been set.

### GetInstanceThreshold

`func (o *ListInstances200ResponseAllOfInstancesInner) GetInstanceThreshold() map[string]interface{}`

GetInstanceThreshold returns the InstanceThreshold field if non-nil, zero value otherwise.

### GetInstanceThresholdOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetInstanceThresholdOk() (*map[string]interface{}, bool)`

GetInstanceThresholdOk returns a tuple with the InstanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceThreshold

`func (o *ListInstances200ResponseAllOfInstancesInner) SetInstanceThreshold(v map[string]interface{})`

SetInstanceThreshold sets InstanceThreshold field to given value.

### HasInstanceThreshold

`func (o *ListInstances200ResponseAllOfInstancesInner) HasInstanceThreshold() bool`

HasInstanceThreshold returns a boolean if a field has been set.

### GetIsBusy

`func (o *ListInstances200ResponseAllOfInstancesInner) GetIsBusy() bool`

GetIsBusy returns the IsBusy field if non-nil, zero value otherwise.

### GetIsBusyOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetIsBusyOk() (*bool, bool)`

GetIsBusyOk returns a tuple with the IsBusy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBusy

`func (o *ListInstances200ResponseAllOfInstancesInner) SetIsBusy(v bool)`

SetIsBusy sets IsBusy field to given value.

### HasIsBusy

`func (o *ListInstances200ResponseAllOfInstancesInner) HasIsBusy() bool`

HasIsBusy returns a boolean if a field has been set.

### GetApps

`func (o *ListInstances200ResponseAllOfInstancesInner) GetApps() []map[string]interface{}`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ListInstances200ResponseAllOfInstancesInner) GetAppsOk() (*[]map[string]interface{}, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ListInstances200ResponseAllOfInstancesInner) SetApps(v []map[string]interface{})`

SetApps sets Apps field to given value.

### HasApps

`func (o *ListInstances200ResponseAllOfInstancesInner) HasApps() bool`

HasApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AddInstance200ResponseAllOfOneOfInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Tenant** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**InstanceType** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceInstanceType**](AddInstance200ResponseAllOfOneOfInstanceInstanceType.md) |  | [optional] 
**Group** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Cloud** | Pointer to [**ListApps200ResponseAllOfAppsInnerBlueprint**](ListApps200ResponseAllOfAppsInnerBlueprint.md) |  | [optional] 
**Cluster** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerCluster**](ListInstances200ResponseAllOfInstancesInnerCluster.md) |  | [optional] 
**Containers** | Pointer to **[]int64** |  | [optional] 
**Servers** | Pointer to **[]int64** |  | [optional] 
**ConnectionInfo** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner**](AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner.md) |  | [optional] 
**Layout** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceLayout**](AddInstance200ResponseAllOfOneOfInstanceLayout.md) |  | [optional] 
**Plan** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceConfig**](AddInstance200ResponseAllOfOneOfInstanceConfig.md) |  | [optional] 
**ConfigGroup** | Pointer to **string** |  | [optional] 
**ConfigId** | Pointer to **string** |  | [optional] 
**ConfigRole** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceVolumesInner**](AddInstance200ResponseAllOfOneOfInstanceVolumesInner.md) |  | [optional] 
**Controllers** | Pointer to [**[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner.md) |  | [optional] 
**Interfaces** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceInterfacesInner**](AddInstance200ResponseAllOfOneOfInstanceInterfacesInner.md) |  | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**InstanceVersion** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceTagsInner**](AddInstance200ResponseAllOfOneOfInstanceTagsInner.md) |  | [optional] 
**Evars** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceEvarsInner**](AddInstance200ResponseAllOfOneOfInstanceEvarsInner.md) |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**CoresPerSocket** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **string** |  | [optional] 
**HourlyCost** | Pointer to **float32** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**InstancePrice** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceInstancePrice**](AddInstance200ResponseAllOfOneOfInstanceInstancePrice.md) |  | [optional] 
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
**Stats** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceStats**](AddInstance200ResponseAllOfOneOfInstanceStats.md) |  | [optional] 
**PowerSchedule** | Pointer to **string** |  | [optional] 
**IsScalable** | Pointer to **bool** |  | [optional] 
**InstanceThreshold** | Pointer to **map[string]interface{}** |  | [optional] 
**IsBusy** | Pointer to **bool** |  | [optional] 
**Apps** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewAddInstance200ResponseAllOfOneOfInstance

`func NewAddInstance200ResponseAllOfOneOfInstance() *AddInstance200ResponseAllOfOneOfInstance`

NewAddInstance200ResponseAllOfOneOfInstance instantiates a new AddInstance200ResponseAllOfOneOfInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstance200ResponseAllOfOneOfInstanceWithDefaults

`func NewAddInstance200ResponseAllOfOneOfInstanceWithDefaults() *AddInstance200ResponseAllOfOneOfInstance`

NewAddInstance200ResponseAllOfOneOfInstanceWithDefaults instantiates a new AddInstance200ResponseAllOfOneOfInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccountId

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetTenant

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetTenant() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetTenantOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetTenant(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetInstanceType

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetInstanceType() AddInstance200ResponseAllOfOneOfInstanceInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetInstanceTypeOk() (*AddInstance200ResponseAllOfOneOfInstanceInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetInstanceType(v AddInstance200ResponseAllOfOneOfInstanceInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetGroup

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetGroup() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetGroupOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetGroup(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetCloud() ListApps200ResponseAllOfAppsInnerBlueprint`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetCloudOk() (*ListApps200ResponseAllOfAppsInnerBlueprint, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetCloud(v ListApps200ResponseAllOfAppsInnerBlueprint)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCluster

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetCluster() ListInstances200ResponseAllOfInstancesInnerCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetClusterOk() (*ListInstances200ResponseAllOfInstancesInnerCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetCluster(v ListInstances200ResponseAllOfInstancesInnerCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetContainers

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetServers

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetServers() []int64`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetServersOk() (*[]int64, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetServers(v []int64)`

SetServers sets Servers field to given value.

### HasServers

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetConnectionInfo

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetConnectionInfo() []AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner`

GetConnectionInfo returns the ConnectionInfo field if non-nil, zero value otherwise.

### GetConnectionInfoOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetConnectionInfoOk() (*[]AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner, bool)`

GetConnectionInfoOk returns a tuple with the ConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionInfo

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetConnectionInfo(v []AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner)`

SetConnectionInfo sets ConnectionInfo field to given value.

### HasConnectionInfo

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasConnectionInfo() bool`

HasConnectionInfo returns a boolean if a field has been set.

### GetLayout

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetLayout() AddInstance200ResponseAllOfOneOfInstanceLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetLayoutOk() (*AddInstance200ResponseAllOfOneOfInstanceLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetLayout(v AddInstance200ResponseAllOfOneOfInstanceLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPlan

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetPlan() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetPlanOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetPlan(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetName

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetConfig

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetConfig() AddInstance200ResponseAllOfOneOfInstanceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetConfigOk() (*AddInstance200ResponseAllOfOneOfInstanceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetConfig(v AddInstance200ResponseAllOfOneOfInstanceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConfigGroup

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetConfigGroup() string`

GetConfigGroup returns the ConfigGroup field if non-nil, zero value otherwise.

### GetConfigGroupOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetConfigGroupOk() (*string, bool)`

GetConfigGroupOk returns a tuple with the ConfigGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGroup

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetConfigGroup(v string)`

SetConfigGroup sets ConfigGroup field to given value.

### HasConfigGroup

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasConfigGroup() bool`

HasConfigGroup returns a boolean if a field has been set.

### GetConfigId

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetConfigRole

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetConfigRole() string`

GetConfigRole returns the ConfigRole field if non-nil, zero value otherwise.

### GetConfigRoleOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetConfigRoleOk() (*string, bool)`

GetConfigRoleOk returns a tuple with the ConfigRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRole

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetConfigRole(v string)`

SetConfigRole sets ConfigRole field to given value.

### HasConfigRole

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasConfigRole() bool`

HasConfigRole returns a boolean if a field has been set.

### GetVolumes

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetVolumes() []AddInstance200ResponseAllOfOneOfInstanceVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetVolumesOk() (*[]AddInstance200ResponseAllOfOneOfInstanceVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetVolumes(v []AddInstance200ResponseAllOfOneOfInstanceVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetControllers() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetControllersOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetControllers(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResourceControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetInterfaces() []AddInstance200ResponseAllOfOneOfInstanceInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetInterfacesOk() (*[]AddInstance200ResponseAllOfOneOfInstanceInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetInterfaces(v []AddInstance200ResponseAllOfOneOfInstanceInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetCustomOptions

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetInstanceVersion

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetLabels

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetTags() []AddInstance200ResponseAllOfOneOfInstanceTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetTagsOk() (*[]AddInstance200ResponseAllOfOneOfInstanceTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetTags(v []AddInstance200ResponseAllOfOneOfInstanceTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEvars

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetEvars() []AddInstance200ResponseAllOfOneOfInstanceEvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetEvarsOk() (*[]AddInstance200ResponseAllOfOneOfInstanceEvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetEvars(v []AddInstance200ResponseAllOfOneOfInstanceEvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetMaxMemory

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCores

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetMaxCpu

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetHourlyCost

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetInstancePrice

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetInstancePrice() AddInstance200ResponseAllOfOneOfInstanceInstancePrice`

GetInstancePrice returns the InstancePrice field if non-nil, zero value otherwise.

### GetInstancePriceOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetInstancePriceOk() (*AddInstance200ResponseAllOfOneOfInstanceInstancePrice, bool)`

GetInstancePriceOk returns a tuple with the InstancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePrice

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetInstancePrice(v AddInstance200ResponseAllOfOneOfInstanceInstancePrice)`

SetInstancePrice sets InstancePrice field to given value.

### HasInstancePrice

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasInstancePrice() bool`

HasInstancePrice returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHostName

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetDomainName

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetFirewallEnabled

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetNetworkLevel

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetNetworkLevel() string`

GetNetworkLevel returns the NetworkLevel field if non-nil, zero value otherwise.

### GetNetworkLevelOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetNetworkLevelOk() (*string, bool)`

GetNetworkLevelOk returns a tuple with the NetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLevel

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetNetworkLevel(v string)`

SetNetworkLevel sets NetworkLevel field to given value.

### HasNetworkLevel

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasNetworkLevel() bool`

HasNetworkLevel returns a boolean if a field has been set.

### GetAutoScale

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetInstanceContext

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### GetCurrentDeployId

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetCurrentDeployId() string`

GetCurrentDeployId returns the CurrentDeployId field if non-nil, zero value otherwise.

### GetCurrentDeployIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetCurrentDeployIdOk() (*string, bool)`

GetCurrentDeployIdOk returns a tuple with the CurrentDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeployId

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetCurrentDeployId(v string)`

SetCurrentDeployId sets CurrentDeployId field to given value.

### HasCurrentDeployId

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasCurrentDeployId() bool`

HasCurrentDeployId returns a boolean if a field has been set.

### GetLocked

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetStatus

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetErrorMessage

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetUserStatus

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetExpireDays

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetExpireDays() int64`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetExpireDaysOk() (*int64, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetExpireDays(v int64)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetRenewDays

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetRenewDays() int64`

GetRenewDays returns the RenewDays field if non-nil, zero value otherwise.

### GetRenewDaysOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetRenewDaysOk() (*int64, bool)`

GetRenewDaysOk returns a tuple with the RenewDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewDays

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetRenewDays(v int64)`

SetRenewDays sets RenewDays field to given value.

### HasRenewDays

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasRenewDays() bool`

HasRenewDays returns a boolean if a field has been set.

### GetExpireCount

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetExpireCount() int64`

GetExpireCount returns the ExpireCount field if non-nil, zero value otherwise.

### GetExpireCountOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetExpireCountOk() (*int64, bool)`

GetExpireCountOk returns a tuple with the ExpireCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireCount

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetExpireCount(v int64)`

SetExpireCount sets ExpireCount field to given value.

### HasExpireCount

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasExpireCount() bool`

HasExpireCount returns a boolean if a field has been set.

### GetExpireDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetExpireDate() time.Time`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetExpireDateOk() (*time.Time, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetExpireDate(v time.Time)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetExpireWarningDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetExpireWarningDate() time.Time`

GetExpireWarningDate returns the ExpireWarningDate field if non-nil, zero value otherwise.

### GetExpireWarningDateOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetExpireWarningDateOk() (*time.Time, bool)`

GetExpireWarningDateOk returns a tuple with the ExpireWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWarningDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetExpireWarningDate(v time.Time)`

SetExpireWarningDate sets ExpireWarningDate field to given value.

### HasExpireWarningDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasExpireWarningDate() bool`

HasExpireWarningDate returns a boolean if a field has been set.

### GetExpireWarningSent

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetExpireWarningSent() bool`

GetExpireWarningSent returns the ExpireWarningSent field if non-nil, zero value otherwise.

### GetExpireWarningSentOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetExpireWarningSentOk() (*bool, bool)`

GetExpireWarningSentOk returns a tuple with the ExpireWarningSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWarningSent

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetExpireWarningSent(v bool)`

SetExpireWarningSent sets ExpireWarningSent field to given value.

### HasExpireWarningSent

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasExpireWarningSent() bool`

HasExpireWarningSent returns a boolean if a field has been set.

### GetShutdownDays

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetShutdownDays() int64`

GetShutdownDays returns the ShutdownDays field if non-nil, zero value otherwise.

### GetShutdownDaysOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetShutdownDaysOk() (*int64, bool)`

GetShutdownDaysOk returns a tuple with the ShutdownDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDays

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetShutdownDays(v int64)`

SetShutdownDays sets ShutdownDays field to given value.

### HasShutdownDays

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasShutdownDays() bool`

HasShutdownDays returns a boolean if a field has been set.

### GetShutdownRenewDays

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetShutdownRenewDays() int64`

GetShutdownRenewDays returns the ShutdownRenewDays field if non-nil, zero value otherwise.

### GetShutdownRenewDaysOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetShutdownRenewDaysOk() (*int64, bool)`

GetShutdownRenewDaysOk returns a tuple with the ShutdownRenewDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewDays

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetShutdownRenewDays(v int64)`

SetShutdownRenewDays sets ShutdownRenewDays field to given value.

### HasShutdownRenewDays

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasShutdownRenewDays() bool`

HasShutdownRenewDays returns a boolean if a field has been set.

### GetShutdownCount

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetShutdownCount() int64`

GetShutdownCount returns the ShutdownCount field if non-nil, zero value otherwise.

### GetShutdownCountOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetShutdownCountOk() (*int64, bool)`

GetShutdownCountOk returns a tuple with the ShutdownCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownCount

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetShutdownCount(v int64)`

SetShutdownCount sets ShutdownCount field to given value.

### HasShutdownCount

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasShutdownCount() bool`

HasShutdownCount returns a boolean if a field has been set.

### GetShutdownDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetShutdownDate() time.Time`

GetShutdownDate returns the ShutdownDate field if non-nil, zero value otherwise.

### GetShutdownDateOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetShutdownDateOk() (*time.Time, bool)`

GetShutdownDateOk returns a tuple with the ShutdownDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetShutdownDate(v time.Time)`

SetShutdownDate sets ShutdownDate field to given value.

### HasShutdownDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasShutdownDate() bool`

HasShutdownDate returns a boolean if a field has been set.

### GetShutdownWarningDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetShutdownWarningDate() time.Time`

GetShutdownWarningDate returns the ShutdownWarningDate field if non-nil, zero value otherwise.

### GetShutdownWarningDateOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetShutdownWarningDateOk() (*time.Time, bool)`

GetShutdownWarningDateOk returns a tuple with the ShutdownWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWarningDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetShutdownWarningDate(v time.Time)`

SetShutdownWarningDate sets ShutdownWarningDate field to given value.

### HasShutdownWarningDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasShutdownWarningDate() bool`

HasShutdownWarningDate returns a boolean if a field has been set.

### GetShutdownWarningSent

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetShutdownWarningSent() bool`

GetShutdownWarningSent returns the ShutdownWarningSent field if non-nil, zero value otherwise.

### GetShutdownWarningSentOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetShutdownWarningSentOk() (*bool, bool)`

GetShutdownWarningSentOk returns a tuple with the ShutdownWarningSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWarningSent

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetShutdownWarningSent(v bool)`

SetShutdownWarningSent sets ShutdownWarningSent field to given value.

### HasShutdownWarningSent

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasShutdownWarningSent() bool`

HasShutdownWarningSent returns a boolean if a field has been set.

### GetRemovalDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetRemovalDate() time.Time`

GetRemovalDate returns the RemovalDate field if non-nil, zero value otherwise.

### GetRemovalDateOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetRemovalDateOk() (*time.Time, bool)`

GetRemovalDateOk returns a tuple with the RemovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetRemovalDate(v time.Time)`

SetRemovalDate sets RemovalDate field to given value.

### HasRemovalDate

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasRemovalDate() bool`

HasRemovalDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetOwner

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetOwner() GetAlerts200ResponseAllOfChecksInnerCreatedBy`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerCreatedBy, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetOwner(v GetAlerts200ResponseAllOfChecksInnerCreatedBy)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNotes

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetStats

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetStats() AddInstance200ResponseAllOfOneOfInstanceStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetStatsOk() (*AddInstance200ResponseAllOfOneOfInstanceStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetStats(v AddInstance200ResponseAllOfOneOfInstanceStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetPowerSchedule

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetPowerSchedule() string`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetPowerScheduleOk() (*string, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetPowerSchedule(v string)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetIsScalable

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetIsScalable() bool`

GetIsScalable returns the IsScalable field if non-nil, zero value otherwise.

### GetIsScalableOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetIsScalableOk() (*bool, bool)`

GetIsScalableOk returns a tuple with the IsScalable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScalable

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetIsScalable(v bool)`

SetIsScalable sets IsScalable field to given value.

### HasIsScalable

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasIsScalable() bool`

HasIsScalable returns a boolean if a field has been set.

### GetInstanceThreshold

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetInstanceThreshold() map[string]interface{}`

GetInstanceThreshold returns the InstanceThreshold field if non-nil, zero value otherwise.

### GetInstanceThresholdOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetInstanceThresholdOk() (*map[string]interface{}, bool)`

GetInstanceThresholdOk returns a tuple with the InstanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceThreshold

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetInstanceThreshold(v map[string]interface{})`

SetInstanceThreshold sets InstanceThreshold field to given value.

### HasInstanceThreshold

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasInstanceThreshold() bool`

HasInstanceThreshold returns a boolean if a field has been set.

### GetIsBusy

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetIsBusy() bool`

GetIsBusy returns the IsBusy field if non-nil, zero value otherwise.

### GetIsBusyOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetIsBusyOk() (*bool, bool)`

GetIsBusyOk returns a tuple with the IsBusy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBusy

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetIsBusy(v bool)`

SetIsBusy sets IsBusy field to given value.

### HasIsBusy

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasIsBusy() bool`

HasIsBusy returns a boolean if a field has been set.

### GetApps

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetApps() []map[string]interface{}`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AddInstance200ResponseAllOfOneOfInstance) GetAppsOk() (*[]map[string]interface{}, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AddInstance200ResponseAllOfOneOfInstance) SetApps(v []map[string]interface{})`

SetApps sets Apps field to given value.

### HasApps

`func (o *AddInstance200ResponseAllOfOneOfInstance) HasApps() bool`

HasApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



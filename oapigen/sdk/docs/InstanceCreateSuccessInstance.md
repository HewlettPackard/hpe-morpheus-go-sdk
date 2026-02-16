# InstanceCreateSuccessInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Tenant** | Pointer to [**InstanceCreateSuccessInstanceTenant**](InstanceCreateSuccessInstanceTenant.md) |  | [optional] 
**InstanceType** | Pointer to [**InstanceCreateSuccessInstanceInstanceType**](InstanceCreateSuccessInstanceInstanceType.md) |  | [optional] 
**Group** | Pointer to [**InstanceCreateSuccessInstanceGroup**](InstanceCreateSuccessInstanceGroup.md) |  | [optional] 
**Cloud** | Pointer to [**InstanceCreateSuccessInstanceCloud**](InstanceCreateSuccessInstanceCloud.md) |  | [optional] 
**Cluster** | Pointer to [**InstanceCreateSuccessInstanceCluster**](InstanceCreateSuccessInstanceCluster.md) |  | [optional] 
**Containers** | Pointer to **[]int64** |  | [optional] 
**Servers** | Pointer to **[]int64** |  | [optional] 
**ConnectionInfo** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner**](AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner.md) |  | [optional] 
**Layout** | Pointer to [**InstanceCreateSuccessInstanceLayout**](InstanceCreateSuccessInstanceLayout.md) |  | [optional] 
**Plan** | Pointer to [**InstanceCreateSuccessInstancePlan**](InstanceCreateSuccessInstancePlan.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Environment** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to [**InstanceCreateSuccessInstanceConfig**](InstanceCreateSuccessInstanceConfig.md) |  | [optional] 
**ConfigGroup** | Pointer to **NullableString** |  | [optional] 
**ConfigId** | Pointer to **NullableString** |  | [optional] 
**ConfigRole** | Pointer to **NullableString** |  | [optional] 
**Volumes** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceVolumesInner**](AddInstance200ResponseAllOfOneOfInstanceVolumesInner.md) |  | [optional] 
**Controllers** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceControllersInner**](AddInstance200ResponseAllOfOneOfInstanceControllersInner.md) |  | [optional] 
**Interfaces** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceInterfacesInner**](AddInstance200ResponseAllOfOneOfInstanceInterfacesInner.md) |  | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**InstanceVersion** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceTagsInner**](AddInstance200ResponseAllOfOneOfInstanceTagsInner.md) |  | [optional] 
**Evars** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceEvarsInner**](AddInstance200ResponseAllOfOneOfInstanceEvarsInner.md) |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**CoresPerSocket** | Pointer to **NullableInt64** |  | [optional] 
**MaxCpu** | Pointer to **NullableInt64** |  | [optional] 
**HourlyCost** | Pointer to **float32** |  | [optional] 
**HourlyPrice** | Pointer to **float32** |  | [optional] 
**InstancePrice** | Pointer to [**InstanceCreateSuccessInstanceInstancePrice**](InstanceCreateSuccessInstanceInstancePrice.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **NullableString** |  | [optional] 
**EnvironmentPrefix** | Pointer to **NullableString** |  | [optional] 
**FirewallEnabled** | Pointer to **bool** |  | [optional] 
**NetworkLevel** | Pointer to **string** |  | [optional] 
**AutoScale** | Pointer to **bool** |  | [optional] 
**InstanceContext** | Pointer to **NullableString** |  | [optional] 
**CurrentDeployId** | Pointer to **NullableString** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusPercent** | Pointer to **NullableString** |  | [optional] 
**StatusEta** | Pointer to **NullableString** |  | [optional] 
**UserStatus** | Pointer to **NullableString** |  | [optional] 
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
**RemovalDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedBy** | Pointer to [**InstanceCreateSuccessInstanceCreatedBy**](InstanceCreateSuccessInstanceCreatedBy.md) |  | [optional] 
**Owner** | Pointer to [**InstanceCreateSuccessInstanceOwner**](InstanceCreateSuccessInstanceOwner.md) |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Stats** | Pointer to [**InstanceCreateSuccessInstanceStats**](InstanceCreateSuccessInstanceStats.md) |  | [optional] 
**PowerSchedule** | Pointer to **NullableString** |  | [optional] 
**IsScalable** | Pointer to **bool** |  | [optional] 
**InstanceThreshold** | Pointer to **map[string]interface{}** |  | [optional] 
**IsBusy** | Pointer to **bool** |  | [optional] 
**Apps** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerDetails** | Pointer to [**[]InstanceContainer4**](InstanceContainer4.md) |  | [optional] 

## Methods

### NewInstanceCreateSuccessInstance

`func NewInstanceCreateSuccessInstance() *InstanceCreateSuccessInstance`

NewInstanceCreateSuccessInstance instantiates a new InstanceCreateSuccessInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceCreateSuccessInstanceWithDefaults

`func NewInstanceCreateSuccessInstanceWithDefaults() *InstanceCreateSuccessInstance`

NewInstanceCreateSuccessInstanceWithDefaults instantiates a new InstanceCreateSuccessInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceCreateSuccessInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceCreateSuccessInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceCreateSuccessInstance) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceCreateSuccessInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *InstanceCreateSuccessInstance) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InstanceCreateSuccessInstance) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InstanceCreateSuccessInstance) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InstanceCreateSuccessInstance) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccountId

`func (o *InstanceCreateSuccessInstance) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InstanceCreateSuccessInstance) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InstanceCreateSuccessInstance) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *InstanceCreateSuccessInstance) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetTenant

`func (o *InstanceCreateSuccessInstance) GetTenant() InstanceCreateSuccessInstanceTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *InstanceCreateSuccessInstance) GetTenantOk() (*InstanceCreateSuccessInstanceTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *InstanceCreateSuccessInstance) SetTenant(v InstanceCreateSuccessInstanceTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *InstanceCreateSuccessInstance) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetInstanceType

`func (o *InstanceCreateSuccessInstance) GetInstanceType() InstanceCreateSuccessInstanceInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *InstanceCreateSuccessInstance) GetInstanceTypeOk() (*InstanceCreateSuccessInstanceInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *InstanceCreateSuccessInstance) SetInstanceType(v InstanceCreateSuccessInstanceInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *InstanceCreateSuccessInstance) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetGroup

`func (o *InstanceCreateSuccessInstance) GetGroup() InstanceCreateSuccessInstanceGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InstanceCreateSuccessInstance) GetGroupOk() (*InstanceCreateSuccessInstanceGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InstanceCreateSuccessInstance) SetGroup(v InstanceCreateSuccessInstanceGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InstanceCreateSuccessInstance) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *InstanceCreateSuccessInstance) GetCloud() InstanceCreateSuccessInstanceCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceCreateSuccessInstance) GetCloudOk() (*InstanceCreateSuccessInstanceCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceCreateSuccessInstance) SetCloud(v InstanceCreateSuccessInstanceCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *InstanceCreateSuccessInstance) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCluster

`func (o *InstanceCreateSuccessInstance) GetCluster() InstanceCreateSuccessInstanceCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InstanceCreateSuccessInstance) GetClusterOk() (*InstanceCreateSuccessInstanceCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InstanceCreateSuccessInstance) SetCluster(v InstanceCreateSuccessInstanceCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *InstanceCreateSuccessInstance) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetContainers

`func (o *InstanceCreateSuccessInstance) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *InstanceCreateSuccessInstance) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *InstanceCreateSuccessInstance) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *InstanceCreateSuccessInstance) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetServers

`func (o *InstanceCreateSuccessInstance) GetServers() []int64`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *InstanceCreateSuccessInstance) GetServersOk() (*[]int64, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *InstanceCreateSuccessInstance) SetServers(v []int64)`

SetServers sets Servers field to given value.

### HasServers

`func (o *InstanceCreateSuccessInstance) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetConnectionInfo

`func (o *InstanceCreateSuccessInstance) GetConnectionInfo() []AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner`

GetConnectionInfo returns the ConnectionInfo field if non-nil, zero value otherwise.

### GetConnectionInfoOk

`func (o *InstanceCreateSuccessInstance) GetConnectionInfoOk() (*[]AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner, bool)`

GetConnectionInfoOk returns a tuple with the ConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionInfo

`func (o *InstanceCreateSuccessInstance) SetConnectionInfo(v []AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner)`

SetConnectionInfo sets ConnectionInfo field to given value.

### HasConnectionInfo

`func (o *InstanceCreateSuccessInstance) HasConnectionInfo() bool`

HasConnectionInfo returns a boolean if a field has been set.

### GetLayout

`func (o *InstanceCreateSuccessInstance) GetLayout() InstanceCreateSuccessInstanceLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InstanceCreateSuccessInstance) GetLayoutOk() (*InstanceCreateSuccessInstanceLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InstanceCreateSuccessInstance) SetLayout(v InstanceCreateSuccessInstanceLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *InstanceCreateSuccessInstance) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPlan

`func (o *InstanceCreateSuccessInstance) GetPlan() InstanceCreateSuccessInstancePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceCreateSuccessInstance) GetPlanOk() (*InstanceCreateSuccessInstancePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceCreateSuccessInstance) SetPlan(v InstanceCreateSuccessInstancePlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *InstanceCreateSuccessInstance) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetName

`func (o *InstanceCreateSuccessInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceCreateSuccessInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceCreateSuccessInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceCreateSuccessInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *InstanceCreateSuccessInstance) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InstanceCreateSuccessInstance) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InstanceCreateSuccessInstance) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InstanceCreateSuccessInstance) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceCreateSuccessInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceCreateSuccessInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceCreateSuccessInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceCreateSuccessInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InstanceCreateSuccessInstance) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InstanceCreateSuccessInstance) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnvironment

`func (o *InstanceCreateSuccessInstance) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *InstanceCreateSuccessInstance) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *InstanceCreateSuccessInstance) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *InstanceCreateSuccessInstance) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *InstanceCreateSuccessInstance) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *InstanceCreateSuccessInstance) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetConfig

`func (o *InstanceCreateSuccessInstance) GetConfig() InstanceCreateSuccessInstanceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InstanceCreateSuccessInstance) GetConfigOk() (*InstanceCreateSuccessInstanceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InstanceCreateSuccessInstance) SetConfig(v InstanceCreateSuccessInstanceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InstanceCreateSuccessInstance) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConfigGroup

`func (o *InstanceCreateSuccessInstance) GetConfigGroup() string`

GetConfigGroup returns the ConfigGroup field if non-nil, zero value otherwise.

### GetConfigGroupOk

`func (o *InstanceCreateSuccessInstance) GetConfigGroupOk() (*string, bool)`

GetConfigGroupOk returns a tuple with the ConfigGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGroup

`func (o *InstanceCreateSuccessInstance) SetConfigGroup(v string)`

SetConfigGroup sets ConfigGroup field to given value.

### HasConfigGroup

`func (o *InstanceCreateSuccessInstance) HasConfigGroup() bool`

HasConfigGroup returns a boolean if a field has been set.

### SetConfigGroupNil

`func (o *InstanceCreateSuccessInstance) SetConfigGroupNil(b bool)`

 SetConfigGroupNil sets the value for ConfigGroup to be an explicit nil

### UnsetConfigGroup
`func (o *InstanceCreateSuccessInstance) UnsetConfigGroup()`

UnsetConfigGroup ensures that no value is present for ConfigGroup, not even an explicit nil
### GetConfigId

`func (o *InstanceCreateSuccessInstance) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *InstanceCreateSuccessInstance) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *InstanceCreateSuccessInstance) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *InstanceCreateSuccessInstance) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### SetConfigIdNil

`func (o *InstanceCreateSuccessInstance) SetConfigIdNil(b bool)`

 SetConfigIdNil sets the value for ConfigId to be an explicit nil

### UnsetConfigId
`func (o *InstanceCreateSuccessInstance) UnsetConfigId()`

UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
### GetConfigRole

`func (o *InstanceCreateSuccessInstance) GetConfigRole() string`

GetConfigRole returns the ConfigRole field if non-nil, zero value otherwise.

### GetConfigRoleOk

`func (o *InstanceCreateSuccessInstance) GetConfigRoleOk() (*string, bool)`

GetConfigRoleOk returns a tuple with the ConfigRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRole

`func (o *InstanceCreateSuccessInstance) SetConfigRole(v string)`

SetConfigRole sets ConfigRole field to given value.

### HasConfigRole

`func (o *InstanceCreateSuccessInstance) HasConfigRole() bool`

HasConfigRole returns a boolean if a field has been set.

### SetConfigRoleNil

`func (o *InstanceCreateSuccessInstance) SetConfigRoleNil(b bool)`

 SetConfigRoleNil sets the value for ConfigRole to be an explicit nil

### UnsetConfigRole
`func (o *InstanceCreateSuccessInstance) UnsetConfigRole()`

UnsetConfigRole ensures that no value is present for ConfigRole, not even an explicit nil
### GetVolumes

`func (o *InstanceCreateSuccessInstance) GetVolumes() []AddInstance200ResponseAllOfOneOfInstanceVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InstanceCreateSuccessInstance) GetVolumesOk() (*[]AddInstance200ResponseAllOfOneOfInstanceVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InstanceCreateSuccessInstance) SetVolumes(v []AddInstance200ResponseAllOfOneOfInstanceVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *InstanceCreateSuccessInstance) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *InstanceCreateSuccessInstance) GetControllers() []AddInstance200ResponseAllOfOneOfInstanceControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *InstanceCreateSuccessInstance) GetControllersOk() (*[]AddInstance200ResponseAllOfOneOfInstanceControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *InstanceCreateSuccessInstance) SetControllers(v []AddInstance200ResponseAllOfOneOfInstanceControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *InstanceCreateSuccessInstance) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *InstanceCreateSuccessInstance) GetInterfaces() []AddInstance200ResponseAllOfOneOfInstanceInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InstanceCreateSuccessInstance) GetInterfacesOk() (*[]AddInstance200ResponseAllOfOneOfInstanceInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InstanceCreateSuccessInstance) SetInterfaces(v []AddInstance200ResponseAllOfOneOfInstanceInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InstanceCreateSuccessInstance) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetCustomOptions

`func (o *InstanceCreateSuccessInstance) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *InstanceCreateSuccessInstance) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *InstanceCreateSuccessInstance) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *InstanceCreateSuccessInstance) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetInstanceVersion

`func (o *InstanceCreateSuccessInstance) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *InstanceCreateSuccessInstance) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *InstanceCreateSuccessInstance) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *InstanceCreateSuccessInstance) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceCreateSuccessInstance) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceCreateSuccessInstance) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceCreateSuccessInstance) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceCreateSuccessInstance) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *InstanceCreateSuccessInstance) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *InstanceCreateSuccessInstance) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetTags

`func (o *InstanceCreateSuccessInstance) GetTags() []AddInstance200ResponseAllOfOneOfInstanceTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstanceCreateSuccessInstance) GetTagsOk() (*[]AddInstance200ResponseAllOfOneOfInstanceTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstanceCreateSuccessInstance) SetTags(v []AddInstance200ResponseAllOfOneOfInstanceTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstanceCreateSuccessInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEvars

`func (o *InstanceCreateSuccessInstance) GetEvars() []AddInstance200ResponseAllOfOneOfInstanceEvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *InstanceCreateSuccessInstance) GetEvarsOk() (*[]AddInstance200ResponseAllOfOneOfInstanceEvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *InstanceCreateSuccessInstance) SetEvars(v []AddInstance200ResponseAllOfOneOfInstanceEvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *InstanceCreateSuccessInstance) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetMaxMemory

`func (o *InstanceCreateSuccessInstance) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *InstanceCreateSuccessInstance) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *InstanceCreateSuccessInstance) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *InstanceCreateSuccessInstance) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *InstanceCreateSuccessInstance) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *InstanceCreateSuccessInstance) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *InstanceCreateSuccessInstance) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *InstanceCreateSuccessInstance) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCores

`func (o *InstanceCreateSuccessInstance) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *InstanceCreateSuccessInstance) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *InstanceCreateSuccessInstance) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *InstanceCreateSuccessInstance) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *InstanceCreateSuccessInstance) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *InstanceCreateSuccessInstance) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *InstanceCreateSuccessInstance) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *InstanceCreateSuccessInstance) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### SetCoresPerSocketNil

`func (o *InstanceCreateSuccessInstance) SetCoresPerSocketNil(b bool)`

 SetCoresPerSocketNil sets the value for CoresPerSocket to be an explicit nil

### UnsetCoresPerSocket
`func (o *InstanceCreateSuccessInstance) UnsetCoresPerSocket()`

UnsetCoresPerSocket ensures that no value is present for CoresPerSocket, not even an explicit nil
### GetMaxCpu

`func (o *InstanceCreateSuccessInstance) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *InstanceCreateSuccessInstance) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *InstanceCreateSuccessInstance) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *InstanceCreateSuccessInstance) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *InstanceCreateSuccessInstance) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *InstanceCreateSuccessInstance) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetHourlyCost

`func (o *InstanceCreateSuccessInstance) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *InstanceCreateSuccessInstance) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *InstanceCreateSuccessInstance) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *InstanceCreateSuccessInstance) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *InstanceCreateSuccessInstance) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *InstanceCreateSuccessInstance) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *InstanceCreateSuccessInstance) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *InstanceCreateSuccessInstance) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetInstancePrice

`func (o *InstanceCreateSuccessInstance) GetInstancePrice() InstanceCreateSuccessInstanceInstancePrice`

GetInstancePrice returns the InstancePrice field if non-nil, zero value otherwise.

### GetInstancePriceOk

`func (o *InstanceCreateSuccessInstance) GetInstancePriceOk() (*InstanceCreateSuccessInstanceInstancePrice, bool)`

GetInstancePriceOk returns a tuple with the InstancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePrice

`func (o *InstanceCreateSuccessInstance) SetInstancePrice(v InstanceCreateSuccessInstanceInstancePrice)`

SetInstancePrice sets InstancePrice field to given value.

### HasInstancePrice

`func (o *InstanceCreateSuccessInstance) HasInstancePrice() bool`

HasInstancePrice returns a boolean if a field has been set.

### GetDateCreated

`func (o *InstanceCreateSuccessInstance) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InstanceCreateSuccessInstance) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InstanceCreateSuccessInstance) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InstanceCreateSuccessInstance) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InstanceCreateSuccessInstance) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InstanceCreateSuccessInstance) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InstanceCreateSuccessInstance) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InstanceCreateSuccessInstance) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHostName

`func (o *InstanceCreateSuccessInstance) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *InstanceCreateSuccessInstance) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *InstanceCreateSuccessInstance) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *InstanceCreateSuccessInstance) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetDomainName

`func (o *InstanceCreateSuccessInstance) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *InstanceCreateSuccessInstance) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *InstanceCreateSuccessInstance) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *InstanceCreateSuccessInstance) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *InstanceCreateSuccessInstance) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *InstanceCreateSuccessInstance) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetEnvironmentPrefix

`func (o *InstanceCreateSuccessInstance) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *InstanceCreateSuccessInstance) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *InstanceCreateSuccessInstance) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *InstanceCreateSuccessInstance) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### SetEnvironmentPrefixNil

`func (o *InstanceCreateSuccessInstance) SetEnvironmentPrefixNil(b bool)`

 SetEnvironmentPrefixNil sets the value for EnvironmentPrefix to be an explicit nil

### UnsetEnvironmentPrefix
`func (o *InstanceCreateSuccessInstance) UnsetEnvironmentPrefix()`

UnsetEnvironmentPrefix ensures that no value is present for EnvironmentPrefix, not even an explicit nil
### GetFirewallEnabled

`func (o *InstanceCreateSuccessInstance) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *InstanceCreateSuccessInstance) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *InstanceCreateSuccessInstance) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *InstanceCreateSuccessInstance) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetNetworkLevel

`func (o *InstanceCreateSuccessInstance) GetNetworkLevel() string`

GetNetworkLevel returns the NetworkLevel field if non-nil, zero value otherwise.

### GetNetworkLevelOk

`func (o *InstanceCreateSuccessInstance) GetNetworkLevelOk() (*string, bool)`

GetNetworkLevelOk returns a tuple with the NetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLevel

`func (o *InstanceCreateSuccessInstance) SetNetworkLevel(v string)`

SetNetworkLevel sets NetworkLevel field to given value.

### HasNetworkLevel

`func (o *InstanceCreateSuccessInstance) HasNetworkLevel() bool`

HasNetworkLevel returns a boolean if a field has been set.

### GetAutoScale

`func (o *InstanceCreateSuccessInstance) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *InstanceCreateSuccessInstance) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *InstanceCreateSuccessInstance) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *InstanceCreateSuccessInstance) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetInstanceContext

`func (o *InstanceCreateSuccessInstance) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *InstanceCreateSuccessInstance) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *InstanceCreateSuccessInstance) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *InstanceCreateSuccessInstance) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### SetInstanceContextNil

`func (o *InstanceCreateSuccessInstance) SetInstanceContextNil(b bool)`

 SetInstanceContextNil sets the value for InstanceContext to be an explicit nil

### UnsetInstanceContext
`func (o *InstanceCreateSuccessInstance) UnsetInstanceContext()`

UnsetInstanceContext ensures that no value is present for InstanceContext, not even an explicit nil
### GetCurrentDeployId

`func (o *InstanceCreateSuccessInstance) GetCurrentDeployId() string`

GetCurrentDeployId returns the CurrentDeployId field if non-nil, zero value otherwise.

### GetCurrentDeployIdOk

`func (o *InstanceCreateSuccessInstance) GetCurrentDeployIdOk() (*string, bool)`

GetCurrentDeployIdOk returns a tuple with the CurrentDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeployId

`func (o *InstanceCreateSuccessInstance) SetCurrentDeployId(v string)`

SetCurrentDeployId sets CurrentDeployId field to given value.

### HasCurrentDeployId

`func (o *InstanceCreateSuccessInstance) HasCurrentDeployId() bool`

HasCurrentDeployId returns a boolean if a field has been set.

### SetCurrentDeployIdNil

`func (o *InstanceCreateSuccessInstance) SetCurrentDeployIdNil(b bool)`

 SetCurrentDeployIdNil sets the value for CurrentDeployId to be an explicit nil

### UnsetCurrentDeployId
`func (o *InstanceCreateSuccessInstance) UnsetCurrentDeployId()`

UnsetCurrentDeployId ensures that no value is present for CurrentDeployId, not even an explicit nil
### GetLocked

`func (o *InstanceCreateSuccessInstance) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *InstanceCreateSuccessInstance) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *InstanceCreateSuccessInstance) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *InstanceCreateSuccessInstance) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetStatus

`func (o *InstanceCreateSuccessInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceCreateSuccessInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceCreateSuccessInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstanceCreateSuccessInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *InstanceCreateSuccessInstance) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *InstanceCreateSuccessInstance) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *InstanceCreateSuccessInstance) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *InstanceCreateSuccessInstance) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *InstanceCreateSuccessInstance) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *InstanceCreateSuccessInstance) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetErrorMessage

`func (o *InstanceCreateSuccessInstance) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *InstanceCreateSuccessInstance) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *InstanceCreateSuccessInstance) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *InstanceCreateSuccessInstance) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *InstanceCreateSuccessInstance) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *InstanceCreateSuccessInstance) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStatusDate

`func (o *InstanceCreateSuccessInstance) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *InstanceCreateSuccessInstance) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *InstanceCreateSuccessInstance) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *InstanceCreateSuccessInstance) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *InstanceCreateSuccessInstance) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *InstanceCreateSuccessInstance) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *InstanceCreateSuccessInstance) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *InstanceCreateSuccessInstance) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### SetStatusPercentNil

`func (o *InstanceCreateSuccessInstance) SetStatusPercentNil(b bool)`

 SetStatusPercentNil sets the value for StatusPercent to be an explicit nil

### UnsetStatusPercent
`func (o *InstanceCreateSuccessInstance) UnsetStatusPercent()`

UnsetStatusPercent ensures that no value is present for StatusPercent, not even an explicit nil
### GetStatusEta

`func (o *InstanceCreateSuccessInstance) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *InstanceCreateSuccessInstance) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *InstanceCreateSuccessInstance) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *InstanceCreateSuccessInstance) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### SetStatusEtaNil

`func (o *InstanceCreateSuccessInstance) SetStatusEtaNil(b bool)`

 SetStatusEtaNil sets the value for StatusEta to be an explicit nil

### UnsetStatusEta
`func (o *InstanceCreateSuccessInstance) UnsetStatusEta()`

UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
### GetUserStatus

`func (o *InstanceCreateSuccessInstance) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *InstanceCreateSuccessInstance) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *InstanceCreateSuccessInstance) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *InstanceCreateSuccessInstance) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### SetUserStatusNil

`func (o *InstanceCreateSuccessInstance) SetUserStatusNil(b bool)`

 SetUserStatusNil sets the value for UserStatus to be an explicit nil

### UnsetUserStatus
`func (o *InstanceCreateSuccessInstance) UnsetUserStatus()`

UnsetUserStatus ensures that no value is present for UserStatus, not even an explicit nil
### GetExpireDays

`func (o *InstanceCreateSuccessInstance) GetExpireDays() int64`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *InstanceCreateSuccessInstance) GetExpireDaysOk() (*int64, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *InstanceCreateSuccessInstance) SetExpireDays(v int64)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *InstanceCreateSuccessInstance) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetRenewDays

`func (o *InstanceCreateSuccessInstance) GetRenewDays() int64`

GetRenewDays returns the RenewDays field if non-nil, zero value otherwise.

### GetRenewDaysOk

`func (o *InstanceCreateSuccessInstance) GetRenewDaysOk() (*int64, bool)`

GetRenewDaysOk returns a tuple with the RenewDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewDays

`func (o *InstanceCreateSuccessInstance) SetRenewDays(v int64)`

SetRenewDays sets RenewDays field to given value.

### HasRenewDays

`func (o *InstanceCreateSuccessInstance) HasRenewDays() bool`

HasRenewDays returns a boolean if a field has been set.

### GetExpireCount

`func (o *InstanceCreateSuccessInstance) GetExpireCount() int64`

GetExpireCount returns the ExpireCount field if non-nil, zero value otherwise.

### GetExpireCountOk

`func (o *InstanceCreateSuccessInstance) GetExpireCountOk() (*int64, bool)`

GetExpireCountOk returns a tuple with the ExpireCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireCount

`func (o *InstanceCreateSuccessInstance) SetExpireCount(v int64)`

SetExpireCount sets ExpireCount field to given value.

### HasExpireCount

`func (o *InstanceCreateSuccessInstance) HasExpireCount() bool`

HasExpireCount returns a boolean if a field has been set.

### GetExpireDate

`func (o *InstanceCreateSuccessInstance) GetExpireDate() time.Time`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *InstanceCreateSuccessInstance) GetExpireDateOk() (*time.Time, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *InstanceCreateSuccessInstance) SetExpireDate(v time.Time)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *InstanceCreateSuccessInstance) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetExpireWarningDate

`func (o *InstanceCreateSuccessInstance) GetExpireWarningDate() time.Time`

GetExpireWarningDate returns the ExpireWarningDate field if non-nil, zero value otherwise.

### GetExpireWarningDateOk

`func (o *InstanceCreateSuccessInstance) GetExpireWarningDateOk() (*time.Time, bool)`

GetExpireWarningDateOk returns a tuple with the ExpireWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWarningDate

`func (o *InstanceCreateSuccessInstance) SetExpireWarningDate(v time.Time)`

SetExpireWarningDate sets ExpireWarningDate field to given value.

### HasExpireWarningDate

`func (o *InstanceCreateSuccessInstance) HasExpireWarningDate() bool`

HasExpireWarningDate returns a boolean if a field has been set.

### GetExpireWarningSent

`func (o *InstanceCreateSuccessInstance) GetExpireWarningSent() bool`

GetExpireWarningSent returns the ExpireWarningSent field if non-nil, zero value otherwise.

### GetExpireWarningSentOk

`func (o *InstanceCreateSuccessInstance) GetExpireWarningSentOk() (*bool, bool)`

GetExpireWarningSentOk returns a tuple with the ExpireWarningSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWarningSent

`func (o *InstanceCreateSuccessInstance) SetExpireWarningSent(v bool)`

SetExpireWarningSent sets ExpireWarningSent field to given value.

### HasExpireWarningSent

`func (o *InstanceCreateSuccessInstance) HasExpireWarningSent() bool`

HasExpireWarningSent returns a boolean if a field has been set.

### GetShutdownDays

`func (o *InstanceCreateSuccessInstance) GetShutdownDays() int64`

GetShutdownDays returns the ShutdownDays field if non-nil, zero value otherwise.

### GetShutdownDaysOk

`func (o *InstanceCreateSuccessInstance) GetShutdownDaysOk() (*int64, bool)`

GetShutdownDaysOk returns a tuple with the ShutdownDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDays

`func (o *InstanceCreateSuccessInstance) SetShutdownDays(v int64)`

SetShutdownDays sets ShutdownDays field to given value.

### HasShutdownDays

`func (o *InstanceCreateSuccessInstance) HasShutdownDays() bool`

HasShutdownDays returns a boolean if a field has been set.

### GetShutdownRenewDays

`func (o *InstanceCreateSuccessInstance) GetShutdownRenewDays() int64`

GetShutdownRenewDays returns the ShutdownRenewDays field if non-nil, zero value otherwise.

### GetShutdownRenewDaysOk

`func (o *InstanceCreateSuccessInstance) GetShutdownRenewDaysOk() (*int64, bool)`

GetShutdownRenewDaysOk returns a tuple with the ShutdownRenewDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewDays

`func (o *InstanceCreateSuccessInstance) SetShutdownRenewDays(v int64)`

SetShutdownRenewDays sets ShutdownRenewDays field to given value.

### HasShutdownRenewDays

`func (o *InstanceCreateSuccessInstance) HasShutdownRenewDays() bool`

HasShutdownRenewDays returns a boolean if a field has been set.

### GetShutdownCount

`func (o *InstanceCreateSuccessInstance) GetShutdownCount() int64`

GetShutdownCount returns the ShutdownCount field if non-nil, zero value otherwise.

### GetShutdownCountOk

`func (o *InstanceCreateSuccessInstance) GetShutdownCountOk() (*int64, bool)`

GetShutdownCountOk returns a tuple with the ShutdownCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownCount

`func (o *InstanceCreateSuccessInstance) SetShutdownCount(v int64)`

SetShutdownCount sets ShutdownCount field to given value.

### HasShutdownCount

`func (o *InstanceCreateSuccessInstance) HasShutdownCount() bool`

HasShutdownCount returns a boolean if a field has been set.

### GetShutdownDate

`func (o *InstanceCreateSuccessInstance) GetShutdownDate() time.Time`

GetShutdownDate returns the ShutdownDate field if non-nil, zero value otherwise.

### GetShutdownDateOk

`func (o *InstanceCreateSuccessInstance) GetShutdownDateOk() (*time.Time, bool)`

GetShutdownDateOk returns a tuple with the ShutdownDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDate

`func (o *InstanceCreateSuccessInstance) SetShutdownDate(v time.Time)`

SetShutdownDate sets ShutdownDate field to given value.

### HasShutdownDate

`func (o *InstanceCreateSuccessInstance) HasShutdownDate() bool`

HasShutdownDate returns a boolean if a field has been set.

### GetShutdownWarningDate

`func (o *InstanceCreateSuccessInstance) GetShutdownWarningDate() time.Time`

GetShutdownWarningDate returns the ShutdownWarningDate field if non-nil, zero value otherwise.

### GetShutdownWarningDateOk

`func (o *InstanceCreateSuccessInstance) GetShutdownWarningDateOk() (*time.Time, bool)`

GetShutdownWarningDateOk returns a tuple with the ShutdownWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWarningDate

`func (o *InstanceCreateSuccessInstance) SetShutdownWarningDate(v time.Time)`

SetShutdownWarningDate sets ShutdownWarningDate field to given value.

### HasShutdownWarningDate

`func (o *InstanceCreateSuccessInstance) HasShutdownWarningDate() bool`

HasShutdownWarningDate returns a boolean if a field has been set.

### GetShutdownWarningSent

`func (o *InstanceCreateSuccessInstance) GetShutdownWarningSent() bool`

GetShutdownWarningSent returns the ShutdownWarningSent field if non-nil, zero value otherwise.

### GetShutdownWarningSentOk

`func (o *InstanceCreateSuccessInstance) GetShutdownWarningSentOk() (*bool, bool)`

GetShutdownWarningSentOk returns a tuple with the ShutdownWarningSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWarningSent

`func (o *InstanceCreateSuccessInstance) SetShutdownWarningSent(v bool)`

SetShutdownWarningSent sets ShutdownWarningSent field to given value.

### HasShutdownWarningSent

`func (o *InstanceCreateSuccessInstance) HasShutdownWarningSent() bool`

HasShutdownWarningSent returns a boolean if a field has been set.

### GetRemovalDate

`func (o *InstanceCreateSuccessInstance) GetRemovalDate() time.Time`

GetRemovalDate returns the RemovalDate field if non-nil, zero value otherwise.

### GetRemovalDateOk

`func (o *InstanceCreateSuccessInstance) GetRemovalDateOk() (*time.Time, bool)`

GetRemovalDateOk returns a tuple with the RemovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalDate

`func (o *InstanceCreateSuccessInstance) SetRemovalDate(v time.Time)`

SetRemovalDate sets RemovalDate field to given value.

### HasRemovalDate

`func (o *InstanceCreateSuccessInstance) HasRemovalDate() bool`

HasRemovalDate returns a boolean if a field has been set.

### SetRemovalDateNil

`func (o *InstanceCreateSuccessInstance) SetRemovalDateNil(b bool)`

 SetRemovalDateNil sets the value for RemovalDate to be an explicit nil

### UnsetRemovalDate
`func (o *InstanceCreateSuccessInstance) UnsetRemovalDate()`

UnsetRemovalDate ensures that no value is present for RemovalDate, not even an explicit nil
### GetCreatedBy

`func (o *InstanceCreateSuccessInstance) GetCreatedBy() InstanceCreateSuccessInstanceCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InstanceCreateSuccessInstance) GetCreatedByOk() (*InstanceCreateSuccessInstanceCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InstanceCreateSuccessInstance) SetCreatedBy(v InstanceCreateSuccessInstanceCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InstanceCreateSuccessInstance) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetOwner

`func (o *InstanceCreateSuccessInstance) GetOwner() InstanceCreateSuccessInstanceOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *InstanceCreateSuccessInstance) GetOwnerOk() (*InstanceCreateSuccessInstanceOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *InstanceCreateSuccessInstance) SetOwner(v InstanceCreateSuccessInstanceOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *InstanceCreateSuccessInstance) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNotes

`func (o *InstanceCreateSuccessInstance) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InstanceCreateSuccessInstance) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InstanceCreateSuccessInstance) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InstanceCreateSuccessInstance) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *InstanceCreateSuccessInstance) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *InstanceCreateSuccessInstance) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetStats

`func (o *InstanceCreateSuccessInstance) GetStats() InstanceCreateSuccessInstanceStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *InstanceCreateSuccessInstance) GetStatsOk() (*InstanceCreateSuccessInstanceStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *InstanceCreateSuccessInstance) SetStats(v InstanceCreateSuccessInstanceStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *InstanceCreateSuccessInstance) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetPowerSchedule

`func (o *InstanceCreateSuccessInstance) GetPowerSchedule() string`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *InstanceCreateSuccessInstance) GetPowerScheduleOk() (*string, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *InstanceCreateSuccessInstance) SetPowerSchedule(v string)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *InstanceCreateSuccessInstance) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### SetPowerScheduleNil

`func (o *InstanceCreateSuccessInstance) SetPowerScheduleNil(b bool)`

 SetPowerScheduleNil sets the value for PowerSchedule to be an explicit nil

### UnsetPowerSchedule
`func (o *InstanceCreateSuccessInstance) UnsetPowerSchedule()`

UnsetPowerSchedule ensures that no value is present for PowerSchedule, not even an explicit nil
### GetIsScalable

`func (o *InstanceCreateSuccessInstance) GetIsScalable() bool`

GetIsScalable returns the IsScalable field if non-nil, zero value otherwise.

### GetIsScalableOk

`func (o *InstanceCreateSuccessInstance) GetIsScalableOk() (*bool, bool)`

GetIsScalableOk returns a tuple with the IsScalable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScalable

`func (o *InstanceCreateSuccessInstance) SetIsScalable(v bool)`

SetIsScalable sets IsScalable field to given value.

### HasIsScalable

`func (o *InstanceCreateSuccessInstance) HasIsScalable() bool`

HasIsScalable returns a boolean if a field has been set.

### GetInstanceThreshold

`func (o *InstanceCreateSuccessInstance) GetInstanceThreshold() map[string]interface{}`

GetInstanceThreshold returns the InstanceThreshold field if non-nil, zero value otherwise.

### GetInstanceThresholdOk

`func (o *InstanceCreateSuccessInstance) GetInstanceThresholdOk() (*map[string]interface{}, bool)`

GetInstanceThresholdOk returns a tuple with the InstanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceThreshold

`func (o *InstanceCreateSuccessInstance) SetInstanceThreshold(v map[string]interface{})`

SetInstanceThreshold sets InstanceThreshold field to given value.

### HasInstanceThreshold

`func (o *InstanceCreateSuccessInstance) HasInstanceThreshold() bool`

HasInstanceThreshold returns a boolean if a field has been set.

### SetInstanceThresholdNil

`func (o *InstanceCreateSuccessInstance) SetInstanceThresholdNil(b bool)`

 SetInstanceThresholdNil sets the value for InstanceThreshold to be an explicit nil

### UnsetInstanceThreshold
`func (o *InstanceCreateSuccessInstance) UnsetInstanceThreshold()`

UnsetInstanceThreshold ensures that no value is present for InstanceThreshold, not even an explicit nil
### GetIsBusy

`func (o *InstanceCreateSuccessInstance) GetIsBusy() bool`

GetIsBusy returns the IsBusy field if non-nil, zero value otherwise.

### GetIsBusyOk

`func (o *InstanceCreateSuccessInstance) GetIsBusyOk() (*bool, bool)`

GetIsBusyOk returns a tuple with the IsBusy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBusy

`func (o *InstanceCreateSuccessInstance) SetIsBusy(v bool)`

SetIsBusy sets IsBusy field to given value.

### HasIsBusy

`func (o *InstanceCreateSuccessInstance) HasIsBusy() bool`

HasIsBusy returns a boolean if a field has been set.

### GetApps

`func (o *InstanceCreateSuccessInstance) GetApps() []map[string]interface{}`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *InstanceCreateSuccessInstance) GetAppsOk() (*[]map[string]interface{}, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *InstanceCreateSuccessInstance) SetApps(v []map[string]interface{})`

SetApps sets Apps field to given value.

### HasApps

`func (o *InstanceCreateSuccessInstance) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *InstanceCreateSuccessInstance) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *InstanceCreateSuccessInstance) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetContainerDetails

`func (o *InstanceCreateSuccessInstance) GetContainerDetails() []InstanceContainer4`

GetContainerDetails returns the ContainerDetails field if non-nil, zero value otherwise.

### GetContainerDetailsOk

`func (o *InstanceCreateSuccessInstance) GetContainerDetailsOk() (*[]InstanceContainer4, bool)`

GetContainerDetailsOk returns a tuple with the ContainerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDetails

`func (o *InstanceCreateSuccessInstance) SetContainerDetails(v []InstanceContainer4)`

SetContainerDetails sets ContainerDetails field to given value.

### HasContainerDetails

`func (o *InstanceCreateSuccessInstance) HasContainerDetails() bool`

HasContainerDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



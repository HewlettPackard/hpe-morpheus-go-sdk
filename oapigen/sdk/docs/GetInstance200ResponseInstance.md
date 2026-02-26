# GetInstance200ResponseInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Tenant** | Pointer to [**GetInstance200ResponseInstanceTenant**](GetInstance200ResponseInstanceTenant.md) |  | [optional] 
**InstanceType** | Pointer to [**GetInstance200ResponseInstanceInstanceType**](GetInstance200ResponseInstanceInstanceType.md) |  | [optional] 
**Group** | Pointer to [**GetInstance200ResponseInstanceGroup**](GetInstance200ResponseInstanceGroup.md) |  | [optional] 
**Cloud** | Pointer to [**GetInstance200ResponseInstanceCloud**](GetInstance200ResponseInstanceCloud.md) |  | [optional] 
**Cluster** | Pointer to [**GetInstance200ResponseInstanceCluster**](GetInstance200ResponseInstanceCluster.md) |  | [optional] 
**Containers** | Pointer to **[]int64** |  | [optional] 
**Servers** | Pointer to **[]int64** |  | [optional] 
**ConnectionInfo** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner**](AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner.md) |  | [optional] 
**Layout** | Pointer to [**GetInstance200ResponseInstanceLayout**](GetInstance200ResponseInstanceLayout.md) |  | [optional] 
**Plan** | Pointer to [**GetInstance200ResponseInstancePlan**](GetInstance200ResponseInstancePlan.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Environment** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to [**GetInstance200ResponseInstanceConfig**](GetInstance200ResponseInstanceConfig.md) |  | [optional] 
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
**InstancePrice** | Pointer to [**GetInstance200ResponseInstanceInstancePrice**](GetInstance200ResponseInstanceInstancePrice.md) |  | [optional] 
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
**CreatedBy** | Pointer to [**GetInstance200ResponseInstanceCreatedBy**](GetInstance200ResponseInstanceCreatedBy.md) |  | [optional] 
**Owner** | Pointer to [**GetInstance200ResponseInstanceOwner**](GetInstance200ResponseInstanceOwner.md) |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Stats** | Pointer to [**GetInstance200ResponseInstanceStats**](GetInstance200ResponseInstanceStats.md) |  | [optional] 
**PowerSchedule** | Pointer to **NullableString** |  | [optional] 
**IsScalable** | Pointer to **bool** |  | [optional] 
**InstanceThreshold** | Pointer to **map[string]interface{}** |  | [optional] 
**IsBusy** | Pointer to **bool** |  | [optional] 
**Apps** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CurrentLoadBalancerInstances** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CurrentLoadBalancerContainersIn** | Pointer to **int32** |  | [optional] 
**CurrentLoadBalancerContainersOut** | Pointer to **int32** |  | [optional] 
**LastDeploy** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerDetails** | Pointer to [**[]InstanceContainer2**](InstanceContainer2.md) |  | [optional] 

## Methods

### NewGetInstance200ResponseInstance

`func NewGetInstance200ResponseInstance() *GetInstance200ResponseInstance`

NewGetInstance200ResponseInstance instantiates a new GetInstance200ResponseInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstance200ResponseInstanceWithDefaults

`func NewGetInstance200ResponseInstanceWithDefaults() *GetInstance200ResponseInstance`

NewGetInstance200ResponseInstanceWithDefaults instantiates a new GetInstance200ResponseInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstance200ResponseInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstance200ResponseInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstance200ResponseInstance) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstance200ResponseInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GetInstance200ResponseInstance) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetInstance200ResponseInstance) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetInstance200ResponseInstance) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetInstance200ResponseInstance) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccountId

`func (o *GetInstance200ResponseInstance) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetInstance200ResponseInstance) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetInstance200ResponseInstance) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetInstance200ResponseInstance) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetTenant

`func (o *GetInstance200ResponseInstance) GetTenant() GetInstance200ResponseInstanceTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *GetInstance200ResponseInstance) GetTenantOk() (*GetInstance200ResponseInstanceTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *GetInstance200ResponseInstance) SetTenant(v GetInstance200ResponseInstanceTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *GetInstance200ResponseInstance) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetInstanceType

`func (o *GetInstance200ResponseInstance) GetInstanceType() GetInstance200ResponseInstanceInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *GetInstance200ResponseInstance) GetInstanceTypeOk() (*GetInstance200ResponseInstanceInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *GetInstance200ResponseInstance) SetInstanceType(v GetInstance200ResponseInstanceInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *GetInstance200ResponseInstance) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetGroup

`func (o *GetInstance200ResponseInstance) GetGroup() GetInstance200ResponseInstanceGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetInstance200ResponseInstance) GetGroupOk() (*GetInstance200ResponseInstanceGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetInstance200ResponseInstance) SetGroup(v GetInstance200ResponseInstanceGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetInstance200ResponseInstance) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *GetInstance200ResponseInstance) GetCloud() GetInstance200ResponseInstanceCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *GetInstance200ResponseInstance) GetCloudOk() (*GetInstance200ResponseInstanceCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *GetInstance200ResponseInstance) SetCloud(v GetInstance200ResponseInstanceCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *GetInstance200ResponseInstance) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCluster

`func (o *GetInstance200ResponseInstance) GetCluster() GetInstance200ResponseInstanceCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *GetInstance200ResponseInstance) GetClusterOk() (*GetInstance200ResponseInstanceCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *GetInstance200ResponseInstance) SetCluster(v GetInstance200ResponseInstanceCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *GetInstance200ResponseInstance) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetContainers

`func (o *GetInstance200ResponseInstance) GetContainers() []int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *GetInstance200ResponseInstance) GetContainersOk() (*[]int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *GetInstance200ResponseInstance) SetContainers(v []int64)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *GetInstance200ResponseInstance) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetServers

`func (o *GetInstance200ResponseInstance) GetServers() []int64`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetInstance200ResponseInstance) GetServersOk() (*[]int64, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetInstance200ResponseInstance) SetServers(v []int64)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetInstance200ResponseInstance) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetConnectionInfo

`func (o *GetInstance200ResponseInstance) GetConnectionInfo() []AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner`

GetConnectionInfo returns the ConnectionInfo field if non-nil, zero value otherwise.

### GetConnectionInfoOk

`func (o *GetInstance200ResponseInstance) GetConnectionInfoOk() (*[]AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner, bool)`

GetConnectionInfoOk returns a tuple with the ConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionInfo

`func (o *GetInstance200ResponseInstance) SetConnectionInfo(v []AddInstance200ResponseAllOfOneOfInstanceConnectionInfoInner)`

SetConnectionInfo sets ConnectionInfo field to given value.

### HasConnectionInfo

`func (o *GetInstance200ResponseInstance) HasConnectionInfo() bool`

HasConnectionInfo returns a boolean if a field has been set.

### GetLayout

`func (o *GetInstance200ResponseInstance) GetLayout() GetInstance200ResponseInstanceLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *GetInstance200ResponseInstance) GetLayoutOk() (*GetInstance200ResponseInstanceLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *GetInstance200ResponseInstance) SetLayout(v GetInstance200ResponseInstanceLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *GetInstance200ResponseInstance) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPlan

`func (o *GetInstance200ResponseInstance) GetPlan() GetInstance200ResponseInstancePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetInstance200ResponseInstance) GetPlanOk() (*GetInstance200ResponseInstancePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetInstance200ResponseInstance) SetPlan(v GetInstance200ResponseInstancePlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetInstance200ResponseInstance) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetName

`func (o *GetInstance200ResponseInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInstance200ResponseInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInstance200ResponseInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInstance200ResponseInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetInstance200ResponseInstance) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetInstance200ResponseInstance) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetInstance200ResponseInstance) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetInstance200ResponseInstance) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *GetInstance200ResponseInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInstance200ResponseInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInstance200ResponseInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetInstance200ResponseInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetInstance200ResponseInstance) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetInstance200ResponseInstance) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnvironment

`func (o *GetInstance200ResponseInstance) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GetInstance200ResponseInstance) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GetInstance200ResponseInstance) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GetInstance200ResponseInstance) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *GetInstance200ResponseInstance) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *GetInstance200ResponseInstance) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetConfig

`func (o *GetInstance200ResponseInstance) GetConfig() GetInstance200ResponseInstanceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetInstance200ResponseInstance) GetConfigOk() (*GetInstance200ResponseInstanceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetInstance200ResponseInstance) SetConfig(v GetInstance200ResponseInstanceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetInstance200ResponseInstance) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConfigGroup

`func (o *GetInstance200ResponseInstance) GetConfigGroup() string`

GetConfigGroup returns the ConfigGroup field if non-nil, zero value otherwise.

### GetConfigGroupOk

`func (o *GetInstance200ResponseInstance) GetConfigGroupOk() (*string, bool)`

GetConfigGroupOk returns a tuple with the ConfigGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGroup

`func (o *GetInstance200ResponseInstance) SetConfigGroup(v string)`

SetConfigGroup sets ConfigGroup field to given value.

### HasConfigGroup

`func (o *GetInstance200ResponseInstance) HasConfigGroup() bool`

HasConfigGroup returns a boolean if a field has been set.

### SetConfigGroupNil

`func (o *GetInstance200ResponseInstance) SetConfigGroupNil(b bool)`

 SetConfigGroupNil sets the value for ConfigGroup to be an explicit nil

### UnsetConfigGroup
`func (o *GetInstance200ResponseInstance) UnsetConfigGroup()`

UnsetConfigGroup ensures that no value is present for ConfigGroup, not even an explicit nil
### GetConfigId

`func (o *GetInstance200ResponseInstance) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *GetInstance200ResponseInstance) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *GetInstance200ResponseInstance) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *GetInstance200ResponseInstance) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### SetConfigIdNil

`func (o *GetInstance200ResponseInstance) SetConfigIdNil(b bool)`

 SetConfigIdNil sets the value for ConfigId to be an explicit nil

### UnsetConfigId
`func (o *GetInstance200ResponseInstance) UnsetConfigId()`

UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
### GetConfigRole

`func (o *GetInstance200ResponseInstance) GetConfigRole() string`

GetConfigRole returns the ConfigRole field if non-nil, zero value otherwise.

### GetConfigRoleOk

`func (o *GetInstance200ResponseInstance) GetConfigRoleOk() (*string, bool)`

GetConfigRoleOk returns a tuple with the ConfigRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRole

`func (o *GetInstance200ResponseInstance) SetConfigRole(v string)`

SetConfigRole sets ConfigRole field to given value.

### HasConfigRole

`func (o *GetInstance200ResponseInstance) HasConfigRole() bool`

HasConfigRole returns a boolean if a field has been set.

### SetConfigRoleNil

`func (o *GetInstance200ResponseInstance) SetConfigRoleNil(b bool)`

 SetConfigRoleNil sets the value for ConfigRole to be an explicit nil

### UnsetConfigRole
`func (o *GetInstance200ResponseInstance) UnsetConfigRole()`

UnsetConfigRole ensures that no value is present for ConfigRole, not even an explicit nil
### GetVolumes

`func (o *GetInstance200ResponseInstance) GetVolumes() []AddInstance200ResponseAllOfOneOfInstanceVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *GetInstance200ResponseInstance) GetVolumesOk() (*[]AddInstance200ResponseAllOfOneOfInstanceVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *GetInstance200ResponseInstance) SetVolumes(v []AddInstance200ResponseAllOfOneOfInstanceVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *GetInstance200ResponseInstance) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetControllers

`func (o *GetInstance200ResponseInstance) GetControllers() []AddInstance200ResponseAllOfOneOfInstanceControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *GetInstance200ResponseInstance) GetControllersOk() (*[]AddInstance200ResponseAllOfOneOfInstanceControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *GetInstance200ResponseInstance) SetControllers(v []AddInstance200ResponseAllOfOneOfInstanceControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *GetInstance200ResponseInstance) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetInstance200ResponseInstance) GetInterfaces() []AddInstance200ResponseAllOfOneOfInstanceInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetInstance200ResponseInstance) GetInterfacesOk() (*[]AddInstance200ResponseAllOfOneOfInstanceInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetInstance200ResponseInstance) SetInterfaces(v []AddInstance200ResponseAllOfOneOfInstanceInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetInstance200ResponseInstance) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetCustomOptions

`func (o *GetInstance200ResponseInstance) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *GetInstance200ResponseInstance) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *GetInstance200ResponseInstance) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *GetInstance200ResponseInstance) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetInstanceVersion

`func (o *GetInstance200ResponseInstance) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *GetInstance200ResponseInstance) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *GetInstance200ResponseInstance) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *GetInstance200ResponseInstance) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetLabels

`func (o *GetInstance200ResponseInstance) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetInstance200ResponseInstance) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetInstance200ResponseInstance) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetInstance200ResponseInstance) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetInstance200ResponseInstance) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetInstance200ResponseInstance) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetTags

`func (o *GetInstance200ResponseInstance) GetTags() []AddInstance200ResponseAllOfOneOfInstanceTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetInstance200ResponseInstance) GetTagsOk() (*[]AddInstance200ResponseAllOfOneOfInstanceTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetInstance200ResponseInstance) SetTags(v []AddInstance200ResponseAllOfOneOfInstanceTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetInstance200ResponseInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEvars

`func (o *GetInstance200ResponseInstance) GetEvars() []AddInstance200ResponseAllOfOneOfInstanceEvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *GetInstance200ResponseInstance) GetEvarsOk() (*[]AddInstance200ResponseAllOfOneOfInstanceEvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *GetInstance200ResponseInstance) SetEvars(v []AddInstance200ResponseAllOfOneOfInstanceEvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *GetInstance200ResponseInstance) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetInstance200ResponseInstance) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetInstance200ResponseInstance) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetInstance200ResponseInstance) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetInstance200ResponseInstance) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GetInstance200ResponseInstance) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetInstance200ResponseInstance) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetInstance200ResponseInstance) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetInstance200ResponseInstance) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxCores

`func (o *GetInstance200ResponseInstance) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *GetInstance200ResponseInstance) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *GetInstance200ResponseInstance) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *GetInstance200ResponseInstance) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *GetInstance200ResponseInstance) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *GetInstance200ResponseInstance) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *GetInstance200ResponseInstance) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *GetInstance200ResponseInstance) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### SetCoresPerSocketNil

`func (o *GetInstance200ResponseInstance) SetCoresPerSocketNil(b bool)`

 SetCoresPerSocketNil sets the value for CoresPerSocket to be an explicit nil

### UnsetCoresPerSocket
`func (o *GetInstance200ResponseInstance) UnsetCoresPerSocket()`

UnsetCoresPerSocket ensures that no value is present for CoresPerSocket, not even an explicit nil
### GetMaxCpu

`func (o *GetInstance200ResponseInstance) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GetInstance200ResponseInstance) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GetInstance200ResponseInstance) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GetInstance200ResponseInstance) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *GetInstance200ResponseInstance) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *GetInstance200ResponseInstance) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetHourlyCost

`func (o *GetInstance200ResponseInstance) GetHourlyCost() float32`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *GetInstance200ResponseInstance) GetHourlyCostOk() (*float32, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *GetInstance200ResponseInstance) SetHourlyCost(v float32)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *GetInstance200ResponseInstance) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyPrice

`func (o *GetInstance200ResponseInstance) GetHourlyPrice() float32`

GetHourlyPrice returns the HourlyPrice field if non-nil, zero value otherwise.

### GetHourlyPriceOk

`func (o *GetInstance200ResponseInstance) GetHourlyPriceOk() (*float32, bool)`

GetHourlyPriceOk returns a tuple with the HourlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyPrice

`func (o *GetInstance200ResponseInstance) SetHourlyPrice(v float32)`

SetHourlyPrice sets HourlyPrice field to given value.

### HasHourlyPrice

`func (o *GetInstance200ResponseInstance) HasHourlyPrice() bool`

HasHourlyPrice returns a boolean if a field has been set.

### GetInstancePrice

`func (o *GetInstance200ResponseInstance) GetInstancePrice() GetInstance200ResponseInstanceInstancePrice`

GetInstancePrice returns the InstancePrice field if non-nil, zero value otherwise.

### GetInstancePriceOk

`func (o *GetInstance200ResponseInstance) GetInstancePriceOk() (*GetInstance200ResponseInstanceInstancePrice, bool)`

GetInstancePriceOk returns a tuple with the InstancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePrice

`func (o *GetInstance200ResponseInstance) SetInstancePrice(v GetInstance200ResponseInstanceInstancePrice)`

SetInstancePrice sets InstancePrice field to given value.

### HasInstancePrice

`func (o *GetInstance200ResponseInstance) HasInstancePrice() bool`

HasInstancePrice returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetInstance200ResponseInstance) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetInstance200ResponseInstance) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetInstance200ResponseInstance) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetInstance200ResponseInstance) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetInstance200ResponseInstance) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetInstance200ResponseInstance) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetInstance200ResponseInstance) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetInstance200ResponseInstance) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetHostName

`func (o *GetInstance200ResponseInstance) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *GetInstance200ResponseInstance) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *GetInstance200ResponseInstance) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *GetInstance200ResponseInstance) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetDomainName

`func (o *GetInstance200ResponseInstance) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *GetInstance200ResponseInstance) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *GetInstance200ResponseInstance) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *GetInstance200ResponseInstance) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *GetInstance200ResponseInstance) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *GetInstance200ResponseInstance) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetEnvironmentPrefix

`func (o *GetInstance200ResponseInstance) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *GetInstance200ResponseInstance) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *GetInstance200ResponseInstance) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *GetInstance200ResponseInstance) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### SetEnvironmentPrefixNil

`func (o *GetInstance200ResponseInstance) SetEnvironmentPrefixNil(b bool)`

 SetEnvironmentPrefixNil sets the value for EnvironmentPrefix to be an explicit nil

### UnsetEnvironmentPrefix
`func (o *GetInstance200ResponseInstance) UnsetEnvironmentPrefix()`

UnsetEnvironmentPrefix ensures that no value is present for EnvironmentPrefix, not even an explicit nil
### GetFirewallEnabled

`func (o *GetInstance200ResponseInstance) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *GetInstance200ResponseInstance) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *GetInstance200ResponseInstance) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *GetInstance200ResponseInstance) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetNetworkLevel

`func (o *GetInstance200ResponseInstance) GetNetworkLevel() string`

GetNetworkLevel returns the NetworkLevel field if non-nil, zero value otherwise.

### GetNetworkLevelOk

`func (o *GetInstance200ResponseInstance) GetNetworkLevelOk() (*string, bool)`

GetNetworkLevelOk returns a tuple with the NetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLevel

`func (o *GetInstance200ResponseInstance) SetNetworkLevel(v string)`

SetNetworkLevel sets NetworkLevel field to given value.

### HasNetworkLevel

`func (o *GetInstance200ResponseInstance) HasNetworkLevel() bool`

HasNetworkLevel returns a boolean if a field has been set.

### GetAutoScale

`func (o *GetInstance200ResponseInstance) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *GetInstance200ResponseInstance) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *GetInstance200ResponseInstance) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *GetInstance200ResponseInstance) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetInstanceContext

`func (o *GetInstance200ResponseInstance) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *GetInstance200ResponseInstance) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *GetInstance200ResponseInstance) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *GetInstance200ResponseInstance) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### SetInstanceContextNil

`func (o *GetInstance200ResponseInstance) SetInstanceContextNil(b bool)`

 SetInstanceContextNil sets the value for InstanceContext to be an explicit nil

### UnsetInstanceContext
`func (o *GetInstance200ResponseInstance) UnsetInstanceContext()`

UnsetInstanceContext ensures that no value is present for InstanceContext, not even an explicit nil
### GetCurrentDeployId

`func (o *GetInstance200ResponseInstance) GetCurrentDeployId() string`

GetCurrentDeployId returns the CurrentDeployId field if non-nil, zero value otherwise.

### GetCurrentDeployIdOk

`func (o *GetInstance200ResponseInstance) GetCurrentDeployIdOk() (*string, bool)`

GetCurrentDeployIdOk returns a tuple with the CurrentDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeployId

`func (o *GetInstance200ResponseInstance) SetCurrentDeployId(v string)`

SetCurrentDeployId sets CurrentDeployId field to given value.

### HasCurrentDeployId

`func (o *GetInstance200ResponseInstance) HasCurrentDeployId() bool`

HasCurrentDeployId returns a boolean if a field has been set.

### SetCurrentDeployIdNil

`func (o *GetInstance200ResponseInstance) SetCurrentDeployIdNil(b bool)`

 SetCurrentDeployIdNil sets the value for CurrentDeployId to be an explicit nil

### UnsetCurrentDeployId
`func (o *GetInstance200ResponseInstance) UnsetCurrentDeployId()`

UnsetCurrentDeployId ensures that no value is present for CurrentDeployId, not even an explicit nil
### GetLocked

`func (o *GetInstance200ResponseInstance) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *GetInstance200ResponseInstance) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *GetInstance200ResponseInstance) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *GetInstance200ResponseInstance) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetStatus

`func (o *GetInstance200ResponseInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInstance200ResponseInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInstance200ResponseInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInstance200ResponseInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetInstance200ResponseInstance) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetInstance200ResponseInstance) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetInstance200ResponseInstance) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetInstance200ResponseInstance) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetInstance200ResponseInstance) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetInstance200ResponseInstance) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetErrorMessage

`func (o *GetInstance200ResponseInstance) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetInstance200ResponseInstance) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetInstance200ResponseInstance) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetInstance200ResponseInstance) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetInstance200ResponseInstance) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetInstance200ResponseInstance) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStatusDate

`func (o *GetInstance200ResponseInstance) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetInstance200ResponseInstance) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetInstance200ResponseInstance) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetInstance200ResponseInstance) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *GetInstance200ResponseInstance) GetStatusPercent() string`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *GetInstance200ResponseInstance) GetStatusPercentOk() (*string, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *GetInstance200ResponseInstance) SetStatusPercent(v string)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *GetInstance200ResponseInstance) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### SetStatusPercentNil

`func (o *GetInstance200ResponseInstance) SetStatusPercentNil(b bool)`

 SetStatusPercentNil sets the value for StatusPercent to be an explicit nil

### UnsetStatusPercent
`func (o *GetInstance200ResponseInstance) UnsetStatusPercent()`

UnsetStatusPercent ensures that no value is present for StatusPercent, not even an explicit nil
### GetStatusEta

`func (o *GetInstance200ResponseInstance) GetStatusEta() string`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetInstance200ResponseInstance) GetStatusEtaOk() (*string, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetInstance200ResponseInstance) SetStatusEta(v string)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetInstance200ResponseInstance) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### SetStatusEtaNil

`func (o *GetInstance200ResponseInstance) SetStatusEtaNil(b bool)`

 SetStatusEtaNil sets the value for StatusEta to be an explicit nil

### UnsetStatusEta
`func (o *GetInstance200ResponseInstance) UnsetStatusEta()`

UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
### GetUserStatus

`func (o *GetInstance200ResponseInstance) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *GetInstance200ResponseInstance) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *GetInstance200ResponseInstance) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *GetInstance200ResponseInstance) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### SetUserStatusNil

`func (o *GetInstance200ResponseInstance) SetUserStatusNil(b bool)`

 SetUserStatusNil sets the value for UserStatus to be an explicit nil

### UnsetUserStatus
`func (o *GetInstance200ResponseInstance) UnsetUserStatus()`

UnsetUserStatus ensures that no value is present for UserStatus, not even an explicit nil
### GetExpireDays

`func (o *GetInstance200ResponseInstance) GetExpireDays() int64`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *GetInstance200ResponseInstance) GetExpireDaysOk() (*int64, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *GetInstance200ResponseInstance) SetExpireDays(v int64)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *GetInstance200ResponseInstance) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetRenewDays

`func (o *GetInstance200ResponseInstance) GetRenewDays() int64`

GetRenewDays returns the RenewDays field if non-nil, zero value otherwise.

### GetRenewDaysOk

`func (o *GetInstance200ResponseInstance) GetRenewDaysOk() (*int64, bool)`

GetRenewDaysOk returns a tuple with the RenewDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewDays

`func (o *GetInstance200ResponseInstance) SetRenewDays(v int64)`

SetRenewDays sets RenewDays field to given value.

### HasRenewDays

`func (o *GetInstance200ResponseInstance) HasRenewDays() bool`

HasRenewDays returns a boolean if a field has been set.

### GetExpireCount

`func (o *GetInstance200ResponseInstance) GetExpireCount() int64`

GetExpireCount returns the ExpireCount field if non-nil, zero value otherwise.

### GetExpireCountOk

`func (o *GetInstance200ResponseInstance) GetExpireCountOk() (*int64, bool)`

GetExpireCountOk returns a tuple with the ExpireCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireCount

`func (o *GetInstance200ResponseInstance) SetExpireCount(v int64)`

SetExpireCount sets ExpireCount field to given value.

### HasExpireCount

`func (o *GetInstance200ResponseInstance) HasExpireCount() bool`

HasExpireCount returns a boolean if a field has been set.

### GetExpireDate

`func (o *GetInstance200ResponseInstance) GetExpireDate() time.Time`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *GetInstance200ResponseInstance) GetExpireDateOk() (*time.Time, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *GetInstance200ResponseInstance) SetExpireDate(v time.Time)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *GetInstance200ResponseInstance) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetExpireWarningDate

`func (o *GetInstance200ResponseInstance) GetExpireWarningDate() time.Time`

GetExpireWarningDate returns the ExpireWarningDate field if non-nil, zero value otherwise.

### GetExpireWarningDateOk

`func (o *GetInstance200ResponseInstance) GetExpireWarningDateOk() (*time.Time, bool)`

GetExpireWarningDateOk returns a tuple with the ExpireWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWarningDate

`func (o *GetInstance200ResponseInstance) SetExpireWarningDate(v time.Time)`

SetExpireWarningDate sets ExpireWarningDate field to given value.

### HasExpireWarningDate

`func (o *GetInstance200ResponseInstance) HasExpireWarningDate() bool`

HasExpireWarningDate returns a boolean if a field has been set.

### GetExpireWarningSent

`func (o *GetInstance200ResponseInstance) GetExpireWarningSent() bool`

GetExpireWarningSent returns the ExpireWarningSent field if non-nil, zero value otherwise.

### GetExpireWarningSentOk

`func (o *GetInstance200ResponseInstance) GetExpireWarningSentOk() (*bool, bool)`

GetExpireWarningSentOk returns a tuple with the ExpireWarningSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWarningSent

`func (o *GetInstance200ResponseInstance) SetExpireWarningSent(v bool)`

SetExpireWarningSent sets ExpireWarningSent field to given value.

### HasExpireWarningSent

`func (o *GetInstance200ResponseInstance) HasExpireWarningSent() bool`

HasExpireWarningSent returns a boolean if a field has been set.

### GetShutdownDays

`func (o *GetInstance200ResponseInstance) GetShutdownDays() int64`

GetShutdownDays returns the ShutdownDays field if non-nil, zero value otherwise.

### GetShutdownDaysOk

`func (o *GetInstance200ResponseInstance) GetShutdownDaysOk() (*int64, bool)`

GetShutdownDaysOk returns a tuple with the ShutdownDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDays

`func (o *GetInstance200ResponseInstance) SetShutdownDays(v int64)`

SetShutdownDays sets ShutdownDays field to given value.

### HasShutdownDays

`func (o *GetInstance200ResponseInstance) HasShutdownDays() bool`

HasShutdownDays returns a boolean if a field has been set.

### GetShutdownRenewDays

`func (o *GetInstance200ResponseInstance) GetShutdownRenewDays() int64`

GetShutdownRenewDays returns the ShutdownRenewDays field if non-nil, zero value otherwise.

### GetShutdownRenewDaysOk

`func (o *GetInstance200ResponseInstance) GetShutdownRenewDaysOk() (*int64, bool)`

GetShutdownRenewDaysOk returns a tuple with the ShutdownRenewDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownRenewDays

`func (o *GetInstance200ResponseInstance) SetShutdownRenewDays(v int64)`

SetShutdownRenewDays sets ShutdownRenewDays field to given value.

### HasShutdownRenewDays

`func (o *GetInstance200ResponseInstance) HasShutdownRenewDays() bool`

HasShutdownRenewDays returns a boolean if a field has been set.

### GetShutdownCount

`func (o *GetInstance200ResponseInstance) GetShutdownCount() int64`

GetShutdownCount returns the ShutdownCount field if non-nil, zero value otherwise.

### GetShutdownCountOk

`func (o *GetInstance200ResponseInstance) GetShutdownCountOk() (*int64, bool)`

GetShutdownCountOk returns a tuple with the ShutdownCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownCount

`func (o *GetInstance200ResponseInstance) SetShutdownCount(v int64)`

SetShutdownCount sets ShutdownCount field to given value.

### HasShutdownCount

`func (o *GetInstance200ResponseInstance) HasShutdownCount() bool`

HasShutdownCount returns a boolean if a field has been set.

### GetShutdownDate

`func (o *GetInstance200ResponseInstance) GetShutdownDate() time.Time`

GetShutdownDate returns the ShutdownDate field if non-nil, zero value otherwise.

### GetShutdownDateOk

`func (o *GetInstance200ResponseInstance) GetShutdownDateOk() (*time.Time, bool)`

GetShutdownDateOk returns a tuple with the ShutdownDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDate

`func (o *GetInstance200ResponseInstance) SetShutdownDate(v time.Time)`

SetShutdownDate sets ShutdownDate field to given value.

### HasShutdownDate

`func (o *GetInstance200ResponseInstance) HasShutdownDate() bool`

HasShutdownDate returns a boolean if a field has been set.

### GetShutdownWarningDate

`func (o *GetInstance200ResponseInstance) GetShutdownWarningDate() time.Time`

GetShutdownWarningDate returns the ShutdownWarningDate field if non-nil, zero value otherwise.

### GetShutdownWarningDateOk

`func (o *GetInstance200ResponseInstance) GetShutdownWarningDateOk() (*time.Time, bool)`

GetShutdownWarningDateOk returns a tuple with the ShutdownWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWarningDate

`func (o *GetInstance200ResponseInstance) SetShutdownWarningDate(v time.Time)`

SetShutdownWarningDate sets ShutdownWarningDate field to given value.

### HasShutdownWarningDate

`func (o *GetInstance200ResponseInstance) HasShutdownWarningDate() bool`

HasShutdownWarningDate returns a boolean if a field has been set.

### GetShutdownWarningSent

`func (o *GetInstance200ResponseInstance) GetShutdownWarningSent() bool`

GetShutdownWarningSent returns the ShutdownWarningSent field if non-nil, zero value otherwise.

### GetShutdownWarningSentOk

`func (o *GetInstance200ResponseInstance) GetShutdownWarningSentOk() (*bool, bool)`

GetShutdownWarningSentOk returns a tuple with the ShutdownWarningSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownWarningSent

`func (o *GetInstance200ResponseInstance) SetShutdownWarningSent(v bool)`

SetShutdownWarningSent sets ShutdownWarningSent field to given value.

### HasShutdownWarningSent

`func (o *GetInstance200ResponseInstance) HasShutdownWarningSent() bool`

HasShutdownWarningSent returns a boolean if a field has been set.

### GetRemovalDate

`func (o *GetInstance200ResponseInstance) GetRemovalDate() time.Time`

GetRemovalDate returns the RemovalDate field if non-nil, zero value otherwise.

### GetRemovalDateOk

`func (o *GetInstance200ResponseInstance) GetRemovalDateOk() (*time.Time, bool)`

GetRemovalDateOk returns a tuple with the RemovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalDate

`func (o *GetInstance200ResponseInstance) SetRemovalDate(v time.Time)`

SetRemovalDate sets RemovalDate field to given value.

### HasRemovalDate

`func (o *GetInstance200ResponseInstance) HasRemovalDate() bool`

HasRemovalDate returns a boolean if a field has been set.

### SetRemovalDateNil

`func (o *GetInstance200ResponseInstance) SetRemovalDateNil(b bool)`

 SetRemovalDateNil sets the value for RemovalDate to be an explicit nil

### UnsetRemovalDate
`func (o *GetInstance200ResponseInstance) UnsetRemovalDate()`

UnsetRemovalDate ensures that no value is present for RemovalDate, not even an explicit nil
### GetCreatedBy

`func (o *GetInstance200ResponseInstance) GetCreatedBy() GetInstance200ResponseInstanceCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetInstance200ResponseInstance) GetCreatedByOk() (*GetInstance200ResponseInstanceCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetInstance200ResponseInstance) SetCreatedBy(v GetInstance200ResponseInstanceCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetInstance200ResponseInstance) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetOwner

`func (o *GetInstance200ResponseInstance) GetOwner() GetInstance200ResponseInstanceOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetInstance200ResponseInstance) GetOwnerOk() (*GetInstance200ResponseInstanceOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetInstance200ResponseInstance) SetOwner(v GetInstance200ResponseInstanceOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetInstance200ResponseInstance) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNotes

`func (o *GetInstance200ResponseInstance) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetInstance200ResponseInstance) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetInstance200ResponseInstance) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetInstance200ResponseInstance) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *GetInstance200ResponseInstance) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *GetInstance200ResponseInstance) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetStats

`func (o *GetInstance200ResponseInstance) GetStats() GetInstance200ResponseInstanceStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetInstance200ResponseInstance) GetStatsOk() (*GetInstance200ResponseInstanceStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetInstance200ResponseInstance) SetStats(v GetInstance200ResponseInstanceStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetInstance200ResponseInstance) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetPowerSchedule

`func (o *GetInstance200ResponseInstance) GetPowerSchedule() string`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *GetInstance200ResponseInstance) GetPowerScheduleOk() (*string, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *GetInstance200ResponseInstance) SetPowerSchedule(v string)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *GetInstance200ResponseInstance) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### SetPowerScheduleNil

`func (o *GetInstance200ResponseInstance) SetPowerScheduleNil(b bool)`

 SetPowerScheduleNil sets the value for PowerSchedule to be an explicit nil

### UnsetPowerSchedule
`func (o *GetInstance200ResponseInstance) UnsetPowerSchedule()`

UnsetPowerSchedule ensures that no value is present for PowerSchedule, not even an explicit nil
### GetIsScalable

`func (o *GetInstance200ResponseInstance) GetIsScalable() bool`

GetIsScalable returns the IsScalable field if non-nil, zero value otherwise.

### GetIsScalableOk

`func (o *GetInstance200ResponseInstance) GetIsScalableOk() (*bool, bool)`

GetIsScalableOk returns a tuple with the IsScalable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScalable

`func (o *GetInstance200ResponseInstance) SetIsScalable(v bool)`

SetIsScalable sets IsScalable field to given value.

### HasIsScalable

`func (o *GetInstance200ResponseInstance) HasIsScalable() bool`

HasIsScalable returns a boolean if a field has been set.

### GetInstanceThreshold

`func (o *GetInstance200ResponseInstance) GetInstanceThreshold() map[string]interface{}`

GetInstanceThreshold returns the InstanceThreshold field if non-nil, zero value otherwise.

### GetInstanceThresholdOk

`func (o *GetInstance200ResponseInstance) GetInstanceThresholdOk() (*map[string]interface{}, bool)`

GetInstanceThresholdOk returns a tuple with the InstanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceThreshold

`func (o *GetInstance200ResponseInstance) SetInstanceThreshold(v map[string]interface{})`

SetInstanceThreshold sets InstanceThreshold field to given value.

### HasInstanceThreshold

`func (o *GetInstance200ResponseInstance) HasInstanceThreshold() bool`

HasInstanceThreshold returns a boolean if a field has been set.

### SetInstanceThresholdNil

`func (o *GetInstance200ResponseInstance) SetInstanceThresholdNil(b bool)`

 SetInstanceThresholdNil sets the value for InstanceThreshold to be an explicit nil

### UnsetInstanceThreshold
`func (o *GetInstance200ResponseInstance) UnsetInstanceThreshold()`

UnsetInstanceThreshold ensures that no value is present for InstanceThreshold, not even an explicit nil
### GetIsBusy

`func (o *GetInstance200ResponseInstance) GetIsBusy() bool`

GetIsBusy returns the IsBusy field if non-nil, zero value otherwise.

### GetIsBusyOk

`func (o *GetInstance200ResponseInstance) GetIsBusyOk() (*bool, bool)`

GetIsBusyOk returns a tuple with the IsBusy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBusy

`func (o *GetInstance200ResponseInstance) SetIsBusy(v bool)`

SetIsBusy sets IsBusy field to given value.

### HasIsBusy

`func (o *GetInstance200ResponseInstance) HasIsBusy() bool`

HasIsBusy returns a boolean if a field has been set.

### GetApps

`func (o *GetInstance200ResponseInstance) GetApps() []map[string]interface{}`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *GetInstance200ResponseInstance) GetAppsOk() (*[]map[string]interface{}, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *GetInstance200ResponseInstance) SetApps(v []map[string]interface{})`

SetApps sets Apps field to given value.

### HasApps

`func (o *GetInstance200ResponseInstance) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *GetInstance200ResponseInstance) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *GetInstance200ResponseInstance) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetCurrentLoadBalancerInstances

`func (o *GetInstance200ResponseInstance) GetCurrentLoadBalancerInstances() []map[string]interface{}`

GetCurrentLoadBalancerInstances returns the CurrentLoadBalancerInstances field if non-nil, zero value otherwise.

### GetCurrentLoadBalancerInstancesOk

`func (o *GetInstance200ResponseInstance) GetCurrentLoadBalancerInstancesOk() (*[]map[string]interface{}, bool)`

GetCurrentLoadBalancerInstancesOk returns a tuple with the CurrentLoadBalancerInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLoadBalancerInstances

`func (o *GetInstance200ResponseInstance) SetCurrentLoadBalancerInstances(v []map[string]interface{})`

SetCurrentLoadBalancerInstances sets CurrentLoadBalancerInstances field to given value.

### HasCurrentLoadBalancerInstances

`func (o *GetInstance200ResponseInstance) HasCurrentLoadBalancerInstances() bool`

HasCurrentLoadBalancerInstances returns a boolean if a field has been set.

### SetCurrentLoadBalancerInstancesNil

`func (o *GetInstance200ResponseInstance) SetCurrentLoadBalancerInstancesNil(b bool)`

 SetCurrentLoadBalancerInstancesNil sets the value for CurrentLoadBalancerInstances to be an explicit nil

### UnsetCurrentLoadBalancerInstances
`func (o *GetInstance200ResponseInstance) UnsetCurrentLoadBalancerInstances()`

UnsetCurrentLoadBalancerInstances ensures that no value is present for CurrentLoadBalancerInstances, not even an explicit nil
### GetCurrentLoadBalancerContainersIn

`func (o *GetInstance200ResponseInstance) GetCurrentLoadBalancerContainersIn() int32`

GetCurrentLoadBalancerContainersIn returns the CurrentLoadBalancerContainersIn field if non-nil, zero value otherwise.

### GetCurrentLoadBalancerContainersInOk

`func (o *GetInstance200ResponseInstance) GetCurrentLoadBalancerContainersInOk() (*int32, bool)`

GetCurrentLoadBalancerContainersInOk returns a tuple with the CurrentLoadBalancerContainersIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLoadBalancerContainersIn

`func (o *GetInstance200ResponseInstance) SetCurrentLoadBalancerContainersIn(v int32)`

SetCurrentLoadBalancerContainersIn sets CurrentLoadBalancerContainersIn field to given value.

### HasCurrentLoadBalancerContainersIn

`func (o *GetInstance200ResponseInstance) HasCurrentLoadBalancerContainersIn() bool`

HasCurrentLoadBalancerContainersIn returns a boolean if a field has been set.

### GetCurrentLoadBalancerContainersOut

`func (o *GetInstance200ResponseInstance) GetCurrentLoadBalancerContainersOut() int32`

GetCurrentLoadBalancerContainersOut returns the CurrentLoadBalancerContainersOut field if non-nil, zero value otherwise.

### GetCurrentLoadBalancerContainersOutOk

`func (o *GetInstance200ResponseInstance) GetCurrentLoadBalancerContainersOutOk() (*int32, bool)`

GetCurrentLoadBalancerContainersOutOk returns a tuple with the CurrentLoadBalancerContainersOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLoadBalancerContainersOut

`func (o *GetInstance200ResponseInstance) SetCurrentLoadBalancerContainersOut(v int32)`

SetCurrentLoadBalancerContainersOut sets CurrentLoadBalancerContainersOut field to given value.

### HasCurrentLoadBalancerContainersOut

`func (o *GetInstance200ResponseInstance) HasCurrentLoadBalancerContainersOut() bool`

HasCurrentLoadBalancerContainersOut returns a boolean if a field has been set.

### GetLastDeploy

`func (o *GetInstance200ResponseInstance) GetLastDeploy() map[string]interface{}`

GetLastDeploy returns the LastDeploy field if non-nil, zero value otherwise.

### GetLastDeployOk

`func (o *GetInstance200ResponseInstance) GetLastDeployOk() (*map[string]interface{}, bool)`

GetLastDeployOk returns a tuple with the LastDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploy

`func (o *GetInstance200ResponseInstance) SetLastDeploy(v map[string]interface{})`

SetLastDeploy sets LastDeploy field to given value.

### HasLastDeploy

`func (o *GetInstance200ResponseInstance) HasLastDeploy() bool`

HasLastDeploy returns a boolean if a field has been set.

### GetContainerDetails

`func (o *GetInstance200ResponseInstance) GetContainerDetails() []InstanceContainer2`

GetContainerDetails returns the ContainerDetails field if non-nil, zero value otherwise.

### GetContainerDetailsOk

`func (o *GetInstance200ResponseInstance) GetContainerDetailsOk() (*[]InstanceContainer2, bool)`

GetContainerDetailsOk returns a tuple with the ContainerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDetails

`func (o *GetInstance200ResponseInstance) SetContainerDetails(v []InstanceContainer2)`

SetContainerDetails sets ContainerDetails field to given value.

### HasContainerDetails

`func (o *GetInstance200ResponseInstance) HasContainerDetails() bool`

HasContainerDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



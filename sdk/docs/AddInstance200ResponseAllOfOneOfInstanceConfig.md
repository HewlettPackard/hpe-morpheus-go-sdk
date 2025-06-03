# AddInstance200ResponseAllOfOneOfInstanceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateUser** | Pointer to **bool** |  | [optional] 
**IsEC2** | Pointer to **bool** |  | [optional] 
**IsVpcSelectable** | Pointer to **bool** |  | [optional] 
**NoAgent** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceConfigNoAgent**](AddInstance200ResponseAllOfOneOfInstanceConfigNoAgent.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]ListInstances200ResponseAllOfInstancesInnerConfigUserGroup**](ListInstances200ResponseAllOfInstancesInnerConfigUserGroup.md) |  | [optional] 
**SmbiosAssetTag** | Pointer to **string** |  | [optional] 
**NestedVirtualization** | Pointer to **string** |  | [optional] 
**VmwareFolderId** | Pointer to **string** |  | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**ResourcePoolId** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId**](AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId.md) |  | [optional] 
**PoolProviderType** | Pointer to **string** |  | [optional] 
**UserGroup** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerConfigUserGroup**](ListInstances200ResponseAllOfInstancesInnerConfigUserGroup.md) |  | [optional] 
**ExpireDays** | Pointer to **string** |  | [optional] 
**ShutdownDays** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceConfigInstanceType**](AddInstance200ResponseAllOfOneOfInstanceConfigInstanceType.md) |  | [optional] 
**Site** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**EnvironmentPrefix** | Pointer to **string** |  | [optional] 
**Layout** | Pointer to [**ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan**](ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**InstanceContext** | Pointer to **string** |  | [optional] 
**MemoryDisplay** | Pointer to **string** |  | [optional] 
**Expose** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CreateBackup** | Pointer to **bool** |  | [optional] 
**Backup** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerConfigBackup**](ListInstances200ResponseAllOfInstancesInnerConfigBackup.md) |  | [optional] 
**ReplicationGroup** | Pointer to [**AddInstance200ResponseAllOfOneOfInstanceConfigReplicationGroup**](AddInstance200ResponseAllOfOneOfInstanceConfigReplicationGroup.md) |  | [optional] 
**LayoutSize** | Pointer to **int64** |  | [optional] 
**LbInstances** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewAddInstance200ResponseAllOfOneOfInstanceConfig

`func NewAddInstance200ResponseAllOfOneOfInstanceConfig() *AddInstance200ResponseAllOfOneOfInstanceConfig`

NewAddInstance200ResponseAllOfOneOfInstanceConfig instantiates a new AddInstance200ResponseAllOfOneOfInstanceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstance200ResponseAllOfOneOfInstanceConfigWithDefaults

`func NewAddInstance200ResponseAllOfOneOfInstanceConfigWithDefaults() *AddInstance200ResponseAllOfOneOfInstanceConfig`

NewAddInstance200ResponseAllOfOneOfInstanceConfigWithDefaults instantiates a new AddInstance200ResponseAllOfOneOfInstanceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateUser

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetIsEC2

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetIsEC2() bool`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetIsEC2Ok() (*bool, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetIsEC2(v bool)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetIsVpcSelectable

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetIsVpcSelectable() bool`

GetIsVpcSelectable returns the IsVpcSelectable field if non-nil, zero value otherwise.

### GetIsVpcSelectableOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetIsVpcSelectableOk() (*bool, bool)`

GetIsVpcSelectableOk returns a tuple with the IsVpcSelectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpcSelectable

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetIsVpcSelectable(v bool)`

SetIsVpcSelectable sets IsVpcSelectable field to given value.

### HasIsVpcSelectable

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasIsVpcSelectable() bool`

HasIsVpcSelectable returns a boolean if a field has been set.

### GetNoAgent

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetNoAgent() AddInstance200ResponseAllOfOneOfInstanceConfigNoAgent`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetNoAgentOk() (*AddInstance200ResponseAllOfOneOfInstanceConfigNoAgent, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetNoAgent(v AddInstance200ResponseAllOfOneOfInstanceConfigNoAgent)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetSecurityGroups() []ListInstances200ResponseAllOfInstancesInnerConfigUserGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetSecurityGroupsOk() (*[]ListInstances200ResponseAllOfInstancesInnerConfigUserGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetSecurityGroups(v []ListInstances200ResponseAllOfInstancesInnerConfigUserGroup)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetSmbiosAssetTag

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetCustomOptions

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetResourcePoolId() AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetResourcePoolIdOk() (*AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetResourcePoolId(v AddInstance200ResponseAllOfOneOfInstanceConfigResourcePoolId)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetPoolProviderType

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetPoolProviderType() string`

GetPoolProviderType returns the PoolProviderType field if non-nil, zero value otherwise.

### GetPoolProviderTypeOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetPoolProviderTypeOk() (*string, bool)`

GetPoolProviderTypeOk returns a tuple with the PoolProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolProviderType

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetPoolProviderType(v string)`

SetPoolProviderType sets PoolProviderType field to given value.

### HasPoolProviderType

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasPoolProviderType() bool`

HasPoolProviderType returns a boolean if a field has been set.

### GetUserGroup

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetUserGroup() ListInstances200ResponseAllOfInstancesInnerConfigUserGroup`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetUserGroupOk() (*ListInstances200ResponseAllOfInstancesInnerConfigUserGroup, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetUserGroup(v ListInstances200ResponseAllOfInstancesInnerConfigUserGroup)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetExpireDays

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetExpireDays() string`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetExpireDaysOk() (*string, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetExpireDays(v string)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetShutdownDays

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetShutdownDays() string`

GetShutdownDays returns the ShutdownDays field if non-nil, zero value otherwise.

### GetShutdownDaysOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetShutdownDaysOk() (*string, bool)`

GetShutdownDaysOk returns a tuple with the ShutdownDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDays

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetShutdownDays(v string)`

SetShutdownDays sets ShutdownDays field to given value.

### HasShutdownDays

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasShutdownDays() bool`

HasShutdownDays returns a boolean if a field has been set.

### GetName

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostName

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetInstanceType

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetInstanceType() AddInstance200ResponseAllOfOneOfInstanceConfigInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetInstanceTypeOk() (*AddInstance200ResponseAllOfOneOfInstanceConfigInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetInstanceType(v AddInstance200ResponseAllOfOneOfInstanceConfigInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetSite

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetSite() GetAlerts200ResponseAllOfChecksInnerAccount`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetSiteOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetSite(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetSite sets Site field to given value.

### HasSite

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetLayout

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetLayout() ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetLayoutOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetLayout(v ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetType

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInstanceContext

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### GetMemoryDisplay

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetMemoryDisplay() string`

GetMemoryDisplay returns the MemoryDisplay field if non-nil, zero value otherwise.

### GetMemoryDisplayOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetMemoryDisplayOk() (*string, bool)`

GetMemoryDisplayOk returns a tuple with the MemoryDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDisplay

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetMemoryDisplay(v string)`

SetMemoryDisplay sets MemoryDisplay field to given value.

### HasMemoryDisplay

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasMemoryDisplay() bool`

HasMemoryDisplay returns a boolean if a field has been set.

### GetExpose

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetExpose() []map[string]interface{}`

GetExpose returns the Expose field if non-nil, zero value otherwise.

### GetExposeOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetExposeOk() (*[]map[string]interface{}, bool)`

GetExposeOk returns a tuple with the Expose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpose

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetExpose(v []map[string]interface{})`

SetExpose sets Expose field to given value.

### HasExpose

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasExpose() bool`

HasExpose returns a boolean if a field has been set.

### GetCreateBackup

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.

### GetBackup

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetBackup() ListInstances200ResponseAllOfInstancesInnerConfigBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetBackupOk() (*ListInstances200ResponseAllOfInstancesInnerConfigBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetBackup(v ListInstances200ResponseAllOfInstancesInnerConfigBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetReplicationGroup

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetReplicationGroup() AddInstance200ResponseAllOfOneOfInstanceConfigReplicationGroup`

GetReplicationGroup returns the ReplicationGroup field if non-nil, zero value otherwise.

### GetReplicationGroupOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetReplicationGroupOk() (*AddInstance200ResponseAllOfOneOfInstanceConfigReplicationGroup, bool)`

GetReplicationGroupOk returns a tuple with the ReplicationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroup

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetReplicationGroup(v AddInstance200ResponseAllOfOneOfInstanceConfigReplicationGroup)`

SetReplicationGroup sets ReplicationGroup field to given value.

### HasReplicationGroup

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasReplicationGroup() bool`

HasReplicationGroup returns a boolean if a field has been set.

### GetLayoutSize

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetLayoutSize() int64`

GetLayoutSize returns the LayoutSize field if non-nil, zero value otherwise.

### GetLayoutSizeOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetLayoutSizeOk() (*int64, bool)`

GetLayoutSizeOk returns a tuple with the LayoutSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSize

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetLayoutSize(v int64)`

SetLayoutSize sets LayoutSize field to given value.

### HasLayoutSize

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasLayoutSize() bool`

HasLayoutSize returns a boolean if a field has been set.

### GetLbInstances

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetLbInstances() []map[string]interface{}`

GetLbInstances returns the LbInstances field if non-nil, zero value otherwise.

### GetLbInstancesOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) GetLbInstancesOk() (*[]map[string]interface{}, bool)`

GetLbInstancesOk returns a tuple with the LbInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbInstances

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) SetLbInstances(v []map[string]interface{})`

SetLbInstances sets LbInstances field to given value.

### HasLbInstances

`func (o *AddInstance200ResponseAllOfOneOfInstanceConfig) HasLbInstances() bool`

HasLbInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



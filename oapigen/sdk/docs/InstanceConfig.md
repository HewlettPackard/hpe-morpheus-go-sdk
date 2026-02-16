# InstanceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateUser** | Pointer to **bool** |  | [optional] 
**IsEC2** | Pointer to **bool** |  | [optional] 
**IsVpcSelectable** | Pointer to **bool** |  | [optional] 
**NoAgent** | Pointer to [**InstanceConfigNoAgent**](InstanceConfigNoAgent.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner**](AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner.md) |  | [optional] 
**KvmHostId** | Pointer to **NullableInt64** |  | [optional] 
**SmbiosAssetTag** | Pointer to **NullableString** |  | [optional] 
**NestedVirtualization** | Pointer to **NullableString** |  | [optional] 
**VmwareFolderId** | Pointer to **string** |  | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**ResourcePoolId** | Pointer to [**InstanceConfigResourcePoolId**](InstanceConfigResourcePoolId.md) |  | [optional] 
**PoolProviderType** | Pointer to **NullableString** |  | [optional] 
**UserGroup** | Pointer to [**InstanceConfigUserGroup**](InstanceConfigUserGroup.md) |  | [optional] 
**ExpireDays** | Pointer to **string** |  | [optional] 
**ShutdownDays** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to [**InstanceConfigInstanceType**](InstanceConfigInstanceType.md) |  | [optional] 
**Site** | Pointer to [**InstanceConfigSite**](InstanceConfigSite.md) |  | [optional] 
**EnvironmentPrefix** | Pointer to **NullableString** |  | [optional] 
**Layout** | Pointer to [**InstanceConfigLayout**](InstanceConfigLayout.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**InstanceContext** | Pointer to **string** |  | [optional] 
**MemoryDisplay** | Pointer to **string** |  | [optional] 
**Expose** | Pointer to **[]int64** |  | [optional] 
**CreateBackup** | Pointer to **bool** |  | [optional] 
**Backup** | Pointer to [**InstanceConfigBackup**](InstanceConfigBackup.md) |  | [optional] 
**ReplicationGroup** | Pointer to [**InstanceConfigReplicationGroup**](InstanceConfigReplicationGroup.md) |  | [optional] 
**LayoutSize** | Pointer to **int64** |  | [optional] 
**LbInstances** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewInstanceConfig

`func NewInstanceConfig() *InstanceConfig`

NewInstanceConfig instantiates a new InstanceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConfigWithDefaults

`func NewInstanceConfigWithDefaults() *InstanceConfig`

NewInstanceConfigWithDefaults instantiates a new InstanceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateUser

`func (o *InstanceConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *InstanceConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *InstanceConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *InstanceConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetIsEC2

`func (o *InstanceConfig) GetIsEC2() bool`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *InstanceConfig) GetIsEC2Ok() (*bool, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *InstanceConfig) SetIsEC2(v bool)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *InstanceConfig) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetIsVpcSelectable

`func (o *InstanceConfig) GetIsVpcSelectable() bool`

GetIsVpcSelectable returns the IsVpcSelectable field if non-nil, zero value otherwise.

### GetIsVpcSelectableOk

`func (o *InstanceConfig) GetIsVpcSelectableOk() (*bool, bool)`

GetIsVpcSelectableOk returns a tuple with the IsVpcSelectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpcSelectable

`func (o *InstanceConfig) SetIsVpcSelectable(v bool)`

SetIsVpcSelectable sets IsVpcSelectable field to given value.

### HasIsVpcSelectable

`func (o *InstanceConfig) HasIsVpcSelectable() bool`

HasIsVpcSelectable returns a boolean if a field has been set.

### GetNoAgent

`func (o *InstanceConfig) GetNoAgent() InstanceConfigNoAgent`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *InstanceConfig) GetNoAgentOk() (*InstanceConfigNoAgent, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *InstanceConfig) SetNoAgent(v InstanceConfigNoAgent)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *InstanceConfig) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *InstanceConfig) GetSecurityGroups() []AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *InstanceConfig) GetSecurityGroupsOk() (*[]AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *InstanceConfig) SetSecurityGroups(v []AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *InstanceConfig) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetKvmHostId

`func (o *InstanceConfig) GetKvmHostId() int64`

GetKvmHostId returns the KvmHostId field if non-nil, zero value otherwise.

### GetKvmHostIdOk

`func (o *InstanceConfig) GetKvmHostIdOk() (*int64, bool)`

GetKvmHostIdOk returns a tuple with the KvmHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmHostId

`func (o *InstanceConfig) SetKvmHostId(v int64)`

SetKvmHostId sets KvmHostId field to given value.

### HasKvmHostId

`func (o *InstanceConfig) HasKvmHostId() bool`

HasKvmHostId returns a boolean if a field has been set.

### SetKvmHostIdNil

`func (o *InstanceConfig) SetKvmHostIdNil(b bool)`

 SetKvmHostIdNil sets the value for KvmHostId to be an explicit nil

### UnsetKvmHostId
`func (o *InstanceConfig) UnsetKvmHostId()`

UnsetKvmHostId ensures that no value is present for KvmHostId, not even an explicit nil
### GetSmbiosAssetTag

`func (o *InstanceConfig) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *InstanceConfig) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *InstanceConfig) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *InstanceConfig) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### SetSmbiosAssetTagNil

`func (o *InstanceConfig) SetSmbiosAssetTagNil(b bool)`

 SetSmbiosAssetTagNil sets the value for SmbiosAssetTag to be an explicit nil

### UnsetSmbiosAssetTag
`func (o *InstanceConfig) UnsetSmbiosAssetTag()`

UnsetSmbiosAssetTag ensures that no value is present for SmbiosAssetTag, not even an explicit nil
### GetNestedVirtualization

`func (o *InstanceConfig) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *InstanceConfig) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *InstanceConfig) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *InstanceConfig) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### SetNestedVirtualizationNil

`func (o *InstanceConfig) SetNestedVirtualizationNil(b bool)`

 SetNestedVirtualizationNil sets the value for NestedVirtualization to be an explicit nil

### UnsetNestedVirtualization
`func (o *InstanceConfig) UnsetNestedVirtualization()`

UnsetNestedVirtualization ensures that no value is present for NestedVirtualization, not even an explicit nil
### GetVmwareFolderId

`func (o *InstanceConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *InstanceConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *InstanceConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *InstanceConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetCustomOptions

`func (o *InstanceConfig) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *InstanceConfig) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *InstanceConfig) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *InstanceConfig) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *InstanceConfig) GetResourcePoolId() InstanceConfigResourcePoolId`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *InstanceConfig) GetResourcePoolIdOk() (*InstanceConfigResourcePoolId, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *InstanceConfig) SetResourcePoolId(v InstanceConfigResourcePoolId)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *InstanceConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetPoolProviderType

`func (o *InstanceConfig) GetPoolProviderType() string`

GetPoolProviderType returns the PoolProviderType field if non-nil, zero value otherwise.

### GetPoolProviderTypeOk

`func (o *InstanceConfig) GetPoolProviderTypeOk() (*string, bool)`

GetPoolProviderTypeOk returns a tuple with the PoolProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolProviderType

`func (o *InstanceConfig) SetPoolProviderType(v string)`

SetPoolProviderType sets PoolProviderType field to given value.

### HasPoolProviderType

`func (o *InstanceConfig) HasPoolProviderType() bool`

HasPoolProviderType returns a boolean if a field has been set.

### SetPoolProviderTypeNil

`func (o *InstanceConfig) SetPoolProviderTypeNil(b bool)`

 SetPoolProviderTypeNil sets the value for PoolProviderType to be an explicit nil

### UnsetPoolProviderType
`func (o *InstanceConfig) UnsetPoolProviderType()`

UnsetPoolProviderType ensures that no value is present for PoolProviderType, not even an explicit nil
### GetUserGroup

`func (o *InstanceConfig) GetUserGroup() InstanceConfigUserGroup`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *InstanceConfig) GetUserGroupOk() (*InstanceConfigUserGroup, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *InstanceConfig) SetUserGroup(v InstanceConfigUserGroup)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *InstanceConfig) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetExpireDays

`func (o *InstanceConfig) GetExpireDays() string`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *InstanceConfig) GetExpireDaysOk() (*string, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *InstanceConfig) SetExpireDays(v string)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *InstanceConfig) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetShutdownDays

`func (o *InstanceConfig) GetShutdownDays() string`

GetShutdownDays returns the ShutdownDays field if non-nil, zero value otherwise.

### GetShutdownDaysOk

`func (o *InstanceConfig) GetShutdownDaysOk() (*string, bool)`

GetShutdownDaysOk returns a tuple with the ShutdownDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDays

`func (o *InstanceConfig) SetShutdownDays(v string)`

SetShutdownDays sets ShutdownDays field to given value.

### HasShutdownDays

`func (o *InstanceConfig) HasShutdownDays() bool`

HasShutdownDays returns a boolean if a field has been set.

### GetName

`func (o *InstanceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostName

`func (o *InstanceConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *InstanceConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *InstanceConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *InstanceConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetInstanceType

`func (o *InstanceConfig) GetInstanceType() InstanceConfigInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *InstanceConfig) GetInstanceTypeOk() (*InstanceConfigInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *InstanceConfig) SetInstanceType(v InstanceConfigInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *InstanceConfig) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetSite

`func (o *InstanceConfig) GetSite() InstanceConfigSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *InstanceConfig) GetSiteOk() (*InstanceConfigSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *InstanceConfig) SetSite(v InstanceConfigSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *InstanceConfig) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *InstanceConfig) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *InstanceConfig) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *InstanceConfig) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *InstanceConfig) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### SetEnvironmentPrefixNil

`func (o *InstanceConfig) SetEnvironmentPrefixNil(b bool)`

 SetEnvironmentPrefixNil sets the value for EnvironmentPrefix to be an explicit nil

### UnsetEnvironmentPrefix
`func (o *InstanceConfig) UnsetEnvironmentPrefix()`

UnsetEnvironmentPrefix ensures that no value is present for EnvironmentPrefix, not even an explicit nil
### GetLayout

`func (o *InstanceConfig) GetLayout() InstanceConfigLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InstanceConfig) GetLayoutOk() (*InstanceConfigLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InstanceConfig) SetLayout(v InstanceConfigLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *InstanceConfig) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetType

`func (o *InstanceConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInstanceContext

`func (o *InstanceConfig) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *InstanceConfig) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *InstanceConfig) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *InstanceConfig) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### GetMemoryDisplay

`func (o *InstanceConfig) GetMemoryDisplay() string`

GetMemoryDisplay returns the MemoryDisplay field if non-nil, zero value otherwise.

### GetMemoryDisplayOk

`func (o *InstanceConfig) GetMemoryDisplayOk() (*string, bool)`

GetMemoryDisplayOk returns a tuple with the MemoryDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDisplay

`func (o *InstanceConfig) SetMemoryDisplay(v string)`

SetMemoryDisplay sets MemoryDisplay field to given value.

### HasMemoryDisplay

`func (o *InstanceConfig) HasMemoryDisplay() bool`

HasMemoryDisplay returns a boolean if a field has been set.

### GetExpose

`func (o *InstanceConfig) GetExpose() []int64`

GetExpose returns the Expose field if non-nil, zero value otherwise.

### GetExposeOk

`func (o *InstanceConfig) GetExposeOk() (*[]int64, bool)`

GetExposeOk returns a tuple with the Expose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpose

`func (o *InstanceConfig) SetExpose(v []int64)`

SetExpose sets Expose field to given value.

### HasExpose

`func (o *InstanceConfig) HasExpose() bool`

HasExpose returns a boolean if a field has been set.

### GetCreateBackup

`func (o *InstanceConfig) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *InstanceConfig) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *InstanceConfig) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *InstanceConfig) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.

### GetBackup

`func (o *InstanceConfig) GetBackup() InstanceConfigBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *InstanceConfig) GetBackupOk() (*InstanceConfigBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *InstanceConfig) SetBackup(v InstanceConfigBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *InstanceConfig) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetReplicationGroup

`func (o *InstanceConfig) GetReplicationGroup() InstanceConfigReplicationGroup`

GetReplicationGroup returns the ReplicationGroup field if non-nil, zero value otherwise.

### GetReplicationGroupOk

`func (o *InstanceConfig) GetReplicationGroupOk() (*InstanceConfigReplicationGroup, bool)`

GetReplicationGroupOk returns a tuple with the ReplicationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroup

`func (o *InstanceConfig) SetReplicationGroup(v InstanceConfigReplicationGroup)`

SetReplicationGroup sets ReplicationGroup field to given value.

### HasReplicationGroup

`func (o *InstanceConfig) HasReplicationGroup() bool`

HasReplicationGroup returns a boolean if a field has been set.

### GetLayoutSize

`func (o *InstanceConfig) GetLayoutSize() int64`

GetLayoutSize returns the LayoutSize field if non-nil, zero value otherwise.

### GetLayoutSizeOk

`func (o *InstanceConfig) GetLayoutSizeOk() (*int64, bool)`

GetLayoutSizeOk returns a tuple with the LayoutSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSize

`func (o *InstanceConfig) SetLayoutSize(v int64)`

SetLayoutSize sets LayoutSize field to given value.

### HasLayoutSize

`func (o *InstanceConfig) HasLayoutSize() bool`

HasLayoutSize returns a boolean if a field has been set.

### GetLbInstances

`func (o *InstanceConfig) GetLbInstances() []map[string]interface{}`

GetLbInstances returns the LbInstances field if non-nil, zero value otherwise.

### GetLbInstancesOk

`func (o *InstanceConfig) GetLbInstancesOk() (*[]map[string]interface{}, bool)`

GetLbInstancesOk returns a tuple with the LbInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbInstances

`func (o *InstanceConfig) SetLbInstances(v []map[string]interface{})`

SetLbInstances sets LbInstances field to given value.

### HasLbInstances

`func (o *InstanceConfig) HasLbInstances() bool`

HasLbInstances returns a boolean if a field has been set.

### SetLbInstancesNil

`func (o *InstanceConfig) SetLbInstancesNil(b bool)`

 SetLbInstancesNil sets the value for LbInstances to be an explicit nil

### UnsetLbInstances
`func (o *InstanceConfig) UnsetLbInstances()`

UnsetLbInstances ensures that no value is present for LbInstances, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



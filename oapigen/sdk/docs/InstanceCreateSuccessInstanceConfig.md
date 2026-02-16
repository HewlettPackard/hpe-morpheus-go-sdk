# InstanceCreateSuccessInstanceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateUser** | Pointer to **bool** |  | [optional] 
**IsEC2** | Pointer to **bool** |  | [optional] 
**IsVpcSelectable** | Pointer to **bool** |  | [optional] 
**NoAgent** | Pointer to [**InstanceCreateSuccessInstanceConfigNoAgent**](InstanceCreateSuccessInstanceConfigNoAgent.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner**](AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner.md) |  | [optional] 
**KvmHostId** | Pointer to **NullableInt64** |  | [optional] 
**SmbiosAssetTag** | Pointer to **NullableString** |  | [optional] 
**NestedVirtualization** | Pointer to **NullableString** |  | [optional] 
**VmwareFolderId** | Pointer to **string** |  | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**ResourcePoolId** | Pointer to [**InstanceCreateSuccessInstanceConfigResourcePoolId**](InstanceCreateSuccessInstanceConfigResourcePoolId.md) |  | [optional] 
**PoolProviderType** | Pointer to **NullableString** |  | [optional] 
**UserGroup** | Pointer to [**InstanceCreateSuccessInstanceConfigUserGroup**](InstanceCreateSuccessInstanceConfigUserGroup.md) |  | [optional] 
**ExpireDays** | Pointer to **string** |  | [optional] 
**ShutdownDays** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to [**InstanceCreateSuccessInstanceConfigInstanceType**](InstanceCreateSuccessInstanceConfigInstanceType.md) |  | [optional] 
**Site** | Pointer to [**InstanceCreateSuccessInstanceConfigSite**](InstanceCreateSuccessInstanceConfigSite.md) |  | [optional] 
**EnvironmentPrefix** | Pointer to **NullableString** |  | [optional] 
**Layout** | Pointer to [**InstanceCreateSuccessInstanceConfigLayout**](InstanceCreateSuccessInstanceConfigLayout.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**InstanceContext** | Pointer to **string** |  | [optional] 
**MemoryDisplay** | Pointer to **string** |  | [optional] 
**Expose** | Pointer to **[]int64** |  | [optional] 
**CreateBackup** | Pointer to **bool** |  | [optional] 
**Backup** | Pointer to [**InstanceCreateSuccessInstanceConfigBackup**](InstanceCreateSuccessInstanceConfigBackup.md) |  | [optional] 
**ReplicationGroup** | Pointer to [**InstanceCreateSuccessInstanceConfigReplicationGroup**](InstanceCreateSuccessInstanceConfigReplicationGroup.md) |  | [optional] 
**LayoutSize** | Pointer to **int64** |  | [optional] 
**LbInstances** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewInstanceCreateSuccessInstanceConfig

`func NewInstanceCreateSuccessInstanceConfig() *InstanceCreateSuccessInstanceConfig`

NewInstanceCreateSuccessInstanceConfig instantiates a new InstanceCreateSuccessInstanceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceCreateSuccessInstanceConfigWithDefaults

`func NewInstanceCreateSuccessInstanceConfigWithDefaults() *InstanceCreateSuccessInstanceConfig`

NewInstanceCreateSuccessInstanceConfigWithDefaults instantiates a new InstanceCreateSuccessInstanceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateUser

`func (o *InstanceCreateSuccessInstanceConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *InstanceCreateSuccessInstanceConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *InstanceCreateSuccessInstanceConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *InstanceCreateSuccessInstanceConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetIsEC2

`func (o *InstanceCreateSuccessInstanceConfig) GetIsEC2() bool`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *InstanceCreateSuccessInstanceConfig) GetIsEC2Ok() (*bool, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *InstanceCreateSuccessInstanceConfig) SetIsEC2(v bool)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *InstanceCreateSuccessInstanceConfig) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetIsVpcSelectable

`func (o *InstanceCreateSuccessInstanceConfig) GetIsVpcSelectable() bool`

GetIsVpcSelectable returns the IsVpcSelectable field if non-nil, zero value otherwise.

### GetIsVpcSelectableOk

`func (o *InstanceCreateSuccessInstanceConfig) GetIsVpcSelectableOk() (*bool, bool)`

GetIsVpcSelectableOk returns a tuple with the IsVpcSelectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpcSelectable

`func (o *InstanceCreateSuccessInstanceConfig) SetIsVpcSelectable(v bool)`

SetIsVpcSelectable sets IsVpcSelectable field to given value.

### HasIsVpcSelectable

`func (o *InstanceCreateSuccessInstanceConfig) HasIsVpcSelectable() bool`

HasIsVpcSelectable returns a boolean if a field has been set.

### GetNoAgent

`func (o *InstanceCreateSuccessInstanceConfig) GetNoAgent() InstanceCreateSuccessInstanceConfigNoAgent`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *InstanceCreateSuccessInstanceConfig) GetNoAgentOk() (*InstanceCreateSuccessInstanceConfigNoAgent, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *InstanceCreateSuccessInstanceConfig) SetNoAgent(v InstanceCreateSuccessInstanceConfigNoAgent)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *InstanceCreateSuccessInstanceConfig) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *InstanceCreateSuccessInstanceConfig) GetSecurityGroups() []AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *InstanceCreateSuccessInstanceConfig) GetSecurityGroupsOk() (*[]AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *InstanceCreateSuccessInstanceConfig) SetSecurityGroups(v []AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *InstanceCreateSuccessInstanceConfig) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetKvmHostId

`func (o *InstanceCreateSuccessInstanceConfig) GetKvmHostId() int64`

GetKvmHostId returns the KvmHostId field if non-nil, zero value otherwise.

### GetKvmHostIdOk

`func (o *InstanceCreateSuccessInstanceConfig) GetKvmHostIdOk() (*int64, bool)`

GetKvmHostIdOk returns a tuple with the KvmHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmHostId

`func (o *InstanceCreateSuccessInstanceConfig) SetKvmHostId(v int64)`

SetKvmHostId sets KvmHostId field to given value.

### HasKvmHostId

`func (o *InstanceCreateSuccessInstanceConfig) HasKvmHostId() bool`

HasKvmHostId returns a boolean if a field has been set.

### SetKvmHostIdNil

`func (o *InstanceCreateSuccessInstanceConfig) SetKvmHostIdNil(b bool)`

 SetKvmHostIdNil sets the value for KvmHostId to be an explicit nil

### UnsetKvmHostId
`func (o *InstanceCreateSuccessInstanceConfig) UnsetKvmHostId()`

UnsetKvmHostId ensures that no value is present for KvmHostId, not even an explicit nil
### GetSmbiosAssetTag

`func (o *InstanceCreateSuccessInstanceConfig) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *InstanceCreateSuccessInstanceConfig) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *InstanceCreateSuccessInstanceConfig) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *InstanceCreateSuccessInstanceConfig) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### SetSmbiosAssetTagNil

`func (o *InstanceCreateSuccessInstanceConfig) SetSmbiosAssetTagNil(b bool)`

 SetSmbiosAssetTagNil sets the value for SmbiosAssetTag to be an explicit nil

### UnsetSmbiosAssetTag
`func (o *InstanceCreateSuccessInstanceConfig) UnsetSmbiosAssetTag()`

UnsetSmbiosAssetTag ensures that no value is present for SmbiosAssetTag, not even an explicit nil
### GetNestedVirtualization

`func (o *InstanceCreateSuccessInstanceConfig) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *InstanceCreateSuccessInstanceConfig) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *InstanceCreateSuccessInstanceConfig) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *InstanceCreateSuccessInstanceConfig) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### SetNestedVirtualizationNil

`func (o *InstanceCreateSuccessInstanceConfig) SetNestedVirtualizationNil(b bool)`

 SetNestedVirtualizationNil sets the value for NestedVirtualization to be an explicit nil

### UnsetNestedVirtualization
`func (o *InstanceCreateSuccessInstanceConfig) UnsetNestedVirtualization()`

UnsetNestedVirtualization ensures that no value is present for NestedVirtualization, not even an explicit nil
### GetVmwareFolderId

`func (o *InstanceCreateSuccessInstanceConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *InstanceCreateSuccessInstanceConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *InstanceCreateSuccessInstanceConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *InstanceCreateSuccessInstanceConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetCustomOptions

`func (o *InstanceCreateSuccessInstanceConfig) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *InstanceCreateSuccessInstanceConfig) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *InstanceCreateSuccessInstanceConfig) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *InstanceCreateSuccessInstanceConfig) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *InstanceCreateSuccessInstanceConfig) GetResourcePoolId() InstanceCreateSuccessInstanceConfigResourcePoolId`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *InstanceCreateSuccessInstanceConfig) GetResourcePoolIdOk() (*InstanceCreateSuccessInstanceConfigResourcePoolId, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *InstanceCreateSuccessInstanceConfig) SetResourcePoolId(v InstanceCreateSuccessInstanceConfigResourcePoolId)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *InstanceCreateSuccessInstanceConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetPoolProviderType

`func (o *InstanceCreateSuccessInstanceConfig) GetPoolProviderType() string`

GetPoolProviderType returns the PoolProviderType field if non-nil, zero value otherwise.

### GetPoolProviderTypeOk

`func (o *InstanceCreateSuccessInstanceConfig) GetPoolProviderTypeOk() (*string, bool)`

GetPoolProviderTypeOk returns a tuple with the PoolProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolProviderType

`func (o *InstanceCreateSuccessInstanceConfig) SetPoolProviderType(v string)`

SetPoolProviderType sets PoolProviderType field to given value.

### HasPoolProviderType

`func (o *InstanceCreateSuccessInstanceConfig) HasPoolProviderType() bool`

HasPoolProviderType returns a boolean if a field has been set.

### SetPoolProviderTypeNil

`func (o *InstanceCreateSuccessInstanceConfig) SetPoolProviderTypeNil(b bool)`

 SetPoolProviderTypeNil sets the value for PoolProviderType to be an explicit nil

### UnsetPoolProviderType
`func (o *InstanceCreateSuccessInstanceConfig) UnsetPoolProviderType()`

UnsetPoolProviderType ensures that no value is present for PoolProviderType, not even an explicit nil
### GetUserGroup

`func (o *InstanceCreateSuccessInstanceConfig) GetUserGroup() InstanceCreateSuccessInstanceConfigUserGroup`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *InstanceCreateSuccessInstanceConfig) GetUserGroupOk() (*InstanceCreateSuccessInstanceConfigUserGroup, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *InstanceCreateSuccessInstanceConfig) SetUserGroup(v InstanceCreateSuccessInstanceConfigUserGroup)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *InstanceCreateSuccessInstanceConfig) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetExpireDays

`func (o *InstanceCreateSuccessInstanceConfig) GetExpireDays() string`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *InstanceCreateSuccessInstanceConfig) GetExpireDaysOk() (*string, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *InstanceCreateSuccessInstanceConfig) SetExpireDays(v string)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *InstanceCreateSuccessInstanceConfig) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetShutdownDays

`func (o *InstanceCreateSuccessInstanceConfig) GetShutdownDays() string`

GetShutdownDays returns the ShutdownDays field if non-nil, zero value otherwise.

### GetShutdownDaysOk

`func (o *InstanceCreateSuccessInstanceConfig) GetShutdownDaysOk() (*string, bool)`

GetShutdownDaysOk returns a tuple with the ShutdownDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDays

`func (o *InstanceCreateSuccessInstanceConfig) SetShutdownDays(v string)`

SetShutdownDays sets ShutdownDays field to given value.

### HasShutdownDays

`func (o *InstanceCreateSuccessInstanceConfig) HasShutdownDays() bool`

HasShutdownDays returns a boolean if a field has been set.

### GetName

`func (o *InstanceCreateSuccessInstanceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceCreateSuccessInstanceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceCreateSuccessInstanceConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceCreateSuccessInstanceConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostName

`func (o *InstanceCreateSuccessInstanceConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *InstanceCreateSuccessInstanceConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *InstanceCreateSuccessInstanceConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *InstanceCreateSuccessInstanceConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetInstanceType

`func (o *InstanceCreateSuccessInstanceConfig) GetInstanceType() InstanceCreateSuccessInstanceConfigInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *InstanceCreateSuccessInstanceConfig) GetInstanceTypeOk() (*InstanceCreateSuccessInstanceConfigInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *InstanceCreateSuccessInstanceConfig) SetInstanceType(v InstanceCreateSuccessInstanceConfigInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *InstanceCreateSuccessInstanceConfig) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetSite

`func (o *InstanceCreateSuccessInstanceConfig) GetSite() InstanceCreateSuccessInstanceConfigSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *InstanceCreateSuccessInstanceConfig) GetSiteOk() (*InstanceCreateSuccessInstanceConfigSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *InstanceCreateSuccessInstanceConfig) SetSite(v InstanceCreateSuccessInstanceConfigSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *InstanceCreateSuccessInstanceConfig) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *InstanceCreateSuccessInstanceConfig) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *InstanceCreateSuccessInstanceConfig) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *InstanceCreateSuccessInstanceConfig) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *InstanceCreateSuccessInstanceConfig) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### SetEnvironmentPrefixNil

`func (o *InstanceCreateSuccessInstanceConfig) SetEnvironmentPrefixNil(b bool)`

 SetEnvironmentPrefixNil sets the value for EnvironmentPrefix to be an explicit nil

### UnsetEnvironmentPrefix
`func (o *InstanceCreateSuccessInstanceConfig) UnsetEnvironmentPrefix()`

UnsetEnvironmentPrefix ensures that no value is present for EnvironmentPrefix, not even an explicit nil
### GetLayout

`func (o *InstanceCreateSuccessInstanceConfig) GetLayout() InstanceCreateSuccessInstanceConfigLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InstanceCreateSuccessInstanceConfig) GetLayoutOk() (*InstanceCreateSuccessInstanceConfigLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InstanceCreateSuccessInstanceConfig) SetLayout(v InstanceCreateSuccessInstanceConfigLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *InstanceCreateSuccessInstanceConfig) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetType

`func (o *InstanceCreateSuccessInstanceConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceCreateSuccessInstanceConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceCreateSuccessInstanceConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceCreateSuccessInstanceConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInstanceContext

`func (o *InstanceCreateSuccessInstanceConfig) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *InstanceCreateSuccessInstanceConfig) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *InstanceCreateSuccessInstanceConfig) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *InstanceCreateSuccessInstanceConfig) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### GetMemoryDisplay

`func (o *InstanceCreateSuccessInstanceConfig) GetMemoryDisplay() string`

GetMemoryDisplay returns the MemoryDisplay field if non-nil, zero value otherwise.

### GetMemoryDisplayOk

`func (o *InstanceCreateSuccessInstanceConfig) GetMemoryDisplayOk() (*string, bool)`

GetMemoryDisplayOk returns a tuple with the MemoryDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDisplay

`func (o *InstanceCreateSuccessInstanceConfig) SetMemoryDisplay(v string)`

SetMemoryDisplay sets MemoryDisplay field to given value.

### HasMemoryDisplay

`func (o *InstanceCreateSuccessInstanceConfig) HasMemoryDisplay() bool`

HasMemoryDisplay returns a boolean if a field has been set.

### GetExpose

`func (o *InstanceCreateSuccessInstanceConfig) GetExpose() []int64`

GetExpose returns the Expose field if non-nil, zero value otherwise.

### GetExposeOk

`func (o *InstanceCreateSuccessInstanceConfig) GetExposeOk() (*[]int64, bool)`

GetExposeOk returns a tuple with the Expose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpose

`func (o *InstanceCreateSuccessInstanceConfig) SetExpose(v []int64)`

SetExpose sets Expose field to given value.

### HasExpose

`func (o *InstanceCreateSuccessInstanceConfig) HasExpose() bool`

HasExpose returns a boolean if a field has been set.

### GetCreateBackup

`func (o *InstanceCreateSuccessInstanceConfig) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *InstanceCreateSuccessInstanceConfig) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *InstanceCreateSuccessInstanceConfig) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *InstanceCreateSuccessInstanceConfig) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.

### GetBackup

`func (o *InstanceCreateSuccessInstanceConfig) GetBackup() InstanceCreateSuccessInstanceConfigBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *InstanceCreateSuccessInstanceConfig) GetBackupOk() (*InstanceCreateSuccessInstanceConfigBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *InstanceCreateSuccessInstanceConfig) SetBackup(v InstanceCreateSuccessInstanceConfigBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *InstanceCreateSuccessInstanceConfig) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetReplicationGroup

`func (o *InstanceCreateSuccessInstanceConfig) GetReplicationGroup() InstanceCreateSuccessInstanceConfigReplicationGroup`

GetReplicationGroup returns the ReplicationGroup field if non-nil, zero value otherwise.

### GetReplicationGroupOk

`func (o *InstanceCreateSuccessInstanceConfig) GetReplicationGroupOk() (*InstanceCreateSuccessInstanceConfigReplicationGroup, bool)`

GetReplicationGroupOk returns a tuple with the ReplicationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroup

`func (o *InstanceCreateSuccessInstanceConfig) SetReplicationGroup(v InstanceCreateSuccessInstanceConfigReplicationGroup)`

SetReplicationGroup sets ReplicationGroup field to given value.

### HasReplicationGroup

`func (o *InstanceCreateSuccessInstanceConfig) HasReplicationGroup() bool`

HasReplicationGroup returns a boolean if a field has been set.

### GetLayoutSize

`func (o *InstanceCreateSuccessInstanceConfig) GetLayoutSize() int64`

GetLayoutSize returns the LayoutSize field if non-nil, zero value otherwise.

### GetLayoutSizeOk

`func (o *InstanceCreateSuccessInstanceConfig) GetLayoutSizeOk() (*int64, bool)`

GetLayoutSizeOk returns a tuple with the LayoutSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSize

`func (o *InstanceCreateSuccessInstanceConfig) SetLayoutSize(v int64)`

SetLayoutSize sets LayoutSize field to given value.

### HasLayoutSize

`func (o *InstanceCreateSuccessInstanceConfig) HasLayoutSize() bool`

HasLayoutSize returns a boolean if a field has been set.

### GetLbInstances

`func (o *InstanceCreateSuccessInstanceConfig) GetLbInstances() []map[string]interface{}`

GetLbInstances returns the LbInstances field if non-nil, zero value otherwise.

### GetLbInstancesOk

`func (o *InstanceCreateSuccessInstanceConfig) GetLbInstancesOk() (*[]map[string]interface{}, bool)`

GetLbInstancesOk returns a tuple with the LbInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbInstances

`func (o *InstanceCreateSuccessInstanceConfig) SetLbInstances(v []map[string]interface{})`

SetLbInstances sets LbInstances field to given value.

### HasLbInstances

`func (o *InstanceCreateSuccessInstanceConfig) HasLbInstances() bool`

HasLbInstances returns a boolean if a field has been set.

### SetLbInstancesNil

`func (o *InstanceCreateSuccessInstanceConfig) SetLbInstancesNil(b bool)`

 SetLbInstancesNil sets the value for LbInstances to be an explicit nil

### UnsetLbInstances
`func (o *InstanceCreateSuccessInstanceConfig) UnsetLbInstances()`

UnsetLbInstances ensures that no value is present for LbInstances, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



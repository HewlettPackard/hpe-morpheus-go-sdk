# GetInstance200ResponseInstanceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateUser** | Pointer to **bool** |  | [optional] 
**IsEC2** | Pointer to **bool** |  | [optional] 
**IsVpcSelectable** | Pointer to **bool** |  | [optional] 
**NoAgent** | Pointer to [**GetInstance200ResponseInstanceConfigNoAgent**](GetInstance200ResponseInstanceConfigNoAgent.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner**](AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner.md) |  | [optional] 
**KvmHostId** | Pointer to **NullableInt64** |  | [optional] 
**SmbiosAssetTag** | Pointer to **NullableString** |  | [optional] 
**NestedVirtualization** | Pointer to **NullableString** |  | [optional] 
**VmwareFolderId** | Pointer to **string** |  | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**ResourcePoolId** | Pointer to [**GetInstance200ResponseInstanceConfigResourcePoolId**](GetInstance200ResponseInstanceConfigResourcePoolId.md) |  | [optional] 
**PoolProviderType** | Pointer to **NullableString** |  | [optional] 
**UserGroup** | Pointer to [**GetInstance200ResponseInstanceConfigUserGroup**](GetInstance200ResponseInstanceConfigUserGroup.md) |  | [optional] 
**ExpireDays** | Pointer to **string** |  | [optional] 
**ShutdownDays** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to [**GetInstance200ResponseInstanceConfigInstanceType**](GetInstance200ResponseInstanceConfigInstanceType.md) |  | [optional] 
**Site** | Pointer to [**GetInstance200ResponseInstanceConfigSite**](GetInstance200ResponseInstanceConfigSite.md) |  | [optional] 
**EnvironmentPrefix** | Pointer to **NullableString** |  | [optional] 
**Layout** | Pointer to [**GetInstance200ResponseInstanceConfigLayout**](GetInstance200ResponseInstanceConfigLayout.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**InstanceContext** | Pointer to **string** |  | [optional] 
**MemoryDisplay** | Pointer to **string** |  | [optional] 
**Expose** | Pointer to **[]int64** |  | [optional] 
**CreateBackup** | Pointer to **bool** |  | [optional] 
**Backup** | Pointer to [**GetInstance200ResponseInstanceConfigBackup**](GetInstance200ResponseInstanceConfigBackup.md) |  | [optional] 
**ReplicationGroup** | Pointer to [**GetInstance200ResponseInstanceConfigReplicationGroup**](GetInstance200ResponseInstanceConfigReplicationGroup.md) |  | [optional] 
**LayoutSize** | Pointer to **int64** |  | [optional] 
**LbInstances** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetInstance200ResponseInstanceConfig

`func NewGetInstance200ResponseInstanceConfig() *GetInstance200ResponseInstanceConfig`

NewGetInstance200ResponseInstanceConfig instantiates a new GetInstance200ResponseInstanceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstance200ResponseInstanceConfigWithDefaults

`func NewGetInstance200ResponseInstanceConfigWithDefaults() *GetInstance200ResponseInstanceConfig`

NewGetInstance200ResponseInstanceConfigWithDefaults instantiates a new GetInstance200ResponseInstanceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateUser

`func (o *GetInstance200ResponseInstanceConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *GetInstance200ResponseInstanceConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *GetInstance200ResponseInstanceConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *GetInstance200ResponseInstanceConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetIsEC2

`func (o *GetInstance200ResponseInstanceConfig) GetIsEC2() bool`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *GetInstance200ResponseInstanceConfig) GetIsEC2Ok() (*bool, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *GetInstance200ResponseInstanceConfig) SetIsEC2(v bool)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *GetInstance200ResponseInstanceConfig) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetIsVpcSelectable

`func (o *GetInstance200ResponseInstanceConfig) GetIsVpcSelectable() bool`

GetIsVpcSelectable returns the IsVpcSelectable field if non-nil, zero value otherwise.

### GetIsVpcSelectableOk

`func (o *GetInstance200ResponseInstanceConfig) GetIsVpcSelectableOk() (*bool, bool)`

GetIsVpcSelectableOk returns a tuple with the IsVpcSelectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpcSelectable

`func (o *GetInstance200ResponseInstanceConfig) SetIsVpcSelectable(v bool)`

SetIsVpcSelectable sets IsVpcSelectable field to given value.

### HasIsVpcSelectable

`func (o *GetInstance200ResponseInstanceConfig) HasIsVpcSelectable() bool`

HasIsVpcSelectable returns a boolean if a field has been set.

### GetNoAgent

`func (o *GetInstance200ResponseInstanceConfig) GetNoAgent() GetInstance200ResponseInstanceConfigNoAgent`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *GetInstance200ResponseInstanceConfig) GetNoAgentOk() (*GetInstance200ResponseInstanceConfigNoAgent, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *GetInstance200ResponseInstanceConfig) SetNoAgent(v GetInstance200ResponseInstanceConfigNoAgent)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *GetInstance200ResponseInstanceConfig) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *GetInstance200ResponseInstanceConfig) GetSecurityGroups() []AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *GetInstance200ResponseInstanceConfig) GetSecurityGroupsOk() (*[]AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *GetInstance200ResponseInstanceConfig) SetSecurityGroups(v []AddInstance200ResponseAllOfOneOfInstanceConfigSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *GetInstance200ResponseInstanceConfig) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetKvmHostId

`func (o *GetInstance200ResponseInstanceConfig) GetKvmHostId() int64`

GetKvmHostId returns the KvmHostId field if non-nil, zero value otherwise.

### GetKvmHostIdOk

`func (o *GetInstance200ResponseInstanceConfig) GetKvmHostIdOk() (*int64, bool)`

GetKvmHostIdOk returns a tuple with the KvmHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmHostId

`func (o *GetInstance200ResponseInstanceConfig) SetKvmHostId(v int64)`

SetKvmHostId sets KvmHostId field to given value.

### HasKvmHostId

`func (o *GetInstance200ResponseInstanceConfig) HasKvmHostId() bool`

HasKvmHostId returns a boolean if a field has been set.

### SetKvmHostIdNil

`func (o *GetInstance200ResponseInstanceConfig) SetKvmHostIdNil(b bool)`

 SetKvmHostIdNil sets the value for KvmHostId to be an explicit nil

### UnsetKvmHostId
`func (o *GetInstance200ResponseInstanceConfig) UnsetKvmHostId()`

UnsetKvmHostId ensures that no value is present for KvmHostId, not even an explicit nil
### GetSmbiosAssetTag

`func (o *GetInstance200ResponseInstanceConfig) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *GetInstance200ResponseInstanceConfig) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *GetInstance200ResponseInstanceConfig) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *GetInstance200ResponseInstanceConfig) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### SetSmbiosAssetTagNil

`func (o *GetInstance200ResponseInstanceConfig) SetSmbiosAssetTagNil(b bool)`

 SetSmbiosAssetTagNil sets the value for SmbiosAssetTag to be an explicit nil

### UnsetSmbiosAssetTag
`func (o *GetInstance200ResponseInstanceConfig) UnsetSmbiosAssetTag()`

UnsetSmbiosAssetTag ensures that no value is present for SmbiosAssetTag, not even an explicit nil
### GetNestedVirtualization

`func (o *GetInstance200ResponseInstanceConfig) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *GetInstance200ResponseInstanceConfig) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *GetInstance200ResponseInstanceConfig) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *GetInstance200ResponseInstanceConfig) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### SetNestedVirtualizationNil

`func (o *GetInstance200ResponseInstanceConfig) SetNestedVirtualizationNil(b bool)`

 SetNestedVirtualizationNil sets the value for NestedVirtualization to be an explicit nil

### UnsetNestedVirtualization
`func (o *GetInstance200ResponseInstanceConfig) UnsetNestedVirtualization()`

UnsetNestedVirtualization ensures that no value is present for NestedVirtualization, not even an explicit nil
### GetVmwareFolderId

`func (o *GetInstance200ResponseInstanceConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *GetInstance200ResponseInstanceConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *GetInstance200ResponseInstanceConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *GetInstance200ResponseInstanceConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetCustomOptions

`func (o *GetInstance200ResponseInstanceConfig) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *GetInstance200ResponseInstanceConfig) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *GetInstance200ResponseInstanceConfig) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *GetInstance200ResponseInstanceConfig) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *GetInstance200ResponseInstanceConfig) GetResourcePoolId() GetInstance200ResponseInstanceConfigResourcePoolId`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *GetInstance200ResponseInstanceConfig) GetResourcePoolIdOk() (*GetInstance200ResponseInstanceConfigResourcePoolId, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *GetInstance200ResponseInstanceConfig) SetResourcePoolId(v GetInstance200ResponseInstanceConfigResourcePoolId)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *GetInstance200ResponseInstanceConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetPoolProviderType

`func (o *GetInstance200ResponseInstanceConfig) GetPoolProviderType() string`

GetPoolProviderType returns the PoolProviderType field if non-nil, zero value otherwise.

### GetPoolProviderTypeOk

`func (o *GetInstance200ResponseInstanceConfig) GetPoolProviderTypeOk() (*string, bool)`

GetPoolProviderTypeOk returns a tuple with the PoolProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolProviderType

`func (o *GetInstance200ResponseInstanceConfig) SetPoolProviderType(v string)`

SetPoolProviderType sets PoolProviderType field to given value.

### HasPoolProviderType

`func (o *GetInstance200ResponseInstanceConfig) HasPoolProviderType() bool`

HasPoolProviderType returns a boolean if a field has been set.

### SetPoolProviderTypeNil

`func (o *GetInstance200ResponseInstanceConfig) SetPoolProviderTypeNil(b bool)`

 SetPoolProviderTypeNil sets the value for PoolProviderType to be an explicit nil

### UnsetPoolProviderType
`func (o *GetInstance200ResponseInstanceConfig) UnsetPoolProviderType()`

UnsetPoolProviderType ensures that no value is present for PoolProviderType, not even an explicit nil
### GetUserGroup

`func (o *GetInstance200ResponseInstanceConfig) GetUserGroup() GetInstance200ResponseInstanceConfigUserGroup`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *GetInstance200ResponseInstanceConfig) GetUserGroupOk() (*GetInstance200ResponseInstanceConfigUserGroup, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *GetInstance200ResponseInstanceConfig) SetUserGroup(v GetInstance200ResponseInstanceConfigUserGroup)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *GetInstance200ResponseInstanceConfig) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetExpireDays

`func (o *GetInstance200ResponseInstanceConfig) GetExpireDays() string`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *GetInstance200ResponseInstanceConfig) GetExpireDaysOk() (*string, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *GetInstance200ResponseInstanceConfig) SetExpireDays(v string)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *GetInstance200ResponseInstanceConfig) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetShutdownDays

`func (o *GetInstance200ResponseInstanceConfig) GetShutdownDays() string`

GetShutdownDays returns the ShutdownDays field if non-nil, zero value otherwise.

### GetShutdownDaysOk

`func (o *GetInstance200ResponseInstanceConfig) GetShutdownDaysOk() (*string, bool)`

GetShutdownDaysOk returns a tuple with the ShutdownDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDays

`func (o *GetInstance200ResponseInstanceConfig) SetShutdownDays(v string)`

SetShutdownDays sets ShutdownDays field to given value.

### HasShutdownDays

`func (o *GetInstance200ResponseInstanceConfig) HasShutdownDays() bool`

HasShutdownDays returns a boolean if a field has been set.

### GetName

`func (o *GetInstance200ResponseInstanceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInstance200ResponseInstanceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInstance200ResponseInstanceConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInstance200ResponseInstanceConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostName

`func (o *GetInstance200ResponseInstanceConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *GetInstance200ResponseInstanceConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *GetInstance200ResponseInstanceConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *GetInstance200ResponseInstanceConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetInstanceType

`func (o *GetInstance200ResponseInstanceConfig) GetInstanceType() GetInstance200ResponseInstanceConfigInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *GetInstance200ResponseInstanceConfig) GetInstanceTypeOk() (*GetInstance200ResponseInstanceConfigInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *GetInstance200ResponseInstanceConfig) SetInstanceType(v GetInstance200ResponseInstanceConfigInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *GetInstance200ResponseInstanceConfig) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetSite

`func (o *GetInstance200ResponseInstanceConfig) GetSite() GetInstance200ResponseInstanceConfigSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GetInstance200ResponseInstanceConfig) GetSiteOk() (*GetInstance200ResponseInstanceConfigSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GetInstance200ResponseInstanceConfig) SetSite(v GetInstance200ResponseInstanceConfigSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *GetInstance200ResponseInstanceConfig) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *GetInstance200ResponseInstanceConfig) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *GetInstance200ResponseInstanceConfig) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *GetInstance200ResponseInstanceConfig) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *GetInstance200ResponseInstanceConfig) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### SetEnvironmentPrefixNil

`func (o *GetInstance200ResponseInstanceConfig) SetEnvironmentPrefixNil(b bool)`

 SetEnvironmentPrefixNil sets the value for EnvironmentPrefix to be an explicit nil

### UnsetEnvironmentPrefix
`func (o *GetInstance200ResponseInstanceConfig) UnsetEnvironmentPrefix()`

UnsetEnvironmentPrefix ensures that no value is present for EnvironmentPrefix, not even an explicit nil
### GetLayout

`func (o *GetInstance200ResponseInstanceConfig) GetLayout() GetInstance200ResponseInstanceConfigLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *GetInstance200ResponseInstanceConfig) GetLayoutOk() (*GetInstance200ResponseInstanceConfigLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *GetInstance200ResponseInstanceConfig) SetLayout(v GetInstance200ResponseInstanceConfigLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *GetInstance200ResponseInstanceConfig) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetType

`func (o *GetInstance200ResponseInstanceConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetInstance200ResponseInstanceConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetInstance200ResponseInstanceConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetInstance200ResponseInstanceConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInstanceContext

`func (o *GetInstance200ResponseInstanceConfig) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *GetInstance200ResponseInstanceConfig) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *GetInstance200ResponseInstanceConfig) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *GetInstance200ResponseInstanceConfig) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### GetMemoryDisplay

`func (o *GetInstance200ResponseInstanceConfig) GetMemoryDisplay() string`

GetMemoryDisplay returns the MemoryDisplay field if non-nil, zero value otherwise.

### GetMemoryDisplayOk

`func (o *GetInstance200ResponseInstanceConfig) GetMemoryDisplayOk() (*string, bool)`

GetMemoryDisplayOk returns a tuple with the MemoryDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDisplay

`func (o *GetInstance200ResponseInstanceConfig) SetMemoryDisplay(v string)`

SetMemoryDisplay sets MemoryDisplay field to given value.

### HasMemoryDisplay

`func (o *GetInstance200ResponseInstanceConfig) HasMemoryDisplay() bool`

HasMemoryDisplay returns a boolean if a field has been set.

### GetExpose

`func (o *GetInstance200ResponseInstanceConfig) GetExpose() []int64`

GetExpose returns the Expose field if non-nil, zero value otherwise.

### GetExposeOk

`func (o *GetInstance200ResponseInstanceConfig) GetExposeOk() (*[]int64, bool)`

GetExposeOk returns a tuple with the Expose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpose

`func (o *GetInstance200ResponseInstanceConfig) SetExpose(v []int64)`

SetExpose sets Expose field to given value.

### HasExpose

`func (o *GetInstance200ResponseInstanceConfig) HasExpose() bool`

HasExpose returns a boolean if a field has been set.

### GetCreateBackup

`func (o *GetInstance200ResponseInstanceConfig) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *GetInstance200ResponseInstanceConfig) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *GetInstance200ResponseInstanceConfig) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *GetInstance200ResponseInstanceConfig) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.

### GetBackup

`func (o *GetInstance200ResponseInstanceConfig) GetBackup() GetInstance200ResponseInstanceConfigBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *GetInstance200ResponseInstanceConfig) GetBackupOk() (*GetInstance200ResponseInstanceConfigBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *GetInstance200ResponseInstanceConfig) SetBackup(v GetInstance200ResponseInstanceConfigBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *GetInstance200ResponseInstanceConfig) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetReplicationGroup

`func (o *GetInstance200ResponseInstanceConfig) GetReplicationGroup() GetInstance200ResponseInstanceConfigReplicationGroup`

GetReplicationGroup returns the ReplicationGroup field if non-nil, zero value otherwise.

### GetReplicationGroupOk

`func (o *GetInstance200ResponseInstanceConfig) GetReplicationGroupOk() (*GetInstance200ResponseInstanceConfigReplicationGroup, bool)`

GetReplicationGroupOk returns a tuple with the ReplicationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroup

`func (o *GetInstance200ResponseInstanceConfig) SetReplicationGroup(v GetInstance200ResponseInstanceConfigReplicationGroup)`

SetReplicationGroup sets ReplicationGroup field to given value.

### HasReplicationGroup

`func (o *GetInstance200ResponseInstanceConfig) HasReplicationGroup() bool`

HasReplicationGroup returns a boolean if a field has been set.

### GetLayoutSize

`func (o *GetInstance200ResponseInstanceConfig) GetLayoutSize() int64`

GetLayoutSize returns the LayoutSize field if non-nil, zero value otherwise.

### GetLayoutSizeOk

`func (o *GetInstance200ResponseInstanceConfig) GetLayoutSizeOk() (*int64, bool)`

GetLayoutSizeOk returns a tuple with the LayoutSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSize

`func (o *GetInstance200ResponseInstanceConfig) SetLayoutSize(v int64)`

SetLayoutSize sets LayoutSize field to given value.

### HasLayoutSize

`func (o *GetInstance200ResponseInstanceConfig) HasLayoutSize() bool`

HasLayoutSize returns a boolean if a field has been set.

### GetLbInstances

`func (o *GetInstance200ResponseInstanceConfig) GetLbInstances() []map[string]interface{}`

GetLbInstances returns the LbInstances field if non-nil, zero value otherwise.

### GetLbInstancesOk

`func (o *GetInstance200ResponseInstanceConfig) GetLbInstancesOk() (*[]map[string]interface{}, bool)`

GetLbInstancesOk returns a tuple with the LbInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbInstances

`func (o *GetInstance200ResponseInstanceConfig) SetLbInstances(v []map[string]interface{})`

SetLbInstances sets LbInstances field to given value.

### HasLbInstances

`func (o *GetInstance200ResponseInstanceConfig) HasLbInstances() bool`

HasLbInstances returns a boolean if a field has been set.

### SetLbInstancesNil

`func (o *GetInstance200ResponseInstanceConfig) SetLbInstancesNil(b bool)`

 SetLbInstancesNil sets the value for LbInstances to be an explicit nil

### UnsetLbInstances
`func (o *GetInstance200ResponseInstanceConfig) UnsetLbInstances()`

UnsetLbInstances ensures that no value is present for LbInstances, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



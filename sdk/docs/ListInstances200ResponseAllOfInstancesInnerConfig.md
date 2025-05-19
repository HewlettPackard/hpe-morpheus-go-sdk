# ListInstances200ResponseAllOfInstancesInnerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateUser** | Pointer to **bool** |  | [optional] 
**IsEC2** | Pointer to **bool** |  | [optional] 
**IsVpcSelectable** | Pointer to **bool** |  | [optional] 
**NoAgent** | Pointer to **bool** |  | [optional] 
**SecurityGroups** | Pointer to [**[]ListInstances200ResponseAllOfInstancesInnerConfigSecurityGroupsInner**](ListInstances200ResponseAllOfInstancesInnerConfigSecurityGroupsInner.md) |  | [optional] 
**SmbiosAssetTag** | Pointer to **string** |  | [optional] 
**NestedVirtualization** | Pointer to **string** |  | [optional] 
**VmwareFolderId** | Pointer to **string** |  | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**PoolProviderType** | Pointer to **string** |  | [optional] 
**UserGroup** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerConfigSecurityGroupsInner**](ListInstances200ResponseAllOfInstancesInnerConfigSecurityGroupsInner.md) |  | [optional] 
**ExpireDays** | Pointer to **string** |  | [optional] 
**ShutdownDays** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerConfigInstanceType**](ListInstances200ResponseAllOfInstancesInnerConfigInstanceType.md) |  | [optional] 
**Site** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**EnvironmentPrefix** | Pointer to **string** |  | [optional] 
**Layout** | Pointer to [**ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan**](ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**InstanceContext** | Pointer to **string** |  | [optional] 
**MemoryDisplay** | Pointer to **string** |  | [optional] 
**Expose** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CreateBackup** | Pointer to **bool** |  | [optional] 
**Backup** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerConfigBackup**](ListInstances200ResponseAllOfInstancesInnerConfigBackup.md) |  | [optional] 
**ReplicationGroup** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerConfigReplicationGroup**](ListInstances200ResponseAllOfInstancesInnerConfigReplicationGroup.md) |  | [optional] 
**LayoutSize** | Pointer to **int64** |  | [optional] 
**LbInstances** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListInstances200ResponseAllOfInstancesInnerConfig

`func NewListInstances200ResponseAllOfInstancesInnerConfig() *ListInstances200ResponseAllOfInstancesInnerConfig`

NewListInstances200ResponseAllOfInstancesInnerConfig instantiates a new ListInstances200ResponseAllOfInstancesInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstances200ResponseAllOfInstancesInnerConfigWithDefaults

`func NewListInstances200ResponseAllOfInstancesInnerConfigWithDefaults() *ListInstances200ResponseAllOfInstancesInnerConfig`

NewListInstances200ResponseAllOfInstancesInnerConfigWithDefaults instantiates a new ListInstances200ResponseAllOfInstancesInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateUser

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetIsEC2

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetIsEC2() bool`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetIsEC2Ok() (*bool, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetIsEC2(v bool)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetIsVpcSelectable

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetIsVpcSelectable() bool`

GetIsVpcSelectable returns the IsVpcSelectable field if non-nil, zero value otherwise.

### GetIsVpcSelectableOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetIsVpcSelectableOk() (*bool, bool)`

GetIsVpcSelectableOk returns a tuple with the IsVpcSelectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpcSelectable

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetIsVpcSelectable(v bool)`

SetIsVpcSelectable sets IsVpcSelectable field to given value.

### HasIsVpcSelectable

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasIsVpcSelectable() bool`

HasIsVpcSelectable returns a boolean if a field has been set.

### GetNoAgent

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetSecurityGroups() []ListInstances200ResponseAllOfInstancesInnerConfigSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetSecurityGroupsOk() (*[]ListInstances200ResponseAllOfInstancesInnerConfigSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetSecurityGroups(v []ListInstances200ResponseAllOfInstancesInnerConfigSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetSmbiosAssetTag

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetCustomOptions

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetPoolProviderType

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetPoolProviderType() string`

GetPoolProviderType returns the PoolProviderType field if non-nil, zero value otherwise.

### GetPoolProviderTypeOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetPoolProviderTypeOk() (*string, bool)`

GetPoolProviderTypeOk returns a tuple with the PoolProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolProviderType

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetPoolProviderType(v string)`

SetPoolProviderType sets PoolProviderType field to given value.

### HasPoolProviderType

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasPoolProviderType() bool`

HasPoolProviderType returns a boolean if a field has been set.

### GetUserGroup

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetUserGroup() ListInstances200ResponseAllOfInstancesInnerConfigSecurityGroupsInner`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetUserGroupOk() (*ListInstances200ResponseAllOfInstancesInnerConfigSecurityGroupsInner, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetUserGroup(v ListInstances200ResponseAllOfInstancesInnerConfigSecurityGroupsInner)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetExpireDays

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetExpireDays() string`

GetExpireDays returns the ExpireDays field if non-nil, zero value otherwise.

### GetExpireDaysOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetExpireDaysOk() (*string, bool)`

GetExpireDaysOk returns a tuple with the ExpireDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDays

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetExpireDays(v string)`

SetExpireDays sets ExpireDays field to given value.

### HasExpireDays

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasExpireDays() bool`

HasExpireDays returns a boolean if a field has been set.

### GetShutdownDays

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetShutdownDays() string`

GetShutdownDays returns the ShutdownDays field if non-nil, zero value otherwise.

### GetShutdownDaysOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetShutdownDaysOk() (*string, bool)`

GetShutdownDaysOk returns a tuple with the ShutdownDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownDays

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetShutdownDays(v string)`

SetShutdownDays sets ShutdownDays field to given value.

### HasShutdownDays

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasShutdownDays() bool`

HasShutdownDays returns a boolean if a field has been set.

### GetName

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostName

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetInstanceType

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetInstanceType() ListInstances200ResponseAllOfInstancesInnerConfigInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetInstanceTypeOk() (*ListInstances200ResponseAllOfInstancesInnerConfigInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetInstanceType(v ListInstances200ResponseAllOfInstancesInnerConfigInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetSite

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetSite() GetAlerts200ResponseAllOfChecksInnerAccount`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetSiteOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetSite(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetSite sets Site field to given value.

### HasSite

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetLayout

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetLayout() ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetLayoutOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetLayout(v ListImageBuilds200ResponseAllOfImageBuildsInnerConfigPlan)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetType

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInstanceContext

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### GetMemoryDisplay

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetMemoryDisplay() string`

GetMemoryDisplay returns the MemoryDisplay field if non-nil, zero value otherwise.

### GetMemoryDisplayOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetMemoryDisplayOk() (*string, bool)`

GetMemoryDisplayOk returns a tuple with the MemoryDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDisplay

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetMemoryDisplay(v string)`

SetMemoryDisplay sets MemoryDisplay field to given value.

### HasMemoryDisplay

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasMemoryDisplay() bool`

HasMemoryDisplay returns a boolean if a field has been set.

### GetExpose

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetExpose() []map[string]interface{}`

GetExpose returns the Expose field if non-nil, zero value otherwise.

### GetExposeOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetExposeOk() (*[]map[string]interface{}, bool)`

GetExposeOk returns a tuple with the Expose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpose

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetExpose(v []map[string]interface{})`

SetExpose sets Expose field to given value.

### HasExpose

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasExpose() bool`

HasExpose returns a boolean if a field has been set.

### GetCreateBackup

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.

### GetBackup

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetBackup() ListInstances200ResponseAllOfInstancesInnerConfigBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetBackupOk() (*ListInstances200ResponseAllOfInstancesInnerConfigBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetBackup(v ListInstances200ResponseAllOfInstancesInnerConfigBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetReplicationGroup

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetReplicationGroup() ListInstances200ResponseAllOfInstancesInnerConfigReplicationGroup`

GetReplicationGroup returns the ReplicationGroup field if non-nil, zero value otherwise.

### GetReplicationGroupOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetReplicationGroupOk() (*ListInstances200ResponseAllOfInstancesInnerConfigReplicationGroup, bool)`

GetReplicationGroupOk returns a tuple with the ReplicationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroup

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetReplicationGroup(v ListInstances200ResponseAllOfInstancesInnerConfigReplicationGroup)`

SetReplicationGroup sets ReplicationGroup field to given value.

### HasReplicationGroup

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasReplicationGroup() bool`

HasReplicationGroup returns a boolean if a field has been set.

### GetLayoutSize

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetLayoutSize() int64`

GetLayoutSize returns the LayoutSize field if non-nil, zero value otherwise.

### GetLayoutSizeOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetLayoutSizeOk() (*int64, bool)`

GetLayoutSizeOk returns a tuple with the LayoutSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSize

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetLayoutSize(v int64)`

SetLayoutSize sets LayoutSize field to given value.

### HasLayoutSize

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasLayoutSize() bool`

HasLayoutSize returns a boolean if a field has been set.

### GetLbInstances

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetLbInstances() []map[string]interface{}`

GetLbInstances returns the LbInstances field if non-nil, zero value otherwise.

### GetLbInstancesOk

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) GetLbInstancesOk() (*[]map[string]interface{}, bool)`

GetLbInstancesOk returns a tuple with the LbInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbInstances

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) SetLbInstances(v []map[string]interface{})`

SetLbInstances sets LbInstances field to given value.

### HasLbInstances

`func (o *ListInstances200ResponseAllOfInstancesInnerConfig) HasLbInstances() bool`

HasLbInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



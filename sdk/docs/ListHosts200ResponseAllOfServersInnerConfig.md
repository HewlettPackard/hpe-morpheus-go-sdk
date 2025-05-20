# ListHosts200ResponseAllOfServersInnerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolProviderType** | Pointer to **string** |  | [optional] 
**IsVpcSelectable** | Pointer to **bool** |  | [optional] 
**SmbiosAssetTag** | Pointer to **string** |  | [optional] 
**IsEC2** | Pointer to **bool** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**HostId** | Pointer to **int64** |  | [optional] 
**CreateUser** | Pointer to [**ListHosts200ResponseAllOfServersInnerConfigCreateUser**](ListHosts200ResponseAllOfServersInnerConfigCreateUser.md) |  | [optional] 
**NestedVirtualization** | Pointer to **string** |  | [optional] 
**VmwareFolderId** | Pointer to **string** |  | [optional] 
**NoAgent** | Pointer to **bool** |  | [optional] 
**PowerScheduleType** | Pointer to **int64** |  | [optional] 

## Methods

### NewListHosts200ResponseAllOfServersInnerConfig

`func NewListHosts200ResponseAllOfServersInnerConfig() *ListHosts200ResponseAllOfServersInnerConfig`

NewListHosts200ResponseAllOfServersInnerConfig instantiates a new ListHosts200ResponseAllOfServersInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHosts200ResponseAllOfServersInnerConfigWithDefaults

`func NewListHosts200ResponseAllOfServersInnerConfigWithDefaults() *ListHosts200ResponseAllOfServersInnerConfig`

NewListHosts200ResponseAllOfServersInnerConfigWithDefaults instantiates a new ListHosts200ResponseAllOfServersInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolProviderType

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetPoolProviderType() string`

GetPoolProviderType returns the PoolProviderType field if non-nil, zero value otherwise.

### GetPoolProviderTypeOk

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetPoolProviderTypeOk() (*string, bool)`

GetPoolProviderTypeOk returns a tuple with the PoolProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolProviderType

`func (o *ListHosts200ResponseAllOfServersInnerConfig) SetPoolProviderType(v string)`

SetPoolProviderType sets PoolProviderType field to given value.

### HasPoolProviderType

`func (o *ListHosts200ResponseAllOfServersInnerConfig) HasPoolProviderType() bool`

HasPoolProviderType returns a boolean if a field has been set.

### GetIsVpcSelectable

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetIsVpcSelectable() bool`

GetIsVpcSelectable returns the IsVpcSelectable field if non-nil, zero value otherwise.

### GetIsVpcSelectableOk

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetIsVpcSelectableOk() (*bool, bool)`

GetIsVpcSelectableOk returns a tuple with the IsVpcSelectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpcSelectable

`func (o *ListHosts200ResponseAllOfServersInnerConfig) SetIsVpcSelectable(v bool)`

SetIsVpcSelectable sets IsVpcSelectable field to given value.

### HasIsVpcSelectable

`func (o *ListHosts200ResponseAllOfServersInnerConfig) HasIsVpcSelectable() bool`

HasIsVpcSelectable returns a boolean if a field has been set.

### GetSmbiosAssetTag

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *ListHosts200ResponseAllOfServersInnerConfig) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *ListHosts200ResponseAllOfServersInnerConfig) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### GetIsEC2

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetIsEC2() bool`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetIsEC2Ok() (*bool, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *ListHosts200ResponseAllOfServersInnerConfig) SetIsEC2(v bool)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *ListHosts200ResponseAllOfServersInnerConfig) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ListHosts200ResponseAllOfServersInnerConfig) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ListHosts200ResponseAllOfServersInnerConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetHostId

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ListHosts200ResponseAllOfServersInnerConfig) SetHostId(v int64)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ListHosts200ResponseAllOfServersInnerConfig) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetCreateUser

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetCreateUser() ListHosts200ResponseAllOfServersInnerConfigCreateUser`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetCreateUserOk() (*ListHosts200ResponseAllOfServersInnerConfigCreateUser, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ListHosts200ResponseAllOfServersInnerConfig) SetCreateUser(v ListHosts200ResponseAllOfServersInnerConfigCreateUser)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ListHosts200ResponseAllOfServersInnerConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *ListHosts200ResponseAllOfServersInnerConfig) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *ListHosts200ResponseAllOfServersInnerConfig) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *ListHosts200ResponseAllOfServersInnerConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *ListHosts200ResponseAllOfServersInnerConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetNoAgent

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *ListHosts200ResponseAllOfServersInnerConfig) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *ListHosts200ResponseAllOfServersInnerConfig) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetPowerScheduleType

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetPowerScheduleType() int64`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *ListHosts200ResponseAllOfServersInnerConfig) GetPowerScheduleTypeOk() (*int64, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *ListHosts200ResponseAllOfServersInnerConfig) SetPowerScheduleType(v int64)`

SetPowerScheduleType sets PowerScheduleType field to given value.

### HasPowerScheduleType

`func (o *ListHosts200ResponseAllOfServersInnerConfig) HasPowerScheduleType() bool`

HasPowerScheduleType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



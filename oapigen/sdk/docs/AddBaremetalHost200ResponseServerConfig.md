# AddBaremetalHost200ResponseServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolProviderType** | Pointer to **NullableString** |  | [optional] 
**IsVpcSelectable** | Pointer to **bool** |  | [optional] 
**SmbiosAssetTag** | Pointer to **NullableString** |  | [optional] 
**IsEC2** | Pointer to **bool** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**HostId** | Pointer to **NullableInt64** |  | [optional] 
**CreateUser** | Pointer to [**AddBaremetalHost200ResponseServerConfigCreateUser**](AddBaremetalHost200ResponseServerConfigCreateUser.md) |  | [optional] 
**NestedVirtualization** | Pointer to **NullableString** |  | [optional] 
**VmwareFolderId** | Pointer to **string** |  | [optional] 
**NoAgent** | Pointer to **bool** |  | [optional] 
**PowerScheduleType** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewAddBaremetalHost200ResponseServerConfig

`func NewAddBaremetalHost200ResponseServerConfig() *AddBaremetalHost200ResponseServerConfig`

NewAddBaremetalHost200ResponseServerConfig instantiates a new AddBaremetalHost200ResponseServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBaremetalHost200ResponseServerConfigWithDefaults

`func NewAddBaremetalHost200ResponseServerConfigWithDefaults() *AddBaremetalHost200ResponseServerConfig`

NewAddBaremetalHost200ResponseServerConfigWithDefaults instantiates a new AddBaremetalHost200ResponseServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolProviderType

`func (o *AddBaremetalHost200ResponseServerConfig) GetPoolProviderType() string`

GetPoolProviderType returns the PoolProviderType field if non-nil, zero value otherwise.

### GetPoolProviderTypeOk

`func (o *AddBaremetalHost200ResponseServerConfig) GetPoolProviderTypeOk() (*string, bool)`

GetPoolProviderTypeOk returns a tuple with the PoolProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolProviderType

`func (o *AddBaremetalHost200ResponseServerConfig) SetPoolProviderType(v string)`

SetPoolProviderType sets PoolProviderType field to given value.

### HasPoolProviderType

`func (o *AddBaremetalHost200ResponseServerConfig) HasPoolProviderType() bool`

HasPoolProviderType returns a boolean if a field has been set.

### SetPoolProviderTypeNil

`func (o *AddBaremetalHost200ResponseServerConfig) SetPoolProviderTypeNil(b bool)`

 SetPoolProviderTypeNil sets the value for PoolProviderType to be an explicit nil

### UnsetPoolProviderType
`func (o *AddBaremetalHost200ResponseServerConfig) UnsetPoolProviderType()`

UnsetPoolProviderType ensures that no value is present for PoolProviderType, not even an explicit nil
### GetIsVpcSelectable

`func (o *AddBaremetalHost200ResponseServerConfig) GetIsVpcSelectable() bool`

GetIsVpcSelectable returns the IsVpcSelectable field if non-nil, zero value otherwise.

### GetIsVpcSelectableOk

`func (o *AddBaremetalHost200ResponseServerConfig) GetIsVpcSelectableOk() (*bool, bool)`

GetIsVpcSelectableOk returns a tuple with the IsVpcSelectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpcSelectable

`func (o *AddBaremetalHost200ResponseServerConfig) SetIsVpcSelectable(v bool)`

SetIsVpcSelectable sets IsVpcSelectable field to given value.

### HasIsVpcSelectable

`func (o *AddBaremetalHost200ResponseServerConfig) HasIsVpcSelectable() bool`

HasIsVpcSelectable returns a boolean if a field has been set.

### GetSmbiosAssetTag

`func (o *AddBaremetalHost200ResponseServerConfig) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *AddBaremetalHost200ResponseServerConfig) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *AddBaremetalHost200ResponseServerConfig) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *AddBaremetalHost200ResponseServerConfig) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### SetSmbiosAssetTagNil

`func (o *AddBaremetalHost200ResponseServerConfig) SetSmbiosAssetTagNil(b bool)`

 SetSmbiosAssetTagNil sets the value for SmbiosAssetTag to be an explicit nil

### UnsetSmbiosAssetTag
`func (o *AddBaremetalHost200ResponseServerConfig) UnsetSmbiosAssetTag()`

UnsetSmbiosAssetTag ensures that no value is present for SmbiosAssetTag, not even an explicit nil
### GetIsEC2

`func (o *AddBaremetalHost200ResponseServerConfig) GetIsEC2() bool`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *AddBaremetalHost200ResponseServerConfig) GetIsEC2Ok() (*bool, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *AddBaremetalHost200ResponseServerConfig) SetIsEC2(v bool)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *AddBaremetalHost200ResponseServerConfig) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *AddBaremetalHost200ResponseServerConfig) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *AddBaremetalHost200ResponseServerConfig) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *AddBaremetalHost200ResponseServerConfig) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *AddBaremetalHost200ResponseServerConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetHostId

`func (o *AddBaremetalHost200ResponseServerConfig) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AddBaremetalHost200ResponseServerConfig) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AddBaremetalHost200ResponseServerConfig) SetHostId(v int64)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AddBaremetalHost200ResponseServerConfig) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### SetHostIdNil

`func (o *AddBaremetalHost200ResponseServerConfig) SetHostIdNil(b bool)`

 SetHostIdNil sets the value for HostId to be an explicit nil

### UnsetHostId
`func (o *AddBaremetalHost200ResponseServerConfig) UnsetHostId()`

UnsetHostId ensures that no value is present for HostId, not even an explicit nil
### GetCreateUser

`func (o *AddBaremetalHost200ResponseServerConfig) GetCreateUser() AddBaremetalHost200ResponseServerConfigCreateUser`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddBaremetalHost200ResponseServerConfig) GetCreateUserOk() (*AddBaremetalHost200ResponseServerConfigCreateUser, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddBaremetalHost200ResponseServerConfig) SetCreateUser(v AddBaremetalHost200ResponseServerConfigCreateUser)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddBaremetalHost200ResponseServerConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *AddBaremetalHost200ResponseServerConfig) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *AddBaremetalHost200ResponseServerConfig) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *AddBaremetalHost200ResponseServerConfig) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *AddBaremetalHost200ResponseServerConfig) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### SetNestedVirtualizationNil

`func (o *AddBaremetalHost200ResponseServerConfig) SetNestedVirtualizationNil(b bool)`

 SetNestedVirtualizationNil sets the value for NestedVirtualization to be an explicit nil

### UnsetNestedVirtualization
`func (o *AddBaremetalHost200ResponseServerConfig) UnsetNestedVirtualization()`

UnsetNestedVirtualization ensures that no value is present for NestedVirtualization, not even an explicit nil
### GetVmwareFolderId

`func (o *AddBaremetalHost200ResponseServerConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *AddBaremetalHost200ResponseServerConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *AddBaremetalHost200ResponseServerConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *AddBaremetalHost200ResponseServerConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetNoAgent

`func (o *AddBaremetalHost200ResponseServerConfig) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *AddBaremetalHost200ResponseServerConfig) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *AddBaremetalHost200ResponseServerConfig) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *AddBaremetalHost200ResponseServerConfig) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetPowerScheduleType

`func (o *AddBaremetalHost200ResponseServerConfig) GetPowerScheduleType() int64`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *AddBaremetalHost200ResponseServerConfig) GetPowerScheduleTypeOk() (*int64, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *AddBaremetalHost200ResponseServerConfig) SetPowerScheduleType(v int64)`

SetPowerScheduleType sets PowerScheduleType field to given value.

### HasPowerScheduleType

`func (o *AddBaremetalHost200ResponseServerConfig) HasPowerScheduleType() bool`

HasPowerScheduleType returns a boolean if a field has been set.

### SetPowerScheduleTypeNil

`func (o *AddBaremetalHost200ResponseServerConfig) SetPowerScheduleTypeNil(b bool)`

 SetPowerScheduleTypeNil sets the value for PowerScheduleType to be an explicit nil

### UnsetPowerScheduleType
`func (o *AddBaremetalHost200ResponseServerConfig) UnsetPowerScheduleType()`

UnsetPowerScheduleType ensures that no value is present for PowerScheduleType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



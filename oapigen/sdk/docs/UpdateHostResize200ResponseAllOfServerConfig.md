# UpdateHostResize200ResponseAllOfServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolProviderType** | Pointer to **NullableString** |  | [optional] 
**IsVpcSelectable** | Pointer to **bool** |  | [optional] 
**SmbiosAssetTag** | Pointer to **NullableString** |  | [optional] 
**IsEC2** | Pointer to **bool** |  | [optional] 
**ResourcePoolId** | Pointer to [**UpdateHostResize200ResponseAllOfServerConfigResourcePoolId**](UpdateHostResize200ResponseAllOfServerConfigResourcePoolId.md) |  | [optional] 
**HostId** | Pointer to **NullableInt64** |  | [optional] 
**CreateUser** | Pointer to [**UpdateHostResize200ResponseAllOfServerConfigCreateUser**](UpdateHostResize200ResponseAllOfServerConfigCreateUser.md) |  | [optional] 
**NestedVirtualization** | Pointer to **NullableString** |  | [optional] 
**VmwareFolderId** | Pointer to **string** |  | [optional] 
**NoAgent** | Pointer to **bool** |  | [optional] 
**PowerScheduleType** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewUpdateHostResize200ResponseAllOfServerConfig

`func NewUpdateHostResize200ResponseAllOfServerConfig() *UpdateHostResize200ResponseAllOfServerConfig`

NewUpdateHostResize200ResponseAllOfServerConfig instantiates a new UpdateHostResize200ResponseAllOfServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostResize200ResponseAllOfServerConfigWithDefaults

`func NewUpdateHostResize200ResponseAllOfServerConfigWithDefaults() *UpdateHostResize200ResponseAllOfServerConfig`

NewUpdateHostResize200ResponseAllOfServerConfigWithDefaults instantiates a new UpdateHostResize200ResponseAllOfServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolProviderType

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetPoolProviderType() string`

GetPoolProviderType returns the PoolProviderType field if non-nil, zero value otherwise.

### GetPoolProviderTypeOk

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetPoolProviderTypeOk() (*string, bool)`

GetPoolProviderTypeOk returns a tuple with the PoolProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolProviderType

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetPoolProviderType(v string)`

SetPoolProviderType sets PoolProviderType field to given value.

### HasPoolProviderType

`func (o *UpdateHostResize200ResponseAllOfServerConfig) HasPoolProviderType() bool`

HasPoolProviderType returns a boolean if a field has been set.

### SetPoolProviderTypeNil

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetPoolProviderTypeNil(b bool)`

 SetPoolProviderTypeNil sets the value for PoolProviderType to be an explicit nil

### UnsetPoolProviderType
`func (o *UpdateHostResize200ResponseAllOfServerConfig) UnsetPoolProviderType()`

UnsetPoolProviderType ensures that no value is present for PoolProviderType, not even an explicit nil
### GetIsVpcSelectable

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetIsVpcSelectable() bool`

GetIsVpcSelectable returns the IsVpcSelectable field if non-nil, zero value otherwise.

### GetIsVpcSelectableOk

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetIsVpcSelectableOk() (*bool, bool)`

GetIsVpcSelectableOk returns a tuple with the IsVpcSelectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpcSelectable

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetIsVpcSelectable(v bool)`

SetIsVpcSelectable sets IsVpcSelectable field to given value.

### HasIsVpcSelectable

`func (o *UpdateHostResize200ResponseAllOfServerConfig) HasIsVpcSelectable() bool`

HasIsVpcSelectable returns a boolean if a field has been set.

### GetSmbiosAssetTag

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *UpdateHostResize200ResponseAllOfServerConfig) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### SetSmbiosAssetTagNil

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetSmbiosAssetTagNil(b bool)`

 SetSmbiosAssetTagNil sets the value for SmbiosAssetTag to be an explicit nil

### UnsetSmbiosAssetTag
`func (o *UpdateHostResize200ResponseAllOfServerConfig) UnsetSmbiosAssetTag()`

UnsetSmbiosAssetTag ensures that no value is present for SmbiosAssetTag, not even an explicit nil
### GetIsEC2

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetIsEC2() bool`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetIsEC2Ok() (*bool, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetIsEC2(v bool)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *UpdateHostResize200ResponseAllOfServerConfig) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetResourcePoolId() UpdateHostResize200ResponseAllOfServerConfigResourcePoolId`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetResourcePoolIdOk() (*UpdateHostResize200ResponseAllOfServerConfigResourcePoolId, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetResourcePoolId(v UpdateHostResize200ResponseAllOfServerConfigResourcePoolId)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *UpdateHostResize200ResponseAllOfServerConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetHostId

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetHostId(v int64)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *UpdateHostResize200ResponseAllOfServerConfig) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### SetHostIdNil

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetHostIdNil(b bool)`

 SetHostIdNil sets the value for HostId to be an explicit nil

### UnsetHostId
`func (o *UpdateHostResize200ResponseAllOfServerConfig) UnsetHostId()`

UnsetHostId ensures that no value is present for HostId, not even an explicit nil
### GetCreateUser

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetCreateUser() UpdateHostResize200ResponseAllOfServerConfigCreateUser`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetCreateUserOk() (*UpdateHostResize200ResponseAllOfServerConfigCreateUser, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetCreateUser(v UpdateHostResize200ResponseAllOfServerConfigCreateUser)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *UpdateHostResize200ResponseAllOfServerConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *UpdateHostResize200ResponseAllOfServerConfig) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### SetNestedVirtualizationNil

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetNestedVirtualizationNil(b bool)`

 SetNestedVirtualizationNil sets the value for NestedVirtualization to be an explicit nil

### UnsetNestedVirtualization
`func (o *UpdateHostResize200ResponseAllOfServerConfig) UnsetNestedVirtualization()`

UnsetNestedVirtualization ensures that no value is present for NestedVirtualization, not even an explicit nil
### GetVmwareFolderId

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *UpdateHostResize200ResponseAllOfServerConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetNoAgent

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *UpdateHostResize200ResponseAllOfServerConfig) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetPowerScheduleType

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetPowerScheduleType() int64`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *UpdateHostResize200ResponseAllOfServerConfig) GetPowerScheduleTypeOk() (*int64, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetPowerScheduleType(v int64)`

SetPowerScheduleType sets PowerScheduleType field to given value.

### HasPowerScheduleType

`func (o *UpdateHostResize200ResponseAllOfServerConfig) HasPowerScheduleType() bool`

HasPowerScheduleType returns a boolean if a field has been set.

### SetPowerScheduleTypeNil

`func (o *UpdateHostResize200ResponseAllOfServerConfig) SetPowerScheduleTypeNil(b bool)`

 SetPowerScheduleTypeNil sets the value for PowerScheduleType to be an explicit nil

### UnsetPowerScheduleType
`func (o *UpdateHostResize200ResponseAllOfServerConfig) UnsetPowerScheduleType()`

UnsetPowerScheduleType ensures that no value is present for PowerScheduleType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



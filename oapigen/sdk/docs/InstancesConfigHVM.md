# InstancesConfigHVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoAgent** | Pointer to **NullableBool** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to false]
**ResourcePoolId** | Pointer to **string** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. | [optional] 
**NestedVirtualization** | Pointer to **string** | Enable Nested Virtualization | [optional] [default to "off"]
**CreateUser** | Pointer to **NullableBool** | Create user | [optional] [default to false]
**PoolProviderType** | Pointer to **string** | The type of pool provider to use for this instance, must be \&quot;mvm\&quot; | [optional] [default to "mvm"]
**KvmHostId** | Pointer to **int64** | The ID of the KVM host to provision the instance on | [optional] 

## Methods

### NewInstancesConfigHVM

`func NewInstancesConfigHVM() *InstancesConfigHVM`

NewInstancesConfigHVM instantiates a new InstancesConfigHVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesConfigHVMWithDefaults

`func NewInstancesConfigHVMWithDefaults() *InstancesConfigHVM`

NewInstancesConfigHVMWithDefaults instantiates a new InstancesConfigHVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoAgent

`func (o *InstancesConfigHVM) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *InstancesConfigHVM) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *InstancesConfigHVM) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *InstancesConfigHVM) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### SetNoAgentNil

`func (o *InstancesConfigHVM) SetNoAgentNil(b bool)`

 SetNoAgentNil sets the value for NoAgent to be an explicit nil

### UnsetNoAgent
`func (o *InstancesConfigHVM) UnsetNoAgent()`

UnsetNoAgent ensures that no value is present for NoAgent, not even an explicit nil
### GetResourcePoolId

`func (o *InstancesConfigHVM) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *InstancesConfigHVM) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *InstancesConfigHVM) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *InstancesConfigHVM) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *InstancesConfigHVM) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *InstancesConfigHVM) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *InstancesConfigHVM) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *InstancesConfigHVM) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetCreateUser

`func (o *InstancesConfigHVM) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *InstancesConfigHVM) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *InstancesConfigHVM) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *InstancesConfigHVM) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### SetCreateUserNil

`func (o *InstancesConfigHVM) SetCreateUserNil(b bool)`

 SetCreateUserNil sets the value for CreateUser to be an explicit nil

### UnsetCreateUser
`func (o *InstancesConfigHVM) UnsetCreateUser()`

UnsetCreateUser ensures that no value is present for CreateUser, not even an explicit nil
### GetPoolProviderType

`func (o *InstancesConfigHVM) GetPoolProviderType() string`

GetPoolProviderType returns the PoolProviderType field if non-nil, zero value otherwise.

### GetPoolProviderTypeOk

`func (o *InstancesConfigHVM) GetPoolProviderTypeOk() (*string, bool)`

GetPoolProviderTypeOk returns a tuple with the PoolProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolProviderType

`func (o *InstancesConfigHVM) SetPoolProviderType(v string)`

SetPoolProviderType sets PoolProviderType field to given value.

### HasPoolProviderType

`func (o *InstancesConfigHVM) HasPoolProviderType() bool`

HasPoolProviderType returns a boolean if a field has been set.

### GetKvmHostId

`func (o *InstancesConfigHVM) GetKvmHostId() int64`

GetKvmHostId returns the KvmHostId field if non-nil, zero value otherwise.

### GetKvmHostIdOk

`func (o *InstancesConfigHVM) GetKvmHostIdOk() (*int64, bool)`

GetKvmHostIdOk returns a tuple with the KvmHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmHostId

`func (o *InstancesConfigHVM) SetKvmHostId(v int64)`

SetKvmHostId sets KvmHostId field to given value.

### HasKvmHostId

`func (o *InstancesConfigHVM) HasKvmHostId() bool`

HasKvmHostId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



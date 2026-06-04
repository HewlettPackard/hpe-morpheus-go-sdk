# CreateLoadBalancerPoolRequestLoadBalancerPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**VipBalance** | Pointer to **string** | Balance Algorithm | [optional] 
**MinActive** | Pointer to **int64** | Min Active Members | [optional] 
**Port** | Pointer to **int64** | Port number | [optional] 
**VipSticky** | Pointer to **string** | Session persistence mode | [optional] 
**VipClientIpMode** | Pointer to **string** | VIP client IP mode | [optional] 
**Partition** | Pointer to **string** | Partition | [optional] 
**Config** | Pointer to [**CreateLoadBalancerPoolRequestLoadBalancerPoolConfig**](CreateLoadBalancerPoolRequestLoadBalancerPoolConfig.md) |  | [optional] 

## Methods

### NewCreateLoadBalancerPoolRequestLoadBalancerPool

`func NewCreateLoadBalancerPoolRequestLoadBalancerPool() *CreateLoadBalancerPoolRequestLoadBalancerPool`

NewCreateLoadBalancerPoolRequestLoadBalancerPool instantiates a new CreateLoadBalancerPoolRequestLoadBalancerPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerPoolRequestLoadBalancerPoolWithDefaults

`func NewCreateLoadBalancerPoolRequestLoadBalancerPoolWithDefaults() *CreateLoadBalancerPoolRequestLoadBalancerPool`

NewCreateLoadBalancerPoolRequestLoadBalancerPoolWithDefaults instantiates a new CreateLoadBalancerPoolRequestLoadBalancerPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVipBalance

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetVipBalance() string`

GetVipBalance returns the VipBalance field if non-nil, zero value otherwise.

### GetVipBalanceOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetVipBalanceOk() (*string, bool)`

GetVipBalanceOk returns a tuple with the VipBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipBalance

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) SetVipBalance(v string)`

SetVipBalance sets VipBalance field to given value.

### HasVipBalance

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) HasVipBalance() bool`

HasVipBalance returns a boolean if a field has been set.

### GetMinActive

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetMinActive() int64`

GetMinActive returns the MinActive field if non-nil, zero value otherwise.

### GetMinActiveOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetMinActiveOk() (*int64, bool)`

GetMinActiveOk returns a tuple with the MinActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinActive

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) SetMinActive(v int64)`

SetMinActive sets MinActive field to given value.

### HasMinActive

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) HasMinActive() bool`

HasMinActive returns a boolean if a field has been set.

### GetPort

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetVipSticky

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetVipSticky() string`

GetVipSticky returns the VipSticky field if non-nil, zero value otherwise.

### GetVipStickyOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetVipStickyOk() (*string, bool)`

GetVipStickyOk returns a tuple with the VipSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSticky

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) SetVipSticky(v string)`

SetVipSticky sets VipSticky field to given value.

### HasVipSticky

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) HasVipSticky() bool`

HasVipSticky returns a boolean if a field has been set.

### GetVipClientIpMode

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetVipClientIpMode() string`

GetVipClientIpMode returns the VipClientIpMode field if non-nil, zero value otherwise.

### GetVipClientIpModeOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetVipClientIpModeOk() (*string, bool)`

GetVipClientIpModeOk returns a tuple with the VipClientIpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipClientIpMode

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) SetVipClientIpMode(v string)`

SetVipClientIpMode sets VipClientIpMode field to given value.

### HasVipClientIpMode

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) HasVipClientIpMode() bool`

HasVipClientIpMode returns a boolean if a field has been set.

### GetPartition

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetConfig

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetConfig() CreateLoadBalancerPoolRequestLoadBalancerPoolConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetConfigOk() (*CreateLoadBalancerPoolRequestLoadBalancerPoolConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) SetConfig(v CreateLoadBalancerPoolRequestLoadBalancerPoolConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



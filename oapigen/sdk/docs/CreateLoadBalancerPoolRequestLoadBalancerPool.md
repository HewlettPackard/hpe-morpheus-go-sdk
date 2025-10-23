# CreateLoadBalancerPoolRequestLoadBalancerPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**VipBalance** | Pointer to **string** | Balance Algorithm | [optional] 
**MinActive** | Pointer to **int64** | Min Active Members | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by type. | [optional] 

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

### GetConfig

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



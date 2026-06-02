# CreateLoadBalancerRequestLoadBalancerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to [**HAProxyLoadBalancerConfigObjectPlan**](HAProxyLoadBalancerConfigObjectPlan.md) |  | [optional] 
**Pool** | Pointer to [**HAProxyLoadBalancerConfigObjectPool**](HAProxyLoadBalancerConfigObjectPool.md) |  | [optional] 
**AdminState** | Pointer to **bool** | If true, the admin state is enabled. | [optional] 
**Size** | Pointer to **string** | Load balancer size. | [optional] 
**Tier1** | Pointer to **string** | Tier 1 gateway provider ID. | [optional] 
**Loglevel** | Pointer to **string** | Log level. | [optional] 

## Methods

### NewCreateLoadBalancerRequestLoadBalancerConfig

`func NewCreateLoadBalancerRequestLoadBalancerConfig() *CreateLoadBalancerRequestLoadBalancerConfig`

NewCreateLoadBalancerRequestLoadBalancerConfig instantiates a new CreateLoadBalancerRequestLoadBalancerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetPlan

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) GetPlan() HAProxyLoadBalancerConfigObjectPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) GetPlanOk() (*HAProxyLoadBalancerConfigObjectPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) SetPlan(v HAProxyLoadBalancerConfigObjectPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPool

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) GetPool() HAProxyLoadBalancerConfigObjectPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) GetPoolOk() (*HAProxyLoadBalancerConfigObjectPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) SetPool(v HAProxyLoadBalancerConfigObjectPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetAdminState

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) GetAdminState() bool`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) GetAdminStateOk() (*bool, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) SetAdminState(v bool)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetSize

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTier1

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) GetTier1() string`

GetTier1 returns the Tier1 field if non-nil, zero value otherwise.

### GetTier1Ok

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) GetTier1Ok() (*string, bool)`

GetTier1Ok returns a tuple with the Tier1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier1

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) SetTier1(v string)`

SetTier1 sets Tier1 field to given value.

### HasTier1

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) HasTier1() bool`

HasTier1 returns a boolean if a field has been set.

### GetLoglevel

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) GetLoglevel() string`

GetLoglevel returns the Loglevel field if non-nil, zero value otherwise.

### GetLoglevelOk

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) GetLoglevelOk() (*string, bool)`

GetLoglevelOk returns a tuple with the Loglevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoglevel

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) SetLoglevel(v string)`

SetLoglevel sets Loglevel field to given value.

### HasLoglevel

`func (o *CreateLoadBalancerRequestLoadBalancerConfig) HasLoglevel() bool`

HasLoglevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



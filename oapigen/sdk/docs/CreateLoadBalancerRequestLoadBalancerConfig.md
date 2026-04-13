# CreateLoadBalancerRequestLoadBalancerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to [**HAProxyLoadBalancerConfigObjectPlan**](HAProxyLoadBalancerConfigObjectPlan.md) |  | [optional] 
**Pool** | Pointer to [**HAProxyLoadBalancerConfigObjectPool**](HAProxyLoadBalancerConfigObjectPool.md) |  | [optional] 

## Methods

### NewCreateLoadBalancerRequestLoadBalancerConfig

`func NewCreateLoadBalancerRequestLoadBalancerConfig() *CreateLoadBalancerRequestLoadBalancerConfig`

NewCreateLoadBalancerRequestLoadBalancerConfig instantiates a new CreateLoadBalancerRequestLoadBalancerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerRequestLoadBalancerConfigWithDefaults

`func NewCreateLoadBalancerRequestLoadBalancerConfigWithDefaults() *CreateLoadBalancerRequestLoadBalancerConfig`

NewCreateLoadBalancerRequestLoadBalancerConfigWithDefaults instantiates a new CreateLoadBalancerRequestLoadBalancerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



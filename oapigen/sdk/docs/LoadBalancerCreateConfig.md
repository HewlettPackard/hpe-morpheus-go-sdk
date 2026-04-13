# LoadBalancerCreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to [**HAProxyLoadBalancerConfigObject1Plan**](HAProxyLoadBalancerConfigObject1Plan.md) |  | [optional] 
**Pool** | Pointer to [**HAProxyLoadBalancerConfigObject1Pool**](HAProxyLoadBalancerConfigObject1Pool.md) |  | [optional] 

## Methods

### NewLoadBalancerCreateConfig

`func NewLoadBalancerCreateConfig() *LoadBalancerCreateConfig`

NewLoadBalancerCreateConfig instantiates a new LoadBalancerCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerCreateConfigWithDefaults

`func NewLoadBalancerCreateConfigWithDefaults() *LoadBalancerCreateConfig`

NewLoadBalancerCreateConfigWithDefaults instantiates a new LoadBalancerCreateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *LoadBalancerCreateConfig) GetPlan() HAProxyLoadBalancerConfigObject1Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *LoadBalancerCreateConfig) GetPlanOk() (*HAProxyLoadBalancerConfigObject1Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *LoadBalancerCreateConfig) SetPlan(v HAProxyLoadBalancerConfigObject1Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *LoadBalancerCreateConfig) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPool

`func (o *LoadBalancerCreateConfig) GetPool() HAProxyLoadBalancerConfigObject1Pool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *LoadBalancerCreateConfig) GetPoolOk() (*HAProxyLoadBalancerConfigObject1Pool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *LoadBalancerCreateConfig) SetPool(v HAProxyLoadBalancerConfigObject1Pool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *LoadBalancerCreateConfig) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LoadBalancerCreateConfigHaproxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to [**LoadBalancerCreateConfigHaproxyPlan**](LoadBalancerCreateConfigHaproxyPlan.md) |  | [optional] 
**Pool** | Pointer to [**LoadBalancerCreateConfigHaproxyPool**](LoadBalancerCreateConfigHaproxyPool.md) |  | [optional] 

## Methods

### NewLoadBalancerCreateConfigHaproxy

`func NewLoadBalancerCreateConfigHaproxy() *LoadBalancerCreateConfigHaproxy`

NewLoadBalancerCreateConfigHaproxy instantiates a new LoadBalancerCreateConfigHaproxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerCreateConfigHaproxyWithDefaults

`func NewLoadBalancerCreateConfigHaproxyWithDefaults() *LoadBalancerCreateConfigHaproxy`

NewLoadBalancerCreateConfigHaproxyWithDefaults instantiates a new LoadBalancerCreateConfigHaproxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *LoadBalancerCreateConfigHaproxy) GetPlan() LoadBalancerCreateConfigHaproxyPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *LoadBalancerCreateConfigHaproxy) GetPlanOk() (*LoadBalancerCreateConfigHaproxyPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *LoadBalancerCreateConfigHaproxy) SetPlan(v LoadBalancerCreateConfigHaproxyPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *LoadBalancerCreateConfigHaproxy) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPool

`func (o *LoadBalancerCreateConfigHaproxy) GetPool() LoadBalancerCreateConfigHaproxyPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *LoadBalancerCreateConfigHaproxy) GetPoolOk() (*LoadBalancerCreateConfigHaproxyPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *LoadBalancerCreateConfigHaproxy) SetPool(v LoadBalancerCreateConfigHaproxyPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *LoadBalancerCreateConfigHaproxy) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# HAProxyLoadBalancerConfigObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to [**HAProxyLoadBalancerConfigObjectPlan**](HAProxyLoadBalancerConfigObjectPlan.md) |  | [optional] 
**Pool** | Pointer to [**HAProxyLoadBalancerConfigObjectPool**](HAProxyLoadBalancerConfigObjectPool.md) |  | [optional] 

## Methods

### NewHAProxyLoadBalancerConfigObject

`func NewHAProxyLoadBalancerConfigObject() *HAProxyLoadBalancerConfigObject`

NewHAProxyLoadBalancerConfigObject instantiates a new HAProxyLoadBalancerConfigObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHAProxyLoadBalancerConfigObjectWithDefaults

`func NewHAProxyLoadBalancerConfigObjectWithDefaults() *HAProxyLoadBalancerConfigObject`

NewHAProxyLoadBalancerConfigObjectWithDefaults instantiates a new HAProxyLoadBalancerConfigObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *HAProxyLoadBalancerConfigObject) GetPlan() HAProxyLoadBalancerConfigObjectPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *HAProxyLoadBalancerConfigObject) GetPlanOk() (*HAProxyLoadBalancerConfigObjectPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *HAProxyLoadBalancerConfigObject) SetPlan(v HAProxyLoadBalancerConfigObjectPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *HAProxyLoadBalancerConfigObject) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPool

`func (o *HAProxyLoadBalancerConfigObject) GetPool() HAProxyLoadBalancerConfigObjectPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *HAProxyLoadBalancerConfigObject) GetPoolOk() (*HAProxyLoadBalancerConfigObjectPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *HAProxyLoadBalancerConfigObject) SetPool(v HAProxyLoadBalancerConfigObjectPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *HAProxyLoadBalancerConfigObject) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



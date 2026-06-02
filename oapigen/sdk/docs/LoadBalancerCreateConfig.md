# LoadBalancerCreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to [**HAProxyLoadBalancerConfigObject1Plan**](HAProxyLoadBalancerConfigObject1Plan.md) |  | [optional] 
**Pool** | Pointer to [**HAProxyLoadBalancerConfigObject1Pool**](HAProxyLoadBalancerConfigObject1Pool.md) |  | [optional] 
**AdminState** | Pointer to **bool** | If true, the admin state is enabled. | [optional] 
**Size** | Pointer to **string** | Load balancer size. | [optional] 
**Tier1** | Pointer to **string** | Tier 1 gateway provider ID. | [optional] 
**Loglevel** | Pointer to **string** | Log level. | [optional] 

## Methods

### NewLoadBalancerCreateConfig

`func NewLoadBalancerCreateConfig() *LoadBalancerCreateConfig`

NewLoadBalancerCreateConfig instantiates a new LoadBalancerCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

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

### GetAdminState

`func (o *LoadBalancerCreateConfig) GetAdminState() bool`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *LoadBalancerCreateConfig) GetAdminStateOk() (*bool, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *LoadBalancerCreateConfig) SetAdminState(v bool)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *LoadBalancerCreateConfig) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetSize

`func (o *LoadBalancerCreateConfig) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LoadBalancerCreateConfig) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LoadBalancerCreateConfig) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *LoadBalancerCreateConfig) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTier1

`func (o *LoadBalancerCreateConfig) GetTier1() string`

GetTier1 returns the Tier1 field if non-nil, zero value otherwise.

### GetTier1Ok

`func (o *LoadBalancerCreateConfig) GetTier1Ok() (*string, bool)`

GetTier1Ok returns a tuple with the Tier1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier1

`func (o *LoadBalancerCreateConfig) SetTier1(v string)`

SetTier1 sets Tier1 field to given value.

### HasTier1

`func (o *LoadBalancerCreateConfig) HasTier1() bool`

HasTier1 returns a boolean if a field has been set.

### GetLoglevel

`func (o *LoadBalancerCreateConfig) GetLoglevel() string`

GetLoglevel returns the Loglevel field if non-nil, zero value otherwise.

### GetLoglevelOk

`func (o *LoadBalancerCreateConfig) GetLoglevelOk() (*string, bool)`

GetLoglevelOk returns a tuple with the Loglevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoglevel

`func (o *LoadBalancerCreateConfig) SetLoglevel(v string)`

SetLoglevel sets Loglevel field to given value.

### HasLoglevel

`func (o *LoadBalancerCreateConfig) HasLoglevel() bool`

HasLoglevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



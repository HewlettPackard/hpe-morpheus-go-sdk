# ListLoadBalancerVirtualServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerInstances** | Pointer to [**[]ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListLoadBalancerVirtualServers200Response

`func NewListLoadBalancerVirtualServers200Response() *ListLoadBalancerVirtualServers200Response`

NewListLoadBalancerVirtualServers200Response instantiates a new ListLoadBalancerVirtualServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerVirtualServers200ResponseWithDefaults

`func NewListLoadBalancerVirtualServers200ResponseWithDefaults() *ListLoadBalancerVirtualServers200Response`

NewListLoadBalancerVirtualServers200ResponseWithDefaults instantiates a new ListLoadBalancerVirtualServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerInstances

`func (o *ListLoadBalancerVirtualServers200Response) GetLoadBalancerInstances() []ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner`

GetLoadBalancerInstances returns the LoadBalancerInstances field if non-nil, zero value otherwise.

### GetLoadBalancerInstancesOk

`func (o *ListLoadBalancerVirtualServers200Response) GetLoadBalancerInstancesOk() (*[]ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner, bool)`

GetLoadBalancerInstancesOk returns a tuple with the LoadBalancerInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerInstances

`func (o *ListLoadBalancerVirtualServers200Response) SetLoadBalancerInstances(v []ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner)`

SetLoadBalancerInstances sets LoadBalancerInstances field to given value.

### HasLoadBalancerInstances

`func (o *ListLoadBalancerVirtualServers200Response) HasLoadBalancerInstances() bool`

HasLoadBalancerInstances returns a boolean if a field has been set.

### GetMeta

`func (o *ListLoadBalancerVirtualServers200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLoadBalancerVirtualServers200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLoadBalancerVirtualServers200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListLoadBalancerVirtualServers200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



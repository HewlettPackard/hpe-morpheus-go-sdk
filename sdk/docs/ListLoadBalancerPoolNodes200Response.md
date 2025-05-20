# ListLoadBalancerPoolNodes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerNodes** | Pointer to [**[]ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner**](ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListLoadBalancerPoolNodes200Response

`func NewListLoadBalancerPoolNodes200Response() *ListLoadBalancerPoolNodes200Response`

NewListLoadBalancerPoolNodes200Response instantiates a new ListLoadBalancerPoolNodes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerPoolNodes200ResponseWithDefaults

`func NewListLoadBalancerPoolNodes200ResponseWithDefaults() *ListLoadBalancerPoolNodes200Response`

NewListLoadBalancerPoolNodes200ResponseWithDefaults instantiates a new ListLoadBalancerPoolNodes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerNodes

`func (o *ListLoadBalancerPoolNodes200Response) GetLoadBalancerNodes() []ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner`

GetLoadBalancerNodes returns the LoadBalancerNodes field if non-nil, zero value otherwise.

### GetLoadBalancerNodesOk

`func (o *ListLoadBalancerPoolNodes200Response) GetLoadBalancerNodesOk() (*[]ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner, bool)`

GetLoadBalancerNodesOk returns a tuple with the LoadBalancerNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerNodes

`func (o *ListLoadBalancerPoolNodes200Response) SetLoadBalancerNodes(v []ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner)`

SetLoadBalancerNodes sets LoadBalancerNodes field to given value.

### HasLoadBalancerNodes

`func (o *ListLoadBalancerPoolNodes200Response) HasLoadBalancerNodes() bool`

HasLoadBalancerNodes returns a boolean if a field has been set.

### GetMeta

`func (o *ListLoadBalancerPoolNodes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLoadBalancerPoolNodes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLoadBalancerPoolNodes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListLoadBalancerPoolNodes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



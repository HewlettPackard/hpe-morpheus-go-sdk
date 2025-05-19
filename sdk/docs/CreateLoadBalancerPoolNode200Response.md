# CreateLoadBalancerPoolNode200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerNode** | Pointer to [**ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner**](ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateLoadBalancerPoolNode200Response

`func NewCreateLoadBalancerPoolNode200Response() *CreateLoadBalancerPoolNode200Response`

NewCreateLoadBalancerPoolNode200Response instantiates a new CreateLoadBalancerPoolNode200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerPoolNode200ResponseWithDefaults

`func NewCreateLoadBalancerPoolNode200ResponseWithDefaults() *CreateLoadBalancerPoolNode200Response`

NewCreateLoadBalancerPoolNode200ResponseWithDefaults instantiates a new CreateLoadBalancerPoolNode200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerNode

`func (o *CreateLoadBalancerPoolNode200Response) GetLoadBalancerNode() ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner`

GetLoadBalancerNode returns the LoadBalancerNode field if non-nil, zero value otherwise.

### GetLoadBalancerNodeOk

`func (o *CreateLoadBalancerPoolNode200Response) GetLoadBalancerNodeOk() (*ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner, bool)`

GetLoadBalancerNodeOk returns a tuple with the LoadBalancerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerNode

`func (o *CreateLoadBalancerPoolNode200Response) SetLoadBalancerNode(v ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner)`

SetLoadBalancerNode sets LoadBalancerNode field to given value.

### HasLoadBalancerNode

`func (o *CreateLoadBalancerPoolNode200Response) HasLoadBalancerNode() bool`

HasLoadBalancerNode returns a boolean if a field has been set.

### GetSuccess

`func (o *CreateLoadBalancerPoolNode200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateLoadBalancerPoolNode200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateLoadBalancerPoolNode200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateLoadBalancerPoolNode200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *CreateLoadBalancerPoolNode200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateLoadBalancerPoolNode200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateLoadBalancerPoolNode200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateLoadBalancerPoolNode200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



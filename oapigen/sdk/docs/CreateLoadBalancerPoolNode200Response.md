# CreateLoadBalancerPoolNode200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerNode** | Pointer to [**CreateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode**](CreateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateLoadBalancerPoolNode200Response

`func NewCreateLoadBalancerPoolNode200Response() *CreateLoadBalancerPoolNode200Response`

NewCreateLoadBalancerPoolNode200Response instantiates a new CreateLoadBalancerPoolNode200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetLoadBalancerNode

`func (o *CreateLoadBalancerPoolNode200Response) GetLoadBalancerNode() CreateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode`

GetLoadBalancerNode returns the LoadBalancerNode field if non-nil, zero value otherwise.

### GetLoadBalancerNodeOk

`func (o *CreateLoadBalancerPoolNode200Response) GetLoadBalancerNodeOk() (*CreateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode, bool)`

GetLoadBalancerNodeOk returns a tuple with the LoadBalancerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerNode

`func (o *CreateLoadBalancerPoolNode200Response) SetLoadBalancerNode(v CreateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode)`

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

### SetMsgNil

`func (o *CreateLoadBalancerPoolNode200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *CreateLoadBalancerPoolNode200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



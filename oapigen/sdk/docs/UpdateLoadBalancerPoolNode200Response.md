# UpdateLoadBalancerPoolNode200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerNode** | Pointer to [**UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode**](UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateLoadBalancerPoolNode200Response

`func NewUpdateLoadBalancerPoolNode200Response() *UpdateLoadBalancerPoolNode200Response`

NewUpdateLoadBalancerPoolNode200Response instantiates a new UpdateLoadBalancerPoolNode200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerPoolNode200ResponseWithDefaults

`func NewUpdateLoadBalancerPoolNode200ResponseWithDefaults() *UpdateLoadBalancerPoolNode200Response`

NewUpdateLoadBalancerPoolNode200ResponseWithDefaults instantiates a new UpdateLoadBalancerPoolNode200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerNode

`func (o *UpdateLoadBalancerPoolNode200Response) GetLoadBalancerNode() UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode`

GetLoadBalancerNode returns the LoadBalancerNode field if non-nil, zero value otherwise.

### GetLoadBalancerNodeOk

`func (o *UpdateLoadBalancerPoolNode200Response) GetLoadBalancerNodeOk() (*UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode, bool)`

GetLoadBalancerNodeOk returns a tuple with the LoadBalancerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerNode

`func (o *UpdateLoadBalancerPoolNode200Response) SetLoadBalancerNode(v UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode)`

SetLoadBalancerNode sets LoadBalancerNode field to given value.

### HasLoadBalancerNode

`func (o *UpdateLoadBalancerPoolNode200Response) HasLoadBalancerNode() bool`

HasLoadBalancerNode returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateLoadBalancerPoolNode200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateLoadBalancerPoolNode200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateLoadBalancerPoolNode200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateLoadBalancerPoolNode200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *UpdateLoadBalancerPoolNode200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UpdateLoadBalancerPoolNode200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UpdateLoadBalancerPoolNode200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *UpdateLoadBalancerPoolNode200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *UpdateLoadBalancerPoolNode200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *UpdateLoadBalancerPoolNode200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



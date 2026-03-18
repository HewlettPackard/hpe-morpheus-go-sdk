# UpdateLoadBalancerPool200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerPool** | Pointer to [**UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool**](UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateLoadBalancerPool200Response

`func NewUpdateLoadBalancerPool200Response() *UpdateLoadBalancerPool200Response`

NewUpdateLoadBalancerPool200Response instantiates a new UpdateLoadBalancerPool200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerPool200ResponseWithDefaults

`func NewUpdateLoadBalancerPool200ResponseWithDefaults() *UpdateLoadBalancerPool200Response`

NewUpdateLoadBalancerPool200ResponseWithDefaults instantiates a new UpdateLoadBalancerPool200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerPool

`func (o *UpdateLoadBalancerPool200Response) GetLoadBalancerPool() UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool`

GetLoadBalancerPool returns the LoadBalancerPool field if non-nil, zero value otherwise.

### GetLoadBalancerPoolOk

`func (o *UpdateLoadBalancerPool200Response) GetLoadBalancerPoolOk() (*UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool, bool)`

GetLoadBalancerPoolOk returns a tuple with the LoadBalancerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPool

`func (o *UpdateLoadBalancerPool200Response) SetLoadBalancerPool(v UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool)`

SetLoadBalancerPool sets LoadBalancerPool field to given value.

### HasLoadBalancerPool

`func (o *UpdateLoadBalancerPool200Response) HasLoadBalancerPool() bool`

HasLoadBalancerPool returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateLoadBalancerPool200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateLoadBalancerPool200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateLoadBalancerPool200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateLoadBalancerPool200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *UpdateLoadBalancerPool200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UpdateLoadBalancerPool200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UpdateLoadBalancerPool200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *UpdateLoadBalancerPool200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *UpdateLoadBalancerPool200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *UpdateLoadBalancerPool200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



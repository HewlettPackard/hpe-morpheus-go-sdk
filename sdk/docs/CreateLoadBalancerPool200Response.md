# CreateLoadBalancerPool200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerPool** | Pointer to [**ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner**](ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateLoadBalancerPool200Response

`func NewCreateLoadBalancerPool200Response() *CreateLoadBalancerPool200Response`

NewCreateLoadBalancerPool200Response instantiates a new CreateLoadBalancerPool200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerPool200ResponseWithDefaults

`func NewCreateLoadBalancerPool200ResponseWithDefaults() *CreateLoadBalancerPool200Response`

NewCreateLoadBalancerPool200ResponseWithDefaults instantiates a new CreateLoadBalancerPool200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerPool

`func (o *CreateLoadBalancerPool200Response) GetLoadBalancerPool() ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner`

GetLoadBalancerPool returns the LoadBalancerPool field if non-nil, zero value otherwise.

### GetLoadBalancerPoolOk

`func (o *CreateLoadBalancerPool200Response) GetLoadBalancerPoolOk() (*ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner, bool)`

GetLoadBalancerPoolOk returns a tuple with the LoadBalancerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPool

`func (o *CreateLoadBalancerPool200Response) SetLoadBalancerPool(v ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner)`

SetLoadBalancerPool sets LoadBalancerPool field to given value.

### HasLoadBalancerPool

`func (o *CreateLoadBalancerPool200Response) HasLoadBalancerPool() bool`

HasLoadBalancerPool returns a boolean if a field has been set.

### GetSuccess

`func (o *CreateLoadBalancerPool200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateLoadBalancerPool200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateLoadBalancerPool200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateLoadBalancerPool200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *CreateLoadBalancerPool200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateLoadBalancerPool200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateLoadBalancerPool200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateLoadBalancerPool200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



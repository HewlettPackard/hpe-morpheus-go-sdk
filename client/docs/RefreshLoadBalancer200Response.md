# RefreshLoadBalancer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancer** | Pointer to [**ListLoadBalancers200ResponseAllOfLoadBalancersInner**](ListLoadBalancers200ResponseAllOfLoadBalancersInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewRefreshLoadBalancer200Response

`func NewRefreshLoadBalancer200Response() *RefreshLoadBalancer200Response`

NewRefreshLoadBalancer200Response instantiates a new RefreshLoadBalancer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshLoadBalancer200ResponseWithDefaults

`func NewRefreshLoadBalancer200ResponseWithDefaults() *RefreshLoadBalancer200Response`

NewRefreshLoadBalancer200ResponseWithDefaults instantiates a new RefreshLoadBalancer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancer

`func (o *RefreshLoadBalancer200Response) GetLoadBalancer() ListLoadBalancers200ResponseAllOfLoadBalancersInner`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *RefreshLoadBalancer200Response) GetLoadBalancerOk() (*ListLoadBalancers200ResponseAllOfLoadBalancersInner, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *RefreshLoadBalancer200Response) SetLoadBalancer(v ListLoadBalancers200ResponseAllOfLoadBalancersInner)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *RefreshLoadBalancer200Response) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetSuccess

`func (o *RefreshLoadBalancer200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RefreshLoadBalancer200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RefreshLoadBalancer200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *RefreshLoadBalancer200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *RefreshLoadBalancer200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *RefreshLoadBalancer200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *RefreshLoadBalancer200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *RefreshLoadBalancer200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



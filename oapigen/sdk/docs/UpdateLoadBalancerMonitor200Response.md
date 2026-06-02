# UpdateLoadBalancerMonitor200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerMonitor** | Pointer to [**UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor**](UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateLoadBalancerMonitor200Response

`func NewUpdateLoadBalancerMonitor200Response() *UpdateLoadBalancerMonitor200Response`

NewUpdateLoadBalancerMonitor200Response instantiates a new UpdateLoadBalancerMonitor200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetLoadBalancerMonitor

`func (o *UpdateLoadBalancerMonitor200Response) GetLoadBalancerMonitor() UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor`

GetLoadBalancerMonitor returns the LoadBalancerMonitor field if non-nil, zero value otherwise.

### GetLoadBalancerMonitorOk

`func (o *UpdateLoadBalancerMonitor200Response) GetLoadBalancerMonitorOk() (*UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor, bool)`

GetLoadBalancerMonitorOk returns a tuple with the LoadBalancerMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerMonitor

`func (o *UpdateLoadBalancerMonitor200Response) SetLoadBalancerMonitor(v UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitor)`

SetLoadBalancerMonitor sets LoadBalancerMonitor field to given value.

### HasLoadBalancerMonitor

`func (o *UpdateLoadBalancerMonitor200Response) HasLoadBalancerMonitor() bool`

HasLoadBalancerMonitor returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateLoadBalancerMonitor200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateLoadBalancerMonitor200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateLoadBalancerMonitor200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateLoadBalancerMonitor200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *UpdateLoadBalancerMonitor200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UpdateLoadBalancerMonitor200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UpdateLoadBalancerMonitor200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *UpdateLoadBalancerMonitor200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *UpdateLoadBalancerMonitor200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *UpdateLoadBalancerMonitor200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



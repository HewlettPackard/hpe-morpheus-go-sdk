# CreateLoadBalancerMonitor200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerMonitor** | Pointer to [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateLoadBalancerMonitor200Response

`func NewCreateLoadBalancerMonitor200Response() *CreateLoadBalancerMonitor200Response`

NewCreateLoadBalancerMonitor200Response instantiates a new CreateLoadBalancerMonitor200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerMonitor200ResponseWithDefaults

`func NewCreateLoadBalancerMonitor200ResponseWithDefaults() *CreateLoadBalancerMonitor200Response`

NewCreateLoadBalancerMonitor200ResponseWithDefaults instantiates a new CreateLoadBalancerMonitor200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerMonitor

`func (o *CreateLoadBalancerMonitor200Response) GetLoadBalancerMonitor() ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner`

GetLoadBalancerMonitor returns the LoadBalancerMonitor field if non-nil, zero value otherwise.

### GetLoadBalancerMonitorOk

`func (o *CreateLoadBalancerMonitor200Response) GetLoadBalancerMonitorOk() (*ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner, bool)`

GetLoadBalancerMonitorOk returns a tuple with the LoadBalancerMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerMonitor

`func (o *CreateLoadBalancerMonitor200Response) SetLoadBalancerMonitor(v ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner)`

SetLoadBalancerMonitor sets LoadBalancerMonitor field to given value.

### HasLoadBalancerMonitor

`func (o *CreateLoadBalancerMonitor200Response) HasLoadBalancerMonitor() bool`

HasLoadBalancerMonitor returns a boolean if a field has been set.

### GetSuccess

`func (o *CreateLoadBalancerMonitor200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateLoadBalancerMonitor200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateLoadBalancerMonitor200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateLoadBalancerMonitor200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *CreateLoadBalancerMonitor200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateLoadBalancerMonitor200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateLoadBalancerMonitor200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateLoadBalancerMonitor200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



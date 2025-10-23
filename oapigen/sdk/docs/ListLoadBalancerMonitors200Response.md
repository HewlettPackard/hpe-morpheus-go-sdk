# ListLoadBalancerMonitors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerMonitors** | Pointer to [**[]ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListLoadBalancerMonitors200Response

`func NewListLoadBalancerMonitors200Response() *ListLoadBalancerMonitors200Response`

NewListLoadBalancerMonitors200Response instantiates a new ListLoadBalancerMonitors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerMonitors200ResponseWithDefaults

`func NewListLoadBalancerMonitors200ResponseWithDefaults() *ListLoadBalancerMonitors200Response`

NewListLoadBalancerMonitors200ResponseWithDefaults instantiates a new ListLoadBalancerMonitors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerMonitors

`func (o *ListLoadBalancerMonitors200Response) GetLoadBalancerMonitors() []ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner`

GetLoadBalancerMonitors returns the LoadBalancerMonitors field if non-nil, zero value otherwise.

### GetLoadBalancerMonitorsOk

`func (o *ListLoadBalancerMonitors200Response) GetLoadBalancerMonitorsOk() (*[]ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner, bool)`

GetLoadBalancerMonitorsOk returns a tuple with the LoadBalancerMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerMonitors

`func (o *ListLoadBalancerMonitors200Response) SetLoadBalancerMonitors(v []ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner)`

SetLoadBalancerMonitors sets LoadBalancerMonitors field to given value.

### HasLoadBalancerMonitors

`func (o *ListLoadBalancerMonitors200Response) HasLoadBalancerMonitors() bool`

HasLoadBalancerMonitors returns a boolean if a field has been set.

### GetMeta

`func (o *ListLoadBalancerMonitors200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLoadBalancerMonitors200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLoadBalancerMonitors200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListLoadBalancerMonitors200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



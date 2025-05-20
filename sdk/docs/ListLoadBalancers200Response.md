# ListLoadBalancers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancers** | Pointer to [**[]ListLoadBalancers200ResponseAllOfLoadBalancersInner**](ListLoadBalancers200ResponseAllOfLoadBalancersInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListLoadBalancers200Response

`func NewListLoadBalancers200Response() *ListLoadBalancers200Response`

NewListLoadBalancers200Response instantiates a new ListLoadBalancers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancers200ResponseWithDefaults

`func NewListLoadBalancers200ResponseWithDefaults() *ListLoadBalancers200Response`

NewListLoadBalancers200ResponseWithDefaults instantiates a new ListLoadBalancers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancers

`func (o *ListLoadBalancers200Response) GetLoadBalancers() []ListLoadBalancers200ResponseAllOfLoadBalancersInner`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *ListLoadBalancers200Response) GetLoadBalancersOk() (*[]ListLoadBalancers200ResponseAllOfLoadBalancersInner, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *ListLoadBalancers200Response) SetLoadBalancers(v []ListLoadBalancers200ResponseAllOfLoadBalancersInner)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *ListLoadBalancers200Response) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetMeta

`func (o *ListLoadBalancers200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLoadBalancers200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLoadBalancers200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListLoadBalancers200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



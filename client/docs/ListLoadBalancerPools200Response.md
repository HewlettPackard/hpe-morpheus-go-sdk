# ListLoadBalancerPools200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerPools** | Pointer to [**[]ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner**](ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListLoadBalancerPools200Response

`func NewListLoadBalancerPools200Response() *ListLoadBalancerPools200Response`

NewListLoadBalancerPools200Response instantiates a new ListLoadBalancerPools200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerPools200ResponseWithDefaults

`func NewListLoadBalancerPools200ResponseWithDefaults() *ListLoadBalancerPools200Response`

NewListLoadBalancerPools200ResponseWithDefaults instantiates a new ListLoadBalancerPools200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerPools

`func (o *ListLoadBalancerPools200Response) GetLoadBalancerPools() []ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner`

GetLoadBalancerPools returns the LoadBalancerPools field if non-nil, zero value otherwise.

### GetLoadBalancerPoolsOk

`func (o *ListLoadBalancerPools200Response) GetLoadBalancerPoolsOk() (*[]ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner, bool)`

GetLoadBalancerPoolsOk returns a tuple with the LoadBalancerPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPools

`func (o *ListLoadBalancerPools200Response) SetLoadBalancerPools(v []ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner)`

SetLoadBalancerPools sets LoadBalancerPools field to given value.

### HasLoadBalancerPools

`func (o *ListLoadBalancerPools200Response) HasLoadBalancerPools() bool`

HasLoadBalancerPools returns a boolean if a field has been set.

### GetMeta

`func (o *ListLoadBalancerPools200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLoadBalancerPools200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLoadBalancerPools200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListLoadBalancerPools200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



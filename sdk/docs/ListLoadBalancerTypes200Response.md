# ListLoadBalancerTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerTypes** | Pointer to [**[]ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner**](ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListLoadBalancerTypes200Response

`func NewListLoadBalancerTypes200Response() *ListLoadBalancerTypes200Response`

NewListLoadBalancerTypes200Response instantiates a new ListLoadBalancerTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerTypes200ResponseWithDefaults

`func NewListLoadBalancerTypes200ResponseWithDefaults() *ListLoadBalancerTypes200Response`

NewListLoadBalancerTypes200ResponseWithDefaults instantiates a new ListLoadBalancerTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerTypes

`func (o *ListLoadBalancerTypes200Response) GetLoadBalancerTypes() []ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner`

GetLoadBalancerTypes returns the LoadBalancerTypes field if non-nil, zero value otherwise.

### GetLoadBalancerTypesOk

`func (o *ListLoadBalancerTypes200Response) GetLoadBalancerTypesOk() (*[]ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner, bool)`

GetLoadBalancerTypesOk returns a tuple with the LoadBalancerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerTypes

`func (o *ListLoadBalancerTypes200Response) SetLoadBalancerTypes(v []ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner)`

SetLoadBalancerTypes sets LoadBalancerTypes field to given value.

### HasLoadBalancerTypes

`func (o *ListLoadBalancerTypes200Response) HasLoadBalancerTypes() bool`

HasLoadBalancerTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListLoadBalancerTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLoadBalancerTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLoadBalancerTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListLoadBalancerTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



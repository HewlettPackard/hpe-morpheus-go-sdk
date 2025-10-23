# ListLoadBalancerProfiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerProfiles** | Pointer to [**[]ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner**](ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListLoadBalancerProfiles200Response

`func NewListLoadBalancerProfiles200Response() *ListLoadBalancerProfiles200Response`

NewListLoadBalancerProfiles200Response instantiates a new ListLoadBalancerProfiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerProfiles200ResponseWithDefaults

`func NewListLoadBalancerProfiles200ResponseWithDefaults() *ListLoadBalancerProfiles200Response`

NewListLoadBalancerProfiles200ResponseWithDefaults instantiates a new ListLoadBalancerProfiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerProfiles

`func (o *ListLoadBalancerProfiles200Response) GetLoadBalancerProfiles() []ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner`

GetLoadBalancerProfiles returns the LoadBalancerProfiles field if non-nil, zero value otherwise.

### GetLoadBalancerProfilesOk

`func (o *ListLoadBalancerProfiles200Response) GetLoadBalancerProfilesOk() (*[]ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner, bool)`

GetLoadBalancerProfilesOk returns a tuple with the LoadBalancerProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerProfiles

`func (o *ListLoadBalancerProfiles200Response) SetLoadBalancerProfiles(v []ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner)`

SetLoadBalancerProfiles sets LoadBalancerProfiles field to given value.

### HasLoadBalancerProfiles

`func (o *ListLoadBalancerProfiles200Response) HasLoadBalancerProfiles() bool`

HasLoadBalancerProfiles returns a boolean if a field has been set.

### GetMeta

`func (o *ListLoadBalancerProfiles200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLoadBalancerProfiles200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLoadBalancerProfiles200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListLoadBalancerProfiles200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



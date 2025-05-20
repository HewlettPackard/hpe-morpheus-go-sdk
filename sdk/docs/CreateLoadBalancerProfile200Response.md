# CreateLoadBalancerProfile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerProfile** | Pointer to [**ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner**](ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateLoadBalancerProfile200Response

`func NewCreateLoadBalancerProfile200Response() *CreateLoadBalancerProfile200Response`

NewCreateLoadBalancerProfile200Response instantiates a new CreateLoadBalancerProfile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerProfile200ResponseWithDefaults

`func NewCreateLoadBalancerProfile200ResponseWithDefaults() *CreateLoadBalancerProfile200Response`

NewCreateLoadBalancerProfile200ResponseWithDefaults instantiates a new CreateLoadBalancerProfile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerProfile

`func (o *CreateLoadBalancerProfile200Response) GetLoadBalancerProfile() ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner`

GetLoadBalancerProfile returns the LoadBalancerProfile field if non-nil, zero value otherwise.

### GetLoadBalancerProfileOk

`func (o *CreateLoadBalancerProfile200Response) GetLoadBalancerProfileOk() (*ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner, bool)`

GetLoadBalancerProfileOk returns a tuple with the LoadBalancerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerProfile

`func (o *CreateLoadBalancerProfile200Response) SetLoadBalancerProfile(v ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner)`

SetLoadBalancerProfile sets LoadBalancerProfile field to given value.

### HasLoadBalancerProfile

`func (o *CreateLoadBalancerProfile200Response) HasLoadBalancerProfile() bool`

HasLoadBalancerProfile returns a boolean if a field has been set.

### GetSuccess

`func (o *CreateLoadBalancerProfile200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateLoadBalancerProfile200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateLoadBalancerProfile200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateLoadBalancerProfile200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *CreateLoadBalancerProfile200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateLoadBalancerProfile200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateLoadBalancerProfile200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateLoadBalancerProfile200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



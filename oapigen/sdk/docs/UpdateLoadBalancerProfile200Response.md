# UpdateLoadBalancerProfile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerProfile** | Pointer to [**UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile**](UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateLoadBalancerProfile200Response

`func NewUpdateLoadBalancerProfile200Response() *UpdateLoadBalancerProfile200Response`

NewUpdateLoadBalancerProfile200Response instantiates a new UpdateLoadBalancerProfile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerProfile200ResponseWithDefaults

`func NewUpdateLoadBalancerProfile200ResponseWithDefaults() *UpdateLoadBalancerProfile200Response`

NewUpdateLoadBalancerProfile200ResponseWithDefaults instantiates a new UpdateLoadBalancerProfile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerProfile

`func (o *UpdateLoadBalancerProfile200Response) GetLoadBalancerProfile() UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile`

GetLoadBalancerProfile returns the LoadBalancerProfile field if non-nil, zero value otherwise.

### GetLoadBalancerProfileOk

`func (o *UpdateLoadBalancerProfile200Response) GetLoadBalancerProfileOk() (*UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile, bool)`

GetLoadBalancerProfileOk returns a tuple with the LoadBalancerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerProfile

`func (o *UpdateLoadBalancerProfile200Response) SetLoadBalancerProfile(v UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile)`

SetLoadBalancerProfile sets LoadBalancerProfile field to given value.

### HasLoadBalancerProfile

`func (o *UpdateLoadBalancerProfile200Response) HasLoadBalancerProfile() bool`

HasLoadBalancerProfile returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateLoadBalancerProfile200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateLoadBalancerProfile200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateLoadBalancerProfile200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateLoadBalancerProfile200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *UpdateLoadBalancerProfile200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UpdateLoadBalancerProfile200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UpdateLoadBalancerProfile200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *UpdateLoadBalancerProfile200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *UpdateLoadBalancerProfile200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *UpdateLoadBalancerProfile200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



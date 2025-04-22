# CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationProfile** | Pointer to **string** | The Load Balancer Application Profile ID The Options API &#x60;/api/options/nsxt/nsxtLBVirtualServerApplicationProfile?loadBalancerId&#x3D;42&amp;loadBalancerInstance.vipProtocol&#x3D;tcp&#x60; can be used to see which options are available.  | [optional] 

## Methods

### NewCreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig

`func NewCreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig() *CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig`

NewCreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig instantiates a new CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfigWithDefaults

`func NewCreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfigWithDefaults() *CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig`

NewCreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfigWithDefaults instantiates a new CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationProfile

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig) GetApplicationProfile() string`

GetApplicationProfile returns the ApplicationProfile field if non-nil, zero value otherwise.

### GetApplicationProfileOk

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig) GetApplicationProfileOk() (*string, bool)`

GetApplicationProfileOk returns a tuple with the ApplicationProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationProfile

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig) SetApplicationProfile(v string)`

SetApplicationProfile sets ApplicationProfile field to given value.

### HasApplicationProfile

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig) HasApplicationProfile() bool`

HasApplicationProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



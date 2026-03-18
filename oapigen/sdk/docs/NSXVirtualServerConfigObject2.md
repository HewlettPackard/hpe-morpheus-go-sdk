# NSXVirtualServerConfigObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationProfile** | Pointer to **NullableString** | The Load Balancer Application Profile ID The Options API &#x60;/api/options/nsxt/nsxtLBVirtualServerApplicationProfile?loadBalancerId&#x3D;42&amp;loadBalancerInstance.vipProtocol&#x3D;tcp&#x60; can be used to see which options are available. | [optional] 

## Methods

### NewNSXVirtualServerConfigObject2

`func NewNSXVirtualServerConfigObject2() *NSXVirtualServerConfigObject2`

NewNSXVirtualServerConfigObject2 instantiates a new NSXVirtualServerConfigObject2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNSXVirtualServerConfigObject2WithDefaults

`func NewNSXVirtualServerConfigObject2WithDefaults() *NSXVirtualServerConfigObject2`

NewNSXVirtualServerConfigObject2WithDefaults instantiates a new NSXVirtualServerConfigObject2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationProfile

`func (o *NSXVirtualServerConfigObject2) GetApplicationProfile() string`

GetApplicationProfile returns the ApplicationProfile field if non-nil, zero value otherwise.

### GetApplicationProfileOk

`func (o *NSXVirtualServerConfigObject2) GetApplicationProfileOk() (*string, bool)`

GetApplicationProfileOk returns a tuple with the ApplicationProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationProfile

`func (o *NSXVirtualServerConfigObject2) SetApplicationProfile(v string)`

SetApplicationProfile sets ApplicationProfile field to given value.

### HasApplicationProfile

`func (o *NSXVirtualServerConfigObject2) HasApplicationProfile() bool`

HasApplicationProfile returns a boolean if a field has been set.

### SetApplicationProfileNil

`func (o *NSXVirtualServerConfigObject2) SetApplicationProfileNil(b bool)`

 SetApplicationProfileNil sets the value for ApplicationProfile to be an explicit nil

### UnsetApplicationProfile
`func (o *NSXVirtualServerConfigObject2) UnsetApplicationProfile()`

UnsetApplicationProfile ensures that no value is present for ApplicationProfile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



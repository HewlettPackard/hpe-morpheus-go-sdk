# CreateLoadBalancerProfileRequestLoadBalancerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**ServiceType** | Pointer to **string** | Service Type | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by type. | [optional] 

## Methods

### NewCreateLoadBalancerProfileRequestLoadBalancerProfile

`func NewCreateLoadBalancerProfileRequestLoadBalancerProfile() *CreateLoadBalancerProfileRequestLoadBalancerProfile`

NewCreateLoadBalancerProfileRequestLoadBalancerProfile instantiates a new CreateLoadBalancerProfileRequestLoadBalancerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerProfileRequestLoadBalancerProfileWithDefaults

`func NewCreateLoadBalancerProfileRequestLoadBalancerProfileWithDefaults() *CreateLoadBalancerProfileRequestLoadBalancerProfile`

NewCreateLoadBalancerProfileRequestLoadBalancerProfileWithDefaults instantiates a new CreateLoadBalancerProfileRequestLoadBalancerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetServiceType

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetConfig

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateLoadBalancerProfileRequestLoadBalancerProfile) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateLoadBalancerProfileRequestLoadBalancerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**ServiceType** | Pointer to **string** | Service Type | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by type. | [optional] 

## Methods

### NewUpdateLoadBalancerProfileRequestLoadBalancerProfile

`func NewUpdateLoadBalancerProfileRequestLoadBalancerProfile() *UpdateLoadBalancerProfileRequestLoadBalancerProfile`

NewUpdateLoadBalancerProfileRequestLoadBalancerProfile instantiates a new UpdateLoadBalancerProfileRequestLoadBalancerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerProfileRequestLoadBalancerProfileWithDefaults

`func NewUpdateLoadBalancerProfileRequestLoadBalancerProfileWithDefaults() *UpdateLoadBalancerProfileRequestLoadBalancerProfile`

NewUpdateLoadBalancerProfileRequestLoadBalancerProfileWithDefaults instantiates a new UpdateLoadBalancerProfileRequestLoadBalancerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetServiceType

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateLoadBalancerProfileRequestLoadBalancerProfile) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



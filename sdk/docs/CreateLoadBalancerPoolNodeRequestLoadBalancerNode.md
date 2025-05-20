# CreateLoadBalancerPoolNodeRequestLoadBalancerNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**IpAddress** | Pointer to **string** | IP Address | [optional] 
**Port** | Pointer to **int32** | Port number | [optional] 
**Weight** | Pointer to **int32** | Weight applied balance algoritm | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by type. | [optional] 

## Methods

### NewCreateLoadBalancerPoolNodeRequestLoadBalancerNode

`func NewCreateLoadBalancerPoolNodeRequestLoadBalancerNode() *CreateLoadBalancerPoolNodeRequestLoadBalancerNode`

NewCreateLoadBalancerPoolNodeRequestLoadBalancerNode instantiates a new CreateLoadBalancerPoolNodeRequestLoadBalancerNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerPoolNodeRequestLoadBalancerNodeWithDefaults

`func NewCreateLoadBalancerPoolNodeRequestLoadBalancerNodeWithDefaults() *CreateLoadBalancerPoolNodeRequestLoadBalancerNode`

NewCreateLoadBalancerPoolNodeRequestLoadBalancerNodeWithDefaults instantiates a new CreateLoadBalancerPoolNodeRequestLoadBalancerNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddress

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPort

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetWeight

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetConfig

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateLoadBalancerPoolNodeRequestLoadBalancerNode) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateNetworkGroupRequestNetworkGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to **[]int64** |  | [optional] 
**Subnets** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateNetworkGroupRequestNetworkGroup

`func NewUpdateNetworkGroupRequestNetworkGroup() *UpdateNetworkGroupRequestNetworkGroup`

NewUpdateNetworkGroupRequestNetworkGroup instantiates a new UpdateNetworkGroupRequestNetworkGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkGroupRequestNetworkGroupWithDefaults

`func NewUpdateNetworkGroupRequestNetworkGroupWithDefaults() *UpdateNetworkGroupRequestNetworkGroup`

NewUpdateNetworkGroupRequestNetworkGroupWithDefaults instantiates a new UpdateNetworkGroupRequestNetworkGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkGroupRequestNetworkGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkGroupRequestNetworkGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkGroupRequestNetworkGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkGroupRequestNetworkGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateNetworkGroupRequestNetworkGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkGroupRequestNetworkGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkGroupRequestNetworkGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkGroupRequestNetworkGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNetworks

`func (o *UpdateNetworkGroupRequestNetworkGroup) GetNetworks() []int64`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *UpdateNetworkGroupRequestNetworkGroup) GetNetworksOk() (*[]int64, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *UpdateNetworkGroupRequestNetworkGroup) SetNetworks(v []int64)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *UpdateNetworkGroupRequestNetworkGroup) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSubnets

`func (o *UpdateNetworkGroupRequestNetworkGroup) GetSubnets() []map[string]interface{}`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *UpdateNetworkGroupRequestNetworkGroup) GetSubnetsOk() (*[]map[string]interface{}, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *UpdateNetworkGroupRequestNetworkGroup) SetSubnets(v []map[string]interface{})`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *UpdateNetworkGroupRequestNetworkGroup) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



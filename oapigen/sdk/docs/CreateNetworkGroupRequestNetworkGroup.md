# CreateNetworkGroupRequestNetworkGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to **[]int64** |  | [optional] 
**Subnets** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewCreateNetworkGroupRequestNetworkGroup

`func NewCreateNetworkGroupRequestNetworkGroup() *CreateNetworkGroupRequestNetworkGroup`

NewCreateNetworkGroupRequestNetworkGroup instantiates a new CreateNetworkGroupRequestNetworkGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkGroupRequestNetworkGroupWithDefaults

`func NewCreateNetworkGroupRequestNetworkGroupWithDefaults() *CreateNetworkGroupRequestNetworkGroup`

NewCreateNetworkGroupRequestNetworkGroupWithDefaults instantiates a new CreateNetworkGroupRequestNetworkGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkGroupRequestNetworkGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkGroupRequestNetworkGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkGroupRequestNetworkGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkGroupRequestNetworkGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateNetworkGroupRequestNetworkGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkGroupRequestNetworkGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkGroupRequestNetworkGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkGroupRequestNetworkGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNetworks

`func (o *CreateNetworkGroupRequestNetworkGroup) GetNetworks() []int64`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *CreateNetworkGroupRequestNetworkGroup) GetNetworksOk() (*[]int64, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *CreateNetworkGroupRequestNetworkGroup) SetNetworks(v []int64)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *CreateNetworkGroupRequestNetworkGroup) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSubnets

`func (o *CreateNetworkGroupRequestNetworkGroup) GetSubnets() []map[string]interface{}`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CreateNetworkGroupRequestNetworkGroup) GetSubnetsOk() (*[]map[string]interface{}, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CreateNetworkGroupRequestNetworkGroup) SetSubnets(v []map[string]interface{})`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *CreateNetworkGroupRequestNetworkGroup) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



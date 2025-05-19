# NetworkGroupsCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to **[]int64** |  | [optional] 
**Subnets** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewNetworkGroupsCreate

`func NewNetworkGroupsCreate() *NetworkGroupsCreate`

NewNetworkGroupsCreate instantiates a new NetworkGroupsCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkGroupsCreateWithDefaults

`func NewNetworkGroupsCreateWithDefaults() *NetworkGroupsCreate`

NewNetworkGroupsCreateWithDefaults instantiates a new NetworkGroupsCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkGroupsCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkGroupsCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkGroupsCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkGroupsCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkGroupsCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkGroupsCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkGroupsCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkGroupsCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNetworks

`func (o *NetworkGroupsCreate) GetNetworks() []int64`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *NetworkGroupsCreate) GetNetworksOk() (*[]int64, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *NetworkGroupsCreate) SetNetworks(v []int64)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *NetworkGroupsCreate) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSubnets

`func (o *NetworkGroupsCreate) GetSubnets() []map[string]interface{}`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *NetworkGroupsCreate) GetSubnetsOk() (*[]map[string]interface{}, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *NetworkGroupsCreate) SetSubnets(v []map[string]interface{})`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *NetworkGroupsCreate) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



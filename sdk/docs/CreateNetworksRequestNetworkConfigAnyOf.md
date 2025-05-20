# CreateNetworksRequestNetworkConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroupId** | **string** | Resource Group Name | 
**SubnetName** | **string** | Subnet Name | 
**SubnetCidr** | **string** | The subnet&#39;s address range in CIDR notation (e.g. 192.168.1.0/24). It must be contained by the address space of the virtual network. | 

## Methods

### NewCreateNetworksRequestNetworkConfigAnyOf

`func NewCreateNetworksRequestNetworkConfigAnyOf(resourceGroupId string, subnetName string, subnetCidr string, ) *CreateNetworksRequestNetworkConfigAnyOf`

NewCreateNetworksRequestNetworkConfigAnyOf instantiates a new CreateNetworksRequestNetworkConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworksRequestNetworkConfigAnyOfWithDefaults

`func NewCreateNetworksRequestNetworkConfigAnyOfWithDefaults() *CreateNetworksRequestNetworkConfigAnyOf`

NewCreateNetworksRequestNetworkConfigAnyOfWithDefaults instantiates a new CreateNetworksRequestNetworkConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroupId

`func (o *CreateNetworksRequestNetworkConfigAnyOf) GetResourceGroupId() string`

GetResourceGroupId returns the ResourceGroupId field if non-nil, zero value otherwise.

### GetResourceGroupIdOk

`func (o *CreateNetworksRequestNetworkConfigAnyOf) GetResourceGroupIdOk() (*string, bool)`

GetResourceGroupIdOk returns a tuple with the ResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupId

`func (o *CreateNetworksRequestNetworkConfigAnyOf) SetResourceGroupId(v string)`

SetResourceGroupId sets ResourceGroupId field to given value.


### GetSubnetName

`func (o *CreateNetworksRequestNetworkConfigAnyOf) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *CreateNetworksRequestNetworkConfigAnyOf) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *CreateNetworksRequestNetworkConfigAnyOf) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.


### GetSubnetCidr

`func (o *CreateNetworksRequestNetworkConfigAnyOf) GetSubnetCidr() string`

GetSubnetCidr returns the SubnetCidr field if non-nil, zero value otherwise.

### GetSubnetCidrOk

`func (o *CreateNetworksRequestNetworkConfigAnyOf) GetSubnetCidrOk() (*string, bool)`

GetSubnetCidrOk returns a tuple with the SubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidr

`func (o *CreateNetworksRequestNetworkConfigAnyOf) SetSubnetCidr(v string)`

SetSubnetCidr sets SubnetCidr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



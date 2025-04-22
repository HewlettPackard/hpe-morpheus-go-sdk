# CreateNetworksRequestNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroupId** | **string** | Resource Group Name | 
**SubnetName** | **string** | Subnet Name | 
**SubnetCidr** | **string** | The subnet&#39;s address range in CIDR notation (e.g. 192.168.1.0/24). It must be contained by the address space of the virtual network. | 
**AvailabilityZone** | **string** | Availability Zone Name | 
**Cidr** | **string** | Network CIDR | 
**AssignPublicIp** | **bool** | Assign public IPs by default. | 
**ZonePool** | [**CreateNetworksRequestNetworkConfigAnyOf2ZonePool**](CreateNetworksRequestNetworkConfigAnyOf2ZonePool.md) |  | 
**Mtu** | **string** | GCP MTU | [default to "1460"]
**AutoCreate** | **bool** | Auto create subnets | [default to true]

## Methods

### NewCreateNetworksRequestNetworkConfig

`func NewCreateNetworksRequestNetworkConfig(resourceGroupId string, subnetName string, subnetCidr string, availabilityZone string, cidr string, assignPublicIp bool, zonePool CreateNetworksRequestNetworkConfigAnyOf2ZonePool, mtu string, autoCreate bool, ) *CreateNetworksRequestNetworkConfig`

NewCreateNetworksRequestNetworkConfig instantiates a new CreateNetworksRequestNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworksRequestNetworkConfigWithDefaults

`func NewCreateNetworksRequestNetworkConfigWithDefaults() *CreateNetworksRequestNetworkConfig`

NewCreateNetworksRequestNetworkConfigWithDefaults instantiates a new CreateNetworksRequestNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroupId

`func (o *CreateNetworksRequestNetworkConfig) GetResourceGroupId() string`

GetResourceGroupId returns the ResourceGroupId field if non-nil, zero value otherwise.

### GetResourceGroupIdOk

`func (o *CreateNetworksRequestNetworkConfig) GetResourceGroupIdOk() (*string, bool)`

GetResourceGroupIdOk returns a tuple with the ResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupId

`func (o *CreateNetworksRequestNetworkConfig) SetResourceGroupId(v string)`

SetResourceGroupId sets ResourceGroupId field to given value.


### GetSubnetName

`func (o *CreateNetworksRequestNetworkConfig) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *CreateNetworksRequestNetworkConfig) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *CreateNetworksRequestNetworkConfig) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.


### GetSubnetCidr

`func (o *CreateNetworksRequestNetworkConfig) GetSubnetCidr() string`

GetSubnetCidr returns the SubnetCidr field if non-nil, zero value otherwise.

### GetSubnetCidrOk

`func (o *CreateNetworksRequestNetworkConfig) GetSubnetCidrOk() (*string, bool)`

GetSubnetCidrOk returns a tuple with the SubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidr

`func (o *CreateNetworksRequestNetworkConfig) SetSubnetCidr(v string)`

SetSubnetCidr sets SubnetCidr field to given value.


### GetAvailabilityZone

`func (o *CreateNetworksRequestNetworkConfig) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateNetworksRequestNetworkConfig) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateNetworksRequestNetworkConfig) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetCidr

`func (o *CreateNetworksRequestNetworkConfig) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateNetworksRequestNetworkConfig) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateNetworksRequestNetworkConfig) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetAssignPublicIp

`func (o *CreateNetworksRequestNetworkConfig) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *CreateNetworksRequestNetworkConfig) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *CreateNetworksRequestNetworkConfig) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.


### GetZonePool

`func (o *CreateNetworksRequestNetworkConfig) GetZonePool() CreateNetworksRequestNetworkConfigAnyOf2ZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *CreateNetworksRequestNetworkConfig) GetZonePoolOk() (*CreateNetworksRequestNetworkConfigAnyOf2ZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *CreateNetworksRequestNetworkConfig) SetZonePool(v CreateNetworksRequestNetworkConfigAnyOf2ZonePool)`

SetZonePool sets ZonePool field to given value.


### GetMtu

`func (o *CreateNetworksRequestNetworkConfig) GetMtu() string`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateNetworksRequestNetworkConfig) GetMtuOk() (*string, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateNetworksRequestNetworkConfig) SetMtu(v string)`

SetMtu sets Mtu field to given value.


### GetAutoCreate

`func (o *CreateNetworksRequestNetworkConfig) GetAutoCreate() bool`

GetAutoCreate returns the AutoCreate field if non-nil, zero value otherwise.

### GetAutoCreateOk

`func (o *CreateNetworksRequestNetworkConfig) GetAutoCreateOk() (*bool, bool)`

GetAutoCreateOk returns a tuple with the AutoCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreate

`func (o *CreateNetworksRequestNetworkConfig) SetAutoCreate(v bool)`

SetAutoCreate sets AutoCreate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



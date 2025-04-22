# CreateNetworksRequestNetworkConfigAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | **string** | Availability Zone Name | 
**Cidr** | **string** | Network CIDR | 
**AssignPublicIp** | **bool** | Assign public IPs by default. | 
**ZonePool** | [**CreateNetworksRequestNetworkConfigAnyOf1ZonePool**](CreateNetworksRequestNetworkConfigAnyOf1ZonePool.md) |  | 

## Methods

### NewCreateNetworksRequestNetworkConfigAnyOf1

`func NewCreateNetworksRequestNetworkConfigAnyOf1(availabilityZone string, cidr string, assignPublicIp bool, zonePool CreateNetworksRequestNetworkConfigAnyOf1ZonePool, ) *CreateNetworksRequestNetworkConfigAnyOf1`

NewCreateNetworksRequestNetworkConfigAnyOf1 instantiates a new CreateNetworksRequestNetworkConfigAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworksRequestNetworkConfigAnyOf1WithDefaults

`func NewCreateNetworksRequestNetworkConfigAnyOf1WithDefaults() *CreateNetworksRequestNetworkConfigAnyOf1`

NewCreateNetworksRequestNetworkConfigAnyOf1WithDefaults instantiates a new CreateNetworksRequestNetworkConfigAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *CreateNetworksRequestNetworkConfigAnyOf1) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateNetworksRequestNetworkConfigAnyOf1) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateNetworksRequestNetworkConfigAnyOf1) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetCidr

`func (o *CreateNetworksRequestNetworkConfigAnyOf1) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateNetworksRequestNetworkConfigAnyOf1) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateNetworksRequestNetworkConfigAnyOf1) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetAssignPublicIp

`func (o *CreateNetworksRequestNetworkConfigAnyOf1) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *CreateNetworksRequestNetworkConfigAnyOf1) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *CreateNetworksRequestNetworkConfigAnyOf1) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.


### GetZonePool

`func (o *CreateNetworksRequestNetworkConfigAnyOf1) GetZonePool() CreateNetworksRequestNetworkConfigAnyOf1ZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *CreateNetworksRequestNetworkConfigAnyOf1) GetZonePoolOk() (*CreateNetworksRequestNetworkConfigAnyOf1ZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *CreateNetworksRequestNetworkConfigAnyOf1) SetZonePool(v CreateNetworksRequestNetworkConfigAnyOf1ZonePool)`

SetZonePool sets ZonePool field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AddClusterRequestClusterServerNetworkInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**AddClusterRequestClusterServerNetworkInterfacesInnerNetwork**](AddClusterRequestClusterServerNetworkInterfacesInnerNetwork.md) |  | 
**NetworkInterfaceTypeId** | Pointer to **int64** | The id of type of the network interface. | [optional] 
**IpAddress** | Pointer to **string** | The ip address. Not applicable when using DHCP or IP Pools. | [optional] 

## Methods

### NewAddClusterRequestClusterServerNetworkInterfacesInner

`func NewAddClusterRequestClusterServerNetworkInterfacesInner(network AddClusterRequestClusterServerNetworkInterfacesInnerNetwork, ) *AddClusterRequestClusterServerNetworkInterfacesInner`

NewAddClusterRequestClusterServerNetworkInterfacesInner instantiates a new AddClusterRequestClusterServerNetworkInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterRequestClusterServerNetworkInterfacesInnerWithDefaults

`func NewAddClusterRequestClusterServerNetworkInterfacesInnerWithDefaults() *AddClusterRequestClusterServerNetworkInterfacesInner`

NewAddClusterRequestClusterServerNetworkInterfacesInnerWithDefaults instantiates a new AddClusterRequestClusterServerNetworkInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *AddClusterRequestClusterServerNetworkInterfacesInner) GetNetwork() AddClusterRequestClusterServerNetworkInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddClusterRequestClusterServerNetworkInterfacesInner) GetNetworkOk() (*AddClusterRequestClusterServerNetworkInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddClusterRequestClusterServerNetworkInterfacesInner) SetNetwork(v AddClusterRequestClusterServerNetworkInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.


### GetNetworkInterfaceTypeId

`func (o *AddClusterRequestClusterServerNetworkInterfacesInner) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *AddClusterRequestClusterServerNetworkInterfacesInner) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *AddClusterRequestClusterServerNetworkInterfacesInner) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *AddClusterRequestClusterServerNetworkInterfacesInner) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### GetIpAddress

`func (o *AddClusterRequestClusterServerNetworkInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AddClusterRequestClusterServerNetworkInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AddClusterRequestClusterServerNetworkInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *AddClusterRequestClusterServerNetworkInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



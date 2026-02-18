# ClusterCreateServerNetworkInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**ClusterCreateServerNetworkInterfacesInnerNetwork**](ClusterCreateServerNetworkInterfacesInnerNetwork.md) |  | 
**NetworkInterfaceTypeId** | Pointer to **int64** | The id of type of the network interface. | [optional] 
**IpAddress** | Pointer to **string** | The ip address. Not applicable when using DHCP or IP Pools. | [optional] 

## Methods

### NewClusterCreateServerNetworkInterfacesInner

`func NewClusterCreateServerNetworkInterfacesInner(network ClusterCreateServerNetworkInterfacesInnerNetwork, ) *ClusterCreateServerNetworkInterfacesInner`

NewClusterCreateServerNetworkInterfacesInner instantiates a new ClusterCreateServerNetworkInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateServerNetworkInterfacesInnerWithDefaults

`func NewClusterCreateServerNetworkInterfacesInnerWithDefaults() *ClusterCreateServerNetworkInterfacesInner`

NewClusterCreateServerNetworkInterfacesInnerWithDefaults instantiates a new ClusterCreateServerNetworkInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *ClusterCreateServerNetworkInterfacesInner) GetNetwork() ClusterCreateServerNetworkInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClusterCreateServerNetworkInterfacesInner) GetNetworkOk() (*ClusterCreateServerNetworkInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClusterCreateServerNetworkInterfacesInner) SetNetwork(v ClusterCreateServerNetworkInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.


### GetNetworkInterfaceTypeId

`func (o *ClusterCreateServerNetworkInterfacesInner) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *ClusterCreateServerNetworkInterfacesInner) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *ClusterCreateServerNetworkInterfacesInner) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *ClusterCreateServerNetworkInterfacesInner) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### GetIpAddress

`func (o *ClusterCreateServerNetworkInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ClusterCreateServerNetworkInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ClusterCreateServerNetworkInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ClusterCreateServerNetworkInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InstancesChildVirtualNetworkInterfacesInner5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InstancesChildVirtualNetworkInterfacesInner5Network**](InstancesChildVirtualNetworkInterfacesInner5Network.md) |  | 
**NetworkInterfaceTypeId** | Pointer to **int64** | The id of type of the network interface. | [optional] 
**IpMode** | Pointer to **string** | The mode for determining ip address. Can be &#39;static&#39;, &#39;dhcp&#39; or empty string. | [optional] [default to ""]
**IpAddress** | Pointer to **string** | The ip address. Not applicable when using DHCP or IP Pools. | [optional] 
**MacAddress** | Pointer to **string** | The MAC address. | [optional] 
**Id** | Pointer to **int64** | The interface id. Applicable when resizing and you want to identify an interface to update that already exists. | [optional] 

## Methods

### NewInstancesChildVirtualNetworkInterfacesInner5

`func NewInstancesChildVirtualNetworkInterfacesInner5(network InstancesChildVirtualNetworkInterfacesInner5Network, ) *InstancesChildVirtualNetworkInterfacesInner5`

NewInstancesChildVirtualNetworkInterfacesInner5 instantiates a new InstancesChildVirtualNetworkInterfacesInner5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesChildVirtualNetworkInterfacesInner5WithDefaults

`func NewInstancesChildVirtualNetworkInterfacesInner5WithDefaults() *InstancesChildVirtualNetworkInterfacesInner5`

NewInstancesChildVirtualNetworkInterfacesInner5WithDefaults instantiates a new InstancesChildVirtualNetworkInterfacesInner5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InstancesChildVirtualNetworkInterfacesInner5) GetNetwork() InstancesChildVirtualNetworkInterfacesInner5Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstancesChildVirtualNetworkInterfacesInner5) GetNetworkOk() (*InstancesChildVirtualNetworkInterfacesInner5Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstancesChildVirtualNetworkInterfacesInner5) SetNetwork(v InstancesChildVirtualNetworkInterfacesInner5Network)`

SetNetwork sets Network field to given value.


### GetNetworkInterfaceTypeId

`func (o *InstancesChildVirtualNetworkInterfacesInner5) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *InstancesChildVirtualNetworkInterfacesInner5) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *InstancesChildVirtualNetworkInterfacesInner5) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *InstancesChildVirtualNetworkInterfacesInner5) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### GetIpMode

`func (o *InstancesChildVirtualNetworkInterfacesInner5) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *InstancesChildVirtualNetworkInterfacesInner5) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *InstancesChildVirtualNetworkInterfacesInner5) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *InstancesChildVirtualNetworkInterfacesInner5) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetIpAddress

`func (o *InstancesChildVirtualNetworkInterfacesInner5) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InstancesChildVirtualNetworkInterfacesInner5) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InstancesChildVirtualNetworkInterfacesInner5) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InstancesChildVirtualNetworkInterfacesInner5) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *InstancesChildVirtualNetworkInterfacesInner5) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *InstancesChildVirtualNetworkInterfacesInner5) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *InstancesChildVirtualNetworkInterfacesInner5) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *InstancesChildVirtualNetworkInterfacesInner5) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetId

`func (o *InstancesChildVirtualNetworkInterfacesInner5) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesChildVirtualNetworkInterfacesInner5) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesChildVirtualNetworkInterfacesInner5) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstancesChildVirtualNetworkInterfacesInner5) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



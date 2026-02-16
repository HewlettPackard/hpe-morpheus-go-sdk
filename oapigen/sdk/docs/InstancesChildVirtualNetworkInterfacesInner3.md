# InstancesChildVirtualNetworkInterfacesInner3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InstancesChildVirtualNetworkInterfacesInner3Network**](InstancesChildVirtualNetworkInterfacesInner3Network.md) |  | 
**NetworkInterfaceTypeId** | Pointer to **int64** | The id of type of the network interface. | [optional] 
**IpMode** | Pointer to **string** | The mode for determining ip address. Can be &#39;static&#39;, &#39;dhcp&#39; or empty string. | [optional] [default to ""]
**IpAddress** | Pointer to **string** | The ip address. Not applicable when using DHCP or IP Pools. | [optional] 
**MacAddress** | Pointer to **string** | The MAC address. | [optional] 
**Id** | Pointer to **int64** | The interface id. Applicable when resizing and you want to identify an interface to update that already exists. | [optional] 

## Methods

### NewInstancesChildVirtualNetworkInterfacesInner3

`func NewInstancesChildVirtualNetworkInterfacesInner3(network InstancesChildVirtualNetworkInterfacesInner3Network, ) *InstancesChildVirtualNetworkInterfacesInner3`

NewInstancesChildVirtualNetworkInterfacesInner3 instantiates a new InstancesChildVirtualNetworkInterfacesInner3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesChildVirtualNetworkInterfacesInner3WithDefaults

`func NewInstancesChildVirtualNetworkInterfacesInner3WithDefaults() *InstancesChildVirtualNetworkInterfacesInner3`

NewInstancesChildVirtualNetworkInterfacesInner3WithDefaults instantiates a new InstancesChildVirtualNetworkInterfacesInner3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InstancesChildVirtualNetworkInterfacesInner3) GetNetwork() InstancesChildVirtualNetworkInterfacesInner3Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstancesChildVirtualNetworkInterfacesInner3) GetNetworkOk() (*InstancesChildVirtualNetworkInterfacesInner3Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstancesChildVirtualNetworkInterfacesInner3) SetNetwork(v InstancesChildVirtualNetworkInterfacesInner3Network)`

SetNetwork sets Network field to given value.


### GetNetworkInterfaceTypeId

`func (o *InstancesChildVirtualNetworkInterfacesInner3) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *InstancesChildVirtualNetworkInterfacesInner3) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *InstancesChildVirtualNetworkInterfacesInner3) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *InstancesChildVirtualNetworkInterfacesInner3) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### GetIpMode

`func (o *InstancesChildVirtualNetworkInterfacesInner3) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *InstancesChildVirtualNetworkInterfacesInner3) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *InstancesChildVirtualNetworkInterfacesInner3) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *InstancesChildVirtualNetworkInterfacesInner3) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetIpAddress

`func (o *InstancesChildVirtualNetworkInterfacesInner3) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InstancesChildVirtualNetworkInterfacesInner3) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InstancesChildVirtualNetworkInterfacesInner3) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InstancesChildVirtualNetworkInterfacesInner3) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *InstancesChildVirtualNetworkInterfacesInner3) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *InstancesChildVirtualNetworkInterfacesInner3) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *InstancesChildVirtualNetworkInterfacesInner3) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *InstancesChildVirtualNetworkInterfacesInner3) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetId

`func (o *InstancesChildVirtualNetworkInterfacesInner3) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesChildVirtualNetworkInterfacesInner3) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesChildVirtualNetworkInterfacesInner3) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstancesChildVirtualNetworkInterfacesInner3) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InstancesNetworkInterfaces6NetworkInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork**](InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork.md) |  | 
**NetworkInterfaceTypeId** | Pointer to **int64** | The id of type of the network interface. | [optional] 
**IpMode** | Pointer to **string** | The mode for determining ip address. Can be &#39;static&#39;, &#39;dhcp&#39; or empty string. | [optional] [default to ""]
**IpAddress** | Pointer to **string** | The ip address. Not applicable when using DHCP or IP Pools. | [optional] 
**MacAddress** | Pointer to **string** | The MAC address. | [optional] 
**Id** | Pointer to **int64** | The interface id. Applicable when resizing and you want to identify an interface to update that already exists. | [optional] 

## Methods

### NewInstancesNetworkInterfaces6NetworkInterfacesInner

`func NewInstancesNetworkInterfaces6NetworkInterfacesInner(network InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork, ) *InstancesNetworkInterfaces6NetworkInterfacesInner`

NewInstancesNetworkInterfaces6NetworkInterfacesInner instantiates a new InstancesNetworkInterfaces6NetworkInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesNetworkInterfaces6NetworkInterfacesInnerWithDefaults

`func NewInstancesNetworkInterfaces6NetworkInterfacesInnerWithDefaults() *InstancesNetworkInterfaces6NetworkInterfacesInner`

NewInstancesNetworkInterfaces6NetworkInterfacesInnerWithDefaults instantiates a new InstancesNetworkInterfaces6NetworkInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) GetNetwork() InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) GetNetworkOk() (*InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) SetNetwork(v InstancesNetworkInterfaces6NetworkInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.


### GetNetworkInterfaceTypeId

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### GetIpMode

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetIpAddress

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetId

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstancesNetworkInterfaces6NetworkInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



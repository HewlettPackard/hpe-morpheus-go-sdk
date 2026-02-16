# InstancesNetworkInterfaces1NetworkInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork**](InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork.md) |  | 
**NetworkInterfaceTypeId** | Pointer to **int64** | The id of type of the network interface. | [optional] 
**IpMode** | Pointer to **string** | The mode for determining ip address. Can be &#39;static&#39;, &#39;dhcp&#39; or empty string. | [optional] [default to ""]
**IpAddress** | Pointer to **string** | The ip address. Not applicable when using DHCP or IP Pools. | [optional] 
**MacAddress** | Pointer to **string** | The MAC address. | [optional] 
**Id** | Pointer to **int64** | The interface id. Applicable when resizing and you want to identify an interface to update that already exists. | [optional] 

## Methods

### NewInstancesNetworkInterfaces1NetworkInterfacesInner

`func NewInstancesNetworkInterfaces1NetworkInterfacesInner(network InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork, ) *InstancesNetworkInterfaces1NetworkInterfacesInner`

NewInstancesNetworkInterfaces1NetworkInterfacesInner instantiates a new InstancesNetworkInterfaces1NetworkInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesNetworkInterfaces1NetworkInterfacesInnerWithDefaults

`func NewInstancesNetworkInterfaces1NetworkInterfacesInnerWithDefaults() *InstancesNetworkInterfaces1NetworkInterfacesInner`

NewInstancesNetworkInterfaces1NetworkInterfacesInnerWithDefaults instantiates a new InstancesNetworkInterfaces1NetworkInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) GetNetwork() InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) GetNetworkOk() (*InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) SetNetwork(v InstancesNetworkInterfaces1NetworkInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.


### GetNetworkInterfaceTypeId

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### GetIpMode

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetIpAddress

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetId

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstancesNetworkInterfaces1NetworkInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



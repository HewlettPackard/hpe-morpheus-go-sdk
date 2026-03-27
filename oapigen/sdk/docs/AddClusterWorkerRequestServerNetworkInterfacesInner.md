# AddClusterWorkerRequestServerNetworkInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**AddClusterWorkerRequestServerNetworkInterfacesInnerNetwork**](AddClusterWorkerRequestServerNetworkInterfacesInnerNetwork.md) |  | 
**NetworkInterfaceTypeId** | Pointer to **int64** | The id of type of the network interface. | [optional] 
**IpAddress** | Pointer to **string** | The ip address. Not applicable when using DHCP or IP Pools. | [optional] 
**IpMode** | Pointer to **string** | The mode for determining ip address. Can be &#39;static&#39;, &#39;dhcp&#39; or empty string. | [optional] [default to ""]

## Methods

### NewAddClusterWorkerRequestServerNetworkInterfacesInner

`func NewAddClusterWorkerRequestServerNetworkInterfacesInner(network AddClusterWorkerRequestServerNetworkInterfacesInnerNetwork, ) *AddClusterWorkerRequestServerNetworkInterfacesInner`

NewAddClusterWorkerRequestServerNetworkInterfacesInner instantiates a new AddClusterWorkerRequestServerNetworkInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterWorkerRequestServerNetworkInterfacesInnerWithDefaults

`func NewAddClusterWorkerRequestServerNetworkInterfacesInnerWithDefaults() *AddClusterWorkerRequestServerNetworkInterfacesInner`

NewAddClusterWorkerRequestServerNetworkInterfacesInnerWithDefaults instantiates a new AddClusterWorkerRequestServerNetworkInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) GetNetwork() AddClusterWorkerRequestServerNetworkInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) GetNetworkOk() (*AddClusterWorkerRequestServerNetworkInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) SetNetwork(v AddClusterWorkerRequestServerNetworkInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.


### GetNetworkInterfaceTypeId

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### GetIpAddress

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpMode

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *AddClusterWorkerRequestServerNetworkInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



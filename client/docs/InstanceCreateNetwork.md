# InstanceCreateNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetwork**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetwork.md) |  | 
**NetworkInterfaceTypeId** | Pointer to **int64** | The id of type of the network interface. | [optional] 
**IpMode** | Pointer to **string** | The mode for determining ip address. Use &#39;static&#39; when specifying an ipAddress, otherwise &#39;dhcp&#39; is used. | [optional] [default to "dhcp"]
**IpAddress** | Pointer to **string** | The ip address. Not applicable when using DHCP or IP Pools. | [optional] 
**Id** | Pointer to **int64** | The interface id. Applicable when resizing and you want to identify an interface to update that already exists. | [optional] 

## Methods

### NewInstanceCreateNetwork

`func NewInstanceCreateNetwork(network AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetwork, ) *InstanceCreateNetwork`

NewInstanceCreateNetwork instantiates a new InstanceCreateNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceCreateNetworkWithDefaults

`func NewInstanceCreateNetworkWithDefaults() *InstanceCreateNetwork`

NewInstanceCreateNetworkWithDefaults instantiates a new InstanceCreateNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InstanceCreateNetwork) GetNetwork() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstanceCreateNetwork) GetNetworkOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstanceCreateNetwork) SetNetwork(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.


### GetNetworkInterfaceTypeId

`func (o *InstanceCreateNetwork) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *InstanceCreateNetwork) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *InstanceCreateNetwork) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *InstanceCreateNetwork) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### GetIpMode

`func (o *InstanceCreateNetwork) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *InstanceCreateNetwork) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *InstanceCreateNetwork) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *InstanceCreateNetwork) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetIpAddress

`func (o *InstanceCreateNetwork) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InstanceCreateNetwork) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InstanceCreateNetwork) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InstanceCreateNetwork) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetId

`func (o *InstanceCreateNetwork) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceCreateNetwork) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceCreateNetwork) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceCreateNetwork) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



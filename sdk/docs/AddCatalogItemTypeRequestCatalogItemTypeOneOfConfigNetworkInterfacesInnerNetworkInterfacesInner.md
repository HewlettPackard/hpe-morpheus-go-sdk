# AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetwork**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetwork.md) |  | 
**NetworkInterfaceTypeId** | Pointer to **int64** | The id of type of the network interface. | [optional] 
**IpMode** | Pointer to **string** | The mode for determining ip address. Use &#39;static&#39; when specifying an ipAddress, otherwise &#39;dhcp&#39; is used. | [optional] [default to "dhcp"]
**IpAddress** | Pointer to **string** | The ip address. Not applicable when using DHCP or IP Pools. | [optional] 
**MacAddress** | Pointer to **string** | The MAC address. | [optional] 
**Id** | Pointer to **int64** | The interface id. Applicable when resizing and you want to identify an interface to update that already exists. | [optional] 

## Methods

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner(network AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetwork, ) *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInnerWithDefaults

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInnerWithDefaults() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInnerWithDefaults instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) GetNetwork() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) GetNetworkOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) SetNetwork(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.


### GetNetworkInterfaceTypeId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### GetIpMode

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetIpAddress

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInnerNetworkInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



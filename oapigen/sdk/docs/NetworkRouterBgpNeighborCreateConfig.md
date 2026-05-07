# NetworkRouterBgpNeighborCreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAddresses** | Pointer to **[]string** | Source IP addresses for the BGP session | [optional] 
**RouterId** | Pointer to **string** | The router identifier (auto-populated for edge routers) | [optional] 
**Interface** | Pointer to **string** | The interface name for the BGP session (distributed routers only) | [optional] 

## Methods

### NewNetworkRouterBgpNeighborCreateConfig

`func NewNetworkRouterBgpNeighborCreateConfig() *NetworkRouterBgpNeighborCreateConfig`

NewNetworkRouterBgpNeighborCreateConfig instantiates a new NetworkRouterBgpNeighborCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRouterBgpNeighborCreateConfigWithDefaults

`func NewNetworkRouterBgpNeighborCreateConfigWithDefaults() *NetworkRouterBgpNeighborCreateConfig`

NewNetworkRouterBgpNeighborCreateConfigWithDefaults instantiates a new NetworkRouterBgpNeighborCreateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAddresses

`func (o *NetworkRouterBgpNeighborCreateConfig) GetSourceAddresses() []string`

GetSourceAddresses returns the SourceAddresses field if non-nil, zero value otherwise.

### GetSourceAddressesOk

`func (o *NetworkRouterBgpNeighborCreateConfig) GetSourceAddressesOk() (*[]string, bool)`

GetSourceAddressesOk returns a tuple with the SourceAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddresses

`func (o *NetworkRouterBgpNeighborCreateConfig) SetSourceAddresses(v []string)`

SetSourceAddresses sets SourceAddresses field to given value.

### HasSourceAddresses

`func (o *NetworkRouterBgpNeighborCreateConfig) HasSourceAddresses() bool`

HasSourceAddresses returns a boolean if a field has been set.

### GetRouterId

`func (o *NetworkRouterBgpNeighborCreateConfig) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *NetworkRouterBgpNeighborCreateConfig) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *NetworkRouterBgpNeighborCreateConfig) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *NetworkRouterBgpNeighborCreateConfig) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetInterface

`func (o *NetworkRouterBgpNeighborCreateConfig) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *NetworkRouterBgpNeighborCreateConfig) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *NetworkRouterBgpNeighborCreateConfig) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *NetworkRouterBgpNeighborCreateConfig) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



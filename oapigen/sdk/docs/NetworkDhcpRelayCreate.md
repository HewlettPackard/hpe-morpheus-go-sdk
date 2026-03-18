# NetworkDhcpRelayCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ServerIpAddresses** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNetworkDhcpRelayCreate

`func NewNetworkDhcpRelayCreate() *NetworkDhcpRelayCreate`

NewNetworkDhcpRelayCreate instantiates a new NetworkDhcpRelayCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDhcpRelayCreateWithDefaults

`func NewNetworkDhcpRelayCreateWithDefaults() *NetworkDhcpRelayCreate`

NewNetworkDhcpRelayCreateWithDefaults instantiates a new NetworkDhcpRelayCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkDhcpRelayCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkDhcpRelayCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkDhcpRelayCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkDhcpRelayCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerIpAddresses

`func (o *NetworkDhcpRelayCreate) GetServerIpAddresses() []string`

GetServerIpAddresses returns the ServerIpAddresses field if non-nil, zero value otherwise.

### GetServerIpAddressesOk

`func (o *NetworkDhcpRelayCreate) GetServerIpAddressesOk() (*[]string, bool)`

GetServerIpAddressesOk returns a tuple with the ServerIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpAddresses

`func (o *NetworkDhcpRelayCreate) SetServerIpAddresses(v []string)`

SetServerIpAddresses sets ServerIpAddresses field to given value.

### HasServerIpAddresses

`func (o *NetworkDhcpRelayCreate) HasServerIpAddresses() bool`

HasServerIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



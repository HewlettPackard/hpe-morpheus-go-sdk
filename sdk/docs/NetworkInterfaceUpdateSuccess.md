# NetworkInterfaceUpdateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInterface** | Pointer to [**UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface**](UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface.md) |  | [optional] 
**InterfaceType** | Pointer to **string** |  | [optional] 
**NetId** | Pointer to **int64** |  | [optional] 
**Server** | Pointer to [**UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer**](UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer.md) |  | [optional] 

## Methods

### NewNetworkInterfaceUpdateSuccess

`func NewNetworkInterfaceUpdateSuccess() *NetworkInterfaceUpdateSuccess`

NewNetworkInterfaceUpdateSuccess instantiates a new NetworkInterfaceUpdateSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceUpdateSuccessWithDefaults

`func NewNetworkInterfaceUpdateSuccessWithDefaults() *NetworkInterfaceUpdateSuccess`

NewNetworkInterfaceUpdateSuccessWithDefaults instantiates a new NetworkInterfaceUpdateSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInterface

`func (o *NetworkInterfaceUpdateSuccess) GetNetworkInterface() UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface`

GetNetworkInterface returns the NetworkInterface field if non-nil, zero value otherwise.

### GetNetworkInterfaceOk

`func (o *NetworkInterfaceUpdateSuccess) GetNetworkInterfaceOk() (*UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface, bool)`

GetNetworkInterfaceOk returns a tuple with the NetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterface

`func (o *NetworkInterfaceUpdateSuccess) SetNetworkInterface(v UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface)`

SetNetworkInterface sets NetworkInterface field to given value.

### HasNetworkInterface

`func (o *NetworkInterfaceUpdateSuccess) HasNetworkInterface() bool`

HasNetworkInterface returns a boolean if a field has been set.

### GetInterfaceType

`func (o *NetworkInterfaceUpdateSuccess) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *NetworkInterfaceUpdateSuccess) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *NetworkInterfaceUpdateSuccess) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *NetworkInterfaceUpdateSuccess) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetNetId

`func (o *NetworkInterfaceUpdateSuccess) GetNetId() int64`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *NetworkInterfaceUpdateSuccess) GetNetIdOk() (*int64, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *NetworkInterfaceUpdateSuccess) SetNetId(v int64)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *NetworkInterfaceUpdateSuccess) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetServer

`func (o *NetworkInterfaceUpdateSuccess) GetServer() UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *NetworkInterfaceUpdateSuccess) GetServerOk() (*UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *NetworkInterfaceUpdateSuccess) SetServer(v UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *NetworkInterfaceUpdateSuccess) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



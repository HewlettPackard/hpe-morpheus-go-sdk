# CreateNetworkDhcpServerRequestNetworkDhcpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerIpAddress** | **string** | Server Address for the DHCP Server | 
**LeaseTime** | **int64** | Optional lease time for the DHCP Server | [default to 86400]
**Name** | **string** | Name | 
**Config** | [**CreateNetworkDhcpServerRequestNetworkDhcpServerConfig**](CreateNetworkDhcpServerRequestNetworkDhcpServerConfig.md) |  | 

## Methods

### NewCreateNetworkDhcpServerRequestNetworkDhcpServer

`func NewCreateNetworkDhcpServerRequestNetworkDhcpServer(serverIpAddress string, leaseTime int64, name string, config CreateNetworkDhcpServerRequestNetworkDhcpServerConfig, ) *CreateNetworkDhcpServerRequestNetworkDhcpServer`

NewCreateNetworkDhcpServerRequestNetworkDhcpServer instantiates a new CreateNetworkDhcpServerRequestNetworkDhcpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDhcpServerRequestNetworkDhcpServerWithDefaults

`func NewCreateNetworkDhcpServerRequestNetworkDhcpServerWithDefaults() *CreateNetworkDhcpServerRequestNetworkDhcpServer`

NewCreateNetworkDhcpServerRequestNetworkDhcpServerWithDefaults instantiates a new CreateNetworkDhcpServerRequestNetworkDhcpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerIpAddress

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServer) GetServerIpAddress() string`

GetServerIpAddress returns the ServerIpAddress field if non-nil, zero value otherwise.

### GetServerIpAddressOk

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServer) GetServerIpAddressOk() (*string, bool)`

GetServerIpAddressOk returns a tuple with the ServerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpAddress

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServer) SetServerIpAddress(v string)`

SetServerIpAddress sets ServerIpAddress field to given value.


### GetLeaseTime

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServer) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServer) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServer) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.


### GetName

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServer) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServer) GetConfig() CreateNetworkDhcpServerRequestNetworkDhcpServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServer) GetConfigOk() (*CreateNetworkDhcpServerRequestNetworkDhcpServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServer) SetConfig(v CreateNetworkDhcpServerRequestNetworkDhcpServerConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



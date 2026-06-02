# NetworkDhcpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**ServerIpAddress** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LeaseTime** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig**](GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig.md) |  | [optional] 
**Owner** | Pointer to [**GetNetworkDhcpServer200ResponseNetworkDhcpServerOwner**](GetNetworkDhcpServer200ResponseNetworkDhcpServerOwner.md) |  | [optional] 
**NetworkServer** | Pointer to [**GetNetworkDhcpServer200ResponseNetworkDhcpServerNetworkServer**](GetNetworkDhcpServer200ResponseNetworkDhcpServerNetworkServer.md) |  | [optional] 

## Methods

### NewNetworkDhcpServer

`func NewNetworkDhcpServer() *NetworkDhcpServer`

NewNetworkDhcpServer instantiates a new NetworkDhcpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *NetworkDhcpServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDhcpServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDhcpServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkDhcpServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *NetworkDhcpServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NetworkDhcpServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NetworkDhcpServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NetworkDhcpServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetProviderId

`func (o *NetworkDhcpServer) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *NetworkDhcpServer) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *NetworkDhcpServer) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *NetworkDhcpServer) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetServerIpAddress

`func (o *NetworkDhcpServer) GetServerIpAddress() string`

GetServerIpAddress returns the ServerIpAddress field if non-nil, zero value otherwise.

### GetServerIpAddressOk

`func (o *NetworkDhcpServer) GetServerIpAddressOk() (*string, bool)`

GetServerIpAddressOk returns a tuple with the ServerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpAddress

`func (o *NetworkDhcpServer) SetServerIpAddress(v string)`

SetServerIpAddress sets ServerIpAddress field to given value.

### HasServerIpAddress

`func (o *NetworkDhcpServer) HasServerIpAddress() bool`

HasServerIpAddress returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NetworkDhcpServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NetworkDhcpServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NetworkDhcpServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NetworkDhcpServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLeaseTime

`func (o *NetworkDhcpServer) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *NetworkDhcpServer) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *NetworkDhcpServer) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *NetworkDhcpServer) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetName

`func (o *NetworkDhcpServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkDhcpServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkDhcpServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkDhcpServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *NetworkDhcpServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NetworkDhcpServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NetworkDhcpServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NetworkDhcpServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkDhcpServer) GetConfig() GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkDhcpServer) GetConfigOk() (*GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkDhcpServer) SetConfig(v GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkDhcpServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOwner

`func (o *NetworkDhcpServer) GetOwner() GetNetworkDhcpServer200ResponseNetworkDhcpServerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NetworkDhcpServer) GetOwnerOk() (*GetNetworkDhcpServer200ResponseNetworkDhcpServerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NetworkDhcpServer) SetOwner(v GetNetworkDhcpServer200ResponseNetworkDhcpServerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NetworkDhcpServer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNetworkServer

`func (o *NetworkDhcpServer) GetNetworkServer() GetNetworkDhcpServer200ResponseNetworkDhcpServerNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *NetworkDhcpServer) GetNetworkServerOk() (*GetNetworkDhcpServer200ResponseNetworkDhcpServerNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *NetworkDhcpServer) SetNetworkServer(v GetNetworkDhcpServer200ResponseNetworkDhcpServerNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *NetworkDhcpServer) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



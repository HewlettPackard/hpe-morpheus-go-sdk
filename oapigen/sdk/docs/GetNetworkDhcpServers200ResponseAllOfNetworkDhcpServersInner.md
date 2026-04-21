# GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner

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
**Config** | Pointer to [**GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerConfig**](GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerConfig.md) |  | [optional] 
**Owner** | Pointer to [**GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner**](GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner.md) |  | [optional] 
**NetworkServer** | Pointer to [**GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerNetworkServer**](GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerNetworkServer.md) |  | [optional] 

## Methods

### NewGetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner

`func NewGetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner() *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner`

NewGetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner instantiates a new GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerWithDefaults

`func NewGetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerWithDefaults() *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner`

NewGetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerWithDefaults instantiates a new GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetProviderId

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetServerIpAddress

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetServerIpAddress() string`

GetServerIpAddress returns the ServerIpAddress field if non-nil, zero value otherwise.

### GetServerIpAddressOk

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetServerIpAddressOk() (*string, bool)`

GetServerIpAddressOk returns a tuple with the ServerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpAddress

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) SetServerIpAddress(v string)`

SetServerIpAddress sets ServerIpAddress field to given value.

### HasServerIpAddress

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) HasServerIpAddress() bool`

HasServerIpAddress returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLeaseTime

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetConfig() GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetConfigOk() (*GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) SetConfig(v GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOwner

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetOwner() GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetOwnerOk() (*GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) SetOwner(v GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNetworkServer

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetNetworkServer() GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) GetNetworkServerOk() (*GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) SetNetworkServer(v GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetNetworkDhcpServer200ResponseNetworkDhcpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**ServerIpAddress** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LeaseTime** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig**](GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig.md) |  | [optional] 
**Owner** | Pointer to [**CreateLoadBalancerRequestLoadBalancerTenantsInner**](CreateLoadBalancerRequestLoadBalancerTenantsInner.md) |  | [optional] 
**NetworkServer** | Pointer to [**CreateLoadBalancerRequestLoadBalancerTenantsInner**](CreateLoadBalancerRequestLoadBalancerTenantsInner.md) |  | [optional] 

## Methods

### NewGetNetworkDhcpServer200ResponseNetworkDhcpServer

`func NewGetNetworkDhcpServer200ResponseNetworkDhcpServer() *GetNetworkDhcpServer200ResponseNetworkDhcpServer`

NewGetNetworkDhcpServer200ResponseNetworkDhcpServer instantiates a new GetNetworkDhcpServer200ResponseNetworkDhcpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkDhcpServer200ResponseNetworkDhcpServerWithDefaults

`func NewGetNetworkDhcpServer200ResponseNetworkDhcpServerWithDefaults() *GetNetworkDhcpServer200ResponseNetworkDhcpServer`

NewGetNetworkDhcpServer200ResponseNetworkDhcpServerWithDefaults instantiates a new GetNetworkDhcpServer200ResponseNetworkDhcpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetProviderId

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetServerIpAddress

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetServerIpAddress() string`

GetServerIpAddress returns the ServerIpAddress field if non-nil, zero value otherwise.

### GetServerIpAddressOk

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetServerIpAddressOk() (*string, bool)`

GetServerIpAddressOk returns a tuple with the ServerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpAddress

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) SetServerIpAddress(v string)`

SetServerIpAddress sets ServerIpAddress field to given value.

### HasServerIpAddress

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) HasServerIpAddress() bool`

HasServerIpAddress returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLeaseTime

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetLeaseTime() int32`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetLeaseTimeOk() (*int32, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) SetLeaseTime(v int32)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetConfig() GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetConfigOk() (*GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) SetConfig(v GetNetworkDhcpServer200ResponseNetworkDhcpServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOwner

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetOwner() CreateLoadBalancerRequestLoadBalancerTenantsInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetOwnerOk() (*CreateLoadBalancerRequestLoadBalancerTenantsInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) SetOwner(v CreateLoadBalancerRequestLoadBalancerTenantsInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNetworkServer

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetNetworkServer() CreateLoadBalancerRequestLoadBalancerTenantsInner`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) GetNetworkServerOk() (*CreateLoadBalancerRequestLoadBalancerTenantsInner, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) SetNetworkServer(v CreateLoadBalancerRequestLoadBalancerTenantsInner)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *GetNetworkDhcpServer200ResponseNetworkDhcpServer) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetNetworkDhcpRelay200ResponseNetworkDhcpRelay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ServerIpAddresses** | Pointer to **[]string** |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**NetworkServer** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 

## Methods

### NewGetNetworkDhcpRelay200ResponseNetworkDhcpRelay

`func NewGetNetworkDhcpRelay200ResponseNetworkDhcpRelay() *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay`

NewGetNetworkDhcpRelay200ResponseNetworkDhcpRelay instantiates a new GetNetworkDhcpRelay200ResponseNetworkDhcpRelay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkDhcpRelay200ResponseNetworkDhcpRelayWithDefaults

`func NewGetNetworkDhcpRelay200ResponseNetworkDhcpRelayWithDefaults() *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay`

NewGetNetworkDhcpRelay200ResponseNetworkDhcpRelayWithDefaults instantiates a new GetNetworkDhcpRelay200ResponseNetworkDhcpRelay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetProviderId

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetServerIpAddresses

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetServerIpAddresses() []string`

GetServerIpAddresses returns the ServerIpAddresses field if non-nil, zero value otherwise.

### GetServerIpAddressesOk

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetServerIpAddressesOk() (*[]string, bool)`

GetServerIpAddressesOk returns a tuple with the ServerIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpAddresses

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetServerIpAddresses(v []string)`

SetServerIpAddresses sets ServerIpAddresses field to given value.

### HasServerIpAddresses

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) HasServerIpAddresses() bool`

HasServerIpAddresses returns a boolean if a field has been set.

### GetOwner

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetOwner() GetAlerts200ResponseAllOfChecksInnerAccount`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetOwner(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNetworkServer

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetNetworkServer() GetAlerts200ResponseAllOfChecksInnerAccount`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetNetworkServerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetNetworkServer(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



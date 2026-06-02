# CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | IP address of the BGP neighbor | 
**Description** | Pointer to **string** | Description of the BGP neighbor | [optional] 
**ForwardingAddress** | Pointer to **string** | Forwarding address (NSX-V distributed router) | [optional] 
**ProtocolAddress** | Pointer to **string** | Protocol address (NSX-V distributed router) | [optional] 
**RemoteAs** | Pointer to **string** | Remote AS number | [optional] 
**Weight** | Pointer to **int64** | Route weight | [optional] [default to 60]
**KeepAlive** | Pointer to **int64** | Keep-alive interval in seconds | [optional] [default to 60]
**HoldDown** | Pointer to **int64** | Hold-down timer in seconds | [optional] [default to 180]
**Password** | Pointer to **string** | BGP session password | [optional] 
**RouteFilteringType** | Pointer to **string** | Address family for route filtering (e.g. IPV4, IPV6) | [optional] 
**RouteFilteringIn** | Pointer to **string** | Inbound route filter name | [optional] 
**RouteFilteringOut** | Pointer to **string** | Outbound route filter name | [optional] 
**BfdEnabled** | Pointer to [**CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborBfdEnabled**](CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborBfdEnabled.md) |  | [optional] 
**BfdInterval** | Pointer to **int64** | BFD interval in milliseconds | [optional] 
**BfdMultiple** | Pointer to **int64** | BFD multiplier | [optional] 
**AllowAsIn** | Pointer to [**CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborAllowAsIn**](CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborAllowAsIn.md) |  | [optional] 
**HopLimit** | Pointer to **int64** | Maximum hop limit | [optional] [default to 1]
**RestartMode** | Pointer to **string** | Graceful restart mode (e.g. HELPER_ONLY, GRACEFUL_RESTART, DISABLE) | [optional] 
**Config** | Pointer to [**CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborConfig**](CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborConfig.md) |  | [optional] 

## Methods

### NewCreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor

`func NewCreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor(ipAddress string, ) *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor`

NewCreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor instantiates a new CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetIpAddress

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetDescription

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForwardingAddress

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetForwardingAddress() string`

GetForwardingAddress returns the ForwardingAddress field if non-nil, zero value otherwise.

### GetForwardingAddressOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetForwardingAddressOk() (*string, bool)`

GetForwardingAddressOk returns a tuple with the ForwardingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingAddress

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetForwardingAddress(v string)`

SetForwardingAddress sets ForwardingAddress field to given value.

### HasForwardingAddress

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasForwardingAddress() bool`

HasForwardingAddress returns a boolean if a field has been set.

### GetProtocolAddress

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetProtocolAddress() string`

GetProtocolAddress returns the ProtocolAddress field if non-nil, zero value otherwise.

### GetProtocolAddressOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetProtocolAddressOk() (*string, bool)`

GetProtocolAddressOk returns a tuple with the ProtocolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAddress

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetProtocolAddress(v string)`

SetProtocolAddress sets ProtocolAddress field to given value.

### HasProtocolAddress

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasProtocolAddress() bool`

HasProtocolAddress returns a boolean if a field has been set.

### GetRemoteAs

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRemoteAs() string`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRemoteAsOk() (*string, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetRemoteAs(v string)`

SetRemoteAs sets RemoteAs field to given value.

### HasRemoteAs

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasRemoteAs() bool`

HasRemoteAs returns a boolean if a field has been set.

### GetWeight

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetKeepAlive

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetKeepAlive() int64`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetKeepAliveOk() (*int64, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetKeepAlive(v int64)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetHoldDown

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetHoldDown() int64`

GetHoldDown returns the HoldDown field if non-nil, zero value otherwise.

### GetHoldDownOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetHoldDownOk() (*int64, bool)`

GetHoldDownOk returns a tuple with the HoldDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldDown

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetHoldDown(v int64)`

SetHoldDown sets HoldDown field to given value.

### HasHoldDown

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasHoldDown() bool`

HasHoldDown returns a boolean if a field has been set.

### GetPassword

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRouteFilteringType

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRouteFilteringType() string`

GetRouteFilteringType returns the RouteFilteringType field if non-nil, zero value otherwise.

### GetRouteFilteringTypeOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRouteFilteringTypeOk() (*string, bool)`

GetRouteFilteringTypeOk returns a tuple with the RouteFilteringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringType

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetRouteFilteringType(v string)`

SetRouteFilteringType sets RouteFilteringType field to given value.

### HasRouteFilteringType

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasRouteFilteringType() bool`

HasRouteFilteringType returns a boolean if a field has been set.

### GetRouteFilteringIn

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRouteFilteringIn() string`

GetRouteFilteringIn returns the RouteFilteringIn field if non-nil, zero value otherwise.

### GetRouteFilteringInOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRouteFilteringInOk() (*string, bool)`

GetRouteFilteringInOk returns a tuple with the RouteFilteringIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringIn

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetRouteFilteringIn(v string)`

SetRouteFilteringIn sets RouteFilteringIn field to given value.

### HasRouteFilteringIn

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasRouteFilteringIn() bool`

HasRouteFilteringIn returns a boolean if a field has been set.

### GetRouteFilteringOut

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRouteFilteringOut() string`

GetRouteFilteringOut returns the RouteFilteringOut field if non-nil, zero value otherwise.

### GetRouteFilteringOutOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRouteFilteringOutOk() (*string, bool)`

GetRouteFilteringOutOk returns a tuple with the RouteFilteringOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringOut

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetRouteFilteringOut(v string)`

SetRouteFilteringOut sets RouteFilteringOut field to given value.

### HasRouteFilteringOut

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasRouteFilteringOut() bool`

HasRouteFilteringOut returns a boolean if a field has been set.

### GetBfdEnabled

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetBfdEnabled() CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborBfdEnabled`

GetBfdEnabled returns the BfdEnabled field if non-nil, zero value otherwise.

### GetBfdEnabledOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetBfdEnabledOk() (*CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborBfdEnabled, bool)`

GetBfdEnabledOk returns a tuple with the BfdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdEnabled

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetBfdEnabled(v CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborBfdEnabled)`

SetBfdEnabled sets BfdEnabled field to given value.

### HasBfdEnabled

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasBfdEnabled() bool`

HasBfdEnabled returns a boolean if a field has been set.

### GetBfdInterval

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetBfdInterval() int64`

GetBfdInterval returns the BfdInterval field if non-nil, zero value otherwise.

### GetBfdIntervalOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetBfdIntervalOk() (*int64, bool)`

GetBfdIntervalOk returns a tuple with the BfdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdInterval

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetBfdInterval(v int64)`

SetBfdInterval sets BfdInterval field to given value.

### HasBfdInterval

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasBfdInterval() bool`

HasBfdInterval returns a boolean if a field has been set.

### GetBfdMultiple

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetBfdMultiple() int64`

GetBfdMultiple returns the BfdMultiple field if non-nil, zero value otherwise.

### GetBfdMultipleOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetBfdMultipleOk() (*int64, bool)`

GetBfdMultipleOk returns a tuple with the BfdMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdMultiple

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetBfdMultiple(v int64)`

SetBfdMultiple sets BfdMultiple field to given value.

### HasBfdMultiple

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasBfdMultiple() bool`

HasBfdMultiple returns a boolean if a field has been set.

### GetAllowAsIn

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetAllowAsIn() CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborAllowAsIn`

GetAllowAsIn returns the AllowAsIn field if non-nil, zero value otherwise.

### GetAllowAsInOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetAllowAsInOk() (*CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborAllowAsIn, bool)`

GetAllowAsInOk returns a tuple with the AllowAsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAsIn

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetAllowAsIn(v CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborAllowAsIn)`

SetAllowAsIn sets AllowAsIn field to given value.

### HasAllowAsIn

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasAllowAsIn() bool`

HasAllowAsIn returns a boolean if a field has been set.

### GetHopLimit

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetHopLimit() int64`

GetHopLimit returns the HopLimit field if non-nil, zero value otherwise.

### GetHopLimitOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetHopLimitOk() (*int64, bool)`

GetHopLimitOk returns a tuple with the HopLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopLimit

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetHopLimit(v int64)`

SetHopLimit sets HopLimit field to given value.

### HasHopLimit

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasHopLimit() bool`

HasHopLimit returns a boolean if a field has been set.

### GetRestartMode

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRestartMode() string`

GetRestartMode returns the RestartMode field if non-nil, zero value otherwise.

### GetRestartModeOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRestartModeOk() (*string, bool)`

GetRestartModeOk returns a tuple with the RestartMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartMode

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetRestartMode(v string)`

SetRestartMode sets RestartMode field to given value.

### HasRestartMode

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasRestartMode() bool`

HasRestartMode returns a boolean if a field has been set.

### GetConfig

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetConfig() CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetConfigOk() (*CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetConfig(v CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NetworkRouterBgpNeighborCreate

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
**BfdEnabled** | Pointer to [**NetworkRouterBgpNeighborCreateBfdEnabled**](NetworkRouterBgpNeighborCreateBfdEnabled.md) |  | [optional] 
**BfdInterval** | Pointer to **int64** | BFD interval in milliseconds | [optional] 
**BfdMultiple** | Pointer to **int64** | BFD multiplier | [optional] 
**AllowAsIn** | Pointer to [**NetworkRouterBgpNeighborCreateAllowAsIn**](NetworkRouterBgpNeighborCreateAllowAsIn.md) |  | [optional] 
**HopLimit** | Pointer to **int64** | Maximum hop limit | [optional] [default to 1]
**RestartMode** | Pointer to **string** | Graceful restart mode (e.g. HELPER_ONLY, GRACEFUL_RESTART, DISABLE) | [optional] 
**Config** | Pointer to [**NetworkRouterBgpNeighborCreateConfig**](NetworkRouterBgpNeighborCreateConfig.md) |  | [optional] 

## Methods

### NewNetworkRouterBgpNeighborCreate

`func NewNetworkRouterBgpNeighborCreate(ipAddress string, ) *NetworkRouterBgpNeighborCreate`

NewNetworkRouterBgpNeighborCreate instantiates a new NetworkRouterBgpNeighborCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetIpAddress

`func (o *NetworkRouterBgpNeighborCreate) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkRouterBgpNeighborCreate) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkRouterBgpNeighborCreate) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetDescription

`func (o *NetworkRouterBgpNeighborCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkRouterBgpNeighborCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkRouterBgpNeighborCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkRouterBgpNeighborCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForwardingAddress

`func (o *NetworkRouterBgpNeighborCreate) GetForwardingAddress() string`

GetForwardingAddress returns the ForwardingAddress field if non-nil, zero value otherwise.

### GetForwardingAddressOk

`func (o *NetworkRouterBgpNeighborCreate) GetForwardingAddressOk() (*string, bool)`

GetForwardingAddressOk returns a tuple with the ForwardingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingAddress

`func (o *NetworkRouterBgpNeighborCreate) SetForwardingAddress(v string)`

SetForwardingAddress sets ForwardingAddress field to given value.

### HasForwardingAddress

`func (o *NetworkRouterBgpNeighborCreate) HasForwardingAddress() bool`

HasForwardingAddress returns a boolean if a field has been set.

### GetProtocolAddress

`func (o *NetworkRouterBgpNeighborCreate) GetProtocolAddress() string`

GetProtocolAddress returns the ProtocolAddress field if non-nil, zero value otherwise.

### GetProtocolAddressOk

`func (o *NetworkRouterBgpNeighborCreate) GetProtocolAddressOk() (*string, bool)`

GetProtocolAddressOk returns a tuple with the ProtocolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAddress

`func (o *NetworkRouterBgpNeighborCreate) SetProtocolAddress(v string)`

SetProtocolAddress sets ProtocolAddress field to given value.

### HasProtocolAddress

`func (o *NetworkRouterBgpNeighborCreate) HasProtocolAddress() bool`

HasProtocolAddress returns a boolean if a field has been set.

### GetRemoteAs

`func (o *NetworkRouterBgpNeighborCreate) GetRemoteAs() string`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *NetworkRouterBgpNeighborCreate) GetRemoteAsOk() (*string, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *NetworkRouterBgpNeighborCreate) SetRemoteAs(v string)`

SetRemoteAs sets RemoteAs field to given value.

### HasRemoteAs

`func (o *NetworkRouterBgpNeighborCreate) HasRemoteAs() bool`

HasRemoteAs returns a boolean if a field has been set.

### GetWeight

`func (o *NetworkRouterBgpNeighborCreate) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *NetworkRouterBgpNeighborCreate) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *NetworkRouterBgpNeighborCreate) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *NetworkRouterBgpNeighborCreate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetKeepAlive

`func (o *NetworkRouterBgpNeighborCreate) GetKeepAlive() int64`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *NetworkRouterBgpNeighborCreate) GetKeepAliveOk() (*int64, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *NetworkRouterBgpNeighborCreate) SetKeepAlive(v int64)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *NetworkRouterBgpNeighborCreate) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetHoldDown

`func (o *NetworkRouterBgpNeighborCreate) GetHoldDown() int64`

GetHoldDown returns the HoldDown field if non-nil, zero value otherwise.

### GetHoldDownOk

`func (o *NetworkRouterBgpNeighborCreate) GetHoldDownOk() (*int64, bool)`

GetHoldDownOk returns a tuple with the HoldDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldDown

`func (o *NetworkRouterBgpNeighborCreate) SetHoldDown(v int64)`

SetHoldDown sets HoldDown field to given value.

### HasHoldDown

`func (o *NetworkRouterBgpNeighborCreate) HasHoldDown() bool`

HasHoldDown returns a boolean if a field has been set.

### GetPassword

`func (o *NetworkRouterBgpNeighborCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NetworkRouterBgpNeighborCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NetworkRouterBgpNeighborCreate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NetworkRouterBgpNeighborCreate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRouteFilteringType

`func (o *NetworkRouterBgpNeighborCreate) GetRouteFilteringType() string`

GetRouteFilteringType returns the RouteFilteringType field if non-nil, zero value otherwise.

### GetRouteFilteringTypeOk

`func (o *NetworkRouterBgpNeighborCreate) GetRouteFilteringTypeOk() (*string, bool)`

GetRouteFilteringTypeOk returns a tuple with the RouteFilteringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringType

`func (o *NetworkRouterBgpNeighborCreate) SetRouteFilteringType(v string)`

SetRouteFilteringType sets RouteFilteringType field to given value.

### HasRouteFilteringType

`func (o *NetworkRouterBgpNeighborCreate) HasRouteFilteringType() bool`

HasRouteFilteringType returns a boolean if a field has been set.

### GetRouteFilteringIn

`func (o *NetworkRouterBgpNeighborCreate) GetRouteFilteringIn() string`

GetRouteFilteringIn returns the RouteFilteringIn field if non-nil, zero value otherwise.

### GetRouteFilteringInOk

`func (o *NetworkRouterBgpNeighborCreate) GetRouteFilteringInOk() (*string, bool)`

GetRouteFilteringInOk returns a tuple with the RouteFilteringIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringIn

`func (o *NetworkRouterBgpNeighborCreate) SetRouteFilteringIn(v string)`

SetRouteFilteringIn sets RouteFilteringIn field to given value.

### HasRouteFilteringIn

`func (o *NetworkRouterBgpNeighborCreate) HasRouteFilteringIn() bool`

HasRouteFilteringIn returns a boolean if a field has been set.

### GetRouteFilteringOut

`func (o *NetworkRouterBgpNeighborCreate) GetRouteFilteringOut() string`

GetRouteFilteringOut returns the RouteFilteringOut field if non-nil, zero value otherwise.

### GetRouteFilteringOutOk

`func (o *NetworkRouterBgpNeighborCreate) GetRouteFilteringOutOk() (*string, bool)`

GetRouteFilteringOutOk returns a tuple with the RouteFilteringOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringOut

`func (o *NetworkRouterBgpNeighborCreate) SetRouteFilteringOut(v string)`

SetRouteFilteringOut sets RouteFilteringOut field to given value.

### HasRouteFilteringOut

`func (o *NetworkRouterBgpNeighborCreate) HasRouteFilteringOut() bool`

HasRouteFilteringOut returns a boolean if a field has been set.

### GetBfdEnabled

`func (o *NetworkRouterBgpNeighborCreate) GetBfdEnabled() NetworkRouterBgpNeighborCreateBfdEnabled`

GetBfdEnabled returns the BfdEnabled field if non-nil, zero value otherwise.

### GetBfdEnabledOk

`func (o *NetworkRouterBgpNeighborCreate) GetBfdEnabledOk() (*NetworkRouterBgpNeighborCreateBfdEnabled, bool)`

GetBfdEnabledOk returns a tuple with the BfdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdEnabled

`func (o *NetworkRouterBgpNeighborCreate) SetBfdEnabled(v NetworkRouterBgpNeighborCreateBfdEnabled)`

SetBfdEnabled sets BfdEnabled field to given value.

### HasBfdEnabled

`func (o *NetworkRouterBgpNeighborCreate) HasBfdEnabled() bool`

HasBfdEnabled returns a boolean if a field has been set.

### GetBfdInterval

`func (o *NetworkRouterBgpNeighborCreate) GetBfdInterval() int64`

GetBfdInterval returns the BfdInterval field if non-nil, zero value otherwise.

### GetBfdIntervalOk

`func (o *NetworkRouterBgpNeighborCreate) GetBfdIntervalOk() (*int64, bool)`

GetBfdIntervalOk returns a tuple with the BfdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdInterval

`func (o *NetworkRouterBgpNeighborCreate) SetBfdInterval(v int64)`

SetBfdInterval sets BfdInterval field to given value.

### HasBfdInterval

`func (o *NetworkRouterBgpNeighborCreate) HasBfdInterval() bool`

HasBfdInterval returns a boolean if a field has been set.

### GetBfdMultiple

`func (o *NetworkRouterBgpNeighborCreate) GetBfdMultiple() int64`

GetBfdMultiple returns the BfdMultiple field if non-nil, zero value otherwise.

### GetBfdMultipleOk

`func (o *NetworkRouterBgpNeighborCreate) GetBfdMultipleOk() (*int64, bool)`

GetBfdMultipleOk returns a tuple with the BfdMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdMultiple

`func (o *NetworkRouterBgpNeighborCreate) SetBfdMultiple(v int64)`

SetBfdMultiple sets BfdMultiple field to given value.

### HasBfdMultiple

`func (o *NetworkRouterBgpNeighborCreate) HasBfdMultiple() bool`

HasBfdMultiple returns a boolean if a field has been set.

### GetAllowAsIn

`func (o *NetworkRouterBgpNeighborCreate) GetAllowAsIn() NetworkRouterBgpNeighborCreateAllowAsIn`

GetAllowAsIn returns the AllowAsIn field if non-nil, zero value otherwise.

### GetAllowAsInOk

`func (o *NetworkRouterBgpNeighborCreate) GetAllowAsInOk() (*NetworkRouterBgpNeighborCreateAllowAsIn, bool)`

GetAllowAsInOk returns a tuple with the AllowAsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAsIn

`func (o *NetworkRouterBgpNeighborCreate) SetAllowAsIn(v NetworkRouterBgpNeighborCreateAllowAsIn)`

SetAllowAsIn sets AllowAsIn field to given value.

### HasAllowAsIn

`func (o *NetworkRouterBgpNeighborCreate) HasAllowAsIn() bool`

HasAllowAsIn returns a boolean if a field has been set.

### GetHopLimit

`func (o *NetworkRouterBgpNeighborCreate) GetHopLimit() int64`

GetHopLimit returns the HopLimit field if non-nil, zero value otherwise.

### GetHopLimitOk

`func (o *NetworkRouterBgpNeighborCreate) GetHopLimitOk() (*int64, bool)`

GetHopLimitOk returns a tuple with the HopLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopLimit

`func (o *NetworkRouterBgpNeighborCreate) SetHopLimit(v int64)`

SetHopLimit sets HopLimit field to given value.

### HasHopLimit

`func (o *NetworkRouterBgpNeighborCreate) HasHopLimit() bool`

HasHopLimit returns a boolean if a field has been set.

### GetRestartMode

`func (o *NetworkRouterBgpNeighborCreate) GetRestartMode() string`

GetRestartMode returns the RestartMode field if non-nil, zero value otherwise.

### GetRestartModeOk

`func (o *NetworkRouterBgpNeighborCreate) GetRestartModeOk() (*string, bool)`

GetRestartModeOk returns a tuple with the RestartMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartMode

`func (o *NetworkRouterBgpNeighborCreate) SetRestartMode(v string)`

SetRestartMode sets RestartMode field to given value.

### HasRestartMode

`func (o *NetworkRouterBgpNeighborCreate) HasRestartMode() bool`

HasRestartMode returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkRouterBgpNeighborCreate) GetConfig() NetworkRouterBgpNeighborCreateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkRouterBgpNeighborCreate) GetConfigOk() (*NetworkRouterBgpNeighborCreateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkRouterBgpNeighborCreate) SetConfig(v NetworkRouterBgpNeighborCreateConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkRouterBgpNeighborCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



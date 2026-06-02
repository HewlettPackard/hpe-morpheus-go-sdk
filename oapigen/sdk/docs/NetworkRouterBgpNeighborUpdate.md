# NetworkRouterBgpNeighborUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | IP address of the BGP neighbor | [optional] 
**Description** | Pointer to **string** | Description of the BGP neighbor | [optional] 
**ForwardingAddress** | Pointer to **string** | Forwarding address (NSX-V distributed router) | [optional] 
**ProtocolAddress** | Pointer to **string** | Protocol address (NSX-V distributed router) | [optional] 
**RemoteAs** | Pointer to **string** | Remote AS number | [optional] 
**Weight** | Pointer to **int64** | Route weight | [optional] 
**KeepAlive** | Pointer to **int64** | Keep-alive interval in seconds | [optional] 
**HoldDown** | Pointer to **int64** | Hold-down timer in seconds | [optional] 
**Password** | Pointer to **string** | BGP session password | [optional] 
**RouteFilteringType** | Pointer to **string** | Address family for route filtering (e.g. IPV4, IPV6) | [optional] 
**RouteFilteringIn** | Pointer to **string** | Inbound route filter name | [optional] 
**RouteFilteringOut** | Pointer to **string** | Outbound route filter name | [optional] 
**BfdEnabled** | Pointer to [**NetworkRouterBgpNeighborUpdateBfdEnabled**](NetworkRouterBgpNeighborUpdateBfdEnabled.md) |  | [optional] 
**BfdInterval** | Pointer to **int64** | BFD interval in milliseconds | [optional] 
**BfdMultiple** | Pointer to **int64** | BFD multiplier | [optional] 
**AllowAsIn** | Pointer to [**NetworkRouterBgpNeighborUpdateAllowAsIn**](NetworkRouterBgpNeighborUpdateAllowAsIn.md) |  | [optional] 
**HopLimit** | Pointer to **int64** | Maximum hop limit | [optional] 
**RestartMode** | Pointer to **string** | Graceful restart mode (e.g. HELPER_ONLY, GRACEFUL_RESTART, DISABLE) | [optional] 
**Config** | Pointer to [**NetworkRouterBgpNeighborUpdateConfig**](NetworkRouterBgpNeighborUpdateConfig.md) |  | [optional] 

## Methods

### NewNetworkRouterBgpNeighborUpdate

`func NewNetworkRouterBgpNeighborUpdate() *NetworkRouterBgpNeighborUpdate`

NewNetworkRouterBgpNeighborUpdate instantiates a new NetworkRouterBgpNeighborUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetIpAddress

`func (o *NetworkRouterBgpNeighborUpdate) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkRouterBgpNeighborUpdate) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkRouterBgpNeighborUpdate) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkRouterBgpNeighborUpdate) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkRouterBgpNeighborUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkRouterBgpNeighborUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkRouterBgpNeighborUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkRouterBgpNeighborUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForwardingAddress

`func (o *NetworkRouterBgpNeighborUpdate) GetForwardingAddress() string`

GetForwardingAddress returns the ForwardingAddress field if non-nil, zero value otherwise.

### GetForwardingAddressOk

`func (o *NetworkRouterBgpNeighborUpdate) GetForwardingAddressOk() (*string, bool)`

GetForwardingAddressOk returns a tuple with the ForwardingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingAddress

`func (o *NetworkRouterBgpNeighborUpdate) SetForwardingAddress(v string)`

SetForwardingAddress sets ForwardingAddress field to given value.

### HasForwardingAddress

`func (o *NetworkRouterBgpNeighborUpdate) HasForwardingAddress() bool`

HasForwardingAddress returns a boolean if a field has been set.

### GetProtocolAddress

`func (o *NetworkRouterBgpNeighborUpdate) GetProtocolAddress() string`

GetProtocolAddress returns the ProtocolAddress field if non-nil, zero value otherwise.

### GetProtocolAddressOk

`func (o *NetworkRouterBgpNeighborUpdate) GetProtocolAddressOk() (*string, bool)`

GetProtocolAddressOk returns a tuple with the ProtocolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAddress

`func (o *NetworkRouterBgpNeighborUpdate) SetProtocolAddress(v string)`

SetProtocolAddress sets ProtocolAddress field to given value.

### HasProtocolAddress

`func (o *NetworkRouterBgpNeighborUpdate) HasProtocolAddress() bool`

HasProtocolAddress returns a boolean if a field has been set.

### GetRemoteAs

`func (o *NetworkRouterBgpNeighborUpdate) GetRemoteAs() string`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *NetworkRouterBgpNeighborUpdate) GetRemoteAsOk() (*string, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *NetworkRouterBgpNeighborUpdate) SetRemoteAs(v string)`

SetRemoteAs sets RemoteAs field to given value.

### HasRemoteAs

`func (o *NetworkRouterBgpNeighborUpdate) HasRemoteAs() bool`

HasRemoteAs returns a boolean if a field has been set.

### GetWeight

`func (o *NetworkRouterBgpNeighborUpdate) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *NetworkRouterBgpNeighborUpdate) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *NetworkRouterBgpNeighborUpdate) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *NetworkRouterBgpNeighborUpdate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetKeepAlive

`func (o *NetworkRouterBgpNeighborUpdate) GetKeepAlive() int64`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *NetworkRouterBgpNeighborUpdate) GetKeepAliveOk() (*int64, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *NetworkRouterBgpNeighborUpdate) SetKeepAlive(v int64)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *NetworkRouterBgpNeighborUpdate) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetHoldDown

`func (o *NetworkRouterBgpNeighborUpdate) GetHoldDown() int64`

GetHoldDown returns the HoldDown field if non-nil, zero value otherwise.

### GetHoldDownOk

`func (o *NetworkRouterBgpNeighborUpdate) GetHoldDownOk() (*int64, bool)`

GetHoldDownOk returns a tuple with the HoldDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldDown

`func (o *NetworkRouterBgpNeighborUpdate) SetHoldDown(v int64)`

SetHoldDown sets HoldDown field to given value.

### HasHoldDown

`func (o *NetworkRouterBgpNeighborUpdate) HasHoldDown() bool`

HasHoldDown returns a boolean if a field has been set.

### GetPassword

`func (o *NetworkRouterBgpNeighborUpdate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NetworkRouterBgpNeighborUpdate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NetworkRouterBgpNeighborUpdate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NetworkRouterBgpNeighborUpdate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRouteFilteringType

`func (o *NetworkRouterBgpNeighborUpdate) GetRouteFilteringType() string`

GetRouteFilteringType returns the RouteFilteringType field if non-nil, zero value otherwise.

### GetRouteFilteringTypeOk

`func (o *NetworkRouterBgpNeighborUpdate) GetRouteFilteringTypeOk() (*string, bool)`

GetRouteFilteringTypeOk returns a tuple with the RouteFilteringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringType

`func (o *NetworkRouterBgpNeighborUpdate) SetRouteFilteringType(v string)`

SetRouteFilteringType sets RouteFilteringType field to given value.

### HasRouteFilteringType

`func (o *NetworkRouterBgpNeighborUpdate) HasRouteFilteringType() bool`

HasRouteFilteringType returns a boolean if a field has been set.

### GetRouteFilteringIn

`func (o *NetworkRouterBgpNeighborUpdate) GetRouteFilteringIn() string`

GetRouteFilteringIn returns the RouteFilteringIn field if non-nil, zero value otherwise.

### GetRouteFilteringInOk

`func (o *NetworkRouterBgpNeighborUpdate) GetRouteFilteringInOk() (*string, bool)`

GetRouteFilteringInOk returns a tuple with the RouteFilteringIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringIn

`func (o *NetworkRouterBgpNeighborUpdate) SetRouteFilteringIn(v string)`

SetRouteFilteringIn sets RouteFilteringIn field to given value.

### HasRouteFilteringIn

`func (o *NetworkRouterBgpNeighborUpdate) HasRouteFilteringIn() bool`

HasRouteFilteringIn returns a boolean if a field has been set.

### GetRouteFilteringOut

`func (o *NetworkRouterBgpNeighborUpdate) GetRouteFilteringOut() string`

GetRouteFilteringOut returns the RouteFilteringOut field if non-nil, zero value otherwise.

### GetRouteFilteringOutOk

`func (o *NetworkRouterBgpNeighborUpdate) GetRouteFilteringOutOk() (*string, bool)`

GetRouteFilteringOutOk returns a tuple with the RouteFilteringOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringOut

`func (o *NetworkRouterBgpNeighborUpdate) SetRouteFilteringOut(v string)`

SetRouteFilteringOut sets RouteFilteringOut field to given value.

### HasRouteFilteringOut

`func (o *NetworkRouterBgpNeighborUpdate) HasRouteFilteringOut() bool`

HasRouteFilteringOut returns a boolean if a field has been set.

### GetBfdEnabled

`func (o *NetworkRouterBgpNeighborUpdate) GetBfdEnabled() NetworkRouterBgpNeighborUpdateBfdEnabled`

GetBfdEnabled returns the BfdEnabled field if non-nil, zero value otherwise.

### GetBfdEnabledOk

`func (o *NetworkRouterBgpNeighborUpdate) GetBfdEnabledOk() (*NetworkRouterBgpNeighborUpdateBfdEnabled, bool)`

GetBfdEnabledOk returns a tuple with the BfdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdEnabled

`func (o *NetworkRouterBgpNeighborUpdate) SetBfdEnabled(v NetworkRouterBgpNeighborUpdateBfdEnabled)`

SetBfdEnabled sets BfdEnabled field to given value.

### HasBfdEnabled

`func (o *NetworkRouterBgpNeighborUpdate) HasBfdEnabled() bool`

HasBfdEnabled returns a boolean if a field has been set.

### GetBfdInterval

`func (o *NetworkRouterBgpNeighborUpdate) GetBfdInterval() int64`

GetBfdInterval returns the BfdInterval field if non-nil, zero value otherwise.

### GetBfdIntervalOk

`func (o *NetworkRouterBgpNeighborUpdate) GetBfdIntervalOk() (*int64, bool)`

GetBfdIntervalOk returns a tuple with the BfdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdInterval

`func (o *NetworkRouterBgpNeighborUpdate) SetBfdInterval(v int64)`

SetBfdInterval sets BfdInterval field to given value.

### HasBfdInterval

`func (o *NetworkRouterBgpNeighborUpdate) HasBfdInterval() bool`

HasBfdInterval returns a boolean if a field has been set.

### GetBfdMultiple

`func (o *NetworkRouterBgpNeighborUpdate) GetBfdMultiple() int64`

GetBfdMultiple returns the BfdMultiple field if non-nil, zero value otherwise.

### GetBfdMultipleOk

`func (o *NetworkRouterBgpNeighborUpdate) GetBfdMultipleOk() (*int64, bool)`

GetBfdMultipleOk returns a tuple with the BfdMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdMultiple

`func (o *NetworkRouterBgpNeighborUpdate) SetBfdMultiple(v int64)`

SetBfdMultiple sets BfdMultiple field to given value.

### HasBfdMultiple

`func (o *NetworkRouterBgpNeighborUpdate) HasBfdMultiple() bool`

HasBfdMultiple returns a boolean if a field has been set.

### GetAllowAsIn

`func (o *NetworkRouterBgpNeighborUpdate) GetAllowAsIn() NetworkRouterBgpNeighborUpdateAllowAsIn`

GetAllowAsIn returns the AllowAsIn field if non-nil, zero value otherwise.

### GetAllowAsInOk

`func (o *NetworkRouterBgpNeighborUpdate) GetAllowAsInOk() (*NetworkRouterBgpNeighborUpdateAllowAsIn, bool)`

GetAllowAsInOk returns a tuple with the AllowAsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAsIn

`func (o *NetworkRouterBgpNeighborUpdate) SetAllowAsIn(v NetworkRouterBgpNeighborUpdateAllowAsIn)`

SetAllowAsIn sets AllowAsIn field to given value.

### HasAllowAsIn

`func (o *NetworkRouterBgpNeighborUpdate) HasAllowAsIn() bool`

HasAllowAsIn returns a boolean if a field has been set.

### GetHopLimit

`func (o *NetworkRouterBgpNeighborUpdate) GetHopLimit() int64`

GetHopLimit returns the HopLimit field if non-nil, zero value otherwise.

### GetHopLimitOk

`func (o *NetworkRouterBgpNeighborUpdate) GetHopLimitOk() (*int64, bool)`

GetHopLimitOk returns a tuple with the HopLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopLimit

`func (o *NetworkRouterBgpNeighborUpdate) SetHopLimit(v int64)`

SetHopLimit sets HopLimit field to given value.

### HasHopLimit

`func (o *NetworkRouterBgpNeighborUpdate) HasHopLimit() bool`

HasHopLimit returns a boolean if a field has been set.

### GetRestartMode

`func (o *NetworkRouterBgpNeighborUpdate) GetRestartMode() string`

GetRestartMode returns the RestartMode field if non-nil, zero value otherwise.

### GetRestartModeOk

`func (o *NetworkRouterBgpNeighborUpdate) GetRestartModeOk() (*string, bool)`

GetRestartModeOk returns a tuple with the RestartMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartMode

`func (o *NetworkRouterBgpNeighborUpdate) SetRestartMode(v string)`

SetRestartMode sets RestartMode field to given value.

### HasRestartMode

`func (o *NetworkRouterBgpNeighborUpdate) HasRestartMode() bool`

HasRestartMode returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkRouterBgpNeighborUpdate) GetConfig() NetworkRouterBgpNeighborUpdateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkRouterBgpNeighborUpdate) GetConfigOk() (*NetworkRouterBgpNeighborUpdateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkRouterBgpNeighborUpdate) SetConfig(v NetworkRouterBgpNeighborUpdateConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkRouterBgpNeighborUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



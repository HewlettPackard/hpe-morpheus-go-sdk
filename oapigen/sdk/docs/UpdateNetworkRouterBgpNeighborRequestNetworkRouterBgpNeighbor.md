# UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor

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
**BfdEnabled** | Pointer to [**UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborBfdEnabled**](UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborBfdEnabled.md) |  | [optional] 
**BfdInterval** | Pointer to **int64** | BFD interval in milliseconds | [optional] 
**BfdMultiple** | Pointer to **int64** | BFD multiplier | [optional] 
**AllowAsIn** | Pointer to [**UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborAllowAsIn**](UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborAllowAsIn.md) |  | [optional] 
**HopLimit** | Pointer to **int64** | Maximum hop limit | [optional] 
**RestartMode** | Pointer to **string** | Graceful restart mode (e.g. HELPER_ONLY, GRACEFUL_RESTART, DISABLE) | [optional] 
**Config** | Pointer to [**UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborConfig**](UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborConfig.md) |  | [optional] 

## Methods

### NewUpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor

`func NewUpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor() *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor`

NewUpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor instantiates a new UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborWithDefaults

`func NewUpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborWithDefaults() *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor`

NewUpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborWithDefaults instantiates a new UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForwardingAddress

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetForwardingAddress() string`

GetForwardingAddress returns the ForwardingAddress field if non-nil, zero value otherwise.

### GetForwardingAddressOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetForwardingAddressOk() (*string, bool)`

GetForwardingAddressOk returns a tuple with the ForwardingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingAddress

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetForwardingAddress(v string)`

SetForwardingAddress sets ForwardingAddress field to given value.

### HasForwardingAddress

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasForwardingAddress() bool`

HasForwardingAddress returns a boolean if a field has been set.

### GetProtocolAddress

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetProtocolAddress() string`

GetProtocolAddress returns the ProtocolAddress field if non-nil, zero value otherwise.

### GetProtocolAddressOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetProtocolAddressOk() (*string, bool)`

GetProtocolAddressOk returns a tuple with the ProtocolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAddress

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetProtocolAddress(v string)`

SetProtocolAddress sets ProtocolAddress field to given value.

### HasProtocolAddress

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasProtocolAddress() bool`

HasProtocolAddress returns a boolean if a field has been set.

### GetRemoteAs

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRemoteAs() string`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRemoteAsOk() (*string, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetRemoteAs(v string)`

SetRemoteAs sets RemoteAs field to given value.

### HasRemoteAs

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasRemoteAs() bool`

HasRemoteAs returns a boolean if a field has been set.

### GetWeight

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetKeepAlive

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetKeepAlive() int64`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetKeepAliveOk() (*int64, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetKeepAlive(v int64)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetHoldDown

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetHoldDown() int64`

GetHoldDown returns the HoldDown field if non-nil, zero value otherwise.

### GetHoldDownOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetHoldDownOk() (*int64, bool)`

GetHoldDownOk returns a tuple with the HoldDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldDown

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetHoldDown(v int64)`

SetHoldDown sets HoldDown field to given value.

### HasHoldDown

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasHoldDown() bool`

HasHoldDown returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRouteFilteringType

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRouteFilteringType() string`

GetRouteFilteringType returns the RouteFilteringType field if non-nil, zero value otherwise.

### GetRouteFilteringTypeOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRouteFilteringTypeOk() (*string, bool)`

GetRouteFilteringTypeOk returns a tuple with the RouteFilteringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringType

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetRouteFilteringType(v string)`

SetRouteFilteringType sets RouteFilteringType field to given value.

### HasRouteFilteringType

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasRouteFilteringType() bool`

HasRouteFilteringType returns a boolean if a field has been set.

### GetRouteFilteringIn

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRouteFilteringIn() string`

GetRouteFilteringIn returns the RouteFilteringIn field if non-nil, zero value otherwise.

### GetRouteFilteringInOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRouteFilteringInOk() (*string, bool)`

GetRouteFilteringInOk returns a tuple with the RouteFilteringIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringIn

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetRouteFilteringIn(v string)`

SetRouteFilteringIn sets RouteFilteringIn field to given value.

### HasRouteFilteringIn

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasRouteFilteringIn() bool`

HasRouteFilteringIn returns a boolean if a field has been set.

### GetRouteFilteringOut

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRouteFilteringOut() string`

GetRouteFilteringOut returns the RouteFilteringOut field if non-nil, zero value otherwise.

### GetRouteFilteringOutOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRouteFilteringOutOk() (*string, bool)`

GetRouteFilteringOutOk returns a tuple with the RouteFilteringOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringOut

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetRouteFilteringOut(v string)`

SetRouteFilteringOut sets RouteFilteringOut field to given value.

### HasRouteFilteringOut

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasRouteFilteringOut() bool`

HasRouteFilteringOut returns a boolean if a field has been set.

### GetBfdEnabled

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetBfdEnabled() UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborBfdEnabled`

GetBfdEnabled returns the BfdEnabled field if non-nil, zero value otherwise.

### GetBfdEnabledOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetBfdEnabledOk() (*UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborBfdEnabled, bool)`

GetBfdEnabledOk returns a tuple with the BfdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdEnabled

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetBfdEnabled(v UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborBfdEnabled)`

SetBfdEnabled sets BfdEnabled field to given value.

### HasBfdEnabled

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasBfdEnabled() bool`

HasBfdEnabled returns a boolean if a field has been set.

### GetBfdInterval

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetBfdInterval() int64`

GetBfdInterval returns the BfdInterval field if non-nil, zero value otherwise.

### GetBfdIntervalOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetBfdIntervalOk() (*int64, bool)`

GetBfdIntervalOk returns a tuple with the BfdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdInterval

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetBfdInterval(v int64)`

SetBfdInterval sets BfdInterval field to given value.

### HasBfdInterval

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasBfdInterval() bool`

HasBfdInterval returns a boolean if a field has been set.

### GetBfdMultiple

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetBfdMultiple() int64`

GetBfdMultiple returns the BfdMultiple field if non-nil, zero value otherwise.

### GetBfdMultipleOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetBfdMultipleOk() (*int64, bool)`

GetBfdMultipleOk returns a tuple with the BfdMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdMultiple

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetBfdMultiple(v int64)`

SetBfdMultiple sets BfdMultiple field to given value.

### HasBfdMultiple

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasBfdMultiple() bool`

HasBfdMultiple returns a boolean if a field has been set.

### GetAllowAsIn

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetAllowAsIn() UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborAllowAsIn`

GetAllowAsIn returns the AllowAsIn field if non-nil, zero value otherwise.

### GetAllowAsInOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetAllowAsInOk() (*UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborAllowAsIn, bool)`

GetAllowAsInOk returns a tuple with the AllowAsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAsIn

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetAllowAsIn(v UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborAllowAsIn)`

SetAllowAsIn sets AllowAsIn field to given value.

### HasAllowAsIn

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasAllowAsIn() bool`

HasAllowAsIn returns a boolean if a field has been set.

### GetHopLimit

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetHopLimit() int64`

GetHopLimit returns the HopLimit field if non-nil, zero value otherwise.

### GetHopLimitOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetHopLimitOk() (*int64, bool)`

GetHopLimitOk returns a tuple with the HopLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopLimit

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetHopLimit(v int64)`

SetHopLimit sets HopLimit field to given value.

### HasHopLimit

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasHopLimit() bool`

HasHopLimit returns a boolean if a field has been set.

### GetRestartMode

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRestartMode() string`

GetRestartMode returns the RestartMode field if non-nil, zero value otherwise.

### GetRestartModeOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetRestartModeOk() (*string, bool)`

GetRestartModeOk returns a tuple with the RestartMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartMode

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetRestartMode(v string)`

SetRestartMode sets RestartMode field to given value.

### HasRestartMode

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasRestartMode() bool`

HasRestartMode returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetConfig() UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) GetConfigOk() (*UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) SetConfig(v UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighborConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateNetworkRouterBgpNeighborRequestNetworkRouterBgpNeighbor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



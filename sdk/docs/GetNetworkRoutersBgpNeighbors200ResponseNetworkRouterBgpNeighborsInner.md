# GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**ForwardingAddress** | Pointer to **string** |  | [optional] 
**ProtocolAddress** | Pointer to **string** |  | [optional] 
**RemoteAs** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int64** |  | [optional] 
**KeepAlive** | Pointer to **int64** |  | [optional] 
**HoldDown** | Pointer to **int64** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**RouteFilteringType** | Pointer to **string** |  | [optional] 
**RouteFilteringIn** | Pointer to **string** |  | [optional] 
**RouteFilteringOut** | Pointer to **string** |  | [optional] 
**BfdEnabled** | Pointer to **bool** |  | [optional] 
**BfdInterval** | Pointer to **int64** |  | [optional] 
**BfdMultiple** | Pointer to **int64** |  | [optional] 
**AllowAsIn** | Pointer to **bool** |  | [optional] 
**HopLimit** | Pointer to **int64** |  | [optional] 
**RestartMode** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**SyncSource** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInnerConfig**](GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInnerConfig.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner

`func NewGetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner() *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner`

NewGetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner instantiates a new GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInnerWithDefaults

`func NewGetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInnerWithDefaults() *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner`

NewGetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInnerWithDefaults instantiates a new GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetForwardingAddress

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetForwardingAddress() string`

GetForwardingAddress returns the ForwardingAddress field if non-nil, zero value otherwise.

### GetForwardingAddressOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetForwardingAddressOk() (*string, bool)`

GetForwardingAddressOk returns a tuple with the ForwardingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingAddress

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetForwardingAddress(v string)`

SetForwardingAddress sets ForwardingAddress field to given value.

### HasForwardingAddress

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasForwardingAddress() bool`

HasForwardingAddress returns a boolean if a field has been set.

### GetProtocolAddress

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetProtocolAddress() string`

GetProtocolAddress returns the ProtocolAddress field if non-nil, zero value otherwise.

### GetProtocolAddressOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetProtocolAddressOk() (*string, bool)`

GetProtocolAddressOk returns a tuple with the ProtocolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAddress

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetProtocolAddress(v string)`

SetProtocolAddress sets ProtocolAddress field to given value.

### HasProtocolAddress

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasProtocolAddress() bool`

HasProtocolAddress returns a boolean if a field has been set.

### GetRemoteAs

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRemoteAs() string`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRemoteAsOk() (*string, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetRemoteAs(v string)`

SetRemoteAs sets RemoteAs field to given value.

### HasRemoteAs

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasRemoteAs() bool`

HasRemoteAs returns a boolean if a field has been set.

### GetWeight

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetKeepAlive

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetKeepAlive() int64`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetKeepAliveOk() (*int64, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetKeepAlive(v int64)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetHoldDown

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetHoldDown() int64`

GetHoldDown returns the HoldDown field if non-nil, zero value otherwise.

### GetHoldDownOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetHoldDownOk() (*int64, bool)`

GetHoldDownOk returns a tuple with the HoldDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldDown

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetHoldDown(v int64)`

SetHoldDown sets HoldDown field to given value.

### HasHoldDown

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasHoldDown() bool`

HasHoldDown returns a boolean if a field has been set.

### GetPassword

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRouteFilteringType

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRouteFilteringType() string`

GetRouteFilteringType returns the RouteFilteringType field if non-nil, zero value otherwise.

### GetRouteFilteringTypeOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRouteFilteringTypeOk() (*string, bool)`

GetRouteFilteringTypeOk returns a tuple with the RouteFilteringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringType

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetRouteFilteringType(v string)`

SetRouteFilteringType sets RouteFilteringType field to given value.

### HasRouteFilteringType

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasRouteFilteringType() bool`

HasRouteFilteringType returns a boolean if a field has been set.

### GetRouteFilteringIn

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRouteFilteringIn() string`

GetRouteFilteringIn returns the RouteFilteringIn field if non-nil, zero value otherwise.

### GetRouteFilteringInOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRouteFilteringInOk() (*string, bool)`

GetRouteFilteringInOk returns a tuple with the RouteFilteringIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringIn

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetRouteFilteringIn(v string)`

SetRouteFilteringIn sets RouteFilteringIn field to given value.

### HasRouteFilteringIn

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasRouteFilteringIn() bool`

HasRouteFilteringIn returns a boolean if a field has been set.

### GetRouteFilteringOut

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRouteFilteringOut() string`

GetRouteFilteringOut returns the RouteFilteringOut field if non-nil, zero value otherwise.

### GetRouteFilteringOutOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRouteFilteringOutOk() (*string, bool)`

GetRouteFilteringOutOk returns a tuple with the RouteFilteringOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringOut

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetRouteFilteringOut(v string)`

SetRouteFilteringOut sets RouteFilteringOut field to given value.

### HasRouteFilteringOut

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasRouteFilteringOut() bool`

HasRouteFilteringOut returns a boolean if a field has been set.

### GetBfdEnabled

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetBfdEnabled() bool`

GetBfdEnabled returns the BfdEnabled field if non-nil, zero value otherwise.

### GetBfdEnabledOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetBfdEnabledOk() (*bool, bool)`

GetBfdEnabledOk returns a tuple with the BfdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdEnabled

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetBfdEnabled(v bool)`

SetBfdEnabled sets BfdEnabled field to given value.

### HasBfdEnabled

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasBfdEnabled() bool`

HasBfdEnabled returns a boolean if a field has been set.

### GetBfdInterval

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetBfdInterval() int64`

GetBfdInterval returns the BfdInterval field if non-nil, zero value otherwise.

### GetBfdIntervalOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetBfdIntervalOk() (*int64, bool)`

GetBfdIntervalOk returns a tuple with the BfdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdInterval

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetBfdInterval(v int64)`

SetBfdInterval sets BfdInterval field to given value.

### HasBfdInterval

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasBfdInterval() bool`

HasBfdInterval returns a boolean if a field has been set.

### GetBfdMultiple

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetBfdMultiple() int64`

GetBfdMultiple returns the BfdMultiple field if non-nil, zero value otherwise.

### GetBfdMultipleOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetBfdMultipleOk() (*int64, bool)`

GetBfdMultipleOk returns a tuple with the BfdMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdMultiple

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetBfdMultiple(v int64)`

SetBfdMultiple sets BfdMultiple field to given value.

### HasBfdMultiple

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasBfdMultiple() bool`

HasBfdMultiple returns a boolean if a field has been set.

### GetAllowAsIn

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetAllowAsIn() bool`

GetAllowAsIn returns the AllowAsIn field if non-nil, zero value otherwise.

### GetAllowAsInOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetAllowAsInOk() (*bool, bool)`

GetAllowAsInOk returns a tuple with the AllowAsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAsIn

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetAllowAsIn(v bool)`

SetAllowAsIn sets AllowAsIn field to given value.

### HasAllowAsIn

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasAllowAsIn() bool`

HasAllowAsIn returns a boolean if a field has been set.

### GetHopLimit

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetHopLimit() int64`

GetHopLimit returns the HopLimit field if non-nil, zero value otherwise.

### GetHopLimitOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetHopLimitOk() (*int64, bool)`

GetHopLimitOk returns a tuple with the HopLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopLimit

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetHopLimit(v int64)`

SetHopLimit sets HopLimit field to given value.

### HasHopLimit

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasHopLimit() bool`

HasHopLimit returns a boolean if a field has been set.

### GetRestartMode

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRestartMode() string`

GetRestartMode returns the RestartMode field if non-nil, zero value otherwise.

### GetRestartModeOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRestartModeOk() (*string, bool)`

GetRestartModeOk returns a tuple with the RestartMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartMode

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetRestartMode(v string)`

SetRestartMode sets RestartMode field to given value.

### HasRestartMode

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasRestartMode() bool`

HasRestartMode returns a boolean if a field has been set.

### GetProviderId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetSyncSource

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetSyncSource() string`

GetSyncSource returns the SyncSource field if non-nil, zero value otherwise.

### GetSyncSourceOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetSyncSourceOk() (*string, bool)`

GetSyncSourceOk returns a tuple with the SyncSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSource

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetSyncSource(v string)`

SetSyncSource sets SyncSource field to given value.

### HasSyncSource

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasSyncSource() bool`

HasSyncSource returns a boolean if a field has been set.

### GetInternalId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRefType

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetConfig

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetConfig() GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetConfigOk() (*GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetConfig(v GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



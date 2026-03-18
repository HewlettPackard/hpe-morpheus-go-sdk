# GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**ForwardingAddress** | Pointer to **NullableString** |  | [optional] 
**ProtocolAddress** | Pointer to **NullableString** |  | [optional] 
**RemoteAs** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int64** |  | [optional] 
**KeepAlive** | Pointer to **int64** |  | [optional] 
**HoldDown** | Pointer to **int64** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
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
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to [**GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighborConfig**](GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighborConfig.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor

`func NewGetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor() *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor`

NewGetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor instantiates a new GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighborWithDefaults

`func NewGetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighborWithDefaults() *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor`

NewGetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighborWithDefaults instantiates a new GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetForwardingAddress

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetForwardingAddress() string`

GetForwardingAddress returns the ForwardingAddress field if non-nil, zero value otherwise.

### GetForwardingAddressOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetForwardingAddressOk() (*string, bool)`

GetForwardingAddressOk returns a tuple with the ForwardingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingAddress

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetForwardingAddress(v string)`

SetForwardingAddress sets ForwardingAddress field to given value.

### HasForwardingAddress

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasForwardingAddress() bool`

HasForwardingAddress returns a boolean if a field has been set.

### SetForwardingAddressNil

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetForwardingAddressNil(b bool)`

 SetForwardingAddressNil sets the value for ForwardingAddress to be an explicit nil

### UnsetForwardingAddress
`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) UnsetForwardingAddress()`

UnsetForwardingAddress ensures that no value is present for ForwardingAddress, not even an explicit nil
### GetProtocolAddress

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetProtocolAddress() string`

GetProtocolAddress returns the ProtocolAddress field if non-nil, zero value otherwise.

### GetProtocolAddressOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetProtocolAddressOk() (*string, bool)`

GetProtocolAddressOk returns a tuple with the ProtocolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAddress

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetProtocolAddress(v string)`

SetProtocolAddress sets ProtocolAddress field to given value.

### HasProtocolAddress

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasProtocolAddress() bool`

HasProtocolAddress returns a boolean if a field has been set.

### SetProtocolAddressNil

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetProtocolAddressNil(b bool)`

 SetProtocolAddressNil sets the value for ProtocolAddress to be an explicit nil

### UnsetProtocolAddress
`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) UnsetProtocolAddress()`

UnsetProtocolAddress ensures that no value is present for ProtocolAddress, not even an explicit nil
### GetRemoteAs

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRemoteAs() string`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRemoteAsOk() (*string, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetRemoteAs(v string)`

SetRemoteAs sets RemoteAs field to given value.

### HasRemoteAs

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasRemoteAs() bool`

HasRemoteAs returns a boolean if a field has been set.

### GetWeight

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetKeepAlive

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetKeepAlive() int64`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetKeepAliveOk() (*int64, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetKeepAlive(v int64)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetHoldDown

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetHoldDown() int64`

GetHoldDown returns the HoldDown field if non-nil, zero value otherwise.

### GetHoldDownOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetHoldDownOk() (*int64, bool)`

GetHoldDownOk returns a tuple with the HoldDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldDown

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetHoldDown(v int64)`

SetHoldDown sets HoldDown field to given value.

### HasHoldDown

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasHoldDown() bool`

HasHoldDown returns a boolean if a field has been set.

### GetPassword

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetRouteFilteringType

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRouteFilteringType() string`

GetRouteFilteringType returns the RouteFilteringType field if non-nil, zero value otherwise.

### GetRouteFilteringTypeOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRouteFilteringTypeOk() (*string, bool)`

GetRouteFilteringTypeOk returns a tuple with the RouteFilteringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringType

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetRouteFilteringType(v string)`

SetRouteFilteringType sets RouteFilteringType field to given value.

### HasRouteFilteringType

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasRouteFilteringType() bool`

HasRouteFilteringType returns a boolean if a field has been set.

### GetRouteFilteringIn

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRouteFilteringIn() string`

GetRouteFilteringIn returns the RouteFilteringIn field if non-nil, zero value otherwise.

### GetRouteFilteringInOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRouteFilteringInOk() (*string, bool)`

GetRouteFilteringInOk returns a tuple with the RouteFilteringIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringIn

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetRouteFilteringIn(v string)`

SetRouteFilteringIn sets RouteFilteringIn field to given value.

### HasRouteFilteringIn

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasRouteFilteringIn() bool`

HasRouteFilteringIn returns a boolean if a field has been set.

### GetRouteFilteringOut

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRouteFilteringOut() string`

GetRouteFilteringOut returns the RouteFilteringOut field if non-nil, zero value otherwise.

### GetRouteFilteringOutOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRouteFilteringOutOk() (*string, bool)`

GetRouteFilteringOutOk returns a tuple with the RouteFilteringOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilteringOut

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetRouteFilteringOut(v string)`

SetRouteFilteringOut sets RouteFilteringOut field to given value.

### HasRouteFilteringOut

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasRouteFilteringOut() bool`

HasRouteFilteringOut returns a boolean if a field has been set.

### GetBfdEnabled

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetBfdEnabled() bool`

GetBfdEnabled returns the BfdEnabled field if non-nil, zero value otherwise.

### GetBfdEnabledOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetBfdEnabledOk() (*bool, bool)`

GetBfdEnabledOk returns a tuple with the BfdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdEnabled

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetBfdEnabled(v bool)`

SetBfdEnabled sets BfdEnabled field to given value.

### HasBfdEnabled

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasBfdEnabled() bool`

HasBfdEnabled returns a boolean if a field has been set.

### GetBfdInterval

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetBfdInterval() int64`

GetBfdInterval returns the BfdInterval field if non-nil, zero value otherwise.

### GetBfdIntervalOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetBfdIntervalOk() (*int64, bool)`

GetBfdIntervalOk returns a tuple with the BfdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdInterval

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetBfdInterval(v int64)`

SetBfdInterval sets BfdInterval field to given value.

### HasBfdInterval

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasBfdInterval() bool`

HasBfdInterval returns a boolean if a field has been set.

### GetBfdMultiple

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetBfdMultiple() int64`

GetBfdMultiple returns the BfdMultiple field if non-nil, zero value otherwise.

### GetBfdMultipleOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetBfdMultipleOk() (*int64, bool)`

GetBfdMultipleOk returns a tuple with the BfdMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdMultiple

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetBfdMultiple(v int64)`

SetBfdMultiple sets BfdMultiple field to given value.

### HasBfdMultiple

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasBfdMultiple() bool`

HasBfdMultiple returns a boolean if a field has been set.

### GetAllowAsIn

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetAllowAsIn() bool`

GetAllowAsIn returns the AllowAsIn field if non-nil, zero value otherwise.

### GetAllowAsInOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetAllowAsInOk() (*bool, bool)`

GetAllowAsInOk returns a tuple with the AllowAsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAsIn

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetAllowAsIn(v bool)`

SetAllowAsIn sets AllowAsIn field to given value.

### HasAllowAsIn

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasAllowAsIn() bool`

HasAllowAsIn returns a boolean if a field has been set.

### GetHopLimit

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetHopLimit() int64`

GetHopLimit returns the HopLimit field if non-nil, zero value otherwise.

### GetHopLimitOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetHopLimitOk() (*int64, bool)`

GetHopLimitOk returns a tuple with the HopLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHopLimit

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetHopLimit(v int64)`

SetHopLimit sets HopLimit field to given value.

### HasHopLimit

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasHopLimit() bool`

HasHopLimit returns a boolean if a field has been set.

### GetRestartMode

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRestartMode() string`

GetRestartMode returns the RestartMode field if non-nil, zero value otherwise.

### GetRestartModeOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRestartModeOk() (*string, bool)`

GetRestartModeOk returns a tuple with the RestartMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartMode

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetRestartMode(v string)`

SetRestartMode sets RestartMode field to given value.

### HasRestartMode

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasRestartMode() bool`

HasRestartMode returns a boolean if a field has been set.

### GetProviderId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetSyncSource

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetSyncSource() string`

GetSyncSource returns the SyncSource field if non-nil, zero value otherwise.

### GetSyncSourceOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetSyncSourceOk() (*string, bool)`

GetSyncSourceOk returns a tuple with the SyncSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSource

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetSyncSource(v string)`

SetSyncSource sets SyncSource field to given value.

### HasSyncSource

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasSyncSource() bool`

HasSyncSource returns a boolean if a field has been set.

### GetInternalId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRefType

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetConfig

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetConfig() GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighborConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetConfigOk() (*GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighborConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetConfig(v GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighborConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkRouterBgpNeighbor200ResponseNetworkRouterBgpNeighbor) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



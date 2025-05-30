# UpdateNetworkRouterRequestNetworkRouter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Type** | Pointer to [**UpdateNetworkRouterRequestNetworkRouterType**](UpdateNetworkRouterRequestNetworkRouterType.md) |  | [optional] 
**Site** | Pointer to [**UpdateNetworkRouterRequestNetworkRouterSite**](UpdateNetworkRouterRequestNetworkRouterSite.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the network router (true, false). Default is on | [optional] 
**Zone** | Pointer to [**UpdateNetworkRouterRequestNetworkRouterZone**](UpdateNetworkRouterRequestNetworkRouterZone.md) |  | [optional] 
**NetworkServer** | Pointer to [**UpdateNetworkRouterRequestNetworkRouterNetworkServer**](UpdateNetworkRouterRequestNetworkRouterNetworkServer.md) |  | [optional] 

## Methods

### NewUpdateNetworkRouterRequestNetworkRouter

`func NewUpdateNetworkRouterRequestNetworkRouter() *UpdateNetworkRouterRequestNetworkRouter`

NewUpdateNetworkRouterRequestNetworkRouter instantiates a new UpdateNetworkRouterRequestNetworkRouter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkRouterRequestNetworkRouterWithDefaults

`func NewUpdateNetworkRouterRequestNetworkRouterWithDefaults() *UpdateNetworkRouterRequestNetworkRouter`

NewUpdateNetworkRouterRequestNetworkRouterWithDefaults instantiates a new UpdateNetworkRouterRequestNetworkRouter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkRouterRequestNetworkRouter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkRouterRequestNetworkRouter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkRouterRequestNetworkRouter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkRouterRequestNetworkRouter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *UpdateNetworkRouterRequestNetworkRouter) GetType() UpdateNetworkRouterRequestNetworkRouterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateNetworkRouterRequestNetworkRouter) GetTypeOk() (*UpdateNetworkRouterRequestNetworkRouterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateNetworkRouterRequestNetworkRouter) SetType(v UpdateNetworkRouterRequestNetworkRouterType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateNetworkRouterRequestNetworkRouter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSite

`func (o *UpdateNetworkRouterRequestNetworkRouter) GetSite() UpdateNetworkRouterRequestNetworkRouterSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *UpdateNetworkRouterRequestNetworkRouter) GetSiteOk() (*UpdateNetworkRouterRequestNetworkRouterSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *UpdateNetworkRouterRequestNetworkRouter) SetSite(v UpdateNetworkRouterRequestNetworkRouterSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *UpdateNetworkRouterRequestNetworkRouter) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateNetworkRouterRequestNetworkRouter) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkRouterRequestNetworkRouter) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkRouterRequestNetworkRouter) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkRouterRequestNetworkRouter) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetZone

`func (o *UpdateNetworkRouterRequestNetworkRouter) GetZone() UpdateNetworkRouterRequestNetworkRouterZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateNetworkRouterRequestNetworkRouter) GetZoneOk() (*UpdateNetworkRouterRequestNetworkRouterZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateNetworkRouterRequestNetworkRouter) SetZone(v UpdateNetworkRouterRequestNetworkRouterZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateNetworkRouterRequestNetworkRouter) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetNetworkServer

`func (o *UpdateNetworkRouterRequestNetworkRouter) GetNetworkServer() UpdateNetworkRouterRequestNetworkRouterNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *UpdateNetworkRouterRequestNetworkRouter) GetNetworkServerOk() (*UpdateNetworkRouterRequestNetworkRouterNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *UpdateNetworkRouterRequestNetworkRouter) SetNetworkServer(v UpdateNetworkRouterRequestNetworkRouterNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *UpdateNetworkRouterRequestNetworkRouter) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



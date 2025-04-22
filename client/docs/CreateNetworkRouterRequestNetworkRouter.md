# CreateNetworkRouterRequestNetworkRouter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**Type** | [**CreateNetworkRouterRequestNetworkRouterType**](CreateNetworkRouterRequestNetworkRouterType.md) |  | 
**Site** | [**CreateNetworkRouterRequestNetworkRouterSite**](CreateNetworkRouterRequestNetworkRouterSite.md) |  | 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the network router (true, false). Default is on | [optional] 
**Zone** | Pointer to [**CreateNetworkRouterRequestNetworkRouterZone**](CreateNetworkRouterRequestNetworkRouterZone.md) |  | [optional] 
**NetworkServer** | Pointer to [**CreateNetworkRouterRequestNetworkRouterNetworkServer**](CreateNetworkRouterRequestNetworkRouterNetworkServer.md) |  | [optional] 

## Methods

### NewCreateNetworkRouterRequestNetworkRouter

`func NewCreateNetworkRouterRequestNetworkRouter(name string, type_ CreateNetworkRouterRequestNetworkRouterType, site CreateNetworkRouterRequestNetworkRouterSite, ) *CreateNetworkRouterRequestNetworkRouter`

NewCreateNetworkRouterRequestNetworkRouter instantiates a new CreateNetworkRouterRequestNetworkRouter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkRouterRequestNetworkRouterWithDefaults

`func NewCreateNetworkRouterRequestNetworkRouterWithDefaults() *CreateNetworkRouterRequestNetworkRouter`

NewCreateNetworkRouterRequestNetworkRouterWithDefaults instantiates a new CreateNetworkRouterRequestNetworkRouter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkRouterRequestNetworkRouter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkRouterRequestNetworkRouter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkRouterRequestNetworkRouter) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateNetworkRouterRequestNetworkRouter) GetType() CreateNetworkRouterRequestNetworkRouterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNetworkRouterRequestNetworkRouter) GetTypeOk() (*CreateNetworkRouterRequestNetworkRouterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNetworkRouterRequestNetworkRouter) SetType(v CreateNetworkRouterRequestNetworkRouterType)`

SetType sets Type field to given value.


### GetSite

`func (o *CreateNetworkRouterRequestNetworkRouter) GetSite() CreateNetworkRouterRequestNetworkRouterSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CreateNetworkRouterRequestNetworkRouter) GetSiteOk() (*CreateNetworkRouterRequestNetworkRouterSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CreateNetworkRouterRequestNetworkRouter) SetSite(v CreateNetworkRouterRequestNetworkRouterSite)`

SetSite sets Site field to given value.


### GetEnabled

`func (o *CreateNetworkRouterRequestNetworkRouter) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateNetworkRouterRequestNetworkRouter) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateNetworkRouterRequestNetworkRouter) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateNetworkRouterRequestNetworkRouter) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetZone

`func (o *CreateNetworkRouterRequestNetworkRouter) GetZone() CreateNetworkRouterRequestNetworkRouterZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CreateNetworkRouterRequestNetworkRouter) GetZoneOk() (*CreateNetworkRouterRequestNetworkRouterZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CreateNetworkRouterRequestNetworkRouter) SetZone(v CreateNetworkRouterRequestNetworkRouterZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CreateNetworkRouterRequestNetworkRouter) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetNetworkServer

`func (o *CreateNetworkRouterRequestNetworkRouter) GetNetworkServer() CreateNetworkRouterRequestNetworkRouterNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *CreateNetworkRouterRequestNetworkRouter) GetNetworkServerOk() (*CreateNetworkRouterRequestNetworkRouterNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *CreateNetworkRouterRequestNetworkRouter) SetNetworkServer(v CreateNetworkRouterRequestNetworkRouterNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *CreateNetworkRouterRequestNetworkRouter) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateNetworkRouterRouteRequestNetworkRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Route name | [optional] 
**Description** | Pointer to **string** | Route description | [optional] 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the route (true, false). Default is off | [optional] [default to false]
**DefaultRoute** | Pointer to **bool** | Can be used to set as default route (true, false). Default is off | [optional] [default to false]
**Source** | **string** | Source address or range | 
**Destination** | **string** | Destination address or range | 
**NetworkMtu** | **string** | MTU | 

## Methods

### NewCreateNetworkRouterRouteRequestNetworkRoute

`func NewCreateNetworkRouterRouteRequestNetworkRoute(source string, destination string, networkMtu string, ) *CreateNetworkRouterRouteRequestNetworkRoute`

NewCreateNetworkRouterRouteRequestNetworkRoute instantiates a new CreateNetworkRouterRouteRequestNetworkRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkRouterRouteRequestNetworkRouteWithDefaults

`func NewCreateNetworkRouterRouteRequestNetworkRouteWithDefaults() *CreateNetworkRouterRouteRequestNetworkRoute`

NewCreateNetworkRouterRouteRequestNetworkRouteWithDefaults instantiates a new CreateNetworkRouterRouteRequestNetworkRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDefaultRoute

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetDefaultRoute() bool`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetDefaultRouteOk() (*bool, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) SetDefaultRoute(v bool)`

SetDefaultRoute sets DefaultRoute field to given value.

### HasDefaultRoute

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) HasDefaultRoute() bool`

HasDefaultRoute returns a boolean if a field has been set.

### GetSource

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) SetSource(v string)`

SetSource sets Source field to given value.


### GetDestination

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetNetworkMtu

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetNetworkMtu() string`

GetNetworkMtu returns the NetworkMtu field if non-nil, zero value otherwise.

### GetNetworkMtuOk

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) GetNetworkMtuOk() (*string, bool)`

GetNetworkMtuOk returns a tuple with the NetworkMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMtu

`func (o *CreateNetworkRouterRouteRequestNetworkRoute) SetNetworkMtu(v string)`

SetNetworkMtu sets NetworkMtu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



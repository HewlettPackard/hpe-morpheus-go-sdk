# NetworkRouterRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**RouteType** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**DefaultRoute** | Pointer to **bool** |  | [optional] 
**NetworkMtu** | Pointer to **string** |  | [optional] 
**ExternalInterface** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**ExternalType** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 

## Methods

### NewNetworkRouterRoute

`func NewNetworkRouterRoute() *NetworkRouterRoute`

NewNetworkRouterRoute instantiates a new NetworkRouterRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRouterRouteWithDefaults

`func NewNetworkRouterRouteWithDefaults() *NetworkRouterRoute`

NewNetworkRouterRouteWithDefaults instantiates a new NetworkRouterRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkRouterRoute) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkRouterRoute) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkRouterRoute) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkRouterRoute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkRouterRoute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkRouterRoute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkRouterRoute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkRouterRoute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *NetworkRouterRoute) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NetworkRouterRoute) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NetworkRouterRoute) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *NetworkRouterRoute) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkRouterRoute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkRouterRoute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkRouterRoute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkRouterRoute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriority

`func (o *NetworkRouterRoute) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NetworkRouterRoute) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NetworkRouterRoute) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NetworkRouterRoute) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRouteType

`func (o *NetworkRouterRoute) GetRouteType() string`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *NetworkRouterRoute) GetRouteTypeOk() (*string, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *NetworkRouterRoute) SetRouteType(v string)`

SetRouteType sets RouteType field to given value.

### HasRouteType

`func (o *NetworkRouterRoute) HasRouteType() bool`

HasRouteType returns a boolean if a field has been set.

### GetSource

`func (o *NetworkRouterRoute) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NetworkRouterRoute) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NetworkRouterRoute) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *NetworkRouterRoute) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceType

`func (o *NetworkRouterRoute) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *NetworkRouterRoute) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *NetworkRouterRoute) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *NetworkRouterRoute) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestination

`func (o *NetworkRouterRoute) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *NetworkRouterRoute) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *NetworkRouterRoute) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *NetworkRouterRoute) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationType

`func (o *NetworkRouterRoute) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *NetworkRouterRoute) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *NetworkRouterRoute) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *NetworkRouterRoute) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDefaultRoute

`func (o *NetworkRouterRoute) GetDefaultRoute() bool`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *NetworkRouterRoute) GetDefaultRouteOk() (*bool, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *NetworkRouterRoute) SetDefaultRoute(v bool)`

SetDefaultRoute sets DefaultRoute field to given value.

### HasDefaultRoute

`func (o *NetworkRouterRoute) HasDefaultRoute() bool`

HasDefaultRoute returns a boolean if a field has been set.

### GetNetworkMtu

`func (o *NetworkRouterRoute) GetNetworkMtu() string`

GetNetworkMtu returns the NetworkMtu field if non-nil, zero value otherwise.

### GetNetworkMtuOk

`func (o *NetworkRouterRoute) GetNetworkMtuOk() (*string, bool)`

GetNetworkMtuOk returns a tuple with the NetworkMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMtu

`func (o *NetworkRouterRoute) SetNetworkMtu(v string)`

SetNetworkMtu sets NetworkMtu field to given value.

### HasNetworkMtu

`func (o *NetworkRouterRoute) HasNetworkMtu() bool`

HasNetworkMtu returns a boolean if a field has been set.

### GetExternalInterface

`func (o *NetworkRouterRoute) GetExternalInterface() string`

GetExternalInterface returns the ExternalInterface field if non-nil, zero value otherwise.

### GetExternalInterfaceOk

`func (o *NetworkRouterRoute) GetExternalInterfaceOk() (*string, bool)`

GetExternalInterfaceOk returns a tuple with the ExternalInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInterface

`func (o *NetworkRouterRoute) SetExternalInterface(v string)`

SetExternalInterface sets ExternalInterface field to given value.

### HasExternalInterface

`func (o *NetworkRouterRoute) HasExternalInterface() bool`

HasExternalInterface returns a boolean if a field has been set.

### GetInternalId

`func (o *NetworkRouterRoute) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *NetworkRouterRoute) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *NetworkRouterRoute) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *NetworkRouterRoute) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *NetworkRouterRoute) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NetworkRouterRoute) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NetworkRouterRoute) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NetworkRouterRoute) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUniqueId

`func (o *NetworkRouterRoute) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *NetworkRouterRoute) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *NetworkRouterRoute) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *NetworkRouterRoute) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetProviderId

`func (o *NetworkRouterRoute) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *NetworkRouterRoute) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *NetworkRouterRoute) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *NetworkRouterRoute) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetExternalType

`func (o *NetworkRouterRoute) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *NetworkRouterRoute) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *NetworkRouterRoute) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *NetworkRouterRoute) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworkRouterRoute) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkRouterRoute) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkRouterRoute) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkRouterRoute) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVisible

`func (o *NetworkRouterRoute) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *NetworkRouterRoute) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *NetworkRouterRoute) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *NetworkRouterRoute) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



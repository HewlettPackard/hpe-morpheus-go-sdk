# GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**IacId** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**ZonePool** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 
**GroupLayer** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner**](GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner.md) |  | [optional] 

## Methods

### NewGetNetworkRouterFirewallRuleGroup200ResponseRuleGroup

`func NewGetNetworkRouterFirewallRuleGroup200ResponseRuleGroup() *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup`

NewGetNetworkRouterFirewallRuleGroup200ResponseRuleGroup instantiates a new GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRouterFirewallRuleGroup200ResponseRuleGroupWithDefaults

`func NewGetNetworkRouterFirewallRuleGroup200ResponseRuleGroupWithDefaults() *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup`

NewGetNetworkRouterFirewallRuleGroup200ResponseRuleGroupWithDefaults instantiates a new GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIacId

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetIacId() string`

GetIacId returns the IacId field if non-nil, zero value otherwise.

### GetIacIdOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetIacIdOk() (*string, bool)`

GetIacIdOk returns a tuple with the IacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacId

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetIacId(v string)`

SetIacId sets IacId field to given value.

### HasIacId

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) HasIacId() bool`

HasIacId returns a boolean if a field has been set.

### SetIacIdNil

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetIacIdNil(b bool)`

 SetIacIdNil sets the value for IacId to be an explicit nil

### UnsetIacId
`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) UnsetIacId()`

UnsetIacId ensures that no value is present for IacId, not even an explicit nil
### GetZone

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetZonePool

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetZonePool() string`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetZonePoolOk() (*string, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetZonePool(v string)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### SetZonePoolNil

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetZonePoolNil(b bool)`

 SetZonePoolNil sets the value for ZonePool to be an explicit nil

### UnsetZonePool
`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) UnsetZonePool()`

UnsetZonePool ensures that no value is present for ZonePool, not even an explicit nil
### GetStatus

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetGroupLayer

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetGroupLayer() string`

GetGroupLayer returns the GroupLayer field if non-nil, zero value otherwise.

### GetGroupLayerOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetGroupLayerOk() (*string, bool)`

GetGroupLayerOk returns a tuple with the GroupLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLayer

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetGroupLayer(v string)`

SetGroupLayer sets GroupLayer field to given value.

### HasGroupLayer

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) HasGroupLayer() bool`

HasGroupLayer returns a boolean if a field has been set.

### GetRules

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetRules() []GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) GetRulesOk() (*[]GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) SetRules(v []GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



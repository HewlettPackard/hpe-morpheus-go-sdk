# GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 
**GroupLayer** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner**](GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner.md) |  | [optional] 

## Methods

### NewGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner

`func NewGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner() *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner`

NewGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner instantiates a new GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerWithDefaults

`func NewGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerWithDefaults() *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner`

NewGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerWithDefaults instantiates a new GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriority

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetGroupLayer

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) GetGroupLayer() string`

GetGroupLayer returns the GroupLayer field if non-nil, zero value otherwise.

### GetGroupLayerOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) GetGroupLayerOk() (*string, bool)`

GetGroupLayerOk returns a tuple with the GroupLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLayer

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) SetGroupLayer(v string)`

SetGroupLayer sets GroupLayer field to given value.

### HasGroupLayer

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) HasGroupLayer() bool`

HasGroupLayer returns a boolean if a field has been set.

### GetRules

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) GetRules() []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) GetRulesOk() (*[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) SetRules(v []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



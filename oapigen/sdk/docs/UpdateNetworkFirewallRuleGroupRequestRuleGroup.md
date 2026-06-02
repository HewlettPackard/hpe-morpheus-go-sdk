# UpdateNetworkFirewallRuleGroupRequestRuleGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Network firewall rule group name | [optional] 
**Description** | Pointer to **NullableString** | Network firewall rule group description | [optional] 
**Priority** | Pointer to **NullableInt64** | Network firewall rule group priority | [optional] 

## Methods

### NewUpdateNetworkFirewallRuleGroupRequestRuleGroup

`func NewUpdateNetworkFirewallRuleGroupRequestRuleGroup() *UpdateNetworkFirewallRuleGroupRequestRuleGroup`

NewUpdateNetworkFirewallRuleGroupRequestRuleGroup instantiates a new UpdateNetworkFirewallRuleGroupRequestRuleGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriority

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *UpdateNetworkFirewallRuleGroupRequestRuleGroup) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



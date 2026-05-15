# UpdateNetworkFirewallRuleRequestRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleGroup** | Pointer to [**UpdateNetworkFirewallRuleRequestRuleRuleGroup**](UpdateNetworkFirewallRuleRequestRuleRuleGroup.md) |  | [optional] 
**Name** | Pointer to **string** | Network firewall rule name | [optional] 
**Description** | Pointer to **NullableString** | Network firewall rule description | [optional] 
**Enabled** | Pointer to **bool** | Use this to set enabled state | [optional] 
**Priority** | Pointer to **NullableInt64** | Network firewall rule priority | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**UpdateNetworkFirewallRuleRequestRuleSources**](UpdateNetworkFirewallRuleRequestRuleSources.md) |  | [optional] 
**Destinations** | Pointer to [**UpdateNetworkFirewallRuleRequestRuleDestinations**](UpdateNetworkFirewallRuleRequestRuleDestinations.md) |  | [optional] 
**Config** | Pointer to [**UpdateNetworkFirewallRuleRequestRuleConfig**](UpdateNetworkFirewallRuleRequestRuleConfig.md) |  | [optional] 
**Scopes** | Pointer to [**UpdateNetworkFirewallRuleRequestRuleScopes**](UpdateNetworkFirewallRuleRequestRuleScopes.md) |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateNetworkFirewallRuleRequestRule

`func NewUpdateNetworkFirewallRuleRequestRule() *UpdateNetworkFirewallRuleRequestRule`

NewUpdateNetworkFirewallRuleRequestRule instantiates a new UpdateNetworkFirewallRuleRequestRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkFirewallRuleRequestRuleWithDefaults

`func NewUpdateNetworkFirewallRuleRequestRuleWithDefaults() *UpdateNetworkFirewallRuleRequestRule`

NewUpdateNetworkFirewallRuleRequestRuleWithDefaults instantiates a new UpdateNetworkFirewallRuleRequestRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleGroup

`func (o *UpdateNetworkFirewallRuleRequestRule) GetRuleGroup() UpdateNetworkFirewallRuleRequestRuleRuleGroup`

GetRuleGroup returns the RuleGroup field if non-nil, zero value otherwise.

### GetRuleGroupOk

`func (o *UpdateNetworkFirewallRuleRequestRule) GetRuleGroupOk() (*UpdateNetworkFirewallRuleRequestRuleRuleGroup, bool)`

GetRuleGroupOk returns a tuple with the RuleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroup

`func (o *UpdateNetworkFirewallRuleRequestRule) SetRuleGroup(v UpdateNetworkFirewallRuleRequestRuleRuleGroup)`

SetRuleGroup sets RuleGroup field to given value.

### HasRuleGroup

`func (o *UpdateNetworkFirewallRuleRequestRule) HasRuleGroup() bool`

HasRuleGroup returns a boolean if a field has been set.

### GetName

`func (o *UpdateNetworkFirewallRuleRequestRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkFirewallRuleRequestRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkFirewallRuleRequestRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkFirewallRuleRequestRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateNetworkFirewallRuleRequestRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkFirewallRuleRequestRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkFirewallRuleRequestRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkFirewallRuleRequestRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateNetworkFirewallRuleRequestRule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateNetworkFirewallRuleRequestRule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *UpdateNetworkFirewallRuleRequestRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkFirewallRuleRequestRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkFirewallRuleRequestRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkFirewallRuleRequestRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *UpdateNetworkFirewallRuleRequestRule) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateNetworkFirewallRuleRequestRule) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateNetworkFirewallRuleRequestRule) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdateNetworkFirewallRuleRequestRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *UpdateNetworkFirewallRuleRequestRule) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *UpdateNetworkFirewallRuleRequestRule) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetDirection

`func (o *UpdateNetworkFirewallRuleRequestRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *UpdateNetworkFirewallRuleRequestRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *UpdateNetworkFirewallRuleRequestRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *UpdateNetworkFirewallRuleRequestRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSources

`func (o *UpdateNetworkFirewallRuleRequestRule) GetSources() UpdateNetworkFirewallRuleRequestRuleSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *UpdateNetworkFirewallRuleRequestRule) GetSourcesOk() (*UpdateNetworkFirewallRuleRequestRuleSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *UpdateNetworkFirewallRuleRequestRule) SetSources(v UpdateNetworkFirewallRuleRequestRuleSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *UpdateNetworkFirewallRuleRequestRule) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *UpdateNetworkFirewallRuleRequestRule) GetDestinations() UpdateNetworkFirewallRuleRequestRuleDestinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *UpdateNetworkFirewallRuleRequestRule) GetDestinationsOk() (*UpdateNetworkFirewallRuleRequestRuleDestinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *UpdateNetworkFirewallRuleRequestRule) SetDestinations(v UpdateNetworkFirewallRuleRequestRuleDestinations)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *UpdateNetworkFirewallRuleRequestRule) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateNetworkFirewallRuleRequestRule) GetConfig() UpdateNetworkFirewallRuleRequestRuleConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateNetworkFirewallRuleRequestRule) GetConfigOk() (*UpdateNetworkFirewallRuleRequestRuleConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateNetworkFirewallRuleRequestRule) SetConfig(v UpdateNetworkFirewallRuleRequestRuleConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateNetworkFirewallRuleRequestRule) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetScopes

`func (o *UpdateNetworkFirewallRuleRequestRule) GetScopes() UpdateNetworkFirewallRuleRequestRuleScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UpdateNetworkFirewallRuleRequestRule) GetScopesOk() (*UpdateNetworkFirewallRuleRequestRuleScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UpdateNetworkFirewallRuleRequestRule) SetScopes(v UpdateNetworkFirewallRuleRequestRuleScopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UpdateNetworkFirewallRuleRequestRule) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPolicy

`func (o *UpdateNetworkFirewallRuleRequestRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateNetworkFirewallRuleRequestRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateNetworkFirewallRuleRequestRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UpdateNetworkFirewallRuleRequestRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



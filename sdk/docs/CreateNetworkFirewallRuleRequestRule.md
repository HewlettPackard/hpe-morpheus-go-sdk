# CreateNetworkFirewallRuleRequestRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleGroup** | Pointer to [**CreateNetworkFirewallRuleRequestRuleRuleGroup**](CreateNetworkFirewallRuleRequestRuleRuleGroup.md) |  | [optional] 
**Name** | **string** | Network firewall rule name | 
**Description** | Pointer to **string** | Network firewall rule description | [optional] 
**Enabled** | Pointer to **bool** | Use this to set enabled state | [optional] 
**Priority** | Pointer to **string** | Network firewall rule priority | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**CreateNetworkFirewallRuleRequestRuleSources**](CreateNetworkFirewallRuleRequestRuleSources.md) |  | [optional] 
**Destinations** | Pointer to [**CreateNetworkFirewallRuleRequestRuleSources**](CreateNetworkFirewallRuleRequestRuleSources.md) |  | [optional] 
**Config** | Pointer to [**CreateNetworkFirewallRuleRequestRuleConfig**](CreateNetworkFirewallRuleRequestRuleConfig.md) |  | [optional] 
**Scopes** | Pointer to [**CreateNetworkFirewallRuleRequestRuleSources**](CreateNetworkFirewallRuleRequestRuleSources.md) |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateNetworkFirewallRuleRequestRule

`func NewCreateNetworkFirewallRuleRequestRule(name string, ) *CreateNetworkFirewallRuleRequestRule`

NewCreateNetworkFirewallRuleRequestRule instantiates a new CreateNetworkFirewallRuleRequestRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFirewallRuleRequestRuleWithDefaults

`func NewCreateNetworkFirewallRuleRequestRuleWithDefaults() *CreateNetworkFirewallRuleRequestRule`

NewCreateNetworkFirewallRuleRequestRuleWithDefaults instantiates a new CreateNetworkFirewallRuleRequestRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleGroup

`func (o *CreateNetworkFirewallRuleRequestRule) GetRuleGroup() CreateNetworkFirewallRuleRequestRuleRuleGroup`

GetRuleGroup returns the RuleGroup field if non-nil, zero value otherwise.

### GetRuleGroupOk

`func (o *CreateNetworkFirewallRuleRequestRule) GetRuleGroupOk() (*CreateNetworkFirewallRuleRequestRuleRuleGroup, bool)`

GetRuleGroupOk returns a tuple with the RuleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroup

`func (o *CreateNetworkFirewallRuleRequestRule) SetRuleGroup(v CreateNetworkFirewallRuleRequestRuleRuleGroup)`

SetRuleGroup sets RuleGroup field to given value.

### HasRuleGroup

`func (o *CreateNetworkFirewallRuleRequestRule) HasRuleGroup() bool`

HasRuleGroup returns a boolean if a field has been set.

### GetName

`func (o *CreateNetworkFirewallRuleRequestRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkFirewallRuleRequestRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkFirewallRuleRequestRule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateNetworkFirewallRuleRequestRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkFirewallRuleRequestRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkFirewallRuleRequestRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkFirewallRuleRequestRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateNetworkFirewallRuleRequestRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateNetworkFirewallRuleRequestRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateNetworkFirewallRuleRequestRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateNetworkFirewallRuleRequestRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *CreateNetworkFirewallRuleRequestRule) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateNetworkFirewallRuleRequestRule) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateNetworkFirewallRuleRequestRule) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateNetworkFirewallRuleRequestRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDirection

`func (o *CreateNetworkFirewallRuleRequestRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CreateNetworkFirewallRuleRequestRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CreateNetworkFirewallRuleRequestRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CreateNetworkFirewallRuleRequestRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSources

`func (o *CreateNetworkFirewallRuleRequestRule) GetSources() CreateNetworkFirewallRuleRequestRuleSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CreateNetworkFirewallRuleRequestRule) GetSourcesOk() (*CreateNetworkFirewallRuleRequestRuleSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CreateNetworkFirewallRuleRequestRule) SetSources(v CreateNetworkFirewallRuleRequestRuleSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CreateNetworkFirewallRuleRequestRule) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *CreateNetworkFirewallRuleRequestRule) GetDestinations() CreateNetworkFirewallRuleRequestRuleSources`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *CreateNetworkFirewallRuleRequestRule) GetDestinationsOk() (*CreateNetworkFirewallRuleRequestRuleSources, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *CreateNetworkFirewallRuleRequestRule) SetDestinations(v CreateNetworkFirewallRuleRequestRuleSources)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *CreateNetworkFirewallRuleRequestRule) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetConfig

`func (o *CreateNetworkFirewallRuleRequestRule) GetConfig() CreateNetworkFirewallRuleRequestRuleConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateNetworkFirewallRuleRequestRule) GetConfigOk() (*CreateNetworkFirewallRuleRequestRuleConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateNetworkFirewallRuleRequestRule) SetConfig(v CreateNetworkFirewallRuleRequestRuleConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateNetworkFirewallRuleRequestRule) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetScopes

`func (o *CreateNetworkFirewallRuleRequestRule) GetScopes() CreateNetworkFirewallRuleRequestRuleSources`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateNetworkFirewallRuleRequestRule) GetScopesOk() (*CreateNetworkFirewallRuleRequestRuleSources, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateNetworkFirewallRuleRequestRule) SetScopes(v CreateNetworkFirewallRuleRequestRuleSources)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CreateNetworkFirewallRuleRequestRule) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPolicy

`func (o *CreateNetworkFirewallRuleRequestRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CreateNetworkFirewallRuleRequestRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CreateNetworkFirewallRuleRequestRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CreateNetworkFirewallRuleRequestRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



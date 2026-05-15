# NetworkFirewallRuleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleGroup** | Pointer to [**NetworkFirewallRuleUpdateRuleGroup**](NetworkFirewallRuleUpdateRuleGroup.md) |  | [optional] 
**Name** | Pointer to **string** | Network firewall rule name | [optional] 
**Description** | Pointer to **NullableString** | Network firewall rule description | [optional] 
**Enabled** | Pointer to **bool** | Use this to set enabled state | [optional] 
**Priority** | Pointer to **NullableInt64** | Network firewall rule priority | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**NetworkFirewallRuleUpdateSources**](NetworkFirewallRuleUpdateSources.md) |  | [optional] 
**Destinations** | Pointer to [**NetworkFirewallRuleUpdateDestinations**](NetworkFirewallRuleUpdateDestinations.md) |  | [optional] 
**Config** | Pointer to [**NetworkFirewallRuleUpdateConfig**](NetworkFirewallRuleUpdateConfig.md) |  | [optional] 
**Scopes** | Pointer to [**NetworkFirewallRuleUpdateScopes**](NetworkFirewallRuleUpdateScopes.md) |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkFirewallRuleUpdate

`func NewNetworkFirewallRuleUpdate() *NetworkFirewallRuleUpdate`

NewNetworkFirewallRuleUpdate instantiates a new NetworkFirewallRuleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFirewallRuleUpdateWithDefaults

`func NewNetworkFirewallRuleUpdateWithDefaults() *NetworkFirewallRuleUpdate`

NewNetworkFirewallRuleUpdateWithDefaults instantiates a new NetworkFirewallRuleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleGroup

`func (o *NetworkFirewallRuleUpdate) GetRuleGroup() NetworkFirewallRuleUpdateRuleGroup`

GetRuleGroup returns the RuleGroup field if non-nil, zero value otherwise.

### GetRuleGroupOk

`func (o *NetworkFirewallRuleUpdate) GetRuleGroupOk() (*NetworkFirewallRuleUpdateRuleGroup, bool)`

GetRuleGroupOk returns a tuple with the RuleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroup

`func (o *NetworkFirewallRuleUpdate) SetRuleGroup(v NetworkFirewallRuleUpdateRuleGroup)`

SetRuleGroup sets RuleGroup field to given value.

### HasRuleGroup

`func (o *NetworkFirewallRuleUpdate) HasRuleGroup() bool`

HasRuleGroup returns a boolean if a field has been set.

### GetName

`func (o *NetworkFirewallRuleUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkFirewallRuleUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkFirewallRuleUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkFirewallRuleUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkFirewallRuleUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkFirewallRuleUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkFirewallRuleUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkFirewallRuleUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkFirewallRuleUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkFirewallRuleUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *NetworkFirewallRuleUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkFirewallRuleUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkFirewallRuleUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkFirewallRuleUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *NetworkFirewallRuleUpdate) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NetworkFirewallRuleUpdate) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NetworkFirewallRuleUpdate) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NetworkFirewallRuleUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *NetworkFirewallRuleUpdate) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *NetworkFirewallRuleUpdate) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetDirection

`func (o *NetworkFirewallRuleUpdate) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *NetworkFirewallRuleUpdate) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *NetworkFirewallRuleUpdate) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *NetworkFirewallRuleUpdate) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSources

`func (o *NetworkFirewallRuleUpdate) GetSources() NetworkFirewallRuleUpdateSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *NetworkFirewallRuleUpdate) GetSourcesOk() (*NetworkFirewallRuleUpdateSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *NetworkFirewallRuleUpdate) SetSources(v NetworkFirewallRuleUpdateSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *NetworkFirewallRuleUpdate) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *NetworkFirewallRuleUpdate) GetDestinations() NetworkFirewallRuleUpdateDestinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *NetworkFirewallRuleUpdate) GetDestinationsOk() (*NetworkFirewallRuleUpdateDestinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *NetworkFirewallRuleUpdate) SetDestinations(v NetworkFirewallRuleUpdateDestinations)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *NetworkFirewallRuleUpdate) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkFirewallRuleUpdate) GetConfig() NetworkFirewallRuleUpdateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkFirewallRuleUpdate) GetConfigOk() (*NetworkFirewallRuleUpdateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkFirewallRuleUpdate) SetConfig(v NetworkFirewallRuleUpdateConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkFirewallRuleUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetScopes

`func (o *NetworkFirewallRuleUpdate) GetScopes() NetworkFirewallRuleUpdateScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *NetworkFirewallRuleUpdate) GetScopesOk() (*NetworkFirewallRuleUpdateScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *NetworkFirewallRuleUpdate) SetScopes(v NetworkFirewallRuleUpdateScopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *NetworkFirewallRuleUpdate) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPolicy

`func (o *NetworkFirewallRuleUpdate) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *NetworkFirewallRuleUpdate) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *NetworkFirewallRuleUpdate) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *NetworkFirewallRuleUpdate) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



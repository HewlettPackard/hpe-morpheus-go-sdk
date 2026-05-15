# GetNetworkFirewallRules200ResponseAllOfRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**RuleGroup** | Pointer to [**NullableGetNetworkFirewallRules200ResponseAllOfRulesInnerRuleGroup**](GetNetworkFirewallRules200ResponseAllOfRulesInnerRuleGroup.md) |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Sources** | Pointer to [**[]GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**Destinations** | Pointer to [**[]GetNetworkFirewallRules200ResponseAllOfRulesInnerDestinationsInner**](GetNetworkFirewallRules200ResponseAllOfRulesInnerDestinationsInner.md) |  | [optional] 
**Applications** | Pointer to [**[]GetNetworkFirewallRules200ResponseAllOfRulesInnerApplicationsInner**](GetNetworkFirewallRules200ResponseAllOfRulesInnerApplicationsInner.md) |  | [optional] 
**Scopes** | Pointer to [**[]GetNetworkFirewallRules200ResponseAllOfRulesInnerScopesInner**](GetNetworkFirewallRules200ResponseAllOfRulesInnerScopesInner.md) |  | [optional] 
**Profiles** | Pointer to [**[]GetNetworkFirewallRules200ResponseAllOfRulesInnerProfilesInner**](GetNetworkFirewallRules200ResponseAllOfRulesInnerProfilesInner.md) |  | [optional] 
**AppliedTargets** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetNetworkFirewallRules200ResponseAllOfRulesInner

`func NewGetNetworkFirewallRules200ResponseAllOfRulesInner() *GetNetworkFirewallRules200ResponseAllOfRulesInner`

NewGetNetworkFirewallRules200ResponseAllOfRulesInner instantiates a new GetNetworkFirewallRules200ResponseAllOfRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFirewallRules200ResponseAllOfRulesInnerWithDefaults

`func NewGetNetworkFirewallRules200ResponseAllOfRulesInnerWithDefaults() *GetNetworkFirewallRules200ResponseAllOfRulesInner`

NewGetNetworkFirewallRules200ResponseAllOfRulesInnerWithDefaults instantiates a new GetNetworkFirewallRules200ResponseAllOfRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDirection

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSourceType

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestinationType

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicy

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPriority

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRuleGroup

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetRuleGroup() GetNetworkFirewallRules200ResponseAllOfRulesInnerRuleGroup`

GetRuleGroup returns the RuleGroup field if non-nil, zero value otherwise.

### GetRuleGroupOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetRuleGroupOk() (*GetNetworkFirewallRules200ResponseAllOfRulesInnerRuleGroup, bool)`

GetRuleGroupOk returns a tuple with the RuleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroup

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetRuleGroup(v GetNetworkFirewallRules200ResponseAllOfRulesInnerRuleGroup)`

SetRuleGroup sets RuleGroup field to given value.

### HasRuleGroup

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasRuleGroup() bool`

HasRuleGroup returns a boolean if a field has been set.

### SetRuleGroupNil

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetRuleGroupNil(b bool)`

 SetRuleGroupNil sets the value for RuleGroup to be an explicit nil

### UnsetRuleGroup
`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) UnsetRuleGroup()`

UnsetRuleGroup ensures that no value is present for RuleGroup, not even an explicit nil
### GetGroupName

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetConfig

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetSources

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetSources() []GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetSourcesOk() (*[]GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetSources(v []GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetDestinations() []GetNetworkFirewallRules200ResponseAllOfRulesInnerDestinationsInner`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetDestinationsOk() (*[]GetNetworkFirewallRules200ResponseAllOfRulesInnerDestinationsInner, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetDestinations(v []GetNetworkFirewallRules200ResponseAllOfRulesInnerDestinationsInner)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetApplications

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetApplications() []GetNetworkFirewallRules200ResponseAllOfRulesInnerApplicationsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetApplicationsOk() (*[]GetNetworkFirewallRules200ResponseAllOfRulesInnerApplicationsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetApplications(v []GetNetworkFirewallRules200ResponseAllOfRulesInnerApplicationsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetScopes

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetScopes() []GetNetworkFirewallRules200ResponseAllOfRulesInnerScopesInner`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetScopesOk() (*[]GetNetworkFirewallRules200ResponseAllOfRulesInnerScopesInner, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetScopes(v []GetNetworkFirewallRules200ResponseAllOfRulesInnerScopesInner)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetProfiles

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetProfiles() []GetNetworkFirewallRules200ResponseAllOfRulesInnerProfilesInner`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetProfilesOk() (*[]GetNetworkFirewallRules200ResponseAllOfRulesInnerProfilesInner, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetProfiles(v []GetNetworkFirewallRules200ResponseAllOfRulesInnerProfilesInner)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetAppliedTargets

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetAppliedTargets() []map[string]interface{}`

GetAppliedTargets returns the AppliedTargets field if non-nil, zero value otherwise.

### GetAppliedTargetsOk

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) GetAppliedTargetsOk() (*[]map[string]interface{}, bool)`

GetAppliedTargetsOk returns a tuple with the AppliedTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTargets

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) SetAppliedTargets(v []map[string]interface{})`

SetAppliedTargets sets AppliedTargets field to given value.

### HasAppliedTargets

`func (o *GetNetworkFirewallRules200ResponseAllOfRulesInner) HasAppliedTargets() bool`

HasAppliedTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



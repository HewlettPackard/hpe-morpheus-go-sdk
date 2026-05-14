# GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**RuleGroup** | Pointer to [**GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerRuleGroup**](GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerRuleGroup.md) |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Sources** | Pointer to [**[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerSourcesInner**](GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerSourcesInner.md) |  | [optional] 
**Destinations** | Pointer to [**[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerDestinationsInner**](GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerDestinationsInner.md) |  | [optional] 
**Applications** | Pointer to [**[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerApplicationsInner**](GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerApplicationsInner.md) |  | [optional] 
**Scopes** | Pointer to [**[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerScopesInner**](GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerScopesInner.md) |  | [optional] 
**Profiles** | Pointer to [**[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerProfilesInner**](GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerProfilesInner.md) |  | [optional] 
**AppliedTargets** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner

`func NewGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner() *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner`

NewGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner instantiates a new GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerWithDefaults

`func NewGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerWithDefaults() *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner`

NewGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerWithDefaults instantiates a new GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDirection

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSourceType

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestinationType

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicy

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPriority

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRuleGroup

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetRuleGroup() GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerRuleGroup`

GetRuleGroup returns the RuleGroup field if non-nil, zero value otherwise.

### GetRuleGroupOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetRuleGroupOk() (*GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerRuleGroup, bool)`

GetRuleGroupOk returns a tuple with the RuleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroup

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetRuleGroup(v GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerRuleGroup)`

SetRuleGroup sets RuleGroup field to given value.

### HasRuleGroup

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasRuleGroup() bool`

HasRuleGroup returns a boolean if a field has been set.

### GetGroupName

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetConfig

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetSources

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetSources() []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetSourcesOk() (*[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetSources(v []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerSourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetDestinations() []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerDestinationsInner`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetDestinationsOk() (*[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerDestinationsInner, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetDestinations(v []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerDestinationsInner)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetApplications

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetApplications() []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerApplicationsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetApplicationsOk() (*[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerApplicationsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetApplications(v []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerApplicationsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetScopes

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetScopes() []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerScopesInner`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetScopesOk() (*[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerScopesInner, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetScopes(v []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerScopesInner)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetProfiles

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetProfiles() []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerProfilesInner`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetProfilesOk() (*[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerProfilesInner, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetProfiles(v []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInnerProfilesInner)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetAppliedTargets

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetAppliedTargets() []map[string]interface{}`

GetAppliedTargets returns the AppliedTargets field if non-nil, zero value otherwise.

### GetAppliedTargetsOk

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) GetAppliedTargetsOk() (*[]map[string]interface{}, bool)`

GetAppliedTargetsOk returns a tuple with the AppliedTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTargets

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) SetAppliedTargets(v []map[string]interface{})`

SetAppliedTargets sets AppliedTargets field to given value.

### HasAppliedTargets

`func (o *GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner) HasAppliedTargets() bool`

HasAppliedTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



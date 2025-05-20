# GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner

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
**RuleGroup** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Sources** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**Destinations** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**Applications** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**Scopes** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**Profiles** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**AppliedTargets** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner

`func NewGetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner() *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner`

NewGetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner instantiates a new GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInnerWithDefaults

`func NewGetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInnerWithDefaults() *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner`

NewGetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInnerWithDefaults instantiates a new GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDirection

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSourceType

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestinationType

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicy

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPriority

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRuleGroup

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetRuleGroup() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetRuleGroup returns the RuleGroup field if non-nil, zero value otherwise.

### GetRuleGroupOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetRuleGroupOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetRuleGroupOk returns a tuple with the RuleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroup

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetRuleGroup(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetRuleGroup sets RuleGroup field to given value.

### HasRuleGroup

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasRuleGroup() bool`

HasRuleGroup returns a boolean if a field has been set.

### GetGroupName

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetConfig

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetSources

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSources() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSourcesOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetSources(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDestinations() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDestinationsOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetDestinations(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetApplications

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetApplications() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetApplicationsOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetApplications(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetScopes

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetScopes() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetScopesOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetScopes(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetProfiles

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetProfiles() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetProfilesOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetProfiles(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetAppliedTargets

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetAppliedTargets() []map[string]interface{}`

GetAppliedTargets returns the AppliedTargets field if non-nil, zero value otherwise.

### GetAppliedTargetsOk

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) GetAppliedTargetsOk() (*[]map[string]interface{}, bool)`

GetAppliedTargetsOk returns a tuple with the AppliedTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTargets

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) SetAppliedTargets(v []map[string]interface{})`

SetAppliedTargets sets AppliedTargets field to given value.

### HasAppliedTargets

`func (o *GetNetworkFirewallRuleGroup200ResponseRuleGroupRulesInner) HasAppliedTargets() bool`

HasAppliedTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetNetworkFirewallRule200ResponseRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**RuleGroup** | Pointer to [**SnapshotsInstance200ResponseSnapshotsInnerZone**](SnapshotsInstance200ResponseSnapshotsInnerZone.md) |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Sources** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**Destinations** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**Applications** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**Scopes** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**Profiles** | Pointer to [**[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner**](ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner.md) |  | [optional] 
**AppliedTargets** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetNetworkFirewallRule200ResponseRule

`func NewGetNetworkFirewallRule200ResponseRule() *GetNetworkFirewallRule200ResponseRule`

NewGetNetworkFirewallRule200ResponseRule instantiates a new GetNetworkFirewallRule200ResponseRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFirewallRule200ResponseRuleWithDefaults

`func NewGetNetworkFirewallRule200ResponseRuleWithDefaults() *GetNetworkFirewallRule200ResponseRule`

NewGetNetworkFirewallRule200ResponseRuleWithDefaults instantiates a new GetNetworkFirewallRule200ResponseRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkFirewallRule200ResponseRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkFirewallRule200ResponseRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkFirewallRule200ResponseRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDirection

`func (o *GetNetworkFirewallRule200ResponseRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetNetworkFirewallRule200ResponseRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetNetworkFirewallRule200ResponseRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSourceType

`func (o *GetNetworkFirewallRule200ResponseRule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GetNetworkFirewallRule200ResponseRule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GetNetworkFirewallRule200ResponseRule) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestinationType

`func (o *GetNetworkFirewallRule200ResponseRule) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *GetNetworkFirewallRule200ResponseRule) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *GetNetworkFirewallRule200ResponseRule) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkFirewallRule200ResponseRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkFirewallRule200ResponseRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkFirewallRule200ResponseRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicy

`func (o *GetNetworkFirewallRule200ResponseRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetNetworkFirewallRule200ResponseRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetNetworkFirewallRule200ResponseRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPriority

`func (o *GetNetworkFirewallRule200ResponseRule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetNetworkFirewallRule200ResponseRule) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetNetworkFirewallRule200ResponseRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkFirewallRule200ResponseRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkFirewallRule200ResponseRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkFirewallRule200ResponseRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRuleGroup

`func (o *GetNetworkFirewallRule200ResponseRule) GetRuleGroup() SnapshotsInstance200ResponseSnapshotsInnerZone`

GetRuleGroup returns the RuleGroup field if non-nil, zero value otherwise.

### GetRuleGroupOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetRuleGroupOk() (*SnapshotsInstance200ResponseSnapshotsInnerZone, bool)`

GetRuleGroupOk returns a tuple with the RuleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroup

`func (o *GetNetworkFirewallRule200ResponseRule) SetRuleGroup(v SnapshotsInstance200ResponseSnapshotsInnerZone)`

SetRuleGroup sets RuleGroup field to given value.

### HasRuleGroup

`func (o *GetNetworkFirewallRule200ResponseRule) HasRuleGroup() bool`

HasRuleGroup returns a boolean if a field has been set.

### GetGroupName

`func (o *GetNetworkFirewallRule200ResponseRule) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GetNetworkFirewallRule200ResponseRule) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GetNetworkFirewallRule200ResponseRule) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetConfig

`func (o *GetNetworkFirewallRule200ResponseRule) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkFirewallRule200ResponseRule) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkFirewallRule200ResponseRule) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetSources

`func (o *GetNetworkFirewallRule200ResponseRule) GetSources() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetSourcesOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GetNetworkFirewallRule200ResponseRule) SetSources(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GetNetworkFirewallRule200ResponseRule) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *GetNetworkFirewallRule200ResponseRule) GetDestinations() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetDestinationsOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *GetNetworkFirewallRule200ResponseRule) SetDestinations(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *GetNetworkFirewallRule200ResponseRule) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetApplications

`func (o *GetNetworkFirewallRule200ResponseRule) GetApplications() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetApplicationsOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *GetNetworkFirewallRule200ResponseRule) SetApplications(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *GetNetworkFirewallRule200ResponseRule) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetScopes

`func (o *GetNetworkFirewallRule200ResponseRule) GetScopes() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetScopesOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetNetworkFirewallRule200ResponseRule) SetScopes(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GetNetworkFirewallRule200ResponseRule) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetProfiles

`func (o *GetNetworkFirewallRule200ResponseRule) GetProfiles() []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetProfilesOk() (*[]ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *GetNetworkFirewallRule200ResponseRule) SetProfiles(v []ListInstanceServicePlans200ResponsePlansInnerAutoOptionsInner)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *GetNetworkFirewallRule200ResponseRule) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetAppliedTargets

`func (o *GetNetworkFirewallRule200ResponseRule) GetAppliedTargets() []map[string]interface{}`

GetAppliedTargets returns the AppliedTargets field if non-nil, zero value otherwise.

### GetAppliedTargetsOk

`func (o *GetNetworkFirewallRule200ResponseRule) GetAppliedTargetsOk() (*[]map[string]interface{}, bool)`

GetAppliedTargetsOk returns a tuple with the AppliedTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTargets

`func (o *GetNetworkFirewallRule200ResponseRule) SetAppliedTargets(v []map[string]interface{})`

SetAppliedTargets sets AppliedTargets field to given value.

### HasAppliedTargets

`func (o *GetNetworkFirewallRule200ResponseRule) HasAppliedTargets() bool`

HasAppliedTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



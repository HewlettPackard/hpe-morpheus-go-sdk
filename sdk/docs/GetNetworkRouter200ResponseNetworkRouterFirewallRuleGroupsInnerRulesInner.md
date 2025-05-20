# GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**RuleType** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **[]string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **[]string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Profiles** | Pointer to **[]string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Application** | Pointer to **string** |  | [optional] 
**ApplicationType** | Pointer to **string** |  | [optional] 
**PortRange** | Pointer to **string** |  | [optional] 
**SourcePortRange** | Pointer to **string** |  | [optional] 
**DestinationPortRange** | Pointer to **string** |  | [optional] 
**SourceGroup** | Pointer to **string** |  | [optional] 
**SourceTier** | Pointer to **string** |  | [optional] 
**Applications** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewGetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner

`func NewGetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner() *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner`

NewGetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner instantiates a new GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInnerWithDefaults

`func NewGetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInnerWithDefaults() *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner`

NewGetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInnerWithDefaults instantiates a new GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroupName

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDirection

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetRuleType

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetPolicy

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSource

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSource() []string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceOk() (*[]string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetSource(v []string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceType

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestination

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDestination() []string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDestinationOk() (*[]string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetDestination(v []string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationType

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetProfiles

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetProfiles() []string`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetProfilesOk() (*[]string, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetProfiles(v []string)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetProtocol

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetApplication

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetApplicationType

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetPortRange

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetSourcePortRange

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourcePortRange() string`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourcePortRangeOk() (*string, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetSourcePortRange(v string)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### GetDestinationPortRange

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDestinationPortRange() string`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDestinationPortRangeOk() (*string, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetDestinationPortRange(v string)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.

### GetSourceGroup

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceGroup() string`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceGroupOk() (*string, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetSourceGroup(v string)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetSourceTier

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceTier() string`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceTierOk() (*string, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetSourceTier(v string)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### GetApplications

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetApplications() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetApplicationsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetApplications(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



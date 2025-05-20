# GetNetworkRoutersFirewallRules200ResponseRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 
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

### NewGetNetworkRoutersFirewallRules200ResponseRulesInner

`func NewGetNetworkRoutersFirewallRules200ResponseRulesInner() *GetNetworkRoutersFirewallRules200ResponseRulesInner`

NewGetNetworkRoutersFirewallRules200ResponseRulesInner instantiates a new GetNetworkRoutersFirewallRules200ResponseRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRoutersFirewallRules200ResponseRulesInnerWithDefaults

`func NewGetNetworkRoutersFirewallRules200ResponseRulesInnerWithDefaults() *GetNetworkRoutersFirewallRules200ResponseRulesInner`

NewGetNetworkRoutersFirewallRules200ResponseRulesInnerWithDefaults instantiates a new GetNetworkRoutersFirewallRules200ResponseRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetGroupName

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDirection

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetRuleType

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetPolicy

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSource

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetSource() []string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetSourceOk() (*[]string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetSource(v []string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceType

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestination

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetDestination() []string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetDestinationOk() (*[]string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetDestination(v []string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationType

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetProfiles

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetProfiles() []string`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetProfilesOk() (*[]string, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetProfiles(v []string)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetProtocol

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetApplication

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetApplicationType

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetPortRange

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetSourcePortRange

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetSourcePortRange() string`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetSourcePortRangeOk() (*string, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetSourcePortRange(v string)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### GetDestinationPortRange

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetDestinationPortRange() string`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetDestinationPortRangeOk() (*string, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetDestinationPortRange(v string)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.

### GetSourceGroup

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetSourceGroup() string`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetSourceGroupOk() (*string, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetSourceGroup(v string)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetSourceTier

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetSourceTier() string`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetSourceTierOk() (*string, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetSourceTier(v string)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### GetApplications

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetApplications() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) GetApplicationsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) SetApplications(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *GetNetworkRoutersFirewallRules200ResponseRulesInner) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



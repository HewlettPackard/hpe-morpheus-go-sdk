# NetworkRouterFirewallRule

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

### NewNetworkRouterFirewallRule

`func NewNetworkRouterFirewallRule() *NetworkRouterFirewallRule`

NewNetworkRouterFirewallRule instantiates a new NetworkRouterFirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRouterFirewallRuleWithDefaults

`func NewNetworkRouterFirewallRuleWithDefaults() *NetworkRouterFirewallRule`

NewNetworkRouterFirewallRuleWithDefaults instantiates a new NetworkRouterFirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkRouterFirewallRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkRouterFirewallRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkRouterFirewallRule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkRouterFirewallRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkRouterFirewallRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkRouterFirewallRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkRouterFirewallRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkRouterFirewallRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *NetworkRouterFirewallRule) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NetworkRouterFirewallRule) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NetworkRouterFirewallRule) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *NetworkRouterFirewallRule) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworkRouterFirewallRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkRouterFirewallRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkRouterFirewallRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkRouterFirewallRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *NetworkRouterFirewallRule) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NetworkRouterFirewallRule) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NetworkRouterFirewallRule) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NetworkRouterFirewallRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetGroupName

`func (o *NetworkRouterFirewallRule) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *NetworkRouterFirewallRule) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *NetworkRouterFirewallRule) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *NetworkRouterFirewallRule) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDirection

`func (o *NetworkRouterFirewallRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *NetworkRouterFirewallRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *NetworkRouterFirewallRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *NetworkRouterFirewallRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetRuleType

`func (o *NetworkRouterFirewallRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *NetworkRouterFirewallRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *NetworkRouterFirewallRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *NetworkRouterFirewallRule) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetPolicy

`func (o *NetworkRouterFirewallRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *NetworkRouterFirewallRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *NetworkRouterFirewallRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *NetworkRouterFirewallRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSource

`func (o *NetworkRouterFirewallRule) GetSource() []string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NetworkRouterFirewallRule) GetSourceOk() (*[]string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NetworkRouterFirewallRule) SetSource(v []string)`

SetSource sets Source field to given value.

### HasSource

`func (o *NetworkRouterFirewallRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceType

`func (o *NetworkRouterFirewallRule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *NetworkRouterFirewallRule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *NetworkRouterFirewallRule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *NetworkRouterFirewallRule) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestination

`func (o *NetworkRouterFirewallRule) GetDestination() []string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *NetworkRouterFirewallRule) GetDestinationOk() (*[]string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *NetworkRouterFirewallRule) SetDestination(v []string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *NetworkRouterFirewallRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationType

`func (o *NetworkRouterFirewallRule) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *NetworkRouterFirewallRule) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *NetworkRouterFirewallRule) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *NetworkRouterFirewallRule) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetProfiles

`func (o *NetworkRouterFirewallRule) GetProfiles() []string`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *NetworkRouterFirewallRule) GetProfilesOk() (*[]string, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *NetworkRouterFirewallRule) SetProfiles(v []string)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *NetworkRouterFirewallRule) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetProtocol

`func (o *NetworkRouterFirewallRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetworkRouterFirewallRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetworkRouterFirewallRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NetworkRouterFirewallRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetApplication

`func (o *NetworkRouterFirewallRule) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *NetworkRouterFirewallRule) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *NetworkRouterFirewallRule) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *NetworkRouterFirewallRule) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetApplicationType

`func (o *NetworkRouterFirewallRule) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *NetworkRouterFirewallRule) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *NetworkRouterFirewallRule) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *NetworkRouterFirewallRule) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetPortRange

`func (o *NetworkRouterFirewallRule) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *NetworkRouterFirewallRule) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *NetworkRouterFirewallRule) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *NetworkRouterFirewallRule) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetSourcePortRange

`func (o *NetworkRouterFirewallRule) GetSourcePortRange() string`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *NetworkRouterFirewallRule) GetSourcePortRangeOk() (*string, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *NetworkRouterFirewallRule) SetSourcePortRange(v string)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *NetworkRouterFirewallRule) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### GetDestinationPortRange

`func (o *NetworkRouterFirewallRule) GetDestinationPortRange() string`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *NetworkRouterFirewallRule) GetDestinationPortRangeOk() (*string, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *NetworkRouterFirewallRule) SetDestinationPortRange(v string)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *NetworkRouterFirewallRule) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.

### GetSourceGroup

`func (o *NetworkRouterFirewallRule) GetSourceGroup() string`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *NetworkRouterFirewallRule) GetSourceGroupOk() (*string, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *NetworkRouterFirewallRule) SetSourceGroup(v string)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *NetworkRouterFirewallRule) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetSourceTier

`func (o *NetworkRouterFirewallRule) GetSourceTier() string`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *NetworkRouterFirewallRule) GetSourceTierOk() (*string, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *NetworkRouterFirewallRule) SetSourceTier(v string)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *NetworkRouterFirewallRule) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### GetApplications

`func (o *NetworkRouterFirewallRule) GetApplications() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *NetworkRouterFirewallRule) GetApplicationsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *NetworkRouterFirewallRule) SetApplications(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *NetworkRouterFirewallRule) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



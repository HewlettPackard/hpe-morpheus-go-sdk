# GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
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
**Protocol** | Pointer to **NullableString** |  | [optional] 
**Application** | Pointer to **NullableString** |  | [optional] 
**ApplicationType** | Pointer to **string** |  | [optional] 
**PortRange** | Pointer to **NullableString** |  | [optional] 
**SourcePortRange** | Pointer to **NullableString** |  | [optional] 
**DestinationPortRange** | Pointer to **NullableString** |  | [optional] 
**SourceGroup** | Pointer to **NullableString** |  | [optional] 
**SourceTier** | Pointer to **NullableString** |  | [optional] 
**Applications** | Pointer to [**[]GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInnerApplicationsInner**](GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInnerApplicationsInner.md) |  | [optional] 

## Methods

### NewGetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner

`func NewGetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner() *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner`

NewGetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner instantiates a new GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetEnabled

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroupName

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDirection

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetRuleType

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetPolicy

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSource

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSource() []string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSourceOk() (*[]string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetSource(v []string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceType

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestination

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDestination() []string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDestinationOk() (*[]string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetDestination(v []string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationType

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetProfiles

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetProfiles() []string`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetProfilesOk() (*[]string, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetProfiles(v []string)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetProtocol

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetApplication

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetApplicationType

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetPortRange

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### SetPortRangeNil

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetPortRangeNil(b bool)`

 SetPortRangeNil sets the value for PortRange to be an explicit nil

### UnsetPortRange
`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) UnsetPortRange()`

UnsetPortRange ensures that no value is present for PortRange, not even an explicit nil
### GetSourcePortRange

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSourcePortRange() string`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSourcePortRangeOk() (*string, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetSourcePortRange(v string)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### SetSourcePortRangeNil

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetSourcePortRangeNil(b bool)`

 SetSourcePortRangeNil sets the value for SourcePortRange to be an explicit nil

### UnsetSourcePortRange
`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) UnsetSourcePortRange()`

UnsetSourcePortRange ensures that no value is present for SourcePortRange, not even an explicit nil
### GetDestinationPortRange

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDestinationPortRange() string`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetDestinationPortRangeOk() (*string, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetDestinationPortRange(v string)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.

### SetDestinationPortRangeNil

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetDestinationPortRangeNil(b bool)`

 SetDestinationPortRangeNil sets the value for DestinationPortRange to be an explicit nil

### UnsetDestinationPortRange
`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) UnsetDestinationPortRange()`

UnsetDestinationPortRange ensures that no value is present for DestinationPortRange, not even an explicit nil
### GetSourceGroup

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSourceGroup() string`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSourceGroupOk() (*string, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetSourceGroup(v string)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### SetSourceGroupNil

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetSourceGroupNil(b bool)`

 SetSourceGroupNil sets the value for SourceGroup to be an explicit nil

### UnsetSourceGroup
`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) UnsetSourceGroup()`

UnsetSourceGroup ensures that no value is present for SourceGroup, not even an explicit nil
### GetSourceTier

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSourceTier() string`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetSourceTierOk() (*string, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetSourceTier(v string)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### SetSourceTierNil

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetSourceTierNil(b bool)`

 SetSourceTierNil sets the value for SourceTier to be an explicit nil

### UnsetSourceTier
`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) UnsetSourceTier()`

UnsetSourceTier ensures that no value is present for SourceTier, not even an explicit nil
### GetApplications

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetApplications() []GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInnerApplicationsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) GetApplicationsOk() (*[]GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInnerApplicationsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) SetApplications(v []GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInnerApplicationsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *GetNetworkRouterFirewallRuleGroup200ResponseRuleGroupRulesInner) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetNetworkRouterFirewallRule200ResponseRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
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
**Protocol** | Pointer to **NullableString** |  | [optional] 
**Application** | Pointer to **NullableString** |  | [optional] 
**ApplicationType** | Pointer to **string** |  | [optional] 
**PortRange** | Pointer to **NullableString** |  | [optional] 
**SourcePortRange** | Pointer to **NullableString** |  | [optional] 
**DestinationPortRange** | Pointer to **NullableString** |  | [optional] 
**SourceGroup** | Pointer to **NullableString** |  | [optional] 
**SourceTier** | Pointer to **NullableString** |  | [optional] 
**Applications** | Pointer to [**[]GetNetworkRouterFirewallRule200ResponseRuleApplicationsInner**](GetNetworkRouterFirewallRule200ResponseRuleApplicationsInner.md) |  | [optional] 

## Methods

### NewGetNetworkRouterFirewallRule200ResponseRule

`func NewGetNetworkRouterFirewallRule200ResponseRule() *GetNetworkRouterFirewallRule200ResponseRule`

NewGetNetworkRouterFirewallRule200ResponseRule instantiates a new GetNetworkRouterFirewallRule200ResponseRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetNetworkRouterFirewallRule200ResponseRule) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetEnabled

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetGroupName

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDirection

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetRuleType

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetPolicy

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSource

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetSource() []string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetSourceOk() (*[]string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetSource(v []string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceType

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestination

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetDestination() []string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetDestinationOk() (*[]string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetDestination(v []string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationType

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetProfiles

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetProfiles() []string`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetProfilesOk() (*[]string, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetProfiles(v []string)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetProtocol

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *GetNetworkRouterFirewallRule200ResponseRule) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetApplication

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *GetNetworkRouterFirewallRule200ResponseRule) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetApplicationType

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetPortRange

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### SetPortRangeNil

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetPortRangeNil(b bool)`

 SetPortRangeNil sets the value for PortRange to be an explicit nil

### UnsetPortRange
`func (o *GetNetworkRouterFirewallRule200ResponseRule) UnsetPortRange()`

UnsetPortRange ensures that no value is present for PortRange, not even an explicit nil
### GetSourcePortRange

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetSourcePortRange() string`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetSourcePortRangeOk() (*string, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetSourcePortRange(v string)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### SetSourcePortRangeNil

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetSourcePortRangeNil(b bool)`

 SetSourcePortRangeNil sets the value for SourcePortRange to be an explicit nil

### UnsetSourcePortRange
`func (o *GetNetworkRouterFirewallRule200ResponseRule) UnsetSourcePortRange()`

UnsetSourcePortRange ensures that no value is present for SourcePortRange, not even an explicit nil
### GetDestinationPortRange

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetDestinationPortRange() string`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetDestinationPortRangeOk() (*string, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetDestinationPortRange(v string)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.

### SetDestinationPortRangeNil

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetDestinationPortRangeNil(b bool)`

 SetDestinationPortRangeNil sets the value for DestinationPortRange to be an explicit nil

### UnsetDestinationPortRange
`func (o *GetNetworkRouterFirewallRule200ResponseRule) UnsetDestinationPortRange()`

UnsetDestinationPortRange ensures that no value is present for DestinationPortRange, not even an explicit nil
### GetSourceGroup

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetSourceGroup() string`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetSourceGroupOk() (*string, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetSourceGroup(v string)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### SetSourceGroupNil

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetSourceGroupNil(b bool)`

 SetSourceGroupNil sets the value for SourceGroup to be an explicit nil

### UnsetSourceGroup
`func (o *GetNetworkRouterFirewallRule200ResponseRule) UnsetSourceGroup()`

UnsetSourceGroup ensures that no value is present for SourceGroup, not even an explicit nil
### GetSourceTier

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetSourceTier() string`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetSourceTierOk() (*string, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetSourceTier(v string)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### SetSourceTierNil

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetSourceTierNil(b bool)`

 SetSourceTierNil sets the value for SourceTier to be an explicit nil

### UnsetSourceTier
`func (o *GetNetworkRouterFirewallRule200ResponseRule) UnsetSourceTier()`

UnsetSourceTier ensures that no value is present for SourceTier, not even an explicit nil
### GetApplications

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetApplications() []GetNetworkRouterFirewallRule200ResponseRuleApplicationsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *GetNetworkRouterFirewallRule200ResponseRule) GetApplicationsOk() (*[]GetNetworkRouterFirewallRule200ResponseRuleApplicationsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *GetNetworkRouterFirewallRule200ResponseRule) SetApplications(v []GetNetworkRouterFirewallRule200ResponseRuleApplicationsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *GetNetworkRouterFirewallRule200ResponseRule) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



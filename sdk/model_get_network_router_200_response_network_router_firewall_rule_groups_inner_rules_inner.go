/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 8.0.7
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner{}

// GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner struct for GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner
type GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner struct {
	Id                   *int64                                                                   `json:"id,omitempty"`
	Name                 *string                                                                  `json:"name,omitempty"`
	Code                 *string                                                                  `json:"code,omitempty"`
	Enabled              *bool                                                                    `json:"enabled,omitempty"`
	GroupName            *string                                                                  `json:"groupName,omitempty"`
	Direction            *string                                                                  `json:"direction,omitempty"`
	RuleType             *string                                                                  `json:"ruleType,omitempty"`
	Policy               *string                                                                  `json:"policy,omitempty"`
	Source               []string                                                                 `json:"source,omitempty"`
	SourceType           *string                                                                  `json:"sourceType,omitempty"`
	Destination          []string                                                                 `json:"destination,omitempty"`
	DestinationType      *string                                                                  `json:"destinationType,omitempty"`
	Profiles             []string                                                                 `json:"profiles,omitempty"`
	Protocol             *string                                                                  `json:"protocol,omitempty"`
	Application          *string                                                                  `json:"application,omitempty"`
	ApplicationType      *string                                                                  `json:"applicationType,omitempty"`
	PortRange            *string                                                                  `json:"portRange,omitempty"`
	SourcePortRange      *string                                                                  `json:"sourcePortRange,omitempty"`
	DestinationPortRange *string                                                                  `json:"destinationPortRange,omitempty"`
	SourceGroup          *string                                                                  `json:"sourceGroup,omitempty"`
	SourceTier           *string                                                                  `json:"sourceTier,omitempty"`
	Applications         []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"applications,omitempty"`
	AdditionalProperties map[string]interface{}                                                   `json:",remain"`
}

type _GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner

// NewGetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner instantiates a new GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner() *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner {
	this := GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner{}
	return &this
}

// NewGetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInnerWithDefaults instantiates a new GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInnerWithDefaults() *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner {
	this := GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetCode(v string) {
	o.Code = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// IsSetGroupName returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetGroupName(v string) {
	o.GroupName = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// IsSetDirection returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetDirection(v string) {
	o.Direction = &v
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetRuleType() string {
	if o == nil || IsNil(o.RuleType) {
		var ret string
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetRuleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RuleType) {
		return nil, false
	}
	return o.RuleType, true
}

// IsSetRuleType returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetRuleType() bool {
	if o != nil && !IsNil(o.RuleType) {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given string and assigns it to the RuleType field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetRuleType(v string) {
	o.RuleType = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetPolicy() string {
	if o == nil || IsNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// IsSetPolicy returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetPolicy(v string) {
	o.Policy = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSource() []string {
	if o == nil || IsNil(o.Source) {
		var ret []string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceOk() ([]string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// IsSetSource returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given []string and assigns it to the Source field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetSource(v []string) {
	o.Source = v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceType() string {
	if o == nil || IsNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// IsSetSourceType returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetSourceType() bool {
	if o != nil && !IsNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetSourceType(v string) {
	o.SourceType = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDestination() []string {
	if o == nil || IsNil(o.Destination) {
		var ret []string
		return ret
	}
	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDestinationOk() ([]string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// IsSetDestination returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given []string and assigns it to the Destination field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetDestination(v []string) {
	o.Destination = v
}

// GetDestinationType returns the DestinationType field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDestinationType() string {
	if o == nil || IsNil(o.DestinationType) {
		var ret string
		return ret
	}
	return *o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDestinationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationType) {
		return nil, false
	}
	return o.DestinationType, true
}

// IsSetDestinationType returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetDestinationType() bool {
	if o != nil && !IsNil(o.DestinationType) {
		return true
	}

	return false
}

// SetDestinationType gets a reference to the given string and assigns it to the DestinationType field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetDestinationType(v string) {
	o.DestinationType = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetProfiles() []string {
	if o == nil || IsNil(o.Profiles) {
		var ret []string
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetProfilesOk() ([]string, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// IsSetProfiles returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []string and assigns it to the Profiles field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetProfiles(v []string) {
	o.Profiles = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// IsSetProtocol returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetApplication() string {
	if o == nil || IsNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetApplicationOk() (*string, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// IsSetApplication returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetApplication(v string) {
	o.Application = &v
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetApplicationType() string {
	if o == nil || IsNil(o.ApplicationType) {
		var ret string
		return ret
	}
	return *o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetApplicationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationType) {
		return nil, false
	}
	return o.ApplicationType, true
}

// IsSetApplicationType returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetApplicationType() bool {
	if o != nil && !IsNil(o.ApplicationType) {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given string and assigns it to the ApplicationType field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetApplicationType(v string) {
	o.ApplicationType = &v
}

// GetPortRange returns the PortRange field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetPortRange() string {
	if o == nil || IsNil(o.PortRange) {
		var ret string
		return ret
	}
	return *o.PortRange
}

// GetPortRangeOk returns a tuple with the PortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetPortRangeOk() (*string, bool) {
	if o == nil || IsNil(o.PortRange) {
		return nil, false
	}
	return o.PortRange, true
}

// IsSetPortRange returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetPortRange() bool {
	if o != nil && !IsNil(o.PortRange) {
		return true
	}

	return false
}

// SetPortRange gets a reference to the given string and assigns it to the PortRange field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetPortRange(v string) {
	o.PortRange = &v
}

// GetSourcePortRange returns the SourcePortRange field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourcePortRange() string {
	if o == nil || IsNil(o.SourcePortRange) {
		var ret string
		return ret
	}
	return *o.SourcePortRange
}

// GetSourcePortRangeOk returns a tuple with the SourcePortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourcePortRangeOk() (*string, bool) {
	if o == nil || IsNil(o.SourcePortRange) {
		return nil, false
	}
	return o.SourcePortRange, true
}

// IsSetSourcePortRange returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetSourcePortRange() bool {
	if o != nil && !IsNil(o.SourcePortRange) {
		return true
	}

	return false
}

// SetSourcePortRange gets a reference to the given string and assigns it to the SourcePortRange field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetSourcePortRange(v string) {
	o.SourcePortRange = &v
}

// GetDestinationPortRange returns the DestinationPortRange field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDestinationPortRange() string {
	if o == nil || IsNil(o.DestinationPortRange) {
		var ret string
		return ret
	}
	return *o.DestinationPortRange
}

// GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetDestinationPortRangeOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationPortRange) {
		return nil, false
	}
	return o.DestinationPortRange, true
}

// IsSetDestinationPortRange returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetDestinationPortRange() bool {
	if o != nil && !IsNil(o.DestinationPortRange) {
		return true
	}

	return false
}

// SetDestinationPortRange gets a reference to the given string and assigns it to the DestinationPortRange field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetDestinationPortRange(v string) {
	o.DestinationPortRange = &v
}

// GetSourceGroup returns the SourceGroup field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceGroup() string {
	if o == nil || IsNil(o.SourceGroup) {
		var ret string
		return ret
	}
	return *o.SourceGroup
}

// GetSourceGroupOk returns a tuple with the SourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceGroupOk() (*string, bool) {
	if o == nil || IsNil(o.SourceGroup) {
		return nil, false
	}
	return o.SourceGroup, true
}

// IsSetSourceGroup returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetSourceGroup() bool {
	if o != nil && !IsNil(o.SourceGroup) {
		return true
	}

	return false
}

// SetSourceGroup gets a reference to the given string and assigns it to the SourceGroup field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetSourceGroup(v string) {
	o.SourceGroup = &v
}

// GetSourceTier returns the SourceTier field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceTier() string {
	if o == nil || IsNil(o.SourceTier) {
		var ret string
		return ret
	}
	return *o.SourceTier
}

// GetSourceTierOk returns a tuple with the SourceTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetSourceTierOk() (*string, bool) {
	if o == nil || IsNil(o.SourceTier) {
		return nil, false
	}
	return o.SourceTier, true
}

// IsSetSourceTier returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetSourceTier() bool {
	if o != nil && !IsNil(o.SourceTier) {
		return true
	}

	return false
}

// SetSourceTier gets a reference to the given string and assigns it to the SourceTier field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetSourceTier(v string) {
	o.SourceTier = &v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetApplications() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Applications) {
		var ret []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) GetApplicationsOk() ([]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// IsSetApplications returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) IsSetApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Applications field.
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) SetApplications(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Applications = v
}

func (o GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.RuleType) {
		toSerialize["ruleType"] = o.RuleType
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.DestinationType) {
		toSerialize["destinationType"] = o.DestinationType
	}
	if !IsNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.ApplicationType) {
		toSerialize["applicationType"] = o.ApplicationType
	}
	if !IsNil(o.PortRange) {
		toSerialize["portRange"] = o.PortRange
	}
	if !IsNil(o.SourcePortRange) {
		toSerialize["sourcePortRange"] = o.SourcePortRange
	}
	if !IsNil(o.DestinationPortRange) {
		toSerialize["destinationPortRange"] = o.DestinationPortRange
	}
	if !IsNil(o.SourceGroup) {
		toSerialize["sourceGroup"] = o.SourceGroup
	}
	if !IsNil(o.SourceTier) {
		toSerialize["sourceTier"] = o.SourceTier
	}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

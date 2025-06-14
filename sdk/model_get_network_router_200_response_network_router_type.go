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

// checks if the GetNetworkRouter200ResponseNetworkRouterType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkRouter200ResponseNetworkRouterType{}

// GetNetworkRouter200ResponseNetworkRouterType struct for GetNetworkRouter200ResponseNetworkRouterType
type GetNetworkRouter200ResponseNetworkRouterType struct {
	Id                       *int64                                                                      `json:"id,omitempty"`
	Code                     *string                                                                     `json:"code,omitempty"`
	Name                     *string                                                                     `json:"name,omitempty"`
	Description              *string                                                                     `json:"description,omitempty"`
	Enabled                  *bool                                                                       `json:"enabled,omitempty"`
	Creatable                *bool                                                                       `json:"creatable,omitempty"`
	Selectable               *bool                                                                       `json:"selectable,omitempty"`
	HasFirewall              *bool                                                                       `json:"hasFirewall,omitempty"`
	HasDhcp                  *bool                                                                       `json:"hasDhcp,omitempty"`
	HasRouting               *bool                                                                       `json:"hasRouting,omitempty"`
	HasNat                   *bool                                                                       `json:"hasNat,omitempty"`
	HasNetworkServer         *bool                                                                       `json:"hasNetworkServer,omitempty"`
	HasFirewallGroups        *bool                                                                       `json:"hasFirewallGroups,omitempty"`
	HasSecurityGroupPriority *bool                                                                       `json:"hasSecurityGroupPriority,omitempty"`
	OptionTypes              []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner `json:"optionTypes,omitempty"`
	RuleOptionTypes          []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner `json:"ruleOptionTypes,omitempty"`
	FirewallGroupOptionTypes []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner `json:"firewallGroupOptionTypes,omitempty"`
	NatOptionTypes           []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner `json:"natOptionTypes,omitempty"`
	AdditionalProperties     map[string]interface{}                                                      `json:",remain"`
}

type _GetNetworkRouter200ResponseNetworkRouterType GetNetworkRouter200ResponseNetworkRouterType

// NewGetNetworkRouter200ResponseNetworkRouterType instantiates a new GetNetworkRouter200ResponseNetworkRouterType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkRouter200ResponseNetworkRouterType() *GetNetworkRouter200ResponseNetworkRouterType {
	this := GetNetworkRouter200ResponseNetworkRouterType{}
	return &this
}

// NewGetNetworkRouter200ResponseNetworkRouterTypeWithDefaults instantiates a new GetNetworkRouter200ResponseNetworkRouterType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkRouter200ResponseNetworkRouterTypeWithDefaults() *GetNetworkRouter200ResponseNetworkRouterType {
	this := GetNetworkRouter200ResponseNetworkRouterType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatable returns the Creatable field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetCreatable() bool {
	if o == nil || IsNil(o.Creatable) {
		var ret bool
		return ret
	}
	return *o.Creatable
}

// GetCreatableOk returns a tuple with the Creatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetCreatableOk() (*bool, bool) {
	if o == nil || IsNil(o.Creatable) {
		return nil, false
	}
	return o.Creatable, true
}

// IsSetCreatable returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetCreatable() bool {
	if o != nil && !IsNil(o.Creatable) {
		return true
	}

	return false
}

// SetCreatable gets a reference to the given bool and assigns it to the Creatable field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetCreatable(v bool) {
	o.Creatable = &v
}

// GetSelectable returns the Selectable field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetSelectable() bool {
	if o == nil || IsNil(o.Selectable) {
		var ret bool
		return ret
	}
	return *o.Selectable
}

// GetSelectableOk returns a tuple with the Selectable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetSelectableOk() (*bool, bool) {
	if o == nil || IsNil(o.Selectable) {
		return nil, false
	}
	return o.Selectable, true
}

// IsSetSelectable returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetSelectable() bool {
	if o != nil && !IsNil(o.Selectable) {
		return true
	}

	return false
}

// SetSelectable gets a reference to the given bool and assigns it to the Selectable field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetSelectable(v bool) {
	o.Selectable = &v
}

// GetHasFirewall returns the HasFirewall field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasFirewall() bool {
	if o == nil || IsNil(o.HasFirewall) {
		var ret bool
		return ret
	}
	return *o.HasFirewall
}

// GetHasFirewallOk returns a tuple with the HasFirewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasFirewallOk() (*bool, bool) {
	if o == nil || IsNil(o.HasFirewall) {
		return nil, false
	}
	return o.HasFirewall, true
}

// IsSetHasFirewall returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetHasFirewall() bool {
	if o != nil && !IsNil(o.HasFirewall) {
		return true
	}

	return false
}

// SetHasFirewall gets a reference to the given bool and assigns it to the HasFirewall field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasFirewall(v bool) {
	o.HasFirewall = &v
}

// GetHasDhcp returns the HasDhcp field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasDhcp() bool {
	if o == nil || IsNil(o.HasDhcp) {
		var ret bool
		return ret
	}
	return *o.HasDhcp
}

// GetHasDhcpOk returns a tuple with the HasDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasDhcpOk() (*bool, bool) {
	if o == nil || IsNil(o.HasDhcp) {
		return nil, false
	}
	return o.HasDhcp, true
}

// IsSetHasDhcp returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetHasDhcp() bool {
	if o != nil && !IsNil(o.HasDhcp) {
		return true
	}

	return false
}

// SetHasDhcp gets a reference to the given bool and assigns it to the HasDhcp field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasDhcp(v bool) {
	o.HasDhcp = &v
}

// GetHasRouting returns the HasRouting field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasRouting() bool {
	if o == nil || IsNil(o.HasRouting) {
		var ret bool
		return ret
	}
	return *o.HasRouting
}

// GetHasRoutingOk returns a tuple with the HasRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasRoutingOk() (*bool, bool) {
	if o == nil || IsNil(o.HasRouting) {
		return nil, false
	}
	return o.HasRouting, true
}

// IsSetHasRouting returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetHasRouting() bool {
	if o != nil && !IsNil(o.HasRouting) {
		return true
	}

	return false
}

// SetHasRouting gets a reference to the given bool and assigns it to the HasRouting field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasRouting(v bool) {
	o.HasRouting = &v
}

// GetHasNat returns the HasNat field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasNat() bool {
	if o == nil || IsNil(o.HasNat) {
		var ret bool
		return ret
	}
	return *o.HasNat
}

// GetHasNatOk returns a tuple with the HasNat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasNatOk() (*bool, bool) {
	if o == nil || IsNil(o.HasNat) {
		return nil, false
	}
	return o.HasNat, true
}

// IsSetHasNat returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetHasNat() bool {
	if o != nil && !IsNil(o.HasNat) {
		return true
	}

	return false
}

// SetHasNat gets a reference to the given bool and assigns it to the HasNat field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasNat(v bool) {
	o.HasNat = &v
}

// GetHasNetworkServer returns the HasNetworkServer field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasNetworkServer() bool {
	if o == nil || IsNil(o.HasNetworkServer) {
		var ret bool
		return ret
	}
	return *o.HasNetworkServer
}

// GetHasNetworkServerOk returns a tuple with the HasNetworkServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasNetworkServerOk() (*bool, bool) {
	if o == nil || IsNil(o.HasNetworkServer) {
		return nil, false
	}
	return o.HasNetworkServer, true
}

// IsSetHasNetworkServer returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetHasNetworkServer() bool {
	if o != nil && !IsNil(o.HasNetworkServer) {
		return true
	}

	return false
}

// SetHasNetworkServer gets a reference to the given bool and assigns it to the HasNetworkServer field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasNetworkServer(v bool) {
	o.HasNetworkServer = &v
}

// GetHasFirewallGroups returns the HasFirewallGroups field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasFirewallGroups() bool {
	if o == nil || IsNil(o.HasFirewallGroups) {
		var ret bool
		return ret
	}
	return *o.HasFirewallGroups
}

// GetHasFirewallGroupsOk returns a tuple with the HasFirewallGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasFirewallGroupsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasFirewallGroups) {
		return nil, false
	}
	return o.HasFirewallGroups, true
}

// IsSetHasFirewallGroups returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetHasFirewallGroups() bool {
	if o != nil && !IsNil(o.HasFirewallGroups) {
		return true
	}

	return false
}

// SetHasFirewallGroups gets a reference to the given bool and assigns it to the HasFirewallGroups field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasFirewallGroups(v bool) {
	o.HasFirewallGroups = &v
}

// GetHasSecurityGroupPriority returns the HasSecurityGroupPriority field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasSecurityGroupPriority() bool {
	if o == nil || IsNil(o.HasSecurityGroupPriority) {
		var ret bool
		return ret
	}
	return *o.HasSecurityGroupPriority
}

// GetHasSecurityGroupPriorityOk returns a tuple with the HasSecurityGroupPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasSecurityGroupPriorityOk() (*bool, bool) {
	if o == nil || IsNil(o.HasSecurityGroupPriority) {
		return nil, false
	}
	return o.HasSecurityGroupPriority, true
}

// IsSetHasSecurityGroupPriority returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetHasSecurityGroupPriority() bool {
	if o != nil && !IsNil(o.HasSecurityGroupPriority) {
		return true
	}

	return false
}

// SetHasSecurityGroupPriority gets a reference to the given bool and assigns it to the HasSecurityGroupPriority field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasSecurityGroupPriority(v bool) {
	o.HasSecurityGroupPriority = &v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner {
	if o == nil || IsNil(o.OptionTypes) {
		var ret []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner
		return ret
	}
	return o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetOptionTypesOk() ([]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool) {
	if o == nil || IsNil(o.OptionTypes) {
		return nil, false
	}
	return o.OptionTypes, true
}

// IsSetOptionTypes returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetOptionTypes() bool {
	if o != nil && !IsNil(o.OptionTypes) {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner and assigns it to the OptionTypes field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner) {
	o.OptionTypes = v
}

// GetRuleOptionTypes returns the RuleOptionTypes field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetRuleOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner {
	if o == nil || IsNil(o.RuleOptionTypes) {
		var ret []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner
		return ret
	}
	return o.RuleOptionTypes
}

// GetRuleOptionTypesOk returns a tuple with the RuleOptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetRuleOptionTypesOk() ([]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool) {
	if o == nil || IsNil(o.RuleOptionTypes) {
		return nil, false
	}
	return o.RuleOptionTypes, true
}

// IsSetRuleOptionTypes returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetRuleOptionTypes() bool {
	if o != nil && !IsNil(o.RuleOptionTypes) {
		return true
	}

	return false
}

// SetRuleOptionTypes gets a reference to the given []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner and assigns it to the RuleOptionTypes field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetRuleOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner) {
	o.RuleOptionTypes = v
}

// GetFirewallGroupOptionTypes returns the FirewallGroupOptionTypes field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetFirewallGroupOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner {
	if o == nil || IsNil(o.FirewallGroupOptionTypes) {
		var ret []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner
		return ret
	}
	return o.FirewallGroupOptionTypes
}

// GetFirewallGroupOptionTypesOk returns a tuple with the FirewallGroupOptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetFirewallGroupOptionTypesOk() ([]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool) {
	if o == nil || IsNil(o.FirewallGroupOptionTypes) {
		return nil, false
	}
	return o.FirewallGroupOptionTypes, true
}

// IsSetFirewallGroupOptionTypes returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetFirewallGroupOptionTypes() bool {
	if o != nil && !IsNil(o.FirewallGroupOptionTypes) {
		return true
	}

	return false
}

// SetFirewallGroupOptionTypes gets a reference to the given []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner and assigns it to the FirewallGroupOptionTypes field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetFirewallGroupOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner) {
	o.FirewallGroupOptionTypes = v
}

// GetNatOptionTypes returns the NatOptionTypes field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetNatOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner {
	if o == nil || IsNil(o.NatOptionTypes) {
		var ret []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner
		return ret
	}
	return o.NatOptionTypes
}

// GetNatOptionTypesOk returns a tuple with the NatOptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) GetNatOptionTypesOk() ([]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool) {
	if o == nil || IsNil(o.NatOptionTypes) {
		return nil, false
	}
	return o.NatOptionTypes, true
}

// IsSetNatOptionTypes returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterType) IsSetNatOptionTypes() bool {
	if o != nil && !IsNil(o.NatOptionTypes) {
		return true
	}

	return false
}

// SetNatOptionTypes gets a reference to the given []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner and assigns it to the NatOptionTypes field.
func (o *GetNetworkRouter200ResponseNetworkRouterType) SetNatOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner) {
	o.NatOptionTypes = v
}

func (o GetNetworkRouter200ResponseNetworkRouterType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkRouter200ResponseNetworkRouterType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Creatable) {
		toSerialize["creatable"] = o.Creatable
	}
	if !IsNil(o.Selectable) {
		toSerialize["selectable"] = o.Selectable
	}
	if !IsNil(o.HasFirewall) {
		toSerialize["hasFirewall"] = o.HasFirewall
	}
	if !IsNil(o.HasDhcp) {
		toSerialize["hasDhcp"] = o.HasDhcp
	}
	if !IsNil(o.HasRouting) {
		toSerialize["hasRouting"] = o.HasRouting
	}
	if !IsNil(o.HasNat) {
		toSerialize["hasNat"] = o.HasNat
	}
	if !IsNil(o.HasNetworkServer) {
		toSerialize["hasNetworkServer"] = o.HasNetworkServer
	}
	if !IsNil(o.HasFirewallGroups) {
		toSerialize["hasFirewallGroups"] = o.HasFirewallGroups
	}
	if !IsNil(o.HasSecurityGroupPriority) {
		toSerialize["hasSecurityGroupPriority"] = o.HasSecurityGroupPriority
	}
	if !IsNil(o.OptionTypes) {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	if !IsNil(o.RuleOptionTypes) {
		toSerialize["ruleOptionTypes"] = o.RuleOptionTypes
	}
	if !IsNil(o.FirewallGroupOptionTypes) {
		toSerialize["firewallGroupOptionTypes"] = o.FirewallGroupOptionTypes
	}
	if !IsNil(o.NatOptionTypes) {
		toSerialize["natOptionTypes"] = o.NatOptionTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetNetworkRouter200ResponseNetworkRouterType) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

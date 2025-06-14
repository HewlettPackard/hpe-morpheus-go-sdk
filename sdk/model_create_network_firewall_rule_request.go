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

// checks if the CreateNetworkFirewallRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirewallRuleRequest{}

// CreateNetworkFirewallRuleRequest struct for CreateNetworkFirewallRuleRequest
type CreateNetworkFirewallRuleRequest struct {
	Rule                 *CreateNetworkFirewallRuleRequestRule `json:"rule,omitempty"`
	AdditionalProperties map[string]interface{}                `json:",remain"`
}

type _CreateNetworkFirewallRuleRequest CreateNetworkFirewallRuleRequest

// NewCreateNetworkFirewallRuleRequest instantiates a new CreateNetworkFirewallRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirewallRuleRequest() *CreateNetworkFirewallRuleRequest {
	this := CreateNetworkFirewallRuleRequest{}
	return &this
}

// NewCreateNetworkFirewallRuleRequestWithDefaults instantiates a new CreateNetworkFirewallRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirewallRuleRequestWithDefaults() *CreateNetworkFirewallRuleRequest {
	this := CreateNetworkFirewallRuleRequest{}
	return &this
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *CreateNetworkFirewallRuleRequest) GetRule() CreateNetworkFirewallRuleRequestRule {
	if o == nil || IsNil(o.Rule) {
		var ret CreateNetworkFirewallRuleRequestRule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirewallRuleRequest) GetRuleOk() (*CreateNetworkFirewallRuleRequestRule, bool) {
	if o == nil || IsNil(o.Rule) {
		return nil, false
	}
	return o.Rule, true
}

// IsSetRule returns a boolean if a field has been set.
func (o *CreateNetworkFirewallRuleRequest) IsSetRule() bool {
	if o != nil && !IsNil(o.Rule) {
		return true
	}

	return false
}

// SetRule gets a reference to the given CreateNetworkFirewallRuleRequestRule and assigns it to the Rule field.
func (o *CreateNetworkFirewallRuleRequest) SetRule(v CreateNetworkFirewallRuleRequestRule) {
	o.Rule = &v
}

func (o CreateNetworkFirewallRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirewallRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rule) {
		toSerialize["rule"] = o.Rule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *CreateNetworkFirewallRuleRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the UpdateSecurityGroupRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSecurityGroupRulesRequest{}

// UpdateSecurityGroupRulesRequest struct for UpdateSecurityGroupRulesRequest
type UpdateSecurityGroupRulesRequest struct {
	Rule                 UpdateSecurityGroupRulesRequestRule `json:"rule"`
	AdditionalProperties map[string]interface{}              `json:",remain"`
}

type _UpdateSecurityGroupRulesRequest UpdateSecurityGroupRulesRequest

// NewUpdateSecurityGroupRulesRequest instantiates a new UpdateSecurityGroupRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSecurityGroupRulesRequest(rule UpdateSecurityGroupRulesRequestRule) *UpdateSecurityGroupRulesRequest {
	this := UpdateSecurityGroupRulesRequest{}
	this.Rule = rule
	return &this
}

// NewUpdateSecurityGroupRulesRequestWithDefaults instantiates a new UpdateSecurityGroupRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSecurityGroupRulesRequestWithDefaults() *UpdateSecurityGroupRulesRequest {
	this := UpdateSecurityGroupRulesRequest{}
	return &this
}

// GetRule returns the Rule field value
func (o *UpdateSecurityGroupRulesRequest) GetRule() UpdateSecurityGroupRulesRequestRule {
	if o == nil {
		var ret UpdateSecurityGroupRulesRequestRule
		return ret
	}

	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *UpdateSecurityGroupRulesRequest) GetRuleOk() (*UpdateSecurityGroupRulesRequestRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value
func (o *UpdateSecurityGroupRulesRequest) SetRule(v UpdateSecurityGroupRulesRequestRule) {
	o.Rule = v
}

func (o UpdateSecurityGroupRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSecurityGroupRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rule"] = o.Rule

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateSecurityGroupRulesRequest) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

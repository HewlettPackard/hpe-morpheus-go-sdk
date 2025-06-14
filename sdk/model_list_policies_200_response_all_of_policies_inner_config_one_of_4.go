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

// checks if the ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4{}

// ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4 - Cluster Resource Name
type ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4 struct {
	AccountIntegrationId *string                `json:"accountIntegrationId,omitempty"`
	ServerNamingType     string                 `json:"serverNamingType"`
	ServerNamingPattern  *string                `json:"serverNamingPattern,omitempty"`
	ServerNamingConflict *bool                  `json:"serverNamingConflict,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4 ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4

// NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4 instantiates a new ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4(serverNamingType string) *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4 {
	this := ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4{}
	this.ServerNamingType = serverNamingType
	return &this
}

// NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4WithDefaults instantiates a new ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4WithDefaults() *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4 {
	this := ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4{}
	return &this
}

// GetAccountIntegrationId returns the AccountIntegrationId field value if set, zero value otherwise.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) GetAccountIntegrationId() string {
	if o == nil || IsNil(o.AccountIntegrationId) {
		var ret string
		return ret
	}
	return *o.AccountIntegrationId
}

// GetAccountIntegrationIdOk returns a tuple with the AccountIntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) GetAccountIntegrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountIntegrationId) {
		return nil, false
	}
	return o.AccountIntegrationId, true
}

// IsSetAccountIntegrationId returns a boolean if a field has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) IsSetAccountIntegrationId() bool {
	if o != nil && !IsNil(o.AccountIntegrationId) {
		return true
	}

	return false
}

// SetAccountIntegrationId gets a reference to the given string and assigns it to the AccountIntegrationId field.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) SetAccountIntegrationId(v string) {
	o.AccountIntegrationId = &v
}

// GetServerNamingType returns the ServerNamingType field value
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) GetServerNamingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerNamingType
}

// GetServerNamingTypeOk returns a tuple with the ServerNamingType field value
// and a boolean to check if the value has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) GetServerNamingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerNamingType, true
}

// SetServerNamingType sets field value
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) SetServerNamingType(v string) {
	o.ServerNamingType = v
}

// GetServerNamingPattern returns the ServerNamingPattern field value if set, zero value otherwise.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) GetServerNamingPattern() string {
	if o == nil || IsNil(o.ServerNamingPattern) {
		var ret string
		return ret
	}
	return *o.ServerNamingPattern
}

// GetServerNamingPatternOk returns a tuple with the ServerNamingPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) GetServerNamingPatternOk() (*string, bool) {
	if o == nil || IsNil(o.ServerNamingPattern) {
		return nil, false
	}
	return o.ServerNamingPattern, true
}

// IsSetServerNamingPattern returns a boolean if a field has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) IsSetServerNamingPattern() bool {
	if o != nil && !IsNil(o.ServerNamingPattern) {
		return true
	}

	return false
}

// SetServerNamingPattern gets a reference to the given string and assigns it to the ServerNamingPattern field.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) SetServerNamingPattern(v string) {
	o.ServerNamingPattern = &v
}

// GetServerNamingConflict returns the ServerNamingConflict field value if set, zero value otherwise.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) GetServerNamingConflict() bool {
	if o == nil || IsNil(o.ServerNamingConflict) {
		var ret bool
		return ret
	}
	return *o.ServerNamingConflict
}

// GetServerNamingConflictOk returns a tuple with the ServerNamingConflict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) GetServerNamingConflictOk() (*bool, bool) {
	if o == nil || IsNil(o.ServerNamingConflict) {
		return nil, false
	}
	return o.ServerNamingConflict, true
}

// IsSetServerNamingConflict returns a boolean if a field has been set.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) IsSetServerNamingConflict() bool {
	if o != nil && !IsNil(o.ServerNamingConflict) {
		return true
	}

	return false
}

// SetServerNamingConflict gets a reference to the given bool and assigns it to the ServerNamingConflict field.
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) SetServerNamingConflict(v bool) {
	o.ServerNamingConflict = &v
}

func (o ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountIntegrationId) {
		toSerialize["accountIntegrationId"] = o.AccountIntegrationId
	}
	toSerialize["serverNamingType"] = o.ServerNamingType
	if !IsNil(o.ServerNamingPattern) {
		toSerialize["serverNamingPattern"] = o.ServerNamingPattern
	}
	if !IsNil(o.ServerNamingConflict) {
		toSerialize["serverNamingConflict"] = o.ServerNamingConflict
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListPolicies200ResponseAllOfPoliciesInnerConfigOneOf4) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the CloudFoundryResourcePoolConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudFoundryResourcePoolConfig{}

// CloudFoundryResourcePoolConfig struct for CloudFoundryResourcePoolConfig
type CloudFoundryResourcePoolConfig struct {
	// Array of manager usernames
	Managers []string `json:"managers,omitempty"`
	// Array of developer usernames
	Developers []string `json:"developers,omitempty"`
	// Array of auditor usernames
	Auditors             []string               `json:"auditors,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _CloudFoundryResourcePoolConfig CloudFoundryResourcePoolConfig

// NewCloudFoundryResourcePoolConfig instantiates a new CloudFoundryResourcePoolConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudFoundryResourcePoolConfig() *CloudFoundryResourcePoolConfig {
	this := CloudFoundryResourcePoolConfig{}
	return &this
}

// NewCloudFoundryResourcePoolConfigWithDefaults instantiates a new CloudFoundryResourcePoolConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFoundryResourcePoolConfigWithDefaults() *CloudFoundryResourcePoolConfig {
	this := CloudFoundryResourcePoolConfig{}
	return &this
}

// GetManagers returns the Managers field value if set, zero value otherwise.
func (o *CloudFoundryResourcePoolConfig) GetManagers() []string {
	if o == nil || IsNil(o.Managers) {
		var ret []string
		return ret
	}
	return o.Managers
}

// GetManagersOk returns a tuple with the Managers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudFoundryResourcePoolConfig) GetManagersOk() ([]string, bool) {
	if o == nil || IsNil(o.Managers) {
		return nil, false
	}
	return o.Managers, true
}

// IsSetManagers returns a boolean if a field has been set.
func (o *CloudFoundryResourcePoolConfig) IsSetManagers() bool {
	if o != nil && !IsNil(o.Managers) {
		return true
	}

	return false
}

// SetManagers gets a reference to the given []string and assigns it to the Managers field.
func (o *CloudFoundryResourcePoolConfig) SetManagers(v []string) {
	o.Managers = v
}

// GetDevelopers returns the Developers field value if set, zero value otherwise.
func (o *CloudFoundryResourcePoolConfig) GetDevelopers() []string {
	if o == nil || IsNil(o.Developers) {
		var ret []string
		return ret
	}
	return o.Developers
}

// GetDevelopersOk returns a tuple with the Developers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudFoundryResourcePoolConfig) GetDevelopersOk() ([]string, bool) {
	if o == nil || IsNil(o.Developers) {
		return nil, false
	}
	return o.Developers, true
}

// IsSetDevelopers returns a boolean if a field has been set.
func (o *CloudFoundryResourcePoolConfig) IsSetDevelopers() bool {
	if o != nil && !IsNil(o.Developers) {
		return true
	}

	return false
}

// SetDevelopers gets a reference to the given []string and assigns it to the Developers field.
func (o *CloudFoundryResourcePoolConfig) SetDevelopers(v []string) {
	o.Developers = v
}

// GetAuditors returns the Auditors field value if set, zero value otherwise.
func (o *CloudFoundryResourcePoolConfig) GetAuditors() []string {
	if o == nil || IsNil(o.Auditors) {
		var ret []string
		return ret
	}
	return o.Auditors
}

// GetAuditorsOk returns a tuple with the Auditors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudFoundryResourcePoolConfig) GetAuditorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Auditors) {
		return nil, false
	}
	return o.Auditors, true
}

// IsSetAuditors returns a boolean if a field has been set.
func (o *CloudFoundryResourcePoolConfig) IsSetAuditors() bool {
	if o != nil && !IsNil(o.Auditors) {
		return true
	}

	return false
}

// SetAuditors gets a reference to the given []string and assigns it to the Auditors field.
func (o *CloudFoundryResourcePoolConfig) SetAuditors(v []string) {
	o.Auditors = v
}

func (o CloudFoundryResourcePoolConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudFoundryResourcePoolConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Managers) {
		toSerialize["managers"] = o.Managers
	}
	if !IsNil(o.Developers) {
		toSerialize["developers"] = o.Developers
	}
	if !IsNil(o.Auditors) {
		toSerialize["auditors"] = o.Auditors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *CloudFoundryResourcePoolConfig) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

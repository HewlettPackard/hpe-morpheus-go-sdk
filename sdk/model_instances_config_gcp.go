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

// checks if the InstancesConfigGCP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstancesConfigGCP{}

// InstancesConfigGCP struct for InstancesConfigGCP
type InstancesConfigGCP struct {
	// Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected.
	NoAgent *bool `json:"noAgent,omitempty"`
	// External ID of the Google zone to use for instance.
	GoogleZoneId *int64 `json:"googleZoneId,omitempty"`
	// External IP Type.  `-1` for ephemeral IP.
	ExternalIpType *int64 `json:"externalIpType,omitempty"`
	// Network Tags
	NetworkTags *string `json:"networkTags,omitempty"`
	// Service Account
	ServiceAccount *string `json:"serviceAccount,omitempty"`
	// Access Scope
	AccessScope          *string                `json:"accessScope,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _InstancesConfigGCP InstancesConfigGCP

// NewInstancesConfigGCP instantiates a new InstancesConfigGCP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstancesConfigGCP() *InstancesConfigGCP {
	this := InstancesConfigGCP{}
	var noAgent bool = false
	this.NoAgent = &noAgent
	return &this
}

// NewInstancesConfigGCPWithDefaults instantiates a new InstancesConfigGCP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstancesConfigGCPWithDefaults() *InstancesConfigGCP {
	this := InstancesConfigGCP{}
	var noAgent bool = false
	this.NoAgent = &noAgent
	return &this
}

// GetNoAgent returns the NoAgent field value if set, zero value otherwise.
func (o *InstancesConfigGCP) GetNoAgent() bool {
	if o == nil || IsNil(o.NoAgent) {
		var ret bool
		return ret
	}
	return *o.NoAgent
}

// GetNoAgentOk returns a tuple with the NoAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesConfigGCP) GetNoAgentOk() (*bool, bool) {
	if o == nil || IsNil(o.NoAgent) {
		return nil, false
	}
	return o.NoAgent, true
}

// IsSetNoAgent returns a boolean if a field has been set.
func (o *InstancesConfigGCP) IsSetNoAgent() bool {
	if o != nil && !IsNil(o.NoAgent) {
		return true
	}

	return false
}

// SetNoAgent gets a reference to the given bool and assigns it to the NoAgent field.
func (o *InstancesConfigGCP) SetNoAgent(v bool) {
	o.NoAgent = &v
}

// GetGoogleZoneId returns the GoogleZoneId field value if set, zero value otherwise.
func (o *InstancesConfigGCP) GetGoogleZoneId() int64 {
	if o == nil || IsNil(o.GoogleZoneId) {
		var ret int64
		return ret
	}
	return *o.GoogleZoneId
}

// GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesConfigGCP) GetGoogleZoneIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GoogleZoneId) {
		return nil, false
	}
	return o.GoogleZoneId, true
}

// IsSetGoogleZoneId returns a boolean if a field has been set.
func (o *InstancesConfigGCP) IsSetGoogleZoneId() bool {
	if o != nil && !IsNil(o.GoogleZoneId) {
		return true
	}

	return false
}

// SetGoogleZoneId gets a reference to the given int64 and assigns it to the GoogleZoneId field.
func (o *InstancesConfigGCP) SetGoogleZoneId(v int64) {
	o.GoogleZoneId = &v
}

// GetExternalIpType returns the ExternalIpType field value if set, zero value otherwise.
func (o *InstancesConfigGCP) GetExternalIpType() int64 {
	if o == nil || IsNil(o.ExternalIpType) {
		var ret int64
		return ret
	}
	return *o.ExternalIpType
}

// GetExternalIpTypeOk returns a tuple with the ExternalIpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesConfigGCP) GetExternalIpTypeOk() (*int64, bool) {
	if o == nil || IsNil(o.ExternalIpType) {
		return nil, false
	}
	return o.ExternalIpType, true
}

// IsSetExternalIpType returns a boolean if a field has been set.
func (o *InstancesConfigGCP) IsSetExternalIpType() bool {
	if o != nil && !IsNil(o.ExternalIpType) {
		return true
	}

	return false
}

// SetExternalIpType gets a reference to the given int64 and assigns it to the ExternalIpType field.
func (o *InstancesConfigGCP) SetExternalIpType(v int64) {
	o.ExternalIpType = &v
}

// GetNetworkTags returns the NetworkTags field value if set, zero value otherwise.
func (o *InstancesConfigGCP) GetNetworkTags() string {
	if o == nil || IsNil(o.NetworkTags) {
		var ret string
		return ret
	}
	return *o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesConfigGCP) GetNetworkTagsOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkTags) {
		return nil, false
	}
	return o.NetworkTags, true
}

// IsSetNetworkTags returns a boolean if a field has been set.
func (o *InstancesConfigGCP) IsSetNetworkTags() bool {
	if o != nil && !IsNil(o.NetworkTags) {
		return true
	}

	return false
}

// SetNetworkTags gets a reference to the given string and assigns it to the NetworkTags field.
func (o *InstancesConfigGCP) SetNetworkTags(v string) {
	o.NetworkTags = &v
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise.
func (o *InstancesConfigGCP) GetServiceAccount() string {
	if o == nil || IsNil(o.ServiceAccount) {
		var ret string
		return ret
	}
	return *o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesConfigGCP) GetServiceAccountOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAccount) {
		return nil, false
	}
	return o.ServiceAccount, true
}

// IsSetServiceAccount returns a boolean if a field has been set.
func (o *InstancesConfigGCP) IsSetServiceAccount() bool {
	if o != nil && !IsNil(o.ServiceAccount) {
		return true
	}

	return false
}

// SetServiceAccount gets a reference to the given string and assigns it to the ServiceAccount field.
func (o *InstancesConfigGCP) SetServiceAccount(v string) {
	o.ServiceAccount = &v
}

// GetAccessScope returns the AccessScope field value if set, zero value otherwise.
func (o *InstancesConfigGCP) GetAccessScope() string {
	if o == nil || IsNil(o.AccessScope) {
		var ret string
		return ret
	}
	return *o.AccessScope
}

// GetAccessScopeOk returns a tuple with the AccessScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesConfigGCP) GetAccessScopeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessScope) {
		return nil, false
	}
	return o.AccessScope, true
}

// IsSetAccessScope returns a boolean if a field has been set.
func (o *InstancesConfigGCP) IsSetAccessScope() bool {
	if o != nil && !IsNil(o.AccessScope) {
		return true
	}

	return false
}

// SetAccessScope gets a reference to the given string and assigns it to the AccessScope field.
func (o *InstancesConfigGCP) SetAccessScope(v string) {
	o.AccessScope = &v
}

func (o InstancesConfigGCP) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstancesConfigGCP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NoAgent) {
		toSerialize["noAgent"] = o.NoAgent
	}
	if !IsNil(o.GoogleZoneId) {
		toSerialize["googleZoneId"] = o.GoogleZoneId
	}
	if !IsNil(o.ExternalIpType) {
		toSerialize["externalIpType"] = o.ExternalIpType
	}
	if !IsNil(o.NetworkTags) {
		toSerialize["networkTags"] = o.NetworkTags
	}
	if !IsNil(o.ServiceAccount) {
		toSerialize["serviceAccount"] = o.ServiceAccount
	}
	if !IsNil(o.AccessScope) {
		toSerialize["accessScope"] = o.AccessScope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *InstancesConfigGCP) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

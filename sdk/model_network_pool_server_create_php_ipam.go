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

// checks if the NetworkPoolServerCreatePhpIpam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkPoolServerCreatePhpIpam{}

// NetworkPoolServerCreatePhpIpam struct for NetworkPoolServerCreatePhpIpam
type NetworkPoolServerCreatePhpIpam struct {
	// Type Code (phpIPAM)
	Type string `json:"type"`
	// Name
	Name string `json:"name"`
	// Can be used to enable / disable the network pool server.
	Enabled *bool `json:"enabled,omitempty"`
	// URL
	ServiceUrl string `json:"serviceUrl"`
	// Username
	ServiceUsername *string `json:"serviceUsername,omitempty"`
	// Password
	ServicePassword *string `json:"servicePassword,omitempty"`
	// Throttle Rate
	ServiceThrottleRate *int64 `json:"serviceThrottleRate,omitempty"`
	// Disable SSL SNI Verification
	IgnoreSsl *bool `json:"ignoreSsl,omitempty"`
	// Network Filter
	NetworkFilter        *string                        `json:"networkFilter,omitempty"`
	Config               PhpIPAMNetworkPoolServerConfig `json:"config"`
	Credential           *NSXNetworkServerCredential    `json:"credential,omitempty"`
	AdditionalProperties map[string]interface{}         `json:",remain"`
}

type _NetworkPoolServerCreatePhpIpam NetworkPoolServerCreatePhpIpam

// NewNetworkPoolServerCreatePhpIpam instantiates a new NetworkPoolServerCreatePhpIpam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPoolServerCreatePhpIpam(type_ string, name string, serviceUrl string, config PhpIPAMNetworkPoolServerConfig) *NetworkPoolServerCreatePhpIpam {
	this := NetworkPoolServerCreatePhpIpam{}
	this.Type = type_
	this.Name = name
	var enabled bool = true
	this.Enabled = &enabled
	this.ServiceUrl = serviceUrl
	var serviceThrottleRate int64 = 0
	this.ServiceThrottleRate = &serviceThrottleRate
	var ignoreSsl bool = true
	this.IgnoreSsl = &ignoreSsl
	this.Config = config
	return &this
}

// NewNetworkPoolServerCreatePhpIpamWithDefaults instantiates a new NetworkPoolServerCreatePhpIpam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPoolServerCreatePhpIpamWithDefaults() *NetworkPoolServerCreatePhpIpam {
	this := NetworkPoolServerCreatePhpIpam{}
	var enabled bool = true
	this.Enabled = &enabled
	var serviceThrottleRate int64 = 0
	this.ServiceThrottleRate = &serviceThrottleRate
	var ignoreSsl bool = true
	this.IgnoreSsl = &ignoreSsl
	return &this
}

// GetType returns the Type field value
func (o *NetworkPoolServerCreatePhpIpam) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreatePhpIpam) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworkPoolServerCreatePhpIpam) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *NetworkPoolServerCreatePhpIpam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreatePhpIpam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkPoolServerCreatePhpIpam) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworkPoolServerCreatePhpIpam) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreatePhpIpam) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *NetworkPoolServerCreatePhpIpam) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworkPoolServerCreatePhpIpam) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServiceUrl returns the ServiceUrl field value
func (o *NetworkPoolServerCreatePhpIpam) GetServiceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreatePhpIpam) GetServiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUrl, true
}

// SetServiceUrl sets field value
func (o *NetworkPoolServerCreatePhpIpam) SetServiceUrl(v string) {
	o.ServiceUrl = v
}

// GetServiceUsername returns the ServiceUsername field value if set, zero value otherwise.
func (o *NetworkPoolServerCreatePhpIpam) GetServiceUsername() string {
	if o == nil || IsNil(o.ServiceUsername) {
		var ret string
		return ret
	}
	return *o.ServiceUsername
}

// GetServiceUsernameOk returns a tuple with the ServiceUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreatePhpIpam) GetServiceUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceUsername) {
		return nil, false
	}
	return o.ServiceUsername, true
}

// IsSetServiceUsername returns a boolean if a field has been set.
func (o *NetworkPoolServerCreatePhpIpam) IsSetServiceUsername() bool {
	if o != nil && !IsNil(o.ServiceUsername) {
		return true
	}

	return false
}

// SetServiceUsername gets a reference to the given string and assigns it to the ServiceUsername field.
func (o *NetworkPoolServerCreatePhpIpam) SetServiceUsername(v string) {
	o.ServiceUsername = &v
}

// GetServicePassword returns the ServicePassword field value if set, zero value otherwise.
func (o *NetworkPoolServerCreatePhpIpam) GetServicePassword() string {
	if o == nil || IsNil(o.ServicePassword) {
		var ret string
		return ret
	}
	return *o.ServicePassword
}

// GetServicePasswordOk returns a tuple with the ServicePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreatePhpIpam) GetServicePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePassword) {
		return nil, false
	}
	return o.ServicePassword, true
}

// IsSetServicePassword returns a boolean if a field has been set.
func (o *NetworkPoolServerCreatePhpIpam) IsSetServicePassword() bool {
	if o != nil && !IsNil(o.ServicePassword) {
		return true
	}

	return false
}

// SetServicePassword gets a reference to the given string and assigns it to the ServicePassword field.
func (o *NetworkPoolServerCreatePhpIpam) SetServicePassword(v string) {
	o.ServicePassword = &v
}

// GetServiceThrottleRate returns the ServiceThrottleRate field value if set, zero value otherwise.
func (o *NetworkPoolServerCreatePhpIpam) GetServiceThrottleRate() int64 {
	if o == nil || IsNil(o.ServiceThrottleRate) {
		var ret int64
		return ret
	}
	return *o.ServiceThrottleRate
}

// GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreatePhpIpam) GetServiceThrottleRateOk() (*int64, bool) {
	if o == nil || IsNil(o.ServiceThrottleRate) {
		return nil, false
	}
	return o.ServiceThrottleRate, true
}

// IsSetServiceThrottleRate returns a boolean if a field has been set.
func (o *NetworkPoolServerCreatePhpIpam) IsSetServiceThrottleRate() bool {
	if o != nil && !IsNil(o.ServiceThrottleRate) {
		return true
	}

	return false
}

// SetServiceThrottleRate gets a reference to the given int64 and assigns it to the ServiceThrottleRate field.
func (o *NetworkPoolServerCreatePhpIpam) SetServiceThrottleRate(v int64) {
	o.ServiceThrottleRate = &v
}

// GetIgnoreSsl returns the IgnoreSsl field value if set, zero value otherwise.
func (o *NetworkPoolServerCreatePhpIpam) GetIgnoreSsl() bool {
	if o == nil || IsNil(o.IgnoreSsl) {
		var ret bool
		return ret
	}
	return *o.IgnoreSsl
}

// GetIgnoreSslOk returns a tuple with the IgnoreSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreatePhpIpam) GetIgnoreSslOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreSsl) {
		return nil, false
	}
	return o.IgnoreSsl, true
}

// IsSetIgnoreSsl returns a boolean if a field has been set.
func (o *NetworkPoolServerCreatePhpIpam) IsSetIgnoreSsl() bool {
	if o != nil && !IsNil(o.IgnoreSsl) {
		return true
	}

	return false
}

// SetIgnoreSsl gets a reference to the given bool and assigns it to the IgnoreSsl field.
func (o *NetworkPoolServerCreatePhpIpam) SetIgnoreSsl(v bool) {
	o.IgnoreSsl = &v
}

// GetNetworkFilter returns the NetworkFilter field value if set, zero value otherwise.
func (o *NetworkPoolServerCreatePhpIpam) GetNetworkFilter() string {
	if o == nil || IsNil(o.NetworkFilter) {
		var ret string
		return ret
	}
	return *o.NetworkFilter
}

// GetNetworkFilterOk returns a tuple with the NetworkFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreatePhpIpam) GetNetworkFilterOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkFilter) {
		return nil, false
	}
	return o.NetworkFilter, true
}

// IsSetNetworkFilter returns a boolean if a field has been set.
func (o *NetworkPoolServerCreatePhpIpam) IsSetNetworkFilter() bool {
	if o != nil && !IsNil(o.NetworkFilter) {
		return true
	}

	return false
}

// SetNetworkFilter gets a reference to the given string and assigns it to the NetworkFilter field.
func (o *NetworkPoolServerCreatePhpIpam) SetNetworkFilter(v string) {
	o.NetworkFilter = &v
}

// GetConfig returns the Config field value
func (o *NetworkPoolServerCreatePhpIpam) GetConfig() PhpIPAMNetworkPoolServerConfig {
	if o == nil {
		var ret PhpIPAMNetworkPoolServerConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreatePhpIpam) GetConfigOk() (*PhpIPAMNetworkPoolServerConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *NetworkPoolServerCreatePhpIpam) SetConfig(v PhpIPAMNetworkPoolServerConfig) {
	o.Config = v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *NetworkPoolServerCreatePhpIpam) GetCredential() NSXNetworkServerCredential {
	if o == nil || IsNil(o.Credential) {
		var ret NSXNetworkServerCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServerCreatePhpIpam) GetCredentialOk() (*NSXNetworkServerCredential, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// IsSetCredential returns a boolean if a field has been set.
func (o *NetworkPoolServerCreatePhpIpam) IsSetCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given NSXNetworkServerCredential and assigns it to the Credential field.
func (o *NetworkPoolServerCreatePhpIpam) SetCredential(v NSXNetworkServerCredential) {
	o.Credential = &v
}

func (o NetworkPoolServerCreatePhpIpam) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkPoolServerCreatePhpIpam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["serviceUrl"] = o.ServiceUrl
	if !IsNil(o.ServiceUsername) {
		toSerialize["serviceUsername"] = o.ServiceUsername
	}
	if !IsNil(o.ServicePassword) {
		toSerialize["servicePassword"] = o.ServicePassword
	}
	if !IsNil(o.ServiceThrottleRate) {
		toSerialize["serviceThrottleRate"] = o.ServiceThrottleRate
	}
	if !IsNil(o.IgnoreSsl) {
		toSerialize["ignoreSsl"] = o.IgnoreSsl
	}
	if !IsNil(o.NetworkFilter) {
		toSerialize["networkFilter"] = o.NetworkFilter
	}
	toSerialize["config"] = o.Config
	if !IsNil(o.Credential) {
		toSerialize["credential"] = o.Credential
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *NetworkPoolServerCreatePhpIpam) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the InfobloxNetworkPoolServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfobloxNetworkPoolServer{}

// InfobloxNetworkPoolServer struct for InfobloxNetworkPoolServer
type InfobloxNetworkPoolServer struct {
	// Type Code (Infoblox)
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
	NetworkFilter *string `json:"networkFilter,omitempty"`
	// Zone Filter
	ZoneFilter *string `json:"zoneFilter,omitempty"`
	// Tenant Match
	TenantMatch *string `json:"tenantMatch,omitempty"`
	// IP Mode
	ServiceMode          *string                          `json:"serviceMode,omitempty"`
	Config               *InfobloxNetworkPoolServerConfig `json:"config,omitempty"`
	Credential           *NSXNetworkServerCredential      `json:"credential,omitempty"`
	AdditionalProperties map[string]interface{}           `json:",remain"`
}

type _InfobloxNetworkPoolServer InfobloxNetworkPoolServer

// NewInfobloxNetworkPoolServer instantiates a new InfobloxNetworkPoolServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfobloxNetworkPoolServer(type_ string, name string, serviceUrl string) *InfobloxNetworkPoolServer {
	this := InfobloxNetworkPoolServer{}
	this.Type = type_
	this.Name = name
	var enabled bool = true
	this.Enabled = &enabled
	this.ServiceUrl = serviceUrl
	var serviceThrottleRate int64 = 0
	this.ServiceThrottleRate = &serviceThrottleRate
	var ignoreSsl bool = true
	this.IgnoreSsl = &ignoreSsl
	var serviceMode string = "static"
	this.ServiceMode = &serviceMode
	return &this
}

// NewInfobloxNetworkPoolServerWithDefaults instantiates a new InfobloxNetworkPoolServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfobloxNetworkPoolServerWithDefaults() *InfobloxNetworkPoolServer {
	this := InfobloxNetworkPoolServer{}
	var enabled bool = true
	this.Enabled = &enabled
	var serviceThrottleRate int64 = 0
	this.ServiceThrottleRate = &serviceThrottleRate
	var ignoreSsl bool = true
	this.IgnoreSsl = &ignoreSsl
	var serviceMode string = "static"
	this.ServiceMode = &serviceMode
	return &this
}

// GetType returns the Type field value
func (o *InfobloxNetworkPoolServer) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InfobloxNetworkPoolServer) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *InfobloxNetworkPoolServer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InfobloxNetworkPoolServer) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InfobloxNetworkPoolServer) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *InfobloxNetworkPoolServer) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InfobloxNetworkPoolServer) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServiceUrl returns the ServiceUrl field value
func (o *InfobloxNetworkPoolServer) GetServiceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetServiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUrl, true
}

// SetServiceUrl sets field value
func (o *InfobloxNetworkPoolServer) SetServiceUrl(v string) {
	o.ServiceUrl = v
}

// GetServiceUsername returns the ServiceUsername field value if set, zero value otherwise.
func (o *InfobloxNetworkPoolServer) GetServiceUsername() string {
	if o == nil || IsNil(o.ServiceUsername) {
		var ret string
		return ret
	}
	return *o.ServiceUsername
}

// GetServiceUsernameOk returns a tuple with the ServiceUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetServiceUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceUsername) {
		return nil, false
	}
	return o.ServiceUsername, true
}

// IsSetServiceUsername returns a boolean if a field has been set.
func (o *InfobloxNetworkPoolServer) IsSetServiceUsername() bool {
	if o != nil && !IsNil(o.ServiceUsername) {
		return true
	}

	return false
}

// SetServiceUsername gets a reference to the given string and assigns it to the ServiceUsername field.
func (o *InfobloxNetworkPoolServer) SetServiceUsername(v string) {
	o.ServiceUsername = &v
}

// GetServicePassword returns the ServicePassword field value if set, zero value otherwise.
func (o *InfobloxNetworkPoolServer) GetServicePassword() string {
	if o == nil || IsNil(o.ServicePassword) {
		var ret string
		return ret
	}
	return *o.ServicePassword
}

// GetServicePasswordOk returns a tuple with the ServicePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetServicePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePassword) {
		return nil, false
	}
	return o.ServicePassword, true
}

// IsSetServicePassword returns a boolean if a field has been set.
func (o *InfobloxNetworkPoolServer) IsSetServicePassword() bool {
	if o != nil && !IsNil(o.ServicePassword) {
		return true
	}

	return false
}

// SetServicePassword gets a reference to the given string and assigns it to the ServicePassword field.
func (o *InfobloxNetworkPoolServer) SetServicePassword(v string) {
	o.ServicePassword = &v
}

// GetServiceThrottleRate returns the ServiceThrottleRate field value if set, zero value otherwise.
func (o *InfobloxNetworkPoolServer) GetServiceThrottleRate() int64 {
	if o == nil || IsNil(o.ServiceThrottleRate) {
		var ret int64
		return ret
	}
	return *o.ServiceThrottleRate
}

// GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetServiceThrottleRateOk() (*int64, bool) {
	if o == nil || IsNil(o.ServiceThrottleRate) {
		return nil, false
	}
	return o.ServiceThrottleRate, true
}

// IsSetServiceThrottleRate returns a boolean if a field has been set.
func (o *InfobloxNetworkPoolServer) IsSetServiceThrottleRate() bool {
	if o != nil && !IsNil(o.ServiceThrottleRate) {
		return true
	}

	return false
}

// SetServiceThrottleRate gets a reference to the given int64 and assigns it to the ServiceThrottleRate field.
func (o *InfobloxNetworkPoolServer) SetServiceThrottleRate(v int64) {
	o.ServiceThrottleRate = &v
}

// GetIgnoreSsl returns the IgnoreSsl field value if set, zero value otherwise.
func (o *InfobloxNetworkPoolServer) GetIgnoreSsl() bool {
	if o == nil || IsNil(o.IgnoreSsl) {
		var ret bool
		return ret
	}
	return *o.IgnoreSsl
}

// GetIgnoreSslOk returns a tuple with the IgnoreSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetIgnoreSslOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreSsl) {
		return nil, false
	}
	return o.IgnoreSsl, true
}

// IsSetIgnoreSsl returns a boolean if a field has been set.
func (o *InfobloxNetworkPoolServer) IsSetIgnoreSsl() bool {
	if o != nil && !IsNil(o.IgnoreSsl) {
		return true
	}

	return false
}

// SetIgnoreSsl gets a reference to the given bool and assigns it to the IgnoreSsl field.
func (o *InfobloxNetworkPoolServer) SetIgnoreSsl(v bool) {
	o.IgnoreSsl = &v
}

// GetNetworkFilter returns the NetworkFilter field value if set, zero value otherwise.
func (o *InfobloxNetworkPoolServer) GetNetworkFilter() string {
	if o == nil || IsNil(o.NetworkFilter) {
		var ret string
		return ret
	}
	return *o.NetworkFilter
}

// GetNetworkFilterOk returns a tuple with the NetworkFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetNetworkFilterOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkFilter) {
		return nil, false
	}
	return o.NetworkFilter, true
}

// IsSetNetworkFilter returns a boolean if a field has been set.
func (o *InfobloxNetworkPoolServer) IsSetNetworkFilter() bool {
	if o != nil && !IsNil(o.NetworkFilter) {
		return true
	}

	return false
}

// SetNetworkFilter gets a reference to the given string and assigns it to the NetworkFilter field.
func (o *InfobloxNetworkPoolServer) SetNetworkFilter(v string) {
	o.NetworkFilter = &v
}

// GetZoneFilter returns the ZoneFilter field value if set, zero value otherwise.
func (o *InfobloxNetworkPoolServer) GetZoneFilter() string {
	if o == nil || IsNil(o.ZoneFilter) {
		var ret string
		return ret
	}
	return *o.ZoneFilter
}

// GetZoneFilterOk returns a tuple with the ZoneFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetZoneFilterOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneFilter) {
		return nil, false
	}
	return o.ZoneFilter, true
}

// IsSetZoneFilter returns a boolean if a field has been set.
func (o *InfobloxNetworkPoolServer) IsSetZoneFilter() bool {
	if o != nil && !IsNil(o.ZoneFilter) {
		return true
	}

	return false
}

// SetZoneFilter gets a reference to the given string and assigns it to the ZoneFilter field.
func (o *InfobloxNetworkPoolServer) SetZoneFilter(v string) {
	o.ZoneFilter = &v
}

// GetTenantMatch returns the TenantMatch field value if set, zero value otherwise.
func (o *InfobloxNetworkPoolServer) GetTenantMatch() string {
	if o == nil || IsNil(o.TenantMatch) {
		var ret string
		return ret
	}
	return *o.TenantMatch
}

// GetTenantMatchOk returns a tuple with the TenantMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetTenantMatchOk() (*string, bool) {
	if o == nil || IsNil(o.TenantMatch) {
		return nil, false
	}
	return o.TenantMatch, true
}

// IsSetTenantMatch returns a boolean if a field has been set.
func (o *InfobloxNetworkPoolServer) IsSetTenantMatch() bool {
	if o != nil && !IsNil(o.TenantMatch) {
		return true
	}

	return false
}

// SetTenantMatch gets a reference to the given string and assigns it to the TenantMatch field.
func (o *InfobloxNetworkPoolServer) SetTenantMatch(v string) {
	o.TenantMatch = &v
}

// GetServiceMode returns the ServiceMode field value if set, zero value otherwise.
func (o *InfobloxNetworkPoolServer) GetServiceMode() string {
	if o == nil || IsNil(o.ServiceMode) {
		var ret string
		return ret
	}
	return *o.ServiceMode
}

// GetServiceModeOk returns a tuple with the ServiceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetServiceModeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceMode) {
		return nil, false
	}
	return o.ServiceMode, true
}

// IsSetServiceMode returns a boolean if a field has been set.
func (o *InfobloxNetworkPoolServer) IsSetServiceMode() bool {
	if o != nil && !IsNil(o.ServiceMode) {
		return true
	}

	return false
}

// SetServiceMode gets a reference to the given string and assigns it to the ServiceMode field.
func (o *InfobloxNetworkPoolServer) SetServiceMode(v string) {
	o.ServiceMode = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *InfobloxNetworkPoolServer) GetConfig() InfobloxNetworkPoolServerConfig {
	if o == nil || IsNil(o.Config) {
		var ret InfobloxNetworkPoolServerConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetConfigOk() (*InfobloxNetworkPoolServerConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *InfobloxNetworkPoolServer) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given InfobloxNetworkPoolServerConfig and assigns it to the Config field.
func (o *InfobloxNetworkPoolServer) SetConfig(v InfobloxNetworkPoolServerConfig) {
	o.Config = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *InfobloxNetworkPoolServer) GetCredential() NSXNetworkServerCredential {
	if o == nil || IsNil(o.Credential) {
		var ret NSXNetworkServerCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfobloxNetworkPoolServer) GetCredentialOk() (*NSXNetworkServerCredential, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// IsSetCredential returns a boolean if a field has been set.
func (o *InfobloxNetworkPoolServer) IsSetCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given NSXNetworkServerCredential and assigns it to the Credential field.
func (o *InfobloxNetworkPoolServer) SetCredential(v NSXNetworkServerCredential) {
	o.Credential = &v
}

func (o InfobloxNetworkPoolServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfobloxNetworkPoolServer) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ZoneFilter) {
		toSerialize["zoneFilter"] = o.ZoneFilter
	}
	if !IsNil(o.TenantMatch) {
		toSerialize["tenantMatch"] = o.TenantMatch
	}
	if !IsNil(o.ServiceMode) {
		toSerialize["serviceMode"] = o.ServiceMode
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Credential) {
		toSerialize["credential"] = o.Credential
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *InfobloxNetworkPoolServer) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

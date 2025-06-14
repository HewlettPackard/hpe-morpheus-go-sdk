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

// checks if the AddIntegrationsRequestOneOf1Integration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddIntegrationsRequestOneOf1Integration{}

// AddIntegrationsRequestOneOf1Integration struct for AddIntegrationsRequestOneOf1Integration
type AddIntegrationsRequestOneOf1Integration struct {
	// Name, a unique identifier for the integration
	Name string `json:"name"`
	// Integration Type Code
	Type string `json:"type"`
	// Set `true` to enable integration
	Enabled *bool `json:"enabled,omitempty"`
	// Pass `false` to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.
	Refresh *bool `json:"refresh,omitempty"`
	// Ansible Git URL
	ServiceUrl string `json:"serviceUrl"`
	// Git Username
	ServiceUsername *string `json:"serviceUsername,omitempty"`
	// Git Password or Token depending on the Git host
	ServicePassword *string `json:"servicePassword,omitempty"`
	// Git Token
	ServiceToken *string `json:"serviceToken,omitempty"`
	// Keypair ID
	ServiceKey           *int64                                         `json:"serviceKey,omitempty"`
	Config               *AddIntegrationsRequestOneOf1IntegrationConfig `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}                         `json:",remain"`
}

type _AddIntegrationsRequestOneOf1Integration AddIntegrationsRequestOneOf1Integration

// NewAddIntegrationsRequestOneOf1Integration instantiates a new AddIntegrationsRequestOneOf1Integration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIntegrationsRequestOneOf1Integration(name string, type_ string, serviceUrl string) *AddIntegrationsRequestOneOf1Integration {
	this := AddIntegrationsRequestOneOf1Integration{}
	this.Name = name
	this.Type = type_
	var refresh bool = true
	this.Refresh = &refresh
	this.ServiceUrl = serviceUrl
	return &this
}

// NewAddIntegrationsRequestOneOf1IntegrationWithDefaults instantiates a new AddIntegrationsRequestOneOf1Integration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIntegrationsRequestOneOf1IntegrationWithDefaults() *AddIntegrationsRequestOneOf1Integration {
	this := AddIntegrationsRequestOneOf1Integration{}
	var refresh bool = true
	this.Refresh = &refresh
	return &this
}

// GetName returns the Name field value
func (o *AddIntegrationsRequestOneOf1Integration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddIntegrationsRequestOneOf1Integration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddIntegrationsRequestOneOf1Integration) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *AddIntegrationsRequestOneOf1Integration) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddIntegrationsRequestOneOf1Integration) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddIntegrationsRequestOneOf1Integration) SetType(v string) {
	o.Type = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AddIntegrationsRequestOneOf1Integration) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIntegrationsRequestOneOf1Integration) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *AddIntegrationsRequestOneOf1Integration) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AddIntegrationsRequestOneOf1Integration) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRefresh returns the Refresh field value if set, zero value otherwise.
func (o *AddIntegrationsRequestOneOf1Integration) GetRefresh() bool {
	if o == nil || IsNil(o.Refresh) {
		var ret bool
		return ret
	}
	return *o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIntegrationsRequestOneOf1Integration) GetRefreshOk() (*bool, bool) {
	if o == nil || IsNil(o.Refresh) {
		return nil, false
	}
	return o.Refresh, true
}

// IsSetRefresh returns a boolean if a field has been set.
func (o *AddIntegrationsRequestOneOf1Integration) IsSetRefresh() bool {
	if o != nil && !IsNil(o.Refresh) {
		return true
	}

	return false
}

// SetRefresh gets a reference to the given bool and assigns it to the Refresh field.
func (o *AddIntegrationsRequestOneOf1Integration) SetRefresh(v bool) {
	o.Refresh = &v
}

// GetServiceUrl returns the ServiceUrl field value
func (o *AddIntegrationsRequestOneOf1Integration) GetServiceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value
// and a boolean to check if the value has been set.
func (o *AddIntegrationsRequestOneOf1Integration) GetServiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUrl, true
}

// SetServiceUrl sets field value
func (o *AddIntegrationsRequestOneOf1Integration) SetServiceUrl(v string) {
	o.ServiceUrl = v
}

// GetServiceUsername returns the ServiceUsername field value if set, zero value otherwise.
func (o *AddIntegrationsRequestOneOf1Integration) GetServiceUsername() string {
	if o == nil || IsNil(o.ServiceUsername) {
		var ret string
		return ret
	}
	return *o.ServiceUsername
}

// GetServiceUsernameOk returns a tuple with the ServiceUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIntegrationsRequestOneOf1Integration) GetServiceUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceUsername) {
		return nil, false
	}
	return o.ServiceUsername, true
}

// IsSetServiceUsername returns a boolean if a field has been set.
func (o *AddIntegrationsRequestOneOf1Integration) IsSetServiceUsername() bool {
	if o != nil && !IsNil(o.ServiceUsername) {
		return true
	}

	return false
}

// SetServiceUsername gets a reference to the given string and assigns it to the ServiceUsername field.
func (o *AddIntegrationsRequestOneOf1Integration) SetServiceUsername(v string) {
	o.ServiceUsername = &v
}

// GetServicePassword returns the ServicePassword field value if set, zero value otherwise.
func (o *AddIntegrationsRequestOneOf1Integration) GetServicePassword() string {
	if o == nil || IsNil(o.ServicePassword) {
		var ret string
		return ret
	}
	return *o.ServicePassword
}

// GetServicePasswordOk returns a tuple with the ServicePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIntegrationsRequestOneOf1Integration) GetServicePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePassword) {
		return nil, false
	}
	return o.ServicePassword, true
}

// IsSetServicePassword returns a boolean if a field has been set.
func (o *AddIntegrationsRequestOneOf1Integration) IsSetServicePassword() bool {
	if o != nil && !IsNil(o.ServicePassword) {
		return true
	}

	return false
}

// SetServicePassword gets a reference to the given string and assigns it to the ServicePassword field.
func (o *AddIntegrationsRequestOneOf1Integration) SetServicePassword(v string) {
	o.ServicePassword = &v
}

// GetServiceToken returns the ServiceToken field value if set, zero value otherwise.
func (o *AddIntegrationsRequestOneOf1Integration) GetServiceToken() string {
	if o == nil || IsNil(o.ServiceToken) {
		var ret string
		return ret
	}
	return *o.ServiceToken
}

// GetServiceTokenOk returns a tuple with the ServiceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIntegrationsRequestOneOf1Integration) GetServiceTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceToken) {
		return nil, false
	}
	return o.ServiceToken, true
}

// IsSetServiceToken returns a boolean if a field has been set.
func (o *AddIntegrationsRequestOneOf1Integration) IsSetServiceToken() bool {
	if o != nil && !IsNil(o.ServiceToken) {
		return true
	}

	return false
}

// SetServiceToken gets a reference to the given string and assigns it to the ServiceToken field.
func (o *AddIntegrationsRequestOneOf1Integration) SetServiceToken(v string) {
	o.ServiceToken = &v
}

// GetServiceKey returns the ServiceKey field value if set, zero value otherwise.
func (o *AddIntegrationsRequestOneOf1Integration) GetServiceKey() int64 {
	if o == nil || IsNil(o.ServiceKey) {
		var ret int64
		return ret
	}
	return *o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIntegrationsRequestOneOf1Integration) GetServiceKeyOk() (*int64, bool) {
	if o == nil || IsNil(o.ServiceKey) {
		return nil, false
	}
	return o.ServiceKey, true
}

// IsSetServiceKey returns a boolean if a field has been set.
func (o *AddIntegrationsRequestOneOf1Integration) IsSetServiceKey() bool {
	if o != nil && !IsNil(o.ServiceKey) {
		return true
	}

	return false
}

// SetServiceKey gets a reference to the given int64 and assigns it to the ServiceKey field.
func (o *AddIntegrationsRequestOneOf1Integration) SetServiceKey(v int64) {
	o.ServiceKey = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AddIntegrationsRequestOneOf1Integration) GetConfig() AddIntegrationsRequestOneOf1IntegrationConfig {
	if o == nil || IsNil(o.Config) {
		var ret AddIntegrationsRequestOneOf1IntegrationConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIntegrationsRequestOneOf1Integration) GetConfigOk() (*AddIntegrationsRequestOneOf1IntegrationConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *AddIntegrationsRequestOneOf1Integration) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given AddIntegrationsRequestOneOf1IntegrationConfig and assigns it to the Config field.
func (o *AddIntegrationsRequestOneOf1Integration) SetConfig(v AddIntegrationsRequestOneOf1IntegrationConfig) {
	o.Config = &v
}

func (o AddIntegrationsRequestOneOf1Integration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddIntegrationsRequestOneOf1Integration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Refresh) {
		toSerialize["refresh"] = o.Refresh
	}
	toSerialize["serviceUrl"] = o.ServiceUrl
	if !IsNil(o.ServiceUsername) {
		toSerialize["serviceUsername"] = o.ServiceUsername
	}
	if !IsNil(o.ServicePassword) {
		toSerialize["servicePassword"] = o.ServicePassword
	}
	if !IsNil(o.ServiceToken) {
		toSerialize["serviceToken"] = o.ServiceToken
	}
	if !IsNil(o.ServiceKey) {
		toSerialize["serviceKey"] = o.ServiceKey
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddIntegrationsRequestOneOf1Integration) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

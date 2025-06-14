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

// checks if the CheckElasticsearchConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckElasticsearchConfig{}

// CheckElasticsearchConfig struct for CheckElasticsearchConfig
type CheckElasticsearchConfig struct {
	// Hostname or IP address of the Elasticsearch server
	EsHost string `json:"esHost"`
	// Port to connect to the HTTP service
	EsPort            string  `json:"esPort"`
	CheckUser         *string `json:"checkUser,omitempty"`
	TextCheckOn       *string `json:"textCheckOn,omitempty"`
	CheckPassword     *string `json:"checkPassword,omitempty"`
	WebTextMatch      *string `json:"webTextMatch,omitempty"`
	CheckPasswordHash *string `json:"checkPasswordHash,omitempty"`
	// Set to on to turn on tunneling
	TunnelOn *string `json:"tunnelOn,omitempty"`
	// Hostname or IP address of the proxy host
	SshHost *string `json:"sshHost,omitempty"`
	// Port for SSH on the proxy host, defaults to 22
	SshPort *int64 `json:"sshPort,omitempty"`
	// SSH user on the proxy host to login as
	SshUser *string `json:"sshUser,omitempty"`
	// Password for user, if not using key based authentication
	SshPassword          *string                `json:"sshPassword,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _CheckElasticsearchConfig CheckElasticsearchConfig

// NewCheckElasticsearchConfig instantiates a new CheckElasticsearchConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckElasticsearchConfig(esHost string, esPort string) *CheckElasticsearchConfig {
	this := CheckElasticsearchConfig{}
	this.EsHost = esHost
	this.EsPort = esPort
	var tunnelOn string = "off"
	this.TunnelOn = &tunnelOn
	return &this
}

// NewCheckElasticsearchConfigWithDefaults instantiates a new CheckElasticsearchConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckElasticsearchConfigWithDefaults() *CheckElasticsearchConfig {
	this := CheckElasticsearchConfig{}
	var tunnelOn string = "off"
	this.TunnelOn = &tunnelOn
	return &this
}

// GetEsHost returns the EsHost field value
func (o *CheckElasticsearchConfig) GetEsHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EsHost
}

// GetEsHostOk returns a tuple with the EsHost field value
// and a boolean to check if the value has been set.
func (o *CheckElasticsearchConfig) GetEsHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EsHost, true
}

// SetEsHost sets field value
func (o *CheckElasticsearchConfig) SetEsHost(v string) {
	o.EsHost = v
}

// GetEsPort returns the EsPort field value
func (o *CheckElasticsearchConfig) GetEsPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EsPort
}

// GetEsPortOk returns a tuple with the EsPort field value
// and a boolean to check if the value has been set.
func (o *CheckElasticsearchConfig) GetEsPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EsPort, true
}

// SetEsPort sets field value
func (o *CheckElasticsearchConfig) SetEsPort(v string) {
	o.EsPort = v
}

// GetCheckUser returns the CheckUser field value if set, zero value otherwise.
func (o *CheckElasticsearchConfig) GetCheckUser() string {
	if o == nil || IsNil(o.CheckUser) {
		var ret string
		return ret
	}
	return *o.CheckUser
}

// GetCheckUserOk returns a tuple with the CheckUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElasticsearchConfig) GetCheckUserOk() (*string, bool) {
	if o == nil || IsNil(o.CheckUser) {
		return nil, false
	}
	return o.CheckUser, true
}

// IsSetCheckUser returns a boolean if a field has been set.
func (o *CheckElasticsearchConfig) IsSetCheckUser() bool {
	if o != nil && !IsNil(o.CheckUser) {
		return true
	}

	return false
}

// SetCheckUser gets a reference to the given string and assigns it to the CheckUser field.
func (o *CheckElasticsearchConfig) SetCheckUser(v string) {
	o.CheckUser = &v
}

// GetTextCheckOn returns the TextCheckOn field value if set, zero value otherwise.
func (o *CheckElasticsearchConfig) GetTextCheckOn() string {
	if o == nil || IsNil(o.TextCheckOn) {
		var ret string
		return ret
	}
	return *o.TextCheckOn
}

// GetTextCheckOnOk returns a tuple with the TextCheckOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElasticsearchConfig) GetTextCheckOnOk() (*string, bool) {
	if o == nil || IsNil(o.TextCheckOn) {
		return nil, false
	}
	return o.TextCheckOn, true
}

// IsSetTextCheckOn returns a boolean if a field has been set.
func (o *CheckElasticsearchConfig) IsSetTextCheckOn() bool {
	if o != nil && !IsNil(o.TextCheckOn) {
		return true
	}

	return false
}

// SetTextCheckOn gets a reference to the given string and assigns it to the TextCheckOn field.
func (o *CheckElasticsearchConfig) SetTextCheckOn(v string) {
	o.TextCheckOn = &v
}

// GetCheckPassword returns the CheckPassword field value if set, zero value otherwise.
func (o *CheckElasticsearchConfig) GetCheckPassword() string {
	if o == nil || IsNil(o.CheckPassword) {
		var ret string
		return ret
	}
	return *o.CheckPassword
}

// GetCheckPasswordOk returns a tuple with the CheckPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElasticsearchConfig) GetCheckPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.CheckPassword) {
		return nil, false
	}
	return o.CheckPassword, true
}

// IsSetCheckPassword returns a boolean if a field has been set.
func (o *CheckElasticsearchConfig) IsSetCheckPassword() bool {
	if o != nil && !IsNil(o.CheckPassword) {
		return true
	}

	return false
}

// SetCheckPassword gets a reference to the given string and assigns it to the CheckPassword field.
func (o *CheckElasticsearchConfig) SetCheckPassword(v string) {
	o.CheckPassword = &v
}

// GetWebTextMatch returns the WebTextMatch field value if set, zero value otherwise.
func (o *CheckElasticsearchConfig) GetWebTextMatch() string {
	if o == nil || IsNil(o.WebTextMatch) {
		var ret string
		return ret
	}
	return *o.WebTextMatch
}

// GetWebTextMatchOk returns a tuple with the WebTextMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElasticsearchConfig) GetWebTextMatchOk() (*string, bool) {
	if o == nil || IsNil(o.WebTextMatch) {
		return nil, false
	}
	return o.WebTextMatch, true
}

// IsSetWebTextMatch returns a boolean if a field has been set.
func (o *CheckElasticsearchConfig) IsSetWebTextMatch() bool {
	if o != nil && !IsNil(o.WebTextMatch) {
		return true
	}

	return false
}

// SetWebTextMatch gets a reference to the given string and assigns it to the WebTextMatch field.
func (o *CheckElasticsearchConfig) SetWebTextMatch(v string) {
	o.WebTextMatch = &v
}

// GetCheckPasswordHash returns the CheckPasswordHash field value if set, zero value otherwise.
func (o *CheckElasticsearchConfig) GetCheckPasswordHash() string {
	if o == nil || IsNil(o.CheckPasswordHash) {
		var ret string
		return ret
	}
	return *o.CheckPasswordHash
}

// GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElasticsearchConfig) GetCheckPasswordHashOk() (*string, bool) {
	if o == nil || IsNil(o.CheckPasswordHash) {
		return nil, false
	}
	return o.CheckPasswordHash, true
}

// IsSetCheckPasswordHash returns a boolean if a field has been set.
func (o *CheckElasticsearchConfig) IsSetCheckPasswordHash() bool {
	if o != nil && !IsNil(o.CheckPasswordHash) {
		return true
	}

	return false
}

// SetCheckPasswordHash gets a reference to the given string and assigns it to the CheckPasswordHash field.
func (o *CheckElasticsearchConfig) SetCheckPasswordHash(v string) {
	o.CheckPasswordHash = &v
}

// GetTunnelOn returns the TunnelOn field value if set, zero value otherwise.
func (o *CheckElasticsearchConfig) GetTunnelOn() string {
	if o == nil || IsNil(o.TunnelOn) {
		var ret string
		return ret
	}
	return *o.TunnelOn
}

// GetTunnelOnOk returns a tuple with the TunnelOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElasticsearchConfig) GetTunnelOnOk() (*string, bool) {
	if o == nil || IsNil(o.TunnelOn) {
		return nil, false
	}
	return o.TunnelOn, true
}

// IsSetTunnelOn returns a boolean if a field has been set.
func (o *CheckElasticsearchConfig) IsSetTunnelOn() bool {
	if o != nil && !IsNil(o.TunnelOn) {
		return true
	}

	return false
}

// SetTunnelOn gets a reference to the given string and assigns it to the TunnelOn field.
func (o *CheckElasticsearchConfig) SetTunnelOn(v string) {
	o.TunnelOn = &v
}

// GetSshHost returns the SshHost field value if set, zero value otherwise.
func (o *CheckElasticsearchConfig) GetSshHost() string {
	if o == nil || IsNil(o.SshHost) {
		var ret string
		return ret
	}
	return *o.SshHost
}

// GetSshHostOk returns a tuple with the SshHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElasticsearchConfig) GetSshHostOk() (*string, bool) {
	if o == nil || IsNil(o.SshHost) {
		return nil, false
	}
	return o.SshHost, true
}

// IsSetSshHost returns a boolean if a field has been set.
func (o *CheckElasticsearchConfig) IsSetSshHost() bool {
	if o != nil && !IsNil(o.SshHost) {
		return true
	}

	return false
}

// SetSshHost gets a reference to the given string and assigns it to the SshHost field.
func (o *CheckElasticsearchConfig) SetSshHost(v string) {
	o.SshHost = &v
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *CheckElasticsearchConfig) GetSshPort() int64 {
	if o == nil || IsNil(o.SshPort) {
		var ret int64
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElasticsearchConfig) GetSshPortOk() (*int64, bool) {
	if o == nil || IsNil(o.SshPort) {
		return nil, false
	}
	return o.SshPort, true
}

// IsSetSshPort returns a boolean if a field has been set.
func (o *CheckElasticsearchConfig) IsSetSshPort() bool {
	if o != nil && !IsNil(o.SshPort) {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given int64 and assigns it to the SshPort field.
func (o *CheckElasticsearchConfig) SetSshPort(v int64) {
	o.SshPort = &v
}

// GetSshUser returns the SshUser field value if set, zero value otherwise.
func (o *CheckElasticsearchConfig) GetSshUser() string {
	if o == nil || IsNil(o.SshUser) {
		var ret string
		return ret
	}
	return *o.SshUser
}

// GetSshUserOk returns a tuple with the SshUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElasticsearchConfig) GetSshUserOk() (*string, bool) {
	if o == nil || IsNil(o.SshUser) {
		return nil, false
	}
	return o.SshUser, true
}

// IsSetSshUser returns a boolean if a field has been set.
func (o *CheckElasticsearchConfig) IsSetSshUser() bool {
	if o != nil && !IsNil(o.SshUser) {
		return true
	}

	return false
}

// SetSshUser gets a reference to the given string and assigns it to the SshUser field.
func (o *CheckElasticsearchConfig) SetSshUser(v string) {
	o.SshUser = &v
}

// GetSshPassword returns the SshPassword field value if set, zero value otherwise.
func (o *CheckElasticsearchConfig) GetSshPassword() string {
	if o == nil || IsNil(o.SshPassword) {
		var ret string
		return ret
	}
	return *o.SshPassword
}

// GetSshPasswordOk returns a tuple with the SshPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElasticsearchConfig) GetSshPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.SshPassword) {
		return nil, false
	}
	return o.SshPassword, true
}

// IsSetSshPassword returns a boolean if a field has been set.
func (o *CheckElasticsearchConfig) IsSetSshPassword() bool {
	if o != nil && !IsNil(o.SshPassword) {
		return true
	}

	return false
}

// SetSshPassword gets a reference to the given string and assigns it to the SshPassword field.
func (o *CheckElasticsearchConfig) SetSshPassword(v string) {
	o.SshPassword = &v
}

func (o CheckElasticsearchConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckElasticsearchConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["esHost"] = o.EsHost
	toSerialize["esPort"] = o.EsPort
	if !IsNil(o.CheckUser) {
		toSerialize["checkUser"] = o.CheckUser
	}
	if !IsNil(o.TextCheckOn) {
		toSerialize["textCheckOn"] = o.TextCheckOn
	}
	if !IsNil(o.CheckPassword) {
		toSerialize["checkPassword"] = o.CheckPassword
	}
	if !IsNil(o.WebTextMatch) {
		toSerialize["webTextMatch"] = o.WebTextMatch
	}
	if !IsNil(o.CheckPasswordHash) {
		toSerialize["checkPasswordHash"] = o.CheckPasswordHash
	}
	if !IsNil(o.TunnelOn) {
		toSerialize["tunnelOn"] = o.TunnelOn
	}
	if !IsNil(o.SshHost) {
		toSerialize["sshHost"] = o.SshHost
	}
	if !IsNil(o.SshPort) {
		toSerialize["sshPort"] = o.SshPort
	}
	if !IsNil(o.SshUser) {
		toSerialize["sshUser"] = o.SshUser
	}
	if !IsNil(o.SshPassword) {
		toSerialize["sshPassword"] = o.SshPassword
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *CheckElasticsearchConfig) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

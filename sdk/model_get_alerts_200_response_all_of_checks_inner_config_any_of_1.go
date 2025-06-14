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

// checks if the GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1{}

// GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1 struct for GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1
type GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1 struct {
	// Hostname or IP address of the database
	DbHost string `json:"dbHost"`
	// Database Port (defaults to default port of DB type selected)
	DbPort string `json:"dbPort"`
	// Database username
	DbUser string `json:"dbUser"`
	// Database password, (all check data is encrypted inside the database)
	DbPassword string `json:"dbPassword"`
	// Database password hash
	DbPasswordHash *string `json:"dbPasswordHash,omitempty"`
	// Database name you would like to connect to
	DbName string `json:"dbName"`
	// Query to test
	DbQuery string `json:"dbQuery"`
	// Can be set to `lt` (less than), `gt` (greater than), `equal` (Equal to) for comparison
	CheckOperator     *string  `json:"checkOperator,omitempty"`
	CheckResult       *float32 `json:"checkResult,omitempty"`
	CheckUser         *string  `json:"checkUser,omitempty"`
	TextCheckOn       *string  `json:"textCheckOn,omitempty"`
	CheckPassword     *string  `json:"checkPassword,omitempty"`
	WebTextMatch      *string  `json:"webTextMatch,omitempty"`
	CheckPasswordHash *string  `json:"checkPasswordHash,omitempty"`
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

type _GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1 GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1

// NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf1 instantiates a new GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf1(dbHost string, dbPort string, dbUser string, dbPassword string, dbName string, dbQuery string) *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1 {
	this := GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1{}
	this.DbHost = dbHost
	this.DbPort = dbPort
	this.DbUser = dbUser
	this.DbPassword = dbPassword
	this.DbName = dbName
	this.DbQuery = dbQuery
	var tunnelOn string = "off"
	this.TunnelOn = &tunnelOn
	return &this
}

// NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf1WithDefaults instantiates a new GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAlerts200ResponseAllOfChecksInnerConfigAnyOf1WithDefaults() *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1 {
	this := GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1{}
	var tunnelOn string = "off"
	this.TunnelOn = &tunnelOn
	return &this
}

// GetDbHost returns the DbHost field value
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbHost
}

// GetDbHostOk returns a tuple with the DbHost field value
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbHost, true
}

// SetDbHost sets field value
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbHost(v string) {
	o.DbHost = v
}

// GetDbPort returns the DbPort field value
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbPort
}

// GetDbPortOk returns a tuple with the DbPort field value
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbPort, true
}

// SetDbPort sets field value
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbPort(v string) {
	o.DbPort = v
}

// GetDbUser returns the DbUser field value
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbUser
}

// GetDbUserOk returns a tuple with the DbUser field value
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbUser, true
}

// SetDbUser sets field value
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbUser(v string) {
	o.DbUser = v
}

// GetDbPassword returns the DbPassword field value
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbPassword
}

// GetDbPasswordOk returns a tuple with the DbPassword field value
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbPassword, true
}

// SetDbPassword sets field value
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbPassword(v string) {
	o.DbPassword = v
}

// GetDbPasswordHash returns the DbPasswordHash field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbPasswordHash() string {
	if o == nil || IsNil(o.DbPasswordHash) {
		var ret string
		return ret
	}
	return *o.DbPasswordHash
}

// GetDbPasswordHashOk returns a tuple with the DbPasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbPasswordHashOk() (*string, bool) {
	if o == nil || IsNil(o.DbPasswordHash) {
		return nil, false
	}
	return o.DbPasswordHash, true
}

// IsSetDbPasswordHash returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetDbPasswordHash() bool {
	if o != nil && !IsNil(o.DbPasswordHash) {
		return true
	}

	return false
}

// SetDbPasswordHash gets a reference to the given string and assigns it to the DbPasswordHash field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbPasswordHash(v string) {
	o.DbPasswordHash = &v
}

// GetDbName returns the DbName field value
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbName, true
}

// SetDbName sets field value
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbName(v string) {
	o.DbName = v
}

// GetDbQuery returns the DbQuery field value
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbQuery
}

// GetDbQueryOk returns a tuple with the DbQuery field value
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetDbQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbQuery, true
}

// SetDbQuery sets field value
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetDbQuery(v string) {
	o.DbQuery = v
}

// GetCheckOperator returns the CheckOperator field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckOperator() string {
	if o == nil || IsNil(o.CheckOperator) {
		var ret string
		return ret
	}
	return *o.CheckOperator
}

// GetCheckOperatorOk returns a tuple with the CheckOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.CheckOperator) {
		return nil, false
	}
	return o.CheckOperator, true
}

// IsSetCheckOperator returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetCheckOperator() bool {
	if o != nil && !IsNil(o.CheckOperator) {
		return true
	}

	return false
}

// SetCheckOperator gets a reference to the given string and assigns it to the CheckOperator field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetCheckOperator(v string) {
	o.CheckOperator = &v
}

// GetCheckResult returns the CheckResult field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckResult() float32 {
	if o == nil || IsNil(o.CheckResult) {
		var ret float32
		return ret
	}
	return *o.CheckResult
}

// GetCheckResultOk returns a tuple with the CheckResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckResultOk() (*float32, bool) {
	if o == nil || IsNil(o.CheckResult) {
		return nil, false
	}
	return o.CheckResult, true
}

// IsSetCheckResult returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetCheckResult() bool {
	if o != nil && !IsNil(o.CheckResult) {
		return true
	}

	return false
}

// SetCheckResult gets a reference to the given float32 and assigns it to the CheckResult field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetCheckResult(v float32) {
	o.CheckResult = &v
}

// GetCheckUser returns the CheckUser field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckUser() string {
	if o == nil || IsNil(o.CheckUser) {
		var ret string
		return ret
	}
	return *o.CheckUser
}

// GetCheckUserOk returns a tuple with the CheckUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckUserOk() (*string, bool) {
	if o == nil || IsNil(o.CheckUser) {
		return nil, false
	}
	return o.CheckUser, true
}

// IsSetCheckUser returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetCheckUser() bool {
	if o != nil && !IsNil(o.CheckUser) {
		return true
	}

	return false
}

// SetCheckUser gets a reference to the given string and assigns it to the CheckUser field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetCheckUser(v string) {
	o.CheckUser = &v
}

// GetTextCheckOn returns the TextCheckOn field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetTextCheckOn() string {
	if o == nil || IsNil(o.TextCheckOn) {
		var ret string
		return ret
	}
	return *o.TextCheckOn
}

// GetTextCheckOnOk returns a tuple with the TextCheckOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetTextCheckOnOk() (*string, bool) {
	if o == nil || IsNil(o.TextCheckOn) {
		return nil, false
	}
	return o.TextCheckOn, true
}

// IsSetTextCheckOn returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetTextCheckOn() bool {
	if o != nil && !IsNil(o.TextCheckOn) {
		return true
	}

	return false
}

// SetTextCheckOn gets a reference to the given string and assigns it to the TextCheckOn field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetTextCheckOn(v string) {
	o.TextCheckOn = &v
}

// GetCheckPassword returns the CheckPassword field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckPassword() string {
	if o == nil || IsNil(o.CheckPassword) {
		var ret string
		return ret
	}
	return *o.CheckPassword
}

// GetCheckPasswordOk returns a tuple with the CheckPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.CheckPassword) {
		return nil, false
	}
	return o.CheckPassword, true
}

// IsSetCheckPassword returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetCheckPassword() bool {
	if o != nil && !IsNil(o.CheckPassword) {
		return true
	}

	return false
}

// SetCheckPassword gets a reference to the given string and assigns it to the CheckPassword field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetCheckPassword(v string) {
	o.CheckPassword = &v
}

// GetWebTextMatch returns the WebTextMatch field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetWebTextMatch() string {
	if o == nil || IsNil(o.WebTextMatch) {
		var ret string
		return ret
	}
	return *o.WebTextMatch
}

// GetWebTextMatchOk returns a tuple with the WebTextMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetWebTextMatchOk() (*string, bool) {
	if o == nil || IsNil(o.WebTextMatch) {
		return nil, false
	}
	return o.WebTextMatch, true
}

// IsSetWebTextMatch returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetWebTextMatch() bool {
	if o != nil && !IsNil(o.WebTextMatch) {
		return true
	}

	return false
}

// SetWebTextMatch gets a reference to the given string and assigns it to the WebTextMatch field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetWebTextMatch(v string) {
	o.WebTextMatch = &v
}

// GetCheckPasswordHash returns the CheckPasswordHash field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckPasswordHash() string {
	if o == nil || IsNil(o.CheckPasswordHash) {
		var ret string
		return ret
	}
	return *o.CheckPasswordHash
}

// GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetCheckPasswordHashOk() (*string, bool) {
	if o == nil || IsNil(o.CheckPasswordHash) {
		return nil, false
	}
	return o.CheckPasswordHash, true
}

// IsSetCheckPasswordHash returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetCheckPasswordHash() bool {
	if o != nil && !IsNil(o.CheckPasswordHash) {
		return true
	}

	return false
}

// SetCheckPasswordHash gets a reference to the given string and assigns it to the CheckPasswordHash field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetCheckPasswordHash(v string) {
	o.CheckPasswordHash = &v
}

// GetTunnelOn returns the TunnelOn field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetTunnelOn() string {
	if o == nil || IsNil(o.TunnelOn) {
		var ret string
		return ret
	}
	return *o.TunnelOn
}

// GetTunnelOnOk returns a tuple with the TunnelOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetTunnelOnOk() (*string, bool) {
	if o == nil || IsNil(o.TunnelOn) {
		return nil, false
	}
	return o.TunnelOn, true
}

// IsSetTunnelOn returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetTunnelOn() bool {
	if o != nil && !IsNil(o.TunnelOn) {
		return true
	}

	return false
}

// SetTunnelOn gets a reference to the given string and assigns it to the TunnelOn field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetTunnelOn(v string) {
	o.TunnelOn = &v
}

// GetSshHost returns the SshHost field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshHost() string {
	if o == nil || IsNil(o.SshHost) {
		var ret string
		return ret
	}
	return *o.SshHost
}

// GetSshHostOk returns a tuple with the SshHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshHostOk() (*string, bool) {
	if o == nil || IsNil(o.SshHost) {
		return nil, false
	}
	return o.SshHost, true
}

// IsSetSshHost returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetSshHost() bool {
	if o != nil && !IsNil(o.SshHost) {
		return true
	}

	return false
}

// SetSshHost gets a reference to the given string and assigns it to the SshHost field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetSshHost(v string) {
	o.SshHost = &v
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshPort() int64 {
	if o == nil || IsNil(o.SshPort) {
		var ret int64
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshPortOk() (*int64, bool) {
	if o == nil || IsNil(o.SshPort) {
		return nil, false
	}
	return o.SshPort, true
}

// IsSetSshPort returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetSshPort() bool {
	if o != nil && !IsNil(o.SshPort) {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given int64 and assigns it to the SshPort field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetSshPort(v int64) {
	o.SshPort = &v
}

// GetSshUser returns the SshUser field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshUser() string {
	if o == nil || IsNil(o.SshUser) {
		var ret string
		return ret
	}
	return *o.SshUser
}

// GetSshUserOk returns a tuple with the SshUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshUserOk() (*string, bool) {
	if o == nil || IsNil(o.SshUser) {
		return nil, false
	}
	return o.SshUser, true
}

// IsSetSshUser returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetSshUser() bool {
	if o != nil && !IsNil(o.SshUser) {
		return true
	}

	return false
}

// SetSshUser gets a reference to the given string and assigns it to the SshUser field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetSshUser(v string) {
	o.SshUser = &v
}

// GetSshPassword returns the SshPassword field value if set, zero value otherwise.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshPassword() string {
	if o == nil || IsNil(o.SshPassword) {
		var ret string
		return ret
	}
	return *o.SshPassword
}

// GetSshPasswordOk returns a tuple with the SshPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) GetSshPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.SshPassword) {
		return nil, false
	}
	return o.SshPassword, true
}

// IsSetSshPassword returns a boolean if a field has been set.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) IsSetSshPassword() bool {
	if o != nil && !IsNil(o.SshPassword) {
		return true
	}

	return false
}

// SetSshPassword gets a reference to the given string and assigns it to the SshPassword field.
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) SetSshPassword(v string) {
	o.SshPassword = &v
}

func (o GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dbHost"] = o.DbHost
	toSerialize["dbPort"] = o.DbPort
	toSerialize["dbUser"] = o.DbUser
	toSerialize["dbPassword"] = o.DbPassword
	if !IsNil(o.DbPasswordHash) {
		toSerialize["dbPasswordHash"] = o.DbPasswordHash
	}
	toSerialize["dbName"] = o.DbName
	toSerialize["dbQuery"] = o.DbQuery
	if !IsNil(o.CheckOperator) {
		toSerialize["checkOperator"] = o.CheckOperator
	}
	if !IsNil(o.CheckResult) {
		toSerialize["checkResult"] = o.CheckResult
	}
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
func (o *GetAlerts200ResponseAllOfChecksInnerConfigAnyOf1) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

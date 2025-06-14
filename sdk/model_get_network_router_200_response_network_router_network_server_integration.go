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
	"time"
)

// checks if the GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration{}

// GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration struct for GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration
type GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration struct {
	Id                   *int64                                                      `json:"id,omitempty"`
	Name                 *string                                                     `json:"name,omitempty"`
	Enabled              *bool                                                       `json:"enabled,omitempty"`
	Type                 *string                                                     `json:"type,omitempty"`
	IntegrationType      *ListBackupSettings200ResponseBackupSettingsDefaultSchedule `json:"integrationType,omitempty"`
	Url                  *string                                                     `json:"url,omitempty"`
	Port                 *string                                                     `json:"port,omitempty"`
	Username             *string                                                     `json:"username,omitempty"`
	Password             *string                                                     `json:"password,omitempty"`
	RefType              *string                                                     `json:"refType,omitempty"`
	RefId                *string                                                     `json:"refId,omitempty"`
	IsPlugin             *bool                                                       `json:"isPlugin,omitempty"`
	Config               map[string]interface{}                                      `json:"config,omitempty"`
	Status               *string                                                     `json:"status,omitempty"`
	StatusDate           *time.Time                                                  `json:"statusDate,omitempty"`
	StatusMessage        *string                                                     `json:"statusMessage,omitempty"`
	LastSync             *time.Time                                                  `json:"lastSync,omitempty"`
	LastSyncDuration     *int64                                                      `json:"lastSyncDuration,omitempty"`
	AdditionalProperties map[string]interface{}                                      `json:",remain"`
}

type _GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration

// NewGetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration instantiates a new GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration() *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration {
	this := GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration{}
	return &this
}

// NewGetNetworkRouter200ResponseNetworkRouterNetworkServerIntegrationWithDefaults instantiates a new GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkRouter200ResponseNetworkRouterNetworkServerIntegrationWithDefaults() *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration {
	this := GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetType(v string) {
	o.Type = &v
}

// GetIntegrationType returns the IntegrationType field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetIntegrationType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule {
	if o == nil || IsNil(o.IntegrationType) {
		var ret ListBackupSettings200ResponseBackupSettingsDefaultSchedule
		return ret
	}
	return *o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetIntegrationTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool) {
	if o == nil || IsNil(o.IntegrationType) {
		return nil, false
	}
	return o.IntegrationType, true
}

// IsSetIntegrationType returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetIntegrationType() bool {
	if o != nil && !IsNil(o.IntegrationType) {
		return true
	}

	return false
}

// SetIntegrationType gets a reference to the given ListBackupSettings200ResponseBackupSettingsDefaultSchedule and assigns it to the IntegrationType field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetIntegrationType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule) {
	o.IntegrationType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// IsSetUrl returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetUrl(v string) {
	o.Url = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// IsSetPort returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetPort(v string) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// IsSetUsername returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// IsSetPassword returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetPassword(v string) {
	o.Password = &v
}

// GetRefType returns the RefType field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetRefType() string {
	if o == nil || IsNil(o.RefType) {
		var ret string
		return ret
	}
	return *o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetRefTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RefType) {
		return nil, false
	}
	return o.RefType, true
}

// IsSetRefType returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetRefType() bool {
	if o != nil && !IsNil(o.RefType) {
		return true
	}

	return false
}

// SetRefType gets a reference to the given string and assigns it to the RefType field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetRefType(v string) {
	o.RefType = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetRefId() string {
	if o == nil || IsNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// IsSetRefId returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetRefId(v string) {
	o.RefId = &v
}

// GetIsPlugin returns the IsPlugin field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetIsPlugin() bool {
	if o == nil || IsNil(o.IsPlugin) {
		var ret bool
		return ret
	}
	return *o.IsPlugin
}

// GetIsPluginOk returns a tuple with the IsPlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetIsPluginOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPlugin) {
		return nil, false
	}
	return o.IsPlugin, true
}

// IsSetIsPlugin returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetIsPlugin() bool {
	if o != nil && !IsNil(o.IsPlugin) {
		return true
	}

	return false
}

// SetIsPlugin gets a reference to the given bool and assigns it to the IsPlugin field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetIsPlugin(v bool) {
	o.IsPlugin = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// IsSetStatus returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDate returns the StatusDate field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetStatusDate() time.Time {
	if o == nil || IsNil(o.StatusDate) {
		var ret time.Time
		return ret
	}
	return *o.StatusDate
}

// GetStatusDateOk returns a tuple with the StatusDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetStatusDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StatusDate) {
		return nil, false
	}
	return o.StatusDate, true
}

// IsSetStatusDate returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetStatusDate() bool {
	if o != nil && !IsNil(o.StatusDate) {
		return true
	}

	return false
}

// SetStatusDate gets a reference to the given time.Time and assigns it to the StatusDate field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetStatusDate(v time.Time) {
	o.StatusDate = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// IsSetStatusMessage returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetLastSync returns the LastSync field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetLastSync() time.Time {
	if o == nil || IsNil(o.LastSync) {
		var ret time.Time
		return ret
	}
	return *o.LastSync
}

// GetLastSyncOk returns a tuple with the LastSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetLastSyncOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSync) {
		return nil, false
	}
	return o.LastSync, true
}

// IsSetLastSync returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetLastSync() bool {
	if o != nil && !IsNil(o.LastSync) {
		return true
	}

	return false
}

// SetLastSync gets a reference to the given time.Time and assigns it to the LastSync field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetLastSync(v time.Time) {
	o.LastSync = &v
}

// GetLastSyncDuration returns the LastSyncDuration field value if set, zero value otherwise.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetLastSyncDuration() int64 {
	if o == nil || IsNil(o.LastSyncDuration) {
		var ret int64
		return ret
	}
	return *o.LastSyncDuration
}

// GetLastSyncDurationOk returns a tuple with the LastSyncDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetLastSyncDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.LastSyncDuration) {
		return nil, false
	}
	return o.LastSyncDuration, true
}

// IsSetLastSyncDuration returns a boolean if a field has been set.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) IsSetLastSyncDuration() bool {
	if o != nil && !IsNil(o.LastSyncDuration) {
		return true
	}

	return false
}

// SetLastSyncDuration gets a reference to the given int64 and assigns it to the LastSyncDuration field.
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetLastSyncDuration(v int64) {
	o.LastSyncDuration = &v
}

func (o GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IntegrationType) {
		toSerialize["integrationType"] = o.IntegrationType
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.RefType) {
		toSerialize["refType"] = o.RefType
	}
	if !IsNil(o.RefId) {
		toSerialize["refId"] = o.RefId
	}
	if !IsNil(o.IsPlugin) {
		toSerialize["isPlugin"] = o.IsPlugin
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDate) {
		toSerialize["statusDate"] = o.StatusDate
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["statusMessage"] = o.StatusMessage
	}
	if !IsNil(o.LastSync) {
		toSerialize["lastSync"] = o.LastSync
	}
	if !IsNil(o.LastSyncDuration) {
		toSerialize["lastSyncDuration"] = o.LastSyncDuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

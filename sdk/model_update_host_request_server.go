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

// checks if the UpdateHostRequestServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateHostRequestServer{}

// UpdateHostRequestServer struct for UpdateHostRequestServer
type UpdateHostRequestServer struct {
	// Unique name scoped to your account for the server.
	Name *string `json:"name,omitempty"`
	// Optional description field.
	Description *string `json:"description,omitempty"`
	// Flag to determine if a host can be selected for provisioning
	Enabled *bool `json:"enabled,omitempty"`
	// Flag to enable/disable managment of internal firewall
	ManageInternalFirewall *bool `json:"manageInternalFirewall,omitempty"`
	// Flag to enable/disable logs
	EnableLogs *bool `json:"enableLogs,omitempty"`
	// SSH Username
	SshUsername *string `json:"sshUsername,omitempty"`
	// SSH Password
	SshPassword *string                                   `json:"sshPassword,omitempty"`
	SshKeyPair  *AddClusterRequestClusterServerSshKeyPair `json:"sshKeyPair,omitempty"`
	// Power schedule ID.
	PowerScheduleType *int64   `json:"powerScheduleType,omitempty"`
	Labels            []string `json:"labels,omitempty"`
	// Metadata tags, Array of objects having a name and value.
	Tags []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner `json:"tags,omitempty"`
	// Add or update value of Metadata tags, Array of objects having a name and value.
	AddTags []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner `json:"addTags,omitempty"`
	// Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed.
	RemoveTags []ListInstances200ResponseAllOfInstancesInnerTagsInner `json:"removeTags,omitempty"`
	// The Type of guest console this server provides such as disabled, vnc, rdp, ssh
	GuestConsoleType *string `json:"guestConsoleType,omitempty"`
	// The optional guest console username if you don't want to use the user defaults
	GuestConsoleUsername *string `json:"guestConsoleUsername,omitempty"`
	// The optional guest console password if not using the accessing users creds
	GuestConsolePassword *string `json:"guestConsolePassword,omitempty"`
	// The port the guest console is being accessed from
	GuestConsolePort *string `json:"guestConsolePort,omitempty"`
	// Can turn off guest console preferences on server in favor of hypervisor console
	GuestConsolePreferred *bool                  `json:"guestConsolePreferred,omitempty"`
	AdditionalProperties  map[string]interface{} `json:",remain"`
}

type _UpdateHostRequestServer UpdateHostRequestServer

// NewUpdateHostRequestServer instantiates a new UpdateHostRequestServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateHostRequestServer() *UpdateHostRequestServer {
	this := UpdateHostRequestServer{}
	var enabled bool = true
	this.Enabled = &enabled
	var manageInternalFirewall bool = true
	this.ManageInternalFirewall = &manageInternalFirewall
	var enableLogs bool = true
	this.EnableLogs = &enableLogs
	var guestConsolePreferred bool = true
	this.GuestConsolePreferred = &guestConsolePreferred
	return &this
}

// NewUpdateHostRequestServerWithDefaults instantiates a new UpdateHostRequestServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateHostRequestServerWithDefaults() *UpdateHostRequestServer {
	this := UpdateHostRequestServer{}
	var enabled bool = true
	this.Enabled = &enabled
	var manageInternalFirewall bool = true
	this.ManageInternalFirewall = &manageInternalFirewall
	var enableLogs bool = true
	this.EnableLogs = &enableLogs
	var guestConsolePreferred bool = true
	this.GuestConsolePreferred = &guestConsolePreferred
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateHostRequestServer) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateHostRequestServer) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateHostRequestServer) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetManageInternalFirewall returns the ManageInternalFirewall field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetManageInternalFirewall() bool {
	if o == nil || IsNil(o.ManageInternalFirewall) {
		var ret bool
		return ret
	}
	return *o.ManageInternalFirewall
}

// GetManageInternalFirewallOk returns a tuple with the ManageInternalFirewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetManageInternalFirewallOk() (*bool, bool) {
	if o == nil || IsNil(o.ManageInternalFirewall) {
		return nil, false
	}
	return o.ManageInternalFirewall, true
}

// IsSetManageInternalFirewall returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetManageInternalFirewall() bool {
	if o != nil && !IsNil(o.ManageInternalFirewall) {
		return true
	}

	return false
}

// SetManageInternalFirewall gets a reference to the given bool and assigns it to the ManageInternalFirewall field.
func (o *UpdateHostRequestServer) SetManageInternalFirewall(v bool) {
	o.ManageInternalFirewall = &v
}

// GetEnableLogs returns the EnableLogs field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetEnableLogs() bool {
	if o == nil || IsNil(o.EnableLogs) {
		var ret bool
		return ret
	}
	return *o.EnableLogs
}

// GetEnableLogsOk returns a tuple with the EnableLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetEnableLogsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableLogs) {
		return nil, false
	}
	return o.EnableLogs, true
}

// IsSetEnableLogs returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetEnableLogs() bool {
	if o != nil && !IsNil(o.EnableLogs) {
		return true
	}

	return false
}

// SetEnableLogs gets a reference to the given bool and assigns it to the EnableLogs field.
func (o *UpdateHostRequestServer) SetEnableLogs(v bool) {
	o.EnableLogs = &v
}

// GetSshUsername returns the SshUsername field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetSshUsername() string {
	if o == nil || IsNil(o.SshUsername) {
		var ret string
		return ret
	}
	return *o.SshUsername
}

// GetSshUsernameOk returns a tuple with the SshUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetSshUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.SshUsername) {
		return nil, false
	}
	return o.SshUsername, true
}

// IsSetSshUsername returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetSshUsername() bool {
	if o != nil && !IsNil(o.SshUsername) {
		return true
	}

	return false
}

// SetSshUsername gets a reference to the given string and assigns it to the SshUsername field.
func (o *UpdateHostRequestServer) SetSshUsername(v string) {
	o.SshUsername = &v
}

// GetSshPassword returns the SshPassword field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetSshPassword() string {
	if o == nil || IsNil(o.SshPassword) {
		var ret string
		return ret
	}
	return *o.SshPassword
}

// GetSshPasswordOk returns a tuple with the SshPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetSshPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.SshPassword) {
		return nil, false
	}
	return o.SshPassword, true
}

// IsSetSshPassword returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetSshPassword() bool {
	if o != nil && !IsNil(o.SshPassword) {
		return true
	}

	return false
}

// SetSshPassword gets a reference to the given string and assigns it to the SshPassword field.
func (o *UpdateHostRequestServer) SetSshPassword(v string) {
	o.SshPassword = &v
}

// GetSshKeyPair returns the SshKeyPair field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetSshKeyPair() AddClusterRequestClusterServerSshKeyPair {
	if o == nil || IsNil(o.SshKeyPair) {
		var ret AddClusterRequestClusterServerSshKeyPair
		return ret
	}
	return *o.SshKeyPair
}

// GetSshKeyPairOk returns a tuple with the SshKeyPair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetSshKeyPairOk() (*AddClusterRequestClusterServerSshKeyPair, bool) {
	if o == nil || IsNil(o.SshKeyPair) {
		return nil, false
	}
	return o.SshKeyPair, true
}

// IsSetSshKeyPair returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetSshKeyPair() bool {
	if o != nil && !IsNil(o.SshKeyPair) {
		return true
	}

	return false
}

// SetSshKeyPair gets a reference to the given AddClusterRequestClusterServerSshKeyPair and assigns it to the SshKeyPair field.
func (o *UpdateHostRequestServer) SetSshKeyPair(v AddClusterRequestClusterServerSshKeyPair) {
	o.SshKeyPair = &v
}

// GetPowerScheduleType returns the PowerScheduleType field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetPowerScheduleType() int64 {
	if o == nil || IsNil(o.PowerScheduleType) {
		var ret int64
		return ret
	}
	return *o.PowerScheduleType
}

// GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetPowerScheduleTypeOk() (*int64, bool) {
	if o == nil || IsNil(o.PowerScheduleType) {
		return nil, false
	}
	return o.PowerScheduleType, true
}

// IsSetPowerScheduleType returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetPowerScheduleType() bool {
	if o != nil && !IsNil(o.PowerScheduleType) {
		return true
	}

	return false
}

// SetPowerScheduleType gets a reference to the given int64 and assigns it to the PowerScheduleType field.
func (o *UpdateHostRequestServer) SetPowerScheduleType(v int64) {
	o.PowerScheduleType = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *UpdateHostRequestServer) SetLabels(v []string) {
	o.Labels = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner {
	if o == nil || IsNil(o.Tags) {
		var ret []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetTagsOk() ([]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// IsSetTags returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner and assigns it to the Tags field.
func (o *UpdateHostRequestServer) SetTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner) {
	o.Tags = v
}

// GetAddTags returns the AddTags field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetAddTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner {
	if o == nil || IsNil(o.AddTags) {
		var ret []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner
		return ret
	}
	return o.AddTags
}

// GetAddTagsOk returns a tuple with the AddTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetAddTagsOk() ([]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool) {
	if o == nil || IsNil(o.AddTags) {
		return nil, false
	}
	return o.AddTags, true
}

// IsSetAddTags returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetAddTags() bool {
	if o != nil && !IsNil(o.AddTags) {
		return true
	}

	return false
}

// SetAddTags gets a reference to the given []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner and assigns it to the AddTags field.
func (o *UpdateHostRequestServer) SetAddTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner) {
	o.AddTags = v
}

// GetRemoveTags returns the RemoveTags field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetRemoveTags() []ListInstances200ResponseAllOfInstancesInnerTagsInner {
	if o == nil || IsNil(o.RemoveTags) {
		var ret []ListInstances200ResponseAllOfInstancesInnerTagsInner
		return ret
	}
	return o.RemoveTags
}

// GetRemoveTagsOk returns a tuple with the RemoveTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetRemoveTagsOk() ([]ListInstances200ResponseAllOfInstancesInnerTagsInner, bool) {
	if o == nil || IsNil(o.RemoveTags) {
		return nil, false
	}
	return o.RemoveTags, true
}

// IsSetRemoveTags returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetRemoveTags() bool {
	if o != nil && !IsNil(o.RemoveTags) {
		return true
	}

	return false
}

// SetRemoveTags gets a reference to the given []ListInstances200ResponseAllOfInstancesInnerTagsInner and assigns it to the RemoveTags field.
func (o *UpdateHostRequestServer) SetRemoveTags(v []ListInstances200ResponseAllOfInstancesInnerTagsInner) {
	o.RemoveTags = v
}

// GetGuestConsoleType returns the GuestConsoleType field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetGuestConsoleType() string {
	if o == nil || IsNil(o.GuestConsoleType) {
		var ret string
		return ret
	}
	return *o.GuestConsoleType
}

// GetGuestConsoleTypeOk returns a tuple with the GuestConsoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetGuestConsoleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GuestConsoleType) {
		return nil, false
	}
	return o.GuestConsoleType, true
}

// IsSetGuestConsoleType returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetGuestConsoleType() bool {
	if o != nil && !IsNil(o.GuestConsoleType) {
		return true
	}

	return false
}

// SetGuestConsoleType gets a reference to the given string and assigns it to the GuestConsoleType field.
func (o *UpdateHostRequestServer) SetGuestConsoleType(v string) {
	o.GuestConsoleType = &v
}

// GetGuestConsoleUsername returns the GuestConsoleUsername field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetGuestConsoleUsername() string {
	if o == nil || IsNil(o.GuestConsoleUsername) {
		var ret string
		return ret
	}
	return *o.GuestConsoleUsername
}

// GetGuestConsoleUsernameOk returns a tuple with the GuestConsoleUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetGuestConsoleUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.GuestConsoleUsername) {
		return nil, false
	}
	return o.GuestConsoleUsername, true
}

// IsSetGuestConsoleUsername returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetGuestConsoleUsername() bool {
	if o != nil && !IsNil(o.GuestConsoleUsername) {
		return true
	}

	return false
}

// SetGuestConsoleUsername gets a reference to the given string and assigns it to the GuestConsoleUsername field.
func (o *UpdateHostRequestServer) SetGuestConsoleUsername(v string) {
	o.GuestConsoleUsername = &v
}

// GetGuestConsolePassword returns the GuestConsolePassword field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetGuestConsolePassword() string {
	if o == nil || IsNil(o.GuestConsolePassword) {
		var ret string
		return ret
	}
	return *o.GuestConsolePassword
}

// GetGuestConsolePasswordOk returns a tuple with the GuestConsolePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetGuestConsolePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.GuestConsolePassword) {
		return nil, false
	}
	return o.GuestConsolePassword, true
}

// IsSetGuestConsolePassword returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetGuestConsolePassword() bool {
	if o != nil && !IsNil(o.GuestConsolePassword) {
		return true
	}

	return false
}

// SetGuestConsolePassword gets a reference to the given string and assigns it to the GuestConsolePassword field.
func (o *UpdateHostRequestServer) SetGuestConsolePassword(v string) {
	o.GuestConsolePassword = &v
}

// GetGuestConsolePort returns the GuestConsolePort field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetGuestConsolePort() string {
	if o == nil || IsNil(o.GuestConsolePort) {
		var ret string
		return ret
	}
	return *o.GuestConsolePort
}

// GetGuestConsolePortOk returns a tuple with the GuestConsolePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetGuestConsolePortOk() (*string, bool) {
	if o == nil || IsNil(o.GuestConsolePort) {
		return nil, false
	}
	return o.GuestConsolePort, true
}

// IsSetGuestConsolePort returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetGuestConsolePort() bool {
	if o != nil && !IsNil(o.GuestConsolePort) {
		return true
	}

	return false
}

// SetGuestConsolePort gets a reference to the given string and assigns it to the GuestConsolePort field.
func (o *UpdateHostRequestServer) SetGuestConsolePort(v string) {
	o.GuestConsolePort = &v
}

// GetGuestConsolePreferred returns the GuestConsolePreferred field value if set, zero value otherwise.
func (o *UpdateHostRequestServer) GetGuestConsolePreferred() bool {
	if o == nil || IsNil(o.GuestConsolePreferred) {
		var ret bool
		return ret
	}
	return *o.GuestConsolePreferred
}

// GetGuestConsolePreferredOk returns a tuple with the GuestConsolePreferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHostRequestServer) GetGuestConsolePreferredOk() (*bool, bool) {
	if o == nil || IsNil(o.GuestConsolePreferred) {
		return nil, false
	}
	return o.GuestConsolePreferred, true
}

// IsSetGuestConsolePreferred returns a boolean if a field has been set.
func (o *UpdateHostRequestServer) IsSetGuestConsolePreferred() bool {
	if o != nil && !IsNil(o.GuestConsolePreferred) {
		return true
	}

	return false
}

// SetGuestConsolePreferred gets a reference to the given bool and assigns it to the GuestConsolePreferred field.
func (o *UpdateHostRequestServer) SetGuestConsolePreferred(v bool) {
	o.GuestConsolePreferred = &v
}

func (o UpdateHostRequestServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateHostRequestServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ManageInternalFirewall) {
		toSerialize["manageInternalFirewall"] = o.ManageInternalFirewall
	}
	if !IsNil(o.EnableLogs) {
		toSerialize["enableLogs"] = o.EnableLogs
	}
	if !IsNil(o.SshUsername) {
		toSerialize["sshUsername"] = o.SshUsername
	}
	if !IsNil(o.SshPassword) {
		toSerialize["sshPassword"] = o.SshPassword
	}
	if !IsNil(o.SshKeyPair) {
		toSerialize["sshKeyPair"] = o.SshKeyPair
	}
	if !IsNil(o.PowerScheduleType) {
		toSerialize["powerScheduleType"] = o.PowerScheduleType
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.AddTags) {
		toSerialize["addTags"] = o.AddTags
	}
	if !IsNil(o.RemoveTags) {
		toSerialize["removeTags"] = o.RemoveTags
	}
	if !IsNil(o.GuestConsoleType) {
		toSerialize["guestConsoleType"] = o.GuestConsoleType
	}
	if !IsNil(o.GuestConsoleUsername) {
		toSerialize["guestConsoleUsername"] = o.GuestConsoleUsername
	}
	if !IsNil(o.GuestConsolePassword) {
		toSerialize["guestConsolePassword"] = o.GuestConsolePassword
	}
	if !IsNil(o.GuestConsolePort) {
		toSerialize["guestConsolePort"] = o.GuestConsolePort
	}
	if !IsNil(o.GuestConsolePreferred) {
		toSerialize["guestConsolePreferred"] = o.GuestConsolePreferred
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateHostRequestServer) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

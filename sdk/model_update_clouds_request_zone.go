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

// checks if the UpdateCloudsRequestZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCloudsRequestZone{}

// UpdateCloudsRequestZone struct for UpdateCloudsRequestZone
type UpdateCloudsRequestZone struct {
	// A unique name scoped to your account for the cloud
	Name string `json:"name"`
	// Optional description field if you want to put more info there
	Description *string `json:"description,omitempty"`
	// Optional code for use with policies
	Code *string `json:"code,omitempty"`
	// Array of label strings, can be used for filtering.
	Labels []string `json:"labels,omitempty"`
	// Optional location for your cloud
	Location *string `json:"location,omitempty"`
	// private or public
	Visibility *string `json:"visibility,omitempty"`
	// Map containing code or id of the cloud type
	ZoneType string `json:"zoneType"`
	// Specifies which Server group this cloud should be assigned to
	GroupId int64 `json:"groupId"`
	// Specifies which Tenant this cloud should be assigned to
	AccountId *int64 `json:"accountId,omitempty"`
	// Can be used to disable the cloud
	Enabled *bool `json:"enabled,omitempty"`
	// Automatically Power on VMs
	AutoRecoverPowerState *bool `json:"autoRecoverPowerState,omitempty"`
	// Scale Priority
	ScalePriority *int64 `json:"scalePriority,omitempty"`
	// Sets the default active state during discovery of new datastores.
	DefaultDatastoreSyncActive *bool `json:"defaultDatastoreSyncActive,omitempty"`
	// Sets the default active state during discovery of new networks.
	DefaultNetworkSyncActive *bool `json:"defaultNetworkSyncActive,omitempty"`
	// Sets the default active state during discovery of new folders.
	DefaultFolderSyncActive *bool `json:"defaultFolderSyncActive,omitempty"`
	// Sets the default active state during discovery of new security groups.
	DefaultSecurityGroupSyncActive *bool `json:"defaultSecurityGroupSyncActive,omitempty"`
	// Sets the default active state during discovery of new resource pools.
	DefaultPoolSyncActive *bool `json:"defaultPoolSyncActive,omitempty"`
	// Sets the default active state during discovery of new plans.
	DefaultPlanSyncActive *bool `json:"defaultPlanSyncActive,omitempty"`
	// Linked Account ID (enter commercial ID to get costing for AWS Govcloud)
	LinkedAccountId *int64 `json:"linkedAccountId,omitempty"`
	// Map containing zone configuration settings. See the section on specific zone types for details.
	Config map[string]interface{} `json:"config,omitempty"`
	// host firewall. `off` or `internal`. a.k.a. \"local firewall\"
	SecurityMode *string `json:"securityMode,omitempty"`
	// Can be used to clear any custom logo and darkLogo, reverting to the defaults for the cloud type
	DefaultCloudLogos *bool `json:"defaultCloudLogos,omitempty"`
	// Map containing Credential ID. `local` means use the values set in the local cloud config instead of associating a credential.
	Credential           map[string]interface{} `json:"credential"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _UpdateCloudsRequestZone UpdateCloudsRequestZone

// NewUpdateCloudsRequestZone instantiates a new UpdateCloudsRequestZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCloudsRequestZone(name string, zoneType string, groupId int64, credential map[string]interface{}) *UpdateCloudsRequestZone {
	this := UpdateCloudsRequestZone{}
	this.Name = name
	var visibility string = "private"
	this.Visibility = &visibility
	this.ZoneType = zoneType
	this.GroupId = groupId
	var enabled bool = true
	this.Enabled = &enabled
	var autoRecoverPowerState bool = false
	this.AutoRecoverPowerState = &autoRecoverPowerState
	var scalePriority int64 = 1
	this.ScalePriority = &scalePriority
	var securityMode string = "off"
	this.SecurityMode = &securityMode
	this.Credential = credential
	return &this
}

// NewUpdateCloudsRequestZoneWithDefaults instantiates a new UpdateCloudsRequestZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCloudsRequestZoneWithDefaults() *UpdateCloudsRequestZone {
	this := UpdateCloudsRequestZone{}
	var visibility string = "private"
	this.Visibility = &visibility
	var zoneType string = "standard"
	this.ZoneType = zoneType
	var enabled bool = true
	this.Enabled = &enabled
	var autoRecoverPowerState bool = false
	this.AutoRecoverPowerState = &autoRecoverPowerState
	var scalePriority int64 = 1
	this.ScalePriority = &scalePriority
	var securityMode string = "off"
	this.SecurityMode = &securityMode
	return &this
}

// GetName returns the Name field value
func (o *UpdateCloudsRequestZone) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateCloudsRequestZone) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateCloudsRequestZone) SetDescription(v string) {
	o.Description = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdateCloudsRequestZone) SetCode(v string) {
	o.Code = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *UpdateCloudsRequestZone) SetLabels(v []string) {
	o.Labels = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// IsSetLocation returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *UpdateCloudsRequestZone) SetLocation(v string) {
	o.Location = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *UpdateCloudsRequestZone) SetVisibility(v string) {
	o.Visibility = &v
}

// GetZoneType returns the ZoneType field value
func (o *UpdateCloudsRequestZone) GetZoneType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZoneType
}

// GetZoneTypeOk returns a tuple with the ZoneType field value
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetZoneTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneType, true
}

// SetZoneType sets field value
func (o *UpdateCloudsRequestZone) SetZoneType(v string) {
	o.ZoneType = v
}

// GetGroupId returns the GroupId field value
func (o *UpdateCloudsRequestZone) GetGroupId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *UpdateCloudsRequestZone) SetGroupId(v int64) {
	o.GroupId = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetAccountId() int64 {
	if o == nil || IsNil(o.AccountId) {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// IsSetAccountId returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *UpdateCloudsRequestZone) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateCloudsRequestZone) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAutoRecoverPowerState returns the AutoRecoverPowerState field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetAutoRecoverPowerState() bool {
	if o == nil || IsNil(o.AutoRecoverPowerState) {
		var ret bool
		return ret
	}
	return *o.AutoRecoverPowerState
}

// GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetAutoRecoverPowerStateOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRecoverPowerState) {
		return nil, false
	}
	return o.AutoRecoverPowerState, true
}

// IsSetAutoRecoverPowerState returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetAutoRecoverPowerState() bool {
	if o != nil && !IsNil(o.AutoRecoverPowerState) {
		return true
	}

	return false
}

// SetAutoRecoverPowerState gets a reference to the given bool and assigns it to the AutoRecoverPowerState field.
func (o *UpdateCloudsRequestZone) SetAutoRecoverPowerState(v bool) {
	o.AutoRecoverPowerState = &v
}

// GetScalePriority returns the ScalePriority field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetScalePriority() int64 {
	if o == nil || IsNil(o.ScalePriority) {
		var ret int64
		return ret
	}
	return *o.ScalePriority
}

// GetScalePriorityOk returns a tuple with the ScalePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetScalePriorityOk() (*int64, bool) {
	if o == nil || IsNil(o.ScalePriority) {
		return nil, false
	}
	return o.ScalePriority, true
}

// IsSetScalePriority returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetScalePriority() bool {
	if o != nil && !IsNil(o.ScalePriority) {
		return true
	}

	return false
}

// SetScalePriority gets a reference to the given int64 and assigns it to the ScalePriority field.
func (o *UpdateCloudsRequestZone) SetScalePriority(v int64) {
	o.ScalePriority = &v
}

// GetDefaultDatastoreSyncActive returns the DefaultDatastoreSyncActive field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetDefaultDatastoreSyncActive() bool {
	if o == nil || IsNil(o.DefaultDatastoreSyncActive) {
		var ret bool
		return ret
	}
	return *o.DefaultDatastoreSyncActive
}

// GetDefaultDatastoreSyncActiveOk returns a tuple with the DefaultDatastoreSyncActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetDefaultDatastoreSyncActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultDatastoreSyncActive) {
		return nil, false
	}
	return o.DefaultDatastoreSyncActive, true
}

// IsSetDefaultDatastoreSyncActive returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetDefaultDatastoreSyncActive() bool {
	if o != nil && !IsNil(o.DefaultDatastoreSyncActive) {
		return true
	}

	return false
}

// SetDefaultDatastoreSyncActive gets a reference to the given bool and assigns it to the DefaultDatastoreSyncActive field.
func (o *UpdateCloudsRequestZone) SetDefaultDatastoreSyncActive(v bool) {
	o.DefaultDatastoreSyncActive = &v
}

// GetDefaultNetworkSyncActive returns the DefaultNetworkSyncActive field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetDefaultNetworkSyncActive() bool {
	if o == nil || IsNil(o.DefaultNetworkSyncActive) {
		var ret bool
		return ret
	}
	return *o.DefaultNetworkSyncActive
}

// GetDefaultNetworkSyncActiveOk returns a tuple with the DefaultNetworkSyncActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetDefaultNetworkSyncActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultNetworkSyncActive) {
		return nil, false
	}
	return o.DefaultNetworkSyncActive, true
}

// IsSetDefaultNetworkSyncActive returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetDefaultNetworkSyncActive() bool {
	if o != nil && !IsNil(o.DefaultNetworkSyncActive) {
		return true
	}

	return false
}

// SetDefaultNetworkSyncActive gets a reference to the given bool and assigns it to the DefaultNetworkSyncActive field.
func (o *UpdateCloudsRequestZone) SetDefaultNetworkSyncActive(v bool) {
	o.DefaultNetworkSyncActive = &v
}

// GetDefaultFolderSyncActive returns the DefaultFolderSyncActive field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetDefaultFolderSyncActive() bool {
	if o == nil || IsNil(o.DefaultFolderSyncActive) {
		var ret bool
		return ret
	}
	return *o.DefaultFolderSyncActive
}

// GetDefaultFolderSyncActiveOk returns a tuple with the DefaultFolderSyncActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetDefaultFolderSyncActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultFolderSyncActive) {
		return nil, false
	}
	return o.DefaultFolderSyncActive, true
}

// IsSetDefaultFolderSyncActive returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetDefaultFolderSyncActive() bool {
	if o != nil && !IsNil(o.DefaultFolderSyncActive) {
		return true
	}

	return false
}

// SetDefaultFolderSyncActive gets a reference to the given bool and assigns it to the DefaultFolderSyncActive field.
func (o *UpdateCloudsRequestZone) SetDefaultFolderSyncActive(v bool) {
	o.DefaultFolderSyncActive = &v
}

// GetDefaultSecurityGroupSyncActive returns the DefaultSecurityGroupSyncActive field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetDefaultSecurityGroupSyncActive() bool {
	if o == nil || IsNil(o.DefaultSecurityGroupSyncActive) {
		var ret bool
		return ret
	}
	return *o.DefaultSecurityGroupSyncActive
}

// GetDefaultSecurityGroupSyncActiveOk returns a tuple with the DefaultSecurityGroupSyncActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetDefaultSecurityGroupSyncActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultSecurityGroupSyncActive) {
		return nil, false
	}
	return o.DefaultSecurityGroupSyncActive, true
}

// IsSetDefaultSecurityGroupSyncActive returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetDefaultSecurityGroupSyncActive() bool {
	if o != nil && !IsNil(o.DefaultSecurityGroupSyncActive) {
		return true
	}

	return false
}

// SetDefaultSecurityGroupSyncActive gets a reference to the given bool and assigns it to the DefaultSecurityGroupSyncActive field.
func (o *UpdateCloudsRequestZone) SetDefaultSecurityGroupSyncActive(v bool) {
	o.DefaultSecurityGroupSyncActive = &v
}

// GetDefaultPoolSyncActive returns the DefaultPoolSyncActive field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetDefaultPoolSyncActive() bool {
	if o == nil || IsNil(o.DefaultPoolSyncActive) {
		var ret bool
		return ret
	}
	return *o.DefaultPoolSyncActive
}

// GetDefaultPoolSyncActiveOk returns a tuple with the DefaultPoolSyncActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetDefaultPoolSyncActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultPoolSyncActive) {
		return nil, false
	}
	return o.DefaultPoolSyncActive, true
}

// IsSetDefaultPoolSyncActive returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetDefaultPoolSyncActive() bool {
	if o != nil && !IsNil(o.DefaultPoolSyncActive) {
		return true
	}

	return false
}

// SetDefaultPoolSyncActive gets a reference to the given bool and assigns it to the DefaultPoolSyncActive field.
func (o *UpdateCloudsRequestZone) SetDefaultPoolSyncActive(v bool) {
	o.DefaultPoolSyncActive = &v
}

// GetDefaultPlanSyncActive returns the DefaultPlanSyncActive field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetDefaultPlanSyncActive() bool {
	if o == nil || IsNil(o.DefaultPlanSyncActive) {
		var ret bool
		return ret
	}
	return *o.DefaultPlanSyncActive
}

// GetDefaultPlanSyncActiveOk returns a tuple with the DefaultPlanSyncActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetDefaultPlanSyncActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultPlanSyncActive) {
		return nil, false
	}
	return o.DefaultPlanSyncActive, true
}

// IsSetDefaultPlanSyncActive returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetDefaultPlanSyncActive() bool {
	if o != nil && !IsNil(o.DefaultPlanSyncActive) {
		return true
	}

	return false
}

// SetDefaultPlanSyncActive gets a reference to the given bool and assigns it to the DefaultPlanSyncActive field.
func (o *UpdateCloudsRequestZone) SetDefaultPlanSyncActive(v bool) {
	o.DefaultPlanSyncActive = &v
}

// GetLinkedAccountId returns the LinkedAccountId field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetLinkedAccountId() int64 {
	if o == nil || IsNil(o.LinkedAccountId) {
		var ret int64
		return ret
	}
	return *o.LinkedAccountId
}

// GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetLinkedAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.LinkedAccountId) {
		return nil, false
	}
	return o.LinkedAccountId, true
}

// IsSetLinkedAccountId returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetLinkedAccountId() bool {
	if o != nil && !IsNil(o.LinkedAccountId) {
		return true
	}

	return false
}

// SetLinkedAccountId gets a reference to the given int64 and assigns it to the LinkedAccountId field.
func (o *UpdateCloudsRequestZone) SetLinkedAccountId(v int64) {
	o.LinkedAccountId = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *UpdateCloudsRequestZone) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetSecurityMode returns the SecurityMode field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetSecurityMode() string {
	if o == nil || IsNil(o.SecurityMode) {
		var ret string
		return ret
	}
	return *o.SecurityMode
}

// GetSecurityModeOk returns a tuple with the SecurityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetSecurityModeOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityMode) {
		return nil, false
	}
	return o.SecurityMode, true
}

// IsSetSecurityMode returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetSecurityMode() bool {
	if o != nil && !IsNil(o.SecurityMode) {
		return true
	}

	return false
}

// SetSecurityMode gets a reference to the given string and assigns it to the SecurityMode field.
func (o *UpdateCloudsRequestZone) SetSecurityMode(v string) {
	o.SecurityMode = &v
}

// GetDefaultCloudLogos returns the DefaultCloudLogos field value if set, zero value otherwise.
func (o *UpdateCloudsRequestZone) GetDefaultCloudLogos() bool {
	if o == nil || IsNil(o.DefaultCloudLogos) {
		var ret bool
		return ret
	}
	return *o.DefaultCloudLogos
}

// GetDefaultCloudLogosOk returns a tuple with the DefaultCloudLogos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetDefaultCloudLogosOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultCloudLogos) {
		return nil, false
	}
	return o.DefaultCloudLogos, true
}

// IsSetDefaultCloudLogos returns a boolean if a field has been set.
func (o *UpdateCloudsRequestZone) IsSetDefaultCloudLogos() bool {
	if o != nil && !IsNil(o.DefaultCloudLogos) {
		return true
	}

	return false
}

// SetDefaultCloudLogos gets a reference to the given bool and assigns it to the DefaultCloudLogos field.
func (o *UpdateCloudsRequestZone) SetDefaultCloudLogos(v bool) {
	o.DefaultCloudLogos = &v
}

// GetCredential returns the Credential field value
func (o *UpdateCloudsRequestZone) GetCredential() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value
// and a boolean to check if the value has been set.
func (o *UpdateCloudsRequestZone) GetCredentialOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Credential, true
}

// SetCredential sets field value
func (o *UpdateCloudsRequestZone) SetCredential(v map[string]interface{}) {
	o.Credential = v
}

func (o UpdateCloudsRequestZone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCloudsRequestZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	toSerialize["zoneType"] = o.ZoneType
	toSerialize["groupId"] = o.GroupId
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.AutoRecoverPowerState) {
		toSerialize["autoRecoverPowerState"] = o.AutoRecoverPowerState
	}
	if !IsNil(o.ScalePriority) {
		toSerialize["scalePriority"] = o.ScalePriority
	}
	if !IsNil(o.DefaultDatastoreSyncActive) {
		toSerialize["defaultDatastoreSyncActive"] = o.DefaultDatastoreSyncActive
	}
	if !IsNil(o.DefaultNetworkSyncActive) {
		toSerialize["defaultNetworkSyncActive"] = o.DefaultNetworkSyncActive
	}
	if !IsNil(o.DefaultFolderSyncActive) {
		toSerialize["defaultFolderSyncActive"] = o.DefaultFolderSyncActive
	}
	if !IsNil(o.DefaultSecurityGroupSyncActive) {
		toSerialize["defaultSecurityGroupSyncActive"] = o.DefaultSecurityGroupSyncActive
	}
	if !IsNil(o.DefaultPoolSyncActive) {
		toSerialize["defaultPoolSyncActive"] = o.DefaultPoolSyncActive
	}
	if !IsNil(o.DefaultPlanSyncActive) {
		toSerialize["defaultPlanSyncActive"] = o.DefaultPlanSyncActive
	}
	if !IsNil(o.LinkedAccountId) {
		toSerialize["linkedAccountId"] = o.LinkedAccountId
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.SecurityMode) {
		toSerialize["securityMode"] = o.SecurityMode
	}
	if !IsNil(o.DefaultCloudLogos) {
		toSerialize["defaultCloudLogos"] = o.DefaultCloudLogos
	}
	toSerialize["credential"] = o.Credential

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateCloudsRequestZone) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

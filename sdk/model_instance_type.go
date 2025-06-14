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

// checks if the InstanceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceType{}

// InstanceType struct for InstanceType
type InstanceType struct {
	Id                   *int64                                                                      `json:"id,omitempty"`
	Account              *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner     `json:"account,omitempty"`
	Name                 *string                                                                     `json:"name,omitempty"`
	Labels               []string                                                                    `json:"labels,omitempty"`
	Code                 *string                                                                     `json:"code,omitempty"`
	Description          *string                                                                     `json:"description,omitempty"`
	ProvisionTypeCode    *string                                                                     `json:"provisionTypeCode,omitempty"`
	Category             *string                                                                     `json:"category,omitempty"`
	Active               *bool                                                                       `json:"active,omitempty"`
	HasProvisioningStep  *bool                                                                       `json:"hasProvisioningStep,omitempty"`
	HasDeployment        *bool                                                                       `json:"hasDeployment,omitempty"`
	HasConfig            *bool                                                                       `json:"hasConfig,omitempty"`
	HasSettings          *bool                                                                       `json:"hasSettings,omitempty"`
	HasAutoScale         *bool                                                                       `json:"hasAutoScale,omitempty"`
	ProxyType            *string                                                                     `json:"proxyType,omitempty"`
	ProxyPort            *string                                                                     `json:"proxyPort,omitempty"`
	ProxyProtocol        *string                                                                     `json:"proxyProtocol,omitempty"`
	EnvironmentPrefix    *string                                                                     `json:"environmentPrefix,omitempty"`
	BackupType           *string                                                                     `json:"backupType,omitempty"`
	Config               map[string]interface{}                                                      `json:"config,omitempty"`
	Visibility           *string                                                                     `json:"visibility,omitempty"`
	Featured             *bool                                                                       `json:"featured,omitempty"`
	Versions             []string                                                                    `json:"versions,omitempty"`
	InstanceTypeLayouts  []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner            `json:"instanceTypeLayouts,omitempty"`
	OptionTypes          []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner `json:"optionTypes,omitempty"`
	EnvironmentVariables []map[string]interface{}                                                    `json:"environmentVariables,omitempty"`
	PriceSets            []map[string]interface{}                                                    `json:"priceSets,omitempty"`
	// Logo image URL
	ImagePath *string `json:"imagePath,omitempty"`
	// Dark logo image URL
	DarkImagePath        *string                `json:"darkImagePath,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _InstanceType InstanceType

// NewInstanceType instantiates a new InstanceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceType() *InstanceType {
	this := InstanceType{}
	return &this
}

// NewInstanceTypeWithDefaults instantiates a new InstanceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceTypeWithDefaults() *InstanceType {
	this := InstanceType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceType) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *InstanceType) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InstanceType) SetId(v int64) {
	o.Id = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *InstanceType) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Account) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// IsSetAccount returns a boolean if a field has been set.
func (o *InstanceType) IsSetAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Account field.
func (o *InstanceType) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Account = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *InstanceType) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceType) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *InstanceType) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *InstanceType) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *InstanceType) SetLabels(v []string) {
	o.Labels = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InstanceType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *InstanceType) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InstanceType) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InstanceType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *InstanceType) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InstanceType) SetDescription(v string) {
	o.Description = &v
}

// GetProvisionTypeCode returns the ProvisionTypeCode field value if set, zero value otherwise.
func (o *InstanceType) GetProvisionTypeCode() string {
	if o == nil || IsNil(o.ProvisionTypeCode) {
		var ret string
		return ret
	}
	return *o.ProvisionTypeCode
}

// GetProvisionTypeCodeOk returns a tuple with the ProvisionTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetProvisionTypeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProvisionTypeCode) {
		return nil, false
	}
	return o.ProvisionTypeCode, true
}

// IsSetProvisionTypeCode returns a boolean if a field has been set.
func (o *InstanceType) IsSetProvisionTypeCode() bool {
	if o != nil && !IsNil(o.ProvisionTypeCode) {
		return true
	}

	return false
}

// SetProvisionTypeCode gets a reference to the given string and assigns it to the ProvisionTypeCode field.
func (o *InstanceType) SetProvisionTypeCode(v string) {
	o.ProvisionTypeCode = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InstanceType) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// IsSetCategory returns a boolean if a field has been set.
func (o *InstanceType) IsSetCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InstanceType) SetCategory(v string) {
	o.Category = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *InstanceType) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// IsSetActive returns a boolean if a field has been set.
func (o *InstanceType) IsSetActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *InstanceType) SetActive(v bool) {
	o.Active = &v
}

// GetHasProvisioningStep returns the HasProvisioningStep field value if set, zero value otherwise.
func (o *InstanceType) GetHasProvisioningStep() bool {
	if o == nil || IsNil(o.HasProvisioningStep) {
		var ret bool
		return ret
	}
	return *o.HasProvisioningStep
}

// GetHasProvisioningStepOk returns a tuple with the HasProvisioningStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetHasProvisioningStepOk() (*bool, bool) {
	if o == nil || IsNil(o.HasProvisioningStep) {
		return nil, false
	}
	return o.HasProvisioningStep, true
}

// IsSetHasProvisioningStep returns a boolean if a field has been set.
func (o *InstanceType) IsSetHasProvisioningStep() bool {
	if o != nil && !IsNil(o.HasProvisioningStep) {
		return true
	}

	return false
}

// SetHasProvisioningStep gets a reference to the given bool and assigns it to the HasProvisioningStep field.
func (o *InstanceType) SetHasProvisioningStep(v bool) {
	o.HasProvisioningStep = &v
}

// GetHasDeployment returns the HasDeployment field value if set, zero value otherwise.
func (o *InstanceType) GetHasDeployment() bool {
	if o == nil || IsNil(o.HasDeployment) {
		var ret bool
		return ret
	}
	return *o.HasDeployment
}

// GetHasDeploymentOk returns a tuple with the HasDeployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetHasDeploymentOk() (*bool, bool) {
	if o == nil || IsNil(o.HasDeployment) {
		return nil, false
	}
	return o.HasDeployment, true
}

// IsSetHasDeployment returns a boolean if a field has been set.
func (o *InstanceType) IsSetHasDeployment() bool {
	if o != nil && !IsNil(o.HasDeployment) {
		return true
	}

	return false
}

// SetHasDeployment gets a reference to the given bool and assigns it to the HasDeployment field.
func (o *InstanceType) SetHasDeployment(v bool) {
	o.HasDeployment = &v
}

// GetHasConfig returns the HasConfig field value if set, zero value otherwise.
func (o *InstanceType) GetHasConfig() bool {
	if o == nil || IsNil(o.HasConfig) {
		var ret bool
		return ret
	}
	return *o.HasConfig
}

// GetHasConfigOk returns a tuple with the HasConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetHasConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.HasConfig) {
		return nil, false
	}
	return o.HasConfig, true
}

// IsSetHasConfig returns a boolean if a field has been set.
func (o *InstanceType) IsSetHasConfig() bool {
	if o != nil && !IsNil(o.HasConfig) {
		return true
	}

	return false
}

// SetHasConfig gets a reference to the given bool and assigns it to the HasConfig field.
func (o *InstanceType) SetHasConfig(v bool) {
	o.HasConfig = &v
}

// GetHasSettings returns the HasSettings field value if set, zero value otherwise.
func (o *InstanceType) GetHasSettings() bool {
	if o == nil || IsNil(o.HasSettings) {
		var ret bool
		return ret
	}
	return *o.HasSettings
}

// GetHasSettingsOk returns a tuple with the HasSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetHasSettingsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasSettings) {
		return nil, false
	}
	return o.HasSettings, true
}

// IsSetHasSettings returns a boolean if a field has been set.
func (o *InstanceType) IsSetHasSettings() bool {
	if o != nil && !IsNil(o.HasSettings) {
		return true
	}

	return false
}

// SetHasSettings gets a reference to the given bool and assigns it to the HasSettings field.
func (o *InstanceType) SetHasSettings(v bool) {
	o.HasSettings = &v
}

// GetHasAutoScale returns the HasAutoScale field value if set, zero value otherwise.
func (o *InstanceType) GetHasAutoScale() bool {
	if o == nil || IsNil(o.HasAutoScale) {
		var ret bool
		return ret
	}
	return *o.HasAutoScale
}

// GetHasAutoScaleOk returns a tuple with the HasAutoScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetHasAutoScaleOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAutoScale) {
		return nil, false
	}
	return o.HasAutoScale, true
}

// IsSetHasAutoScale returns a boolean if a field has been set.
func (o *InstanceType) IsSetHasAutoScale() bool {
	if o != nil && !IsNil(o.HasAutoScale) {
		return true
	}

	return false
}

// SetHasAutoScale gets a reference to the given bool and assigns it to the HasAutoScale field.
func (o *InstanceType) SetHasAutoScale(v bool) {
	o.HasAutoScale = &v
}

// GetProxyType returns the ProxyType field value if set, zero value otherwise.
func (o *InstanceType) GetProxyType() string {
	if o == nil || IsNil(o.ProxyType) {
		var ret string
		return ret
	}
	return *o.ProxyType
}

// GetProxyTypeOk returns a tuple with the ProxyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetProxyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyType) {
		return nil, false
	}
	return o.ProxyType, true
}

// IsSetProxyType returns a boolean if a field has been set.
func (o *InstanceType) IsSetProxyType() bool {
	if o != nil && !IsNil(o.ProxyType) {
		return true
	}

	return false
}

// SetProxyType gets a reference to the given string and assigns it to the ProxyType field.
func (o *InstanceType) SetProxyType(v string) {
	o.ProxyType = &v
}

// GetProxyPort returns the ProxyPort field value if set, zero value otherwise.
func (o *InstanceType) GetProxyPort() string {
	if o == nil || IsNil(o.ProxyPort) {
		var ret string
		return ret
	}
	return *o.ProxyPort
}

// GetProxyPortOk returns a tuple with the ProxyPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetProxyPortOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyPort) {
		return nil, false
	}
	return o.ProxyPort, true
}

// IsSetProxyPort returns a boolean if a field has been set.
func (o *InstanceType) IsSetProxyPort() bool {
	if o != nil && !IsNil(o.ProxyPort) {
		return true
	}

	return false
}

// SetProxyPort gets a reference to the given string and assigns it to the ProxyPort field.
func (o *InstanceType) SetProxyPort(v string) {
	o.ProxyPort = &v
}

// GetProxyProtocol returns the ProxyProtocol field value if set, zero value otherwise.
func (o *InstanceType) GetProxyProtocol() string {
	if o == nil || IsNil(o.ProxyProtocol) {
		var ret string
		return ret
	}
	return *o.ProxyProtocol
}

// GetProxyProtocolOk returns a tuple with the ProxyProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetProxyProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyProtocol) {
		return nil, false
	}
	return o.ProxyProtocol, true
}

// IsSetProxyProtocol returns a boolean if a field has been set.
func (o *InstanceType) IsSetProxyProtocol() bool {
	if o != nil && !IsNil(o.ProxyProtocol) {
		return true
	}

	return false
}

// SetProxyProtocol gets a reference to the given string and assigns it to the ProxyProtocol field.
func (o *InstanceType) SetProxyProtocol(v string) {
	o.ProxyProtocol = &v
}

// GetEnvironmentPrefix returns the EnvironmentPrefix field value if set, zero value otherwise.
func (o *InstanceType) GetEnvironmentPrefix() string {
	if o == nil || IsNil(o.EnvironmentPrefix) {
		var ret string
		return ret
	}
	return *o.EnvironmentPrefix
}

// GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetEnvironmentPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentPrefix) {
		return nil, false
	}
	return o.EnvironmentPrefix, true
}

// IsSetEnvironmentPrefix returns a boolean if a field has been set.
func (o *InstanceType) IsSetEnvironmentPrefix() bool {
	if o != nil && !IsNil(o.EnvironmentPrefix) {
		return true
	}

	return false
}

// SetEnvironmentPrefix gets a reference to the given string and assigns it to the EnvironmentPrefix field.
func (o *InstanceType) SetEnvironmentPrefix(v string) {
	o.EnvironmentPrefix = &v
}

// GetBackupType returns the BackupType field value if set, zero value otherwise.
func (o *InstanceType) GetBackupType() string {
	if o == nil || IsNil(o.BackupType) {
		var ret string
		return ret
	}
	return *o.BackupType
}

// GetBackupTypeOk returns a tuple with the BackupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetBackupTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BackupType) {
		return nil, false
	}
	return o.BackupType, true
}

// IsSetBackupType returns a boolean if a field has been set.
func (o *InstanceType) IsSetBackupType() bool {
	if o != nil && !IsNil(o.BackupType) {
		return true
	}

	return false
}

// SetBackupType gets a reference to the given string and assigns it to the BackupType field.
func (o *InstanceType) SetBackupType(v string) {
	o.BackupType = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *InstanceType) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *InstanceType) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *InstanceType) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *InstanceType) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *InstanceType) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *InstanceType) SetVisibility(v string) {
	o.Visibility = &v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *InstanceType) GetFeatured() bool {
	if o == nil || IsNil(o.Featured) {
		var ret bool
		return ret
	}
	return *o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetFeaturedOk() (*bool, bool) {
	if o == nil || IsNil(o.Featured) {
		return nil, false
	}
	return o.Featured, true
}

// IsSetFeatured returns a boolean if a field has been set.
func (o *InstanceType) IsSetFeatured() bool {
	if o != nil && !IsNil(o.Featured) {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given bool and assigns it to the Featured field.
func (o *InstanceType) SetFeatured(v bool) {
	o.Featured = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *InstanceType) GetVersions() []string {
	if o == nil || IsNil(o.Versions) {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// IsSetVersions returns a boolean if a field has been set.
func (o *InstanceType) IsSetVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *InstanceType) SetVersions(v []string) {
	o.Versions = v
}

// GetInstanceTypeLayouts returns the InstanceTypeLayouts field value if set, zero value otherwise.
func (o *InstanceType) GetInstanceTypeLayouts() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner {
	if o == nil || IsNil(o.InstanceTypeLayouts) {
		var ret []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner
		return ret
	}
	return o.InstanceTypeLayouts
}

// GetInstanceTypeLayoutsOk returns a tuple with the InstanceTypeLayouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetInstanceTypeLayoutsOk() ([]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner, bool) {
	if o == nil || IsNil(o.InstanceTypeLayouts) {
		return nil, false
	}
	return o.InstanceTypeLayouts, true
}

// IsSetInstanceTypeLayouts returns a boolean if a field has been set.
func (o *InstanceType) IsSetInstanceTypeLayouts() bool {
	if o != nil && !IsNil(o.InstanceTypeLayouts) {
		return true
	}

	return false
}

// SetInstanceTypeLayouts gets a reference to the given []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner and assigns it to the InstanceTypeLayouts field.
func (o *InstanceType) SetInstanceTypeLayouts(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) {
	o.InstanceTypeLayouts = v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *InstanceType) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner {
	if o == nil || IsNil(o.OptionTypes) {
		var ret []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner
		return ret
	}
	return o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetOptionTypesOk() ([]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool) {
	if o == nil || IsNil(o.OptionTypes) {
		return nil, false
	}
	return o.OptionTypes, true
}

// IsSetOptionTypes returns a boolean if a field has been set.
func (o *InstanceType) IsSetOptionTypes() bool {
	if o != nil && !IsNil(o.OptionTypes) {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner and assigns it to the OptionTypes field.
func (o *InstanceType) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner) {
	o.OptionTypes = v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *InstanceType) GetEnvironmentVariables() []map[string]interface{} {
	if o == nil || IsNil(o.EnvironmentVariables) {
		var ret []map[string]interface{}
		return ret
	}
	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetEnvironmentVariablesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// IsSetEnvironmentVariables returns a boolean if a field has been set.
func (o *InstanceType) IsSetEnvironmentVariables() bool {
	if o != nil && !IsNil(o.EnvironmentVariables) {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given []map[string]interface{} and assigns it to the EnvironmentVariables field.
func (o *InstanceType) SetEnvironmentVariables(v []map[string]interface{}) {
	o.EnvironmentVariables = v
}

// GetPriceSets returns the PriceSets field value if set, zero value otherwise.
func (o *InstanceType) GetPriceSets() []map[string]interface{} {
	if o == nil || IsNil(o.PriceSets) {
		var ret []map[string]interface{}
		return ret
	}
	return o.PriceSets
}

// GetPriceSetsOk returns a tuple with the PriceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetPriceSetsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.PriceSets) {
		return nil, false
	}
	return o.PriceSets, true
}

// IsSetPriceSets returns a boolean if a field has been set.
func (o *InstanceType) IsSetPriceSets() bool {
	if o != nil && !IsNil(o.PriceSets) {
		return true
	}

	return false
}

// SetPriceSets gets a reference to the given []map[string]interface{} and assigns it to the PriceSets field.
func (o *InstanceType) SetPriceSets(v []map[string]interface{}) {
	o.PriceSets = v
}

// GetImagePath returns the ImagePath field value if set, zero value otherwise.
func (o *InstanceType) GetImagePath() string {
	if o == nil || IsNil(o.ImagePath) {
		var ret string
		return ret
	}
	return *o.ImagePath
}

// GetImagePathOk returns a tuple with the ImagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetImagePathOk() (*string, bool) {
	if o == nil || IsNil(o.ImagePath) {
		return nil, false
	}
	return o.ImagePath, true
}

// IsSetImagePath returns a boolean if a field has been set.
func (o *InstanceType) IsSetImagePath() bool {
	if o != nil && !IsNil(o.ImagePath) {
		return true
	}

	return false
}

// SetImagePath gets a reference to the given string and assigns it to the ImagePath field.
func (o *InstanceType) SetImagePath(v string) {
	o.ImagePath = &v
}

// GetDarkImagePath returns the DarkImagePath field value if set, zero value otherwise.
func (o *InstanceType) GetDarkImagePath() string {
	if o == nil || IsNil(o.DarkImagePath) {
		var ret string
		return ret
	}
	return *o.DarkImagePath
}

// GetDarkImagePathOk returns a tuple with the DarkImagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceType) GetDarkImagePathOk() (*string, bool) {
	if o == nil || IsNil(o.DarkImagePath) {
		return nil, false
	}
	return o.DarkImagePath, true
}

// IsSetDarkImagePath returns a boolean if a field has been set.
func (o *InstanceType) IsSetDarkImagePath() bool {
	if o != nil && !IsNil(o.DarkImagePath) {
		return true
	}

	return false
}

// SetDarkImagePath gets a reference to the given string and assigns it to the DarkImagePath field.
func (o *InstanceType) SetDarkImagePath(v string) {
	o.DarkImagePath = &v
}

func (o InstanceType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ProvisionTypeCode) {
		toSerialize["provisionTypeCode"] = o.ProvisionTypeCode
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.HasProvisioningStep) {
		toSerialize["hasProvisioningStep"] = o.HasProvisioningStep
	}
	if !IsNil(o.HasDeployment) {
		toSerialize["hasDeployment"] = o.HasDeployment
	}
	if !IsNil(o.HasConfig) {
		toSerialize["hasConfig"] = o.HasConfig
	}
	if !IsNil(o.HasSettings) {
		toSerialize["hasSettings"] = o.HasSettings
	}
	if !IsNil(o.HasAutoScale) {
		toSerialize["hasAutoScale"] = o.HasAutoScale
	}
	if !IsNil(o.ProxyType) {
		toSerialize["proxyType"] = o.ProxyType
	}
	if !IsNil(o.ProxyPort) {
		toSerialize["proxyPort"] = o.ProxyPort
	}
	if !IsNil(o.ProxyProtocol) {
		toSerialize["proxyProtocol"] = o.ProxyProtocol
	}
	if !IsNil(o.EnvironmentPrefix) {
		toSerialize["environmentPrefix"] = o.EnvironmentPrefix
	}
	if !IsNil(o.BackupType) {
		toSerialize["backupType"] = o.BackupType
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.Featured) {
		toSerialize["featured"] = o.Featured
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.InstanceTypeLayouts) {
		toSerialize["instanceTypeLayouts"] = o.InstanceTypeLayouts
	}
	if !IsNil(o.OptionTypes) {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	if !IsNil(o.EnvironmentVariables) {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	if !IsNil(o.PriceSets) {
		toSerialize["priceSets"] = o.PriceSets
	}
	if !IsNil(o.ImagePath) {
		toSerialize["imagePath"] = o.ImagePath
	}
	if !IsNil(o.DarkImagePath) {
		toSerialize["darkImagePath"] = o.DarkImagePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *InstanceType) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

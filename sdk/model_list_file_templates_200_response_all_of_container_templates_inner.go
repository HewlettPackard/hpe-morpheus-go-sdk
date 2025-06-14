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

// checks if the ListFileTemplates200ResponseAllOfContainerTemplatesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFileTemplates200ResponseAllOfContainerTemplatesInner{}

// ListFileTemplates200ResponseAllOfContainerTemplatesInner struct for ListFileTemplates200ResponseAllOfContainerTemplatesInner
type ListFileTemplates200ResponseAllOfContainerTemplatesInner struct {
	Id                   *int64                                             `json:"id,omitempty"`
	Code                 *string                                            `json:"code,omitempty"`
	Account              *GetAlerts200ResponseAllOfCheckGroupsInnerInstance `json:"account,omitempty"`
	Name                 *string                                            `json:"name,omitempty"`
	Labels               []string                                           `json:"labels,omitempty"`
	FileName             *string                                            `json:"fileName,omitempty"`
	FilePath             *string                                            `json:"filePath,omitempty"`
	TemplateType         *string                                            `json:"templateType,omitempty"`
	TemplatePhase        *string                                            `json:"templatePhase,omitempty"`
	Template             *string                                            `json:"template,omitempty"`
	Category             *string                                            `json:"category,omitempty"`
	SettingCategory      *string                                            `json:"settingCategory,omitempty"`
	SettingName          *string                                            `json:"settingName,omitempty"`
	AutoRun              *bool                                              `json:"autoRun,omitempty"`
	RunOnScale           *bool                                              `json:"runOnScale,omitempty"`
	RunOnDeploy          *bool                                              `json:"runOnDeploy,omitempty"`
	FileOwner            *string                                            `json:"fileOwner,omitempty"`
	FileGroup            *string                                            `json:"fileGroup,omitempty"`
	Permissions          *string                                            `json:"permissions,omitempty"`
	DateCreated          *time.Time                                         `json:"dateCreated,omitempty"`
	LastUpdated          *time.Time                                         `json:"lastUpdated,omitempty"`
	AdditionalProperties map[string]interface{}                             `json:",remain"`
}

type _ListFileTemplates200ResponseAllOfContainerTemplatesInner ListFileTemplates200ResponseAllOfContainerTemplatesInner

// NewListFileTemplates200ResponseAllOfContainerTemplatesInner instantiates a new ListFileTemplates200ResponseAllOfContainerTemplatesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFileTemplates200ResponseAllOfContainerTemplatesInner() *ListFileTemplates200ResponseAllOfContainerTemplatesInner {
	this := ListFileTemplates200ResponseAllOfContainerTemplatesInner{}
	return &this
}

// NewListFileTemplates200ResponseAllOfContainerTemplatesInnerWithDefaults instantiates a new ListFileTemplates200ResponseAllOfContainerTemplatesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFileTemplates200ResponseAllOfContainerTemplatesInnerWithDefaults() *ListFileTemplates200ResponseAllOfContainerTemplatesInner {
	this := ListFileTemplates200ResponseAllOfContainerTemplatesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetCode(v string) {
	o.Code = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance {
	if o == nil || IsNil(o.Account) {
		var ret GetAlerts200ResponseAllOfCheckGroupsInnerInstance
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// IsSetAccount returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given GetAlerts200ResponseAllOfCheckGroupsInnerInstance and assigns it to the Account field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance) {
	o.Account = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetLabels(v []string) {
	o.Labels = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// IsSetFileName returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetFileName(v string) {
	o.FileName = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// IsSetFilePath returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetFilePath(v string) {
	o.FilePath = &v
}

// GetTemplateType returns the TemplateType field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetTemplateType() string {
	if o == nil || IsNil(o.TemplateType) {
		var ret string
		return ret
	}
	return *o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetTemplateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateType) {
		return nil, false
	}
	return o.TemplateType, true
}

// IsSetTemplateType returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetTemplateType() bool {
	if o != nil && !IsNil(o.TemplateType) {
		return true
	}

	return false
}

// SetTemplateType gets a reference to the given string and assigns it to the TemplateType field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetTemplateType(v string) {
	o.TemplateType = &v
}

// GetTemplatePhase returns the TemplatePhase field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetTemplatePhase() string {
	if o == nil || IsNil(o.TemplatePhase) {
		var ret string
		return ret
	}
	return *o.TemplatePhase
}

// GetTemplatePhaseOk returns a tuple with the TemplatePhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetTemplatePhaseOk() (*string, bool) {
	if o == nil || IsNil(o.TemplatePhase) {
		return nil, false
	}
	return o.TemplatePhase, true
}

// IsSetTemplatePhase returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetTemplatePhase() bool {
	if o != nil && !IsNil(o.TemplatePhase) {
		return true
	}

	return false
}

// SetTemplatePhase gets a reference to the given string and assigns it to the TemplatePhase field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetTemplatePhase(v string) {
	o.TemplatePhase = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// IsSetTemplate returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetTemplate(v string) {
	o.Template = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// IsSetCategory returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetCategory(v string) {
	o.Category = &v
}

// GetSettingCategory returns the SettingCategory field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetSettingCategory() string {
	if o == nil || IsNil(o.SettingCategory) {
		var ret string
		return ret
	}
	return *o.SettingCategory
}

// GetSettingCategoryOk returns a tuple with the SettingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetSettingCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.SettingCategory) {
		return nil, false
	}
	return o.SettingCategory, true
}

// IsSetSettingCategory returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetSettingCategory() bool {
	if o != nil && !IsNil(o.SettingCategory) {
		return true
	}

	return false
}

// SetSettingCategory gets a reference to the given string and assigns it to the SettingCategory field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetSettingCategory(v string) {
	o.SettingCategory = &v
}

// GetSettingName returns the SettingName field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetSettingName() string {
	if o == nil || IsNil(o.SettingName) {
		var ret string
		return ret
	}
	return *o.SettingName
}

// GetSettingNameOk returns a tuple with the SettingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetSettingNameOk() (*string, bool) {
	if o == nil || IsNil(o.SettingName) {
		return nil, false
	}
	return o.SettingName, true
}

// IsSetSettingName returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetSettingName() bool {
	if o != nil && !IsNil(o.SettingName) {
		return true
	}

	return false
}

// SetSettingName gets a reference to the given string and assigns it to the SettingName field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetSettingName(v string) {
	o.SettingName = &v
}

// GetAutoRun returns the AutoRun field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetAutoRun() bool {
	if o == nil || IsNil(o.AutoRun) {
		var ret bool
		return ret
	}
	return *o.AutoRun
}

// GetAutoRunOk returns a tuple with the AutoRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetAutoRunOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRun) {
		return nil, false
	}
	return o.AutoRun, true
}

// IsSetAutoRun returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetAutoRun() bool {
	if o != nil && !IsNil(o.AutoRun) {
		return true
	}

	return false
}

// SetAutoRun gets a reference to the given bool and assigns it to the AutoRun field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetAutoRun(v bool) {
	o.AutoRun = &v
}

// GetRunOnScale returns the RunOnScale field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetRunOnScale() bool {
	if o == nil || IsNil(o.RunOnScale) {
		var ret bool
		return ret
	}
	return *o.RunOnScale
}

// GetRunOnScaleOk returns a tuple with the RunOnScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetRunOnScaleOk() (*bool, bool) {
	if o == nil || IsNil(o.RunOnScale) {
		return nil, false
	}
	return o.RunOnScale, true
}

// IsSetRunOnScale returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetRunOnScale() bool {
	if o != nil && !IsNil(o.RunOnScale) {
		return true
	}

	return false
}

// SetRunOnScale gets a reference to the given bool and assigns it to the RunOnScale field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetRunOnScale(v bool) {
	o.RunOnScale = &v
}

// GetRunOnDeploy returns the RunOnDeploy field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetRunOnDeploy() bool {
	if o == nil || IsNil(o.RunOnDeploy) {
		var ret bool
		return ret
	}
	return *o.RunOnDeploy
}

// GetRunOnDeployOk returns a tuple with the RunOnDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetRunOnDeployOk() (*bool, bool) {
	if o == nil || IsNil(o.RunOnDeploy) {
		return nil, false
	}
	return o.RunOnDeploy, true
}

// IsSetRunOnDeploy returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetRunOnDeploy() bool {
	if o != nil && !IsNil(o.RunOnDeploy) {
		return true
	}

	return false
}

// SetRunOnDeploy gets a reference to the given bool and assigns it to the RunOnDeploy field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetRunOnDeploy(v bool) {
	o.RunOnDeploy = &v
}

// GetFileOwner returns the FileOwner field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFileOwner() string {
	if o == nil || IsNil(o.FileOwner) {
		var ret string
		return ret
	}
	return *o.FileOwner
}

// GetFileOwnerOk returns a tuple with the FileOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFileOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.FileOwner) {
		return nil, false
	}
	return o.FileOwner, true
}

// IsSetFileOwner returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetFileOwner() bool {
	if o != nil && !IsNil(o.FileOwner) {
		return true
	}

	return false
}

// SetFileOwner gets a reference to the given string and assigns it to the FileOwner field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetFileOwner(v string) {
	o.FileOwner = &v
}

// GetFileGroup returns the FileGroup field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFileGroup() string {
	if o == nil || IsNil(o.FileGroup) {
		var ret string
		return ret
	}
	return *o.FileGroup
}

// GetFileGroupOk returns a tuple with the FileGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFileGroupOk() (*string, bool) {
	if o == nil || IsNil(o.FileGroup) {
		return nil, false
	}
	return o.FileGroup, true
}

// IsSetFileGroup returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetFileGroup() bool {
	if o != nil && !IsNil(o.FileGroup) {
		return true
	}

	return false
}

// SetFileGroup gets a reference to the given string and assigns it to the FileGroup field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetFileGroup(v string) {
	o.FileGroup = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetPermissions() string {
	if o == nil || IsNil(o.Permissions) {
		var ret string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetPermissionsOk() (*string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// IsSetPermissions returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given string and assigns it to the Permissions field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetPermissions(v string) {
	o.Permissions = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o ListFileTemplates200ResponseAllOfContainerTemplatesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFileTemplates200ResponseAllOfContainerTemplatesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
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
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.FilePath) {
		toSerialize["filePath"] = o.FilePath
	}
	if !IsNil(o.TemplateType) {
		toSerialize["templateType"] = o.TemplateType
	}
	if !IsNil(o.TemplatePhase) {
		toSerialize["templatePhase"] = o.TemplatePhase
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.SettingCategory) {
		toSerialize["settingCategory"] = o.SettingCategory
	}
	if !IsNil(o.SettingName) {
		toSerialize["settingName"] = o.SettingName
	}
	if !IsNil(o.AutoRun) {
		toSerialize["autoRun"] = o.AutoRun
	}
	if !IsNil(o.RunOnScale) {
		toSerialize["runOnScale"] = o.RunOnScale
	}
	if !IsNil(o.RunOnDeploy) {
		toSerialize["runOnDeploy"] = o.RunOnDeploy
	}
	if !IsNil(o.FileOwner) {
		toSerialize["fileOwner"] = o.FileOwner
	}
	if !IsNil(o.FileGroup) {
		toSerialize["fileGroup"] = o.FileGroup
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

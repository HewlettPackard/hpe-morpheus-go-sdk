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

// checks if the UpdateFileTemplateRequestContainerTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFileTemplateRequestContainerTemplate{}

// UpdateFileTemplateRequestContainerTemplate struct for UpdateFileTemplateRequestContainerTemplate
type UpdateFileTemplateRequestContainerTemplate struct {
	// File template name
	Name *string `json:"name,omitempty"`
	// Array of label strings, can be used for filtering.
	Labels []string `json:"labels,omitempty"`
	// Filename for the file template
	FileName *string `json:"fileName,omitempty"`
	// Path for the file template
	FilePath *string `json:"filePath,omitempty"`
	// Category
	Category *string `json:"category,omitempty"`
	// Template Phase, provision, start, etc.
	TemplatePhase *string `json:"templatePhase,omitempty"`
	// Template content, that is, the file template content itself.
	Template *string `json:"template,omitempty"`
	// File Owner
	FileOwner *int64 `json:"fileOwner,omitempty"`
	// Setting Name
	SettingName *string `json:"settingName,omitempty"`
	// Setting Category
	SettingCategory      *string                `json:"settingCategory,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _UpdateFileTemplateRequestContainerTemplate UpdateFileTemplateRequestContainerTemplate

// NewUpdateFileTemplateRequestContainerTemplate instantiates a new UpdateFileTemplateRequestContainerTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFileTemplateRequestContainerTemplate() *UpdateFileTemplateRequestContainerTemplate {
	this := UpdateFileTemplateRequestContainerTemplate{}
	return &this
}

// NewUpdateFileTemplateRequestContainerTemplateWithDefaults instantiates a new UpdateFileTemplateRequestContainerTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFileTemplateRequestContainerTemplateWithDefaults() *UpdateFileTemplateRequestContainerTemplate {
	this := UpdateFileTemplateRequestContainerTemplate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateFileTemplateRequestContainerTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateFileTemplateRequestContainerTemplate) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateFileTemplateRequestContainerTemplate) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *UpdateFileTemplateRequestContainerTemplate) SetLabels(v []string) {
	o.Labels = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *UpdateFileTemplateRequestContainerTemplate) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// IsSetFileName returns a boolean if a field has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) IsSetFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *UpdateFileTemplateRequestContainerTemplate) SetFileName(v string) {
	o.FileName = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *UpdateFileTemplateRequestContainerTemplate) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// IsSetFilePath returns a boolean if a field has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) IsSetFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *UpdateFileTemplateRequestContainerTemplate) SetFilePath(v string) {
	o.FilePath = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *UpdateFileTemplateRequestContainerTemplate) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// IsSetCategory returns a boolean if a field has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) IsSetCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *UpdateFileTemplateRequestContainerTemplate) SetCategory(v string) {
	o.Category = &v
}

// GetTemplatePhase returns the TemplatePhase field value if set, zero value otherwise.
func (o *UpdateFileTemplateRequestContainerTemplate) GetTemplatePhase() string {
	if o == nil || IsNil(o.TemplatePhase) {
		var ret string
		return ret
	}
	return *o.TemplatePhase
}

// GetTemplatePhaseOk returns a tuple with the TemplatePhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) GetTemplatePhaseOk() (*string, bool) {
	if o == nil || IsNil(o.TemplatePhase) {
		return nil, false
	}
	return o.TemplatePhase, true
}

// IsSetTemplatePhase returns a boolean if a field has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) IsSetTemplatePhase() bool {
	if o != nil && !IsNil(o.TemplatePhase) {
		return true
	}

	return false
}

// SetTemplatePhase gets a reference to the given string and assigns it to the TemplatePhase field.
func (o *UpdateFileTemplateRequestContainerTemplate) SetTemplatePhase(v string) {
	o.TemplatePhase = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *UpdateFileTemplateRequestContainerTemplate) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// IsSetTemplate returns a boolean if a field has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) IsSetTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *UpdateFileTemplateRequestContainerTemplate) SetTemplate(v string) {
	o.Template = &v
}

// GetFileOwner returns the FileOwner field value if set, zero value otherwise.
func (o *UpdateFileTemplateRequestContainerTemplate) GetFileOwner() int64 {
	if o == nil || IsNil(o.FileOwner) {
		var ret int64
		return ret
	}
	return *o.FileOwner
}

// GetFileOwnerOk returns a tuple with the FileOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) GetFileOwnerOk() (*int64, bool) {
	if o == nil || IsNil(o.FileOwner) {
		return nil, false
	}
	return o.FileOwner, true
}

// IsSetFileOwner returns a boolean if a field has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) IsSetFileOwner() bool {
	if o != nil && !IsNil(o.FileOwner) {
		return true
	}

	return false
}

// SetFileOwner gets a reference to the given int64 and assigns it to the FileOwner field.
func (o *UpdateFileTemplateRequestContainerTemplate) SetFileOwner(v int64) {
	o.FileOwner = &v
}

// GetSettingName returns the SettingName field value if set, zero value otherwise.
func (o *UpdateFileTemplateRequestContainerTemplate) GetSettingName() string {
	if o == nil || IsNil(o.SettingName) {
		var ret string
		return ret
	}
	return *o.SettingName
}

// GetSettingNameOk returns a tuple with the SettingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) GetSettingNameOk() (*string, bool) {
	if o == nil || IsNil(o.SettingName) {
		return nil, false
	}
	return o.SettingName, true
}

// IsSetSettingName returns a boolean if a field has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) IsSetSettingName() bool {
	if o != nil && !IsNil(o.SettingName) {
		return true
	}

	return false
}

// SetSettingName gets a reference to the given string and assigns it to the SettingName field.
func (o *UpdateFileTemplateRequestContainerTemplate) SetSettingName(v string) {
	o.SettingName = &v
}

// GetSettingCategory returns the SettingCategory field value if set, zero value otherwise.
func (o *UpdateFileTemplateRequestContainerTemplate) GetSettingCategory() string {
	if o == nil || IsNil(o.SettingCategory) {
		var ret string
		return ret
	}
	return *o.SettingCategory
}

// GetSettingCategoryOk returns a tuple with the SettingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) GetSettingCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.SettingCategory) {
		return nil, false
	}
	return o.SettingCategory, true
}

// IsSetSettingCategory returns a boolean if a field has been set.
func (o *UpdateFileTemplateRequestContainerTemplate) IsSetSettingCategory() bool {
	if o != nil && !IsNil(o.SettingCategory) {
		return true
	}

	return false
}

// SetSettingCategory gets a reference to the given string and assigns it to the SettingCategory field.
func (o *UpdateFileTemplateRequestContainerTemplate) SetSettingCategory(v string) {
	o.SettingCategory = &v
}

func (o UpdateFileTemplateRequestContainerTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFileTemplateRequestContainerTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.TemplatePhase) {
		toSerialize["templatePhase"] = o.TemplatePhase
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.FileOwner) {
		toSerialize["fileOwner"] = o.FileOwner
	}
	if !IsNil(o.SettingName) {
		toSerialize["settingName"] = o.SettingName
	}
	if !IsNil(o.SettingCategory) {
		toSerialize["settingCategory"] = o.SettingCategory
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateFileTemplateRequestContainerTemplate) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

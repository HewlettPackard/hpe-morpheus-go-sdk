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

// checks if the InstanceTypeUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceTypeUpdate{}

// InstanceTypeUpdate struct for InstanceTypeUpdate
type InstanceTypeUpdate struct {
	// Instance type name
	Name *string `json:"name,omitempty"`
	// Instance type description
	Description *string `json:"description,omitempty"`
	// Array of label strings, can be used for filtering.
	Labels []string `json:"labels,omitempty"`
	// Instance type code
	Code *string `json:"code,omitempty"`
	// Category
	Category *string `json:"category,omitempty"`
	// Visibility
	Visibility *string `json:"visibility,omitempty"`
	// Featured
	Featured *bool `json:"featured,omitempty"`
	// Enable Settings
	HasSettings *bool `json:"hasSettings,omitempty"`
	// Enable Scaling (Horizontal)
	HasAutoScale *bool `json:"hasAutoScale,omitempty"`
	// Supports Deployments
	HasDeployment *bool `json:"hasDeployment,omitempty"`
	// Environment Prefix, can be used to make exported evars unique.
	EnvironmentPrefix *string `json:"environmentPrefix,omitempty"`
	// Array of instance type env variables.
	EnvironmentVariables []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner `json:"environmentVariables,omitempty"`
	// Array of price set objects
	PriceSets []AddInstanceTypeRequestInstanceTypePriceSetsInner `json:"priceSets,omitempty"`
	// Array of instance type option type IDs
	OptionTypes          []int64                `json:"optionTypes,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _InstanceTypeUpdate InstanceTypeUpdate

// NewInstanceTypeUpdate instantiates a new InstanceTypeUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceTypeUpdate() *InstanceTypeUpdate {
	this := InstanceTypeUpdate{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// NewInstanceTypeUpdateWithDefaults instantiates a new InstanceTypeUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceTypeUpdateWithDefaults() *InstanceTypeUpdate {
	this := InstanceTypeUpdate{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceTypeUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InstanceTypeUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *InstanceTypeUpdate) SetLabels(v []string) {
	o.Labels = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InstanceTypeUpdate) SetCode(v string) {
	o.Code = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// IsSetCategory returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InstanceTypeUpdate) SetCategory(v string) {
	o.Category = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *InstanceTypeUpdate) SetVisibility(v string) {
	o.Visibility = &v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetFeatured() bool {
	if o == nil || IsNil(o.Featured) {
		var ret bool
		return ret
	}
	return *o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetFeaturedOk() (*bool, bool) {
	if o == nil || IsNil(o.Featured) {
		return nil, false
	}
	return o.Featured, true
}

// IsSetFeatured returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetFeatured() bool {
	if o != nil && !IsNil(o.Featured) {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given bool and assigns it to the Featured field.
func (o *InstanceTypeUpdate) SetFeatured(v bool) {
	o.Featured = &v
}

// GetHasSettings returns the HasSettings field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetHasSettings() bool {
	if o == nil || IsNil(o.HasSettings) {
		var ret bool
		return ret
	}
	return *o.HasSettings
}

// GetHasSettingsOk returns a tuple with the HasSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetHasSettingsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasSettings) {
		return nil, false
	}
	return o.HasSettings, true
}

// IsSetHasSettings returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetHasSettings() bool {
	if o != nil && !IsNil(o.HasSettings) {
		return true
	}

	return false
}

// SetHasSettings gets a reference to the given bool and assigns it to the HasSettings field.
func (o *InstanceTypeUpdate) SetHasSettings(v bool) {
	o.HasSettings = &v
}

// GetHasAutoScale returns the HasAutoScale field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetHasAutoScale() bool {
	if o == nil || IsNil(o.HasAutoScale) {
		var ret bool
		return ret
	}
	return *o.HasAutoScale
}

// GetHasAutoScaleOk returns a tuple with the HasAutoScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetHasAutoScaleOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAutoScale) {
		return nil, false
	}
	return o.HasAutoScale, true
}

// IsSetHasAutoScale returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetHasAutoScale() bool {
	if o != nil && !IsNil(o.HasAutoScale) {
		return true
	}

	return false
}

// SetHasAutoScale gets a reference to the given bool and assigns it to the HasAutoScale field.
func (o *InstanceTypeUpdate) SetHasAutoScale(v bool) {
	o.HasAutoScale = &v
}

// GetHasDeployment returns the HasDeployment field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetHasDeployment() bool {
	if o == nil || IsNil(o.HasDeployment) {
		var ret bool
		return ret
	}
	return *o.HasDeployment
}

// GetHasDeploymentOk returns a tuple with the HasDeployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetHasDeploymentOk() (*bool, bool) {
	if o == nil || IsNil(o.HasDeployment) {
		return nil, false
	}
	return o.HasDeployment, true
}

// IsSetHasDeployment returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetHasDeployment() bool {
	if o != nil && !IsNil(o.HasDeployment) {
		return true
	}

	return false
}

// SetHasDeployment gets a reference to the given bool and assigns it to the HasDeployment field.
func (o *InstanceTypeUpdate) SetHasDeployment(v bool) {
	o.HasDeployment = &v
}

// GetEnvironmentPrefix returns the EnvironmentPrefix field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetEnvironmentPrefix() string {
	if o == nil || IsNil(o.EnvironmentPrefix) {
		var ret string
		return ret
	}
	return *o.EnvironmentPrefix
}

// GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetEnvironmentPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentPrefix) {
		return nil, false
	}
	return o.EnvironmentPrefix, true
}

// IsSetEnvironmentPrefix returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetEnvironmentPrefix() bool {
	if o != nil && !IsNil(o.EnvironmentPrefix) {
		return true
	}

	return false
}

// SetEnvironmentPrefix gets a reference to the given string and assigns it to the EnvironmentPrefix field.
func (o *InstanceTypeUpdate) SetEnvironmentPrefix(v string) {
	o.EnvironmentPrefix = &v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetEnvironmentVariables() []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner {
	if o == nil || IsNil(o.EnvironmentVariables) {
		var ret []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner
		return ret
	}
	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetEnvironmentVariablesOk() ([]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// IsSetEnvironmentVariables returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetEnvironmentVariables() bool {
	if o != nil && !IsNil(o.EnvironmentVariables) {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner and assigns it to the EnvironmentVariables field.
func (o *InstanceTypeUpdate) SetEnvironmentVariables(v []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner) {
	o.EnvironmentVariables = v
}

// GetPriceSets returns the PriceSets field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetPriceSets() []AddInstanceTypeRequestInstanceTypePriceSetsInner {
	if o == nil || IsNil(o.PriceSets) {
		var ret []AddInstanceTypeRequestInstanceTypePriceSetsInner
		return ret
	}
	return o.PriceSets
}

// GetPriceSetsOk returns a tuple with the PriceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetPriceSetsOk() ([]AddInstanceTypeRequestInstanceTypePriceSetsInner, bool) {
	if o == nil || IsNil(o.PriceSets) {
		return nil, false
	}
	return o.PriceSets, true
}

// IsSetPriceSets returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetPriceSets() bool {
	if o != nil && !IsNil(o.PriceSets) {
		return true
	}

	return false
}

// SetPriceSets gets a reference to the given []AddInstanceTypeRequestInstanceTypePriceSetsInner and assigns it to the PriceSets field.
func (o *InstanceTypeUpdate) SetPriceSets(v []AddInstanceTypeRequestInstanceTypePriceSetsInner) {
	o.PriceSets = v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetOptionTypes() []int64 {
	if o == nil || IsNil(o.OptionTypes) {
		var ret []int64
		return ret
	}
	return o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetOptionTypesOk() ([]int64, bool) {
	if o == nil || IsNil(o.OptionTypes) {
		return nil, false
	}
	return o.OptionTypes, true
}

// IsSetOptionTypes returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) IsSetOptionTypes() bool {
	if o != nil && !IsNil(o.OptionTypes) {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []int64 and assigns it to the OptionTypes field.
func (o *InstanceTypeUpdate) SetOptionTypes(v []int64) {
	o.OptionTypes = v
}

func (o InstanceTypeUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceTypeUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.Featured) {
		toSerialize["featured"] = o.Featured
	}
	if !IsNil(o.HasSettings) {
		toSerialize["hasSettings"] = o.HasSettings
	}
	if !IsNil(o.HasAutoScale) {
		toSerialize["hasAutoScale"] = o.HasAutoScale
	}
	if !IsNil(o.HasDeployment) {
		toSerialize["hasDeployment"] = o.HasDeployment
	}
	if !IsNil(o.EnvironmentPrefix) {
		toSerialize["environmentPrefix"] = o.EnvironmentPrefix
	}
	if !IsNil(o.EnvironmentVariables) {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	if !IsNil(o.PriceSets) {
		toSerialize["priceSets"] = o.PriceSets
	}
	if !IsNil(o.OptionTypes) {
		toSerialize["optionTypes"] = o.OptionTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *InstanceTypeUpdate) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

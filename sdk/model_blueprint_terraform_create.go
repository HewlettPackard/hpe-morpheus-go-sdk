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

// checks if the BlueprintTerraformCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintTerraformCreate{}

// BlueprintTerraformCreate struct for BlueprintTerraformCreate
type BlueprintTerraformCreate struct {
	// A name for the blueprint
	Name string `json:"name"`
	// Path to display image. Defaults to an internal Morpheus image.
	Image *string `json:"image,omitempty"`
	// Blueprint Type
	Type string `json:"type"`
	// Array of label strings, can be used for filtering.
	Labels               []string                           `json:"labels,omitempty"`
	Terraform            AddBlueprintRequestOneOf5Terraform `json:"terraform"`
	Config               *AddBlueprintRequestOneOf5Config   `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}             `json:",remain"`
}

type _BlueprintTerraformCreate BlueprintTerraformCreate

// NewBlueprintTerraformCreate instantiates a new BlueprintTerraformCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintTerraformCreate(name string, type_ string, terraform AddBlueprintRequestOneOf5Terraform) *BlueprintTerraformCreate {
	this := BlueprintTerraformCreate{}
	this.Name = name
	this.Type = type_
	this.Terraform = terraform
	return &this
}

// NewBlueprintTerraformCreateWithDefaults instantiates a new BlueprintTerraformCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintTerraformCreateWithDefaults() *BlueprintTerraformCreate {
	this := BlueprintTerraformCreate{}
	return &this
}

// GetName returns the Name field value
func (o *BlueprintTerraformCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BlueprintTerraformCreate) SetName(v string) {
	o.Name = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BlueprintTerraformCreate) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreate) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// IsSetImage returns a boolean if a field has been set.
func (o *BlueprintTerraformCreate) IsSetImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BlueprintTerraformCreate) SetImage(v string) {
	o.Image = &v
}

// GetType returns the Type field value
func (o *BlueprintTerraformCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BlueprintTerraformCreate) SetType(v string) {
	o.Type = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *BlueprintTerraformCreate) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreate) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *BlueprintTerraformCreate) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *BlueprintTerraformCreate) SetLabels(v []string) {
	o.Labels = v
}

// GetTerraform returns the Terraform field value
func (o *BlueprintTerraformCreate) GetTerraform() AddBlueprintRequestOneOf5Terraform {
	if o == nil {
		var ret AddBlueprintRequestOneOf5Terraform
		return ret
	}

	return o.Terraform
}

// GetTerraformOk returns a tuple with the Terraform field value
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreate) GetTerraformOk() (*AddBlueprintRequestOneOf5Terraform, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Terraform, true
}

// SetTerraform sets field value
func (o *BlueprintTerraformCreate) SetTerraform(v AddBlueprintRequestOneOf5Terraform) {
	o.Terraform = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *BlueprintTerraformCreate) GetConfig() AddBlueprintRequestOneOf5Config {
	if o == nil || IsNil(o.Config) {
		var ret AddBlueprintRequestOneOf5Config
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreate) GetConfigOk() (*AddBlueprintRequestOneOf5Config, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *BlueprintTerraformCreate) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given AddBlueprintRequestOneOf5Config and assigns it to the Config field.
func (o *BlueprintTerraformCreate) SetConfig(v AddBlueprintRequestOneOf5Config) {
	o.Config = &v
}

func (o BlueprintTerraformCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintTerraformCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["terraform"] = o.Terraform
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *BlueprintTerraformCreate) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

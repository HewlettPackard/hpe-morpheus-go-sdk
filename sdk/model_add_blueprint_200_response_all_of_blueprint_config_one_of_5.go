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

// checks if the AddBlueprint200ResponseAllOfBlueprintConfigOneOf5 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddBlueprint200ResponseAllOfBlueprintConfigOneOf5{}

// AddBlueprint200ResponseAllOfBlueprintConfigOneOf5 struct for AddBlueprint200ResponseAllOfBlueprintConfigOneOf5
type AddBlueprint200ResponseAllOfBlueprintConfigOneOf5 struct {
	// A name for the blueprint
	Name *string `json:"name,omitempty"`
	// Path to display image. Defaults to an internal Morpheus image.
	Image *string `json:"image,omitempty"`
	// Blueprint Type
	Type      *string                             `json:"type,omitempty"`
	Terraform *AddBlueprintRequestOneOf5Terraform `json:"terraform,omitempty"`
	Config    *AddBlueprintRequestOneOf5Config    `json:"config,omitempty"`
	// Private or Public Access
	Visibility *string `json:"visibility,omitempty"`
	// Resource Permission Block
	ResourcePermission map[string]interface{} `json:"resourcePermission,omitempty"`
	// Owner
	Owner map[string]interface{} `json:"owner,omitempty"`
	// Tenant
	Tenant               map[string]interface{} `json:"tenant,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddBlueprint200ResponseAllOfBlueprintConfigOneOf5 AddBlueprint200ResponseAllOfBlueprintConfigOneOf5

// NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf5 instantiates a new AddBlueprint200ResponseAllOfBlueprintConfigOneOf5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf5() *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5 {
	this := AddBlueprint200ResponseAllOfBlueprintConfigOneOf5{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf5WithDefaults instantiates a new AddBlueprint200ResponseAllOfBlueprintConfigOneOf5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf5WithDefaults() *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5 {
	this := AddBlueprint200ResponseAllOfBlueprintConfigOneOf5{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) SetName(v string) {
	o.Name = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// IsSetImage returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) IsSetImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) SetImage(v string) {
	o.Image = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) SetType(v string) {
	o.Type = &v
}

// GetTerraform returns the Terraform field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetTerraform() AddBlueprintRequestOneOf5Terraform {
	if o == nil || IsNil(o.Terraform) {
		var ret AddBlueprintRequestOneOf5Terraform
		return ret
	}
	return *o.Terraform
}

// GetTerraformOk returns a tuple with the Terraform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetTerraformOk() (*AddBlueprintRequestOneOf5Terraform, bool) {
	if o == nil || IsNil(o.Terraform) {
		return nil, false
	}
	return o.Terraform, true
}

// IsSetTerraform returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) IsSetTerraform() bool {
	if o != nil && !IsNil(o.Terraform) {
		return true
	}

	return false
}

// SetTerraform gets a reference to the given AddBlueprintRequestOneOf5Terraform and assigns it to the Terraform field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) SetTerraform(v AddBlueprintRequestOneOf5Terraform) {
	o.Terraform = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetConfig() AddBlueprintRequestOneOf5Config {
	if o == nil || IsNil(o.Config) {
		var ret AddBlueprintRequestOneOf5Config
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetConfigOk() (*AddBlueprintRequestOneOf5Config, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given AddBlueprintRequestOneOf5Config and assigns it to the Config field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) SetConfig(v AddBlueprintRequestOneOf5Config) {
	o.Config = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) SetVisibility(v string) {
	o.Visibility = &v
}

// GetResourcePermission returns the ResourcePermission field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetResourcePermission() map[string]interface{} {
	if o == nil || IsNil(o.ResourcePermission) {
		var ret map[string]interface{}
		return ret
	}
	return o.ResourcePermission
}

// GetResourcePermissionOk returns a tuple with the ResourcePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetResourcePermissionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ResourcePermission) {
		return map[string]interface{}{}, false
	}
	return o.ResourcePermission, true
}

// IsSetResourcePermission returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) IsSetResourcePermission() bool {
	if o != nil && !IsNil(o.ResourcePermission) {
		return true
	}

	return false
}

// SetResourcePermission gets a reference to the given map[string]interface{} and assigns it to the ResourcePermission field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) SetResourcePermission(v map[string]interface{}) {
	o.ResourcePermission = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetOwner() map[string]interface{} {
	if o == nil || IsNil(o.Owner) {
		var ret map[string]interface{}
		return ret
	}
	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetOwnerOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Owner) {
		return map[string]interface{}{}, false
	}
	return o.Owner, true
}

// IsSetOwner returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) IsSetOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given map[string]interface{} and assigns it to the Owner field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) SetOwner(v map[string]interface{}) {
	o.Owner = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetTenant() map[string]interface{} {
	if o == nil || IsNil(o.Tenant) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) GetTenantOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tenant) {
		return map[string]interface{}{}, false
	}
	return o.Tenant, true
}

// IsSetTenant returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) IsSetTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given map[string]interface{} and assigns it to the Tenant field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) SetTenant(v map[string]interface{}) {
	o.Tenant = v
}

func (o AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Terraform) {
		toSerialize["terraform"] = o.Terraform
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.ResourcePermission) {
		toSerialize["resourcePermission"] = o.ResourcePermission
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf5) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

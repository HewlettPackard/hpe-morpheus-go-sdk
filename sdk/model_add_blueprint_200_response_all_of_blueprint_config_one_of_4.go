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

// checks if the AddBlueprint200ResponseAllOfBlueprintConfigOneOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddBlueprint200ResponseAllOfBlueprintConfigOneOf4{}

// AddBlueprint200ResponseAllOfBlueprintConfigOneOf4 struct for AddBlueprint200ResponseAllOfBlueprintConfigOneOf4
type AddBlueprint200ResponseAllOfBlueprintConfigOneOf4 struct {
	// A name for the blueprint
	Name *string `json:"name,omitempty"`
	// Blueprint Type
	Type   *string                                                  `json:"type,omitempty"`
	Config *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config `json:"config,omitempty"`
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

type _AddBlueprint200ResponseAllOfBlueprintConfigOneOf4 AddBlueprint200ResponseAllOfBlueprintConfigOneOf4

// NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf4 instantiates a new AddBlueprint200ResponseAllOfBlueprintConfigOneOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf4() *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4 {
	this := AddBlueprint200ResponseAllOfBlueprintConfigOneOf4{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf4WithDefaults instantiates a new AddBlueprint200ResponseAllOfBlueprintConfigOneOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf4WithDefaults() *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4 {
	this := AddBlueprint200ResponseAllOfBlueprintConfigOneOf4{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetType(v string) {
	o.Type = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetConfig() AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config {
	if o == nil || IsNil(o.Config) {
		var ret AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetConfigOk() (*AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config and assigns it to the Config field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetConfig(v AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config) {
	o.Config = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetVisibility(v string) {
	o.Visibility = &v
}

// GetResourcePermission returns the ResourcePermission field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetResourcePermission() map[string]interface{} {
	if o == nil || IsNil(o.ResourcePermission) {
		var ret map[string]interface{}
		return ret
	}
	return o.ResourcePermission
}

// GetResourcePermissionOk returns a tuple with the ResourcePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetResourcePermissionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ResourcePermission) {
		return map[string]interface{}{}, false
	}
	return o.ResourcePermission, true
}

// IsSetResourcePermission returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) IsSetResourcePermission() bool {
	if o != nil && !IsNil(o.ResourcePermission) {
		return true
	}

	return false
}

// SetResourcePermission gets a reference to the given map[string]interface{} and assigns it to the ResourcePermission field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetResourcePermission(v map[string]interface{}) {
	o.ResourcePermission = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetOwner() map[string]interface{} {
	if o == nil || IsNil(o.Owner) {
		var ret map[string]interface{}
		return ret
	}
	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetOwnerOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Owner) {
		return map[string]interface{}{}, false
	}
	return o.Owner, true
}

// IsSetOwner returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) IsSetOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given map[string]interface{} and assigns it to the Owner field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetOwner(v map[string]interface{}) {
	o.Owner = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetTenant() map[string]interface{} {
	if o == nil || IsNil(o.Tenant) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetTenantOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tenant) {
		return map[string]interface{}{}, false
	}
	return o.Tenant, true
}

// IsSetTenant returns a boolean if a field has been set.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) IsSetTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given map[string]interface{} and assigns it to the Tenant field.
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetTenant(v map[string]interface{}) {
	o.Tenant = v
}

func (o AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
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
func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

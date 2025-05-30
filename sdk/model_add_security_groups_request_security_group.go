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
	"bytes"
	"fmt"
)

// checks if the AddSecurityGroupsRequestSecurityGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSecurityGroupsRequestSecurityGroup{}

// AddSecurityGroupsRequestSecurityGroup struct for AddSecurityGroupsRequestSecurityGroup
type AddSecurityGroupsRequestSecurityGroup struct {
	// Name for your security group
	Name string `json:"name"`
	// Optional description field
	Description *string `json:"description,omitempty"`
	// Scoped Cloud ID
	ZoneId int64 `json:"zoneId"`
	// Set to `false` to disable a security group.
	Active *bool `json:"active,omitempty"`
	CustomOptions *AddSecurityGroupsRequestSecurityGroupCustomOptions `json:"customOptions,omitempty"`
	TenantPermissions *AddSecurityGroupsRequestSecurityGroupTenantPermissions `json:"tenantPermissions,omitempty"`
	ResourcePermissions *UpdateCloudDatastoresRequestDatastoreResourcePermissions `json:"resourcePermissions,omitempty"`
}

type _AddSecurityGroupsRequestSecurityGroup AddSecurityGroupsRequestSecurityGroup

// NewAddSecurityGroupsRequestSecurityGroup instantiates a new AddSecurityGroupsRequestSecurityGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSecurityGroupsRequestSecurityGroup(name string, zoneId int64) *AddSecurityGroupsRequestSecurityGroup {
	this := AddSecurityGroupsRequestSecurityGroup{}
	this.Name = name
	this.ZoneId = zoneId
	return &this
}

// NewAddSecurityGroupsRequestSecurityGroupWithDefaults instantiates a new AddSecurityGroupsRequestSecurityGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSecurityGroupsRequestSecurityGroupWithDefaults() *AddSecurityGroupsRequestSecurityGroup {
	this := AddSecurityGroupsRequestSecurityGroup{}
	return &this
}

// GetName returns the Name field value
func (o *AddSecurityGroupsRequestSecurityGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddSecurityGroupsRequestSecurityGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddSecurityGroupsRequestSecurityGroup) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSecurityGroupsRequestSecurityGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSecurityGroupsRequestSecurityGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *AddSecurityGroupsRequestSecurityGroup) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSecurityGroupsRequestSecurityGroup) SetDescription(v string) {
	o.Description = &v
}

// GetZoneId returns the ZoneId field value
func (o *AddSecurityGroupsRequestSecurityGroup) GetZoneId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value
// and a boolean to check if the value has been set.
func (o *AddSecurityGroupsRequestSecurityGroup) GetZoneIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneId, true
}

// SetZoneId sets field value
func (o *AddSecurityGroupsRequestSecurityGroup) SetZoneId(v int64) {
	o.ZoneId = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AddSecurityGroupsRequestSecurityGroup) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSecurityGroupsRequestSecurityGroup) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// IsSetActive returns a boolean if a field has been set.
func (o *AddSecurityGroupsRequestSecurityGroup) IsSetActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AddSecurityGroupsRequestSecurityGroup) SetActive(v bool) {
	o.Active = &v
}

// GetCustomOptions returns the CustomOptions field value if set, zero value otherwise.
func (o *AddSecurityGroupsRequestSecurityGroup) GetCustomOptions() AddSecurityGroupsRequestSecurityGroupCustomOptions {
	if o == nil || IsNil(o.CustomOptions) {
		var ret AddSecurityGroupsRequestSecurityGroupCustomOptions
		return ret
	}
	return *o.CustomOptions
}

// GetCustomOptionsOk returns a tuple with the CustomOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSecurityGroupsRequestSecurityGroup) GetCustomOptionsOk() (*AddSecurityGroupsRequestSecurityGroupCustomOptions, bool) {
	if o == nil || IsNil(o.CustomOptions) {
		return nil, false
	}
	return o.CustomOptions, true
}

// IsSetCustomOptions returns a boolean if a field has been set.
func (o *AddSecurityGroupsRequestSecurityGroup) IsSetCustomOptions() bool {
	if o != nil && !IsNil(o.CustomOptions) {
		return true
	}

	return false
}

// SetCustomOptions gets a reference to the given AddSecurityGroupsRequestSecurityGroupCustomOptions and assigns it to the CustomOptions field.
func (o *AddSecurityGroupsRequestSecurityGroup) SetCustomOptions(v AddSecurityGroupsRequestSecurityGroupCustomOptions) {
	o.CustomOptions = &v
}

// GetTenantPermissions returns the TenantPermissions field value if set, zero value otherwise.
func (o *AddSecurityGroupsRequestSecurityGroup) GetTenantPermissions() AddSecurityGroupsRequestSecurityGroupTenantPermissions {
	if o == nil || IsNil(o.TenantPermissions) {
		var ret AddSecurityGroupsRequestSecurityGroupTenantPermissions
		return ret
	}
	return *o.TenantPermissions
}

// GetTenantPermissionsOk returns a tuple with the TenantPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSecurityGroupsRequestSecurityGroup) GetTenantPermissionsOk() (*AddSecurityGroupsRequestSecurityGroupTenantPermissions, bool) {
	if o == nil || IsNil(o.TenantPermissions) {
		return nil, false
	}
	return o.TenantPermissions, true
}

// IsSetTenantPermissions returns a boolean if a field has been set.
func (o *AddSecurityGroupsRequestSecurityGroup) IsSetTenantPermissions() bool {
	if o != nil && !IsNil(o.TenantPermissions) {
		return true
	}

	return false
}

// SetTenantPermissions gets a reference to the given AddSecurityGroupsRequestSecurityGroupTenantPermissions and assigns it to the TenantPermissions field.
func (o *AddSecurityGroupsRequestSecurityGroup) SetTenantPermissions(v AddSecurityGroupsRequestSecurityGroupTenantPermissions) {
	o.TenantPermissions = &v
}

// GetResourcePermissions returns the ResourcePermissions field value if set, zero value otherwise.
func (o *AddSecurityGroupsRequestSecurityGroup) GetResourcePermissions() UpdateCloudDatastoresRequestDatastoreResourcePermissions {
	if o == nil || IsNil(o.ResourcePermissions) {
		var ret UpdateCloudDatastoresRequestDatastoreResourcePermissions
		return ret
	}
	return *o.ResourcePermissions
}

// GetResourcePermissionsOk returns a tuple with the ResourcePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSecurityGroupsRequestSecurityGroup) GetResourcePermissionsOk() (*UpdateCloudDatastoresRequestDatastoreResourcePermissions, bool) {
	if o == nil || IsNil(o.ResourcePermissions) {
		return nil, false
	}
	return o.ResourcePermissions, true
}

// IsSetResourcePermissions returns a boolean if a field has been set.
func (o *AddSecurityGroupsRequestSecurityGroup) IsSetResourcePermissions() bool {
	if o != nil && !IsNil(o.ResourcePermissions) {
		return true
	}

	return false
}

// SetResourcePermissions gets a reference to the given UpdateCloudDatastoresRequestDatastoreResourcePermissions and assigns it to the ResourcePermissions field.
func (o *AddSecurityGroupsRequestSecurityGroup) SetResourcePermissions(v UpdateCloudDatastoresRequestDatastoreResourcePermissions) {
	o.ResourcePermissions = &v
}

func (o AddSecurityGroupsRequestSecurityGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSecurityGroupsRequestSecurityGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["zoneId"] = o.ZoneId
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.CustomOptions) {
		toSerialize["customOptions"] = o.CustomOptions
	}
	if !IsNil(o.TenantPermissions) {
		toSerialize["tenantPermissions"] = o.TenantPermissions
	}
	if !IsNil(o.ResourcePermissions) {
		toSerialize["resourcePermissions"] = o.ResourcePermissions
	}
	return toSerialize, nil
}

func (o *AddSecurityGroupsRequestSecurityGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"zoneId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddSecurityGroupsRequestSecurityGroup := _AddSecurityGroupsRequestSecurityGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddSecurityGroupsRequestSecurityGroup)

	if err != nil {
		return err
	}

	*o = AddSecurityGroupsRequestSecurityGroup(varAddSecurityGroupsRequestSecurityGroup)

	return err
}

type NullableAddSecurityGroupsRequestSecurityGroup struct {
	value *AddSecurityGroupsRequestSecurityGroup
	isSet bool
}

func (v NullableAddSecurityGroupsRequestSecurityGroup) Get() *AddSecurityGroupsRequestSecurityGroup {
	return v.value
}

func (v *NullableAddSecurityGroupsRequestSecurityGroup) Set(val *AddSecurityGroupsRequestSecurityGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSecurityGroupsRequestSecurityGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSecurityGroupsRequestSecurityGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSecurityGroupsRequestSecurityGroup(val *AddSecurityGroupsRequestSecurityGroup) *NullableAddSecurityGroupsRequestSecurityGroup {
	return &NullableAddSecurityGroupsRequestSecurityGroup{value: val, isSet: true}
}

func (v NullableAddSecurityGroupsRequestSecurityGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSecurityGroupsRequestSecurityGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1{}

// AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1 struct for AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1
type AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1 struct {
	// Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected.
	NoAgent *bool `json:"noAgent,omitempty"`
	// id of the resource group to be used, can be prefixed with `pool-`. A resource pool group can be specified instead by prefixing its ID with `poolGroup-`.
	ResourcePoolId *string `json:"resourcePoolId,omitempty"`
	// Specific host to deploy to if so desired.
	HostId *string `json:"hostId,omitempty"`
	// Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used.
	SmbiosAssetTag *string `json:"smbiosAssetTag,omitempty"`
	// Enable Nested Virtualization
	NestedVirtualization *string `json:"nestedVirtualization,omitempty"`
	// VMWare Folder External ID (as a String) or ID (as an Integer or String)
	VmwareFolderId       *string                `json:"vmwareFolderId,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1 AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1

// NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1 instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1 {
	this := AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1{}
	var noAgent bool = false
	this.NoAgent = &noAgent
	var nestedVirtualization string = "off"
	this.NestedVirtualization = &nestedVirtualization
	return &this
}

// NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1WithDefaults instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1WithDefaults() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1 {
	this := AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1{}
	var noAgent bool = false
	this.NoAgent = &noAgent
	var nestedVirtualization string = "off"
	this.NestedVirtualization = &nestedVirtualization
	return &this
}

// GetNoAgent returns the NoAgent field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetNoAgent() bool {
	if o == nil || IsNil(o.NoAgent) {
		var ret bool
		return ret
	}
	return *o.NoAgent
}

// GetNoAgentOk returns a tuple with the NoAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetNoAgentOk() (*bool, bool) {
	if o == nil || IsNil(o.NoAgent) {
		return nil, false
	}
	return o.NoAgent, true
}

// IsSetNoAgent returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) IsSetNoAgent() bool {
	if o != nil && !IsNil(o.NoAgent) {
		return true
	}

	return false
}

// SetNoAgent gets a reference to the given bool and assigns it to the NoAgent field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) SetNoAgent(v bool) {
	o.NoAgent = &v
}

// GetResourcePoolId returns the ResourcePoolId field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetResourcePoolId() string {
	if o == nil || IsNil(o.ResourcePoolId) {
		var ret string
		return ret
	}
	return *o.ResourcePoolId
}

// GetResourcePoolIdOk returns a tuple with the ResourcePoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetResourcePoolIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourcePoolId) {
		return nil, false
	}
	return o.ResourcePoolId, true
}

// IsSetResourcePoolId returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) IsSetResourcePoolId() bool {
	if o != nil && !IsNil(o.ResourcePoolId) {
		return true
	}

	return false
}

// SetResourcePoolId gets a reference to the given string and assigns it to the ResourcePoolId field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) SetResourcePoolId(v string) {
	o.ResourcePoolId = &v
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetHostId() string {
	if o == nil || IsNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetHostIdOk() (*string, bool) {
	if o == nil || IsNil(o.HostId) {
		return nil, false
	}
	return o.HostId, true
}

// IsSetHostId returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) IsSetHostId() bool {
	if o != nil && !IsNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) SetHostId(v string) {
	o.HostId = &v
}

// GetSmbiosAssetTag returns the SmbiosAssetTag field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetSmbiosAssetTag() string {
	if o == nil || IsNil(o.SmbiosAssetTag) {
		var ret string
		return ret
	}
	return *o.SmbiosAssetTag
}

// GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetSmbiosAssetTagOk() (*string, bool) {
	if o == nil || IsNil(o.SmbiosAssetTag) {
		return nil, false
	}
	return o.SmbiosAssetTag, true
}

// IsSetSmbiosAssetTag returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) IsSetSmbiosAssetTag() bool {
	if o != nil && !IsNil(o.SmbiosAssetTag) {
		return true
	}

	return false
}

// SetSmbiosAssetTag gets a reference to the given string and assigns it to the SmbiosAssetTag field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) SetSmbiosAssetTag(v string) {
	o.SmbiosAssetTag = &v
}

// GetNestedVirtualization returns the NestedVirtualization field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetNestedVirtualization() string {
	if o == nil || IsNil(o.NestedVirtualization) {
		var ret string
		return ret
	}
	return *o.NestedVirtualization
}

// GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetNestedVirtualizationOk() (*string, bool) {
	if o == nil || IsNil(o.NestedVirtualization) {
		return nil, false
	}
	return o.NestedVirtualization, true
}

// IsSetNestedVirtualization returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) IsSetNestedVirtualization() bool {
	if o != nil && !IsNil(o.NestedVirtualization) {
		return true
	}

	return false
}

// SetNestedVirtualization gets a reference to the given string and assigns it to the NestedVirtualization field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) SetNestedVirtualization(v string) {
	o.NestedVirtualization = &v
}

// GetVmwareFolderId returns the VmwareFolderId field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetVmwareFolderId() string {
	if o == nil || IsNil(o.VmwareFolderId) {
		var ret string
		return ret
	}
	return *o.VmwareFolderId
}

// GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetVmwareFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.VmwareFolderId) {
		return nil, false
	}
	return o.VmwareFolderId, true
}

// IsSetVmwareFolderId returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) IsSetVmwareFolderId() bool {
	if o != nil && !IsNil(o.VmwareFolderId) {
		return true
	}

	return false
}

// SetVmwareFolderId gets a reference to the given string and assigns it to the VmwareFolderId field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) SetVmwareFolderId(v string) {
	o.VmwareFolderId = &v
}

func (o AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NoAgent) {
		toSerialize["noAgent"] = o.NoAgent
	}
	if !IsNil(o.ResourcePoolId) {
		toSerialize["resourcePoolId"] = o.ResourcePoolId
	}
	if !IsNil(o.HostId) {
		toSerialize["hostId"] = o.HostId
	}
	if !IsNil(o.SmbiosAssetTag) {
		toSerialize["smbiosAssetTag"] = o.SmbiosAssetTag
	}
	if !IsNil(o.NestedVirtualization) {
		toSerialize["nestedVirtualization"] = o.NestedVirtualization
	}
	if !IsNil(o.VmwareFolderId) {
		toSerialize["vmwareFolderId"] = o.VmwareFolderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

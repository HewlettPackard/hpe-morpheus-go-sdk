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

// checks if the InstancesCloneImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstancesCloneImage{}

// InstancesCloneImage struct for InstancesCloneImage
type InstancesCloneImage struct {
	// Image Template Name
	TemplateName *string `json:"templateName,omitempty"`
	// Zone Folder externalId. This is required for VMware
	ZoneFolder           *string                `json:"zoneFolder,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _InstancesCloneImage InstancesCloneImage

// NewInstancesCloneImage instantiates a new InstancesCloneImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstancesCloneImage() *InstancesCloneImage {
	this := InstancesCloneImage{}
	var templateName string = "{server.name}-{timestamp}"
	this.TemplateName = &templateName
	return &this
}

// NewInstancesCloneImageWithDefaults instantiates a new InstancesCloneImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstancesCloneImageWithDefaults() *InstancesCloneImage {
	this := InstancesCloneImage{}
	var templateName string = "{server.name}-{timestamp}"
	this.TemplateName = &templateName
	return &this
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *InstancesCloneImage) GetTemplateName() string {
	if o == nil || IsNil(o.TemplateName) {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesCloneImage) GetTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateName) {
		return nil, false
	}
	return o.TemplateName, true
}

// IsSetTemplateName returns a boolean if a field has been set.
func (o *InstancesCloneImage) IsSetTemplateName() bool {
	if o != nil && !IsNil(o.TemplateName) {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *InstancesCloneImage) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetZoneFolder returns the ZoneFolder field value if set, zero value otherwise.
func (o *InstancesCloneImage) GetZoneFolder() string {
	if o == nil || IsNil(o.ZoneFolder) {
		var ret string
		return ret
	}
	return *o.ZoneFolder
}

// GetZoneFolderOk returns a tuple with the ZoneFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesCloneImage) GetZoneFolderOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneFolder) {
		return nil, false
	}
	return o.ZoneFolder, true
}

// IsSetZoneFolder returns a boolean if a field has been set.
func (o *InstancesCloneImage) IsSetZoneFolder() bool {
	if o != nil && !IsNil(o.ZoneFolder) {
		return true
	}

	return false
}

// SetZoneFolder gets a reference to the given string and assigns it to the ZoneFolder field.
func (o *InstancesCloneImage) SetZoneFolder(v string) {
	o.ZoneFolder = &v
}

func (o InstancesCloneImage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstancesCloneImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplateName) {
		toSerialize["templateName"] = o.TemplateName
	}
	if !IsNil(o.ZoneFolder) {
		toSerialize["zoneFolder"] = o.ZoneFolder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *InstancesCloneImage) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

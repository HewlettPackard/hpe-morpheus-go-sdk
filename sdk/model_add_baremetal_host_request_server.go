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

// checks if the AddBaremetalHostRequestServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddBaremetalHostRequestServer{}

// AddBaremetalHostRequestServer struct for AddBaremetalHostRequestServer
type AddBaremetalHostRequestServer struct {
	Cloud                *AddBaremetalHostRequestServerCloud             `json:"cloud,omitempty"`
	ComputeServerType    *AddBaremetalHostRequestServerComputeServerType `json:"computeServerType,omitempty"`
	Description          *string                                         `json:"description,omitempty"`
	Group                *AddBaremetalHostRequestServerGroup             `json:"group,omitempty"`
	Labels               []string                                        `json:"labels,omitempty"`
	Name                 *string                                         `json:"name,omitempty"`
	Visibility           *string                                         `json:"visibility,omitempty"`
	Config               *AddBaremetalHostRequestServerConfig            `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}                          `json:",remain"`
}

type _AddBaremetalHostRequestServer AddBaremetalHostRequestServer

// NewAddBaremetalHostRequestServer instantiates a new AddBaremetalHostRequestServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddBaremetalHostRequestServer() *AddBaremetalHostRequestServer {
	this := AddBaremetalHostRequestServer{}
	return &this
}

// NewAddBaremetalHostRequestServerWithDefaults instantiates a new AddBaremetalHostRequestServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddBaremetalHostRequestServerWithDefaults() *AddBaremetalHostRequestServer {
	this := AddBaremetalHostRequestServer{}
	return &this
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *AddBaremetalHostRequestServer) GetCloud() AddBaremetalHostRequestServerCloud {
	if o == nil || IsNil(o.Cloud) {
		var ret AddBaremetalHostRequestServerCloud
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBaremetalHostRequestServer) GetCloudOk() (*AddBaremetalHostRequestServerCloud, bool) {
	if o == nil || IsNil(o.Cloud) {
		return nil, false
	}
	return o.Cloud, true
}

// IsSetCloud returns a boolean if a field has been set.
func (o *AddBaremetalHostRequestServer) IsSetCloud() bool {
	if o != nil && !IsNil(o.Cloud) {
		return true
	}

	return false
}

// SetCloud gets a reference to the given AddBaremetalHostRequestServerCloud and assigns it to the Cloud field.
func (o *AddBaremetalHostRequestServer) SetCloud(v AddBaremetalHostRequestServerCloud) {
	o.Cloud = &v
}

// GetComputeServerType returns the ComputeServerType field value if set, zero value otherwise.
func (o *AddBaremetalHostRequestServer) GetComputeServerType() AddBaremetalHostRequestServerComputeServerType {
	if o == nil || IsNil(o.ComputeServerType) {
		var ret AddBaremetalHostRequestServerComputeServerType
		return ret
	}
	return *o.ComputeServerType
}

// GetComputeServerTypeOk returns a tuple with the ComputeServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBaremetalHostRequestServer) GetComputeServerTypeOk() (*AddBaremetalHostRequestServerComputeServerType, bool) {
	if o == nil || IsNil(o.ComputeServerType) {
		return nil, false
	}
	return o.ComputeServerType, true
}

// IsSetComputeServerType returns a boolean if a field has been set.
func (o *AddBaremetalHostRequestServer) IsSetComputeServerType() bool {
	if o != nil && !IsNil(o.ComputeServerType) {
		return true
	}

	return false
}

// SetComputeServerType gets a reference to the given AddBaremetalHostRequestServerComputeServerType and assigns it to the ComputeServerType field.
func (o *AddBaremetalHostRequestServer) SetComputeServerType(v AddBaremetalHostRequestServerComputeServerType) {
	o.ComputeServerType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddBaremetalHostRequestServer) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBaremetalHostRequestServer) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *AddBaremetalHostRequestServer) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddBaremetalHostRequestServer) SetDescription(v string) {
	o.Description = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *AddBaremetalHostRequestServer) GetGroup() AddBaremetalHostRequestServerGroup {
	if o == nil || IsNil(o.Group) {
		var ret AddBaremetalHostRequestServerGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBaremetalHostRequestServer) GetGroupOk() (*AddBaremetalHostRequestServerGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// IsSetGroup returns a boolean if a field has been set.
func (o *AddBaremetalHostRequestServer) IsSetGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given AddBaremetalHostRequestServerGroup and assigns it to the Group field.
func (o *AddBaremetalHostRequestServer) SetGroup(v AddBaremetalHostRequestServerGroup) {
	o.Group = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AddBaremetalHostRequestServer) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBaremetalHostRequestServer) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *AddBaremetalHostRequestServer) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *AddBaremetalHostRequestServer) SetLabels(v []string) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddBaremetalHostRequestServer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBaremetalHostRequestServer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *AddBaremetalHostRequestServer) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddBaremetalHostRequestServer) SetName(v string) {
	o.Name = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *AddBaremetalHostRequestServer) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBaremetalHostRequestServer) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *AddBaremetalHostRequestServer) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *AddBaremetalHostRequestServer) SetVisibility(v string) {
	o.Visibility = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AddBaremetalHostRequestServer) GetConfig() AddBaremetalHostRequestServerConfig {
	if o == nil || IsNil(o.Config) {
		var ret AddBaremetalHostRequestServerConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBaremetalHostRequestServer) GetConfigOk() (*AddBaremetalHostRequestServerConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *AddBaremetalHostRequestServer) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given AddBaremetalHostRequestServerConfig and assigns it to the Config field.
func (o *AddBaremetalHostRequestServer) SetConfig(v AddBaremetalHostRequestServerConfig) {
	o.Config = &v
}

func (o AddBaremetalHostRequestServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddBaremetalHostRequestServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cloud) {
		toSerialize["cloud"] = o.Cloud
	}
	if !IsNil(o.ComputeServerType) {
		toSerialize["computeServerType"] = o.ComputeServerType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddBaremetalHostRequestServer) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

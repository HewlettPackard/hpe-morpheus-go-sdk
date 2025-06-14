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

// checks if the ListSecurityPackages200ResponseAllOfSecurityPackagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSecurityPackages200ResponseAllOfSecurityPackagesInner{}

// ListSecurityPackages200ResponseAllOfSecurityPackagesInner struct for ListSecurityPackages200ResponseAllOfSecurityPackagesInner
type ListSecurityPackages200ResponseAllOfSecurityPackagesInner struct {
	Id                   *int64                                                      `json:"id,omitempty"`
	Name                 *string                                                     `json:"name,omitempty"`
	Labels               []string                                                    `json:"labels,omitempty"`
	Description          *string                                                     `json:"description,omitempty"`
	Type                 *ListBackupSettings200ResponseBackupSettingsDefaultSchedule `json:"type,omitempty"`
	Enabled              *bool                                                       `json:"enabled,omitempty"`
	Url                  *string                                                     `json:"url,omitempty"`
	Uuid                 *string                                                     `json:"uuid,omitempty"`
	DateCreated          *time.Time                                                  `json:"dateCreated,omitempty"`
	LastUpdated          *time.Time                                                  `json:"lastUpdated,omitempty"`
	Config               map[string]interface{}                                      `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}                                      `json:",remain"`
}

type _ListSecurityPackages200ResponseAllOfSecurityPackagesInner ListSecurityPackages200ResponseAllOfSecurityPackagesInner

// NewListSecurityPackages200ResponseAllOfSecurityPackagesInner instantiates a new ListSecurityPackages200ResponseAllOfSecurityPackagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSecurityPackages200ResponseAllOfSecurityPackagesInner() *ListSecurityPackages200ResponseAllOfSecurityPackagesInner {
	this := ListSecurityPackages200ResponseAllOfSecurityPackagesInner{}
	return &this
}

// NewListSecurityPackages200ResponseAllOfSecurityPackagesInnerWithDefaults instantiates a new ListSecurityPackages200ResponseAllOfSecurityPackagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSecurityPackages200ResponseAllOfSecurityPackagesInnerWithDefaults() *ListSecurityPackages200ResponseAllOfSecurityPackagesInner {
	this := ListSecurityPackages200ResponseAllOfSecurityPackagesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetLabels(v []string) {
	o.Labels = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule {
	if o == nil || IsNil(o.Type) {
		var ret ListBackupSettings200ResponseBackupSettingsDefaultSchedule
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ListBackupSettings200ResponseBackupSettingsDefaultSchedule and assigns it to the Type field.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// IsSetUrl returns a boolean if a field has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) IsSetUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetUrl(v string) {
	o.Url = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// IsSetUuid returns a boolean if a field has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) IsSetUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetUuid(v string) {
	o.Uuid = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) SetConfig(v map[string]interface{}) {
	o.Config = v
}

func (o ListSecurityPackages200ResponseAllOfSecurityPackagesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSecurityPackages200ResponseAllOfSecurityPackagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListSecurityPackages200ResponseAllOfSecurityPackagesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

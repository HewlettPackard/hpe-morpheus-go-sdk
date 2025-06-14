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

// checks if the UpdateInstanceRequestInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInstanceRequestInstance{}

// UpdateInstanceRequestInstance struct for UpdateInstanceRequestInstance
type UpdateInstanceRequestInstance struct {
	// Unique name scoped to your account for the instance.
	Name *string `json:"name,omitempty"`
	// Optional description field.
	Description *string `json:"description,omitempty"`
	// Environment
	InstanceContext *string `json:"instanceContext,omitempty"`
	// Array of strings (keywords).
	Labels []string `json:"labels,omitempty"`
	// Metadata tags, Array of objects having a name and value.
	Tags []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner `json:"tags,omitempty"`
	// Add or update value of Metadata tags, Array of objects having a name and value.
	AddTags []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner `json:"addTags,omitempty"`
	// Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed.
	RemoveTags []ListInstances200ResponseAllOfInstancesInnerTagsInner `json:"removeTags,omitempty"`
	// Power schedule ID.
	PowerScheduleType *int64                             `json:"powerScheduleType,omitempty"`
	Site              *UpdateInstanceRequestInstanceSite `json:"site,omitempty"`
	// User ID, can be used to change instance owner.
	OwnerId *int64 `json:"ownerId,omitempty"`
	// Name used in the UI for display
	DisplayName          *string                `json:"displayName,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _UpdateInstanceRequestInstance UpdateInstanceRequestInstance

// NewUpdateInstanceRequestInstance instantiates a new UpdateInstanceRequestInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInstanceRequestInstance() *UpdateInstanceRequestInstance {
	this := UpdateInstanceRequestInstance{}
	return &this
}

// NewUpdateInstanceRequestInstanceWithDefaults instantiates a new UpdateInstanceRequestInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInstanceRequestInstanceWithDefaults() *UpdateInstanceRequestInstance {
	this := UpdateInstanceRequestInstance{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateInstanceRequestInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceRequestInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *UpdateInstanceRequestInstance) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateInstanceRequestInstance) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateInstanceRequestInstance) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceRequestInstance) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *UpdateInstanceRequestInstance) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateInstanceRequestInstance) SetDescription(v string) {
	o.Description = &v
}

// GetInstanceContext returns the InstanceContext field value if set, zero value otherwise.
func (o *UpdateInstanceRequestInstance) GetInstanceContext() string {
	if o == nil || IsNil(o.InstanceContext) {
		var ret string
		return ret
	}
	return *o.InstanceContext
}

// GetInstanceContextOk returns a tuple with the InstanceContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceRequestInstance) GetInstanceContextOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceContext) {
		return nil, false
	}
	return o.InstanceContext, true
}

// IsSetInstanceContext returns a boolean if a field has been set.
func (o *UpdateInstanceRequestInstance) IsSetInstanceContext() bool {
	if o != nil && !IsNil(o.InstanceContext) {
		return true
	}

	return false
}

// SetInstanceContext gets a reference to the given string and assigns it to the InstanceContext field.
func (o *UpdateInstanceRequestInstance) SetInstanceContext(v string) {
	o.InstanceContext = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateInstanceRequestInstance) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceRequestInstance) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *UpdateInstanceRequestInstance) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *UpdateInstanceRequestInstance) SetLabels(v []string) {
	o.Labels = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateInstanceRequestInstance) GetTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner {
	if o == nil || IsNil(o.Tags) {
		var ret []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceRequestInstance) GetTagsOk() ([]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// IsSetTags returns a boolean if a field has been set.
func (o *UpdateInstanceRequestInstance) IsSetTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner and assigns it to the Tags field.
func (o *UpdateInstanceRequestInstance) SetTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner) {
	o.Tags = v
}

// GetAddTags returns the AddTags field value if set, zero value otherwise.
func (o *UpdateInstanceRequestInstance) GetAddTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner {
	if o == nil || IsNil(o.AddTags) {
		var ret []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner
		return ret
	}
	return o.AddTags
}

// GetAddTagsOk returns a tuple with the AddTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceRequestInstance) GetAddTagsOk() ([]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool) {
	if o == nil || IsNil(o.AddTags) {
		return nil, false
	}
	return o.AddTags, true
}

// IsSetAddTags returns a boolean if a field has been set.
func (o *UpdateInstanceRequestInstance) IsSetAddTags() bool {
	if o != nil && !IsNil(o.AddTags) {
		return true
	}

	return false
}

// SetAddTags gets a reference to the given []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner and assigns it to the AddTags field.
func (o *UpdateInstanceRequestInstance) SetAddTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner) {
	o.AddTags = v
}

// GetRemoveTags returns the RemoveTags field value if set, zero value otherwise.
func (o *UpdateInstanceRequestInstance) GetRemoveTags() []ListInstances200ResponseAllOfInstancesInnerTagsInner {
	if o == nil || IsNil(o.RemoveTags) {
		var ret []ListInstances200ResponseAllOfInstancesInnerTagsInner
		return ret
	}
	return o.RemoveTags
}

// GetRemoveTagsOk returns a tuple with the RemoveTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceRequestInstance) GetRemoveTagsOk() ([]ListInstances200ResponseAllOfInstancesInnerTagsInner, bool) {
	if o == nil || IsNil(o.RemoveTags) {
		return nil, false
	}
	return o.RemoveTags, true
}

// IsSetRemoveTags returns a boolean if a field has been set.
func (o *UpdateInstanceRequestInstance) IsSetRemoveTags() bool {
	if o != nil && !IsNil(o.RemoveTags) {
		return true
	}

	return false
}

// SetRemoveTags gets a reference to the given []ListInstances200ResponseAllOfInstancesInnerTagsInner and assigns it to the RemoveTags field.
func (o *UpdateInstanceRequestInstance) SetRemoveTags(v []ListInstances200ResponseAllOfInstancesInnerTagsInner) {
	o.RemoveTags = v
}

// GetPowerScheduleType returns the PowerScheduleType field value if set, zero value otherwise.
func (o *UpdateInstanceRequestInstance) GetPowerScheduleType() int64 {
	if o == nil || IsNil(o.PowerScheduleType) {
		var ret int64
		return ret
	}
	return *o.PowerScheduleType
}

// GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceRequestInstance) GetPowerScheduleTypeOk() (*int64, bool) {
	if o == nil || IsNil(o.PowerScheduleType) {
		return nil, false
	}
	return o.PowerScheduleType, true
}

// IsSetPowerScheduleType returns a boolean if a field has been set.
func (o *UpdateInstanceRequestInstance) IsSetPowerScheduleType() bool {
	if o != nil && !IsNil(o.PowerScheduleType) {
		return true
	}

	return false
}

// SetPowerScheduleType gets a reference to the given int64 and assigns it to the PowerScheduleType field.
func (o *UpdateInstanceRequestInstance) SetPowerScheduleType(v int64) {
	o.PowerScheduleType = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *UpdateInstanceRequestInstance) GetSite() UpdateInstanceRequestInstanceSite {
	if o == nil || IsNil(o.Site) {
		var ret UpdateInstanceRequestInstanceSite
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceRequestInstance) GetSiteOk() (*UpdateInstanceRequestInstanceSite, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// IsSetSite returns a boolean if a field has been set.
func (o *UpdateInstanceRequestInstance) IsSetSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given UpdateInstanceRequestInstanceSite and assigns it to the Site field.
func (o *UpdateInstanceRequestInstance) SetSite(v UpdateInstanceRequestInstanceSite) {
	o.Site = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *UpdateInstanceRequestInstance) GetOwnerId() int64 {
	if o == nil || IsNil(o.OwnerId) {
		var ret int64
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceRequestInstance) GetOwnerIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// IsSetOwnerId returns a boolean if a field has been set.
func (o *UpdateInstanceRequestInstance) IsSetOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given int64 and assigns it to the OwnerId field.
func (o *UpdateInstanceRequestInstance) SetOwnerId(v int64) {
	o.OwnerId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UpdateInstanceRequestInstance) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceRequestInstance) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// IsSetDisplayName returns a boolean if a field has been set.
func (o *UpdateInstanceRequestInstance) IsSetDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UpdateInstanceRequestInstance) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o UpdateInstanceRequestInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInstanceRequestInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InstanceContext) {
		toSerialize["instanceContext"] = o.InstanceContext
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.AddTags) {
		toSerialize["addTags"] = o.AddTags
	}
	if !IsNil(o.RemoveTags) {
		toSerialize["removeTags"] = o.RemoveTags
	}
	if !IsNil(o.PowerScheduleType) {
		toSerialize["powerScheduleType"] = o.PowerScheduleType
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateInstanceRequestInstance) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

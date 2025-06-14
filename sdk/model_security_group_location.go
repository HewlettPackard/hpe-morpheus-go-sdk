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

// checks if the SecurityGroupLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupLocation{}

// SecurityGroupLocation struct for SecurityGroupLocation
type SecurityGroupLocation struct {
	Id                   *int64                                                                  `json:"id,omitempty"`
	Name                 *string                                                                 `json:"name,omitempty"`
	Description          *string                                                                 `json:"description,omitempty"`
	ExternalId           *string                                                                 `json:"externalId,omitempty"`
	IacId                *string                                                                 `json:"iacId,omitempty"`
	Zone                 *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"zone,omitempty"`
	ZonePool             *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"zonePool,omitempty"`
	Status               *string                                                                 `json:"status,omitempty"`
	Priority             *string                                                                 `json:"priority,omitempty"`
	GroupLayer           *string                                                                 `json:"groupLayer,omitempty"`
	AdditionalProperties map[string]interface{}                                                  `json:",remain"`
}

type _SecurityGroupLocation SecurityGroupLocation

// NewSecurityGroupLocation instantiates a new SecurityGroupLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupLocation() *SecurityGroupLocation {
	this := SecurityGroupLocation{}
	return &this
}

// NewSecurityGroupLocationWithDefaults instantiates a new SecurityGroupLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupLocationWithDefaults() *SecurityGroupLocation {
	this := SecurityGroupLocation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityGroupLocation) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLocation) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *SecurityGroupLocation) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SecurityGroupLocation) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityGroupLocation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLocation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *SecurityGroupLocation) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityGroupLocation) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityGroupLocation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLocation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *SecurityGroupLocation) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityGroupLocation) SetDescription(v string) {
	o.Description = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *SecurityGroupLocation) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLocation) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// IsSetExternalId returns a boolean if a field has been set.
func (o *SecurityGroupLocation) IsSetExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *SecurityGroupLocation) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetIacId returns the IacId field value if set, zero value otherwise.
func (o *SecurityGroupLocation) GetIacId() string {
	if o == nil || IsNil(o.IacId) {
		var ret string
		return ret
	}
	return *o.IacId
}

// GetIacIdOk returns a tuple with the IacId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLocation) GetIacIdOk() (*string, bool) {
	if o == nil || IsNil(o.IacId) {
		return nil, false
	}
	return o.IacId, true
}

// IsSetIacId returns a boolean if a field has been set.
func (o *SecurityGroupLocation) IsSetIacId() bool {
	if o != nil && !IsNil(o.IacId) {
		return true
	}

	return false
}

// SetIacId gets a reference to the given string and assigns it to the IacId field.
func (o *SecurityGroupLocation) SetIacId(v string) {
	o.IacId = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *SecurityGroupLocation) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Zone) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLocation) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// IsSetZone returns a boolean if a field has been set.
func (o *SecurityGroupLocation) IsSetZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Zone field.
func (o *SecurityGroupLocation) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Zone = &v
}

// GetZonePool returns the ZonePool field value if set, zero value otherwise.
func (o *SecurityGroupLocation) GetZonePool() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.ZonePool) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.ZonePool
}

// GetZonePoolOk returns a tuple with the ZonePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLocation) GetZonePoolOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.ZonePool) {
		return nil, false
	}
	return o.ZonePool, true
}

// IsSetZonePool returns a boolean if a field has been set.
func (o *SecurityGroupLocation) IsSetZonePool() bool {
	if o != nil && !IsNil(o.ZonePool) {
		return true
	}

	return false
}

// SetZonePool gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the ZonePool field.
func (o *SecurityGroupLocation) SetZonePool(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.ZonePool = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecurityGroupLocation) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLocation) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// IsSetStatus returns a boolean if a field has been set.
func (o *SecurityGroupLocation) IsSetStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SecurityGroupLocation) SetStatus(v string) {
	o.Status = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SecurityGroupLocation) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLocation) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// IsSetPriority returns a boolean if a field has been set.
func (o *SecurityGroupLocation) IsSetPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *SecurityGroupLocation) SetPriority(v string) {
	o.Priority = &v
}

// GetGroupLayer returns the GroupLayer field value if set, zero value otherwise.
func (o *SecurityGroupLocation) GetGroupLayer() string {
	if o == nil || IsNil(o.GroupLayer) {
		var ret string
		return ret
	}
	return *o.GroupLayer
}

// GetGroupLayerOk returns a tuple with the GroupLayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupLocation) GetGroupLayerOk() (*string, bool) {
	if o == nil || IsNil(o.GroupLayer) {
		return nil, false
	}
	return o.GroupLayer, true
}

// IsSetGroupLayer returns a boolean if a field has been set.
func (o *SecurityGroupLocation) IsSetGroupLayer() bool {
	if o != nil && !IsNil(o.GroupLayer) {
		return true
	}

	return false
}

// SetGroupLayer gets a reference to the given string and assigns it to the GroupLayer field.
func (o *SecurityGroupLocation) SetGroupLayer(v string) {
	o.GroupLayer = &v
}

func (o SecurityGroupLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityGroupLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.IacId) {
		toSerialize["iacId"] = o.IacId
	}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}
	if !IsNil(o.ZonePool) {
		toSerialize["zonePool"] = o.ZonePool
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.GroupLayer) {
		toSerialize["groupLayer"] = o.GroupLayer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *SecurityGroupLocation) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

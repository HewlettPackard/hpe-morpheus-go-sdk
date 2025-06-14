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

// checks if the ListCloudDatastores200ResponseAllOfDatastoresInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCloudDatastores200ResponseAllOfDatastoresInner{}

// ListCloudDatastores200ResponseAllOfDatastoresInner struct for ListCloudDatastores200ResponseAllOfDatastoresInner
type ListCloudDatastores200ResponseAllOfDatastoresInner struct {
	Id                   *int64                                                                  `json:"id,omitempty"`
	Name                 *string                                                                 `json:"name,omitempty"`
	Zone                 *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"zone,omitempty"`
	Type                 *string                                                                 `json:"type,omitempty"`
	FreeSpace            *int64                                                                  `json:"freeSpace,omitempty"`
	Online               *bool                                                                   `json:"online,omitempty"`
	Active               *bool                                                                   `json:"active,omitempty"`
	Visibility           *string                                                                 `json:"visibility,omitempty"`
	Tenants              []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner        `json:"tenants,omitempty"`
	ResourcePermission   *ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission   `json:"resourcePermission,omitempty"`
	AdditionalProperties map[string]interface{}                                                  `json:",remain"`
}

type _ListCloudDatastores200ResponseAllOfDatastoresInner ListCloudDatastores200ResponseAllOfDatastoresInner

// NewListCloudDatastores200ResponseAllOfDatastoresInner instantiates a new ListCloudDatastores200ResponseAllOfDatastoresInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCloudDatastores200ResponseAllOfDatastoresInner() *ListCloudDatastores200ResponseAllOfDatastoresInner {
	this := ListCloudDatastores200ResponseAllOfDatastoresInner{}
	return &this
}

// NewListCloudDatastores200ResponseAllOfDatastoresInnerWithDefaults instantiates a new ListCloudDatastores200ResponseAllOfDatastoresInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCloudDatastores200ResponseAllOfDatastoresInnerWithDefaults() *ListCloudDatastores200ResponseAllOfDatastoresInner {
	this := ListCloudDatastores200ResponseAllOfDatastoresInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) SetName(v string) {
	o.Name = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Zone) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// IsSetZone returns a boolean if a field has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) IsSetZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Zone field.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Zone = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) SetType(v string) {
	o.Type = &v
}

// GetFreeSpace returns the FreeSpace field value if set, zero value otherwise.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetFreeSpace() int64 {
	if o == nil || IsNil(o.FreeSpace) {
		var ret int64
		return ret
	}
	return *o.FreeSpace
}

// GetFreeSpaceOk returns a tuple with the FreeSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetFreeSpaceOk() (*int64, bool) {
	if o == nil || IsNil(o.FreeSpace) {
		return nil, false
	}
	return o.FreeSpace, true
}

// IsSetFreeSpace returns a boolean if a field has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) IsSetFreeSpace() bool {
	if o != nil && !IsNil(o.FreeSpace) {
		return true
	}

	return false
}

// SetFreeSpace gets a reference to the given int64 and assigns it to the FreeSpace field.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) SetFreeSpace(v int64) {
	o.FreeSpace = &v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetOnline() bool {
	if o == nil || IsNil(o.Online) {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetOnlineOk() (*bool, bool) {
	if o == nil || IsNil(o.Online) {
		return nil, false
	}
	return o.Online, true
}

// IsSetOnline returns a boolean if a field has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) IsSetOnline() bool {
	if o != nil && !IsNil(o.Online) {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) SetOnline(v bool) {
	o.Online = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// IsSetActive returns a boolean if a field has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) IsSetActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) SetActive(v bool) {
	o.Active = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) SetVisibility(v string) {
	o.Visibility = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetTenants() []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner {
	if o == nil || IsNil(o.Tenants) {
		var ret []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetTenantsOk() ([]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// IsSetTenants returns a boolean if a field has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) IsSetTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner and assigns it to the Tenants field.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) SetTenants(v []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner) {
	o.Tenants = v
}

// GetResourcePermission returns the ResourcePermission field value if set, zero value otherwise.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetResourcePermission() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission {
	if o == nil || IsNil(o.ResourcePermission) {
		var ret ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission
		return ret
	}
	return *o.ResourcePermission
}

// GetResourcePermissionOk returns a tuple with the ResourcePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) GetResourcePermissionOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool) {
	if o == nil || IsNil(o.ResourcePermission) {
		return nil, false
	}
	return o.ResourcePermission, true
}

// IsSetResourcePermission returns a boolean if a field has been set.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) IsSetResourcePermission() bool {
	if o != nil && !IsNil(o.ResourcePermission) {
		return true
	}

	return false
}

// SetResourcePermission gets a reference to the given ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission and assigns it to the ResourcePermission field.
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) SetResourcePermission(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission) {
	o.ResourcePermission = &v
}

func (o ListCloudDatastores200ResponseAllOfDatastoresInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCloudDatastores200ResponseAllOfDatastoresInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.FreeSpace) {
		toSerialize["freeSpace"] = o.FreeSpace
	}
	if !IsNil(o.Online) {
		toSerialize["online"] = o.Online
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}
	if !IsNil(o.ResourcePermission) {
		toSerialize["resourcePermission"] = o.ResourcePermission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListCloudDatastores200ResponseAllOfDatastoresInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the GetClusterMasters200ResponseMastersInnerVolumesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetClusterMasters200ResponseMastersInnerVolumesInner{}

// GetClusterMasters200ResponseMastersInnerVolumesInner struct for GetClusterMasters200ResponseMastersInnerVolumesInner
type GetClusterMasters200ResponseMastersInnerVolumesInner struct {
	Id                   *int64                 `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	ControllerId         *string                `json:"controllerId,omitempty"`
	ControllerMountPoint *string                `json:"controllerMountPoint,omitempty"`
	Resizeable           *bool                  `json:"resizeable,omitempty"`
	PlanResizable        *bool                  `json:"planResizable,omitempty"`
	RootVolume           *bool                  `json:"rootVolume,omitempty"`
	UnitNumber           *string                `json:"unitNumber,omitempty"`
	TypeId               *int64                 `json:"typeId,omitempty"`
	ConfigurableIOPS     *bool                  `json:"configurableIOPS,omitempty"`
	DatastoreId          *string                `json:"datastoreId,omitempty"`
	MaxStorage           *int64                 `json:"maxStorage,omitempty"`
	DisplayOrder         *int64                 `json:"displayOrder,omitempty"`
	MaxIOPS              *string                `json:"maxIOPS,omitempty"`
	Uuid                 *string                `json:"uuid,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _GetClusterMasters200ResponseMastersInnerVolumesInner GetClusterMasters200ResponseMastersInnerVolumesInner

// NewGetClusterMasters200ResponseMastersInnerVolumesInner instantiates a new GetClusterMasters200ResponseMastersInnerVolumesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetClusterMasters200ResponseMastersInnerVolumesInner() *GetClusterMasters200ResponseMastersInnerVolumesInner {
	this := GetClusterMasters200ResponseMastersInnerVolumesInner{}
	return &this
}

// NewGetClusterMasters200ResponseMastersInnerVolumesInnerWithDefaults instantiates a new GetClusterMasters200ResponseMastersInnerVolumesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetClusterMasters200ResponseMastersInnerVolumesInnerWithDefaults() *GetClusterMasters200ResponseMastersInnerVolumesInner {
	this := GetClusterMasters200ResponseMastersInnerVolumesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetName(v string) {
	o.Name = &v
}

// GetControllerId returns the ControllerId field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetControllerId() string {
	if o == nil || IsNil(o.ControllerId) {
		var ret string
		return ret
	}
	return *o.ControllerId
}

// GetControllerIdOk returns a tuple with the ControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetControllerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerId) {
		return nil, false
	}
	return o.ControllerId, true
}

// IsSetControllerId returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetControllerId() bool {
	if o != nil && !IsNil(o.ControllerId) {
		return true
	}

	return false
}

// SetControllerId gets a reference to the given string and assigns it to the ControllerId field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetControllerId(v string) {
	o.ControllerId = &v
}

// GetControllerMountPoint returns the ControllerMountPoint field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetControllerMountPoint() string {
	if o == nil || IsNil(o.ControllerMountPoint) {
		var ret string
		return ret
	}
	return *o.ControllerMountPoint
}

// GetControllerMountPointOk returns a tuple with the ControllerMountPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetControllerMountPointOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerMountPoint) {
		return nil, false
	}
	return o.ControllerMountPoint, true
}

// IsSetControllerMountPoint returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetControllerMountPoint() bool {
	if o != nil && !IsNil(o.ControllerMountPoint) {
		return true
	}

	return false
}

// SetControllerMountPoint gets a reference to the given string and assigns it to the ControllerMountPoint field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetControllerMountPoint(v string) {
	o.ControllerMountPoint = &v
}

// GetResizeable returns the Resizeable field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetResizeable() bool {
	if o == nil || IsNil(o.Resizeable) {
		var ret bool
		return ret
	}
	return *o.Resizeable
}

// GetResizeableOk returns a tuple with the Resizeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetResizeableOk() (*bool, bool) {
	if o == nil || IsNil(o.Resizeable) {
		return nil, false
	}
	return o.Resizeable, true
}

// IsSetResizeable returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetResizeable() bool {
	if o != nil && !IsNil(o.Resizeable) {
		return true
	}

	return false
}

// SetResizeable gets a reference to the given bool and assigns it to the Resizeable field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetResizeable(v bool) {
	o.Resizeable = &v
}

// GetPlanResizable returns the PlanResizable field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetPlanResizable() bool {
	if o == nil || IsNil(o.PlanResizable) {
		var ret bool
		return ret
	}
	return *o.PlanResizable
}

// GetPlanResizableOk returns a tuple with the PlanResizable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetPlanResizableOk() (*bool, bool) {
	if o == nil || IsNil(o.PlanResizable) {
		return nil, false
	}
	return o.PlanResizable, true
}

// IsSetPlanResizable returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetPlanResizable() bool {
	if o != nil && !IsNil(o.PlanResizable) {
		return true
	}

	return false
}

// SetPlanResizable gets a reference to the given bool and assigns it to the PlanResizable field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetPlanResizable(v bool) {
	o.PlanResizable = &v
}

// GetRootVolume returns the RootVolume field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetRootVolume() bool {
	if o == nil || IsNil(o.RootVolume) {
		var ret bool
		return ret
	}
	return *o.RootVolume
}

// GetRootVolumeOk returns a tuple with the RootVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetRootVolumeOk() (*bool, bool) {
	if o == nil || IsNil(o.RootVolume) {
		return nil, false
	}
	return o.RootVolume, true
}

// IsSetRootVolume returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetRootVolume() bool {
	if o != nil && !IsNil(o.RootVolume) {
		return true
	}

	return false
}

// SetRootVolume gets a reference to the given bool and assigns it to the RootVolume field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetRootVolume(v bool) {
	o.RootVolume = &v
}

// GetUnitNumber returns the UnitNumber field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetUnitNumber() string {
	if o == nil || IsNil(o.UnitNumber) {
		var ret string
		return ret
	}
	return *o.UnitNumber
}

// GetUnitNumberOk returns a tuple with the UnitNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetUnitNumberOk() (*string, bool) {
	if o == nil || IsNil(o.UnitNumber) {
		return nil, false
	}
	return o.UnitNumber, true
}

// IsSetUnitNumber returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetUnitNumber() bool {
	if o != nil && !IsNil(o.UnitNumber) {
		return true
	}

	return false
}

// SetUnitNumber gets a reference to the given string and assigns it to the UnitNumber field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetUnitNumber(v string) {
	o.UnitNumber = &v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetTypeId() int64 {
	if o == nil || IsNil(o.TypeId) {
		var ret int64
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetTypeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TypeId) {
		return nil, false
	}
	return o.TypeId, true
}

// IsSetTypeId returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetTypeId() bool {
	if o != nil && !IsNil(o.TypeId) {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given int64 and assigns it to the TypeId field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetTypeId(v int64) {
	o.TypeId = &v
}

// GetConfigurableIOPS returns the ConfigurableIOPS field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetConfigurableIOPS() bool {
	if o == nil || IsNil(o.ConfigurableIOPS) {
		var ret bool
		return ret
	}
	return *o.ConfigurableIOPS
}

// GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetConfigurableIOPSOk() (*bool, bool) {
	if o == nil || IsNil(o.ConfigurableIOPS) {
		return nil, false
	}
	return o.ConfigurableIOPS, true
}

// IsSetConfigurableIOPS returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetConfigurableIOPS() bool {
	if o != nil && !IsNil(o.ConfigurableIOPS) {
		return true
	}

	return false
}

// SetConfigurableIOPS gets a reference to the given bool and assigns it to the ConfigurableIOPS field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetConfigurableIOPS(v bool) {
	o.ConfigurableIOPS = &v
}

// GetDatastoreId returns the DatastoreId field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetDatastoreId() string {
	if o == nil || IsNil(o.DatastoreId) {
		var ret string
		return ret
	}
	return *o.DatastoreId
}

// GetDatastoreIdOk returns a tuple with the DatastoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetDatastoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatastoreId) {
		return nil, false
	}
	return o.DatastoreId, true
}

// IsSetDatastoreId returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetDatastoreId() bool {
	if o != nil && !IsNil(o.DatastoreId) {
		return true
	}

	return false
}

// SetDatastoreId gets a reference to the given string and assigns it to the DatastoreId field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetDatastoreId(v string) {
	o.DatastoreId = &v
}

// GetMaxStorage returns the MaxStorage field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetMaxStorage() int64 {
	if o == nil || IsNil(o.MaxStorage) {
		var ret int64
		return ret
	}
	return *o.MaxStorage
}

// GetMaxStorageOk returns a tuple with the MaxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetMaxStorageOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxStorage) {
		return nil, false
	}
	return o.MaxStorage, true
}

// IsSetMaxStorage returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetMaxStorage() bool {
	if o != nil && !IsNil(o.MaxStorage) {
		return true
	}

	return false
}

// SetMaxStorage gets a reference to the given int64 and assigns it to the MaxStorage field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetMaxStorage(v int64) {
	o.MaxStorage = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetDisplayOrder() int64 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret int64
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetDisplayOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// IsSetDisplayOrder returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int64 and assigns it to the DisplayOrder field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetDisplayOrder(v int64) {
	o.DisplayOrder = &v
}

// GetMaxIOPS returns the MaxIOPS field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetMaxIOPS() string {
	if o == nil || IsNil(o.MaxIOPS) {
		var ret string
		return ret
	}
	return *o.MaxIOPS
}

// GetMaxIOPSOk returns a tuple with the MaxIOPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetMaxIOPSOk() (*string, bool) {
	if o == nil || IsNil(o.MaxIOPS) {
		return nil, false
	}
	return o.MaxIOPS, true
}

// IsSetMaxIOPS returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetMaxIOPS() bool {
	if o != nil && !IsNil(o.MaxIOPS) {
		return true
	}

	return false
}

// SetMaxIOPS gets a reference to the given string and assigns it to the MaxIOPS field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetMaxIOPS(v string) {
	o.MaxIOPS = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// IsSetUuid returns a boolean if a field has been set.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) IsSetUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) SetUuid(v string) {
	o.Uuid = &v
}

func (o GetClusterMasters200ResponseMastersInnerVolumesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetClusterMasters200ResponseMastersInnerVolumesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ControllerId) {
		toSerialize["controllerId"] = o.ControllerId
	}
	if !IsNil(o.ControllerMountPoint) {
		toSerialize["controllerMountPoint"] = o.ControllerMountPoint
	}
	if !IsNil(o.Resizeable) {
		toSerialize["resizeable"] = o.Resizeable
	}
	if !IsNil(o.PlanResizable) {
		toSerialize["planResizable"] = o.PlanResizable
	}
	if !IsNil(o.RootVolume) {
		toSerialize["rootVolume"] = o.RootVolume
	}
	if !IsNil(o.UnitNumber) {
		toSerialize["unitNumber"] = o.UnitNumber
	}
	if !IsNil(o.TypeId) {
		toSerialize["typeId"] = o.TypeId
	}
	if !IsNil(o.ConfigurableIOPS) {
		toSerialize["configurableIOPS"] = o.ConfigurableIOPS
	}
	if !IsNil(o.DatastoreId) {
		toSerialize["datastoreId"] = o.DatastoreId
	}
	if !IsNil(o.MaxStorage) {
		toSerialize["maxStorage"] = o.MaxStorage
	}
	if !IsNil(o.DisplayOrder) {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if !IsNil(o.MaxIOPS) {
		toSerialize["maxIOPS"] = o.MaxIOPS
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetClusterMasters200ResponseMastersInnerVolumesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner{}

// AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner struct for AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner
type AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner struct {
	Id                   *int64                 `json:"id,omitempty"`
	Size                 *int64                 `json:"size,omitempty"`
	MaxIOPS              *string                `json:"maxIOPS,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	RootVolume           *bool                  `json:"rootVolume,omitempty"`
	StorageType          *int64                 `json:"storageType,omitempty"`
	DatastoreId          *string                `json:"datastoreId,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner

// NewAddImageBuild200ResponseAllOfImageBuildConfigVolumesInner instantiates a new AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddImageBuild200ResponseAllOfImageBuildConfigVolumesInner() *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner {
	this := AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner{}
	return &this
}

// NewAddImageBuild200ResponseAllOfImageBuildConfigVolumesInnerWithDefaults instantiates a new AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddImageBuild200ResponseAllOfImageBuildConfigVolumesInnerWithDefaults() *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner {
	this := AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) SetId(v int64) {
	o.Id = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// IsSetSize returns a boolean if a field has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) IsSetSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) SetSize(v int64) {
	o.Size = &v
}

// GetMaxIOPS returns the MaxIOPS field value if set, zero value otherwise.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetMaxIOPS() string {
	if o == nil || IsNil(o.MaxIOPS) {
		var ret string
		return ret
	}
	return *o.MaxIOPS
}

// GetMaxIOPSOk returns a tuple with the MaxIOPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetMaxIOPSOk() (*string, bool) {
	if o == nil || IsNil(o.MaxIOPS) {
		return nil, false
	}
	return o.MaxIOPS, true
}

// IsSetMaxIOPS returns a boolean if a field has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) IsSetMaxIOPS() bool {
	if o != nil && !IsNil(o.MaxIOPS) {
		return true
	}

	return false
}

// SetMaxIOPS gets a reference to the given string and assigns it to the MaxIOPS field.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) SetMaxIOPS(v string) {
	o.MaxIOPS = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) SetName(v string) {
	o.Name = &v
}

// GetRootVolume returns the RootVolume field value if set, zero value otherwise.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetRootVolume() bool {
	if o == nil || IsNil(o.RootVolume) {
		var ret bool
		return ret
	}
	return *o.RootVolume
}

// GetRootVolumeOk returns a tuple with the RootVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetRootVolumeOk() (*bool, bool) {
	if o == nil || IsNil(o.RootVolume) {
		return nil, false
	}
	return o.RootVolume, true
}

// IsSetRootVolume returns a boolean if a field has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) IsSetRootVolume() bool {
	if o != nil && !IsNil(o.RootVolume) {
		return true
	}

	return false
}

// SetRootVolume gets a reference to the given bool and assigns it to the RootVolume field.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) SetRootVolume(v bool) {
	o.RootVolume = &v
}

// GetStorageType returns the StorageType field value if set, zero value otherwise.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetStorageType() int64 {
	if o == nil || IsNil(o.StorageType) {
		var ret int64
		return ret
	}
	return *o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetStorageTypeOk() (*int64, bool) {
	if o == nil || IsNil(o.StorageType) {
		return nil, false
	}
	return o.StorageType, true
}

// IsSetStorageType returns a boolean if a field has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) IsSetStorageType() bool {
	if o != nil && !IsNil(o.StorageType) {
		return true
	}

	return false
}

// SetStorageType gets a reference to the given int64 and assigns it to the StorageType field.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) SetStorageType(v int64) {
	o.StorageType = &v
}

// GetDatastoreId returns the DatastoreId field value if set, zero value otherwise.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetDatastoreId() string {
	if o == nil || IsNil(o.DatastoreId) {
		var ret string
		return ret
	}
	return *o.DatastoreId
}

// GetDatastoreIdOk returns a tuple with the DatastoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) GetDatastoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatastoreId) {
		return nil, false
	}
	return o.DatastoreId, true
}

// IsSetDatastoreId returns a boolean if a field has been set.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) IsSetDatastoreId() bool {
	if o != nil && !IsNil(o.DatastoreId) {
		return true
	}

	return false
}

// SetDatastoreId gets a reference to the given string and assigns it to the DatastoreId field.
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) SetDatastoreId(v string) {
	o.DatastoreId = &v
}

func (o AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.MaxIOPS) {
		toSerialize["maxIOPS"] = o.MaxIOPS
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RootVolume) {
		toSerialize["rootVolume"] = o.RootVolume
	}
	if !IsNil(o.StorageType) {
		toSerialize["storageType"] = o.StorageType
	}
	if !IsNil(o.DatastoreId) {
		toSerialize["datastoreId"] = o.DatastoreId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddImageBuild200ResponseAllOfImageBuildConfigVolumesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

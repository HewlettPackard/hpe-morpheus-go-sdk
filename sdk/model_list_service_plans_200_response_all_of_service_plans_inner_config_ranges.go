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

// checks if the ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges{}

// ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges struct for ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges
type ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges struct {
	MinStorage           *string                `json:"minStorage,omitempty"`
	MaxStorage           *string                `json:"maxStorage,omitempty"`
	MinPerDiskSize       *string                `json:"minPerDiskSize,omitempty"`
	MaxPerDiskSize       *string                `json:"maxPerDiskSize,omitempty"`
	MinMemory            *float32               `json:"minMemory,omitempty"`
	MaxMemory            *float32               `json:"maxMemory,omitempty"`
	MinCores             *string                `json:"minCores,omitempty"`
	MaxCores             *string                `json:"maxCores,omitempty"`
	MinSockets           *string                `json:"minSockets,omitempty"`
	MaxSockets           *string                `json:"maxSockets,omitempty"`
	MinCoresPerSocket    *string                `json:"minCoresPerSocket,omitempty"`
	MaxCoresPerSocket    *string                `json:"maxCoresPerSocket,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges

// NewListServicePlans200ResponseAllOfServicePlansInnerConfigRanges instantiates a new ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServicePlans200ResponseAllOfServicePlansInnerConfigRanges() *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges {
	this := ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges{}
	return &this
}

// NewListServicePlans200ResponseAllOfServicePlansInnerConfigRangesWithDefaults instantiates a new ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServicePlans200ResponseAllOfServicePlansInnerConfigRangesWithDefaults() *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges {
	this := ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges{}
	return &this
}

// GetMinStorage returns the MinStorage field value if set, zero value otherwise.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMinStorage() string {
	if o == nil || IsNil(o.MinStorage) {
		var ret string
		return ret
	}
	return *o.MinStorage
}

// GetMinStorageOk returns a tuple with the MinStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMinStorageOk() (*string, bool) {
	if o == nil || IsNil(o.MinStorage) {
		return nil, false
	}
	return o.MinStorage, true
}

// IsSetMinStorage returns a boolean if a field has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) IsSetMinStorage() bool {
	if o != nil && !IsNil(o.MinStorage) {
		return true
	}

	return false
}

// SetMinStorage gets a reference to the given string and assigns it to the MinStorage field.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) SetMinStorage(v string) {
	o.MinStorage = &v
}

// GetMaxStorage returns the MaxStorage field value if set, zero value otherwise.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMaxStorage() string {
	if o == nil || IsNil(o.MaxStorage) {
		var ret string
		return ret
	}
	return *o.MaxStorage
}

// GetMaxStorageOk returns a tuple with the MaxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMaxStorageOk() (*string, bool) {
	if o == nil || IsNil(o.MaxStorage) {
		return nil, false
	}
	return o.MaxStorage, true
}

// IsSetMaxStorage returns a boolean if a field has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) IsSetMaxStorage() bool {
	if o != nil && !IsNil(o.MaxStorage) {
		return true
	}

	return false
}

// SetMaxStorage gets a reference to the given string and assigns it to the MaxStorage field.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) SetMaxStorage(v string) {
	o.MaxStorage = &v
}

// GetMinPerDiskSize returns the MinPerDiskSize field value if set, zero value otherwise.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMinPerDiskSize() string {
	if o == nil || IsNil(o.MinPerDiskSize) {
		var ret string
		return ret
	}
	return *o.MinPerDiskSize
}

// GetMinPerDiskSizeOk returns a tuple with the MinPerDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMinPerDiskSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MinPerDiskSize) {
		return nil, false
	}
	return o.MinPerDiskSize, true
}

// IsSetMinPerDiskSize returns a boolean if a field has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) IsSetMinPerDiskSize() bool {
	if o != nil && !IsNil(o.MinPerDiskSize) {
		return true
	}

	return false
}

// SetMinPerDiskSize gets a reference to the given string and assigns it to the MinPerDiskSize field.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) SetMinPerDiskSize(v string) {
	o.MinPerDiskSize = &v
}

// GetMaxPerDiskSize returns the MaxPerDiskSize field value if set, zero value otherwise.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMaxPerDiskSize() string {
	if o == nil || IsNil(o.MaxPerDiskSize) {
		var ret string
		return ret
	}
	return *o.MaxPerDiskSize
}

// GetMaxPerDiskSizeOk returns a tuple with the MaxPerDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMaxPerDiskSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxPerDiskSize) {
		return nil, false
	}
	return o.MaxPerDiskSize, true
}

// IsSetMaxPerDiskSize returns a boolean if a field has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) IsSetMaxPerDiskSize() bool {
	if o != nil && !IsNil(o.MaxPerDiskSize) {
		return true
	}

	return false
}

// SetMaxPerDiskSize gets a reference to the given string and assigns it to the MaxPerDiskSize field.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) SetMaxPerDiskSize(v string) {
	o.MaxPerDiskSize = &v
}

// GetMinMemory returns the MinMemory field value if set, zero value otherwise.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMinMemory() float32 {
	if o == nil || IsNil(o.MinMemory) {
		var ret float32
		return ret
	}
	return *o.MinMemory
}

// GetMinMemoryOk returns a tuple with the MinMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMinMemoryOk() (*float32, bool) {
	if o == nil || IsNil(o.MinMemory) {
		return nil, false
	}
	return o.MinMemory, true
}

// IsSetMinMemory returns a boolean if a field has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) IsSetMinMemory() bool {
	if o != nil && !IsNil(o.MinMemory) {
		return true
	}

	return false
}

// SetMinMemory gets a reference to the given float32 and assigns it to the MinMemory field.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) SetMinMemory(v float32) {
	o.MinMemory = &v
}

// GetMaxMemory returns the MaxMemory field value if set, zero value otherwise.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMaxMemory() float32 {
	if o == nil || IsNil(o.MaxMemory) {
		var ret float32
		return ret
	}
	return *o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMaxMemoryOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxMemory) {
		return nil, false
	}
	return o.MaxMemory, true
}

// IsSetMaxMemory returns a boolean if a field has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) IsSetMaxMemory() bool {
	if o != nil && !IsNil(o.MaxMemory) {
		return true
	}

	return false
}

// SetMaxMemory gets a reference to the given float32 and assigns it to the MaxMemory field.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) SetMaxMemory(v float32) {
	o.MaxMemory = &v
}

// GetMinCores returns the MinCores field value if set, zero value otherwise.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMinCores() string {
	if o == nil || IsNil(o.MinCores) {
		var ret string
		return ret
	}
	return *o.MinCores
}

// GetMinCoresOk returns a tuple with the MinCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMinCoresOk() (*string, bool) {
	if o == nil || IsNil(o.MinCores) {
		return nil, false
	}
	return o.MinCores, true
}

// IsSetMinCores returns a boolean if a field has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) IsSetMinCores() bool {
	if o != nil && !IsNil(o.MinCores) {
		return true
	}

	return false
}

// SetMinCores gets a reference to the given string and assigns it to the MinCores field.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) SetMinCores(v string) {
	o.MinCores = &v
}

// GetMaxCores returns the MaxCores field value if set, zero value otherwise.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMaxCores() string {
	if o == nil || IsNil(o.MaxCores) {
		var ret string
		return ret
	}
	return *o.MaxCores
}

// GetMaxCoresOk returns a tuple with the MaxCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMaxCoresOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCores) {
		return nil, false
	}
	return o.MaxCores, true
}

// IsSetMaxCores returns a boolean if a field has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) IsSetMaxCores() bool {
	if o != nil && !IsNil(o.MaxCores) {
		return true
	}

	return false
}

// SetMaxCores gets a reference to the given string and assigns it to the MaxCores field.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) SetMaxCores(v string) {
	o.MaxCores = &v
}

// GetMinSockets returns the MinSockets field value if set, zero value otherwise.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMinSockets() string {
	if o == nil || IsNil(o.MinSockets) {
		var ret string
		return ret
	}
	return *o.MinSockets
}

// GetMinSocketsOk returns a tuple with the MinSockets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMinSocketsOk() (*string, bool) {
	if o == nil || IsNil(o.MinSockets) {
		return nil, false
	}
	return o.MinSockets, true
}

// IsSetMinSockets returns a boolean if a field has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) IsSetMinSockets() bool {
	if o != nil && !IsNil(o.MinSockets) {
		return true
	}

	return false
}

// SetMinSockets gets a reference to the given string and assigns it to the MinSockets field.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) SetMinSockets(v string) {
	o.MinSockets = &v
}

// GetMaxSockets returns the MaxSockets field value if set, zero value otherwise.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMaxSockets() string {
	if o == nil || IsNil(o.MaxSockets) {
		var ret string
		return ret
	}
	return *o.MaxSockets
}

// GetMaxSocketsOk returns a tuple with the MaxSockets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMaxSocketsOk() (*string, bool) {
	if o == nil || IsNil(o.MaxSockets) {
		return nil, false
	}
	return o.MaxSockets, true
}

// IsSetMaxSockets returns a boolean if a field has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) IsSetMaxSockets() bool {
	if o != nil && !IsNil(o.MaxSockets) {
		return true
	}

	return false
}

// SetMaxSockets gets a reference to the given string and assigns it to the MaxSockets field.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) SetMaxSockets(v string) {
	o.MaxSockets = &v
}

// GetMinCoresPerSocket returns the MinCoresPerSocket field value if set, zero value otherwise.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMinCoresPerSocket() string {
	if o == nil || IsNil(o.MinCoresPerSocket) {
		var ret string
		return ret
	}
	return *o.MinCoresPerSocket
}

// GetMinCoresPerSocketOk returns a tuple with the MinCoresPerSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMinCoresPerSocketOk() (*string, bool) {
	if o == nil || IsNil(o.MinCoresPerSocket) {
		return nil, false
	}
	return o.MinCoresPerSocket, true
}

// IsSetMinCoresPerSocket returns a boolean if a field has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) IsSetMinCoresPerSocket() bool {
	if o != nil && !IsNil(o.MinCoresPerSocket) {
		return true
	}

	return false
}

// SetMinCoresPerSocket gets a reference to the given string and assigns it to the MinCoresPerSocket field.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) SetMinCoresPerSocket(v string) {
	o.MinCoresPerSocket = &v
}

// GetMaxCoresPerSocket returns the MaxCoresPerSocket field value if set, zero value otherwise.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMaxCoresPerSocket() string {
	if o == nil || IsNil(o.MaxCoresPerSocket) {
		var ret string
		return ret
	}
	return *o.MaxCoresPerSocket
}

// GetMaxCoresPerSocketOk returns a tuple with the MaxCoresPerSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) GetMaxCoresPerSocketOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCoresPerSocket) {
		return nil, false
	}
	return o.MaxCoresPerSocket, true
}

// IsSetMaxCoresPerSocket returns a boolean if a field has been set.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) IsSetMaxCoresPerSocket() bool {
	if o != nil && !IsNil(o.MaxCoresPerSocket) {
		return true
	}

	return false
}

// SetMaxCoresPerSocket gets a reference to the given string and assigns it to the MaxCoresPerSocket field.
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) SetMaxCoresPerSocket(v string) {
	o.MaxCoresPerSocket = &v
}

func (o ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinStorage) {
		toSerialize["minStorage"] = o.MinStorage
	}
	if !IsNil(o.MaxStorage) {
		toSerialize["maxStorage"] = o.MaxStorage
	}
	if !IsNil(o.MinPerDiskSize) {
		toSerialize["minPerDiskSize"] = o.MinPerDiskSize
	}
	if !IsNil(o.MaxPerDiskSize) {
		toSerialize["maxPerDiskSize"] = o.MaxPerDiskSize
	}
	if !IsNil(o.MinMemory) {
		toSerialize["minMemory"] = o.MinMemory
	}
	if !IsNil(o.MaxMemory) {
		toSerialize["maxMemory"] = o.MaxMemory
	}
	if !IsNil(o.MinCores) {
		toSerialize["minCores"] = o.MinCores
	}
	if !IsNil(o.MaxCores) {
		toSerialize["maxCores"] = o.MaxCores
	}
	if !IsNil(o.MinSockets) {
		toSerialize["minSockets"] = o.MinSockets
	}
	if !IsNil(o.MaxSockets) {
		toSerialize["maxSockets"] = o.MaxSockets
	}
	if !IsNil(o.MinCoresPerSocket) {
		toSerialize["minCoresPerSocket"] = o.MinCoresPerSocket
	}
	if !IsNil(o.MaxCoresPerSocket) {
		toSerialize["maxCoresPerSocket"] = o.MaxCoresPerSocket
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListServicePlans200ResponseAllOfServicePlansInnerConfigRanges) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

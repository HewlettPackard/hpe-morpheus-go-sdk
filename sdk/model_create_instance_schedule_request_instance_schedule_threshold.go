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

// checks if the CreateInstanceScheduleRequestInstanceScheduleThreshold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInstanceScheduleRequestInstanceScheduleThreshold{}

// CreateInstanceScheduleRequestInstanceScheduleThreshold struct for CreateInstanceScheduleRequestInstanceScheduleThreshold
type CreateInstanceScheduleRequestInstanceScheduleThreshold struct {
	// Source Scale Threshold to apply as a template. All threshold settings with be copied from this threshold, and can be overridden by also passing each setting explicitly.
	SourceThresholdId *int64 `json:"sourceThresholdId,omitempty"`
	// Auto Upscale
	AutoUp *bool `json:"autoUp,omitempty"`
	// Auto Downscale
	AutoDown *bool `json:"autoDown,omitempty"`
	// The minimum number of nodes to scale down to
	MinCount *int32 `json:"minCount,omitempty"`
	// The maximum number of nodes to scale up to
	MaxCount *int32 `json:"maxCount,omitempty"`
	// Enable CPU Threshold
	CpuEnabled *bool `json:"cpuEnabled,omitempty"`
	// Min CPU (%)
	MinCpu *float64 `json:"minCpu,omitempty"`
	// Max CPU (%)
	MaxCpu *float64 `json:"maxCpu,omitempty"`
	// Enable Memory Threshold
	MemoryEnabled *bool `json:"memoryEnabled,omitempty"`
	// Min Memory (%)
	MinMemory *float64 `json:"minMemory,omitempty"`
	// Max Memory (%)
	MaxMemory *float64 `json:"maxMemory,omitempty"`
	// Enable Disk Threshold
	DiskEnabled *bool `json:"diskEnabled,omitempty"`
	// Min Disk (%)
	MinDisk *float64 `json:"minDisk,omitempty"`
	// Max Disk (%)
	MaxDisk              *float64               `json:"maxDisk,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _CreateInstanceScheduleRequestInstanceScheduleThreshold CreateInstanceScheduleRequestInstanceScheduleThreshold

// NewCreateInstanceScheduleRequestInstanceScheduleThreshold instantiates a new CreateInstanceScheduleRequestInstanceScheduleThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInstanceScheduleRequestInstanceScheduleThreshold() *CreateInstanceScheduleRequestInstanceScheduleThreshold {
	this := CreateInstanceScheduleRequestInstanceScheduleThreshold{}
	var autoUp bool = false
	this.AutoUp = &autoUp
	var autoDown bool = false
	this.AutoDown = &autoDown
	var cpuEnabled bool = false
	this.CpuEnabled = &cpuEnabled
	var minCpu float64 = 0
	this.MinCpu = &minCpu
	var maxCpu float64 = 0
	this.MaxCpu = &maxCpu
	var memoryEnabled bool = false
	this.MemoryEnabled = &memoryEnabled
	var minMemory float64 = 0
	this.MinMemory = &minMemory
	var maxMemory float64 = 0
	this.MaxMemory = &maxMemory
	var diskEnabled bool = false
	this.DiskEnabled = &diskEnabled
	var minDisk float64 = 0
	this.MinDisk = &minDisk
	var maxDisk float64 = 0
	this.MaxDisk = &maxDisk
	return &this
}

// NewCreateInstanceScheduleRequestInstanceScheduleThresholdWithDefaults instantiates a new CreateInstanceScheduleRequestInstanceScheduleThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInstanceScheduleRequestInstanceScheduleThresholdWithDefaults() *CreateInstanceScheduleRequestInstanceScheduleThreshold {
	this := CreateInstanceScheduleRequestInstanceScheduleThreshold{}
	var autoUp bool = false
	this.AutoUp = &autoUp
	var autoDown bool = false
	this.AutoDown = &autoDown
	var cpuEnabled bool = false
	this.CpuEnabled = &cpuEnabled
	var minCpu float64 = 0
	this.MinCpu = &minCpu
	var maxCpu float64 = 0
	this.MaxCpu = &maxCpu
	var memoryEnabled bool = false
	this.MemoryEnabled = &memoryEnabled
	var minMemory float64 = 0
	this.MinMemory = &minMemory
	var maxMemory float64 = 0
	this.MaxMemory = &maxMemory
	var diskEnabled bool = false
	this.DiskEnabled = &diskEnabled
	var minDisk float64 = 0
	this.MinDisk = &minDisk
	var maxDisk float64 = 0
	this.MaxDisk = &maxDisk
	return &this
}

// GetSourceThresholdId returns the SourceThresholdId field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetSourceThresholdId() int64 {
	if o == nil || IsNil(o.SourceThresholdId) {
		var ret int64
		return ret
	}
	return *o.SourceThresholdId
}

// GetSourceThresholdIdOk returns a tuple with the SourceThresholdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetSourceThresholdIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SourceThresholdId) {
		return nil, false
	}
	return o.SourceThresholdId, true
}

// IsSetSourceThresholdId returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetSourceThresholdId() bool {
	if o != nil && !IsNil(o.SourceThresholdId) {
		return true
	}

	return false
}

// SetSourceThresholdId gets a reference to the given int64 and assigns it to the SourceThresholdId field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetSourceThresholdId(v int64) {
	o.SourceThresholdId = &v
}

// GetAutoUp returns the AutoUp field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetAutoUp() bool {
	if o == nil || IsNil(o.AutoUp) {
		var ret bool
		return ret
	}
	return *o.AutoUp
}

// GetAutoUpOk returns a tuple with the AutoUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetAutoUpOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoUp) {
		return nil, false
	}
	return o.AutoUp, true
}

// IsSetAutoUp returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetAutoUp() bool {
	if o != nil && !IsNil(o.AutoUp) {
		return true
	}

	return false
}

// SetAutoUp gets a reference to the given bool and assigns it to the AutoUp field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetAutoUp(v bool) {
	o.AutoUp = &v
}

// GetAutoDown returns the AutoDown field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetAutoDown() bool {
	if o == nil || IsNil(o.AutoDown) {
		var ret bool
		return ret
	}
	return *o.AutoDown
}

// GetAutoDownOk returns a tuple with the AutoDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetAutoDownOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoDown) {
		return nil, false
	}
	return o.AutoDown, true
}

// IsSetAutoDown returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetAutoDown() bool {
	if o != nil && !IsNil(o.AutoDown) {
		return true
	}

	return false
}

// SetAutoDown gets a reference to the given bool and assigns it to the AutoDown field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetAutoDown(v bool) {
	o.AutoDown = &v
}

// GetMinCount returns the MinCount field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMinCount() int32 {
	if o == nil || IsNil(o.MinCount) {
		var ret int32
		return ret
	}
	return *o.MinCount
}

// GetMinCountOk returns a tuple with the MinCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMinCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MinCount) {
		return nil, false
	}
	return o.MinCount, true
}

// IsSetMinCount returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetMinCount() bool {
	if o != nil && !IsNil(o.MinCount) {
		return true
	}

	return false
}

// SetMinCount gets a reference to the given int32 and assigns it to the MinCount field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetMinCount(v int32) {
	o.MinCount = &v
}

// GetMaxCount returns the MaxCount field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMaxCount() int32 {
	if o == nil || IsNil(o.MaxCount) {
		var ret int32
		return ret
	}
	return *o.MaxCount
}

// GetMaxCountOk returns a tuple with the MaxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMaxCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxCount) {
		return nil, false
	}
	return o.MaxCount, true
}

// IsSetMaxCount returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetMaxCount() bool {
	if o != nil && !IsNil(o.MaxCount) {
		return true
	}

	return false
}

// SetMaxCount gets a reference to the given int32 and assigns it to the MaxCount field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetMaxCount(v int32) {
	o.MaxCount = &v
}

// GetCpuEnabled returns the CpuEnabled field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetCpuEnabled() bool {
	if o == nil || IsNil(o.CpuEnabled) {
		var ret bool
		return ret
	}
	return *o.CpuEnabled
}

// GetCpuEnabledOk returns a tuple with the CpuEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetCpuEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CpuEnabled) {
		return nil, false
	}
	return o.CpuEnabled, true
}

// IsSetCpuEnabled returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetCpuEnabled() bool {
	if o != nil && !IsNil(o.CpuEnabled) {
		return true
	}

	return false
}

// SetCpuEnabled gets a reference to the given bool and assigns it to the CpuEnabled field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetCpuEnabled(v bool) {
	o.CpuEnabled = &v
}

// GetMinCpu returns the MinCpu field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMinCpu() float64 {
	if o == nil || IsNil(o.MinCpu) {
		var ret float64
		return ret
	}
	return *o.MinCpu
}

// GetMinCpuOk returns a tuple with the MinCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMinCpuOk() (*float64, bool) {
	if o == nil || IsNil(o.MinCpu) {
		return nil, false
	}
	return o.MinCpu, true
}

// IsSetMinCpu returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetMinCpu() bool {
	if o != nil && !IsNil(o.MinCpu) {
		return true
	}

	return false
}

// SetMinCpu gets a reference to the given float64 and assigns it to the MinCpu field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetMinCpu(v float64) {
	o.MinCpu = &v
}

// GetMaxCpu returns the MaxCpu field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMaxCpu() float64 {
	if o == nil || IsNil(o.MaxCpu) {
		var ret float64
		return ret
	}
	return *o.MaxCpu
}

// GetMaxCpuOk returns a tuple with the MaxCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMaxCpuOk() (*float64, bool) {
	if o == nil || IsNil(o.MaxCpu) {
		return nil, false
	}
	return o.MaxCpu, true
}

// IsSetMaxCpu returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetMaxCpu() bool {
	if o != nil && !IsNil(o.MaxCpu) {
		return true
	}

	return false
}

// SetMaxCpu gets a reference to the given float64 and assigns it to the MaxCpu field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetMaxCpu(v float64) {
	o.MaxCpu = &v
}

// GetMemoryEnabled returns the MemoryEnabled field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMemoryEnabled() bool {
	if o == nil || IsNil(o.MemoryEnabled) {
		var ret bool
		return ret
	}
	return *o.MemoryEnabled
}

// GetMemoryEnabledOk returns a tuple with the MemoryEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMemoryEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MemoryEnabled) {
		return nil, false
	}
	return o.MemoryEnabled, true
}

// IsSetMemoryEnabled returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetMemoryEnabled() bool {
	if o != nil && !IsNil(o.MemoryEnabled) {
		return true
	}

	return false
}

// SetMemoryEnabled gets a reference to the given bool and assigns it to the MemoryEnabled field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetMemoryEnabled(v bool) {
	o.MemoryEnabled = &v
}

// GetMinMemory returns the MinMemory field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMinMemory() float64 {
	if o == nil || IsNil(o.MinMemory) {
		var ret float64
		return ret
	}
	return *o.MinMemory
}

// GetMinMemoryOk returns a tuple with the MinMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMinMemoryOk() (*float64, bool) {
	if o == nil || IsNil(o.MinMemory) {
		return nil, false
	}
	return o.MinMemory, true
}

// IsSetMinMemory returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetMinMemory() bool {
	if o != nil && !IsNil(o.MinMemory) {
		return true
	}

	return false
}

// SetMinMemory gets a reference to the given float64 and assigns it to the MinMemory field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetMinMemory(v float64) {
	o.MinMemory = &v
}

// GetMaxMemory returns the MaxMemory field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMaxMemory() float64 {
	if o == nil || IsNil(o.MaxMemory) {
		var ret float64
		return ret
	}
	return *o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMaxMemoryOk() (*float64, bool) {
	if o == nil || IsNil(o.MaxMemory) {
		return nil, false
	}
	return o.MaxMemory, true
}

// IsSetMaxMemory returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetMaxMemory() bool {
	if o != nil && !IsNil(o.MaxMemory) {
		return true
	}

	return false
}

// SetMaxMemory gets a reference to the given float64 and assigns it to the MaxMemory field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetMaxMemory(v float64) {
	o.MaxMemory = &v
}

// GetDiskEnabled returns the DiskEnabled field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetDiskEnabled() bool {
	if o == nil || IsNil(o.DiskEnabled) {
		var ret bool
		return ret
	}
	return *o.DiskEnabled
}

// GetDiskEnabledOk returns a tuple with the DiskEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetDiskEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DiskEnabled) {
		return nil, false
	}
	return o.DiskEnabled, true
}

// IsSetDiskEnabled returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetDiskEnabled() bool {
	if o != nil && !IsNil(o.DiskEnabled) {
		return true
	}

	return false
}

// SetDiskEnabled gets a reference to the given bool and assigns it to the DiskEnabled field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetDiskEnabled(v bool) {
	o.DiskEnabled = &v
}

// GetMinDisk returns the MinDisk field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMinDisk() float64 {
	if o == nil || IsNil(o.MinDisk) {
		var ret float64
		return ret
	}
	return *o.MinDisk
}

// GetMinDiskOk returns a tuple with the MinDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMinDiskOk() (*float64, bool) {
	if o == nil || IsNil(o.MinDisk) {
		return nil, false
	}
	return o.MinDisk, true
}

// IsSetMinDisk returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetMinDisk() bool {
	if o != nil && !IsNil(o.MinDisk) {
		return true
	}

	return false
}

// SetMinDisk gets a reference to the given float64 and assigns it to the MinDisk field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetMinDisk(v float64) {
	o.MinDisk = &v
}

// GetMaxDisk returns the MaxDisk field value if set, zero value otherwise.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMaxDisk() float64 {
	if o == nil || IsNil(o.MaxDisk) {
		var ret float64
		return ret
	}
	return *o.MaxDisk
}

// GetMaxDiskOk returns a tuple with the MaxDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) GetMaxDiskOk() (*float64, bool) {
	if o == nil || IsNil(o.MaxDisk) {
		return nil, false
	}
	return o.MaxDisk, true
}

// IsSetMaxDisk returns a boolean if a field has been set.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) IsSetMaxDisk() bool {
	if o != nil && !IsNil(o.MaxDisk) {
		return true
	}

	return false
}

// SetMaxDisk gets a reference to the given float64 and assigns it to the MaxDisk field.
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) SetMaxDisk(v float64) {
	o.MaxDisk = &v
}

func (o CreateInstanceScheduleRequestInstanceScheduleThreshold) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInstanceScheduleRequestInstanceScheduleThreshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceThresholdId) {
		toSerialize["sourceThresholdId"] = o.SourceThresholdId
	}
	if !IsNil(o.AutoUp) {
		toSerialize["autoUp"] = o.AutoUp
	}
	if !IsNil(o.AutoDown) {
		toSerialize["autoDown"] = o.AutoDown
	}
	if !IsNil(o.MinCount) {
		toSerialize["minCount"] = o.MinCount
	}
	if !IsNil(o.MaxCount) {
		toSerialize["maxCount"] = o.MaxCount
	}
	if !IsNil(o.CpuEnabled) {
		toSerialize["cpuEnabled"] = o.CpuEnabled
	}
	if !IsNil(o.MinCpu) {
		toSerialize["minCpu"] = o.MinCpu
	}
	if !IsNil(o.MaxCpu) {
		toSerialize["maxCpu"] = o.MaxCpu
	}
	if !IsNil(o.MemoryEnabled) {
		toSerialize["memoryEnabled"] = o.MemoryEnabled
	}
	if !IsNil(o.MinMemory) {
		toSerialize["minMemory"] = o.MinMemory
	}
	if !IsNil(o.MaxMemory) {
		toSerialize["maxMemory"] = o.MaxMemory
	}
	if !IsNil(o.DiskEnabled) {
		toSerialize["diskEnabled"] = o.DiskEnabled
	}
	if !IsNil(o.MinDisk) {
		toSerialize["minDisk"] = o.MinDisk
	}
	if !IsNil(o.MaxDisk) {
		toSerialize["maxDisk"] = o.MaxDisk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *CreateInstanceScheduleRequestInstanceScheduleThreshold) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

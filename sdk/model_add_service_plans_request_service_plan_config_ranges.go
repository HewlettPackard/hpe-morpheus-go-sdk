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

// checks if the AddServicePlansRequestServicePlanConfigRanges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddServicePlansRequestServicePlanConfigRanges{}

// AddServicePlansRequestServicePlanConfigRanges struct for AddServicePlansRequestServicePlanConfigRanges
type AddServicePlansRequestServicePlanConfigRanges struct {
	// Custom min storage in GB
	MinStorage *string `json:"minStorage,omitempty"`
	// Custom max storage in GB
	MaxStorage *string `json:"maxStorage,omitempty"`
	// Custom min per disk size in GB
	MinPerDiskSize *string `json:"minPerDiskSize,omitempty"`
	// Custom max per disk size in GB
	MaxPerDiskSize *string `json:"maxPerDiskSize,omitempty"`
	// Custom min memory in bytes
	MinMemory *int64 `json:"minMemory,omitempty"`
	// Custom max memory in bytes
	MaxMemory *int64 `json:"maxMemory,omitempty"`
	// Custom min cores
	MinCores *string `json:"minCores,omitempty"`
	// Custom max cores
	MaxCores *string `json:"maxCores,omitempty"`
	// Custom min sockets
	MinSockets *string `json:"minSockets,omitempty"`
	// Custom max sockets
	MaxSockets *string `json:"maxSockets,omitempty"`
	// Custom min cores allowed per socket
	MinCoresPerSocket *string `json:"minCoresPerSocket,omitempty"`
	// Custom max cores allowed per socket
	MaxCoresPerSocket    *string                `json:"maxCoresPerSocket,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddServicePlansRequestServicePlanConfigRanges AddServicePlansRequestServicePlanConfigRanges

// NewAddServicePlansRequestServicePlanConfigRanges instantiates a new AddServicePlansRequestServicePlanConfigRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddServicePlansRequestServicePlanConfigRanges() *AddServicePlansRequestServicePlanConfigRanges {
	this := AddServicePlansRequestServicePlanConfigRanges{}
	return &this
}

// NewAddServicePlansRequestServicePlanConfigRangesWithDefaults instantiates a new AddServicePlansRequestServicePlanConfigRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddServicePlansRequestServicePlanConfigRangesWithDefaults() *AddServicePlansRequestServicePlanConfigRanges {
	this := AddServicePlansRequestServicePlanConfigRanges{}
	return &this
}

// GetMinStorage returns the MinStorage field value if set, zero value otherwise.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinStorage() string {
	if o == nil || IsNil(o.MinStorage) {
		var ret string
		return ret
	}
	return *o.MinStorage
}

// GetMinStorageOk returns a tuple with the MinStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinStorageOk() (*string, bool) {
	if o == nil || IsNil(o.MinStorage) {
		return nil, false
	}
	return o.MinStorage, true
}

// IsSetMinStorage returns a boolean if a field has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) IsSetMinStorage() bool {
	if o != nil && !IsNil(o.MinStorage) {
		return true
	}

	return false
}

// SetMinStorage gets a reference to the given string and assigns it to the MinStorage field.
func (o *AddServicePlansRequestServicePlanConfigRanges) SetMinStorage(v string) {
	o.MinStorage = &v
}

// GetMaxStorage returns the MaxStorage field value if set, zero value otherwise.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxStorage() string {
	if o == nil || IsNil(o.MaxStorage) {
		var ret string
		return ret
	}
	return *o.MaxStorage
}

// GetMaxStorageOk returns a tuple with the MaxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxStorageOk() (*string, bool) {
	if o == nil || IsNil(o.MaxStorage) {
		return nil, false
	}
	return o.MaxStorage, true
}

// IsSetMaxStorage returns a boolean if a field has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) IsSetMaxStorage() bool {
	if o != nil && !IsNil(o.MaxStorage) {
		return true
	}

	return false
}

// SetMaxStorage gets a reference to the given string and assigns it to the MaxStorage field.
func (o *AddServicePlansRequestServicePlanConfigRanges) SetMaxStorage(v string) {
	o.MaxStorage = &v
}

// GetMinPerDiskSize returns the MinPerDiskSize field value if set, zero value otherwise.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinPerDiskSize() string {
	if o == nil || IsNil(o.MinPerDiskSize) {
		var ret string
		return ret
	}
	return *o.MinPerDiskSize
}

// GetMinPerDiskSizeOk returns a tuple with the MinPerDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinPerDiskSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MinPerDiskSize) {
		return nil, false
	}
	return o.MinPerDiskSize, true
}

// IsSetMinPerDiskSize returns a boolean if a field has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) IsSetMinPerDiskSize() bool {
	if o != nil && !IsNil(o.MinPerDiskSize) {
		return true
	}

	return false
}

// SetMinPerDiskSize gets a reference to the given string and assigns it to the MinPerDiskSize field.
func (o *AddServicePlansRequestServicePlanConfigRanges) SetMinPerDiskSize(v string) {
	o.MinPerDiskSize = &v
}

// GetMaxPerDiskSize returns the MaxPerDiskSize field value if set, zero value otherwise.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxPerDiskSize() string {
	if o == nil || IsNil(o.MaxPerDiskSize) {
		var ret string
		return ret
	}
	return *o.MaxPerDiskSize
}

// GetMaxPerDiskSizeOk returns a tuple with the MaxPerDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxPerDiskSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxPerDiskSize) {
		return nil, false
	}
	return o.MaxPerDiskSize, true
}

// IsSetMaxPerDiskSize returns a boolean if a field has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) IsSetMaxPerDiskSize() bool {
	if o != nil && !IsNil(o.MaxPerDiskSize) {
		return true
	}

	return false
}

// SetMaxPerDiskSize gets a reference to the given string and assigns it to the MaxPerDiskSize field.
func (o *AddServicePlansRequestServicePlanConfigRanges) SetMaxPerDiskSize(v string) {
	o.MaxPerDiskSize = &v
}

// GetMinMemory returns the MinMemory field value if set, zero value otherwise.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinMemory() int64 {
	if o == nil || IsNil(o.MinMemory) {
		var ret int64
		return ret
	}
	return *o.MinMemory
}

// GetMinMemoryOk returns a tuple with the MinMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinMemoryOk() (*int64, bool) {
	if o == nil || IsNil(o.MinMemory) {
		return nil, false
	}
	return o.MinMemory, true
}

// IsSetMinMemory returns a boolean if a field has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) IsSetMinMemory() bool {
	if o != nil && !IsNil(o.MinMemory) {
		return true
	}

	return false
}

// SetMinMemory gets a reference to the given int64 and assigns it to the MinMemory field.
func (o *AddServicePlansRequestServicePlanConfigRanges) SetMinMemory(v int64) {
	o.MinMemory = &v
}

// GetMaxMemory returns the MaxMemory field value if set, zero value otherwise.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxMemory() int64 {
	if o == nil || IsNil(o.MaxMemory) {
		var ret int64
		return ret
	}
	return *o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxMemoryOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxMemory) {
		return nil, false
	}
	return o.MaxMemory, true
}

// IsSetMaxMemory returns a boolean if a field has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) IsSetMaxMemory() bool {
	if o != nil && !IsNil(o.MaxMemory) {
		return true
	}

	return false
}

// SetMaxMemory gets a reference to the given int64 and assigns it to the MaxMemory field.
func (o *AddServicePlansRequestServicePlanConfigRanges) SetMaxMemory(v int64) {
	o.MaxMemory = &v
}

// GetMinCores returns the MinCores field value if set, zero value otherwise.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinCores() string {
	if o == nil || IsNil(o.MinCores) {
		var ret string
		return ret
	}
	return *o.MinCores
}

// GetMinCoresOk returns a tuple with the MinCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinCoresOk() (*string, bool) {
	if o == nil || IsNil(o.MinCores) {
		return nil, false
	}
	return o.MinCores, true
}

// IsSetMinCores returns a boolean if a field has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) IsSetMinCores() bool {
	if o != nil && !IsNil(o.MinCores) {
		return true
	}

	return false
}

// SetMinCores gets a reference to the given string and assigns it to the MinCores field.
func (o *AddServicePlansRequestServicePlanConfigRanges) SetMinCores(v string) {
	o.MinCores = &v
}

// GetMaxCores returns the MaxCores field value if set, zero value otherwise.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxCores() string {
	if o == nil || IsNil(o.MaxCores) {
		var ret string
		return ret
	}
	return *o.MaxCores
}

// GetMaxCoresOk returns a tuple with the MaxCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxCoresOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCores) {
		return nil, false
	}
	return o.MaxCores, true
}

// IsSetMaxCores returns a boolean if a field has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) IsSetMaxCores() bool {
	if o != nil && !IsNil(o.MaxCores) {
		return true
	}

	return false
}

// SetMaxCores gets a reference to the given string and assigns it to the MaxCores field.
func (o *AddServicePlansRequestServicePlanConfigRanges) SetMaxCores(v string) {
	o.MaxCores = &v
}

// GetMinSockets returns the MinSockets field value if set, zero value otherwise.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinSockets() string {
	if o == nil || IsNil(o.MinSockets) {
		var ret string
		return ret
	}
	return *o.MinSockets
}

// GetMinSocketsOk returns a tuple with the MinSockets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinSocketsOk() (*string, bool) {
	if o == nil || IsNil(o.MinSockets) {
		return nil, false
	}
	return o.MinSockets, true
}

// IsSetMinSockets returns a boolean if a field has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) IsSetMinSockets() bool {
	if o != nil && !IsNil(o.MinSockets) {
		return true
	}

	return false
}

// SetMinSockets gets a reference to the given string and assigns it to the MinSockets field.
func (o *AddServicePlansRequestServicePlanConfigRanges) SetMinSockets(v string) {
	o.MinSockets = &v
}

// GetMaxSockets returns the MaxSockets field value if set, zero value otherwise.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxSockets() string {
	if o == nil || IsNil(o.MaxSockets) {
		var ret string
		return ret
	}
	return *o.MaxSockets
}

// GetMaxSocketsOk returns a tuple with the MaxSockets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxSocketsOk() (*string, bool) {
	if o == nil || IsNil(o.MaxSockets) {
		return nil, false
	}
	return o.MaxSockets, true
}

// IsSetMaxSockets returns a boolean if a field has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) IsSetMaxSockets() bool {
	if o != nil && !IsNil(o.MaxSockets) {
		return true
	}

	return false
}

// SetMaxSockets gets a reference to the given string and assigns it to the MaxSockets field.
func (o *AddServicePlansRequestServicePlanConfigRanges) SetMaxSockets(v string) {
	o.MaxSockets = &v
}

// GetMinCoresPerSocket returns the MinCoresPerSocket field value if set, zero value otherwise.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinCoresPerSocket() string {
	if o == nil || IsNil(o.MinCoresPerSocket) {
		var ret string
		return ret
	}
	return *o.MinCoresPerSocket
}

// GetMinCoresPerSocketOk returns a tuple with the MinCoresPerSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinCoresPerSocketOk() (*string, bool) {
	if o == nil || IsNil(o.MinCoresPerSocket) {
		return nil, false
	}
	return o.MinCoresPerSocket, true
}

// IsSetMinCoresPerSocket returns a boolean if a field has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) IsSetMinCoresPerSocket() bool {
	if o != nil && !IsNil(o.MinCoresPerSocket) {
		return true
	}

	return false
}

// SetMinCoresPerSocket gets a reference to the given string and assigns it to the MinCoresPerSocket field.
func (o *AddServicePlansRequestServicePlanConfigRanges) SetMinCoresPerSocket(v string) {
	o.MinCoresPerSocket = &v
}

// GetMaxCoresPerSocket returns the MaxCoresPerSocket field value if set, zero value otherwise.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxCoresPerSocket() string {
	if o == nil || IsNil(o.MaxCoresPerSocket) {
		var ret string
		return ret
	}
	return *o.MaxCoresPerSocket
}

// GetMaxCoresPerSocketOk returns a tuple with the MaxCoresPerSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxCoresPerSocketOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCoresPerSocket) {
		return nil, false
	}
	return o.MaxCoresPerSocket, true
}

// IsSetMaxCoresPerSocket returns a boolean if a field has been set.
func (o *AddServicePlansRequestServicePlanConfigRanges) IsSetMaxCoresPerSocket() bool {
	if o != nil && !IsNil(o.MaxCoresPerSocket) {
		return true
	}

	return false
}

// SetMaxCoresPerSocket gets a reference to the given string and assigns it to the MaxCoresPerSocket field.
func (o *AddServicePlansRequestServicePlanConfigRanges) SetMaxCoresPerSocket(v string) {
	o.MaxCoresPerSocket = &v
}

func (o AddServicePlansRequestServicePlanConfigRanges) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddServicePlansRequestServicePlanConfigRanges) ToMap() (map[string]interface{}, error) {
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
func (o *AddServicePlansRequestServicePlanConfigRanges) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

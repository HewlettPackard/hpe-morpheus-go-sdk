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

// checks if the ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges{}

// ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges struct for ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges
type ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges struct {
	MinStorage           *string                `json:"minStorage,omitempty"`
	MaxStorage           *string                `json:"maxStorage,omitempty"`
	MinMemory            *string                `json:"minMemory,omitempty"`
	MaxMemory            *string                `json:"maxMemory,omitempty"`
	MinCores             *string                `json:"minCores,omitempty"`
	MaxCores             *string                `json:"maxCores,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges

// NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges instantiates a new ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges() *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges {
	this := ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges{}
	return &this
}

// NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRangesWithDefaults instantiates a new ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRangesWithDefaults() *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges {
	this := ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges{}
	return &this
}

// GetMinStorage returns the MinStorage field value if set, zero value otherwise.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) GetMinStorage() string {
	if o == nil || IsNil(o.MinStorage) {
		var ret string
		return ret
	}
	return *o.MinStorage
}

// GetMinStorageOk returns a tuple with the MinStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) GetMinStorageOk() (*string, bool) {
	if o == nil || IsNil(o.MinStorage) {
		return nil, false
	}
	return o.MinStorage, true
}

// IsSetMinStorage returns a boolean if a field has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) IsSetMinStorage() bool {
	if o != nil && !IsNil(o.MinStorage) {
		return true
	}

	return false
}

// SetMinStorage gets a reference to the given string and assigns it to the MinStorage field.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) SetMinStorage(v string) {
	o.MinStorage = &v
}

// GetMaxStorage returns the MaxStorage field value if set, zero value otherwise.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) GetMaxStorage() string {
	if o == nil || IsNil(o.MaxStorage) {
		var ret string
		return ret
	}
	return *o.MaxStorage
}

// GetMaxStorageOk returns a tuple with the MaxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) GetMaxStorageOk() (*string, bool) {
	if o == nil || IsNil(o.MaxStorage) {
		return nil, false
	}
	return o.MaxStorage, true
}

// IsSetMaxStorage returns a boolean if a field has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) IsSetMaxStorage() bool {
	if o != nil && !IsNil(o.MaxStorage) {
		return true
	}

	return false
}

// SetMaxStorage gets a reference to the given string and assigns it to the MaxStorage field.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) SetMaxStorage(v string) {
	o.MaxStorage = &v
}

// GetMinMemory returns the MinMemory field value if set, zero value otherwise.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) GetMinMemory() string {
	if o == nil || IsNil(o.MinMemory) {
		var ret string
		return ret
	}
	return *o.MinMemory
}

// GetMinMemoryOk returns a tuple with the MinMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) GetMinMemoryOk() (*string, bool) {
	if o == nil || IsNil(o.MinMemory) {
		return nil, false
	}
	return o.MinMemory, true
}

// IsSetMinMemory returns a boolean if a field has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) IsSetMinMemory() bool {
	if o != nil && !IsNil(o.MinMemory) {
		return true
	}

	return false
}

// SetMinMemory gets a reference to the given string and assigns it to the MinMemory field.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) SetMinMemory(v string) {
	o.MinMemory = &v
}

// GetMaxMemory returns the MaxMemory field value if set, zero value otherwise.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) GetMaxMemory() string {
	if o == nil || IsNil(o.MaxMemory) {
		var ret string
		return ret
	}
	return *o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) GetMaxMemoryOk() (*string, bool) {
	if o == nil || IsNil(o.MaxMemory) {
		return nil, false
	}
	return o.MaxMemory, true
}

// IsSetMaxMemory returns a boolean if a field has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) IsSetMaxMemory() bool {
	if o != nil && !IsNil(o.MaxMemory) {
		return true
	}

	return false
}

// SetMaxMemory gets a reference to the given string and assigns it to the MaxMemory field.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) SetMaxMemory(v string) {
	o.MaxMemory = &v
}

// GetMinCores returns the MinCores field value if set, zero value otherwise.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) GetMinCores() string {
	if o == nil || IsNil(o.MinCores) {
		var ret string
		return ret
	}
	return *o.MinCores
}

// GetMinCoresOk returns a tuple with the MinCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) GetMinCoresOk() (*string, bool) {
	if o == nil || IsNil(o.MinCores) {
		return nil, false
	}
	return o.MinCores, true
}

// IsSetMinCores returns a boolean if a field has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) IsSetMinCores() bool {
	if o != nil && !IsNil(o.MinCores) {
		return true
	}

	return false
}

// SetMinCores gets a reference to the given string and assigns it to the MinCores field.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) SetMinCores(v string) {
	o.MinCores = &v
}

// GetMaxCores returns the MaxCores field value if set, zero value otherwise.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) GetMaxCores() string {
	if o == nil || IsNil(o.MaxCores) {
		var ret string
		return ret
	}
	return *o.MaxCores
}

// GetMaxCoresOk returns a tuple with the MaxCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) GetMaxCoresOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCores) {
		return nil, false
	}
	return o.MaxCores, true
}

// IsSetMaxCores returns a boolean if a field has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) IsSetMaxCores() bool {
	if o != nil && !IsNil(o.MaxCores) {
		return true
	}

	return false
}

// SetMaxCores gets a reference to the given string and assigns it to the MaxCores field.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) SetMaxCores(v string) {
	o.MaxCores = &v
}

func (o ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinStorage) {
		toSerialize["minStorage"] = o.MinStorage
	}
	if !IsNil(o.MaxStorage) {
		toSerialize["maxStorage"] = o.MaxStorage
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfigRanges) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

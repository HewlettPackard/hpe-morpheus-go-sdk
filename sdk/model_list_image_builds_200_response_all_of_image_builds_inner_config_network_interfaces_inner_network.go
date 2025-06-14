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

// checks if the ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork{}

// ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork struct for ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork
type ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork struct {
	IdName               *string                                                                 `json:"idName,omitempty"`
	Pool                 *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"pool,omitempty"`
	Id                   *string                                                                 `json:"id,omitempty"`
	HasPool              *bool                                                                   `json:"hasPool,omitempty"`
	AdditionalProperties map[string]interface{}                                                  `json:",remain"`
}

type _ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork

// NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork instantiates a new ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork() *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork {
	this := ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork{}
	return &this
}

// NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetworkWithDefaults instantiates a new ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetworkWithDefaults() *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork {
	this := ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork{}
	return &this
}

// GetIdName returns the IdName field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) GetIdName() string {
	if o == nil || IsNil(o.IdName) {
		var ret string
		return ret
	}
	return *o.IdName
}

// GetIdNameOk returns a tuple with the IdName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) GetIdNameOk() (*string, bool) {
	if o == nil || IsNil(o.IdName) {
		return nil, false
	}
	return o.IdName, true
}

// IsSetIdName returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) IsSetIdName() bool {
	if o != nil && !IsNil(o.IdName) {
		return true
	}

	return false
}

// SetIdName gets a reference to the given string and assigns it to the IdName field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) SetIdName(v string) {
	o.IdName = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) GetPool() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Pool) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) GetPoolOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// IsSetPool returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) IsSetPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Pool field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) SetPool(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Pool = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) SetId(v string) {
	o.Id = &v
}

// GetHasPool returns the HasPool field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) GetHasPool() bool {
	if o == nil || IsNil(o.HasPool) {
		var ret bool
		return ret
	}
	return *o.HasPool
}

// GetHasPoolOk returns a tuple with the HasPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) GetHasPoolOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPool) {
		return nil, false
	}
	return o.HasPool, true
}

// IsSetHasPool returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) IsSetHasPool() bool {
	if o != nil && !IsNil(o.HasPool) {
		return true
	}

	return false
}

// SetHasPool gets a reference to the given bool and assigns it to the HasPool field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) SetHasPool(v bool) {
	o.HasPool = &v
}

func (o ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdName) {
		toSerialize["idName"] = o.IdName
	}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.HasPool) {
		toSerialize["hasPool"] = o.HasPool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

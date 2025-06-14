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

// checks if the IntegrationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationType{}

// IntegrationType struct for IntegrationType
type IntegrationType struct {
	Id                      *int64                 `json:"id,omitempty"`
	Code                    *string                `json:"code,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	Description             *string                `json:"description,omitempty"`
	Category                *string                `json:"category,omitempty"`
	Enabled                 *bool                  `json:"enabled,omitempty"`
	Creatable               *bool                  `json:"creatable,omitempty"`
	HasCMDB                 *bool                  `json:"hasCMDB,omitempty"`
	HasCMDBDiscovery        *bool                  `json:"hasCMDBDiscovery,omitempty"`
	HasCM                   *bool                  `json:"hasCM,omitempty"`
	HasDNS                  *bool                  `json:"hasDNS,omitempty"`
	HasApprovals            *bool                  `json:"hasApprovals,omitempty"`
	HasDeleteApprovals      *bool                  `json:"hasDeleteApprovals,omitempty"`
	HasReconfigureApprovals *bool                  `json:"hasReconfigureApprovals,omitempty"`
	IsPlugin                *bool                  `json:"isPlugin,omitempty"`
	AdditionalProperties    map[string]interface{} `json:",remain"`
}

type _IntegrationType IntegrationType

// NewIntegrationType instantiates a new IntegrationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationType() *IntegrationType {
	this := IntegrationType{}
	return &this
}

// NewIntegrationTypeWithDefaults instantiates a new IntegrationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationTypeWithDefaults() *IntegrationType {
	this := IntegrationType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationType) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *IntegrationType) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IntegrationType) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *IntegrationType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *IntegrationType) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *IntegrationType) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IntegrationType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *IntegrationType) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IntegrationType) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IntegrationType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *IntegrationType) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IntegrationType) SetDescription(v string) {
	o.Description = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *IntegrationType) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// IsSetCategory returns a boolean if a field has been set.
func (o *IntegrationType) IsSetCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *IntegrationType) SetCategory(v string) {
	o.Category = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IntegrationType) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *IntegrationType) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IntegrationType) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatable returns the Creatable field value if set, zero value otherwise.
func (o *IntegrationType) GetCreatable() bool {
	if o == nil || IsNil(o.Creatable) {
		var ret bool
		return ret
	}
	return *o.Creatable
}

// GetCreatableOk returns a tuple with the Creatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetCreatableOk() (*bool, bool) {
	if o == nil || IsNil(o.Creatable) {
		return nil, false
	}
	return o.Creatable, true
}

// IsSetCreatable returns a boolean if a field has been set.
func (o *IntegrationType) IsSetCreatable() bool {
	if o != nil && !IsNil(o.Creatable) {
		return true
	}

	return false
}

// SetCreatable gets a reference to the given bool and assigns it to the Creatable field.
func (o *IntegrationType) SetCreatable(v bool) {
	o.Creatable = &v
}

// GetHasCMDB returns the HasCMDB field value if set, zero value otherwise.
func (o *IntegrationType) GetHasCMDB() bool {
	if o == nil || IsNil(o.HasCMDB) {
		var ret bool
		return ret
	}
	return *o.HasCMDB
}

// GetHasCMDBOk returns a tuple with the HasCMDB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetHasCMDBOk() (*bool, bool) {
	if o == nil || IsNil(o.HasCMDB) {
		return nil, false
	}
	return o.HasCMDB, true
}

// IsSetHasCMDB returns a boolean if a field has been set.
func (o *IntegrationType) IsSetHasCMDB() bool {
	if o != nil && !IsNil(o.HasCMDB) {
		return true
	}

	return false
}

// SetHasCMDB gets a reference to the given bool and assigns it to the HasCMDB field.
func (o *IntegrationType) SetHasCMDB(v bool) {
	o.HasCMDB = &v
}

// GetHasCMDBDiscovery returns the HasCMDBDiscovery field value if set, zero value otherwise.
func (o *IntegrationType) GetHasCMDBDiscovery() bool {
	if o == nil || IsNil(o.HasCMDBDiscovery) {
		var ret bool
		return ret
	}
	return *o.HasCMDBDiscovery
}

// GetHasCMDBDiscoveryOk returns a tuple with the HasCMDBDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetHasCMDBDiscoveryOk() (*bool, bool) {
	if o == nil || IsNil(o.HasCMDBDiscovery) {
		return nil, false
	}
	return o.HasCMDBDiscovery, true
}

// IsSetHasCMDBDiscovery returns a boolean if a field has been set.
func (o *IntegrationType) IsSetHasCMDBDiscovery() bool {
	if o != nil && !IsNil(o.HasCMDBDiscovery) {
		return true
	}

	return false
}

// SetHasCMDBDiscovery gets a reference to the given bool and assigns it to the HasCMDBDiscovery field.
func (o *IntegrationType) SetHasCMDBDiscovery(v bool) {
	o.HasCMDBDiscovery = &v
}

// GetHasCM returns the HasCM field value if set, zero value otherwise.
func (o *IntegrationType) GetHasCM() bool {
	if o == nil || IsNil(o.HasCM) {
		var ret bool
		return ret
	}
	return *o.HasCM
}

// GetHasCMOk returns a tuple with the HasCM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetHasCMOk() (*bool, bool) {
	if o == nil || IsNil(o.HasCM) {
		return nil, false
	}
	return o.HasCM, true
}

// IsSetHasCM returns a boolean if a field has been set.
func (o *IntegrationType) IsSetHasCM() bool {
	if o != nil && !IsNil(o.HasCM) {
		return true
	}

	return false
}

// SetHasCM gets a reference to the given bool and assigns it to the HasCM field.
func (o *IntegrationType) SetHasCM(v bool) {
	o.HasCM = &v
}

// GetHasDNS returns the HasDNS field value if set, zero value otherwise.
func (o *IntegrationType) GetHasDNS() bool {
	if o == nil || IsNil(o.HasDNS) {
		var ret bool
		return ret
	}
	return *o.HasDNS
}

// GetHasDNSOk returns a tuple with the HasDNS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetHasDNSOk() (*bool, bool) {
	if o == nil || IsNil(o.HasDNS) {
		return nil, false
	}
	return o.HasDNS, true
}

// IsSetHasDNS returns a boolean if a field has been set.
func (o *IntegrationType) IsSetHasDNS() bool {
	if o != nil && !IsNil(o.HasDNS) {
		return true
	}

	return false
}

// SetHasDNS gets a reference to the given bool and assigns it to the HasDNS field.
func (o *IntegrationType) SetHasDNS(v bool) {
	o.HasDNS = &v
}

// GetHasApprovals returns the HasApprovals field value if set, zero value otherwise.
func (o *IntegrationType) GetHasApprovals() bool {
	if o == nil || IsNil(o.HasApprovals) {
		var ret bool
		return ret
	}
	return *o.HasApprovals
}

// GetHasApprovalsOk returns a tuple with the HasApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetHasApprovalsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasApprovals) {
		return nil, false
	}
	return o.HasApprovals, true
}

// IsSetHasApprovals returns a boolean if a field has been set.
func (o *IntegrationType) IsSetHasApprovals() bool {
	if o != nil && !IsNil(o.HasApprovals) {
		return true
	}

	return false
}

// SetHasApprovals gets a reference to the given bool and assigns it to the HasApprovals field.
func (o *IntegrationType) SetHasApprovals(v bool) {
	o.HasApprovals = &v
}

// GetHasDeleteApprovals returns the HasDeleteApprovals field value if set, zero value otherwise.
func (o *IntegrationType) GetHasDeleteApprovals() bool {
	if o == nil || IsNil(o.HasDeleteApprovals) {
		var ret bool
		return ret
	}
	return *o.HasDeleteApprovals
}

// GetHasDeleteApprovalsOk returns a tuple with the HasDeleteApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetHasDeleteApprovalsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasDeleteApprovals) {
		return nil, false
	}
	return o.HasDeleteApprovals, true
}

// IsSetHasDeleteApprovals returns a boolean if a field has been set.
func (o *IntegrationType) IsSetHasDeleteApprovals() bool {
	if o != nil && !IsNil(o.HasDeleteApprovals) {
		return true
	}

	return false
}

// SetHasDeleteApprovals gets a reference to the given bool and assigns it to the HasDeleteApprovals field.
func (o *IntegrationType) SetHasDeleteApprovals(v bool) {
	o.HasDeleteApprovals = &v
}

// GetHasReconfigureApprovals returns the HasReconfigureApprovals field value if set, zero value otherwise.
func (o *IntegrationType) GetHasReconfigureApprovals() bool {
	if o == nil || IsNil(o.HasReconfigureApprovals) {
		var ret bool
		return ret
	}
	return *o.HasReconfigureApprovals
}

// GetHasReconfigureApprovalsOk returns a tuple with the HasReconfigureApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetHasReconfigureApprovalsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasReconfigureApprovals) {
		return nil, false
	}
	return o.HasReconfigureApprovals, true
}

// IsSetHasReconfigureApprovals returns a boolean if a field has been set.
func (o *IntegrationType) IsSetHasReconfigureApprovals() bool {
	if o != nil && !IsNil(o.HasReconfigureApprovals) {
		return true
	}

	return false
}

// SetHasReconfigureApprovals gets a reference to the given bool and assigns it to the HasReconfigureApprovals field.
func (o *IntegrationType) SetHasReconfigureApprovals(v bool) {
	o.HasReconfigureApprovals = &v
}

// GetIsPlugin returns the IsPlugin field value if set, zero value otherwise.
func (o *IntegrationType) GetIsPlugin() bool {
	if o == nil || IsNil(o.IsPlugin) {
		var ret bool
		return ret
	}
	return *o.IsPlugin
}

// GetIsPluginOk returns a tuple with the IsPlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationType) GetIsPluginOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPlugin) {
		return nil, false
	}
	return o.IsPlugin, true
}

// IsSetIsPlugin returns a boolean if a field has been set.
func (o *IntegrationType) IsSetIsPlugin() bool {
	if o != nil && !IsNil(o.IsPlugin) {
		return true
	}

	return false
}

// SetIsPlugin gets a reference to the given bool and assigns it to the IsPlugin field.
func (o *IntegrationType) SetIsPlugin(v bool) {
	o.IsPlugin = &v
}

func (o IntegrationType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Creatable) {
		toSerialize["creatable"] = o.Creatable
	}
	if !IsNil(o.HasCMDB) {
		toSerialize["hasCMDB"] = o.HasCMDB
	}
	if !IsNil(o.HasCMDBDiscovery) {
		toSerialize["hasCMDBDiscovery"] = o.HasCMDBDiscovery
	}
	if !IsNil(o.HasCM) {
		toSerialize["hasCM"] = o.HasCM
	}
	if !IsNil(o.HasDNS) {
		toSerialize["hasDNS"] = o.HasDNS
	}
	if !IsNil(o.HasApprovals) {
		toSerialize["hasApprovals"] = o.HasApprovals
	}
	if !IsNil(o.HasDeleteApprovals) {
		toSerialize["hasDeleteApprovals"] = o.HasDeleteApprovals
	}
	if !IsNil(o.HasReconfigureApprovals) {
		toSerialize["hasReconfigureApprovals"] = o.HasReconfigureApprovals
	}
	if !IsNil(o.IsPlugin) {
		toSerialize["isPlugin"] = o.IsPlugin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *IntegrationType) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

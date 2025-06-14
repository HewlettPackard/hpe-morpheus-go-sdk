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

// checks if the ListTaskTypes200ResponseTaskTypesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTaskTypes200ResponseTaskTypesInner{}

// ListTaskTypes200ResponseTaskTypesInner struct for ListTaskTypes200ResponseTaskTypesInner
type ListTaskTypes200ResponseTaskTypesInner struct {
	Id                   *int64                                                   `json:"id,omitempty"`
	Code                 *string                                                  `json:"code,omitempty"`
	Name                 *string                                                  `json:"name,omitempty"`
	Category             *string                                                  `json:"category,omitempty"`
	Description          *string                                                  `json:"description,omitempty"`
	Scriptable           *bool                                                    `json:"scriptable,omitempty"`
	Enabled              *bool                                                    `json:"enabled,omitempty"`
	HasResults           *bool                                                    `json:"hasResults,omitempty"`
	AllowExecuteLocal    *bool                                                    `json:"allowExecuteLocal,omitempty"`
	AllowExecuteRemote   *bool                                                    `json:"allowExecuteRemote,omitempty"`
	AllowExecuteResource *bool                                                    `json:"allowExecuteResource,omitempty"`
	AllowLocalRepo       *bool                                                    `json:"allowLocalRepo,omitempty"`
	AllowRemoteKeyAuth   *bool                                                    `json:"allowRemoteKeyAuth,omitempty"`
	OptionTypes          []ListTaskTypes200ResponseTaskTypesInnerOptionTypesInner `json:"optionTypes,omitempty"`
	AdditionalProperties map[string]interface{}                                   `json:",remain"`
}

type _ListTaskTypes200ResponseTaskTypesInner ListTaskTypes200ResponseTaskTypesInner

// NewListTaskTypes200ResponseTaskTypesInner instantiates a new ListTaskTypes200ResponseTaskTypesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTaskTypes200ResponseTaskTypesInner() *ListTaskTypes200ResponseTaskTypesInner {
	this := ListTaskTypes200ResponseTaskTypesInner{}
	return &this
}

// NewListTaskTypes200ResponseTaskTypesInnerWithDefaults instantiates a new ListTaskTypes200ResponseTaskTypesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTaskTypes200ResponseTaskTypesInnerWithDefaults() *ListTaskTypes200ResponseTaskTypesInner {
	this := ListTaskTypes200ResponseTaskTypesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetName(v string) {
	o.Name = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// IsSetCategory returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetDescription(v string) {
	o.Description = &v
}

// GetScriptable returns the Scriptable field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetScriptable() bool {
	if o == nil || IsNil(o.Scriptable) {
		var ret bool
		return ret
	}
	return *o.Scriptable
}

// GetScriptableOk returns a tuple with the Scriptable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetScriptableOk() (*bool, bool) {
	if o == nil || IsNil(o.Scriptable) {
		return nil, false
	}
	return o.Scriptable, true
}

// IsSetScriptable returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetScriptable() bool {
	if o != nil && !IsNil(o.Scriptable) {
		return true
	}

	return false
}

// SetScriptable gets a reference to the given bool and assigns it to the Scriptable field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetScriptable(v bool) {
	o.Scriptable = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHasResults returns the HasResults field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetHasResults() bool {
	if o == nil || IsNil(o.HasResults) {
		var ret bool
		return ret
	}
	return *o.HasResults
}

// GetHasResultsOk returns a tuple with the HasResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetHasResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasResults) {
		return nil, false
	}
	return o.HasResults, true
}

// IsSetHasResults returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetHasResults() bool {
	if o != nil && !IsNil(o.HasResults) {
		return true
	}

	return false
}

// SetHasResults gets a reference to the given bool and assigns it to the HasResults field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetHasResults(v bool) {
	o.HasResults = &v
}

// GetAllowExecuteLocal returns the AllowExecuteLocal field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowExecuteLocal() bool {
	if o == nil || IsNil(o.AllowExecuteLocal) {
		var ret bool
		return ret
	}
	return *o.AllowExecuteLocal
}

// GetAllowExecuteLocalOk returns a tuple with the AllowExecuteLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowExecuteLocalOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowExecuteLocal) {
		return nil, false
	}
	return o.AllowExecuteLocal, true
}

// IsSetAllowExecuteLocal returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetAllowExecuteLocal() bool {
	if o != nil && !IsNil(o.AllowExecuteLocal) {
		return true
	}

	return false
}

// SetAllowExecuteLocal gets a reference to the given bool and assigns it to the AllowExecuteLocal field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetAllowExecuteLocal(v bool) {
	o.AllowExecuteLocal = &v
}

// GetAllowExecuteRemote returns the AllowExecuteRemote field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowExecuteRemote() bool {
	if o == nil || IsNil(o.AllowExecuteRemote) {
		var ret bool
		return ret
	}
	return *o.AllowExecuteRemote
}

// GetAllowExecuteRemoteOk returns a tuple with the AllowExecuteRemote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowExecuteRemoteOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowExecuteRemote) {
		return nil, false
	}
	return o.AllowExecuteRemote, true
}

// IsSetAllowExecuteRemote returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetAllowExecuteRemote() bool {
	if o != nil && !IsNil(o.AllowExecuteRemote) {
		return true
	}

	return false
}

// SetAllowExecuteRemote gets a reference to the given bool and assigns it to the AllowExecuteRemote field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetAllowExecuteRemote(v bool) {
	o.AllowExecuteRemote = &v
}

// GetAllowExecuteResource returns the AllowExecuteResource field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowExecuteResource() bool {
	if o == nil || IsNil(o.AllowExecuteResource) {
		var ret bool
		return ret
	}
	return *o.AllowExecuteResource
}

// GetAllowExecuteResourceOk returns a tuple with the AllowExecuteResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowExecuteResourceOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowExecuteResource) {
		return nil, false
	}
	return o.AllowExecuteResource, true
}

// IsSetAllowExecuteResource returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetAllowExecuteResource() bool {
	if o != nil && !IsNil(o.AllowExecuteResource) {
		return true
	}

	return false
}

// SetAllowExecuteResource gets a reference to the given bool and assigns it to the AllowExecuteResource field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetAllowExecuteResource(v bool) {
	o.AllowExecuteResource = &v
}

// GetAllowLocalRepo returns the AllowLocalRepo field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowLocalRepo() bool {
	if o == nil || IsNil(o.AllowLocalRepo) {
		var ret bool
		return ret
	}
	return *o.AllowLocalRepo
}

// GetAllowLocalRepoOk returns a tuple with the AllowLocalRepo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowLocalRepoOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowLocalRepo) {
		return nil, false
	}
	return o.AllowLocalRepo, true
}

// IsSetAllowLocalRepo returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetAllowLocalRepo() bool {
	if o != nil && !IsNil(o.AllowLocalRepo) {
		return true
	}

	return false
}

// SetAllowLocalRepo gets a reference to the given bool and assigns it to the AllowLocalRepo field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetAllowLocalRepo(v bool) {
	o.AllowLocalRepo = &v
}

// GetAllowRemoteKeyAuth returns the AllowRemoteKeyAuth field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowRemoteKeyAuth() bool {
	if o == nil || IsNil(o.AllowRemoteKeyAuth) {
		var ret bool
		return ret
	}
	return *o.AllowRemoteKeyAuth
}

// GetAllowRemoteKeyAuthOk returns a tuple with the AllowRemoteKeyAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetAllowRemoteKeyAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowRemoteKeyAuth) {
		return nil, false
	}
	return o.AllowRemoteKeyAuth, true
}

// IsSetAllowRemoteKeyAuth returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetAllowRemoteKeyAuth() bool {
	if o != nil && !IsNil(o.AllowRemoteKeyAuth) {
		return true
	}

	return false
}

// SetAllowRemoteKeyAuth gets a reference to the given bool and assigns it to the AllowRemoteKeyAuth field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetAllowRemoteKeyAuth(v bool) {
	o.AllowRemoteKeyAuth = &v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetOptionTypes() []ListTaskTypes200ResponseTaskTypesInnerOptionTypesInner {
	if o == nil || IsNil(o.OptionTypes) {
		var ret []ListTaskTypes200ResponseTaskTypesInnerOptionTypesInner
		return ret
	}
	return o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) GetOptionTypesOk() ([]ListTaskTypes200ResponseTaskTypesInnerOptionTypesInner, bool) {
	if o == nil || IsNil(o.OptionTypes) {
		return nil, false
	}
	return o.OptionTypes, true
}

// IsSetOptionTypes returns a boolean if a field has been set.
func (o *ListTaskTypes200ResponseTaskTypesInner) IsSetOptionTypes() bool {
	if o != nil && !IsNil(o.OptionTypes) {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []ListTaskTypes200ResponseTaskTypesInnerOptionTypesInner and assigns it to the OptionTypes field.
func (o *ListTaskTypes200ResponseTaskTypesInner) SetOptionTypes(v []ListTaskTypes200ResponseTaskTypesInnerOptionTypesInner) {
	o.OptionTypes = v
}

func (o ListTaskTypes200ResponseTaskTypesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTaskTypes200ResponseTaskTypesInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Scriptable) {
		toSerialize["scriptable"] = o.Scriptable
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.HasResults) {
		toSerialize["hasResults"] = o.HasResults
	}
	if !IsNil(o.AllowExecuteLocal) {
		toSerialize["allowExecuteLocal"] = o.AllowExecuteLocal
	}
	if !IsNil(o.AllowExecuteRemote) {
		toSerialize["allowExecuteRemote"] = o.AllowExecuteRemote
	}
	if !IsNil(o.AllowExecuteResource) {
		toSerialize["allowExecuteResource"] = o.AllowExecuteResource
	}
	if !IsNil(o.AllowLocalRepo) {
		toSerialize["allowLocalRepo"] = o.AllowLocalRepo
	}
	if !IsNil(o.AllowRemoteKeyAuth) {
		toSerialize["allowRemoteKeyAuth"] = o.AllowRemoteKeyAuth
	}
	if !IsNil(o.OptionTypes) {
		toSerialize["optionTypes"] = o.OptionTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListTaskTypes200ResponseTaskTypesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

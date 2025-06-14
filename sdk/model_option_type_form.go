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

// checks if the OptionTypeForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionTypeForm{}

// OptionTypeForm struct for OptionTypeForm
type OptionTypeForm struct {
	Id                   *int64                                                            `json:"id,omitempty"`
	Name                 *string                                                           `json:"name,omitempty"`
	Code                 *string                                                           `json:"code,omitempty"`
	Description          *string                                                           `json:"description,omitempty"`
	Context              *string                                                           `json:"context,omitempty"`
	Locked               *bool                                                             `json:"locked,omitempty"`
	Labels               []string                                                          `json:"labels,omitempty"`
	Options              []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner     `json:"options,omitempty"`
	FieldGroups          []ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner `json:"fieldGroups,omitempty"`
	AdditionalProperties map[string]interface{}                                            `json:",remain"`
}

type _OptionTypeForm OptionTypeForm

// NewOptionTypeForm instantiates a new OptionTypeForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionTypeForm() *OptionTypeForm {
	this := OptionTypeForm{}
	return &this
}

// NewOptionTypeFormWithDefaults instantiates a new OptionTypeForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionTypeFormWithDefaults() *OptionTypeForm {
	this := OptionTypeForm{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OptionTypeForm) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeForm) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *OptionTypeForm) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OptionTypeForm) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OptionTypeForm) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeForm) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *OptionTypeForm) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OptionTypeForm) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptionTypeForm) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeForm) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *OptionTypeForm) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptionTypeForm) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OptionTypeForm) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeForm) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *OptionTypeForm) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OptionTypeForm) SetDescription(v string) {
	o.Description = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *OptionTypeForm) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeForm) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// IsSetContext returns a boolean if a field has been set.
func (o *OptionTypeForm) IsSetContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *OptionTypeForm) SetContext(v string) {
	o.Context = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *OptionTypeForm) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeForm) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// IsSetLocked returns a boolean if a field has been set.
func (o *OptionTypeForm) IsSetLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *OptionTypeForm) SetLocked(v bool) {
	o.Locked = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *OptionTypeForm) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeForm) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *OptionTypeForm) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *OptionTypeForm) SetLabels(v []string) {
	o.Labels = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *OptionTypeForm) GetOptions() []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner {
	if o == nil || IsNil(o.Options) {
		var ret []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeForm) GetOptionsOk() ([]ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// IsSetOptions returns a boolean if a field has been set.
func (o *OptionTypeForm) IsSetOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner and assigns it to the Options field.
func (o *OptionTypeForm) SetOptions(v []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) {
	o.Options = v
}

// GetFieldGroups returns the FieldGroups field value if set, zero value otherwise.
func (o *OptionTypeForm) GetFieldGroups() []ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner {
	if o == nil || IsNil(o.FieldGroups) {
		var ret []ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner
		return ret
	}
	return o.FieldGroups
}

// GetFieldGroupsOk returns a tuple with the FieldGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionTypeForm) GetFieldGroupsOk() ([]ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner, bool) {
	if o == nil || IsNil(o.FieldGroups) {
		return nil, false
	}
	return o.FieldGroups, true
}

// IsSetFieldGroups returns a boolean if a field has been set.
func (o *OptionTypeForm) IsSetFieldGroups() bool {
	if o != nil && !IsNil(o.FieldGroups) {
		return true
	}

	return false
}

// SetFieldGroups gets a reference to the given []ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner and assigns it to the FieldGroups field.
func (o *OptionTypeForm) SetFieldGroups(v []ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner) {
	o.FieldGroups = v
}

func (o OptionTypeForm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionTypeForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.FieldGroups) {
		toSerialize["fieldGroups"] = o.FieldGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *OptionTypeForm) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

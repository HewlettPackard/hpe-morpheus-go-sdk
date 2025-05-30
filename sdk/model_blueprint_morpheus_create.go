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
	"bytes"
	"fmt"
)

// checks if the BlueprintMorpheusCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintMorpheusCreate{}

// BlueprintMorpheusCreate struct for BlueprintMorpheusCreate
type BlueprintMorpheusCreate struct {
	// A name for the blueprint
	Name string `json:"name"`
	// Blueprint Type
	Type string `json:"type"`
	// Array of label strings, can be used for filtering.
	Labels []string `json:"labels,omitempty"`
	// Tier definitions - Create in UI to view a baseline for object
	Tiers map[string]interface{} `json:"tiers"`
}

type _BlueprintMorpheusCreate BlueprintMorpheusCreate

// NewBlueprintMorpheusCreate instantiates a new BlueprintMorpheusCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintMorpheusCreate(name string, type_ string, tiers map[string]interface{}) *BlueprintMorpheusCreate {
	this := BlueprintMorpheusCreate{}
	this.Name = name
	this.Type = type_
	this.Tiers = tiers
	return &this
}

// NewBlueprintMorpheusCreateWithDefaults instantiates a new BlueprintMorpheusCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintMorpheusCreateWithDefaults() *BlueprintMorpheusCreate {
	this := BlueprintMorpheusCreate{}
	return &this
}

// GetName returns the Name field value
func (o *BlueprintMorpheusCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlueprintMorpheusCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BlueprintMorpheusCreate) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *BlueprintMorpheusCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BlueprintMorpheusCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BlueprintMorpheusCreate) SetType(v string) {
	o.Type = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *BlueprintMorpheusCreate) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintMorpheusCreate) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *BlueprintMorpheusCreate) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *BlueprintMorpheusCreate) SetLabels(v []string) {
	o.Labels = v
}

// GetTiers returns the Tiers field value
func (o *BlueprintMorpheusCreate) GetTiers() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value
// and a boolean to check if the value has been set.
func (o *BlueprintMorpheusCreate) GetTiersOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Tiers, true
}

// SetTiers sets field value
func (o *BlueprintMorpheusCreate) SetTiers(v map[string]interface{}) {
	o.Tiers = v
}

func (o BlueprintMorpheusCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintMorpheusCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["tiers"] = o.Tiers
	return toSerialize, nil
}

func (o *BlueprintMorpheusCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"tiers",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBlueprintMorpheusCreate := _BlueprintMorpheusCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlueprintMorpheusCreate)

	if err != nil {
		return err
	}

	*o = BlueprintMorpheusCreate(varBlueprintMorpheusCreate)

	return err
}

type NullableBlueprintMorpheusCreate struct {
	value *BlueprintMorpheusCreate
	isSet bool
}

func (v NullableBlueprintMorpheusCreate) Get() *BlueprintMorpheusCreate {
	return v.value
}

func (v *NullableBlueprintMorpheusCreate) Set(val *BlueprintMorpheusCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintMorpheusCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintMorpheusCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintMorpheusCreate(val *BlueprintMorpheusCreate) *NullableBlueprintMorpheusCreate {
	return &NullableBlueprintMorpheusCreate{value: val, isSet: true}
}

func (v NullableBlueprintMorpheusCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintMorpheusCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



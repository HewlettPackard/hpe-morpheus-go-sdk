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

// checks if the BlueprintCFTCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintCFTCreate{}

// BlueprintCFTCreate struct for BlueprintCFTCreate
type BlueprintCFTCreate struct {
	// A name for the blueprint
	Name string `json:"name"`
	// Blueprint Type
	Type string `json:"type"`
	// Array of label strings, can be used for filtering.
	Labels []string `json:"labels,omitempty"`
	CloudFormation AddBlueprintRequestOneOf1CloudFormation `json:"cloudFormation"`
}

type _BlueprintCFTCreate BlueprintCFTCreate

// NewBlueprintCFTCreate instantiates a new BlueprintCFTCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintCFTCreate(name string, type_ string, cloudFormation AddBlueprintRequestOneOf1CloudFormation) *BlueprintCFTCreate {
	this := BlueprintCFTCreate{}
	this.Name = name
	this.Type = type_
	this.CloudFormation = cloudFormation
	return &this
}

// NewBlueprintCFTCreateWithDefaults instantiates a new BlueprintCFTCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintCFTCreateWithDefaults() *BlueprintCFTCreate {
	this := BlueprintCFTCreate{}
	return &this
}

// GetName returns the Name field value
func (o *BlueprintCFTCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlueprintCFTCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BlueprintCFTCreate) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *BlueprintCFTCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BlueprintCFTCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BlueprintCFTCreate) SetType(v string) {
	o.Type = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *BlueprintCFTCreate) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCFTCreate) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *BlueprintCFTCreate) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *BlueprintCFTCreate) SetLabels(v []string) {
	o.Labels = v
}

// GetCloudFormation returns the CloudFormation field value
func (o *BlueprintCFTCreate) GetCloudFormation() AddBlueprintRequestOneOf1CloudFormation {
	if o == nil {
		var ret AddBlueprintRequestOneOf1CloudFormation
		return ret
	}

	return o.CloudFormation
}

// GetCloudFormationOk returns a tuple with the CloudFormation field value
// and a boolean to check if the value has been set.
func (o *BlueprintCFTCreate) GetCloudFormationOk() (*AddBlueprintRequestOneOf1CloudFormation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudFormation, true
}

// SetCloudFormation sets field value
func (o *BlueprintCFTCreate) SetCloudFormation(v AddBlueprintRequestOneOf1CloudFormation) {
	o.CloudFormation = v
}

func (o BlueprintCFTCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintCFTCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["cloudFormation"] = o.CloudFormation
	return toSerialize, nil
}

func (o *BlueprintCFTCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"cloudFormation",
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

	varBlueprintCFTCreate := _BlueprintCFTCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlueprintCFTCreate)

	if err != nil {
		return err
	}

	*o = BlueprintCFTCreate(varBlueprintCFTCreate)

	return err
}

type NullableBlueprintCFTCreate struct {
	value *BlueprintCFTCreate
	isSet bool
}

func (v NullableBlueprintCFTCreate) Get() *BlueprintCFTCreate {
	return v.value
}

func (v *NullableBlueprintCFTCreate) Set(val *BlueprintCFTCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintCFTCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintCFTCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintCFTCreate(val *BlueprintCFTCreate) *NullableBlueprintCFTCreate {
	return &NullableBlueprintCFTCreate{value: val, isSet: true}
}

func (v NullableBlueprintCFTCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintCFTCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the Activity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Activity{}

// Activity struct for Activity
type Activity struct {
	Activity             []ListActivity200ResponseAllOfActivityInnerActivityInner `json:"activity,omitempty"`
	AdditionalProperties map[string]interface{}                                   `json:",remain"`
}

type _Activity Activity

// NewActivity instantiates a new Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity() *Activity {
	this := Activity{}
	return &this
}

// NewActivityWithDefaults instantiates a new Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityWithDefaults() *Activity {
	this := Activity{}
	return &this
}

// GetActivity returns the Activity field value if set, zero value otherwise.
func (o *Activity) GetActivity() []ListActivity200ResponseAllOfActivityInnerActivityInner {
	if o == nil || IsNil(o.Activity) {
		var ret []ListActivity200ResponseAllOfActivityInnerActivityInner
		return ret
	}
	return o.Activity
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetActivityOk() ([]ListActivity200ResponseAllOfActivityInnerActivityInner, bool) {
	if o == nil || IsNil(o.Activity) {
		return nil, false
	}
	return o.Activity, true
}

// IsSetActivity returns a boolean if a field has been set.
func (o *Activity) IsSetActivity() bool {
	if o != nil && !IsNil(o.Activity) {
		return true
	}

	return false
}

// SetActivity gets a reference to the given []ListActivity200ResponseAllOfActivityInnerActivityInner and assigns it to the Activity field.
func (o *Activity) SetActivity(v []ListActivity200ResponseAllOfActivityInnerActivityInner) {
	o.Activity = v
}

func (o Activity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Activity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Activity) {
		toSerialize["activity"] = o.Activity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *Activity) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

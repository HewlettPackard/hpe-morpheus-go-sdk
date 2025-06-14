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

// checks if the ResizeInstanceRequestInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResizeInstanceRequestInstance{}

// ResizeInstanceRequestInstance The map containing the id of the service plan you wish to apply to the containers in this instance.
type ResizeInstanceRequestInstance struct {
	Plan                 *ResizeInstanceRequestInstancePlan `json:"plan,omitempty"`
	AdditionalProperties map[string]interface{}             `json:",remain"`
}

type _ResizeInstanceRequestInstance ResizeInstanceRequestInstance

// NewResizeInstanceRequestInstance instantiates a new ResizeInstanceRequestInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResizeInstanceRequestInstance() *ResizeInstanceRequestInstance {
	this := ResizeInstanceRequestInstance{}
	return &this
}

// NewResizeInstanceRequestInstanceWithDefaults instantiates a new ResizeInstanceRequestInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResizeInstanceRequestInstanceWithDefaults() *ResizeInstanceRequestInstance {
	this := ResizeInstanceRequestInstance{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *ResizeInstanceRequestInstance) GetPlan() ResizeInstanceRequestInstancePlan {
	if o == nil || IsNil(o.Plan) {
		var ret ResizeInstanceRequestInstancePlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResizeInstanceRequestInstance) GetPlanOk() (*ResizeInstanceRequestInstancePlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// IsSetPlan returns a boolean if a field has been set.
func (o *ResizeInstanceRequestInstance) IsSetPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given ResizeInstanceRequestInstancePlan and assigns it to the Plan field.
func (o *ResizeInstanceRequestInstance) SetPlan(v ResizeInstanceRequestInstancePlan) {
	o.Plan = &v
}

func (o ResizeInstanceRequestInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResizeInstanceRequestInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ResizeInstanceRequestInstance) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

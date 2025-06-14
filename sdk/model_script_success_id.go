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

// checks if the ScriptSuccessId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScriptSuccessId{}

// ScriptSuccessId struct for ScriptSuccessId
type ScriptSuccessId struct {
	Success              *bool                                        `json:"success,omitempty"`
	ContainerScript      *GetAlerts200ResponseAllOfChecksInnerAccount `json:"containerScript,omitempty"`
	AdditionalProperties map[string]interface{}                       `json:",remain"`
}

type _ScriptSuccessId ScriptSuccessId

// NewScriptSuccessId instantiates a new ScriptSuccessId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptSuccessId() *ScriptSuccessId {
	this := ScriptSuccessId{}
	return &this
}

// NewScriptSuccessIdWithDefaults instantiates a new ScriptSuccessId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptSuccessIdWithDefaults() *ScriptSuccessId {
	this := ScriptSuccessId{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ScriptSuccessId) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptSuccessId) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *ScriptSuccessId) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ScriptSuccessId) SetSuccess(v bool) {
	o.Success = &v
}

// GetContainerScript returns the ContainerScript field value if set, zero value otherwise.
func (o *ScriptSuccessId) GetContainerScript() GetAlerts200ResponseAllOfChecksInnerAccount {
	if o == nil || IsNil(o.ContainerScript) {
		var ret GetAlerts200ResponseAllOfChecksInnerAccount
		return ret
	}
	return *o.ContainerScript
}

// GetContainerScriptOk returns a tuple with the ContainerScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptSuccessId) GetContainerScriptOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool) {
	if o == nil || IsNil(o.ContainerScript) {
		return nil, false
	}
	return o.ContainerScript, true
}

// IsSetContainerScript returns a boolean if a field has been set.
func (o *ScriptSuccessId) IsSetContainerScript() bool {
	if o != nil && !IsNil(o.ContainerScript) {
		return true
	}

	return false
}

// SetContainerScript gets a reference to the given GetAlerts200ResponseAllOfChecksInnerAccount and assigns it to the ContainerScript field.
func (o *ScriptSuccessId) SetContainerScript(v GetAlerts200ResponseAllOfChecksInnerAccount) {
	o.ContainerScript = &v
}

func (o ScriptSuccessId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScriptSuccessId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.ContainerScript) {
		toSerialize["containerScript"] = o.ContainerScript
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ScriptSuccessId) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

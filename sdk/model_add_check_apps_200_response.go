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

// checks if the AddCheckApps200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCheckApps200Response{}

// AddCheckApps200Response struct for AddCheckApps200Response
type AddCheckApps200Response struct {
	CheckApp             *GetAlerts200ResponseAllOfAppsInner `json:"checkApp,omitempty"`
	Success              *bool                               `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}              `json:",remain"`
}

type _AddCheckApps200Response AddCheckApps200Response

// NewAddCheckApps200Response instantiates a new AddCheckApps200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCheckApps200Response() *AddCheckApps200Response {
	this := AddCheckApps200Response{}
	return &this
}

// NewAddCheckApps200ResponseWithDefaults instantiates a new AddCheckApps200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCheckApps200ResponseWithDefaults() *AddCheckApps200Response {
	this := AddCheckApps200Response{}
	return &this
}

// GetCheckApp returns the CheckApp field value if set, zero value otherwise.
func (o *AddCheckApps200Response) GetCheckApp() GetAlerts200ResponseAllOfAppsInner {
	if o == nil || IsNil(o.CheckApp) {
		var ret GetAlerts200ResponseAllOfAppsInner
		return ret
	}
	return *o.CheckApp
}

// GetCheckAppOk returns a tuple with the CheckApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCheckApps200Response) GetCheckAppOk() (*GetAlerts200ResponseAllOfAppsInner, bool) {
	if o == nil || IsNil(o.CheckApp) {
		return nil, false
	}
	return o.CheckApp, true
}

// IsSetCheckApp returns a boolean if a field has been set.
func (o *AddCheckApps200Response) IsSetCheckApp() bool {
	if o != nil && !IsNil(o.CheckApp) {
		return true
	}

	return false
}

// SetCheckApp gets a reference to the given GetAlerts200ResponseAllOfAppsInner and assigns it to the CheckApp field.
func (o *AddCheckApps200Response) SetCheckApp(v GetAlerts200ResponseAllOfAppsInner) {
	o.CheckApp = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AddCheckApps200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCheckApps200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *AddCheckApps200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AddCheckApps200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o AddCheckApps200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCheckApps200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CheckApp) {
		toSerialize["checkApp"] = o.CheckApp
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddCheckApps200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

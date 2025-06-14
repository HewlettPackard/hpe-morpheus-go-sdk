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

// checks if the GetLicense200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLicense200Response{}

// GetLicense200Response struct for GetLicense200Response
type GetLicense200Response struct {
	License *GetLicense200ResponseLicense `json:"license,omitempty"`
	// List of all the installed licenses
	InstalledLicenses    []GetLicense200ResponseInstalledLicensesInner `json:"installedLicenses,omitempty"`
	CurrentUsage         *GetLicense200ResponseCurrentUsage            `json:"currentUsage,omitempty"`
	AdditionalProperties map[string]interface{}                        `json:",remain"`
}

type _GetLicense200Response GetLicense200Response

// NewGetLicense200Response instantiates a new GetLicense200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLicense200Response() *GetLicense200Response {
	this := GetLicense200Response{}
	return &this
}

// NewGetLicense200ResponseWithDefaults instantiates a new GetLicense200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLicense200ResponseWithDefaults() *GetLicense200Response {
	this := GetLicense200Response{}
	return &this
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *GetLicense200Response) GetLicense() GetLicense200ResponseLicense {
	if o == nil || IsNil(o.License) {
		var ret GetLicense200ResponseLicense
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLicense200Response) GetLicenseOk() (*GetLicense200ResponseLicense, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// IsSetLicense returns a boolean if a field has been set.
func (o *GetLicense200Response) IsSetLicense() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given GetLicense200ResponseLicense and assigns it to the License field.
func (o *GetLicense200Response) SetLicense(v GetLicense200ResponseLicense) {
	o.License = &v
}

// GetInstalledLicenses returns the InstalledLicenses field value if set, zero value otherwise.
func (o *GetLicense200Response) GetInstalledLicenses() []GetLicense200ResponseInstalledLicensesInner {
	if o == nil || IsNil(o.InstalledLicenses) {
		var ret []GetLicense200ResponseInstalledLicensesInner
		return ret
	}
	return o.InstalledLicenses
}

// GetInstalledLicensesOk returns a tuple with the InstalledLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLicense200Response) GetInstalledLicensesOk() ([]GetLicense200ResponseInstalledLicensesInner, bool) {
	if o == nil || IsNil(o.InstalledLicenses) {
		return nil, false
	}
	return o.InstalledLicenses, true
}

// IsSetInstalledLicenses returns a boolean if a field has been set.
func (o *GetLicense200Response) IsSetInstalledLicenses() bool {
	if o != nil && !IsNil(o.InstalledLicenses) {
		return true
	}

	return false
}

// SetInstalledLicenses gets a reference to the given []GetLicense200ResponseInstalledLicensesInner and assigns it to the InstalledLicenses field.
func (o *GetLicense200Response) SetInstalledLicenses(v []GetLicense200ResponseInstalledLicensesInner) {
	o.InstalledLicenses = v
}

// GetCurrentUsage returns the CurrentUsage field value if set, zero value otherwise.
func (o *GetLicense200Response) GetCurrentUsage() GetLicense200ResponseCurrentUsage {
	if o == nil || IsNil(o.CurrentUsage) {
		var ret GetLicense200ResponseCurrentUsage
		return ret
	}
	return *o.CurrentUsage
}

// GetCurrentUsageOk returns a tuple with the CurrentUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLicense200Response) GetCurrentUsageOk() (*GetLicense200ResponseCurrentUsage, bool) {
	if o == nil || IsNil(o.CurrentUsage) {
		return nil, false
	}
	return o.CurrentUsage, true
}

// IsSetCurrentUsage returns a boolean if a field has been set.
func (o *GetLicense200Response) IsSetCurrentUsage() bool {
	if o != nil && !IsNil(o.CurrentUsage) {
		return true
	}

	return false
}

// SetCurrentUsage gets a reference to the given GetLicense200ResponseCurrentUsage and assigns it to the CurrentUsage field.
func (o *GetLicense200Response) SetCurrentUsage(v GetLicense200ResponseCurrentUsage) {
	o.CurrentUsage = &v
}

func (o GetLicense200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLicense200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.License) {
		toSerialize["license"] = o.License
	}
	if !IsNil(o.InstalledLicenses) {
		toSerialize["installedLicenses"] = o.InstalledLicenses
	}
	if !IsNil(o.CurrentUsage) {
		toSerialize["currentUsage"] = o.CurrentUsage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetLicense200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

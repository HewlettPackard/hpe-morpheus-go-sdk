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

// checks if the GetCertificate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCertificate200Response{}

// GetCertificate200Response struct for GetCertificate200Response
type GetCertificate200Response struct {
	Certificate          *ListCertificates200ResponseCertificatesInner `json:"certificate,omitempty"`
	AdditionalProperties map[string]interface{}                        `json:",remain"`
}

type _GetCertificate200Response GetCertificate200Response

// NewGetCertificate200Response instantiates a new GetCertificate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCertificate200Response() *GetCertificate200Response {
	this := GetCertificate200Response{}
	return &this
}

// NewGetCertificate200ResponseWithDefaults instantiates a new GetCertificate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCertificate200ResponseWithDefaults() *GetCertificate200Response {
	this := GetCertificate200Response{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *GetCertificate200Response) GetCertificate() ListCertificates200ResponseCertificatesInner {
	if o == nil || IsNil(o.Certificate) {
		var ret ListCertificates200ResponseCertificatesInner
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCertificate200Response) GetCertificateOk() (*ListCertificates200ResponseCertificatesInner, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// IsSetCertificate returns a boolean if a field has been set.
func (o *GetCertificate200Response) IsSetCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given ListCertificates200ResponseCertificatesInner and assigns it to the Certificate field.
func (o *GetCertificate200Response) SetCertificate(v ListCertificates200ResponseCertificatesInner) {
	o.Certificate = &v
}

func (o GetCertificate200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCertificate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetCertificate200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

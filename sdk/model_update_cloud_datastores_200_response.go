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

// checks if the UpdateCloudDatastores200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCloudDatastores200Response{}

// UpdateCloudDatastores200Response struct for UpdateCloudDatastores200Response
type UpdateCloudDatastores200Response struct {
	Datastore            *UpdateCloudDatastores200ResponseDatastore `json:"datastore,omitempty"`
	AdditionalProperties map[string]interface{}                     `json:",remain"`
}

type _UpdateCloudDatastores200Response UpdateCloudDatastores200Response

// NewUpdateCloudDatastores200Response instantiates a new UpdateCloudDatastores200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCloudDatastores200Response() *UpdateCloudDatastores200Response {
	this := UpdateCloudDatastores200Response{}
	return &this
}

// NewUpdateCloudDatastores200ResponseWithDefaults instantiates a new UpdateCloudDatastores200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCloudDatastores200ResponseWithDefaults() *UpdateCloudDatastores200Response {
	this := UpdateCloudDatastores200Response{}
	return &this
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *UpdateCloudDatastores200Response) GetDatastore() UpdateCloudDatastores200ResponseDatastore {
	if o == nil || IsNil(o.Datastore) {
		var ret UpdateCloudDatastores200ResponseDatastore
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCloudDatastores200Response) GetDatastoreOk() (*UpdateCloudDatastores200ResponseDatastore, bool) {
	if o == nil || IsNil(o.Datastore) {
		return nil, false
	}
	return o.Datastore, true
}

// IsSetDatastore returns a boolean if a field has been set.
func (o *UpdateCloudDatastores200Response) IsSetDatastore() bool {
	if o != nil && !IsNil(o.Datastore) {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given UpdateCloudDatastores200ResponseDatastore and assigns it to the Datastore field.
func (o *UpdateCloudDatastores200Response) SetDatastore(v UpdateCloudDatastores200ResponseDatastore) {
	o.Datastore = &v
}

func (o UpdateCloudDatastores200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCloudDatastores200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Datastore) {
		toSerialize["datastore"] = o.Datastore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateCloudDatastores200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

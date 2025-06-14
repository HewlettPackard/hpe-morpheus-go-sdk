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

// checks if the GetArchiveFileDetail200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetArchiveFileDetail200Response{}

// GetArchiveFileDetail200Response struct for GetArchiveFileDetail200Response
type GetArchiveFileDetail200Response struct {
	ArchiveFile          *GetArchiveBucket200ResponseArchiveFilesInner `json:"archiveFile,omitempty"`
	IsOwner              *bool                                         `json:"isOwner,omitempty"`
	AdditionalProperties map[string]interface{}                        `json:",remain"`
}

type _GetArchiveFileDetail200Response GetArchiveFileDetail200Response

// NewGetArchiveFileDetail200Response instantiates a new GetArchiveFileDetail200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetArchiveFileDetail200Response() *GetArchiveFileDetail200Response {
	this := GetArchiveFileDetail200Response{}
	return &this
}

// NewGetArchiveFileDetail200ResponseWithDefaults instantiates a new GetArchiveFileDetail200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetArchiveFileDetail200ResponseWithDefaults() *GetArchiveFileDetail200Response {
	this := GetArchiveFileDetail200Response{}
	return &this
}

// GetArchiveFile returns the ArchiveFile field value if set, zero value otherwise.
func (o *GetArchiveFileDetail200Response) GetArchiveFile() GetArchiveBucket200ResponseArchiveFilesInner {
	if o == nil || IsNil(o.ArchiveFile) {
		var ret GetArchiveBucket200ResponseArchiveFilesInner
		return ret
	}
	return *o.ArchiveFile
}

// GetArchiveFileOk returns a tuple with the ArchiveFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetArchiveFileDetail200Response) GetArchiveFileOk() (*GetArchiveBucket200ResponseArchiveFilesInner, bool) {
	if o == nil || IsNil(o.ArchiveFile) {
		return nil, false
	}
	return o.ArchiveFile, true
}

// IsSetArchiveFile returns a boolean if a field has been set.
func (o *GetArchiveFileDetail200Response) IsSetArchiveFile() bool {
	if o != nil && !IsNil(o.ArchiveFile) {
		return true
	}

	return false
}

// SetArchiveFile gets a reference to the given GetArchiveBucket200ResponseArchiveFilesInner and assigns it to the ArchiveFile field.
func (o *GetArchiveFileDetail200Response) SetArchiveFile(v GetArchiveBucket200ResponseArchiveFilesInner) {
	o.ArchiveFile = &v
}

// GetIsOwner returns the IsOwner field value if set, zero value otherwise.
func (o *GetArchiveFileDetail200Response) GetIsOwner() bool {
	if o == nil || IsNil(o.IsOwner) {
		var ret bool
		return ret
	}
	return *o.IsOwner
}

// GetIsOwnerOk returns a tuple with the IsOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetArchiveFileDetail200Response) GetIsOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOwner) {
		return nil, false
	}
	return o.IsOwner, true
}

// IsSetIsOwner returns a boolean if a field has been set.
func (o *GetArchiveFileDetail200Response) IsSetIsOwner() bool {
	if o != nil && !IsNil(o.IsOwner) {
		return true
	}

	return false
}

// SetIsOwner gets a reference to the given bool and assigns it to the IsOwner field.
func (o *GetArchiveFileDetail200Response) SetIsOwner(v bool) {
	o.IsOwner = &v
}

func (o GetArchiveFileDetail200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetArchiveFileDetail200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArchiveFile) {
		toSerialize["archiveFile"] = o.ArchiveFile
	}
	if !IsNil(o.IsOwner) {
		toSerialize["isOwner"] = o.IsOwner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetArchiveFileDetail200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

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

// checks if the GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket{}

// GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket struct for GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket
type GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket struct {
	Id                   *int64                 `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	IsPublic             *bool                  `json:"isPublic,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket

// NewGetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket instantiates a new GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket() *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket {
	this := GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket{}
	return &this
}

// NewGetArchiveBucket200ResponseArchiveFilesInnerArchiveBucketWithDefaults instantiates a new GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetArchiveBucket200ResponseArchiveFilesInnerArchiveBucketWithDefaults() *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket {
	this := GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) SetName(v string) {
	o.Name = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// IsSetIsPublic returns a boolean if a field has been set.
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) IsSetIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) SetIsPublic(v bool) {
	o.IsPublic = &v
}

func (o GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IsPublic) {
		toSerialize["isPublic"] = o.IsPublic
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

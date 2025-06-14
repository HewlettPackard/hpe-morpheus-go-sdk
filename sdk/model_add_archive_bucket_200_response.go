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

// checks if the AddArchiveBucket200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddArchiveBucket200Response{}

// AddArchiveBucket200Response struct for AddArchiveBucket200Response
type AddArchiveBucket200Response struct {
	ArchiveBucket        *ListArchiveBuckets200ResponseAllOfArchiveBucketsInner `json:"archiveBucket,omitempty"`
	Success              *bool                                                  `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}                                 `json:",remain"`
}

type _AddArchiveBucket200Response AddArchiveBucket200Response

// NewAddArchiveBucket200Response instantiates a new AddArchiveBucket200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddArchiveBucket200Response() *AddArchiveBucket200Response {
	this := AddArchiveBucket200Response{}
	return &this
}

// NewAddArchiveBucket200ResponseWithDefaults instantiates a new AddArchiveBucket200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddArchiveBucket200ResponseWithDefaults() *AddArchiveBucket200Response {
	this := AddArchiveBucket200Response{}
	return &this
}

// GetArchiveBucket returns the ArchiveBucket field value if set, zero value otherwise.
func (o *AddArchiveBucket200Response) GetArchiveBucket() ListArchiveBuckets200ResponseAllOfArchiveBucketsInner {
	if o == nil || IsNil(o.ArchiveBucket) {
		var ret ListArchiveBuckets200ResponseAllOfArchiveBucketsInner
		return ret
	}
	return *o.ArchiveBucket
}

// GetArchiveBucketOk returns a tuple with the ArchiveBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddArchiveBucket200Response) GetArchiveBucketOk() (*ListArchiveBuckets200ResponseAllOfArchiveBucketsInner, bool) {
	if o == nil || IsNil(o.ArchiveBucket) {
		return nil, false
	}
	return o.ArchiveBucket, true
}

// IsSetArchiveBucket returns a boolean if a field has been set.
func (o *AddArchiveBucket200Response) IsSetArchiveBucket() bool {
	if o != nil && !IsNil(o.ArchiveBucket) {
		return true
	}

	return false
}

// SetArchiveBucket gets a reference to the given ListArchiveBuckets200ResponseAllOfArchiveBucketsInner and assigns it to the ArchiveBucket field.
func (o *AddArchiveBucket200Response) SetArchiveBucket(v ListArchiveBuckets200ResponseAllOfArchiveBucketsInner) {
	o.ArchiveBucket = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AddArchiveBucket200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddArchiveBucket200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *AddArchiveBucket200Response) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AddArchiveBucket200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o AddArchiveBucket200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddArchiveBucket200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArchiveBucket) {
		toSerialize["archiveBucket"] = o.ArchiveBucket
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddArchiveBucket200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

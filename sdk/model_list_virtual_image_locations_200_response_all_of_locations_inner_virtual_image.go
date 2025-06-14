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

// checks if the ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage{}

// ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage struct for ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage
type ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage struct {
	Id                   *int64                 `json:"id,omitempty"`
	Code                 *string                `json:"code,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	ImageType            *string                `json:"imageType,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage

// NewListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage instantiates a new ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage() *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage {
	this := ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage{}
	return &this
}

// NewListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImageWithDefaults instantiates a new ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImageWithDefaults() *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage {
	this := ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) SetName(v string) {
	o.Name = &v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) GetImageType() string {
	if o == nil || IsNil(o.ImageType) {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) GetImageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ImageType) {
		return nil, false
	}
	return o.ImageType, true
}

// IsSetImageType returns a boolean if a field has been set.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) IsSetImageType() bool {
	if o != nil && !IsNil(o.ImageType) {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) SetImageType(v string) {
	o.ImageType = &v
}

func (o ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ImageType) {
		toSerialize["imageType"] = o.ImageType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListVirtualImageLocations200ResponseAllOfLocationsInnerVirtualImage) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

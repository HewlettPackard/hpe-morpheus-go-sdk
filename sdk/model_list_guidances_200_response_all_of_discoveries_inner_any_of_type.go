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

// checks if the ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType{}

// ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType struct for ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType
type ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType struct {
	Id                   *int64                 `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Code                 *string                `json:"code,omitempty"`
	Title                *string                `json:"title,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType

// NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType instantiates a new ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType() *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType {
	this := ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType{}
	return &this
}

// NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfTypeWithDefaults instantiates a new ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfTypeWithDefaults() *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType {
	this := ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) SetCode(v string) {
	o.Code = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// IsSetTitle returns a boolean if a field has been set.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) IsSetTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) SetTitle(v string) {
	o.Title = &v
}

func (o ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

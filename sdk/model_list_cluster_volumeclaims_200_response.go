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

// checks if the ListClusterVolumeclaims200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListClusterVolumeclaims200Response{}

// ListClusterVolumeclaims200Response struct for ListClusterVolumeclaims200Response
type ListClusterVolumeclaims200Response struct {
	Volumeclaims         []ListClusterVolumes200ResponseAllOfVolumesInner `json:"volumeclaims,omitempty"`
	Meta                 *ListActivity200ResponseAllOfMeta                `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}                           `json:",remain"`
}

type _ListClusterVolumeclaims200Response ListClusterVolumeclaims200Response

// NewListClusterVolumeclaims200Response instantiates a new ListClusterVolumeclaims200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListClusterVolumeclaims200Response() *ListClusterVolumeclaims200Response {
	this := ListClusterVolumeclaims200Response{}
	return &this
}

// NewListClusterVolumeclaims200ResponseWithDefaults instantiates a new ListClusterVolumeclaims200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListClusterVolumeclaims200ResponseWithDefaults() *ListClusterVolumeclaims200Response {
	this := ListClusterVolumeclaims200Response{}
	return &this
}

// GetVolumeclaims returns the Volumeclaims field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200Response) GetVolumeclaims() []ListClusterVolumes200ResponseAllOfVolumesInner {
	if o == nil || IsNil(o.Volumeclaims) {
		var ret []ListClusterVolumes200ResponseAllOfVolumesInner
		return ret
	}
	return o.Volumeclaims
}

// GetVolumeclaimsOk returns a tuple with the Volumeclaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200Response) GetVolumeclaimsOk() ([]ListClusterVolumes200ResponseAllOfVolumesInner, bool) {
	if o == nil || IsNil(o.Volumeclaims) {
		return nil, false
	}
	return o.Volumeclaims, true
}

// IsSetVolumeclaims returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200Response) IsSetVolumeclaims() bool {
	if o != nil && !IsNil(o.Volumeclaims) {
		return true
	}

	return false
}

// SetVolumeclaims gets a reference to the given []ListClusterVolumes200ResponseAllOfVolumesInner and assigns it to the Volumeclaims field.
func (o *ListClusterVolumeclaims200Response) SetVolumeclaims(v []ListClusterVolumes200ResponseAllOfVolumesInner) {
	o.Volumeclaims = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200Response) GetMeta() ListActivity200ResponseAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ListActivity200ResponseAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// IsSetMeta returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200Response) IsSetMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ListActivity200ResponseAllOfMeta and assigns it to the Meta field.
func (o *ListClusterVolumeclaims200Response) SetMeta(v ListActivity200ResponseAllOfMeta) {
	o.Meta = &v
}

func (o ListClusterVolumeclaims200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListClusterVolumeclaims200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Volumeclaims) {
		toSerialize["volumeclaims"] = o.Volumeclaims
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListClusterVolumeclaims200Response) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

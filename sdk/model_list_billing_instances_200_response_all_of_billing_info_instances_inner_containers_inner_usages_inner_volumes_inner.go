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

// checks if the ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner{}

// ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner struct for ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner
type ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner struct {
	Size                 *int64                                                                                                        `json:"size,omitempty"`
	TypeCode             *string                                                                                                       `json:"typeCode,omitempty"`
	Datastore            *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore `json:"datastore,omitempty"`
	AdditionalProperties map[string]interface{}                                                                                        `json:",remain"`
}

type _ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner

// NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner instantiates a new ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner() *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner {
	this := ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner{}
	return &this
}

// NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInnerWithDefaults instantiates a new ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInnerWithDefaults() *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner {
	this := ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// IsSetSize returns a boolean if a field has been set.
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) IsSetSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) SetSize(v int64) {
	o.Size = &v
}

// GetTypeCode returns the TypeCode field value if set, zero value otherwise.
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) GetTypeCode() string {
	if o == nil || IsNil(o.TypeCode) {
		var ret string
		return ret
	}
	return *o.TypeCode
}

// GetTypeCodeOk returns a tuple with the TypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) GetTypeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TypeCode) {
		return nil, false
	}
	return o.TypeCode, true
}

// IsSetTypeCode returns a boolean if a field has been set.
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) IsSetTypeCode() bool {
	if o != nil && !IsNil(o.TypeCode) {
		return true
	}

	return false
}

// SetTypeCode gets a reference to the given string and assigns it to the TypeCode field.
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) SetTypeCode(v string) {
	o.TypeCode = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) GetDatastore() ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore {
	if o == nil || IsNil(o.Datastore) {
		var ret ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) GetDatastoreOk() (*ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore, bool) {
	if o == nil || IsNil(o.Datastore) {
		return nil, false
	}
	return o.Datastore, true
}

// IsSetDatastore returns a boolean if a field has been set.
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) IsSetDatastore() bool {
	if o != nil && !IsNil(o.Datastore) {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore and assigns it to the Datastore field.
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) SetDatastore(v ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) {
	o.Datastore = &v
}

func (o ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.TypeCode) {
		toSerialize["typeCode"] = o.TypeCode
	}
	if !IsNil(o.Datastore) {
		toSerialize["datastore"] = o.Datastore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListBillingInstances200ResponseAllOfBillingInfoInstancesInnerContainersInnerUsagesInnerVolumesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

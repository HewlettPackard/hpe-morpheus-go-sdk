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

// checks if the ListBackupResults200ResponseAllOfResultsInnerCreatedBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBackupResults200ResponseAllOfResultsInnerCreatedBy{}

// ListBackupResults200ResponseAllOfResultsInnerCreatedBy struct for ListBackupResults200ResponseAllOfResultsInnerCreatedBy
type ListBackupResults200ResponseAllOfResultsInnerCreatedBy struct {
	// User ID
	Id *int64 `json:"id,omitempty"`
	// Username
	Username             *string                `json:"username,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListBackupResults200ResponseAllOfResultsInnerCreatedBy ListBackupResults200ResponseAllOfResultsInnerCreatedBy

// NewListBackupResults200ResponseAllOfResultsInnerCreatedBy instantiates a new ListBackupResults200ResponseAllOfResultsInnerCreatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBackupResults200ResponseAllOfResultsInnerCreatedBy() *ListBackupResults200ResponseAllOfResultsInnerCreatedBy {
	this := ListBackupResults200ResponseAllOfResultsInnerCreatedBy{}
	return &this
}

// NewListBackupResults200ResponseAllOfResultsInnerCreatedByWithDefaults instantiates a new ListBackupResults200ResponseAllOfResultsInnerCreatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBackupResults200ResponseAllOfResultsInnerCreatedByWithDefaults() *ListBackupResults200ResponseAllOfResultsInnerCreatedBy {
	this := ListBackupResults200ResponseAllOfResultsInnerCreatedBy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListBackupResults200ResponseAllOfResultsInnerCreatedBy) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBackupResults200ResponseAllOfResultsInnerCreatedBy) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListBackupResults200ResponseAllOfResultsInnerCreatedBy) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListBackupResults200ResponseAllOfResultsInnerCreatedBy) SetId(v int64) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ListBackupResults200ResponseAllOfResultsInnerCreatedBy) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBackupResults200ResponseAllOfResultsInnerCreatedBy) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// IsSetUsername returns a boolean if a field has been set.
func (o *ListBackupResults200ResponseAllOfResultsInnerCreatedBy) IsSetUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ListBackupResults200ResponseAllOfResultsInnerCreatedBy) SetUsername(v string) {
	o.Username = &v
}

func (o ListBackupResults200ResponseAllOfResultsInnerCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListBackupResults200ResponseAllOfResultsInnerCreatedBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListBackupResults200ResponseAllOfResultsInnerCreatedBy) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

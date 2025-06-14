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

// checks if the ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner{}

// ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner struct for ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner
type ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner struct {
	SourceRoleName       *string                                                                     `json:"sourceRoleName,omitempty"`
	SourceRoleFqn        *string                                                                     `json:"sourceRoleFqn,omitempty"`
	MappedRole           *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole `json:"mappedRole,omitempty"`
	AdditionalProperties map[string]interface{}                                                      `json:",remain"`
}

type _ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner

// NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner instantiates a new ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner() *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner {
	this := ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner{}
	return &this
}

// NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInnerWithDefaults instantiates a new ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInnerWithDefaults() *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner {
	this := ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner{}
	return &this
}

// GetSourceRoleName returns the SourceRoleName field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) GetSourceRoleName() string {
	if o == nil || IsNil(o.SourceRoleName) {
		var ret string
		return ret
	}
	return *o.SourceRoleName
}

// GetSourceRoleNameOk returns a tuple with the SourceRoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) GetSourceRoleNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceRoleName) {
		return nil, false
	}
	return o.SourceRoleName, true
}

// IsSetSourceRoleName returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) IsSetSourceRoleName() bool {
	if o != nil && !IsNil(o.SourceRoleName) {
		return true
	}

	return false
}

// SetSourceRoleName gets a reference to the given string and assigns it to the SourceRoleName field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) SetSourceRoleName(v string) {
	o.SourceRoleName = &v
}

// GetSourceRoleFqn returns the SourceRoleFqn field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) GetSourceRoleFqn() string {
	if o == nil || IsNil(o.SourceRoleFqn) {
		var ret string
		return ret
	}
	return *o.SourceRoleFqn
}

// GetSourceRoleFqnOk returns a tuple with the SourceRoleFqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) GetSourceRoleFqnOk() (*string, bool) {
	if o == nil || IsNil(o.SourceRoleFqn) {
		return nil, false
	}
	return o.SourceRoleFqn, true
}

// IsSetSourceRoleFqn returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) IsSetSourceRoleFqn() bool {
	if o != nil && !IsNil(o.SourceRoleFqn) {
		return true
	}

	return false
}

// SetSourceRoleFqn gets a reference to the given string and assigns it to the SourceRoleFqn field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) SetSourceRoleFqn(v string) {
	o.SourceRoleFqn = &v
}

// GetMappedRole returns the MappedRole field value if set, zero value otherwise.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) GetMappedRole() ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole {
	if o == nil || IsNil(o.MappedRole) {
		var ret ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole
		return ret
	}
	return *o.MappedRole
}

// GetMappedRoleOk returns a tuple with the MappedRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) GetMappedRoleOk() (*ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole, bool) {
	if o == nil || IsNil(o.MappedRole) {
		return nil, false
	}
	return o.MappedRole, true
}

// IsSetMappedRole returns a boolean if a field has been set.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) IsSetMappedRole() bool {
	if o != nil && !IsNil(o.MappedRole) {
		return true
	}

	return false
}

// SetMappedRole gets a reference to the given ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole and assigns it to the MappedRole field.
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) SetMappedRole(v ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfDefaultAccountRole) {
	o.MappedRole = &v
}

func (o ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceRoleName) {
		toSerialize["sourceRoleName"] = o.SourceRoleName
	}
	if !IsNil(o.SourceRoleFqn) {
		toSerialize["sourceRoleFqn"] = o.SourceRoleFqn
	}
	if !IsNil(o.MappedRole) {
		toSerialize["mappedRole"] = o.MappedRole
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOfRoleMappingsInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

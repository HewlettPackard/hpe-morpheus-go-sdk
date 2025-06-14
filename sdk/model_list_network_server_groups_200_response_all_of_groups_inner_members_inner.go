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

// checks if the ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner{}

// ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner struct for ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner
type ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner struct {
	Id                   *int64                   `json:"id,omitempty"`
	Category             *string                  `json:"category,omitempty"`
	Type                 *string                  `json:"type,omitempty"`
	MemberName           *string                  `json:"memberName,omitempty"`
	MemberType           *string                  `json:"memberType,omitempty"`
	MemberValue          *string                  `json:"memberValue,omitempty"`
	MemberExpression     *string                  `json:"memberExpression,omitempty"`
	DisplayOrder         *int64                   `json:"displayOrder,omitempty"`
	InternalId           *string                  `json:"internalId,omitempty"`
	ExternalId           *string                  `json:"externalId,omitempty"`
	Members              []map[string]interface{} `json:"members,omitempty"`
	AdditionalProperties map[string]interface{}   `json:",remain"`
}

type _ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner

// NewListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner instantiates a new ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner() *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner {
	this := ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner{}
	return &this
}

// NewListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInnerWithDefaults instantiates a new ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInnerWithDefaults() *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner {
	this := ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) SetId(v int64) {
	o.Id = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// IsSetCategory returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) IsSetCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) SetCategory(v string) {
	o.Category = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) SetType(v string) {
	o.Type = &v
}

// GetMemberName returns the MemberName field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetMemberName() string {
	if o == nil || IsNil(o.MemberName) {
		var ret string
		return ret
	}
	return *o.MemberName
}

// GetMemberNameOk returns a tuple with the MemberName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetMemberNameOk() (*string, bool) {
	if o == nil || IsNil(o.MemberName) {
		return nil, false
	}
	return o.MemberName, true
}

// IsSetMemberName returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) IsSetMemberName() bool {
	if o != nil && !IsNil(o.MemberName) {
		return true
	}

	return false
}

// SetMemberName gets a reference to the given string and assigns it to the MemberName field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) SetMemberName(v string) {
	o.MemberName = &v
}

// GetMemberType returns the MemberType field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetMemberType() string {
	if o == nil || IsNil(o.MemberType) {
		var ret string
		return ret
	}
	return *o.MemberType
}

// GetMemberTypeOk returns a tuple with the MemberType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetMemberTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MemberType) {
		return nil, false
	}
	return o.MemberType, true
}

// IsSetMemberType returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) IsSetMemberType() bool {
	if o != nil && !IsNil(o.MemberType) {
		return true
	}

	return false
}

// SetMemberType gets a reference to the given string and assigns it to the MemberType field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) SetMemberType(v string) {
	o.MemberType = &v
}

// GetMemberValue returns the MemberValue field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetMemberValue() string {
	if o == nil || IsNil(o.MemberValue) {
		var ret string
		return ret
	}
	return *o.MemberValue
}

// GetMemberValueOk returns a tuple with the MemberValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetMemberValueOk() (*string, bool) {
	if o == nil || IsNil(o.MemberValue) {
		return nil, false
	}
	return o.MemberValue, true
}

// IsSetMemberValue returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) IsSetMemberValue() bool {
	if o != nil && !IsNil(o.MemberValue) {
		return true
	}

	return false
}

// SetMemberValue gets a reference to the given string and assigns it to the MemberValue field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) SetMemberValue(v string) {
	o.MemberValue = &v
}

// GetMemberExpression returns the MemberExpression field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetMemberExpression() string {
	if o == nil || IsNil(o.MemberExpression) {
		var ret string
		return ret
	}
	return *o.MemberExpression
}

// GetMemberExpressionOk returns a tuple with the MemberExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetMemberExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.MemberExpression) {
		return nil, false
	}
	return o.MemberExpression, true
}

// IsSetMemberExpression returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) IsSetMemberExpression() bool {
	if o != nil && !IsNil(o.MemberExpression) {
		return true
	}

	return false
}

// SetMemberExpression gets a reference to the given string and assigns it to the MemberExpression field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) SetMemberExpression(v string) {
	o.MemberExpression = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetDisplayOrder() int64 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret int64
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetDisplayOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// IsSetDisplayOrder returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) IsSetDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int64 and assigns it to the DisplayOrder field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) SetDisplayOrder(v int64) {
	o.DisplayOrder = &v
}

// GetInternalId returns the InternalId field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetInternalId() string {
	if o == nil || IsNil(o.InternalId) {
		var ret string
		return ret
	}
	return *o.InternalId
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetInternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.InternalId) {
		return nil, false
	}
	return o.InternalId, true
}

// IsSetInternalId returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) IsSetInternalId() bool {
	if o != nil && !IsNil(o.InternalId) {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given string and assigns it to the InternalId field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) SetInternalId(v string) {
	o.InternalId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// IsSetExternalId returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) IsSetExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetMembers() []map[string]interface{} {
	if o == nil || IsNil(o.Members) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) GetMembersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// IsSetMembers returns a boolean if a field has been set.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) IsSetMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []map[string]interface{} and assigns it to the Members field.
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) SetMembers(v []map[string]interface{}) {
	o.Members = v
}

func (o ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.MemberName) {
		toSerialize["memberName"] = o.MemberName
	}
	if !IsNil(o.MemberType) {
		toSerialize["memberType"] = o.MemberType
	}
	if !IsNil(o.MemberValue) {
		toSerialize["memberValue"] = o.MemberValue
	}
	if !IsNil(o.MemberExpression) {
		toSerialize["memberExpression"] = o.MemberExpression
	}
	if !IsNil(o.DisplayOrder) {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if !IsNil(o.InternalId) {
		toSerialize["internalId"] = o.InternalId
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

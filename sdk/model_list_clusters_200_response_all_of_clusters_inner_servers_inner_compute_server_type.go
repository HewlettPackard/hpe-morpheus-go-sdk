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

// checks if the ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType{}

// ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType struct for ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType
type ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType struct {
	Id                   *int64                 `json:"id,omitempty"`
	Code                 *string                `json:"code,omitempty"`
	NodeType             *string                `json:"nodeType,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType

// NewListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType instantiates a new ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType() *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType {
	this := ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType{}
	return &this
}

// NewListClusters200ResponseAllOfClustersInnerServersInnerComputeServerTypeWithDefaults instantiates a new ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListClusters200ResponseAllOfClustersInnerServersInnerComputeServerTypeWithDefaults() *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType {
	this := ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) SetCode(v string) {
	o.Code = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) GetNodeType() string {
	if o == nil || IsNil(o.NodeType) {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) GetNodeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NodeType) {
		return nil, false
	}
	return o.NodeType, true
}

// IsSetNodeType returns a boolean if a field has been set.
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) IsSetNodeType() bool {
	if o != nil && !IsNil(o.NodeType) {
		return true
	}

	return false
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) SetNodeType(v string) {
	o.NodeType = &v
}

func (o ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.NodeType) {
		toSerialize["nodeType"] = o.NodeType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListClusters200ResponseAllOfClustersInnerServersInnerComputeServerType) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

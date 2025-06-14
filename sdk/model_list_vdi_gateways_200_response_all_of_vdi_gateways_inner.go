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
	"time"
)

// checks if the ListVDIGateways200ResponseAllOfVdiGatewaysInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListVDIGateways200ResponseAllOfVdiGatewaysInner{}

// ListVDIGateways200ResponseAllOfVdiGatewaysInner struct for ListVDIGateways200ResponseAllOfVdiGatewaysInner
type ListVDIGateways200ResponseAllOfVdiGatewaysInner struct {
	Id                   *int64                 `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Description          *string                `json:"description,omitempty"`
	GatewayUrl           *string                `json:"gatewayUrl,omitempty"`
	ApiKey               *string                `json:"apiKey,omitempty"`
	DateCreated          *time.Time             `json:"dateCreated,omitempty"`
	LastUpdated          *time.Time             `json:"lastUpdated,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ListVDIGateways200ResponseAllOfVdiGatewaysInner ListVDIGateways200ResponseAllOfVdiGatewaysInner

// NewListVDIGateways200ResponseAllOfVdiGatewaysInner instantiates a new ListVDIGateways200ResponseAllOfVdiGatewaysInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVDIGateways200ResponseAllOfVdiGatewaysInner() *ListVDIGateways200ResponseAllOfVdiGatewaysInner {
	this := ListVDIGateways200ResponseAllOfVdiGatewaysInner{}
	return &this
}

// NewListVDIGateways200ResponseAllOfVdiGatewaysInnerWithDefaults instantiates a new ListVDIGateways200ResponseAllOfVdiGatewaysInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVDIGateways200ResponseAllOfVdiGatewaysInnerWithDefaults() *ListVDIGateways200ResponseAllOfVdiGatewaysInner {
	this := ListVDIGateways200ResponseAllOfVdiGatewaysInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) SetDescription(v string) {
	o.Description = &v
}

// GetGatewayUrl returns the GatewayUrl field value if set, zero value otherwise.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetGatewayUrl() string {
	if o == nil || IsNil(o.GatewayUrl) {
		var ret string
		return ret
	}
	return *o.GatewayUrl
}

// GetGatewayUrlOk returns a tuple with the GatewayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetGatewayUrlOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayUrl) {
		return nil, false
	}
	return o.GatewayUrl, true
}

// IsSetGatewayUrl returns a boolean if a field has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) IsSetGatewayUrl() bool {
	if o != nil && !IsNil(o.GatewayUrl) {
		return true
	}

	return false
}

// SetGatewayUrl gets a reference to the given string and assigns it to the GatewayUrl field.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) SetGatewayUrl(v string) {
	o.GatewayUrl = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// IsSetApiKey returns a boolean if a field has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) IsSetApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o ListVDIGateways200ResponseAllOfVdiGatewaysInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListVDIGateways200ResponseAllOfVdiGatewaysInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GatewayUrl) {
		toSerialize["gatewayUrl"] = o.GatewayUrl
	}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListVDIGateways200ResponseAllOfVdiGatewaysInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

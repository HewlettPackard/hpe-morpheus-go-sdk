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

// checks if the ListPreseedScripts200ResponseAllOfPreseedScriptsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPreseedScripts200ResponseAllOfPreseedScriptsInner{}

// ListPreseedScripts200ResponseAllOfPreseedScriptsInner struct for ListPreseedScripts200ResponseAllOfPreseedScriptsInner
type ListPreseedScripts200ResponseAllOfPreseedScriptsInner struct {
	Id                   *int64                                                                  `json:"id,omitempty"`
	Account              *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"account,omitempty"`
	FileName             *string                                                                 `json:"fileName,omitempty"`
	Description          *string                                                                 `json:"description,omitempty"`
	Content              *string                                                                 `json:"content,omitempty"`
	CreatedBy            *GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy                  `json:"createdBy,omitempty"`
	AdditionalProperties map[string]interface{}                                                  `json:",remain"`
}

type _ListPreseedScripts200ResponseAllOfPreseedScriptsInner ListPreseedScripts200ResponseAllOfPreseedScriptsInner

// NewListPreseedScripts200ResponseAllOfPreseedScriptsInner instantiates a new ListPreseedScripts200ResponseAllOfPreseedScriptsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPreseedScripts200ResponseAllOfPreseedScriptsInner() *ListPreseedScripts200ResponseAllOfPreseedScriptsInner {
	this := ListPreseedScripts200ResponseAllOfPreseedScriptsInner{}
	return &this
}

// NewListPreseedScripts200ResponseAllOfPreseedScriptsInnerWithDefaults instantiates a new ListPreseedScripts200ResponseAllOfPreseedScriptsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPreseedScripts200ResponseAllOfPreseedScriptsInnerWithDefaults() *ListPreseedScripts200ResponseAllOfPreseedScriptsInner {
	this := ListPreseedScripts200ResponseAllOfPreseedScriptsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) SetId(v int64) {
	o.Id = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Account) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// IsSetAccount returns a boolean if a field has been set.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) IsSetAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Account field.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Account = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// IsSetFileName returns a boolean if a field has been set.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) IsSetFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) SetFileName(v string) {
	o.FileName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) SetDescription(v string) {
	o.Description = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// IsSetContent returns a boolean if a field has been set.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) IsSetContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) SetContent(v string) {
	o.Content = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) GetCreatedBy() GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy {
	if o == nil || IsNil(o.CreatedBy) {
		var ret GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) GetCreatedByOk() (*GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// IsSetCreatedBy returns a boolean if a field has been set.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) IsSetCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy and assigns it to the CreatedBy field.
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) SetCreatedBy(v GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy) {
	o.CreatedBy = &v
}

func (o ListPreseedScripts200ResponseAllOfPreseedScriptsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPreseedScripts200ResponseAllOfPreseedScriptsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListPreseedScripts200ResponseAllOfPreseedScriptsInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

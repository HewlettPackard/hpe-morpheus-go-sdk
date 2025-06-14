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

// checks if the ListPlugins200ResponseAllOfPluginsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPlugins200ResponseAllOfPluginsInner{}

// ListPlugins200ResponseAllOfPluginsInner struct for ListPlugins200ResponseAllOfPluginsInner
type ListPlugins200ResponseAllOfPluginsInner struct {
	Id                    *int64                                                                      `json:"id,omitempty"`
	Name                  *string                                                                     `json:"name,omitempty"`
	Code                  *string                                                                     `json:"code,omitempty"`
	Description           *string                                                                     `json:"description,omitempty"`
	Version               *string                                                                     `json:"version,omitempty"`
	Enabled               *bool                                                                       `json:"enabled,omitempty"`
	Author                *string                                                                     `json:"author,omitempty"`
	WebsiteUrl            *string                                                                     `json:"websiteUrl,omitempty"`
	SourceCodeLocationUrl *string                                                                     `json:"sourceCodeLocationUrl,omitempty"`
	IssueTrackerUrl       *string                                                                     `json:"issueTrackerUrl,omitempty"`
	Valid                 *bool                                                                       `json:"valid,omitempty"`
	HasValidUpdate        *bool                                                                       `json:"hasValidUpdate,omitempty"`
	Status                *string                                                                     `json:"status,omitempty"`
	StatusMessage         *string                                                                     `json:"statusMessage,omitempty"`
	Providers             []GetAppState200ResponseAllOfSpecsInnerTemplate                             `json:"providers,omitempty"`
	Config                map[string]interface{}                                                      `json:"config,omitempty"`
	OptionTypes           []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner `json:"optionTypes,omitempty"`
	DateCreated           *time.Time                                                                  `json:"dateCreated,omitempty"`
	LastUpdated           *time.Time                                                                  `json:"lastUpdated,omitempty"`
	AdditionalProperties  map[string]interface{}                                                      `json:",remain"`
}

type _ListPlugins200ResponseAllOfPluginsInner ListPlugins200ResponseAllOfPluginsInner

// NewListPlugins200ResponseAllOfPluginsInner instantiates a new ListPlugins200ResponseAllOfPluginsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPlugins200ResponseAllOfPluginsInner() *ListPlugins200ResponseAllOfPluginsInner {
	this := ListPlugins200ResponseAllOfPluginsInner{}
	return &this
}

// NewListPlugins200ResponseAllOfPluginsInnerWithDefaults instantiates a new ListPlugins200ResponseAllOfPluginsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPlugins200ResponseAllOfPluginsInnerWithDefaults() *ListPlugins200ResponseAllOfPluginsInner {
	this := ListPlugins200ResponseAllOfPluginsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// IsSetVersion returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetVersion(v string) {
	o.Version = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// IsSetAuthor returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetAuthor(v string) {
	o.Author = &v
}

// GetWebsiteUrl returns the WebsiteUrl field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetWebsiteUrl() string {
	if o == nil || IsNil(o.WebsiteUrl) {
		var ret string
		return ret
	}
	return *o.WebsiteUrl
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetWebsiteUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebsiteUrl) {
		return nil, false
	}
	return o.WebsiteUrl, true
}

// IsSetWebsiteUrl returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetWebsiteUrl() bool {
	if o != nil && !IsNil(o.WebsiteUrl) {
		return true
	}

	return false
}

// SetWebsiteUrl gets a reference to the given string and assigns it to the WebsiteUrl field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetWebsiteUrl(v string) {
	o.WebsiteUrl = &v
}

// GetSourceCodeLocationUrl returns the SourceCodeLocationUrl field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetSourceCodeLocationUrl() string {
	if o == nil || IsNil(o.SourceCodeLocationUrl) {
		var ret string
		return ret
	}
	return *o.SourceCodeLocationUrl
}

// GetSourceCodeLocationUrlOk returns a tuple with the SourceCodeLocationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetSourceCodeLocationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCodeLocationUrl) {
		return nil, false
	}
	return o.SourceCodeLocationUrl, true
}

// IsSetSourceCodeLocationUrl returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetSourceCodeLocationUrl() bool {
	if o != nil && !IsNil(o.SourceCodeLocationUrl) {
		return true
	}

	return false
}

// SetSourceCodeLocationUrl gets a reference to the given string and assigns it to the SourceCodeLocationUrl field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetSourceCodeLocationUrl(v string) {
	o.SourceCodeLocationUrl = &v
}

// GetIssueTrackerUrl returns the IssueTrackerUrl field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetIssueTrackerUrl() string {
	if o == nil || IsNil(o.IssueTrackerUrl) {
		var ret string
		return ret
	}
	return *o.IssueTrackerUrl
}

// GetIssueTrackerUrlOk returns a tuple with the IssueTrackerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetIssueTrackerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IssueTrackerUrl) {
		return nil, false
	}
	return o.IssueTrackerUrl, true
}

// IsSetIssueTrackerUrl returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetIssueTrackerUrl() bool {
	if o != nil && !IsNil(o.IssueTrackerUrl) {
		return true
	}

	return false
}

// SetIssueTrackerUrl gets a reference to the given string and assigns it to the IssueTrackerUrl field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetIssueTrackerUrl(v string) {
	o.IssueTrackerUrl = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// IsSetValid returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetValid(v bool) {
	o.Valid = &v
}

// GetHasValidUpdate returns the HasValidUpdate field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetHasValidUpdate() bool {
	if o == nil || IsNil(o.HasValidUpdate) {
		var ret bool
		return ret
	}
	return *o.HasValidUpdate
}

// GetHasValidUpdateOk returns a tuple with the HasValidUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetHasValidUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.HasValidUpdate) {
		return nil, false
	}
	return o.HasValidUpdate, true
}

// IsSetHasValidUpdate returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetHasValidUpdate() bool {
	if o != nil && !IsNil(o.HasValidUpdate) {
		return true
	}

	return false
}

// SetHasValidUpdate gets a reference to the given bool and assigns it to the HasValidUpdate field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetHasValidUpdate(v bool) {
	o.HasValidUpdate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// IsSetStatus returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// IsSetStatusMessage returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetProviders() []GetAppState200ResponseAllOfSpecsInnerTemplate {
	if o == nil || IsNil(o.Providers) {
		var ret []GetAppState200ResponseAllOfSpecsInnerTemplate
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetProvidersOk() ([]GetAppState200ResponseAllOfSpecsInnerTemplate, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// IsSetProviders returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []GetAppState200ResponseAllOfSpecsInnerTemplate and assigns it to the Providers field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetProviders(v []GetAppState200ResponseAllOfSpecsInnerTemplate) {
	o.Providers = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner {
	if o == nil || IsNil(o.OptionTypes) {
		var ret []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner
		return ret
	}
	return o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetOptionTypesOk() ([]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool) {
	if o == nil || IsNil(o.OptionTypes) {
		return nil, false
	}
	return o.OptionTypes, true
}

// IsSetOptionTypes returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetOptionTypes() bool {
	if o != nil && !IsNil(o.OptionTypes) {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner and assigns it to the OptionTypes field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner) {
	o.OptionTypes = v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *ListPlugins200ResponseAllOfPluginsInner) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ListPlugins200ResponseAllOfPluginsInner) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o ListPlugins200ResponseAllOfPluginsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPlugins200ResponseAllOfPluginsInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.WebsiteUrl) {
		toSerialize["websiteUrl"] = o.WebsiteUrl
	}
	if !IsNil(o.SourceCodeLocationUrl) {
		toSerialize["sourceCodeLocationUrl"] = o.SourceCodeLocationUrl
	}
	if !IsNil(o.IssueTrackerUrl) {
		toSerialize["issueTrackerUrl"] = o.IssueTrackerUrl
	}
	if !IsNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	if !IsNil(o.HasValidUpdate) {
		toSerialize["hasValidUpdate"] = o.HasValidUpdate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["statusMessage"] = o.StatusMessage
	}
	if !IsNil(o.Providers) {
		toSerialize["providers"] = o.Providers
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.OptionTypes) {
		toSerialize["optionTypes"] = o.OptionTypes
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
func (o *ListPlugins200ResponseAllOfPluginsInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

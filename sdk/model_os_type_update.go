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

// checks if the OsTypeUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsTypeUpdate{}

// OsTypeUpdate struct for OsTypeUpdate
type OsTypeUpdate struct {
	// The name of the osType.
	Name *string `json:"name,omitempty"`
	// The description of the osType.
	Description *string `json:"description,omitempty"`
	// The platform of the osType.
	Platform *string `json:"platform,omitempty"`
	// The category of the osType.
	Category *string `json:"category,omitempty"`
	// The vendor of the osType.
	Vendor *string `json:"vendor,omitempty"`
	// The osName of the osType.
	OsName *string `json:"osName,omitempty"`
	// The osVersion of the osType.
	OsVersion *string `json:"osVersion,omitempty"`
	// The osCodename of the osType.
	OsCodename *string `json:"osCodename,omitempty"`
	// The family of the osType.
	OsFamily *string `json:"osFamily,omitempty"`
	// The bitCount/architecture of the osType.
	BitCount *int64 `json:"bitCount,omitempty"`
	// The version of CloudInit being used.
	CloudInitVersion *string `json:"cloudInitVersion,omitempty"`
	// Whether the morpheus agent is installed.
	InstallAgent         *bool                  `json:"installAgent,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _OsTypeUpdate OsTypeUpdate

// NewOsTypeUpdate instantiates a new OsTypeUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsTypeUpdate() *OsTypeUpdate {
	this := OsTypeUpdate{}
	return &this
}

// NewOsTypeUpdateWithDefaults instantiates a new OsTypeUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsTypeUpdateWithDefaults() *OsTypeUpdate {
	this := OsTypeUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsTypeUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *OsTypeUpdate) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsTypeUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OsTypeUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *OsTypeUpdate) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OsTypeUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *OsTypeUpdate) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeUpdate) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// IsSetPlatform returns a boolean if a field has been set.
func (o *OsTypeUpdate) IsSetPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *OsTypeUpdate) SetPlatform(v string) {
	o.Platform = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *OsTypeUpdate) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeUpdate) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// IsSetCategory returns a boolean if a field has been set.
func (o *OsTypeUpdate) IsSetCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *OsTypeUpdate) SetCategory(v string) {
	o.Category = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *OsTypeUpdate) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeUpdate) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// IsSetVendor returns a boolean if a field has been set.
func (o *OsTypeUpdate) IsSetVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *OsTypeUpdate) SetVendor(v string) {
	o.Vendor = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *OsTypeUpdate) GetOsName() string {
	if o == nil || IsNil(o.OsName) {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeUpdate) GetOsNameOk() (*string, bool) {
	if o == nil || IsNil(o.OsName) {
		return nil, false
	}
	return o.OsName, true
}

// IsSetOsName returns a boolean if a field has been set.
func (o *OsTypeUpdate) IsSetOsName() bool {
	if o != nil && !IsNil(o.OsName) {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *OsTypeUpdate) SetOsName(v string) {
	o.OsName = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *OsTypeUpdate) GetOsVersion() string {
	if o == nil || IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeUpdate) GetOsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// IsSetOsVersion returns a boolean if a field has been set.
func (o *OsTypeUpdate) IsSetOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *OsTypeUpdate) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetOsCodename returns the OsCodename field value if set, zero value otherwise.
func (o *OsTypeUpdate) GetOsCodename() string {
	if o == nil || IsNil(o.OsCodename) {
		var ret string
		return ret
	}
	return *o.OsCodename
}

// GetOsCodenameOk returns a tuple with the OsCodename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeUpdate) GetOsCodenameOk() (*string, bool) {
	if o == nil || IsNil(o.OsCodename) {
		return nil, false
	}
	return o.OsCodename, true
}

// IsSetOsCodename returns a boolean if a field has been set.
func (o *OsTypeUpdate) IsSetOsCodename() bool {
	if o != nil && !IsNil(o.OsCodename) {
		return true
	}

	return false
}

// SetOsCodename gets a reference to the given string and assigns it to the OsCodename field.
func (o *OsTypeUpdate) SetOsCodename(v string) {
	o.OsCodename = &v
}

// GetOsFamily returns the OsFamily field value if set, zero value otherwise.
func (o *OsTypeUpdate) GetOsFamily() string {
	if o == nil || IsNil(o.OsFamily) {
		var ret string
		return ret
	}
	return *o.OsFamily
}

// GetOsFamilyOk returns a tuple with the OsFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeUpdate) GetOsFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.OsFamily) {
		return nil, false
	}
	return o.OsFamily, true
}

// IsSetOsFamily returns a boolean if a field has been set.
func (o *OsTypeUpdate) IsSetOsFamily() bool {
	if o != nil && !IsNil(o.OsFamily) {
		return true
	}

	return false
}

// SetOsFamily gets a reference to the given string and assigns it to the OsFamily field.
func (o *OsTypeUpdate) SetOsFamily(v string) {
	o.OsFamily = &v
}

// GetBitCount returns the BitCount field value if set, zero value otherwise.
func (o *OsTypeUpdate) GetBitCount() int64 {
	if o == nil || IsNil(o.BitCount) {
		var ret int64
		return ret
	}
	return *o.BitCount
}

// GetBitCountOk returns a tuple with the BitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeUpdate) GetBitCountOk() (*int64, bool) {
	if o == nil || IsNil(o.BitCount) {
		return nil, false
	}
	return o.BitCount, true
}

// IsSetBitCount returns a boolean if a field has been set.
func (o *OsTypeUpdate) IsSetBitCount() bool {
	if o != nil && !IsNil(o.BitCount) {
		return true
	}

	return false
}

// SetBitCount gets a reference to the given int64 and assigns it to the BitCount field.
func (o *OsTypeUpdate) SetBitCount(v int64) {
	o.BitCount = &v
}

// GetCloudInitVersion returns the CloudInitVersion field value if set, zero value otherwise.
func (o *OsTypeUpdate) GetCloudInitVersion() string {
	if o == nil || IsNil(o.CloudInitVersion) {
		var ret string
		return ret
	}
	return *o.CloudInitVersion
}

// GetCloudInitVersionOk returns a tuple with the CloudInitVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeUpdate) GetCloudInitVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CloudInitVersion) {
		return nil, false
	}
	return o.CloudInitVersion, true
}

// IsSetCloudInitVersion returns a boolean if a field has been set.
func (o *OsTypeUpdate) IsSetCloudInitVersion() bool {
	if o != nil && !IsNil(o.CloudInitVersion) {
		return true
	}

	return false
}

// SetCloudInitVersion gets a reference to the given string and assigns it to the CloudInitVersion field.
func (o *OsTypeUpdate) SetCloudInitVersion(v string) {
	o.CloudInitVersion = &v
}

// GetInstallAgent returns the InstallAgent field value if set, zero value otherwise.
func (o *OsTypeUpdate) GetInstallAgent() bool {
	if o == nil || IsNil(o.InstallAgent) {
		var ret bool
		return ret
	}
	return *o.InstallAgent
}

// GetInstallAgentOk returns a tuple with the InstallAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsTypeUpdate) GetInstallAgentOk() (*bool, bool) {
	if o == nil || IsNil(o.InstallAgent) {
		return nil, false
	}
	return o.InstallAgent, true
}

// IsSetInstallAgent returns a boolean if a field has been set.
func (o *OsTypeUpdate) IsSetInstallAgent() bool {
	if o != nil && !IsNil(o.InstallAgent) {
		return true
	}

	return false
}

// SetInstallAgent gets a reference to the given bool and assigns it to the InstallAgent field.
func (o *OsTypeUpdate) SetInstallAgent(v bool) {
	o.InstallAgent = &v
}

func (o OsTypeUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsTypeUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.OsName) {
		toSerialize["osName"] = o.OsName
	}
	if !IsNil(o.OsVersion) {
		toSerialize["osVersion"] = o.OsVersion
	}
	if !IsNil(o.OsCodename) {
		toSerialize["osCodename"] = o.OsCodename
	}
	if !IsNil(o.OsFamily) {
		toSerialize["osFamily"] = o.OsFamily
	}
	if !IsNil(o.BitCount) {
		toSerialize["bitCount"] = o.BitCount
	}
	if !IsNil(o.CloudInitVersion) {
		toSerialize["cloudInitVersion"] = o.CloudInitVersion
	}
	if !IsNil(o.InstallAgent) {
		toSerialize["installAgent"] = o.InstallAgent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *OsTypeUpdate) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

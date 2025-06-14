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

// checks if the ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config{}

// ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config struct for ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config
type ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config struct {
	Databags             []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3ConfigDatabagsInner `json:"databags,omitempty"`
	Endpoint             *string                                                                      `json:"endpoint,omitempty"`
	Org                  *string                                                                      `json:"org,omitempty"`
	ChefUser             *string                                                                      `json:"chefUser,omitempty"`
	UserKey              *string                                                                      `json:"userKey,omitempty"`
	OrgKey               *string                                                                      `json:"orgKey,omitempty"`
	Version              *string                                                                      `json:"version,omitempty"`
	ChefUseFqdn          *bool                                                                        `json:"chefUseFqdn,omitempty"`
	WindowsVersion       *string                                                                      `json:"windowsVersion,omitempty"`
	WindowsInstallUrl    *string                                                                      `json:"windowsInstallUrl,omitempty"`
	UserKeyHash          *string                                                                      `json:"userKeyHash,omitempty"`
	OrgKeyHash           *string                                                                      `json:"orgKeyHash,omitempty"`
	AdditionalProperties map[string]interface{}                                                       `json:",remain"`
}

type _ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config

// NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config {
	this := ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config{}
	return &this
}

// NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3ConfigWithDefaults instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3ConfigWithDefaults() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config {
	this := ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config{}
	return &this
}

// GetDatabags returns the Databags field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetDatabags() []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3ConfigDatabagsInner {
	if o == nil || IsNil(o.Databags) {
		var ret []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3ConfigDatabagsInner
		return ret
	}
	return o.Databags
}

// GetDatabagsOk returns a tuple with the Databags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetDatabagsOk() ([]ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3ConfigDatabagsInner, bool) {
	if o == nil || IsNil(o.Databags) {
		return nil, false
	}
	return o.Databags, true
}

// IsSetDatabags returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) IsSetDatabags() bool {
	if o != nil && !IsNil(o.Databags) {
		return true
	}

	return false
}

// SetDatabags gets a reference to the given []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3ConfigDatabagsInner and assigns it to the Databags field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) SetDatabags(v []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3ConfigDatabagsInner) {
	o.Databags = v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// IsSetEndpoint returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) IsSetEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetOrg() string {
	if o == nil || IsNil(o.Org) {
		var ret string
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetOrgOk() (*string, bool) {
	if o == nil || IsNil(o.Org) {
		return nil, false
	}
	return o.Org, true
}

// IsSetOrg returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) IsSetOrg() bool {
	if o != nil && !IsNil(o.Org) {
		return true
	}

	return false
}

// SetOrg gets a reference to the given string and assigns it to the Org field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) SetOrg(v string) {
	o.Org = &v
}

// GetChefUser returns the ChefUser field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetChefUser() string {
	if o == nil || IsNil(o.ChefUser) {
		var ret string
		return ret
	}
	return *o.ChefUser
}

// GetChefUserOk returns a tuple with the ChefUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetChefUserOk() (*string, bool) {
	if o == nil || IsNil(o.ChefUser) {
		return nil, false
	}
	return o.ChefUser, true
}

// IsSetChefUser returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) IsSetChefUser() bool {
	if o != nil && !IsNil(o.ChefUser) {
		return true
	}

	return false
}

// SetChefUser gets a reference to the given string and assigns it to the ChefUser field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) SetChefUser(v string) {
	o.ChefUser = &v
}

// GetUserKey returns the UserKey field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetUserKey() string {
	if o == nil || IsNil(o.UserKey) {
		var ret string
		return ret
	}
	return *o.UserKey
}

// GetUserKeyOk returns a tuple with the UserKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetUserKeyOk() (*string, bool) {
	if o == nil || IsNil(o.UserKey) {
		return nil, false
	}
	return o.UserKey, true
}

// IsSetUserKey returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) IsSetUserKey() bool {
	if o != nil && !IsNil(o.UserKey) {
		return true
	}

	return false
}

// SetUserKey gets a reference to the given string and assigns it to the UserKey field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) SetUserKey(v string) {
	o.UserKey = &v
}

// GetOrgKey returns the OrgKey field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetOrgKey() string {
	if o == nil || IsNil(o.OrgKey) {
		var ret string
		return ret
	}
	return *o.OrgKey
}

// GetOrgKeyOk returns a tuple with the OrgKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetOrgKeyOk() (*string, bool) {
	if o == nil || IsNil(o.OrgKey) {
		return nil, false
	}
	return o.OrgKey, true
}

// IsSetOrgKey returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) IsSetOrgKey() bool {
	if o != nil && !IsNil(o.OrgKey) {
		return true
	}

	return false
}

// SetOrgKey gets a reference to the given string and assigns it to the OrgKey field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) SetOrgKey(v string) {
	o.OrgKey = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// IsSetVersion returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) IsSetVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) SetVersion(v string) {
	o.Version = &v
}

// GetChefUseFqdn returns the ChefUseFqdn field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetChefUseFqdn() bool {
	if o == nil || IsNil(o.ChefUseFqdn) {
		var ret bool
		return ret
	}
	return *o.ChefUseFqdn
}

// GetChefUseFqdnOk returns a tuple with the ChefUseFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetChefUseFqdnOk() (*bool, bool) {
	if o == nil || IsNil(o.ChefUseFqdn) {
		return nil, false
	}
	return o.ChefUseFqdn, true
}

// IsSetChefUseFqdn returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) IsSetChefUseFqdn() bool {
	if o != nil && !IsNil(o.ChefUseFqdn) {
		return true
	}

	return false
}

// SetChefUseFqdn gets a reference to the given bool and assigns it to the ChefUseFqdn field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) SetChefUseFqdn(v bool) {
	o.ChefUseFqdn = &v
}

// GetWindowsVersion returns the WindowsVersion field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetWindowsVersion() string {
	if o == nil || IsNil(o.WindowsVersion) {
		var ret string
		return ret
	}
	return *o.WindowsVersion
}

// GetWindowsVersionOk returns a tuple with the WindowsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetWindowsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsVersion) {
		return nil, false
	}
	return o.WindowsVersion, true
}

// IsSetWindowsVersion returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) IsSetWindowsVersion() bool {
	if o != nil && !IsNil(o.WindowsVersion) {
		return true
	}

	return false
}

// SetWindowsVersion gets a reference to the given string and assigns it to the WindowsVersion field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) SetWindowsVersion(v string) {
	o.WindowsVersion = &v
}

// GetWindowsInstallUrl returns the WindowsInstallUrl field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetWindowsInstallUrl() string {
	if o == nil || IsNil(o.WindowsInstallUrl) {
		var ret string
		return ret
	}
	return *o.WindowsInstallUrl
}

// GetWindowsInstallUrlOk returns a tuple with the WindowsInstallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetWindowsInstallUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsInstallUrl) {
		return nil, false
	}
	return o.WindowsInstallUrl, true
}

// IsSetWindowsInstallUrl returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) IsSetWindowsInstallUrl() bool {
	if o != nil && !IsNil(o.WindowsInstallUrl) {
		return true
	}

	return false
}

// SetWindowsInstallUrl gets a reference to the given string and assigns it to the WindowsInstallUrl field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) SetWindowsInstallUrl(v string) {
	o.WindowsInstallUrl = &v
}

// GetUserKeyHash returns the UserKeyHash field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetUserKeyHash() string {
	if o == nil || IsNil(o.UserKeyHash) {
		var ret string
		return ret
	}
	return *o.UserKeyHash
}

// GetUserKeyHashOk returns a tuple with the UserKeyHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetUserKeyHashOk() (*string, bool) {
	if o == nil || IsNil(o.UserKeyHash) {
		return nil, false
	}
	return o.UserKeyHash, true
}

// IsSetUserKeyHash returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) IsSetUserKeyHash() bool {
	if o != nil && !IsNil(o.UserKeyHash) {
		return true
	}

	return false
}

// SetUserKeyHash gets a reference to the given string and assigns it to the UserKeyHash field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) SetUserKeyHash(v string) {
	o.UserKeyHash = &v
}

// GetOrgKeyHash returns the OrgKeyHash field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetOrgKeyHash() string {
	if o == nil || IsNil(o.OrgKeyHash) {
		var ret string
		return ret
	}
	return *o.OrgKeyHash
}

// GetOrgKeyHashOk returns a tuple with the OrgKeyHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) GetOrgKeyHashOk() (*string, bool) {
	if o == nil || IsNil(o.OrgKeyHash) {
		return nil, false
	}
	return o.OrgKeyHash, true
}

// IsSetOrgKeyHash returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) IsSetOrgKeyHash() bool {
	if o != nil && !IsNil(o.OrgKeyHash) {
		return true
	}

	return false
}

// SetOrgKeyHash gets a reference to the given string and assigns it to the OrgKeyHash field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) SetOrgKeyHash(v string) {
	o.OrgKeyHash = &v
}

func (o ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Databags) {
		toSerialize["databags"] = o.Databags
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.Org) {
		toSerialize["org"] = o.Org
	}
	if !IsNil(o.ChefUser) {
		toSerialize["chefUser"] = o.ChefUser
	}
	if !IsNil(o.UserKey) {
		toSerialize["userKey"] = o.UserKey
	}
	if !IsNil(o.OrgKey) {
		toSerialize["orgKey"] = o.OrgKey
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.ChefUseFqdn) {
		toSerialize["chefUseFqdn"] = o.ChefUseFqdn
	}
	if !IsNil(o.WindowsVersion) {
		toSerialize["windowsVersion"] = o.WindowsVersion
	}
	if !IsNil(o.WindowsInstallUrl) {
		toSerialize["windowsInstallUrl"] = o.WindowsInstallUrl
	}
	if !IsNil(o.UserKeyHash) {
		toSerialize["userKeyHash"] = o.UserKeyHash
	}
	if !IsNil(o.OrgKeyHash) {
		toSerialize["orgKeyHash"] = o.OrgKeyHash
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

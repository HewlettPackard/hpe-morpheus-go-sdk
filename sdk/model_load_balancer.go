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

// checks if the LoadBalancer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancer{}

// LoadBalancer struct for LoadBalancer
type LoadBalancer struct {
	Id                   *int64                                                                   `json:"id,omitempty"`
	Uuid                 *string                                                                  `json:"uuid,omitempty"`
	Name                 *string                                                                  `json:"name,omitempty"`
	AccountId            *int64                                                                   `json:"accountId,omitempty"`
	Cloud                *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner  `json:"cloud,omitempty"`
	Type                 *ListBackupSettings200ResponseBackupSettingsDefaultSchedule              `json:"type,omitempty"`
	Owner                *GetAlerts200ResponseAllOfCheckGroupsInnerInstance                       `json:"owner,omitempty"`
	Visibility           *string                                                                  `json:"visibility,omitempty"`
	Description          *string                                                                  `json:"description,omitempty"`
	Host                 *string                                                                  `json:"host,omitempty"`
	Port                 *int64                                                                   `json:"port,omitempty"`
	Username             *string                                                                  `json:"username,omitempty"`
	Ip                   *string                                                                  `json:"ip,omitempty"`
	InternalIp           *string                                                                  `json:"internalIp,omitempty"`
	ExternalIp           *string                                                                  `json:"externalIp,omitempty"`
	ApiPort              *string                                                                  `json:"apiPort,omitempty"`
	AdminPort            *string                                                                  `json:"adminPort,omitempty"`
	SslEnabled           *bool                                                                    `json:"sslEnabled,omitempty"`
	SslCert              *string                                                                  `json:"sslCert,omitempty"`
	Config               map[string]interface{}                                                   `json:"config,omitempty"`
	DateCreated          *time.Time                                                               `json:"dateCreated,omitempty"`
	LastUpdated          *time.Time                                                               `json:"lastUpdated,omitempty"`
	Credential           *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential        `json:"credential,omitempty"`
	Tenants              []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"tenants,omitempty"`
	ResourcePermission   *ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission    `json:"resourcePermission,omitempty"`
	AdditionalProperties map[string]interface{}                                                   `json:",remain"`
}

type _LoadBalancer LoadBalancer

// NewLoadBalancer instantiates a new LoadBalancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancer() *LoadBalancer {
	this := LoadBalancer{}
	return &this
}

// NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerWithDefaults() *LoadBalancer {
	this := LoadBalancer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LoadBalancer) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *LoadBalancer) SetId(v int64) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *LoadBalancer) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// IsSetUuid returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *LoadBalancer) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadBalancer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadBalancer) SetName(v string) {
	o.Name = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *LoadBalancer) GetAccountId() int64 {
	if o == nil || IsNil(o.AccountId) {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// IsSetAccountId returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *LoadBalancer) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *LoadBalancer) GetCloud() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Cloud) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetCloudOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Cloud) {
		return nil, false
	}
	return o.Cloud, true
}

// IsSetCloud returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetCloud() bool {
	if o != nil && !IsNil(o.Cloud) {
		return true
	}

	return false
}

// SetCloud gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Cloud field.
func (o *LoadBalancer) SetCloud(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Cloud = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LoadBalancer) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule {
	if o == nil || IsNil(o.Type) {
		var ret ListBackupSettings200ResponseBackupSettingsDefaultSchedule
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ListBackupSettings200ResponseBackupSettingsDefaultSchedule and assigns it to the Type field.
func (o *LoadBalancer) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule) {
	o.Type = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *LoadBalancer) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance {
	if o == nil || IsNil(o.Owner) {
		var ret GetAlerts200ResponseAllOfCheckGroupsInnerInstance
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// IsSetOwner returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given GetAlerts200ResponseAllOfCheckGroupsInnerInstance and assigns it to the Owner field.
func (o *LoadBalancer) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance) {
	o.Owner = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *LoadBalancer) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *LoadBalancer) SetVisibility(v string) {
	o.Visibility = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LoadBalancer) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LoadBalancer) SetDescription(v string) {
	o.Description = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *LoadBalancer) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// IsSetHost returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *LoadBalancer) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LoadBalancer) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// IsSetPort returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *LoadBalancer) SetPort(v int64) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *LoadBalancer) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// IsSetUsername returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *LoadBalancer) SetUsername(v string) {
	o.Username = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *LoadBalancer) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// IsSetIp returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *LoadBalancer) SetIp(v string) {
	o.Ip = &v
}

// GetInternalIp returns the InternalIp field value if set, zero value otherwise.
func (o *LoadBalancer) GetInternalIp() string {
	if o == nil || IsNil(o.InternalIp) {
		var ret string
		return ret
	}
	return *o.InternalIp
}

// GetInternalIpOk returns a tuple with the InternalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetInternalIpOk() (*string, bool) {
	if o == nil || IsNil(o.InternalIp) {
		return nil, false
	}
	return o.InternalIp, true
}

// IsSetInternalIp returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetInternalIp() bool {
	if o != nil && !IsNil(o.InternalIp) {
		return true
	}

	return false
}

// SetInternalIp gets a reference to the given string and assigns it to the InternalIp field.
func (o *LoadBalancer) SetInternalIp(v string) {
	o.InternalIp = &v
}

// GetExternalIp returns the ExternalIp field value if set, zero value otherwise.
func (o *LoadBalancer) GetExternalIp() string {
	if o == nil || IsNil(o.ExternalIp) {
		var ret string
		return ret
	}
	return *o.ExternalIp
}

// GetExternalIpOk returns a tuple with the ExternalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetExternalIpOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalIp) {
		return nil, false
	}
	return o.ExternalIp, true
}

// IsSetExternalIp returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetExternalIp() bool {
	if o != nil && !IsNil(o.ExternalIp) {
		return true
	}

	return false
}

// SetExternalIp gets a reference to the given string and assigns it to the ExternalIp field.
func (o *LoadBalancer) SetExternalIp(v string) {
	o.ExternalIp = &v
}

// GetApiPort returns the ApiPort field value if set, zero value otherwise.
func (o *LoadBalancer) GetApiPort() string {
	if o == nil || IsNil(o.ApiPort) {
		var ret string
		return ret
	}
	return *o.ApiPort
}

// GetApiPortOk returns a tuple with the ApiPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetApiPortOk() (*string, bool) {
	if o == nil || IsNil(o.ApiPort) {
		return nil, false
	}
	return o.ApiPort, true
}

// IsSetApiPort returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetApiPort() bool {
	if o != nil && !IsNil(o.ApiPort) {
		return true
	}

	return false
}

// SetApiPort gets a reference to the given string and assigns it to the ApiPort field.
func (o *LoadBalancer) SetApiPort(v string) {
	o.ApiPort = &v
}

// GetAdminPort returns the AdminPort field value if set, zero value otherwise.
func (o *LoadBalancer) GetAdminPort() string {
	if o == nil || IsNil(o.AdminPort) {
		var ret string
		return ret
	}
	return *o.AdminPort
}

// GetAdminPortOk returns a tuple with the AdminPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetAdminPortOk() (*string, bool) {
	if o == nil || IsNil(o.AdminPort) {
		return nil, false
	}
	return o.AdminPort, true
}

// IsSetAdminPort returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetAdminPort() bool {
	if o != nil && !IsNil(o.AdminPort) {
		return true
	}

	return false
}

// SetAdminPort gets a reference to the given string and assigns it to the AdminPort field.
func (o *LoadBalancer) SetAdminPort(v string) {
	o.AdminPort = &v
}

// GetSslEnabled returns the SslEnabled field value if set, zero value otherwise.
func (o *LoadBalancer) GetSslEnabled() bool {
	if o == nil || IsNil(o.SslEnabled) {
		var ret bool
		return ret
	}
	return *o.SslEnabled
}

// GetSslEnabledOk returns a tuple with the SslEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetSslEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SslEnabled) {
		return nil, false
	}
	return o.SslEnabled, true
}

// IsSetSslEnabled returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetSslEnabled() bool {
	if o != nil && !IsNil(o.SslEnabled) {
		return true
	}

	return false
}

// SetSslEnabled gets a reference to the given bool and assigns it to the SslEnabled field.
func (o *LoadBalancer) SetSslEnabled(v bool) {
	o.SslEnabled = &v
}

// GetSslCert returns the SslCert field value if set, zero value otherwise.
func (o *LoadBalancer) GetSslCert() string {
	if o == nil || IsNil(o.SslCert) {
		var ret string
		return ret
	}
	return *o.SslCert
}

// GetSslCertOk returns a tuple with the SslCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetSslCertOk() (*string, bool) {
	if o == nil || IsNil(o.SslCert) {
		return nil, false
	}
	return o.SslCert, true
}

// IsSetSslCert returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetSslCert() bool {
	if o != nil && !IsNil(o.SslCert) {
		return true
	}

	return false
}

// SetSslCert gets a reference to the given string and assigns it to the SslCert field.
func (o *LoadBalancer) SetSslCert(v string) {
	o.SslCert = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *LoadBalancer) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *LoadBalancer) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *LoadBalancer) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *LoadBalancer) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *LoadBalancer) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *LoadBalancer) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *LoadBalancer) GetCredential() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential {
	if o == nil || IsNil(o.Credential) {
		var ret ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetCredentialOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// IsSetCredential returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential and assigns it to the Credential field.
func (o *LoadBalancer) SetCredential(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential) {
	o.Credential = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *LoadBalancer) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Tenants) {
		var ret []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetTenantsOk() ([]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// IsSetTenants returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Tenants field.
func (o *LoadBalancer) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Tenants = v
}

// GetResourcePermission returns the ResourcePermission field value if set, zero value otherwise.
func (o *LoadBalancer) GetResourcePermission() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission {
	if o == nil || IsNil(o.ResourcePermission) {
		var ret ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission
		return ret
	}
	return *o.ResourcePermission
}

// GetResourcePermissionOk returns a tuple with the ResourcePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetResourcePermissionOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool) {
	if o == nil || IsNil(o.ResourcePermission) {
		return nil, false
	}
	return o.ResourcePermission, true
}

// IsSetResourcePermission returns a boolean if a field has been set.
func (o *LoadBalancer) IsSetResourcePermission() bool {
	if o != nil && !IsNil(o.ResourcePermission) {
		return true
	}

	return false
}

// SetResourcePermission gets a reference to the given ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission and assigns it to the ResourcePermission field.
func (o *LoadBalancer) SetResourcePermission(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission) {
	o.ResourcePermission = &v
}

func (o LoadBalancer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.Cloud) {
		toSerialize["cloud"] = o.Cloud
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.InternalIp) {
		toSerialize["internalIp"] = o.InternalIp
	}
	if !IsNil(o.ExternalIp) {
		toSerialize["externalIp"] = o.ExternalIp
	}
	if !IsNil(o.ApiPort) {
		toSerialize["apiPort"] = o.ApiPort
	}
	if !IsNil(o.AdminPort) {
		toSerialize["adminPort"] = o.AdminPort
	}
	if !IsNil(o.SslEnabled) {
		toSerialize["sslEnabled"] = o.SslEnabled
	}
	if !IsNil(o.SslCert) {
		toSerialize["sslCert"] = o.SslCert
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Credential) {
		toSerialize["credential"] = o.Credential
	}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}
	if !IsNil(o.ResourcePermission) {
		toSerialize["resourcePermission"] = o.ResourcePermission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *LoadBalancer) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

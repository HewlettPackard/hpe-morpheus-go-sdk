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

// checks if the ZoneAwsConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneAwsConfig{}

// ZoneAwsConfig struct for ZoneAwsConfig
type ZoneAwsConfig struct {
	Endpoint             *string                                                       `json:"endpoint,omitempty"`
	AccessKey            *string                                                       `json:"accessKey,omitempty"`
	SecretKey            *string                                                       `json:"secretKey,omitempty"`
	UseHostCredentials   *string                                                       `json:"useHostCredentials,omitempty"`
	StsAssumeRole        *string                                                       `json:"stsAssumeRole,omitempty"`
	IsVpc                *string                                                       `json:"isVpc,omitempty"`
	Vpc                  *string                                                       `json:"vpc,omitempty"`
	ImageStoreId         *string                                                       `json:"imageStoreId,omitempty"`
	EbsEncryption        *string                                                       `json:"ebsEncryption,omitempty"`
	CostingReport        *string                                                       `json:"costingReport,omitempty"`
	CostingFolder        *string                                                       `json:"costingFolder,omitempty"`
	CostingBucket        *string                                                       `json:"costingBucket,omitempty"`
	CostingBucketName    *string                                                       `json:"costingBucketName,omitempty"`
	CostingRegion        *string                                                       `json:"costingRegion,omitempty"`
	CostingAccessKey     *string                                                       `json:"costingAccessKey,omitempty"`
	CostingSecretKey     *string                                                       `json:"costingSecretKey,omitempty"`
	CostingReportName    *string                                                       `json:"costingReportName,omitempty"`
	ApplianceUrl         *string                                                       `json:"applianceUrl,omitempty"`
	DatacenterName       *string                                                       `json:"datacenterName,omitempty"`
	NetworkServerId      *string                                                       `json:"networkServer.id,omitempty"`
	NetworkServer        *ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer `json:"networkServer,omitempty"`
	SecurityServer       *string                                                       `json:"securityServer,omitempty"`
	CertificateProvider  *string                                                       `json:"certificateProvider,omitempty"`
	BackupMode           *string                                                       `json:"backupMode,omitempty"`
	ReplicationMode      *string                                                       `json:"replicationMode,omitempty"`
	DnsIntegrationId     *string                                                       `json:"dnsIntegrationId,omitempty"`
	ServiceRegistryId    *string                                                       `json:"serviceRegistryId,omitempty"`
	ConfigManagementId   *string                                                       `json:"configManagementId,omitempty"`
	ConfigCmdbDiscovery  *bool                                                         `json:"configCmdbDiscovery,omitempty"`
	SecretKeyHash        *string                                                       `json:"secretKeyHash,omitempty"`
	CostingSecretKeyHash *string                                                       `json:"costingSecretKeyHash,omitempty"`
	AdditionalProperties map[string]interface{}                                        `json:",remain"`
}

type _ZoneAwsConfig ZoneAwsConfig

// NewZoneAwsConfig instantiates a new ZoneAwsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneAwsConfig() *ZoneAwsConfig {
	this := ZoneAwsConfig{}
	return &this
}

// NewZoneAwsConfigWithDefaults instantiates a new ZoneAwsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneAwsConfigWithDefaults() *ZoneAwsConfig {
	this := ZoneAwsConfig{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// IsSetEndpoint returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *ZoneAwsConfig) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// IsSetAccessKey returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *ZoneAwsConfig) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// IsSetSecretKey returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *ZoneAwsConfig) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetUseHostCredentials returns the UseHostCredentials field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetUseHostCredentials() string {
	if o == nil || IsNil(o.UseHostCredentials) {
		var ret string
		return ret
	}
	return *o.UseHostCredentials
}

// GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetUseHostCredentialsOk() (*string, bool) {
	if o == nil || IsNil(o.UseHostCredentials) {
		return nil, false
	}
	return o.UseHostCredentials, true
}

// IsSetUseHostCredentials returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetUseHostCredentials() bool {
	if o != nil && !IsNil(o.UseHostCredentials) {
		return true
	}

	return false
}

// SetUseHostCredentials gets a reference to the given string and assigns it to the UseHostCredentials field.
func (o *ZoneAwsConfig) SetUseHostCredentials(v string) {
	o.UseHostCredentials = &v
}

// GetStsAssumeRole returns the StsAssumeRole field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetStsAssumeRole() string {
	if o == nil || IsNil(o.StsAssumeRole) {
		var ret string
		return ret
	}
	return *o.StsAssumeRole
}

// GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetStsAssumeRoleOk() (*string, bool) {
	if o == nil || IsNil(o.StsAssumeRole) {
		return nil, false
	}
	return o.StsAssumeRole, true
}

// IsSetStsAssumeRole returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetStsAssumeRole() bool {
	if o != nil && !IsNil(o.StsAssumeRole) {
		return true
	}

	return false
}

// SetStsAssumeRole gets a reference to the given string and assigns it to the StsAssumeRole field.
func (o *ZoneAwsConfig) SetStsAssumeRole(v string) {
	o.StsAssumeRole = &v
}

// GetIsVpc returns the IsVpc field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetIsVpc() string {
	if o == nil || IsNil(o.IsVpc) {
		var ret string
		return ret
	}
	return *o.IsVpc
}

// GetIsVpcOk returns a tuple with the IsVpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetIsVpcOk() (*string, bool) {
	if o == nil || IsNil(o.IsVpc) {
		return nil, false
	}
	return o.IsVpc, true
}

// IsSetIsVpc returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetIsVpc() bool {
	if o != nil && !IsNil(o.IsVpc) {
		return true
	}

	return false
}

// SetIsVpc gets a reference to the given string and assigns it to the IsVpc field.
func (o *ZoneAwsConfig) SetIsVpc(v string) {
	o.IsVpc = &v
}

// GetVpc returns the Vpc field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetVpc() string {
	if o == nil || IsNil(o.Vpc) {
		var ret string
		return ret
	}
	return *o.Vpc
}

// GetVpcOk returns a tuple with the Vpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetVpcOk() (*string, bool) {
	if o == nil || IsNil(o.Vpc) {
		return nil, false
	}
	return o.Vpc, true
}

// IsSetVpc returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetVpc() bool {
	if o != nil && !IsNil(o.Vpc) {
		return true
	}

	return false
}

// SetVpc gets a reference to the given string and assigns it to the Vpc field.
func (o *ZoneAwsConfig) SetVpc(v string) {
	o.Vpc = &v
}

// GetImageStoreId returns the ImageStoreId field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetImageStoreId() string {
	if o == nil || IsNil(o.ImageStoreId) {
		var ret string
		return ret
	}
	return *o.ImageStoreId
}

// GetImageStoreIdOk returns a tuple with the ImageStoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetImageStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageStoreId) {
		return nil, false
	}
	return o.ImageStoreId, true
}

// IsSetImageStoreId returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetImageStoreId() bool {
	if o != nil && !IsNil(o.ImageStoreId) {
		return true
	}

	return false
}

// SetImageStoreId gets a reference to the given string and assigns it to the ImageStoreId field.
func (o *ZoneAwsConfig) SetImageStoreId(v string) {
	o.ImageStoreId = &v
}

// GetEbsEncryption returns the EbsEncryption field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetEbsEncryption() string {
	if o == nil || IsNil(o.EbsEncryption) {
		var ret string
		return ret
	}
	return *o.EbsEncryption
}

// GetEbsEncryptionOk returns a tuple with the EbsEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetEbsEncryptionOk() (*string, bool) {
	if o == nil || IsNil(o.EbsEncryption) {
		return nil, false
	}
	return o.EbsEncryption, true
}

// IsSetEbsEncryption returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetEbsEncryption() bool {
	if o != nil && !IsNil(o.EbsEncryption) {
		return true
	}

	return false
}

// SetEbsEncryption gets a reference to the given string and assigns it to the EbsEncryption field.
func (o *ZoneAwsConfig) SetEbsEncryption(v string) {
	o.EbsEncryption = &v
}

// GetCostingReport returns the CostingReport field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetCostingReport() string {
	if o == nil || IsNil(o.CostingReport) {
		var ret string
		return ret
	}
	return *o.CostingReport
}

// GetCostingReportOk returns a tuple with the CostingReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetCostingReportOk() (*string, bool) {
	if o == nil || IsNil(o.CostingReport) {
		return nil, false
	}
	return o.CostingReport, true
}

// IsSetCostingReport returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetCostingReport() bool {
	if o != nil && !IsNil(o.CostingReport) {
		return true
	}

	return false
}

// SetCostingReport gets a reference to the given string and assigns it to the CostingReport field.
func (o *ZoneAwsConfig) SetCostingReport(v string) {
	o.CostingReport = &v
}

// GetCostingFolder returns the CostingFolder field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetCostingFolder() string {
	if o == nil || IsNil(o.CostingFolder) {
		var ret string
		return ret
	}
	return *o.CostingFolder
}

// GetCostingFolderOk returns a tuple with the CostingFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetCostingFolderOk() (*string, bool) {
	if o == nil || IsNil(o.CostingFolder) {
		return nil, false
	}
	return o.CostingFolder, true
}

// IsSetCostingFolder returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetCostingFolder() bool {
	if o != nil && !IsNil(o.CostingFolder) {
		return true
	}

	return false
}

// SetCostingFolder gets a reference to the given string and assigns it to the CostingFolder field.
func (o *ZoneAwsConfig) SetCostingFolder(v string) {
	o.CostingFolder = &v
}

// GetCostingBucket returns the CostingBucket field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetCostingBucket() string {
	if o == nil || IsNil(o.CostingBucket) {
		var ret string
		return ret
	}
	return *o.CostingBucket
}

// GetCostingBucketOk returns a tuple with the CostingBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetCostingBucketOk() (*string, bool) {
	if o == nil || IsNil(o.CostingBucket) {
		return nil, false
	}
	return o.CostingBucket, true
}

// IsSetCostingBucket returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetCostingBucket() bool {
	if o != nil && !IsNil(o.CostingBucket) {
		return true
	}

	return false
}

// SetCostingBucket gets a reference to the given string and assigns it to the CostingBucket field.
func (o *ZoneAwsConfig) SetCostingBucket(v string) {
	o.CostingBucket = &v
}

// GetCostingBucketName returns the CostingBucketName field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetCostingBucketName() string {
	if o == nil || IsNil(o.CostingBucketName) {
		var ret string
		return ret
	}
	return *o.CostingBucketName
}

// GetCostingBucketNameOk returns a tuple with the CostingBucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetCostingBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.CostingBucketName) {
		return nil, false
	}
	return o.CostingBucketName, true
}

// IsSetCostingBucketName returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetCostingBucketName() bool {
	if o != nil && !IsNil(o.CostingBucketName) {
		return true
	}

	return false
}

// SetCostingBucketName gets a reference to the given string and assigns it to the CostingBucketName field.
func (o *ZoneAwsConfig) SetCostingBucketName(v string) {
	o.CostingBucketName = &v
}

// GetCostingRegion returns the CostingRegion field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetCostingRegion() string {
	if o == nil || IsNil(o.CostingRegion) {
		var ret string
		return ret
	}
	return *o.CostingRegion
}

// GetCostingRegionOk returns a tuple with the CostingRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetCostingRegionOk() (*string, bool) {
	if o == nil || IsNil(o.CostingRegion) {
		return nil, false
	}
	return o.CostingRegion, true
}

// IsSetCostingRegion returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetCostingRegion() bool {
	if o != nil && !IsNil(o.CostingRegion) {
		return true
	}

	return false
}

// SetCostingRegion gets a reference to the given string and assigns it to the CostingRegion field.
func (o *ZoneAwsConfig) SetCostingRegion(v string) {
	o.CostingRegion = &v
}

// GetCostingAccessKey returns the CostingAccessKey field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetCostingAccessKey() string {
	if o == nil || IsNil(o.CostingAccessKey) {
		var ret string
		return ret
	}
	return *o.CostingAccessKey
}

// GetCostingAccessKeyOk returns a tuple with the CostingAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetCostingAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.CostingAccessKey) {
		return nil, false
	}
	return o.CostingAccessKey, true
}

// IsSetCostingAccessKey returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetCostingAccessKey() bool {
	if o != nil && !IsNil(o.CostingAccessKey) {
		return true
	}

	return false
}

// SetCostingAccessKey gets a reference to the given string and assigns it to the CostingAccessKey field.
func (o *ZoneAwsConfig) SetCostingAccessKey(v string) {
	o.CostingAccessKey = &v
}

// GetCostingSecretKey returns the CostingSecretKey field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetCostingSecretKey() string {
	if o == nil || IsNil(o.CostingSecretKey) {
		var ret string
		return ret
	}
	return *o.CostingSecretKey
}

// GetCostingSecretKeyOk returns a tuple with the CostingSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetCostingSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.CostingSecretKey) {
		return nil, false
	}
	return o.CostingSecretKey, true
}

// IsSetCostingSecretKey returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetCostingSecretKey() bool {
	if o != nil && !IsNil(o.CostingSecretKey) {
		return true
	}

	return false
}

// SetCostingSecretKey gets a reference to the given string and assigns it to the CostingSecretKey field.
func (o *ZoneAwsConfig) SetCostingSecretKey(v string) {
	o.CostingSecretKey = &v
}

// GetCostingReportName returns the CostingReportName field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetCostingReportName() string {
	if o == nil || IsNil(o.CostingReportName) {
		var ret string
		return ret
	}
	return *o.CostingReportName
}

// GetCostingReportNameOk returns a tuple with the CostingReportName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetCostingReportNameOk() (*string, bool) {
	if o == nil || IsNil(o.CostingReportName) {
		return nil, false
	}
	return o.CostingReportName, true
}

// IsSetCostingReportName returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetCostingReportName() bool {
	if o != nil && !IsNil(o.CostingReportName) {
		return true
	}

	return false
}

// SetCostingReportName gets a reference to the given string and assigns it to the CostingReportName field.
func (o *ZoneAwsConfig) SetCostingReportName(v string) {
	o.CostingReportName = &v
}

// GetApplianceUrl returns the ApplianceUrl field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetApplianceUrl() string {
	if o == nil || IsNil(o.ApplianceUrl) {
		var ret string
		return ret
	}
	return *o.ApplianceUrl
}

// GetApplianceUrlOk returns a tuple with the ApplianceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetApplianceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ApplianceUrl) {
		return nil, false
	}
	return o.ApplianceUrl, true
}

// IsSetApplianceUrl returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetApplianceUrl() bool {
	if o != nil && !IsNil(o.ApplianceUrl) {
		return true
	}

	return false
}

// SetApplianceUrl gets a reference to the given string and assigns it to the ApplianceUrl field.
func (o *ZoneAwsConfig) SetApplianceUrl(v string) {
	o.ApplianceUrl = &v
}

// GetDatacenterName returns the DatacenterName field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetDatacenterName() string {
	if o == nil || IsNil(o.DatacenterName) {
		var ret string
		return ret
	}
	return *o.DatacenterName
}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetDatacenterNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatacenterName) {
		return nil, false
	}
	return o.DatacenterName, true
}

// IsSetDatacenterName returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetDatacenterName() bool {
	if o != nil && !IsNil(o.DatacenterName) {
		return true
	}

	return false
}

// SetDatacenterName gets a reference to the given string and assigns it to the DatacenterName field.
func (o *ZoneAwsConfig) SetDatacenterName(v string) {
	o.DatacenterName = &v
}

// GetNetworkServerId returns the NetworkServerId field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetNetworkServerId() string {
	if o == nil || IsNil(o.NetworkServerId) {
		var ret string
		return ret
	}
	return *o.NetworkServerId
}

// GetNetworkServerIdOk returns a tuple with the NetworkServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetNetworkServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkServerId) {
		return nil, false
	}
	return o.NetworkServerId, true
}

// IsSetNetworkServerId returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetNetworkServerId() bool {
	if o != nil && !IsNil(o.NetworkServerId) {
		return true
	}

	return false
}

// SetNetworkServerId gets a reference to the given string and assigns it to the NetworkServerId field.
func (o *ZoneAwsConfig) SetNetworkServerId(v string) {
	o.NetworkServerId = &v
}

// GetNetworkServer returns the NetworkServer field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetNetworkServer() ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer {
	if o == nil || IsNil(o.NetworkServer) {
		var ret ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer
		return ret
	}
	return *o.NetworkServer
}

// GetNetworkServerOk returns a tuple with the NetworkServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetNetworkServerOk() (*ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer, bool) {
	if o == nil || IsNil(o.NetworkServer) {
		return nil, false
	}
	return o.NetworkServer, true
}

// IsSetNetworkServer returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetNetworkServer() bool {
	if o != nil && !IsNil(o.NetworkServer) {
		return true
	}

	return false
}

// SetNetworkServer gets a reference to the given ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer and assigns it to the NetworkServer field.
func (o *ZoneAwsConfig) SetNetworkServer(v ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer) {
	o.NetworkServer = &v
}

// GetSecurityServer returns the SecurityServer field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetSecurityServer() string {
	if o == nil || IsNil(o.SecurityServer) {
		var ret string
		return ret
	}
	return *o.SecurityServer
}

// GetSecurityServerOk returns a tuple with the SecurityServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetSecurityServerOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityServer) {
		return nil, false
	}
	return o.SecurityServer, true
}

// IsSetSecurityServer returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetSecurityServer() bool {
	if o != nil && !IsNil(o.SecurityServer) {
		return true
	}

	return false
}

// SetSecurityServer gets a reference to the given string and assigns it to the SecurityServer field.
func (o *ZoneAwsConfig) SetSecurityServer(v string) {
	o.SecurityServer = &v
}

// GetCertificateProvider returns the CertificateProvider field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetCertificateProvider() string {
	if o == nil || IsNil(o.CertificateProvider) {
		var ret string
		return ret
	}
	return *o.CertificateProvider
}

// GetCertificateProviderOk returns a tuple with the CertificateProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetCertificateProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateProvider) {
		return nil, false
	}
	return o.CertificateProvider, true
}

// IsSetCertificateProvider returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetCertificateProvider() bool {
	if o != nil && !IsNil(o.CertificateProvider) {
		return true
	}

	return false
}

// SetCertificateProvider gets a reference to the given string and assigns it to the CertificateProvider field.
func (o *ZoneAwsConfig) SetCertificateProvider(v string) {
	o.CertificateProvider = &v
}

// GetBackupMode returns the BackupMode field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetBackupMode() string {
	if o == nil || IsNil(o.BackupMode) {
		var ret string
		return ret
	}
	return *o.BackupMode
}

// GetBackupModeOk returns a tuple with the BackupMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetBackupModeOk() (*string, bool) {
	if o == nil || IsNil(o.BackupMode) {
		return nil, false
	}
	return o.BackupMode, true
}

// IsSetBackupMode returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetBackupMode() bool {
	if o != nil && !IsNil(o.BackupMode) {
		return true
	}

	return false
}

// SetBackupMode gets a reference to the given string and assigns it to the BackupMode field.
func (o *ZoneAwsConfig) SetBackupMode(v string) {
	o.BackupMode = &v
}

// GetReplicationMode returns the ReplicationMode field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetReplicationMode() string {
	if o == nil || IsNil(o.ReplicationMode) {
		var ret string
		return ret
	}
	return *o.ReplicationMode
}

// GetReplicationModeOk returns a tuple with the ReplicationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetReplicationModeOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationMode) {
		return nil, false
	}
	return o.ReplicationMode, true
}

// IsSetReplicationMode returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetReplicationMode() bool {
	if o != nil && !IsNil(o.ReplicationMode) {
		return true
	}

	return false
}

// SetReplicationMode gets a reference to the given string and assigns it to the ReplicationMode field.
func (o *ZoneAwsConfig) SetReplicationMode(v string) {
	o.ReplicationMode = &v
}

// GetDnsIntegrationId returns the DnsIntegrationId field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetDnsIntegrationId() string {
	if o == nil || IsNil(o.DnsIntegrationId) {
		var ret string
		return ret
	}
	return *o.DnsIntegrationId
}

// GetDnsIntegrationIdOk returns a tuple with the DnsIntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetDnsIntegrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.DnsIntegrationId) {
		return nil, false
	}
	return o.DnsIntegrationId, true
}

// IsSetDnsIntegrationId returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetDnsIntegrationId() bool {
	if o != nil && !IsNil(o.DnsIntegrationId) {
		return true
	}

	return false
}

// SetDnsIntegrationId gets a reference to the given string and assigns it to the DnsIntegrationId field.
func (o *ZoneAwsConfig) SetDnsIntegrationId(v string) {
	o.DnsIntegrationId = &v
}

// GetServiceRegistryId returns the ServiceRegistryId field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetServiceRegistryId() string {
	if o == nil || IsNil(o.ServiceRegistryId) {
		var ret string
		return ret
	}
	return *o.ServiceRegistryId
}

// GetServiceRegistryIdOk returns a tuple with the ServiceRegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetServiceRegistryIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceRegistryId) {
		return nil, false
	}
	return o.ServiceRegistryId, true
}

// IsSetServiceRegistryId returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetServiceRegistryId() bool {
	if o != nil && !IsNil(o.ServiceRegistryId) {
		return true
	}

	return false
}

// SetServiceRegistryId gets a reference to the given string and assigns it to the ServiceRegistryId field.
func (o *ZoneAwsConfig) SetServiceRegistryId(v string) {
	o.ServiceRegistryId = &v
}

// GetConfigManagementId returns the ConfigManagementId field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetConfigManagementId() string {
	if o == nil || IsNil(o.ConfigManagementId) {
		var ret string
		return ret
	}
	return *o.ConfigManagementId
}

// GetConfigManagementIdOk returns a tuple with the ConfigManagementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetConfigManagementIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigManagementId) {
		return nil, false
	}
	return o.ConfigManagementId, true
}

// IsSetConfigManagementId returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetConfigManagementId() bool {
	if o != nil && !IsNil(o.ConfigManagementId) {
		return true
	}

	return false
}

// SetConfigManagementId gets a reference to the given string and assigns it to the ConfigManagementId field.
func (o *ZoneAwsConfig) SetConfigManagementId(v string) {
	o.ConfigManagementId = &v
}

// GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetConfigCmdbDiscovery() bool {
	if o == nil || IsNil(o.ConfigCmdbDiscovery) {
		var ret bool
		return ret
	}
	return *o.ConfigCmdbDiscovery
}

// GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetConfigCmdbDiscoveryOk() (*bool, bool) {
	if o == nil || IsNil(o.ConfigCmdbDiscovery) {
		return nil, false
	}
	return o.ConfigCmdbDiscovery, true
}

// IsSetConfigCmdbDiscovery returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetConfigCmdbDiscovery() bool {
	if o != nil && !IsNil(o.ConfigCmdbDiscovery) {
		return true
	}

	return false
}

// SetConfigCmdbDiscovery gets a reference to the given bool and assigns it to the ConfigCmdbDiscovery field.
func (o *ZoneAwsConfig) SetConfigCmdbDiscovery(v bool) {
	o.ConfigCmdbDiscovery = &v
}

// GetSecretKeyHash returns the SecretKeyHash field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetSecretKeyHash() string {
	if o == nil || IsNil(o.SecretKeyHash) {
		var ret string
		return ret
	}
	return *o.SecretKeyHash
}

// GetSecretKeyHashOk returns a tuple with the SecretKeyHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetSecretKeyHashOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKeyHash) {
		return nil, false
	}
	return o.SecretKeyHash, true
}

// IsSetSecretKeyHash returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetSecretKeyHash() bool {
	if o != nil && !IsNil(o.SecretKeyHash) {
		return true
	}

	return false
}

// SetSecretKeyHash gets a reference to the given string and assigns it to the SecretKeyHash field.
func (o *ZoneAwsConfig) SetSecretKeyHash(v string) {
	o.SecretKeyHash = &v
}

// GetCostingSecretKeyHash returns the CostingSecretKeyHash field value if set, zero value otherwise.
func (o *ZoneAwsConfig) GetCostingSecretKeyHash() string {
	if o == nil || IsNil(o.CostingSecretKeyHash) {
		var ret string
		return ret
	}
	return *o.CostingSecretKeyHash
}

// GetCostingSecretKeyHashOk returns a tuple with the CostingSecretKeyHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAwsConfig) GetCostingSecretKeyHashOk() (*string, bool) {
	if o == nil || IsNil(o.CostingSecretKeyHash) {
		return nil, false
	}
	return o.CostingSecretKeyHash, true
}

// IsSetCostingSecretKeyHash returns a boolean if a field has been set.
func (o *ZoneAwsConfig) IsSetCostingSecretKeyHash() bool {
	if o != nil && !IsNil(o.CostingSecretKeyHash) {
		return true
	}

	return false
}

// SetCostingSecretKeyHash gets a reference to the given string and assigns it to the CostingSecretKeyHash field.
func (o *ZoneAwsConfig) SetCostingSecretKeyHash(v string) {
	o.CostingSecretKeyHash = &v
}

func (o ZoneAwsConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneAwsConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.AccessKey) {
		toSerialize["accessKey"] = o.AccessKey
	}
	if !IsNil(o.SecretKey) {
		toSerialize["secretKey"] = o.SecretKey
	}
	if !IsNil(o.UseHostCredentials) {
		toSerialize["useHostCredentials"] = o.UseHostCredentials
	}
	if !IsNil(o.StsAssumeRole) {
		toSerialize["stsAssumeRole"] = o.StsAssumeRole
	}
	if !IsNil(o.IsVpc) {
		toSerialize["isVpc"] = o.IsVpc
	}
	if !IsNil(o.Vpc) {
		toSerialize["vpc"] = o.Vpc
	}
	if !IsNil(o.ImageStoreId) {
		toSerialize["imageStoreId"] = o.ImageStoreId
	}
	if !IsNil(o.EbsEncryption) {
		toSerialize["ebsEncryption"] = o.EbsEncryption
	}
	if !IsNil(o.CostingReport) {
		toSerialize["costingReport"] = o.CostingReport
	}
	if !IsNil(o.CostingFolder) {
		toSerialize["costingFolder"] = o.CostingFolder
	}
	if !IsNil(o.CostingBucket) {
		toSerialize["costingBucket"] = o.CostingBucket
	}
	if !IsNil(o.CostingBucketName) {
		toSerialize["costingBucketName"] = o.CostingBucketName
	}
	if !IsNil(o.CostingRegion) {
		toSerialize["costingRegion"] = o.CostingRegion
	}
	if !IsNil(o.CostingAccessKey) {
		toSerialize["costingAccessKey"] = o.CostingAccessKey
	}
	if !IsNil(o.CostingSecretKey) {
		toSerialize["costingSecretKey"] = o.CostingSecretKey
	}
	if !IsNil(o.CostingReportName) {
		toSerialize["costingReportName"] = o.CostingReportName
	}
	if !IsNil(o.ApplianceUrl) {
		toSerialize["applianceUrl"] = o.ApplianceUrl
	}
	if !IsNil(o.DatacenterName) {
		toSerialize["datacenterName"] = o.DatacenterName
	}
	if !IsNil(o.NetworkServerId) {
		toSerialize["networkServer.id"] = o.NetworkServerId
	}
	if !IsNil(o.NetworkServer) {
		toSerialize["networkServer"] = o.NetworkServer
	}
	if !IsNil(o.SecurityServer) {
		toSerialize["securityServer"] = o.SecurityServer
	}
	if !IsNil(o.CertificateProvider) {
		toSerialize["certificateProvider"] = o.CertificateProvider
	}
	if !IsNil(o.BackupMode) {
		toSerialize["backupMode"] = o.BackupMode
	}
	if !IsNil(o.ReplicationMode) {
		toSerialize["replicationMode"] = o.ReplicationMode
	}
	if !IsNil(o.DnsIntegrationId) {
		toSerialize["dnsIntegrationId"] = o.DnsIntegrationId
	}
	if !IsNil(o.ServiceRegistryId) {
		toSerialize["serviceRegistryId"] = o.ServiceRegistryId
	}
	if !IsNil(o.ConfigManagementId) {
		toSerialize["configManagementId"] = o.ConfigManagementId
	}
	if !IsNil(o.ConfigCmdbDiscovery) {
		toSerialize["configCmdbDiscovery"] = o.ConfigCmdbDiscovery
	}
	if !IsNil(o.SecretKeyHash) {
		toSerialize["secretKeyHash"] = o.SecretKeyHash
	}
	if !IsNil(o.CostingSecretKeyHash) {
		toSerialize["costingSecretKeyHash"] = o.CostingSecretKeyHash
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ZoneAwsConfig) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache

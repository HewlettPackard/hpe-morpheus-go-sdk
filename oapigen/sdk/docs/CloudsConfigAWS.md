# CloudsConfigAWS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **string** |  | [optional] 
**AccessKey** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**UseHostCredentials** | Pointer to **string** |  | [optional] 
**StsAssumeRole** | Pointer to **string** |  | [optional] 
**IsVpc** | Pointer to **string** |  | [optional] 
**Vpc** | Pointer to **string** |  | [optional] 
**ImageStoreId** | Pointer to **string** |  | [optional] 
**EbsEncryption** | Pointer to **string** |  | [optional] 
**CostingReport** | Pointer to **string** |  | [optional] 
**CostingFolder** | Pointer to **string** |  | [optional] 
**CostingBucket** | Pointer to **string** |  | [optional] 
**CostingBucketName** | Pointer to **string** |  | [optional] 
**CostingRegion** | Pointer to **string** |  | [optional] 
**CostingAccessKey** | Pointer to **string** |  | [optional] 
**CostingSecretKey** | Pointer to **NullableString** |  | [optional] 
**CostingReportName** | Pointer to **string** |  | [optional] 
**ApplianceUrl** | Pointer to **string** |  | [optional] 
**DatacenterName** | Pointer to **string** |  | [optional] 
**NetworkServerId** | Pointer to **string** |  | [optional] 
**NetworkServer** | Pointer to [**AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer**](AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer.md) |  | [optional] 
**SecurityServer** | Pointer to **NullableString** |  | [optional] 
**CertificateProvider** | Pointer to **string** |  | [optional] 
**BackupMode** | Pointer to **string** |  | [optional] 
**ReplicationMode** | Pointer to **string** |  | [optional] 
**DnsIntegrationId** | Pointer to **string** |  | [optional] 
**ServiceRegistryId** | Pointer to **string** |  | [optional] 
**ConfigManagementId** | Pointer to **string** |  | [optional] 
**ConfigCmdbDiscovery** | Pointer to **bool** |  | [optional] 
**SecretKeyHash** | Pointer to **string** |  | [optional] 
**CostingSecretKeyHash** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCloudsConfigAWS

`func NewCloudsConfigAWS() *CloudsConfigAWS`

NewCloudsConfigAWS instantiates a new CloudsConfigAWS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudsConfigAWSWithDefaults

`func NewCloudsConfigAWSWithDefaults() *CloudsConfigAWS`

NewCloudsConfigAWSWithDefaults instantiates a new CloudsConfigAWS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *CloudsConfigAWS) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CloudsConfigAWS) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CloudsConfigAWS) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *CloudsConfigAWS) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAccessKey

`func (o *CloudsConfigAWS) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *CloudsConfigAWS) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *CloudsConfigAWS) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *CloudsConfigAWS) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *CloudsConfigAWS) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *CloudsConfigAWS) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *CloudsConfigAWS) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *CloudsConfigAWS) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *CloudsConfigAWS) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *CloudsConfigAWS) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *CloudsConfigAWS) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *CloudsConfigAWS) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *CloudsConfigAWS) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *CloudsConfigAWS) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *CloudsConfigAWS) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *CloudsConfigAWS) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetIsVpc

`func (o *CloudsConfigAWS) GetIsVpc() string`

GetIsVpc returns the IsVpc field if non-nil, zero value otherwise.

### GetIsVpcOk

`func (o *CloudsConfigAWS) GetIsVpcOk() (*string, bool)`

GetIsVpcOk returns a tuple with the IsVpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpc

`func (o *CloudsConfigAWS) SetIsVpc(v string)`

SetIsVpc sets IsVpc field to given value.

### HasIsVpc

`func (o *CloudsConfigAWS) HasIsVpc() bool`

HasIsVpc returns a boolean if a field has been set.

### GetVpc

`func (o *CloudsConfigAWS) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *CloudsConfigAWS) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *CloudsConfigAWS) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *CloudsConfigAWS) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetImageStoreId

`func (o *CloudsConfigAWS) GetImageStoreId() string`

GetImageStoreId returns the ImageStoreId field if non-nil, zero value otherwise.

### GetImageStoreIdOk

`func (o *CloudsConfigAWS) GetImageStoreIdOk() (*string, bool)`

GetImageStoreIdOk returns a tuple with the ImageStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStoreId

`func (o *CloudsConfigAWS) SetImageStoreId(v string)`

SetImageStoreId sets ImageStoreId field to given value.

### HasImageStoreId

`func (o *CloudsConfigAWS) HasImageStoreId() bool`

HasImageStoreId returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *CloudsConfigAWS) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *CloudsConfigAWS) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *CloudsConfigAWS) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *CloudsConfigAWS) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetCostingReport

`func (o *CloudsConfigAWS) GetCostingReport() string`

GetCostingReport returns the CostingReport field if non-nil, zero value otherwise.

### GetCostingReportOk

`func (o *CloudsConfigAWS) GetCostingReportOk() (*string, bool)`

GetCostingReportOk returns a tuple with the CostingReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReport

`func (o *CloudsConfigAWS) SetCostingReport(v string)`

SetCostingReport sets CostingReport field to given value.

### HasCostingReport

`func (o *CloudsConfigAWS) HasCostingReport() bool`

HasCostingReport returns a boolean if a field has been set.

### GetCostingFolder

`func (o *CloudsConfigAWS) GetCostingFolder() string`

GetCostingFolder returns the CostingFolder field if non-nil, zero value otherwise.

### GetCostingFolderOk

`func (o *CloudsConfigAWS) GetCostingFolderOk() (*string, bool)`

GetCostingFolderOk returns a tuple with the CostingFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingFolder

`func (o *CloudsConfigAWS) SetCostingFolder(v string)`

SetCostingFolder sets CostingFolder field to given value.

### HasCostingFolder

`func (o *CloudsConfigAWS) HasCostingFolder() bool`

HasCostingFolder returns a boolean if a field has been set.

### GetCostingBucket

`func (o *CloudsConfigAWS) GetCostingBucket() string`

GetCostingBucket returns the CostingBucket field if non-nil, zero value otherwise.

### GetCostingBucketOk

`func (o *CloudsConfigAWS) GetCostingBucketOk() (*string, bool)`

GetCostingBucketOk returns a tuple with the CostingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucket

`func (o *CloudsConfigAWS) SetCostingBucket(v string)`

SetCostingBucket sets CostingBucket field to given value.

### HasCostingBucket

`func (o *CloudsConfigAWS) HasCostingBucket() bool`

HasCostingBucket returns a boolean if a field has been set.

### GetCostingBucketName

`func (o *CloudsConfigAWS) GetCostingBucketName() string`

GetCostingBucketName returns the CostingBucketName field if non-nil, zero value otherwise.

### GetCostingBucketNameOk

`func (o *CloudsConfigAWS) GetCostingBucketNameOk() (*string, bool)`

GetCostingBucketNameOk returns a tuple with the CostingBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucketName

`func (o *CloudsConfigAWS) SetCostingBucketName(v string)`

SetCostingBucketName sets CostingBucketName field to given value.

### HasCostingBucketName

`func (o *CloudsConfigAWS) HasCostingBucketName() bool`

HasCostingBucketName returns a boolean if a field has been set.

### GetCostingRegion

`func (o *CloudsConfigAWS) GetCostingRegion() string`

GetCostingRegion returns the CostingRegion field if non-nil, zero value otherwise.

### GetCostingRegionOk

`func (o *CloudsConfigAWS) GetCostingRegionOk() (*string, bool)`

GetCostingRegionOk returns a tuple with the CostingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingRegion

`func (o *CloudsConfigAWS) SetCostingRegion(v string)`

SetCostingRegion sets CostingRegion field to given value.

### HasCostingRegion

`func (o *CloudsConfigAWS) HasCostingRegion() bool`

HasCostingRegion returns a boolean if a field has been set.

### GetCostingAccessKey

`func (o *CloudsConfigAWS) GetCostingAccessKey() string`

GetCostingAccessKey returns the CostingAccessKey field if non-nil, zero value otherwise.

### GetCostingAccessKeyOk

`func (o *CloudsConfigAWS) GetCostingAccessKeyOk() (*string, bool)`

GetCostingAccessKeyOk returns a tuple with the CostingAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingAccessKey

`func (o *CloudsConfigAWS) SetCostingAccessKey(v string)`

SetCostingAccessKey sets CostingAccessKey field to given value.

### HasCostingAccessKey

`func (o *CloudsConfigAWS) HasCostingAccessKey() bool`

HasCostingAccessKey returns a boolean if a field has been set.

### GetCostingSecretKey

`func (o *CloudsConfigAWS) GetCostingSecretKey() string`

GetCostingSecretKey returns the CostingSecretKey field if non-nil, zero value otherwise.

### GetCostingSecretKeyOk

`func (o *CloudsConfigAWS) GetCostingSecretKeyOk() (*string, bool)`

GetCostingSecretKeyOk returns a tuple with the CostingSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKey

`func (o *CloudsConfigAWS) SetCostingSecretKey(v string)`

SetCostingSecretKey sets CostingSecretKey field to given value.

### HasCostingSecretKey

`func (o *CloudsConfigAWS) HasCostingSecretKey() bool`

HasCostingSecretKey returns a boolean if a field has been set.

### SetCostingSecretKeyNil

`func (o *CloudsConfigAWS) SetCostingSecretKeyNil(b bool)`

 SetCostingSecretKeyNil sets the value for CostingSecretKey to be an explicit nil

### UnsetCostingSecretKey
`func (o *CloudsConfigAWS) UnsetCostingSecretKey()`

UnsetCostingSecretKey ensures that no value is present for CostingSecretKey, not even an explicit nil
### GetCostingReportName

`func (o *CloudsConfigAWS) GetCostingReportName() string`

GetCostingReportName returns the CostingReportName field if non-nil, zero value otherwise.

### GetCostingReportNameOk

`func (o *CloudsConfigAWS) GetCostingReportNameOk() (*string, bool)`

GetCostingReportNameOk returns a tuple with the CostingReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReportName

`func (o *CloudsConfigAWS) SetCostingReportName(v string)`

SetCostingReportName sets CostingReportName field to given value.

### HasCostingReportName

`func (o *CloudsConfigAWS) HasCostingReportName() bool`

HasCostingReportName returns a boolean if a field has been set.

### GetApplianceUrl

`func (o *CloudsConfigAWS) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *CloudsConfigAWS) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *CloudsConfigAWS) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *CloudsConfigAWS) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *CloudsConfigAWS) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *CloudsConfigAWS) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *CloudsConfigAWS) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *CloudsConfigAWS) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *CloudsConfigAWS) GetNetworkServerId() string`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *CloudsConfigAWS) GetNetworkServerIdOk() (*string, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *CloudsConfigAWS) SetNetworkServerId(v string)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *CloudsConfigAWS) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetNetworkServer

`func (o *CloudsConfigAWS) GetNetworkServer() AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *CloudsConfigAWS) GetNetworkServerOk() (*AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *CloudsConfigAWS) SetNetworkServer(v AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *CloudsConfigAWS) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetSecurityServer

`func (o *CloudsConfigAWS) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *CloudsConfigAWS) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *CloudsConfigAWS) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *CloudsConfigAWS) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### SetSecurityServerNil

`func (o *CloudsConfigAWS) SetSecurityServerNil(b bool)`

 SetSecurityServerNil sets the value for SecurityServer to be an explicit nil

### UnsetSecurityServer
`func (o *CloudsConfigAWS) UnsetSecurityServer()`

UnsetSecurityServer ensures that no value is present for SecurityServer, not even an explicit nil
### GetCertificateProvider

`func (o *CloudsConfigAWS) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *CloudsConfigAWS) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *CloudsConfigAWS) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *CloudsConfigAWS) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetBackupMode

`func (o *CloudsConfigAWS) GetBackupMode() string`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *CloudsConfigAWS) GetBackupModeOk() (*string, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *CloudsConfigAWS) SetBackupMode(v string)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *CloudsConfigAWS) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### GetReplicationMode

`func (o *CloudsConfigAWS) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *CloudsConfigAWS) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *CloudsConfigAWS) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *CloudsConfigAWS) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### GetDnsIntegrationId

`func (o *CloudsConfigAWS) GetDnsIntegrationId() string`

GetDnsIntegrationId returns the DnsIntegrationId field if non-nil, zero value otherwise.

### GetDnsIntegrationIdOk

`func (o *CloudsConfigAWS) GetDnsIntegrationIdOk() (*string, bool)`

GetDnsIntegrationIdOk returns a tuple with the DnsIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrationId

`func (o *CloudsConfigAWS) SetDnsIntegrationId(v string)`

SetDnsIntegrationId sets DnsIntegrationId field to given value.

### HasDnsIntegrationId

`func (o *CloudsConfigAWS) HasDnsIntegrationId() bool`

HasDnsIntegrationId returns a boolean if a field has been set.

### GetServiceRegistryId

`func (o *CloudsConfigAWS) GetServiceRegistryId() string`

GetServiceRegistryId returns the ServiceRegistryId field if non-nil, zero value otherwise.

### GetServiceRegistryIdOk

`func (o *CloudsConfigAWS) GetServiceRegistryIdOk() (*string, bool)`

GetServiceRegistryIdOk returns a tuple with the ServiceRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRegistryId

`func (o *CloudsConfigAWS) SetServiceRegistryId(v string)`

SetServiceRegistryId sets ServiceRegistryId field to given value.

### HasServiceRegistryId

`func (o *CloudsConfigAWS) HasServiceRegistryId() bool`

HasServiceRegistryId returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *CloudsConfigAWS) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *CloudsConfigAWS) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *CloudsConfigAWS) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *CloudsConfigAWS) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetConfigCmdbDiscovery

`func (o *CloudsConfigAWS) GetConfigCmdbDiscovery() bool`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *CloudsConfigAWS) GetConfigCmdbDiscoveryOk() (*bool, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *CloudsConfigAWS) SetConfigCmdbDiscovery(v bool)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *CloudsConfigAWS) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetSecretKeyHash

`func (o *CloudsConfigAWS) GetSecretKeyHash() string`

GetSecretKeyHash returns the SecretKeyHash field if non-nil, zero value otherwise.

### GetSecretKeyHashOk

`func (o *CloudsConfigAWS) GetSecretKeyHashOk() (*string, bool)`

GetSecretKeyHashOk returns a tuple with the SecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyHash

`func (o *CloudsConfigAWS) SetSecretKeyHash(v string)`

SetSecretKeyHash sets SecretKeyHash field to given value.

### HasSecretKeyHash

`func (o *CloudsConfigAWS) HasSecretKeyHash() bool`

HasSecretKeyHash returns a boolean if a field has been set.

### GetCostingSecretKeyHash

`func (o *CloudsConfigAWS) GetCostingSecretKeyHash() string`

GetCostingSecretKeyHash returns the CostingSecretKeyHash field if non-nil, zero value otherwise.

### GetCostingSecretKeyHashOk

`func (o *CloudsConfigAWS) GetCostingSecretKeyHashOk() (*string, bool)`

GetCostingSecretKeyHashOk returns a tuple with the CostingSecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKeyHash

`func (o *CloudsConfigAWS) SetCostingSecretKeyHash(v string)`

SetCostingSecretKeyHash sets CostingSecretKeyHash field to given value.

### HasCostingSecretKeyHash

`func (o *CloudsConfigAWS) HasCostingSecretKeyHash() bool`

HasCostingSecretKeyHash returns a boolean if a field has been set.

### SetCostingSecretKeyHashNil

`func (o *CloudsConfigAWS) SetCostingSecretKeyHashNil(b bool)`

 SetCostingSecretKeyHashNil sets the value for CostingSecretKeyHash to be an explicit nil

### UnsetCostingSecretKeyHash
`func (o *CloudsConfigAWS) UnsetCostingSecretKeyHash()`

UnsetCostingSecretKeyHash ensures that no value is present for CostingSecretKeyHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



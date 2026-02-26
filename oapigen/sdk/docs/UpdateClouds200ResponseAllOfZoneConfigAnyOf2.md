# UpdateClouds200ResponseAllOfZoneConfigAnyOf2

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

### NewUpdateClouds200ResponseAllOfZoneConfigAnyOf2

`func NewUpdateClouds200ResponseAllOfZoneConfigAnyOf2() *UpdateClouds200ResponseAllOfZoneConfigAnyOf2`

NewUpdateClouds200ResponseAllOfZoneConfigAnyOf2 instantiates a new UpdateClouds200ResponseAllOfZoneConfigAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClouds200ResponseAllOfZoneConfigAnyOf2WithDefaults

`func NewUpdateClouds200ResponseAllOfZoneConfigAnyOf2WithDefaults() *UpdateClouds200ResponseAllOfZoneConfigAnyOf2`

NewUpdateClouds200ResponseAllOfZoneConfigAnyOf2WithDefaults instantiates a new UpdateClouds200ResponseAllOfZoneConfigAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAccessKey

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetIsVpc

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetIsVpc() string`

GetIsVpc returns the IsVpc field if non-nil, zero value otherwise.

### GetIsVpcOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetIsVpcOk() (*string, bool)`

GetIsVpcOk returns a tuple with the IsVpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpc

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetIsVpc(v string)`

SetIsVpc sets IsVpc field to given value.

### HasIsVpc

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasIsVpc() bool`

HasIsVpc returns a boolean if a field has been set.

### GetVpc

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetImageStoreId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetImageStoreId() string`

GetImageStoreId returns the ImageStoreId field if non-nil, zero value otherwise.

### GetImageStoreIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetImageStoreIdOk() (*string, bool)`

GetImageStoreIdOk returns a tuple with the ImageStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStoreId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetImageStoreId(v string)`

SetImageStoreId sets ImageStoreId field to given value.

### HasImageStoreId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasImageStoreId() bool`

HasImageStoreId returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetCostingReport

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingReport() string`

GetCostingReport returns the CostingReport field if non-nil, zero value otherwise.

### GetCostingReportOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingReportOk() (*string, bool)`

GetCostingReportOk returns a tuple with the CostingReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReport

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetCostingReport(v string)`

SetCostingReport sets CostingReport field to given value.

### HasCostingReport

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasCostingReport() bool`

HasCostingReport returns a boolean if a field has been set.

### GetCostingFolder

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingFolder() string`

GetCostingFolder returns the CostingFolder field if non-nil, zero value otherwise.

### GetCostingFolderOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingFolderOk() (*string, bool)`

GetCostingFolderOk returns a tuple with the CostingFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingFolder

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetCostingFolder(v string)`

SetCostingFolder sets CostingFolder field to given value.

### HasCostingFolder

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasCostingFolder() bool`

HasCostingFolder returns a boolean if a field has been set.

### GetCostingBucket

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingBucket() string`

GetCostingBucket returns the CostingBucket field if non-nil, zero value otherwise.

### GetCostingBucketOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingBucketOk() (*string, bool)`

GetCostingBucketOk returns a tuple with the CostingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucket

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetCostingBucket(v string)`

SetCostingBucket sets CostingBucket field to given value.

### HasCostingBucket

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasCostingBucket() bool`

HasCostingBucket returns a boolean if a field has been set.

### GetCostingBucketName

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingBucketName() string`

GetCostingBucketName returns the CostingBucketName field if non-nil, zero value otherwise.

### GetCostingBucketNameOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingBucketNameOk() (*string, bool)`

GetCostingBucketNameOk returns a tuple with the CostingBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucketName

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetCostingBucketName(v string)`

SetCostingBucketName sets CostingBucketName field to given value.

### HasCostingBucketName

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasCostingBucketName() bool`

HasCostingBucketName returns a boolean if a field has been set.

### GetCostingRegion

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingRegion() string`

GetCostingRegion returns the CostingRegion field if non-nil, zero value otherwise.

### GetCostingRegionOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingRegionOk() (*string, bool)`

GetCostingRegionOk returns a tuple with the CostingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingRegion

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetCostingRegion(v string)`

SetCostingRegion sets CostingRegion field to given value.

### HasCostingRegion

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasCostingRegion() bool`

HasCostingRegion returns a boolean if a field has been set.

### GetCostingAccessKey

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingAccessKey() string`

GetCostingAccessKey returns the CostingAccessKey field if non-nil, zero value otherwise.

### GetCostingAccessKeyOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingAccessKeyOk() (*string, bool)`

GetCostingAccessKeyOk returns a tuple with the CostingAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingAccessKey

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetCostingAccessKey(v string)`

SetCostingAccessKey sets CostingAccessKey field to given value.

### HasCostingAccessKey

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasCostingAccessKey() bool`

HasCostingAccessKey returns a boolean if a field has been set.

### GetCostingSecretKey

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingSecretKey() string`

GetCostingSecretKey returns the CostingSecretKey field if non-nil, zero value otherwise.

### GetCostingSecretKeyOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingSecretKeyOk() (*string, bool)`

GetCostingSecretKeyOk returns a tuple with the CostingSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKey

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetCostingSecretKey(v string)`

SetCostingSecretKey sets CostingSecretKey field to given value.

### HasCostingSecretKey

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasCostingSecretKey() bool`

HasCostingSecretKey returns a boolean if a field has been set.

### SetCostingSecretKeyNil

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetCostingSecretKeyNil(b bool)`

 SetCostingSecretKeyNil sets the value for CostingSecretKey to be an explicit nil

### UnsetCostingSecretKey
`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) UnsetCostingSecretKey()`

UnsetCostingSecretKey ensures that no value is present for CostingSecretKey, not even an explicit nil
### GetCostingReportName

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingReportName() string`

GetCostingReportName returns the CostingReportName field if non-nil, zero value otherwise.

### GetCostingReportNameOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingReportNameOk() (*string, bool)`

GetCostingReportNameOk returns a tuple with the CostingReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReportName

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetCostingReportName(v string)`

SetCostingReportName sets CostingReportName field to given value.

### HasCostingReportName

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasCostingReportName() bool`

HasCostingReportName returns a boolean if a field has been set.

### GetApplianceUrl

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetNetworkServerId() string`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetNetworkServerIdOk() (*string, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetNetworkServerId(v string)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetNetworkServer

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetNetworkServer() AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetNetworkServerOk() (*AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetNetworkServer(v AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetSecurityServer

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### SetSecurityServerNil

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetSecurityServerNil(b bool)`

 SetSecurityServerNil sets the value for SecurityServer to be an explicit nil

### UnsetSecurityServer
`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) UnsetSecurityServer()`

UnsetSecurityServer ensures that no value is present for SecurityServer, not even an explicit nil
### GetCertificateProvider

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetBackupMode

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetBackupMode() string`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetBackupModeOk() (*string, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetBackupMode(v string)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### GetReplicationMode

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### GetDnsIntegrationId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetDnsIntegrationId() string`

GetDnsIntegrationId returns the DnsIntegrationId field if non-nil, zero value otherwise.

### GetDnsIntegrationIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetDnsIntegrationIdOk() (*string, bool)`

GetDnsIntegrationIdOk returns a tuple with the DnsIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrationId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetDnsIntegrationId(v string)`

SetDnsIntegrationId sets DnsIntegrationId field to given value.

### HasDnsIntegrationId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasDnsIntegrationId() bool`

HasDnsIntegrationId returns a boolean if a field has been set.

### GetServiceRegistryId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetServiceRegistryId() string`

GetServiceRegistryId returns the ServiceRegistryId field if non-nil, zero value otherwise.

### GetServiceRegistryIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetServiceRegistryIdOk() (*string, bool)`

GetServiceRegistryIdOk returns a tuple with the ServiceRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRegistryId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetServiceRegistryId(v string)`

SetServiceRegistryId sets ServiceRegistryId field to given value.

### HasServiceRegistryId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasServiceRegistryId() bool`

HasServiceRegistryId returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetConfigCmdbDiscovery

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetConfigCmdbDiscovery() bool`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetConfigCmdbDiscoveryOk() (*bool, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetConfigCmdbDiscovery(v bool)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetSecretKeyHash

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetSecretKeyHash() string`

GetSecretKeyHash returns the SecretKeyHash field if non-nil, zero value otherwise.

### GetSecretKeyHashOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetSecretKeyHashOk() (*string, bool)`

GetSecretKeyHashOk returns a tuple with the SecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyHash

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetSecretKeyHash(v string)`

SetSecretKeyHash sets SecretKeyHash field to given value.

### HasSecretKeyHash

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasSecretKeyHash() bool`

HasSecretKeyHash returns a boolean if a field has been set.

### GetCostingSecretKeyHash

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingSecretKeyHash() string`

GetCostingSecretKeyHash returns the CostingSecretKeyHash field if non-nil, zero value otherwise.

### GetCostingSecretKeyHashOk

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) GetCostingSecretKeyHashOk() (*string, bool)`

GetCostingSecretKeyHashOk returns a tuple with the CostingSecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKeyHash

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetCostingSecretKeyHash(v string)`

SetCostingSecretKeyHash sets CostingSecretKeyHash field to given value.

### HasCostingSecretKeyHash

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) HasCostingSecretKeyHash() bool`

HasCostingSecretKeyHash returns a boolean if a field has been set.

### SetCostingSecretKeyHashNil

`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) SetCostingSecretKeyHashNil(b bool)`

 SetCostingSecretKeyHashNil sets the value for CostingSecretKeyHash to be an explicit nil

### UnsetCostingSecretKeyHash
`func (o *UpdateClouds200ResponseAllOfZoneConfigAnyOf2) UnsetCostingSecretKeyHash()`

UnsetCostingSecretKeyHash ensures that no value is present for CostingSecretKeyHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



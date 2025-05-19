# ListClouds200ResponseAllOfZonesInnerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiUrl** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Datacenter** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **string** |  | [optional] 
**ResourcePool** | Pointer to **string** |  | [optional] 
**RpcMode** | Pointer to **string** |  | [optional] 
**HideHostSelection** | Pointer to **string** |  | [optional] 
**ImportExisting** | Pointer to **string** |  | [optional] 
**EnableVnc** | Pointer to **string** |  | [optional] 
**EnableDiskTypeSelection** | Pointer to **string** |  | [optional] 
**EnableNetworkTypeSelection** | Pointer to **string** |  | [optional] 
**DiskStorageType** | Pointer to **string** |  | [optional] 
**ApplianceUrl** | Pointer to **string** |  | [optional] 
**DatacenterName** | Pointer to **string** |  | [optional] 
**NetworkServerId** | Pointer to **string** |  | [optional] 
**NetworkServer** | Pointer to [**ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer**](ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer.md) |  | [optional] 
**SecurityMode** | Pointer to **string** |  | [optional] 
**CertificateProvider** | Pointer to **string** |  | [optional] 
**BackupMode** | Pointer to **string** |  | [optional] 
**ReplicationMode** | Pointer to **string** |  | [optional] 
**DnsIntegrationId** | Pointer to **string** |  | [optional] 
**ConfigCmdbId** | Pointer to **string** |  | [optional] 
**ConfigManagementId** | Pointer to **string** |  | [optional] 
**ConfigCmId** | Pointer to **string** |  | [optional] 
**SecurityServer** | Pointer to **string** |  | [optional] 
**ServiceRegistryId** | Pointer to **string** |  | [optional] 
**KubeUrl** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**DatacenterId** | Pointer to **string** |  | [optional] 
**ConfigCmdbDiscovery** | Pointer to **bool** |  | [optional] 
**DistributedWorkerId** | Pointer to **string** |  | [optional] 
**PasswordHash** | Pointer to **string** |  | [optional] 
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
**CostingSecretKey** | Pointer to **string** |  | [optional] 
**CostingReportName** | Pointer to **string** |  | [optional] 
**SecretKeyHash** | Pointer to **string** |  | [optional] 
**CostingSecretKeyHash** | Pointer to **string** |  | [optional] 
**SubscriberId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**ResourceGroup** | Pointer to **string** |  | [optional] 
**InventoryLevel** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**CloudType** | Pointer to **string** |  | [optional] 
**DiskEncryption** | Pointer to **string** |  | [optional] 
**EncryptionSet** | Pointer to **string** |  | [optional] 
**CspTenantId** | Pointer to **string** |  | [optional] 
**CspClientId** | Pointer to **string** |  | [optional] 
**CspClientSecret** | Pointer to **string** |  | [optional] 
**CspCustomer** | Pointer to **string** |  | [optional] 
**AzureCostingMode** | Pointer to **string** |  | [optional] 
**ClientSecretHash** | Pointer to **string** |  | [optional] 
**CspClientSecretHash** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**ClientEmail** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**GoogleRegionId** | Pointer to **string** |  | [optional] 
**PrivateKeyHash** | Pointer to **string** |  | [optional] 

## Methods

### NewListClouds200ResponseAllOfZonesInnerConfig

`func NewListClouds200ResponseAllOfZonesInnerConfig() *ListClouds200ResponseAllOfZonesInnerConfig`

NewListClouds200ResponseAllOfZonesInnerConfig instantiates a new ListClouds200ResponseAllOfZonesInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClouds200ResponseAllOfZonesInnerConfigWithDefaults

`func NewListClouds200ResponseAllOfZonesInnerConfigWithDefaults() *ListClouds200ResponseAllOfZonesInnerConfig`

NewListClouds200ResponseAllOfZonesInnerConfigWithDefaults instantiates a new ListClouds200ResponseAllOfZonesInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUrl

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetUsername

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDatacenter

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetCluster

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetResourcePool

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetRpcMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetRpcMode() string`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetRpcModeOk() (*string, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetRpcMode(v string)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### GetHideHostSelection

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetHideHostSelection() string`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetHideHostSelectionOk() (*string, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetHideHostSelection(v string)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### GetImportExisting

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetImportExisting() string`

GetImportExisting returns the ImportExisting field if non-nil, zero value otherwise.

### GetImportExistingOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetImportExistingOk() (*string, bool)`

GetImportExistingOk returns a tuple with the ImportExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportExisting

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetImportExisting(v string)`

SetImportExisting sets ImportExisting field to given value.

### HasImportExisting

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasImportExisting() bool`

HasImportExisting returns a boolean if a field has been set.

### GetEnableVnc

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetEnableVnc() string`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetEnableVncOk() (*string, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetEnableVnc(v string)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### GetEnableDiskTypeSelection

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetEnableDiskTypeSelection() string`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetEnableDiskTypeSelectionOk() (*string, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetEnableDiskTypeSelection(v string)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### GetDiskStorageType

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDiskStorageType() string`

GetDiskStorageType returns the DiskStorageType field if non-nil, zero value otherwise.

### GetDiskStorageTypeOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDiskStorageTypeOk() (*string, bool)`

GetDiskStorageTypeOk returns a tuple with the DiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorageType

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetDiskStorageType(v string)`

SetDiskStorageType sets DiskStorageType field to given value.

### HasDiskStorageType

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasDiskStorageType() bool`

HasDiskStorageType returns a boolean if a field has been set.

### GetApplianceUrl

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetNetworkServerId() string`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetNetworkServerIdOk() (*string, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetNetworkServerId(v string)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetNetworkServer

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetNetworkServer() ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetNetworkServerOk() (*ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetNetworkServer(v ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetSecurityMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetBackupMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetBackupMode() string`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetBackupModeOk() (*string, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetBackupMode(v string)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### GetReplicationMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### GetDnsIntegrationId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDnsIntegrationId() string`

GetDnsIntegrationId returns the DnsIntegrationId field if non-nil, zero value otherwise.

### GetDnsIntegrationIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDnsIntegrationIdOk() (*string, bool)`

GetDnsIntegrationIdOk returns a tuple with the DnsIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrationId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetDnsIntegrationId(v string)`

SetDnsIntegrationId sets DnsIntegrationId field to given value.

### HasDnsIntegrationId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasDnsIntegrationId() bool`

HasDnsIntegrationId returns a boolean if a field has been set.

### GetConfigCmdbId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetConfigCmdbId() string`

GetConfigCmdbId returns the ConfigCmdbId field if non-nil, zero value otherwise.

### GetConfigCmdbIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetConfigCmdbIdOk() (*string, bool)`

GetConfigCmdbIdOk returns a tuple with the ConfigCmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetConfigCmdbId(v string)`

SetConfigCmdbId sets ConfigCmdbId field to given value.

### HasConfigCmdbId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasConfigCmdbId() bool`

HasConfigCmdbId returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetConfigCmId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetConfigCmId() string`

GetConfigCmId returns the ConfigCmId field if non-nil, zero value otherwise.

### GetConfigCmIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetConfigCmIdOk() (*string, bool)`

GetConfigCmIdOk returns a tuple with the ConfigCmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetConfigCmId(v string)`

SetConfigCmId sets ConfigCmId field to given value.

### HasConfigCmId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasConfigCmId() bool`

HasConfigCmId returns a boolean if a field has been set.

### GetSecurityServer

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetServiceRegistryId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetServiceRegistryId() string`

GetServiceRegistryId returns the ServiceRegistryId field if non-nil, zero value otherwise.

### GetServiceRegistryIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetServiceRegistryIdOk() (*string, bool)`

GetServiceRegistryIdOk returns a tuple with the ServiceRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRegistryId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetServiceRegistryId(v string)`

SetServiceRegistryId sets ServiceRegistryId field to given value.

### HasServiceRegistryId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasServiceRegistryId() bool`

HasServiceRegistryId returns a boolean if a field has been set.

### GetKubeUrl

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetKubeUrl() string`

GetKubeUrl returns the KubeUrl field if non-nil, zero value otherwise.

### GetKubeUrlOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetKubeUrlOk() (*string, bool)`

GetKubeUrlOk returns a tuple with the KubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeUrl

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetKubeUrl(v string)`

SetKubeUrl sets KubeUrl field to given value.

### HasKubeUrl

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasKubeUrl() bool`

HasKubeUrl returns a boolean if a field has been set.

### GetApiVersion

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDatacenterId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetConfigCmdbDiscovery

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetConfigCmdbDiscovery() bool`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetConfigCmdbDiscoveryOk() (*bool, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetConfigCmdbDiscovery(v bool)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetDistributedWorkerId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDistributedWorkerId() string`

GetDistributedWorkerId returns the DistributedWorkerId field if non-nil, zero value otherwise.

### GetDistributedWorkerIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDistributedWorkerIdOk() (*string, bool)`

GetDistributedWorkerIdOk returns a tuple with the DistributedWorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedWorkerId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetDistributedWorkerId(v string)`

SetDistributedWorkerId sets DistributedWorkerId field to given value.

### HasDistributedWorkerId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasDistributedWorkerId() bool`

HasDistributedWorkerId returns a boolean if a field has been set.

### GetPasswordHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetEndpoint

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAccessKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetIsVpc

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetIsVpc() string`

GetIsVpc returns the IsVpc field if non-nil, zero value otherwise.

### GetIsVpcOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetIsVpcOk() (*string, bool)`

GetIsVpcOk returns a tuple with the IsVpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpc

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetIsVpc(v string)`

SetIsVpc sets IsVpc field to given value.

### HasIsVpc

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasIsVpc() bool`

HasIsVpc returns a boolean if a field has been set.

### GetVpc

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetImageStoreId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetImageStoreId() string`

GetImageStoreId returns the ImageStoreId field if non-nil, zero value otherwise.

### GetImageStoreIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetImageStoreIdOk() (*string, bool)`

GetImageStoreIdOk returns a tuple with the ImageStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStoreId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetImageStoreId(v string)`

SetImageStoreId sets ImageStoreId field to given value.

### HasImageStoreId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasImageStoreId() bool`

HasImageStoreId returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetCostingReport

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingReport() string`

GetCostingReport returns the CostingReport field if non-nil, zero value otherwise.

### GetCostingReportOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingReportOk() (*string, bool)`

GetCostingReportOk returns a tuple with the CostingReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReport

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCostingReport(v string)`

SetCostingReport sets CostingReport field to given value.

### HasCostingReport

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCostingReport() bool`

HasCostingReport returns a boolean if a field has been set.

### GetCostingFolder

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingFolder() string`

GetCostingFolder returns the CostingFolder field if non-nil, zero value otherwise.

### GetCostingFolderOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingFolderOk() (*string, bool)`

GetCostingFolderOk returns a tuple with the CostingFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingFolder

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCostingFolder(v string)`

SetCostingFolder sets CostingFolder field to given value.

### HasCostingFolder

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCostingFolder() bool`

HasCostingFolder returns a boolean if a field has been set.

### GetCostingBucket

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingBucket() string`

GetCostingBucket returns the CostingBucket field if non-nil, zero value otherwise.

### GetCostingBucketOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingBucketOk() (*string, bool)`

GetCostingBucketOk returns a tuple with the CostingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucket

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCostingBucket(v string)`

SetCostingBucket sets CostingBucket field to given value.

### HasCostingBucket

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCostingBucket() bool`

HasCostingBucket returns a boolean if a field has been set.

### GetCostingBucketName

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingBucketName() string`

GetCostingBucketName returns the CostingBucketName field if non-nil, zero value otherwise.

### GetCostingBucketNameOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingBucketNameOk() (*string, bool)`

GetCostingBucketNameOk returns a tuple with the CostingBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucketName

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCostingBucketName(v string)`

SetCostingBucketName sets CostingBucketName field to given value.

### HasCostingBucketName

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCostingBucketName() bool`

HasCostingBucketName returns a boolean if a field has been set.

### GetCostingRegion

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingRegion() string`

GetCostingRegion returns the CostingRegion field if non-nil, zero value otherwise.

### GetCostingRegionOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingRegionOk() (*string, bool)`

GetCostingRegionOk returns a tuple with the CostingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingRegion

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCostingRegion(v string)`

SetCostingRegion sets CostingRegion field to given value.

### HasCostingRegion

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCostingRegion() bool`

HasCostingRegion returns a boolean if a field has been set.

### GetCostingAccessKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingAccessKey() string`

GetCostingAccessKey returns the CostingAccessKey field if non-nil, zero value otherwise.

### GetCostingAccessKeyOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingAccessKeyOk() (*string, bool)`

GetCostingAccessKeyOk returns a tuple with the CostingAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingAccessKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCostingAccessKey(v string)`

SetCostingAccessKey sets CostingAccessKey field to given value.

### HasCostingAccessKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCostingAccessKey() bool`

HasCostingAccessKey returns a boolean if a field has been set.

### GetCostingSecretKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingSecretKey() string`

GetCostingSecretKey returns the CostingSecretKey field if non-nil, zero value otherwise.

### GetCostingSecretKeyOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingSecretKeyOk() (*string, bool)`

GetCostingSecretKeyOk returns a tuple with the CostingSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCostingSecretKey(v string)`

SetCostingSecretKey sets CostingSecretKey field to given value.

### HasCostingSecretKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCostingSecretKey() bool`

HasCostingSecretKey returns a boolean if a field has been set.

### GetCostingReportName

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingReportName() string`

GetCostingReportName returns the CostingReportName field if non-nil, zero value otherwise.

### GetCostingReportNameOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingReportNameOk() (*string, bool)`

GetCostingReportNameOk returns a tuple with the CostingReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReportName

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCostingReportName(v string)`

SetCostingReportName sets CostingReportName field to given value.

### HasCostingReportName

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCostingReportName() bool`

HasCostingReportName returns a boolean if a field has been set.

### GetSecretKeyHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetSecretKeyHash() string`

GetSecretKeyHash returns the SecretKeyHash field if non-nil, zero value otherwise.

### GetSecretKeyHashOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetSecretKeyHashOk() (*string, bool)`

GetSecretKeyHashOk returns a tuple with the SecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetSecretKeyHash(v string)`

SetSecretKeyHash sets SecretKeyHash field to given value.

### HasSecretKeyHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasSecretKeyHash() bool`

HasSecretKeyHash returns a boolean if a field has been set.

### GetCostingSecretKeyHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingSecretKeyHash() string`

GetCostingSecretKeyHash returns the CostingSecretKeyHash field if non-nil, zero value otherwise.

### GetCostingSecretKeyHashOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCostingSecretKeyHashOk() (*string, bool)`

GetCostingSecretKeyHashOk returns a tuple with the CostingSecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKeyHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCostingSecretKeyHash(v string)`

SetCostingSecretKeyHash sets CostingSecretKeyHash field to given value.

### HasCostingSecretKeyHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCostingSecretKeyHash() bool`

HasCostingSecretKeyHash returns a boolean if a field has been set.

### GetSubscriberId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetSubscriberId() string`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetSubscriberIdOk() (*string, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetSubscriberId(v string)`

SetSubscriberId sets SubscriberId field to given value.

### HasSubscriberId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasSubscriberId() bool`

HasSubscriberId returns a boolean if a field has been set.

### GetTenantId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetClientId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetResourceGroup

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetAccountType

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetCloudType

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetDiskEncryption

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDiskEncryption() string`

GetDiskEncryption returns the DiskEncryption field if non-nil, zero value otherwise.

### GetDiskEncryptionOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetDiskEncryptionOk() (*string, bool)`

GetDiskEncryptionOk returns a tuple with the DiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryption

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetDiskEncryption(v string)`

SetDiskEncryption sets DiskEncryption field to given value.

### HasDiskEncryption

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasDiskEncryption() bool`

HasDiskEncryption returns a boolean if a field has been set.

### GetEncryptionSet

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetEncryptionSet() string`

GetEncryptionSet returns the EncryptionSet field if non-nil, zero value otherwise.

### GetEncryptionSetOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetEncryptionSetOk() (*string, bool)`

GetEncryptionSetOk returns a tuple with the EncryptionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSet

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetEncryptionSet(v string)`

SetEncryptionSet sets EncryptionSet field to given value.

### HasEncryptionSet

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasEncryptionSet() bool`

HasEncryptionSet returns a boolean if a field has been set.

### GetCspTenantId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCspTenantId() string`

GetCspTenantId returns the CspTenantId field if non-nil, zero value otherwise.

### GetCspTenantIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCspTenantIdOk() (*string, bool)`

GetCspTenantIdOk returns a tuple with the CspTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspTenantId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCspTenantId(v string)`

SetCspTenantId sets CspTenantId field to given value.

### HasCspTenantId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCspTenantId() bool`

HasCspTenantId returns a boolean if a field has been set.

### GetCspClientId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCspClientId() string`

GetCspClientId returns the CspClientId field if non-nil, zero value otherwise.

### GetCspClientIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCspClientIdOk() (*string, bool)`

GetCspClientIdOk returns a tuple with the CspClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspClientId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCspClientId(v string)`

SetCspClientId sets CspClientId field to given value.

### HasCspClientId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCspClientId() bool`

HasCspClientId returns a boolean if a field has been set.

### GetCspClientSecret

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCspClientSecret() string`

GetCspClientSecret returns the CspClientSecret field if non-nil, zero value otherwise.

### GetCspClientSecretOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCspClientSecretOk() (*string, bool)`

GetCspClientSecretOk returns a tuple with the CspClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspClientSecret

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCspClientSecret(v string)`

SetCspClientSecret sets CspClientSecret field to given value.

### HasCspClientSecret

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCspClientSecret() bool`

HasCspClientSecret returns a boolean if a field has been set.

### GetCspCustomer

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### GetAzureCostingMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetAzureCostingMode() string`

GetAzureCostingMode returns the AzureCostingMode field if non-nil, zero value otherwise.

### GetAzureCostingModeOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetAzureCostingModeOk() (*string, bool)`

GetAzureCostingModeOk returns a tuple with the AzureCostingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCostingMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetAzureCostingMode(v string)`

SetAzureCostingMode sets AzureCostingMode field to given value.

### HasAzureCostingMode

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasAzureCostingMode() bool`

HasAzureCostingMode returns a boolean if a field has been set.

### GetClientSecretHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetClientSecretHash() string`

GetClientSecretHash returns the ClientSecretHash field if non-nil, zero value otherwise.

### GetClientSecretHashOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetClientSecretHashOk() (*string, bool)`

GetClientSecretHashOk returns a tuple with the ClientSecretHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetClientSecretHash(v string)`

SetClientSecretHash sets ClientSecretHash field to given value.

### HasClientSecretHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasClientSecretHash() bool`

HasClientSecretHash returns a boolean if a field has been set.

### GetCspClientSecretHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCspClientSecretHash() string`

GetCspClientSecretHash returns the CspClientSecretHash field if non-nil, zero value otherwise.

### GetCspClientSecretHashOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetCspClientSecretHashOk() (*string, bool)`

GetCspClientSecretHashOk returns a tuple with the CspClientSecretHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspClientSecretHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetCspClientSecretHash(v string)`

SetCspClientSecretHash sets CspClientSecretHash field to given value.

### HasCspClientSecretHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasCspClientSecretHash() bool`

HasCspClientSecretHash returns a boolean if a field has been set.

### GetPrivateKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetClientEmail

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetClientEmail() string`

GetClientEmail returns the ClientEmail field if non-nil, zero value otherwise.

### GetClientEmailOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetClientEmailOk() (*string, bool)`

GetClientEmailOk returns a tuple with the ClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmail

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetClientEmail(v string)`

SetClientEmail sets ClientEmail field to given value.

### HasClientEmail

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasClientEmail() bool`

HasClientEmail returns a boolean if a field has been set.

### GetProjectId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetGoogleRegionId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetGoogleRegionId() string`

GetGoogleRegionId returns the GoogleRegionId field if non-nil, zero value otherwise.

### GetGoogleRegionIdOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetGoogleRegionIdOk() (*string, bool)`

GetGoogleRegionIdOk returns a tuple with the GoogleRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleRegionId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetGoogleRegionId(v string)`

SetGoogleRegionId sets GoogleRegionId field to given value.

### HasGoogleRegionId

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasGoogleRegionId() bool`

HasGoogleRegionId returns a boolean if a field has been set.

### GetPrivateKeyHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetPrivateKeyHash() string`

GetPrivateKeyHash returns the PrivateKeyHash field if non-nil, zero value otherwise.

### GetPrivateKeyHashOk

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) GetPrivateKeyHashOk() (*string, bool)`

GetPrivateKeyHashOk returns a tuple with the PrivateKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) SetPrivateKeyHash(v string)`

SetPrivateKeyHash sets PrivateKeyHash field to given value.

### HasPrivateKeyHash

`func (o *ListClouds200ResponseAllOfZonesInnerConfig) HasPrivateKeyHash() bool`

HasPrivateKeyHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AddClouds200ResponseAllOfZoneConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | Pointer to **string** |  | [optional] 
**DatacenterName** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**InventoryLevel** | Pointer to **string** |  | [optional] 
**ConsoleKeymap** | Pointer to **string** |  | [optional] 
**BackupMode** | Pointer to **string** |  | [optional] 
**CertificateProvider** | Pointer to **string** |  | [optional] 
**ConfigCmdbDiscovery** | Pointer to **bool** |  | [optional] 
**EnableNetworkTypeSelection** | Pointer to **string** |  | [optional] 
**KubeUrl** | Pointer to **string** |  | [optional] 
**NetworkServer** | Pointer to [**AddClouds200ResponseAllOfZoneConfigAnyOf3NetworkServer**](AddClouds200ResponseAllOfZoneConfigAnyOf3NetworkServer.md) |  | [optional] 
**NetworkServerId** | Pointer to **string** |  | [optional] 
**ReplicationMode** | Pointer to **string** |  | [optional] 
**SecurityServer** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**ClusterRef** | Pointer to **string** |  | [optional] 
**ConfigCmId** | Pointer to **string** |  | [optional] 
**ConfigCmdbId** | Pointer to **string** |  | [optional] 
**ConfigManagementId** | Pointer to **string** |  | [optional] 
**Datacenter** | Pointer to **string** |  | [optional] 
**DatacenterId** | Pointer to **string** |  | [optional] 
**DiskStorageType** | Pointer to **string** |  | [optional] 
**DistributedWorkerId** | Pointer to **string** |  | [optional] 
**DnsIntegrationId** | Pointer to **string** |  | [optional] 
**EnableDiskTypeSelection** | Pointer to **string** |  | [optional] 
**EnableStorageTypeSelection** | Pointer to **string** |  | [optional] 
**EnableVnc** | Pointer to **string** |  | [optional] 
**HideHostSelection** | Pointer to **string** |  | [optional] 
**ImportExisting** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordHash** | Pointer to **string** |  | [optional] 
**ResourcePool** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **string** |  | [optional] 
**RpcMode** | Pointer to **string** |  | [optional] 
**SecurityMode** | Pointer to **string** |  | [optional] 
**ServiceRegistryId** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
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
**AccountType** | Pointer to **string** |  | [optional] 
**CloudType** | Pointer to **string** |  | [optional] 
**StorageAccount** | Pointer to **string** |  | [optional] 
**DiskEncryption** | Pointer to **string** |  | [optional] 
**EncryptionSet** | Pointer to **string** |  | [optional] 
**CspTenantId** | Pointer to **string** |  | [optional] 
**CspClientId** | Pointer to **string** |  | [optional] 
**CspClientSecret** | Pointer to **string** |  | [optional] 
**CspCustomer** | Pointer to **string** |  | [optional] 
**AzureCostingMode** | Pointer to **string** |  | [optional] 
**ClientSecretHash** | Pointer to **string** |  | [optional] 
**CspClientSecretHash** | Pointer to **string** |  | [optional] 

## Methods

### NewAddClouds200ResponseAllOfZoneConfig

`func NewAddClouds200ResponseAllOfZoneConfig() *AddClouds200ResponseAllOfZoneConfig`

NewAddClouds200ResponseAllOfZoneConfig instantiates a new AddClouds200ResponseAllOfZoneConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetApplianceUrl

`func (o *AddClouds200ResponseAllOfZoneConfig) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *AddClouds200ResponseAllOfZoneConfig) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *AddClouds200ResponseAllOfZoneConfig) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *AddClouds200ResponseAllOfZoneConfig) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *AddClouds200ResponseAllOfZoneConfig) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *AddClouds200ResponseAllOfZoneConfig) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *AddClouds200ResponseAllOfZoneConfig) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *AddClouds200ResponseAllOfZoneConfig) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *AddClouds200ResponseAllOfZoneConfig) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *AddClouds200ResponseAllOfZoneConfig) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *AddClouds200ResponseAllOfZoneConfig) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetBackupMode

`func (o *AddClouds200ResponseAllOfZoneConfig) GetBackupMode() string`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetBackupModeOk() (*string, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *AddClouds200ResponseAllOfZoneConfig) SetBackupMode(v string)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *AddClouds200ResponseAllOfZoneConfig) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetConfigCmdbDiscovery

`func (o *AddClouds200ResponseAllOfZoneConfig) GetConfigCmdbDiscovery() bool`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetConfigCmdbDiscoveryOk() (*bool, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *AddClouds200ResponseAllOfZoneConfig) SetConfigCmdbDiscovery(v bool)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *AddClouds200ResponseAllOfZoneConfig) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *AddClouds200ResponseAllOfZoneConfig) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *AddClouds200ResponseAllOfZoneConfig) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### GetKubeUrl

`func (o *AddClouds200ResponseAllOfZoneConfig) GetKubeUrl() string`

GetKubeUrl returns the KubeUrl field if non-nil, zero value otherwise.

### GetKubeUrlOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetKubeUrlOk() (*string, bool)`

GetKubeUrlOk returns a tuple with the KubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeUrl

`func (o *AddClouds200ResponseAllOfZoneConfig) SetKubeUrl(v string)`

SetKubeUrl sets KubeUrl field to given value.

### HasKubeUrl

`func (o *AddClouds200ResponseAllOfZoneConfig) HasKubeUrl() bool`

HasKubeUrl returns a boolean if a field has been set.

### GetNetworkServer

`func (o *AddClouds200ResponseAllOfZoneConfig) GetNetworkServer() AddClouds200ResponseAllOfZoneConfigAnyOf3NetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetNetworkServerOk() (*AddClouds200ResponseAllOfZoneConfigAnyOf3NetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *AddClouds200ResponseAllOfZoneConfig) SetNetworkServer(v AddClouds200ResponseAllOfZoneConfigAnyOf3NetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *AddClouds200ResponseAllOfZoneConfig) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetNetworkServerId() string`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetNetworkServerIdOk() (*string, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetNetworkServerId(v string)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetReplicationMode

`func (o *AddClouds200ResponseAllOfZoneConfig) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *AddClouds200ResponseAllOfZoneConfig) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *AddClouds200ResponseAllOfZoneConfig) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### GetSecurityServer

`func (o *AddClouds200ResponseAllOfZoneConfig) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *AddClouds200ResponseAllOfZoneConfig) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *AddClouds200ResponseAllOfZoneConfig) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetApiUrl

`func (o *AddClouds200ResponseAllOfZoneConfig) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *AddClouds200ResponseAllOfZoneConfig) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *AddClouds200ResponseAllOfZoneConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetApiVersion

`func (o *AddClouds200ResponseAllOfZoneConfig) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AddClouds200ResponseAllOfZoneConfig) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AddClouds200ResponseAllOfZoneConfig) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCluster

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterRef

`func (o *AddClouds200ResponseAllOfZoneConfig) GetClusterRef() string`

GetClusterRef returns the ClusterRef field if non-nil, zero value otherwise.

### GetClusterRefOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetClusterRefOk() (*string, bool)`

GetClusterRefOk returns a tuple with the ClusterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRef

`func (o *AddClouds200ResponseAllOfZoneConfig) SetClusterRef(v string)`

SetClusterRef sets ClusterRef field to given value.

### HasClusterRef

`func (o *AddClouds200ResponseAllOfZoneConfig) HasClusterRef() bool`

HasClusterRef returns a boolean if a field has been set.

### GetConfigCmId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetConfigCmId() string`

GetConfigCmId returns the ConfigCmId field if non-nil, zero value otherwise.

### GetConfigCmIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetConfigCmIdOk() (*string, bool)`

GetConfigCmIdOk returns a tuple with the ConfigCmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetConfigCmId(v string)`

SetConfigCmId sets ConfigCmId field to given value.

### HasConfigCmId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasConfigCmId() bool`

HasConfigCmId returns a boolean if a field has been set.

### GetConfigCmdbId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetConfigCmdbId() string`

GetConfigCmdbId returns the ConfigCmdbId field if non-nil, zero value otherwise.

### GetConfigCmdbIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetConfigCmdbIdOk() (*string, bool)`

GetConfigCmdbIdOk returns a tuple with the ConfigCmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetConfigCmdbId(v string)`

SetConfigCmdbId sets ConfigCmdbId field to given value.

### HasConfigCmdbId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasConfigCmdbId() bool`

HasConfigCmdbId returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetDatacenter

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *AddClouds200ResponseAllOfZoneConfig) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *AddClouds200ResponseAllOfZoneConfig) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetDatacenterId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetDiskStorageType

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDiskStorageType() string`

GetDiskStorageType returns the DiskStorageType field if non-nil, zero value otherwise.

### GetDiskStorageTypeOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDiskStorageTypeOk() (*string, bool)`

GetDiskStorageTypeOk returns a tuple with the DiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorageType

`func (o *AddClouds200ResponseAllOfZoneConfig) SetDiskStorageType(v string)`

SetDiskStorageType sets DiskStorageType field to given value.

### HasDiskStorageType

`func (o *AddClouds200ResponseAllOfZoneConfig) HasDiskStorageType() bool`

HasDiskStorageType returns a boolean if a field has been set.

### GetDistributedWorkerId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDistributedWorkerId() string`

GetDistributedWorkerId returns the DistributedWorkerId field if non-nil, zero value otherwise.

### GetDistributedWorkerIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDistributedWorkerIdOk() (*string, bool)`

GetDistributedWorkerIdOk returns a tuple with the DistributedWorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedWorkerId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetDistributedWorkerId(v string)`

SetDistributedWorkerId sets DistributedWorkerId field to given value.

### HasDistributedWorkerId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasDistributedWorkerId() bool`

HasDistributedWorkerId returns a boolean if a field has been set.

### GetDnsIntegrationId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDnsIntegrationId() string`

GetDnsIntegrationId returns the DnsIntegrationId field if non-nil, zero value otherwise.

### GetDnsIntegrationIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDnsIntegrationIdOk() (*string, bool)`

GetDnsIntegrationIdOk returns a tuple with the DnsIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrationId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetDnsIntegrationId(v string)`

SetDnsIntegrationId sets DnsIntegrationId field to given value.

### HasDnsIntegrationId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasDnsIntegrationId() bool`

HasDnsIntegrationId returns a boolean if a field has been set.

### GetEnableDiskTypeSelection

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEnableDiskTypeSelection() string`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEnableDiskTypeSelectionOk() (*string, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *AddClouds200ResponseAllOfZoneConfig) SetEnableDiskTypeSelection(v string)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *AddClouds200ResponseAllOfZoneConfig) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### GetEnableStorageTypeSelection

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEnableStorageTypeSelection() string`

GetEnableStorageTypeSelection returns the EnableStorageTypeSelection field if non-nil, zero value otherwise.

### GetEnableStorageTypeSelectionOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEnableStorageTypeSelectionOk() (*string, bool)`

GetEnableStorageTypeSelectionOk returns a tuple with the EnableStorageTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStorageTypeSelection

`func (o *AddClouds200ResponseAllOfZoneConfig) SetEnableStorageTypeSelection(v string)`

SetEnableStorageTypeSelection sets EnableStorageTypeSelection field to given value.

### HasEnableStorageTypeSelection

`func (o *AddClouds200ResponseAllOfZoneConfig) HasEnableStorageTypeSelection() bool`

HasEnableStorageTypeSelection returns a boolean if a field has been set.

### GetEnableVnc

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEnableVnc() string`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEnableVncOk() (*string, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *AddClouds200ResponseAllOfZoneConfig) SetEnableVnc(v string)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *AddClouds200ResponseAllOfZoneConfig) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### GetHideHostSelection

`func (o *AddClouds200ResponseAllOfZoneConfig) GetHideHostSelection() string`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetHideHostSelectionOk() (*string, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *AddClouds200ResponseAllOfZoneConfig) SetHideHostSelection(v string)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *AddClouds200ResponseAllOfZoneConfig) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### GetImportExisting

`func (o *AddClouds200ResponseAllOfZoneConfig) GetImportExisting() string`

GetImportExisting returns the ImportExisting field if non-nil, zero value otherwise.

### GetImportExistingOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetImportExistingOk() (*string, bool)`

GetImportExistingOk returns a tuple with the ImportExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportExisting

`func (o *AddClouds200ResponseAllOfZoneConfig) SetImportExisting(v string)`

SetImportExisting sets ImportExisting field to given value.

### HasImportExisting

`func (o *AddClouds200ResponseAllOfZoneConfig) HasImportExisting() bool`

HasImportExisting returns a boolean if a field has been set.

### GetPassword

`func (o *AddClouds200ResponseAllOfZoneConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddClouds200ResponseAllOfZoneConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddClouds200ResponseAllOfZoneConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *AddClouds200ResponseAllOfZoneConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *AddClouds200ResponseAllOfZoneConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *AddClouds200ResponseAllOfZoneConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetResourcePool

`func (o *AddClouds200ResponseAllOfZoneConfig) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *AddClouds200ResponseAllOfZoneConfig) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *AddClouds200ResponseAllOfZoneConfig) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetRpcMode

`func (o *AddClouds200ResponseAllOfZoneConfig) GetRpcMode() string`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetRpcModeOk() (*string, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *AddClouds200ResponseAllOfZoneConfig) SetRpcMode(v string)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *AddClouds200ResponseAllOfZoneConfig) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### GetSecurityMode

`func (o *AddClouds200ResponseAllOfZoneConfig) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *AddClouds200ResponseAllOfZoneConfig) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *AddClouds200ResponseAllOfZoneConfig) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetServiceRegistryId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetServiceRegistryId() string`

GetServiceRegistryId returns the ServiceRegistryId field if non-nil, zero value otherwise.

### GetServiceRegistryIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetServiceRegistryIdOk() (*string, bool)`

GetServiceRegistryIdOk returns a tuple with the ServiceRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRegistryId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetServiceRegistryId(v string)`

SetServiceRegistryId sets ServiceRegistryId field to given value.

### HasServiceRegistryId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasServiceRegistryId() bool`

HasServiceRegistryId returns a boolean if a field has been set.

### GetUsername

`func (o *AddClouds200ResponseAllOfZoneConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddClouds200ResponseAllOfZoneConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddClouds200ResponseAllOfZoneConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEndpoint

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AddClouds200ResponseAllOfZoneConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AddClouds200ResponseAllOfZoneConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAccessKey

`func (o *AddClouds200ResponseAllOfZoneConfig) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AddClouds200ResponseAllOfZoneConfig) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *AddClouds200ResponseAllOfZoneConfig) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *AddClouds200ResponseAllOfZoneConfig) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AddClouds200ResponseAllOfZoneConfig) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AddClouds200ResponseAllOfZoneConfig) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *AddClouds200ResponseAllOfZoneConfig) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *AddClouds200ResponseAllOfZoneConfig) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *AddClouds200ResponseAllOfZoneConfig) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *AddClouds200ResponseAllOfZoneConfig) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *AddClouds200ResponseAllOfZoneConfig) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *AddClouds200ResponseAllOfZoneConfig) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetIsVpc

`func (o *AddClouds200ResponseAllOfZoneConfig) GetIsVpc() string`

GetIsVpc returns the IsVpc field if non-nil, zero value otherwise.

### GetIsVpcOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetIsVpcOk() (*string, bool)`

GetIsVpcOk returns a tuple with the IsVpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpc

`func (o *AddClouds200ResponseAllOfZoneConfig) SetIsVpc(v string)`

SetIsVpc sets IsVpc field to given value.

### HasIsVpc

`func (o *AddClouds200ResponseAllOfZoneConfig) HasIsVpc() bool`

HasIsVpc returns a boolean if a field has been set.

### GetVpc

`func (o *AddClouds200ResponseAllOfZoneConfig) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *AddClouds200ResponseAllOfZoneConfig) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *AddClouds200ResponseAllOfZoneConfig) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetImageStoreId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetImageStoreId() string`

GetImageStoreId returns the ImageStoreId field if non-nil, zero value otherwise.

### GetImageStoreIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetImageStoreIdOk() (*string, bool)`

GetImageStoreIdOk returns a tuple with the ImageStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStoreId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetImageStoreId(v string)`

SetImageStoreId sets ImageStoreId field to given value.

### HasImageStoreId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasImageStoreId() bool`

HasImageStoreId returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *AddClouds200ResponseAllOfZoneConfig) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *AddClouds200ResponseAllOfZoneConfig) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetCostingReport

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingReport() string`

GetCostingReport returns the CostingReport field if non-nil, zero value otherwise.

### GetCostingReportOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingReportOk() (*string, bool)`

GetCostingReportOk returns a tuple with the CostingReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReport

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCostingReport(v string)`

SetCostingReport sets CostingReport field to given value.

### HasCostingReport

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCostingReport() bool`

HasCostingReport returns a boolean if a field has been set.

### GetCostingFolder

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingFolder() string`

GetCostingFolder returns the CostingFolder field if non-nil, zero value otherwise.

### GetCostingFolderOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingFolderOk() (*string, bool)`

GetCostingFolderOk returns a tuple with the CostingFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingFolder

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCostingFolder(v string)`

SetCostingFolder sets CostingFolder field to given value.

### HasCostingFolder

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCostingFolder() bool`

HasCostingFolder returns a boolean if a field has been set.

### GetCostingBucket

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingBucket() string`

GetCostingBucket returns the CostingBucket field if non-nil, zero value otherwise.

### GetCostingBucketOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingBucketOk() (*string, bool)`

GetCostingBucketOk returns a tuple with the CostingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucket

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCostingBucket(v string)`

SetCostingBucket sets CostingBucket field to given value.

### HasCostingBucket

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCostingBucket() bool`

HasCostingBucket returns a boolean if a field has been set.

### GetCostingBucketName

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingBucketName() string`

GetCostingBucketName returns the CostingBucketName field if non-nil, zero value otherwise.

### GetCostingBucketNameOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingBucketNameOk() (*string, bool)`

GetCostingBucketNameOk returns a tuple with the CostingBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucketName

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCostingBucketName(v string)`

SetCostingBucketName sets CostingBucketName field to given value.

### HasCostingBucketName

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCostingBucketName() bool`

HasCostingBucketName returns a boolean if a field has been set.

### GetCostingRegion

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingRegion() string`

GetCostingRegion returns the CostingRegion field if non-nil, zero value otherwise.

### GetCostingRegionOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingRegionOk() (*string, bool)`

GetCostingRegionOk returns a tuple with the CostingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingRegion

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCostingRegion(v string)`

SetCostingRegion sets CostingRegion field to given value.

### HasCostingRegion

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCostingRegion() bool`

HasCostingRegion returns a boolean if a field has been set.

### GetCostingAccessKey

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingAccessKey() string`

GetCostingAccessKey returns the CostingAccessKey field if non-nil, zero value otherwise.

### GetCostingAccessKeyOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingAccessKeyOk() (*string, bool)`

GetCostingAccessKeyOk returns a tuple with the CostingAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingAccessKey

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCostingAccessKey(v string)`

SetCostingAccessKey sets CostingAccessKey field to given value.

### HasCostingAccessKey

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCostingAccessKey() bool`

HasCostingAccessKey returns a boolean if a field has been set.

### GetCostingSecretKey

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingSecretKey() string`

GetCostingSecretKey returns the CostingSecretKey field if non-nil, zero value otherwise.

### GetCostingSecretKeyOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingSecretKeyOk() (*string, bool)`

GetCostingSecretKeyOk returns a tuple with the CostingSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKey

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCostingSecretKey(v string)`

SetCostingSecretKey sets CostingSecretKey field to given value.

### HasCostingSecretKey

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCostingSecretKey() bool`

HasCostingSecretKey returns a boolean if a field has been set.

### GetCostingReportName

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingReportName() string`

GetCostingReportName returns the CostingReportName field if non-nil, zero value otherwise.

### GetCostingReportNameOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingReportNameOk() (*string, bool)`

GetCostingReportNameOk returns a tuple with the CostingReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReportName

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCostingReportName(v string)`

SetCostingReportName sets CostingReportName field to given value.

### HasCostingReportName

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCostingReportName() bool`

HasCostingReportName returns a boolean if a field has been set.

### GetSecretKeyHash

`func (o *AddClouds200ResponseAllOfZoneConfig) GetSecretKeyHash() string`

GetSecretKeyHash returns the SecretKeyHash field if non-nil, zero value otherwise.

### GetSecretKeyHashOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetSecretKeyHashOk() (*string, bool)`

GetSecretKeyHashOk returns a tuple with the SecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyHash

`func (o *AddClouds200ResponseAllOfZoneConfig) SetSecretKeyHash(v string)`

SetSecretKeyHash sets SecretKeyHash field to given value.

### HasSecretKeyHash

`func (o *AddClouds200ResponseAllOfZoneConfig) HasSecretKeyHash() bool`

HasSecretKeyHash returns a boolean if a field has been set.

### GetCostingSecretKeyHash

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingSecretKeyHash() string`

GetCostingSecretKeyHash returns the CostingSecretKeyHash field if non-nil, zero value otherwise.

### GetCostingSecretKeyHashOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCostingSecretKeyHashOk() (*string, bool)`

GetCostingSecretKeyHashOk returns a tuple with the CostingSecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKeyHash

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCostingSecretKeyHash(v string)`

SetCostingSecretKeyHash sets CostingSecretKeyHash field to given value.

### HasCostingSecretKeyHash

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCostingSecretKeyHash() bool`

HasCostingSecretKeyHash returns a boolean if a field has been set.

### GetSubscriberId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetSubscriberId() string`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetSubscriberIdOk() (*string, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetSubscriberId(v string)`

SetSubscriberId sets SubscriberId field to given value.

### HasSubscriberId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasSubscriberId() bool`

HasSubscriberId returns a boolean if a field has been set.

### GetTenantId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetClientId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AddClouds200ResponseAllOfZoneConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AddClouds200ResponseAllOfZoneConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AddClouds200ResponseAllOfZoneConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetResourceGroup

`func (o *AddClouds200ResponseAllOfZoneConfig) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AddClouds200ResponseAllOfZoneConfig) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AddClouds200ResponseAllOfZoneConfig) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetAccountType

`func (o *AddClouds200ResponseAllOfZoneConfig) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AddClouds200ResponseAllOfZoneConfig) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AddClouds200ResponseAllOfZoneConfig) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetCloudType

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetStorageAccount

`func (o *AddClouds200ResponseAllOfZoneConfig) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *AddClouds200ResponseAllOfZoneConfig) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *AddClouds200ResponseAllOfZoneConfig) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### GetDiskEncryption

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDiskEncryption() string`

GetDiskEncryption returns the DiskEncryption field if non-nil, zero value otherwise.

### GetDiskEncryptionOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetDiskEncryptionOk() (*string, bool)`

GetDiskEncryptionOk returns a tuple with the DiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryption

`func (o *AddClouds200ResponseAllOfZoneConfig) SetDiskEncryption(v string)`

SetDiskEncryption sets DiskEncryption field to given value.

### HasDiskEncryption

`func (o *AddClouds200ResponseAllOfZoneConfig) HasDiskEncryption() bool`

HasDiskEncryption returns a boolean if a field has been set.

### GetEncryptionSet

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEncryptionSet() string`

GetEncryptionSet returns the EncryptionSet field if non-nil, zero value otherwise.

### GetEncryptionSetOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetEncryptionSetOk() (*string, bool)`

GetEncryptionSetOk returns a tuple with the EncryptionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSet

`func (o *AddClouds200ResponseAllOfZoneConfig) SetEncryptionSet(v string)`

SetEncryptionSet sets EncryptionSet field to given value.

### HasEncryptionSet

`func (o *AddClouds200ResponseAllOfZoneConfig) HasEncryptionSet() bool`

HasEncryptionSet returns a boolean if a field has been set.

### GetCspTenantId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCspTenantId() string`

GetCspTenantId returns the CspTenantId field if non-nil, zero value otherwise.

### GetCspTenantIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCspTenantIdOk() (*string, bool)`

GetCspTenantIdOk returns a tuple with the CspTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspTenantId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCspTenantId(v string)`

SetCspTenantId sets CspTenantId field to given value.

### HasCspTenantId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCspTenantId() bool`

HasCspTenantId returns a boolean if a field has been set.

### GetCspClientId

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCspClientId() string`

GetCspClientId returns the CspClientId field if non-nil, zero value otherwise.

### GetCspClientIdOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCspClientIdOk() (*string, bool)`

GetCspClientIdOk returns a tuple with the CspClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspClientId

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCspClientId(v string)`

SetCspClientId sets CspClientId field to given value.

### HasCspClientId

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCspClientId() bool`

HasCspClientId returns a boolean if a field has been set.

### GetCspClientSecret

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCspClientSecret() string`

GetCspClientSecret returns the CspClientSecret field if non-nil, zero value otherwise.

### GetCspClientSecretOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCspClientSecretOk() (*string, bool)`

GetCspClientSecretOk returns a tuple with the CspClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspClientSecret

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCspClientSecret(v string)`

SetCspClientSecret sets CspClientSecret field to given value.

### HasCspClientSecret

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCspClientSecret() bool`

HasCspClientSecret returns a boolean if a field has been set.

### GetCspCustomer

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### GetAzureCostingMode

`func (o *AddClouds200ResponseAllOfZoneConfig) GetAzureCostingMode() string`

GetAzureCostingMode returns the AzureCostingMode field if non-nil, zero value otherwise.

### GetAzureCostingModeOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetAzureCostingModeOk() (*string, bool)`

GetAzureCostingModeOk returns a tuple with the AzureCostingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCostingMode

`func (o *AddClouds200ResponseAllOfZoneConfig) SetAzureCostingMode(v string)`

SetAzureCostingMode sets AzureCostingMode field to given value.

### HasAzureCostingMode

`func (o *AddClouds200ResponseAllOfZoneConfig) HasAzureCostingMode() bool`

HasAzureCostingMode returns a boolean if a field has been set.

### GetClientSecretHash

`func (o *AddClouds200ResponseAllOfZoneConfig) GetClientSecretHash() string`

GetClientSecretHash returns the ClientSecretHash field if non-nil, zero value otherwise.

### GetClientSecretHashOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetClientSecretHashOk() (*string, bool)`

GetClientSecretHashOk returns a tuple with the ClientSecretHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretHash

`func (o *AddClouds200ResponseAllOfZoneConfig) SetClientSecretHash(v string)`

SetClientSecretHash sets ClientSecretHash field to given value.

### HasClientSecretHash

`func (o *AddClouds200ResponseAllOfZoneConfig) HasClientSecretHash() bool`

HasClientSecretHash returns a boolean if a field has been set.

### GetCspClientSecretHash

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCspClientSecretHash() string`

GetCspClientSecretHash returns the CspClientSecretHash field if non-nil, zero value otherwise.

### GetCspClientSecretHashOk

`func (o *AddClouds200ResponseAllOfZoneConfig) GetCspClientSecretHashOk() (*string, bool)`

GetCspClientSecretHashOk returns a tuple with the CspClientSecretHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspClientSecretHash

`func (o *AddClouds200ResponseAllOfZoneConfig) SetCspClientSecretHash(v string)`

SetCspClientSecretHash sets CspClientSecretHash field to given value.

### HasCspClientSecretHash

`func (o *AddClouds200ResponseAllOfZoneConfig) HasCspClientSecretHash() bool`

HasCspClientSecretHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



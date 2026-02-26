# ZoneConfig

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
**NetworkServer** | Pointer to [**AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer**](AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer.md) |  | [optional] 
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

## Methods

### NewZoneConfig

`func NewZoneConfig() *ZoneConfig`

NewZoneConfig instantiates a new ZoneConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneConfigWithDefaults

`func NewZoneConfigWithDefaults() *ZoneConfig`

NewZoneConfigWithDefaults instantiates a new ZoneConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *ZoneConfig) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ZoneConfig) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ZoneConfig) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ZoneConfig) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *ZoneConfig) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *ZoneConfig) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *ZoneConfig) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *ZoneConfig) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *ZoneConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ZoneConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ZoneConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ZoneConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *ZoneConfig) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *ZoneConfig) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *ZoneConfig) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *ZoneConfig) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *ZoneConfig) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *ZoneConfig) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *ZoneConfig) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *ZoneConfig) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetBackupMode

`func (o *ZoneConfig) GetBackupMode() string`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *ZoneConfig) GetBackupModeOk() (*string, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *ZoneConfig) SetBackupMode(v string)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *ZoneConfig) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *ZoneConfig) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *ZoneConfig) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *ZoneConfig) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *ZoneConfig) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetConfigCmdbDiscovery

`func (o *ZoneConfig) GetConfigCmdbDiscovery() bool`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *ZoneConfig) GetConfigCmdbDiscoveryOk() (*bool, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *ZoneConfig) SetConfigCmdbDiscovery(v bool)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *ZoneConfig) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *ZoneConfig) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *ZoneConfig) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *ZoneConfig) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *ZoneConfig) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### GetKubeUrl

`func (o *ZoneConfig) GetKubeUrl() string`

GetKubeUrl returns the KubeUrl field if non-nil, zero value otherwise.

### GetKubeUrlOk

`func (o *ZoneConfig) GetKubeUrlOk() (*string, bool)`

GetKubeUrlOk returns a tuple with the KubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeUrl

`func (o *ZoneConfig) SetKubeUrl(v string)`

SetKubeUrl sets KubeUrl field to given value.

### HasKubeUrl

`func (o *ZoneConfig) HasKubeUrl() bool`

HasKubeUrl returns a boolean if a field has been set.

### GetNetworkServer

`func (o *ZoneConfig) GetNetworkServer() AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *ZoneConfig) GetNetworkServerOk() (*AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *ZoneConfig) SetNetworkServer(v AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *ZoneConfig) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *ZoneConfig) GetNetworkServerId() string`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *ZoneConfig) GetNetworkServerIdOk() (*string, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *ZoneConfig) SetNetworkServerId(v string)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *ZoneConfig) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetReplicationMode

`func (o *ZoneConfig) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *ZoneConfig) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *ZoneConfig) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *ZoneConfig) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### GetSecurityServer

`func (o *ZoneConfig) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *ZoneConfig) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *ZoneConfig) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *ZoneConfig) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetApiUrl

`func (o *ZoneConfig) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ZoneConfig) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ZoneConfig) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ZoneConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetApiVersion

`func (o *ZoneConfig) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ZoneConfig) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ZoneConfig) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ZoneConfig) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCluster

`func (o *ZoneConfig) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ZoneConfig) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ZoneConfig) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ZoneConfig) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterRef

`func (o *ZoneConfig) GetClusterRef() string`

GetClusterRef returns the ClusterRef field if non-nil, zero value otherwise.

### GetClusterRefOk

`func (o *ZoneConfig) GetClusterRefOk() (*string, bool)`

GetClusterRefOk returns a tuple with the ClusterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRef

`func (o *ZoneConfig) SetClusterRef(v string)`

SetClusterRef sets ClusterRef field to given value.

### HasClusterRef

`func (o *ZoneConfig) HasClusterRef() bool`

HasClusterRef returns a boolean if a field has been set.

### GetConfigCmId

`func (o *ZoneConfig) GetConfigCmId() string`

GetConfigCmId returns the ConfigCmId field if non-nil, zero value otherwise.

### GetConfigCmIdOk

`func (o *ZoneConfig) GetConfigCmIdOk() (*string, bool)`

GetConfigCmIdOk returns a tuple with the ConfigCmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmId

`func (o *ZoneConfig) SetConfigCmId(v string)`

SetConfigCmId sets ConfigCmId field to given value.

### HasConfigCmId

`func (o *ZoneConfig) HasConfigCmId() bool`

HasConfigCmId returns a boolean if a field has been set.

### GetConfigCmdbId

`func (o *ZoneConfig) GetConfigCmdbId() string`

GetConfigCmdbId returns the ConfigCmdbId field if non-nil, zero value otherwise.

### GetConfigCmdbIdOk

`func (o *ZoneConfig) GetConfigCmdbIdOk() (*string, bool)`

GetConfigCmdbIdOk returns a tuple with the ConfigCmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbId

`func (o *ZoneConfig) SetConfigCmdbId(v string)`

SetConfigCmdbId sets ConfigCmdbId field to given value.

### HasConfigCmdbId

`func (o *ZoneConfig) HasConfigCmdbId() bool`

HasConfigCmdbId returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *ZoneConfig) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *ZoneConfig) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *ZoneConfig) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *ZoneConfig) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetDatacenter

`func (o *ZoneConfig) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *ZoneConfig) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *ZoneConfig) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *ZoneConfig) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetDatacenterId

`func (o *ZoneConfig) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *ZoneConfig) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *ZoneConfig) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *ZoneConfig) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetDiskStorageType

`func (o *ZoneConfig) GetDiskStorageType() string`

GetDiskStorageType returns the DiskStorageType field if non-nil, zero value otherwise.

### GetDiskStorageTypeOk

`func (o *ZoneConfig) GetDiskStorageTypeOk() (*string, bool)`

GetDiskStorageTypeOk returns a tuple with the DiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorageType

`func (o *ZoneConfig) SetDiskStorageType(v string)`

SetDiskStorageType sets DiskStorageType field to given value.

### HasDiskStorageType

`func (o *ZoneConfig) HasDiskStorageType() bool`

HasDiskStorageType returns a boolean if a field has been set.

### GetDistributedWorkerId

`func (o *ZoneConfig) GetDistributedWorkerId() string`

GetDistributedWorkerId returns the DistributedWorkerId field if non-nil, zero value otherwise.

### GetDistributedWorkerIdOk

`func (o *ZoneConfig) GetDistributedWorkerIdOk() (*string, bool)`

GetDistributedWorkerIdOk returns a tuple with the DistributedWorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedWorkerId

`func (o *ZoneConfig) SetDistributedWorkerId(v string)`

SetDistributedWorkerId sets DistributedWorkerId field to given value.

### HasDistributedWorkerId

`func (o *ZoneConfig) HasDistributedWorkerId() bool`

HasDistributedWorkerId returns a boolean if a field has been set.

### GetDnsIntegrationId

`func (o *ZoneConfig) GetDnsIntegrationId() string`

GetDnsIntegrationId returns the DnsIntegrationId field if non-nil, zero value otherwise.

### GetDnsIntegrationIdOk

`func (o *ZoneConfig) GetDnsIntegrationIdOk() (*string, bool)`

GetDnsIntegrationIdOk returns a tuple with the DnsIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrationId

`func (o *ZoneConfig) SetDnsIntegrationId(v string)`

SetDnsIntegrationId sets DnsIntegrationId field to given value.

### HasDnsIntegrationId

`func (o *ZoneConfig) HasDnsIntegrationId() bool`

HasDnsIntegrationId returns a boolean if a field has been set.

### GetEnableDiskTypeSelection

`func (o *ZoneConfig) GetEnableDiskTypeSelection() string`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *ZoneConfig) GetEnableDiskTypeSelectionOk() (*string, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *ZoneConfig) SetEnableDiskTypeSelection(v string)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *ZoneConfig) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### GetEnableStorageTypeSelection

`func (o *ZoneConfig) GetEnableStorageTypeSelection() string`

GetEnableStorageTypeSelection returns the EnableStorageTypeSelection field if non-nil, zero value otherwise.

### GetEnableStorageTypeSelectionOk

`func (o *ZoneConfig) GetEnableStorageTypeSelectionOk() (*string, bool)`

GetEnableStorageTypeSelectionOk returns a tuple with the EnableStorageTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStorageTypeSelection

`func (o *ZoneConfig) SetEnableStorageTypeSelection(v string)`

SetEnableStorageTypeSelection sets EnableStorageTypeSelection field to given value.

### HasEnableStorageTypeSelection

`func (o *ZoneConfig) HasEnableStorageTypeSelection() bool`

HasEnableStorageTypeSelection returns a boolean if a field has been set.

### GetEnableVnc

`func (o *ZoneConfig) GetEnableVnc() string`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *ZoneConfig) GetEnableVncOk() (*string, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *ZoneConfig) SetEnableVnc(v string)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *ZoneConfig) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### GetHideHostSelection

`func (o *ZoneConfig) GetHideHostSelection() string`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *ZoneConfig) GetHideHostSelectionOk() (*string, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *ZoneConfig) SetHideHostSelection(v string)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *ZoneConfig) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### GetImportExisting

`func (o *ZoneConfig) GetImportExisting() string`

GetImportExisting returns the ImportExisting field if non-nil, zero value otherwise.

### GetImportExistingOk

`func (o *ZoneConfig) GetImportExistingOk() (*string, bool)`

GetImportExistingOk returns a tuple with the ImportExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportExisting

`func (o *ZoneConfig) SetImportExisting(v string)`

SetImportExisting sets ImportExisting field to given value.

### HasImportExisting

`func (o *ZoneConfig) HasImportExisting() bool`

HasImportExisting returns a boolean if a field has been set.

### GetPassword

`func (o *ZoneConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ZoneConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ZoneConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ZoneConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *ZoneConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *ZoneConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *ZoneConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *ZoneConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetResourcePool

`func (o *ZoneConfig) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ZoneConfig) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ZoneConfig) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ZoneConfig) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ZoneConfig) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ZoneConfig) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ZoneConfig) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ZoneConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetRpcMode

`func (o *ZoneConfig) GetRpcMode() string`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *ZoneConfig) GetRpcModeOk() (*string, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *ZoneConfig) SetRpcMode(v string)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *ZoneConfig) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### GetSecurityMode

`func (o *ZoneConfig) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *ZoneConfig) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *ZoneConfig) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *ZoneConfig) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetServiceRegistryId

`func (o *ZoneConfig) GetServiceRegistryId() string`

GetServiceRegistryId returns the ServiceRegistryId field if non-nil, zero value otherwise.

### GetServiceRegistryIdOk

`func (o *ZoneConfig) GetServiceRegistryIdOk() (*string, bool)`

GetServiceRegistryIdOk returns a tuple with the ServiceRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRegistryId

`func (o *ZoneConfig) SetServiceRegistryId(v string)`

SetServiceRegistryId sets ServiceRegistryId field to given value.

### HasServiceRegistryId

`func (o *ZoneConfig) HasServiceRegistryId() bool`

HasServiceRegistryId returns a boolean if a field has been set.

### GetUsername

`func (o *ZoneConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ZoneConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ZoneConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ZoneConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEndpoint

`func (o *ZoneConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ZoneConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ZoneConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ZoneConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAccessKey

`func (o *ZoneConfig) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ZoneConfig) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ZoneConfig) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ZoneConfig) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *ZoneConfig) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ZoneConfig) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ZoneConfig) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ZoneConfig) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *ZoneConfig) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *ZoneConfig) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *ZoneConfig) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *ZoneConfig) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *ZoneConfig) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *ZoneConfig) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *ZoneConfig) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *ZoneConfig) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetIsVpc

`func (o *ZoneConfig) GetIsVpc() string`

GetIsVpc returns the IsVpc field if non-nil, zero value otherwise.

### GetIsVpcOk

`func (o *ZoneConfig) GetIsVpcOk() (*string, bool)`

GetIsVpcOk returns a tuple with the IsVpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpc

`func (o *ZoneConfig) SetIsVpc(v string)`

SetIsVpc sets IsVpc field to given value.

### HasIsVpc

`func (o *ZoneConfig) HasIsVpc() bool`

HasIsVpc returns a boolean if a field has been set.

### GetVpc

`func (o *ZoneConfig) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *ZoneConfig) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *ZoneConfig) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *ZoneConfig) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetImageStoreId

`func (o *ZoneConfig) GetImageStoreId() string`

GetImageStoreId returns the ImageStoreId field if non-nil, zero value otherwise.

### GetImageStoreIdOk

`func (o *ZoneConfig) GetImageStoreIdOk() (*string, bool)`

GetImageStoreIdOk returns a tuple with the ImageStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStoreId

`func (o *ZoneConfig) SetImageStoreId(v string)`

SetImageStoreId sets ImageStoreId field to given value.

### HasImageStoreId

`func (o *ZoneConfig) HasImageStoreId() bool`

HasImageStoreId returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *ZoneConfig) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *ZoneConfig) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *ZoneConfig) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *ZoneConfig) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetCostingReport

`func (o *ZoneConfig) GetCostingReport() string`

GetCostingReport returns the CostingReport field if non-nil, zero value otherwise.

### GetCostingReportOk

`func (o *ZoneConfig) GetCostingReportOk() (*string, bool)`

GetCostingReportOk returns a tuple with the CostingReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReport

`func (o *ZoneConfig) SetCostingReport(v string)`

SetCostingReport sets CostingReport field to given value.

### HasCostingReport

`func (o *ZoneConfig) HasCostingReport() bool`

HasCostingReport returns a boolean if a field has been set.

### GetCostingFolder

`func (o *ZoneConfig) GetCostingFolder() string`

GetCostingFolder returns the CostingFolder field if non-nil, zero value otherwise.

### GetCostingFolderOk

`func (o *ZoneConfig) GetCostingFolderOk() (*string, bool)`

GetCostingFolderOk returns a tuple with the CostingFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingFolder

`func (o *ZoneConfig) SetCostingFolder(v string)`

SetCostingFolder sets CostingFolder field to given value.

### HasCostingFolder

`func (o *ZoneConfig) HasCostingFolder() bool`

HasCostingFolder returns a boolean if a field has been set.

### GetCostingBucket

`func (o *ZoneConfig) GetCostingBucket() string`

GetCostingBucket returns the CostingBucket field if non-nil, zero value otherwise.

### GetCostingBucketOk

`func (o *ZoneConfig) GetCostingBucketOk() (*string, bool)`

GetCostingBucketOk returns a tuple with the CostingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucket

`func (o *ZoneConfig) SetCostingBucket(v string)`

SetCostingBucket sets CostingBucket field to given value.

### HasCostingBucket

`func (o *ZoneConfig) HasCostingBucket() bool`

HasCostingBucket returns a boolean if a field has been set.

### GetCostingBucketName

`func (o *ZoneConfig) GetCostingBucketName() string`

GetCostingBucketName returns the CostingBucketName field if non-nil, zero value otherwise.

### GetCostingBucketNameOk

`func (o *ZoneConfig) GetCostingBucketNameOk() (*string, bool)`

GetCostingBucketNameOk returns a tuple with the CostingBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucketName

`func (o *ZoneConfig) SetCostingBucketName(v string)`

SetCostingBucketName sets CostingBucketName field to given value.

### HasCostingBucketName

`func (o *ZoneConfig) HasCostingBucketName() bool`

HasCostingBucketName returns a boolean if a field has been set.

### GetCostingRegion

`func (o *ZoneConfig) GetCostingRegion() string`

GetCostingRegion returns the CostingRegion field if non-nil, zero value otherwise.

### GetCostingRegionOk

`func (o *ZoneConfig) GetCostingRegionOk() (*string, bool)`

GetCostingRegionOk returns a tuple with the CostingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingRegion

`func (o *ZoneConfig) SetCostingRegion(v string)`

SetCostingRegion sets CostingRegion field to given value.

### HasCostingRegion

`func (o *ZoneConfig) HasCostingRegion() bool`

HasCostingRegion returns a boolean if a field has been set.

### GetCostingAccessKey

`func (o *ZoneConfig) GetCostingAccessKey() string`

GetCostingAccessKey returns the CostingAccessKey field if non-nil, zero value otherwise.

### GetCostingAccessKeyOk

`func (o *ZoneConfig) GetCostingAccessKeyOk() (*string, bool)`

GetCostingAccessKeyOk returns a tuple with the CostingAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingAccessKey

`func (o *ZoneConfig) SetCostingAccessKey(v string)`

SetCostingAccessKey sets CostingAccessKey field to given value.

### HasCostingAccessKey

`func (o *ZoneConfig) HasCostingAccessKey() bool`

HasCostingAccessKey returns a boolean if a field has been set.

### GetCostingSecretKey

`func (o *ZoneConfig) GetCostingSecretKey() string`

GetCostingSecretKey returns the CostingSecretKey field if non-nil, zero value otherwise.

### GetCostingSecretKeyOk

`func (o *ZoneConfig) GetCostingSecretKeyOk() (*string, bool)`

GetCostingSecretKeyOk returns a tuple with the CostingSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKey

`func (o *ZoneConfig) SetCostingSecretKey(v string)`

SetCostingSecretKey sets CostingSecretKey field to given value.

### HasCostingSecretKey

`func (o *ZoneConfig) HasCostingSecretKey() bool`

HasCostingSecretKey returns a boolean if a field has been set.

### GetCostingReportName

`func (o *ZoneConfig) GetCostingReportName() string`

GetCostingReportName returns the CostingReportName field if non-nil, zero value otherwise.

### GetCostingReportNameOk

`func (o *ZoneConfig) GetCostingReportNameOk() (*string, bool)`

GetCostingReportNameOk returns a tuple with the CostingReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReportName

`func (o *ZoneConfig) SetCostingReportName(v string)`

SetCostingReportName sets CostingReportName field to given value.

### HasCostingReportName

`func (o *ZoneConfig) HasCostingReportName() bool`

HasCostingReportName returns a boolean if a field has been set.

### GetSecretKeyHash

`func (o *ZoneConfig) GetSecretKeyHash() string`

GetSecretKeyHash returns the SecretKeyHash field if non-nil, zero value otherwise.

### GetSecretKeyHashOk

`func (o *ZoneConfig) GetSecretKeyHashOk() (*string, bool)`

GetSecretKeyHashOk returns a tuple with the SecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyHash

`func (o *ZoneConfig) SetSecretKeyHash(v string)`

SetSecretKeyHash sets SecretKeyHash field to given value.

### HasSecretKeyHash

`func (o *ZoneConfig) HasSecretKeyHash() bool`

HasSecretKeyHash returns a boolean if a field has been set.

### GetCostingSecretKeyHash

`func (o *ZoneConfig) GetCostingSecretKeyHash() string`

GetCostingSecretKeyHash returns the CostingSecretKeyHash field if non-nil, zero value otherwise.

### GetCostingSecretKeyHashOk

`func (o *ZoneConfig) GetCostingSecretKeyHashOk() (*string, bool)`

GetCostingSecretKeyHashOk returns a tuple with the CostingSecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKeyHash

`func (o *ZoneConfig) SetCostingSecretKeyHash(v string)`

SetCostingSecretKeyHash sets CostingSecretKeyHash field to given value.

### HasCostingSecretKeyHash

`func (o *ZoneConfig) HasCostingSecretKeyHash() bool`

HasCostingSecretKeyHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



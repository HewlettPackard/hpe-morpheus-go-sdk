# UpdateClouds200ResponseAllOfZoneConfig

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

### NewUpdateClouds200ResponseAllOfZoneConfig

`func NewUpdateClouds200ResponseAllOfZoneConfig() *UpdateClouds200ResponseAllOfZoneConfig`

NewUpdateClouds200ResponseAllOfZoneConfig instantiates a new UpdateClouds200ResponseAllOfZoneConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClouds200ResponseAllOfZoneConfigWithDefaults

`func NewUpdateClouds200ResponseAllOfZoneConfigWithDefaults() *UpdateClouds200ResponseAllOfZoneConfig`

NewUpdateClouds200ResponseAllOfZoneConfigWithDefaults instantiates a new UpdateClouds200ResponseAllOfZoneConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetBackupMode

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetBackupMode() string`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetBackupModeOk() (*string, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetBackupMode(v string)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetConfigCmdbDiscovery

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetConfigCmdbDiscovery() bool`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetConfigCmdbDiscoveryOk() (*bool, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetConfigCmdbDiscovery(v bool)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### GetKubeUrl

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetKubeUrl() string`

GetKubeUrl returns the KubeUrl field if non-nil, zero value otherwise.

### GetKubeUrlOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetKubeUrlOk() (*string, bool)`

GetKubeUrlOk returns a tuple with the KubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeUrl

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetKubeUrl(v string)`

SetKubeUrl sets KubeUrl field to given value.

### HasKubeUrl

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasKubeUrl() bool`

HasKubeUrl returns a boolean if a field has been set.

### GetNetworkServer

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetNetworkServer() AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetNetworkServerOk() (*AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetNetworkServer(v AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetNetworkServerId() string`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetNetworkServerIdOk() (*string, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetNetworkServerId(v string)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetReplicationMode

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### GetSecurityServer

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetApiUrl

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetApiVersion

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCluster

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterRef

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetClusterRef() string`

GetClusterRef returns the ClusterRef field if non-nil, zero value otherwise.

### GetClusterRefOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetClusterRefOk() (*string, bool)`

GetClusterRefOk returns a tuple with the ClusterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRef

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetClusterRef(v string)`

SetClusterRef sets ClusterRef field to given value.

### HasClusterRef

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasClusterRef() bool`

HasClusterRef returns a boolean if a field has been set.

### GetConfigCmId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetConfigCmId() string`

GetConfigCmId returns the ConfigCmId field if non-nil, zero value otherwise.

### GetConfigCmIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetConfigCmIdOk() (*string, bool)`

GetConfigCmIdOk returns a tuple with the ConfigCmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetConfigCmId(v string)`

SetConfigCmId sets ConfigCmId field to given value.

### HasConfigCmId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasConfigCmId() bool`

HasConfigCmId returns a boolean if a field has been set.

### GetConfigCmdbId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetConfigCmdbId() string`

GetConfigCmdbId returns the ConfigCmdbId field if non-nil, zero value otherwise.

### GetConfigCmdbIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetConfigCmdbIdOk() (*string, bool)`

GetConfigCmdbIdOk returns a tuple with the ConfigCmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetConfigCmdbId(v string)`

SetConfigCmdbId sets ConfigCmdbId field to given value.

### HasConfigCmdbId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasConfigCmdbId() bool`

HasConfigCmdbId returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetDatacenter

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetDatacenterId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetDiskStorageType

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetDiskStorageType() string`

GetDiskStorageType returns the DiskStorageType field if non-nil, zero value otherwise.

### GetDiskStorageTypeOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetDiskStorageTypeOk() (*string, bool)`

GetDiskStorageTypeOk returns a tuple with the DiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorageType

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetDiskStorageType(v string)`

SetDiskStorageType sets DiskStorageType field to given value.

### HasDiskStorageType

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasDiskStorageType() bool`

HasDiskStorageType returns a boolean if a field has been set.

### GetDistributedWorkerId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetDistributedWorkerId() string`

GetDistributedWorkerId returns the DistributedWorkerId field if non-nil, zero value otherwise.

### GetDistributedWorkerIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetDistributedWorkerIdOk() (*string, bool)`

GetDistributedWorkerIdOk returns a tuple with the DistributedWorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedWorkerId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetDistributedWorkerId(v string)`

SetDistributedWorkerId sets DistributedWorkerId field to given value.

### HasDistributedWorkerId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasDistributedWorkerId() bool`

HasDistributedWorkerId returns a boolean if a field has been set.

### GetDnsIntegrationId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetDnsIntegrationId() string`

GetDnsIntegrationId returns the DnsIntegrationId field if non-nil, zero value otherwise.

### GetDnsIntegrationIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetDnsIntegrationIdOk() (*string, bool)`

GetDnsIntegrationIdOk returns a tuple with the DnsIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrationId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetDnsIntegrationId(v string)`

SetDnsIntegrationId sets DnsIntegrationId field to given value.

### HasDnsIntegrationId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasDnsIntegrationId() bool`

HasDnsIntegrationId returns a boolean if a field has been set.

### GetEnableDiskTypeSelection

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetEnableDiskTypeSelection() string`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetEnableDiskTypeSelectionOk() (*string, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetEnableDiskTypeSelection(v string)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### GetEnableStorageTypeSelection

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetEnableStorageTypeSelection() string`

GetEnableStorageTypeSelection returns the EnableStorageTypeSelection field if non-nil, zero value otherwise.

### GetEnableStorageTypeSelectionOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetEnableStorageTypeSelectionOk() (*string, bool)`

GetEnableStorageTypeSelectionOk returns a tuple with the EnableStorageTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStorageTypeSelection

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetEnableStorageTypeSelection(v string)`

SetEnableStorageTypeSelection sets EnableStorageTypeSelection field to given value.

### HasEnableStorageTypeSelection

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasEnableStorageTypeSelection() bool`

HasEnableStorageTypeSelection returns a boolean if a field has been set.

### GetEnableVnc

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetEnableVnc() string`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetEnableVncOk() (*string, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetEnableVnc(v string)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### GetHideHostSelection

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetHideHostSelection() string`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetHideHostSelectionOk() (*string, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetHideHostSelection(v string)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### GetImportExisting

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetImportExisting() string`

GetImportExisting returns the ImportExisting field if non-nil, zero value otherwise.

### GetImportExistingOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetImportExistingOk() (*string, bool)`

GetImportExistingOk returns a tuple with the ImportExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportExisting

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetImportExisting(v string)`

SetImportExisting sets ImportExisting field to given value.

### HasImportExisting

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasImportExisting() bool`

HasImportExisting returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetResourcePool

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetRpcMode

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetRpcMode() string`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetRpcModeOk() (*string, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetRpcMode(v string)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### GetSecurityMode

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetServiceRegistryId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetServiceRegistryId() string`

GetServiceRegistryId returns the ServiceRegistryId field if non-nil, zero value otherwise.

### GetServiceRegistryIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetServiceRegistryIdOk() (*string, bool)`

GetServiceRegistryIdOk returns a tuple with the ServiceRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRegistryId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetServiceRegistryId(v string)`

SetServiceRegistryId sets ServiceRegistryId field to given value.

### HasServiceRegistryId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasServiceRegistryId() bool`

HasServiceRegistryId returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEndpoint

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAccessKey

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetIsVpc

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetIsVpc() string`

GetIsVpc returns the IsVpc field if non-nil, zero value otherwise.

### GetIsVpcOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetIsVpcOk() (*string, bool)`

GetIsVpcOk returns a tuple with the IsVpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpc

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetIsVpc(v string)`

SetIsVpc sets IsVpc field to given value.

### HasIsVpc

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasIsVpc() bool`

HasIsVpc returns a boolean if a field has been set.

### GetVpc

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetImageStoreId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetImageStoreId() string`

GetImageStoreId returns the ImageStoreId field if non-nil, zero value otherwise.

### GetImageStoreIdOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetImageStoreIdOk() (*string, bool)`

GetImageStoreIdOk returns a tuple with the ImageStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStoreId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetImageStoreId(v string)`

SetImageStoreId sets ImageStoreId field to given value.

### HasImageStoreId

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasImageStoreId() bool`

HasImageStoreId returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetCostingReport

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingReport() string`

GetCostingReport returns the CostingReport field if non-nil, zero value otherwise.

### GetCostingReportOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingReportOk() (*string, bool)`

GetCostingReportOk returns a tuple with the CostingReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReport

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetCostingReport(v string)`

SetCostingReport sets CostingReport field to given value.

### HasCostingReport

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasCostingReport() bool`

HasCostingReport returns a boolean if a field has been set.

### GetCostingFolder

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingFolder() string`

GetCostingFolder returns the CostingFolder field if non-nil, zero value otherwise.

### GetCostingFolderOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingFolderOk() (*string, bool)`

GetCostingFolderOk returns a tuple with the CostingFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingFolder

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetCostingFolder(v string)`

SetCostingFolder sets CostingFolder field to given value.

### HasCostingFolder

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasCostingFolder() bool`

HasCostingFolder returns a boolean if a field has been set.

### GetCostingBucket

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingBucket() string`

GetCostingBucket returns the CostingBucket field if non-nil, zero value otherwise.

### GetCostingBucketOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingBucketOk() (*string, bool)`

GetCostingBucketOk returns a tuple with the CostingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucket

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetCostingBucket(v string)`

SetCostingBucket sets CostingBucket field to given value.

### HasCostingBucket

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasCostingBucket() bool`

HasCostingBucket returns a boolean if a field has been set.

### GetCostingBucketName

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingBucketName() string`

GetCostingBucketName returns the CostingBucketName field if non-nil, zero value otherwise.

### GetCostingBucketNameOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingBucketNameOk() (*string, bool)`

GetCostingBucketNameOk returns a tuple with the CostingBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucketName

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetCostingBucketName(v string)`

SetCostingBucketName sets CostingBucketName field to given value.

### HasCostingBucketName

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasCostingBucketName() bool`

HasCostingBucketName returns a boolean if a field has been set.

### GetCostingRegion

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingRegion() string`

GetCostingRegion returns the CostingRegion field if non-nil, zero value otherwise.

### GetCostingRegionOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingRegionOk() (*string, bool)`

GetCostingRegionOk returns a tuple with the CostingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingRegion

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetCostingRegion(v string)`

SetCostingRegion sets CostingRegion field to given value.

### HasCostingRegion

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasCostingRegion() bool`

HasCostingRegion returns a boolean if a field has been set.

### GetCostingAccessKey

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingAccessKey() string`

GetCostingAccessKey returns the CostingAccessKey field if non-nil, zero value otherwise.

### GetCostingAccessKeyOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingAccessKeyOk() (*string, bool)`

GetCostingAccessKeyOk returns a tuple with the CostingAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingAccessKey

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetCostingAccessKey(v string)`

SetCostingAccessKey sets CostingAccessKey field to given value.

### HasCostingAccessKey

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasCostingAccessKey() bool`

HasCostingAccessKey returns a boolean if a field has been set.

### GetCostingSecretKey

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingSecretKey() string`

GetCostingSecretKey returns the CostingSecretKey field if non-nil, zero value otherwise.

### GetCostingSecretKeyOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingSecretKeyOk() (*string, bool)`

GetCostingSecretKeyOk returns a tuple with the CostingSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKey

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetCostingSecretKey(v string)`

SetCostingSecretKey sets CostingSecretKey field to given value.

### HasCostingSecretKey

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasCostingSecretKey() bool`

HasCostingSecretKey returns a boolean if a field has been set.

### GetCostingReportName

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingReportName() string`

GetCostingReportName returns the CostingReportName field if non-nil, zero value otherwise.

### GetCostingReportNameOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingReportNameOk() (*string, bool)`

GetCostingReportNameOk returns a tuple with the CostingReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReportName

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetCostingReportName(v string)`

SetCostingReportName sets CostingReportName field to given value.

### HasCostingReportName

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasCostingReportName() bool`

HasCostingReportName returns a boolean if a field has been set.

### GetSecretKeyHash

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetSecretKeyHash() string`

GetSecretKeyHash returns the SecretKeyHash field if non-nil, zero value otherwise.

### GetSecretKeyHashOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetSecretKeyHashOk() (*string, bool)`

GetSecretKeyHashOk returns a tuple with the SecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyHash

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetSecretKeyHash(v string)`

SetSecretKeyHash sets SecretKeyHash field to given value.

### HasSecretKeyHash

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasSecretKeyHash() bool`

HasSecretKeyHash returns a boolean if a field has been set.

### GetCostingSecretKeyHash

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingSecretKeyHash() string`

GetCostingSecretKeyHash returns the CostingSecretKeyHash field if non-nil, zero value otherwise.

### GetCostingSecretKeyHashOk

`func (o *UpdateClouds200ResponseAllOfZoneConfig) GetCostingSecretKeyHashOk() (*string, bool)`

GetCostingSecretKeyHashOk returns a tuple with the CostingSecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKeyHash

`func (o *UpdateClouds200ResponseAllOfZoneConfig) SetCostingSecretKeyHash(v string)`

SetCostingSecretKeyHash sets CostingSecretKeyHash field to given value.

### HasCostingSecretKeyHash

`func (o *UpdateClouds200ResponseAllOfZoneConfig) HasCostingSecretKeyHash() bool`

HasCostingSecretKeyHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



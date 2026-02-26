# GetClouds200ResponseZoneConfig

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

### NewGetClouds200ResponseZoneConfig

`func NewGetClouds200ResponseZoneConfig() *GetClouds200ResponseZoneConfig`

NewGetClouds200ResponseZoneConfig instantiates a new GetClouds200ResponseZoneConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClouds200ResponseZoneConfigWithDefaults

`func NewGetClouds200ResponseZoneConfigWithDefaults() *GetClouds200ResponseZoneConfig`

NewGetClouds200ResponseZoneConfigWithDefaults instantiates a new GetClouds200ResponseZoneConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *GetClouds200ResponseZoneConfig) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *GetClouds200ResponseZoneConfig) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *GetClouds200ResponseZoneConfig) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *GetClouds200ResponseZoneConfig) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *GetClouds200ResponseZoneConfig) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *GetClouds200ResponseZoneConfig) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *GetClouds200ResponseZoneConfig) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *GetClouds200ResponseZoneConfig) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *GetClouds200ResponseZoneConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetClouds200ResponseZoneConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetClouds200ResponseZoneConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetClouds200ResponseZoneConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *GetClouds200ResponseZoneConfig) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *GetClouds200ResponseZoneConfig) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *GetClouds200ResponseZoneConfig) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *GetClouds200ResponseZoneConfig) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *GetClouds200ResponseZoneConfig) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *GetClouds200ResponseZoneConfig) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *GetClouds200ResponseZoneConfig) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *GetClouds200ResponseZoneConfig) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetBackupMode

`func (o *GetClouds200ResponseZoneConfig) GetBackupMode() string`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *GetClouds200ResponseZoneConfig) GetBackupModeOk() (*string, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *GetClouds200ResponseZoneConfig) SetBackupMode(v string)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *GetClouds200ResponseZoneConfig) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *GetClouds200ResponseZoneConfig) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *GetClouds200ResponseZoneConfig) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *GetClouds200ResponseZoneConfig) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *GetClouds200ResponseZoneConfig) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetConfigCmdbDiscovery

`func (o *GetClouds200ResponseZoneConfig) GetConfigCmdbDiscovery() bool`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *GetClouds200ResponseZoneConfig) GetConfigCmdbDiscoveryOk() (*bool, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *GetClouds200ResponseZoneConfig) SetConfigCmdbDiscovery(v bool)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *GetClouds200ResponseZoneConfig) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *GetClouds200ResponseZoneConfig) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *GetClouds200ResponseZoneConfig) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *GetClouds200ResponseZoneConfig) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *GetClouds200ResponseZoneConfig) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### GetKubeUrl

`func (o *GetClouds200ResponseZoneConfig) GetKubeUrl() string`

GetKubeUrl returns the KubeUrl field if non-nil, zero value otherwise.

### GetKubeUrlOk

`func (o *GetClouds200ResponseZoneConfig) GetKubeUrlOk() (*string, bool)`

GetKubeUrlOk returns a tuple with the KubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeUrl

`func (o *GetClouds200ResponseZoneConfig) SetKubeUrl(v string)`

SetKubeUrl sets KubeUrl field to given value.

### HasKubeUrl

`func (o *GetClouds200ResponseZoneConfig) HasKubeUrl() bool`

HasKubeUrl returns a boolean if a field has been set.

### GetNetworkServer

`func (o *GetClouds200ResponseZoneConfig) GetNetworkServer() AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *GetClouds200ResponseZoneConfig) GetNetworkServerOk() (*AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *GetClouds200ResponseZoneConfig) SetNetworkServer(v AddClouds200ResponseAllOfZoneConfigAnyOf2NetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *GetClouds200ResponseZoneConfig) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *GetClouds200ResponseZoneConfig) GetNetworkServerId() string`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *GetClouds200ResponseZoneConfig) GetNetworkServerIdOk() (*string, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *GetClouds200ResponseZoneConfig) SetNetworkServerId(v string)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *GetClouds200ResponseZoneConfig) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetReplicationMode

`func (o *GetClouds200ResponseZoneConfig) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *GetClouds200ResponseZoneConfig) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *GetClouds200ResponseZoneConfig) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *GetClouds200ResponseZoneConfig) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### GetSecurityServer

`func (o *GetClouds200ResponseZoneConfig) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *GetClouds200ResponseZoneConfig) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *GetClouds200ResponseZoneConfig) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *GetClouds200ResponseZoneConfig) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetApiUrl

`func (o *GetClouds200ResponseZoneConfig) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *GetClouds200ResponseZoneConfig) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *GetClouds200ResponseZoneConfig) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *GetClouds200ResponseZoneConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetApiVersion

`func (o *GetClouds200ResponseZoneConfig) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetClouds200ResponseZoneConfig) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetClouds200ResponseZoneConfig) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *GetClouds200ResponseZoneConfig) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCluster

`func (o *GetClouds200ResponseZoneConfig) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *GetClouds200ResponseZoneConfig) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *GetClouds200ResponseZoneConfig) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *GetClouds200ResponseZoneConfig) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterRef

`func (o *GetClouds200ResponseZoneConfig) GetClusterRef() string`

GetClusterRef returns the ClusterRef field if non-nil, zero value otherwise.

### GetClusterRefOk

`func (o *GetClouds200ResponseZoneConfig) GetClusterRefOk() (*string, bool)`

GetClusterRefOk returns a tuple with the ClusterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRef

`func (o *GetClouds200ResponseZoneConfig) SetClusterRef(v string)`

SetClusterRef sets ClusterRef field to given value.

### HasClusterRef

`func (o *GetClouds200ResponseZoneConfig) HasClusterRef() bool`

HasClusterRef returns a boolean if a field has been set.

### GetConfigCmId

`func (o *GetClouds200ResponseZoneConfig) GetConfigCmId() string`

GetConfigCmId returns the ConfigCmId field if non-nil, zero value otherwise.

### GetConfigCmIdOk

`func (o *GetClouds200ResponseZoneConfig) GetConfigCmIdOk() (*string, bool)`

GetConfigCmIdOk returns a tuple with the ConfigCmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmId

`func (o *GetClouds200ResponseZoneConfig) SetConfigCmId(v string)`

SetConfigCmId sets ConfigCmId field to given value.

### HasConfigCmId

`func (o *GetClouds200ResponseZoneConfig) HasConfigCmId() bool`

HasConfigCmId returns a boolean if a field has been set.

### GetConfigCmdbId

`func (o *GetClouds200ResponseZoneConfig) GetConfigCmdbId() string`

GetConfigCmdbId returns the ConfigCmdbId field if non-nil, zero value otherwise.

### GetConfigCmdbIdOk

`func (o *GetClouds200ResponseZoneConfig) GetConfigCmdbIdOk() (*string, bool)`

GetConfigCmdbIdOk returns a tuple with the ConfigCmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbId

`func (o *GetClouds200ResponseZoneConfig) SetConfigCmdbId(v string)`

SetConfigCmdbId sets ConfigCmdbId field to given value.

### HasConfigCmdbId

`func (o *GetClouds200ResponseZoneConfig) HasConfigCmdbId() bool`

HasConfigCmdbId returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *GetClouds200ResponseZoneConfig) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *GetClouds200ResponseZoneConfig) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *GetClouds200ResponseZoneConfig) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *GetClouds200ResponseZoneConfig) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetDatacenter

`func (o *GetClouds200ResponseZoneConfig) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *GetClouds200ResponseZoneConfig) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *GetClouds200ResponseZoneConfig) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *GetClouds200ResponseZoneConfig) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetDatacenterId

`func (o *GetClouds200ResponseZoneConfig) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *GetClouds200ResponseZoneConfig) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *GetClouds200ResponseZoneConfig) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *GetClouds200ResponseZoneConfig) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetDiskStorageType

`func (o *GetClouds200ResponseZoneConfig) GetDiskStorageType() string`

GetDiskStorageType returns the DiskStorageType field if non-nil, zero value otherwise.

### GetDiskStorageTypeOk

`func (o *GetClouds200ResponseZoneConfig) GetDiskStorageTypeOk() (*string, bool)`

GetDiskStorageTypeOk returns a tuple with the DiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorageType

`func (o *GetClouds200ResponseZoneConfig) SetDiskStorageType(v string)`

SetDiskStorageType sets DiskStorageType field to given value.

### HasDiskStorageType

`func (o *GetClouds200ResponseZoneConfig) HasDiskStorageType() bool`

HasDiskStorageType returns a boolean if a field has been set.

### GetDistributedWorkerId

`func (o *GetClouds200ResponseZoneConfig) GetDistributedWorkerId() string`

GetDistributedWorkerId returns the DistributedWorkerId field if non-nil, zero value otherwise.

### GetDistributedWorkerIdOk

`func (o *GetClouds200ResponseZoneConfig) GetDistributedWorkerIdOk() (*string, bool)`

GetDistributedWorkerIdOk returns a tuple with the DistributedWorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedWorkerId

`func (o *GetClouds200ResponseZoneConfig) SetDistributedWorkerId(v string)`

SetDistributedWorkerId sets DistributedWorkerId field to given value.

### HasDistributedWorkerId

`func (o *GetClouds200ResponseZoneConfig) HasDistributedWorkerId() bool`

HasDistributedWorkerId returns a boolean if a field has been set.

### GetDnsIntegrationId

`func (o *GetClouds200ResponseZoneConfig) GetDnsIntegrationId() string`

GetDnsIntegrationId returns the DnsIntegrationId field if non-nil, zero value otherwise.

### GetDnsIntegrationIdOk

`func (o *GetClouds200ResponseZoneConfig) GetDnsIntegrationIdOk() (*string, bool)`

GetDnsIntegrationIdOk returns a tuple with the DnsIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrationId

`func (o *GetClouds200ResponseZoneConfig) SetDnsIntegrationId(v string)`

SetDnsIntegrationId sets DnsIntegrationId field to given value.

### HasDnsIntegrationId

`func (o *GetClouds200ResponseZoneConfig) HasDnsIntegrationId() bool`

HasDnsIntegrationId returns a boolean if a field has been set.

### GetEnableDiskTypeSelection

`func (o *GetClouds200ResponseZoneConfig) GetEnableDiskTypeSelection() string`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *GetClouds200ResponseZoneConfig) GetEnableDiskTypeSelectionOk() (*string, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *GetClouds200ResponseZoneConfig) SetEnableDiskTypeSelection(v string)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *GetClouds200ResponseZoneConfig) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### GetEnableStorageTypeSelection

`func (o *GetClouds200ResponseZoneConfig) GetEnableStorageTypeSelection() string`

GetEnableStorageTypeSelection returns the EnableStorageTypeSelection field if non-nil, zero value otherwise.

### GetEnableStorageTypeSelectionOk

`func (o *GetClouds200ResponseZoneConfig) GetEnableStorageTypeSelectionOk() (*string, bool)`

GetEnableStorageTypeSelectionOk returns a tuple with the EnableStorageTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStorageTypeSelection

`func (o *GetClouds200ResponseZoneConfig) SetEnableStorageTypeSelection(v string)`

SetEnableStorageTypeSelection sets EnableStorageTypeSelection field to given value.

### HasEnableStorageTypeSelection

`func (o *GetClouds200ResponseZoneConfig) HasEnableStorageTypeSelection() bool`

HasEnableStorageTypeSelection returns a boolean if a field has been set.

### GetEnableVnc

`func (o *GetClouds200ResponseZoneConfig) GetEnableVnc() string`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *GetClouds200ResponseZoneConfig) GetEnableVncOk() (*string, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *GetClouds200ResponseZoneConfig) SetEnableVnc(v string)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *GetClouds200ResponseZoneConfig) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### GetHideHostSelection

`func (o *GetClouds200ResponseZoneConfig) GetHideHostSelection() string`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *GetClouds200ResponseZoneConfig) GetHideHostSelectionOk() (*string, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *GetClouds200ResponseZoneConfig) SetHideHostSelection(v string)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *GetClouds200ResponseZoneConfig) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### GetImportExisting

`func (o *GetClouds200ResponseZoneConfig) GetImportExisting() string`

GetImportExisting returns the ImportExisting field if non-nil, zero value otherwise.

### GetImportExistingOk

`func (o *GetClouds200ResponseZoneConfig) GetImportExistingOk() (*string, bool)`

GetImportExistingOk returns a tuple with the ImportExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportExisting

`func (o *GetClouds200ResponseZoneConfig) SetImportExisting(v string)`

SetImportExisting sets ImportExisting field to given value.

### HasImportExisting

`func (o *GetClouds200ResponseZoneConfig) HasImportExisting() bool`

HasImportExisting returns a boolean if a field has been set.

### GetPassword

`func (o *GetClouds200ResponseZoneConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetClouds200ResponseZoneConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetClouds200ResponseZoneConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetClouds200ResponseZoneConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *GetClouds200ResponseZoneConfig) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *GetClouds200ResponseZoneConfig) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *GetClouds200ResponseZoneConfig) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *GetClouds200ResponseZoneConfig) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetResourcePool

`func (o *GetClouds200ResponseZoneConfig) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *GetClouds200ResponseZoneConfig) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *GetClouds200ResponseZoneConfig) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *GetClouds200ResponseZoneConfig) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *GetClouds200ResponseZoneConfig) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *GetClouds200ResponseZoneConfig) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *GetClouds200ResponseZoneConfig) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *GetClouds200ResponseZoneConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetRpcMode

`func (o *GetClouds200ResponseZoneConfig) GetRpcMode() string`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *GetClouds200ResponseZoneConfig) GetRpcModeOk() (*string, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *GetClouds200ResponseZoneConfig) SetRpcMode(v string)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *GetClouds200ResponseZoneConfig) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### GetSecurityMode

`func (o *GetClouds200ResponseZoneConfig) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *GetClouds200ResponseZoneConfig) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *GetClouds200ResponseZoneConfig) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *GetClouds200ResponseZoneConfig) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetServiceRegistryId

`func (o *GetClouds200ResponseZoneConfig) GetServiceRegistryId() string`

GetServiceRegistryId returns the ServiceRegistryId field if non-nil, zero value otherwise.

### GetServiceRegistryIdOk

`func (o *GetClouds200ResponseZoneConfig) GetServiceRegistryIdOk() (*string, bool)`

GetServiceRegistryIdOk returns a tuple with the ServiceRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRegistryId

`func (o *GetClouds200ResponseZoneConfig) SetServiceRegistryId(v string)`

SetServiceRegistryId sets ServiceRegistryId field to given value.

### HasServiceRegistryId

`func (o *GetClouds200ResponseZoneConfig) HasServiceRegistryId() bool`

HasServiceRegistryId returns a boolean if a field has been set.

### GetUsername

`func (o *GetClouds200ResponseZoneConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetClouds200ResponseZoneConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetClouds200ResponseZoneConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetClouds200ResponseZoneConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEndpoint

`func (o *GetClouds200ResponseZoneConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GetClouds200ResponseZoneConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GetClouds200ResponseZoneConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *GetClouds200ResponseZoneConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAccessKey

`func (o *GetClouds200ResponseZoneConfig) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *GetClouds200ResponseZoneConfig) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *GetClouds200ResponseZoneConfig) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *GetClouds200ResponseZoneConfig) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *GetClouds200ResponseZoneConfig) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *GetClouds200ResponseZoneConfig) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *GetClouds200ResponseZoneConfig) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *GetClouds200ResponseZoneConfig) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *GetClouds200ResponseZoneConfig) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *GetClouds200ResponseZoneConfig) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *GetClouds200ResponseZoneConfig) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *GetClouds200ResponseZoneConfig) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *GetClouds200ResponseZoneConfig) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *GetClouds200ResponseZoneConfig) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *GetClouds200ResponseZoneConfig) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *GetClouds200ResponseZoneConfig) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetIsVpc

`func (o *GetClouds200ResponseZoneConfig) GetIsVpc() string`

GetIsVpc returns the IsVpc field if non-nil, zero value otherwise.

### GetIsVpcOk

`func (o *GetClouds200ResponseZoneConfig) GetIsVpcOk() (*string, bool)`

GetIsVpcOk returns a tuple with the IsVpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpc

`func (o *GetClouds200ResponseZoneConfig) SetIsVpc(v string)`

SetIsVpc sets IsVpc field to given value.

### HasIsVpc

`func (o *GetClouds200ResponseZoneConfig) HasIsVpc() bool`

HasIsVpc returns a boolean if a field has been set.

### GetVpc

`func (o *GetClouds200ResponseZoneConfig) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *GetClouds200ResponseZoneConfig) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *GetClouds200ResponseZoneConfig) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *GetClouds200ResponseZoneConfig) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetImageStoreId

`func (o *GetClouds200ResponseZoneConfig) GetImageStoreId() string`

GetImageStoreId returns the ImageStoreId field if non-nil, zero value otherwise.

### GetImageStoreIdOk

`func (o *GetClouds200ResponseZoneConfig) GetImageStoreIdOk() (*string, bool)`

GetImageStoreIdOk returns a tuple with the ImageStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStoreId

`func (o *GetClouds200ResponseZoneConfig) SetImageStoreId(v string)`

SetImageStoreId sets ImageStoreId field to given value.

### HasImageStoreId

`func (o *GetClouds200ResponseZoneConfig) HasImageStoreId() bool`

HasImageStoreId returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *GetClouds200ResponseZoneConfig) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *GetClouds200ResponseZoneConfig) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *GetClouds200ResponseZoneConfig) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *GetClouds200ResponseZoneConfig) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetCostingReport

`func (o *GetClouds200ResponseZoneConfig) GetCostingReport() string`

GetCostingReport returns the CostingReport field if non-nil, zero value otherwise.

### GetCostingReportOk

`func (o *GetClouds200ResponseZoneConfig) GetCostingReportOk() (*string, bool)`

GetCostingReportOk returns a tuple with the CostingReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReport

`func (o *GetClouds200ResponseZoneConfig) SetCostingReport(v string)`

SetCostingReport sets CostingReport field to given value.

### HasCostingReport

`func (o *GetClouds200ResponseZoneConfig) HasCostingReport() bool`

HasCostingReport returns a boolean if a field has been set.

### GetCostingFolder

`func (o *GetClouds200ResponseZoneConfig) GetCostingFolder() string`

GetCostingFolder returns the CostingFolder field if non-nil, zero value otherwise.

### GetCostingFolderOk

`func (o *GetClouds200ResponseZoneConfig) GetCostingFolderOk() (*string, bool)`

GetCostingFolderOk returns a tuple with the CostingFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingFolder

`func (o *GetClouds200ResponseZoneConfig) SetCostingFolder(v string)`

SetCostingFolder sets CostingFolder field to given value.

### HasCostingFolder

`func (o *GetClouds200ResponseZoneConfig) HasCostingFolder() bool`

HasCostingFolder returns a boolean if a field has been set.

### GetCostingBucket

`func (o *GetClouds200ResponseZoneConfig) GetCostingBucket() string`

GetCostingBucket returns the CostingBucket field if non-nil, zero value otherwise.

### GetCostingBucketOk

`func (o *GetClouds200ResponseZoneConfig) GetCostingBucketOk() (*string, bool)`

GetCostingBucketOk returns a tuple with the CostingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucket

`func (o *GetClouds200ResponseZoneConfig) SetCostingBucket(v string)`

SetCostingBucket sets CostingBucket field to given value.

### HasCostingBucket

`func (o *GetClouds200ResponseZoneConfig) HasCostingBucket() bool`

HasCostingBucket returns a boolean if a field has been set.

### GetCostingBucketName

`func (o *GetClouds200ResponseZoneConfig) GetCostingBucketName() string`

GetCostingBucketName returns the CostingBucketName field if non-nil, zero value otherwise.

### GetCostingBucketNameOk

`func (o *GetClouds200ResponseZoneConfig) GetCostingBucketNameOk() (*string, bool)`

GetCostingBucketNameOk returns a tuple with the CostingBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucketName

`func (o *GetClouds200ResponseZoneConfig) SetCostingBucketName(v string)`

SetCostingBucketName sets CostingBucketName field to given value.

### HasCostingBucketName

`func (o *GetClouds200ResponseZoneConfig) HasCostingBucketName() bool`

HasCostingBucketName returns a boolean if a field has been set.

### GetCostingRegion

`func (o *GetClouds200ResponseZoneConfig) GetCostingRegion() string`

GetCostingRegion returns the CostingRegion field if non-nil, zero value otherwise.

### GetCostingRegionOk

`func (o *GetClouds200ResponseZoneConfig) GetCostingRegionOk() (*string, bool)`

GetCostingRegionOk returns a tuple with the CostingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingRegion

`func (o *GetClouds200ResponseZoneConfig) SetCostingRegion(v string)`

SetCostingRegion sets CostingRegion field to given value.

### HasCostingRegion

`func (o *GetClouds200ResponseZoneConfig) HasCostingRegion() bool`

HasCostingRegion returns a boolean if a field has been set.

### GetCostingAccessKey

`func (o *GetClouds200ResponseZoneConfig) GetCostingAccessKey() string`

GetCostingAccessKey returns the CostingAccessKey field if non-nil, zero value otherwise.

### GetCostingAccessKeyOk

`func (o *GetClouds200ResponseZoneConfig) GetCostingAccessKeyOk() (*string, bool)`

GetCostingAccessKeyOk returns a tuple with the CostingAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingAccessKey

`func (o *GetClouds200ResponseZoneConfig) SetCostingAccessKey(v string)`

SetCostingAccessKey sets CostingAccessKey field to given value.

### HasCostingAccessKey

`func (o *GetClouds200ResponseZoneConfig) HasCostingAccessKey() bool`

HasCostingAccessKey returns a boolean if a field has been set.

### GetCostingSecretKey

`func (o *GetClouds200ResponseZoneConfig) GetCostingSecretKey() string`

GetCostingSecretKey returns the CostingSecretKey field if non-nil, zero value otherwise.

### GetCostingSecretKeyOk

`func (o *GetClouds200ResponseZoneConfig) GetCostingSecretKeyOk() (*string, bool)`

GetCostingSecretKeyOk returns a tuple with the CostingSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKey

`func (o *GetClouds200ResponseZoneConfig) SetCostingSecretKey(v string)`

SetCostingSecretKey sets CostingSecretKey field to given value.

### HasCostingSecretKey

`func (o *GetClouds200ResponseZoneConfig) HasCostingSecretKey() bool`

HasCostingSecretKey returns a boolean if a field has been set.

### GetCostingReportName

`func (o *GetClouds200ResponseZoneConfig) GetCostingReportName() string`

GetCostingReportName returns the CostingReportName field if non-nil, zero value otherwise.

### GetCostingReportNameOk

`func (o *GetClouds200ResponseZoneConfig) GetCostingReportNameOk() (*string, bool)`

GetCostingReportNameOk returns a tuple with the CostingReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReportName

`func (o *GetClouds200ResponseZoneConfig) SetCostingReportName(v string)`

SetCostingReportName sets CostingReportName field to given value.

### HasCostingReportName

`func (o *GetClouds200ResponseZoneConfig) HasCostingReportName() bool`

HasCostingReportName returns a boolean if a field has been set.

### GetSecretKeyHash

`func (o *GetClouds200ResponseZoneConfig) GetSecretKeyHash() string`

GetSecretKeyHash returns the SecretKeyHash field if non-nil, zero value otherwise.

### GetSecretKeyHashOk

`func (o *GetClouds200ResponseZoneConfig) GetSecretKeyHashOk() (*string, bool)`

GetSecretKeyHashOk returns a tuple with the SecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyHash

`func (o *GetClouds200ResponseZoneConfig) SetSecretKeyHash(v string)`

SetSecretKeyHash sets SecretKeyHash field to given value.

### HasSecretKeyHash

`func (o *GetClouds200ResponseZoneConfig) HasSecretKeyHash() bool`

HasSecretKeyHash returns a boolean if a field has been set.

### GetCostingSecretKeyHash

`func (o *GetClouds200ResponseZoneConfig) GetCostingSecretKeyHash() string`

GetCostingSecretKeyHash returns the CostingSecretKeyHash field if non-nil, zero value otherwise.

### GetCostingSecretKeyHashOk

`func (o *GetClouds200ResponseZoneConfig) GetCostingSecretKeyHashOk() (*string, bool)`

GetCostingSecretKeyHashOk returns a tuple with the CostingSecretKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecretKeyHash

`func (o *GetClouds200ResponseZoneConfig) SetCostingSecretKeyHash(v string)`

SetCostingSecretKeyHash sets CostingSecretKeyHash field to given value.

### HasCostingSecretKeyHash

`func (o *GetClouds200ResponseZoneConfig) HasCostingSecretKeyHash() bool`

HasCostingSecretKeyHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



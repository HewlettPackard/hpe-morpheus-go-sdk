# GetClouds200ResponseZoneConfigAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | Pointer to **NullableString** |  | [optional] 
**DatacenterName** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**InventoryLevel** | Pointer to **NullableString** |  | [optional] 
**ConsoleKeymap** | Pointer to **NullableString** |  | [optional] 
**ApiUrl** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **NullableString** |  | [optional] 
**BackupMode** | Pointer to **NullableString** |  | [optional] 
**CertificateProvider** | Pointer to **NullableString** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**ClusterRef** | Pointer to **string** |  | [optional] 
**ConfigCmId** | Pointer to **NullableString** |  | [optional] 
**ConfigCmdbDiscovery** | Pointer to **bool** |  | [optional] 
**ConfigCmdbId** | Pointer to **NullableString** |  | [optional] 
**ConfigManagementId** | Pointer to **NullableString** |  | [optional] 
**Datacenter** | Pointer to **string** |  | [optional] 
**DatacenterId** | Pointer to **NullableString** |  | [optional] 
**DiskStorageType** | Pointer to **NullableString** |  | [optional] 
**DistributedWorkerId** | Pointer to **NullableString** |  | [optional] 
**DnsIntegrationId** | Pointer to **NullableString** |  | [optional] 
**EnableDiskTypeSelection** | Pointer to **NullableString** |  | [optional] 
**EnableNetworkTypeSelection** | Pointer to **NullableString** |  | [optional] 
**EnableStorageTypeSelection** | Pointer to **NullableString** |  | [optional] 
**EnableVnc** | Pointer to **NullableString** |  | [optional] 
**HideHostSelection** | Pointer to **NullableString** |  | [optional] 
**ImportExisting** | Pointer to **NullableString** |  | [optional] 
**KubeUrl** | Pointer to **NullableString** |  | [optional] 
**NetworkServer** | Pointer to [**AddClouds200ResponseAllOfZoneConfigAnyOf1NetworkServer**](AddClouds200ResponseAllOfZoneConfigAnyOf1NetworkServer.md) |  | [optional] 
**NetworkServerId** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 
**ReplicationMode** | Pointer to **NullableString** |  | [optional] 
**ResourcePool** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **string** |  | [optional] 
**RpcMode** | Pointer to **string** |  | [optional] 
**SecurityMode** | Pointer to **string** |  | [optional] 
**SecurityServer** | Pointer to **NullableString** |  | [optional] 
**ServiceRegistryId** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewGetClouds200ResponseZoneConfigAnyOf1

`func NewGetClouds200ResponseZoneConfigAnyOf1() *GetClouds200ResponseZoneConfigAnyOf1`

NewGetClouds200ResponseZoneConfigAnyOf1 instantiates a new GetClouds200ResponseZoneConfigAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClouds200ResponseZoneConfigAnyOf1WithDefaults

`func NewGetClouds200ResponseZoneConfigAnyOf1WithDefaults() *GetClouds200ResponseZoneConfigAnyOf1`

NewGetClouds200ResponseZoneConfigAnyOf1WithDefaults instantiates a new GetClouds200ResponseZoneConfigAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### SetApplianceUrlNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetApplianceUrlNil(b bool)`

 SetApplianceUrlNil sets the value for ApplianceUrl to be an explicit nil

### UnsetApplianceUrl
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetApplianceUrl()`

UnsetApplianceUrl ensures that no value is present for ApplianceUrl, not even an explicit nil
### GetDatacenterName

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### SetDatacenterNameNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetDatacenterNameNil(b bool)`

 SetDatacenterNameNil sets the value for DatacenterName to be an explicit nil

### UnsetDatacenterName
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetDatacenterName()`

UnsetDatacenterName ensures that no value is present for DatacenterName, not even an explicit nil
### GetExternalId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInventoryLevel

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### SetInventoryLevelNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetInventoryLevelNil(b bool)`

 SetInventoryLevelNil sets the value for InventoryLevel to be an explicit nil

### UnsetInventoryLevel
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetInventoryLevel()`

UnsetInventoryLevel ensures that no value is present for InventoryLevel, not even an explicit nil
### GetConsoleKeymap

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### SetConsoleKeymapNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetConsoleKeymapNil(b bool)`

 SetConsoleKeymapNil sets the value for ConsoleKeymap to be an explicit nil

### UnsetConsoleKeymap
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetConsoleKeymap()`

UnsetConsoleKeymap ensures that no value is present for ConsoleKeymap, not even an explicit nil
### GetApiUrl

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetApiVersion

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetBackupMode

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetBackupMode() string`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetBackupModeOk() (*string, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetBackupMode(v string)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### SetBackupModeNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetBackupModeNil(b bool)`

 SetBackupModeNil sets the value for BackupMode to be an explicit nil

### UnsetBackupMode
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetBackupMode()`

UnsetBackupMode ensures that no value is present for BackupMode, not even an explicit nil
### GetCertificateProvider

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### SetCertificateProviderNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetCertificateProviderNil(b bool)`

 SetCertificateProviderNil sets the value for CertificateProvider to be an explicit nil

### UnsetCertificateProvider
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetCertificateProvider()`

UnsetCertificateProvider ensures that no value is present for CertificateProvider, not even an explicit nil
### GetCluster

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterRef

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetClusterRef() string`

GetClusterRef returns the ClusterRef field if non-nil, zero value otherwise.

### GetClusterRefOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetClusterRefOk() (*string, bool)`

GetClusterRefOk returns a tuple with the ClusterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRef

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetClusterRef(v string)`

SetClusterRef sets ClusterRef field to given value.

### HasClusterRef

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasClusterRef() bool`

HasClusterRef returns a boolean if a field has been set.

### GetConfigCmId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetConfigCmId() string`

GetConfigCmId returns the ConfigCmId field if non-nil, zero value otherwise.

### GetConfigCmIdOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetConfigCmIdOk() (*string, bool)`

GetConfigCmIdOk returns a tuple with the ConfigCmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetConfigCmId(v string)`

SetConfigCmId sets ConfigCmId field to given value.

### HasConfigCmId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasConfigCmId() bool`

HasConfigCmId returns a boolean if a field has been set.

### SetConfigCmIdNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetConfigCmIdNil(b bool)`

 SetConfigCmIdNil sets the value for ConfigCmId to be an explicit nil

### UnsetConfigCmId
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetConfigCmId()`

UnsetConfigCmId ensures that no value is present for ConfigCmId, not even an explicit nil
### GetConfigCmdbDiscovery

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetConfigCmdbDiscovery() bool`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetConfigCmdbDiscoveryOk() (*bool, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetConfigCmdbDiscovery(v bool)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetConfigCmdbId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetConfigCmdbId() string`

GetConfigCmdbId returns the ConfigCmdbId field if non-nil, zero value otherwise.

### GetConfigCmdbIdOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetConfigCmdbIdOk() (*string, bool)`

GetConfigCmdbIdOk returns a tuple with the ConfigCmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetConfigCmdbId(v string)`

SetConfigCmdbId sets ConfigCmdbId field to given value.

### HasConfigCmdbId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasConfigCmdbId() bool`

HasConfigCmdbId returns a boolean if a field has been set.

### SetConfigCmdbIdNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetConfigCmdbIdNil(b bool)`

 SetConfigCmdbIdNil sets the value for ConfigCmdbId to be an explicit nil

### UnsetConfigCmdbId
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetConfigCmdbId()`

UnsetConfigCmdbId ensures that no value is present for ConfigCmdbId, not even an explicit nil
### GetConfigManagementId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### SetConfigManagementIdNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetConfigManagementIdNil(b bool)`

 SetConfigManagementIdNil sets the value for ConfigManagementId to be an explicit nil

### UnsetConfigManagementId
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetConfigManagementId()`

UnsetConfigManagementId ensures that no value is present for ConfigManagementId, not even an explicit nil
### GetDatacenter

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetDatacenterId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### SetDatacenterIdNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetDatacenterIdNil(b bool)`

 SetDatacenterIdNil sets the value for DatacenterId to be an explicit nil

### UnsetDatacenterId
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetDatacenterId()`

UnsetDatacenterId ensures that no value is present for DatacenterId, not even an explicit nil
### GetDiskStorageType

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetDiskStorageType() string`

GetDiskStorageType returns the DiskStorageType field if non-nil, zero value otherwise.

### GetDiskStorageTypeOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetDiskStorageTypeOk() (*string, bool)`

GetDiskStorageTypeOk returns a tuple with the DiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorageType

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetDiskStorageType(v string)`

SetDiskStorageType sets DiskStorageType field to given value.

### HasDiskStorageType

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasDiskStorageType() bool`

HasDiskStorageType returns a boolean if a field has been set.

### SetDiskStorageTypeNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetDiskStorageTypeNil(b bool)`

 SetDiskStorageTypeNil sets the value for DiskStorageType to be an explicit nil

### UnsetDiskStorageType
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetDiskStorageType()`

UnsetDiskStorageType ensures that no value is present for DiskStorageType, not even an explicit nil
### GetDistributedWorkerId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetDistributedWorkerId() string`

GetDistributedWorkerId returns the DistributedWorkerId field if non-nil, zero value otherwise.

### GetDistributedWorkerIdOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetDistributedWorkerIdOk() (*string, bool)`

GetDistributedWorkerIdOk returns a tuple with the DistributedWorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedWorkerId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetDistributedWorkerId(v string)`

SetDistributedWorkerId sets DistributedWorkerId field to given value.

### HasDistributedWorkerId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasDistributedWorkerId() bool`

HasDistributedWorkerId returns a boolean if a field has been set.

### SetDistributedWorkerIdNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetDistributedWorkerIdNil(b bool)`

 SetDistributedWorkerIdNil sets the value for DistributedWorkerId to be an explicit nil

### UnsetDistributedWorkerId
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetDistributedWorkerId()`

UnsetDistributedWorkerId ensures that no value is present for DistributedWorkerId, not even an explicit nil
### GetDnsIntegrationId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetDnsIntegrationId() string`

GetDnsIntegrationId returns the DnsIntegrationId field if non-nil, zero value otherwise.

### GetDnsIntegrationIdOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetDnsIntegrationIdOk() (*string, bool)`

GetDnsIntegrationIdOk returns a tuple with the DnsIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrationId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetDnsIntegrationId(v string)`

SetDnsIntegrationId sets DnsIntegrationId field to given value.

### HasDnsIntegrationId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasDnsIntegrationId() bool`

HasDnsIntegrationId returns a boolean if a field has been set.

### SetDnsIntegrationIdNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetDnsIntegrationIdNil(b bool)`

 SetDnsIntegrationIdNil sets the value for DnsIntegrationId to be an explicit nil

### UnsetDnsIntegrationId
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetDnsIntegrationId()`

UnsetDnsIntegrationId ensures that no value is present for DnsIntegrationId, not even an explicit nil
### GetEnableDiskTypeSelection

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetEnableDiskTypeSelection() string`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetEnableDiskTypeSelectionOk() (*string, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetEnableDiskTypeSelection(v string)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### SetEnableDiskTypeSelectionNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetEnableDiskTypeSelectionNil(b bool)`

 SetEnableDiskTypeSelectionNil sets the value for EnableDiskTypeSelection to be an explicit nil

### UnsetEnableDiskTypeSelection
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetEnableDiskTypeSelection()`

UnsetEnableDiskTypeSelection ensures that no value is present for EnableDiskTypeSelection, not even an explicit nil
### GetEnableNetworkTypeSelection

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### SetEnableNetworkTypeSelectionNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetEnableNetworkTypeSelectionNil(b bool)`

 SetEnableNetworkTypeSelectionNil sets the value for EnableNetworkTypeSelection to be an explicit nil

### UnsetEnableNetworkTypeSelection
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetEnableNetworkTypeSelection()`

UnsetEnableNetworkTypeSelection ensures that no value is present for EnableNetworkTypeSelection, not even an explicit nil
### GetEnableStorageTypeSelection

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetEnableStorageTypeSelection() string`

GetEnableStorageTypeSelection returns the EnableStorageTypeSelection field if non-nil, zero value otherwise.

### GetEnableStorageTypeSelectionOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetEnableStorageTypeSelectionOk() (*string, bool)`

GetEnableStorageTypeSelectionOk returns a tuple with the EnableStorageTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStorageTypeSelection

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetEnableStorageTypeSelection(v string)`

SetEnableStorageTypeSelection sets EnableStorageTypeSelection field to given value.

### HasEnableStorageTypeSelection

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasEnableStorageTypeSelection() bool`

HasEnableStorageTypeSelection returns a boolean if a field has been set.

### SetEnableStorageTypeSelectionNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetEnableStorageTypeSelectionNil(b bool)`

 SetEnableStorageTypeSelectionNil sets the value for EnableStorageTypeSelection to be an explicit nil

### UnsetEnableStorageTypeSelection
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetEnableStorageTypeSelection()`

UnsetEnableStorageTypeSelection ensures that no value is present for EnableStorageTypeSelection, not even an explicit nil
### GetEnableVnc

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetEnableVnc() string`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetEnableVncOk() (*string, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetEnableVnc(v string)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### SetEnableVncNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetEnableVncNil(b bool)`

 SetEnableVncNil sets the value for EnableVnc to be an explicit nil

### UnsetEnableVnc
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetEnableVnc()`

UnsetEnableVnc ensures that no value is present for EnableVnc, not even an explicit nil
### GetHideHostSelection

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetHideHostSelection() string`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetHideHostSelectionOk() (*string, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetHideHostSelection(v string)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### SetHideHostSelectionNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetHideHostSelectionNil(b bool)`

 SetHideHostSelectionNil sets the value for HideHostSelection to be an explicit nil

### UnsetHideHostSelection
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetHideHostSelection()`

UnsetHideHostSelection ensures that no value is present for HideHostSelection, not even an explicit nil
### GetImportExisting

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetImportExisting() string`

GetImportExisting returns the ImportExisting field if non-nil, zero value otherwise.

### GetImportExistingOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetImportExistingOk() (*string, bool)`

GetImportExistingOk returns a tuple with the ImportExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportExisting

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetImportExisting(v string)`

SetImportExisting sets ImportExisting field to given value.

### HasImportExisting

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasImportExisting() bool`

HasImportExisting returns a boolean if a field has been set.

### SetImportExistingNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetImportExistingNil(b bool)`

 SetImportExistingNil sets the value for ImportExisting to be an explicit nil

### UnsetImportExisting
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetImportExisting()`

UnsetImportExisting ensures that no value is present for ImportExisting, not even an explicit nil
### GetKubeUrl

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetKubeUrl() string`

GetKubeUrl returns the KubeUrl field if non-nil, zero value otherwise.

### GetKubeUrlOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetKubeUrlOk() (*string, bool)`

GetKubeUrlOk returns a tuple with the KubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeUrl

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetKubeUrl(v string)`

SetKubeUrl sets KubeUrl field to given value.

### HasKubeUrl

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasKubeUrl() bool`

HasKubeUrl returns a boolean if a field has been set.

### SetKubeUrlNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetKubeUrlNil(b bool)`

 SetKubeUrlNil sets the value for KubeUrl to be an explicit nil

### UnsetKubeUrl
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetKubeUrl()`

UnsetKubeUrl ensures that no value is present for KubeUrl, not even an explicit nil
### GetNetworkServer

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetNetworkServer() AddClouds200ResponseAllOfZoneConfigAnyOf1NetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetNetworkServerOk() (*AddClouds200ResponseAllOfZoneConfigAnyOf1NetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetNetworkServer(v AddClouds200ResponseAllOfZoneConfigAnyOf1NetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetNetworkServerId() string`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetNetworkServerIdOk() (*string, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetNetworkServerId(v string)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetPassword

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetReplicationMode

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### SetReplicationModeNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetReplicationModeNil(b bool)`

 SetReplicationModeNil sets the value for ReplicationMode to be an explicit nil

### UnsetReplicationMode
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetReplicationMode()`

UnsetReplicationMode ensures that no value is present for ReplicationMode, not even an explicit nil
### GetResourcePool

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetRpcMode

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetRpcMode() string`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetRpcModeOk() (*string, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetRpcMode(v string)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### GetSecurityMode

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetSecurityServer

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### SetSecurityServerNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetSecurityServerNil(b bool)`

 SetSecurityServerNil sets the value for SecurityServer to be an explicit nil

### UnsetSecurityServer
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetSecurityServer()`

UnsetSecurityServer ensures that no value is present for SecurityServer, not even an explicit nil
### GetServiceRegistryId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetServiceRegistryId() string`

GetServiceRegistryId returns the ServiceRegistryId field if non-nil, zero value otherwise.

### GetServiceRegistryIdOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetServiceRegistryIdOk() (*string, bool)`

GetServiceRegistryIdOk returns a tuple with the ServiceRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRegistryId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetServiceRegistryId(v string)`

SetServiceRegistryId sets ServiceRegistryId field to given value.

### HasServiceRegistryId

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasServiceRegistryId() bool`

HasServiceRegistryId returns a boolean if a field has been set.

### SetServiceRegistryIdNil

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetServiceRegistryIdNil(b bool)`

 SetServiceRegistryIdNil sets the value for ServiceRegistryId to be an explicit nil

### UnsetServiceRegistryId
`func (o *GetClouds200ResponseZoneConfigAnyOf1) UnsetServiceRegistryId()`

UnsetServiceRegistryId ensures that no value is present for ServiceRegistryId, not even an explicit nil
### GetUsername

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetClouds200ResponseZoneConfigAnyOf1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetClouds200ResponseZoneConfigAnyOf1) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetClouds200ResponseZoneConfigAnyOf1) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



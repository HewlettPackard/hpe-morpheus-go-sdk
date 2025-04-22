# ZoneAzureConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriberId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**ResourceGroup** | Pointer to **string** |  | [optional] 
**ImportExisting** | Pointer to **string** |  | [optional] 
**InventoryLevel** | Pointer to **string** |  | [optional] 
**ApplianceUrl** | Pointer to **string** |  | [optional] 
**DatacenterName** | Pointer to **string** |  | [optional] 
**NetworkServerId** | Pointer to **string** |  | [optional] 
**NetworkServer** | Pointer to [**ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer**](ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer.md) |  | [optional] 
**SecurityMode** | Pointer to **string** |  | [optional] 
**CertificateProvider** | Pointer to **string** |  | [optional] 
**BackupMode** | Pointer to **string** |  | [optional] 
**ReplicationMode** | Pointer to **string** |  | [optional] 
**DnsIntegrationId** | Pointer to **string** |  | [optional] 
**ConfigManagementId** | Pointer to **string** |  | [optional] 
**ConfigCmdbId** | Pointer to **string** |  | [optional] 
**SecurityServer** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**ServiceRegistryId** | Pointer to **string** |  | [optional] 
**CloudType** | Pointer to **string** |  | [optional] 
**RpcMode** | Pointer to **string** |  | [optional] 
**DiskEncryption** | Pointer to **string** |  | [optional] 
**EncryptionSet** | Pointer to **string** |  | [optional] 
**CspTenantId** | Pointer to **string** |  | [optional] 
**CspClientId** | Pointer to **string** |  | [optional] 
**CspClientSecret** | Pointer to **string** |  | [optional] 
**CspCustomer** | Pointer to **string** |  | [optional] 
**ConfigCmdbDiscovery** | Pointer to **bool** |  | [optional] 
**AzureCostingMode** | Pointer to **string** |  | [optional] 
**ClientSecretHash** | Pointer to **string** |  | [optional] 
**CspClientSecretHash** | Pointer to **string** |  | [optional] 

## Methods

### NewZoneAzureConfig

`func NewZoneAzureConfig() *ZoneAzureConfig`

NewZoneAzureConfig instantiates a new ZoneAzureConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAzureConfigWithDefaults

`func NewZoneAzureConfigWithDefaults() *ZoneAzureConfig`

NewZoneAzureConfigWithDefaults instantiates a new ZoneAzureConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriberId

`func (o *ZoneAzureConfig) GetSubscriberId() string`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *ZoneAzureConfig) GetSubscriberIdOk() (*string, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *ZoneAzureConfig) SetSubscriberId(v string)`

SetSubscriberId sets SubscriberId field to given value.

### HasSubscriberId

`func (o *ZoneAzureConfig) HasSubscriberId() bool`

HasSubscriberId returns a boolean if a field has been set.

### GetTenantId

`func (o *ZoneAzureConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ZoneAzureConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ZoneAzureConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ZoneAzureConfig) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetClientId

`func (o *ZoneAzureConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ZoneAzureConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ZoneAzureConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ZoneAzureConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ZoneAzureConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ZoneAzureConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ZoneAzureConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ZoneAzureConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetResourceGroup

`func (o *ZoneAzureConfig) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *ZoneAzureConfig) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *ZoneAzureConfig) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *ZoneAzureConfig) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetImportExisting

`func (o *ZoneAzureConfig) GetImportExisting() string`

GetImportExisting returns the ImportExisting field if non-nil, zero value otherwise.

### GetImportExistingOk

`func (o *ZoneAzureConfig) GetImportExistingOk() (*string, bool)`

GetImportExistingOk returns a tuple with the ImportExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportExisting

`func (o *ZoneAzureConfig) SetImportExisting(v string)`

SetImportExisting sets ImportExisting field to given value.

### HasImportExisting

`func (o *ZoneAzureConfig) HasImportExisting() bool`

HasImportExisting returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *ZoneAzureConfig) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *ZoneAzureConfig) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *ZoneAzureConfig) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *ZoneAzureConfig) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetApplianceUrl

`func (o *ZoneAzureConfig) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ZoneAzureConfig) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ZoneAzureConfig) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ZoneAzureConfig) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *ZoneAzureConfig) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *ZoneAzureConfig) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *ZoneAzureConfig) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *ZoneAzureConfig) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *ZoneAzureConfig) GetNetworkServerId() string`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *ZoneAzureConfig) GetNetworkServerIdOk() (*string, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *ZoneAzureConfig) SetNetworkServerId(v string)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *ZoneAzureConfig) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetNetworkServer

`func (o *ZoneAzureConfig) GetNetworkServer() ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *ZoneAzureConfig) GetNetworkServerOk() (*ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *ZoneAzureConfig) SetNetworkServer(v ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *ZoneAzureConfig) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetSecurityMode

`func (o *ZoneAzureConfig) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *ZoneAzureConfig) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *ZoneAzureConfig) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *ZoneAzureConfig) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *ZoneAzureConfig) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *ZoneAzureConfig) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *ZoneAzureConfig) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *ZoneAzureConfig) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetBackupMode

`func (o *ZoneAzureConfig) GetBackupMode() string`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *ZoneAzureConfig) GetBackupModeOk() (*string, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *ZoneAzureConfig) SetBackupMode(v string)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *ZoneAzureConfig) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### GetReplicationMode

`func (o *ZoneAzureConfig) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *ZoneAzureConfig) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *ZoneAzureConfig) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *ZoneAzureConfig) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### GetDnsIntegrationId

`func (o *ZoneAzureConfig) GetDnsIntegrationId() string`

GetDnsIntegrationId returns the DnsIntegrationId field if non-nil, zero value otherwise.

### GetDnsIntegrationIdOk

`func (o *ZoneAzureConfig) GetDnsIntegrationIdOk() (*string, bool)`

GetDnsIntegrationIdOk returns a tuple with the DnsIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrationId

`func (o *ZoneAzureConfig) SetDnsIntegrationId(v string)`

SetDnsIntegrationId sets DnsIntegrationId field to given value.

### HasDnsIntegrationId

`func (o *ZoneAzureConfig) HasDnsIntegrationId() bool`

HasDnsIntegrationId returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *ZoneAzureConfig) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *ZoneAzureConfig) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *ZoneAzureConfig) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *ZoneAzureConfig) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetConfigCmdbId

`func (o *ZoneAzureConfig) GetConfigCmdbId() string`

GetConfigCmdbId returns the ConfigCmdbId field if non-nil, zero value otherwise.

### GetConfigCmdbIdOk

`func (o *ZoneAzureConfig) GetConfigCmdbIdOk() (*string, bool)`

GetConfigCmdbIdOk returns a tuple with the ConfigCmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbId

`func (o *ZoneAzureConfig) SetConfigCmdbId(v string)`

SetConfigCmdbId sets ConfigCmdbId field to given value.

### HasConfigCmdbId

`func (o *ZoneAzureConfig) HasConfigCmdbId() bool`

HasConfigCmdbId returns a boolean if a field has been set.

### GetSecurityServer

`func (o *ZoneAzureConfig) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *ZoneAzureConfig) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *ZoneAzureConfig) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *ZoneAzureConfig) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetAccountType

`func (o *ZoneAzureConfig) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ZoneAzureConfig) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ZoneAzureConfig) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ZoneAzureConfig) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetServiceRegistryId

`func (o *ZoneAzureConfig) GetServiceRegistryId() string`

GetServiceRegistryId returns the ServiceRegistryId field if non-nil, zero value otherwise.

### GetServiceRegistryIdOk

`func (o *ZoneAzureConfig) GetServiceRegistryIdOk() (*string, bool)`

GetServiceRegistryIdOk returns a tuple with the ServiceRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRegistryId

`func (o *ZoneAzureConfig) SetServiceRegistryId(v string)`

SetServiceRegistryId sets ServiceRegistryId field to given value.

### HasServiceRegistryId

`func (o *ZoneAzureConfig) HasServiceRegistryId() bool`

HasServiceRegistryId returns a boolean if a field has been set.

### GetCloudType

`func (o *ZoneAzureConfig) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ZoneAzureConfig) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ZoneAzureConfig) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *ZoneAzureConfig) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetRpcMode

`func (o *ZoneAzureConfig) GetRpcMode() string`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *ZoneAzureConfig) GetRpcModeOk() (*string, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *ZoneAzureConfig) SetRpcMode(v string)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *ZoneAzureConfig) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### GetDiskEncryption

`func (o *ZoneAzureConfig) GetDiskEncryption() string`

GetDiskEncryption returns the DiskEncryption field if non-nil, zero value otherwise.

### GetDiskEncryptionOk

`func (o *ZoneAzureConfig) GetDiskEncryptionOk() (*string, bool)`

GetDiskEncryptionOk returns a tuple with the DiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryption

`func (o *ZoneAzureConfig) SetDiskEncryption(v string)`

SetDiskEncryption sets DiskEncryption field to given value.

### HasDiskEncryption

`func (o *ZoneAzureConfig) HasDiskEncryption() bool`

HasDiskEncryption returns a boolean if a field has been set.

### GetEncryptionSet

`func (o *ZoneAzureConfig) GetEncryptionSet() string`

GetEncryptionSet returns the EncryptionSet field if non-nil, zero value otherwise.

### GetEncryptionSetOk

`func (o *ZoneAzureConfig) GetEncryptionSetOk() (*string, bool)`

GetEncryptionSetOk returns a tuple with the EncryptionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSet

`func (o *ZoneAzureConfig) SetEncryptionSet(v string)`

SetEncryptionSet sets EncryptionSet field to given value.

### HasEncryptionSet

`func (o *ZoneAzureConfig) HasEncryptionSet() bool`

HasEncryptionSet returns a boolean if a field has been set.

### GetCspTenantId

`func (o *ZoneAzureConfig) GetCspTenantId() string`

GetCspTenantId returns the CspTenantId field if non-nil, zero value otherwise.

### GetCspTenantIdOk

`func (o *ZoneAzureConfig) GetCspTenantIdOk() (*string, bool)`

GetCspTenantIdOk returns a tuple with the CspTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspTenantId

`func (o *ZoneAzureConfig) SetCspTenantId(v string)`

SetCspTenantId sets CspTenantId field to given value.

### HasCspTenantId

`func (o *ZoneAzureConfig) HasCspTenantId() bool`

HasCspTenantId returns a boolean if a field has been set.

### GetCspClientId

`func (o *ZoneAzureConfig) GetCspClientId() string`

GetCspClientId returns the CspClientId field if non-nil, zero value otherwise.

### GetCspClientIdOk

`func (o *ZoneAzureConfig) GetCspClientIdOk() (*string, bool)`

GetCspClientIdOk returns a tuple with the CspClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspClientId

`func (o *ZoneAzureConfig) SetCspClientId(v string)`

SetCspClientId sets CspClientId field to given value.

### HasCspClientId

`func (o *ZoneAzureConfig) HasCspClientId() bool`

HasCspClientId returns a boolean if a field has been set.

### GetCspClientSecret

`func (o *ZoneAzureConfig) GetCspClientSecret() string`

GetCspClientSecret returns the CspClientSecret field if non-nil, zero value otherwise.

### GetCspClientSecretOk

`func (o *ZoneAzureConfig) GetCspClientSecretOk() (*string, bool)`

GetCspClientSecretOk returns a tuple with the CspClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspClientSecret

`func (o *ZoneAzureConfig) SetCspClientSecret(v string)`

SetCspClientSecret sets CspClientSecret field to given value.

### HasCspClientSecret

`func (o *ZoneAzureConfig) HasCspClientSecret() bool`

HasCspClientSecret returns a boolean if a field has been set.

### GetCspCustomer

`func (o *ZoneAzureConfig) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *ZoneAzureConfig) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *ZoneAzureConfig) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *ZoneAzureConfig) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### GetConfigCmdbDiscovery

`func (o *ZoneAzureConfig) GetConfigCmdbDiscovery() bool`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *ZoneAzureConfig) GetConfigCmdbDiscoveryOk() (*bool, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *ZoneAzureConfig) SetConfigCmdbDiscovery(v bool)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *ZoneAzureConfig) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetAzureCostingMode

`func (o *ZoneAzureConfig) GetAzureCostingMode() string`

GetAzureCostingMode returns the AzureCostingMode field if non-nil, zero value otherwise.

### GetAzureCostingModeOk

`func (o *ZoneAzureConfig) GetAzureCostingModeOk() (*string, bool)`

GetAzureCostingModeOk returns a tuple with the AzureCostingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCostingMode

`func (o *ZoneAzureConfig) SetAzureCostingMode(v string)`

SetAzureCostingMode sets AzureCostingMode field to given value.

### HasAzureCostingMode

`func (o *ZoneAzureConfig) HasAzureCostingMode() bool`

HasAzureCostingMode returns a boolean if a field has been set.

### GetClientSecretHash

`func (o *ZoneAzureConfig) GetClientSecretHash() string`

GetClientSecretHash returns the ClientSecretHash field if non-nil, zero value otherwise.

### GetClientSecretHashOk

`func (o *ZoneAzureConfig) GetClientSecretHashOk() (*string, bool)`

GetClientSecretHashOk returns a tuple with the ClientSecretHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretHash

`func (o *ZoneAzureConfig) SetClientSecretHash(v string)`

SetClientSecretHash sets ClientSecretHash field to given value.

### HasClientSecretHash

`func (o *ZoneAzureConfig) HasClientSecretHash() bool`

HasClientSecretHash returns a boolean if a field has been set.

### GetCspClientSecretHash

`func (o *ZoneAzureConfig) GetCspClientSecretHash() string`

GetCspClientSecretHash returns the CspClientSecretHash field if non-nil, zero value otherwise.

### GetCspClientSecretHashOk

`func (o *ZoneAzureConfig) GetCspClientSecretHashOk() (*string, bool)`

GetCspClientSecretHashOk returns a tuple with the CspClientSecretHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspClientSecretHash

`func (o *ZoneAzureConfig) SetCspClientSecretHash(v string)`

SetCspClientSecretHash sets CspClientSecretHash field to given value.

### HasCspClientSecretHash

`func (o *ZoneAzureConfig) HasCspClientSecretHash() bool`

HasCspClientSecretHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



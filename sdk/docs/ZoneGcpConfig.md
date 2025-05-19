# ZoneGcpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | Pointer to **string** |  | [optional] 
**ClientEmail** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**GoogleRegionId** | Pointer to **string** |  | [optional] 
**ImportExisting** | Pointer to **string** |  | [optional] 
**ApplianceUrl** | Pointer to **string** |  | [optional] 
**DatacenterName** | Pointer to **string** |  | [optional] 
**NetworkServerId** | Pointer to **string** |  | [optional] 
**NetworkServer** | Pointer to [**ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer**](ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer.md) |  | [optional] 
**SecurityServer** | Pointer to **string** |  | [optional] 
**CertificateProvider** | Pointer to **string** |  | [optional] 
**BackupMode** | Pointer to **string** |  | [optional] 
**ReplicationMode** | Pointer to **string** |  | [optional] 
**DnsIntegrationId** | Pointer to **string** |  | [optional] 
**ServiceRegistryId** | Pointer to **string** |  | [optional] 
**ConfigManagementId** | Pointer to **string** |  | [optional] 
**PrivateKeyHash** | Pointer to **string** |  | [optional] 

## Methods

### NewZoneGcpConfig

`func NewZoneGcpConfig() *ZoneGcpConfig`

NewZoneGcpConfig instantiates a new ZoneGcpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneGcpConfigWithDefaults

`func NewZoneGcpConfigWithDefaults() *ZoneGcpConfig`

NewZoneGcpConfigWithDefaults instantiates a new ZoneGcpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *ZoneGcpConfig) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ZoneGcpConfig) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ZoneGcpConfig) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *ZoneGcpConfig) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetClientEmail

`func (o *ZoneGcpConfig) GetClientEmail() string`

GetClientEmail returns the ClientEmail field if non-nil, zero value otherwise.

### GetClientEmailOk

`func (o *ZoneGcpConfig) GetClientEmailOk() (*string, bool)`

GetClientEmailOk returns a tuple with the ClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmail

`func (o *ZoneGcpConfig) SetClientEmail(v string)`

SetClientEmail sets ClientEmail field to given value.

### HasClientEmail

`func (o *ZoneGcpConfig) HasClientEmail() bool`

HasClientEmail returns a boolean if a field has been set.

### GetProjectId

`func (o *ZoneGcpConfig) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ZoneGcpConfig) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ZoneGcpConfig) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ZoneGcpConfig) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetGoogleRegionId

`func (o *ZoneGcpConfig) GetGoogleRegionId() string`

GetGoogleRegionId returns the GoogleRegionId field if non-nil, zero value otherwise.

### GetGoogleRegionIdOk

`func (o *ZoneGcpConfig) GetGoogleRegionIdOk() (*string, bool)`

GetGoogleRegionIdOk returns a tuple with the GoogleRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleRegionId

`func (o *ZoneGcpConfig) SetGoogleRegionId(v string)`

SetGoogleRegionId sets GoogleRegionId field to given value.

### HasGoogleRegionId

`func (o *ZoneGcpConfig) HasGoogleRegionId() bool`

HasGoogleRegionId returns a boolean if a field has been set.

### GetImportExisting

`func (o *ZoneGcpConfig) GetImportExisting() string`

GetImportExisting returns the ImportExisting field if non-nil, zero value otherwise.

### GetImportExistingOk

`func (o *ZoneGcpConfig) GetImportExistingOk() (*string, bool)`

GetImportExistingOk returns a tuple with the ImportExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportExisting

`func (o *ZoneGcpConfig) SetImportExisting(v string)`

SetImportExisting sets ImportExisting field to given value.

### HasImportExisting

`func (o *ZoneGcpConfig) HasImportExisting() bool`

HasImportExisting returns a boolean if a field has been set.

### GetApplianceUrl

`func (o *ZoneGcpConfig) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ZoneGcpConfig) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ZoneGcpConfig) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ZoneGcpConfig) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *ZoneGcpConfig) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *ZoneGcpConfig) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *ZoneGcpConfig) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *ZoneGcpConfig) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *ZoneGcpConfig) GetNetworkServerId() string`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *ZoneGcpConfig) GetNetworkServerIdOk() (*string, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *ZoneGcpConfig) SetNetworkServerId(v string)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *ZoneGcpConfig) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetNetworkServer

`func (o *ZoneGcpConfig) GetNetworkServer() ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *ZoneGcpConfig) GetNetworkServerOk() (*ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *ZoneGcpConfig) SetNetworkServer(v ListClouds200ResponseAllOfZonesInnerConfigAnyOfNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *ZoneGcpConfig) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetSecurityServer

`func (o *ZoneGcpConfig) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *ZoneGcpConfig) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *ZoneGcpConfig) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *ZoneGcpConfig) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *ZoneGcpConfig) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *ZoneGcpConfig) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *ZoneGcpConfig) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *ZoneGcpConfig) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetBackupMode

`func (o *ZoneGcpConfig) GetBackupMode() string`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *ZoneGcpConfig) GetBackupModeOk() (*string, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *ZoneGcpConfig) SetBackupMode(v string)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *ZoneGcpConfig) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### GetReplicationMode

`func (o *ZoneGcpConfig) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *ZoneGcpConfig) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *ZoneGcpConfig) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *ZoneGcpConfig) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### GetDnsIntegrationId

`func (o *ZoneGcpConfig) GetDnsIntegrationId() string`

GetDnsIntegrationId returns the DnsIntegrationId field if non-nil, zero value otherwise.

### GetDnsIntegrationIdOk

`func (o *ZoneGcpConfig) GetDnsIntegrationIdOk() (*string, bool)`

GetDnsIntegrationIdOk returns a tuple with the DnsIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrationId

`func (o *ZoneGcpConfig) SetDnsIntegrationId(v string)`

SetDnsIntegrationId sets DnsIntegrationId field to given value.

### HasDnsIntegrationId

`func (o *ZoneGcpConfig) HasDnsIntegrationId() bool`

HasDnsIntegrationId returns a boolean if a field has been set.

### GetServiceRegistryId

`func (o *ZoneGcpConfig) GetServiceRegistryId() string`

GetServiceRegistryId returns the ServiceRegistryId field if non-nil, zero value otherwise.

### GetServiceRegistryIdOk

`func (o *ZoneGcpConfig) GetServiceRegistryIdOk() (*string, bool)`

GetServiceRegistryIdOk returns a tuple with the ServiceRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRegistryId

`func (o *ZoneGcpConfig) SetServiceRegistryId(v string)`

SetServiceRegistryId sets ServiceRegistryId field to given value.

### HasServiceRegistryId

`func (o *ZoneGcpConfig) HasServiceRegistryId() bool`

HasServiceRegistryId returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *ZoneGcpConfig) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *ZoneGcpConfig) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *ZoneGcpConfig) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *ZoneGcpConfig) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetPrivateKeyHash

`func (o *ZoneGcpConfig) GetPrivateKeyHash() string`

GetPrivateKeyHash returns the PrivateKeyHash field if non-nil, zero value otherwise.

### GetPrivateKeyHashOk

`func (o *ZoneGcpConfig) GetPrivateKeyHashOk() (*string, bool)`

GetPrivateKeyHashOk returns a tuple with the PrivateKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyHash

`func (o *ZoneGcpConfig) SetPrivateKeyHash(v string)`

SetPrivateKeyHash sets PrivateKeyHash field to given value.

### HasPrivateKeyHash

`func (o *ZoneGcpConfig) HasPrivateKeyHash() bool`

HasPrivateKeyHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



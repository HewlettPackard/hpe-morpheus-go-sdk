# ZoneConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | Pointer to **NullableString** |  | [optional] 
**DatacenterName** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**InventoryLevel** | Pointer to **NullableString** |  | [optional] 
**ConsoleKeymap** | Pointer to **NullableString** |  | [optional] 
**BackupMode** | Pointer to **NullableString** |  | [optional] 
**CertificateProvider** | Pointer to **NullableString** |  | [optional] 
**ConfigCmdbDiscovery** | Pointer to **bool** |  | [optional] 
**EnableNetworkTypeSelection** | Pointer to **NullableString** |  | [optional] 
**KubeUrl** | Pointer to **NullableString** |  | [optional] 
**NetworkServer** | Pointer to [**AddClouds200ResponseAllOfZoneConfigAnyOfNetworkServer**](AddClouds200ResponseAllOfZoneConfigAnyOfNetworkServer.md) |  | [optional] 
**NetworkServerId** | Pointer to **string** |  | [optional] 
**ReplicationMode** | Pointer to **NullableString** |  | [optional] 
**SecurityServer** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewZoneConfigAnyOf

`func NewZoneConfigAnyOf() *ZoneConfigAnyOf`

NewZoneConfigAnyOf instantiates a new ZoneConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneConfigAnyOfWithDefaults

`func NewZoneConfigAnyOfWithDefaults() *ZoneConfigAnyOf`

NewZoneConfigAnyOfWithDefaults instantiates a new ZoneConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *ZoneConfigAnyOf) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ZoneConfigAnyOf) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ZoneConfigAnyOf) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ZoneConfigAnyOf) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### SetApplianceUrlNil

`func (o *ZoneConfigAnyOf) SetApplianceUrlNil(b bool)`

 SetApplianceUrlNil sets the value for ApplianceUrl to be an explicit nil

### UnsetApplianceUrl
`func (o *ZoneConfigAnyOf) UnsetApplianceUrl()`

UnsetApplianceUrl ensures that no value is present for ApplianceUrl, not even an explicit nil
### GetDatacenterName

`func (o *ZoneConfigAnyOf) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *ZoneConfigAnyOf) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *ZoneConfigAnyOf) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *ZoneConfigAnyOf) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### SetDatacenterNameNil

`func (o *ZoneConfigAnyOf) SetDatacenterNameNil(b bool)`

 SetDatacenterNameNil sets the value for DatacenterName to be an explicit nil

### UnsetDatacenterName
`func (o *ZoneConfigAnyOf) UnsetDatacenterName()`

UnsetDatacenterName ensures that no value is present for DatacenterName, not even an explicit nil
### GetExternalId

`func (o *ZoneConfigAnyOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ZoneConfigAnyOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ZoneConfigAnyOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ZoneConfigAnyOf) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ZoneConfigAnyOf) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ZoneConfigAnyOf) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInventoryLevel

`func (o *ZoneConfigAnyOf) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *ZoneConfigAnyOf) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *ZoneConfigAnyOf) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *ZoneConfigAnyOf) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### SetInventoryLevelNil

`func (o *ZoneConfigAnyOf) SetInventoryLevelNil(b bool)`

 SetInventoryLevelNil sets the value for InventoryLevel to be an explicit nil

### UnsetInventoryLevel
`func (o *ZoneConfigAnyOf) UnsetInventoryLevel()`

UnsetInventoryLevel ensures that no value is present for InventoryLevel, not even an explicit nil
### GetConsoleKeymap

`func (o *ZoneConfigAnyOf) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *ZoneConfigAnyOf) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *ZoneConfigAnyOf) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *ZoneConfigAnyOf) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### SetConsoleKeymapNil

`func (o *ZoneConfigAnyOf) SetConsoleKeymapNil(b bool)`

 SetConsoleKeymapNil sets the value for ConsoleKeymap to be an explicit nil

### UnsetConsoleKeymap
`func (o *ZoneConfigAnyOf) UnsetConsoleKeymap()`

UnsetConsoleKeymap ensures that no value is present for ConsoleKeymap, not even an explicit nil
### GetBackupMode

`func (o *ZoneConfigAnyOf) GetBackupMode() string`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *ZoneConfigAnyOf) GetBackupModeOk() (*string, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *ZoneConfigAnyOf) SetBackupMode(v string)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *ZoneConfigAnyOf) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### SetBackupModeNil

`func (o *ZoneConfigAnyOf) SetBackupModeNil(b bool)`

 SetBackupModeNil sets the value for BackupMode to be an explicit nil

### UnsetBackupMode
`func (o *ZoneConfigAnyOf) UnsetBackupMode()`

UnsetBackupMode ensures that no value is present for BackupMode, not even an explicit nil
### GetCertificateProvider

`func (o *ZoneConfigAnyOf) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *ZoneConfigAnyOf) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *ZoneConfigAnyOf) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *ZoneConfigAnyOf) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### SetCertificateProviderNil

`func (o *ZoneConfigAnyOf) SetCertificateProviderNil(b bool)`

 SetCertificateProviderNil sets the value for CertificateProvider to be an explicit nil

### UnsetCertificateProvider
`func (o *ZoneConfigAnyOf) UnsetCertificateProvider()`

UnsetCertificateProvider ensures that no value is present for CertificateProvider, not even an explicit nil
### GetConfigCmdbDiscovery

`func (o *ZoneConfigAnyOf) GetConfigCmdbDiscovery() bool`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *ZoneConfigAnyOf) GetConfigCmdbDiscoveryOk() (*bool, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *ZoneConfigAnyOf) SetConfigCmdbDiscovery(v bool)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *ZoneConfigAnyOf) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *ZoneConfigAnyOf) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *ZoneConfigAnyOf) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *ZoneConfigAnyOf) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *ZoneConfigAnyOf) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### SetEnableNetworkTypeSelectionNil

`func (o *ZoneConfigAnyOf) SetEnableNetworkTypeSelectionNil(b bool)`

 SetEnableNetworkTypeSelectionNil sets the value for EnableNetworkTypeSelection to be an explicit nil

### UnsetEnableNetworkTypeSelection
`func (o *ZoneConfigAnyOf) UnsetEnableNetworkTypeSelection()`

UnsetEnableNetworkTypeSelection ensures that no value is present for EnableNetworkTypeSelection, not even an explicit nil
### GetKubeUrl

`func (o *ZoneConfigAnyOf) GetKubeUrl() string`

GetKubeUrl returns the KubeUrl field if non-nil, zero value otherwise.

### GetKubeUrlOk

`func (o *ZoneConfigAnyOf) GetKubeUrlOk() (*string, bool)`

GetKubeUrlOk returns a tuple with the KubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeUrl

`func (o *ZoneConfigAnyOf) SetKubeUrl(v string)`

SetKubeUrl sets KubeUrl field to given value.

### HasKubeUrl

`func (o *ZoneConfigAnyOf) HasKubeUrl() bool`

HasKubeUrl returns a boolean if a field has been set.

### SetKubeUrlNil

`func (o *ZoneConfigAnyOf) SetKubeUrlNil(b bool)`

 SetKubeUrlNil sets the value for KubeUrl to be an explicit nil

### UnsetKubeUrl
`func (o *ZoneConfigAnyOf) UnsetKubeUrl()`

UnsetKubeUrl ensures that no value is present for KubeUrl, not even an explicit nil
### GetNetworkServer

`func (o *ZoneConfigAnyOf) GetNetworkServer() AddClouds200ResponseAllOfZoneConfigAnyOfNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *ZoneConfigAnyOf) GetNetworkServerOk() (*AddClouds200ResponseAllOfZoneConfigAnyOfNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *ZoneConfigAnyOf) SetNetworkServer(v AddClouds200ResponseAllOfZoneConfigAnyOfNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *ZoneConfigAnyOf) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *ZoneConfigAnyOf) GetNetworkServerId() string`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *ZoneConfigAnyOf) GetNetworkServerIdOk() (*string, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *ZoneConfigAnyOf) SetNetworkServerId(v string)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *ZoneConfigAnyOf) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetReplicationMode

`func (o *ZoneConfigAnyOf) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *ZoneConfigAnyOf) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *ZoneConfigAnyOf) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *ZoneConfigAnyOf) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.

### SetReplicationModeNil

`func (o *ZoneConfigAnyOf) SetReplicationModeNil(b bool)`

 SetReplicationModeNil sets the value for ReplicationMode to be an explicit nil

### UnsetReplicationMode
`func (o *ZoneConfigAnyOf) UnsetReplicationMode()`

UnsetReplicationMode ensures that no value is present for ReplicationMode, not even an explicit nil
### GetSecurityServer

`func (o *ZoneConfigAnyOf) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *ZoneConfigAnyOf) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *ZoneConfigAnyOf) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *ZoneConfigAnyOf) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### SetSecurityServerNil

`func (o *ZoneConfigAnyOf) SetSecurityServerNil(b bool)`

 SetSecurityServerNil sets the value for SecurityServer to be an explicit nil

### UnsetSecurityServer
`func (o *ZoneConfigAnyOf) UnsetSecurityServer()`

UnsetSecurityServer ensures that no value is present for SecurityServer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



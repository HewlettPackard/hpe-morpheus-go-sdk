# ZoneCreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | Pointer to **string** | The URL used by workloads provisioned in the cloud for interacting with the Morpheus appliance. | [optional] 
**DatacenterName** | Pointer to **string** | A custom name used to reference the datacenter for the cloud. | [optional] 
**ExternalId** | Pointer to **string** | The external id of the cloud | [optional] 
**InventoryLevel** | Pointer to **string** | Whether to import existing virtual machines. | [optional] 
**ConsoleKeymap** | Pointer to **string** | The keyboard layout to use for the console | [optional] 
**Endpoint** | **string** | AWS endpoint | 
**AccessKey** | Pointer to **string** | AWS access key | [optional] 
**SecretKey** | Pointer to **string** | AWS secret key | [optional] 
**UseHostCredentials** | Pointer to **string** | Whether to use the IAM profile associated with the Morpheus server or not | [optional] [default to "on"]
**EbsEncryption** | Pointer to **string** | Determines whether to configure default EBS volume encryption or not | [optional] [default to "on"]
**StsAssumeRole** | Pointer to **string** | The AWS IAM role ARN to assume for authentication | [optional] 
**ConfigManagementId** | Pointer to **string** | The id of the configuration management integration associated with the vSphere cloud. | [optional] 
**Vpc** | Pointer to **string** | The VPC ID for a specific VPC | [optional] 
**VdiGateway** | Pointer to **string** | The VDI gateway for this cloud to use for provisioning virtual desktops. | [optional] 
**CmdbConfig** | Pointer to **string** | The CMDB configuration for this cloud to use for syncing with the CMDB. | [optional] 
**ChangeManagementConfig** | Pointer to **string** | The change management configuration for this cloud to use for syncing with the change management system. | [optional] 
**NetworkMode** | Pointer to **string** | Whether to use public or private IP addresses for provisioning and managing instances in this cloud. | [optional] 
**ApiProxy** | Pointer to **string** | The API proxy to use for API calls to the cloud. | [optional] 
**Region** | Pointer to **string** | The AWS region to use for this cloud. | [optional] 
**Credentials** | Pointer to **string** | The credentials to use for this cloud. | [optional] 
**CostingBucket** | Pointer to **string** | The S3 bucket to use for storing costing reports. | [optional] 
**CostingFolder** | Pointer to **string** | The folder within the S3 bucket to use for storing costing reports. | [optional] 
**CostingReportName** | Pointer to **string** | The name of the costing report to generate. | [optional] 
**CostingKey** | Pointer to **string** | The AWS access key to use for generating costing reports. | [optional] 
**CostingSecret** | Pointer to **string** | The AWS secret key to use for generating costing reports. | [optional] 
**Domain** | Pointer to **string** | The domain to use for this cloud. | [optional] 
**Timezone** | Pointer to **string** | The timezone to use for this cloud. | [optional] 
**SecurityServer** | Pointer to **string** | The security server to use for this cloud, or &#39;off&#39; to not use a security server. | [optional] 
**Guidance** | Pointer to **string** | Optional guidance field if you want to put more info there | [optional] 
**Costing** | Pointer to **string** | Whether to enable costing for this cloud or not. | [optional] 
**ConfigCmdbDiscovery** | Pointer to **string** | The CMDB discovery configuration to use for this cloud. | [optional] 
**Logo** | Pointer to **string** | The logo to use for this cloud. | [optional] 
**DarkModeLogo** | Pointer to **string** | The logo to use for this cloud in dark mode. | [optional] 
**Proxy** | Pointer to **string** | The proxy to use for this cloud. | [optional] 
**BypassProxyForCloud** | Pointer to **string** | Whether to bypass the proxy for API calls to the cloud or not. | [optional] 
**NoProxy** | Pointer to **string** | A comma separated list of hosts to bypass the proxy for when making API calls to the cloud. | [optional] 
**UserDataLinux** | Pointer to **string** | The user data script to use for provisioning instances in this cloud. | [optional] 
**SubscriberId** | Pointer to **string** | Azure subscriber id | [optional] 
**TenantId** | Pointer to **string** | Azure tenant id | [optional] 
**ClientId** | Pointer to **string** | Azure client id | [optional] 
**ClientSecret** | Pointer to **string** | Azure client secret | [optional] 
**ResourceGroup** | Pointer to **string** | Azure resource group | [optional] 
**RpcMode** | Pointer to [**NullableZoneCreateConfigAnyOf3RpcMode**](ZoneCreateConfigAnyOf3RpcMode.md) |  | [optional] 
**CertificateProvider** | Pointer to **string** | Certificate provider | [optional] [default to "internal"]
**EnableNetworkTypeSelection** | Pointer to [**NullableZoneCreateConfigAnyOf3EnableNetworkTypeSelection**](ZoneCreateConfigAnyOf3EnableNetworkTypeSelection.md) |  | [optional] 
**ApiUrl** | **string** | The SDK URL of the vCenter server. | 
**ApiVersion** | **string** | The SDK version of the vCenter server. | 
**Datacenter** | **string** | The vSphere datacenter to add. | 
**Cluster** | Pointer to **string** | The name of the vSphere cluster | [optional] [default to "all"]
**ResourcePool** | Pointer to **string** | The name of the vSphere resource pool | [optional] 
**StorageType** | Pointer to **string** | The default vSphere VMDK type for virtual machines | [optional] [default to "thin"]
**EnableVnc** | Pointer to [**NullableZoneCreateConfigAnyOf3EnableVnc**](ZoneCreateConfigAnyOf3EnableVnc.md) |  | [optional] 
**HideHostSelection** | Pointer to [**NullableZoneCreateConfigAnyOf3HideHostSelection**](ZoneCreateConfigAnyOf3HideHostSelection.md) |  | [optional] 
**EnableDiskTypeSelection** | Pointer to [**NullableZoneCreateConfigAnyOf3EnableDiskTypeSelection**](ZoneCreateConfigAnyOf3EnableDiskTypeSelection.md) |  | [optional] 
**EnableStorageTypeSelection** | Pointer to [**NullableZoneCreateConfigAnyOf3EnableStorageTypeSelection**](ZoneCreateConfigAnyOf3EnableStorageTypeSelection.md) |  | [optional] 
**Username** | Pointer to **string** | Username. | [optional] 
**Password** | Pointer to **string** | Password to apply to the user | [optional] 

## Methods

### NewZoneCreateConfig

`func NewZoneCreateConfig(endpoint string, apiUrl string, apiVersion string, datacenter string, ) *ZoneCreateConfig`

NewZoneCreateConfig instantiates a new ZoneCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCreateConfigWithDefaults

`func NewZoneCreateConfigWithDefaults() *ZoneCreateConfig`

NewZoneCreateConfigWithDefaults instantiates a new ZoneCreateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *ZoneCreateConfig) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ZoneCreateConfig) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ZoneCreateConfig) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ZoneCreateConfig) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *ZoneCreateConfig) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *ZoneCreateConfig) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *ZoneCreateConfig) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *ZoneCreateConfig) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *ZoneCreateConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ZoneCreateConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ZoneCreateConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ZoneCreateConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *ZoneCreateConfig) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *ZoneCreateConfig) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *ZoneCreateConfig) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *ZoneCreateConfig) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *ZoneCreateConfig) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *ZoneCreateConfig) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *ZoneCreateConfig) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *ZoneCreateConfig) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetEndpoint

`func (o *ZoneCreateConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ZoneCreateConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ZoneCreateConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetAccessKey

`func (o *ZoneCreateConfig) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ZoneCreateConfig) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ZoneCreateConfig) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ZoneCreateConfig) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *ZoneCreateConfig) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ZoneCreateConfig) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ZoneCreateConfig) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ZoneCreateConfig) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *ZoneCreateConfig) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *ZoneCreateConfig) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *ZoneCreateConfig) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *ZoneCreateConfig) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *ZoneCreateConfig) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *ZoneCreateConfig) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *ZoneCreateConfig) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *ZoneCreateConfig) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *ZoneCreateConfig) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *ZoneCreateConfig) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *ZoneCreateConfig) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *ZoneCreateConfig) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *ZoneCreateConfig) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *ZoneCreateConfig) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *ZoneCreateConfig) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *ZoneCreateConfig) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetVpc

`func (o *ZoneCreateConfig) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *ZoneCreateConfig) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *ZoneCreateConfig) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *ZoneCreateConfig) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetVdiGateway

`func (o *ZoneCreateConfig) GetVdiGateway() string`

GetVdiGateway returns the VdiGateway field if non-nil, zero value otherwise.

### GetVdiGatewayOk

`func (o *ZoneCreateConfig) GetVdiGatewayOk() (*string, bool)`

GetVdiGatewayOk returns a tuple with the VdiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiGateway

`func (o *ZoneCreateConfig) SetVdiGateway(v string)`

SetVdiGateway sets VdiGateway field to given value.

### HasVdiGateway

`func (o *ZoneCreateConfig) HasVdiGateway() bool`

HasVdiGateway returns a boolean if a field has been set.

### GetCmdbConfig

`func (o *ZoneCreateConfig) GetCmdbConfig() string`

GetCmdbConfig returns the CmdbConfig field if non-nil, zero value otherwise.

### GetCmdbConfigOk

`func (o *ZoneCreateConfig) GetCmdbConfigOk() (*string, bool)`

GetCmdbConfigOk returns a tuple with the CmdbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdbConfig

`func (o *ZoneCreateConfig) SetCmdbConfig(v string)`

SetCmdbConfig sets CmdbConfig field to given value.

### HasCmdbConfig

`func (o *ZoneCreateConfig) HasCmdbConfig() bool`

HasCmdbConfig returns a boolean if a field has been set.

### GetChangeManagementConfig

`func (o *ZoneCreateConfig) GetChangeManagementConfig() string`

GetChangeManagementConfig returns the ChangeManagementConfig field if non-nil, zero value otherwise.

### GetChangeManagementConfigOk

`func (o *ZoneCreateConfig) GetChangeManagementConfigOk() (*string, bool)`

GetChangeManagementConfigOk returns a tuple with the ChangeManagementConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeManagementConfig

`func (o *ZoneCreateConfig) SetChangeManagementConfig(v string)`

SetChangeManagementConfig sets ChangeManagementConfig field to given value.

### HasChangeManagementConfig

`func (o *ZoneCreateConfig) HasChangeManagementConfig() bool`

HasChangeManagementConfig returns a boolean if a field has been set.

### GetNetworkMode

`func (o *ZoneCreateConfig) GetNetworkMode() string`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *ZoneCreateConfig) GetNetworkModeOk() (*string, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *ZoneCreateConfig) SetNetworkMode(v string)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *ZoneCreateConfig) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### GetApiProxy

`func (o *ZoneCreateConfig) GetApiProxy() string`

GetApiProxy returns the ApiProxy field if non-nil, zero value otherwise.

### GetApiProxyOk

`func (o *ZoneCreateConfig) GetApiProxyOk() (*string, bool)`

GetApiProxyOk returns a tuple with the ApiProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProxy

`func (o *ZoneCreateConfig) SetApiProxy(v string)`

SetApiProxy sets ApiProxy field to given value.

### HasApiProxy

`func (o *ZoneCreateConfig) HasApiProxy() bool`

HasApiProxy returns a boolean if a field has been set.

### GetRegion

`func (o *ZoneCreateConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ZoneCreateConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ZoneCreateConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ZoneCreateConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCredentials

`func (o *ZoneCreateConfig) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ZoneCreateConfig) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ZoneCreateConfig) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ZoneCreateConfig) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetCostingBucket

`func (o *ZoneCreateConfig) GetCostingBucket() string`

GetCostingBucket returns the CostingBucket field if non-nil, zero value otherwise.

### GetCostingBucketOk

`func (o *ZoneCreateConfig) GetCostingBucketOk() (*string, bool)`

GetCostingBucketOk returns a tuple with the CostingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucket

`func (o *ZoneCreateConfig) SetCostingBucket(v string)`

SetCostingBucket sets CostingBucket field to given value.

### HasCostingBucket

`func (o *ZoneCreateConfig) HasCostingBucket() bool`

HasCostingBucket returns a boolean if a field has been set.

### GetCostingFolder

`func (o *ZoneCreateConfig) GetCostingFolder() string`

GetCostingFolder returns the CostingFolder field if non-nil, zero value otherwise.

### GetCostingFolderOk

`func (o *ZoneCreateConfig) GetCostingFolderOk() (*string, bool)`

GetCostingFolderOk returns a tuple with the CostingFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingFolder

`func (o *ZoneCreateConfig) SetCostingFolder(v string)`

SetCostingFolder sets CostingFolder field to given value.

### HasCostingFolder

`func (o *ZoneCreateConfig) HasCostingFolder() bool`

HasCostingFolder returns a boolean if a field has been set.

### GetCostingReportName

`func (o *ZoneCreateConfig) GetCostingReportName() string`

GetCostingReportName returns the CostingReportName field if non-nil, zero value otherwise.

### GetCostingReportNameOk

`func (o *ZoneCreateConfig) GetCostingReportNameOk() (*string, bool)`

GetCostingReportNameOk returns a tuple with the CostingReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReportName

`func (o *ZoneCreateConfig) SetCostingReportName(v string)`

SetCostingReportName sets CostingReportName field to given value.

### HasCostingReportName

`func (o *ZoneCreateConfig) HasCostingReportName() bool`

HasCostingReportName returns a boolean if a field has been set.

### GetCostingKey

`func (o *ZoneCreateConfig) GetCostingKey() string`

GetCostingKey returns the CostingKey field if non-nil, zero value otherwise.

### GetCostingKeyOk

`func (o *ZoneCreateConfig) GetCostingKeyOk() (*string, bool)`

GetCostingKeyOk returns a tuple with the CostingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingKey

`func (o *ZoneCreateConfig) SetCostingKey(v string)`

SetCostingKey sets CostingKey field to given value.

### HasCostingKey

`func (o *ZoneCreateConfig) HasCostingKey() bool`

HasCostingKey returns a boolean if a field has been set.

### GetCostingSecret

`func (o *ZoneCreateConfig) GetCostingSecret() string`

GetCostingSecret returns the CostingSecret field if non-nil, zero value otherwise.

### GetCostingSecretOk

`func (o *ZoneCreateConfig) GetCostingSecretOk() (*string, bool)`

GetCostingSecretOk returns a tuple with the CostingSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecret

`func (o *ZoneCreateConfig) SetCostingSecret(v string)`

SetCostingSecret sets CostingSecret field to given value.

### HasCostingSecret

`func (o *ZoneCreateConfig) HasCostingSecret() bool`

HasCostingSecret returns a boolean if a field has been set.

### GetDomain

`func (o *ZoneCreateConfig) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ZoneCreateConfig) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ZoneCreateConfig) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ZoneCreateConfig) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTimezone

`func (o *ZoneCreateConfig) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ZoneCreateConfig) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ZoneCreateConfig) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ZoneCreateConfig) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSecurityServer

`func (o *ZoneCreateConfig) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *ZoneCreateConfig) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *ZoneCreateConfig) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *ZoneCreateConfig) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetGuidance

`func (o *ZoneCreateConfig) GetGuidance() string`

GetGuidance returns the Guidance field if non-nil, zero value otherwise.

### GetGuidanceOk

`func (o *ZoneCreateConfig) GetGuidanceOk() (*string, bool)`

GetGuidanceOk returns a tuple with the Guidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidance

`func (o *ZoneCreateConfig) SetGuidance(v string)`

SetGuidance sets Guidance field to given value.

### HasGuidance

`func (o *ZoneCreateConfig) HasGuidance() bool`

HasGuidance returns a boolean if a field has been set.

### GetCosting

`func (o *ZoneCreateConfig) GetCosting() string`

GetCosting returns the Costing field if non-nil, zero value otherwise.

### GetCostingOk

`func (o *ZoneCreateConfig) GetCostingOk() (*string, bool)`

GetCostingOk returns a tuple with the Costing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosting

`func (o *ZoneCreateConfig) SetCosting(v string)`

SetCosting sets Costing field to given value.

### HasCosting

`func (o *ZoneCreateConfig) HasCosting() bool`

HasCosting returns a boolean if a field has been set.

### GetConfigCmdbDiscovery

`func (o *ZoneCreateConfig) GetConfigCmdbDiscovery() string`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *ZoneCreateConfig) GetConfigCmdbDiscoveryOk() (*string, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *ZoneCreateConfig) SetConfigCmdbDiscovery(v string)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *ZoneCreateConfig) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetLogo

`func (o *ZoneCreateConfig) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ZoneCreateConfig) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ZoneCreateConfig) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ZoneCreateConfig) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetDarkModeLogo

`func (o *ZoneCreateConfig) GetDarkModeLogo() string`

GetDarkModeLogo returns the DarkModeLogo field if non-nil, zero value otherwise.

### GetDarkModeLogoOk

`func (o *ZoneCreateConfig) GetDarkModeLogoOk() (*string, bool)`

GetDarkModeLogoOk returns a tuple with the DarkModeLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkModeLogo

`func (o *ZoneCreateConfig) SetDarkModeLogo(v string)`

SetDarkModeLogo sets DarkModeLogo field to given value.

### HasDarkModeLogo

`func (o *ZoneCreateConfig) HasDarkModeLogo() bool`

HasDarkModeLogo returns a boolean if a field has been set.

### GetProxy

`func (o *ZoneCreateConfig) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ZoneCreateConfig) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ZoneCreateConfig) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ZoneCreateConfig) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetBypassProxyForCloud

`func (o *ZoneCreateConfig) GetBypassProxyForCloud() string`

GetBypassProxyForCloud returns the BypassProxyForCloud field if non-nil, zero value otherwise.

### GetBypassProxyForCloudOk

`func (o *ZoneCreateConfig) GetBypassProxyForCloudOk() (*string, bool)`

GetBypassProxyForCloudOk returns a tuple with the BypassProxyForCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassProxyForCloud

`func (o *ZoneCreateConfig) SetBypassProxyForCloud(v string)`

SetBypassProxyForCloud sets BypassProxyForCloud field to given value.

### HasBypassProxyForCloud

`func (o *ZoneCreateConfig) HasBypassProxyForCloud() bool`

HasBypassProxyForCloud returns a boolean if a field has been set.

### GetNoProxy

`func (o *ZoneCreateConfig) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *ZoneCreateConfig) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *ZoneCreateConfig) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *ZoneCreateConfig) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetUserDataLinux

`func (o *ZoneCreateConfig) GetUserDataLinux() string`

GetUserDataLinux returns the UserDataLinux field if non-nil, zero value otherwise.

### GetUserDataLinuxOk

`func (o *ZoneCreateConfig) GetUserDataLinuxOk() (*string, bool)`

GetUserDataLinuxOk returns a tuple with the UserDataLinux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataLinux

`func (o *ZoneCreateConfig) SetUserDataLinux(v string)`

SetUserDataLinux sets UserDataLinux field to given value.

### HasUserDataLinux

`func (o *ZoneCreateConfig) HasUserDataLinux() bool`

HasUserDataLinux returns a boolean if a field has been set.

### GetSubscriberId

`func (o *ZoneCreateConfig) GetSubscriberId() string`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *ZoneCreateConfig) GetSubscriberIdOk() (*string, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *ZoneCreateConfig) SetSubscriberId(v string)`

SetSubscriberId sets SubscriberId field to given value.

### HasSubscriberId

`func (o *ZoneCreateConfig) HasSubscriberId() bool`

HasSubscriberId returns a boolean if a field has been set.

### GetTenantId

`func (o *ZoneCreateConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ZoneCreateConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ZoneCreateConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ZoneCreateConfig) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetClientId

`func (o *ZoneCreateConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ZoneCreateConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ZoneCreateConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ZoneCreateConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ZoneCreateConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ZoneCreateConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ZoneCreateConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ZoneCreateConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetResourceGroup

`func (o *ZoneCreateConfig) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *ZoneCreateConfig) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *ZoneCreateConfig) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *ZoneCreateConfig) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetRpcMode

`func (o *ZoneCreateConfig) GetRpcMode() ZoneCreateConfigAnyOf3RpcMode`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *ZoneCreateConfig) GetRpcModeOk() (*ZoneCreateConfigAnyOf3RpcMode, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *ZoneCreateConfig) SetRpcMode(v ZoneCreateConfigAnyOf3RpcMode)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *ZoneCreateConfig) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### SetRpcModeNil

`func (o *ZoneCreateConfig) SetRpcModeNil(b bool)`

 SetRpcModeNil sets the value for RpcMode to be an explicit nil

### UnsetRpcMode
`func (o *ZoneCreateConfig) UnsetRpcMode()`

UnsetRpcMode ensures that no value is present for RpcMode, not even an explicit nil
### GetCertificateProvider

`func (o *ZoneCreateConfig) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *ZoneCreateConfig) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *ZoneCreateConfig) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *ZoneCreateConfig) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *ZoneCreateConfig) GetEnableNetworkTypeSelection() ZoneCreateConfigAnyOf3EnableNetworkTypeSelection`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *ZoneCreateConfig) GetEnableNetworkTypeSelectionOk() (*ZoneCreateConfigAnyOf3EnableNetworkTypeSelection, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *ZoneCreateConfig) SetEnableNetworkTypeSelection(v ZoneCreateConfigAnyOf3EnableNetworkTypeSelection)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *ZoneCreateConfig) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### SetEnableNetworkTypeSelectionNil

`func (o *ZoneCreateConfig) SetEnableNetworkTypeSelectionNil(b bool)`

 SetEnableNetworkTypeSelectionNil sets the value for EnableNetworkTypeSelection to be an explicit nil

### UnsetEnableNetworkTypeSelection
`func (o *ZoneCreateConfig) UnsetEnableNetworkTypeSelection()`

UnsetEnableNetworkTypeSelection ensures that no value is present for EnableNetworkTypeSelection, not even an explicit nil
### GetApiUrl

`func (o *ZoneCreateConfig) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ZoneCreateConfig) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ZoneCreateConfig) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.


### GetApiVersion

`func (o *ZoneCreateConfig) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ZoneCreateConfig) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ZoneCreateConfig) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetDatacenter

`func (o *ZoneCreateConfig) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *ZoneCreateConfig) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *ZoneCreateConfig) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.


### GetCluster

`func (o *ZoneCreateConfig) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ZoneCreateConfig) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ZoneCreateConfig) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ZoneCreateConfig) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetResourcePool

`func (o *ZoneCreateConfig) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ZoneCreateConfig) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ZoneCreateConfig) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ZoneCreateConfig) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetStorageType

`func (o *ZoneCreateConfig) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *ZoneCreateConfig) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *ZoneCreateConfig) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *ZoneCreateConfig) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetEnableVnc

`func (o *ZoneCreateConfig) GetEnableVnc() ZoneCreateConfigAnyOf3EnableVnc`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *ZoneCreateConfig) GetEnableVncOk() (*ZoneCreateConfigAnyOf3EnableVnc, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *ZoneCreateConfig) SetEnableVnc(v ZoneCreateConfigAnyOf3EnableVnc)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *ZoneCreateConfig) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### SetEnableVncNil

`func (o *ZoneCreateConfig) SetEnableVncNil(b bool)`

 SetEnableVncNil sets the value for EnableVnc to be an explicit nil

### UnsetEnableVnc
`func (o *ZoneCreateConfig) UnsetEnableVnc()`

UnsetEnableVnc ensures that no value is present for EnableVnc, not even an explicit nil
### GetHideHostSelection

`func (o *ZoneCreateConfig) GetHideHostSelection() ZoneCreateConfigAnyOf3HideHostSelection`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *ZoneCreateConfig) GetHideHostSelectionOk() (*ZoneCreateConfigAnyOf3HideHostSelection, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *ZoneCreateConfig) SetHideHostSelection(v ZoneCreateConfigAnyOf3HideHostSelection)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *ZoneCreateConfig) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### SetHideHostSelectionNil

`func (o *ZoneCreateConfig) SetHideHostSelectionNil(b bool)`

 SetHideHostSelectionNil sets the value for HideHostSelection to be an explicit nil

### UnsetHideHostSelection
`func (o *ZoneCreateConfig) UnsetHideHostSelection()`

UnsetHideHostSelection ensures that no value is present for HideHostSelection, not even an explicit nil
### GetEnableDiskTypeSelection

`func (o *ZoneCreateConfig) GetEnableDiskTypeSelection() ZoneCreateConfigAnyOf3EnableDiskTypeSelection`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *ZoneCreateConfig) GetEnableDiskTypeSelectionOk() (*ZoneCreateConfigAnyOf3EnableDiskTypeSelection, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *ZoneCreateConfig) SetEnableDiskTypeSelection(v ZoneCreateConfigAnyOf3EnableDiskTypeSelection)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *ZoneCreateConfig) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### SetEnableDiskTypeSelectionNil

`func (o *ZoneCreateConfig) SetEnableDiskTypeSelectionNil(b bool)`

 SetEnableDiskTypeSelectionNil sets the value for EnableDiskTypeSelection to be an explicit nil

### UnsetEnableDiskTypeSelection
`func (o *ZoneCreateConfig) UnsetEnableDiskTypeSelection()`

UnsetEnableDiskTypeSelection ensures that no value is present for EnableDiskTypeSelection, not even an explicit nil
### GetEnableStorageTypeSelection

`func (o *ZoneCreateConfig) GetEnableStorageTypeSelection() ZoneCreateConfigAnyOf3EnableStorageTypeSelection`

GetEnableStorageTypeSelection returns the EnableStorageTypeSelection field if non-nil, zero value otherwise.

### GetEnableStorageTypeSelectionOk

`func (o *ZoneCreateConfig) GetEnableStorageTypeSelectionOk() (*ZoneCreateConfigAnyOf3EnableStorageTypeSelection, bool)`

GetEnableStorageTypeSelectionOk returns a tuple with the EnableStorageTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStorageTypeSelection

`func (o *ZoneCreateConfig) SetEnableStorageTypeSelection(v ZoneCreateConfigAnyOf3EnableStorageTypeSelection)`

SetEnableStorageTypeSelection sets EnableStorageTypeSelection field to given value.

### HasEnableStorageTypeSelection

`func (o *ZoneCreateConfig) HasEnableStorageTypeSelection() bool`

HasEnableStorageTypeSelection returns a boolean if a field has been set.

### SetEnableStorageTypeSelectionNil

`func (o *ZoneCreateConfig) SetEnableStorageTypeSelectionNil(b bool)`

 SetEnableStorageTypeSelectionNil sets the value for EnableStorageTypeSelection to be an explicit nil

### UnsetEnableStorageTypeSelection
`func (o *ZoneCreateConfig) UnsetEnableStorageTypeSelection()`

UnsetEnableStorageTypeSelection ensures that no value is present for EnableStorageTypeSelection, not even an explicit nil
### GetUsername

`func (o *ZoneCreateConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ZoneCreateConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ZoneCreateConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ZoneCreateConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ZoneCreateConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ZoneCreateConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ZoneCreateConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ZoneCreateConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



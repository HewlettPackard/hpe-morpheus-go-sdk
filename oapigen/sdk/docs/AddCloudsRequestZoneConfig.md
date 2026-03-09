# AddCloudsRequestZoneConfig

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
**RpcMode** | Pointer to **string** |  | [optional] 
**CertificateProvider** | Pointer to **string** | Certificate provider | [optional] [default to "internal"]
**EnableNetworkTypeSelection** | Pointer to **string** |  | [optional] 
**ApiUrl** | **string** | The SDK URL of the vCenter server. | 
**ApiVersion** | **string** | The SDK version of the vCenter server. | 
**Datacenter** | **string** | The vSphere datacenter to add. | 
**Cluster** | Pointer to **string** | The name of the vSphere cluster | [optional] [default to "all"]
**ResourcePool** | Pointer to **string** | The name of the vSphere resource pool | [optional] 
**StorageType** | Pointer to **string** | The default vSphere VMDK type for virtual machines | [optional] [default to "thin"]
**EnableVnc** | Pointer to **string** |  | [optional] 
**HideHostSelection** | Pointer to **string** |  | [optional] 
**EnableDiskTypeSelection** | Pointer to **string** |  | [optional] 
**EnableStorageTypeSelection** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** | Username. | [optional] 
**Password** | Pointer to **string** | Password to apply to the user | [optional] 

## Methods

### NewAddCloudsRequestZoneConfig

`func NewAddCloudsRequestZoneConfig(endpoint string, apiUrl string, apiVersion string, datacenter string, ) *AddCloudsRequestZoneConfig`

NewAddCloudsRequestZoneConfig instantiates a new AddCloudsRequestZoneConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCloudsRequestZoneConfigWithDefaults

`func NewAddCloudsRequestZoneConfigWithDefaults() *AddCloudsRequestZoneConfig`

NewAddCloudsRequestZoneConfigWithDefaults instantiates a new AddCloudsRequestZoneConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *AddCloudsRequestZoneConfig) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *AddCloudsRequestZoneConfig) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *AddCloudsRequestZoneConfig) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *AddCloudsRequestZoneConfig) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *AddCloudsRequestZoneConfig) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *AddCloudsRequestZoneConfig) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *AddCloudsRequestZoneConfig) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *AddCloudsRequestZoneConfig) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *AddCloudsRequestZoneConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddCloudsRequestZoneConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddCloudsRequestZoneConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddCloudsRequestZoneConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *AddCloudsRequestZoneConfig) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *AddCloudsRequestZoneConfig) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *AddCloudsRequestZoneConfig) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *AddCloudsRequestZoneConfig) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *AddCloudsRequestZoneConfig) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *AddCloudsRequestZoneConfig) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *AddCloudsRequestZoneConfig) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *AddCloudsRequestZoneConfig) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetEndpoint

`func (o *AddCloudsRequestZoneConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AddCloudsRequestZoneConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AddCloudsRequestZoneConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetAccessKey

`func (o *AddCloudsRequestZoneConfig) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AddCloudsRequestZoneConfig) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AddCloudsRequestZoneConfig) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *AddCloudsRequestZoneConfig) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *AddCloudsRequestZoneConfig) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AddCloudsRequestZoneConfig) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AddCloudsRequestZoneConfig) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AddCloudsRequestZoneConfig) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *AddCloudsRequestZoneConfig) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *AddCloudsRequestZoneConfig) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *AddCloudsRequestZoneConfig) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *AddCloudsRequestZoneConfig) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *AddCloudsRequestZoneConfig) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *AddCloudsRequestZoneConfig) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *AddCloudsRequestZoneConfig) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *AddCloudsRequestZoneConfig) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *AddCloudsRequestZoneConfig) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *AddCloudsRequestZoneConfig) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *AddCloudsRequestZoneConfig) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *AddCloudsRequestZoneConfig) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *AddCloudsRequestZoneConfig) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *AddCloudsRequestZoneConfig) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *AddCloudsRequestZoneConfig) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *AddCloudsRequestZoneConfig) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetVpc

`func (o *AddCloudsRequestZoneConfig) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *AddCloudsRequestZoneConfig) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *AddCloudsRequestZoneConfig) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *AddCloudsRequestZoneConfig) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetVdiGateway

`func (o *AddCloudsRequestZoneConfig) GetVdiGateway() string`

GetVdiGateway returns the VdiGateway field if non-nil, zero value otherwise.

### GetVdiGatewayOk

`func (o *AddCloudsRequestZoneConfig) GetVdiGatewayOk() (*string, bool)`

GetVdiGatewayOk returns a tuple with the VdiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiGateway

`func (o *AddCloudsRequestZoneConfig) SetVdiGateway(v string)`

SetVdiGateway sets VdiGateway field to given value.

### HasVdiGateway

`func (o *AddCloudsRequestZoneConfig) HasVdiGateway() bool`

HasVdiGateway returns a boolean if a field has been set.

### GetCmdbConfig

`func (o *AddCloudsRequestZoneConfig) GetCmdbConfig() string`

GetCmdbConfig returns the CmdbConfig field if non-nil, zero value otherwise.

### GetCmdbConfigOk

`func (o *AddCloudsRequestZoneConfig) GetCmdbConfigOk() (*string, bool)`

GetCmdbConfigOk returns a tuple with the CmdbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdbConfig

`func (o *AddCloudsRequestZoneConfig) SetCmdbConfig(v string)`

SetCmdbConfig sets CmdbConfig field to given value.

### HasCmdbConfig

`func (o *AddCloudsRequestZoneConfig) HasCmdbConfig() bool`

HasCmdbConfig returns a boolean if a field has been set.

### GetChangeManagementConfig

`func (o *AddCloudsRequestZoneConfig) GetChangeManagementConfig() string`

GetChangeManagementConfig returns the ChangeManagementConfig field if non-nil, zero value otherwise.

### GetChangeManagementConfigOk

`func (o *AddCloudsRequestZoneConfig) GetChangeManagementConfigOk() (*string, bool)`

GetChangeManagementConfigOk returns a tuple with the ChangeManagementConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeManagementConfig

`func (o *AddCloudsRequestZoneConfig) SetChangeManagementConfig(v string)`

SetChangeManagementConfig sets ChangeManagementConfig field to given value.

### HasChangeManagementConfig

`func (o *AddCloudsRequestZoneConfig) HasChangeManagementConfig() bool`

HasChangeManagementConfig returns a boolean if a field has been set.

### GetNetworkMode

`func (o *AddCloudsRequestZoneConfig) GetNetworkMode() string`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *AddCloudsRequestZoneConfig) GetNetworkModeOk() (*string, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *AddCloudsRequestZoneConfig) SetNetworkMode(v string)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *AddCloudsRequestZoneConfig) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### GetApiProxy

`func (o *AddCloudsRequestZoneConfig) GetApiProxy() string`

GetApiProxy returns the ApiProxy field if non-nil, zero value otherwise.

### GetApiProxyOk

`func (o *AddCloudsRequestZoneConfig) GetApiProxyOk() (*string, bool)`

GetApiProxyOk returns a tuple with the ApiProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProxy

`func (o *AddCloudsRequestZoneConfig) SetApiProxy(v string)`

SetApiProxy sets ApiProxy field to given value.

### HasApiProxy

`func (o *AddCloudsRequestZoneConfig) HasApiProxy() bool`

HasApiProxy returns a boolean if a field has been set.

### GetRegion

`func (o *AddCloudsRequestZoneConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddCloudsRequestZoneConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddCloudsRequestZoneConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddCloudsRequestZoneConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCredentials

`func (o *AddCloudsRequestZoneConfig) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AddCloudsRequestZoneConfig) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AddCloudsRequestZoneConfig) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AddCloudsRequestZoneConfig) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetCostingBucket

`func (o *AddCloudsRequestZoneConfig) GetCostingBucket() string`

GetCostingBucket returns the CostingBucket field if non-nil, zero value otherwise.

### GetCostingBucketOk

`func (o *AddCloudsRequestZoneConfig) GetCostingBucketOk() (*string, bool)`

GetCostingBucketOk returns a tuple with the CostingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucket

`func (o *AddCloudsRequestZoneConfig) SetCostingBucket(v string)`

SetCostingBucket sets CostingBucket field to given value.

### HasCostingBucket

`func (o *AddCloudsRequestZoneConfig) HasCostingBucket() bool`

HasCostingBucket returns a boolean if a field has been set.

### GetCostingFolder

`func (o *AddCloudsRequestZoneConfig) GetCostingFolder() string`

GetCostingFolder returns the CostingFolder field if non-nil, zero value otherwise.

### GetCostingFolderOk

`func (o *AddCloudsRequestZoneConfig) GetCostingFolderOk() (*string, bool)`

GetCostingFolderOk returns a tuple with the CostingFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingFolder

`func (o *AddCloudsRequestZoneConfig) SetCostingFolder(v string)`

SetCostingFolder sets CostingFolder field to given value.

### HasCostingFolder

`func (o *AddCloudsRequestZoneConfig) HasCostingFolder() bool`

HasCostingFolder returns a boolean if a field has been set.

### GetCostingReportName

`func (o *AddCloudsRequestZoneConfig) GetCostingReportName() string`

GetCostingReportName returns the CostingReportName field if non-nil, zero value otherwise.

### GetCostingReportNameOk

`func (o *AddCloudsRequestZoneConfig) GetCostingReportNameOk() (*string, bool)`

GetCostingReportNameOk returns a tuple with the CostingReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReportName

`func (o *AddCloudsRequestZoneConfig) SetCostingReportName(v string)`

SetCostingReportName sets CostingReportName field to given value.

### HasCostingReportName

`func (o *AddCloudsRequestZoneConfig) HasCostingReportName() bool`

HasCostingReportName returns a boolean if a field has been set.

### GetCostingKey

`func (o *AddCloudsRequestZoneConfig) GetCostingKey() string`

GetCostingKey returns the CostingKey field if non-nil, zero value otherwise.

### GetCostingKeyOk

`func (o *AddCloudsRequestZoneConfig) GetCostingKeyOk() (*string, bool)`

GetCostingKeyOk returns a tuple with the CostingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingKey

`func (o *AddCloudsRequestZoneConfig) SetCostingKey(v string)`

SetCostingKey sets CostingKey field to given value.

### HasCostingKey

`func (o *AddCloudsRequestZoneConfig) HasCostingKey() bool`

HasCostingKey returns a boolean if a field has been set.

### GetCostingSecret

`func (o *AddCloudsRequestZoneConfig) GetCostingSecret() string`

GetCostingSecret returns the CostingSecret field if non-nil, zero value otherwise.

### GetCostingSecretOk

`func (o *AddCloudsRequestZoneConfig) GetCostingSecretOk() (*string, bool)`

GetCostingSecretOk returns a tuple with the CostingSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecret

`func (o *AddCloudsRequestZoneConfig) SetCostingSecret(v string)`

SetCostingSecret sets CostingSecret field to given value.

### HasCostingSecret

`func (o *AddCloudsRequestZoneConfig) HasCostingSecret() bool`

HasCostingSecret returns a boolean if a field has been set.

### GetDomain

`func (o *AddCloudsRequestZoneConfig) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AddCloudsRequestZoneConfig) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AddCloudsRequestZoneConfig) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AddCloudsRequestZoneConfig) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTimezone

`func (o *AddCloudsRequestZoneConfig) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AddCloudsRequestZoneConfig) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AddCloudsRequestZoneConfig) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AddCloudsRequestZoneConfig) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSecurityServer

`func (o *AddCloudsRequestZoneConfig) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *AddCloudsRequestZoneConfig) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *AddCloudsRequestZoneConfig) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *AddCloudsRequestZoneConfig) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetGuidance

`func (o *AddCloudsRequestZoneConfig) GetGuidance() string`

GetGuidance returns the Guidance field if non-nil, zero value otherwise.

### GetGuidanceOk

`func (o *AddCloudsRequestZoneConfig) GetGuidanceOk() (*string, bool)`

GetGuidanceOk returns a tuple with the Guidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidance

`func (o *AddCloudsRequestZoneConfig) SetGuidance(v string)`

SetGuidance sets Guidance field to given value.

### HasGuidance

`func (o *AddCloudsRequestZoneConfig) HasGuidance() bool`

HasGuidance returns a boolean if a field has been set.

### GetCosting

`func (o *AddCloudsRequestZoneConfig) GetCosting() string`

GetCosting returns the Costing field if non-nil, zero value otherwise.

### GetCostingOk

`func (o *AddCloudsRequestZoneConfig) GetCostingOk() (*string, bool)`

GetCostingOk returns a tuple with the Costing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosting

`func (o *AddCloudsRequestZoneConfig) SetCosting(v string)`

SetCosting sets Costing field to given value.

### HasCosting

`func (o *AddCloudsRequestZoneConfig) HasCosting() bool`

HasCosting returns a boolean if a field has been set.

### GetConfigCmdbDiscovery

`func (o *AddCloudsRequestZoneConfig) GetConfigCmdbDiscovery() string`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *AddCloudsRequestZoneConfig) GetConfigCmdbDiscoveryOk() (*string, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *AddCloudsRequestZoneConfig) SetConfigCmdbDiscovery(v string)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *AddCloudsRequestZoneConfig) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetLogo

`func (o *AddCloudsRequestZoneConfig) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *AddCloudsRequestZoneConfig) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *AddCloudsRequestZoneConfig) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *AddCloudsRequestZoneConfig) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetDarkModeLogo

`func (o *AddCloudsRequestZoneConfig) GetDarkModeLogo() string`

GetDarkModeLogo returns the DarkModeLogo field if non-nil, zero value otherwise.

### GetDarkModeLogoOk

`func (o *AddCloudsRequestZoneConfig) GetDarkModeLogoOk() (*string, bool)`

GetDarkModeLogoOk returns a tuple with the DarkModeLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkModeLogo

`func (o *AddCloudsRequestZoneConfig) SetDarkModeLogo(v string)`

SetDarkModeLogo sets DarkModeLogo field to given value.

### HasDarkModeLogo

`func (o *AddCloudsRequestZoneConfig) HasDarkModeLogo() bool`

HasDarkModeLogo returns a boolean if a field has been set.

### GetProxy

`func (o *AddCloudsRequestZoneConfig) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AddCloudsRequestZoneConfig) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AddCloudsRequestZoneConfig) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AddCloudsRequestZoneConfig) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetBypassProxyForCloud

`func (o *AddCloudsRequestZoneConfig) GetBypassProxyForCloud() string`

GetBypassProxyForCloud returns the BypassProxyForCloud field if non-nil, zero value otherwise.

### GetBypassProxyForCloudOk

`func (o *AddCloudsRequestZoneConfig) GetBypassProxyForCloudOk() (*string, bool)`

GetBypassProxyForCloudOk returns a tuple with the BypassProxyForCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassProxyForCloud

`func (o *AddCloudsRequestZoneConfig) SetBypassProxyForCloud(v string)`

SetBypassProxyForCloud sets BypassProxyForCloud field to given value.

### HasBypassProxyForCloud

`func (o *AddCloudsRequestZoneConfig) HasBypassProxyForCloud() bool`

HasBypassProxyForCloud returns a boolean if a field has been set.

### GetNoProxy

`func (o *AddCloudsRequestZoneConfig) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *AddCloudsRequestZoneConfig) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *AddCloudsRequestZoneConfig) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *AddCloudsRequestZoneConfig) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetUserDataLinux

`func (o *AddCloudsRequestZoneConfig) GetUserDataLinux() string`

GetUserDataLinux returns the UserDataLinux field if non-nil, zero value otherwise.

### GetUserDataLinuxOk

`func (o *AddCloudsRequestZoneConfig) GetUserDataLinuxOk() (*string, bool)`

GetUserDataLinuxOk returns a tuple with the UserDataLinux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataLinux

`func (o *AddCloudsRequestZoneConfig) SetUserDataLinux(v string)`

SetUserDataLinux sets UserDataLinux field to given value.

### HasUserDataLinux

`func (o *AddCloudsRequestZoneConfig) HasUserDataLinux() bool`

HasUserDataLinux returns a boolean if a field has been set.

### GetSubscriberId

`func (o *AddCloudsRequestZoneConfig) GetSubscriberId() string`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *AddCloudsRequestZoneConfig) GetSubscriberIdOk() (*string, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *AddCloudsRequestZoneConfig) SetSubscriberId(v string)`

SetSubscriberId sets SubscriberId field to given value.

### HasSubscriberId

`func (o *AddCloudsRequestZoneConfig) HasSubscriberId() bool`

HasSubscriberId returns a boolean if a field has been set.

### GetTenantId

`func (o *AddCloudsRequestZoneConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AddCloudsRequestZoneConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AddCloudsRequestZoneConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AddCloudsRequestZoneConfig) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetClientId

`func (o *AddCloudsRequestZoneConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AddCloudsRequestZoneConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AddCloudsRequestZoneConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AddCloudsRequestZoneConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AddCloudsRequestZoneConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AddCloudsRequestZoneConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AddCloudsRequestZoneConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AddCloudsRequestZoneConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetResourceGroup

`func (o *AddCloudsRequestZoneConfig) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AddCloudsRequestZoneConfig) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AddCloudsRequestZoneConfig) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AddCloudsRequestZoneConfig) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetRpcMode

`func (o *AddCloudsRequestZoneConfig) GetRpcMode() string`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *AddCloudsRequestZoneConfig) GetRpcModeOk() (*string, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *AddCloudsRequestZoneConfig) SetRpcMode(v string)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *AddCloudsRequestZoneConfig) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *AddCloudsRequestZoneConfig) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *AddCloudsRequestZoneConfig) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *AddCloudsRequestZoneConfig) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *AddCloudsRequestZoneConfig) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *AddCloudsRequestZoneConfig) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *AddCloudsRequestZoneConfig) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *AddCloudsRequestZoneConfig) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *AddCloudsRequestZoneConfig) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### GetApiUrl

`func (o *AddCloudsRequestZoneConfig) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *AddCloudsRequestZoneConfig) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *AddCloudsRequestZoneConfig) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.


### GetApiVersion

`func (o *AddCloudsRequestZoneConfig) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AddCloudsRequestZoneConfig) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AddCloudsRequestZoneConfig) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetDatacenter

`func (o *AddCloudsRequestZoneConfig) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *AddCloudsRequestZoneConfig) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *AddCloudsRequestZoneConfig) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.


### GetCluster

`func (o *AddCloudsRequestZoneConfig) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AddCloudsRequestZoneConfig) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AddCloudsRequestZoneConfig) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *AddCloudsRequestZoneConfig) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetResourcePool

`func (o *AddCloudsRequestZoneConfig) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *AddCloudsRequestZoneConfig) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *AddCloudsRequestZoneConfig) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *AddCloudsRequestZoneConfig) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetStorageType

`func (o *AddCloudsRequestZoneConfig) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *AddCloudsRequestZoneConfig) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *AddCloudsRequestZoneConfig) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *AddCloudsRequestZoneConfig) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetEnableVnc

`func (o *AddCloudsRequestZoneConfig) GetEnableVnc() string`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *AddCloudsRequestZoneConfig) GetEnableVncOk() (*string, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *AddCloudsRequestZoneConfig) SetEnableVnc(v string)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *AddCloudsRequestZoneConfig) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### GetHideHostSelection

`func (o *AddCloudsRequestZoneConfig) GetHideHostSelection() string`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *AddCloudsRequestZoneConfig) GetHideHostSelectionOk() (*string, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *AddCloudsRequestZoneConfig) SetHideHostSelection(v string)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *AddCloudsRequestZoneConfig) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### GetEnableDiskTypeSelection

`func (o *AddCloudsRequestZoneConfig) GetEnableDiskTypeSelection() string`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *AddCloudsRequestZoneConfig) GetEnableDiskTypeSelectionOk() (*string, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *AddCloudsRequestZoneConfig) SetEnableDiskTypeSelection(v string)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *AddCloudsRequestZoneConfig) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### GetEnableStorageTypeSelection

`func (o *AddCloudsRequestZoneConfig) GetEnableStorageTypeSelection() string`

GetEnableStorageTypeSelection returns the EnableStorageTypeSelection field if non-nil, zero value otherwise.

### GetEnableStorageTypeSelectionOk

`func (o *AddCloudsRequestZoneConfig) GetEnableStorageTypeSelectionOk() (*string, bool)`

GetEnableStorageTypeSelectionOk returns a tuple with the EnableStorageTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStorageTypeSelection

`func (o *AddCloudsRequestZoneConfig) SetEnableStorageTypeSelection(v string)`

SetEnableStorageTypeSelection sets EnableStorageTypeSelection field to given value.

### HasEnableStorageTypeSelection

`func (o *AddCloudsRequestZoneConfig) HasEnableStorageTypeSelection() bool`

HasEnableStorageTypeSelection returns a boolean if a field has been set.

### GetUsername

`func (o *AddCloudsRequestZoneConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddCloudsRequestZoneConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddCloudsRequestZoneConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddCloudsRequestZoneConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AddCloudsRequestZoneConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddCloudsRequestZoneConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddCloudsRequestZoneConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddCloudsRequestZoneConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ZoneCreateConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | Pointer to **string** | The URL used by workloads provisioned in the cloud for interacting with the Morpheus appliance. | [optional] 
**DatacenterName** | Pointer to **string** | A custom name used to reference the datacenter for the cloud. | [optional] 
**ExternalId** | Pointer to **NullableString** | The external id of the cloud | [optional] 
**InventoryLevel** | Pointer to **string** | Whether to import existing virtual machines. | [optional] 
**ConsoleKeymap** | Pointer to **string** | The keyboard layout to use for the console | [optional] 
**Endpoint** | **string** | AWS endpoint | 
**AccessKey** | Pointer to **string** | AWS access key | [optional] 
**SecretKey** | Pointer to **string** | AWS secret key | [optional] 
**UseHostCredentials** | Pointer to **string** | Whether to use the IAM profile associated with the Morpheus server or not | [optional] [default to "on"]
**EbsEncryption** | Pointer to **string** | Determines whether to configure default EBS volume encryption or not | [optional] [default to "on"]
**StsAssumeRole** | Pointer to **string** | The AWS IAM role ARN to assume for authentication | [optional] 
**ConfigManagementId** | Pointer to **string** | The id of the configuration management integration associated with the AWS cloud | [optional] 
**Vpc** | Pointer to **string** | The VPC ID for a specific VPC | [optional] 
**VdiGateway** | Pointer to **string** | The VDI gateway for this cloud to use for provisioning virtual desktops. | [optional] 
**CmdbConfig** | Pointer to **string** | The CMDB configuration for this cloud to use for syncing with the CMDB. | [optional] 
**ChangeManagementConfig** | Pointer to **string** | The change management configuration for this cloud to use for syncing with the change management system. | [optional] 
**NetworkMode** | Pointer to **NullableString** | Whether to use public or private IP addresses for provisioning and managing instances in this cloud. | [optional] 
**ApiProxy** | Pointer to **NullableString** | The API proxy to use for API calls to the cloud. | [optional] 
**Region** | Pointer to **string** | The AWS region to use for this cloud. | [optional] 
**Credentials** | Pointer to **string** | The credentials to use for this cloud. | [optional] 
**CostingBucket** | Pointer to **string** | The S3 bucket to use for storing costing reports. | [optional] 
**CostingFolder** | Pointer to **NullableString** | The folder within the S3 bucket to use for storing costing reports. | [optional] 
**CostingReportName** | Pointer to **NullableString** | The name of the costing report to generate. | [optional] 
**CostingKey** | Pointer to **NullableString** | The AWS access key to use for generating costing reports. | [optional] 
**CostingSecret** | Pointer to **NullableString** | The AWS secret key to use for generating costing reports. | [optional] 
**Domain** | Pointer to **string** | The domain to use for this cloud. | [optional] 
**Timezone** | Pointer to **string** | The timezone to use for this cloud. | [optional] 
**SecurityServer** | Pointer to **string** | The security server to use for this cloud, or &#39;off&#39; to not use a security server. | [optional] 
**Guidance** | Pointer to **string** | Optional guidance field if you want to put more info there | [optional] 
**Costing** | Pointer to **string** | Whether to enable costing for this cloud or not. | [optional] 
**ConfigCmdbDiscovery** | Pointer to **string** | The CMDB discovery configuration to use for this cloud. | [optional] 
**Logo** | Pointer to **NullableString** | The logo to use for this cloud. | [optional] 
**DarkModeLogo** | Pointer to **NullableString** | The logo to use for this cloud in dark mode. | [optional] 
**Proxy** | Pointer to **string** | The proxy to use for this cloud. | [optional] 
**BypassProxyForCloud** | Pointer to **string** | Whether to bypass the proxy for API calls to the cloud or not. | [optional] 
**NoProxy** | Pointer to **string** | A comma separated list of hosts to bypass the proxy for when making API calls to the cloud. | [optional] 
**UserDataLinux** | Pointer to **string** | The user data script to use for provisioning instances in this cloud. | [optional] 

## Methods

### NewZoneCreateConfigAnyOf

`func NewZoneCreateConfigAnyOf(endpoint string, ) *ZoneCreateConfigAnyOf`

NewZoneCreateConfigAnyOf instantiates a new ZoneCreateConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCreateConfigAnyOfWithDefaults

`func NewZoneCreateConfigAnyOfWithDefaults() *ZoneCreateConfigAnyOf`

NewZoneCreateConfigAnyOfWithDefaults instantiates a new ZoneCreateConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *ZoneCreateConfigAnyOf) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ZoneCreateConfigAnyOf) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ZoneCreateConfigAnyOf) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ZoneCreateConfigAnyOf) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *ZoneCreateConfigAnyOf) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *ZoneCreateConfigAnyOf) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *ZoneCreateConfigAnyOf) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *ZoneCreateConfigAnyOf) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *ZoneCreateConfigAnyOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ZoneCreateConfigAnyOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ZoneCreateConfigAnyOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ZoneCreateConfigAnyOf) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ZoneCreateConfigAnyOf) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ZoneCreateConfigAnyOf) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInventoryLevel

`func (o *ZoneCreateConfigAnyOf) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *ZoneCreateConfigAnyOf) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *ZoneCreateConfigAnyOf) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *ZoneCreateConfigAnyOf) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *ZoneCreateConfigAnyOf) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *ZoneCreateConfigAnyOf) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *ZoneCreateConfigAnyOf) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *ZoneCreateConfigAnyOf) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetEndpoint

`func (o *ZoneCreateConfigAnyOf) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ZoneCreateConfigAnyOf) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ZoneCreateConfigAnyOf) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetAccessKey

`func (o *ZoneCreateConfigAnyOf) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ZoneCreateConfigAnyOf) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ZoneCreateConfigAnyOf) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ZoneCreateConfigAnyOf) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *ZoneCreateConfigAnyOf) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ZoneCreateConfigAnyOf) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ZoneCreateConfigAnyOf) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ZoneCreateConfigAnyOf) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *ZoneCreateConfigAnyOf) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *ZoneCreateConfigAnyOf) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *ZoneCreateConfigAnyOf) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *ZoneCreateConfigAnyOf) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *ZoneCreateConfigAnyOf) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *ZoneCreateConfigAnyOf) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *ZoneCreateConfigAnyOf) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *ZoneCreateConfigAnyOf) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *ZoneCreateConfigAnyOf) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *ZoneCreateConfigAnyOf) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *ZoneCreateConfigAnyOf) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *ZoneCreateConfigAnyOf) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *ZoneCreateConfigAnyOf) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *ZoneCreateConfigAnyOf) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *ZoneCreateConfigAnyOf) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *ZoneCreateConfigAnyOf) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetVpc

`func (o *ZoneCreateConfigAnyOf) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *ZoneCreateConfigAnyOf) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *ZoneCreateConfigAnyOf) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *ZoneCreateConfigAnyOf) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetVdiGateway

`func (o *ZoneCreateConfigAnyOf) GetVdiGateway() string`

GetVdiGateway returns the VdiGateway field if non-nil, zero value otherwise.

### GetVdiGatewayOk

`func (o *ZoneCreateConfigAnyOf) GetVdiGatewayOk() (*string, bool)`

GetVdiGatewayOk returns a tuple with the VdiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiGateway

`func (o *ZoneCreateConfigAnyOf) SetVdiGateway(v string)`

SetVdiGateway sets VdiGateway field to given value.

### HasVdiGateway

`func (o *ZoneCreateConfigAnyOf) HasVdiGateway() bool`

HasVdiGateway returns a boolean if a field has been set.

### GetCmdbConfig

`func (o *ZoneCreateConfigAnyOf) GetCmdbConfig() string`

GetCmdbConfig returns the CmdbConfig field if non-nil, zero value otherwise.

### GetCmdbConfigOk

`func (o *ZoneCreateConfigAnyOf) GetCmdbConfigOk() (*string, bool)`

GetCmdbConfigOk returns a tuple with the CmdbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdbConfig

`func (o *ZoneCreateConfigAnyOf) SetCmdbConfig(v string)`

SetCmdbConfig sets CmdbConfig field to given value.

### HasCmdbConfig

`func (o *ZoneCreateConfigAnyOf) HasCmdbConfig() bool`

HasCmdbConfig returns a boolean if a field has been set.

### GetChangeManagementConfig

`func (o *ZoneCreateConfigAnyOf) GetChangeManagementConfig() string`

GetChangeManagementConfig returns the ChangeManagementConfig field if non-nil, zero value otherwise.

### GetChangeManagementConfigOk

`func (o *ZoneCreateConfigAnyOf) GetChangeManagementConfigOk() (*string, bool)`

GetChangeManagementConfigOk returns a tuple with the ChangeManagementConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeManagementConfig

`func (o *ZoneCreateConfigAnyOf) SetChangeManagementConfig(v string)`

SetChangeManagementConfig sets ChangeManagementConfig field to given value.

### HasChangeManagementConfig

`func (o *ZoneCreateConfigAnyOf) HasChangeManagementConfig() bool`

HasChangeManagementConfig returns a boolean if a field has been set.

### GetNetworkMode

`func (o *ZoneCreateConfigAnyOf) GetNetworkMode() string`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *ZoneCreateConfigAnyOf) GetNetworkModeOk() (*string, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *ZoneCreateConfigAnyOf) SetNetworkMode(v string)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *ZoneCreateConfigAnyOf) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### SetNetworkModeNil

`func (o *ZoneCreateConfigAnyOf) SetNetworkModeNil(b bool)`

 SetNetworkModeNil sets the value for NetworkMode to be an explicit nil

### UnsetNetworkMode
`func (o *ZoneCreateConfigAnyOf) UnsetNetworkMode()`

UnsetNetworkMode ensures that no value is present for NetworkMode, not even an explicit nil
### GetApiProxy

`func (o *ZoneCreateConfigAnyOf) GetApiProxy() string`

GetApiProxy returns the ApiProxy field if non-nil, zero value otherwise.

### GetApiProxyOk

`func (o *ZoneCreateConfigAnyOf) GetApiProxyOk() (*string, bool)`

GetApiProxyOk returns a tuple with the ApiProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProxy

`func (o *ZoneCreateConfigAnyOf) SetApiProxy(v string)`

SetApiProxy sets ApiProxy field to given value.

### HasApiProxy

`func (o *ZoneCreateConfigAnyOf) HasApiProxy() bool`

HasApiProxy returns a boolean if a field has been set.

### SetApiProxyNil

`func (o *ZoneCreateConfigAnyOf) SetApiProxyNil(b bool)`

 SetApiProxyNil sets the value for ApiProxy to be an explicit nil

### UnsetApiProxy
`func (o *ZoneCreateConfigAnyOf) UnsetApiProxy()`

UnsetApiProxy ensures that no value is present for ApiProxy, not even an explicit nil
### GetRegion

`func (o *ZoneCreateConfigAnyOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ZoneCreateConfigAnyOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ZoneCreateConfigAnyOf) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ZoneCreateConfigAnyOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCredentials

`func (o *ZoneCreateConfigAnyOf) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ZoneCreateConfigAnyOf) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ZoneCreateConfigAnyOf) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ZoneCreateConfigAnyOf) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetCostingBucket

`func (o *ZoneCreateConfigAnyOf) GetCostingBucket() string`

GetCostingBucket returns the CostingBucket field if non-nil, zero value otherwise.

### GetCostingBucketOk

`func (o *ZoneCreateConfigAnyOf) GetCostingBucketOk() (*string, bool)`

GetCostingBucketOk returns a tuple with the CostingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucket

`func (o *ZoneCreateConfigAnyOf) SetCostingBucket(v string)`

SetCostingBucket sets CostingBucket field to given value.

### HasCostingBucket

`func (o *ZoneCreateConfigAnyOf) HasCostingBucket() bool`

HasCostingBucket returns a boolean if a field has been set.

### GetCostingFolder

`func (o *ZoneCreateConfigAnyOf) GetCostingFolder() string`

GetCostingFolder returns the CostingFolder field if non-nil, zero value otherwise.

### GetCostingFolderOk

`func (o *ZoneCreateConfigAnyOf) GetCostingFolderOk() (*string, bool)`

GetCostingFolderOk returns a tuple with the CostingFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingFolder

`func (o *ZoneCreateConfigAnyOf) SetCostingFolder(v string)`

SetCostingFolder sets CostingFolder field to given value.

### HasCostingFolder

`func (o *ZoneCreateConfigAnyOf) HasCostingFolder() bool`

HasCostingFolder returns a boolean if a field has been set.

### SetCostingFolderNil

`func (o *ZoneCreateConfigAnyOf) SetCostingFolderNil(b bool)`

 SetCostingFolderNil sets the value for CostingFolder to be an explicit nil

### UnsetCostingFolder
`func (o *ZoneCreateConfigAnyOf) UnsetCostingFolder()`

UnsetCostingFolder ensures that no value is present for CostingFolder, not even an explicit nil
### GetCostingReportName

`func (o *ZoneCreateConfigAnyOf) GetCostingReportName() string`

GetCostingReportName returns the CostingReportName field if non-nil, zero value otherwise.

### GetCostingReportNameOk

`func (o *ZoneCreateConfigAnyOf) GetCostingReportNameOk() (*string, bool)`

GetCostingReportNameOk returns a tuple with the CostingReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReportName

`func (o *ZoneCreateConfigAnyOf) SetCostingReportName(v string)`

SetCostingReportName sets CostingReportName field to given value.

### HasCostingReportName

`func (o *ZoneCreateConfigAnyOf) HasCostingReportName() bool`

HasCostingReportName returns a boolean if a field has been set.

### SetCostingReportNameNil

`func (o *ZoneCreateConfigAnyOf) SetCostingReportNameNil(b bool)`

 SetCostingReportNameNil sets the value for CostingReportName to be an explicit nil

### UnsetCostingReportName
`func (o *ZoneCreateConfigAnyOf) UnsetCostingReportName()`

UnsetCostingReportName ensures that no value is present for CostingReportName, not even an explicit nil
### GetCostingKey

`func (o *ZoneCreateConfigAnyOf) GetCostingKey() string`

GetCostingKey returns the CostingKey field if non-nil, zero value otherwise.

### GetCostingKeyOk

`func (o *ZoneCreateConfigAnyOf) GetCostingKeyOk() (*string, bool)`

GetCostingKeyOk returns a tuple with the CostingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingKey

`func (o *ZoneCreateConfigAnyOf) SetCostingKey(v string)`

SetCostingKey sets CostingKey field to given value.

### HasCostingKey

`func (o *ZoneCreateConfigAnyOf) HasCostingKey() bool`

HasCostingKey returns a boolean if a field has been set.

### SetCostingKeyNil

`func (o *ZoneCreateConfigAnyOf) SetCostingKeyNil(b bool)`

 SetCostingKeyNil sets the value for CostingKey to be an explicit nil

### UnsetCostingKey
`func (o *ZoneCreateConfigAnyOf) UnsetCostingKey()`

UnsetCostingKey ensures that no value is present for CostingKey, not even an explicit nil
### GetCostingSecret

`func (o *ZoneCreateConfigAnyOf) GetCostingSecret() string`

GetCostingSecret returns the CostingSecret field if non-nil, zero value otherwise.

### GetCostingSecretOk

`func (o *ZoneCreateConfigAnyOf) GetCostingSecretOk() (*string, bool)`

GetCostingSecretOk returns a tuple with the CostingSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecret

`func (o *ZoneCreateConfigAnyOf) SetCostingSecret(v string)`

SetCostingSecret sets CostingSecret field to given value.

### HasCostingSecret

`func (o *ZoneCreateConfigAnyOf) HasCostingSecret() bool`

HasCostingSecret returns a boolean if a field has been set.

### SetCostingSecretNil

`func (o *ZoneCreateConfigAnyOf) SetCostingSecretNil(b bool)`

 SetCostingSecretNil sets the value for CostingSecret to be an explicit nil

### UnsetCostingSecret
`func (o *ZoneCreateConfigAnyOf) UnsetCostingSecret()`

UnsetCostingSecret ensures that no value is present for CostingSecret, not even an explicit nil
### GetDomain

`func (o *ZoneCreateConfigAnyOf) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ZoneCreateConfigAnyOf) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ZoneCreateConfigAnyOf) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ZoneCreateConfigAnyOf) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTimezone

`func (o *ZoneCreateConfigAnyOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ZoneCreateConfigAnyOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ZoneCreateConfigAnyOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ZoneCreateConfigAnyOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSecurityServer

`func (o *ZoneCreateConfigAnyOf) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *ZoneCreateConfigAnyOf) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *ZoneCreateConfigAnyOf) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *ZoneCreateConfigAnyOf) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetGuidance

`func (o *ZoneCreateConfigAnyOf) GetGuidance() string`

GetGuidance returns the Guidance field if non-nil, zero value otherwise.

### GetGuidanceOk

`func (o *ZoneCreateConfigAnyOf) GetGuidanceOk() (*string, bool)`

GetGuidanceOk returns a tuple with the Guidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidance

`func (o *ZoneCreateConfigAnyOf) SetGuidance(v string)`

SetGuidance sets Guidance field to given value.

### HasGuidance

`func (o *ZoneCreateConfigAnyOf) HasGuidance() bool`

HasGuidance returns a boolean if a field has been set.

### GetCosting

`func (o *ZoneCreateConfigAnyOf) GetCosting() string`

GetCosting returns the Costing field if non-nil, zero value otherwise.

### GetCostingOk

`func (o *ZoneCreateConfigAnyOf) GetCostingOk() (*string, bool)`

GetCostingOk returns a tuple with the Costing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosting

`func (o *ZoneCreateConfigAnyOf) SetCosting(v string)`

SetCosting sets Costing field to given value.

### HasCosting

`func (o *ZoneCreateConfigAnyOf) HasCosting() bool`

HasCosting returns a boolean if a field has been set.

### GetConfigCmdbDiscovery

`func (o *ZoneCreateConfigAnyOf) GetConfigCmdbDiscovery() string`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *ZoneCreateConfigAnyOf) GetConfigCmdbDiscoveryOk() (*string, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *ZoneCreateConfigAnyOf) SetConfigCmdbDiscovery(v string)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *ZoneCreateConfigAnyOf) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetLogo

`func (o *ZoneCreateConfigAnyOf) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ZoneCreateConfigAnyOf) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ZoneCreateConfigAnyOf) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ZoneCreateConfigAnyOf) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *ZoneCreateConfigAnyOf) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *ZoneCreateConfigAnyOf) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetDarkModeLogo

`func (o *ZoneCreateConfigAnyOf) GetDarkModeLogo() string`

GetDarkModeLogo returns the DarkModeLogo field if non-nil, zero value otherwise.

### GetDarkModeLogoOk

`func (o *ZoneCreateConfigAnyOf) GetDarkModeLogoOk() (*string, bool)`

GetDarkModeLogoOk returns a tuple with the DarkModeLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkModeLogo

`func (o *ZoneCreateConfigAnyOf) SetDarkModeLogo(v string)`

SetDarkModeLogo sets DarkModeLogo field to given value.

### HasDarkModeLogo

`func (o *ZoneCreateConfigAnyOf) HasDarkModeLogo() bool`

HasDarkModeLogo returns a boolean if a field has been set.

### SetDarkModeLogoNil

`func (o *ZoneCreateConfigAnyOf) SetDarkModeLogoNil(b bool)`

 SetDarkModeLogoNil sets the value for DarkModeLogo to be an explicit nil

### UnsetDarkModeLogo
`func (o *ZoneCreateConfigAnyOf) UnsetDarkModeLogo()`

UnsetDarkModeLogo ensures that no value is present for DarkModeLogo, not even an explicit nil
### GetProxy

`func (o *ZoneCreateConfigAnyOf) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ZoneCreateConfigAnyOf) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ZoneCreateConfigAnyOf) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ZoneCreateConfigAnyOf) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetBypassProxyForCloud

`func (o *ZoneCreateConfigAnyOf) GetBypassProxyForCloud() string`

GetBypassProxyForCloud returns the BypassProxyForCloud field if non-nil, zero value otherwise.

### GetBypassProxyForCloudOk

`func (o *ZoneCreateConfigAnyOf) GetBypassProxyForCloudOk() (*string, bool)`

GetBypassProxyForCloudOk returns a tuple with the BypassProxyForCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassProxyForCloud

`func (o *ZoneCreateConfigAnyOf) SetBypassProxyForCloud(v string)`

SetBypassProxyForCloud sets BypassProxyForCloud field to given value.

### HasBypassProxyForCloud

`func (o *ZoneCreateConfigAnyOf) HasBypassProxyForCloud() bool`

HasBypassProxyForCloud returns a boolean if a field has been set.

### GetNoProxy

`func (o *ZoneCreateConfigAnyOf) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *ZoneCreateConfigAnyOf) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *ZoneCreateConfigAnyOf) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *ZoneCreateConfigAnyOf) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetUserDataLinux

`func (o *ZoneCreateConfigAnyOf) GetUserDataLinux() string`

GetUserDataLinux returns the UserDataLinux field if non-nil, zero value otherwise.

### GetUserDataLinuxOk

`func (o *ZoneCreateConfigAnyOf) GetUserDataLinuxOk() (*string, bool)`

GetUserDataLinuxOk returns a tuple with the UserDataLinux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataLinux

`func (o *ZoneCreateConfigAnyOf) SetUserDataLinux(v string)`

SetUserDataLinux sets UserDataLinux field to given value.

### HasUserDataLinux

`func (o *ZoneCreateConfigAnyOf) HasUserDataLinux() bool`

HasUserDataLinux returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



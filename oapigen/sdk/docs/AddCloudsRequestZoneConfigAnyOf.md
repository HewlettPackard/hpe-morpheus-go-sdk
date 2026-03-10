# AddCloudsRequestZoneConfigAnyOf

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

### NewAddCloudsRequestZoneConfigAnyOf

`func NewAddCloudsRequestZoneConfigAnyOf(endpoint string, ) *AddCloudsRequestZoneConfigAnyOf`

NewAddCloudsRequestZoneConfigAnyOf instantiates a new AddCloudsRequestZoneConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCloudsRequestZoneConfigAnyOfWithDefaults

`func NewAddCloudsRequestZoneConfigAnyOfWithDefaults() *AddCloudsRequestZoneConfigAnyOf`

NewAddCloudsRequestZoneConfigAnyOfWithDefaults instantiates a new AddCloudsRequestZoneConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *AddCloudsRequestZoneConfigAnyOf) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *AddCloudsRequestZoneConfigAnyOf) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *AddCloudsRequestZoneConfigAnyOf) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *AddCloudsRequestZoneConfigAnyOf) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *AddCloudsRequestZoneConfigAnyOf) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *AddCloudsRequestZoneConfigAnyOf) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *AddCloudsRequestZoneConfigAnyOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddCloudsRequestZoneConfigAnyOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddCloudsRequestZoneConfigAnyOf) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AddCloudsRequestZoneConfigAnyOf) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AddCloudsRequestZoneConfigAnyOf) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInventoryLevel

`func (o *AddCloudsRequestZoneConfigAnyOf) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *AddCloudsRequestZoneConfigAnyOf) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *AddCloudsRequestZoneConfigAnyOf) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *AddCloudsRequestZoneConfigAnyOf) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *AddCloudsRequestZoneConfigAnyOf) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *AddCloudsRequestZoneConfigAnyOf) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetEndpoint

`func (o *AddCloudsRequestZoneConfigAnyOf) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AddCloudsRequestZoneConfigAnyOf) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetAccessKey

`func (o *AddCloudsRequestZoneConfigAnyOf) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AddCloudsRequestZoneConfigAnyOf) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *AddCloudsRequestZoneConfigAnyOf) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *AddCloudsRequestZoneConfigAnyOf) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AddCloudsRequestZoneConfigAnyOf) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AddCloudsRequestZoneConfigAnyOf) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *AddCloudsRequestZoneConfigAnyOf) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *AddCloudsRequestZoneConfigAnyOf) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *AddCloudsRequestZoneConfigAnyOf) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *AddCloudsRequestZoneConfigAnyOf) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *AddCloudsRequestZoneConfigAnyOf) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *AddCloudsRequestZoneConfigAnyOf) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *AddCloudsRequestZoneConfigAnyOf) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *AddCloudsRequestZoneConfigAnyOf) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *AddCloudsRequestZoneConfigAnyOf) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *AddCloudsRequestZoneConfigAnyOf) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *AddCloudsRequestZoneConfigAnyOf) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *AddCloudsRequestZoneConfigAnyOf) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetVpc

`func (o *AddCloudsRequestZoneConfigAnyOf) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *AddCloudsRequestZoneConfigAnyOf) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *AddCloudsRequestZoneConfigAnyOf) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetVdiGateway

`func (o *AddCloudsRequestZoneConfigAnyOf) GetVdiGateway() string`

GetVdiGateway returns the VdiGateway field if non-nil, zero value otherwise.

### GetVdiGatewayOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetVdiGatewayOk() (*string, bool)`

GetVdiGatewayOk returns a tuple with the VdiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiGateway

`func (o *AddCloudsRequestZoneConfigAnyOf) SetVdiGateway(v string)`

SetVdiGateway sets VdiGateway field to given value.

### HasVdiGateway

`func (o *AddCloudsRequestZoneConfigAnyOf) HasVdiGateway() bool`

HasVdiGateway returns a boolean if a field has been set.

### GetCmdbConfig

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCmdbConfig() string`

GetCmdbConfig returns the CmdbConfig field if non-nil, zero value otherwise.

### GetCmdbConfigOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCmdbConfigOk() (*string, bool)`

GetCmdbConfigOk returns a tuple with the CmdbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdbConfig

`func (o *AddCloudsRequestZoneConfigAnyOf) SetCmdbConfig(v string)`

SetCmdbConfig sets CmdbConfig field to given value.

### HasCmdbConfig

`func (o *AddCloudsRequestZoneConfigAnyOf) HasCmdbConfig() bool`

HasCmdbConfig returns a boolean if a field has been set.

### GetChangeManagementConfig

`func (o *AddCloudsRequestZoneConfigAnyOf) GetChangeManagementConfig() string`

GetChangeManagementConfig returns the ChangeManagementConfig field if non-nil, zero value otherwise.

### GetChangeManagementConfigOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetChangeManagementConfigOk() (*string, bool)`

GetChangeManagementConfigOk returns a tuple with the ChangeManagementConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeManagementConfig

`func (o *AddCloudsRequestZoneConfigAnyOf) SetChangeManagementConfig(v string)`

SetChangeManagementConfig sets ChangeManagementConfig field to given value.

### HasChangeManagementConfig

`func (o *AddCloudsRequestZoneConfigAnyOf) HasChangeManagementConfig() bool`

HasChangeManagementConfig returns a boolean if a field has been set.

### GetNetworkMode

`func (o *AddCloudsRequestZoneConfigAnyOf) GetNetworkMode() string`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetNetworkModeOk() (*string, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *AddCloudsRequestZoneConfigAnyOf) SetNetworkMode(v string)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *AddCloudsRequestZoneConfigAnyOf) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### SetNetworkModeNil

`func (o *AddCloudsRequestZoneConfigAnyOf) SetNetworkModeNil(b bool)`

 SetNetworkModeNil sets the value for NetworkMode to be an explicit nil

### UnsetNetworkMode
`func (o *AddCloudsRequestZoneConfigAnyOf) UnsetNetworkMode()`

UnsetNetworkMode ensures that no value is present for NetworkMode, not even an explicit nil
### GetApiProxy

`func (o *AddCloudsRequestZoneConfigAnyOf) GetApiProxy() string`

GetApiProxy returns the ApiProxy field if non-nil, zero value otherwise.

### GetApiProxyOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetApiProxyOk() (*string, bool)`

GetApiProxyOk returns a tuple with the ApiProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProxy

`func (o *AddCloudsRequestZoneConfigAnyOf) SetApiProxy(v string)`

SetApiProxy sets ApiProxy field to given value.

### HasApiProxy

`func (o *AddCloudsRequestZoneConfigAnyOf) HasApiProxy() bool`

HasApiProxy returns a boolean if a field has been set.

### SetApiProxyNil

`func (o *AddCloudsRequestZoneConfigAnyOf) SetApiProxyNil(b bool)`

 SetApiProxyNil sets the value for ApiProxy to be an explicit nil

### UnsetApiProxy
`func (o *AddCloudsRequestZoneConfigAnyOf) UnsetApiProxy()`

UnsetApiProxy ensures that no value is present for ApiProxy, not even an explicit nil
### GetRegion

`func (o *AddCloudsRequestZoneConfigAnyOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddCloudsRequestZoneConfigAnyOf) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddCloudsRequestZoneConfigAnyOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCredentials

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AddCloudsRequestZoneConfigAnyOf) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AddCloudsRequestZoneConfigAnyOf) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetCostingBucket

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCostingBucket() string`

GetCostingBucket returns the CostingBucket field if non-nil, zero value otherwise.

### GetCostingBucketOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCostingBucketOk() (*string, bool)`

GetCostingBucketOk returns a tuple with the CostingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingBucket

`func (o *AddCloudsRequestZoneConfigAnyOf) SetCostingBucket(v string)`

SetCostingBucket sets CostingBucket field to given value.

### HasCostingBucket

`func (o *AddCloudsRequestZoneConfigAnyOf) HasCostingBucket() bool`

HasCostingBucket returns a boolean if a field has been set.

### GetCostingFolder

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCostingFolder() string`

GetCostingFolder returns the CostingFolder field if non-nil, zero value otherwise.

### GetCostingFolderOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCostingFolderOk() (*string, bool)`

GetCostingFolderOk returns a tuple with the CostingFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingFolder

`func (o *AddCloudsRequestZoneConfigAnyOf) SetCostingFolder(v string)`

SetCostingFolder sets CostingFolder field to given value.

### HasCostingFolder

`func (o *AddCloudsRequestZoneConfigAnyOf) HasCostingFolder() bool`

HasCostingFolder returns a boolean if a field has been set.

### SetCostingFolderNil

`func (o *AddCloudsRequestZoneConfigAnyOf) SetCostingFolderNil(b bool)`

 SetCostingFolderNil sets the value for CostingFolder to be an explicit nil

### UnsetCostingFolder
`func (o *AddCloudsRequestZoneConfigAnyOf) UnsetCostingFolder()`

UnsetCostingFolder ensures that no value is present for CostingFolder, not even an explicit nil
### GetCostingReportName

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCostingReportName() string`

GetCostingReportName returns the CostingReportName field if non-nil, zero value otherwise.

### GetCostingReportNameOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCostingReportNameOk() (*string, bool)`

GetCostingReportNameOk returns a tuple with the CostingReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingReportName

`func (o *AddCloudsRequestZoneConfigAnyOf) SetCostingReportName(v string)`

SetCostingReportName sets CostingReportName field to given value.

### HasCostingReportName

`func (o *AddCloudsRequestZoneConfigAnyOf) HasCostingReportName() bool`

HasCostingReportName returns a boolean if a field has been set.

### SetCostingReportNameNil

`func (o *AddCloudsRequestZoneConfigAnyOf) SetCostingReportNameNil(b bool)`

 SetCostingReportNameNil sets the value for CostingReportName to be an explicit nil

### UnsetCostingReportName
`func (o *AddCloudsRequestZoneConfigAnyOf) UnsetCostingReportName()`

UnsetCostingReportName ensures that no value is present for CostingReportName, not even an explicit nil
### GetCostingKey

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCostingKey() string`

GetCostingKey returns the CostingKey field if non-nil, zero value otherwise.

### GetCostingKeyOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCostingKeyOk() (*string, bool)`

GetCostingKeyOk returns a tuple with the CostingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingKey

`func (o *AddCloudsRequestZoneConfigAnyOf) SetCostingKey(v string)`

SetCostingKey sets CostingKey field to given value.

### HasCostingKey

`func (o *AddCloudsRequestZoneConfigAnyOf) HasCostingKey() bool`

HasCostingKey returns a boolean if a field has been set.

### SetCostingKeyNil

`func (o *AddCloudsRequestZoneConfigAnyOf) SetCostingKeyNil(b bool)`

 SetCostingKeyNil sets the value for CostingKey to be an explicit nil

### UnsetCostingKey
`func (o *AddCloudsRequestZoneConfigAnyOf) UnsetCostingKey()`

UnsetCostingKey ensures that no value is present for CostingKey, not even an explicit nil
### GetCostingSecret

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCostingSecret() string`

GetCostingSecret returns the CostingSecret field if non-nil, zero value otherwise.

### GetCostingSecretOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCostingSecretOk() (*string, bool)`

GetCostingSecretOk returns a tuple with the CostingSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingSecret

`func (o *AddCloudsRequestZoneConfigAnyOf) SetCostingSecret(v string)`

SetCostingSecret sets CostingSecret field to given value.

### HasCostingSecret

`func (o *AddCloudsRequestZoneConfigAnyOf) HasCostingSecret() bool`

HasCostingSecret returns a boolean if a field has been set.

### SetCostingSecretNil

`func (o *AddCloudsRequestZoneConfigAnyOf) SetCostingSecretNil(b bool)`

 SetCostingSecretNil sets the value for CostingSecret to be an explicit nil

### UnsetCostingSecret
`func (o *AddCloudsRequestZoneConfigAnyOf) UnsetCostingSecret()`

UnsetCostingSecret ensures that no value is present for CostingSecret, not even an explicit nil
### GetDomain

`func (o *AddCloudsRequestZoneConfigAnyOf) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AddCloudsRequestZoneConfigAnyOf) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AddCloudsRequestZoneConfigAnyOf) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTimezone

`func (o *AddCloudsRequestZoneConfigAnyOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AddCloudsRequestZoneConfigAnyOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AddCloudsRequestZoneConfigAnyOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSecurityServer

`func (o *AddCloudsRequestZoneConfigAnyOf) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *AddCloudsRequestZoneConfigAnyOf) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *AddCloudsRequestZoneConfigAnyOf) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetGuidance

`func (o *AddCloudsRequestZoneConfigAnyOf) GetGuidance() string`

GetGuidance returns the Guidance field if non-nil, zero value otherwise.

### GetGuidanceOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetGuidanceOk() (*string, bool)`

GetGuidanceOk returns a tuple with the Guidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidance

`func (o *AddCloudsRequestZoneConfigAnyOf) SetGuidance(v string)`

SetGuidance sets Guidance field to given value.

### HasGuidance

`func (o *AddCloudsRequestZoneConfigAnyOf) HasGuidance() bool`

HasGuidance returns a boolean if a field has been set.

### GetCosting

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCosting() string`

GetCosting returns the Costing field if non-nil, zero value otherwise.

### GetCostingOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetCostingOk() (*string, bool)`

GetCostingOk returns a tuple with the Costing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosting

`func (o *AddCloudsRequestZoneConfigAnyOf) SetCosting(v string)`

SetCosting sets Costing field to given value.

### HasCosting

`func (o *AddCloudsRequestZoneConfigAnyOf) HasCosting() bool`

HasCosting returns a boolean if a field has been set.

### GetConfigCmdbDiscovery

`func (o *AddCloudsRequestZoneConfigAnyOf) GetConfigCmdbDiscovery() string`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetConfigCmdbDiscoveryOk() (*string, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *AddCloudsRequestZoneConfigAnyOf) SetConfigCmdbDiscovery(v string)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *AddCloudsRequestZoneConfigAnyOf) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetLogo

`func (o *AddCloudsRequestZoneConfigAnyOf) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *AddCloudsRequestZoneConfigAnyOf) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *AddCloudsRequestZoneConfigAnyOf) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *AddCloudsRequestZoneConfigAnyOf) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *AddCloudsRequestZoneConfigAnyOf) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetDarkModeLogo

`func (o *AddCloudsRequestZoneConfigAnyOf) GetDarkModeLogo() string`

GetDarkModeLogo returns the DarkModeLogo field if non-nil, zero value otherwise.

### GetDarkModeLogoOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetDarkModeLogoOk() (*string, bool)`

GetDarkModeLogoOk returns a tuple with the DarkModeLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkModeLogo

`func (o *AddCloudsRequestZoneConfigAnyOf) SetDarkModeLogo(v string)`

SetDarkModeLogo sets DarkModeLogo field to given value.

### HasDarkModeLogo

`func (o *AddCloudsRequestZoneConfigAnyOf) HasDarkModeLogo() bool`

HasDarkModeLogo returns a boolean if a field has been set.

### SetDarkModeLogoNil

`func (o *AddCloudsRequestZoneConfigAnyOf) SetDarkModeLogoNil(b bool)`

 SetDarkModeLogoNil sets the value for DarkModeLogo to be an explicit nil

### UnsetDarkModeLogo
`func (o *AddCloudsRequestZoneConfigAnyOf) UnsetDarkModeLogo()`

UnsetDarkModeLogo ensures that no value is present for DarkModeLogo, not even an explicit nil
### GetProxy

`func (o *AddCloudsRequestZoneConfigAnyOf) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AddCloudsRequestZoneConfigAnyOf) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AddCloudsRequestZoneConfigAnyOf) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetBypassProxyForCloud

`func (o *AddCloudsRequestZoneConfigAnyOf) GetBypassProxyForCloud() string`

GetBypassProxyForCloud returns the BypassProxyForCloud field if non-nil, zero value otherwise.

### GetBypassProxyForCloudOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetBypassProxyForCloudOk() (*string, bool)`

GetBypassProxyForCloudOk returns a tuple with the BypassProxyForCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassProxyForCloud

`func (o *AddCloudsRequestZoneConfigAnyOf) SetBypassProxyForCloud(v string)`

SetBypassProxyForCloud sets BypassProxyForCloud field to given value.

### HasBypassProxyForCloud

`func (o *AddCloudsRequestZoneConfigAnyOf) HasBypassProxyForCloud() bool`

HasBypassProxyForCloud returns a boolean if a field has been set.

### GetNoProxy

`func (o *AddCloudsRequestZoneConfigAnyOf) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *AddCloudsRequestZoneConfigAnyOf) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *AddCloudsRequestZoneConfigAnyOf) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetUserDataLinux

`func (o *AddCloudsRequestZoneConfigAnyOf) GetUserDataLinux() string`

GetUserDataLinux returns the UserDataLinux field if non-nil, zero value otherwise.

### GetUserDataLinuxOk

`func (o *AddCloudsRequestZoneConfigAnyOf) GetUserDataLinuxOk() (*string, bool)`

GetUserDataLinuxOk returns a tuple with the UserDataLinux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataLinux

`func (o *AddCloudsRequestZoneConfigAnyOf) SetUserDataLinux(v string)`

SetUserDataLinux sets UserDataLinux field to given value.

### HasUserDataLinux

`func (o *AddCloudsRequestZoneConfigAnyOf) HasUserDataLinux() bool`

HasUserDataLinux returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



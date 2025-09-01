# AddCloudsRequestZoneConfig

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
**ConfigManagementId** | Pointer to **string** | The id of the configuration management integration associated with the vSphere cloud. | [optional] 
**Vpc** | Pointer to **string** | The VPC ID for a specific VPC | [optional] 
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

### SetExternalIdNil

`func (o *AddCloudsRequestZoneConfig) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AddCloudsRequestZoneConfig) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
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



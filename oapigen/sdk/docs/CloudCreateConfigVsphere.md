# CloudCreateConfigVsphere

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | Pointer to **string** | The URL used by workloads provisioned in the cloud for interacting with the Morpheus appliance. | [optional] 
**DatacenterName** | Pointer to **string** | A custom name used to reference the datacenter for the cloud. | [optional] 
**ExternalId** | Pointer to **NullableString** | The external id of the cloud | [optional] 
**InventoryLevel** | Pointer to **string** | Whether to import existing virtual machines. | [optional] 
**ConsoleKeymap** | Pointer to **string** | The keyboard layout to use for the console | [optional] 
**ApiUrl** | **string** | The SDK URL of the vCenter server. | 
**ApiVersion** | **string** | The SDK version of the vCenter server. | 
**Datacenter** | **string** | The vSphere datacenter to add. | 
**Cluster** | Pointer to **string** | The name of the vSphere cluster | [optional] [default to "all"]
**ConfigManagementId** | Pointer to **string** | The id of the configuration management integration associated with the vSphere cloud. | [optional] 
**ResourcePool** | Pointer to **string** | The name of the vSphere resource pool | [optional] 
**RpcMode** | Pointer to [**NullableCloudCreateConfigVsphereRpcMode**](CloudCreateConfigVsphereRpcMode.md) |  | [optional] 
**StorageType** | Pointer to **string** | The default vSphere VMDK type for virtual machines | [optional] [default to "thin"]
**CertificateProvider** | Pointer to **string** | Certificate provider | [optional] [default to "internal"]
**EnableVnc** | Pointer to **NullableString** |  | [optional] 
**HideHostSelection** | Pointer to **NullableString** |  | [optional] 
**EnableDiskTypeSelection** | Pointer to **NullableString** |  | [optional] 
**EnableStorageTypeSelection** | Pointer to **NullableString** |  | [optional] 
**EnableNetworkTypeSelection** | Pointer to [**NullableCloudCreateConfigVsphereEnableNetworkTypeSelection**](CloudCreateConfigVsphereEnableNetworkTypeSelection.md) |  | [optional] 
**Username** | Pointer to **string** | Username. | [optional] 
**Password** | Pointer to **string** | Password to apply to the user | [optional] 

## Methods

### NewCloudCreateConfigVsphere

`func NewCloudCreateConfigVsphere(apiUrl string, apiVersion string, datacenter string, ) *CloudCreateConfigVsphere`

NewCloudCreateConfigVsphere instantiates a new CloudCreateConfigVsphere object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCreateConfigVsphereWithDefaults

`func NewCloudCreateConfigVsphereWithDefaults() *CloudCreateConfigVsphere`

NewCloudCreateConfigVsphereWithDefaults instantiates a new CloudCreateConfigVsphere object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *CloudCreateConfigVsphere) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *CloudCreateConfigVsphere) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *CloudCreateConfigVsphere) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *CloudCreateConfigVsphere) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *CloudCreateConfigVsphere) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *CloudCreateConfigVsphere) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *CloudCreateConfigVsphere) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *CloudCreateConfigVsphere) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudCreateConfigVsphere) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudCreateConfigVsphere) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudCreateConfigVsphere) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudCreateConfigVsphere) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CloudCreateConfigVsphere) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CloudCreateConfigVsphere) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInventoryLevel

`func (o *CloudCreateConfigVsphere) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *CloudCreateConfigVsphere) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *CloudCreateConfigVsphere) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *CloudCreateConfigVsphere) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *CloudCreateConfigVsphere) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *CloudCreateConfigVsphere) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *CloudCreateConfigVsphere) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *CloudCreateConfigVsphere) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetApiUrl

`func (o *CloudCreateConfigVsphere) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *CloudCreateConfigVsphere) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *CloudCreateConfigVsphere) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.


### GetApiVersion

`func (o *CloudCreateConfigVsphere) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudCreateConfigVsphere) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudCreateConfigVsphere) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetDatacenter

`func (o *CloudCreateConfigVsphere) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *CloudCreateConfigVsphere) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *CloudCreateConfigVsphere) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.


### GetCluster

`func (o *CloudCreateConfigVsphere) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CloudCreateConfigVsphere) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CloudCreateConfigVsphere) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *CloudCreateConfigVsphere) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *CloudCreateConfigVsphere) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *CloudCreateConfigVsphere) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *CloudCreateConfigVsphere) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *CloudCreateConfigVsphere) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetResourcePool

`func (o *CloudCreateConfigVsphere) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *CloudCreateConfigVsphere) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *CloudCreateConfigVsphere) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *CloudCreateConfigVsphere) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetRpcMode

`func (o *CloudCreateConfigVsphere) GetRpcMode() CloudCreateConfigVsphereRpcMode`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *CloudCreateConfigVsphere) GetRpcModeOk() (*CloudCreateConfigVsphereRpcMode, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *CloudCreateConfigVsphere) SetRpcMode(v CloudCreateConfigVsphereRpcMode)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *CloudCreateConfigVsphere) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### SetRpcModeNil

`func (o *CloudCreateConfigVsphere) SetRpcModeNil(b bool)`

 SetRpcModeNil sets the value for RpcMode to be an explicit nil

### UnsetRpcMode
`func (o *CloudCreateConfigVsphere) UnsetRpcMode()`

UnsetRpcMode ensures that no value is present for RpcMode, not even an explicit nil
### GetStorageType

`func (o *CloudCreateConfigVsphere) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *CloudCreateConfigVsphere) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *CloudCreateConfigVsphere) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *CloudCreateConfigVsphere) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *CloudCreateConfigVsphere) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *CloudCreateConfigVsphere) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *CloudCreateConfigVsphere) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *CloudCreateConfigVsphere) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetEnableVnc

`func (o *CloudCreateConfigVsphere) GetEnableVnc() string`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *CloudCreateConfigVsphere) GetEnableVncOk() (*string, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *CloudCreateConfigVsphere) SetEnableVnc(v string)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *CloudCreateConfigVsphere) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### SetEnableVncNil

`func (o *CloudCreateConfigVsphere) SetEnableVncNil(b bool)`

 SetEnableVncNil sets the value for EnableVnc to be an explicit nil

### UnsetEnableVnc
`func (o *CloudCreateConfigVsphere) UnsetEnableVnc()`

UnsetEnableVnc ensures that no value is present for EnableVnc, not even an explicit nil
### GetHideHostSelection

`func (o *CloudCreateConfigVsphere) GetHideHostSelection() string`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *CloudCreateConfigVsphere) GetHideHostSelectionOk() (*string, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *CloudCreateConfigVsphere) SetHideHostSelection(v string)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *CloudCreateConfigVsphere) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### SetHideHostSelectionNil

`func (o *CloudCreateConfigVsphere) SetHideHostSelectionNil(b bool)`

 SetHideHostSelectionNil sets the value for HideHostSelection to be an explicit nil

### UnsetHideHostSelection
`func (o *CloudCreateConfigVsphere) UnsetHideHostSelection()`

UnsetHideHostSelection ensures that no value is present for HideHostSelection, not even an explicit nil
### GetEnableDiskTypeSelection

`func (o *CloudCreateConfigVsphere) GetEnableDiskTypeSelection() string`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *CloudCreateConfigVsphere) GetEnableDiskTypeSelectionOk() (*string, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *CloudCreateConfigVsphere) SetEnableDiskTypeSelection(v string)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *CloudCreateConfigVsphere) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### SetEnableDiskTypeSelectionNil

`func (o *CloudCreateConfigVsphere) SetEnableDiskTypeSelectionNil(b bool)`

 SetEnableDiskTypeSelectionNil sets the value for EnableDiskTypeSelection to be an explicit nil

### UnsetEnableDiskTypeSelection
`func (o *CloudCreateConfigVsphere) UnsetEnableDiskTypeSelection()`

UnsetEnableDiskTypeSelection ensures that no value is present for EnableDiskTypeSelection, not even an explicit nil
### GetEnableStorageTypeSelection

`func (o *CloudCreateConfigVsphere) GetEnableStorageTypeSelection() string`

GetEnableStorageTypeSelection returns the EnableStorageTypeSelection field if non-nil, zero value otherwise.

### GetEnableStorageTypeSelectionOk

`func (o *CloudCreateConfigVsphere) GetEnableStorageTypeSelectionOk() (*string, bool)`

GetEnableStorageTypeSelectionOk returns a tuple with the EnableStorageTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStorageTypeSelection

`func (o *CloudCreateConfigVsphere) SetEnableStorageTypeSelection(v string)`

SetEnableStorageTypeSelection sets EnableStorageTypeSelection field to given value.

### HasEnableStorageTypeSelection

`func (o *CloudCreateConfigVsphere) HasEnableStorageTypeSelection() bool`

HasEnableStorageTypeSelection returns a boolean if a field has been set.

### SetEnableStorageTypeSelectionNil

`func (o *CloudCreateConfigVsphere) SetEnableStorageTypeSelectionNil(b bool)`

 SetEnableStorageTypeSelectionNil sets the value for EnableStorageTypeSelection to be an explicit nil

### UnsetEnableStorageTypeSelection
`func (o *CloudCreateConfigVsphere) UnsetEnableStorageTypeSelection()`

UnsetEnableStorageTypeSelection ensures that no value is present for EnableStorageTypeSelection, not even an explicit nil
### GetEnableNetworkTypeSelection

`func (o *CloudCreateConfigVsphere) GetEnableNetworkTypeSelection() CloudCreateConfigVsphereEnableNetworkTypeSelection`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *CloudCreateConfigVsphere) GetEnableNetworkTypeSelectionOk() (*CloudCreateConfigVsphereEnableNetworkTypeSelection, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *CloudCreateConfigVsphere) SetEnableNetworkTypeSelection(v CloudCreateConfigVsphereEnableNetworkTypeSelection)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *CloudCreateConfigVsphere) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### SetEnableNetworkTypeSelectionNil

`func (o *CloudCreateConfigVsphere) SetEnableNetworkTypeSelectionNil(b bool)`

 SetEnableNetworkTypeSelectionNil sets the value for EnableNetworkTypeSelection to be an explicit nil

### UnsetEnableNetworkTypeSelection
`func (o *CloudCreateConfigVsphere) UnsetEnableNetworkTypeSelection()`

UnsetEnableNetworkTypeSelection ensures that no value is present for EnableNetworkTypeSelection, not even an explicit nil
### GetUsername

`func (o *CloudCreateConfigVsphere) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CloudCreateConfigVsphere) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CloudCreateConfigVsphere) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CloudCreateConfigVsphere) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *CloudCreateConfigVsphere) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CloudCreateConfigVsphere) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CloudCreateConfigVsphere) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CloudCreateConfigVsphere) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ZoneCreateConfigAnyOf3

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
**RpcMode** | Pointer to [**NullableZoneCreateConfigAnyOf3RpcMode**](ZoneCreateConfigAnyOf3RpcMode.md) |  | [optional] 
**StorageType** | Pointer to **string** | The default vSphere VMDK type for virtual machines | [optional] [default to "thin"]
**CertificateProvider** | Pointer to **string** | Certificate provider | [optional] [default to "internal"]
**EnableVnc** | Pointer to [**NullableZoneCreateConfigAnyOf3EnableVnc**](ZoneCreateConfigAnyOf3EnableVnc.md) |  | [optional] 
**HideHostSelection** | Pointer to [**NullableZoneCreateConfigAnyOf3HideHostSelection**](ZoneCreateConfigAnyOf3HideHostSelection.md) |  | [optional] 
**EnableDiskTypeSelection** | Pointer to [**NullableZoneCreateConfigAnyOf3EnableDiskTypeSelection**](ZoneCreateConfigAnyOf3EnableDiskTypeSelection.md) |  | [optional] 
**EnableStorageTypeSelection** | Pointer to [**NullableZoneCreateConfigAnyOf3EnableStorageTypeSelection**](ZoneCreateConfigAnyOf3EnableStorageTypeSelection.md) |  | [optional] 
**EnableNetworkTypeSelection** | Pointer to [**NullableZoneCreateConfigAnyOf3EnableNetworkTypeSelection**](ZoneCreateConfigAnyOf3EnableNetworkTypeSelection.md) |  | [optional] 
**Username** | Pointer to **string** | Username. | [optional] 
**Password** | Pointer to **string** | Password to apply to the user | [optional] 

## Methods

### NewZoneCreateConfigAnyOf3

`func NewZoneCreateConfigAnyOf3(apiUrl string, apiVersion string, datacenter string, ) *ZoneCreateConfigAnyOf3`

NewZoneCreateConfigAnyOf3 instantiates a new ZoneCreateConfigAnyOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCreateConfigAnyOf3WithDefaults

`func NewZoneCreateConfigAnyOf3WithDefaults() *ZoneCreateConfigAnyOf3`

NewZoneCreateConfigAnyOf3WithDefaults instantiates a new ZoneCreateConfigAnyOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *ZoneCreateConfigAnyOf3) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ZoneCreateConfigAnyOf3) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ZoneCreateConfigAnyOf3) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ZoneCreateConfigAnyOf3) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *ZoneCreateConfigAnyOf3) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *ZoneCreateConfigAnyOf3) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *ZoneCreateConfigAnyOf3) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *ZoneCreateConfigAnyOf3) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *ZoneCreateConfigAnyOf3) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ZoneCreateConfigAnyOf3) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ZoneCreateConfigAnyOf3) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ZoneCreateConfigAnyOf3) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ZoneCreateConfigAnyOf3) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ZoneCreateConfigAnyOf3) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInventoryLevel

`func (o *ZoneCreateConfigAnyOf3) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *ZoneCreateConfigAnyOf3) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *ZoneCreateConfigAnyOf3) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *ZoneCreateConfigAnyOf3) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *ZoneCreateConfigAnyOf3) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *ZoneCreateConfigAnyOf3) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *ZoneCreateConfigAnyOf3) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *ZoneCreateConfigAnyOf3) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetApiUrl

`func (o *ZoneCreateConfigAnyOf3) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ZoneCreateConfigAnyOf3) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ZoneCreateConfigAnyOf3) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.


### GetApiVersion

`func (o *ZoneCreateConfigAnyOf3) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ZoneCreateConfigAnyOf3) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ZoneCreateConfigAnyOf3) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetDatacenter

`func (o *ZoneCreateConfigAnyOf3) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *ZoneCreateConfigAnyOf3) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *ZoneCreateConfigAnyOf3) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.


### GetCluster

`func (o *ZoneCreateConfigAnyOf3) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ZoneCreateConfigAnyOf3) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ZoneCreateConfigAnyOf3) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ZoneCreateConfigAnyOf3) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *ZoneCreateConfigAnyOf3) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *ZoneCreateConfigAnyOf3) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *ZoneCreateConfigAnyOf3) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *ZoneCreateConfigAnyOf3) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetResourcePool

`func (o *ZoneCreateConfigAnyOf3) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ZoneCreateConfigAnyOf3) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ZoneCreateConfigAnyOf3) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ZoneCreateConfigAnyOf3) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetRpcMode

`func (o *ZoneCreateConfigAnyOf3) GetRpcMode() ZoneCreateConfigAnyOf3RpcMode`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *ZoneCreateConfigAnyOf3) GetRpcModeOk() (*ZoneCreateConfigAnyOf3RpcMode, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *ZoneCreateConfigAnyOf3) SetRpcMode(v ZoneCreateConfigAnyOf3RpcMode)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *ZoneCreateConfigAnyOf3) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### SetRpcModeNil

`func (o *ZoneCreateConfigAnyOf3) SetRpcModeNil(b bool)`

 SetRpcModeNil sets the value for RpcMode to be an explicit nil

### UnsetRpcMode
`func (o *ZoneCreateConfigAnyOf3) UnsetRpcMode()`

UnsetRpcMode ensures that no value is present for RpcMode, not even an explicit nil
### GetStorageType

`func (o *ZoneCreateConfigAnyOf3) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *ZoneCreateConfigAnyOf3) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *ZoneCreateConfigAnyOf3) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *ZoneCreateConfigAnyOf3) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *ZoneCreateConfigAnyOf3) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *ZoneCreateConfigAnyOf3) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *ZoneCreateConfigAnyOf3) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *ZoneCreateConfigAnyOf3) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetEnableVnc

`func (o *ZoneCreateConfigAnyOf3) GetEnableVnc() ZoneCreateConfigAnyOf3EnableVnc`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *ZoneCreateConfigAnyOf3) GetEnableVncOk() (*ZoneCreateConfigAnyOf3EnableVnc, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *ZoneCreateConfigAnyOf3) SetEnableVnc(v ZoneCreateConfigAnyOf3EnableVnc)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *ZoneCreateConfigAnyOf3) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### SetEnableVncNil

`func (o *ZoneCreateConfigAnyOf3) SetEnableVncNil(b bool)`

 SetEnableVncNil sets the value for EnableVnc to be an explicit nil

### UnsetEnableVnc
`func (o *ZoneCreateConfigAnyOf3) UnsetEnableVnc()`

UnsetEnableVnc ensures that no value is present for EnableVnc, not even an explicit nil
### GetHideHostSelection

`func (o *ZoneCreateConfigAnyOf3) GetHideHostSelection() ZoneCreateConfigAnyOf3HideHostSelection`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *ZoneCreateConfigAnyOf3) GetHideHostSelectionOk() (*ZoneCreateConfigAnyOf3HideHostSelection, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *ZoneCreateConfigAnyOf3) SetHideHostSelection(v ZoneCreateConfigAnyOf3HideHostSelection)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *ZoneCreateConfigAnyOf3) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### SetHideHostSelectionNil

`func (o *ZoneCreateConfigAnyOf3) SetHideHostSelectionNil(b bool)`

 SetHideHostSelectionNil sets the value for HideHostSelection to be an explicit nil

### UnsetHideHostSelection
`func (o *ZoneCreateConfigAnyOf3) UnsetHideHostSelection()`

UnsetHideHostSelection ensures that no value is present for HideHostSelection, not even an explicit nil
### GetEnableDiskTypeSelection

`func (o *ZoneCreateConfigAnyOf3) GetEnableDiskTypeSelection() ZoneCreateConfigAnyOf3EnableDiskTypeSelection`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *ZoneCreateConfigAnyOf3) GetEnableDiskTypeSelectionOk() (*ZoneCreateConfigAnyOf3EnableDiskTypeSelection, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *ZoneCreateConfigAnyOf3) SetEnableDiskTypeSelection(v ZoneCreateConfigAnyOf3EnableDiskTypeSelection)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *ZoneCreateConfigAnyOf3) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### SetEnableDiskTypeSelectionNil

`func (o *ZoneCreateConfigAnyOf3) SetEnableDiskTypeSelectionNil(b bool)`

 SetEnableDiskTypeSelectionNil sets the value for EnableDiskTypeSelection to be an explicit nil

### UnsetEnableDiskTypeSelection
`func (o *ZoneCreateConfigAnyOf3) UnsetEnableDiskTypeSelection()`

UnsetEnableDiskTypeSelection ensures that no value is present for EnableDiskTypeSelection, not even an explicit nil
### GetEnableStorageTypeSelection

`func (o *ZoneCreateConfigAnyOf3) GetEnableStorageTypeSelection() ZoneCreateConfigAnyOf3EnableStorageTypeSelection`

GetEnableStorageTypeSelection returns the EnableStorageTypeSelection field if non-nil, zero value otherwise.

### GetEnableStorageTypeSelectionOk

`func (o *ZoneCreateConfigAnyOf3) GetEnableStorageTypeSelectionOk() (*ZoneCreateConfigAnyOf3EnableStorageTypeSelection, bool)`

GetEnableStorageTypeSelectionOk returns a tuple with the EnableStorageTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStorageTypeSelection

`func (o *ZoneCreateConfigAnyOf3) SetEnableStorageTypeSelection(v ZoneCreateConfigAnyOf3EnableStorageTypeSelection)`

SetEnableStorageTypeSelection sets EnableStorageTypeSelection field to given value.

### HasEnableStorageTypeSelection

`func (o *ZoneCreateConfigAnyOf3) HasEnableStorageTypeSelection() bool`

HasEnableStorageTypeSelection returns a boolean if a field has been set.

### SetEnableStorageTypeSelectionNil

`func (o *ZoneCreateConfigAnyOf3) SetEnableStorageTypeSelectionNil(b bool)`

 SetEnableStorageTypeSelectionNil sets the value for EnableStorageTypeSelection to be an explicit nil

### UnsetEnableStorageTypeSelection
`func (o *ZoneCreateConfigAnyOf3) UnsetEnableStorageTypeSelection()`

UnsetEnableStorageTypeSelection ensures that no value is present for EnableStorageTypeSelection, not even an explicit nil
### GetEnableNetworkTypeSelection

`func (o *ZoneCreateConfigAnyOf3) GetEnableNetworkTypeSelection() ZoneCreateConfigAnyOf3EnableNetworkTypeSelection`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *ZoneCreateConfigAnyOf3) GetEnableNetworkTypeSelectionOk() (*ZoneCreateConfigAnyOf3EnableNetworkTypeSelection, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *ZoneCreateConfigAnyOf3) SetEnableNetworkTypeSelection(v ZoneCreateConfigAnyOf3EnableNetworkTypeSelection)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *ZoneCreateConfigAnyOf3) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### SetEnableNetworkTypeSelectionNil

`func (o *ZoneCreateConfigAnyOf3) SetEnableNetworkTypeSelectionNil(b bool)`

 SetEnableNetworkTypeSelectionNil sets the value for EnableNetworkTypeSelection to be an explicit nil

### UnsetEnableNetworkTypeSelection
`func (o *ZoneCreateConfigAnyOf3) UnsetEnableNetworkTypeSelection()`

UnsetEnableNetworkTypeSelection ensures that no value is present for EnableNetworkTypeSelection, not even an explicit nil
### GetUsername

`func (o *ZoneCreateConfigAnyOf3) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ZoneCreateConfigAnyOf3) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ZoneCreateConfigAnyOf3) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ZoneCreateConfigAnyOf3) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ZoneCreateConfigAnyOf3) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ZoneCreateConfigAnyOf3) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ZoneCreateConfigAnyOf3) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ZoneCreateConfigAnyOf3) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AddCloudsRequestZoneConfigAnyOfOneOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiUrl** | **string** | The SDK URL of the vCenter server. | 
**ApiVersion** | **string** | The SDK version of the vCenter server. | 
**Datacenter** | **string** | The vSphere datacenter to add. | 
**Cluster** | Pointer to **string** | The name of the vSphere cluster | [optional] [default to "all"]
**ConfigManagementId** | Pointer to **string** | The id of the configuration management integration associated with the vSphere cloud. | [optional] 
**ResourcePool** | Pointer to **string** | The name of the vSphere resource pool | [optional] 
**RpcMode** | Pointer to **NullableString** |  | [optional] 
**ConsoleKeymap** | Pointer to **string** | The keyboard layout | [optional] [default to "us"]
**StorageType** | Pointer to **string** | The default vSphere VMDK type for virtual machines | [optional] [default to "thin"]
**CertificateProvider** | Pointer to **string** | Certificate provider | [optional] [default to "internal"]
**EnableVnc** | Pointer to **NullableString** |  | [optional] 
**HideHostSelection** | Pointer to **NullableString** |  | [optional] 
**EnableDiskTypeSelection** | Pointer to **NullableString** |  | [optional] 
**EnableStorageTypeSelection** | Pointer to **NullableString** |  | [optional] 
**EnableNetworkTypeSelection** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAddCloudsRequestZoneConfigAnyOfOneOf3

`func NewAddCloudsRequestZoneConfigAnyOfOneOf3(apiUrl string, apiVersion string, datacenter string, ) *AddCloudsRequestZoneConfigAnyOfOneOf3`

NewAddCloudsRequestZoneConfigAnyOfOneOf3 instantiates a new AddCloudsRequestZoneConfigAnyOfOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCloudsRequestZoneConfigAnyOfOneOf3WithDefaults

`func NewAddCloudsRequestZoneConfigAnyOfOneOf3WithDefaults() *AddCloudsRequestZoneConfigAnyOfOneOf3`

NewAddCloudsRequestZoneConfigAnyOfOneOf3WithDefaults instantiates a new AddCloudsRequestZoneConfigAnyOfOneOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUrl

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.


### GetApiVersion

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetDatacenter

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.


### GetCluster

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetResourcePool

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetRpcMode

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetRpcMode() string`

GetRpcMode returns the RpcMode field if non-nil, zero value otherwise.

### GetRpcModeOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetRpcModeOk() (*string, bool)`

GetRpcModeOk returns a tuple with the RpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcMode

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetRpcMode(v string)`

SetRpcMode sets RpcMode field to given value.

### HasRpcMode

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) HasRpcMode() bool`

HasRpcMode returns a boolean if a field has been set.

### SetRpcModeNil

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetRpcModeNil(b bool)`

 SetRpcModeNil sets the value for RpcMode to be an explicit nil

### UnsetRpcMode
`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) UnsetRpcMode()`

UnsetRpcMode ensures that no value is present for RpcMode, not even an explicit nil
### GetConsoleKeymap

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetStorageType

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetEnableVnc

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetEnableVnc() string`

GetEnableVnc returns the EnableVnc field if non-nil, zero value otherwise.

### GetEnableVncOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetEnableVncOk() (*string, bool)`

GetEnableVncOk returns a tuple with the EnableVnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVnc

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetEnableVnc(v string)`

SetEnableVnc sets EnableVnc field to given value.

### HasEnableVnc

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) HasEnableVnc() bool`

HasEnableVnc returns a boolean if a field has been set.

### SetEnableVncNil

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetEnableVncNil(b bool)`

 SetEnableVncNil sets the value for EnableVnc to be an explicit nil

### UnsetEnableVnc
`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) UnsetEnableVnc()`

UnsetEnableVnc ensures that no value is present for EnableVnc, not even an explicit nil
### GetHideHostSelection

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetHideHostSelection() string`

GetHideHostSelection returns the HideHostSelection field if non-nil, zero value otherwise.

### GetHideHostSelectionOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetHideHostSelectionOk() (*string, bool)`

GetHideHostSelectionOk returns a tuple with the HideHostSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideHostSelection

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetHideHostSelection(v string)`

SetHideHostSelection sets HideHostSelection field to given value.

### HasHideHostSelection

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) HasHideHostSelection() bool`

HasHideHostSelection returns a boolean if a field has been set.

### SetHideHostSelectionNil

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetHideHostSelectionNil(b bool)`

 SetHideHostSelectionNil sets the value for HideHostSelection to be an explicit nil

### UnsetHideHostSelection
`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) UnsetHideHostSelection()`

UnsetHideHostSelection ensures that no value is present for HideHostSelection, not even an explicit nil
### GetEnableDiskTypeSelection

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetEnableDiskTypeSelection() string`

GetEnableDiskTypeSelection returns the EnableDiskTypeSelection field if non-nil, zero value otherwise.

### GetEnableDiskTypeSelectionOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetEnableDiskTypeSelectionOk() (*string, bool)`

GetEnableDiskTypeSelectionOk returns a tuple with the EnableDiskTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDiskTypeSelection

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetEnableDiskTypeSelection(v string)`

SetEnableDiskTypeSelection sets EnableDiskTypeSelection field to given value.

### HasEnableDiskTypeSelection

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) HasEnableDiskTypeSelection() bool`

HasEnableDiskTypeSelection returns a boolean if a field has been set.

### SetEnableDiskTypeSelectionNil

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetEnableDiskTypeSelectionNil(b bool)`

 SetEnableDiskTypeSelectionNil sets the value for EnableDiskTypeSelection to be an explicit nil

### UnsetEnableDiskTypeSelection
`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) UnsetEnableDiskTypeSelection()`

UnsetEnableDiskTypeSelection ensures that no value is present for EnableDiskTypeSelection, not even an explicit nil
### GetEnableStorageTypeSelection

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetEnableStorageTypeSelection() string`

GetEnableStorageTypeSelection returns the EnableStorageTypeSelection field if non-nil, zero value otherwise.

### GetEnableStorageTypeSelectionOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetEnableStorageTypeSelectionOk() (*string, bool)`

GetEnableStorageTypeSelectionOk returns a tuple with the EnableStorageTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStorageTypeSelection

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetEnableStorageTypeSelection(v string)`

SetEnableStorageTypeSelection sets EnableStorageTypeSelection field to given value.

### HasEnableStorageTypeSelection

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) HasEnableStorageTypeSelection() bool`

HasEnableStorageTypeSelection returns a boolean if a field has been set.

### SetEnableStorageTypeSelectionNil

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetEnableStorageTypeSelectionNil(b bool)`

 SetEnableStorageTypeSelectionNil sets the value for EnableStorageTypeSelection to be an explicit nil

### UnsetEnableStorageTypeSelection
`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) UnsetEnableStorageTypeSelection()`

UnsetEnableStorageTypeSelection ensures that no value is present for EnableStorageTypeSelection, not even an explicit nil
### GetEnableNetworkTypeSelection

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### SetEnableNetworkTypeSelectionNil

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) SetEnableNetworkTypeSelectionNil(b bool)`

 SetEnableNetworkTypeSelectionNil sets the value for EnableNetworkTypeSelection to be an explicit nil

### UnsetEnableNetworkTypeSelection
`func (o *AddCloudsRequestZoneConfigAnyOfOneOf3) UnsetEnableNetworkTypeSelection()`

UnsetEnableNetworkTypeSelection ensures that no value is present for EnableNetworkTypeSelection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



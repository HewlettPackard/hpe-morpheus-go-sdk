# CloudHVMConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | Pointer to **NullableString** |  | [optional] 
**CertificateProvider** | Pointer to **NullableString** |  | [optional] 
**ConfigCmdbDiscovery** | Pointer to **bool** |  | [optional] 
**ConsoleKeymap** | Pointer to **NullableString** |  | [optional] 
**DatacenterName** | Pointer to **NullableString** |  | [optional] 
**EnableNetworkTypeSelection** | Pointer to **bool** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**InventoryLevel** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCloudHVMConfig

`func NewCloudHVMConfig() *CloudHVMConfig`

NewCloudHVMConfig instantiates a new CloudHVMConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudHVMConfigWithDefaults

`func NewCloudHVMConfigWithDefaults() *CloudHVMConfig`

NewCloudHVMConfigWithDefaults instantiates a new CloudHVMConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *CloudHVMConfig) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *CloudHVMConfig) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *CloudHVMConfig) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *CloudHVMConfig) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### SetApplianceUrlNil

`func (o *CloudHVMConfig) SetApplianceUrlNil(b bool)`

 SetApplianceUrlNil sets the value for ApplianceUrl to be an explicit nil

### UnsetApplianceUrl
`func (o *CloudHVMConfig) UnsetApplianceUrl()`

UnsetApplianceUrl ensures that no value is present for ApplianceUrl, not even an explicit nil
### GetCertificateProvider

`func (o *CloudHVMConfig) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *CloudHVMConfig) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *CloudHVMConfig) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *CloudHVMConfig) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### SetCertificateProviderNil

`func (o *CloudHVMConfig) SetCertificateProviderNil(b bool)`

 SetCertificateProviderNil sets the value for CertificateProvider to be an explicit nil

### UnsetCertificateProvider
`func (o *CloudHVMConfig) UnsetCertificateProvider()`

UnsetCertificateProvider ensures that no value is present for CertificateProvider, not even an explicit nil
### GetConfigCmdbDiscovery

`func (o *CloudHVMConfig) GetConfigCmdbDiscovery() bool`

GetConfigCmdbDiscovery returns the ConfigCmdbDiscovery field if non-nil, zero value otherwise.

### GetConfigCmdbDiscoveryOk

`func (o *CloudHVMConfig) GetConfigCmdbDiscoveryOk() (*bool, bool)`

GetConfigCmdbDiscoveryOk returns a tuple with the ConfigCmdbDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCmdbDiscovery

`func (o *CloudHVMConfig) SetConfigCmdbDiscovery(v bool)`

SetConfigCmdbDiscovery sets ConfigCmdbDiscovery field to given value.

### HasConfigCmdbDiscovery

`func (o *CloudHVMConfig) HasConfigCmdbDiscovery() bool`

HasConfigCmdbDiscovery returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *CloudHVMConfig) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *CloudHVMConfig) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *CloudHVMConfig) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *CloudHVMConfig) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### SetConsoleKeymapNil

`func (o *CloudHVMConfig) SetConsoleKeymapNil(b bool)`

 SetConsoleKeymapNil sets the value for ConsoleKeymap to be an explicit nil

### UnsetConsoleKeymap
`func (o *CloudHVMConfig) UnsetConsoleKeymap()`

UnsetConsoleKeymap ensures that no value is present for ConsoleKeymap, not even an explicit nil
### GetDatacenterName

`func (o *CloudHVMConfig) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *CloudHVMConfig) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *CloudHVMConfig) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *CloudHVMConfig) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### SetDatacenterNameNil

`func (o *CloudHVMConfig) SetDatacenterNameNil(b bool)`

 SetDatacenterNameNil sets the value for DatacenterName to be an explicit nil

### UnsetDatacenterName
`func (o *CloudHVMConfig) UnsetDatacenterName()`

UnsetDatacenterName ensures that no value is present for DatacenterName, not even an explicit nil
### GetEnableNetworkTypeSelection

`func (o *CloudHVMConfig) GetEnableNetworkTypeSelection() bool`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *CloudHVMConfig) GetEnableNetworkTypeSelectionOk() (*bool, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *CloudHVMConfig) SetEnableNetworkTypeSelection(v bool)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *CloudHVMConfig) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudHVMConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudHVMConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudHVMConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudHVMConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CloudHVMConfig) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CloudHVMConfig) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInventoryLevel

`func (o *CloudHVMConfig) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *CloudHVMConfig) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *CloudHVMConfig) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *CloudHVMConfig) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### SetInventoryLevelNil

`func (o *CloudHVMConfig) SetInventoryLevelNil(b bool)`

 SetInventoryLevelNil sets the value for InventoryLevel to be an explicit nil

### UnsetInventoryLevel
`func (o *CloudHVMConfig) UnsetInventoryLevel()`

UnsetInventoryLevel ensures that no value is present for InventoryLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



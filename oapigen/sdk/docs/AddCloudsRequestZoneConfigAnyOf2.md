# AddCloudsRequestZoneConfigAnyOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | Pointer to **string** | The URL used by workloads provisioned in the cloud for interacting with the Morpheus appliance. | [optional] 
**DatacenterName** | Pointer to **string** | A custom name used to reference the datacenter for the cloud. | [optional] 
**ExternalId** | Pointer to **NullableString** | The external id of the cloud | [optional] 
**InventoryLevel** | Pointer to **string** | Whether to import existing virtual machines. | [optional] 
**ConsoleKeymap** | Pointer to **string** | The keyboard layout to use for the console | [optional] 
**CertificateProvider** | Pointer to **string** | Certificate provider | [optional] [default to "internal"]
**EnableNetworkTypeSelection** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAddCloudsRequestZoneConfigAnyOf2

`func NewAddCloudsRequestZoneConfigAnyOf2() *AddCloudsRequestZoneConfigAnyOf2`

NewAddCloudsRequestZoneConfigAnyOf2 instantiates a new AddCloudsRequestZoneConfigAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCloudsRequestZoneConfigAnyOf2WithDefaults

`func NewAddCloudsRequestZoneConfigAnyOf2WithDefaults() *AddCloudsRequestZoneConfigAnyOf2`

NewAddCloudsRequestZoneConfigAnyOf2WithDefaults instantiates a new AddCloudsRequestZoneConfigAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *AddCloudsRequestZoneConfigAnyOf2) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *AddCloudsRequestZoneConfigAnyOf2) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetDatacenterName

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *AddCloudsRequestZoneConfigAnyOf2) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *AddCloudsRequestZoneConfigAnyOf2) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetExternalId

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddCloudsRequestZoneConfigAnyOf2) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddCloudsRequestZoneConfigAnyOf2) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AddCloudsRequestZoneConfigAnyOf2) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AddCloudsRequestZoneConfigAnyOf2) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInventoryLevel

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *AddCloudsRequestZoneConfigAnyOf2) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *AddCloudsRequestZoneConfigAnyOf2) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *AddCloudsRequestZoneConfigAnyOf2) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *AddCloudsRequestZoneConfigAnyOf2) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetCertificateProvider

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *AddCloudsRequestZoneConfigAnyOf2) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *AddCloudsRequestZoneConfigAnyOf2) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *AddCloudsRequestZoneConfigAnyOf2) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *AddCloudsRequestZoneConfigAnyOf2) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *AddCloudsRequestZoneConfigAnyOf2) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### SetEnableNetworkTypeSelectionNil

`func (o *AddCloudsRequestZoneConfigAnyOf2) SetEnableNetworkTypeSelectionNil(b bool)`

 SetEnableNetworkTypeSelectionNil sets the value for EnableNetworkTypeSelection to be an explicit nil

### UnsetEnableNetworkTypeSelection
`func (o *AddCloudsRequestZoneConfigAnyOf2) UnsetEnableNetworkTypeSelection()`

UnsetEnableNetworkTypeSelection ensures that no value is present for EnableNetworkTypeSelection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudCreateConfigStandard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateProvider** | Pointer to **string** | Certificate provider | [optional] [default to "internal"]
**EnableNetworkTypeSelection** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCloudCreateConfigStandard

`func NewCloudCreateConfigStandard() *CloudCreateConfigStandard`

NewCloudCreateConfigStandard instantiates a new CloudCreateConfigStandard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCreateConfigStandardWithDefaults

`func NewCloudCreateConfigStandardWithDefaults() *CloudCreateConfigStandard`

NewCloudCreateConfigStandardWithDefaults instantiates a new CloudCreateConfigStandard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateProvider

`func (o *CloudCreateConfigStandard) GetCertificateProvider() string`

GetCertificateProvider returns the CertificateProvider field if non-nil, zero value otherwise.

### GetCertificateProviderOk

`func (o *CloudCreateConfigStandard) GetCertificateProviderOk() (*string, bool)`

GetCertificateProviderOk returns a tuple with the CertificateProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProvider

`func (o *CloudCreateConfigStandard) SetCertificateProvider(v string)`

SetCertificateProvider sets CertificateProvider field to given value.

### HasCertificateProvider

`func (o *CloudCreateConfigStandard) HasCertificateProvider() bool`

HasCertificateProvider returns a boolean if a field has been set.

### GetEnableNetworkTypeSelection

`func (o *CloudCreateConfigStandard) GetEnableNetworkTypeSelection() string`

GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field if non-nil, zero value otherwise.

### GetEnableNetworkTypeSelectionOk

`func (o *CloudCreateConfigStandard) GetEnableNetworkTypeSelectionOk() (*string, bool)`

GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNetworkTypeSelection

`func (o *CloudCreateConfigStandard) SetEnableNetworkTypeSelection(v string)`

SetEnableNetworkTypeSelection sets EnableNetworkTypeSelection field to given value.

### HasEnableNetworkTypeSelection

`func (o *CloudCreateConfigStandard) HasEnableNetworkTypeSelection() bool`

HasEnableNetworkTypeSelection returns a boolean if a field has been set.

### SetEnableNetworkTypeSelectionNil

`func (o *CloudCreateConfigStandard) SetEnableNetworkTypeSelectionNil(b bool)`

 SetEnableNetworkTypeSelectionNil sets the value for EnableNetworkTypeSelection to be an explicit nil

### UnsetEnableNetworkTypeSelection
`func (o *CloudCreateConfigStandard) UnsetEnableNetworkTypeSelection()`

UnsetEnableNetworkTypeSelection ensures that no value is present for EnableNetworkTypeSelection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



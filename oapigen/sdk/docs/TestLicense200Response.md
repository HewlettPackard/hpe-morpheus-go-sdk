# TestLicense200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**License** | Pointer to [**InstallLicense200ResponseLicense**](InstallLicense200ResponseLicense.md) |  | [optional] 
**InstalledLicenses** | Pointer to [**[]InstallLicense200ResponseInstalledLicensesInner**](InstallLicense200ResponseInstalledLicensesInner.md) | List of all the installed licenses | [optional] 
**CurrentUsage** | Pointer to [**InstallLicense200ResponseCurrentUsage**](InstallLicense200ResponseCurrentUsage.md) |  | [optional] 

## Methods

### NewTestLicense200Response

`func NewTestLicense200Response() *TestLicense200Response`

NewTestLicense200Response instantiates a new TestLicense200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestLicense200ResponseWithDefaults

`func NewTestLicense200ResponseWithDefaults() *TestLicense200Response`

NewTestLicense200ResponseWithDefaults instantiates a new TestLicense200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicense

`func (o *TestLicense200Response) GetLicense() InstallLicense200ResponseLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *TestLicense200Response) GetLicenseOk() (*InstallLicense200ResponseLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *TestLicense200Response) SetLicense(v InstallLicense200ResponseLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *TestLicense200Response) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetInstalledLicenses

`func (o *TestLicense200Response) GetInstalledLicenses() []InstallLicense200ResponseInstalledLicensesInner`

GetInstalledLicenses returns the InstalledLicenses field if non-nil, zero value otherwise.

### GetInstalledLicensesOk

`func (o *TestLicense200Response) GetInstalledLicensesOk() (*[]InstallLicense200ResponseInstalledLicensesInner, bool)`

GetInstalledLicensesOk returns a tuple with the InstalledLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledLicenses

`func (o *TestLicense200Response) SetInstalledLicenses(v []InstallLicense200ResponseInstalledLicensesInner)`

SetInstalledLicenses sets InstalledLicenses field to given value.

### HasInstalledLicenses

`func (o *TestLicense200Response) HasInstalledLicenses() bool`

HasInstalledLicenses returns a boolean if a field has been set.

### GetCurrentUsage

`func (o *TestLicense200Response) GetCurrentUsage() InstallLicense200ResponseCurrentUsage`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *TestLicense200Response) GetCurrentUsageOk() (*InstallLicense200ResponseCurrentUsage, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *TestLicense200Response) SetCurrentUsage(v InstallLicense200ResponseCurrentUsage)`

SetCurrentUsage sets CurrentUsage field to given value.

### HasCurrentUsage

`func (o *TestLicense200Response) HasCurrentUsage() bool`

HasCurrentUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



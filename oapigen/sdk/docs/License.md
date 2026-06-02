# License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**License** | Pointer to [**InstallLicense200ResponseLicense**](InstallLicense200ResponseLicense.md) |  | [optional] 
**InstalledLicenses** | Pointer to [**[]InstallLicense200ResponseInstalledLicensesInner**](InstallLicense200ResponseInstalledLicensesInner.md) | List of all the installed licenses | [optional] 
**CurrentUsage** | Pointer to [**InstallLicense200ResponseCurrentUsage**](InstallLicense200ResponseCurrentUsage.md) |  | [optional] 

## Methods

### NewLicense

`func NewLicense() *License`

NewLicense instantiates a new License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetLicense

`func (o *License) GetLicense() InstallLicense200ResponseLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *License) GetLicenseOk() (*InstallLicense200ResponseLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *License) SetLicense(v InstallLicense200ResponseLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *License) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetInstalledLicenses

`func (o *License) GetInstalledLicenses() []InstallLicense200ResponseInstalledLicensesInner`

GetInstalledLicenses returns the InstalledLicenses field if non-nil, zero value otherwise.

### GetInstalledLicensesOk

`func (o *License) GetInstalledLicensesOk() (*[]InstallLicense200ResponseInstalledLicensesInner, bool)`

GetInstalledLicensesOk returns a tuple with the InstalledLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledLicenses

`func (o *License) SetInstalledLicenses(v []InstallLicense200ResponseInstalledLicensesInner)`

SetInstalledLicenses sets InstalledLicenses field to given value.

### HasInstalledLicenses

`func (o *License) HasInstalledLicenses() bool`

HasInstalledLicenses returns a boolean if a field has been set.

### GetCurrentUsage

`func (o *License) GetCurrentUsage() InstallLicense200ResponseCurrentUsage`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *License) GetCurrentUsageOk() (*InstallLicense200ResponseCurrentUsage, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *License) SetCurrentUsage(v InstallLicense200ResponseCurrentUsage)`

SetCurrentUsage sets CurrentUsage field to given value.

### HasCurrentUsage

`func (o *License) HasCurrentUsage() bool`

HasCurrentUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



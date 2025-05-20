# GetLicense200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**License** | Pointer to [**GetLicense200ResponseLicense**](GetLicense200ResponseLicense.md) |  | [optional] 
**InstalledLicenses** | Pointer to [**[]GetLicense200ResponseInstalledLicensesInner**](GetLicense200ResponseInstalledLicensesInner.md) | List of all the installed licenses | [optional] 
**CurrentUsage** | Pointer to [**GetLicense200ResponseCurrentUsage**](GetLicense200ResponseCurrentUsage.md) |  | [optional] 

## Methods

### NewGetLicense200Response

`func NewGetLicense200Response() *GetLicense200Response`

NewGetLicense200Response instantiates a new GetLicense200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLicense200ResponseWithDefaults

`func NewGetLicense200ResponseWithDefaults() *GetLicense200Response`

NewGetLicense200ResponseWithDefaults instantiates a new GetLicense200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicense

`func (o *GetLicense200Response) GetLicense() GetLicense200ResponseLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *GetLicense200Response) GetLicenseOk() (*GetLicense200ResponseLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *GetLicense200Response) SetLicense(v GetLicense200ResponseLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *GetLicense200Response) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetInstalledLicenses

`func (o *GetLicense200Response) GetInstalledLicenses() []GetLicense200ResponseInstalledLicensesInner`

GetInstalledLicenses returns the InstalledLicenses field if non-nil, zero value otherwise.

### GetInstalledLicensesOk

`func (o *GetLicense200Response) GetInstalledLicensesOk() (*[]GetLicense200ResponseInstalledLicensesInner, bool)`

GetInstalledLicensesOk returns a tuple with the InstalledLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledLicenses

`func (o *GetLicense200Response) SetInstalledLicenses(v []GetLicense200ResponseInstalledLicensesInner)`

SetInstalledLicenses sets InstalledLicenses field to given value.

### HasInstalledLicenses

`func (o *GetLicense200Response) HasInstalledLicenses() bool`

HasInstalledLicenses returns a boolean if a field has been set.

### GetCurrentUsage

`func (o *GetLicense200Response) GetCurrentUsage() GetLicense200ResponseCurrentUsage`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *GetLicense200Response) GetCurrentUsageOk() (*GetLicense200ResponseCurrentUsage, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *GetLicense200Response) SetCurrentUsage(v GetLicense200ResponseCurrentUsage)`

SetCurrentUsage sets CurrentUsage field to given value.

### HasCurrentUsage

`func (o *GetLicense200Response) HasCurrentUsage() bool`

HasCurrentUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InstallLicenseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**License** | **string** | License Key. This is a long and unique string of your Morpheus license. License keys can be requested from the [Morpheus Hub](https://morpheushub.com). | 
**InstallAction** | Pointer to **string** | Install Action can be passed as &#39;add&#39; to stack the license. By default all currently installed licenses are removed and replaced by the new one. | [optional] [default to "replace"]

## Methods

### NewInstallLicenseRequest

`func NewInstallLicenseRequest(license string, ) *InstallLicenseRequest`

NewInstallLicenseRequest instantiates a new InstallLicenseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallLicenseRequestWithDefaults

`func NewInstallLicenseRequestWithDefaults() *InstallLicenseRequest`

NewInstallLicenseRequestWithDefaults instantiates a new InstallLicenseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicense

`func (o *InstallLicenseRequest) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *InstallLicenseRequest) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *InstallLicenseRequest) SetLicense(v string)`

SetLicense sets License field to given value.


### GetInstallAction

`func (o *InstallLicenseRequest) GetInstallAction() string`

GetInstallAction returns the InstallAction field if non-nil, zero value otherwise.

### GetInstallActionOk

`func (o *InstallLicenseRequest) GetInstallActionOk() (*string, bool)`

GetInstallActionOk returns a tuple with the InstallAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAction

`func (o *InstallLicenseRequest) SetInstallAction(v string)`

SetInstallAction sets InstallAction field to given value.

### HasInstallAction

`func (o *InstallLicenseRequest) HasInstallAction() bool`

HasInstallAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



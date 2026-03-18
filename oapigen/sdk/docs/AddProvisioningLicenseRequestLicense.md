# AddProvisioningLicenseRequestLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**LicenseType** | **string** | License Type - The license type code. | 
**LicenseKey** | **string** | License Key - The license key, to be kept a secret. | 
**OrgName** | Pointer to **string** | Org Name - The Organization Name (if applicable) related to the license key | [optional] 
**FullName** | Pointer to **string** | Full Name - The Full Name (if applicable) related to the license key | [optional] 
**LicenseVersion** | Pointer to **string** | License Version | [optional] 
**Copies** | Pointer to **int64** | Copies - The number of times the key can be used. | [optional] [default to 1]
**Description** | Pointer to **string** | Description | [optional] 
**VirtualImages** | Pointer to **[]int64** | Virtual Images - Array of Virtual Image IDs to associate with license. | [optional] 
**Tenants** | Pointer to **[]int64** | Tenants - Array of tenants that are allowed to use the key. | [optional] 

## Methods

### NewAddProvisioningLicenseRequestLicense

`func NewAddProvisioningLicenseRequestLicense(name string, licenseType string, licenseKey string, ) *AddProvisioningLicenseRequestLicense`

NewAddProvisioningLicenseRequestLicense instantiates a new AddProvisioningLicenseRequestLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddProvisioningLicenseRequestLicenseWithDefaults

`func NewAddProvisioningLicenseRequestLicenseWithDefaults() *AddProvisioningLicenseRequestLicense`

NewAddProvisioningLicenseRequestLicenseWithDefaults instantiates a new AddProvisioningLicenseRequestLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddProvisioningLicenseRequestLicense) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddProvisioningLicenseRequestLicense) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddProvisioningLicenseRequestLicense) SetName(v string)`

SetName sets Name field to given value.


### GetLicenseType

`func (o *AddProvisioningLicenseRequestLicense) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *AddProvisioningLicenseRequestLicense) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *AddProvisioningLicenseRequestLicense) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.


### GetLicenseKey

`func (o *AddProvisioningLicenseRequestLicense) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *AddProvisioningLicenseRequestLicense) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *AddProvisioningLicenseRequestLicense) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.


### GetOrgName

`func (o *AddProvisioningLicenseRequestLicense) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *AddProvisioningLicenseRequestLicense) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *AddProvisioningLicenseRequestLicense) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *AddProvisioningLicenseRequestLicense) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetFullName

`func (o *AddProvisioningLicenseRequestLicense) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AddProvisioningLicenseRequestLicense) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AddProvisioningLicenseRequestLicense) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AddProvisioningLicenseRequestLicense) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetLicenseVersion

`func (o *AddProvisioningLicenseRequestLicense) GetLicenseVersion() string`

GetLicenseVersion returns the LicenseVersion field if non-nil, zero value otherwise.

### GetLicenseVersionOk

`func (o *AddProvisioningLicenseRequestLicense) GetLicenseVersionOk() (*string, bool)`

GetLicenseVersionOk returns a tuple with the LicenseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseVersion

`func (o *AddProvisioningLicenseRequestLicense) SetLicenseVersion(v string)`

SetLicenseVersion sets LicenseVersion field to given value.

### HasLicenseVersion

`func (o *AddProvisioningLicenseRequestLicense) HasLicenseVersion() bool`

HasLicenseVersion returns a boolean if a field has been set.

### GetCopies

`func (o *AddProvisioningLicenseRequestLicense) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *AddProvisioningLicenseRequestLicense) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *AddProvisioningLicenseRequestLicense) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *AddProvisioningLicenseRequestLicense) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetDescription

`func (o *AddProvisioningLicenseRequestLicense) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddProvisioningLicenseRequestLicense) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddProvisioningLicenseRequestLicense) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddProvisioningLicenseRequestLicense) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVirtualImages

`func (o *AddProvisioningLicenseRequestLicense) GetVirtualImages() []int64`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *AddProvisioningLicenseRequestLicense) GetVirtualImagesOk() (*[]int64, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *AddProvisioningLicenseRequestLicense) SetVirtualImages(v []int64)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *AddProvisioningLicenseRequestLicense) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetTenants

`func (o *AddProvisioningLicenseRequestLicense) GetTenants() []int64`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *AddProvisioningLicenseRequestLicense) GetTenantsOk() (*[]int64, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *AddProvisioningLicenseRequestLicense) SetTenants(v []int64)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *AddProvisioningLicenseRequestLicense) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



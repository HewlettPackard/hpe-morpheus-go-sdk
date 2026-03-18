# ProvisioningLicensesCreate

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

### NewProvisioningLicensesCreate

`func NewProvisioningLicensesCreate(name string, licenseType string, licenseKey string, ) *ProvisioningLicensesCreate`

NewProvisioningLicensesCreate instantiates a new ProvisioningLicensesCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningLicensesCreateWithDefaults

`func NewProvisioningLicensesCreateWithDefaults() *ProvisioningLicensesCreate`

NewProvisioningLicensesCreateWithDefaults instantiates a new ProvisioningLicensesCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProvisioningLicensesCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningLicensesCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningLicensesCreate) SetName(v string)`

SetName sets Name field to given value.


### GetLicenseType

`func (o *ProvisioningLicensesCreate) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *ProvisioningLicensesCreate) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *ProvisioningLicensesCreate) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.


### GetLicenseKey

`func (o *ProvisioningLicensesCreate) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *ProvisioningLicensesCreate) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *ProvisioningLicensesCreate) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.


### GetOrgName

`func (o *ProvisioningLicensesCreate) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ProvisioningLicensesCreate) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ProvisioningLicensesCreate) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *ProvisioningLicensesCreate) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetFullName

`func (o *ProvisioningLicensesCreate) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ProvisioningLicensesCreate) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ProvisioningLicensesCreate) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ProvisioningLicensesCreate) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetLicenseVersion

`func (o *ProvisioningLicensesCreate) GetLicenseVersion() string`

GetLicenseVersion returns the LicenseVersion field if non-nil, zero value otherwise.

### GetLicenseVersionOk

`func (o *ProvisioningLicensesCreate) GetLicenseVersionOk() (*string, bool)`

GetLicenseVersionOk returns a tuple with the LicenseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseVersion

`func (o *ProvisioningLicensesCreate) SetLicenseVersion(v string)`

SetLicenseVersion sets LicenseVersion field to given value.

### HasLicenseVersion

`func (o *ProvisioningLicensesCreate) HasLicenseVersion() bool`

HasLicenseVersion returns a boolean if a field has been set.

### GetCopies

`func (o *ProvisioningLicensesCreate) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *ProvisioningLicensesCreate) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *ProvisioningLicensesCreate) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *ProvisioningLicensesCreate) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetDescription

`func (o *ProvisioningLicensesCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProvisioningLicensesCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProvisioningLicensesCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProvisioningLicensesCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVirtualImages

`func (o *ProvisioningLicensesCreate) GetVirtualImages() []int64`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *ProvisioningLicensesCreate) GetVirtualImagesOk() (*[]int64, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *ProvisioningLicensesCreate) SetVirtualImages(v []int64)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *ProvisioningLicensesCreate) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetTenants

`func (o *ProvisioningLicensesCreate) GetTenants() []int64`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ProvisioningLicensesCreate) GetTenantsOk() (*[]int64, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ProvisioningLicensesCreate) SetTenants(v []int64)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ProvisioningLicensesCreate) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateProvisioningLicenseRequestLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**LicenseVersion** | Pointer to **string** | License Version | [optional] 
**Copies** | Pointer to **int64** | Copies - The number of times the key can be used. | [optional] [default to 1]
**Description** | Pointer to **string** | Description | [optional] 
**VirtualImages** | Pointer to **[]int64** | Virtual Images - Array of Virtual Image IDs to associate with license. | [optional] 
**Tenants** | Pointer to **[]int64** | Tenants - Array of tenants that are allowed to use the key. | [optional] 

## Methods

### NewUpdateProvisioningLicenseRequestLicense

`func NewUpdateProvisioningLicenseRequestLicense() *UpdateProvisioningLicenseRequestLicense`

NewUpdateProvisioningLicenseRequestLicense instantiates a new UpdateProvisioningLicenseRequestLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProvisioningLicenseRequestLicenseWithDefaults

`func NewUpdateProvisioningLicenseRequestLicenseWithDefaults() *UpdateProvisioningLicenseRequestLicense`

NewUpdateProvisioningLicenseRequestLicenseWithDefaults instantiates a new UpdateProvisioningLicenseRequestLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateProvisioningLicenseRequestLicense) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProvisioningLicenseRequestLicense) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProvisioningLicenseRequestLicense) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProvisioningLicenseRequestLicense) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLicenseVersion

`func (o *UpdateProvisioningLicenseRequestLicense) GetLicenseVersion() string`

GetLicenseVersion returns the LicenseVersion field if non-nil, zero value otherwise.

### GetLicenseVersionOk

`func (o *UpdateProvisioningLicenseRequestLicense) GetLicenseVersionOk() (*string, bool)`

GetLicenseVersionOk returns a tuple with the LicenseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseVersion

`func (o *UpdateProvisioningLicenseRequestLicense) SetLicenseVersion(v string)`

SetLicenseVersion sets LicenseVersion field to given value.

### HasLicenseVersion

`func (o *UpdateProvisioningLicenseRequestLicense) HasLicenseVersion() bool`

HasLicenseVersion returns a boolean if a field has been set.

### GetCopies

`func (o *UpdateProvisioningLicenseRequestLicense) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *UpdateProvisioningLicenseRequestLicense) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *UpdateProvisioningLicenseRequestLicense) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *UpdateProvisioningLicenseRequestLicense) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateProvisioningLicenseRequestLicense) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProvisioningLicenseRequestLicense) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProvisioningLicenseRequestLicense) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProvisioningLicenseRequestLicense) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVirtualImages

`func (o *UpdateProvisioningLicenseRequestLicense) GetVirtualImages() []int64`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *UpdateProvisioningLicenseRequestLicense) GetVirtualImagesOk() (*[]int64, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *UpdateProvisioningLicenseRequestLicense) SetVirtualImages(v []int64)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *UpdateProvisioningLicenseRequestLicense) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateProvisioningLicenseRequestLicense) GetTenants() []int64`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateProvisioningLicenseRequestLicense) GetTenantsOk() (*[]int64, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateProvisioningLicenseRequestLicense) SetTenants(v []int64)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateProvisioningLicenseRequestLicense) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



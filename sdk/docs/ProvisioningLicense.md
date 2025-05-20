# ProvisioningLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**LicenseType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**LicenseKey** | Pointer to **string** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**LicenseVersion** | Pointer to **string** |  | [optional] 
**Copies** | Pointer to **int64** |  | [optional] 
**ReservationCount** | Pointer to **int64** |  | [optional] 
**Tenants** | Pointer to **[]map[string]interface{}** |  | [optional] 
**VirtualImages** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewProvisioningLicense

`func NewProvisioningLicense() *ProvisioningLicense`

NewProvisioningLicense instantiates a new ProvisioningLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningLicenseWithDefaults

`func NewProvisioningLicenseWithDefaults() *ProvisioningLicense`

NewProvisioningLicenseWithDefaults instantiates a new ProvisioningLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvisioningLicense) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningLicense) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningLicense) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisioningLicense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProvisioningLicense) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningLicense) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningLicense) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisioningLicense) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProvisioningLicense) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProvisioningLicense) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProvisioningLicense) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProvisioningLicense) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicenseType

`func (o *ProvisioningLicense) GetLicenseType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *ProvisioningLicense) GetLicenseTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *ProvisioningLicense) SetLicenseType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *ProvisioningLicense) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetLicenseKey

`func (o *ProvisioningLicense) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *ProvisioningLicense) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *ProvisioningLicense) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *ProvisioningLicense) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetOrgName

`func (o *ProvisioningLicense) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ProvisioningLicense) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ProvisioningLicense) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *ProvisioningLicense) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetFullName

`func (o *ProvisioningLicense) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ProvisioningLicense) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ProvisioningLicense) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ProvisioningLicense) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetLicenseVersion

`func (o *ProvisioningLicense) GetLicenseVersion() string`

GetLicenseVersion returns the LicenseVersion field if non-nil, zero value otherwise.

### GetLicenseVersionOk

`func (o *ProvisioningLicense) GetLicenseVersionOk() (*string, bool)`

GetLicenseVersionOk returns a tuple with the LicenseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseVersion

`func (o *ProvisioningLicense) SetLicenseVersion(v string)`

SetLicenseVersion sets LicenseVersion field to given value.

### HasLicenseVersion

`func (o *ProvisioningLicense) HasLicenseVersion() bool`

HasLicenseVersion returns a boolean if a field has been set.

### GetCopies

`func (o *ProvisioningLicense) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *ProvisioningLicense) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *ProvisioningLicense) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *ProvisioningLicense) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetReservationCount

`func (o *ProvisioningLicense) GetReservationCount() int64`

GetReservationCount returns the ReservationCount field if non-nil, zero value otherwise.

### GetReservationCountOk

`func (o *ProvisioningLicense) GetReservationCountOk() (*int64, bool)`

GetReservationCountOk returns a tuple with the ReservationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationCount

`func (o *ProvisioningLicense) SetReservationCount(v int64)`

SetReservationCount sets ReservationCount field to given value.

### HasReservationCount

`func (o *ProvisioningLicense) HasReservationCount() bool`

HasReservationCount returns a boolean if a field has been set.

### GetTenants

`func (o *ProvisioningLicense) GetTenants() []map[string]interface{}`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ProvisioningLicense) GetTenantsOk() (*[]map[string]interface{}, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ProvisioningLicense) SetTenants(v []map[string]interface{})`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ProvisioningLicense) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetVirtualImages

`func (o *ProvisioningLicense) GetVirtualImages() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *ProvisioningLicense) GetVirtualImagesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *ProvisioningLicense) SetVirtualImages(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *ProvisioningLicense) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetAccount

`func (o *ProvisioningLicense) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ProvisioningLicense) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ProvisioningLicense) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ProvisioningLicense) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



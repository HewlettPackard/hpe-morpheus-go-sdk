# ListProvisioningLicenses200ResponseAllOfLicensesInner

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

### NewListProvisioningLicenses200ResponseAllOfLicensesInner

`func NewListProvisioningLicenses200ResponseAllOfLicensesInner() *ListProvisioningLicenses200ResponseAllOfLicensesInner`

NewListProvisioningLicenses200ResponseAllOfLicensesInner instantiates a new ListProvisioningLicenses200ResponseAllOfLicensesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProvisioningLicenses200ResponseAllOfLicensesInnerWithDefaults

`func NewListProvisioningLicenses200ResponseAllOfLicensesInnerWithDefaults() *ListProvisioningLicenses200ResponseAllOfLicensesInner`

NewListProvisioningLicenses200ResponseAllOfLicensesInnerWithDefaults instantiates a new ListProvisioningLicenses200ResponseAllOfLicensesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicenseType

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetLicenseType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetLicenseTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetLicenseType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetLicenseKey

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetOrgName

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetFullName

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetLicenseVersion

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetLicenseVersion() string`

GetLicenseVersion returns the LicenseVersion field if non-nil, zero value otherwise.

### GetLicenseVersionOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetLicenseVersionOk() (*string, bool)`

GetLicenseVersionOk returns a tuple with the LicenseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseVersion

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetLicenseVersion(v string)`

SetLicenseVersion sets LicenseVersion field to given value.

### HasLicenseVersion

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasLicenseVersion() bool`

HasLicenseVersion returns a boolean if a field has been set.

### GetCopies

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetReservationCount

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetReservationCount() int64`

GetReservationCount returns the ReservationCount field if non-nil, zero value otherwise.

### GetReservationCountOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetReservationCountOk() (*int64, bool)`

GetReservationCountOk returns a tuple with the ReservationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationCount

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetReservationCount(v int64)`

SetReservationCount sets ReservationCount field to given value.

### HasReservationCount

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasReservationCount() bool`

HasReservationCount returns a boolean if a field has been set.

### GetTenants

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetTenants() []map[string]interface{}`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetTenantsOk() (*[]map[string]interface{}, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetTenants(v []map[string]interface{})`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetVirtualImages

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetVirtualImages() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetVirtualImagesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetVirtualImages(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetAccount

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListProvisioningLicenses200ResponseAllOfLicensesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



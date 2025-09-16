# GetProvisioningLicense200ResponseLicense

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
**VirtualImages** | Pointer to [**[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 

## Methods

### NewGetProvisioningLicense200ResponseLicense

`func NewGetProvisioningLicense200ResponseLicense() *GetProvisioningLicense200ResponseLicense`

NewGetProvisioningLicense200ResponseLicense instantiates a new GetProvisioningLicense200ResponseLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProvisioningLicense200ResponseLicenseWithDefaults

`func NewGetProvisioningLicense200ResponseLicenseWithDefaults() *GetProvisioningLicense200ResponseLicense`

NewGetProvisioningLicense200ResponseLicenseWithDefaults instantiates a new GetProvisioningLicense200ResponseLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetProvisioningLicense200ResponseLicense) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetProvisioningLicense200ResponseLicense) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetProvisioningLicense200ResponseLicense) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetProvisioningLicense200ResponseLicense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetProvisioningLicense200ResponseLicense) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetProvisioningLicense200ResponseLicense) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetProvisioningLicense200ResponseLicense) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetProvisioningLicense200ResponseLicense) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetProvisioningLicense200ResponseLicense) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetProvisioningLicense200ResponseLicense) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetProvisioningLicense200ResponseLicense) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetProvisioningLicense200ResponseLicense) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicenseType

`func (o *GetProvisioningLicense200ResponseLicense) GetLicenseType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *GetProvisioningLicense200ResponseLicense) GetLicenseTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *GetProvisioningLicense200ResponseLicense) SetLicenseType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *GetProvisioningLicense200ResponseLicense) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetLicenseKey

`func (o *GetProvisioningLicense200ResponseLicense) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *GetProvisioningLicense200ResponseLicense) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *GetProvisioningLicense200ResponseLicense) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *GetProvisioningLicense200ResponseLicense) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetOrgName

`func (o *GetProvisioningLicense200ResponseLicense) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *GetProvisioningLicense200ResponseLicense) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *GetProvisioningLicense200ResponseLicense) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *GetProvisioningLicense200ResponseLicense) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetFullName

`func (o *GetProvisioningLicense200ResponseLicense) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GetProvisioningLicense200ResponseLicense) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GetProvisioningLicense200ResponseLicense) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GetProvisioningLicense200ResponseLicense) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetLicenseVersion

`func (o *GetProvisioningLicense200ResponseLicense) GetLicenseVersion() string`

GetLicenseVersion returns the LicenseVersion field if non-nil, zero value otherwise.

### GetLicenseVersionOk

`func (o *GetProvisioningLicense200ResponseLicense) GetLicenseVersionOk() (*string, bool)`

GetLicenseVersionOk returns a tuple with the LicenseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseVersion

`func (o *GetProvisioningLicense200ResponseLicense) SetLicenseVersion(v string)`

SetLicenseVersion sets LicenseVersion field to given value.

### HasLicenseVersion

`func (o *GetProvisioningLicense200ResponseLicense) HasLicenseVersion() bool`

HasLicenseVersion returns a boolean if a field has been set.

### GetCopies

`func (o *GetProvisioningLicense200ResponseLicense) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *GetProvisioningLicense200ResponseLicense) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *GetProvisioningLicense200ResponseLicense) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *GetProvisioningLicense200ResponseLicense) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetReservationCount

`func (o *GetProvisioningLicense200ResponseLicense) GetReservationCount() int64`

GetReservationCount returns the ReservationCount field if non-nil, zero value otherwise.

### GetReservationCountOk

`func (o *GetProvisioningLicense200ResponseLicense) GetReservationCountOk() (*int64, bool)`

GetReservationCountOk returns a tuple with the ReservationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationCount

`func (o *GetProvisioningLicense200ResponseLicense) SetReservationCount(v int64)`

SetReservationCount sets ReservationCount field to given value.

### HasReservationCount

`func (o *GetProvisioningLicense200ResponseLicense) HasReservationCount() bool`

HasReservationCount returns a boolean if a field has been set.

### GetTenants

`func (o *GetProvisioningLicense200ResponseLicense) GetTenants() []map[string]interface{}`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetProvisioningLicense200ResponseLicense) GetTenantsOk() (*[]map[string]interface{}, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetProvisioningLicense200ResponseLicense) SetTenants(v []map[string]interface{})`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetProvisioningLicense200ResponseLicense) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetVirtualImages

`func (o *GetProvisioningLicense200ResponseLicense) GetVirtualImages() []GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *GetProvisioningLicense200ResponseLicense) GetVirtualImagesOk() (*[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *GetProvisioningLicense200ResponseLicense) SetVirtualImages(v []GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *GetProvisioningLicense200ResponseLicense) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetAccount

`func (o *GetProvisioningLicense200ResponseLicense) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetProvisioningLicense200ResponseLicense) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetProvisioningLicense200ResponseLicense) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetProvisioningLicense200ResponseLicense) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



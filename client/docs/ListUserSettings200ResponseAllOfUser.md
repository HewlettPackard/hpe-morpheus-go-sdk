# ListUserSettings200ResponseAllOfUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**LinuxUsername** | Pointer to **string** |  | [optional] 
**LinuxPassword** | Pointer to **string** |  | [optional] 
**LinuxKeyPairId** | Pointer to **int64** |  | [optional] 
**WindowsUsername** | Pointer to **string** |  | [optional] 
**WindowsPassword** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**DesktopBackground** | Pointer to **string** |  | [optional] 
**ReceiveNotifications** | Pointer to **bool** |  | [optional] 
**DefaultGroup** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**DefaultCloud** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**DefaultPersona** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**IsUsing2FA** | Pointer to **bool** |  | [optional] 
**Tenant** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewListUserSettings200ResponseAllOfUser

`func NewListUserSettings200ResponseAllOfUser() *ListUserSettings200ResponseAllOfUser`

NewListUserSettings200ResponseAllOfUser instantiates a new ListUserSettings200ResponseAllOfUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserSettings200ResponseAllOfUserWithDefaults

`func NewListUserSettings200ResponseAllOfUserWithDefaults() *ListUserSettings200ResponseAllOfUser`

NewListUserSettings200ResponseAllOfUserWithDefaults instantiates a new ListUserSettings200ResponseAllOfUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListUserSettings200ResponseAllOfUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListUserSettings200ResponseAllOfUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListUserSettings200ResponseAllOfUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListUserSettings200ResponseAllOfUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *ListUserSettings200ResponseAllOfUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListUserSettings200ResponseAllOfUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListUserSettings200ResponseAllOfUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ListUserSettings200ResponseAllOfUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *ListUserSettings200ResponseAllOfUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ListUserSettings200ResponseAllOfUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ListUserSettings200ResponseAllOfUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ListUserSettings200ResponseAllOfUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ListUserSettings200ResponseAllOfUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ListUserSettings200ResponseAllOfUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ListUserSettings200ResponseAllOfUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ListUserSettings200ResponseAllOfUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *ListUserSettings200ResponseAllOfUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ListUserSettings200ResponseAllOfUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ListUserSettings200ResponseAllOfUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ListUserSettings200ResponseAllOfUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *ListUserSettings200ResponseAllOfUser) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *ListUserSettings200ResponseAllOfUser) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *ListUserSettings200ResponseAllOfUser) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *ListUserSettings200ResponseAllOfUser) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### GetLinuxPassword

`func (o *ListUserSettings200ResponseAllOfUser) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *ListUserSettings200ResponseAllOfUser) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *ListUserSettings200ResponseAllOfUser) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *ListUserSettings200ResponseAllOfUser) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### GetLinuxKeyPairId

`func (o *ListUserSettings200ResponseAllOfUser) GetLinuxKeyPairId() int64`

GetLinuxKeyPairId returns the LinuxKeyPairId field if non-nil, zero value otherwise.

### GetLinuxKeyPairIdOk

`func (o *ListUserSettings200ResponseAllOfUser) GetLinuxKeyPairIdOk() (*int64, bool)`

GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPairId

`func (o *ListUserSettings200ResponseAllOfUser) SetLinuxKeyPairId(v int64)`

SetLinuxKeyPairId sets LinuxKeyPairId field to given value.

### HasLinuxKeyPairId

`func (o *ListUserSettings200ResponseAllOfUser) HasLinuxKeyPairId() bool`

HasLinuxKeyPairId returns a boolean if a field has been set.

### GetWindowsUsername

`func (o *ListUserSettings200ResponseAllOfUser) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *ListUserSettings200ResponseAllOfUser) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *ListUserSettings200ResponseAllOfUser) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *ListUserSettings200ResponseAllOfUser) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### GetWindowsPassword

`func (o *ListUserSettings200ResponseAllOfUser) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *ListUserSettings200ResponseAllOfUser) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *ListUserSettings200ResponseAllOfUser) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *ListUserSettings200ResponseAllOfUser) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### GetAvatar

`func (o *ListUserSettings200ResponseAllOfUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ListUserSettings200ResponseAllOfUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ListUserSettings200ResponseAllOfUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ListUserSettings200ResponseAllOfUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetDesktopBackground

`func (o *ListUserSettings200ResponseAllOfUser) GetDesktopBackground() string`

GetDesktopBackground returns the DesktopBackground field if non-nil, zero value otherwise.

### GetDesktopBackgroundOk

`func (o *ListUserSettings200ResponseAllOfUser) GetDesktopBackgroundOk() (*string, bool)`

GetDesktopBackgroundOk returns a tuple with the DesktopBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopBackground

`func (o *ListUserSettings200ResponseAllOfUser) SetDesktopBackground(v string)`

SetDesktopBackground sets DesktopBackground field to given value.

### HasDesktopBackground

`func (o *ListUserSettings200ResponseAllOfUser) HasDesktopBackground() bool`

HasDesktopBackground returns a boolean if a field has been set.

### GetReceiveNotifications

`func (o *ListUserSettings200ResponseAllOfUser) GetReceiveNotifications() bool`

GetReceiveNotifications returns the ReceiveNotifications field if non-nil, zero value otherwise.

### GetReceiveNotificationsOk

`func (o *ListUserSettings200ResponseAllOfUser) GetReceiveNotificationsOk() (*bool, bool)`

GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveNotifications

`func (o *ListUserSettings200ResponseAllOfUser) SetReceiveNotifications(v bool)`

SetReceiveNotifications sets ReceiveNotifications field to given value.

### HasReceiveNotifications

`func (o *ListUserSettings200ResponseAllOfUser) HasReceiveNotifications() bool`

HasReceiveNotifications returns a boolean if a field has been set.

### GetDefaultGroup

`func (o *ListUserSettings200ResponseAllOfUser) GetDefaultGroup() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *ListUserSettings200ResponseAllOfUser) GetDefaultGroupOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *ListUserSettings200ResponseAllOfUser) SetDefaultGroup(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetDefaultGroup sets DefaultGroup field to given value.

### HasDefaultGroup

`func (o *ListUserSettings200ResponseAllOfUser) HasDefaultGroup() bool`

HasDefaultGroup returns a boolean if a field has been set.

### GetDefaultCloud

`func (o *ListUserSettings200ResponseAllOfUser) GetDefaultCloud() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetDefaultCloud returns the DefaultCloud field if non-nil, zero value otherwise.

### GetDefaultCloudOk

`func (o *ListUserSettings200ResponseAllOfUser) GetDefaultCloudOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetDefaultCloudOk returns a tuple with the DefaultCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCloud

`func (o *ListUserSettings200ResponseAllOfUser) SetDefaultCloud(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetDefaultCloud sets DefaultCloud field to given value.

### HasDefaultCloud

`func (o *ListUserSettings200ResponseAllOfUser) HasDefaultCloud() bool`

HasDefaultCloud returns a boolean if a field has been set.

### GetDefaultPersona

`func (o *ListUserSettings200ResponseAllOfUser) GetDefaultPersona() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *ListUserSettings200ResponseAllOfUser) GetDefaultPersonaOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *ListUserSettings200ResponseAllOfUser) SetDefaultPersona(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *ListUserSettings200ResponseAllOfUser) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### GetIsUsing2FA

`func (o *ListUserSettings200ResponseAllOfUser) GetIsUsing2FA() bool`

GetIsUsing2FA returns the IsUsing2FA field if non-nil, zero value otherwise.

### GetIsUsing2FAOk

`func (o *ListUserSettings200ResponseAllOfUser) GetIsUsing2FAOk() (*bool, bool)`

GetIsUsing2FAOk returns a tuple with the IsUsing2FA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsing2FA

`func (o *ListUserSettings200ResponseAllOfUser) SetIsUsing2FA(v bool)`

SetIsUsing2FA sets IsUsing2FA field to given value.

### HasIsUsing2FA

`func (o *ListUserSettings200ResponseAllOfUser) HasIsUsing2FA() bool`

HasIsUsing2FA returns a boolean if a field has been set.

### GetTenant

`func (o *ListUserSettings200ResponseAllOfUser) GetTenant() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ListUserSettings200ResponseAllOfUser) GetTenantOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ListUserSettings200ResponseAllOfUser) SetTenant(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ListUserSettings200ResponseAllOfUser) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



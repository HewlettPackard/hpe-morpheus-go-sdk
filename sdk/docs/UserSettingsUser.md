# UserSettingsUser

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
**LinuxKeyPairId** | Pointer to **NullableInt64** |  | [optional] 
**WindowsUsername** | Pointer to **string** |  | [optional] 
**WindowsPassword** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**DesktopBackground** | Pointer to **string** |  | [optional] 
**ReceiveNotifications** | Pointer to **bool** |  | [optional] 
**DefaultGroup** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**DefaultCloud** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**DefaultPersona** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**IsUsing2FA** | Pointer to **bool** |  | [optional] 
**Tenant** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 

## Methods

### NewUserSettingsUser

`func NewUserSettingsUser() *UserSettingsUser`

NewUserSettingsUser instantiates a new UserSettingsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsUserWithDefaults

`func NewUserSettingsUserWithDefaults() *UserSettingsUser`

NewUserSettingsUserWithDefaults instantiates a new UserSettingsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSettingsUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSettingsUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSettingsUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserSettingsUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *UserSettingsUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSettingsUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSettingsUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserSettingsUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *UserSettingsUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserSettingsUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserSettingsUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserSettingsUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserSettingsUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserSettingsUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserSettingsUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserSettingsUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *UserSettingsUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSettingsUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSettingsUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserSettingsUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *UserSettingsUser) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *UserSettingsUser) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *UserSettingsUser) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *UserSettingsUser) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### GetLinuxPassword

`func (o *UserSettingsUser) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *UserSettingsUser) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *UserSettingsUser) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *UserSettingsUser) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### GetLinuxKeyPairId

`func (o *UserSettingsUser) GetLinuxKeyPairId() int64`

GetLinuxKeyPairId returns the LinuxKeyPairId field if non-nil, zero value otherwise.

### GetLinuxKeyPairIdOk

`func (o *UserSettingsUser) GetLinuxKeyPairIdOk() (*int64, bool)`

GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPairId

`func (o *UserSettingsUser) SetLinuxKeyPairId(v int64)`

SetLinuxKeyPairId sets LinuxKeyPairId field to given value.

### HasLinuxKeyPairId

`func (o *UserSettingsUser) HasLinuxKeyPairId() bool`

HasLinuxKeyPairId returns a boolean if a field has been set.

### SetLinuxKeyPairIdNil

`func (o *UserSettingsUser) SetLinuxKeyPairIdNil(b bool)`

 SetLinuxKeyPairIdNil sets the value for LinuxKeyPairId to be an explicit nil

### UnsetLinuxKeyPairId
`func (o *UserSettingsUser) UnsetLinuxKeyPairId()`

UnsetLinuxKeyPairId ensures that no value is present for LinuxKeyPairId, not even an explicit nil
### GetWindowsUsername

`func (o *UserSettingsUser) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *UserSettingsUser) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *UserSettingsUser) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *UserSettingsUser) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### GetWindowsPassword

`func (o *UserSettingsUser) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *UserSettingsUser) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *UserSettingsUser) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *UserSettingsUser) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### GetAvatar

`func (o *UserSettingsUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserSettingsUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserSettingsUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserSettingsUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetDesktopBackground

`func (o *UserSettingsUser) GetDesktopBackground() string`

GetDesktopBackground returns the DesktopBackground field if non-nil, zero value otherwise.

### GetDesktopBackgroundOk

`func (o *UserSettingsUser) GetDesktopBackgroundOk() (*string, bool)`

GetDesktopBackgroundOk returns a tuple with the DesktopBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopBackground

`func (o *UserSettingsUser) SetDesktopBackground(v string)`

SetDesktopBackground sets DesktopBackground field to given value.

### HasDesktopBackground

`func (o *UserSettingsUser) HasDesktopBackground() bool`

HasDesktopBackground returns a boolean if a field has been set.

### GetReceiveNotifications

`func (o *UserSettingsUser) GetReceiveNotifications() bool`

GetReceiveNotifications returns the ReceiveNotifications field if non-nil, zero value otherwise.

### GetReceiveNotificationsOk

`func (o *UserSettingsUser) GetReceiveNotificationsOk() (*bool, bool)`

GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveNotifications

`func (o *UserSettingsUser) SetReceiveNotifications(v bool)`

SetReceiveNotifications sets ReceiveNotifications field to given value.

### HasReceiveNotifications

`func (o *UserSettingsUser) HasReceiveNotifications() bool`

HasReceiveNotifications returns a boolean if a field has been set.

### GetDefaultGroup

`func (o *UserSettingsUser) GetDefaultGroup() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *UserSettingsUser) GetDefaultGroupOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *UserSettingsUser) SetDefaultGroup(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetDefaultGroup sets DefaultGroup field to given value.

### HasDefaultGroup

`func (o *UserSettingsUser) HasDefaultGroup() bool`

HasDefaultGroup returns a boolean if a field has been set.

### GetDefaultCloud

`func (o *UserSettingsUser) GetDefaultCloud() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetDefaultCloud returns the DefaultCloud field if non-nil, zero value otherwise.

### GetDefaultCloudOk

`func (o *UserSettingsUser) GetDefaultCloudOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetDefaultCloudOk returns a tuple with the DefaultCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCloud

`func (o *UserSettingsUser) SetDefaultCloud(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetDefaultCloud sets DefaultCloud field to given value.

### HasDefaultCloud

`func (o *UserSettingsUser) HasDefaultCloud() bool`

HasDefaultCloud returns a boolean if a field has been set.

### GetDefaultPersona

`func (o *UserSettingsUser) GetDefaultPersona() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *UserSettingsUser) GetDefaultPersonaOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *UserSettingsUser) SetDefaultPersona(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *UserSettingsUser) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### GetIsUsing2FA

`func (o *UserSettingsUser) GetIsUsing2FA() bool`

GetIsUsing2FA returns the IsUsing2FA field if non-nil, zero value otherwise.

### GetIsUsing2FAOk

`func (o *UserSettingsUser) GetIsUsing2FAOk() (*bool, bool)`

GetIsUsing2FAOk returns a tuple with the IsUsing2FA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsing2FA

`func (o *UserSettingsUser) SetIsUsing2FA(v bool)`

SetIsUsing2FA sets IsUsing2FA field to given value.

### HasIsUsing2FA

`func (o *UserSettingsUser) HasIsUsing2FA() bool`

HasIsUsing2FA returns a boolean if a field has been set.

### GetTenant

`func (o *UserSettingsUser) GetTenant() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *UserSettingsUser) GetTenantOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *UserSettingsUser) SetTenant(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *UserSettingsUser) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



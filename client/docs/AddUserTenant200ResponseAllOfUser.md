# AddUserTenant200ResponseAllOfUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ReceiveNotifications** | Pointer to **bool** |  | [optional] 
**IsUsing2FA** | Pointer to **bool** |  | [optional] 
**AccountExpired** | Pointer to **bool** |  | [optional] 
**AccountLocked** | Pointer to **bool** |  | [optional] 
**PasswordExpired** | Pointer to **bool** |  | [optional] 
**LoginCount** | Pointer to **int64** |  | [optional] 
**LoginAttempts** | Pointer to **int64** |  | [optional] 
**LastLoginDate** | Pointer to **time.Time** |  | [optional] 
**Roles** | Pointer to [**[]AddUserTenant200ResponseAllOfUserRolesInner**](AddUserTenant200ResponseAllOfUserRolesInner.md) |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**LinuxUsername** | Pointer to **string** |  | [optional] 
**LinuxPassword** | Pointer to **string** |  | [optional] 
**LinuxKeyPairId** | Pointer to **int64** |  | [optional] 
**WindowsUsername** | Pointer to **string** |  | [optional] 
**WindowsPassword** | Pointer to **string** |  | [optional] 
**DefaultPersona** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Access** | Pointer to [**AddUserTenant200ResponseAllOfUserAccess**](AddUserTenant200ResponseAllOfUserAccess.md) |  | [optional] 

## Methods

### NewAddUserTenant200ResponseAllOfUser

`func NewAddUserTenant200ResponseAllOfUser() *AddUserTenant200ResponseAllOfUser`

NewAddUserTenant200ResponseAllOfUser instantiates a new AddUserTenant200ResponseAllOfUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserTenant200ResponseAllOfUserWithDefaults

`func NewAddUserTenant200ResponseAllOfUserWithDefaults() *AddUserTenant200ResponseAllOfUser`

NewAddUserTenant200ResponseAllOfUserWithDefaults instantiates a new AddUserTenant200ResponseAllOfUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddUserTenant200ResponseAllOfUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddUserTenant200ResponseAllOfUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddUserTenant200ResponseAllOfUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddUserTenant200ResponseAllOfUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *AddUserTenant200ResponseAllOfUser) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddUserTenant200ResponseAllOfUser) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddUserTenant200ResponseAllOfUser) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddUserTenant200ResponseAllOfUser) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUsername

`func (o *AddUserTenant200ResponseAllOfUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddUserTenant200ResponseAllOfUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddUserTenant200ResponseAllOfUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddUserTenant200ResponseAllOfUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *AddUserTenant200ResponseAllOfUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddUserTenant200ResponseAllOfUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddUserTenant200ResponseAllOfUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AddUserTenant200ResponseAllOfUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *AddUserTenant200ResponseAllOfUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddUserTenant200ResponseAllOfUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddUserTenant200ResponseAllOfUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddUserTenant200ResponseAllOfUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *AddUserTenant200ResponseAllOfUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AddUserTenant200ResponseAllOfUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AddUserTenant200ResponseAllOfUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AddUserTenant200ResponseAllOfUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AddUserTenant200ResponseAllOfUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AddUserTenant200ResponseAllOfUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AddUserTenant200ResponseAllOfUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AddUserTenant200ResponseAllOfUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEnabled

`func (o *AddUserTenant200ResponseAllOfUser) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddUserTenant200ResponseAllOfUser) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddUserTenant200ResponseAllOfUser) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddUserTenant200ResponseAllOfUser) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetReceiveNotifications

`func (o *AddUserTenant200ResponseAllOfUser) GetReceiveNotifications() bool`

GetReceiveNotifications returns the ReceiveNotifications field if non-nil, zero value otherwise.

### GetReceiveNotificationsOk

`func (o *AddUserTenant200ResponseAllOfUser) GetReceiveNotificationsOk() (*bool, bool)`

GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveNotifications

`func (o *AddUserTenant200ResponseAllOfUser) SetReceiveNotifications(v bool)`

SetReceiveNotifications sets ReceiveNotifications field to given value.

### HasReceiveNotifications

`func (o *AddUserTenant200ResponseAllOfUser) HasReceiveNotifications() bool`

HasReceiveNotifications returns a boolean if a field has been set.

### GetIsUsing2FA

`func (o *AddUserTenant200ResponseAllOfUser) GetIsUsing2FA() bool`

GetIsUsing2FA returns the IsUsing2FA field if non-nil, zero value otherwise.

### GetIsUsing2FAOk

`func (o *AddUserTenant200ResponseAllOfUser) GetIsUsing2FAOk() (*bool, bool)`

GetIsUsing2FAOk returns a tuple with the IsUsing2FA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsing2FA

`func (o *AddUserTenant200ResponseAllOfUser) SetIsUsing2FA(v bool)`

SetIsUsing2FA sets IsUsing2FA field to given value.

### HasIsUsing2FA

`func (o *AddUserTenant200ResponseAllOfUser) HasIsUsing2FA() bool`

HasIsUsing2FA returns a boolean if a field has been set.

### GetAccountExpired

`func (o *AddUserTenant200ResponseAllOfUser) GetAccountExpired() bool`

GetAccountExpired returns the AccountExpired field if non-nil, zero value otherwise.

### GetAccountExpiredOk

`func (o *AddUserTenant200ResponseAllOfUser) GetAccountExpiredOk() (*bool, bool)`

GetAccountExpiredOk returns a tuple with the AccountExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpired

`func (o *AddUserTenant200ResponseAllOfUser) SetAccountExpired(v bool)`

SetAccountExpired sets AccountExpired field to given value.

### HasAccountExpired

`func (o *AddUserTenant200ResponseAllOfUser) HasAccountExpired() bool`

HasAccountExpired returns a boolean if a field has been set.

### GetAccountLocked

`func (o *AddUserTenant200ResponseAllOfUser) GetAccountLocked() bool`

GetAccountLocked returns the AccountLocked field if non-nil, zero value otherwise.

### GetAccountLockedOk

`func (o *AddUserTenant200ResponseAllOfUser) GetAccountLockedOk() (*bool, bool)`

GetAccountLockedOk returns a tuple with the AccountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLocked

`func (o *AddUserTenant200ResponseAllOfUser) SetAccountLocked(v bool)`

SetAccountLocked sets AccountLocked field to given value.

### HasAccountLocked

`func (o *AddUserTenant200ResponseAllOfUser) HasAccountLocked() bool`

HasAccountLocked returns a boolean if a field has been set.

### GetPasswordExpired

`func (o *AddUserTenant200ResponseAllOfUser) GetPasswordExpired() bool`

GetPasswordExpired returns the PasswordExpired field if non-nil, zero value otherwise.

### GetPasswordExpiredOk

`func (o *AddUserTenant200ResponseAllOfUser) GetPasswordExpiredOk() (*bool, bool)`

GetPasswordExpiredOk returns a tuple with the PasswordExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpired

`func (o *AddUserTenant200ResponseAllOfUser) SetPasswordExpired(v bool)`

SetPasswordExpired sets PasswordExpired field to given value.

### HasPasswordExpired

`func (o *AddUserTenant200ResponseAllOfUser) HasPasswordExpired() bool`

HasPasswordExpired returns a boolean if a field has been set.

### GetLoginCount

`func (o *AddUserTenant200ResponseAllOfUser) GetLoginCount() int64`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *AddUserTenant200ResponseAllOfUser) GetLoginCountOk() (*int64, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *AddUserTenant200ResponseAllOfUser) SetLoginCount(v int64)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *AddUserTenant200ResponseAllOfUser) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetLoginAttempts

`func (o *AddUserTenant200ResponseAllOfUser) GetLoginAttempts() int64`

GetLoginAttempts returns the LoginAttempts field if non-nil, zero value otherwise.

### GetLoginAttemptsOk

`func (o *AddUserTenant200ResponseAllOfUser) GetLoginAttemptsOk() (*int64, bool)`

GetLoginAttemptsOk returns a tuple with the LoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttempts

`func (o *AddUserTenant200ResponseAllOfUser) SetLoginAttempts(v int64)`

SetLoginAttempts sets LoginAttempts field to given value.

### HasLoginAttempts

`func (o *AddUserTenant200ResponseAllOfUser) HasLoginAttempts() bool`

HasLoginAttempts returns a boolean if a field has been set.

### GetLastLoginDate

`func (o *AddUserTenant200ResponseAllOfUser) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *AddUserTenant200ResponseAllOfUser) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *AddUserTenant200ResponseAllOfUser) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *AddUserTenant200ResponseAllOfUser) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### GetRoles

`func (o *AddUserTenant200ResponseAllOfUser) GetRoles() []AddUserTenant200ResponseAllOfUserRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AddUserTenant200ResponseAllOfUser) GetRolesOk() (*[]AddUserTenant200ResponseAllOfUserRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AddUserTenant200ResponseAllOfUser) SetRoles(v []AddUserTenant200ResponseAllOfUserRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AddUserTenant200ResponseAllOfUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAccount

`func (o *AddUserTenant200ResponseAllOfUser) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddUserTenant200ResponseAllOfUser) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddUserTenant200ResponseAllOfUser) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddUserTenant200ResponseAllOfUser) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *AddUserTenant200ResponseAllOfUser) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *AddUserTenant200ResponseAllOfUser) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *AddUserTenant200ResponseAllOfUser) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *AddUserTenant200ResponseAllOfUser) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### GetLinuxPassword

`func (o *AddUserTenant200ResponseAllOfUser) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *AddUserTenant200ResponseAllOfUser) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *AddUserTenant200ResponseAllOfUser) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *AddUserTenant200ResponseAllOfUser) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### GetLinuxKeyPairId

`func (o *AddUserTenant200ResponseAllOfUser) GetLinuxKeyPairId() int64`

GetLinuxKeyPairId returns the LinuxKeyPairId field if non-nil, zero value otherwise.

### GetLinuxKeyPairIdOk

`func (o *AddUserTenant200ResponseAllOfUser) GetLinuxKeyPairIdOk() (*int64, bool)`

GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPairId

`func (o *AddUserTenant200ResponseAllOfUser) SetLinuxKeyPairId(v int64)`

SetLinuxKeyPairId sets LinuxKeyPairId field to given value.

### HasLinuxKeyPairId

`func (o *AddUserTenant200ResponseAllOfUser) HasLinuxKeyPairId() bool`

HasLinuxKeyPairId returns a boolean if a field has been set.

### GetWindowsUsername

`func (o *AddUserTenant200ResponseAllOfUser) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *AddUserTenant200ResponseAllOfUser) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *AddUserTenant200ResponseAllOfUser) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *AddUserTenant200ResponseAllOfUser) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### GetWindowsPassword

`func (o *AddUserTenant200ResponseAllOfUser) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *AddUserTenant200ResponseAllOfUser) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *AddUserTenant200ResponseAllOfUser) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *AddUserTenant200ResponseAllOfUser) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### GetDefaultPersona

`func (o *AddUserTenant200ResponseAllOfUser) GetDefaultPersona() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *AddUserTenant200ResponseAllOfUser) GetDefaultPersonaOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *AddUserTenant200ResponseAllOfUser) SetDefaultPersona(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *AddUserTenant200ResponseAllOfUser) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddUserTenant200ResponseAllOfUser) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddUserTenant200ResponseAllOfUser) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddUserTenant200ResponseAllOfUser) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddUserTenant200ResponseAllOfUser) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddUserTenant200ResponseAllOfUser) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddUserTenant200ResponseAllOfUser) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddUserTenant200ResponseAllOfUser) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddUserTenant200ResponseAllOfUser) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccess

`func (o *AddUserTenant200ResponseAllOfUser) GetAccess() AddUserTenant200ResponseAllOfUserAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AddUserTenant200ResponseAllOfUser) GetAccessOk() (*AddUserTenant200ResponseAllOfUserAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AddUserTenant200ResponseAllOfUser) SetAccess(v AddUserTenant200ResponseAllOfUserAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AddUserTenant200ResponseAllOfUser) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



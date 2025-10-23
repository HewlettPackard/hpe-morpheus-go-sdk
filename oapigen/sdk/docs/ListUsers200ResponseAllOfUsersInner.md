# ListUsers200ResponseAllOfUsersInner

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
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**LinuxUsername** | Pointer to **NullableString** |  | [optional] 
**LinuxPassword** | Pointer to **NullableString** |  | [optional] 
**LinuxKeyPairId** | Pointer to **NullableInt64** |  | [optional] 
**WindowsUsername** | Pointer to **NullableString** |  | [optional] 
**WindowsPassword** | Pointer to **NullableString** |  | [optional] 
**DefaultPersona** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Access** | Pointer to [**AddUserTenant200ResponseAllOfUserAccess**](AddUserTenant200ResponseAllOfUserAccess.md) |  | [optional] 

## Methods

### NewListUsers200ResponseAllOfUsersInner

`func NewListUsers200ResponseAllOfUsersInner() *ListUsers200ResponseAllOfUsersInner`

NewListUsers200ResponseAllOfUsersInner instantiates a new ListUsers200ResponseAllOfUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsers200ResponseAllOfUsersInnerWithDefaults

`func NewListUsers200ResponseAllOfUsersInnerWithDefaults() *ListUsers200ResponseAllOfUsersInner`

NewListUsers200ResponseAllOfUsersInnerWithDefaults instantiates a new ListUsers200ResponseAllOfUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListUsers200ResponseAllOfUsersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListUsers200ResponseAllOfUsersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListUsers200ResponseAllOfUsersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *ListUsers200ResponseAllOfUsersInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListUsers200ResponseAllOfUsersInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListUsers200ResponseAllOfUsersInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUsername

`func (o *ListUsers200ResponseAllOfUsersInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListUsers200ResponseAllOfUsersInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ListUsers200ResponseAllOfUsersInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListUsers200ResponseAllOfUsersInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListUsers200ResponseAllOfUsersInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListUsers200ResponseAllOfUsersInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *ListUsers200ResponseAllOfUsersInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ListUsers200ResponseAllOfUsersInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ListUsers200ResponseAllOfUsersInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *ListUsers200ResponseAllOfUsersInner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ListUsers200ResponseAllOfUsersInner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ListUsers200ResponseAllOfUsersInner) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ListUsers200ResponseAllOfUsersInner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ListUsers200ResponseAllOfUsersInner) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ListUsers200ResponseAllOfUsersInner) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEnabled

`func (o *ListUsers200ResponseAllOfUsersInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListUsers200ResponseAllOfUsersInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListUsers200ResponseAllOfUsersInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetReceiveNotifications

`func (o *ListUsers200ResponseAllOfUsersInner) GetReceiveNotifications() bool`

GetReceiveNotifications returns the ReceiveNotifications field if non-nil, zero value otherwise.

### GetReceiveNotificationsOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetReceiveNotificationsOk() (*bool, bool)`

GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveNotifications

`func (o *ListUsers200ResponseAllOfUsersInner) SetReceiveNotifications(v bool)`

SetReceiveNotifications sets ReceiveNotifications field to given value.

### HasReceiveNotifications

`func (o *ListUsers200ResponseAllOfUsersInner) HasReceiveNotifications() bool`

HasReceiveNotifications returns a boolean if a field has been set.

### GetIsUsing2FA

`func (o *ListUsers200ResponseAllOfUsersInner) GetIsUsing2FA() bool`

GetIsUsing2FA returns the IsUsing2FA field if non-nil, zero value otherwise.

### GetIsUsing2FAOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetIsUsing2FAOk() (*bool, bool)`

GetIsUsing2FAOk returns a tuple with the IsUsing2FA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsing2FA

`func (o *ListUsers200ResponseAllOfUsersInner) SetIsUsing2FA(v bool)`

SetIsUsing2FA sets IsUsing2FA field to given value.

### HasIsUsing2FA

`func (o *ListUsers200ResponseAllOfUsersInner) HasIsUsing2FA() bool`

HasIsUsing2FA returns a boolean if a field has been set.

### GetAccountExpired

`func (o *ListUsers200ResponseAllOfUsersInner) GetAccountExpired() bool`

GetAccountExpired returns the AccountExpired field if non-nil, zero value otherwise.

### GetAccountExpiredOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetAccountExpiredOk() (*bool, bool)`

GetAccountExpiredOk returns a tuple with the AccountExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpired

`func (o *ListUsers200ResponseAllOfUsersInner) SetAccountExpired(v bool)`

SetAccountExpired sets AccountExpired field to given value.

### HasAccountExpired

`func (o *ListUsers200ResponseAllOfUsersInner) HasAccountExpired() bool`

HasAccountExpired returns a boolean if a field has been set.

### GetAccountLocked

`func (o *ListUsers200ResponseAllOfUsersInner) GetAccountLocked() bool`

GetAccountLocked returns the AccountLocked field if non-nil, zero value otherwise.

### GetAccountLockedOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetAccountLockedOk() (*bool, bool)`

GetAccountLockedOk returns a tuple with the AccountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLocked

`func (o *ListUsers200ResponseAllOfUsersInner) SetAccountLocked(v bool)`

SetAccountLocked sets AccountLocked field to given value.

### HasAccountLocked

`func (o *ListUsers200ResponseAllOfUsersInner) HasAccountLocked() bool`

HasAccountLocked returns a boolean if a field has been set.

### GetPasswordExpired

`func (o *ListUsers200ResponseAllOfUsersInner) GetPasswordExpired() bool`

GetPasswordExpired returns the PasswordExpired field if non-nil, zero value otherwise.

### GetPasswordExpiredOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetPasswordExpiredOk() (*bool, bool)`

GetPasswordExpiredOk returns a tuple with the PasswordExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpired

`func (o *ListUsers200ResponseAllOfUsersInner) SetPasswordExpired(v bool)`

SetPasswordExpired sets PasswordExpired field to given value.

### HasPasswordExpired

`func (o *ListUsers200ResponseAllOfUsersInner) HasPasswordExpired() bool`

HasPasswordExpired returns a boolean if a field has been set.

### GetLoginCount

`func (o *ListUsers200ResponseAllOfUsersInner) GetLoginCount() int64`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetLoginCountOk() (*int64, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *ListUsers200ResponseAllOfUsersInner) SetLoginCount(v int64)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *ListUsers200ResponseAllOfUsersInner) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetLoginAttempts

`func (o *ListUsers200ResponseAllOfUsersInner) GetLoginAttempts() int64`

GetLoginAttempts returns the LoginAttempts field if non-nil, zero value otherwise.

### GetLoginAttemptsOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetLoginAttemptsOk() (*int64, bool)`

GetLoginAttemptsOk returns a tuple with the LoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttempts

`func (o *ListUsers200ResponseAllOfUsersInner) SetLoginAttempts(v int64)`

SetLoginAttempts sets LoginAttempts field to given value.

### HasLoginAttempts

`func (o *ListUsers200ResponseAllOfUsersInner) HasLoginAttempts() bool`

HasLoginAttempts returns a boolean if a field has been set.

### GetLastLoginDate

`func (o *ListUsers200ResponseAllOfUsersInner) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *ListUsers200ResponseAllOfUsersInner) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *ListUsers200ResponseAllOfUsersInner) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### GetRoles

`func (o *ListUsers200ResponseAllOfUsersInner) GetRoles() []AddUserTenant200ResponseAllOfUserRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetRolesOk() (*[]AddUserTenant200ResponseAllOfUserRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ListUsers200ResponseAllOfUsersInner) SetRoles(v []AddUserTenant200ResponseAllOfUserRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ListUsers200ResponseAllOfUsersInner) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAccount

`func (o *ListUsers200ResponseAllOfUsersInner) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListUsers200ResponseAllOfUsersInner) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListUsers200ResponseAllOfUsersInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *ListUsers200ResponseAllOfUsersInner) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *ListUsers200ResponseAllOfUsersInner) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *ListUsers200ResponseAllOfUsersInner) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### SetLinuxUsernameNil

`func (o *ListUsers200ResponseAllOfUsersInner) SetLinuxUsernameNil(b bool)`

 SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil

### UnsetLinuxUsername
`func (o *ListUsers200ResponseAllOfUsersInner) UnsetLinuxUsername()`

UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil
### GetLinuxPassword

`func (o *ListUsers200ResponseAllOfUsersInner) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *ListUsers200ResponseAllOfUsersInner) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *ListUsers200ResponseAllOfUsersInner) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### SetLinuxPasswordNil

`func (o *ListUsers200ResponseAllOfUsersInner) SetLinuxPasswordNil(b bool)`

 SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil

### UnsetLinuxPassword
`func (o *ListUsers200ResponseAllOfUsersInner) UnsetLinuxPassword()`

UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
### GetLinuxKeyPairId

`func (o *ListUsers200ResponseAllOfUsersInner) GetLinuxKeyPairId() int64`

GetLinuxKeyPairId returns the LinuxKeyPairId field if non-nil, zero value otherwise.

### GetLinuxKeyPairIdOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetLinuxKeyPairIdOk() (*int64, bool)`

GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPairId

`func (o *ListUsers200ResponseAllOfUsersInner) SetLinuxKeyPairId(v int64)`

SetLinuxKeyPairId sets LinuxKeyPairId field to given value.

### HasLinuxKeyPairId

`func (o *ListUsers200ResponseAllOfUsersInner) HasLinuxKeyPairId() bool`

HasLinuxKeyPairId returns a boolean if a field has been set.

### SetLinuxKeyPairIdNil

`func (o *ListUsers200ResponseAllOfUsersInner) SetLinuxKeyPairIdNil(b bool)`

 SetLinuxKeyPairIdNil sets the value for LinuxKeyPairId to be an explicit nil

### UnsetLinuxKeyPairId
`func (o *ListUsers200ResponseAllOfUsersInner) UnsetLinuxKeyPairId()`

UnsetLinuxKeyPairId ensures that no value is present for LinuxKeyPairId, not even an explicit nil
### GetWindowsUsername

`func (o *ListUsers200ResponseAllOfUsersInner) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *ListUsers200ResponseAllOfUsersInner) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *ListUsers200ResponseAllOfUsersInner) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### SetWindowsUsernameNil

`func (o *ListUsers200ResponseAllOfUsersInner) SetWindowsUsernameNil(b bool)`

 SetWindowsUsernameNil sets the value for WindowsUsername to be an explicit nil

### UnsetWindowsUsername
`func (o *ListUsers200ResponseAllOfUsersInner) UnsetWindowsUsername()`

UnsetWindowsUsername ensures that no value is present for WindowsUsername, not even an explicit nil
### GetWindowsPassword

`func (o *ListUsers200ResponseAllOfUsersInner) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *ListUsers200ResponseAllOfUsersInner) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *ListUsers200ResponseAllOfUsersInner) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### SetWindowsPasswordNil

`func (o *ListUsers200ResponseAllOfUsersInner) SetWindowsPasswordNil(b bool)`

 SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil

### UnsetWindowsPassword
`func (o *ListUsers200ResponseAllOfUsersInner) UnsetWindowsPassword()`

UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
### GetDefaultPersona

`func (o *ListUsers200ResponseAllOfUsersInner) GetDefaultPersona() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetDefaultPersonaOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *ListUsers200ResponseAllOfUsersInner) SetDefaultPersona(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *ListUsers200ResponseAllOfUsersInner) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListUsers200ResponseAllOfUsersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListUsers200ResponseAllOfUsersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListUsers200ResponseAllOfUsersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListUsers200ResponseAllOfUsersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListUsers200ResponseAllOfUsersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListUsers200ResponseAllOfUsersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccess

`func (o *ListUsers200ResponseAllOfUsersInner) GetAccess() AddUserTenant200ResponseAllOfUserAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ListUsers200ResponseAllOfUsersInner) GetAccessOk() (*AddUserTenant200ResponseAllOfUserAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ListUsers200ResponseAllOfUsersInner) SetAccess(v AddUserTenant200ResponseAllOfUserAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ListUsers200ResponseAllOfUsersInner) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



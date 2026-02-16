# AddUser200ResponseAllOfUser

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
**Roles** | Pointer to [**[]AddUser200ResponseAllOfUserRolesInner**](AddUser200ResponseAllOfUserRolesInner.md) |  | [optional] 
**Account** | Pointer to [**AddUser200ResponseAllOfUserAccount**](AddUser200ResponseAllOfUserAccount.md) |  | [optional] 
**LinuxUsername** | Pointer to **NullableString** |  | [optional] 
**LinuxPassword** | Pointer to **NullableString** |  | [optional] 
**LinuxKeyPairId** | Pointer to **NullableInt64** |  | [optional] 
**WindowsUsername** | Pointer to **NullableString** |  | [optional] 
**WindowsPassword** | Pointer to **NullableString** |  | [optional] 
**DefaultPersona** | Pointer to [**AddUser200ResponseAllOfUserDefaultPersona**](AddUser200ResponseAllOfUserDefaultPersona.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Access** | Pointer to [**AddUser200ResponseAllOfUserAccess**](AddUser200ResponseAllOfUserAccess.md) |  | [optional] 

## Methods

### NewAddUser200ResponseAllOfUser

`func NewAddUser200ResponseAllOfUser() *AddUser200ResponseAllOfUser`

NewAddUser200ResponseAllOfUser instantiates a new AddUser200ResponseAllOfUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUser200ResponseAllOfUserWithDefaults

`func NewAddUser200ResponseAllOfUserWithDefaults() *AddUser200ResponseAllOfUser`

NewAddUser200ResponseAllOfUserWithDefaults instantiates a new AddUser200ResponseAllOfUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddUser200ResponseAllOfUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddUser200ResponseAllOfUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddUser200ResponseAllOfUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddUser200ResponseAllOfUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *AddUser200ResponseAllOfUser) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddUser200ResponseAllOfUser) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddUser200ResponseAllOfUser) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddUser200ResponseAllOfUser) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUsername

`func (o *AddUser200ResponseAllOfUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddUser200ResponseAllOfUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddUser200ResponseAllOfUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddUser200ResponseAllOfUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *AddUser200ResponseAllOfUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddUser200ResponseAllOfUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddUser200ResponseAllOfUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AddUser200ResponseAllOfUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *AddUser200ResponseAllOfUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddUser200ResponseAllOfUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddUser200ResponseAllOfUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddUser200ResponseAllOfUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *AddUser200ResponseAllOfUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AddUser200ResponseAllOfUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AddUser200ResponseAllOfUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AddUser200ResponseAllOfUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AddUser200ResponseAllOfUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AddUser200ResponseAllOfUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AddUser200ResponseAllOfUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AddUser200ResponseAllOfUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEnabled

`func (o *AddUser200ResponseAllOfUser) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddUser200ResponseAllOfUser) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddUser200ResponseAllOfUser) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddUser200ResponseAllOfUser) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetReceiveNotifications

`func (o *AddUser200ResponseAllOfUser) GetReceiveNotifications() bool`

GetReceiveNotifications returns the ReceiveNotifications field if non-nil, zero value otherwise.

### GetReceiveNotificationsOk

`func (o *AddUser200ResponseAllOfUser) GetReceiveNotificationsOk() (*bool, bool)`

GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveNotifications

`func (o *AddUser200ResponseAllOfUser) SetReceiveNotifications(v bool)`

SetReceiveNotifications sets ReceiveNotifications field to given value.

### HasReceiveNotifications

`func (o *AddUser200ResponseAllOfUser) HasReceiveNotifications() bool`

HasReceiveNotifications returns a boolean if a field has been set.

### GetIsUsing2FA

`func (o *AddUser200ResponseAllOfUser) GetIsUsing2FA() bool`

GetIsUsing2FA returns the IsUsing2FA field if non-nil, zero value otherwise.

### GetIsUsing2FAOk

`func (o *AddUser200ResponseAllOfUser) GetIsUsing2FAOk() (*bool, bool)`

GetIsUsing2FAOk returns a tuple with the IsUsing2FA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsing2FA

`func (o *AddUser200ResponseAllOfUser) SetIsUsing2FA(v bool)`

SetIsUsing2FA sets IsUsing2FA field to given value.

### HasIsUsing2FA

`func (o *AddUser200ResponseAllOfUser) HasIsUsing2FA() bool`

HasIsUsing2FA returns a boolean if a field has been set.

### GetAccountExpired

`func (o *AddUser200ResponseAllOfUser) GetAccountExpired() bool`

GetAccountExpired returns the AccountExpired field if non-nil, zero value otherwise.

### GetAccountExpiredOk

`func (o *AddUser200ResponseAllOfUser) GetAccountExpiredOk() (*bool, bool)`

GetAccountExpiredOk returns a tuple with the AccountExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpired

`func (o *AddUser200ResponseAllOfUser) SetAccountExpired(v bool)`

SetAccountExpired sets AccountExpired field to given value.

### HasAccountExpired

`func (o *AddUser200ResponseAllOfUser) HasAccountExpired() bool`

HasAccountExpired returns a boolean if a field has been set.

### GetAccountLocked

`func (o *AddUser200ResponseAllOfUser) GetAccountLocked() bool`

GetAccountLocked returns the AccountLocked field if non-nil, zero value otherwise.

### GetAccountLockedOk

`func (o *AddUser200ResponseAllOfUser) GetAccountLockedOk() (*bool, bool)`

GetAccountLockedOk returns a tuple with the AccountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLocked

`func (o *AddUser200ResponseAllOfUser) SetAccountLocked(v bool)`

SetAccountLocked sets AccountLocked field to given value.

### HasAccountLocked

`func (o *AddUser200ResponseAllOfUser) HasAccountLocked() bool`

HasAccountLocked returns a boolean if a field has been set.

### GetPasswordExpired

`func (o *AddUser200ResponseAllOfUser) GetPasswordExpired() bool`

GetPasswordExpired returns the PasswordExpired field if non-nil, zero value otherwise.

### GetPasswordExpiredOk

`func (o *AddUser200ResponseAllOfUser) GetPasswordExpiredOk() (*bool, bool)`

GetPasswordExpiredOk returns a tuple with the PasswordExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpired

`func (o *AddUser200ResponseAllOfUser) SetPasswordExpired(v bool)`

SetPasswordExpired sets PasswordExpired field to given value.

### HasPasswordExpired

`func (o *AddUser200ResponseAllOfUser) HasPasswordExpired() bool`

HasPasswordExpired returns a boolean if a field has been set.

### GetLoginCount

`func (o *AddUser200ResponseAllOfUser) GetLoginCount() int64`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *AddUser200ResponseAllOfUser) GetLoginCountOk() (*int64, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *AddUser200ResponseAllOfUser) SetLoginCount(v int64)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *AddUser200ResponseAllOfUser) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetLoginAttempts

`func (o *AddUser200ResponseAllOfUser) GetLoginAttempts() int64`

GetLoginAttempts returns the LoginAttempts field if non-nil, zero value otherwise.

### GetLoginAttemptsOk

`func (o *AddUser200ResponseAllOfUser) GetLoginAttemptsOk() (*int64, bool)`

GetLoginAttemptsOk returns a tuple with the LoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttempts

`func (o *AddUser200ResponseAllOfUser) SetLoginAttempts(v int64)`

SetLoginAttempts sets LoginAttempts field to given value.

### HasLoginAttempts

`func (o *AddUser200ResponseAllOfUser) HasLoginAttempts() bool`

HasLoginAttempts returns a boolean if a field has been set.

### GetLastLoginDate

`func (o *AddUser200ResponseAllOfUser) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *AddUser200ResponseAllOfUser) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *AddUser200ResponseAllOfUser) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *AddUser200ResponseAllOfUser) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### GetRoles

`func (o *AddUser200ResponseAllOfUser) GetRoles() []AddUser200ResponseAllOfUserRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AddUser200ResponseAllOfUser) GetRolesOk() (*[]AddUser200ResponseAllOfUserRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AddUser200ResponseAllOfUser) SetRoles(v []AddUser200ResponseAllOfUserRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AddUser200ResponseAllOfUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAccount

`func (o *AddUser200ResponseAllOfUser) GetAccount() AddUser200ResponseAllOfUserAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddUser200ResponseAllOfUser) GetAccountOk() (*AddUser200ResponseAllOfUserAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddUser200ResponseAllOfUser) SetAccount(v AddUser200ResponseAllOfUserAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddUser200ResponseAllOfUser) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *AddUser200ResponseAllOfUser) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *AddUser200ResponseAllOfUser) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *AddUser200ResponseAllOfUser) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *AddUser200ResponseAllOfUser) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### SetLinuxUsernameNil

`func (o *AddUser200ResponseAllOfUser) SetLinuxUsernameNil(b bool)`

 SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil

### UnsetLinuxUsername
`func (o *AddUser200ResponseAllOfUser) UnsetLinuxUsername()`

UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil
### GetLinuxPassword

`func (o *AddUser200ResponseAllOfUser) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *AddUser200ResponseAllOfUser) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *AddUser200ResponseAllOfUser) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *AddUser200ResponseAllOfUser) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### SetLinuxPasswordNil

`func (o *AddUser200ResponseAllOfUser) SetLinuxPasswordNil(b bool)`

 SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil

### UnsetLinuxPassword
`func (o *AddUser200ResponseAllOfUser) UnsetLinuxPassword()`

UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
### GetLinuxKeyPairId

`func (o *AddUser200ResponseAllOfUser) GetLinuxKeyPairId() int64`

GetLinuxKeyPairId returns the LinuxKeyPairId field if non-nil, zero value otherwise.

### GetLinuxKeyPairIdOk

`func (o *AddUser200ResponseAllOfUser) GetLinuxKeyPairIdOk() (*int64, bool)`

GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPairId

`func (o *AddUser200ResponseAllOfUser) SetLinuxKeyPairId(v int64)`

SetLinuxKeyPairId sets LinuxKeyPairId field to given value.

### HasLinuxKeyPairId

`func (o *AddUser200ResponseAllOfUser) HasLinuxKeyPairId() bool`

HasLinuxKeyPairId returns a boolean if a field has been set.

### SetLinuxKeyPairIdNil

`func (o *AddUser200ResponseAllOfUser) SetLinuxKeyPairIdNil(b bool)`

 SetLinuxKeyPairIdNil sets the value for LinuxKeyPairId to be an explicit nil

### UnsetLinuxKeyPairId
`func (o *AddUser200ResponseAllOfUser) UnsetLinuxKeyPairId()`

UnsetLinuxKeyPairId ensures that no value is present for LinuxKeyPairId, not even an explicit nil
### GetWindowsUsername

`func (o *AddUser200ResponseAllOfUser) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *AddUser200ResponseAllOfUser) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *AddUser200ResponseAllOfUser) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *AddUser200ResponseAllOfUser) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### SetWindowsUsernameNil

`func (o *AddUser200ResponseAllOfUser) SetWindowsUsernameNil(b bool)`

 SetWindowsUsernameNil sets the value for WindowsUsername to be an explicit nil

### UnsetWindowsUsername
`func (o *AddUser200ResponseAllOfUser) UnsetWindowsUsername()`

UnsetWindowsUsername ensures that no value is present for WindowsUsername, not even an explicit nil
### GetWindowsPassword

`func (o *AddUser200ResponseAllOfUser) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *AddUser200ResponseAllOfUser) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *AddUser200ResponseAllOfUser) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *AddUser200ResponseAllOfUser) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### SetWindowsPasswordNil

`func (o *AddUser200ResponseAllOfUser) SetWindowsPasswordNil(b bool)`

 SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil

### UnsetWindowsPassword
`func (o *AddUser200ResponseAllOfUser) UnsetWindowsPassword()`

UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
### GetDefaultPersona

`func (o *AddUser200ResponseAllOfUser) GetDefaultPersona() AddUser200ResponseAllOfUserDefaultPersona`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *AddUser200ResponseAllOfUser) GetDefaultPersonaOk() (*AddUser200ResponseAllOfUserDefaultPersona, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *AddUser200ResponseAllOfUser) SetDefaultPersona(v AddUser200ResponseAllOfUserDefaultPersona)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *AddUser200ResponseAllOfUser) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddUser200ResponseAllOfUser) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddUser200ResponseAllOfUser) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddUser200ResponseAllOfUser) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddUser200ResponseAllOfUser) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddUser200ResponseAllOfUser) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddUser200ResponseAllOfUser) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddUser200ResponseAllOfUser) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddUser200ResponseAllOfUser) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccess

`func (o *AddUser200ResponseAllOfUser) GetAccess() AddUser200ResponseAllOfUserAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AddUser200ResponseAllOfUser) GetAccessOk() (*AddUser200ResponseAllOfUserAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AddUser200ResponseAllOfUser) SetAccess(v AddUser200ResponseAllOfUserAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AddUser200ResponseAllOfUser) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



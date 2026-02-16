# UpdateUser200ResponseAllOfUser

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
**Roles** | Pointer to [**[]UpdateUser200ResponseAllOfUserRolesInner**](UpdateUser200ResponseAllOfUserRolesInner.md) |  | [optional] 
**Account** | Pointer to [**UpdateUser200ResponseAllOfUserAccount**](UpdateUser200ResponseAllOfUserAccount.md) |  | [optional] 
**LinuxUsername** | Pointer to **NullableString** |  | [optional] 
**LinuxPassword** | Pointer to **NullableString** |  | [optional] 
**LinuxKeyPairId** | Pointer to **NullableInt64** |  | [optional] 
**WindowsUsername** | Pointer to **NullableString** |  | [optional] 
**WindowsPassword** | Pointer to **NullableString** |  | [optional] 
**DefaultPersona** | Pointer to [**UpdateUser200ResponseAllOfUserDefaultPersona**](UpdateUser200ResponseAllOfUserDefaultPersona.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Access** | Pointer to [**UpdateUser200ResponseAllOfUserAccess**](UpdateUser200ResponseAllOfUserAccess.md) |  | [optional] 

## Methods

### NewUpdateUser200ResponseAllOfUser

`func NewUpdateUser200ResponseAllOfUser() *UpdateUser200ResponseAllOfUser`

NewUpdateUser200ResponseAllOfUser instantiates a new UpdateUser200ResponseAllOfUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUser200ResponseAllOfUserWithDefaults

`func NewUpdateUser200ResponseAllOfUserWithDefaults() *UpdateUser200ResponseAllOfUser`

NewUpdateUser200ResponseAllOfUserWithDefaults instantiates a new UpdateUser200ResponseAllOfUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateUser200ResponseAllOfUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateUser200ResponseAllOfUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateUser200ResponseAllOfUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateUser200ResponseAllOfUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateUser200ResponseAllOfUser) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateUser200ResponseAllOfUser) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateUser200ResponseAllOfUser) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateUser200ResponseAllOfUser) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateUser200ResponseAllOfUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateUser200ResponseAllOfUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateUser200ResponseAllOfUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateUser200ResponseAllOfUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateUser200ResponseAllOfUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateUser200ResponseAllOfUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateUser200ResponseAllOfUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateUser200ResponseAllOfUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateUser200ResponseAllOfUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUser200ResponseAllOfUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUser200ResponseAllOfUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUser200ResponseAllOfUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *UpdateUser200ResponseAllOfUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateUser200ResponseAllOfUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateUser200ResponseAllOfUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateUser200ResponseAllOfUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UpdateUser200ResponseAllOfUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateUser200ResponseAllOfUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateUser200ResponseAllOfUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateUser200ResponseAllOfUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateUser200ResponseAllOfUser) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateUser200ResponseAllOfUser) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateUser200ResponseAllOfUser) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateUser200ResponseAllOfUser) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetReceiveNotifications

`func (o *UpdateUser200ResponseAllOfUser) GetReceiveNotifications() bool`

GetReceiveNotifications returns the ReceiveNotifications field if non-nil, zero value otherwise.

### GetReceiveNotificationsOk

`func (o *UpdateUser200ResponseAllOfUser) GetReceiveNotificationsOk() (*bool, bool)`

GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveNotifications

`func (o *UpdateUser200ResponseAllOfUser) SetReceiveNotifications(v bool)`

SetReceiveNotifications sets ReceiveNotifications field to given value.

### HasReceiveNotifications

`func (o *UpdateUser200ResponseAllOfUser) HasReceiveNotifications() bool`

HasReceiveNotifications returns a boolean if a field has been set.

### GetIsUsing2FA

`func (o *UpdateUser200ResponseAllOfUser) GetIsUsing2FA() bool`

GetIsUsing2FA returns the IsUsing2FA field if non-nil, zero value otherwise.

### GetIsUsing2FAOk

`func (o *UpdateUser200ResponseAllOfUser) GetIsUsing2FAOk() (*bool, bool)`

GetIsUsing2FAOk returns a tuple with the IsUsing2FA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsing2FA

`func (o *UpdateUser200ResponseAllOfUser) SetIsUsing2FA(v bool)`

SetIsUsing2FA sets IsUsing2FA field to given value.

### HasIsUsing2FA

`func (o *UpdateUser200ResponseAllOfUser) HasIsUsing2FA() bool`

HasIsUsing2FA returns a boolean if a field has been set.

### GetAccountExpired

`func (o *UpdateUser200ResponseAllOfUser) GetAccountExpired() bool`

GetAccountExpired returns the AccountExpired field if non-nil, zero value otherwise.

### GetAccountExpiredOk

`func (o *UpdateUser200ResponseAllOfUser) GetAccountExpiredOk() (*bool, bool)`

GetAccountExpiredOk returns a tuple with the AccountExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpired

`func (o *UpdateUser200ResponseAllOfUser) SetAccountExpired(v bool)`

SetAccountExpired sets AccountExpired field to given value.

### HasAccountExpired

`func (o *UpdateUser200ResponseAllOfUser) HasAccountExpired() bool`

HasAccountExpired returns a boolean if a field has been set.

### GetAccountLocked

`func (o *UpdateUser200ResponseAllOfUser) GetAccountLocked() bool`

GetAccountLocked returns the AccountLocked field if non-nil, zero value otherwise.

### GetAccountLockedOk

`func (o *UpdateUser200ResponseAllOfUser) GetAccountLockedOk() (*bool, bool)`

GetAccountLockedOk returns a tuple with the AccountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLocked

`func (o *UpdateUser200ResponseAllOfUser) SetAccountLocked(v bool)`

SetAccountLocked sets AccountLocked field to given value.

### HasAccountLocked

`func (o *UpdateUser200ResponseAllOfUser) HasAccountLocked() bool`

HasAccountLocked returns a boolean if a field has been set.

### GetPasswordExpired

`func (o *UpdateUser200ResponseAllOfUser) GetPasswordExpired() bool`

GetPasswordExpired returns the PasswordExpired field if non-nil, zero value otherwise.

### GetPasswordExpiredOk

`func (o *UpdateUser200ResponseAllOfUser) GetPasswordExpiredOk() (*bool, bool)`

GetPasswordExpiredOk returns a tuple with the PasswordExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpired

`func (o *UpdateUser200ResponseAllOfUser) SetPasswordExpired(v bool)`

SetPasswordExpired sets PasswordExpired field to given value.

### HasPasswordExpired

`func (o *UpdateUser200ResponseAllOfUser) HasPasswordExpired() bool`

HasPasswordExpired returns a boolean if a field has been set.

### GetLoginCount

`func (o *UpdateUser200ResponseAllOfUser) GetLoginCount() int64`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *UpdateUser200ResponseAllOfUser) GetLoginCountOk() (*int64, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *UpdateUser200ResponseAllOfUser) SetLoginCount(v int64)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *UpdateUser200ResponseAllOfUser) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetLoginAttempts

`func (o *UpdateUser200ResponseAllOfUser) GetLoginAttempts() int64`

GetLoginAttempts returns the LoginAttempts field if non-nil, zero value otherwise.

### GetLoginAttemptsOk

`func (o *UpdateUser200ResponseAllOfUser) GetLoginAttemptsOk() (*int64, bool)`

GetLoginAttemptsOk returns a tuple with the LoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttempts

`func (o *UpdateUser200ResponseAllOfUser) SetLoginAttempts(v int64)`

SetLoginAttempts sets LoginAttempts field to given value.

### HasLoginAttempts

`func (o *UpdateUser200ResponseAllOfUser) HasLoginAttempts() bool`

HasLoginAttempts returns a boolean if a field has been set.

### GetLastLoginDate

`func (o *UpdateUser200ResponseAllOfUser) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *UpdateUser200ResponseAllOfUser) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *UpdateUser200ResponseAllOfUser) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *UpdateUser200ResponseAllOfUser) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateUser200ResponseAllOfUser) GetRoles() []UpdateUser200ResponseAllOfUserRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateUser200ResponseAllOfUser) GetRolesOk() (*[]UpdateUser200ResponseAllOfUserRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateUser200ResponseAllOfUser) SetRoles(v []UpdateUser200ResponseAllOfUserRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateUser200ResponseAllOfUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateUser200ResponseAllOfUser) GetAccount() UpdateUser200ResponseAllOfUserAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateUser200ResponseAllOfUser) GetAccountOk() (*UpdateUser200ResponseAllOfUserAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateUser200ResponseAllOfUser) SetAccount(v UpdateUser200ResponseAllOfUserAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateUser200ResponseAllOfUser) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *UpdateUser200ResponseAllOfUser) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *UpdateUser200ResponseAllOfUser) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *UpdateUser200ResponseAllOfUser) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *UpdateUser200ResponseAllOfUser) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### SetLinuxUsernameNil

`func (o *UpdateUser200ResponseAllOfUser) SetLinuxUsernameNil(b bool)`

 SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil

### UnsetLinuxUsername
`func (o *UpdateUser200ResponseAllOfUser) UnsetLinuxUsername()`

UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil
### GetLinuxPassword

`func (o *UpdateUser200ResponseAllOfUser) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *UpdateUser200ResponseAllOfUser) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *UpdateUser200ResponseAllOfUser) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *UpdateUser200ResponseAllOfUser) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### SetLinuxPasswordNil

`func (o *UpdateUser200ResponseAllOfUser) SetLinuxPasswordNil(b bool)`

 SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil

### UnsetLinuxPassword
`func (o *UpdateUser200ResponseAllOfUser) UnsetLinuxPassword()`

UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
### GetLinuxKeyPairId

`func (o *UpdateUser200ResponseAllOfUser) GetLinuxKeyPairId() int64`

GetLinuxKeyPairId returns the LinuxKeyPairId field if non-nil, zero value otherwise.

### GetLinuxKeyPairIdOk

`func (o *UpdateUser200ResponseAllOfUser) GetLinuxKeyPairIdOk() (*int64, bool)`

GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPairId

`func (o *UpdateUser200ResponseAllOfUser) SetLinuxKeyPairId(v int64)`

SetLinuxKeyPairId sets LinuxKeyPairId field to given value.

### HasLinuxKeyPairId

`func (o *UpdateUser200ResponseAllOfUser) HasLinuxKeyPairId() bool`

HasLinuxKeyPairId returns a boolean if a field has been set.

### SetLinuxKeyPairIdNil

`func (o *UpdateUser200ResponseAllOfUser) SetLinuxKeyPairIdNil(b bool)`

 SetLinuxKeyPairIdNil sets the value for LinuxKeyPairId to be an explicit nil

### UnsetLinuxKeyPairId
`func (o *UpdateUser200ResponseAllOfUser) UnsetLinuxKeyPairId()`

UnsetLinuxKeyPairId ensures that no value is present for LinuxKeyPairId, not even an explicit nil
### GetWindowsUsername

`func (o *UpdateUser200ResponseAllOfUser) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *UpdateUser200ResponseAllOfUser) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *UpdateUser200ResponseAllOfUser) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *UpdateUser200ResponseAllOfUser) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### SetWindowsUsernameNil

`func (o *UpdateUser200ResponseAllOfUser) SetWindowsUsernameNil(b bool)`

 SetWindowsUsernameNil sets the value for WindowsUsername to be an explicit nil

### UnsetWindowsUsername
`func (o *UpdateUser200ResponseAllOfUser) UnsetWindowsUsername()`

UnsetWindowsUsername ensures that no value is present for WindowsUsername, not even an explicit nil
### GetWindowsPassword

`func (o *UpdateUser200ResponseAllOfUser) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *UpdateUser200ResponseAllOfUser) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *UpdateUser200ResponseAllOfUser) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *UpdateUser200ResponseAllOfUser) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### SetWindowsPasswordNil

`func (o *UpdateUser200ResponseAllOfUser) SetWindowsPasswordNil(b bool)`

 SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil

### UnsetWindowsPassword
`func (o *UpdateUser200ResponseAllOfUser) UnsetWindowsPassword()`

UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
### GetDefaultPersona

`func (o *UpdateUser200ResponseAllOfUser) GetDefaultPersona() UpdateUser200ResponseAllOfUserDefaultPersona`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *UpdateUser200ResponseAllOfUser) GetDefaultPersonaOk() (*UpdateUser200ResponseAllOfUserDefaultPersona, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *UpdateUser200ResponseAllOfUser) SetDefaultPersona(v UpdateUser200ResponseAllOfUserDefaultPersona)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *UpdateUser200ResponseAllOfUser) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateUser200ResponseAllOfUser) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateUser200ResponseAllOfUser) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateUser200ResponseAllOfUser) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateUser200ResponseAllOfUser) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateUser200ResponseAllOfUser) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateUser200ResponseAllOfUser) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateUser200ResponseAllOfUser) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateUser200ResponseAllOfUser) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccess

`func (o *UpdateUser200ResponseAllOfUser) GetAccess() UpdateUser200ResponseAllOfUserAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateUser200ResponseAllOfUser) GetAccessOk() (*UpdateUser200ResponseAllOfUserAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateUser200ResponseAllOfUser) SetAccess(v UpdateUser200ResponseAllOfUserAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *UpdateUser200ResponseAllOfUser) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



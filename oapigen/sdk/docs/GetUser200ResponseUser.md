# GetUser200ResponseUser

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
**Roles** | Pointer to [**[]GetUser200ResponseUserRolesInner**](GetUser200ResponseUserRolesInner.md) |  | [optional] 
**Account** | Pointer to [**GetUser200ResponseUserAccount**](GetUser200ResponseUserAccount.md) |  | [optional] 
**LinuxUsername** | Pointer to **NullableString** |  | [optional] 
**LinuxPassword** | Pointer to **NullableString** |  | [optional] 
**LinuxKeyPairId** | Pointer to **NullableInt64** |  | [optional] 
**WindowsUsername** | Pointer to **NullableString** |  | [optional] 
**WindowsPassword** | Pointer to **NullableString** |  | [optional] 
**DefaultPersona** | Pointer to [**GetUser200ResponseUserDefaultPersona**](GetUser200ResponseUserDefaultPersona.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Access** | Pointer to [**GetUser200ResponseUserAccess**](GetUser200ResponseUserAccess.md) |  | [optional] 

## Methods

### NewGetUser200ResponseUser

`func NewGetUser200ResponseUser() *GetUser200ResponseUser`

NewGetUser200ResponseUser instantiates a new GetUser200ResponseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUser200ResponseUserWithDefaults

`func NewGetUser200ResponseUserWithDefaults() *GetUser200ResponseUser`

NewGetUser200ResponseUserWithDefaults instantiates a new GetUser200ResponseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetUser200ResponseUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUser200ResponseUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUser200ResponseUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetUser200ResponseUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetUser200ResponseUser) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetUser200ResponseUser) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetUser200ResponseUser) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetUser200ResponseUser) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUsername

`func (o *GetUser200ResponseUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetUser200ResponseUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetUser200ResponseUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetUser200ResponseUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetUser200ResponseUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetUser200ResponseUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetUser200ResponseUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetUser200ResponseUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *GetUser200ResponseUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetUser200ResponseUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetUser200ResponseUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetUser200ResponseUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *GetUser200ResponseUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *GetUser200ResponseUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *GetUser200ResponseUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *GetUser200ResponseUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *GetUser200ResponseUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GetUser200ResponseUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GetUser200ResponseUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *GetUser200ResponseUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEnabled

`func (o *GetUser200ResponseUser) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetUser200ResponseUser) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetUser200ResponseUser) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetUser200ResponseUser) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetReceiveNotifications

`func (o *GetUser200ResponseUser) GetReceiveNotifications() bool`

GetReceiveNotifications returns the ReceiveNotifications field if non-nil, zero value otherwise.

### GetReceiveNotificationsOk

`func (o *GetUser200ResponseUser) GetReceiveNotificationsOk() (*bool, bool)`

GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveNotifications

`func (o *GetUser200ResponseUser) SetReceiveNotifications(v bool)`

SetReceiveNotifications sets ReceiveNotifications field to given value.

### HasReceiveNotifications

`func (o *GetUser200ResponseUser) HasReceiveNotifications() bool`

HasReceiveNotifications returns a boolean if a field has been set.

### GetIsUsing2FA

`func (o *GetUser200ResponseUser) GetIsUsing2FA() bool`

GetIsUsing2FA returns the IsUsing2FA field if non-nil, zero value otherwise.

### GetIsUsing2FAOk

`func (o *GetUser200ResponseUser) GetIsUsing2FAOk() (*bool, bool)`

GetIsUsing2FAOk returns a tuple with the IsUsing2FA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsing2FA

`func (o *GetUser200ResponseUser) SetIsUsing2FA(v bool)`

SetIsUsing2FA sets IsUsing2FA field to given value.

### HasIsUsing2FA

`func (o *GetUser200ResponseUser) HasIsUsing2FA() bool`

HasIsUsing2FA returns a boolean if a field has been set.

### GetAccountExpired

`func (o *GetUser200ResponseUser) GetAccountExpired() bool`

GetAccountExpired returns the AccountExpired field if non-nil, zero value otherwise.

### GetAccountExpiredOk

`func (o *GetUser200ResponseUser) GetAccountExpiredOk() (*bool, bool)`

GetAccountExpiredOk returns a tuple with the AccountExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpired

`func (o *GetUser200ResponseUser) SetAccountExpired(v bool)`

SetAccountExpired sets AccountExpired field to given value.

### HasAccountExpired

`func (o *GetUser200ResponseUser) HasAccountExpired() bool`

HasAccountExpired returns a boolean if a field has been set.

### GetAccountLocked

`func (o *GetUser200ResponseUser) GetAccountLocked() bool`

GetAccountLocked returns the AccountLocked field if non-nil, zero value otherwise.

### GetAccountLockedOk

`func (o *GetUser200ResponseUser) GetAccountLockedOk() (*bool, bool)`

GetAccountLockedOk returns a tuple with the AccountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLocked

`func (o *GetUser200ResponseUser) SetAccountLocked(v bool)`

SetAccountLocked sets AccountLocked field to given value.

### HasAccountLocked

`func (o *GetUser200ResponseUser) HasAccountLocked() bool`

HasAccountLocked returns a boolean if a field has been set.

### GetPasswordExpired

`func (o *GetUser200ResponseUser) GetPasswordExpired() bool`

GetPasswordExpired returns the PasswordExpired field if non-nil, zero value otherwise.

### GetPasswordExpiredOk

`func (o *GetUser200ResponseUser) GetPasswordExpiredOk() (*bool, bool)`

GetPasswordExpiredOk returns a tuple with the PasswordExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpired

`func (o *GetUser200ResponseUser) SetPasswordExpired(v bool)`

SetPasswordExpired sets PasswordExpired field to given value.

### HasPasswordExpired

`func (o *GetUser200ResponseUser) HasPasswordExpired() bool`

HasPasswordExpired returns a boolean if a field has been set.

### GetLoginCount

`func (o *GetUser200ResponseUser) GetLoginCount() int64`

GetLoginCount returns the LoginCount field if non-nil, zero value otherwise.

### GetLoginCountOk

`func (o *GetUser200ResponseUser) GetLoginCountOk() (*int64, bool)`

GetLoginCountOk returns a tuple with the LoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCount

`func (o *GetUser200ResponseUser) SetLoginCount(v int64)`

SetLoginCount sets LoginCount field to given value.

### HasLoginCount

`func (o *GetUser200ResponseUser) HasLoginCount() bool`

HasLoginCount returns a boolean if a field has been set.

### GetLoginAttempts

`func (o *GetUser200ResponseUser) GetLoginAttempts() int64`

GetLoginAttempts returns the LoginAttempts field if non-nil, zero value otherwise.

### GetLoginAttemptsOk

`func (o *GetUser200ResponseUser) GetLoginAttemptsOk() (*int64, bool)`

GetLoginAttemptsOk returns a tuple with the LoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttempts

`func (o *GetUser200ResponseUser) SetLoginAttempts(v int64)`

SetLoginAttempts sets LoginAttempts field to given value.

### HasLoginAttempts

`func (o *GetUser200ResponseUser) HasLoginAttempts() bool`

HasLoginAttempts returns a boolean if a field has been set.

### GetLastLoginDate

`func (o *GetUser200ResponseUser) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *GetUser200ResponseUser) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *GetUser200ResponseUser) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *GetUser200ResponseUser) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### GetRoles

`func (o *GetUser200ResponseUser) GetRoles() []GetUser200ResponseUserRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GetUser200ResponseUser) GetRolesOk() (*[]GetUser200ResponseUserRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GetUser200ResponseUser) SetRoles(v []GetUser200ResponseUserRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GetUser200ResponseUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAccount

`func (o *GetUser200ResponseUser) GetAccount() GetUser200ResponseUserAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetUser200ResponseUser) GetAccountOk() (*GetUser200ResponseUserAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetUser200ResponseUser) SetAccount(v GetUser200ResponseUserAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetUser200ResponseUser) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *GetUser200ResponseUser) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *GetUser200ResponseUser) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *GetUser200ResponseUser) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *GetUser200ResponseUser) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### SetLinuxUsernameNil

`func (o *GetUser200ResponseUser) SetLinuxUsernameNil(b bool)`

 SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil

### UnsetLinuxUsername
`func (o *GetUser200ResponseUser) UnsetLinuxUsername()`

UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil
### GetLinuxPassword

`func (o *GetUser200ResponseUser) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *GetUser200ResponseUser) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *GetUser200ResponseUser) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *GetUser200ResponseUser) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### SetLinuxPasswordNil

`func (o *GetUser200ResponseUser) SetLinuxPasswordNil(b bool)`

 SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil

### UnsetLinuxPassword
`func (o *GetUser200ResponseUser) UnsetLinuxPassword()`

UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
### GetLinuxKeyPairId

`func (o *GetUser200ResponseUser) GetLinuxKeyPairId() int64`

GetLinuxKeyPairId returns the LinuxKeyPairId field if non-nil, zero value otherwise.

### GetLinuxKeyPairIdOk

`func (o *GetUser200ResponseUser) GetLinuxKeyPairIdOk() (*int64, bool)`

GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPairId

`func (o *GetUser200ResponseUser) SetLinuxKeyPairId(v int64)`

SetLinuxKeyPairId sets LinuxKeyPairId field to given value.

### HasLinuxKeyPairId

`func (o *GetUser200ResponseUser) HasLinuxKeyPairId() bool`

HasLinuxKeyPairId returns a boolean if a field has been set.

### SetLinuxKeyPairIdNil

`func (o *GetUser200ResponseUser) SetLinuxKeyPairIdNil(b bool)`

 SetLinuxKeyPairIdNil sets the value for LinuxKeyPairId to be an explicit nil

### UnsetLinuxKeyPairId
`func (o *GetUser200ResponseUser) UnsetLinuxKeyPairId()`

UnsetLinuxKeyPairId ensures that no value is present for LinuxKeyPairId, not even an explicit nil
### GetWindowsUsername

`func (o *GetUser200ResponseUser) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *GetUser200ResponseUser) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *GetUser200ResponseUser) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *GetUser200ResponseUser) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### SetWindowsUsernameNil

`func (o *GetUser200ResponseUser) SetWindowsUsernameNil(b bool)`

 SetWindowsUsernameNil sets the value for WindowsUsername to be an explicit nil

### UnsetWindowsUsername
`func (o *GetUser200ResponseUser) UnsetWindowsUsername()`

UnsetWindowsUsername ensures that no value is present for WindowsUsername, not even an explicit nil
### GetWindowsPassword

`func (o *GetUser200ResponseUser) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *GetUser200ResponseUser) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *GetUser200ResponseUser) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *GetUser200ResponseUser) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### SetWindowsPasswordNil

`func (o *GetUser200ResponseUser) SetWindowsPasswordNil(b bool)`

 SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil

### UnsetWindowsPassword
`func (o *GetUser200ResponseUser) UnsetWindowsPassword()`

UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
### GetDefaultPersona

`func (o *GetUser200ResponseUser) GetDefaultPersona() GetUser200ResponseUserDefaultPersona`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *GetUser200ResponseUser) GetDefaultPersonaOk() (*GetUser200ResponseUserDefaultPersona, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *GetUser200ResponseUser) SetDefaultPersona(v GetUser200ResponseUserDefaultPersona)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *GetUser200ResponseUser) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetUser200ResponseUser) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetUser200ResponseUser) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetUser200ResponseUser) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetUser200ResponseUser) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetUser200ResponseUser) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetUser200ResponseUser) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetUser200ResponseUser) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetUser200ResponseUser) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccess

`func (o *GetUser200ResponseUser) GetAccess() GetUser200ResponseUserAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *GetUser200ResponseUser) GetAccessOk() (*GetUser200ResponseUserAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *GetUser200ResponseUser) SetAccess(v GetUser200ResponseUserAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *GetUser200ResponseUser) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


